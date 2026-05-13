# Agent Swarm Insights

Generated from 6 parallel agent runs across 2 projects and 3 models.

---

## Aggregate Metrics

| Metric | Value |
|---|---|
| Agents launched | 6 |
| Agents with orchestrator completion | 6 |
| Agents with substantial analysis | 6 |
| Agents with commandset design | 6 |
| Total tool calls | 389 |
| Total tool errors | 14 |
| Error rate | 3.6% |

## Model Comparison

### kimi-k2.6

- Runs: 1
- Orchestrator-completed: 1
- Avg tool calls: 62
- Avg analysis files: 9.0
- Avg events: 43850

### glm-5

- Runs: 1
- Orchestrator-completed: 1
- Avg tool calls: 64
- Avg analysis files: 9.0
- Avg events: 25242

### deepseek-v4-pro

- Runs: 1
- Orchestrator-completed: 1
- Avg tool calls: 63
- Avg analysis files: 9.0
- Avg events: 25891

### kimi-k2.6-qwen36

- Runs: 1
- Orchestrator-completed: 1
- Avg tool calls: 77
- Avg analysis files: 9.0
- Avg events: 22885

### glm-5-qwen36

- Runs: 1
- Orchestrator-completed: 1
- Avg tool calls: 59
- Avg analysis files: 9.0
- Avg events: 16468

### deepseek-v4-pro-qwen36

- Runs: 1
- Orchestrator-completed: 1
- Avg tool calls: 64
- Avg analysis files: 9.0
- Avg events: 27568

## Key Findings

1. **All agents produced analysis output**: Every agent wrote at least some 0-analysis files, indicating the skill instructions were sufficiently clear for project exploration.

2. **GLM-5 most efficient**: GLM-5 agents had the lowest event counts while still completing both phases, suggesting more focused exploration.

3. **Kimi K2.6 most verbose**: kimi-k2.6-juju generated over 24,000 events (2GB session log), indicating extensive but potentially unfocused exploration.

4. **DeepSeek V4 Pro inconsistent**: deepseek-v4-pro-juju stalled in analysis phase initially but eventually completed; deepseek-v4-pro-qwen36-snap produced the most design files (6).

5. **Orchestrator reliability issues**: 3 of 6 exit events were missed by the Node.js orchestrator due to large stdout streams. This is a technical limitation, not a model issue.

## Recommendations for the Skill

1. **Chunk the workflow**: Split analyze-cli into smaller sub-tasks to prevent context overflow and excessive exploration.

2. **Add file-count validation**: Instruct agents to verify they have written all required files before declaring completion.

3. **Provide project size hints**: Large projects like juju should include a "known entry points" list to reduce exploratory tool calls.

4. **Model-specific tuning**: Kimi K2.6 benefits from tighter step-by-step constraints; GLM-5 works well with high-level goals.

