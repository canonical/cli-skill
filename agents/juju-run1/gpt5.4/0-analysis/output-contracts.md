# Juju CLI Output Contracts

## Overview

Juju has partial but meaningful support for machine-readable output. The shared framework in `cmd/cmd/output.go` provides `json`, `yaml`, and sometimes `smart`; individual commands add their own formatter maps and human/tabular renderers.

The result is not one uniform output contract. It is a family of local contracts with a shared flag convention.

## Framework-level behavior

### Shared formatter support

The framework provides these core formatters:
- `json`
- `yaml`
- `smart`

Commands opt into output support by wiring `cmd.Output.AddFlags(...)` in `SetFlags()`.

### Shared flags

When a command uses `cmd.Output`, it generally gets:
- `--format`
- `-o`, `--output`

### Serialisable mode

In the framework, `json` and `yaml` are marked serialisable. When one of those formats is requested, `SuperCommand` marks the context as machine-oriented and adjusts fatal-error handling:
- if a command fails before writing stdout output
- and the selected format is serialisable
- the framework writes an empty value such as `{}` to stdout
- while still reporting the error on stderr

This is useful for parsers that demand valid JSON/YAML, but it is an unusual contract and should be considered part of the CLI semantics.

## Commands that support `--format`

Verified from generated docs, the following built-ins expose `--format`:

`actions`, `cancel-task`, `charm-resources`, `clouds`, `config`, `constraints`, `controller-config`, `controllers`, `credentials`, `debug-log`, `disabled-commands`, `exec`, `find`, `find-offers`, `firewall-rules`, `info`, `machines`, `model-config`, `model-constraints`, `model-defaults`, `models`, `move-to-space`, `offers`, `operations`, `regions`, `resources`, `run`, `secret-backends`, `secrets`, `show-application`, `show-cloud`, `show-controller`, `show-credential`, `show-machine`, `show-model`, `show-offer`, `show-operation`, `show-secret`, `show-secret-backend`, `show-space`, `show-status-log`, `show-storage`, `show-task`, `show-unit`, `show-user`, `spaces`, `status`, `storage`, `storage-pools`, `subnets`, `users`, `version`, `whoami`

That is a substantial set, but still far from the full CLI.

## Major output families

### YAML / JSON only or YAML / JSON plus simple human output

Examples:
- `config` defaults to YAML and supports JSON
- `whoami` supports formatted structured output
- many `show-*` and list commands support YAML/JSON plus a package-specific human default

These are the easiest commands to script against, but their schemas are not centrally versioned.

### Tabular-first reporting commands

Examples:
- `controller-config`
- `model-config`
- `find`
- `controllers`
- `models`
- `storage-pools`
- `machines`

These typically default to `tabular` and fall back to structured formats when explicitly requested.

### Status-specific multi-format contract

`status` is the richest formatter surface in the CLI.

Supported formats:
- `tabular` (default)
- `line`
- `short`
- `oneline`
- `summary`
- `json`
- `yaml`

Behavior differences matter:
- non-tabular formats implicitly include relations/integrations and storage data
- `--relations`, `--integrations`, and `--storage` become redundant outside `tabular`
- `status` emits informational notices about ignored flags when those flags are specified in non-tabular formats

That makes `status` powerful, but not a simple stable one-schema interface.

### Dynamic default output

`actions` has a special case:
- default output is effectively tabular for normal listing
- but if `--schema` is used, tabular is invalid and the command falls back to YAML unless an explicit format is chosen

This is a good example of output semantics being command-specific rather than standardized across the CLI.

## Contract characteristics by family

### Config commands

- `config`: default `yaml`; supports `json`
- `model-config`: default `tabular`; supports `json`, `yaml`
- `controller-config`: default `tabular`; supports `json`, `yaml`

Single-key retrieval on model/controller config with default tabular output uses `FormatSmart`, which can emit a plain scalar rather than a table/object. That is ergonomic for humans but means the contract shape changes depending on both the requested format and whether one key or all keys were requested.

### Search / listing commands

Many list/search commands support tabular, JSON, and YAML, but with custom tabular renderers and command-specific field sets. Examples include `find`, `actions`, `resources`, `spaces`, and `offers`.

### Mutating commands with machine-readable acknowledgements

A few mutating commands also support formatted output, for example `move-to-space` and `run`. That is useful for automation, but not consistently available across all mutators.

## Inconsistencies

### Not all list/show commands support format flags

There is no universal rule that list/show commands must support `--format`. Many do, but not all operational or mutating commands return structured acknowledgements.

### Default format varies by family

- app config defaults to YAML
- controller/model config default to tabular
- status defaults to tabular
- some commands default to package-specific human renderers
- some commands use custom names like `human`, `summary`, `short`

This is workable, but users cannot assume a single default philosophy.

### Schema stability is implied, not documented

Commands expose JSON/YAML, but the repo does not provide a central stability statement for field names, missing fields, ordering, or backward compatibility for those schemas.

### Usage and docs metadata leak output details

Several commands embed format hints in the positional usage string instead of leaving them solely in the options table. Verified examples include `move-to-space`, `spaces`, and `subnets`.

### Error behavior under machine formats is non-obvious

The framework writes empty structured output on fatal errors for serialisable formats. That preserves parseability, but the behavior is easy to miss because it is implemented in the supercommand layer rather than documented per command.

## Practical guidance for users and automation

Most scriptable command families today are:
- `status` with explicit `--format=json`
- `show-*` and list commands that explicitly advertise JSON/YAML
- config retrieval commands with explicit `--format`
- user/controller/model metadata commands like `whoami`, `controllers`, and `models`

For automation, explicit `--format=json` is safer than relying on defaults. For config commands, callers should assume that default tabular or smart output is human-oriented only.

## Assessment

Strengths:
- a real shared output framework exists
- JSON/YAML support is broad enough to be useful
- serialisable mode is handled deliberately, not accidentally
- `status` offers a rich reporting contract

Weaknesses:
- output support is incomplete across the full surface
- defaults and shape conventions vary by family
- machine-readable schema guarantees are undocumented
- some command metadata around format support is malformed or duplicated

The current output model is good enough for advanced operators, but not yet disciplined enough to count as one coherent CLI-wide contract.
