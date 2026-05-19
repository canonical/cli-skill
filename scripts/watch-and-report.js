#!/usr/bin/env node
const fs = require('node:fs');
const path = require('node:path');
const { spawn } = require('node:child_process');

const AGENTS = [
  'kimi-k2.6-juju', 'glm-5-juju', 'deepseek-v4-pro-juju',
  'kimi-k2.6-qwen36-snap', 'glm-5-qwen36-snap', 'deepseek-v4-pro-qwen36-snap',
];

function isDone(agentName) {
  const log = fs.existsSync('/project/agents/orchestrator.log') ? fs.readFileSync('/project/agents/orchestrator.log', 'utf-8') : '';
  return log.includes(`[DONE] ${agentName}`);
}

let lastDoneCount = 0;

function poll() {
  const done = AGENTS.filter(isDone);
  if (done.length > lastDoneCount) {
    lastDoneCount = done.length;
    console.log(`[${new Date().toISOString()}] ${done.length}/${AGENTS.length} agents done. Regenerating reports...`);
    const child = spawn('node', ['/project/generate-reports.js'], {
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
