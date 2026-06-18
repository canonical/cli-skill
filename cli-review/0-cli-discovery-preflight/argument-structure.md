# Argument Structure

The CLI uses long-form flags (no short aliases), with Cobra defaults (`--help`, whitespace/equals value forms). Global arguments on `todo` are transport/output oriented, while command-local arguments define object selection and mutation.

## Common patterns

- Positional selectors for object IDs (e.g. `<todo-id>`, `<sink-id>`, `<schedule-id>`)
- Long flags for optional mutation/filtering
- Repeatable flags for multivalue input (`--event`, `--sink`, `--schedule`)
- Optional output mode switches (`--format`, `--rfc3339`)

## `todo` global arguments

- `--host <string>` optional, default `127.0.0.1`
- `--port <int>` optional, default `44180`
- `--timeout <duration>` optional, default `10s`
- `--format <table|json>` optional, default `table`
- `--rfc3339` optional boolean, default `false`

## `todo` command arguments

### `list`
- `--state <open|closed|reopened|rejected>` optional

### `show <todo-id>`
- `<todo-id>` required positional

### `create <todo-id> <title>`
- `<todo-id>` required positional
- `<title>` required positional (minimum one token)
- `--due <datetime>` optional
- `--schedule <spec>` optional repeatable

### `update <todo-id>`
- `<todo-id>` required positional
- `--title <string>` optional
- `--due <datetime>` optional
- `--clear-due` optional boolean

### `close|reopen|reject <todo-id>`
- `<todo-id>` required positional

### `sinks`
- `--enabled <true|false>` optional
- `--event <upcoming|overdue>` optional

### `sink <sink-id>`
- `<sink-id>` required positional

### `create-sink <sink-id>`
- `<sink-id>` required positional
- `--url <url>` required by runtime validation
- `--event <upcoming|overdue>` optional repeatable

### `update-sink <sink-id>`
- `<sink-id>` required positional
- `--url <url>` optional
- `--event <upcoming|overdue>` optional repeatable
- `--clear-events` optional boolean

### `delete-sink|enable-sink|disable-sink <sink-id>`
- `<sink-id>` required positional

### `schedules`
- `--todo <todo-id>` optional
- `--kind <upcoming|overdue>` optional
- `--status <active|sent>` optional
- `--target <sink|motd>` optional

### `schedule <schedule-id>`
- `<schedule-id>` required positional

### `add-schedule <schedule-id>`
- `<schedule-id>` required positional
- `--todo <todo-id>` required by runtime validation
- `--kind <upcoming|overdue>` required by runtime validation
- `--before <duration>` optional
- `--every <duration>` optional
- `--sink <sink-id>` optional repeatable
- `--motd` optional boolean

### `remove-schedule <schedule-id>`
- `<schedule-id>` required positional

### `reminder-status`
- no command-specific args

### `status`
- no command-specific args

### `version`
- no command-specific args

## `todod` command arguments

### `start`
- `--host <string>` optional, default `127.0.0.1`
- `--port <int>` optional, default `44180`
- `--db <path>` optional (defaults to user data path)

### `stop`
- `--host <string>` optional
- `--port <int>` optional

### `status`
- `--host <string>` optional
- `--port <int>` optional
- `--format <table|json>` optional

### `version`
- no command-specific args

## Special arguments

- `create` uses `cobra.MinimumNArgs(2)` with free-form title captured from all tokens after `<todo-id>`.
- Inline schedule grammar in `--schedule` is custom and parsed manually (`kind[:before=...|every=...][:motd|sink=...]`).
- Schedule duration strings accept relative/human forms via shared parser logic.
- `--enabled` is modeled as a string on `sinks` (`"true"`/`"false"`), not a boolean flag.
