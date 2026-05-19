# Architecture

## Summary

The qwen36 snap exposes a thin Go CLI front end over a snap-packaged inference runtime. The dominant architecture style is a client-server CLI: user-facing commands mutate snap configuration, query snap state, or launch an interactive client, while the long-running `qwen36.server` daemon starts the selected inference engine and serves an OpenAI-compatible HTTP API locally. A secondary classification is layered CLI application: snap hooks seed configuration, the CLI mediates access to snapd config, shell wrappers translate config into process launches, and engine manifests describe hardware-specific runtime variants.

## Tech stack

| Layer | Technology | Evidence |
|---|---|---|
| Packaging | Snap, snapcraft, components | `snap/snapcraft.yaml` defines apps, hooks, parts, and components |
| CLI front end | Go binary at `bin/qwen36` | `snap/snapcraft.yaml` `cli` part builds `bin/qwen36` from a private Go submodule |
| Interactive chat client | Go | `go-chat-client` part builds and `apps/chat.sh` execs `go-chat-client` |
| Control scripts | POSIX shell / bash | `apps/*.sh`, `engines/*/server`, `snap/hooks/*` |
| Runtime config | snapd config via `snapctl get/set` wrapped by `qwen36 get/set` | hooks and scripts consistently read and write through the CLI |
| Engine metadata | YAML | `engines/cpu/engine.yaml`, `engines/cuda/engine.yaml` |
| Inference server | `llama-server` from `llama.cpp` | engine server scripts exec `$server/usr/local/bin/llama-server` |
| Machine-readable plumbing | `yq`, `jq`, `wget`, `nc` | used by hooks and health checks |

## Primary architecture style: client-server CLI

The command surface splits cleanly into three roles:

1. Control-plane commands: `use-engine`, `show-engine`, `get`, and `set` read or mutate persisted snap state.
2. Data-plane client command: `chat` resolves the current local API endpoint and execs `go-chat-client`.
3. Service process: `qwen36.server` runs the selected engine-specific launcher and exposes the HTTP inference API.

That shape is client-server rather than monolithic because the user-facing CLI is not the inference runtime itself. It configures and targets a separately running daemon, and even the chat flow relies on the daemon being healthy before proceeding.

## Secondary architecture style: layered CLI application

A second useful classification is layered CLI application:

1. Snap packaging layer: `snapcraft.yaml` declares the CLI app, server daemon, hooks, and installable components.
2. Configuration layer: install hook writes package defaults and engine selection into snap config.
3. Command layer: the Go CLI exposes verbs like `get`, `set`, `show-engine`, and `use-engine`.
4. Script orchestration layer: shell wrappers derive environment variables, wait for readiness, and translate YAML/config into process flags.
5. Runtime layer: `llama-server` performs inference against the selected model and projector artifacts.

## Control flow

### Installation and bootstrap

1. The install hook sets generic defaults with `qwen36 set --package`:
   - `http.port=8326`
   - `http.host=127.0.0.1`
   - `verbose=false`
2. If hardware inspection is available, the install hook runs `qwen36 use-engine --auto --assume-yes`.
3. The selected engine determines which snap components and configuration keys are required later by the daemon.

### Daemon startup

1. `qwen36.server` runs `apps/server.sh`.
2. `server.sh` calls `qwen36 show-engine`, parses `.components[]`, and waits until all required components are present under `$SNAP_COMPONENTS`.
3. Once components are available, `server.sh` parses `.name` from the same YAML and execs `engines/<name>/server`.
4. The engine-specific launcher sources `common.sh`, resolves component paths through `qwen36 get`, sources model/mmproj init scripts, and finally execs `llama-server` with engine-specific flags.

### Interactive chat

1. `qwen36 chat` ultimately relies on `apps/chat.sh`.
2. `chat.sh` reads `http.port`, `http.base-path`, and optional `model-name` through `qwen36 get`.
3. It exports `OPENAI_BASE_URL` and `MODEL_NAME`, then execs `go-chat-client`.
4. `apps/wait-for-server.sh` provides the readiness gate used in the chat chain, polling the engine health check for up to 60 seconds.

## Architecture diagram

```text
User
  |
  v
qwen36 CLI (Go)
  |\
  | \__ chat -> apps/chat.sh -> go-chat-client -> localhost OpenAI-compatible API
  |
  |__ get/set/use-engine/show-engine -> snap config + engine metadata
                                   |
                                   v
                            apps/server.sh
                                   |
                                   v
                        engines/<cpu|cuda>/server
                                   |
                                   v
                              llama-server
                                   |
                                   v
                    model + mmproj + engine components
```

## Boundaries and responsibilities

| Boundary | Responsibility |
|---|---|
| Go CLI | Parse user commands, expose stable verbs, wrap snap config and engine selection logic |
| Snap hooks | Seed defaults and perform automatic engine selection during installation |
| Shell wrappers | Bridge persisted config to executable runtime behavior |
| Engine manifests | Describe hardware requirements, required components, and engine-specific config values |
| Engine launchers | Translate config into concrete `llama-server` arguments |
| Runtime components | Supply the engine binary, model file, and projector file |

## Architectural strengths

- Packaging, configuration, and runtime concerns are separated cleanly.
- Engine selection is data-driven through per-engine YAML manifests rather than hard-coded in one large script.
- The daemon startup path treats component availability as a first-class prerequisite and waits explicitly.
- `show-engine` provides a machine-readable bridge between control-plane selection and runtime launch.

## Architectural constraints and gaps

- The Go CLI submodule is not present in this checkout, so command parsing and validation logic cannot be directly audited here.
- The shell/runtime path depends on `show-engine` YAML fields `name` and `components[]`; that output contract is load-bearing but undocumented.
- Configuration is centralized in snapd rather than a tool-local config file, which is operationally simple but reduces portability outside the snap environment.
- Chat and daemon behavior are tightly coupled to the local server API shape and selected engine config.

## Assessment

This is a pragmatic snap-native control-plane architecture: a small CLI and hook layer persist state, shell wrappers operationalize that state, and the actual inference work happens in a daemonized engine process. The command set is small, but the architecture already has clear extension seams in engine manifests, snap components, and wrapper scripts.