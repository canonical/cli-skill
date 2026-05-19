# Juju CLI Architecture Analysis

## Technology Stack Summary

Juju is a comprehensive application orchestration platform written in **Go**. The CLI is a key component of the Juju ecosystem, providing the primary user interface for managing cloud infrastructure, Kubernetes clusters, applications, and their lifecycles.

### Core Technologies

| Component | Technology |
|-----------|------------|
| Language | Go 1.22+ |
| CLI Framework | Custom `cmd` package using `github.com/juju/gnuflag` |
| Configuration | YAML-based, XDG-compliant (`~/.local/share/juju/`) |
| API Communication | JSON-RPC over HTTPS |
| Logging | `github.com/juju/loggo/v3` |
| Testing | `github.com/juju/tc` (test checkers) |

### Key Dependencies

- `github.com/juju/gnuflag` - Flag parsing (GNU-style)
- `github.com/juju/errors` - Error handling with stack traces
- `github.com/juju/loggo/v3` - Structured logging
- `github.com/juju/collections` - Set and collection utilities
- `github.com/juju/names/v6` - Entity naming conventions

## Architecture Style

**Primary Style: Client-Server CLI**

The Juju CLI operates as a client that communicates with one or more Juju controllers. Each controller manages multiple models, and the CLI provides a unified interface to interact with this distributed system.

**Secondary Style: Layered CLI Application**

The CLI follows a layered architecture with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                      cmd/juju (CLI Entry)                    │
│  - main.go: Entry point                                      │
│  - commands/main.go: Command registration and orchestration  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    cmd/juju/<domain>/                        │
│  - action/, application/, model/, user/, cloud/, etc.       │
│  - Command implementations with business logic orchestration │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    cmd/modelcmd/                             │
│  - ModelCommandBase: Model-aware command base               │
│  - ControllerCommandBase: Controller-aware command base     │
│  - API connection management                                 │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                        api/                                  │
│  - API client implementations                               │
│  - api/client/: Typed clients (application, model, etc.)    │
│  - api/jujuclient/: Client store for controller/model state │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   RPC over HTTPS                             │
│  - JSON-RPC protocol                                        │
│  - Macaroon-based authentication                            │
└─────────────────────────────────────────────────────────────┘
```

## Architectural Layers

### 1. Command Layer (`cmd/juju/`)

Each subdirectory represents a domain:
- `action/` - Action execution and management
- `application/` - Application lifecycle (deploy, remove, scale)
- `model/` - Model configuration and management
- `user/` - User authentication and authorization
- `cloud/` - Cloud and credential management
- `controller/` - Controller lifecycle
- `storage/` - Storage management
- `secrets/` - Secret management
- `space/` - Network space management
- `ssh/` - SSH and SCP functionality

### 2. Base Command Layer (`cmd/modelcmd/`, `cmd/cmd/`)

- `modelcmd.ModelCommandBase`: Provides model context (current model, controller)
- `cmd.CommandBase`: Base implementation for all commands
- `cmd.SuperCommand`: Implements command hierarchy and delegation

### 3. API Client Layer (`api/`)

- Typed clients for each API facade
- Connection management and authentication
- Client store for persistent controller/model information

### 4. Core Layer (`core/`)

Shared internal logic including:
- Entity types (application, unit, machine, model)
- Life cycle management
- Constraints and configuration types

### 5. Domain Layer (`domain/`)

Business workflows implemented as services:
- State persistence behind interfaces
- Transactional operations

## Command Execution Flow

```
User Input → main.go → commands.Main()
                              │
                              ▼
                    NewJujuCommand() → SuperCommand
                              │
                              ▼
                    SuperCommand.Init()
                              │
                              ▼
                    Command.Init(args)
                              │
                              ▼
                    Command.Run(ctx)
                              │
                              ▼
                    API Client → Controller
```

## Key Design Patterns

### 1. SuperCommand Pattern

The `SuperCommand` acts as a command router, delegating to subcommands:
- Supports nested command hierarchies
- Provides global flags to all subcommands
- Implements help and documentation generation

### 2. Model-Aware Commands

Commands that operate on models inherit from `ModelCommandBase`:
- Automatic model resolution (current, specified, or default)
- Controller context management
- API connection establishment

### 3. Plugin Architecture

External plugins discovered via PATH:
- Plugins named `juju-*` are automatically discovered
- Full argument pass-through
- Environment variable context injection (JUJU_MODEL, JUJU_CONTROLLER)

### 4. API Facade Pattern

Each API client corresponds to a server-side facade:
- Thin client wrappers
- Type-safe API calls
- Automatic serialization/deserialization

## Configuration Architecture

### Client-Side Storage

```
~/.local/share/juju/
├── accounts.yaml         # Controller account details
├── controllers.yaml      # Controller definitions
├── credentials.yaml      # Cloud credentials
├── models.yaml           # Model mappings
├── clouds.yaml           # User-defined clouds
├── aliases               # Command aliases
└── ssh/                  # SSH keys
```

### Configuration Precedence

1. Command-line flags (highest priority)
2. Environment variables (JUJU_*)
3. Configuration files
4. Defaults (lowest priority)

## Connection Model

### Multi-Controller Support

The CLI can manage multiple controllers simultaneously:
- Each controller has isolated credentials and models
- `switch` command changes active controller/model
- `JUJU_CONTROLLER` and `JUJU_MODEL` environment variables override

### Authentication

- Macaroon-based authentication
- Browser-based login for interactive sessions
- Token-based authentication for automation

## Interactive Mode

When run without arguments, Juju enters an interactive REPL:
- Command history
- Auto-completion
- Context-aware help

## Output Formatting

Multiple output formats supported:
- `tabular`: Human-readable tables (default)
- `yaml`: YAML format for scripting
- `json`: JSON format for programmatic consumption
- `line`/`oneline`: Compact single-line output

## Logging Architecture

- Hierarchical logging via `loggo`
- Configurable log levels per module
- File and stderr output options
- Debug mode for verbose diagnostics

## Error Handling

- Structured errors with stack traces
- Silent errors for expected failures
- Passthrough exit codes for plugins
- Contextual error wrapping

## Threading and Concurrency

- Context propagation for cancellation
- Worker patterns for background operations
- No unmanaged goroutines in command implementations
