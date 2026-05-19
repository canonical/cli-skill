# Juju CLI Error Model And Exit Codes

## Framework-level exit codes

Juju's exit-code model is defined primarily by `cmd/cmd/cmd.go` and the top-level supercommand flow.

Verified behavior:
- `0`: successful execution
- `0`: help output requested via `gnuflag.ErrHelp`
- `1`: command `Run()` returned an error after successful parse/init
- `2`: parse/init/usage failure handled by `handleCommandError`
- plugin passthrough: arbitrary exit code if a plugin returns `RcPassthroughError`

This is a simple and mostly conventional scheme, but it is not semantically rich.

## Error lifecycle

### Parse and init failures

Errors raised during:
- flag parsing
- argument parsing
- command `Init()`

are handled before execution and return exit code `2`.

Examples:
- missing required operands: `no application name specified`
- invalid flag combinations: `cannot mix --no-color and --color`
- illegal safety combinations: `--timeout can only be used with --force (dangerous)`
- extra operands: `unrecognized args: ...`

### Runtime failures

Errors returned from `Run()` generally produce exit code `1`.

Examples include:
- API connection failures
- remote validation failures
- blocked operations
- missing resources on the controller
- remote service errors from Charmhub or controller facades

### Silent or shaped failures

Two special cases alter the normal path:
- `ErrSilent`: suppresses duplicate error printing
- `RcPassthroughError`: returns the embedded exit code directly

The supercommand frequently converts already-printed errors into `ErrSilent` so they are not printed again by `cmd.Main`.

## Supercommand-specific error handling

The supercommand adds important behavior that plain commands do not:
- deprecated command warnings before execution
- machine-format error shaping for JSON/YAML requests
- closest-command suggestions for unrecognized commands
- plugin fallback for unknown subcommands

### Unknown commands

If an unrecognized command is close to a known command, Juju returns a friendly suggestion of the form:
- `juju: "<arg>" is not a juju command ... Did you mean: ...`

If the command is truly unknown and no plugin matches, it falls back to the default unrecognized-command error.

### Plugin errors

For plugin commands:
- missing executable falls back to normal unknown-command handling
- plugin subprocess exit codes can pass through to the Juju process unchanged

This is good for shell integration.

## Error categories in the current CLI

### 1. Usage / validation errors

These are generally explicit and local.
Examples:
- `no model specified`
- `controller name must be specified`
- `expected type to be charm or bundle`
- `cannot set and reset key "..." simultaneously`

Quality: generally good.

### 2. Context-resolution errors

These arise in `modelcmd` when controller/model focus cannot be resolved.
Examples:
- no current model in focus
- unknown controller
- model removed from the controller
- redirected / migrated model messages with recovery steps

Quality: strong. These errors are among the better user-facing errors in the CLI because they usually include a next step.

### 3. Connection and authentication errors

Examples include:
- `cannot connect to API`
- `no controller API addresses; is bootstrap still in progress?`
- `cannot connect to k8s api server; try running 'juju update-k8s --client <k8s cloud name>'`
- login-related auth failures

Quality: mixed but often actionable. Some of the `modelcmd` connection wrappers do a good job of annotating lower-level failures.

### 4. Blocked-operation errors

Juju has an explicit operation-blocking model for destructive or change operations.
Commands often run returned errors through `block.ProcessBlockedError(...)`.

Effect:
- operator receives a domain-specific blocked message
- action is denied consistently across several mutation commands

Quality: conceptually strong, though not always obvious unless the operator already knows about disabled command sets.

### 5. Remote resource / not-found errors

Examples:
- key not found in controller config
- no matching charms for a Charmhub query
- missing cloud/controller/model/application resources

Quality: mostly acceptable, though phrasing and casing vary.

### 6. Long-running destructive errors

Destroy commands produce richer guidance than average.
Examples include:
- explicit storage-handling requirements
- force/no-wait warnings
- operation blocked remediation
- recommendation to use `enable-destroy-controller`

Quality: high relative to the rest of the surface.

## Error message quality assessment

### Strengths

- Many argument errors are concrete and immediate.
- `modelcmd` adds genuinely useful remediation guidance for missing/migrated models.
- Destructive commands explain why they are blocked and how to retry safely.
- Unknown-command suggestions improve recoverability.

### Weaknesses

- Error grammar is inconsistent across packages.
- Exit codes do not distinguish validation, auth, network, or server categories beyond `1` vs `2`.
- Some user-facing errors remain thin wrappers around lower-level API failures.
- Machine-readable error contracts are not formalized for JSON/YAML modes.

## Exit-code implications for automation

For scripts, the practical contract is:
- `0` means success or help output
- `2` means usage or parse/init failure
- `1` means runtime failure
- plugin commands may return their own code

That is enough for coarse shell control flow, but not enough for robust machine policy decisions without parsing stderr.

## Inconsistencies and notable quirks

- The supercommand may print errors itself and then return `ErrSilent`, while plain commands rely on `cmd.Main` for printing. That split is invisible to users but increases complexity.
- JSON/YAML requests can still emit errors on stderr while stdout receives an empty structured value. That is sensible for parsers, but surprising if undocumented.
- Some commands lean heavily on annotated errors, others return short raw errors.

## Overall assessment

Juju's error model is operationally competent:
- parse vs runtime failures are cleanly separated
- several domain-specific error paths are well designed
- plugin exit-code passthrough is pragmatic

The main weakness is semantic shallowness. Exit codes are too coarse, and error-message style is not standardized across the command surface. For human operators the result is acceptable; for automation and contract clarity it is only moderate.
