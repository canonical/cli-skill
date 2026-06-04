# Architecture Decisions

1. Canonical-first, adapter-thin model
- `cli-skill/` is the single source of truth.
- Agent-specific files are treated as generated adapters, not primary definitions.

2. Command-per-file contract
- Skill behavior is split by command file for on-demand loading (no monolithic skill prompt).
- Shared logic is isolated into internal modules under `cli-skill/shared/`.

3. Fixed preflight artifact location
- CLI discovery preflight always writes to `cli-review/0-cli-discovery-preflight/`.
- This path is intentionally stable for downstream command chaining and workflow automation.

4. Strict scope for `/cli-review`
- `/cli-review` is compliance-only.
- It must evaluate only against `cli-skill/references/cli-standard.md`; non-standard checks are explicitly excluded.

5. Deferred command evolution
- `propose/rename` are intentionally scaffold-only future commands.
- They are separated from current active command execution to keep v1 deterministic.

6. Manifest-driven cross-agent sync
- `cli-skill/schemas/commands.manifest.yaml` is the command routing contract.
- Adapter entrypoints are regenerated from the manifest via `scripts/sync-cli-skill-adapters.js`.

7. Drift prevention at commit time
- Pre-commit hook enforces adapter sync and blocks commit when generated adapter files changed but are unstaged.
- This keeps canonical + adapter views synchronized.

8. GitHub PR reporting strategy
- Reusable workflow posts markdown report to PR and upserts (updates in place) using a stable hidden marker.
- Concurrency is PR-scoped to avoid duplicate comment races during rapid updates.

9. Packaged workflow entrypoint
- `cli-skill-build` is the simplified consumer-facing workflow wrapper over the reusable workflow.
- Defaults target OpenRouter with minimal required configuration.

10. Reserved workflow secret naming
- Reusable workflow uses `gh_token` (not `github_token`) for `workflow_call` secret interface to avoid reserved-name collision.

11. Repository-root path policy for scripts
- Operational scripts resolve repository root at runtime via git (`rev-parse --show-toplevel`) with fallback to script-relative root.
- Hard-coded absolute workspace paths are intentionally avoided.
