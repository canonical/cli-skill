# Architecture

## Tech Stack Summary

The `qwen36` CLI is built as a **Go application** using the [Cobra](https://github.com/spf13/cobra) command framework. It is packaged as a **snap** (via `snapcraft`) and deployed on `core24` (Ubuntu 24.04 base). The CLI binary is compiled from the `inference-snaps-cli` repository and shipped under `bin/qwen36`. The snap also bundles:

- **llama.cpp** (build tag `b8157`) as the inference engine, with CPU and CUDA variants.
- A **go-chat-client** (`v1.0.0-beta.1`) for OpenAI-compatible chat interactions.
- System tools: `wget`, `jq`, `yq` for model initialization and hook scripts.
- PCI and GPU detection tools: `pciutils`, `nvidia-utils-580`, `clinfo`.

## Architecture Style

### Primary: Layered CLI Application

The CLI follows a **layered architecture** with well-defined tiers:

| Layer | Package/Module | Responsibility |
|-------|---------------|----------------|
| **Command layer** | `cmd/cli/main.go`, `cmd/cli/commands/` | Cobra command definitions, argument parsing, user interaction |
| **Common/shared layer** | `cmd/cli/common/` | Shared logic: engine scoring, endpoints, services, prompts, spinners, suggestions |
| **Domain packages** | `pkg/engines/`, `pkg/selector/`, `pkg/hardware_info/` | Engine manifest loading, compatibility scoring, hardware detection |
| **Infrastructure layer** | `pkg/storage/`, `pkg/snap/`, `pkg/snap_store/` | Configuration persistence (snapctl), snap management, snap store queries |
| **External services** | `pkg/webui/` | Web UI static file server |

Data flows from the infrastructure/domain layers up through the common layer to the command layer, which renders output (YAML/JSON/table) and manages user interaction.

### Secondary: Plugin-based Architecture (Engine System)

The CLI has a **plugin-based engine selection system**. Engines are defined as YAML manifests in `$SNAP/engines/`. Each engine declares:
- Hardware requirements (CPU architecture, PCI devices, memory, disk space)
- Required snap components (e.g., `llamacpp`, `llamacpp-cuda`, model weights)
- Default configurations (`configurations` map)

At runtime, the CLI discovers all engine manifests, scores them against the host hardware, and lets the user select or auto-select the best compatible engine. New engines can be added by dropping new manifest files — no code changes needed. This is a **data-driven plugin model** rather than a binary plugin model.

## Key Design Decisions

- **Snap-native persistence**: Configuration and cache are stored via `snapctl set/get/unset` (snapd's configuration API), not direct file I/O. This ties the CLI closely to the snap ecosystem.
- **Single binary, multiple commands**: All functionality lives in one `qwen36` binary with a flat top-level command hierarchy (no deeply nested subcommands beyond `debug`).
- **Conditional feature enablement**: Chat and WebUI commands are toggled via the `ADDITIONAL_FEATURES` environment variable, allowing the same binary to serve different snap packages.
- **Machine-readable output**: Commands that display data (`status`, `show-machine`, `show-engine`, `list-engines`, `version`) support both human-readable (YAML/table) and machine-readable (JSON) formats via a `--format` flag.
- **Hardware-awareness as first-class feature**: The CLI performs deep hardware probing (CPU flags via `/proc/cpuinfo`, PCI devices via `lspci`, GPU memory via `clinfo`, disk via `statfs`) to determine engine compatibility.
