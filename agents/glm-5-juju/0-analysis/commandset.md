# Juju CLI Command Set

## Overview

Juju provides a comprehensive CLI with 150+ commands organized into functional domains. Commands follow a verb-noun naming pattern (e.g., `add-model`, `deploy-application`) with some exceptions for legacy or special-case commands.

## Command Hierarchy

The CLI uses a flat command structure with no subcommand nesting. All commands are direct children of the `juju` root command. However, commands are logically grouped by domain in the implementation and help output.

## Full Command List

### Infrastructure Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `bootstrap` | Initializes a cloud environment by creating a controller | `cmd/juju/commands/bootstrap.go` - Provisions controller machine, installs Juju agent |
| `add-cloud` | Adds a user-defined cloud to Juju | `cmd/juju/cloud/add.go` - Stores cloud definition in local config |
| `remove-cloud` | Removes a cloud from Juju | `cmd/juju/cloud/remove.go` - Deletes cloud from local configuration |
| `update-cloud` | Updates cloud information available to Juju | `cmd/juju/cloud/update.go` - Refreshes cloud definition |
| `update-public-clouds` | Updates public cloud information | `cmd/juju/cloud/updatepublic.go` - Fetches latest public cloud metadata |
| `show-cloud` | Shows detailed information for a cloud | `cmd/juju/cloud/show.go` - Displays cloud configuration and regions |
| `clouds` / `list-clouds` | Lists all clouds available to Juju | `cmd/juju/cloud/list.go` - Shows configured clouds |
| `regions` / `list-regions` | Lists regions for a given cloud | `cmd/juju/cloud/regions.go` - Shows available regions |
| `add-k8s` | Adds a Kubernetes endpoint and credential to Juju | `cmd/juju/caas/add.go` - Registers Kubernetes cluster |
| `remove-k8s` | Removes a k8s cloud from Juju | `cmd/juju/caas/remove.go` - Unregisters Kubernetes cluster |
| `update-k8s` | Updates an existing Kubernetes endpoint | `cmd/juju/caas/update.go` - Modifies Kubernetes cloud config |

### Credential Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-credential` | Adds a credential for a cloud | `cmd/juju/cloud/addcredential.go` - Stores credentials locally and optionally uploads to controller |
| `remove-credential` | Removes Juju credentials for a cloud | `cmd/juju/cloud/removecredential.go` - Deletes credential |
| `update-credential` | Updates a controller credential | `cmd/juju/cloud/updatecredential.go` - Refreshes credential on controller |
| `show-credential` | Shows credential information | `cmd/juju/cloud/showcredential.go` - Displays credential details |
| `credentials` / `list-credentials` | Lists Juju credentials for a cloud | `cmd/juju/cloud/listcredentials.go` - Shows stored credentials |
| `autoload-credentials` | Auto-detects and adds cloud credentials | `cmd/juju/cloud/detect.go` - Discovers credentials from environment |
| `default-credential` | Gets, sets, or unsets default credential | `cmd/juju/cloud/defaultcredential.go` - Manages default credential |
| `default-region` | Gets, sets, or unsets default region | `cmd/juju/cloud/defaultregion.go` - Manages default region |

### Controller Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `controllers` / `list-controllers` | Lists all controllers | `cmd/juju/controller/list.go` - Shows registered controllers |
| `show-controller` | Shows detailed controller information | `cmd/juju/controller/show.go` - Displays controller config and status |
| `destroy-controller` | Destroys a controller | `cmd/juju/controller/destroy.go` - Terminates controller and all models |
| `kill-controller` | Forcibly terminates controller | `cmd/juju/controller/kill.go` - Force destroys without cleanup |
| `enable-destroy-controller` | Enables destroy-controller command | `cmd/juju/controller/enabledestroy.go` - Removes protection blocks |
| `register` | Registers a controller | `cmd/juju/controller/register.go` - Adds controller to local config |
| `unregister` | Unregisters a Juju controller | `cmd/juju/controller/unregister.go` - Removes from local config |
| `controller-config` | Displays or sets controller configuration | `cmd/juju/controller/config.go` - Manages controller settings |
| `upgrade-controller` | Upgrades Juju on a controller | `cmd/juju/commands/upgradecontroller.go` - Upgrades agent binaries |

### Model Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-model` | Adds a workload model | `cmd/juju/controller/addmodel.go` - Creates new model |
| `models` / `list-models` | Lists models a user can access | `cmd/juju/model/show.go` - Shows available models |
| `show-model` | Shows information about a model | `cmd/juju/model/show.go` - Displays model details |
| `destroy-model` | Terminates all machines/containers for a model | `cmd/juju/model/destroy.go` - Removes model and resources |
| `model-config` | Displays or sets model configuration | `cmd/juju/model/config.go` - Manages model settings |
| `model-defaults` | Displays or sets default configuration for new models | `cmd/juju/model/defaults.go` - Manages model defaults |
| `model-constraints` | Displays machine constraints for a model | `cmd/juju/model/constraints.go` - Shows model constraints |
| `set-model-constraints` | Sets machine constraints on a model | `cmd/juju/model/constraints.go` - Sets model-level constraints |
| `migrate` | Migrates a workload model to another controller | `cmd/juju/commands/migrate.go` - Moves model between controllers |
| `switch` | Selects or identifies current controller and model | `cmd/juju/commands/switch.go` - Changes active context |
| `upgrade-model` | Upgrades Juju on all machines in a model | `cmd/juju/commands/upgrademodel.go` - Upgrades model agents |

### Application Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `deploy` | Deploys a new application or bundle | `cmd/juju/application/deploy.go` - Creates application from charm |
| `remove-application` | Removes applications from the model | `cmd/juju/application/removeapplication.go` - Deletes application |
| `config` | Get, set, or reset configuration for an application | `cmd/juju/application/config.go` - Manages application config |
| `constraints` | Displays machine constraints for an application | `cmd/juju/application/constraints.go` - Shows app constraints |
| `set-constraints` | Sets machine constraints for an application | `cmd/juju/application/constraints.go` - Sets app constraints |
| `show-application` | Displays information about an application | `cmd/juju/application/show.go` - Shows application details |
| `refresh` | Refreshes an application's charm | `cmd/juju/application/refresh.go` - Upgrades charm version |
| `expose` | Makes an application publicly available | `cmd/juju/application/expose.go` - Opens network access |
| `unexpose` | Removes public availability for an application | `cmd/juju/application/unexpose.go` - Closes network access |
| `trust` | Sets trust status of an application to true | `cmd/juju/application/trust.go` - Grants credential access |
| `bind` | Changes bindings for a deployed application | `cmd/juju/application/bind.go` - Configures endpoint-space bindings |
| `scale-application` | Sets desired number of k8s application units | `cmd/juju/application/scaleapplication.go` - Scales Kubernetes apps |

### Unit Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-unit` | Adds one or more units to a deployed application | `cmd/juju/application/addunit.go` - Scales out application |
| `remove-unit` | Removes application units from the model | `cmd/juju/application/removeunit.go` - Scales in application |
| `show-unit` | Displays information about a unit | `cmd/juju/application/showunit.go` - Shows unit details |
| `resolved` / `resolve` | Marks unit errors resolved and re-executes hooks | `cmd/juju/application/resolved.go` - Clears error state |

### Machine Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-machine` | Provisions a new machine or assigns one to the model | `cmd/juju/machine/add.go` - Creates machine instance |
| `remove-machine` | Removes one or more machines from a model | `cmd/juju/machine/remove.go` - Terminates machine |
| `machines` / `list-machines` | Lists machines in a model | `cmd/juju/machine/list.go` - Shows machine status |
| `show-machine` | Shows a machine's status | `cmd/juju/machine/show.go` - Displays machine details |
| `retry-provisioning` | Retries provisioning for failed machines | `cmd/juju/model/retryprovisioning.go` - Retries failed provisioning |

### Integration (Relation) Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `integrate` / `relate` | Integrates two applications | `cmd/juju/application/integrate.go` - Creates relation |
| `remove-relation` | Removes an existing relation between applications | `cmd/juju/application/removerelation.go` - Deletes relation |
| `suspend-relation` | Suspends a relation to an application offer | `cmd/juju/application/suspendrelation.go` - Pauses relation |
| `resume-relation` | Resumes a suspended relation | `cmd/juju/application/resumerelation.go` - Resumes relation |

### Cross-Model Relations

| Command | Description | Implementation |
|---------|-------------|----------------|
| `offer` | Offers application endpoints for use in other models | `cmd/juju/crossmodel/offer.go` - Creates cross-model offer |
| `remove-offer` | Removes one or more offers | `cmd/juju/crossmodel/removeoffer.go` - Deletes offer |
| `show-offer` | Shows extended information about an offer | `cmd/juju/crossmodel/showoffer.go` - Displays offer details |
| `offers` / `list-offers` | Lists shared endpoints | `cmd/juju/crossmodel/list.go` - Shows available offers |
| `find-offers` | Finds offered application endpoints | `cmd/juju/crossmodel/find.go` - Discovers offers |
| `consume` | Adds a remote offer to the model | `cmd/juju/application/consume.go` - Consumes remote offer |
| `remove-saas` | Removes consumed applications from the model | `cmd/juju/application/removesaas.go` - Removes SAAS |

### Action Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `actions` / `list-actions` | Lists actions defined for an application | `cmd/juju/action/list.go` - Shows available actions |
| `show-action` | Shows detailed information about an action | `cmd/juju/action/show.go` - Displays action spec |
| `run` | Runs an action on a specified unit | `cmd/juju/action/run.go` - Executes action |
| `cancel-task` | Cancels pending or running tasks | `cmd/juju/action/cancel.go` - Aborts action execution |
| `operations` / `list-operations` | Lists operations | `cmd/juju/action/listoperations.go` - Shows action history |
| `show-operation` | Shows results of an operation | `cmd/juju/action/showoperation.go` - Displays operation results |
| `show-task` | Shows results of a task by ID | `cmd/juju/action/showtask.go` - Displays task results |
| `exec` | Runs commands on remote targets | `cmd/juju/action/exec.go` - Executes commands on units/machines |

### Storage Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-storage` | Adds storage to a unit after deployment | `cmd/juju/storage/add.go` - Attaches storage |
| `storage` / `list-storage` | Lists storage details | `cmd/juju/storage/list.go` - Shows storage instances |
| `show-storage` | Shows storage instance information | `cmd/juju/storage/show.go` - Displays storage details |
| `remove-storage` | Removes storage from the model | `cmd/juju/storage/remove.go` - Deletes storage |
| `attach-storage` | Attaches existing storage to a unit | `cmd/juju/storage/attach.go` - Connects storage to unit |
| `detach-storage` | Detaches storage from units | `cmd/juju/storage/detach.go` - Disconnects storage |
| `import-filesystem` | Imports a filesystem into the model | `cmd/juju/storage/importfilesystem.go` - Imports external FS |
| `create-storage-pool` | Creates or defines a storage pool | `cmd/juju/storage/poolcreate.go` - Creates pool |
| `storage-pools` / `list-storage-pools` | Lists storage pools | `cmd/juju/storage/poollist.go` - Shows pools |
| `remove-storage-pool` | Removes a storage pool | `cmd/juju/storage/poolremove.go` - Deletes pool |
| `update-storage-pool` | Updates storage pool attributes | `cmd/juju/storage/poolupdate.go` - Modifies pool |

### Secret Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `secrets` / `list-secrets` | Lists secrets available in the model | `cmd/juju/secrets/list.go` - Shows secrets |
| `show-secret` | Shows details for a specific secret | `cmd/juju/secrets/show.go` - Displays secret info |
| `add-secret` | Adds a new secret | `cmd/juju/secrets/add.go` - Creates secret |
| `update-secret` | Updates an existing secret | `cmd/juju/secrets/update.go` - Modifies secret |
| `remove-secret` | Removes an existing secret | `cmd/juju/secrets/remove.go` - Deletes secret |
| `grant-secret` | Grants access to a secret | `cmd/juju/secrets/grant.go` - Authorizes access |
| `revoke-secret` | Revokes access to a secret | `cmd/juju/secrets/revoke.go` - Removes access |

### Secret Backend Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `secret-backends` / `list-secret-backends` | Lists secret backends | `cmd/juju/secretbackends/list.go` - Shows backends |
| `add-secret-backend` | Adds a secret backend to the controller | `cmd/juju/secretbackends/add.go` - Registers backend |
| `update-secret-backend` | Updates a secret backend | `cmd/juju/secretbackends/update.go` - Modifies backend |
| `remove-secret-backend` | Removes a secret backend | `cmd/juju/secretbackends/remove.go` - Deletes backend |
| `show-secret-backend` | Displays the specified secret backend | `cmd/juju/secretbackends/show.go` - Shows backend info |
| `model-secret-backend` | Displays or sets secret backend for a model | `cmd/juju/secretbackends/modelbackend.go` - Sets model backend |

### Network Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-space` | Adds a new network space | `cmd/juju/space/add.go` - Creates space |
| `spaces` / `list-spaces` | Lists network spaces | `cmd/juju/space/list.go` - Shows spaces |
| `show-space` | Shows information about a network space | `cmd/juju/space/show.go` - Displays space details |
| `remove-space` | Removes a network space | `cmd/juju/space/remove.go` - Deletes space |
| `rename-space` | Renames a network space | `cmd/juju/space/rename.go` - Changes space name |
| `move-to-space` | Updates a network space's CIDR | `cmd/juju/space/move.go` - Modifies space subnets |
| `reload-spaces` | Reloads spaces and subnets from substrate | `cmd/juju/space/reload.go` - Refreshes from cloud |
| `subnets` / `list-subnets` | Lists subnets known to Juju | `cmd/juju/subnet/list.go` - Shows subnets |

### User Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-user` | Adds a Juju user to a controller | `cmd/juju/user/add.go` - Creates user account |
| `remove-user` | Deletes a Juju user from a controller | `cmd/juju/user/remove.go` - Removes user |
| `users` / `list-users` | Lists Juju users | `cmd/juju/user/list.go` - Shows users |
| `show-user` | Shows information about a user | `cmd/juju/user/info.go` - Displays user details |
| `enable-user` | Re-enables a previously disabled user | `cmd/juju/user/disenable.go` - Activates user |
| `disable-user` | Disables a Juju user | `cmd/juju/user/disenable.go` - Deactivates user |
| `change-user-password` | Changes password for a user | `cmd/juju/user/change_password.go` - Updates password |
| `login` | Logs a user in to a controller | `cmd/juju/user/login.go` - Authenticates user |
| `logout` | Logs a Juju user out of a controller | `cmd/juju/user/logout.go` - Ends session |
| `whoami` | Prints current login details | `cmd/juju/user/whoami.go` - Shows current identity |
| `grant` | Grants access level to a user | `cmd/juju/model/grantrevoke.go` - Authorizes user |
| `revoke` | Revokes access from a user | `cmd/juju/model/grantrevoke.go` - Removes access |
| `grant-cloud` | Grants access to a cloud | `cmd/juju/model/grantrevokecloud.go` - Authorizes cloud access |
| `revoke-cloud` | Revokes access to a cloud | `cmd/juju/model/grantrevokecloud.go` - Removes cloud access |

### SSH Key Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `add-ssh-key` | Adds a public SSH key to a model | `cmd/juju/sshkeys/addkeys.go` - Registers key |
| `ssh-keys` / `list-ssh-keys` | Lists SSH keys for a model | `cmd/juju/sshkeys/listkeys.go` - Shows keys |
| `remove-ssh-key` | Removes a public SSH key from a model | `cmd/juju/sshkeys/removekeys.go` - Deletes key |
| `import-ssh-key` | Adds SSH key from trusted source | `cmd/juju/sshkeys/importkeys.go` - Imports from GitHub/LP |

### SSH and Debugging

| Command | Description | Implementation |
|---------|-------------|----------------|
| `ssh` | Initiates SSH session or executes command | `cmd/juju/ssh/ssh.go` - Connects to machine/container |
| `scp` | Securely transfers files within a model | `cmd/juju/ssh/scp.go` - Copies files |
| `debug-log` | Displays log messages for a model | `cmd/juju/commands/debuglog.go` - Streams logs |
| `debug-hooks` | Launches tmux session to debug hooks | `cmd/juju/ssh/debughooks.go` - Debug mode |
| `debug-hook` | Alias for debug-hooks | `cmd/juju/ssh/debughooks.go` - Debug alias |
| `debug-code` | Launches tmux session to debug code | `cmd/juju/ssh/debugcode.go` - Code debug mode |

### Status and Reporting

| Command | Description | Implementation |
|---------|-------------|----------------|
| `status` | Reports model, machines, applications, units status | `cmd/juju/status/status.go` - Shows full state |
| `show-status-log` | Outputs past statuses for entity | `cmd/juju/status/history.go` - Shows status history |
| `firewall-rules` / `list-firewall-rules` | Prints firewall rules | `cmd/juju/firewall/list.go` - Shows rules |
| `set-firewall-rule` | Sets a firewall rule | `cmd/juju/firewall/set.go` - Configures rule |

### Backup Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `create-backup` | Creates a backup | `cmd/juju/backups/create.go` - Generates backup |
| `download-backup` | Downloads a backup archive file | `cmd/juju/backups/download.go` - Fetches backup |

### Resource Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `resources` / `list-resources` | Shows resources for an application or unit | `cmd/juju/resource/list.go` - Lists resources |
| `attach-resource` | Updates a resource for an application | `cmd/juju/resource/upload.go` - Uploads resource |
| `charm-resources` / `list-charm-resources` | Displays resources for a charm | `cmd/juju/resource/charmresources.go` - Shows charm resources |

### CharmHub Commands

| Command | Description | Implementation |
|---------|-------------|----------------|
| `info` | Displays detailed information about charms | `cmd/juju/charmhub/info.go` - Shows charm info |
| `find` | Queries CharmHub for available charms/bundles | `cmd/juju/charmhub/find.go` - Searches charms |
| `download` | Locates and downloads a CharmHub charm | `cmd/juju/charmhub/download.go` - Fetches charm |

### Bundle Management

| Command | Description | Implementation |
|---------|-------------|----------------|
| `export-bundle` | Exports current model as reusable bundle | `cmd/juju/model/exportbundle.go` - Generates bundle |
| `diff-bundle` | Compares bundle with model and reports differences | `cmd/juju/application/diffbundle.go` - Shows diff |

### Operation Protection

| Command | Description | Implementation |
|---------|-------------|----------------|
| `enable-command` | Enables previously disabled commands | `cmd/juju/block/enable.go` - Unblocks command |
| `disable-command` | Disables commands for the model | `cmd/juju/block/disable.go` - Blocks command |
| `disabled-commands` / `list-disabled-commands` | Lists disabled commands | `cmd/juju/block/list.go` - Shows blocked commands |

### Utility Commands

| Command | Description | Implementation |
|---------|-------------|----------------|
| `version` | Prints Juju CLI client version | `cmd/juju/commands/version.go` - Shows version |
| `help` | Shows help on a command or topic | `cmd/cmd/supercommand.go` - Help system |
| `documentation` | Generates documentation for all commands | `cmd/cmd/documentation.go` - Doc generator |
| `dashboard` | Prints or opens Juju Dashboard | `cmd/juju/dashboard/dashboard.go` - Dashboard access |
| `sync-agent-binary` | Copies agent binaries to local controller | `cmd/juju/commands/synctools.go` - Syncs tools |
| `set-credential` | Relates a remote credential to a model | `cmd/juju/model/setcredential.go` - Sets model credential |

### Developer Commands (Feature Flagged)

| Command | Description | Implementation |
|---------|-------------|----------------|
| `dump` | Dumps model state (developer mode) | `cmd/juju/model/dump.go` - Debug dump |
| `dump-db` | Dumps database state (developer mode) | `cmd/juju/model/dumpdb.go` - DB dump |

### Help Commands

| Command | Description | Implementation |
|---------|-------------|----------------|
| `help-hook-commands` | Shows help on Juju charm hook commands | `cmd/juju/commands/help_hook_commands.go` - Hook help |
| `help-action-commands` | Shows help on Juju charm action commands | `cmd/juju/commands/help_action_commands.go` - Action help |

## Command Aliases

Many commands have list-* aliases for backward compatibility:

| Primary Command | Alias Commands |
|-----------------|----------------|
| `actions` | `list-actions` |
| `clouds` | `list-clouds` |
| `controllers` | `list-controllers` |
| `credentials` | `list-credentials` |
| `disabled-commands` | `list-disabled-commands` |
| `firewall-rules` | `list-firewall-rules` |
| `machines` | `list-machines` |
| `models` | `list-models` |
| `offers` | `list-offers` |
| `operations` | `list-operations` |
| `regions` | `list-regions` |
| `resources` | `list-resources` |
| `secret-backends` | `list-secret-backends` |
| `secrets` | `list-secrets` |
| `spaces` | `list-spaces` |
| `ssh-keys` | `list-ssh-keys` |
| `storage` | `list-storage` |
| `storage-pools` | `list-storage-pools` |
| `subnets` | `list-subnets` |
| `users` | `list-users` |
| `integrate` | `relate` |
| `resolved` | `resolve` |
| `debug-hooks` | `debug-hook` |

## Command Naming Patterns

### Verb-Noun Pattern (Dominant)

Most commands follow the pattern: `<verb>-<noun>`

- **add-**: Creates new resources (add-model, add-user, add-machine)
- **remove-**: Deletes resources (remove-application, remove-unit, remove-cloud)
- **show-**: Displays detailed information (show-model, show-application, show-unit)
- **list-**: Enumerates resources (list-models, list-users - now deprecated aliases)
- **update-**: Modifies existing resources (update-cloud, update-credential)
- **set-**: Sets configuration values (set-constraints, set-firewall-rule)
- **grant-**: Authorizes access (grant, grant-cloud, grant-secret)
- **revoke-**: Removes access (revoke, revoke-cloud, revoke-secret)
- **enable-**: Activates features (enable-user, enable-command)
- **disable-**: Deactivates features (disable-user, disable-command)

### Standalone Commands (Exceptions)

Several commands don't follow the verb-noun pattern:

- `bootstrap` - Initializes entire cloud environment
- `deploy` - Creates application from charm
- `integrate` - Creates relation between applications
- `expose`/`unexpose` - Network visibility
- `trust` - Security context
- `switch` - Context switching
- `whoami` - Identity query
- `ssh`/`scp` - Remote access
- `login`/`logout` - Authentication
- `refresh` - Charm upgrade
- `run` - Action execution

### Deprecated Aliases

The CLI maintains backward compatibility through alias registration:
- `list-*` → plural noun form (e.g., `list-models` → `models`)
- `relate` → `integrate`
- `resolve` → `resolved`
