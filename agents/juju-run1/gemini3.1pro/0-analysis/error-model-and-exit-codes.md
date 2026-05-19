# Error Model and Exit Codes

## Exit Code Conventions

### CLI Binary (`qwen36`)

| Exit Code | Meaning | Context |
|-----------|---------|---------|
| 0 | Success | All commands on success |
| non-zero | Failure | Exact codes not documented; inferred from Go conventions (likely 1 for general errors) |

### Health Check Script (`check-server-llamacpp.sh`)

| Exit Code | Meaning | Context |
|-----------|---------|---------|
| 0 | Server healthy | llama-server process running, port open, completions API responding correctly |
| 1 | Server starting | Server is running but still loading model (HTTP 503 "Loading model") or port not yet open |
| 2 | Server failed | Process not running, empty API response, or API returned a non-recoverable error |

### Server Script (`server.sh`)

| Exit Code | Meaning | Context |
|-----------|---------|---------|
| 0 | Clean shutdown | Normal server exit |
| 1 | Timeout waiting for components | Required snap components not installed within 3600s |

### Wait-for-Server Script (`wait-for-server.sh`)

| Exit Code | Meaning | Context |
|-----------|---------|---------|
| 0 | Server ready | Health check returned 0 |
| 1 | Timeout or failure | Server didn't start within 60s, or health check returned 2 (failed) |

### Engine Server Scripts (`engines/*/server`)

| Exit Code | Meaning | Context |
|-----------|---------|---------|
| 1 | Missing component | Required component directory (`server`, `model`, `mmproj`) does not exist |
| *(passthrough)* | llama-server exit | Exit code from `exec llama-server ...` is passed through |

## Error Categories

### Missing Components

**Messages**:
```
Missing component: $model
Missing component: $mmproj
Missing component: $server
```

**Context**: Engine common.sh scripts check that snap component directories exist before attempting to launch the server. This occurs when components haven't been installed yet.

**Recovery**: Wait for component installation (`snap changes` to monitor progress).

### Component Installation Timeout

**Message**:
```
Error: timed out after 3600s while waiting for required components: [component1 component2]
Please use "snap changes" to monitor the progress and start the service once all components are installed.
```

**Context**: `server.sh` waits up to 1 hour for required components. On timeout, stops the snap service.

**Recovery**: Monitor `snap changes`, restart service after components install.

### Server Not Running

**Messages**:
```
Server is not running or failed. Please check the logs.
Timed out waiting for server to start. Check the server logs and try again.
```

**Context**: `wait-for-server.sh` reports these when the health check indicates failure or timeout.

### GPU Layers Not Configured

**Message**:
```
gpu-layers snap option is not set
```

**Context**: CUDA engine server script requires `gpu-layers` to be configured. Exits with code 1.

### Hardware Detection Unavailable

**Message**:
```
hardware-observe interface not auto connected. Skip auto engine selection.
```

**Context**: Install hook cannot auto-detect hardware without the `hardware-observe` interface. Non-fatal — engine must be selected manually.

## Error Output Channels

- **CLI errors**: Likely to stderr (standard Go convention), but unconfirmed due to private source
- **Script errors**: Mix of stdout and stderr; some scripts redirect stderr to syslog
- **Syslog integration**: Install and post-refresh hooks redirect both stdout and stderr to syslog via `logger`

## Unobserved Error Behaviors

Since the CLI source is private, the following are unknown:
- Error messages from `qwen36 get` with invalid keys
- Error messages from `qwen36 set` with invalid values
- Error format from `qwen36 use-engine` with invalid engine name
- Whether errors include suggestions or next-step guidance
- Whether `--help` is supported and what it outputs
