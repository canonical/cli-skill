
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

1. Sync adapters from the manifest:

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
- model: gpt-5.4

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
