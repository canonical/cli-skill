# Error Model and Exit Codes

## Global Exit Behavior

The root command is configured with `SilenceUsage: true`, so usage text is **not** printed on runtime errors. Errors are returned from `RunE` handlers, printed to stderr by Cobra, and the process exits with status `1`.

In `main.go`:
```go
err := rootCmd.Execute()
if err != nil {
    os.Exit(1)
}
```

**All non-success exits are `1`**. There are no differentiated exit codes for specific error categories. This limits scriptability.

## Error Categories and Representative Messages

### Permission Errors
- **Trigger**: non-root user runs `set`, `unset`, `use-engine`, `prune-cache`
- **Message**: `permission denied, try again with sudo`
- **Source**: `common.ErrPermissionDenied`
- **Exit code**: `1`

### No Active Engine
- **Trigger**: commands that require an active engine when none is selected (`status`, `chat`, `webui`, `prune-cache`, `run`, `serve-webui`)
- **Message**: `no active engine`
- **Source**: `common.ErrNoActiveEngine`
- **Exit code**: `1`

### Config Key Not Found
- **Trigger**: `get <key>` on unknown key (unless passthrough), `unset <key>` on unknown key
- **Message**: `key "<key>" is not found` followed by a suggestion line
- **Exit code**: `1`

### Engine Manifest Errors
- **Trigger**: `show-engine <name>` with unknown engine, `use-engine <name>` with unknown engine, `prune-cache --engine=<name>` with unknown engine
- **Message**: `engine "<name>" does not exist` or `"<name>" not found`
- **Exit code**: `1`

### Component Installation Failures
- **Trigger**: `use-engine` or `fix` fails to install a snap component
- **Messages**:
  - `timed out while installing "<component>"`
  - `snap not known to the store`
  - `installing "<component>": <snapd error>`
- **Exit code**: `1`

### Server Unavailability (Chat/WebUI)
- **Trigger**: `chat` or `webui` when server is not active or endpoint unreachable
- **Messages**:
  - `server not active`
  - `connection refused` + suggestions
  - `no models available on server`
- **Exit code**: `1`

### Argument Validation
- **Trigger**: wrong number of args, mutually exclusive flags, invalid key-value syntax
- **Messages**:
  - `engine name not specified`
  - `cannot specify both engine name and --auto flag`
  - `expected key=value, got "..."`
  - `duplicate key: "..."`
- **Exit code**: `1`

### Unknown Format
- **Trigger**: `--format=<bad>` on commands with format flag
- **Message**: `unknown format "<bad>"`
- **Exit code**: `1`

### Internal / IO Errors
- **Trigger**: YAML/JSON marshal failure, snapctl communication failure, filesystem errors
- **Messages**: prefixed with the failing operation, e.g., `getting machine info: ...`, `scoring engines: ...`
- **Exit code**: `1`

## Per-Command Differences

| Command | Unique Error Patterns |
|---|---|
| `chat` | Connection refused → suggests logs and startup; model lookup timeout |
| `webui` | Service inactive → suggests `snap start`; `xdg-open` fork failure |
| `use-engine` | User cancellation during component install prompt → exits `0` (not an error) |
| `prune-cache` | Cannot prune active engine → explicit guard |
| `run` | Child process exit code is propagated directly via `exec.Cmd.Run()` — this is the **only** path where the CLI may exit with a code other than `0` or `1`, inheriting the subprocess's status. |
| `debug select-engine` | Pipe errors (`decoding hardware info`) when stdin is missing or invalid |

## Missing Exit Code Differentiation

Because all errors return `1`, scripts cannot reliably distinguish between:
- permission denied (should retry with sudo)
- no active engine (should select one)
- transient network/service failure (should retry)
- invalid user input (should fix command)

**Recommendation**: introduce distinct exit codes (e.g., `2` for permission, `3` for missing engine, `4` for server unavailable, `126` for subprocess exec failure) to improve automation robustness.
