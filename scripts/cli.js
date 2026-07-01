#!/usr/bin/env node

const { Command } = require("commander");
const fs = require("node:fs");
const path = require("node:path");
const crypto = require("node:crypto");
const { execSync } = require("node:child_process");

const SUPPORTED_AGENTS = ["copilot", "pi", "claude", "opencode"];
const STATE_FILE = ".cli-skill-install-state.json";
const HELP_APPENDIX = `

Auto-detection:
  When --agents is omitted, installer detects agents from repository markers.
  If nothing is detected, only cli-skill core is installed.

Managed updates:
  Existing unmanaged files are never overwritten unless --force is set.
  Existing managed files are auto-updated only when unchanged since last install.
`;

function writeOut(str) {
  process.stdout.write(str);
}

function writeErr(str) {
  process.stderr.write(str);
}

const AGENT_MARKERS = {
  copilot: {
    paths: [
      ".github/copilot-instructions.md",
      ".github/prompts",
    ],
    contentHints: [/copilot/i, /github\s+copilot/i],
  },
  pi: {
    paths: [
      ".pi",
      ".pi/skills",
    ],
    contentHints: [/pi\s+coding\s+agent/i],
  },
  claude: {
    paths: [
      "CLAUDE.md",
      ".claude",
      ".claude/commands",
    ],
    contentHints: [/claude\s+code/i, /anthropic\s+claude/i],
  },
  opencode: {
    paths: [
      "opencode.json",
      ".opencode",
      ".opencode/commands",
    ],
    contentHints: [/\bopencode\b/i],
  },
};

function normalizeBool(value) {
  if (!value) return false;
  const lowered = String(value).toLowerCase();
  return lowered === "1" || lowered === "true" || lowered === "yes";
}

function parseAgents(raw) {
  if (!raw) return [];
  return String(raw)
    .split(",")
    .map((part) => part.trim().toLowerCase())
    .filter(Boolean)
    .filter((agent) => SUPPORTED_AGENTS.includes(agent));
}

function fileContainsAnyPattern(filePath, patterns) {
  if (!fs.existsSync(filePath)) return false;
  let text = "";
  try {
    text = fs.readFileSync(filePath, "utf-8");
  } catch {
    return false;
  }

  return patterns.some((pattern) => pattern.test(text));
}

function detectAgents(targetDir) {
  const detected = [];
  const reasons = {};
  const hintFiles = ["AGENTS.md"];

  for (const agent of SUPPORTED_AGENTS) {
    const marker = AGENT_MARKERS[agent];
    reasons[agent] = [];

    for (const relPath of marker.paths) {
      if (fs.existsSync(path.join(targetDir, relPath))) {
        reasons[agent].push(`marker:${relPath}`);
      }
    }

    for (const relPath of hintFiles) {
      const absPath = path.join(targetDir, relPath);
      if (fileContainsAnyPattern(absPath, marker.contentHints)) {
        reasons[agent].push(`hint:${relPath}`);
      }
    }

    if (reasons[agent].length > 0) {
      detected.push(agent);
    }
  }

  return {
    agents: detected,
    reasons,
  };
}

function hashBuffer(buffer) {
  return crypto.createHash("sha256").update(buffer).digest("hex");
}

function hashFile(filePath) {
  return hashBuffer(fs.readFileSync(filePath));
}

function readState(targetDir) {
  const statePath = path.join(targetDir, STATE_FILE);

  if (!fs.existsSync(statePath)) {
    return { path: statePath, data: { version: 1, files: {} } };
  }

  try {
    const parsed = JSON.parse(fs.readFileSync(statePath, "utf-8"));
    if (!parsed || typeof parsed !== "object" || typeof parsed.files !== "object") {
      return { path: statePath, data: { version: 1, files: {} } };
    }
    return { path: statePath, data: parsed };
  } catch {
    return { path: statePath, data: { version: 1, files: {} } };
  }
}

function writeState(statePath, stateData, options, logs) {
  if (options.dryRun) {
    logs.push(`skip state write (dry-run): ${path.relative(options.targetDir, statePath)}`);
    return;
  }

  fs.mkdirSync(path.dirname(statePath), { recursive: true });
  fs.writeFileSync(statePath, `${JSON.stringify(stateData, null, 2)}\n`, "utf-8");
  logs.push(`wrote state: ${path.relative(options.targetDir, statePath)}`);
}

function listFilesRecursive(rootDir) {
  const out = [];

  function walk(currentDir) {
    const entries = fs.readdirSync(currentDir, { withFileTypes: true });
    for (const entry of entries) {
      const absPath = path.join(currentDir, entry.name);
      if (entry.isDirectory()) {
        walk(absPath);
      } else if (entry.isFile()) {
        out.push(absPath);
      }
    }
  }

  walk(rootDir);
  return out;
}

function shouldOverwriteManaged(dstFile, srcHash, options, stateData) {
  if (!fs.existsSync(dstFile)) {
    return { overwrite: true, reason: "create" };
  }

  if (options.force) {
    return { overwrite: true, reason: "force" };
  }

  const rel = path.relative(options.targetDir, dstFile);
  const currentHash = hashFile(dstFile);
  const tracked = stateData.files[rel];

  if (currentHash === srcHash) {
    return { overwrite: false, reason: "already-current" };
  }

  if (!tracked) {
    return { overwrite: false, reason: "unmanaged-existing" };
  }

  if (tracked.installedHash === currentHash) {
    return { overwrite: true, reason: "managed-update" };
  }

  return { overwrite: false, reason: "managed-modified" };
}

function copyManagedFile(srcFile, dstFile, options, logs, stateData) {
  const srcHash = hashFile(srcFile);
  const decision = shouldOverwriteManaged(dstFile, srcHash, options, stateData);

  if (!decision.overwrite) {
    const rel = path.relative(options.targetDir, dstFile);
    if (decision.reason === "already-current") {
      logs.push(`skip file (up-to-date): ${rel}`);
      stateData.files[rel] = {
        sourceHash: srcHash,
        installedHash: srcHash,
      };
      return;
    }

    if (decision.reason === "managed-modified") {
      logs.push(`skip file (managed but modified): ${rel}`);
      return;
    }

    logs.push(`skip file (exists unmanaged): ${rel}`);
    return;
  }

  if (options.dryRun) {
    logs.push(`${decision.reason}: ${path.relative(options.targetDir, srcFile)} -> ${path.relative(options.targetDir, dstFile)}`);
    const rel = path.relative(options.targetDir, dstFile);
    stateData.files[rel] = {
      sourceHash: srcHash,
      installedHash: srcHash,
    };
    return;
  }

  fs.mkdirSync(path.dirname(dstFile), { recursive: true });
  fs.copyFileSync(srcFile, dstFile);
  logs.push(`${decision.reason}: ${path.relative(options.targetDir, dstFile)}`);
  const rel = path.relative(options.targetDir, dstFile);
  stateData.files[rel] = {
    sourceHash: srcHash,
    installedHash: srcHash,
  };
}

function copyManagedTree(srcRootDir, dstRootDir, options, logs, stateData) {
  const files = listFilesRecursive(srcRootDir);
  for (const srcFile of files) {
    const relFromSrc = path.relative(srcRootDir, srcFile);
    const dstFile = path.join(dstRootDir, relFromSrc);
    copyManagedFile(srcFile, dstFile, options, logs, stateData);
  }
}

function findGitTopLevel(startDir) {
  try {
    const result = execSync("git rev-parse --show-toplevel", {
      cwd: startDir,
      encoding: "utf-8",
      stdio: ["ignore", "pipe", "ignore"],
    }).trim();

    if (!result) return null;
    return path.resolve(result);
  } catch {
    return null;
  }
}

function resolveTargetDir(targetDirArg) {
  if (targetDirArg) {
    return {
      targetDir: path.resolve(targetDirArg),
      targetSource: "--target",
    };
  }

  const startDir = path.resolve(process.env.INIT_CWD || process.cwd());
  const gitTopLevel = findGitTopLevel(startDir);
  if (gitTopLevel) {
    return {
      targetDir: gitTopLevel,
      targetSource: "git-root",
    };
  }

  return {
    targetDir: startDir,
    targetSource: "cwd",
  };
}

function installSkill({ agentsArg, dryRun = false, force = false, postinstall = false, targetDirArg = null }) {
  const packageRoot = path.resolve(__dirname, "..");
  const { targetDir, targetSource } = resolveTargetDir(targetDirArg);
  const logs = [];
  const warnings = [];
  const state = readState(targetDir);

  if (!fs.existsSync(path.join(packageRoot, "cli-skill"))) {
    throw new Error("Package payload missing cli-skill directory.");
  }

  const envAgents = parseAgents(process.env.CLI_SKILL_AGENTS);
  const argAgents = parseAgents(agentsArg);
  const detection = detectAgents(targetDir);
  const detected = detection.agents;

  let selectionSource = "auto-detect";
  let selectedAgents = detected;

  if (argAgents.length > 0) {
    selectionSource = "--agents";
    selectedAgents = argAgents;
  } else if (envAgents.length > 0) {
    selectionSource = "CLI_SKILL_AGENTS";
    selectedAgents = envAgents;
  }

  selectedAgents = [...new Set(selectedAgents)];

  const resolvedForce = force || normalizeBool(process.env.CLI_SKILL_FORCE);
  const options = { dryRun, force: resolvedForce, targetDir };

  logs.push(`target: ${targetDir}`);
  logs.push(`target source: ${targetSource}`);
  logs.push(`agents: ${selectedAgents.join(",")}`);
  logs.push(`agents source: ${selectionSource}`);
  logs.push(`mode: ${options.dryRun ? "dry-run" : "write"}${options.force ? ", force" : ""}`);
  if (selectionSource === "auto-detect" && selectedAgents.length > 0) {
    for (const agent of selectedAgents) {
      logs.push(`detected ${agent}: ${detection.reasons[agent].join(", ")}`);
    }
  }

  const srcCliSkill = path.join(packageRoot, "cli-skill");
  const dstCliSkill = path.join(targetDir, "cli-skill");
  copyManagedTree(srcCliSkill, dstCliSkill, options, logs, state.data);

  const installs = {
    copilot: [
      {
        src: path.join(packageRoot, "cli-skill", "stubs", "entrypoints", "github-skill", "SKILL.md"),
        dst: path.join(targetDir, ".github", "skills", "cli-skill", "SKILL.md"),
      },
    ],
    pi: [
      {
        src: path.join(packageRoot, "cli-skill", "stubs", "entrypoints", "pi-skill", "SKILL.md"),
        dst: path.join(targetDir, ".pi", "skills", "cli-skill", "SKILL.md"),
      },
    ],
    claude: [
      {
        src: path.join(packageRoot, "cli-skill", "stubs", "agents", "claude-code", "commands.yaml"),
        dst: path.join(targetDir, ".claude", "commands", "cli-skill.yaml"),
      },
    ],
    opencode: [
      {
        src: path.join(packageRoot, "cli-skill", "stubs", "agents", "opencode", "commands.json"),
        dst: path.join(targetDir, ".opencode", "commands", "cli-skill.json"),
      },
    ],
  };

  for (const agent of selectedAgents) {
    const mappings = installs[agent] || [];
    logs.push(`installing adapter: ${agent}`);

    for (const mapping of mappings) {
      if (!fs.existsSync(mapping.src)) {
        warnings.push(`warning: missing source for ${agent}: ${mapping.src}`);
        continue;
      }

      copyManagedFile(mapping.src, mapping.dst, options, logs, state.data);
    }
  }

  writeState(state.path, state.data, options, logs);

  if (postinstall) {
    logs.push("postinstall complete");
    logs.push("tip: run `npx canonical-cli-skill install --dry-run` to preview updates later");
  }

  if (detected.length === 0 && argAgents.length === 0 && envAgents.length === 0) {
    warnings.push("warning: no known agent markers detected, installed cli-skill core only (set --agents to force adapter install)");
  }

  return { logs, warnings };
}

function runInstall(options, command) {
  try {
    const result = installSkill({
      agentsArg: options.agents || null,
      dryRun: Boolean(options.dryRun),
      force: Boolean(options.force),
      postinstall: Boolean(options.postinstall),
      targetDirArg: options.target || null,
    });

    result.logs.forEach((line) => writeOut(`${line}\n`));

    if (result.warnings.length > 0) {
      result.warnings.forEach((line) => writeErr(`${line}\n`));
    }
  } catch (error) {
    command.error(`Install failed: ${error instanceof Error ? error.message : String(error)}`, {
      exitCode: 1,
      code: "cli-skill.install-failed",
    });
  }
}

function main() {
  const program = new Command();

  program
    .name("canonical-cli-skill")
    .description("Install and maintain cli-skill files for supported coding agents.")
    .configureOutput({
      writeOut,
      writeErr,
    })
    .showHelpAfterError()
    .addHelpText("after", HELP_APPENDIX);

  program
    .command("install")
    .description("Install cli-skill core and detected/selected agent adapter files.")
    .option("--dry-run", "Show changes without writing files")
    .option("--force", "Overwrite existing files")
    .option("--postinstall", "Internal flag used by npm postinstall")
    .option("--agents <list>", "Comma-separated agent list (copilot,pi,claude,opencode)")
    .option("--target <path>", "Target directory for installation")
    .action((options, command) => {
      runInstall(options, command);
    });

  program.parse(process.argv);
}

main();
