# Architecture

## Tech Stack

| Layer | Technology |
|-------|-----------|
| CLI binary | Go (private submodule, compiled via `go` plugin) |
| Shell scripts | Bash (chat.sh, server.sh, wait-for-server.sh, check-server-llamacpp.sh, completion.bash) |
| Inference engine | llama.cpp (C/C++, cmake-built, CPU and CUDA variants) |
| Chat client | go-chat-client v1.0.0-beta.1 (external Go binary) |
| Configuration | snapd snapctl get/set (snap configuration system) |
| Packaging | Snapcraft (core24 base, strict confinement) |
| Model format | GGUF (Qwen3.6-35B-A3B-UD-Q4_K_XL, ~22GB) |
| Hardware detection | lspci, clinfo, lscpu (via pciutils, clinfo, nvidia-utils-580 stage-packages) |
| Service management | snapd daemon (systemd unit) |
| Dependencies | wget, jq, yq (runtime); OpenCL ICD (GPU detection) |

## Architecture Style

**Primary: Client-Server CLI**

The qwen36 snap follows a client-server architecture where:

- The **server** (`qwen36.server`) is a long-running daemon process that loads the LLM into memory and exposes an OpenAI-compatible HTTP API (llama-server).
- The **CLI client** (`qwen36` binary + shell wrappers) sends requests to the local server over HTTP (localhost:port/v1/).
- The `chat` command launches `go-chat-client` which connects to the server's completions API.
- Configuration is mediated through snapd's `snapctl get/set` mechanism, acting as a shared configuration store between daemon and CLI.

**Secondary: Layered CLI Application**

Within the CLI itself, there is a layered structure:

1. **Go binary layer** — the `qwen36` binary handles `use-engine`, `show-engine`, `get`, `set`, and `completion` commands with hardware detection and engine selection logic.
2. **Shell script layer** — wraps the Go binary for service orchestration (`server.sh` waits for components, selects engine, execs the server binary; `chat.sh` reads config and launches go-chat-client).
3. **Engine layer** — per-engine directories (`engines/cpu/`, `engines/cuda/`) contain `engine.yaml` (declarations), `server` (launch script), `common.sh` (component resolution), and `check-server` (health check).
4. **Component layer** — snap components provide hot-swappable model weights, multimodal projectors, and engine binaries.

## Key Architectural Characteristics

- **Component-based deployment**: Model, mmproj, and engine are independent snap components that can be installed/updated separately.
- **Engine abstraction**: Engine selection is decoupled from the server launch; `engine.yaml` declares hardware requirements and component dependencies.
- **Configuration via snapctl**: All runtime config (port, host, verbose, gpu-layers, engine name) is stored in snapd's configuration system, providing atomic get/set and persistence across restarts.
- **Auto-detection on install**: The `install` hook probes hardware and auto-selects the best engine via `use-engine --auto`.
