# Argument Structure

## Common Patterns

1. **Key-value positional arguments**: `set` uses `key=value` as positional args (one or more). `get` uses an optional single key.
2. **Named engine positional**: `use-engine`, `show-engine` accept an optional engine name positionally.
3. **No positional args**: Most commands (`status`, `chat`, `webui`, `list-engines`, `show-machine`, `prune-cache`, `version`) take no positional arguments.
4. **Format flag**: A consistent `--format` flag across observation commands (status, show-engine, list-engines, show-machine, version).
5. **Root-only mutation**: All write commands (set, unset, use-engine, prune-cache) require root and fail with a permission error otherwise.
6. **Confirmation suppression**: `--assume-yes` on mutating commands.
7. **Restart control**: `--no-restart` on commands that would trigger a service restart.

## Complete Argument Map

### `status`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `--format` | string | no | `yaml` | Output format: `json`, `yaml` |
| `--wait-for-components` | bool | no | `false` | Wait for engine components to be installed |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `chat`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

No command-specific flags.

### `webui`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

No command-specific flags.

### `get [<key>]`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<key>` | positional string | no | (all) | Config key to retrieve |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `set <key=value>...`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<key=value>` | positional string(s) | yes (≥1) | — | One or more key=value pairs |
| `--assume-yes` | bool | no | `false` | Skip confirmation prompts |
| `--no-restart` | bool | no | `false` | Do not restart the snap service |
| `--package` | bool | no | `false` | Set package-layer config (hidden) |
| `--engine` | bool | no | `false` | Set engine-layer config (hidden) |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `unset <key>`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<key>` | positional string | yes | — | Config key to unset |
| `--assume-yes` | bool | no | `false` | Skip confirmation prompts |
| `--no-restart` | bool | no | `false` | Do not restart the snap service |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `list-engines`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `--format` | string | no | `table` | Output format: `table`, `json` |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `show-engine [<engine>]`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<engine>` | positional string | no | (active) | Engine name to show |
| `--format` | string | no | `yaml` | Output format: `json`, `yaml` |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `use-engine [<engine>]`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<engine>` | positional string | conditional | — | Engine name (required unless --auto or --fix) |
| `--auto` | bool | no | `false` | Auto-select best compatible engine |
| `--fix` | bool | no | `false` | Fix issues with currently active engine |
| `--assume-yes` | bool | no | `false` | Skip confirmation prompts |
| `--no-restart` | bool | no | `false` | Do not restart after engine change |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `show-machine`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `--format` | string | no | `yaml` | Output format: `json`, `yaml` |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `prune-cache`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `--engine` | string | no | `""` | Target a specific engine's caches |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `version`

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `--format` | string | no | `yaml` | Output format: `json`, `yaml` |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `run <command>` (hidden)

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<command>` | positional string(s) | yes (≥1) | — | Command and args to execute |
| `--wait-for-components` | bool | no | `false` | (deprecated) Wait for components |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

### `serve-webui <static-files-dir>` (hidden)

| Argument | Type | Required | Default | Description |
|----------|------|----------|---------|-------------|
| `<static-files-dir>` | positional string | yes | — | Path to static files |
| `--port` | int | no | `8081` | HTTP bind port |
| `--host` | string | no | `localhost` | HTTP bind address |
| `--capabilities` | string | no | `text` | Comma-separated capabilities list |
| `-v, --verbose` | bool | no | `false` | Enable verbose logging (global) |

## Special Arguments

### The `--` separator for `run`

The `run` command uses `--` to separate CLI flags from the subprocess command and its arguments, following POSIX conventions.

### Key=value syntax for `set`

The `set` command uses `key=value` as a positional argument pattern (not a flag value). Multiple pairs are accepted. Keys use dot-notation (e.g., `http.port`). This mirrors `snap set` conventions.

### Mutually exclusive flags on `use-engine`

`--auto`, `--fix`, and the positional `<engine>` argument are mutually exclusive. Providing both `--auto` and an engine name results in an error. The code validates these combinations at runtime rather than via cobra's built-in exclusivity.

### Hidden flags on `set`

The `--package` and `--engine` flags on `set` are hidden from help. They select the config layer (package defaults vs engine defaults vs user overrides). These are used internally by snap hooks and are not intended for end users.
