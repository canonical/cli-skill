# Juju CLI Architecture

## Tech Stack

| Component | Technology |
|---|---|
| Language | Go (1.23+) |
| CLI Framework | Custom `cmd` package (`github.com/juju/juju/cmd/cmd`) |
| Flag Parsing | `github.com/juju/gnuflag` (GNU getopt-compatible) |
| API Transport | WebSocket over HTTPS |
| API Serialization | JSON-RPC |
| Client Store | YAML files in `~/.local/share/juju/` |
| Output Formatting | Custom `cmd.Output` with yaml, json, and tabular formatters |
| Logging | `github.com/juju/loggo/v3` |
| Interactive Shell | Custom REPL in `cmd/juju/commands/repl.go` |

## Architecture Style

**Primary style: Client-server CLI**

The Juju CLI is a thin client that communicates with a Juju controller via a WebSocket API. Nearly all commands perform RPC calls to a controller process. Local state is minimal and limited to client-side configuration (controller endpoints, credentials, model metadata).

**Secondary style: Plugin-based architecture**

The CLI supports external plugins via a missing-command callback: if a command is not found in the built-in registry, the CLI attempts to execute `juju-<command>` from `$PATH`. This allows third-party extensions without modifying core code.

**Additional characteristics:**

- **Supercommand pattern**: A single binary (`juju`) dispatches to ~120 flat subcommands. There is no nested subcommand hierarchy beyond the top level.
- **Layered CLI application**: Commands are organized in domain packages (`cmd/juju/<domain>/`) but flattened at runtime into a single namespace.
- **Command bus architecture**: The `registerCommands()` function in `cmd/juju/commands/main.go` acts as a manual command bus, wiring concrete command constructors into the supercommand registry.

## Key Components

### Command Framework (`cmd/cmd/`)

The custom `cmd` package provides:
- `Command` interface: `Info()`, `SetFlags()`, `Init()`, `Run()`
- `SuperCommand`: Container for subcommands with help topic generation
- `Context`: Execution context carrying stdin/stdout/stderr, working directory, and formatter state
- `Output`: Pluggable output formatter supporting yaml, json, and tabular output
- `MissingCallback`: Hook for unrecognized commands (used for plugin dispatch)

### Model Command Base (`cmd/modelcmd/`)

Most commands embed `ModelCommandBase`, which provides:
- Controller and model resolution from client store or flags (`-c`, `-m`)
- API root connection management
- WebSocket dialing with TLS and macaroon authentication
- Model type detection (IAAS vs CAAS)

### Client Store (`api/jujuclient/`)

YAML-backed local storage for:
- Controllers (`controllers.yaml`)
- Models (`models.yaml`)
- Accounts (`accounts.yaml`)
- Cookies (`cookies/`)

### Command Registration (`cmd/juju/commands/main.go`)

The `registerCommands()` function is the single source of truth for the command set. Commands are grouped by domain in the source code but registered into a flat namespace. There is no automatic discovery; every command must be manually registered.

### Plugin System

External plugins are resolved via `RunPlugin()` in the missing command callback. If `juju <cmd>` is not found, the CLI tries to execute `juju-<cmd>` from the user's PATH.

### Interactive Shell

Running `juju` with no arguments launches a REPL that delegates to the same command registry. The REPL is implemented in `cmd/juju/commands/repl.go`.

## Entry Points

| Binary | Path | Purpose |
|---|---|---|
| `juju` | `cmd/juju/main.go` | Primary CLI client |
| `jujud` | `cmd/jujud/main.go` | Controller/agent daemon |
| `juju-containeragent` | `cmd/containeragent/main_nix.go` | Kubernetes sidecar agent |

## Data Flow

```
User → juju CLI → SuperCommand → Command.Init() → ModelCommandBase.NewAPIRoot() → WebSocket → Controller API → State (MongoDB)
                              ↓
                        ClientStore (YAML)
```

## Notable Design Decisions

1. **Flat command namespace**: Despite ~120 commands, there are no nested subcommands (e.g., `juju cloud add` vs `juju add-cloud`). The CLI uses `verb-noun` and `noun` conventions rather than hierarchical grouping.

2. **Manual registration**: Every command must be explicitly registered in `registerCommands()`. There is no reflection-based or annotation-based discovery.

3. **Feature flags**: Some commands are gated behind `featureflag.DeveloperMode` (e.g., `dump-model`, `dump-db`).

4. **Embedded mode**: A subset of commands can run in "embedded" mode for the Juju Dashboard, controlled by a whitelist.

5. **User aliases**: The supercommand supports user-defined aliases from `~/.local/share/juju/aliases`.
