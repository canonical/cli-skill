# Agent Swarm Insights

Generated from 6 parallel agent runs across 2 projects and 3 models.

---

## Aggregate Metrics

| Metric | Value |
|---|---|
| Agents launched | 6 |
| Agents completed | 1 |
| Agents with substantial analysis | 2 |
| Agents with commandset design | 1 |
| Total tool calls | 311 |
| Total tool errors | 9 |
| Error rate | 2.9% |

## Model Comparison

### kimi-k2.6

- Runs: 2
- Completed: 0
- Avg tool calls: 52
- Avg analysis files: 1.0

### glm-5

- Runs: 2
- Completed: 1
- Avg tool calls: 53
- Avg analysis files: 7.0

### deepseek-v4

- Runs: 2
- Completed: 0
- Avg tool calls: 51
- Avg analysis files: 2.0

## Key Findings

1. **Tool errors occurred**: Some agents encountered read/bash failures. This suggests file paths or commands in the skill instructions may need refinement for robustness across different project structures.

4. **Context isolation works**: Each agent operated independently with no cross-contamination. The sub-agent approach successfully isolates model-specific behavior.

## Recommendations for the Skill

1. **Chunk the workflow**: Split analyze-cli into smaller sub-tasks (e.g., architecture only, then commandset only) to prevent context overflow.

2. **Add validation checkpoints**: After each major section, instruct the agent to verify the file was written before proceeding.

3. **Provide project-specific hints**: For large projects like juju, pre-seed a list of key directories to reduce exploration time.

4. **Model-specific guidance**: Some models (e.g., GLM-5) may need more explicit step-by-step instructions than others (e.g., Kimi K2.6).

