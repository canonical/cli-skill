# Architecture Analysis: inference-snaps-cli (qwen36)

## Tech Stack Summary

- **Language**: Go 1.26+
- **CLI Framework**: Cobra (spf13/cobra)
- **Build Tool**: snapcraft (CMake for llama.cpp, Go plugin for CLI)
- **Target Platform**: Ubuntu core24 snap, strict confinement
- **Key Libraries**:
  - `github.com/canonical/go-snapctl` — snapd control plane integration
  - `github.com/openai/openai-go/v3` — OpenAI-compatible chat client
  - `github.com/olekukonko/tablewriter` — terminal table rendering
  - `gopkg.in/yaml.v3` — YAML I/O
  - `github.com/fatih/color` — colored terminal output

## Architecture Style

**Primary style**: **Plugin-based / Component-driven CLI** (Microkernel command host variant)

The CLI acts as a command host that discovers and manages external engine manifests and snap components. The core binary contains the command framework, configuration logic, and chat client, but the actual inference engines are loaded dynamically from manifest files on disk (under `$SNAP/engines/`). Each engine declares its required snap components, environment variables, server settings, and filesystem layouts. The CLI orchestrates component installation, environment setup, and service lifecycle without embedding engine-specific logic.

**Secondary style**: **Client-server CLI** (for chat/webui commands)

Several top-level commands (`chat`, `webui`) are thin clients that connect to a local OpenAI-compatible HTTP server running inside the snap daemon. These commands validate service health, wait for endpoints, and stream completions but do not perform inference locally.

## Structural Overview

```
main.go (root cobra.Command)
├── Command Groups (visual only; Cobra groups)
│   ├── basic:      status, chat*, webui*
│   ├── config:     get, set, unset
│   └── engine:     list-engines, show-engine, use-engine
├── Ungrouped:      show-machine, prune-cache, version
├── Hidden:         run, serve-webui
└── Hidden group:   debug
    ├── validate-engines
    ├── select-engine
    ├── chat
    └── serve-webui
```

- `*`: Added conditionally based on `ADDITIONAL_FEATURES` env var.

## Key Packages

| Package | Responsibility |
|---|---|
| `cmd/cli/commands` | Cobra command definitions and user-facing RunE handlers |
| `cmd/cli/common` | Shared helpers: engine scoring, endpoints, prompts, service status |
| `pkg/engines` | Manifest loading and validation (plugin boundary) |
| `pkg/hardware_info` | Host capability introspection (CPU, memory, disk, PCI devices) |
| `pkg/selector` | Engine compatibility scoring and selection algorithm |
| `pkg/storage` | Configuration and cache persistence via `snapctl` |
| `pkg/snap` / `pkg/snap_store` | Snapd interaction wrappers (services, components, store metadata) |
| `pkg/webui` | Embedded HTTP server for Web UI static files |

## Platform Binding

The CLI is tightly bound to the snapd ecosystem:
- Uses `snapctl` as the only configuration backend (`pkg/storage/snapctl_storage.go`)
- Uses `snapctl.Services()`, `snapctl.InstallComponents()`, `snapctl.RemoveComponents()` for lifecycle
- Reads `SNAP`, `SNAP_INSTANCE_NAME`, `SNAP_REVISION`, `SNAP_COMPONENTS` for paths
- Requires root for all mutating operations (enforced by `utils.IsRootUser()`)

This design makes the CLI non-portable outside of a snap context; running standalone loses config persistence and service management.

## State Model

The CLI maintains two logical stores backed by `snapctl`:
1. **Cache** (`cache.` prefix): transient runtime state, currently only `active-engine`
2. **Config** (`config.` prefix): layered configuration with precedence
   - `package` (lowest) → `engine` → `user` (highest)

Configuration keys are dot-separated (e.g., `http.port`, `passthrough.environment.my-key`).

## Execution Model

- **Foreground**: chat, webui launch, run subcommand
- **Background daemon**: server startup handled by snapd systemd unit (`snapctl start`)
- **Hooks**: install/post-refresh hooks call the CLI itself to initialize defaults
