# Extensibility Model

## Overview

Juju supports two primary extension mechanisms: **external plugins** (via `PATH` discovery) and **internal command registration** (via the Go type system). There is no dynamic plugin loading at runtime.

## 1. External Plugin System

### Discovery Mechanism

When a user runs `juju <subcommand>` and the subcommand does not match any registered command, the `MissingCallback` flow is triggered:

1. The framework calls `RunPlugin()`, which searches `$PATH` for an executable named `juju-<subcommand>`
2. If found, the plugin is executed with the remaining arguments
3. Common Juju flags (`-m`, `--model`, `-c`, `--controller`) are extracted from the argument list and passed as environment variables or arguments to the plugin

### Plugin Conventions

- Plugin names **must** match `^juju-[a-zA-Z].*`
- Plugins are standalone executables on `$PATH`
- Plugins receive the full argument list after the subcommand
- Plugins are responsible for their own flag parsing, help, and error handling
- There is no SDK for plugin development; plugins must implement their own API client using the Juju Go client library or REST API

### Limitations

- No version checking between plugins and the Juju client
- No plugin installation/management commands
- Plugins cannot hook into the Juju command lifecycle or middleware
- Plugin output is not captured or formatted by Juju
- No plugin sandboxing or permission model

## 2. Internal Command Registration

### Registration Path

New commands are added by:

1. Creating a struct that embeds the appropriate base type (`modelcmd.ModelCommandBase`, `modelcmd.ControllerCommandBase`, etc.)
2. Implementing the `cmd.Command` interface:
   - `Info() *cmd.Info` — command metadata (name, purpose, doc, examples, see-also)
   - `Init(args []string) error` — argument validation
   - `Run(ctx *cmd.Context) error` — command execution
   - `SetFlags(f *gnuflag.FlagSet)` — flag registration
3. Calling `r.Register(NewMyCommand())` in `registerCommands()` in `cmd/juju/commands/main.go`

### Base Type Hierarchy (Extension Points)

```
cmd.Command (interface)
└── cmd.CommandBase              // Base functionality, embedding
    ├── modelcmd.CommandBase     // + client store, cookie jar, HTTP client
    │   ├── ModelCommandBase     // + model selection (-m flag)
    │   │   └── OptionalModelCommandBase  // model optional
    │   ├── ControllerCommandBase         // controller only
    │   │   └── OptionalControllerCommandBase
    │   └── bare CommandBase usage        // no model/controller needed
    └── cmd.SuperCommand          // Command that hosts subcommands
```

### Key Embedding Interfaces

| Interface | Purpose |
|-----------|---------|
| `cmd.Command` | Core: Info, Init, Run, SetFlags |
| `cmd.Output` | Adds `--format` and `-o/--output` flags |
| `cmd.Log` | Adds logging configuration flags |
| `modelcmd.Command` | Adds API connection, client store |
| `supportsEmbedded` | Marks command as embeddable in dashboard |
| `hasClientStore` | Receives the client store reference |

### Middleware / Hooks

The SuperCommand framework provides:

- **`NotifyRun`**: Called before any subcommand is executed. Used for the first-run welcome message.
- **`MissingCallback`**: Called when a subcommand is not found. Delegates to the plugin system.
- **`UserAliasesFilename`**: Supports user-defined command aliases via a file (default: `~/.local/share/juju/aliases`).
- **`RegisterDeprecated`**: Registers a command as deprecated with a deprecation check.
- **`RegisterSuperAlias`**: Creates an alias for a subcommand.

### Feature Flag Gating

Commands can be gated behind feature flags:

```go
if featureflag.Enabled(featureflag.DeveloperMode) {
    r.Register(model.NewDumpCommand())
    r.Register(model.NewDumpDBCommand())
}
```

Feature flags are set via the `JUJU_FEATURES` environment variable.

## 3. Command Discovery (Help System)

- `juju help commands` — lists all registered commands (not plugins)
- `juju help <command>` — shows detailed help for a command
- `juju help topics` — shows help topics (`basics`, etc.)
- `juju help-action-commands` — lists commands available inside `juju run` operations
- `juju help-hook-commands` — lists commands available inside charm hooks
- Help output can be generated as Markdown via `cmd.PrintMarkdown()`

### Help Topics

Additional topics can be registered via:

```go
jcmd.AddHelpTopic("basics", "Basic Help Summary", usageHelp)
jcmd.AddHelpTopicCallback(name, short, longCallback)
```

Currently, the only registered help topic is `"basics"`.

## 4. Embedded Commands (Dashboard)

Commands can be marked as embeddable for the Juju Dashboard web UI. The `supportsEmbedded` interface provides `SetEmbedded(bool)` for commands that can run in-browser. A whitelist controls which commands are available via the Dashboard API. This is configured in the apiserver and enforced in `jujuCommandRegistry.Register()`.

## 5. Extension Boundaries

| Area | Extensibility | Notes |
|------|--------------|-------|
| New top-level commands | ✅ Easy | Register a new command struct |
| New subcommand groups | ✅ Possible but unused | SuperCommand supports nesting, but juju only has 1 level |
| Plugins | ✅ Via PATH | No SDK, no dependency tracking |
| Output formatters | ✅ Extensible | Register via `cmd.DefaultFormatters` |
| Help topics | ✅ Extensible | `AddHelpTopic()` / `AddHelpTopicCallback()` |
| Cloud providers | ⚠️ Through jujud | Not directly CLI-extensible |
| CharmHub source | ⚠️ Hardcoded | CharmHub is the only source for `find`/`info`/`download` |
| Login providers | ✅ Extensible | `cmd/internal/loginprovider/` provides pluggable login strategies |
| Alias files | ✅ | User-defined aliases in `~/.local/share/juju/aliases` |
| Config backends | ⚠️ Limited | SECRET backends have add/remove commands but types are fixed |
| Firewall rules | ⚠️ Limited | Only well-known service names |

## 6. Adding a New Command (Checklist)

1. Create a new Go package under `cmd/juju/<domain>/` (or file in existing package)
2. Define a struct embedding the appropriate base type
3. Implement `Info()`, `Init()`, `Run()`, `SetFlags()`
4. May implement optional interfaces:
   - `supportsEmbedded` — for dashboard access
   - `hasClientStore` — for accessing the client store
   - Feature flag check — for gated features
5. Register in `registerCommands()` in `cmd/juju/commands/main.go`
6. Add test file (`*_test.go`) with `tc` test suite
7. Run `make install` to rebuild
