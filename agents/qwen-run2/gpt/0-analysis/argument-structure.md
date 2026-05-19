# Argument Structure

## Introduction

The qwen36 command surface uses a small number of argument patterns:

- bare workflow commands with no arguments: `chat`, `show-engine`
- verb-noun mutation with either an enum positional or a mode flag: `use-engine`
- config-key commands using dot-notated keys: `get <key>` and `set <key>=<value>`
- a nested shell-integration command: `completion bash`
- a daemon service entrypoint with no user-facing arguments: `qwen36.server`

The CLI is conservative about flags in the observable surface. Only two flags are evidenced in repository sources: `--auto` and `--assume-yes` on `use-engine`, plus `--package` on `set` for packaging-time defaults.

## Complete command and argument map

| Command | Positional arguments | Flags / subcommands | Required? | Default | Type | Accepted values | Aliases | Env var mapping |
|---|---|---|---|---|---|---|---|---|
| `qwen36 chat` | none observed | none observed | no | start interactive chat against current local server | command | n/a | none observed | none observed |
| `qwen36 use-engine` | optional engine selector when not using `--auto` | `--auto`, `--assume-yes` | yes: one selection mode is required | none for manual invocation | positional enum + boolean flags | positional values observed: `cpu`, `cuda`; flag mode: `--auto` | none observed | none observed |
| `qwen36 show-engine` | none | none observed | no | show current engine description | command | n/a | none observed | none observed |
| `qwen36 get` | `<key>` | none observed | yes | none | string positional | observed keys: `http.port`, `http.host`, `http.base-path`, `verbose`, `server`, `model`, `multimodel-projector`, `gpu-layers`, `model-name` | none observed | none observed |
| `qwen36 set` | `<key>=<value>` | `--package` | yes | none | assignment positional + boolean flag | observed keys share the same key space as `get`; values are shell strings persisted into snap config | none observed | none observed |
| `qwen36 completion bash` | `bash` | `completion` command with shell target subcommand/argument | yes | none | string positional under meta-command | only `bash` is evidenced | none observed | none observed |
| `qwen36.server` | none observed | none observed | no | run daemon entrypoint | service command | n/a | none observed | none observed |

## Per-command details

### `qwen36 chat`

- Observed syntax: `qwen36 chat`
- No flags or positionals are evidenced in the repository.
- Behavioral dependency: requires a healthy local server selected through current engine config.
- Runtime values are indirect rather than CLI arguments: `chat.sh` resolves `http.port`, `http.base-path`, and optional `model-name` through `qwen36 get`.

### `qwen36 use-engine`

Observed forms:

- `qwen36 use-engine cpu`
- `qwen36 use-engine cuda`
- `qwen36 use-engine --auto`
- `qwen36 use-engine --auto --assume-yes`

Argument metadata:

| Element | Kind | Required | Description |
|---|---|---|---|
| `cpu` / `cuda` | positional enum | required unless `--auto` is used | Explicitly select the engine manifest |
| `--auto` | boolean flag | required unless an explicit engine value is used | Auto-detect the best available engine |
| `--assume-yes` | boolean flag | optional | Suppress confirmation prompts; observed in the install hook |

Constraints inferred from usage:

- `--auto` and an explicit engine value appear to represent alternative modes.
- `--assume-yes` implies the command may otherwise prompt interactively in some flows.
- No short flags are evidenced.

### `qwen36 show-engine`

- Observed syntax: `qwen36 show-engine`
- No flags or positionals are evidenced.
- Output is YAML and is parsed by scripts with `yq`, so arguments are intentionally minimal.

### `qwen36 get`

- Observed syntax: `qwen36 get <key>`
- The positional `<key>` is required.
- Key naming pattern is dot notation for structured config, for example:
  - `http.port`
  - `http.host`
  - `http.base-path`
- Non-HTTP keys observed in scripts:
  - `verbose`
  - `server`
  - `model`
  - `multimodel-projector`
  - `gpu-layers`
  - `model-name`

Inferred argument behavior:

- The key is a string token, not a constrained enum in the visible sources.
- The command is intended for scalar lookups used in shell substitution.
- Optional keys exist: scripts explicitly suppress errors for `model-name` when unset.

### `qwen36 set`

- Observed syntax: `qwen36 set <key>=<value>`
- Observed syntax in hook context: `qwen36 set --package <key>=<value>`

Argument metadata:

| Element | Kind | Required | Description |
|---|---|---|---|
| `<key>=<value>` | positional assignment | yes | Persist one config value into snap config |
| `--package` | boolean flag | optional | Mark the write as a packaging default rather than a normal user write |

Observed examples:

- `qwen36 set http.port=8326`
- `qwen36 set --package http.port="8326"`
- `qwen36 set --package http.host="127.0.0.1"`
- `qwen36 set --package verbose="false"`

Unverified points:

- Whether multiple assignments in one invocation are accepted by the Go wrapper.
- Whether value validation exists per key.

### `qwen36 completion bash`

- Observed syntax: `qwen36 completion bash`
- `bash` is required in the only evidenced form.
- The completion script consumes its stdout as a whitespace-delimited word list.
- No evidence of `zsh`, `fish`, or other shell targets exists in this repository.

### `qwen36.server`

- Snap app/service syntax: `qwen36.server`
- No user-facing flags or arguments are evidenced.
- Startup behavior is configured entirely through snap config and engine metadata, not command-line parameters.

## Special arguments

### Dot-notated config keys

`get` and `set` both operate on keys such as `http.port` and `http.base-path`. This is a structural exception relative to the rest of the CLI because users must know internal config key names rather than object names.

### Inline `key=value` assignment

`set` uses a single positional `key=value` token rather than separate `--key value` or `<key> <value>` arguments. This is efficient, but it makes shell quoting and whitespace handling the caller’s responsibility.

### Packaging-only mutation flag

`--package` on `set` is not a normal user workflow flag. It appears in the install hook and establishes package defaults. This is a non-standard structural exception because it exposes an installation-scope write mode alongside the regular command.

### Implicit mode selection in `use-engine`

`use-engine` supports two structural modes:

- explicit engine via positional enum (`cpu`, `cuda`)
- implicit hardware selection via `--auto`

That makes it the only command in the set where the required input shape changes depending on the flag set.

### Service surface outside the normal grammar

`qwen36.server` is not part of the top-level `qwen36 ...` hierarchy. It is a snap app/service entrypoint and should be understood as operational surface rather than a typical end-user CLI verb.

## Common argument patterns and gaps

What is consistent:

- no short flags are exposed in the observed surface
- state-inspection commands are argument-light
- mutation commands take one focused input shape each

What is missing or undocumented:

- no documented `unset` or reset path for config keys
- no documented way to enumerate valid keys for `get`/`set`
- no documented list of supported shells for `completion`
- no direct user-level arguments for daemon status, logs, or health