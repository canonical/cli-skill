const { spawn } = require('node:child_process');
const fs = require('node:fs');
const path = require('node:path');
const os = require('node:os');
const { execSync } = require('node:child_process');

function getRepoRoot() {
  try {
    return execSync('git rev-parse --show-toplevel', { encoding: 'utf-8' }).trim();
  } catch {
    return path.resolve(__dirname, '..');
  }
}

const ROOT = getRepoRoot();

const PROJECTS = [
  { name: 'juju', dir: path.join(ROOT, 'juju') },
  { name: 'qwen36-snap', dir: path.join(ROOT, 'qwen36-snap') },
];

const MODELS = [
  { name: 'kimi-k2.6', modelId: 'moonshotai/kimi-k2.6' },
  { name: 'glm-5', modelId: 'z-ai/glm-5' },
  { name: 'deepseek-v4-pro', modelId: 'deepseek/deepseek-v4-pro' },
];

const AGENTS_DIR = path.join(ROOT, 'agents');

dirname = path.dirname;

// Ensure skill is available in each project
PROJECTS.forEach(p => {
  const skillDir = path.join(p.dir, '.pi', 'skills', 'cli-review');
  const skillSrc = path.join(ROOT, '.github', 'skills', 'cli-review');
  if (!fs.existsSync(skillDir)) {
    fs.mkdirSync(skillDir, { recursive: true });
    fs.copyFileSync(path.join(skillSrc, 'SKILL.md'), path.join(skillDir, 'SKILL.md'));
    fs.mkdirSync(path.join(skillDir, 'standard'), { recursive: true });
    fs.copyFileSync(path.join(skillSrc, 'standard', 'README.md'), path.join(skillDir, 'standard', 'README.md'));
    fs.mkdirSync(path.join(skillDir, 'deprecation'), { recursive: true });
    fs.copyFileSync(path.join(skillSrc, 'deprecation', 'README.md'), path.join(skillDir, 'deprecation', 'README.md'));
  }
});

// Create agent output directories and run
const agents = [];
for (const proj of PROJECTS) {
  for (const mod of MODELS) {
    const agentName = `${mod.name}-${proj.name}`;
    const agentDir = path.join(AGENTS_DIR, agentName);
    fs.mkdirSync(agentDir, { recursive: true });
    fs.mkdirSync(path.join(agentDir, '0-analysis'), { recursive: true });
    fs.mkdirSync(path.join(agentDir, '1-command-design'), { recursive: true });

    const sysPrompt = `You are an autonomous CLI analysis agent named "${agentName}".\n\nYour task:\n1. Run the cli-review skill workflow for the project at ${proj.dir}.\n2. Execute analyze-cli in THREE phases. Complete each phase fully before starting the next:\n   - Phase 1 (Structure Discovery): Write architecture.md, commandset.md, argument-structure.md to 0-analysis/. Verify all 3 exist before continuing.\n   - Phase 2 (Behavioral Analysis): Re-read your Phase 1 files, then write configuration-model.md, output-contracts.md, error-model-and-exit-codes.md, safety-model.md. Verify all 4 exist before continuing.\n   - Phase 3 (Meta-Analysis): Re-read your Phase 1+2 files, then write extensibility-model.md, documentation-quality-gaps.md. Verify all 9 files exist.\n3. Then, execute discuss-commandset: create a 1-command-design/commandset-shape.md with all six sections.\n4. Write ALL outputs to your agent directory: ${agentDir}/\n5. Do NOT reuse or reference any pre-existing analysis files. Do your own independent work.\n6. Be thorough and complete all sections.\n\nThe cli-review skill has been copied into this project at .pi/skills/cli-review/. It will be auto-discovered.\n`;

    const tmpDir = fs.mkdtempSync(path.join(os.tmpdir(), 'pi-agent-'));
    const promptFile = path.join(tmpDir, 'prompt.md');
    fs.writeFileSync(promptFile, sysPrompt, { mode: 0o600 });

    const args = [
      '--provider', 'openrouter',
      '--mode', 'json',
      '-p',
      '--no-session',
      '--model', mod.modelId,
      '--append-system-prompt', promptFile,
      `Run cli-review for ${proj.name}. Create all analysis and commandset files in ${agentDir}/. Do independent fresh work.`,
    ];

    agents.push({
      name: agentName,
      dir: agentDir,
      tmpDir,
      promptFile,
      project: proj,
      model: mod,
      args,
      process: null,
      exitCode: null,
      stdout: '',
      stderr: '',
    });
  }
}

function runAgent(agent) {
  return new Promise((resolve) => {
    console.log(`[START] ${agent.name} -> ${agent.dir}`);
    const sessionPath = path.join(agent.dir, 'session.jsonl');
    const stdoutStream = fs.createWriteStream(sessionPath);
    const proc = spawn('pi', agent.args, {
      cwd: agent.project.dir,
      shell: false,
      stdio: ['ignore', 'pipe', 'pipe'],
    });
    agent.process = proc;
    proc.stdout.pipe(stdoutStream);

    proc.stderr.on('data', (data) => {
      agent.stderr += data.toString();
    });
    proc.on('close', (code) => {
      stdoutStream.end();
      agent.exitCode = code ?? 0;
      fs.writeFileSync(path.join(agent.dir, 'stderr.log'), agent.stderr, 'utf-8');
      // Cleanup temp
      try { fs.unlinkSync(agent.promptFile); } catch {}
      try { fs.rmdirSync(agent.tmpDir); } catch {}
      console.log(`[DONE] ${agent.name} exit=${agent.exitCode}`);
      resolve();
    });
    proc.on('error', (err) => {
      stdoutStream.end();
      agent.exitCode = 1;
      fs.writeFileSync(path.join(agent.dir, 'stderr.log'), agent.stderr + '\n' + err.message, 'utf-8');
      try { fs.unlinkSync(agent.promptFile); } catch {}
      try { fs.rmdirSync(agent.tmpDir); } catch {}
      console.log(`[ERROR] ${agent.name} ${err.message}`);
      resolve();
    });
  });
}

async function main() {
  console.log(`Spawning ${agents.length} agents...`);
  await Promise.all(agents.map(runAgent));

  // Summary
  console.log('\n=== SUMMARY ===');
  for (const a of agents) {
    const hasOutput = fs.existsSync(path.join(a.dir, '0-analysis')) &&
      fs.readdirSync(path.join(a.dir, '0-analysis')).length > 0;
    console.log(`${a.name}: exit=${a.exitCode}, produced_files=${hasOutput ? 'yes' : 'no'}`);
  }
}

main().catch(err => {
  console.error('Orchestrator failed:', err);
  process.exit(1);
});
