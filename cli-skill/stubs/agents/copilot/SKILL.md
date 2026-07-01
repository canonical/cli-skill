---
name: cli-skill
description: "Copilot adapter for cross-agent CLI skill commands."
---

# Copilot Adapter

This adapter maps slash-style command intents to command files in `cli-skill/`.

## Resolve Order

1. Read `cli-skill/schemas/commands.manifest.yaml`
2. Resolve command to file path
3. Run `cli-skill/shared/cli-discovery-preflight.md` before any analysis command
4. Execute requested command workflow

## Command Routing

- `/cli-review` -> `cli-skill/commands/cli-review.md`
- `/cli-semantic-analysis` -> `cli-skill/commands/cli-semantic-analysis.md`
- `/cli-heuristic-analysis` -> `cli-skill/commands/cli-heuristic-analysis.md`
- `/cli-check-help` -> `cli-skill/commands/cli-check-help.md`
- `/cli-behavioral-analysis` -> `cli-skill/commands/cli-behavioral-analysis.md`
- `/install-cli-pr-workflow` -> `cli-skill/commands/install-cli-pr-workflow.md`

## Future Command Stubs

- `/cli-propose-command` -> `cli-skill/future-commands/cli-propose-command.md`
- `/cli-rename-command` -> `cli-skill/future-commands/cli-rename-command.md`
