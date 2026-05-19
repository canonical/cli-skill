# Configuration Model

## Configuration Sources and Storage

The `qwen36` CLI uses **snapctl** as its sole configuration backend. All configuration reads and writes go through `snapctl get/set/unset` commands via the `go-snapctl` library. The storage layer is abstracted through the `storage.Config` interface with a single implementation: `SnapctlStorage`.

### Storage Backend

```
snapctl get <snap-name> config.<type>.<key>
snapctl set <snap-name> config.<type>.<key>=<value>
snapctl unset <snap-name> config.<type>.<key>
```

All keys are namespaced under `config.` prefix, followed by the config type (`package`, `engine`, or `user`), then the dotted key path.

## Configuration Precedence Tiers

Three tiers, from lowest to highest precedence:

| Tier | Config Type | Set By | Behavior |
|------|------------|--------|----------|
| 1 (lowest) | `PackageConfig` (`"package"`) | Snap package installation/hooks | Base defaults shipped with the snap |
| 2 | `EngineConfig` (`"engine"`) | `use-engine` / engine manifest `configurations` field | Engine-specific overrides applied when switching engines |
| 3 (highest) | `UserConfig` (`"user"`) | `qwen36 set <key=value>` | User overrides, persists across engine switches |

### Precedence Resolution

When reading a config value via `Config.Get(key)` or `Config.GetAll()`:
1. Load all tiers starting from lowest precedence.
2. Each subsequent tier **overwrites** matching keys from lower tiers.
3. The result is a flat map with dot-separated keys (e.g., `http.port`).

When a user unsets a key via `qwen36 unset <key>`, only the `UserConfig` tier is removed, exposing whatever value exists in `EngineConfig` or `PackageConfig`.

## Key Configuration Values

### Known Config Keys

| Key | Description | Set By |
|-----|-------------|--------|
| `http.port` | Port for the inference server HTTP endpoint | Package/Engine |
| `webui.http.port` | Port for the web UI HTTP server | Package/Engine |
| `openai` | OpenAI-compatible server endpoint name (referenced internally) | Engine |
| `protocol` | Server protocol (`http`/`https`) per component server | Engine |
| `base-path` | Base path for each server endpoint | Engine |
| `passthrough.*` | Passthrough environment variables (e.g., `passthrough.environment.my-key`) | User |

### Passthrough Keys

Keys prefixed with `passthrough.` follow special semantics:
- In `set`, they bypass the "key must be found" validation (allowing new passthrough keys).
- In `run`, `passthrough.environment.<key>` values are converted to uppercase underscored environment variables (e.g., `passthrough.environment.my-var` → `MY_VAR`) and injected into the subprocess environment.

## Engine Configuration Application

When `use-engine` switches to a new engine:
1. **Unset**: All previous `EngineConfig` values are cleared (`config.Unset(".", EngineConfig)` — the `"."` key clears the entire engine config namespace).
2. **Unset user overrides**: User overrides for keys defined in the old engine's manifest are removed.
3. **Set**: New engine's `configurations` from its YAML manifest are written to `EngineConfig`.
4. **Restart prompt**: User is prompted to restart the snap to apply changes.

## Cache (Separate from Config)

The CLI also maintains a **cache** namespace via snapctl (prefix `cache.`):
- `cache.active-engine`: Stores the name of the currently active engine (set by `use-engine`, read by many commands).

The cache is **not** subject to the three-tier precedence model — it's a simple key-value store with atomic get/set semantics.

## Surprising Behaviors

### `unset` only operates on `UserConfig`

There is **no way** for users to unset engine or package configs via the CLI. `unset` always targets `UserConfig`. The internal `--package` and `--engine` flags on `set` are hidden and undocumented for end users. This means engine-level configs can only be changed by switching engines, not by directly manipulating them.

### No config file support

The CLI has **no config file** (no `.yaml`, `.json`, `.toml`). All configuration is snapctl-backed. This is by design — the CLI is snap-native — but it means there's no way to view or edit configs outside the snapctl interface.

### `get` with no args shows merged view

`qwen36 get` (no key) shows all configs after applying precedence. `qwen36 get <key>` shows the resolved value after precedence. There is **no command to view a specific tier** (e.g., "show me only the user overrides").

### `no-restart` default is `false` but must be explicit

On `set` and `unset`, the default behavior is to prompt the user to restart. Users can suppress this with `--no-restart` or auto-confirm with `--assume-yes`. On `use-engine`, the same behavior applies with the same flags.

### Environment variable for feature gating

The `ADDITIONAL_FEATURES` environment variable is read at startup to determine whether `chat` and `webui` commands are available. This is a configuration mechanism outside the snapctl config system — it's a process-environment-level toggle rather than a persisted config.
