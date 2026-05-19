# Juju CLI Extensibility Model

## Built-in extension paths

1. Registering new built-in commands
- Add command constructor and register in `registerCommands` (`cmd/juju/commands/main.go`).
- Command must implement standard command interface (`Info`, `SetFlags`, `Init`, `Run`).

2. Command aliases and deprecation transitions
- Framework supports alias registration and deprecation/obsolescence checks.
- Enables migration without immediate breaking rename/removal.

3. Documentation generation integration
- Built-in `documentation` command and scripts generate markdown docs from registered commands.
- New commands become documentable through the same pipeline.

## Plugin model

### Discovery and invocation

- Unknown subcommands can be treated as plugins named `juju-<subcommand>` discovered on PATH.
- Plugin pattern is constrained by executable naming regex.

### Context propagation

- Common Juju scoping args (`-m/--model`, `-c/--controller`) are extracted and exported as environment variables to plugin process.
- Stdio is wired through, making plugins first-class in shell workflows.

### Exit semantics

- Plugin exit codes are propagated back to caller using passthrough error handling.

## Extension boundaries

- Plugin commands extend command surface but do not automatically join built-in command registry internals.
- Built-in command additions require source change and rebuild; plugin additions do not.
- Embedded/dashboard allow-listing can restrict which registered commands are available in embedded contexts.

## Middleware/hooks touchpoints

- Missing-command callback path provides the interception point for plugin execution.
- Notify hooks in super-command support side effects around command run/help flows.

## Risks and governance considerations

- PATH-based plugin discovery introduces supply-chain and environment-order sensitivity.
- Naming collisions between future built-ins and existing plugins may affect discoverability.
- Recommendation: reserve names and establish plugin signing/provenance guidance in operational docs.

## Evidence pointers

- Built-in registry: `cmd/juju/commands/main.go`
- Missing callback and dispatch internals: `cmd/cmd/supercommand.go`
- Plugin discovery/execution: `cmd/juju/commands/plugin.go`
- Documentation generator: `cmd/cmd/documentation.go`, `scripts/md-gen/juju-cli-commands/main.go`
