---
name: cli-skill
description: "Cross-agent CLI skill with command-per-file architecture. Primary command: /cli-review."
---

# cli-skill

This skill is designed to work across multiple agents (Copilot, Claude Code, Pi Coding Agent, and OpenCode) using a command-per-file architecture.

## Architecture

- One command definition per file under `commands/`
- Shared internal helper modules under `shared/`
- Future command scaffolds under `future-commands/`
- Commands are loaded on demand; avoid loading all command files at once

## Shared Internal Module

Before any analysis command, run:

- `shared/cli-discovery-preflight.md`

The preflight output path is fixed to:

- `cli-review/0-cli-discovery-preflight/`

## Active Commands

- `/cli-review` -> `commands/cli-review.md`
- `/cli-semantic-analysis` -> `commands/cli-semantic-analysis.md`
- `/cli-heuristic-analysis` -> `commands/cli-heuristic-analysis.md`
- `/cli-check-help` -> `commands/cli-check-help.md`
- `/install-cli-pr-workflow` -> `commands/install-cli-pr-workflow.md`

## Future Commands (Scaffolding Only)

- `/cli-propose-command` -> `future-commands/cli-propose-command.md`
- `/cli-rename-command` -> `future-commands/cli-rename-command.md`
