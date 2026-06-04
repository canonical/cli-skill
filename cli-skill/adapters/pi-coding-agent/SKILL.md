---
name: cli-skill
description: "Pi Coding Agent adapter for command-per-file CLI workflows."
---

# Pi Coding Agent Adapter

This adapter is intentionally thin and delegates all behavior to files under `cli-skill/`.

## Startup

- Load `cli-skill/schemas/commands.manifest.yaml`
- Route user intent to one command file
- Always run `cli-skill/shared/cli-discovery-preflight.md` first

## Commands

- `/cli-review` -> `cli-skill/commands/cli-review.md`
- `/cli-semantic-analysis` -> `cli-skill/commands/cli-semantic-analysis.md`
- `/cli-heuristic-analysis` -> `cli-skill/commands/cli-heuristic-analysis.md`
- `/cli-check-help` -> `cli-skill/commands/cli-check-help.md`
