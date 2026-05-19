# Extensibility Model

## Adding New Engines

The primary extension point in the qwen36 snap is adding new inference engines. The process is:

### Registration Path

1. **Create an engine directory**: `engines/<engine-name>/` containing:
   - `engine.yaml` — declarative metadata: name, description, hardware requirements (`devices`, `memory`, `disk-space`), required snap components, and configuration key-value pairs.
   - `server` — bash launch script that sources `common.sh` and execs the inference binary.
   - `common.sh` — shared component resolution logic (resolves model/mmproj/server component paths).
   - `check-server` — health check script (exit 0 = healthy, 1 = starting, 2 = failed).

2. **Create a snap component**: Define a new component in `snapcraft.yaml` under `components:` for the engine binary (e.g., `llamacpp-cuda`). The component is built separately and staged into `(component/<name>)`.

3. **Declare in snapcraft.yaml parts**: Add a build part for the engine binary.

### Command Discovery

- The Go CLI binary discovers engines by scanning `$SNAP/engines/*/engine.yaml` at runtime.
- Engine names are derived from the `name:` field in each `engine.yaml`.
- `use-engine` lists available engines and validates the user's choice against detected hardware.
- Tab completion (via `completion bash`) includes discovered engine names.

### Hardware Matching

The `engine.yaml` schema defines hardware requirements:
- `devices.anyof` / `devices.allof` — lists of device requirements (CPU architecture/flags, GPU vendor-id/vram).
- `memory` — minimum system RAM.
- `disk-space` — minimum disk space.

The Go binary evaluates these against `lscpu`, `lspci`, `clinfo`, and memory info to determine compatibility.

## Adding New Commands

### Go CLI binary
New commands are added to the private Go CLI submodule (`cli/`). Since this submodule is not checked out, the exact registration mechanism is unknown, but based on the snap structure:
- The binary is built from `cli/` source and installed as `bin/qwen36`.
- New subcommands are Go functions registered with whatever CLI framework the binary uses (likely cobra or similar).

### Shell script commands
Shell scripts in `apps/` are registered as snap apps in `snapcraft.yaml` under `apps:`. Adding a new script-based command requires:
1. Add the script to `apps/`.
2. Register it in `snapcraft.yaml` `apps:` section.
3. Declare required plugs and environment.

## Adding New Models

Models are added as snap components:
1. Create a component source directory under `components/<model-name>/` with:
   - `init` — a bash script that exports `MODEL_FILE` (the path to the GGUF file).
   - `README.md` — model documentation.
2. Declare the component in `snapcraft.yaml` under `components:` and `parts:`.
3. Reference the model component name in the engine.yaml `configurations.model` field.

## Extension Boundaries

| What | How | Boundary |
|------|-----|----------|
| New engine | Directory + component + snapcraft.yaml | Requires snap rebuild |
| New model | Component + init script + snapcraft.yaml | Requires snap rebuild |
| New CLI command | Go code in cli/ submodule | Requires snap rebuild |
| New shell command | Script in apps/ + snapcraft.yaml | Requires snap rebuild |
| New config key | Go CLI code + script usage | Requires snap rebuild |
| Runtime engine switch | `use-engine` command | User action, no rebuild |
| Runtime config change | `set` command | User action, no rebuild |

## Plugin System

There is no plugin system. All extensions require modifying snap source and rebuilding. The component system provides modularity for deployment (independent install/update of engines, models, and projectors) but not for runtime extensibility.

## Middleware / Hooks

| Hook | Trigger | Purpose |
|------|---------|---------|
| `install` | First snap install | Set package defaults, auto-select engine |
| `post-refresh` | After snap refresh | Currently a no-op (just sets up logging) |

Snap hooks are the only lifecycle extension points. There are no pre/post command hooks, no middleware pipeline, and no event system within the CLI itself.
