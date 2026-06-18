# Argument Structure

## Introduction and Common Patterns

The client `todo` utilizes consistent global formatting options across almost all query-based commands to configure formatting and time standards.
Additionally, commands that list plural objects (`sinks`, `schedules`, `list`) or manage individual ones support specific filtering or specification flags.

---

## Global and Formatting Flags

Almost all `todo` subcommands accept output formatting flags:

- `--format` (string, default: `"table"`, accepted: `"table"`, `"json"`) - Sets the output formatting style.
- `--rfc3339` (boolean, default: `false`) - Force dates to display in strict RFC3339 format instead of relative text.

---

## Command Argument Specification Map

### `todo` Commands

#### `todo list`
- Flags:
  - `--state` (string, optional, accepted: `"open"`, `"closed"`, `"reopened"`, `"rejected"`) - Filter todos by state.
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo show <todo-id>`
- Positional:
  - `<todo-id>` (string, required) - The ID of the todo to retrieve.
- Flags:
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo create <title>`
- Positional:
  - `<title>` (string/string array, required) - The title of the todo. Multiple space-separated arguments are joined into a single title.
- Flags:
  - `--due` (string, optional) - Optional due date, supports RFC3339 or human-readable formats.
  - `--schedule` (string array, optional, repeatable) - Schedule configuration specifications.
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo update <todo-id>`
- Positional:
  - `<todo-id>` (string, required) - The ID of the todo to update.
- Flags:
  - `--title` (string, optional) - New title.
  - `--due` (string, optional) - New due date.
  - `--clear-due` (boolean, optional, default: `false`) - Clear the existing due date.
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo close <todo-id>` / `todo reopen <todo-id>` / `todo reject <todo-id>`
- Positional:
  - `<todo-id>` (string, required) - The ID of the target todo.
- Flags:
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo sinks`
- Flags:
  - `--enabled` (string, optional, accepted: `"true"`, `"false"`) - Filter by enabled state.
  - `--event` (string, optional, accepted: `"upcoming"`, `"overdue"`) - Filter by subscription event.
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo sink <sink-id>`
- Positional:
  - `<sink-id>` (string, required) - The unique sink ID.
- Flags:
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo create-sink <sink-id>`
- Positional:
  - `<sink-id>` (string, required) - The unique sink ID to assign.
- Flags:
  - `--url` (string, required) - Webhook URL target.
  - `--event` (string array, optional, repeatable, accepted: `"upcoming"`, `"overdue"`) - Subscription event.
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo delete-sink <sink-id>`
- Positional:
  - `<sink-id>` (string, required) - ID of the sink to delete.

#### `todo schedules`
- Flags:
  - `--todo` (string, optional) - Filter by associated todo ID.
  - `--kind` (string, optional, accepted: `"upcoming"`, `"overdue"`) - Filter by schedule type.
  - `--status` (string, optional, accepted: `"active"`, `"sent"`) - Filter by schedule status.
  - `--target` (string, optional, accepted: `"sink"`, `"motd"`) - Filter by notification target.
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo schedule <schedule-id>`
- Positional:
  - `<schedule-id>` (string, required) - Unique schedule identifier.
- Flags:
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo add-schedule <schedule-id>`
- Positional:
  - `<schedule-id>` (string, required) - ID of the schedule to register.
- Flags:
  - `--todo` (string, required) - Associated todo ID.
  - `--kind` (string, required, accepted: `"upcoming"`, `"overdue"`) - Schedule kind.
  - `--before` (string, optional, default: `"24h"`) - Lead time offset before due date.
  - `--every` (string, optional, default: `"1d"`) - Recurrence frequency.
  - `--sink` (string array, optional, repeatable) - Associated target webhook sink ID(s).
  - `--motd` (boolean, optional, default: `false`) - Deliver notifications through MOTD.
  - `--format`, `--rfc3339` (Global format overrides)

#### `todo remove-schedule <schedule-id>`
- Positional:
  - `<schedule-id>` (string, required) - ID of the schedule to remove.

#### `todo status`
- Flags:
  - `--format` (string, default: `"table"`, accepted: `"table"`, `"json"`) - Sets the output format.

---

### `todod` Commands

#### `todod start`
- Flags:
  - `--db` (string, optional) - Custom SQLite database file path.

#### `todod status`
- Flags:
  - `--format` (string, default: `"table"`, accepted: `"table"`, `"json"`) - Sets the output format.

#### `todod stop`
- None.

---

## Special Arguments

- **`--` separator**: Double-dash separator is not explicitly intercepted for custom command execution by default, but it is supported by standard POSIX guidelines and Cobra's default argument processor.
