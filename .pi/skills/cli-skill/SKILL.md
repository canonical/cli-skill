---
name: cli-skill
description: "Pi entrypoint for cross-agent CLI command workflows."
---

# Entrypoint

Canonical files are in `cli-skill/`.

## Dispatch

- Manifest: `cli-skill/schemas/commands.manifest.yaml`
- Shared preflight: `cli-skill/shared/cli-discovery-preflight.md`
- Commands: `cli-skill/commands/*.md`

## Supported Commands

- `/cli-review`
- `/cli-semantic-analysis`
- `/cli-heuristic-analysis`
- `/cli-check-help`
- `/cli-behavioral-analysis`
- `/install-cli-pr-workflow`
