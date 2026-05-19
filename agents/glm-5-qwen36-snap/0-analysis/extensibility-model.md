# qwen36 Extensibility Model

## Summary

This snap is extensible at the packaging and engine-definition level, but not at the CLI parser level from the public repository because the `cli/` submodule is private and empty here.

## Observable Extension Boundaries

### 1. Engine Extension Boundary

The cleanest extension point is the `engines/` directory. Each engine currently provides:

- `engine.yaml` - Engine manifest with metadata, requirements, and defaults
- `common.sh` - Shared initialization script
- `server` - Server launcher script
- `check-server` - Health check script

The manifest supplies:

- Engine name
- Descriptive metadata (description, vendor, grade)
- Hardware requirements (devices, memory, disk-space)
- Required snap components
- Default configuration values

**To add a new engine, you need:**

1. A new `engines/<name>/engine.yaml` manifest
2. Launcher and health-check scripts under `engines/<name>/`
3. Matching component definitions in `snap/snapcraft.yaml`
4. CLI support so `use-engine` can discover or accept the new engine

### 2. Component Extension Boundary

Snap components already isolate:

- Server binaries (`llamacpp`, `llamacpp-cuda`)
- Model weights (`model-qwen36-35b-a3b-ud-q4-k-xl`)
- Multimodal projector weights (`mmproj-qwen36-35b-a3b-f16`)

That makes it straightforward to swap artifacts without rewriting launcher logic, as long as the CLI and engine manifests agree on component names.

### 3. Configuration Key Boundary

The shell scripts are decoupled from storage details by using `qwen36 get`. New engine or runtime features can be exposed by:

- Adding a new config key
- Teaching the CLI to persist and read it
- Consuming it in shell wrappers or engine launchers

The current example is `gpu-layers`, which only matters for the CUDA path.

## Non-Observable Extension Boundaries

These likely exist but cannot be audited:

- Command registration in the Go CLI
- Shared middleware for prompting, validation, and output formatting
- Completion generation internals
- Config backend abstraction
- Plugin or extension system

## Discovery Model

The public repository suggests that runtime discovery happens through selected engine state rather than directory scanning alone:

- `use-engine` chooses an engine
- `show-engine` returns the selected engine description
- `server.sh` trusts `show-engine` rather than enumerating `engines/*`

That means adding a new engine is not just a file-system operation. The CLI must know how to surface it.

## Extension Patterns

| Extension Type | Mechanism | Documentation | Discoverability |
|----------------|-----------|---------------|-----------------|
| New engine | Add `engines/<name>/engine.yaml` and scripts | Partially documented via existing engine examples | CLI must be updated to recognize new engine |
| New component | Add to `snap/snapcraft.yaml` components section | Documented in snapcraft docs | Requires engine manifest update |
| New config key | Add key to CLI config store | Undocumented | Only discovered via code inspection |
| New command | Modify private CLI binary | Not possible from public repo | N/A |
| New shell target for completion | Modify private CLI binary | Not possible from public repo | N/A |

## Constraints On Adding New Commands

Without the private `cli/` source, the project cannot publicly:

- Add new top-level commands
- Adjust parser rules
- Improve help text
- Expose new completion targets
- Add new flags to existing commands

So the current repo is **packaging-extensible but CLI-surface-constrained**.

## Engine Manifest Schema

Based on observed engine manifests, the schema is:

```yaml
name: <string>                    # Engine identifier
description: <string>             # Human-readable description
vendor: <string>                  # Vendor name
grade: <stable|devel>             # Stability grade
devices:                          # Hardware requirements
  anyof:                          # OR condition
    - type: <cpu|gpu>
      architecture: <string>      # e.g., amd64, arm64
      flags: [<string>, ...]      # Required CPU flags
      features: [<string>, ...]   # Required CPU features
      vendor-id: <string>         # GPU vendor ID (hex)
      vram: <string>              # Required VRAM (e.g., "24G")
memory: <string>                  # Required RAM (e.g., "32G")
disk-space: <string>              # Required disk (e.g., "25G")
components:                       # Required snap components
  - <component-name>
  - <component-name>
configurations:                   # Engine-specific config defaults
  server: <component-name>
  model: <component-name>
  multimodel-projector: <component-name>
  http.base-path: <string>
  gpu-layers: <integer>           # CUDA only
```

## Recommendation

If this project expects public contributors to add engines or config keys, it should publish at least:

- The engine discovery contract (how CLI discovers available engines)
- The shape of `show-engine` YAML
- The config precedence model
- Minimal help output or command grammar reference
- A template for adding new engines
