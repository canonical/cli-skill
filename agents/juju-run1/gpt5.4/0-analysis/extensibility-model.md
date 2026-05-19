# Juju CLI Extensibility Model

## Overview

Juju is extensible in three distinct ways:
- adding built-in commands in the source tree and registry
- dynamic PATH plugins using the `juju-<name>` convention
- client-local aliases loaded from the Juju aliases file

This is a pragmatic model rather than a heavy plugin API. The built-in extensibility story is strong for Juju developers, while third-party extension is intentionally simple and subprocess-based.

## 1. Built-in command extensibility

The main built-in extension point is `registerCommands` in `cmd/juju/commands/main.go`.

To add a built-in command, a developer generally:
1. implements a new command type under the relevant `cmd/juju/<domain>` package
2. wires `Info`, `SetFlags`, `Init`, and `Run`
3. registers the constructor in `registerCommands`
4. optionally ensures the command is whitelisted for embedded/controller-hosted use

Strengths:
- central registry is easy to audit
- source package boundaries are clear
- generated help and markdown docs derive automatically from command metadata

Constraint:
- the registry is flat and manual; there is no declarative command manifest or typed schema for the surface

## 2. Framework aliases and compatibility shims

The `SuperCommand` supports:
- command aliases through `Info().Aliases`
- `RegisterAlias`
- `RegisterSuperAlias`
- deprecation checks for aliases and commands

This is the internal compatibility mechanism for evolving the CLI without immediately breaking old names.

It is more of a migration facility than an open extension API, but it is important for command-surface maintenance.

## 3. User aliases

Juju loads user aliases from `osenv.JujuXDGDataHomePath("aliases")`.

The alias file format is simple:
- `name = cmd [args...]`
- blank lines and `#` comments are ignored
- values are tokenized with `strings.Fields`

Behavioral implications:
- users can create shortcuts or defaulted command expansions
- alias processing is client-local
- alias expansion is textual and simple, not semantic
- `--no-alias` is supported by the supercommand to bypass alias processing

Strengths:
- extremely lightweight
- good for operator ergonomics

Constraints:
- no validation beyond parseability
- no structured help integration for user aliases
- quoting/field-splitting is intentionally simple and therefore limited

## 4. PATH plugins

Unknown commands are routed through the missing-command callback, which is wrapped by `RunPlugin` when Juju is not embedded.

### Discovery model

A plugin is any executable on `PATH` matching the `juju-<name>` convention.

Discovery behavior:
- `findPlugins()` scans every directory in `PATH`
- files must match `^juju-[a-zA-Z]`
- files must be executable
- names are sorted for deterministic output

### Invocation model

When the user runs `juju foo ...` and no built-in command matches:
1. Juju tries to execute `juju-foo`
2. selected common arguments are extracted from the argument list
3. `-m/--model` and `-c/--controller` are parsed and propagated via environment variables
4. stdin/stdout/stderr are attached directly to the subprocess
5. plugin exit codes can pass through to the Juju process

### Help/discoverability model

Plugins are expected to support `--description`.
`GetPluginDescriptions()` runs plugins in parallel with `--description` to collect one-line descriptions for help surfaces.

Strengths:
- simple distribution model
- no in-process ABI or SDK burden
- shell-friendly exit-code passthrough
- parallel description collection avoids serial plugin startup penalties

Constraints:
- plugin contract is intentionally thin
- no typed capability negotiation
- no structured completion/help API beyond `--description`
- only model/controller context is forwarded explicitly; richer Juju integration is left to the plugin author
- plugin name collisions with future built-ins are possible in principle, though built-ins win by being registered first

## 5. Documentation extensibility

The built-in `documentation` command can generate markdown for the command tree from `Info()` metadata. This is not an extension mechanism by itself, but it lowers the cost of adding commands because docs can be regenerated from source.

However, the documentation system only reflects what is encoded in `Info()` and flags. It does not automatically guarantee semantic completeness or accuracy.

## 6. Embedded-command constraints

The registry has an embedded mode and a whitelist / exclusion mechanism. That matters for commands exposed via controller API calls or embedded environments:
- commands can be denied in embedded mode
- excluded commands produce explicit unsupported-command errors

This is an internal extension boundary: the same command host can be reused in more restricted contexts.

## 7. Extensibility strengths

- Very low friction for built-ins.
- PATH plugins are dead simple to distribute and debug.
- User aliases provide lightweight local customization.
- Aliases and deprecation hooks support command-surface evolution.
- Generated docs reduce integration cost for new built-ins.

## 8. Extensibility constraints and gaps

- No formal plugin SDK or typed plugin manifest.
- No stable structured metadata format for commands, beyond human help and markdown generation.
- User aliases are local and text-based, not portable command objects.
- Help/discovery for plugins is one-line-description only.
- Because the CLI is flat, adding new built-ins increases top-level namespace pressure quickly.

## Assessment

Juju's extensibility model is deliberately pragmatic:
- built-in commands are easy for maintainers to add
- operators can bend the CLI with aliases
- third parties can add plugin commands without linking against Juju internals

That model is a good fit for a mature Go CLI. The tradeoff is that extensibility is broad but shallow: there are many extension points, but few deep guarantees about integration, metadata, or lifecycle. For Juju's current scale, the bigger challenge is not adding commands. It is keeping the command surface coherent as it grows.
