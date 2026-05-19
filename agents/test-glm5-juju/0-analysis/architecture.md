# Juju CLI Architecture Analysis

## Tech Stack Summary

The Juju CLI is built with the following technology stack:

| Component | Technology | Purpose |
|-----------|------------|---------|
| Language | Go 1.21+ | Primary implementation language |
| CLI Framework | Custom (cmd/cmd package) | Command parsing, help, output formatting |
| Flag Parsing | gnuflag (github.com/juju/gnuflag) | GNU-style flag handling with long/short forms |
| API Client | api/jujuclient | Controller/model connection management |
| Configuration | XDG Base Directory | User data storage (~/.local/share/juju) |
| Logging | loggo (github.com/juju/loggo) | Structured logging |
| Output Formatting | Multiple formatters | YAML, JSON, tabular, line formats |
| Charm Hub | Charmhub API | Remote charm/bundle discovery |

## Architecture Style

**Primary Style: Client-Server CLI with Plugin Extension**

The Juju CLI follows a client-server architecture where the CLI client (`juju`) communicates with a Juju controller via an API. The controller manages models, applications, and cloud resources.

**Secondary Style: Plugin-based Architecture**

The CLI supports external plugins that are discovered and executed via the filesystem (executables named `juju-*` in PATH).

## Architecture Diagram

```
┌─────────────────────────────────────────────────────────────────┐
│                         juju CLI                                 │
├─────────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────────────┐   │
│  │ SuperCommand │  │   Plugins    │  │    REPL/Interactive  │   │
│  │   (main.go)  │  │ (plugin.go)  │  │      (repl.go)       │   │
│  └──────┬───────┘  └──────┬───────┘  └──────────┬───────────┘   │
│         │                 │                      │               │
│  ┌──────▼─────────────────▼──────────────────────▼───────────┐   │
│  │                   cmd/cmd Package                          │   │
│  │  ┌─────────┐  ┌──────────┐  ┌─────────┐  ┌─────────────┐  │   │
│  │  │Command  │  │SuperCmd  │  │Context  │  │Formatters   │  │   │
│  │  │Interface│  │          │  │(stdio)  │  │(json/yaml)  │  │   │
│  │  └─────────┘  └──────────┘  └─────────┘  └─────────────┘  │   │
│  └────────────────────────────────────────────────────────────┘   │
│                              │                                   │
│  ┌───────────────────────────▼────────────────────────────────┐   │
│  │              Command Domain Packages                        │   │
│  │  ┌─────────┐ ┌────────┐ ┌─────────┐ ┌────────┐ ┌────────┐  │   │
│  │  │action/  │ │cloud/  │ │model/   │ │user/   │ │ssh/    │  │   │
│  │  │application/ │controller/ │storage/ │secrets/ │...    │  │   │
│  │  └─────────┘ └────────┘ └─────────┘ └────────┘ └────────┘  │   │
│  └────────────────────────────────────────────────────────────┘   │
│                              │                                   │
└──────────────────────────────┼───────────────────────────────────────┘
                               │
                               │ API Connection
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                     Juju Controller API                          │
│  ┌───────────┐  ┌──────────┐  ┌───────────┐  ┌──────────────┐   │
│  │ Models    │  │ Apps     │  │ Machines   │  │ Clouds      │   │
│  │ Users     │  │ Secrets  │  │ Storage    │  │ Credentials │   │
│  └───────────┘  └──────────┘  └───────────┘  └──────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

## Key Architectural Components

### 1. Command Framework (cmd/cmd)

The `cmd/cmd` package provides the foundational CLI infrastructure:

- **Command Interface**: Defines the contract for all commands (`Info()`, `SetFlags()`, `Init()`, `Run()`)
- **SuperCommand**: Hierarchical command composition (e.g., `juju add-model`)
- **Context**: Execution context with I/O streams, environment variables
- **Formatters**: Output formatting (tabular, JSON, YAML, line, oneline, summary)

### 2. Command Registration (cmd/juju/commands/main.go)

Commands are registered in `registerCommands()`:

```go
func registerCommands(r commandRegistry) {
    r.Register(newBootstrapCommand())
    r.Register(application.NewDeployCommand())
    r.Register(model.NewConfigCommand())
    // ... 100+ commands
}
```

### 3. Plugin System (cmd/juju/commands/plugin.go)

External plugins are discovered via:
- Pattern matching: `^juju-[a-zA-Z]`
- Executed with inherited environment (JUJU_CONTROLLER, JUJU_MODEL)
- Fallback to MissingCallback for "did you mean?" suggestions

### 4. Configuration Storage

User configuration stored in XDG-compliant locations:
- Default: `~/.local/share/juju/`
- Override: `JUJU_DATA` environment variable
- Files: `controllers.yaml`, `models.yaml`, `accounts.yaml`, `credentials.yaml`

### 5. API Client Store

The `api/jujuclient` package manages:
- Connection caching
- Credential management  
- Controller/model UUID resolution

## Design Patterns Used

| Pattern | Usage |
|---------|-------|
| **Command Pattern** | Each command encapsulates a request as an object |
| **Factory Pattern** | `NewXCommand()` functions create command instances |
| **Strategy Pattern** | Multiple output formatters (tabular, JSON, YAML) |
| **Decorator Pattern** | `ModelCommand` wraps `Command` with model context |
| **Template Method** | `CommandBase` provides default implementations |
| **Registry Pattern** | Commands registered in a central registry |

## Execution Flow

```
main() → Main() → NewJujuCommand() → registerCommands()
    │
    ▼
cmd.Main(jcmd, ctx, args)
    │
    ├── Parse flags (SuperCommand)
    │
    ├── Resolve subcommand
    │   ├── Check registered commands
    │   ├── Check user aliases
    │   └── Invoke MissingCallback (plugin search)
    │
    ├── Init command arguments
    │
    └── Run(ctx)
        │
        ├── Connect to API (if needed)
        ├── Execute operation
        ├── Format output
        └── Return exit code
```

## Error Handling Model

The CLI uses a structured error handling approach:

| Exit Code | Meaning |
|-----------|---------|
| 0 | Success |
| 1 | Command error (displayed to user) |
| 2 | Initialization/parse error |
| N | Passthrough from plugins (RcPassthroughError) |

Special errors:
- `ErrSilent`: Exit 1, no error message
- `RcPassthroughError`: Propagate exit code from plugin

## Extensibility Points

1. **Plugin System**: Add `juju-<name>` executables to PATH
2. **User Aliases**: Define command aliases in `~/.local/share/juju/aliases`
3. **Feature Flags**: Enable experimental features via `JUJU_DEV_FEATURE_FLAGS`
4. **Embedded Commands**: Whitelisted commands available via controller API

## Dependencies

Key external dependencies:
- `github.com/juju/gnuflag`: Flag parsing
- `github.com/juju/loggo`: Logging
- `github.com/juju/errors`: Error wrapping
- `github.com/juju/collections`: Set utilities
- `github.com/juju/utils/v4`: General utilities
