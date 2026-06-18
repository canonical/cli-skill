# Command Set

Full list of CLI commands and hierarchy.

## Client CLI (`todo`)

### Todos Commands
- **`todo list`**
  - **Short Description**: List todos.
  - **How it Works**: Fetches the list of todos from the daemon client via `ListTodos(ctx, state)`. Supports filtering by state and displays the results in a tabular format (ID, STATE, DUE, TITLE) or JSON.
- **`todo show <todo-id>`**
  - **Short Description**: Show details of a specific todo.
  - **How it Works**: Fetches the specified todo via `ShowTodo(ctx, id)` and prints it as a key-value detail or JSON.
- **`todo create <title>`**
  - **Short Description**: Create a new todo.
  - **How it Works**: Parses optional due date and inline schedules from flags, then calls `CreateTodo(ctx, req)` on the client.
- **`todo update <todo-id>`**
  - **Short Description**: Update a todo.
  - **How it Works**: Parses title, due date, or clear-due changes and calls `UpdateTodo(ctx, id, req)`.
- **`todo close <todo-id>`**
  - **Short Description**: Close a todo.
  - **How it Works**: Marks a todo as closed using `CloseTodo(ctx, id)`.
- **`todo reopen <todo-id>`**
  - **Short Description**: Reopen a closed or rejected todo.
  - **How it Works**: Reopens a todo using `ReopenTodo(ctx, id)`.
- **`todo reject <todo-id>`**
  - **Short Description**: Reject a todo.
  - **How it Works**: Marks a todo as rejected using `RejectTodo(ctx, id)`.

### Sinks Commands
- **`todo sinks`**
  - **Short Description**: List sinks.
  - **How it Works**: Lists configured sinks via `ListSinks(ctx, enabled, event)`.
- **`todo sink <sink-id>`**
  - **Short Description**: Show a sink.
  - **How it Works**: Shows details of a specific sink via `ShowSink(ctx, id)`.
- **`todo create-sink <sink-id>`**
  - **Short Description**: Create a sink.
  - **How it Works**: Creates a new webhook sink using `CreateSink(ctx, req)`.
- **`todo delete-sink <sink-id>`**
  - **Short Description**: Delete a sink.
  - **How it Works**: Deletes a sink via `DeleteSink(ctx, id)`.

### Schedules Commands
- **`todo schedules`**
  - **Short Description**: List schedules.
  - **How it Works**: Lists configured reminder schedules via `ListSchedules(ctx, todo, kind, status, target)`.
- **`todo schedule <schedule-id>`**
  - **Short Description**: Show a schedule.
  - **How it Works**: Shows details of a specific schedule via `ShowSchedule(ctx, id)`.
- **`todo add-schedule <schedule-id>`**
  - **Short Description**: Add an immutable schedule.
  - **How it Works**: Adds a reminder schedule connected to a todo using `AddSchedule(ctx, req)`.
- **`todo remove-schedule <schedule-id>`**
  - **Short Description**: Remove a schedule.
  - **How it Works**: Removes a schedule via `RemoveSchedule(ctx, id)`.

### Other Commands
- **`todo status`**
  - **Short Description**: Show todo system status.
  - **How it Works**: Queries system status via `Status(ctx)` and formats the result as a text block or JSON.
- **`todo reminder-status`**
  - **Short Description**: Print pending MOTD reminder messages.
  - **How it Works**: Fetches MOTD reminder messages via `MOTDMessage(ctx)` and prints them.
- **`todo version`**
  - **Short Description**: Show client version.
  - **How it Works**: Prints the constant version of the client.
- **`todo todos`**
  - **Short Description**: Concepts and help topic about what todos mean in this app.
  - **How it Works**: Explains the core todo models and concept hierarchies.

---

## Daemon CLI (`todod`)

- **`todod start`**
  - **Short Description**: Start the daemon.
  - **How it Works**: Initializes and runs the daemon SQLite server listening on a UNIX socket.
- **`todod stop`**
  - **Short Description**: Stop the daemon.
  - **How it Works**: Sends a POST request to `/shutdown` of the running daemon.
- **`todod status`**
  - **Short Description**: Show daemon status.
  - **How it Works**: Sends a GET request to `/status` of the running daemon.
- **`todod version`**
  - **Short Description**: Show daemon version.
  - **How it Works**: Prints the version of the daemon.
