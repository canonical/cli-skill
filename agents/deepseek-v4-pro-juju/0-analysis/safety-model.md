# Safety Model

## Overview

Juju provides safety mechanisms for destructive operations through several layered protections. The primary safeguards are: confirmation prompts, `--force` flags for override, `--dry-run` for preview, destroy protection blocks, and `--no-wait` for non-blocking teardown.

## Destructive Operations

Operations classified as destructive (data loss, resource termination):

| Category | Commands | Safeguards |
|----------|----------|------------|
| Resource removal | `remove-application`, `remove-unit`, `remove-machine`, `remove-storage`, `remove-saas`, `remove-relation`, `remove-offer`, `remove-space`, `remove-secret`, `remove-secret-backend`, `remove-k8s` | Confirmation prompt, `--force`, `--dry-run` (some) |
| Model destruction | `destroy-model` | Confirmation prompt, destroy protection blocks, `--force`, `--no-wait` |
| Controller destruction | `destroy-controller`, `kill-controller` | Strongest protection: `enable-destroy-controller` must be run first, then confirmation, `--t`/`--timeout` |
| User removal | `remove-user` | Confirmation prompt (`--yes` forced), `--force` |
| Credential removal | `remove-credential` | No confirmation prompt (low-risk local operation) |
| Cloud removal | `remove-cloud` | Cannot remove if in use by controllers |
| Application refresh | `refresh` | `--force`, `--force-units`, `--force-base` for bypassing safety checks |
| Bundle deployment | `deploy` (bundle) | `--dry-run` for preview |
| Force unit removal | `remove-unit --force` | Bypasses operational state checks |

## Safety Mechanisms

### 1. Confirmation Prompts

The `modelcmd` package provides confirmation utilities. Destructive commands that target models/controllers prompt for interactive confirmation:

```
Are you sure you want to continue? (y/N):
```

- `--yes` (`-y`) flag: Skip confirmation prompt (auto-answer yes)
- `--no-prompt`: Some commands support this for non-interactive use

Commands with confirmation prompts:
- `destroy-model` — warns about data loss, requires confirmation
- `destroy-controller` — warns about all models, requires confirmation
- `kill-controller` — warns about forced destruction
- `remove-user` — requires confirmation
- `remove-offer` — requires `--yes` when `--force` is used
- `disable-user` — no confirmation needed (reversible)
- `enable-destroy-controller` — enables the ability to destroy

### 2. Destroy Protection Blocks

The `block` package provides operation-level protection:

- `disable-command destroy-model` — prevents model destruction
- `disable-command remove-object` — prevents removal of applications, units, machines
- `disable-command all` — disables all destructive operations
- `enable-command` — re-enables blocked operations
- `disabled-commands` — lists currently blocked operations

This is a server-side protection that must be bypassed by an admin.

### 3. `--force` Flag

The `--force` flag bypasses normal safety checks:

| Command | What `--force` does |
|---------|---------------------|
| `remove-application --force` | Removes even if units are in error state; removes subordinates |
| `remove-unit --force` | Removes even if in error state; may remove machine |
| `remove-machine --force` | Removes even if units are present |
| `remove-storage --force` | Force-detaches storage before removal |
| `remove-relation --force` | Bypasses operational checks |
| `remove-saas --force` | Force-removes consumed offer |
| `destroy-model --force` | Force-destroys model ignoring errors; may need `--no-wait` |
| `destroy-controller --force` | Force-destroys even with errors |
| `kill-controller` | Implicitly forces destruction without server cooperation |
| `add-cloud --force` | Adds cloud even if credentials fail validation |
| `deploy --force` | Deploys charm even if base/LXD profile checks fail |
| `refresh --force` | Bypasses LXD profile allow list checks |
| `refresh --force-units` | Refreshes units even if in error state |
| `refresh --force-base` | Refreshes even if base not supported by new charm |
| `bind --force` | Binds endpoints even if space not available to all units |
| `update-credential --force` | Updates ignoring validation errors |
| `update-k8s --force` | Forces cloud update |
| `remove-k8s --force` | Force-removes cloud from controller |
| `detach-storage --force` | Force-detaches storage |

### 4. `--dry-run` Flag

Preview without performing the operation:

| Command | What `--dry-run` shows |
|---------|------------------------|
| `deploy --dry-run` | Shows what would be deployed (charm, units, machines) |
| `remove-application --dry-run` | Shows what would be removed |
| `remove-unit --dry-run` | Shows what would be removed |

Notes:
- `remove-unit --dry-run` is not supported for Kubernetes units (returns error)
- `remove-application --dry-run` may not be supported by older controllers (returns `errDryRunNotSupportedByController`)

### 5. `--no-wait` Flag

For teardown operations, `--no-wait` returns immediately without waiting for completion:

```
juju remove-application --force --no-wait <app>
```

This is faster but leaves resources in a "dying" state. Users must monitor the status separately. `--no-wait` always requires `--force`.

### 6. Timeouts

Destructive operations that can hang support explicit timeouts:

| Command | Flag | Default |
|---------|------|---------|
| `destroy-model` | `--timeout` / `-t` | Unset (blocks until complete) |
| `destroy-controller` | `--model-timeout` | Unset |
| `kill-controller` | `--timeout` / `-t` | 5 minutes |
| `upgrade-model` | `--timeout` | 10 minutes |
| `upgrade-controller` | `--timeout` | 10 minutes |

### 7. Recovery Operations

Commands for recovering from problematic states:

| Command | Purpose |
|---------|---------|
| `enable-destroy-controller` | Re-enables controller destruction if previously blocked |
| `retry-provisioning` | Retries provisioning of failed machines |
| `resolved` | Marks unit errors as resolved (optionally re-runs failed hooks) |
| `resolved --no-retry` | Marks resolved without re-running hooks |
| `resolved --all` | Resolves all units in error |
| `kill-controller` | Last-resort controller destruction when unreachable |

## Risk Level Summary

| Risk Level | Commands | Default Protection |
|------------|----------|-------------------|
| **High** (controller loss) | `destroy-controller`, `kill-controller` | `enable-destroy-controller` gate + confirmation + timeout |
| **Medium** (model/resource loss) | `destroy-model`, `remove-application`, `remove-unit`, `remove-machine` | Confirmation prompt, `--force`, `--dry-run` |
| **Low** (local state) | `remove-credential`, `remove-cloud`, `unregister` | Local-only, no controller impact |
| **Reversible** (state toggle) | `disable-user`, `disable-command`, `suspend-relation`, `unexpose` | Can be undone via `enable-*` / `resume-*` |
| **Data loss** (storage) | `remove-storage`, `detach-storage` | Confirmation, `--force`, `--destroy-storage` vs `--release-storage` |

## Gaps & Concerns

1. **No undo**: Juju has no built-in undo/revert mechanism. Once a destructive command completes, recovery is manual.
2. **No backup prompt**: `remove-application`, `remove-machine`, `remove-unit` do not prompt for confirmation by default (only `destroy-model`/`destroy-controller` do).
3. **`--force --no-wait` is powerful**: The combination allows rapid, irreversible destruction. A warning or additional confirmation could be appropriate.
4. **Dry-run coverage**: Only `deploy`, `remove-application`, and `remove-unit` support `--dry-run`. Missing coverage for `destroy-model`, `remove-relation`, `remove-machine`.
5. **`--yes` vs `--force` semantics**: These are sometimes conflated. `--yes` answers the prompt; `--force` bypasses operational checks. Users may not understand the distinction.
