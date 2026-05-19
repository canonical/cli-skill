# Juju CLI Error Model and Exit Codes

## Exit Code Conventions

| Exit Code | Meaning |
|---|---|
| `0` | Success |
| `1` | General error (command failed, API error, validation error) |
| `2` | CLI usage error (bad args, unrecognized command, missing required arg) |

The exit codes are defined in the custom `cmd` package. `cmd.Main()` translates error types to exit codes:
- `ErrSilent` or `RcPassthroughError` → exit code from wrapped process
- `UnrecognizedCommand` → exit code 2
- Any other error → exit code 1

## Error Categories

### 1. CLI Parsing Errors (Exit Code 2)

| Scenario | Representative Message |
|---|---|
| Unrecognized command | `juju: "foo" is not a juju command. See "juju --help".` |
| Missing required argument | `ERROR missing application name` |
| Unknown flag | `error: flag provided but not defined: -foo` |
| Flag value error | `error: invalid value "foo" for flag --format` |
| Invalid constraint syntax | `invalid constraint syntax: ...` |

### 2. API Connection Errors (Exit Code 1)

| Scenario | Representative Message |
|---|---|
| No current controller | `No selected controller.` |
| Controller unreachable | `cannot connect to API: ...` |
| Authentication failed | `invalid entity name or password` |
| Macaroon discharge required | (triggers browser-based auth or fails with `-B`) |
| TLS certificate error | `x509: certificate signed by unknown authority` |

### 3. Model/Entity Not Found (Exit Code 1)

| Scenario | Representative Message |
|---|---|
| Model not found | `model "foo" not found` |
| Application not found | `application "foo" not found` |
| Unit not found | `unit "foo/0" not found` |
| Machine not found | `machine 0 not found` |
| Cloud not found | `cloud foo not found` |

### 4. Permission Errors (Exit Code 1)

| Scenario | Representative Message |
|---|---|
| Insufficient model access | `permission denied` |
| Cannot grant access | `you do not have permission to grant ...` |
| Cannot modify controller config | `restricted controller configuration` |

### 5. State/Provisioning Errors (Exit Code 1)

| Scenario | Representative Message |
|---|---|
| Machine not provisioned | `machine 0 is not provisioned` |
| Unit not started | `unit ... has not started` |
| Application already exists | `application already exists` |
| Charm not found | `charm "foo" not found` |
| Storage not found | `storage ... not found` |

### 6. Destruction/Blocking Errors (Exit Code 1)

| Scenario | Representative Message |
|---|---|
| Destroy blocked by disabled commands | `cannot destroy model: model destruction is disabled` |
| Destroy with persistent storage | `destroying storage ...` (warning, not error) |
| Force required | `some machines have not been destroyed, use --force` |

### 7. Action/Operation Errors (Exit Code 1)

| Scenario | Representative Message |
|---|---|
| Action failed | `action failed with error: ...` |
| Action timeout | `timed out waiting for action to complete` |
| Invalid action parameters | `parameter "foo" not found` |

## Error Message Style

Juju errors generally follow these patterns:
- Prefix: `ERROR ` (uppercase, with trailing space)
- No contractions: "cannot" instead of "can't"
- Passive voice for missing args: `Invalid URL: ...`
- Active voice for guidance: `Create a new controller using "juju bootstrap" ...`

## Per-Command Differences

### Bootstrap
- Complex multi-stage error reporting
- May fail mid-process, leaving partial infrastructure
- Error messages include provider-specific details (e.g., AWS credential errors)

### SSH / SCP / Debug Hooks
- Passthrough of ssh exit codes
- Connection failures are distinct from remote command failures

### Exec
- Aggregates errors across multiple targets
- Per-unit errors are displayed, but overall exit code may still be 0 if some succeeded

### Plugin Commands
- Exit code is passed through from the plugin process
- Error output comes directly from the plugin

## Deprecation Warnings

Deprecation warnings are printed to stderr and do not affect the exit code:
```
warning: "foo bar" is deprecated and will be removed in the 4.0 release
```

## Silent Errors

Some commands return `ErrSilent` to suppress default error output. This is used when:
- The command has already printed a custom error message
- The error is expected and handled gracefully
