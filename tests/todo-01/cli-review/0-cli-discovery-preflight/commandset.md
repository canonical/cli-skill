# Command Set

## Tool: `todo`

### Primary-object commands (todos)

- `list-todos`
  - What: List todo items.
  - How: Reads optional `--state`; calls `client.ListTodos` -> daemon `GET /todos`.
- `show <todo-id>`
  - What: Show one todo.
  - How: Calls `client.ShowTodo` -> daemon `GET /todos/{id}`.
- `create <todo-id> <title>`
  - What: Create a todo, optional due date and inline schedules.
  - How: Parses `--due` and repeatable `--schedule`, builds `CreateTodoRequest`, calls `POST /todos`.
- `update <todo-id>`
  - What: Update title and/or due date.
  - How: Parses `--title`, `--due`, `--clear-due`, calls `PATCH /todos/{id}`.
- `close <todo-id>`
  - What: Transition todo to closed.
  - How: Calls `POST /todos/{id}/close`.
- `reopen <todo-id>`
  - What: Transition todo to reopened.
  - How: Calls `POST /todos/{id}/reopen`.
- `reject <todo-id>`
  - What: Transition todo to rejected.
  - How: Calls `POST /todos/{id}/reject`.

### Secondary-object commands (sinks)

- `sinks`
  - What: List sinks.
  - How: Reads `--enabled`, `--event`; calls `GET /sinks`.
- `sink <sink-id>`
  - What: Show one sink.
  - How: Calls `GET /sinks/{id}`.
- `create-sink <sink-id>`
  - What: Create sink.
  - How: Uses `--url`, repeatable `--event`; calls `POST /sinks`.
- `update-sink <sink-id>`
  - What: Update sink values.
  - How: Uses `--url`, `--event`, `--clear-events`; calls `PATCH /sinks/{id}`.
- `delete-sink <sink-id>`
  - What: Delete sink.
  - How: Calls `DELETE /sinks/{id}`.
- `enable-sink <sink-id>`
  - What: Enable sink.
  - How: Calls `POST /sinks/{id}/enable`.
- `disable-sink <sink-id>`
  - What: Disable sink.
  - How: Calls `POST /sinks/{id}/disable`.

### Secondary-object commands (schedules)

- `schedules`
  - What: List schedules.
  - How: Uses `--todo`, `--kind`, `--status`, `--target`; calls `GET /schedules`.
- `schedule <schedule-id>`
  - What: Show one schedule.
  - How: Calls `GET /schedules/{id}`.
- `add-schedule <schedule-id>`
  - What: Add immutable schedule.
  - How: Uses `--todo`, `--kind`, `--before`, `--every`, `--sink`, `--motd`; calls `POST /schedules`.
- `remove-schedule <schedule-id>`
  - What: Remove schedule.
  - How: Calls `DELETE /schedules/{id}`.

### State/info commands

- `reminder-status`
  - What: Print queued MOTD reminder messages.
  - How: Calls `GET /motd/message`.
- `status`
  - What: Show system status and MOTD setup hint.
  - How: Calls `GET /status`; renders table/json.
- `version`
  - What: Show client version.
  - How: Prints local constant.

## Tool: `todod`

- `start`
  - What: Start daemon.
  - How: Reads `--host`, `--port`, `--db`; constructs server and runs HTTP service.
- `stop`
  - What: Stop daemon.
  - How: Calls `POST /shutdown`.
- `status`
  - What: Show daemon status.
  - How: Calls `GET /status`; renders table/json.
- `version`
  - What: Show daemon version.
  - How: Prints local constant.
