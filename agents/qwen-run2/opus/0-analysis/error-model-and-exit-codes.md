# Error Model and Exit Codes

## Exit Code Map

| Exit Code | Meaning | Used By |
|-----------|---------|---------|
| 0 | Success | All commands on success |
| 1 | General error / failure | All commands on failure; server.sh on component timeout; wait-for-server.sh on timeout; cuda/server on missing gpu-layers |
| 2 | Server failed / fatal error | check-server-llamacpp.sh (server not running or unrecoverable) |
| *passthrough* | Upstream exit code | wait-for-server.sh passes unexpected check-server exit codes through |

## Error Categories

### Component Missing Errors

| Context | Message Pattern | Exit Code | Recovery |
|---------|----------------|-----------|----------|
| common.sh | `Missing component: $model` | 1 | Install the snap component |
| common.sh | `Missing component: $mmproj` | 1 | Install the snap component |
| common.sh | `Missing component: $server` | 1 | Install the snap component |
| server.sh | `Error: timed out after ${elapsed}s while waiting for required components: [${missing_components[*]}]` | 1 | Wait for snap component install to complete, then restart |

### Server Health Check Errors (check-server-llamacpp.sh)

| Condition | Message | Exit Code | Meaning |
|-----------|---------|-----------|---------|
| llama-server not running | `llama-server process is not running` (debug only) | 2 | Fatal — server must be restarted |
| Port not listening | `llama-server is not listening on the configured port` (debug only) | 1 | Transient — server still starting |
| Empty API response | `Empty response from server` (debug only) | 2 | Fatal — server hung or crashed |
| API error: loading model | `Loading model` (debug only) | 1 | Transient — model still loading |
| API error: other | `Server error: $error_message` | 2 | Fatal — server in error state |
| Empty chat text | `Empty chat response` (debug only) | 2 | Fatal — unexpected server behavior |

### Wait-for-Server Errors (wait-for-server.sh)

| Condition | Message | Exit Code |
|-----------|---------|-----------|
| 60s timeout elapsed | `Timed out waiting for server to start. Check the server logs and try again.` | 1 |
| Server not running (exit 2 from check) | `Server is not running or failed. Please check the logs.` | 1 |
| Unexpected check exit code | `Unexpected exit code when waiting for server: $exit_code` | *passthrough* |

### Configuration Errors

| Context | Condition | Message Pattern | Exit Code |
|---------|-----------|----------------|-----------|
| cuda/server | `gpu-layers` not set | `gpu-layers snap option is not set` | 1 |
| get (snapctl) | Invalid key | snapctl error message | non-zero |
| set (snapctl) | Invalid key/value | snapctl error message | non-zero |

### Engine Selection Errors

| Condition | Likely Message | Exit Code |
|-----------|----------------|-----------|
| No compatible engine found | (Go binary error, not in shell scripts) | 1 |
| Hardware detection fails | (Go binary error) | 1 |
| Invalid engine name | (Go binary error) | 1 |

## Per-Command Error Behavior

| Command | Error Sources | Stderr Used | Exit Code Range |
|---------|--------------|-------------|-----------------|
| `chat` | Server not running, timeout, network error | Yes | 0, 1 |
| `use-engine` | No compatible hardware, invalid engine name | Yes | 0, 1 |
| `show-engine` | No engine configured | Yes | 0, 1 |
| `get` | Invalid key, key not set | Yes (from snapctl) | 0, 1 |
| `set` | Invalid key/value format | Yes (from snapctl) | 0, 1 |
| `completion` | Invalid shell name | Yes | 0, 1 |
| `server` (daemon) | Component missing, timeout, server crash | Yes (to journal) | 0, 1, 2 |

## Observations

- **No custom exit codes**: The CLI uses only 0, 1, and 2. There is no distinction between user errors (bad input), system errors (missing hardware), and runtime errors (server crash) in exit codes.
- **Debug-gated messages**: Many diagnostic messages in `check-server-llamacpp.sh` are only shown when the `DEBUG` env var is set, making troubleshooting harder in normal operation.
- **No structured error output**: All errors are free-form strings on stderr. No JSON error format, no error codes in messages.
- **Inconsistent error guidance**: Some messages include recovery advice ("Check the server logs and try again"), while others give only the error condition.
