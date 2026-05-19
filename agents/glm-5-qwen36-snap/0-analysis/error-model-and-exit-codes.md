# qwen36 Error Model And Exit Codes

## Confidence Level

CLI-native exit codes are not documented and the CLI source is not available, so this file separates observed shell-wrapper behavior from inferred CLI behavior.

## Observed Error Categories

| Category | Representative Message | Observed Source | Exit Code |
|----------|------------------------|-----------------|-----------|
| Missing required snap component | `Missing component: <path>` | `engines/*/common.sh` | `1` |
| Missing GPU configuration | `gpu-layers snap option is not set` | `engines/cuda/server` | `1` |
| Timed out waiting for components | `Error: timed out after ... while waiting for required components` | `apps/server.sh` | `1` |
| Server startup timeout | `Timed out waiting for server to start. Check the server logs and try again.` | `apps/wait-for-server.sh` | `1` |
| Server not running / failed | `Server is not running or failed. Please check the logs.` | `apps/wait-for-server.sh` | `1` |
| Unexpected helper failure | `Unexpected exit code when waiting for server: <n>` | `apps/wait-for-server.sh` | propagates helper code |
| Health check: server still starting | no user-facing output unless `DEBUG` set | `apps/check-server-llamacpp.sh` | `1` |
| Health check: server failed | `Server error: <message>` or debug text | `apps/check-server-llamacpp.sh` | `2` |
| Health check: server healthy | debug text only | `apps/check-server-llamacpp.sh` | `0` |

## Inferred CLI Error Categories

These are strongly implied but not directly evidenced:

| Command | Likely Failure Case | Evidence Level | Expected Behavior |
|---------|---------------------|----------------|-------------------|
| `qwen36 get <key>` | Unknown key or unset key | medium | Non-zero exit with stderr message; scripts treat `model-name` as optionally absent by swallowing stderr and exit code |
| `qwen36 set <key>=<value>` | Invalid key, invalid value, malformed assignment | low | Non-zero exit with stderr message |
| `qwen36 use-engine <name>` | Unsupported engine, incompatible hardware, missing confirmation | medium | Non-zero exit and/or confirmation prompt |
| `qwen36 show-engine` | No engine selected, corrupted config | low | Non-zero exit with stderr message |
| `qwen36 completion bash` | Unsupported shell target or internal completion failure | low | Non-zero exit, probably silent when used by shell completion wrapper |

## Exit Code Guidance By Command

### `chat`

- No direct evidence for CLI exit codes.
- Operationally depends on server availability and chat client startup.
- Should return non-zero if server is unreachable or client fails to start.

### `use-engine`

- No direct exit code evidence.
- Presence of `--assume-yes` implies the command otherwise may wait for or require interactive confirmation.
- Engine misselection can surface later as runtime errors in the daemon path.

### `show-engine`

- Must succeed for daemon startup to work.
- Changing output shape is effectively a breaking error for consumers even if exit code stays zero.

### `get`

- `model-name` is explicitly treated as optional in scripts via `2>/dev/null || true`.
- That pattern implies at least some keys can fail lookup without being fatal to the overall application.
- Unknown key should likely return non-zero exit code.

### `set`

- Shell automation should assume exit status is the only reliable success signal.
- No validation errors documented.

### `completion bash`

- The completer suppresses stderr with `2>/dev/null`, so user-visible errors are intentionally hidden during tab completion.
- Should fail silently for interactive completion scenarios.

## Exit Code Conventions (Per DE013)

Per the Canonical CLI standards, exit codes should follow:

- `0`: Success
- `1`: Usage error (invalid arguments, missing required arguments)
- `2`: Runtime error (command failed, resource not found)

The health check script uses `2` for "server failed" which aligns with this convention.

## Recommendations For Error Contract Documentation

1. Document exit code meanings for every leaf command.
2. Distinguish configuration lookup miss from parser misuse.
3. Document whether `use-engine --auto` can fail due to unsupported hardware versus missing permissions.
4. Publish the stderr wording that scripts should not parse.
5. Reserve distinct exit codes for usage error versus runtime failure, especially for `get`, `set`, and `use-engine`.
6. Consider adding `--json` output format for machine-parseable error messages.
