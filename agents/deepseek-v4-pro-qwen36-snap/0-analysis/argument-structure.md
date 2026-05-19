# Argument Structure

## Introduction

The `qwen36` CLI uses a **flat top-level command** model with positional arguments, flags, and a single nested subcommand group (`debug`). The overall pattern follows standard Go/Cobra conventions:

- **Positional arguments** are used where the argument is the primary object of the command (e.g., `get <key>`, `set <key=value>`, `use-engine <engine>`).
- **Flags** use long-form `--flag` with values separated by space or `=`. Short flags are limited to the single global `-v` / `--verbose`.
- **Machine-readable output** is controlled by a `--format` flag (`yaml` or `json`) on data-display commands.
- **Safety flags** (`--assume-yes`, `--no-restart`) appear on mutation commands (`set`, `unset`, `use-engine`, `prune-cache`).
- **Hidden flags** (`--package`, `--engine` on `set`) exist for internal/test use.

Common patterns:
- Root-level `--verbose`/`-v` persists across all commands.
- `--format` uses `yaml` as default, `json` as alternative.
- Mutation commands require `sudo` (root) and prompt for confirmation before destructive actions.

## Complete Argument Map

### Global / Root Flags

| Flag | Short | Type | Default | Description | Applies To |
|------|-------|------|---------|-------------|------------|
| `--verbose` | `-v` | bool | `false` | Enable verbose logging | All commands |

---

### `status`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `--format` | string | No | `yaml` | `json`, `yaml` | Output format |
| `--wait-for-components` | bool | No | `false` | — | Wait for engine components before reporting |

No positional arguments.

---

### `chat`

No flags. No positional arguments.

Conditional: only present when `ADDITIONAL_FEATURES` env var includes `"chat"`.

---

### `webui`

No flags. No positional arguments.

Conditional: only present when `ADDITIONAL_FEATURES` env var includes `"webui"`.

---

### `get [<key>]`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<key>` | string (positional) | No | — | Any valid config key | Config key to retrieve; if omitted, prints all |

No flags.

---

### `set <key=value>...`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<key=value>` | string (positional, repeatable) | Yes (≥1) | — | `key=value` pairs | Configuration to set |
| `--assume-yes` | bool | No | `false` | — | Skip restart confirmation prompt |
| `--no-restart` | bool | No | `false` | — | Don't prompt to restart after changes |
| `--package` | bool (hidden) | No | `false` | — | Set package-level config (internal) |
| `--engine` | bool (hidden) | No | `false` | — | Set engine-level config (internal) |

Requires root. Duplicate keys in the same invocation are rejected.

---

### `unset <key>`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<key>` | string (positional) | Yes | — | Any valid config key | Config key to unset |
| `--assume-yes` | bool | No | `false` | — | Skip restart confirmation prompt |
| `--no-restart` | bool | No | `false` | — | Don't prompt to restart after changes |

Requires root.

---

### `list-engines`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `--format` | string | No | `table` | `table`, `json` | Output format |

No positional arguments.

---

### `show-engine [<engine>]`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<engine>` | string (positional) | No | — | Any installed engine name | Engine to show; if omitted, shows active engine |
| `--format` | string | No | `yaml` | `json`, `yaml` | Output format |

Supports shell completion for engine names via `engines.LoadManifests()`.

---

### `use-engine [<engine>]`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<engine>` | string (positional) | Conditional | — | Any installed engine name | Engine to switch to; required unless `--auto` or `--fix` |
| `--auto` | bool | No | `false` | — | Automatically select the best compatible engine |
| `--fix` | bool | No | `false` | — | Fix issues with the currently active engine |
| `--assume-yes` | bool | No | `false` | — | Skip prompts for component installation and restart |
| `--no-restart` | bool | No | `false` | — | Don't prompt to restart after switching engine |

Requires root. Mutually exclusive: `--auto` and `--fix` each cannot be combined with a positional engine name. `--auto` and `--fix` cannot be combined with each other.

Supports shell completion for engine names.

---

### `show-machine`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `--format` | string | No | `yaml` | `json`, `yaml` | Output format |

No positional arguments.

---

### `prune-cache`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `--engine` | string | No | `""` | Any inactive engine name | Prune a specific engine's components instead of all inactive |

No positional arguments. Requires root. Interactively prompts for confirmation (unless piped). Cannot prune the active engine.

---

### `version`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `--format` | string | No | `yaml` | `json`, `yaml` | Output format |

No positional arguments.

---

### `run <command>` (hidden)

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<command>` | string (positional) | Yes | — | Any executable | Command to run in engine environment |
| `[args...]` | string (positional, variadic) | No | — | — | Arguments passed to the command |
| `--wait-for-components` | bool (deprecated) | No | `false` | — | Now a no-op; components always waited for |
| `--` | separator | No | — | — | Separator between `run` flags and subprocess command+args |

---

### `serve-webui <static-files-dir>` (hidden)

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<static-files-dir>` | string (positional) | Yes | — | A directory path | Directory containing web UI static files |
| `--port` | int | No | `8081` | Any valid port | HTTP bind port |
| `--host` | string | No | `localhost` | Hostname/IP | HTTP bind address |
| `--capabilities` | string | No | `text` | Comma-separated capability names | Server capabilities to advertise |

---

### `debug` subcommand group

#### `debug validate-engines <manifest>...`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<manifest>` | string (positional, repeatable) | Yes (≥1) | — | File paths | Engine manifest YAML files to validate |

#### `debug select-engine`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `--format` | string | No | `yaml` | `json`, `yaml` | Output format |
| `--engines` | string | No | `$SNAP/engines` | Directory path | Engine manifests directory |

Reads hardware info as YAML from **stdin**.

#### `debug chat`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `--base-url` | string | **Yes** | `""` | OpenAI-compatible URL | Base URL of the OpenAI-compatible server |
| `--model` | string | No | `""` | Model name | Name of the model to use |

#### `debug serve-webui <static-files-dir>`

| Argument | Type | Required | Default | Accepted Values | Description |
|----------|------|----------|---------|-----------------|-------------|
| `<static-files-dir>` | string (positional) | Yes | — | A directory path | Directory containing web UI static files |
| `--base-url` | string | No | `http://localhost:8080/v1` | OpenAI-compatible URL | Base URL for the chat API |
| `--port` | int | No | `8081` | Any valid port | HTTP bind port |

---

## Special Arguments

### The `--auto` / `--fix` flags on `use-engine`

These flags change the command's argument expectations dynamically:
- `use-engine <engine>` — requires exactly 1 positional arg.
- `use-engine --auto` — forbids any positional arg.
- `use-engine --fix` — forbids any positional arg.

If both `--auto`/`--fix` and an engine name are provided, the command errors. This is a **flag-dependent arity pattern** — unusual but well-documented in the help text.

### The `--` separator on `run`

`run` uses `--` to separate its own flags from the subprocess command. This follows the POSIX convention:
```
qwen36 run -- echo "Hello World!"
```
Cobra's `Args: cobra.MinimumNArgs(1)` combined with `--` handling works correctly.

### The `--assume-yes` and `--no-restart` flags

These appear on `set`, `unset`, `use-engine`, and (in spirit) `prune-cache`. They control interactive prompting behavior:
- `--assume-yes`: auto-confirms restart prompts.
- `--no-restart`: suppresses restart prompts entirely.

`prune-cache` lacks `--assume-yes` — it uses `utils.IsTerminalOutput()` to detect non-interactive mode instead.

### The `--format` flag inconsistency

Commands that produce data output support `--format`:
- `status`: `yaml` (default), `json`
- `show-machine`: `yaml` (default), `json`
- `show-engine`: `yaml` (default), `json`
- `version`: `yaml` (default), `json`
- `list-engines`: `table` (default), `json`
- `debug select-engine`: `yaml` (default), `json`

`list-engines` is the **only command with a `table` format** and the **only command whose default is not `yaml`**.

### Environment Variable Feature Gating

`chat` and `webui` commands are conditionally compiled-in at runtime via `ADDITIONAL_FEATURES` env var. This is not a flag — it's an environment-variable-driven feature flag that determines whether commands appear at all.
