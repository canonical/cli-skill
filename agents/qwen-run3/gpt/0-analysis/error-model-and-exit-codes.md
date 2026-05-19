# qwen36 Error Model And Exit Codes

## Exit Code Model

At the Cobra root, all command-handler errors collapse to exit code `1`:

- `rootCmd.Execute()` returns an error
- `main()` calls `os.Exit(1)` for any error

That means the public CLI does not currently expose differentiated exit codes for usage errors versus runtime failures.

## Known Exceptions

The main exception is hidden `run`, where the child process exit status is propagated by `exec.Cmd.Run()`. In that case the CLI may exit with the subprocess status rather than a generic `1`.

## Error Categories

| Category | Representative Message | Where it comes from | Exit Code |
|---|---|---|---|
| Permission denied | `permission denied, try again with sudo` | `set`, `unset`, `use-engine`, `prune-cache` gate on root user | `1` |
| No active engine | `no active engine` | commands requiring a selected engine | `1` |
| Unknown config key | `key "<key>" is not found` plus suggestion | `get`, `set`, `unset` | `1` |
| Invalid assignment syntax | `expected key=value, got "..."` | `set` parser | `1` |
| Duplicate config key | `duplicate key: "..."` | `set` parser | `1` |
| Missing engine argument | `engine name not specified` | `use-engine` | `1` |
| Mutually exclusive selection modes | `cannot specify both engine name and --auto flag` | `use-engine` | `1` |
| Unknown engine | `engine "<name>" does not exist` or `"<name>" not found` | `show-engine`, `use-engine`, `prune-cache` | `1` |
| Unknown output format | `unknown format "..."` | `status`, `show-engine`, `show-machine`, `version`, `debug select-engine`, `list-engines` | `1` |
| Component install failure | `timed out while installing "<component>"` and related snapd errors | `use-engine` | `1` |
| Server inactive | `server not active` with a start suggestion | `chat` | `1` |
| Connection refused | `connection refused` plus startup/log suggestions | shared chat client | `1` |
| No models available on server | `no models available on server` | shared chat client | `1` |
| Web UI service inactive | `<service> not active` plus start suggestion | `webui` | `1` |
| Validation failure | `not all manifests are valid` | `debug validate-engines` | `1` |

## Per-Command Differences

### `status`

- may fail because no active engine exists
- may fail because component metadata contains deprecated or unsupported server config
- can optionally wait for components and therefore fail only after a long timeout

### `chat`

- has the richest error guidance
- handshake failures include suggestions to retry later and inspect snap logs
- server readiness errors are retried for up to 60 seconds before failing

### `webui`

- validates both `server` and `server-webui` service states
- also depends on chat-server readiness, so it inherits chat-client endpoint failures

### `use-engine`

- cancellation during component confirmation is not an error and returns success
- fix mode suppresses `no active engine` by treating it as a no-op
- manifest-not-found during `--fix` falls back to `--auto`

### `prune-cache`

- explicitly rejects pruning the active engine
- user cancellation is a success path with `Cancelled. No changes applied.`

### Hidden `run`

- may return child process exit codes directly
- also carries the risk of leaving temporary symlinks behind on abnormal termination

## Representative Shell-Level Exit Codes Outside The CLI

The shell wrappers under `apps/` and `engines/` use additional conventions:

- `apps/check-server-llamacpp.sh`: `0` healthy, `1` still starting, `2` failed
- `apps/wait-for-server.sh`: collapses several helper states into `1` for the user-facing wrapper
- `apps/server.sh`: exits `1` on component timeout and stops the snap service

Those codes matter operationally, but they are not the same as the main Cobra command exit model.

## Assessment

Strength:

- messages are usually specific and often include next-step guidance

Weakness:

- exit code `1` is overloaded for almost everything
- scripts cannot distinguish bad input, missing engine, permission failure, or transient service issues without parsing stderr
