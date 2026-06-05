# Command: /cli-behavioral-analysis

## Execution Order

1. Run `shared/cli-discovery-preflight.md`
2. Use preflight outputs from `cli-review/0-cli-discovery-preflight/`
3. Analyze CLI runtime behavior, contracts, and safety properties.

**Before starting this command**, re-read your preflight outputs (`architecture.md`, `commandset.md`, `argument-structure.md`) to load them into context. Reference specific commands and arguments from those files rather than re-exploring the codebase from scratch.

Files to produce (in order):

4. `cli-review/1-behavioral-analysis/00-configuration-model.md`
	- Describe config sources and precedence: flags, env vars, config files, defaults.
	- Note command-specific overrides and any surprising precedence behavior.

5. `cli-review/1-behavioral-analysis/01-output-contracts.md`
	- Describe output formats by command (human-readable and machine-readable).
	- Document stability expectations for output fields and parseability guidance.

6. `cli-review/1-behavioral-analysis/02-error-model-and-exit-codes.md`
	- Map error categories to representative messages and exit codes.
	- Include per-command or per-command-group differences.

7. `cli-review/1-behavioral-analysis/03-safety-model.md`
	- Describe destructive operations, confirmations, dry-run support, force flags, and recovery behavior.

**Phase 1 checkpoint**: Verify that all four Phase 1 files exist in `cli-review/1-behavioral-analysis/` and are non-empty. Do not proceed until confirmed.
