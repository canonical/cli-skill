# cli-skill

Canonical cli-skill provides a CLI review and design skill. It can run in Copilot, Claude Code, Pi Coding Agent, and OpenCode. 
Its main purpose is to provide a review against the [Canonical CLI standard](cli-skill/references/cli-standard.md). The review
can be triggered locally from your agent, or as a github workflow that will execute the /cli-review command using pi-coding-agent.

## Supported models

The review will run on a wide range of models. Depending on the model, you will usually have different finding, especially unrated findings
can vary due to the analytic capabilities of the model used.

### Latest frontier models 
These models (e.g. Opus, GPT, Gemini Pro) will usually perform well without further instructions.

### Open weight and smaller special-purpose models
Some open weight and smaller special-purpose models also run with reliable results. Examples are **DeepSeek V4 Pro, Gemini Flash 3.5+, Mimi2.5**

Other models may perform less reliably, this includes some with excellent coding performance: qwen3-coder, gemma4, nemotron3, kimi2.7-code, minimax-m3

## Install via npm

This repository is also packaged as `canonical-cli-skill` for consumers that want to install the skill files into their own repo.

```bash
npx canonical-cli-skill install
```

By default, the installer resolves the target directory to the top-level of the current Git repository, so generated paths are anchored at the repo root even when you run it from a subdirectory. If there is no Git repository, it falls back to the current directory. Use `--target <path>` to override this behavior. Existing unmanaged files are not overwritten unless `--force` is used, and managed files are only updated when they have not been edited since the last install.

Run the skill in your agent:
```
/cli-review
```

## Commands

| Command | Description |
|---------|-------------|
| /cli-review | CLI Review [cli-skill/commands/cli-review.md](cli-skill/commands/cli-review.md) |
| /cli-semantic-analysis | Semantic analysis of words used in commands and flags [cli-semantic-analysis.md](cli-skill/commands/cli-semantic-analysis.md) |
| /cli-check-help | CLI Help analysis and considerations [cli-check-help.md](cli-skill/commands/cli-check-help.md) |
| /cli-behavioral-analysis | Run some analysis on the architecture and setup of the CLI, [cli-behavioral-analysis.md](cli-skill/commands/cli-behavioral-analysis.md) |
| /install-cli-pr-workflow | Install the Github workflow in your repository [install-cli-pr-workflow.md](cli-skill/commands/install-cli-pr-workflow.md) |

### Planned future commands:

| Command | Description |
|---------|-------------|
| /cli-heuristic-analysis | Analyze CLI using UX heuristics [cli-heuristic-analysis.md](cli-skill/future-commands/cli-heuristic-analysis.md) |
| /cli-propose-command | Proposition helper [cli-propose-command.md](cli-skill/future-commands/cli-propose-command.md) |
| /cli-rename-command | Renaming helper [cli-rename-command.md](cli-skill/future-commands/cli-rename-command.md) |

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
- model: openrouter/fusion

It publishes a markdown report in PR comments and updates the same report on subsequent PR updates.

### 3. Manual Runs

Run the workflow from the Actions tab with `workflow_dispatch`.

Optional inputs:

- pr_number (set to post or update a PR report comment)
- command (/cli-review, /cli-check-help, /cli-semantic-analysis, /cli-heuristic-analysis)
- target_path
- fusion_analysis_models (comma-separated panel models for Fusion)
- fusion_synthesis_model (synthesis model for Fusion)

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
    uses: canonical/cli-skill/.github/workflows/cli-skill-review-reusable.yml@v1
    with:
      provider: openrouter
      pr_number: ${{ github.event.pull_request.number }}
      # Optional customization, choosing a different model is likely to affect analysis results
      model: google/gemini-3.5-flash

      # Optional input-based path filters (comma or newline separated globs) to save on token cost
      cli_paths_include: |
        cmd/**
      cli_paths_exclude: |
        **/*.md
    secrets:
      llm_token: ${{ secrets.PROVIDER_API_KEY }}
      gh_token: ${{ secrets.GITHUB_TOKEN }}
```

### 2. Configure secrets in the consumer repository

- PROVIDER_API_KEY: configure your provider's API key here for accessing models

The workflow uses the repository-provided GITHUB_TOKEN for PR APIs.

### 3. Inputs contract

Required inputs:

- provider

Optional inputs:

- command (default: /cli-review)
- target_path (default: .)
- model (default: openrouter/fusion)
- thinking_level (default: medium)
- fusion_analysis_models (default: ~anthropic/claude-opus-latest,~openai/gpt-latest,~google/gemini-pro-latest)
- fusion_synthesis_model (default: ~openai/gpt-latest)
- pr_number
- post_pr_comment (default: true)
- fail_on_agent_error (default: true)
- cli_paths_include (default: empty)
- cli_paths_exclude (default: empty)
- enforce_cli_metadata (default: false)

Required secrets:

- llm_token
- gh_token

### 4. Version pinning and version history

Use a tagged major version (for example, @v1) for a stable interface. Use a commit SHA pin for maximum reproducibility.

v1 - current stable version <-- use this
