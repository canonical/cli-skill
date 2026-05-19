# Argument Structure

Overview: Map of commands to their argument and flag structure, derived from `cli/cmd/cli/commands` source.

## Common global flags
- `--verbose` / `-v`: global persistent flag for verbose logging

## Per-command arguments

- `status`
  - args: none
  - flags:
    - `--format` (string) default `yaml` accepted: `json`,`yaml`
    - `--wait-for-components` (bool)

- `chat`
  - args: none
  - flags: none

- `webui`
  - args: none
  - flags: none

- `list-engines`
  - args: none
  - flags:
    - `--format` (string) default `table` accepted: `table`,`json`

- `show-engine [<engine>]`
  - args: optional engine name (max 1)
  - flags:
    - `--format` (string) default `yaml` accepted: `json`,`yaml`

- `use-engine [<engine>]`
  - args: optional engine name (max 1)
  - flags:
    - `--auto` (bool)
    - `--fix` (bool)
    - `--assume-yes` (bool)
    - `--no-restart` (bool)

- `get [<key>]`
  - args: optional key (max 1)
  - flags: none

- `set <key=value>...`
  - args: one or more key=value pairs
  - flags:
    - `--package` (hidden)
    - `--engine` (hidden)
    - `--assume-yes` (bool)
    - `--no-restart` (bool)

- `unset <key>`
  - args: single key
  - flags:
    - `--assume-yes` (bool)
    - `--no-restart` (bool)

- `show-machine`
  - args: none
  - flags:
    - `--format` (string) default `yaml` accepted: `json`,`yaml`

- `prune-cache`
  - args: none
  - flags:
    - `--engine` (string) optional engine name

- `version`
  - args: none
  - flags:
    - `--format` (string) default `yaml` accepted: `json`,`yaml`

- `run <command>` (hidden)
  - args: command + args (passthrough)
  - flags:
    - `--wait-for-components` (deprecated)

- `serve-webui <static-files-dir>` (hidden)
  - args: exact 1 (static files dir)
  - flags:
    - `--port` (int) default 8081
    - `--host` (string) default `localhost`
    - `--capabilities` (comma-separated string)

## Special arguments
- `set` accepts multiple `key=value` pairs; values may contain `=` and are split on first `=`.
- `run` uses `--` to separate CLI flags from the subprocess arguments (Cobra standard).
