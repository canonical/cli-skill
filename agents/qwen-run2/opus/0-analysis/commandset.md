# Command Set

## Overview

The qwen36 snap exposes commands through two mechanisms:
1. **Go CLI binary** (`qwen36`) — provides subcommands for engine management, configuration, and shell completions.
2. **Snap app declarations** — `qwen36` (default app) and `qwen36.server` (daemon).

The chat functionality is implemented as a shell script (`chat.sh`) invoked by the Go binary's `chat` subcommand.

## Commands

### qwen36 chat

| Field | Value |
|-------|-------|
| **Name** | `chat` |
| **Description** | Launch an interactive chat session with the Qwen3.6 model |
| **How it works** | The Go binary invokes `$SNAP/bin/chat.sh`. The script calls `qwen36 get http.port`, `qwen36 get model-name`, and `qwen36 get http.base-path` to resolve the server endpoint. It sets `OPENAI_BASE_URL` and `MODEL_NAME` environment variables, then execs `$SNAP/bin/go-chat-client`. Before connecting, `wait-for-server.sh` polls the server health check (via the engine's `check-server` script) with a 60-second timeout. |

### qwen36 use-engine

| Field | Value |
|-------|-------|
| **Name** | `use-engine` |
| **Description** | Select the inference engine to use (cpu, cuda, or auto-detect) |
| **How it works** | The Go binary reads available `engines/*/engine.yaml` files, evaluates hardware requirements (CPU flags, GPU vendor/VRAM, RAM, disk) against the current system, and writes the selected engine's configurations to snapctl. With `--auto`, it scores engines by capability and selects the best match. Writes `server`, `model`, `multimodel-projector`, `http.base-path`, and optionally `gpu-layers` to snap config. May restart the server daemon after selection. |

### qwen36 show-engine

| Field | Value |
|-------|-------|
| **Name** | `show-engine` |
| **Description** | Display the currently configured engine and its properties |
| **How it works** | The Go binary reads the current engine name from snap config, locates the corresponding `engine.yaml` file, and outputs its contents (YAML format) to stdout. Other scripts parse this output with `yq` to extract the engine name and component lists. |

### qwen36 get

| Field | Value |
|-------|-------|
| **Name** | `get` |
| **Description** | Read a snap configuration value |
| **How it works** | The Go binary wraps `snapctl get <key>` and prints the value to stdout. Used extensively by shell scripts (chat.sh, server scripts, common.sh) to retrieve runtime configuration such as `http.port`, `http.host`, `verbose`, `model-name`, `server`, `model`, `multimodel-projector`, `gpu-layers`, and `http.base-path`. |

### qwen36 set

| Field | Value |
|-------|-------|
| **Name** | `set` |
| **Description** | Write a snap configuration value |
| **How it works** | The Go binary wraps `snapctl set <key>=<value>`. Supports a `--package` flag (used in install hook) to set values that are not user-modifiable. Configuration keys follow a dot-separated hierarchy (e.g., `http.port`, `http.host`). |

### qwen36 completion

| Field | Value |
|-------|-------|
| **Name** | `completion` |
| **Description** | Generate shell completion scripts |
| **How it works** | The Go binary outputs completion words for the specified shell. Currently only `bash` is supported as a subcommand argument. The generated completions are used by `completion.bash` (registered via the `completer` field in snapcraft.yaml) to provide tab-completion for the `qwen36` command. |

### qwen36.server (daemon)

| Field | Value |
|-------|-------|
| **Name** | `qwen36.server` |
| **Description** | Long-running inference server daemon |
| **How it works** | Declared as a `daemon: simple` snap app. Runs `server.sh` which: (1) waits for all required snap components to be installed (polling every 10s, up to 3600s timeout), (2) reads the engine name via `qwen36 show-engine | yq .name`, (3) execs the engine-specific server script (`$SNAP/engines/$engine/server`). The engine server script sources `common.sh` to resolve component paths, then launches `llama-server` with model, port, host, context-size, and sampling parameters. |

## Command Hierarchy

```
qwen36
├── chat                    (interactive session)
├── use-engine              (engine selection)
│   ├── cpu                 (positional: select CPU engine)
│   ├── cuda                (positional: select CUDA engine)
│   └── --auto              (flag: auto-detect best engine)
├── show-engine             (display current engine)
├── get <key>               (read config)
├── set <key>=<value>       (write config)
└── completion
    └── bash                (generate bash completions)

qwen36.server               (daemon, not a user CLI command)
```
