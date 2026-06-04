const fs = require('node:fs');
const path = require('node:path');
const readline = require('node:readline');
const { execSync } = require('node:child_process');

function getRepoRoot() {
  try {
    return execSync('git rev-parse --show-toplevel', { encoding: 'utf-8' }).trim();
  } catch {
    return path.resolve(__dirname, '..');
  }
}

const ROOT = getRepoRoot();
const AGENTS_ROOT = path.join(ROOT, 'agents');

const AGENTS = [
  'kimi-k2.6-juju', 'glm-5-juju', 'deepseek-v4-pro-juju',
  'kimi-k2.6-qwen36-snap', 'glm-5-qwen36-snap', 'deepseek-v4-pro-qwen36-snap',
];

function streamLines(filePath) {
  return new Promise((resolve) => {
    if (!fs.existsSync(filePath)) { resolve(0); return; }
    let count = 0;
    const rl = readline.createInterface({ input: fs.createReadStream(filePath), crlfDelay: Infinity });
    rl.on('line', () => count++);
    rl.on('close', () => resolve(count));
    rl.on('error', () => resolve(0));
  });
}

async function streamStats(agentName) {
  const dir = path.join(AGENTS_ROOT, agentName);
  const sessionFile = path.join(dir, 'session.jsonl');
  const analysisDir = path.join(dir, '0-analysis');
  const designDir = path.join(dir, '1-command-design');
  
  const sessionLines = await streamLines(sessionFile);
  const analysisFiles = fs.existsSync(analysisDir) ? fs.readdirSync(analysisDir).filter(f => f.endsWith('.md')).length : 0;
  const designFiles = fs.existsSync(designDir) ? fs.readdirSync(designDir).filter(f => f.endsWith('.md')).length : 0;
  
  const orchestratorPath = path.join(AGENTS_ROOT, 'orchestrator.log');
  const orchestratorLog = fs.existsSync(orchestratorPath) ? fs.readFileSync(orchestratorPath, 'utf-8') : '';
  const isDone = orchestratorLog.includes(`[DONE] ${agentName}`);
  
  let toolCalls = 0;
  let toolErrors = 0;
  
  if (fs.existsSync(sessionFile)) {
    const rl = readline.createInterface({ input: fs.createReadStream(sessionFile), crlfDelay: Infinity });
    for await (const line of rl) {
      if (!line.trim()) continue;
      try {
        const ev = JSON.parse(line);
        if (ev.type === 'tool_execution_start') toolCalls++;
        if (ev.type === 'tool_execution_end' && ev.isError) toolErrors++;
      } catch {}
    }
  }
  
  return { agentName, sessionLines, analysisFiles, designFiles, isDone, toolCalls, toolErrors };
}

async function status() {
  console.log('=== Agent Status ===\n');
  let done = 0;
  let running = 0;
  for (const name of AGENTS) {
    const r = await streamStats(name);
    const icon = r.isDone ? '✅' : '🔄';
    console.log(`${icon} ${r.agentName}`);
    console.log(`   events: ${r.sessionLines.toLocaleString()} | tools: ${r.toolCalls} | errors: ${r.toolErrors}`);
    console.log(`   0-analysis: ${r.analysisFiles} files | 1-command-design: ${r.designFiles} files`);
    if (r.isDone) done++; else running++;
  }
  console.log(`\nDone: ${done} | Running: ${running}`);
}

status().catch(console.error);
