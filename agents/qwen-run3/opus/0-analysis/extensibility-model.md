# Extensibility Model

## Adding New Engines

The primary extension point is adding new inference engines:

1. **Create engine directory**: `engines/<name>/` with:
   - `engine.yaml`: Manifest declaring hardware requirements, components, and default config
   - `server`: Shell script to start the engine's server process
   - `check-server`: Shell script for health checking
   - `common.sh`: Shared environment setup (model/component paths)

2. **Add snap component**: In `snapcraft.yaml`, declare a new component part that builds the engine binary.

3. **No code changes required in the CLI**: The CLI discovers engines by scanning `$SNAP/engines/*/engine.yaml`. New engines are automatically listed, scored, and selectable.

## Engine Manifest Schema

```yaml
name: <engine-name>
description: <human description>
vendor: <vendor>
grade: stable|devel
devices:
  anyof|allof:
    - type: cpu|gpu
      architecture: amd64|arm64
      flags: [...]       # CPU flags
      vendor-id: 0x...   # GPU vendor
      vram: <size>       # GPU VRAM requirement
memory: <size>           # System RAM requirement
disk-space: <size>       # Disk space requirement
components:
  - <component-name>     # Required snap components
configurations:
  <key>: <value>         # Config applied when engine is activated
```

## Adding New CLI Commands

New commands are added by:

1. Creating a new file in `cli/cmd/cli/commands/<command>.go`
2. Implementing a function returning `*cobra.Command`
3. Registering it in `main.go` via `addCommandGroup()` or `addCommands()`

The CLI uses cobra's standard registration pattern — no plugin discovery, no dynamic loading.

## Extension Boundaries

- **No plugin system**: The CLI does not support external plugins, extensions, or command discovery from PATH.
- **No middleware/hooks**: There is no pre/post command hook mechanism beyond cobra's `PersistentPreRunE`.
- **No API versioning**: The CLI config keys and output schemas are not versioned; changes are breaking.
- **Snap components as extension**: The snap component system provides the primary extensibility — new engines and models can be installed independently without rebuilding the snap.

## Command Discovery

The CLI uses cobra's built-in features:
- `--help` on any command
- Shell completion via `completion bash` (bash completion script in `apps/completion.bash`)
- Hidden commands are excluded from help but accessible if known
- Command groups are displayed in help with headers ("Basic Commands:", "Configuration Commands:", "Management Commands:")

## Debug Extension Point

The hidden `debug` command group provides a namespace for developer/testing commands that can be added without polluting the user-facing command set.
