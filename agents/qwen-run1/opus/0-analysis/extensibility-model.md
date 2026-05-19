# Extensibility Model

## Adding New Engines

The primary extension point is the engine system. New engines are added by:

1. **Create an engine directory** under `engines/<engine-name>/` containing:
   - `engine.yaml` — declarative specification with hardware requirements, component list, and default configurations
   - `server` — executable script that launches the inference server
   - `common.sh` — shared setup (component path resolution, validation, LD_LIBRARY_PATH)
   - `check-server` — health check script (optional, for readiness probes)

2. **Add a snap component** for the new inference backend (e.g., a new build of llama.cpp with different backends, or an entirely different inference engine).

3. **Register in snapcraft.yaml** — add the component definition under `components:` section.

### Engine Discovery

Engines are discovered by the CLI at runtime by scanning `$SNAP/engines/*/engine.yaml`. No code changes to the CLI binary are needed to add a new engine — the engine name from the YAML becomes a valid argument to `use-engine`.

### Engine YAML Schema

```yaml
name: <engine-name>          # Used as the identifier
description: <human text>    # Shown during engine selection
vendor: <string>
grade: stable|devel
devices:                     # Hardware requirements
  anyof|allof:
    - type: cpu|gpu
      architecture: amd64|arm64
      flags: [...]           # CPU instruction set requirements
      vendor-id: <hex>       # GPU vendor
      vram: <size>           # GPU memory requirement
memory: <size>               # System RAM requirement
disk-space: <size>           # Storage requirement
components:                  # Snap components this engine needs
  - <component-name>
configurations:              # Config values to set when this engine is selected
  server: <component>
  model: <component>
  multimodel-projector: <component>
  http.base-path: <path>
  gpu-layers: <int>          # Optional, GPU-specific
```

## Adding New Commands

Since the CLI source is a private Go submodule, the command registration mechanism is not observable. Based on the Go plugin and typical Go CLI frameworks (cobra, urfavecli), new commands would be added in the Go source code and require rebuilding the binary.

**No plugin system** is exposed for external command registration.

## Adding New Configuration Keys

New configuration keys can be introduced by:
1. Using them in engine YAML `configurations:` blocks (automatically set during engine selection)
2. Referencing them in server/app scripts via `qwen36 get <key>`
3. No schema validation is observed — the snap option system accepts arbitrary keys

## Adding New Models

New models are added as snap components:
1. Create a component directory under `components/<model-name>/` with an `init` script that exports `MODEL_FILE`
2. Register the component in `snapcraft.yaml`
3. Reference the component name in the engine YAML's `configurations.model` field

## Extension Boundaries

| Extension | Requires Code Change | Requires Snap Rebuild |
|-----------|---------------------|----------------------|
| New engine | No (add files only) | Yes (new component in snapcraft.yaml) |
| New model | No (add files only) | Yes (new component in snapcraft.yaml) |
| New CLI command | Yes (Go source) | Yes |
| New config key | No | No (if consumed by scripts only) |
| New output format | Yes (Go source) | Yes |

## Hooks and Middleware

- **Snap hooks** (`install`, `post-refresh`) provide lifecycle extension points managed by snapd
- **No pre/post command hooks** in the CLI itself
- **No middleware pattern** observed in the shell orchestration layer
