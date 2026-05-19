# Configuration Model

## Configuration Sources and Precedence

The configuration system has three layers, in order of precedence (highest first):

1. **User-level snap options** (`qwen36 set key=value`) — highest precedence, set by the user
2. **Package-level snap options** (`qwen36 set --package key=value`) — defaults set during installation
3. **Engine YAML configurations** (`engines/*/engine.yaml` → `configurations:` block) — declarative defaults bundled with the engine definition

### How Precedence Works

- The install hook sets package-level defaults for `http.port`, `http.host`, and `verbose`.
- When an engine is selected via `use-engine`, its `configurations:` block values (e.g., `server`, `model`, `multimodel-projector`, `http.base-path`, `gpu-layers`) are written as snap options.
- Users can override any value with `qwen36 set key=value` (without `--package`).
- All reads via `qwen36 get key` return the effective value (user override > package default).

## Configuration Keys

| Key | Type | Default | Set By | Description |
|-----|------|---------|--------|-------------|
| `http.port` | integer | `8326` | install hook | Server TCP port |
| `http.host` | string | `127.0.0.1` | install hook | Server bind address |
| `http.base-path` | string | `v1` | engine config | API URL base path |
| `model-name` | string | *(empty)* | engine/user | Model identifier for API responses |
| `verbose` | boolean | `false` | install hook | Enable verbose llama-server logging |
| `server` | string | *(engine-dependent)* | use-engine | Server component name (`llamacpp`, `llamacpp-cuda`) |
| `model` | string | *(engine-dependent)* | use-engine | Model component name |
| `multimodel-projector` | string | *(engine-dependent)* | use-engine | Multimodal projector component name |
| `gpu-layers` | integer | *(unset for CPU)* | use-engine (CUDA) | GPU layer offload count |

## No Environment Variable Mapping

There is no direct env-var-to-config mapping for the CLI itself. However, shell scripts set environment variables for downstream tools:
- `OPENAI_BASE_URL` — constructed from `http.port` + `http.base-path` (used by go-chat-client)
- `MODEL_NAME` — from `model-name` (used by go-chat-client)
- `LD_LIBRARY_PATH` — extended for engine shared libraries
- `OCL_ICD_VENDORS` — for OpenCL/Intel GPU detection

## No Config File

There is no file-based configuration (e.g., `~/.config/qwen36/config.yaml`). All configuration is stored in snap's internal option system, accessible only via `snapctl get/set` or the `qwen36 get/set` wrapper.

## Surprising Behaviors

1. **Engine selection writes multiple keys**: `use-engine` is not just an engine toggle — it writes `server`, `model`, `multimodel-projector`, `http.base-path`, and optionally `gpu-layers` from the engine's YAML configurations block.
2. **`model-name` may be unset**: Scripts handle this gracefully with `2>/dev/null || true`, but `get model-name` may produce an error or empty output.
3. **No `unset` command**: There's no observed way to reset a key to its default. The `set` command requires a value.
