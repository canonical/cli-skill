---
name: cli-skill
description: "GitHub Copilot entrypoint for the cross-agent CLI skill."
---

# Entrypoint

This is an adapter entrypoint. The canonical implementation lives in `cli-skill/`.

## Load Order

1. Read `cli-skill/schemas/commands.manifest.yaml`
2. Resolve command file for the requested command
3. Run `cli-skill/shared/cli-discovery-preflight.md`
4. Execute command workflow file

## Commands

- `/cli-review`
- `/cli-semantic-analysis`
- `/cli-heuristic-analysis`
- `/cli-check-help`
