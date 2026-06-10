
# cli-skill

cli-skill is a command based, cross agent CLI review framework. It uses one canonical command set and thin adapters so the same workflows can run in Copilot, Claude Code, Pi Coding Agent, and OpenCode.

## Project Structure

- Canonical skill index: [cli-skill/SKILL.md](cli-skill/SKILL.md)
- Command manifest: [cli-skill/schemas/commands.manifest.yaml](cli-skill/schemas/commands.manifest.yaml)
- Shared preflight module: [cli-skill/shared/cli-discovery-preflight.md](cli-skill/shared/cli-discovery-preflight.md)
- Standard reference: [cli-skill/references/cli-standard.md](cli-skill/references/cli-standard.md)

## Commands

- [cli-skill/commands/cli-review.md](cli-skill/commands/cli-review.md)
- [cli-skill/commands/cli-check-help.md](cli-skill/commands/cli-check-help.md)
- [cli-skill/commands/cli-semantic-analysis.md](cli-skill/commands/cli-semantic-analysis.md)
- [cli-skill/commands/cli-heuristic-analysis.md](cli-skill/commands/cli-heuristic-analysis.md)

Future scaffolds:

- [cli-skill/future-commands/cli-propose-command.md](cli-skill/future-commands/cli-propose-command.md)
- [cli-skill/future-commands/cli-rename-command.md](cli-skill/future-commands/cli-rename-command.md)

## Use with Different Agents

1. Sync adapters from the manifest (single maintained helper script):

	node scripts/sync-cli-skill-adapters.js

2. Copilot (GitHub skills path):

- Entry point: [.github/skills/cli-skill/SKILL.md](.github/skills/cli-skill/SKILL.md)
- Adapter file: [cli-skill/adapters/copilot/SKILL.md](cli-skill/adapters/copilot/SKILL.md)

3. Pi Coding Agent:

- Entry point: [.pi/skills/cli-skill/SKILL.md](.pi/skills/cli-skill/SKILL.md)
- Adapter file: [cli-skill/adapters/pi-coding-agent/SKILL.md](cli-skill/adapters/pi-coding-agent/SKILL.md)

4. Claude Code:

- Adapter file: [cli-skill/adapters/claude-code/commands.yaml](cli-skill/adapters/claude-code/commands.yaml)

5. OpenCode:

- Adapter file: [cli-skill/adapters/opencode/commands.json](cli-skill/adapters/opencode/commands.json)

6. Adapter docs:

- [cli-skill/adapters/README.md](cli-skill/adapters/README.md)

## Quick Start for GitHub Action: cli-skill-build

Use the packaged workflow at [.github/workflows/cli-skill-build.yml](.github/workflows/cli-skill-build.yml).

### 1. Configure Secret

Add this repository secret:

- OPENROUTER_API_KEY

The workflow uses GITHUB_TOKEN for PR comment updates.

### 2. Automatic PR Reviews

On pull requests (opened, synchronize, reopened, ready_for_review), the workflow runs:

- command: /cli-review
- provider: openrouter
- model: moonshotai/kimi-k2.6

It publishes a markdown report in PR comments and updates the same report on subsequent PR updates.

### 3. Manual Runs

Run the workflow from the Actions tab with workflow_dispatch.

Optional inputs:

- pr_number (set to post or update a PR report comment)
- command (/cli-review, /cli-check-help, /cli-semantic-analysis, /cli-heuristic-analysis)
- target_path

### 4. Notes

- /cli-review is standards compliance only and evaluates against [cli-skill/references/cli-standard.md](cli-skill/references/cli-standard.md).
- Session exports are uploaded as workflow artifacts.

## Use from Another Repository (One-File Setup)

Another repository can consume the reusable workflow by adding a single workflow file. No helper scripts, skill files, or metadata files are required in the consumer repository.

### 1. Add a workflow in the consumer repository

Create a workflow file in the consumer repository (for example, .github/workflows/cli-review.yml):

```yaml
name: CLI Skill Review

on:
  pull_request:
    types: [opened, synchronize, reopened, ready_for_review]

permissions:
  contents: read
  pull-requests: write

jobs:
  cli-review:
    uses: shaftoe/cli-skill/.github/workflows/cli-skill-review-reusable.yml@v1
    with:
      command: /cli-review
      provider: openrouter
      model: moonshotai/kimi-k2.6
      thinking_level: medium
      pr_number: ${{ github.event.pull_request.number }}
      post_pr_comment: true
      fail_on_agent_error: true

      # Optional input-based path filters (comma or newline separated globs)
      cli_paths_include: |
        cmd/**
        internal/cli/**
        scripts/demo_cli.py
      cli_paths_exclude: |
        docs/archive/**
        **/*.md

      # Optional strict mode. When true, knowledge_base/CLI.md is required
      # if cli_paths_include is not provided.
      enforce_cli_metadata: false
    secrets:
      llm_token: ${{ secrets.OPENROUTER_API_KEY }}
      gh_token: ${{ secrets.GITHUB_TOKEN }}
```

### 2. Configure secrets in the consumer repository

- OPENROUTER_API_KEY

The workflow uses the repository-provided GITHUB_TOKEN for PR APIs.

### 3. Inputs contract

Required inputs:

- provider

Optional inputs:

- command (default: /cli-review)
- target_path (default: .)
- model (default: moonshotai/kimi-k2.6)
- thinking_level (default: medium)
- pr_number
- post_pr_comment (default: true)
- fail_on_agent_error (default: true)
- cli_paths_include (default: empty)
- cli_paths_exclude (default: empty)
- enforce_cli_metadata (default: false)

Required secrets:

- llm_token
- gh_token

### 4. Version pinning

Use a tagged major version (for example, @v1) for a stable interface. Use a commit SHA pin for maximum reproducibility.
