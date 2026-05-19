# Juju CLI Argument Structure

## Overview

Juju uses a mostly consistent argument model built on `gnuflag` plus freeform `Info.Args` usage strings. The dominant pattern is:

`juju <verb-or-verb-noun-command> [global flags] [command flags] [positionals]`

The architecture is consistent, but the surface semantics are not fully normalized. The CLI mixes:
- pure verb commands: `deploy`, `consume`, `offer`, `register`, `switch`
- verb-noun commands: `add-model`, `remove-unit`, `show-controller`
- noun-ish list commands: `models`, `controllers`, `spaces`, `users`
- meta/debug commands: `debug-log`, `help-hook-commands`, `documentation`

## Common positional patterns

### Single required resource identifiers

This is the most common pattern.
Examples:
- `juju show-space <name>`
- `juju show-secret <ID>|<name>`
- `juju remove-cloud <cloud name>`
- `juju show-machine <machineID> ...`

### Resource plus optional override name

Common in deploy and cross-model flows.
Examples:
- `juju deploy <charm or bundle> [<application name>]`
- `juju consume <remote offer path> [<local application name>]`
- `juju offer ... [offer-name]`

### Multiple resource operands

Used for additive or destructive bulk operations.
Examples:
- `juju add-ssh-key <ssh key> ...`
- `juju remove-machine <machine number> ...`
- `juju detach-storage <storage> [<storage> ...]`
- `juju retry-provisioning <machine> [...]`

### Two-sided relations / pairwise operations

These use structurally asymmetric but semantically paired operands.
Examples:
- `juju integrate <application>[:<endpoint>] <application>[:<endpoint>]`
- `juju remove-relation <application1>[:<relation name1>] <application2>[:<relation name2>] | <relation-id>`
- `juju grant <user name> <permission> [<model name> ... | <offer url> ...]`

### Freeform command tail

A small set of commands accept a shell-like or opaque tail.
Examples:
- `juju exec <commands>`
- `juju ssh <[user@]target> [openssh options] [command]`
- `juju disable-command <command set> [message...]`

These are the least structured positionals in the surface and have the highest ambiguity cost.

## Common flag patterns

### Context-selection flags

These are injected by wrappers rather than repeated manually.
- `-m`, `--model`: common on model-scoped commands
- `-c`, `--controller`: common on controller-scoped commands
- `-B`, `--no-browser-login`: appears on many commands that may trigger browser auth

Resolution precedence is generally:
1. explicit CLI flag
2. command-specific identifier parsing such as `controller:model`
3. relevant environment variable or local client store state

### Output flags

Many but not all reporting commands use the shared `cmd.Output` helper:
- `--format`
- `-o`, `--output`

Formats are command-specific. The framework defaults are `smart`, `yaml`, and `json`, but many commands register their own formatter maps such as `tabular`, `summary`, `short`, or `human`.

### File-input flags

Juju uses `cmd.FileVar` for path-or-stdin semantics.
Common pattern:
- `--file path/to/file.yaml`
- `--file -` to read stdin

This is used in config flows and other YAML-backed update commands.

### Repeated multi-value flags

A number of commands use append-style semantics.
Examples:
- repeated `--config` in `deploy`
- repeated `--resource`
- repeated `--overlay`
- repeated `--reset` values merged later in config commands

### Boolean override / safety flags

Common safety toggles include:
- `--force`
- `--no-prompt`
- `--no-wait`
- `--destroy-storage`
- `--release-storage`
- `--dry-run`

The semantics are not globally standardized; each command interprets them locally.

## Flag value types in the framework

The framework and Juju helpers support several recurring value shapes:
- plain strings and booleans via `gnuflag`
- durations, such as destroy timeouts and retry delays
- comma-separated string lists via `cmd.StringsValue`
- repeated append-only string values via `cmd.AppendStringsValue`
- key/value maps via `cmd.StringMap`
- file-or-stdin handles via `cmd.FileVar`

This gives Juju good expressive range without inventing too many ad hoc parsers.

## High-frequency structural conventions

### Name-or-ID duals

Many commands accept more than one identifier type:
- `<ID>|<name>` for secrets
- task ID or prefix for `cancel-task`
- model names or `controller:model`
- cloud names plus optional region or credential suffixes

This is efficient for operators, but it increases the semantic load of help text.

### Inline `key=value` updates

A major CLI pattern across Juju is inline updates:
- app config: `juju config mysql foo=bar`
- model config: `juju model-config ftp-proxy=...`
- controller config: `juju controller-config auditing-enabled=true`
- storage pool updates: `juju update-storage-pool <name> k=v`
- action parameters: `juju run ... key=value`

This is one of the strongest cross-command conventions in the CLI.

### CSV-within-a-single-flag

Some commands expect comma-delimited items inside one argument rather than repeated flags or repeated positionals:
- `--reset key1,key2`
- relation ID lists in `resume-relation`
- CIDR allowlists in firewall rules
- some placement and space-binding values

This is script-friendly, but inconsistent with other parts of the surface that prefer repeated values.

## Special arguments

### Embedded flags inside `Args` strings

Some commands encode flags inside `Info.Args`, which leaks presentation concerns into the usage signature. Verified examples:
- `move-to-space`: `Args` includes `[--format yaml|json]`
- `spaces`: help signature includes format/output details as if positional structure
- `subnets`: same pattern

This is structurally unusual because flags should already be represented in the options table.

### Commands whose `Args` string already contains `[options]`

Several Charmhub commands embed `[options]` directly in `Info.Args`, and then the framework also prepends `[options]` from flag handling. Resulting help signatures include duplicated `[options]`.
Verified examples:
- `find`
- `download`
- `info`

### Commands with no explicit usage in generated markdown

The markdown documentation generator only emits a Usage section when `Info.Args` is non-empty. That means zero-arg commands like `clouds`, `controllers`, `whoami`, and `dashboard` lose an explicit usage signature in generated docs even though `juju help <command>` does show one.

### Dynamic default formatting

`actions` uses a synthetic `default` formatter because its default output changes based on `--schema`. That is an implementation smell in argument semantics: output defaults are not purely declarative.

### Special help behavior

`help` is not a normal top-level command in the registry. It is a built-in supercommand feature, and `juju`, `juju --help`, and `juju help` all enter specialized code paths. That makes its behavior harder to inventory mechanically than normal commands.

## Inconsistencies

### Naming structure

- Some families use explicit verb-noun pairs: `add-model`, `show-controller`, `remove-space`.
- Others use bare verbs where the noun is implied by flags or positionals: `deploy`, `offer`, `consume`, `trust`.
- Reporting commands mix noun plurals and `show-*` verbs: `models`, `users`, `storage-pools`, `show-model`, `show-user`, `show-storage`.

### Separator conventions

- Some bulk arguments are repeated positionals.
- Some are comma-separated in a single positional.
- Some are repeated flags.
- Some use `key=value` maps.

All four patterns are reasonable independently, but the CLI does not clearly reserve one convention per semantic category.

### Confirmation semantics

`--no-prompt` generally means confirmation bypass for destructive flows, but in `login` it means read the password from stdin without prompting. The flag name is reused for materially different operator intent.

### Error grammar around arguments

Argument errors are generally clear and local, but the tone varies:
- `no application name specified`
- `cannot mix --no-color and --color`
- `expected type to be charm or bundle`
- `cannot set and reset key ... simultaneously`

There is no uniform grammar for parse/validation failures.

## Net assessment

The Juju CLI has strong recurring patterns around:
- model/controller scoping
- inline key/value updates
- YAML file import
- list/show/report formatting

The biggest structural weakness is that the CLI relies on handwritten `Args` strings for usage contracts. That creates drift, duplicate option markers, and occasional leakage of formatting hints into positional syntax. The framework is strong enough to support more standardization than the current command set actually uses.
