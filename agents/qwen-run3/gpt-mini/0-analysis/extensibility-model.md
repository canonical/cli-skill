# Extensibility Model

How new commands are added:
- The CLI uses Cobra with a `commands` package. New commands are functions that return `*cobra.Command` and are wired into `main.go` via `addCommandGroup` or `addCommands`.

Plugin/extension points:
- Engines and models are snap components. Adding a new engine is done by adding a component under `engines/` with a manifest; the CLI discovers engines by reading manifests from the `EnginesDir`.
- The snap `cli` binary is built from a Go module; additional functionality should be added as Go code in `cmd/cli/commands` or `pkg` packages.

Registration paths:
- Commands are registered centrally in `main.go`.
- Engine manifests loaded by `pkg/engines`; the `UseEngine`, `ListEngines` commands rely on that.

Middleware/hooks:
- PersistentPreRunE (`persistentPreRunE`) processes global flags like `--verbose`.
- Common helper functions in `cmd/cli/common` provide shared behavior (progress spinners, config helpers).

Notes about contributing:
- Follow Cobra patterns; keep new subcommands consistent with DE013 (verbs as commands) and avoid mixing short/long flag styles.
