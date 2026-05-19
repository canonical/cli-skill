# Juju CLI Architecture

## Scope and method

This analysis is based on static code and repository documentation in `/project/juju`.
Runtime probing (`go run ./cmd/juju ...`) was not available because Go was not installed in the environment.

## Tech stack summary

- Language: Go
- CLI framework: internal command framework in `cmd/cmd`
- Flag parser: `github.com/juju/gnuflag`
- Output formatting: built-in formatters (`smart`, `yaml`, `json`) plus command-specific formatters
- Plugin integration: executable discovery via `juju-*` binaries on `PATH`
- Docs generation: built-in `juju documentation` command and markdown generator scripts

## Architecture style classification

Primary style: Layered CLI application

Secondary styles:
- Plugin-based architecture
- Command bus architecture (super-command command registration + dispatch)

## Layer breakdown

1. Entry and bootstrap
- `cmd/juju/commands/main.go`
- Initializes context, XDG data home, proxy, feature flags.
- Selects REPL/help behavior and forwards to the super-command.

2. Command registry and dispatch
- `cmd/juju/commands/main.go` (`registerCommands`)
- `cmd/cmd/supercommand.go`
- Registers top-level commands and routes argv to a concrete command implementation.

3. Domain command implementations
- `cmd/juju/<domain>/...` packages (application, model, controller, cloud, storage, secrets, etc.)
- Commands expose `Info`, `SetFlags`, `Init`, `Run` and often share base classes.

4. Shared command infrastructure
- `cmd/cmd/cmd.go`: command lifecycle, context, exit behavior
- `cmd/cmd/output.go`: machine/human output handling
- `cmd/modelcmd/*`: model/controller resolution and confirmation helpers

5. Extension surface
- `cmd/juju/commands/plugin.go`: missing-command fallback to `juju-<subcommand>` plugin executables

## Request path (typical)

1. argv enters `Main` in `cmd/juju/commands/main.go`.
2. bootstrap checks run (context, local data home, proxy, optional cloud metadata update).
3. `NewJujuCommand...` builds a `SuperCommand` and `registerCommands` installs all known subcommands.
4. `SuperCommand.Init` parses global flags, resolves command/alias/help/plugin.
5. Concrete command `Init` and `Run` execute domain logic.
6. `cmd.Main` translates errors to stable shell exit behavior.

## Observed strengths

- Clear command lifecycle contract across all commands.
- Centralized registration makes discoverability and documentation generation tractable.
- Plugin fallback preserves extension without recompiling CLI.
- Shared output and error plumbing encourages consistency.

## Constraints / tradeoffs

- Registry is long and hand-maintained, making accidental naming drift possible.
- Some behavior depends on runtime context (controller/model store, env vars), so static review cannot fully validate all edge paths.
