# Error Model and Exit Codes

## Overview

The `qwen36` CLI uses a **single exit code convention**: `0` for success, `1` for all errors. There is no differentiation between error categories via exit codes (no usage of exit code `2` for argument errors, no exit code `127` for command-not-found, etc.).

The `main()` function in `cmd/cli/main.go` unconditionally calls `os.Exit(1)` when `rootCmd.Execute()` returns an error:
```go
err := rootCmd.Execute()
if err != nil {
    os.Exit(1)
}
```

All error messages go to **stderr** (via Cobra's default behavior or explicit `fmt.Fprintf(os.Stderr, ...)` calls).

## Error Categories

### 1. Permission Errors

| Pattern | Message | Exit Code |
|---------|---------|-----------|
| Non-root user runs privileged command | `permission denied, try again with sudo` | 1 |

**Affected commands**: `set`, `unset`, `use-engine`, `prune-cache`

The `ErrPermissionDenied` sentinel error is defined in `common/errors.go` and checked by each command that requires root via `utils.IsRootUser()`.

### 2. Engine/State Errors

| Pattern | Message | Exit Code |
|---------|---------|-----------|
| No active engine set | `no active engine` | 1 |
| Engine not found | `"<name>" not found` + optional verbose detail | 1 |
| Active engine manifest not found | `active engine manifest not found` | 1 |
| Attempt to prune active engine | `cannot prune the active engine "<name>"` | 1 |

The `ErrNoActiveEngine` sentinel is returned when `Cache.GetActiveEngine()` returns an empty string.

### 3. Configuration Errors

| Pattern | Message | Exit Code |
|---------|---------|-----------|
| Key not found (`get`) | `key "<key>" is not found` + suggestion to run `get` | 1 |
| Key not found (`set` — non-passthrough) | `key "<key>" is not found` + suggestion | 1 |
| Key not found (`unset`) | `key "<key>" is not found` + suggestion | 1 |
| Invalid key=value format | `expected key=value, got "<input>"` | 1 |
| Key starts with `=` | `key must not start with an equal sign` | 1 |
| Duplicate key in `set` | `duplicate key: "<key>"` | 1 |
| Reading/writing config failure | `getting value of "<key>": <err>` | 1 |

### 4. Argument/Flag Errors

| Pattern | Message | Exit Code |
|---------|---------|-----------|
| Missing required argument | Cobra-generated (e.g., `accepts 1 arg(s), received 0`) | 1 |
| `--auto` with engine name | `cannot specify both engine name and --auto flag` | 1 |
| `--fix` with engine name | `cannot specify both engine name and --fix flag` | 1 |
| `use-engine` without name | `engine name not specified` | 1 |
| Unknown format | `unknown format "<fmt>"` | 1 |
| `run` without command | `unexpected number of arguments, expected at least 1 got 0` | 1 |
| Too many args to `show-engine` | `invalid number of arguments` | 1 |

### 5. Service/Infrastructure Errors

| Pattern | Message | Exit Code |
|---------|---------|-----------|
| Server not active | `server not active` + suggestion to start | 1 |
| WebUI server not active | `<service> not active` + suggestion to start | 1 |
| Service not found | `<service>: service not found` | 1 |
| Cannot retrieve services | `Error: retrieving snap services: <err>` | 1 (from main) |
| OpenAI endpoint not found | `"openai" not found in server endpoints` | 1 |
| `SNAP_COMPONENTS` env var not set | `SNAP_COMPONENTS env var not set` | 1 |
| Component not installed | Timeout or install errors with descriptive messages | 1 |

### 6. Hardware Probe Warnings (Non-Fatal)

| Pattern | Message | Exit Code |
|---------|---------|-----------|
| `clinfo` not available | `Warning: <msg>` on stderr (only in `--verbose`) | 0 |
| Cannot query component sizes | `Warning: unable to query component sizes: <err>` on stderr (only in `--verbose`) | 0 |

Warnings are printed only when `--verbose` is set and do not cause non-zero exit.

### 7. Wrapped/Chained Errors

The CLI uses Go's `fmt.Errorf("context: %w", err)` extensively for error wrapping. Inner errors from the domain packages (`pkg/engines`, `pkg/storage`, `pkg/hardware_info`) are surfaced with contextual prefixes. Examples:
- `looking up active engine: <err>`
- `scoring engines: <err>`
- `loading engine manifest: <err>`
- `waiting for component: <err>`
- `getting OpenAI base URL: <err>`

### 8. Interactive Cancellation

When a user responds "no" to a confirmation prompt:
- `Cancelled. No changes applied.` (printed to stdout)
- Exit code **0** (not an error — user intentionally cancelled)

This applies to `use-engine` (component installation), `prune-cache`, and restart prompts.

## Per-Command Error Differences

| Command | Unique Error Patterns |
|---------|----------------------|
| `use-engine --fix` | If `ErrNoActiveEngine`, returns `nil` (success) — "nothing to fix" is not an error |
| `use-engine --auto` | Prints incompatibility reasons per-engine when `--verbose` |
| `prune-cache` | Single-engine vs all-engines mode uses different confirmation prompts; `--engine` with non-existent name yields `"<name>" not found` |
| `run` | Deprecated `--wait-for-components` flag marked via `cobraCmd.Flags().MarkDeprecated()` — using it prints a deprecation warning but does not cause error |
| `debug validate-engines` | Returns error `"not all manifests are valid"` only after checking all manifests (validates all, reports all failures) |
| `debug chat` | Explicitly checks `--base-url` is non-empty; returns `"the --base-url parameter is required"` |
| `debug select-engine` | Reads from stdin — YAML decode errors produce `"decoding hardware info: <err>"` |

## Exit Code Summary Table

| Exit Code | Meaning | When |
|-----------|---------|------|
| `0` | Success | Command completed successfully, or user cancelled an interactive prompt |
| `1` | Error | Any error — permission, argument, runtime, engine, config, service, etc. |

## Notable Gaps

### No exit code distinctions

All errors return exit code 1. There is no differentiation between:
- Argument errors (conventional exit code 2 in many tools)
- Runtime/operational errors (exit code 1)
- Configuration errors

Scripts cannot distinguish between "you typed the command wrong" and "the server is down" based on exit code alone.

### No exit code for "no active engine" as success

`status` returns exit code 1 with `no active engine` when no engine is configured. This could be argued as a valid state that should return exit code 0 with a descriptive message (like `git status` on an uninitialized repo) — especially since `use-engine --fix` treats `ErrNoActiveEngine` as success.

### Deprecation warnings use MarkDeprecated

The `--wait-for-components` flag on `run` is deprecated via Cobra's `MarkDeprecated()`. When used, Cobra prints: `Flag --wait-for-components has been deprecated, "run" always waits for components.` This goes to stderr. The command still executes normally (exit code 0 if the subprocess succeeds).

### No `--quiet` mode

There is no way to suppress informational output. All commands print to stdout/stderr unconditionally (except warnings which require `--verbose`). The `list-engines` command prints `"No engines found."` to stderr in the empty state.
