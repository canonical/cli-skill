# Juju CLI Architecture

## Scope and evidence

This analysis is based on the current Juju source tree, with the primary entrypoint and command registry in `cmd/juju/commands/main.go`, the generic CLI framework in `cmd/cmd`, command wrappers in `cmd/modelcmd`, and per-command implementations under `cmd/juju/*`. I also cross-checked the runtime-generated command documentation produced by `juju documentation --split` and the command help output from a built CLI binary.

## Tech stack

- Language: Go
- CLI framework: Juju's in-repo `cmd/cmd` package, built on `github.com/juju/gnuflag`
- Command host binary: `cmd/juju/main.go`
- Command registration: `cmd/juju/commands/main.go`
- Command wrappers: `cmd/modelcmd` and controller/model-specific bases
- API access: Juju client and facade packages under `api/*`
- Formatting/output helpers: `cmd/cmd/output.go`, `cmd/cmd/markdown.go`, `core/output`
- Extension hooks: plugin dispatch in `cmd/juju/commands/plugin.go`, local aliases through `cmd/cmd/aliasfile.go`

## Primary architecture style

Primary style: layered CLI application.

Secondary styles:
- command host / command registry
- thin client-server CLI for most operational commands
- limited plugin host for unknown subcommands

The CLI is not a shell-script style command bag. It is a structured command host with a central registry, shared wrappers for model/controller resolution, and a large number of thin command objects that mostly validate input, open the right API facade, and format results.

## Layer breakdown

### 1. Process and bootstrap layer

`cmd/juju/main.go` seeds randomness, configures logging, and hands control to `commands.Main(os.Args)`.

`cmd/juju/commands/main.go` then:
- creates the default command context
- initializes Juju's XDG data directory
- installs proxy settings into the default transport
- optionally refreshes public cloud metadata on first run
- diverts bare `juju` and single-token help invocations into the interactive REPL / help path
- rewrites bare `--version` to the explicit `version` command
- constructs the top-level `SuperCommand`

### 2. Command host and framework layer

`cmd/cmd` provides the framework primitives:
- `Command` and `CommandBase`
- `Info` for usage/help metadata
- `SuperCommand` for command dispatch, aliases, deprecation, and missing-command handling
- output formatting via `Output`
- markdown documentation generation via `documentationCommand`
- common error handling and exit codes in `cmd.Main`

Important framework behaviors:
- parse/init failures return exit code `2`
- runtime command failures return exit code `1`
- `gnuflag.ErrHelp` returns `0`
- plugin exit codes can be passed through verbatim using `RcPassthroughError`
- machine-oriented output formats (`json`, `yaml`) trigger error-shaping logic in the supercommand so consumers can still read an empty serialisable value from stdout

### 3. Shared Juju command wrappers

`cmd/modelcmd` is the next important layer. It gives commands a consistent way to:
- resolve current controller and model from local client state
- accept `-m/--model` and `-c/--controller` when appropriate
- open API connections
- manage login state, redirects, migrated-model handling, and store cleanup
- add confirmation helpers for destructive operations

This layer is where a lot of the CLI's actual ergonomics live. Many top-level commands are thin wrappers around `modelcmd.Wrap`, `WrapController`, or related helpers.

### 4. Command implementation layer

Each top-level functional area has its own package under `cmd/juju`:
- `application`, `model`, `controller`, `cloud`, `machine`, `storage`, `space`, `secrets`, `secretbackends`, `user`, `status`, `action`, `ssh`, `resource`, `crossmodel`, and so on.

Most commands follow the same shape:
- declare a small command struct with parsed flag state
- expose `Info()` and `SetFlags()`
- validate args in `Init()`
- resolve API clients in `Run()`
- format results through `cmd.Output` or package-specific tabular renderers

### 5. API / server interaction layer

Operational commands are usually client stubs over Juju facades:
- `application` commands use application and model-config clients
- `controller` commands use controller-specific clients
- `model` commands use modelmanager or modelupgrader clients
- `status` uses the legacy status client facade
- `cloud` commands use cloud/credential clients and local client store state

This keeps most CLI commands thin. The tradeoff is that behavior details and validation sometimes live partly in the client, partly in wrappers, and partly in server-side facades.

## Typical request path

A representative path for `juju model-config ftp-proxy=10.0.0.1:8000` is:

1. `main()` calls `commands.Main`.
2. `commands.Main` constructs the Juju `SuperCommand`.
3. `registerCommands` registers `model.NewConfigCommand()`.
4. `cmd.Main` parses top-level flags and dispatches to the selected subcommand.
5. `modelcmd.Wrap` resolves controller/model context from `-m`, environment, or local client store.
6. `model/config.go` parses file/reset/key=value inputs through `cmd/juju/config.ConfigCommandBase`.
7. The command opens the relevant API facade.
8. The command sends the request, then renders the result using the selected formatter.

For plugin commands, the path diverges after step 4:
- an unknown subcommand hits the missing callback
- `RunPlugin` looks for an executable named `juju-<subcommand>` in `PATH`
- only selected common flags are parsed and forwarded as environment (`JUJU_MODEL`, `JUJU_CONTROLLER`)
- the plugin runs as a subprocess with stdin/stdout/stderr attached

## Strengths

- Strong central registry. `registerCommands` gives one concrete place to inventory the built-in command surface.
- Consistent command object shape. The `Info/SetFlags/Init/Run` contract keeps most commands mechanically similar.
- Good wrapper reuse. `modelcmd` removes a large amount of repetitive controller/model-resolution code.
- Built-in documentation generation. `documentationCommand` can emit markdown directly from command metadata.
- Machine output support exists across a meaningful slice of the CLI through shared `Output` helpers.
- Extensibility is pragmatic. PATH-based plugins and user alias files are simple and require little infrastructure.

## Constraints and architectural debt

- The command registry is flat. Commands are top-level rather than deeply hierarchical, so discoverability relies heavily on naming discipline.
- Command families are package-grouped in source, but not represented as nested subcommands in the UX.
- The `Info.Args` string is freeform. It is easy for commands to leak formatting hints or duplicate `[options]` into usage text.
- Output contracts are fragmented. Some commands use shared formatters, others package-specific table writers, and schema stability is not documented centrally.
- Safety policy is not centralized. Confirmation, force flags, block checks, and dry-run behavior are implemented command by command.
- Exit codes are framework-consistent but semantically shallow: mostly `0`, `1`, `2`, plus plugin passthrough.
- Some framework features are slightly leaky. Examples include special handling for `help`, dynamic default formatting in `actions`, and serialisable error shaping in `SuperCommand`.

## Architectural consequences for command design

The current architecture favors incremental addition of top-level commands. That makes Juju flexible, but it also leads to surface-area sprawl:
- verb selection becomes more important than hierarchy
- command discoverability depends on consistent naming and docs quality
- asymmetries are easy to introduce because no structural grammar enforces complement pairs
- machine output and safety semantics drift unless they are actively standardized

That tension shows up throughout the current CLI: operational breadth is strong, but command semantics, naming, and output guarantees are not normalized to the same degree as the registry and wrapper plumbing.
