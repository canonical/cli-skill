# Juju CLI Extensibility Model

## Adding New Commands

### Manual Registration

New commands must be explicitly registered in `cmd/juju/commands/main.go`:

```go
func registerCommands(r commandRegistry) {
    // ... existing commands ...
    r.Register(mycommand.NewMyCommand())
}
```

There is no automatic discovery via reflection, annotations, or file naming conventions. Every command requires a code change to the central registry.

### Command Interface

New commands must implement the `cmd.Command` interface:

```go
type Command interface {
    IsSuperCommand() bool
    Info() *Info
    SetFlags(f *gnuflag.FlagSet)
    Init(args []string) error
    Run(ctx *Context) error
    AllowInterspersedFlags() bool
}
```

For commands that need API access, embed `modelcmd.ModelCommandBase`.

### Command Base Types

| Base Type | Use Case |
|---|---|
| `cmd.CommandBase` | Simple commands with no API access |
| `modelcmd.ModelCommandBase` | Commands that operate within a model context |
| `modelcmd.ControllerCommandBase` | Commands that operate at controller scope |
| `modelcmd.OptionalControllerCommand` | Commands that may work without a controller |

### Package Organization

Commands are grouped by domain in `cmd/juju/<domain>/`. The package name typically matches the domain (e.g., `cloud`, `storage`, `secrets`).

## Plugin System

### External Plugins

The CLI supports external plugins via the missing command callback:

```go
missingCallback = RunPlugin(missingCallback)
```

When a command is not found in the registry, the CLI attempts to execute `juju-<command>` from `$PATH`.

### Plugin Behavior

- Plugin name: `juju-<subcommand>` (e.g., `juju-kubectl` for `juju kubectl`)
- Arguments: All args after the subcommand are passed through
- Exit code: Passed through directly
- Stdout/Stderr: Passed through directly
- No built-in help integration for plugins

### Plugin Limitations

- Plugins cannot register flags in `juju --help`
- Plugins do not inherit model context or authentication
- No plugin directory or management commands exist

## Embedded Commands

A subset of commands can run in "embedded" mode for the Juju Dashboard. This is controlled by:

1. The `embedded` parameter in `NewJujuCommandWithStore()`
2. A whitelist of allowed embedded commands
3. The `supportsEmbedded` interface on commands

Commands that are not whitelisted return:
```
juju "foo" is not supported when run via a controller API call
```

## Command Aliases

Aliases are defined in two ways:

### Built-in Aliases

Set in the command's `Info().Aliases`:
```go
func (c *MyCommand) Info() *cmd.Info {
    return &cmd.Info{
        Name:    "my-cmd",
        Aliases: []string{"my-command"},
        // ...
    }
}
```

### User Aliases

Users can define aliases in `~/.local/share/juju/aliases`:
```
mydeploy = deploy --channel=stable
```

These are loaded by the supercommand at startup.

## Deprecation Mechanism

The command registry supports deprecated commands:

```go
r.RegisterDeprecated(mycommand.NewOldCommand(), mycommand.DeprecationCheck)
```

Deprecated commands:
- Still function normally
- Print a deprecation warning to stderr
- Are hidden from `juju help commands` (but shown in docs)

There is no automatic migration path; users must manually update scripts.

## Feature Flags

Commands can be conditionally registered using feature flags:

```go
if featureflag.Enabled(featureflag.DeveloperMode) {
    r.Register(model.NewDumpCommand())
}
```

Feature flags are set via the `JUJU_FEATURE_FLAGS` environment variable.

## Middleware and Hooks

There is no formal middleware system. However, the command framework provides several hooks:

| Hook | Mechanism |
|---|---|
| Pre-run notification | `SuperCommandParams.NotifyRun` |
| Pre-help notification | `SuperCommandParams.NotifyHelp` |
| Missing command | `SuperCommandParams.MissingCallback` |
| Post-init validation | `Init()` method on each command |

## Extension Boundaries

### What Can Be Extended

- New top-level commands via plugin binaries
- New built-in commands via code changes
- User aliases for existing commands
- Output formatters (custom `cmd.Formatter` implementations)

### What Cannot Be Extended

- New global flags without modifying `modelcmd` base types
- New output formats without modifying `cmd/output.go`
- Subcommand nesting (the framework supports it but Juju uses a flat namespace)
- Command interception or middleware wrapping

## Documentation Generation

The CLI can generate markdown documentation for all registered commands:

```bash
go run ./cmd/juju documentation --split --no-index --out /tmp/juju-cli-docs
```

This iterates the command registry and produces one markdown file per command. The documentation generator is in `cmd/cmd/documentation.go`.
