# Architecture

The `qwen36` CLI is built primarily with **Go**, accompanied by supporting wrappers in **Bash**. It serves as the user-facing entrypoint for the `qwen36` snap package.

**Architecture Style**: Layered CLI application
The CLI provides a high-level configuration and lifecycle management interface (setting ports, choosing inference engines), delegating the actual generative AI inference workload to bundled C/C++ microservices (namely `llama.cpp` using CPU or CUDA), and a separate Go-based `go-chat-client` for interactive shell chat. Communication between layers relies on local network ports (HTTP API).