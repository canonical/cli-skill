# Extensibility Model

## Plugin System
- **Discovery**: On startup, if a command is unrecognized, the CLI searches `$PATH` for an executable named `juju-<command>`.
- **Invocation**: `RunPlugin` extracts known Juju flags (`--model`, `--controller`) and passes the remaining args to the plugin.
- **Environment Injection**: `JUJU_CONTROLLER` and `JUJU_MODEL` are injected into the plugin environment based on parsed flags.
- **Exit-code Passthrough**: Plugin exit codes are preserved via `utils.RcPassthroughError`.
- **Description**: Plugins can implement `--description` to appear in `juju help plugins`.

## User Aliases
- **File**: `~/.local/share/juju/aliases`
- **Format**: `alias = command [args...]`
- **Processing**: Aliases are expanded before subcommand lookup in `SuperCommand.Init`.
- **Disable**: `--no-alias` skips alias expansion.

## Command Registration
- **Path**: `cmd/juju/commands/main.go` → `registerCommands`.
- **Method**: Commands implement `cmd.Command` (Info, SetFlags, Init, Run) and are registered via `r.Register(...)`.
- **Deprecation**: `RegisterDeprecated` and `RegisterAlias` support phased removal.
- **Embedded Whitelist**: Only commands in `allowedEmbeddedCommands` are exposed to the Juju Dashboard; others are excluded with a clear error.

## Middleware / Hooks
- **NotifyRun**: Called before every command execution; used for version logging.
- **NotifyHelp**: Called before help output; used to load plugin descriptions dynamically.
- **MissingCallback**: Provides fuzzy-match suggestions (`FindClosestSubCommand`) and plugin fallback.
- **Logging**: `cmd.Log` wraps `loggo`, configured via env var and flags.

## Extension Boundaries
- The CLI does not support dynamically loaded Go modules or shared libraries.
- Extensions must be external executables (plugins) or patches to `registerCommands`.
- No hook system exists for pre/post command interception beyond `NotifyRun`.

## Recommendations
- Expose a formal `juju plugin install` mechanism instead of relying on raw `$PATH` management.
- Consider a `juju hook` system for pre-command validation or audit logging.
