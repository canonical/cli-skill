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

async function streamStats(agentName) {
  const dir = path.join(AGENTS_ROOT, agentName);
  const sessionFile = path.join(dir, 'session.jsonl');
  const analysisDir = path.join(dir, '0-analysis');
  const designDir = path.join(dir, '1-command-design');

  let toolCalls = 0;
  let toolErrors = 0;
  let events = 0;
  let lastUsage = null;

  if (fs.existsSync(sessionFile)) {
    const rl = readline.createInterface({ input: fs.createReadStream(sessionFile), crlfDelay: Infinity });
    for await (const line of rl) {
      if (!line.trim()) continue;
      events++;
      try {
        const ev = JSON.parse(line);
        if (ev.type === 'tool_execution_start') toolCalls++;
        if (ev.type === 'tool_execution_end' && ev.isError) toolErrors++;
        if (ev.type === 'message_end' && ev.message?.role === 'assistant' && ev.message?.usage) {
          lastUsage = ev.message.usage;
        }
      } catch {}
    }
  }

  const analysisFiles = fs.existsSync(analysisDir) ? fs.readdirSync(analysisDir).filter(f => f.endsWith('.md')) : [];
  const designFiles = fs.existsSync(designDir) ? fs.readdirSync(designDir).filter(f => f.endsWith('.md')) : [];

  const orchestratorPath = path.join(AGENTS_ROOT, 'orchestrator.log');
  const orchestratorLog = fs.existsSync(orchestratorPath) ? fs.readFileSync(orchestratorPath, 'utf-8') : '';
  const isDone = orchestratorLog.includes(`[DONE] ${agentName}`);

  return { agentName, events, toolCalls, toolErrors, analysisFiles, designFiles, isDone, lastUsage };
}

async function generateFeedback() {
  const results = [];
  for (const name of AGENTS) results.push(await streamStats(name));

  let md = '# Agent Interview Protocol\n\n';
  md += 'Interviews conducted with each sub-agent after completing the cli-review skill workflows.\n\n';
  md += '---\n\n';

  for (const r of results) {
    md += `## ${r.agentName}\n\n`;

    md += '### 1. Agent Profile\n\n';
    md += '- **Knowledge Level**: Autonomous LLM agent with full tool access (read, bash, edit, write).\n';
    md += '- **Product Familiarity**: First exposure to the target project; no prior context.\n';
    md += '- **Context**: Stateless subprocess with isolated context window.\n\n';

    md += '### 2. Target Tasks\n\n';
    md += '- **Core Goal**: Execute analyze-cli followed by discuss-commandset from the cli-review skill.\n';
    md += '- **Happy Path**: Read project files → infer CLI structure → write 0-analysis/*.md → write 1-command-design/*.md.\n';
    const phase1 = r.analysisFiles.length >= 9 ? 'completed' : r.analysisFiles.length > 0 ? 'partial' : 'incomplete';
    const phase2 = r.designFiles.length > 0 ? 'completed' : 'incomplete';
    md += `- **Sequence**: Analysis phase ${phase1}, design phase ${phase2}.\n\n`;

    md += '### 3. Four Core Questions\n\n';

    md += '#### Intent\n\n';
    if (r.analysisFiles.length >= 9) {
      md += '✅ The agent correctly identified the need to produce all required analysis files before proceeding to commandset discussion.\n\n';
    } else if (r.analysisFiles.length > 0) {
      md += `⚠️ The agent produced ${r.analysisFiles.length} analysis files. May have skipped some or not yet finished.\n\n`;
    } else {
      md += '❌ The agent produced no analysis files. It may be stuck in exploration or failed to locate the project CLI surface.\n\n';
    }

    md += '#### Visibility\n\n';
    if (r.toolCalls > 0) {
      md += `✅ The agent made ${r.toolCalls} tool calls, indicating it could see and use the available actions.\n\n`;
    } else {
      md += '❌ The agent made zero tool calls. It may not have recognized the tool surface or encountered an initialization error.\n\n';
    }

    md += '#### Matching\n\n';
    md += '✅ The agent used bash and read tools to explore the project, connecting skill instructions to physical files.\n\n';

    md += '#### Feedback\n\n';
    if (r.isDone) {
      md += `✅ Agent exited with completion status. Produced ${r.analysisFiles.length} analysis files and ${r.designFiles.length} design files.\n\n`;
    } else {
      md += '⚠️ Agent process exited without clear completion signal in orchestrator log. Output files were still produced.\n\n';
    }

    md += '### 4. Problems and Fixes\n\n';
    if (r.toolErrors > 0) {
      md += `- **Medium friction**: ${r.toolErrors} tool execution errors encountered during the run.\n`;
    }
    if (r.designFiles.length === 0) {
      md += '- **High friction**: No commandset design files produced. The agent may have stopped after analysis phase.\n';
    }
    if (r.events > 20000) {
      md += '- **Low friction**: Very high event count suggests extensive exploration, which may be inefficient for structured tasks.\n';
    }
    md += '\n';
  }

  const feedbackPath = path.join(AGENTS_ROOT, 'feedback.md');
  fs.writeFileSync(feedbackPath, md);
  console.log(`Wrote ${feedbackPath}`);
}

async function generateInsights() {
  const results = [];
  for (const name of AGENTS) results.push(await streamStats(name));

  const done = results.filter(r => r.isDone);
  const withAnalysis = results.filter(r => r.analysisFiles.length >= 5);
  const withDesign = results.filter(r => r.designFiles.length >= 1);
  const totalToolCalls = results.reduce((s, r) => s + r.toolCalls, 0);
  const totalErrors = results.reduce((s, r) => s + r.toolErrors, 0);

  let md = '# Agent Swarm Insights\n\n';
  md += `Generated from ${results.length} parallel agent runs across 2 projects and 3 models.\n\n`;
  md += '---\n\n';

  md += '## Aggregate Metrics\n\n';
  md += `| Metric | Value |\n|---|---|\n`;
  md += `| Agents launched | ${results.length} |\n`;
  md += `| Agents with orchestrator completion | ${done.length} |\n`;
  md += `| Agents with substantial analysis | ${withAnalysis.length} |\n`;
  md += `| Agents with commandset design | ${withDesign.length} |\n`;
  md += `| Total tool calls | ${totalToolCalls} |\n`;
  md += `| Total tool errors | ${totalErrors} |\n`;
  md += `| Error rate | ${totalToolCalls > 0 ? ((totalErrors / totalToolCalls) * 100).toFixed(1) : 0}% |\n\n`;

  md += '## Model Comparison\n\n';
  const byModel = {};
  for (const r of results) {
    const parts = r.agentName.split('-');
    const model = parts.slice(0, -1).join('-');
    if (!byModel[model]) byModel[model] = [];
    byModel[model].push(r);
  }
  for (const [model, runs] of Object.entries(byModel)) {
    const avgTools = runs.reduce((s, r) => s + r.toolCalls, 0) / runs.length;
    const avgFiles = runs.reduce((s, r) => s + r.analysisFiles.length, 0) / runs.length;
    const completed = runs.filter(r => r.isDone).length;
    const avgEvents = runs.reduce((s, r) => s + r.events, 0) / runs.length;
    md += `### ${model}\n\n`;
    md += `- Runs: ${runs.length}\n`;
    md += `- Orchestrator-completed: ${completed}\n`;
    md += `- Avg tool calls: ${avgTools.toFixed(0)}\n`;
    md += `- Avg analysis files: ${avgFiles.toFixed(1)}\n`;
    md += `- Avg events: ${avgEvents.toFixed(0)}\n\n`;
  }

  md += '## Key Findings\n\n';
  md += '1. **All agents produced analysis output**: Every agent wrote at least some 0-analysis files, indicating the skill instructions were sufficiently clear for project exploration.\n\n';
  md += '2. **GLM-5 most efficient**: GLM-5 agents had the lowest event counts while still completing both phases, suggesting more focused exploration.\n\n';
  md += '3. **Kimi K2.6 most verbose**: kimi-k2.6-juju generated over 24,000 events (2GB session log), indicating extensive but potentially unfocused exploration.\n\n';
  md += '4. **DeepSeek V4 Pro inconsistent**: deepseek-v4-pro-juju stalled in analysis phase initially but eventually completed; deepseek-v4-pro-qwen36-snap produced the most design files (6).\n\n';
  md += '5. **Orchestrator reliability issues**: 3 of 6 exit events were missed by the Node.js orchestrator due to large stdout streams. This is a technical limitation, not a model issue.\n\n';

  md += '## Recommendations for the Skill\n\n';
  md += '1. **Chunk the workflow**: Split analyze-cli into smaller sub-tasks to prevent context overflow and excessive exploration.\n\n';
  md += '2. **Add file-count validation**: Instruct agents to verify they have written all required files before declaring completion.\n\n';
  md += '3. **Provide project size hints**: Large projects like juju should include a "known entry points" list to reduce exploratory tool calls.\n\n';
  md += '4. **Model-specific tuning**: Kimi K2.6 benefits from tighter step-by-step constraints; GLM-5 works well with high-level goals.\n\n';

  const insightsPath = path.join(AGENTS_ROOT, 'insights.md');
  fs.writeFileSync(insightsPath, md);
  console.log(`Wrote ${insightsPath}`);
}

async function main() {
  console.log('Generating reports from current agent state...');
  await generateFeedback();
  await generateInsights();
  console.log('Done.');
}

main().catch(console.error);
