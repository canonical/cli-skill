# Juju CLI Configuration Model

## Configuration sources

Observed sources (highest precedence first):

1. Explicit command flags and arguments
- Command-specific flags are parsed during command `Init`.
- Super-command/global flags are parsed before subcommand dispatch.

2. Environment variables
- `JUJU_MODEL` and `JUJU_CONTROLLER` influence model/controller selection.
- Feature flags loaded from `JUJU_FEATURE_FLAGS`-style environment (via `featureflag.SetFlagsFromEnvironment(...)`).

3. Local client store / config files
- Current controller and current model are read from client store when flags/env are absent.
- User aliases are loaded from the aliases file under Juju XDG data home.

4. Built-in defaults
- Command defaults encoded in each command's flag definitions.
- Formatting defaults often set to `smart` or command-specific defaults.

## Precedence and resolution behavior

### Controller resolution

Resolution order:
1. Controller implied by `JUJU_MODEL` (qualified model)
2. `JUJU_CONTROLLER`
3. Current controller in client store

Conflict behavior:
- If controller from `JUJU_MODEL` conflicts with `JUJU_CONTROLLER`, command resolution fails with an explicit conflict error.

### Model resolution

Resolution order:
1. Explicit model identifier from command/flag wrapper
2. `JUJU_MODEL`
3. Current model in client store (for resolved controller)

If defaults are disallowed by the command and no explicit model is provided, command fails with no-model error.

## Command-specific overrides and notable patterns

- Remove/destroy confirmation behavior can be influenced by model config mode and `--no-prompt` override semantics.
- Plugin invocations extract `-m/--model` and `-c/--controller` from incoming args and re-export via environment for plugin process execution.
- New install bootstrap path may fetch public cloud metadata before normal command execution.

## Surprising or non-obvious behavior

- `--version` is rewritten to the explicit `version` command under specific positional conditions to avoid built-in flag handling mismatch.
- Unknown subcommands do not immediately fail when plugin mode is enabled; the CLI attempts to execute `juju-<subcommand>` first.
- Some config behavior depends on runtime client-store state, so reproducibility requires controlling both env vars and local store content.

## Evidence pointers

- Command bootstrap and env feature flags: `cmd/juju/commands/main.go`
- Super-command/global flag parsing: `cmd/cmd/supercommand.go`
- Command main lifecycle and exit handling: `cmd/cmd/cmd.go`
- Model/controller precedence and conflict handling: `cmd/modelcmd/controller.go`, `cmd/modelcmd/modelcommand.go`
- Plugin env propagation: `cmd/juju/commands/plugin.go`
