# Error Model And Exit Codes

## Summary

The checked-in shell layer exposes a clearer error model than the missing Go CLI source does. The most concrete exit-code contract in the repository is the engine health check protocol: `0` means healthy, `1` means still starting, and `2` means failed. User-facing CLI command exit codes for invalid syntax or validation failures cannot be fully verified without the private Go submodule, so those are documented as unknown or inferred rather than asserted.

## Error categories

| Category | Typical source | Representative symptom | Known exit code behavior |
|---|---|---|---|
| Invalid or missing config value | `qwen36 get`, engine launchers | missing key, empty required key, malformed effective runtime config | non-zero for command failures is likely; explicit `1` for shell launch errors |
| Engine startup not ready yet | `check-server-llamacpp.sh` | server process exists but model still loading or port not yet open | explicit `1` |
| Engine startup failed | `check-server-llamacpp.sh`, engine launchers | missing process, empty API response, non-loading API error, missing component | explicit `2` in health check; explicit `1` in launch scripts |
| Chat blocked by daemon unavailability | `wait-for-server.sh` | timeout or failed health check | explicit `1` |
| Component installation incomplete | `apps/server.sh` | required snap components not yet present | explicit `1` after timeout |
| Command-line validation failure | missing Go CLI source | invalid engine name, invalid key, malformed arguments | non-zero expected but exact codes/messages unverified |

## Per-command error model

| Command | Error / failure condition | Representative message or behavior | Exit code |
|---|---|---|---|
| `qwen36 chat` | server not ready within 60 seconds | `Timed out waiting for server to start. Check the server logs and try again.` | `1` from `wait-for-server.sh` |
| `qwen36 chat` | server not running or health check reports failure | `Server is not running or failed. Please check the logs.` | `1` from `wait-for-server.sh` |
| `qwen36 use-engine` | unsupported engine / incompatible hardware / validation error | not directly visible in checked-in source | unknown, but should be non-zero |
| `qwen36 show-engine` | no engine selected / unreadable manifest / CLI validation error | not directly visible in checked-in source | unknown, but should be non-zero |
| `qwen36 get <key>` | unknown or unset key | `model-name` reads are wrapped with `2>/dev/null || true`, implying stderr + non-zero are possible | unknown from Go source; callers expect possible failure |
| `qwen36 set <key>=<value>` | malformed assignment, invalid key, invalid value | not directly visible in checked-in source | unknown, but should be non-zero |
| `qwen36 completion bash` | unsupported shell target / CLI validation error | stderr is suppressed by the completer script | unknown from Go source |
| `qwen36.server` | missing required snap components after wait deadline | `Error: timed out after ... while waiting for required components: [...]` | `1` |
| `qwen36.server` | missing model, mmproj, or server component directory | `Missing component: ...` | `1` |
| `qwen36.server` | CUDA selected but `gpu-layers` unset | `gpu-layers snap option is not set` | `1` |

## Explicit exit-code contracts in shell scripts

### Health check contract

`apps/check-server-llamacpp.sh` defines a three-state protocol:

| Exit code | Meaning | Condition |
|---|---|---|
| `0` | server is healthy | process exists, port is open, API returns a non-error completion response |
| `1` | server is still starting | port closed but process exists, or API returns `{"error":{"message":"Loading model"...}}` |
| `2` | server failed or is unavailable | process missing, empty API response, unexpected API error, or invalid completion payload |

Representative messages:

- debug only: `llama-server process is not running`
- debug only: `llama-server is not listening on the configured port`
- debug only: `Loading model`
- stderr on real server error: `Server error: <message>`

### Wait wrapper contract

`apps/wait-for-server.sh` consumes the health-check protocol and maps it into user-facing chat gating:

| Health result | User-facing behavior | Exit code |
|---|---|---|
| `0` | proceed without message, or after printing a newline if dots were shown | `0` |
| `1` repeatedly until timeout | prints `Waiting for server ...` followed by dots | still running |
| timeout after repeated `1` | prints timeout guidance | `1` |
| `2` | prints failure guidance immediately | `1` |
| unexpected code | prints `Unexpected exit code when waiting for server: <n>` | propagates unexpected code |

### Daemon wait-for-components contract

`apps/server.sh` waits up to 3600 seconds for required components:

| Condition | Message | Exit code |
|---|---|---|
| missing components but still within timeout | `Waiting for required snap components: [...] (elapsed/max)` | still running |
| timeout reached with components still missing | `Error: timed out after ... while waiting for required components: [...]` | `1` |
| timeout remediation | `Please use "snap changes" to monitor the progress and start the service once all components are installed.` | informational text before exit |

## Error propagation model

The runtime error path is layered:

1. CLI or hook chooses an engine and persisted config.
2. Daemon startup checks component availability.
3. Engine launcher checks required component directories and critical config values.
4. Health check probes both process state and HTTP response validity.
5. Chat wait wrapper translates health-check states into user-facing readiness messages.

This layering is good because it separates setup failure, warmup delay, and true runtime failure into different messages and exit states.

## Gaps and risks

- User-command exit codes for `get`, `set`, `show-engine`, `use-engine`, and `completion bash` are not documented and not directly visible from source.
- `model-name` is treated as optionally missing, but that optionality is implicit rather than specified.
- The daemon and health-check scripts use precise exit codes, but those contracts are not described in user docs.
- Because `completion bash` suppresses stderr, users may get silent failure if the completion command changes shape.

## Recommendations for documenting the current model

- Document the `0/1/2` health-check contract as an internal runtime interface.
- Document that `chat` may fail with readiness errors even when the top-level command syntax is valid.
- Document that `show-engine` and `get` are the safest verification commands after mutation.
- Treat future changes to exit codes and machine-readable stderr/stdout paths as breaking per the deprecation spec.