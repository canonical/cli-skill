# Argument Structure: qwen36 / inference-snaps-cli

## Introduction

The CLI follows a **positional-argument + flag** pattern typical of Cobra applications. Most commands accept zero or one positional arguments; only `set`, `run`, `validate-engines`, and `serve-webui` accept multiple positional values. Flags are almost exclusively optional, with sensible defaults documented in flag descriptions. There is no subcommand-specific persistent flag beyond the global `--verbose`.

Common argument patterns:
- **Key-value config setting**: `set <key=value>` (exactly one `=` delimiter, value may contain `=`)
- **Optional key lookup**: `get [<key>]` — omission returns all configs
- **Optional engine name**: `show-engine [<engine>]` and `use-engine [<engine>]` — omission or `--auto` infers the target
- **Format serialization**: `--format` flag present on most information/inspection commands (`status`, `show-engine`, `show-machine`, `version`, `debug select-engine`)
- **Boolean confirmation flags**: `--assume-yes`, `--no-restart` appear on mutating commands (`set`, `unset`, `use-engine`) to suppress interactive prompts and service restarts

## Global Flags

| Flag | Shorthand | Type | Default | Env Var | Description |
|---|---|---|---|---|---|
| `--verbose` | `-v` | bool | `false` | none | Enable verbose logging |

## Per-Command Argument Map

### `status`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `--format` | no | string | `yaml` | `json`, `yaml` | — | — |
| `--wait-for-components` | no | bool | `false` | — | — | — |

### `chat`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| *(none)* | — | — | — | — | — | — |

### `webui`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| *(none)* | — | — | — | — | — | — |

### `get [<key>]`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `key` (positional) | no | string | — | any known config key | — | — |

### `set <key=value>...`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `key=value` (positional) | yes (≥1) | string | — | any config key | — | — |
| `--package` | no | bool | `false` | — | — | — |
| `--engine` | no | bool | `false` | — | — | — |
| `--assume-yes` | no | bool | `false` | — | — | — |
| `--no-restart` | no | bool | `false` | — | — | — |

*Notes*: `--package` and `--engine` are **hidden** flags. They are mutually exclusive in practice (switch statement on `cmd.packageConfig` / `cmd.engineConfig`).

### `unset <key>`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `key` (positional) | yes | string | — | any known config key | — | — |
| `--assume-yes` | no | bool | `false` | — | — | — |
| `--no-restart` | no | bool | `false` | — | — | — |

### `list-engines`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `--format` | no | string | `table` | `table`, `json` | — | — |

### `show-engine [<engine>]`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `engine` (positional) | no | string | active engine | any engine name from manifests | — | — |
| `--format` | no | string | `yaml` | `json`, `yaml` | — | — |

*Tab completion*: dynamically loaded from engine manifest names.

### `use-engine [<engine>]`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `engine` (positional) | no | string | — | any engine name from manifests | — | — |
| `--auto` | no | bool | `false` | — | — | — |
| `--fix` | no | bool | `false` | — | — | — |
| `--assume-yes` | no | bool | `false` | — | — | — |
| `--no-restart` | no | bool | `false` | — | — | — |

*Constraints*: cannot specify both a positional engine name and `--auto`; cannot specify both positional name and `--fix`. Exactly one mode must be chosen (explicit name, `--auto`, or `--fix`). No argument at all and no flag yields "engine name not specified".

### `show-machine`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `--format` | no | string | `yaml` | `json`, `yaml` | — | — |

### `prune-cache`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `--engine` | no | string | "" (all inactive) | any engine name from manifests | — | — |

*Constraints*: cannot prune the active engine. With no `--engine`, prunes all inactive engines.

### `version`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `--format` | no | string | `yaml` | `json`, `yaml` | — | — |

### `run <command>`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `command` (positional, ≥1) | yes | string | — | any executable | — | — |
| `--wait-for-components` | no | bool | `false` | — | — | — |

*Note*: `--wait-for-components` is **deprecated** with message `"run" always waits for components.`

### `serve-webui <static-files-dir>`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `static-files-dir` (positional) | yes | string | — | filesystem path | — | — |
| `--port` | no | int | `8081` | 1–65535 | — | — |
| `--host` | no | string | `localhost` | valid IP or hostname | — | — |
| `--capabilities` | no | string | `text` | comma-separated from `webui.SupportedCapabilities()` | — | — |

### `debug validate-engines <manifest>...`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `manifest` (positional, ≥1) | yes | string | — | filesystem path | — | — |

### `debug select-engine`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `--format` | no | string | `yaml` | `json`, `yaml` | — | — |
| `--engines` | no | string | `$SNAP/engines` | directory path | — | — |

*Note*: Reads hardware info YAML from **stdin** (stream input, not a flag or positional arg).

### `debug chat`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `--base-url` | yes | string | "" | valid HTTP(S) URL | — | — |
| `--model` | no | string | "" | any model name | — | — |

### `debug serve-webui <static-files-dir>`
| Arg/Flag | Required | Type | Default | Accepted Values | Aliases | Env Var |
|---|---|---|---|---|---|---|
| `static-files-dir` (positional) | yes | string | — | filesystem path | — | — |
| `--base-url` | no | string | `http://localhost:8080/v1` | valid HTTP(S) URL | — | — |
| `--port` | no | int | `8081` | 1–65535 | — | — |

## Special Arguments

### 1. Key-Value Syntax with Embedded Equal Signs (`set`)
The `set` command splits positional arguments on the **first** `=` only. This allows values like `passthrough.environment.LD_LIBRARY_PATH=/custom/lib`. The parser rejects arguments starting with `=` and requires exactly two parts after the split. Duplicate keys are rejected.

### 2. Passthrough Configuration Keys
Any key prefixed with `passthrough.` is exempt from key-existence validation in `set` and `unset`. During `run`, passthrough keys under `passthrough.environment.<key>` are translated to environment variables (`<KEY>` with hyphens converted to underscores). This provides an escape hatch for arbitrary engine environment injection.

### 3. Deprecation-Only Flag (`run --wait-for-components`)
`--wait-for-components` on `run` exists solely to satisfy prior users. It always behaves as if true; the flag is ignored but accepted to avoid breaking scripts. This is an unusual pattern: the flag is functionally a no-op.

### 4. Feature-Gated Command Registration (`chat`, `webui`)
Top-level `chat` and `webui` commands are dynamically added to the command tree based on the `ADDITIONAL_FEATURES` environment variable at process startup. They do not appear in help or tab completion when the feature is absent. This is structural dynamism driven by the environment, not by flags.

### 5. Hidden Flags on `set` (`--package`, `--engine`)
These flags are excluded from help text. They exist to allow snap hooks and internal scripts to write to lower-precedence configuration layers without user visibility. Because they are hidden, users cannot discover them from `--help`.

### 6. Dynamic Tab Completion for Engine Names
`show-engine`, `use-engine`, and `prune-cache --engine` all provide dynamic shell completion sourced from engine manifest filenames in `$SNAP/engines/`. This means completion is filesystem-dependent and will vary across installations.
