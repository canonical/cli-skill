# Configuration Model

Sources of configuration (observed in `pkg/storage`, `commands/set.go`, `SetEngineConfig` calls and snap parts):

- Flags: per-command flags (e.g., `--format`, `--auto`). These take precedence for that invocation.
- Environment variables: used for snap environment (e.g., `SNAP`, `SNAP_INSTANCE_NAME`, `ARCH_TRIPLET`).
- Persistent config store: `pkg/storage` exposes `Config` with scopes: `UserConfig`, `EngineConfig`, `PackageConfig`.
- Engine-provided configuration: `common.SetEngineConfig` writes engine-specific defaults during `use-engine`.

Precedence (inferred):
1. Command-line flags
2. Environment variables for runtime values (where used by code)
3. User-config stored in `UserConfig` via `set` command
4. Engine-specific config (`EngineConfig`) applied when an engine is selected
5. Package defaults (`PackageConfig`) shipped with the snap

Notes:
- `set` and `unset` support engine/package/user scopes (hidden flags `--engine`, `--package`) and document restart semantics. `set` will prompt to restart unless `--no-restart`.
- Passthrough keys (`passthrough.*`) are supported and used by `run` to set environment variables into subprocesses.
