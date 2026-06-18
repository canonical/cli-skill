# Architecture

## Stack Summary

- Language: Go
- CLI framework: Cobra
- Transport: HTTP (localhost daemon API)
- Persistence: SQLite
- Binaries: `todo` (client CLI), `todod` (daemon CLI)

## Primary Architecture Style

Layered CLI application.

- Presentation layer: Cobra command tree in `cmd/todo/main.go` and `cmd/todod/main.go`
- Service layer: daemon logic in `internal/daemon/service.go`
- Persistence layer: SQLite-backed store in `internal/store/store.go`

## Secondary Style

Client-server CLI.

- `todo` is an HTTP client to `todod`
- `todod` hosts routes (`/todos`, `/sinks`, `/schedules`, `/status`, `/motd/message`) and runs scheduler loop

## Runtime Flow

1. User runs `todo <command>`
2. Command validates args/flags and sends HTTP request to daemon
3. Daemon service applies domain rules and writes to SQLite
4. Scheduler loop dispatches reminders to MOTD queue and/or sinks
