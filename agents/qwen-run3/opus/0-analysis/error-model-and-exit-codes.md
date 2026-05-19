# Error Model and Exit Codes

## Exit Code Conventions

| Exit Code | Meaning | Used By |
|-----------|---------|---------|
| 0 | Success | All commands |
| 1 | General error | All commands (cobra default), server scripts |
| 2 | Server failed / unrecoverable | `check-server-llamacpp.sh` (server process died or bad response) |

The CLI uses cobra's default behavior: if `RunE` returns a non-nil error, cobra prints the error to stderr and `main()` calls `os.Exit(1)`. There is no differentiated exit code scheme beyond 0/1 for the Go CLI itself.

The shell scripts (`check-server-llamacpp.sh`) use a three-tier scheme:
- `0`: Server healthy
- `1`: Server still starting (retry)
- `2`: Server failed (do not retry)

## Error Categories

### Permission Errors

```
Error: permission denied. Please re-run with sudo.
```

Triggered by: `set`, `unset`, `use-engine`, `prune-cache` when not running as root. Uses a shared `common.ErrPermissionDenied` sentinel.

### Configuration Errors

```
Error: key "foo" is not found
```

Triggered by: `get`, `unset` when the key doesn't exist. Includes a suggestion hint about available keys.

```
Error: getting value of "foo": <underlying error>
```

Triggered by: `get` when the config store read fails.

### Engine Errors

```
Error: engine "foo" does not exist
```

Triggered by: `show-engine`, `use-engine` with an invalid engine name.

```
Error: no active engine. Use "qwen36 use-engine" to select one.
```

Triggered by: `show-engine` (no arg), `prune-cache` when no engine has been configured.

```
Error: cannot specify both engine name and --auto flag
```

Triggered by: `use-engine` with mutually exclusive inputs.

### Service Errors

```
Error: server not active

  Start the server with: sudo snap start qwen36.server
```

Triggered by: `chat` when the inference server is not running. Includes next-step guidance.

### Component/Timeout Errors

```
Error: timed out after 3600s while waiting for required components: [llamacpp model-qwen36-35b-a3b-ud-q4-k-xl]
```

Triggered by: `server.sh` when snap components are not installed within the timeout.

### Server Health Check Errors

```
Server error: <message from llama-server>
```

Triggered by: `check-server-llamacpp.sh` when the server returns an error response.

## Error Output Channel

All errors are written to **stderr** (cobra convention). The CLI uses `SilenceUsage: true` on the root command, preventing cobra from printing usage on errors.

## Error Recovery Guidance

| Error | User Action |
|-------|-------------|
| Permission denied | Re-run with `sudo` |
| No active engine | Run `qwen36 use-engine --auto` |
| Server not active | Run `sudo snap start qwen36.server` |
| Component timeout | Check `snap changes` and restart once components arrive |
| Key not found | Check available keys with `qwen36 get` |

## Gaps

1. **No structured error format**: Errors are plain text. There is no `--format json` for error output.
2. **No differentiated exit codes**: All Go CLI errors exit with code 1 regardless of category. Per POSIX and Canonical conventions, usage errors should use exit code 2.
3. **No error codes/IDs**: Errors are identified only by message text, making automated error handling fragile.
