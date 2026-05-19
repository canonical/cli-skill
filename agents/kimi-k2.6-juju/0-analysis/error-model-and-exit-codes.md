# Error Model and Exit Codes

## Exit Code Map
| Code | Meaning | Trigger |
|---|---|---|
| 0 | Success | Command completed without error |
| 1 | General error | `Run()` returned a non-silent error; printed to stderr |
| 2 | Flag/Init/Help error | `gnuflag.ErrHelp`, `ErrSilent`, or `CheckEmpty` failure |
| N | Plugin passthrough | `utils.RcPassthroughError` propagates the exit code of a `juju-*` plugin executable |

## Error Categories
1. **Argument/Flag Errors** (`gnuflag` parse errors, `CheckEmpty`, missing positional args) → exit 2, message printed via `WriteError`.
2. **API Errors** (NotFound, Unauthorized, Forbidden, BadRequest) → exit 1. Common messages:
   - `unrecognized command: juju <cmd>` (typo)
   - `You do not have permission to ...` (authorization)
   - `model <name> not found`
3. **Network/Timeout Errors** → exit 1, often wrapped with `errors.Trace`.
4. **Terms Agreement** (`TermsRequiredError`) → exit 1 with suggestion to run `juju agree`.
5. **Validation Errors** (constraint syntax, base image, credential mismatch) → exit 1.
6. **Missing Model/Controller Context** (`MissingModelNameError`) → exit 2 with suggested correction.

## Per-Command Differences
- `bootstrap` can emit warnings (non-fatal) about public cloud downloads; fatal bootstrap failures exit 1.
- `destroy-controller` and `kill-controller` may exit 1 if resources remain after timeout, even with `--force`.
- `exec` / `ssh` / `scp` passthrough remote command exit codes? No, Juju wraps them; remote errors are reported as Juju errors (exit 1).
- Plugin commands preserve the original executable exit code via `RcPassthroughError`.

## Stability
- Exit codes 0, 1, 2 are stable and used consistently across all commands.
- No command uses custom exit codes beyond plugin passthrough.
