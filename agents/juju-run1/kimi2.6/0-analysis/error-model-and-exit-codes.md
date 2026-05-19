# Juju CLI Error Model and Exit Codes

## Error categories

1. Argument and parsing errors
- Unknown/invalid flags, extra args, and malformed usage.
- Typically surfaced via command init/flag parsing errors.

2. Command resolution errors
- Unrecognized command names.
- Near-match suggestions may be provided.
- Missing command may be delegated to plugin lookup before final failure.

3. Environment/config resolution errors
- Missing or conflicting controller/model context.
- Missing local store/config state.

4. Runtime execution errors
- API failures, permission failures, transport/connectivity issues.
- Command-specific domain validation and state errors.

5. User-aborted operations
- Confirmation prompts can return explicit aborted errors.

## Exit code behavior

Framework-level behavior observed:

- `0`: successful command execution.
- `1`: runtime error returned from `Run` (non-silent, non-pass-through).
- `2`: command setup/init/help/flag related errors and silent error mode in command handling path.
- Passthrough: commands/plugins may return wrapped exit codes (`RcPassthroughError`), which are propagated as-is.

## Per-command/group differences

- Plugin commands (`juju-<name>`) pass through plugin process exit codes.
- Internal runner commands and some command families intentionally use passthrough error codes to preserve remote/underlying semantics.

## Representative messages

- Unrecognized command: framework default unrecognized command errors.
- Not-found with suggestion: custom message includes nearest known command and help hint.
- Controller/env conflict: explicit message when `JUJU_MODEL` and `JUJU_CONTROLLER` imply different controllers.

## Reliability implications

- Scripted callers should key behavior on exit code first, then parse machine-formatted stdout when available.
- For plugin flows, treat non-zero codes as command-defined rather than framework-normalized.

## Evidence pointers

- Main command lifecycle and exit mapping: `cmd/cmd/cmd.go`
- Super-command missing handling and error emission: `cmd/cmd/supercommand.go`
- Plugin exit code passthrough: `cmd/juju/commands/plugin.go`
- Confirmation abort helper: `cmd/helpers.go`
