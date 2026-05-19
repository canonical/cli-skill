# Configuration Model

## Config Sources and Precedence

The `qwen36` CLI uses a layered configuration model backed by `snapctl` (the snap daemon's configuration subsystem). Precedence from highest to lowest:

1. **User config** (highest priority) — set by `qwen36 set <key=value>`
2. **Engine config** — set by `qwen36 set --engine <key=value>` during engine activation
3. **Package config** (lowest priority) — set by `qwen36 set --package <key=value>` during snap install hook

When a key is read via `qwen36 get <key>`, the highest-priority layer wins. When `qwen36 unset <key>` is called, only the user layer is cleared, revealing the engine or package default.

## Config Keys

Based on the install hook and engine manifests, the known configuration keys are:

| Key | Default (package) | Engine Override | Description |
|-----|-------------------|----------------|-------------|
| `http.port` | `8326` | — | Port for the inference server |
| `http.host` | `127.0.0.1` | — | Bind address for the inference server |
| `http.base-path` | — | `v1` (both engines) | Base path for the OpenAI API |
| `verbose` | `false` | — | Enable verbose server logging |
| `server` | — | `llamacpp` or `llamacpp-cuda` | Active server component name |
| `model` | — | `model-qwen36-35b-a3b-ud-q4-k-xl` | Active model component name |
| `multimodel-projector` | — | `mmproj-qwen36-35b-a3b-f16` | Active mmproj component name |
| `model-name` | — | (optional) | Model name for OpenAI API |
| `gpu-layers` | — | `99` (cuda only) | Number of layers to offload to GPU |

## Config Source: Engine Manifests

Each engine (`engines/<name>/engine.yaml`) has a `configurations` section that is applied when the engine is activated via `use-engine`. These become the engine-layer config values.

## Config Source: Flags

Flags like `--verbose` on the root command are runtime-only and do not persist. They override the stored `verbose` config for that invocation only.

## Config Source: Environment Variables

No explicit environment variable mapping exists in the CLI itself. However:
- `SNAP_COMPONENTS` — set by the snap environment, points to installed components
- `SNAP`, `SNAP_INSTANCE_NAME`, `SNAP_VERSION`, `SNAP_REVISION` — standard snap environment
- `LD_LIBRARY_PATH` — extended by engine scripts for shared library loading
- `NO_COLOR` — not explicitly handled (potential gap)

## Surprising Precedence Behavior

1. **Hidden layer flags**: The `--package` and `--engine` flags on `set` allow internal tools (snap hooks) to write to lower-priority layers. End users cannot easily discover that their `set` is writing to the "user" layer while the install hook wrote to "package" layer.

2. **Engine switch replaces engine config**: When switching engines via `use-engine`, the entire engine configuration layer is replaced with values from the new engine's manifest. Previous engine-layer values are lost without warning.

3. **No env var overrides**: Unlike many CLIs, there is no `QWEN36_HTTP_PORT` or similar env var that users can set to override config without modifying the snap state. All configuration is persisted in snapd.
