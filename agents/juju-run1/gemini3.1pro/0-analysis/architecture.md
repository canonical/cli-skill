# Architecture

## Tech Stack

- **Language**: Go 1.24 (CLI binary)
- **Build system**: Snapcraft (snap packaging)
- **Inference backend**: llama.cpp (C/C++, built via CMake)
- **Chat client**: go-chat-client (Go, external dependency)
- **Configuration**: Snap options via `snapctl` (managed by the Go CLI)
- **Shell scripts**: Bash (server orchestration, engine management, health checks)
- **Target platform**: Ubuntu (snap confinement: strict, base: core24)

## Architecture Style

**Primary: Client-server CLI**

The `qwen36` binary acts as a client that configures and manages a long-running `llama-server` daemon process. The server exposes an OpenAI-compatible HTTP API (`/v1/completions`, `/v1/models`). The `chat` command connects to this server via the API. Configuration is stored in snap options and engine YAML files.

**Secondary: Layered CLI application**

The CLI is layered into:

1. **CLI layer** (Go binary) — command parsing, configuration get/set, engine selection, shell completion
2. **Orchestration layer** (Bash scripts) — server startup, component readiness checks, health probes
3. **Engine layer** (engine YAML + server scripts) — hardware-specific inference configuration
4. **Runtime layer** (llama.cpp) — actual model inference via llama-server

## Component Diagram

```
┌─────────────────────────────────────────────────────┐
│  User                                                │
│    qwen36 chat / qwen36 set / qwen36 use-engine     │
└─────────────┬───────────────────────────────────────┘
              │
┌─────────────▼───────────────────────────────────────┐
│  CLI Binary (Go)                                     │
│    Commands: chat, get, set, show-engine,            │
│              use-engine, completion                   │
│    Config: snap options (snapctl get/set)            │
└─────────────┬───────────────────────────────────────┘
              │
┌─────────────▼───────────────────────────────────────┐
│  Shell Scripts (Bash)                                │
│    server.sh, chat.sh, wait-for-server.sh,          │
│    check-server-llamacpp.sh                          │
└─────────────┬───────────────────────────────────────┘
              │
┌─────────────▼───────────────────────────────────────┐
│  Engine (llama.cpp / llama.cpp-cuda)                 │
│    llama-server process (daemon)                     │
│    OpenAI-compatible API on localhost:8326           │
└─────────────────────────────────────────────────────┘
```

## Key Architectural Decisions

1. **Snap components for modularity**: Model weights, inference engines, and multimodal projectors are separate snap components that can be installed/updated independently.
2. **Engine abstraction**: Hardware-specific configurations (CPU vs CUDA) are encapsulated in engine YAML files with declarative hardware requirements.
3. **Auto-detection at install**: The install hook probes hardware capabilities and selects the best engine automatically.
4. **Configuration via snap options**: All runtime config is stored in snap's built-in option system, accessible via `snapctl get/set` (wrapped by the CLI's `get`/`set` commands).
