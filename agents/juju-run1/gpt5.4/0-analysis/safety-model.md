# Juju CLI Safety Model

## Overview

Juju has a real safety model, but it is not uniformly applied across the whole CLI. Safety is strongest around destructive model/controller lifecycle operations and around the command-blocking system. It is weaker for broad mutation commands that change application, cloud, or networking state without preview support.

The main safety mechanisms are:
- confirmation prompts
- `--no-prompt` bypasses
- `--force` overrides
- `--no-wait` acceleration for destructive operations
- `--destroy-storage` / `--release-storage` explicit storage decisions
- `--dry-run` on selected commands such as `deploy`
- operation blocking through `disable-command`

## Strongest safety area: model and controller destruction

### `destroy-model`

Verified protections:
- prompts for confirmation by default
- explicit summary warning to stderr before destruction
- details of affected resources in the warning output
- requires an explicit storage decision when detachable storage remains
- `--force` and `--no-wait` are validated together
- `--timeout` is only legal with `--force`
- help text strongly warns about cleanup consequences

### `destroy-controller`

Verified protections:
- prompts for confirmation by default
- refuses to proceed if hosted models remain unless `--destroy-all-models` is specified
- requires explicit `--destroy-storage` or `--release-storage` when persistent storage remains
- warns extensively about `--force`, `--no-wait`, and `--model-timeout`
- includes specialized remediation if disabled commands are blocking destruction

### `kill-controller`

This is the intentionally dangerous escape hatch. It exists as a separate verb instead of overloading `destroy-controller --force` indefinitely. That is a good safety design because it keeps the dangerous path lexically distinct.

## Confirmation model

Juju uses explicit confirmation helpers for several destructive commands through `DestroyConfirmationCommandBase` and `UserConfirmName` flows.

Observed use cases include:
- `destroy-model`
- `destroy-controller`
- `unregister`

The model is usually:
1. print a warning with the named target
2. require explicit confirmation unless `--no-prompt` is set

This is effective for irreversible actions, though it is not used on every impactful mutator.

## Force flags

`--force` exists on several commands, but semantics vary.

Examples:
- `deploy --force`: bypasses deployment checks such as supported base or profile allow-list checks
- `destroy-model --force`: ignore operational errors and allow harsher teardown
- `destroy-controller --force`: same, but at controller scope
- `move-to-space --force`: allow moving subnets even if they are in use elsewhere

This inconsistency is typical in large CLIs, but it means `--force` is a family resemblance, not a strict contract.

## Dry-run support

Verified dry-run support is limited.

### `deploy --dry-run`

`deploy` explicitly supports `--dry-run` and documents it as showing what deployment would do without applying it.

This is valuable because deploy is one of the highest-impact commands in the CLI.

### Gaps

There is no comparable first-class dry-run support across most other mutating surfaces:
- no generic dry-run for `refresh`
- no dry-run for `config`, `model-config`, or `controller-config`
- no dry-run for relation changes, removals, or access-control changes
- no dry-run for storage movement/removal commands

That is one of the biggest consistency gaps in the current safety model.

## Command blocking as a policy safety layer

Juju's distinctive safety feature is the command-blocking model:
- `disable-command`
- `enable-command`
- `disabled-commands`
- `enable-destroy-controller`

Block groups include:
- `destroy-model`
- `remove-object`
- `all`

This protects models from accidental or policy-disallowed mutation. It is stronger than ad hoc confirmations because it changes whether commands are even allowed to succeed.

Strengths:
- model-scoped safety policy
- can protect destructive and change-heavy command families
- some commands can surface blocked-operation guidance

Weaknesses:
- the mental model is not obvious to new users
- bypass behavior via `--force` exists for some commands, which reduces uniformity
- command grouping names like `remove-object` are functional but not especially discoverable

## Storage safety

Storage-related destructive commands are more explicit than average.

Patterns include:
- forcing the operator to choose between destroy vs release in model/controller destruction
- explicit `detach-storage` and `attach-storage` verbs
- `import-filesystem` to reuse released storage

This is good lifecycle hygiene. The CLI explicitly models irreversible destruction versus relinquishing management.

## Authentication and trust safety

Some sensitive flows have their own safeguards:
- `login --trust` explicitly opts into trusting a controller CA certificate
- `trust` is an explicit post-deploy command rather than an implicit side effect
- `deploy --trust` makes privileged intent explicit at deployment time

That is better than hidden escalation, though there is not one unified privilege-impact warning system.

## Safety weaknesses by family

### Application and relation mutation

Commands like:
- `config`
- `bind`
- `integrate`
- `remove-relation`
- `refresh`
- `unexpose`

generally rely on argument validation and server-side correctness rather than confirmation, preview, or rollback UX.

### Cloud and credential mutation

Commands such as:
- `update-cloud`
- `remove-cloud`
- `update-credential`
- `remove-credential`

have relatively light safety scaffolding compared with model/controller destruction.

### Network mutation

`move-to-space` exposes `--force`, but there is no preview mode. It can produce structured change output after the fact, but not a dry-run plan.

## Inconsistencies

- `--no-prompt` usually means skip destructive confirmation, but `login --no-prompt` means read password from stdin without an interactive prompt.
- `--force` semantics vary widely by command family.
- only a small subset of mutators offer dry-run behavior.
- not every irreversible or high-impact command uses a confirmation gate.

## Assessment

Strengths:
- very strong safety around destroy flows
- explicit storage-decision requirements are well designed
- command-blocking is a distinctive and genuinely useful policy control
- dangerous escape hatches are lexically separated in some cases, such as `kill-controller`

Weaknesses:
- dry-run support is sparse
- mutation safety outside destroy flows is uneven
- `--force` is not standardized enough to communicate predictable semantics
- safety guidance is concentrated in help text rather than surfaced consistently before execution

## Net conclusion

Juju's safety model is serious where the project knows the stakes are highest: controller and model teardown, storage handling, and policy blocking. The next maturity step would be to extend that discipline sideways across non-destructive-but-impactful mutations, especially by standardizing previews, dry-runs, and force semantics.
