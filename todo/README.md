# todo / todod

Task management CLI and daemon implemented in Go with Cobra and SQLite.

## Components

- `todo`: CLI client for task, schedule, sink, reminder-status, and status commands.
- `todod`: user-space daemon exposing an HTTP API and managing persistence/scheduling.

## Defaults

- Transport: HTTP over Unix socket at `/run/todod.socket`
- Database: `~/.local/share/todo/todo.db`
- Date input: RFC3339 and human-readable dates
- Date output: human-readable by default, with `--rfc3339` to force RFC3339 output

## Commands

### Todos (primary object)

- `todo list [--state ...]`
- `todo show <todo-id>`
- `todo create <title> [--due <date>] [--schedule <spec>...]`
- `todo update <todo-id> [--title ...] [--due ...] [--clear-due]`
- `todo close <todo-id>`
- `todo reopen <todo-id>`
- `todo reject <todo-id>`

### Sinks (secondary object, immutable)

- `todo sinks [--enabled true|false] [--event upcoming|overdue]`
- `todo sink <sink-id>`
- `todo create-sink <sink-id> --url <url> [--event ...]`
- `todo delete-sink <sink-id>`

### Schedules (secondary object, immutable)

- `todo schedules [--todo <id>] [--kind ...] [--status ...] [--target sink|motd]`
- `todo schedule <schedule-id>`
- `todo add-schedule <schedule-id> --todo <todo-id> --kind upcoming|overdue [--before <dur>] [--every <dur>] [--sink <id>...] [--motd]`
- `todo remove-schedule <schedule-id>`

### MOTD + status

- `todo reminder-status`
- `todo status`

When MOTD integration is not detected, `todo status` prints:

`To show todo reminders on login, run: echo 'todo reminder-status' >> ~/.profile`

### Daemon

- `todod start [--db ...]`
- `todod status [--format table|json]`
- `todod stop`
- `todod version`

## Schedule semantics

- `upcoming`: one-shot reminder before due date (`--before`, default `24h`)
- `overdue`: repeating reminder after due date (`--every`, default `1d`)
- targets: sink(s), MOTD, or both
- status: `active` or `sent`; upcoming schedules transition to `sent` when all targets are delivered
