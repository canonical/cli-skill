const fs = require('node:fs');
const path = require('node:path');
const { execSync } = require('node:child_process');

function getRepoRoot() {
  try {
    return execSync('git rev-parse --show-toplevel', { encoding: 'utf-8' }).trim();
  } catch {
    return path.resolve(__dirname, '..');
  }
}

const ROOT = getRepoRoot();
const MANIFEST_PATH = path.join(ROOT, 'cli-skill', 'schemas', 'commands.manifest.yaml');

function parseManifest(yamlText) {
  const lines = yamlText.split(/\r?\n/);

  const result = {
    version: null,
    skill: null,
    preflightModule: null,
    preflightOutputDir: null,
    commands: [],
    futureCommands: [],
  };

  let section = 'root';
  let current = null;

  for (const rawLine of lines) {
    const line = rawLine.trim();
    if (!line || line.startsWith('#')) continue;

    if (line.startsWith('version:')) {
      result.version = line.split(':').slice(1).join(':').trim();
      continue;
    }
    if (line.startsWith('skill:')) {
      result.skill = line.split(':').slice(1).join(':').trim();
      continue;
    }

    if (line === 'shared:') {
      section = 'shared';
      continue;
    }
    if (line === 'commands:') {
      section = 'commands';
      current = null;
      continue;
    }
    if (line === 'futureCommands:') {
      section = 'futureCommands';
      current = null;
      continue;
    }

    if (section === 'shared') {
      if (line.startsWith('module:')) {
        result.preflightModule = line.split(':').slice(1).join(':').trim();
      }
      if (line.startsWith('outputDir:')) {
        result.preflightOutputDir = line.split(':').slice(1).join(':').trim();
      }
      continue;
    }

    if (section === 'commands' || section === 'futureCommands') {
      if (line.startsWith('- name:')) {
        const name = line.replace('- name:', '').trim();
        current = { name, file: '', description: '' };
        if (section === 'commands') {
          result.commands.push(current);
        } else {
          result.futureCommands.push(current);
        }
        continue;
      }

      if (current && line.startsWith('file:')) {
        current.file = line.split(':').slice(1).join(':').trim();
        continue;
      }

      if (current && line.startsWith('description:')) {
        current.description = line.split(':').slice(1).join(':').trim();
        continue;
      }
    }
  }

  if (!result.skill || !result.preflightModule || result.commands.length === 0) {
    throw new Error('Manifest parse failed: missing required fields (skill/preflight/commands).');
  }

  return result;
}

function ensureDir(dirPath) {
  fs.mkdirSync(dirPath, { recursive: true });
}

function writeFileSafe(filePath, content) {
  ensureDir(path.dirname(filePath));
  fs.writeFileSync(filePath, content, 'utf-8');
}

function renderCommandMap(commands) {
  return commands.map((c) => `- \`${c.name}\` -> \`${c.file}\``).join('\n');
}

function renderFutureMap(commands) {
  if (!commands.length) return '- none';
  return commands.map((c) => `- \`${c.name}\` -> \`${c.file}\``).join('\n');
}

function toClaudeRoutes(commands) {
  return commands.map((c) => `  ${c.name}: ${c.file}`).join('\n');
}

function toOpenCodeRoutes(commands) {
  return commands
    .map((c, i) => `    "${c.name}": "${c.file}"${i < commands.length - 1 ? ',' : ''}`)
    .join('\n');
}

function sync(manifest) {
  const activeRoutes = renderCommandMap(manifest.commands);
  const futureRoutes = renderFutureMap(manifest.futureCommands);

  const copilotAdapter = `---
name: ${manifest.skill}
description: "Copilot adapter for cross-agent CLI skill commands."
---

# Copilot Adapter

This adapter maps slash-style command intents to command files in \`cli-skill/\`.

## Resolve Order

1. Read \`cli-skill/schemas/commands.manifest.yaml\`
2. Resolve command to file path
3. Run \`${manifest.preflightModule}\` before any analysis command
4. Execute requested command workflow

## Command Routing

${activeRoutes}

## Future Command Stubs

${futureRoutes}
`;

  const claudeAdapter = `version: ${manifest.version || 1}
agent: claude-code
skill: ${manifest.skill}

bootstrap:
  manifest: cli-skill/schemas/commands.manifest.yaml
  preflight: ${manifest.preflightModule}

routes:
${toClaudeRoutes(manifest.commands)}

future:
${toClaudeRoutes(manifest.futureCommands)}
`;

  const piAdapter = `---
name: ${manifest.skill}
description: "Pi Coding Agent adapter for command-per-file CLI workflows."
---

# Pi Coding Agent Adapter

This adapter is intentionally thin and delegates all behavior to files under \`cli-skill/\`.

## Startup

- Load \`cli-skill/schemas/commands.manifest.yaml\`
- Route user intent to one command file
- Always run \`${manifest.preflightModule}\` first

## Commands

${activeRoutes}
`;

  const openCodeAdapter = `{
  "version": ${Number(manifest.version || 1)},
  "agent": "opencode",
  "skill": "${manifest.skill}",
  "manifest": "cli-skill/schemas/commands.manifest.yaml",
  "preflight": "${manifest.preflightModule}",
  "commands": {
${toOpenCodeRoutes(manifest.commands)}
  },
  "future": {
${toOpenCodeRoutes(manifest.futureCommands)}
  }
}
`;

  const githubEntrypoint = `---
name: ${manifest.skill}
description: "GitHub Copilot entrypoint for the cross-agent CLI skill."
---

# Entrypoint

This is an adapter entrypoint. The canonical implementation lives in \`cli-skill/\`.

## Load Order

1. Read \`cli-skill/schemas/commands.manifest.yaml\`
2. Resolve command file for the requested command
3. Run \`${manifest.preflightModule}\`
4. Execute command workflow file

## Commands

${manifest.commands.map((c) => `- \`${c.name}\``).join('\n')}
`;

  const piEntrypoint = `---
name: ${manifest.skill}
description: "Pi entrypoint for cross-agent CLI command workflows."
---

# Entrypoint

Canonical files are in \`cli-skill/\`.

## Dispatch

- Manifest: \`cli-skill/schemas/commands.manifest.yaml\`
- Shared preflight: \`${manifest.preflightModule}\`
- Commands: \`cli-skill/commands/*.md\`

## Supported Commands

${manifest.commands.map((c) => `- \`${c.name}\``).join('\n')}
`;

  writeFileSafe(path.join(ROOT, 'cli-skill', 'adapters', 'copilot', 'SKILL.md'), copilotAdapter);
  writeFileSafe(path.join(ROOT, 'cli-skill', 'adapters', 'claude-code', 'commands.yaml'), claudeAdapter);
  writeFileSafe(path.join(ROOT, 'cli-skill', 'adapters', 'pi-coding-agent', 'SKILL.md'), piAdapter);
  writeFileSafe(path.join(ROOT, 'cli-skill', 'adapters', 'opencode', 'commands.json'), openCodeAdapter);
  writeFileSafe(path.join(ROOT, '.github', 'skills', 'cli-skill', 'SKILL.md'), githubEntrypoint);
  writeFileSafe(path.join(ROOT, '.pi', 'skills', 'cli-skill', 'SKILL.md'), piEntrypoint);

  return {
    filesWritten: [
      'cli-skill/adapters/copilot/SKILL.md',
      'cli-skill/adapters/claude-code/commands.yaml',
      'cli-skill/adapters/pi-coding-agent/SKILL.md',
      'cli-skill/adapters/opencode/commands.json',
      '.github/skills/cli-skill/SKILL.md',
      '.pi/skills/cli-skill/SKILL.md',
    ],
  };
}

function main() {
  const manifestText = fs.readFileSync(MANIFEST_PATH, 'utf-8');
  const manifest = parseManifest(manifestText);
  const result = sync(manifest);

  console.log('Synced cli-skill adapters from manifest:');
  for (const file of result.filesWritten) {
    console.log(`- ${file}`);
  }
  console.log(`Preflight output dir (fixed): ${manifest.preflightOutputDir}`);
}

main();
