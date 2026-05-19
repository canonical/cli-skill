# qwen36 CLI Architecture

## Summary

The `qwen36` command is a snap-packaged Go CLI that serves as the control plane for a local inference service running the Qwen3.6-35B-A3B Vision Language Model. The observable stack consists of:

- A private Go binary installed as `bin/qwen36`
- Snap packaging and lifecycle hooks (`snap/snapcraft.yaml`, `snap/hooks/`)
- Shell wrappers in `apps/` that read CLI-managed configuration and launch or probe the model server
- Engine manifests and launcher scripts in `engines/cpu/` and `engines/cuda/`
- A daemon app (`server`) that starts the selected inference engine (llama.cpp)
- Model and projector components installed as snap components

## Primary Architecture Style

**Primary style: Layered CLI application**

The layers are:

1. **Command layer**: The `qwen36` binary exposes user commands (`chat`, `use-engine`, `show-engine`, `get`, `set`, `completion bash`)
2. **Configuration layer**: Shell scripts resolve runtime values through `qwen36 get <key>`, backed by a persistent configuration store managed by the CLI
3. **Engine selection layer**: Engine manifests describe hardware requirements, required snap components, and engine-specific default configuration
4. **Runtime layer**: `apps/server.sh` resolves the selected engine and execs the engine-specific `server` script

## Secondary Architecture Style

**Secondary style: Client-server CLI**

The CLI itself is not the inference runtime. It configures and targets a local daemonized server exposed by the snap's `server` app. The chat client (`go-chat-client`) uses config values from the CLI to connect to `http://localhost:<port>/<base-path>`.

## Control Flow

### Engine Startup Path

1. Snap install hook seeds package defaults with `qwen36 set --package` for `http.port`, `http.host`, and `verbose`
2. The install hook runs `qwen36 use-engine --auto --assume-yes` when hardware inspection is available
3. `apps/server.sh` calls `qwen36 show-engine`, parses YAML fields `.components[]` and `.name`, waits for required snap components, and execs `engines/<name>/server`
4. The selected engine server script reads resolved config via `qwen36 get ...` and starts `llama-server` with the correct model, projector, host, port, and optional GPU flags

### Chat Path

1. `qwen36 chat` is documented as the interactive entrypoint
2. The snap environment sets `CHAT=$SNAP/bin/chat.sh`, indicating the CLI delegates chat startup to that wrapper
3. `apps/chat.sh` resolves `http.port`, `http.base-path`, and optional `model-name` via `qwen36 get` and exports `OPENAI_BASE_URL` and `MODEL_NAME` before execing `go-chat-client`

## Observable Boundaries

### Inside the Private CLI

Not observable in this repository:

- Parser implementation
- Help text
- Validation rules for arguments and keys
- Actual storage backend for configuration
- Exit code mapping for CLI-native failures

### Outside the Private CLI

Observable and relied upon by packaging:

- `show-engine` returns YAML with at least `name` and `components`
- `get` returns scalar values for keys used by shell scripts
- `set` can write package-scoped defaults when `--package` is present
- `use-engine` can auto-detect hardware and can skip confirmation with `--assume-yes`
- `completion bash` emits a token list usable by `compgen -W`

## Architectural Strengths

- Clear separation between CLI control plane and model-serving data plane
- Engine-specific launch behavior is isolated in per-engine directories
- Snap components allow model, projector, and server binaries to be composed independently
- Runtime shell wrappers are thin and mostly declarative
- Configuration is centralized through the `get`/`set` interface

## Architectural Risks

- The CLI is a single opaque dependency—help quality, error contracts, and argument validation cannot be audited from source
- Scripts depend on output shape from `show-engine` and `get` but the contracts are undocumented
- Configuration layering appears important, but precedence rules are not documented
- Completion is delegated back into the CLI, so completion quality is only as good as the private binary's completion emitter
- No observable mechanism for users to discover supported engines from within the CLI
