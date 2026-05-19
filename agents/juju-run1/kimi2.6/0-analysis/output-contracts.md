# Juju CLI Output Contracts

## Output modes

Global/shared formatter model:
- `smart` (default in many commands): human-oriented and type-adaptive
- `yaml`: machine-serializable
- `json`: machine-serializable

Shared output flags (where enabled by command):
- `--format <name>`
- `-o, --output <path>`

## Contract expectations by output type

1. Human-readable output
- Optimized for operator readability and task guidance.
- May include warnings, context messages, and tabular output.
- Not guaranteed stable for strict parsing.

2. Machine-readable output (`yaml`/`json`)
- Intended for scripts and integrations.
- CLI framework has explicit serializable-mode behavior to suppress noisy output patterns.
- On command failure, framework attempts to emit empty object output for machine formats (when feasible) rather than mixed human text.

## Stability guidance

- For automation, prefer `--format=json` or `--format=yaml` on commands that support formatter flags.
- Do not parse human tabular output in scripts.
- Treat field-level schema as command-specific; verify per-command docs and tests before relying on undocumented fields.

## Command-group differences

- Status-oriented commands expose richer formatter choices (for example tabular/summary/line variants in status domain).
- Some commands in docs show `N/A` usage entries, indicating generated docs are incomplete for that command page; these should be validated against command code/help.

## Parseability recommendations

- Use `--format=json` with explicit `--output` redirection where deterministic capture is required.
- When consuming stderr/stdout, account for warnings/deprecation messages outside formatter payloads unless command is in serializable mode.

## Evidence pointers

- Formatter definitions and output plumbing: `cmd/cmd/output.go`
- Serializable-mode and machine-format error handling: `cmd/cmd/supercommand.go`
- General command lifecycle and stderr error writer: `cmd/cmd/cmd.go`
- Status formatter variations: `cmd/juju/status/status.go`, `cmd/juju/status/output_tabular.go`, `cmd/juju/status/output_summary.go`
