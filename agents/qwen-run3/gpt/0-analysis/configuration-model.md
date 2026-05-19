# qwen36 Configuration Model

## Sources And Precedence

The CLI implements a clear three-layer precedence model in `pkg/storage/config.go`.

From lowest to highest precedence:

1. package config
2. engine config
3. user config

All layers are stored via `snapctl` under the `config` namespace, then flattened into dot-separated keys for reads.

## Config Sources

| Source | Storage Prefix | Written By | Typical Use |
|---|---|---|---|
| Package defaults | `config.package.*` | snap hooks via `set --package` | install-time defaults such as `http.port` and `http.host` |
| Engine defaults | `config.engine.*` | `use-engine` via manifest `configurations` | selected engine settings such as `server`, `model`, `http.base-path`, `gpu-layers` |
| User overrides | `config.user.*` | public `set` command | user-level overrides of any supported key |
| Runtime fallbacks | not stored | shell wrappers | fallback to `v1` when `http.base-path` is empty |

## Observed Keys

| Key | Likely Type | Source(s) | Notes |
|---|---|---|---|
| `http.port` | integer-like string | package, user | install hook sets `8326` |
| `http.host` | string | package, user | install hook sets `127.0.0.1` |
| `verbose` | boolean-like string | package, user | install hook sets `false` |
| `http.base-path` | string | engine, user | both shipped engines set `v1` |
| `server` | string | engine | engine component name |
| `model` | string | engine | model component name |
| `multimodel-projector` | string | engine | projector component name |
| `gpu-layers` | integer-like string | engine, user | set by the CUDA manifest |
| `model-name` | string | engine or runtime | optional in chat and health-check wrappers |
| `webui.http.port` | integer-like string | package or engine in reusable CLI snap | supported by code, not initialized in qwen36 root hooks |
| `passthrough.environment.*` | arbitrary string | user | exported during `run` |

## Precedence Behavior

### Reads

- `get` and `get all` read the merged effective view.
- lower layers are overwritten by higher layers.
- provenance is not exposed, so users cannot see which layer produced a value.

### Writes

- plain `set` writes to user config
- hidden `set --package` writes to package config
- hidden `set --engine` writes to engine config
- `unset` only removes keys from user config

## Command-Specific Overrides

### `set`

- validates existing keys unless the key begins with `passthrough.`
- accepts multiple key/value pairs in one invocation
- only restarts when at least one effective value changed

### `unset`

- looks up the current effective value first
- removes only the user-layer entry
- if a lower layer exists, the key remains visible after `unset`

### `use-engine`

- unsets the prior engine config layer globally
- optionally removes matching user overrides for the prior engine's keys
- writes the new engine manifest `configurations` map into engine config
- updates `cache.active-engine`

## Surprising Precedence Behavior

1. `unset` does not delete a key from the merged config if engine or package defaults still provide it.
2. engine switches wipe and rewrite the whole engine layer, not only keys that changed.
3. the public CLI offers no way to inspect raw layer values or provenance.
4. `set` performs key-existence checks but not type validation, so invalid port-like strings can be stored and only fail later at runtime.

## Command-Specific Overrides And Restart Semantics

- `set`, `unset`, and `use-engine` all rely on restart prompts rather than automatic service state detection.
- `use-engine` can change many keys at once, making it a high-level config mutation command rather than a simple selector.
- shell wrappers such as `apps/chat.sh` and `apps/check-server-llamacpp.sh` still apply hardcoded fallbacks for missing `http.base-path`, so the runtime is tolerant of partial config.

## Missing Documentation

A complete user-facing config reference should document:

- every supported key
- which layer typically owns it
- valid types and ranges
- whether changes require restart
- whether `use-engine` may overwrite the key
