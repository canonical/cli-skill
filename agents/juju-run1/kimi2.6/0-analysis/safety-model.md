# Juju CLI Safety Model

## Safety primitives

1. Confirmation prompts
- Shared helpers provide explicit interactive confirmation (`Continue [y/N]?`) and typed-name confirmation for destructive actions.
- Default behavior is conservative: non-affirmative input aborts.

2. No-prompt escape hatches
- Remove/destroy command bases provide `--no-prompt` to skip confirmation prompts.
- For remove operations, model mode/config can require prompts unless explicitly overridden.

3. Disabled command mechanism
- Command families can be disabled/enabled at model/controller level (`disable-command`, `enable-command`, `disabled-commands`) as an operational guardrail.

4. Explicit targeting
- Many destructive commands require explicit resource identifiers (`<model>`, `<machine>`, `<application>`), reducing accidental broad actions.

## Destructive operation families

High-risk categories include:
- `destroy-*` (model/controller teardown)
- `remove-*` (applications, units, machines, storage, users, offers, clouds, credentials, secrets)
- operations that change network exposure (`expose`, `unexpose`) or security posture (`trust`)

## Dry-run and force semantics

- Some command families include command-specific force/no-prompt semantics.
- Dry-run behavior is command-specific and not globally enforced by the framework.
- Users should validate each destructive command page before automation.

## Recovery and rollback characteristics

- Many destructive operations are not automatically reversible.
- Backups and export paths (`create-backup`, `download-backup`, `export-bundle`) are available to support recovery planning.

## Safety risks observed

- `--no-prompt` flags can bypass interactive guardrails in scripts; governance controls should limit their use in production automation.
- Plugin command execution inherits trust from PATH resolution; environments should harden PATH and plugin provenance.

## Recommendations

- Require machine-readable output and explicit target qualification in automation.
- Prefer pre-flight backup/export for controller/model destructive workflows.
- Gate no-prompt/force usage via CI policy or wrapper tooling.

## Evidence pointers

- Confirmation helpers: `cmd/helpers.go`, `cmd/modelcmd/confirmation.go`
- Command registration including safety-related commands: `cmd/juju/commands/main.go`
- Plugin execution path: `cmd/juju/commands/plugin.go`
