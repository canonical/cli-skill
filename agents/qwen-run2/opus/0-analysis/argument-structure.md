# Argument Structure

## Common Argument Patterns

The qwen36 CLI uses a small, consistent set of argument patterns:

1. **Subcommand dispatch**: The first positional argument is always a subcommand (`chat`, `use-engine`, `show-engine`, `get`, `set`, `completion`).
2. **Key-value assignment via `=`**: The `set` command uses `<key>=<value>` syntax for configuration writes (e.g., `http.port=8326`).
3. **Bare key lookup**: The `get` command takes a single positional key argument (e.g., `http.port`).
4. **Engine name as positional**: `use-engine` accepts an optional engine name (`cpu`, `cuda`) or a flag (`--auto`).
5. **Shell name as positional**: `completion` takes a shell name (`bash`) as its argument.
6. **No global flags**: There are no global flags (like `--verbose` or `--help`) documented at the top level, though the Go CLI likely supports `--help` implicitly.

## Complete Argument Map

### qwen36 chat

| Argument | Position | Required | Type | Default | Accepted Values | Aliases | Env Var |
|----------|----------|----------|------|---------|-----------------|---------|---------|
| *(none)* | — | — | — | — | — | — | — |

The `chat` command takes no arguments. All configuration is read from snap config keys internally.

### qwen36 use-engine

| Argument | Position | Required | Type | Default | Accepted Values | Aliases | Env Var |
|----------|----------|----------|------|---------|-----------------|---------|---------|
| `engine` | 1 | No (if `--auto` used) | string | — | `cpu`, `cuda` | — | — |
| `--auto` | flag | No | boolean | false | — | — | — |
| `--assume-yes` | flag | No | boolean | false | — | `-y` (likely) | — |

Notes:
- Either a positional engine name OR `--auto` must be provided.
- `--assume-yes` suppresses interactive confirmation prompts (used in install hook).
- When `--auto` is specified, the CLI evaluates all available engine.yaml files against detected hardware.

### qwen36 show-engine

| Argument | Position | Required | Type | Default | Accepted Values | Aliases | Env Var |
|----------|----------|----------|------|---------|-----------------|---------|---------|
| *(none)* | — | — | — | — | — | — | — |

No arguments. Outputs the current engine's YAML configuration to stdout.

### qwen36 get

| Argument | Position | Required | Type | Default | Accepted Values | Aliases | Env Var |
|----------|----------|----------|------|---------|-----------------|---------|---------|
| `key` | 1 | Yes | string | — | Any valid snap config key | — | — |

Known configuration keys:
- `http.port` (default: `8326`)
- `http.host` (default: `127.0.0.1`)
- `http.base-path` (default: `v1`)
- `verbose` (default: `false`)
- `model-name` (optional, may not be set)
- `server` (component name, e.g., `llamacpp`)
- `model` (component name, e.g., `model-qwen36-35b-a3b-ud-q4-k-xl`)
- `multimodel-projector` (component name, e.g., `mmproj-qwen36-35b-a3b-f16`)
- `gpu-layers` (integer, CUDA only, e.g., `99`)

### qwen36 set

| Argument | Position | Required | Type | Default | Accepted Values | Aliases | Env Var |
|----------|----------|----------|------|---------|-----------------|---------|---------|
| `key=value` | 1 | Yes | string | — | `<key>=<value>` | — | — |
| `--package` | flag | No | boolean | false | — | — | — |

Notes:
- The `--package` flag marks values as package-level defaults (used in install hook, not user-modifiable via `snap set`).
- Multiple key=value pairs may be accepted in a single invocation (based on `snapctl set` behavior).

### qwen36 completion

| Argument | Position | Required | Type | Default | Accepted Values | Aliases | Env Var |
|----------|----------|----------|------|---------|-----------------|---------|---------|
| `shell` | 1 | Yes | string | — | `bash` | — | — |

Only `bash` is currently supported.

## Special Arguments

### Engine YAML as implicit argument source

The `use-engine` command does not take explicit flags for the configurations it writes. Instead, it reads the selected engine's `engine.yaml` file and applies all `configurations:` entries as snap config keys. This means the "arguments" to engine selection are defined declaratively in YAML rather than on the command line.

### The `--package` flag on `set`

The `--package` flag is structural because it changes the persistence scope of the configuration value. Package-level settings survive snap refreshes and are not overridable via `snap set`. This creates a two-tier configuration model that is invisible to the user unless they inspect the install hook.

### Daemon pass-through arguments

The `qwen36.server` daemon's launch script (`server.sh`) passes `"$@"` to the engine server script, which in turn passes `"$@"` to `llama-server`. This means arbitrary llama-server flags can theoretically be passed to the daemon, but there is no documented interface for this and no snap app declaration exposes these arguments to users.
