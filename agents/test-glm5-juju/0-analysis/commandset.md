# Juju CLI Command Set

## Overview

The Juju CLI provides 159 commands organized into functional domains. This document catalogs all commands with their descriptions and implementation locations.

## Command Hierarchy

The CLI follows a flat command structure where all commands are top-level (no nested subcommands beyond the primary `juju` command). Commands are grouped by functional domain in the codebase.

## Complete Command List

### Action Commands (action/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `actions` | List actions defined for an application | `status/` (uses status) |
| `cancel-task` | Cancel a running action | `action/cancel.go` |
| `exec` | Run an action on an application/unit | `action/exec.go` |
| `operations` | List action operations | `action/listoperations.go` |
| `run` | Run an action on an application/unit | `action/run.go` |
| `show-action` | Show action definition | `action/show.go` |
| `show-operation` | Show action operation details | `action/showoperation.go` |
| `show-task` | Show action task details | `action/showtask.go` |

### Application Commands (application/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-unit` | Add units to an application | `application/addunit.go` |
| `bind` | Bind application endpoints to spaces | `application/bind.go` |
| `config` | Get or set application configuration | `application/config.go` |
| `consume` | Add a remote application to the model | `application/consume.go` |
| `deploy` | Deploy a new application or bundle | `application/deploy.go` |
| `diff-bundle` | Compare a bundle with the model | `application/diffbundle.go` |
| `expose` | Make an application publicly available | `application/expose.go` |
| `integrate` | Integrate two applications | `application/integrate.go` |
| `refresh` | Refresh an application's charm | `application/refresh.go` |
| `remove-application` | Remove an application from the model | `application/removeapplication.go` |
| `remove-relation` | Remove a relation between applications | `application/removerelation.go` |
| `remove-saas` | Remove a consumed remote application | `application/removesaas.go` |
| `remove-unit` | Remove units from an application | `application/removeunit.go` |
| `resolved` | Mark unit errors as resolved | `application/resolved.go` |
| `resume-relation` | Resume a suspended relation | `application/resumerelation.go` |
| `scale-application` | Scale a Kubernetes application | `application/scaleapplication.go` |
| `show-application` | Show application details | `application/show.go` |
| `show-unit` | Show unit details | `application/showunit.go` |
| `suspend-relation` | Suspend a relation | `application/suspendrelation.go` |
| `trust` | Grant charm access to credentials | `application/trust.go` |
| `unexpose` | Remove public exposure | `application/unexpose.go` |

### Backup Commands (backups/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `create-backup` | Create a backup of the controller | `backups/create.go` |
| `download-backup` | Download a backup file | `backups/download.go` |

### Block Commands (block/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `disable-command` | Disable a destructive operation | `block/disable.go` |
| `disabled-commands` | List disabled commands | `block/list.go` |
| `enable-command` | Enable a destructive operation | `block/enable.go` |

### Cloud Commands (cloud/, caas/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-cloud` | Add a cloud definition | `cloud/add.go` |
| `add-credential` | Add a credential for a cloud | `cloud/addcredential.go` |
| `add-k8s` | Add a Kubernetes cloud | `caas/add.go` |
| `autoload-credentials` | Auto-detect cloud credentials | `cloud/detectcredentials.go` |
| `clouds` | List available clouds | `cloud/list.go` |
| `credentials` | List credentials for a cloud | `cloud/listcredentials.go` |
| `default-credential` | Set default credential | `cloud/defaultcredential.go` |
| `default-region` | Set default region | `cloud/defaultregion.go` |
| `regions` | List regions for a cloud | `cloud/regions.go` |
| `remove-cloud` | Remove a cloud definition | `cloud/remove.go` |
| `remove-credential` | Remove a credential | `cloud/removecredential.go` |
| `remove-k8s` | Remove a Kubernetes cloud | `caas/remove.go` |
| `show-cloud` | Show cloud details | `cloud/show.go` |
| `show-credential` | Show credential details | `cloud/showcredential.go` |
| `update-cloud` | Update a cloud definition | `cloud/updatecloud.go` |
| `update-credential` | Update a credential | `cloud/updatecredential.go` |
| `update-k8s` | Update a Kubernetes cloud | `caas/update.go` |
| `update-public-clouds` | Update public cloud definitions | `cloud/updatepublicclouds.go` |

### Controller Commands (controller/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-model` | Add a model to the controller | `controller/addmodel.go` |
| `controller-config` | View or set controller configuration | `controller/config.go` |
| `controllers` | List all controllers | `controller/listcontrollers.go` |
| `destroy-controller` | Destroy a controller | `controller/destroy.go` |
| `enable-destroy-controller` | Enable controller destruction | `controller/enabledestroy.go` |
| `kill-controller` | Forcibly kill a controller | `controller/kill.go` |
| `list-models` | List models in a controller | `controller/listmodels.go` |
| `register` | Register a controller | `controller/register.go` |
| `show-controller` | Show controller details | `controller/showcontroller.go` |
| `unregister` | Unregister a controller | `controller/unregister.go` |

### Cross-Model Commands (crossmodel/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `find-offers` | Find available offers | `crossmodel/find.go` |
| `offer` | Create an offer for an endpoint | `crossmodel/offer.go` |
| `offers` | List offers in the model | `crossmodel/list.go` |
| `remove-offer` | Remove an offer | `crossmodel/remove.go` |
| `show-offer` | Show offer details | `crossmodel/show.go` |

### Charm Hub Commands (charmhub/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `download` | Download a charm from Charmhub | `charmhub/download.go` |
| `find` | Search for charms | `charmhub/find.go` |
| `info` | Show charm information | `charmhub/info.go` |

### Dashboard Commands (dashboard/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `dashboard` | Open the Juju Dashboard | `dashboard/dashboard.go` |

### Firewall Commands (firewall/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `firewall-rules` | List firewall rules | `firewall/listrules.go` |
| `set-firewall-rule` | Set a firewall rule | `firewall/setrule.go` |

### Machine Commands (machine/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-machine` | Add a machine to the model | `machine/add.go` |
| `machines` | List machines in the model | `machine/list.go` |
| `remove-machine` | Remove a machine from the model | `machine/remove.go` |
| `show-machine` | Show machine details | `machine/show.go` |

### Model Commands (model/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `config` | Get or set model configuration | `model/config.go` |
| `constraints` | View model constraints | `model/constraints.go` |
| `defaults` | View or set model defaults | `model/defaults.go` |
| `destroy-model` | Destroy a model | `model/destroy.go` |
| `export-bundle` | Export model as bundle | `model/exportbundle.go` |
| `grant` | Grant model access to a user | `model/grantrevoke.go` |
| `grant-cloud` | Grant cloud access to a user | `model/grantrevokecloud.go` |
| `migrate` | Migrate model to another controller | `commands/migrate.go` |
| `model-config` | (alias for config) | `model/config.go` |
| `model-constraints` | (alias for constraints) | `model/constraints.go` |
| `model-defaults` | (alias for defaults) | `model/defaults.go` |
| `model-secret-backend` | Set model secret backend | `secretbackends/modelsecretbackend.go` |
| `models` | List models | `model/show.go` (via status) |
| `retry-provisioning` | Retry machine provisioning | `model/retryprovisioning.go` |
| `revoke` | Revoke model access from a user | `model/grantrevoke.go` |
| `revoke-cloud` | Revoke cloud access from a user | `model/grantrevokecloud.go` |
| `set-constraints` | Set application constraints | `application/constraints.go` |
| `set-credential` | Set model credential | `model/setcredential.go` |
| `set-model-constraints` | Set model constraints | `model/constraints.go` |
| `show-model` | Show model details | `model/show.go` |

### Resource Commands (resource/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `attach-resource` | Attach a resource to an application | `resource/upload.go` |
| `charm-resources` | Show charm resource definitions | `resource/charmresources.go` |
| `resources` | List resources for an application | `resource/list.go` |

### Secret Commands (secrets/, secretbackends/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-secret` | Add a new secret | `secrets/add.go` |
| `add-secret-backend` | Add a secret backend | `secretbackends/add.go` |
| `grant-secret` | Grant access to a secret | `secrets/grantrevoke.go` |
| `remove-secret` | Remove a secret | `secrets/remove.go` |
| `remove-secret-backend` | Remove a secret backend | `secretbackends/remove.go` |
| `revoke-secret` | Revoke access to a secret | `secrets/grantrevoke.go` |
| `secret-backends` | List secret backends | `secretbackends/list.go` |
| `secrets` | List secrets | `secrets/list.go` |
| `show-secret` | Show secret details | `secrets/show.go` |
| `show-secret-backend` | Show secret backend details | `secretbackends/show.go` |
| `update-secret` | Update a secret | `secrets/update.go` |
| `update-secret-backend` | Update a secret backend | `secretbackends/update.go` |

### Space Commands (space/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-space` | Add a network space | `space/add.go` |
| `move-to-space` | Move subnets to a space | `space/move.go` |
| `reload-spaces` | Reload spaces from provider | `space/reload.go` |
| `remove-space` | Remove a network space | `space/remove.go` |
| `rename-space` | Rename a network space | `space/rename.go` |
| `show-space` | Show space details | `space/show.go` |
| `spaces` | List network spaces | `space/list.go` |

### SSH Commands (ssh/, sshkeys/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-ssh-key` | Add an SSH key to the model | `sshkeys/add_sshkeys.go` |
| `debug-code` | Debug charm code remotely | `ssh/debugcode.go` |
| `debug-hooks` | Debug charm hooks remotely | `ssh/debughooks.go` |
| `import-ssh-key` | Import SSH keys from a source | `sshkeys/import_sshkeys.go` |
| `remove-ssh-key` | Remove an SSH key | `sshkeys/remove_sshkeys.go` |
| `scp` | Secure copy files to/from machines | `ssh/scp.go` |
| `ssh` | SSH into a machine or container | `ssh/ssh.go` |
| `ssh-keys` | List SSH keys | `sshkeys/list_sshkeys.go` |

### Storage Commands (storage/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-storage` | Add storage to a unit | `storage/add.go` |
| `attach-storage` | Attach storage to a unit | `storage/attach.go` |
| `create-storage-pool` | Create a storage pool | `storage/poolcreate.go` |
| `detach-storage` | Detach storage from a unit | `storage/detach.go` |
| `import-filesystem` | Import a filesystem | `storage/import.go` |
| `remove-storage` | Remove storage | `storage/remove.go` |
| `remove-storage-pool` | Remove a storage pool | `storage/pooldelete.go` |
| `show-storage` | Show storage details | `storage/show.go` |
| `storage` | List storage | `storage/list.go` |
| `storage-pools` | List storage pools | `storage/poollist.go` |
| `update-storage-pool` | Update a storage pool | `storage/poolupdate.go` |

### Subnet Commands (subnet/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `subnets` | List subnets | `subnet/list.go` |

### Status/Debug Commands (status/, commands/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `bootstrap` | Initialize a cloud environment | `commands/bootstrap.go` |
| `debug-log` | View debug logs from the model | `commands/debuglog.go` |
| `show-status-log` | Show status change history | `status/history.go` |
| `status` | Show current model status | `status/status.go` |
| `switch` | Switch active model | `commands/switch.go` |
| `whoami` | Show current user and model | `user/whoami.go` |

### User Commands (user/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-user` | Add a user to the controller | `user/add.go` |
| `change-user-password` | Change a user's password | `user/change_password.go` |
| `disable-user` | Disable a user account | `user/disenable.go` |
| `enable-user` | Enable a user account | `user/disenable.go` |
| `login` | Log in to a controller | `user/login.go` |
| `logout` | Log out from a controller | `user/logout.go` |
| `remove-user` | Remove a user from the controller | `user/remove.go` |
| `show-user` | Show user details | `user/info.go` |
| `users` | List users | `user/list.go` |

### Upgrade Commands (commands/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `sync-agent-binary` | Sync agent binaries | `commands/synctools.go` |
| `upgrade-controller` | Upgrade the controller | `commands/upgradecontroller.go` |
| `upgrade-model` | Upgrade the model | `commands/upgrademodel.go` |

### Help/Documentation Commands (commands/)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `documentation` | Generate documentation | `cmd/cmd/documentation.go` |
| `help` | Show help | Built into SuperCommand |
| `help-action-commands` | Help for action commands | `commands/help_action_commands.go` |
| `help-hook-commands` | Help for hook commands | `commands/help_hook_commands.go` |
| `version` | Show version | `commands/version.go` |

### Miscellaneous Commands

| Command | Description | Implementation |
|---------|-------------|----------------|
| `relate` | Alias for integrate | `application/integrate.go` |

## Command Aliases

Several commands have aliases for backward compatibility or convenience:

| Alias | Primary Command |
|-------|-----------------|
| `relate` | `integrate` |
| `model-config` | `config` |
| `model-constraints` | `constraints` |
| `model-defaults` | `defaults` |

## Implementation Summary

| Domain | Package | Command Count |
|--------|---------|---------------|
| Actions | `action/` | 8 |
| Applications | `application/` | 21 |
| Backups | `backups/` | 2 |
| Blocks | `block/` | 3 |
| Clouds | `cloud/`, `caas/` | 18 |
| Controllers | `controller/` | 10 |
| Cross-model | `crossmodel/` | 5 |
| Charm Hub | `charmhub/` | 3 |
| Dashboard | `dashboard/` | 1 |
| Firewall | `firewall/` | 2 |
| Machines | `machine/` | 4 |
| Models | `model/` | 21 |
| Resources | `resource/` | 3 |
| Secrets | `secrets/`, `secretbackends/` | 12 |
| Spaces | `space/` | 7 |
| SSH | `ssh/`, `sshkeys/` | 8 |
| Storage | `storage/` | 11 |
| Subnets | `subnet/` | 1 |
| Status/Debug | `status/`, `commands/` | 6 |
| Users | `user/` | 10 |
| Upgrades | `commands/` | 3 |
| Help | `commands/` | 5 |
| **Total** | | **159** |

## Notes

1. The `actions` command is implemented in `status/` as it uses status output formatting.
2. The `models` command redirects to `juju status` with model filtering.
3. Some commands like `dump` and `dumpdb` are only available when `J