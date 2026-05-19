# Argument Structure

## Common Patterns

1. **Subcommand-first**: All commands use `qwen36 <command>` as the entry point with no global flags observed before the subcommand.
2. **Key=value positional for `set`**: The `set` command uses `key=value` as a positional argument rather than flags.
3. **Bare key positional for `get`**: The `get` command takes the configuration key as a bare positional argument.
4. **Positional engine name for `use-engine`**: Engine selection accepts engine name as a positional argument OR the `--auto` flag (mutually exclusive).
5. **Subcommand for `completion`**: The `completion` command takes a shell name (`bash`) as a positional argument.

## Complete Argument Map

### `qwen36 chat`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| *(none)* | — | — | — | No arguments; connects to the running server using stored config |

### `qwen36 use-engine`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<engine-name>` | positional | conditional | — | Engine to select (`cpu`, `cuda`). Required unless `--auto` is specified |
| `--auto` | flag (boolean) | conditional | false | Automatically detect and select the best engine for current hardware |
| `--assume-yes` | flag (boolean) | optional | false | Skip interactive confirmation prompt |

Notes:
- `<engine-name>` and `--auto` are mutually exclusive
- Available engine names are determined by YAML files in `$SNAP/engines/*/engine.yaml`

### `qwen36 show-engine`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| *(none)* | — | — | — | No arguments; outputs current engine YAML to stdout |

### `qwen36 get`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<key>` | positional | required | — | Configuration key to retrieve |

Known keys:
- `http.port` — Server listening port (default: `8326`)
- `http.host` — Server bind address (default: `127.0.0.1`)
- `http.base-path` — API base path (default: `v1`)
- `model-name` — Served model name (may be empty)
- `verbose` — Verbose logging (`true`/`false`, default: `false`)
- `server` — Server component name (e.g., `llamacpp`, `llamacpp-cuda`)
- `model` — Model component name (e.g., `model-qwen36-35b-a3b-ud-q4-k-xl`)
- `multimodel-projector` — Multimodal projector component name
- `gpu-layers` — Number of layers offloaded to GPU (CUDA only, e.g., `99`)

### `qwen36 set`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<key>=<value>` | positional | required | — | Configuration key-value pair to set |
| `--package` | flag (boolean) | optional | false | Set as package-level default (used during install hook) |

### `qwen36 completion`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<shell>` | positional | required | — | Shell type to generate completions for (known: `bash`) |

## Special Arguments

### Engine name validation
The `use-engine` positional argument is validated against available engine directories (`$SNAP/engines/*/engine.yaml`). Invalid engine names presumably produce an error.

### Dot-notation keys
Configuration keys use dot-notation (e.g., `http.port`) which maps to snap's hierarchical option system (`snapctl get http.port`).

### `--package` flag semantics
The `--package` flag on `set` distinguishes between package-level defaults (set during installation, lower precedence) and user-level overrides (set by the user, higher precedence). This mirrors snap's option precedence model.

### Stderr suppression in scripts
Several scripts redirect stderr when calling `qwen36 get model-name 2>/dev/null`, suggesting this key may produce warnings or errors when unset.
