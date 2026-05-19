# Safety Model

## Destructive Operations

| Command | Destructive? | What it destroys | Confirmation | Dry-run | Force/Skip |
|---------|-------------|-----------------|--------------|---------|------------|
| `set` | Partially | Overwrites config values | Yes (interactive prompt showing old→new) | No | `--assume-yes` |
| `unset` | Partially | Removes user config (reverts to default) | Yes (shows revert value) | No | `--assume-yes` |
| `use-engine` | Partially | Replaces engine config layer, may remove/install components | Yes (shows engine details) | No | `--assume-yes` |
| `prune-cache` | Yes | Removes snap components (engine binaries, potentially GBs) | Yes (lists components to remove) | No | No |

## Confirmation Behavior

Mutating commands (`set`, `unset`, `use-engine`, `prune-cache`) prompt interactively before applying changes. The `--assume-yes` flag suppresses these prompts.

- `set`: Shows current value vs. new value, asks "Continue? [y/N]"
- `unset`: Shows which value will revert and to what default
- `use-engine`: Shows engine details and required component installations
- `prune-cache`: Lists all components that will be removed with sizes

## Service Restart Behavior

Commands that modify configuration or engine selection will **automatically restart the snap service** unless `--no-restart` is specified. This means:

- `set http.port=9000` → restarts the inference server (brief downtime)
- `use-engine cuda` → installs components + restarts server

The `--no-restart` flag allows batching multiple config changes before a single manual restart.

## Root Requirement

All write operations require root (`sudo`). The CLI checks `os.Geteuid() == 0` and returns a clear error message directing the user to use `sudo`. This prevents accidental modification by unprivileged users.

## Recovery Behavior

| Failure Scenario | Recovery |
|-----------------|----------|
| Engine switch fails mid-component-install | Previous engine config remains; re-run `use-engine --fix` |
| Config set causes server crash | `unset` the key to revert to default |
| Component not available in store | `use-engine --auto` falls back to next compatible engine |
| Server timeout on start | `snap logs qwen36.server` for diagnosis; `snap restart qwen36.server` |

## Gaps and Observations

1. **No `--dry-run` support**: No command offers a preview-only mode. Users cannot see what `use-engine --auto` would select without actually performing the switch.
2. **No undo/rollback**: Once `prune-cache` removes components, they must be re-downloaded from the store. There is no local backup or undo.
3. **Restart is opt-out, not opt-in**: The default behavior of restarting the server after any config change is surprising for users making exploratory changes. Many CLIs default to no-restart and require explicit action.
4. **`--assume-yes` without `--assume-no`**: There is no way to script a "refuse all" path for testing what would be prompted.
