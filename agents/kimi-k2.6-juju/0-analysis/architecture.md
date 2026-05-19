# Architecture

## Tech Stack
- **Language**: Go (>=1.22 inferred from `go.mod`).
- **Build System**: Make + Go modules.
- **CLI Framework**: Custom `cmd` package (`github.com/juju/juju/cmd/cmd`) built on `github.com/juju/gnuflag`, providing a `SuperCommand` dispatcher and plugin fallback.
- **Client-Server**: The `juju` CLI is a thick client that communicates with a Juju controller via HTTP/JSON API (api package).
- **Persistence**: Client-side state stored in `~/.local/share/juju/` (controllers.yaml, models.yaml, accounts.yaml, cookies).
- **Embedded Docs**: Sphinx-based documentation in `docs/`, plus auto-generated command reference via `juju documentation`.

## Architecture Style
- **Primary**: **Client-server CLI** — the CLI translates user intent into API calls against a controller.
- **Secondary**: **Plugin-based architecture** — unknown commands fall back to `juju-<command>` executables on `$PATH`.
- **Additional**: **Layered CLI application** — clear separation between `cmd/juju` (presentation), `api` (client transport), `apiserver` (server), and `domain` (business logic). The CLI layer also uses a `SuperCommand` registry (microkernel command host pattern) where commands self-register via `registerCommands`.

## Key Components
| Layer | Path | Responsibility |
|---|---|---|
| CLI Entry | `cmd/juju/main.go` | Bootstrap logging, proxy setup, REPL detection |
| Command Registry | `cmd/juju/commands/main.go` | `registerCommands` maps every top-level command |
| SuperCommand | `cmd/cmd/supercommand.go` | Dispatch, aliases, deprecation, help, plugin fallback |
| API Client | `api/` | HTTP client, facades, authentication |
| Client Store | `api/jujuclient` | File-backed controller/model/account state |
| Plugins | `cmd/juju/commands/plugin.go` | PATH discovery, passthrough of args/env |
