# Juju CLI Safety Model

## Destructive Operations

Juju has several categories of destructive operations, with varying levels of protection.

### High-Impact Destruction Commands

| Command | Destructive Effect | Confirmation Required | Force Flag | Dry-Run |
|---|---|---|---|---|
| `destroy-controller` | Destroys controller and all models | Yes (`-y` or prompt) | `--force`, `--no-wait` | No |
| `destroy-model` | Destroys model and all resources | Yes (`-y` or prompt) | `--force`, `--no-wait` | No |
| `kill-controller` | Forcibly terminates controller | No (implicitly forceful) | `--timeout` | No |
| `remove-application` | Removes application and units | No | `--force`, `--no-wait` | No |
| `remove-unit` | Removes units | No | `--force`, `--no-wait`, `--destroy-storage` | No |
| `remove-machine` | Removes machines | No | `--force`, `--no-wait`, `--no-keep-instance` | No |
| `remove-relation` | Removes relation | No | `--force` | No |
| `remove-storage` | Removes storage | No | `--force` | No |
| `remove-offer` | Removes cross-model offers | No | `--force` | No |
| `remove-cloud` | Removes cloud definition | No | | No |
| `remove-credential` | Removes credential | No | | No |
| `remove-user` | Removes user from controller | No | | No |
| `unregister` | Unregisters controller from client | Yes (`-y` or prompt) | | No |
| `revoke` | Revokes user access | No | | No |
| `revoke-cloud` | Revokes cloud access | No | | No |
| `revoke-secret` | Revokes secret access | No | | No |

### Confirmation Mechanisms

#### Explicit Confirmation Prompt

`destroy-controller`, `destroy-model`, and `unregister` prompt for confirmation:
```
$ juju destroy-model mymodel
WARNING! This command will destroy the "mymodel" model.
This includes all machines, applications, data and other resources.

Continue [y/N]?
```

The `-y` / `--yes` flag bypasses the prompt:
```
juju destroy-model mymodel -y
```

#### Block Commands (Operation Protection)

Administrators can disable destructive commands at the model or controller level:

```
juju disable-command destroy-model "Maintenance window"
juju enable-command destroy-model
```

Blocked commands:
| Block Name | Commands Blocked |
|---|---|
| `destroy-model` | `destroy-model` |
| `remove-object` | `remove-application`, `remove-machine`, `remove-unit`, `remove-relation` |
| `all` | All removal and destruction commands |

When a blocked command is attempted:
```
ERROR cannot destroy model: model destruction is disabled
```

The `enable-destroy-controller` command exists specifically to remove controller-level blocks before destruction.

### Force Flags

The `--force` flag accelerates destruction by bypassing graceful shutdown:
- `--force` on `destroy-model`: Skip waiting for agents to report
- `--force` on `remove-application`: Remove immediately, ignore errors
- `--force` on `remove-unit`: Remove immediately, ignore hook errors
- `--force` on `remove-machine`: Remove immediately, ignore agent presence
- `--force` on `remove-relation`: Remove immediately, ignore hook errors

The `--no-wait` flag is often paired with `--force` to not wait for operations to complete in the background.

### Storage Destruction Behavior

Storage has nuanced destruction semantics:

| Flag | Behavior |
|---|---|
| `--destroy-storage` | Destroy persistent storage volumes |
| `--release-storage` | Detach but do not destroy storage |
| Default (no flag) | Detach storage, do not destroy |

This applies to `destroy-model` and `remove-application`.

### Machine Termination

By default, `remove-machine` destroys the cloud instance. The `--no-keep-instance` flag is the default behavior for IAAS; for CAAS, pods are always destroyed.

### Recovery Behavior

Juju does not have a built-in "undo" or "recycle bin" for destroyed resources:
- Destroyed models cannot be recovered
- Destroyed applications can be re-deployed from charm/bundle, but data is lost
- Destroyed machines release cloud instances; data on ephemeral disks is lost
- Relations can be re-created with `integrate`

### Safe Defaults

- Bootstrap does not destroy existing resources
- Deploy does not overwrite existing applications (fails with "already exists")
- Model config changes apply immediately but are reversible
- Application config can be reverted by setting back to default

### Credential Safety

Cloud credentials are never transmitted in full in output:
- `show-credential` masks secret values unless `--show-secrets` is used
- Credential files are stored with 0600 permissions

## Safety Checklist

1. **Pre-destruction review**: `status`, `storage`, and `show-model` should be checked before destruction
2. **Backups**: `create-backup` exists for model-level backup before major changes
3. **Blocks**: Use `disable-command` to prevent accidental destruction in production
4. **Force awareness**: `--force` should be used sparingly and understood to bypass safeguards
5. **Storage awareness**: Always specify `--destroy-storage` or `--release-storage` explicitly when desired
