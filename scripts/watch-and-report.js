#!/usr/bin/env node
const fs = require('node:fs');
const path = require('node:path');
const { spawn } = require('node:child_process');
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
const GENERATE_REPORTS_SCRIPT = path.join(ROOT, 'scripts', 'generate-reports.js');

const AGENTS = [
  'kimi-k2.6-juju', 'glm-5-juju', 'deepseek-v4-pro-juju',
  'kimi-k2.6-qwen36-snap', 'glm-5-qwen36-snap', 'deepseek-v4-pro-qwen36-snap',
];

function isDone(agentName) {
  const orchestratorPath = path.join(AGENTS_ROOT, 'orchestrator.log');
  const log = fs.existsSync(orchestratorPath) ? fs.readFileSync(orchestratorPath, 'utf-8') : '';
  return log.includes(`[DONE] ${agentName}`);
}

let lastDoneCount = 0;

function poll() {
  const done = AGENTS.filter(isDone);
  if (done.length > lastDoneCount) {
    lastDoneCount = done.length;
    console.log(`[${new Date().toISOString()}] ${done.length}/${AGENTS.length} agents done. Regenerating reports...`);
    const child = spawn('node', [GENERATE_REPORTS_SCRIPT], {
      stdio: ['ignore', 'pipe', 'pipe'],
    });
    child.stdout.on('data', d => process.stdout.write(d));
    child.stderr.on('data', d => process.stderr.write(d));
  }
  if (done.length < AGENTS.length) {
    setTimeout(poll, 30000); // check every 30s
  } else {
    console.log(`[${new Date().toISOString()}] All agents complete. Final reports generated.`);
    process.exit(0);
  }
}

console.log('Watcher started. Checking every 30 seconds...');
poll();
