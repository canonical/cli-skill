# Configuration Model: inference-snaps-cli

## Sources and Precedence

Configuration is stored via `snapctl` under the snap's configuration namespace. The CLI implements a **three-layer precedence stack** (lowest to highest):

1. **Package config** (`config.package.*`) — set during install/post-refresh hooks
2. **Engine config** (`config.engine.*`) — set when an engine is activated via `use-engine`
3. **User config** (`config.user.*`) — set explicitly by the user via `set`

When `Get()` or `GetAll()` is called, the storage backend loads all three layers and merges them in order, with higher layers overwriting lower ones. The result is presented as a flat dot-separated key map (e.g., `http.port`, `model-name`).

## Configuration Keys (Observed)

| Key | Typical Value | Set By | Scope |
|---|---|---|---|
| `http.port` | `8326` | install hook (`set --package`) | package |
| `http.host` | `127.0.0.1` | install hook (`set --package`) | package |
| `verbose` | `false` | install hook (`set --package`) | package |
| `http.base-path` | `v1` | engine manifest / user | engine / user |
| `model-name` | (model ID) | engine manifest / user | engine / user |
| `passthrough.environment.*` | arbitrary | user | user |
| `webui.http.port` | (varies) | engine manifest / user | engine / user |

## Command-Specific Overrides

### `set` command
- Without `--package` or `--engine`, writes to **UserConfig**.
- `--package` writes to **PackageConfig** (hidden flag, used by hooks).
- `--engine` writes to **EngineConfig** (hidden flag, used internally during engine switch).
- Validation: keys must already exist unless prefixed with `passthrough.`. Duplicate keys in a single invocation are rejected.
- Side effect: unless `--no-restart` is used, the snap daemon is restarted after any user-visible config change.

### `unset` command
- Only operates on **UserConfig**.
- If the key does not exist in the merged view, returns a "key not found" error with a suggestion to run `get`.
- Side effect: if the effective value changes after unsetting, prompts for daemon restart (respects `--no-restart`).

### `use-engine` command
- Unsets the previous engine's **EngineConfig** and **UserConfig** overrides.
- Writes the new engine's manifest `configurations` map to **EngineConfig** via `SetDocument()`.
- Does not touch **PackageConfig**.

## Storage Backend

- **Hardcoded backend**: `NewSnapctlStorage()` is the only implementation of the `storage` interface.
- No local files, no XDG directories, no `/etc/<name>.conf`.
- Configuration lifecycle is tied to the snap revision: if the snap is removed, config persists in snapd unless purged.

## Surprising Precedence Behavior

1. **Engine config overrides package config without visibility**: When a user runs `get`, the merged view is shown. There is no way from the CLI to see which layer provided a given value. This makes debugging config sources impossible without external `snapctl` inspection.
2. **Unset only targets user layer**: If a value is set in package or engine config, `unset` cannot remove it; it only reverts the effective value to the next lower layer. The help text says "reverting to package or engine default," but users may be confused why `unset` does not make the key disappear entirely when a lower-layer default exists.
3. **Full-wipe on engine change**: `UnsetEngineConfig` unsets ALL engine configs with `key="."` (wildcard), then unsets per-key user overrides. This is a blunt instrument: if two engines share config keys, switching between them still wipes and rewrites the entire engine layer.
4. **No config schema or validation**: Beyond key existence, there is no type checking, range validation, or format enforcement. Setting `http.port=not-a-number` succeeds silently and will fail later when the server tries to bind.
