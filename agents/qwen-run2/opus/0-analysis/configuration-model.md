# Configuration Model

## Configuration Sources and Precedence

The qwen36 snap uses a single-source configuration model centered on snapd's configuration system (`snapctl get`/`snapctl set`). There is no config file, no environment variable override mechanism, and no flag-based runtime override (with one exception).

### Source Hierarchy (highest to lowest precedence)

| Priority | Source | Mechanism | Scope |
|----------|--------|-----------|-------|
| 1 | `snap set qwen36 <key>=<value>` | User sets via snap CLI | Persists across restarts |
| 2 | `qwen36 set <key>=<value>` | CLI wrapper around snapctl set | Same as above |
| 3 | `qwen36 use-engine` | Writes engine-specific configs from engine.yaml | Overwrites engine-related keys |
| 4 | `qwen36 set --package <key>=<value>` | Install hook sets package defaults | Lowest priority, overridable by user |
| 5 | Hardcoded defaults in scripts | Fallback values in shell scripts | Used when key is unset |

### Effective Configuration Flow

1. **Install time**: The `install` hook runs `qwen36 set --package` to establish defaults (`http.port=8326`, `http.host=127.0.0.1`, `verbose=false`), then `qwen36 use-engine --auto --assume-yes` which writes engine-specific keys.
2. **User override**: Users can run `qwen36 set http.port=9000` or `snap set qwen36 http.port=9000` to override any value.
3. **Engine switch**: Running `qwen36 use-engine cuda` overwrites engine-related keys (`server`, `model`, `multimodel-projector`, `http.base-path`, `gpu-layers`) from the engine.yaml's `configurations:` block.
4. **Runtime read**: Scripts and the server read config via `qwen36 get <key>` at launch time. There is no hot-reload; the server must be restarted to pick up changes.

## Configuration Keys

| Key | Type | Default | Set By | Used By |
|-----|------|---------|--------|---------|
| `http.port` | integer | `8326` | install hook | server scripts, chat.sh, check-server |
| `http.host` | string | `127.0.0.1` | install hook | server scripts |
| `http.base-path` | string | `v1` | use-engine (from engine.yaml) | chat.sh, check-server |
| `verbose` | boolean | `false` | install hook | server scripts (adds `--verbose` flag) |
| `model-name` | string | *(unset)* | unclear (may be set by CLI) | chat.sh, check-server (optional) |
| `server` | string | e.g., `llamacpp` | use-engine | common.sh (component path) |
| `model` | string | e.g., `model-qwen36-35b-a3b-ud-q4-k-xl` | use-engine | common.sh (component path) |
| `multimodel-projector` | string | e.g., `mmproj-qwen36-35b-a3b-f16` | use-engine | common.sh (component path) |
| `gpu-layers` | integer | `99` (cuda only) | use-engine | cuda/server script |

## Surprising Behaviors

### No environment variable support
Unlike most CLI tools, there is no `QWEN36_PORT` or similar env var to override configuration. All config must go through snapctl.

### Engine switch clobbers multiple keys atomically
Running `use-engine` writes 4-5 keys at once from the engine.yaml `configurations:` block. There is no way to switch engines without also resetting `http.base-path`, `server`, `model`, and `multimodel-projector`.

### `--package` flag creates invisible persistence tier
Values set with `--package` in the install hook behave differently from user-set values during snap refreshes, but this distinction is invisible to `qwen36 get` output.

### Fallback in chat.sh for http.base-path
`chat.sh` has a hardcoded fallback: if `http.base-path` is empty, it defaults to `v1`. This creates a shadow default that could diverge from the engine.yaml declaration.

### model-name is silently optional
Several scripts use `2>/dev/null || true` when reading `model-name`, making it a config key that may or may not exist. Its purpose and how it gets set is undocumented.
