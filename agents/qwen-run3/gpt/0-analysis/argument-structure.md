# qwen36 Argument Structure

## Introduction

The public command surface uses a small set of consistent argument patterns:

- zero-argument inspection commands: `status`, `chat`, `webui`, `show-machine`, `version`
- optional-single-argument inspection commands: `get [<key>]`, `show-engine [<engine>]`
- key-value mutation: `set <key=value>...`
- single-key revert: `unset <key>`
- engine selection with modal flags: `use-engine [<engine>]`
- format flags on most inspection commands, but not on config commands

The CLI mostly avoids deep positional structures, but it does contain a few non-standard or hidden patterns that matter for scripting.

## Public Command Argument Map

| Command | Positional Arguments | Flags | Required | Defaults | Accepted Values / Notes | Env Var Mapping |
|---|---|---|---|---|---|---|
| `status` | none | `--format`, `--wait-for-components` | no args | `--format=yaml`, `--wait-for-components=false` | `json`, `yaml`; wait flag can block for up to 3600s | none |
| `chat` | none | none | no args | none | top-level command is feature-gated | none |
| `webui` | none | none | no args | none | top-level command is feature-gated | none |
| `get [<key>]` | optional `key` | none | key optional | no key means all resolved config | scalar key lookup or subtree lookup | none |
| `set <key=value>...` | one or more `key=value` items | `--assume-yes`, `--no-restart` | at least one assignment | both flags default false | hidden `--package` and `--engine` also exist but are not user-discoverable from help | none |
| `unset <key>` | exactly one `key` | `--assume-yes`, `--no-restart` | yes | both flags default false | only removes user-layer overrides | none |
| `list-engines` | none | `--format` | no args | `--format=table` | `table`, `json` | none |
| `show-engine [<engine>]` | optional `engine` | `--format` | no | active engine if omitted | completion comes from manifest names | none |
| `use-engine [<engine>]` | optional `engine` | `--auto`, `--fix`, `--assume-yes`, `--no-restart` | engine or exactly one mode flag required | all flags false | cannot combine explicit engine with `--auto` or `--fix` | none |
| `show-machine` | none | `--format` | no args | `--format=yaml` | `json`, `yaml` | none |
| `prune-cache` | none | `--engine` | no args | empty engine means all inactive engines | named engine must not be active | none |
| `version` | none | `--format` | no args | `--format=yaml` | `json`, `yaml` | none |

## Hidden And Debug Arguments

| Command | Positional Arguments | Flags | Notes |
|---|---|---|---|
| `run <command>` | one or more words forming the child command | deprecated `--wait-for-components` | command always waits for components; help text still shows the deprecated flag in examples |
| `serve-webui <static-files-dir>` | exactly one path | `--port`, `--host`, `--capabilities` | hidden internal server command |
| `debug validate-engines <manifest>...` | one or more manifest paths | none | schema validation utility |
| `debug select-engine` | none on argv; reads hardware YAML from stdin | `--format`, `--engines` | unusual stdin-driven command shape |
| `debug chat` | none | `--base-url`, `--model` | `--base-url` is required at runtime |
| `debug serve-webui <static-files-dir>` | exactly one path | `--base-url`, `--port` | localhost-only debug server |

## Special Arguments

### `set` key-value parsing

- parsing splits on the first `=` only
- values may themselves contain `=`
- a value starting with `=` is rejected because the key would be empty
- duplicate keys in one invocation are rejected

### `use-engine` modal behavior

`use-engine` is not a normal positional command. It has three mutually exclusive modes:

1. explicit engine: `use-engine cpu`
2. auto selection: `use-engine --auto`
3. repair current engine: `use-engine --fix`

This matters because the help text does not explain the exclusivity rules; only runtime validation does.

### hidden config-layer flags

`set` defines hidden flags:

- `--package`: write to package config
- `--engine`: write to engine config

These are operationally important because hooks use them, but they are not part of the discoverable public interface.

### passthrough keys

Keys under `passthrough.environment.*` are treated specially:

- they bypass normal key-existence validation
- they are transformed into environment variables during `run`
- hyphens are converted to underscores and names are uppercased

## Common Patterns And Gaps

- inspection commands expose `--format`, but `get` does not; it always prints scalar text or YAML
- config mutation commands have restart-suppression flags, but no dry-run or validation-only mode
- the public surface exposes no explicit `--help` examples for `unset`, `prune-cache`, or `show-machine`
- `run` still teaches a deprecated flag in its example block
