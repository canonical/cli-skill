# Architecture

Short summary: The `qwen36` CLI is a Go-based wrapper (Cobra) provided as a snap app that orchestrates an inference snap composed of components and engines. The CLI is a thin client that delegates to snap-controlled services and local helper libraries.

Primary style: Layered CLI application (thin client + local libraries + service orchestration).

Secondary styles:
- Plugin-based / component-driven distribution (snap components for engines and models).
- Client-server interaction: commands interact with a local `server` daemon (App `server` uses a simple daemon).

Tech stack:
- Language: Go (Cobra CLI)
- Packaging: snapcraft (snap with components and apps)
- Engine integration: components (llama.cpp, cuda variant) staged as snap components
- Helpers: shell scripts in `apps/` for chat and server controls

Notes:
- The CLI relies heavily on snapd APIs (go-snapctl) to manage components and services. Configuration and state persist in a local storage implementation under `pkg/storage`.
