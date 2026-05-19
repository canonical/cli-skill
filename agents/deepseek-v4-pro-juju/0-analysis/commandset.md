# Command Set — Juju CLI

## Overview

The `juju` CLI surfaces approximately 130 top-level commands. All commands are registered into a single flat namespace (no subcommand nesting). Commands are organized by domain in the source tree but appear as peers at the CLI level.

## Full Command List

### Bootstrap & Platform Setup

| Command | Description | Code Path |
|---------|-------------|-----------|
| `bootstrap` | Initializes a cloud environment by creating a controller on a chosen cloud. | `cmd/juju/commands/bootstrap.go` → `environs/bootstrap` |
| `add-k8s` | Adds a Kubernetes cloud endpoint and credential to Juju. | `cmd/juju/caas/add.go` |
| `update-k8s` | Updates an existing Kubernetes cloud endpoint. | `cmd/juju/caas/update.go` |
| `remove-k8s` | Removes a Kubernetes cloud from Juju. | `cmd/juju/caas/remove.go` |

### Cloud Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-cloud` | Adds a user-defined cloud to Juju. | `cmd/juju/cloud/add.go` → `cloudfile.WritePersonalCloudMetadata` |
| `update-cloud` | Updates an existing cloud definition. | `cmd/juju/cloud/updatecloud.go` |
| `remove-cloud` | Removes a user-defined cloud. | `cmd/juju/cloud/remove.go` |
| `clouds` | Lists all known clouds (public + user-defined). | `cmd/juju/cloud/list.go` |
| `show-cloud` | Shows details of a specific cloud. | `cmd/juju/cloud/show.go` |
| `regions` | Lists regions for a given cloud. | `cmd/juju/cloud/regions.go` |
| `update-public-clouds` | Updates cached public cloud metadata. | `cmd/juju/cloud/updatepublicclouds.go` |

### Credential Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-credential` | Adds or replaces credentials for a cloud. | `cmd/juju/cloud/addcredential.go` |
| `update-credential` | Updates credentials for a cloud. | `cmd/juju/cloud/updatecredential.go` |
| `remove-credential` | Removes credentials for a cloud. | `cmd/juju/cloud/removecredential.go` |
| `credentials` | Lists all stored credentials. | `cmd/juju/cloud/listcredentials.go` |
| `show-credential` | Shows credential details. | `cmd/juju/cloud/showcredential.go` |
| `default-credential` | Sets the default credential for a cloud. | `cmd/juju/cloud/defaultcredential.go` |
| `autoload-credentials` | Detects and imports cloud credentials from the local environment. | `cmd/juju/cloud/detectcredentials.go` |

### Region Configuration

| Command | Description | Code Path |
|---------|-------------|-----------|
| `default-region` | Sets the default region for a cloud. | `cmd/juju/cloud/defaultregion.go` |

### Controller Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `controllers` | Lists all known controllers. | `cmd/juju/controller/listcontrollers.go` |
| `show-controller` | Shows details of a controller. | `cmd/juju/controller/showcontroller.go` |
| `register` | Registers a controller with Juju (from invitation URL). | `cmd/juju/controller/register.go` |
| `unregister` | Removes a controller from the local cache. | `cmd/juju/controller/unregister.go` |
| `destroy-controller` | Destroys a controller and all its models. | `cmd/juju/controller/destroy.go` |
| `kill-controller` | Forces destruction of an unreachable controller. | `cmd/juju/controller/kill.go` |
| `enable-destroy-controller` | Re-enables controller destruction if previously blocked. | `cmd/juju/controller/enabledestroy.go` |
| `controller-config` | Gets or sets controller configuration. | `cmd/juju/controller/config.go` |
| `migrate` | Migrates a model from one controller to another. | `cmd/juju/commands/migrate.go` |
| `upgrade-controller` | Upgrades the controller to a newer Juju version. | `cmd/juju/commands/upgradecontroller.go` |

### Model Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-model` | Adds a new workload model to a controller. | `cmd/juju/controller/addmodel.go` |
| `models` | Lists models a user can access on a controller. | `cmd/juju/controller/listmodels.go` |
| `show-model` | Displays details of the current model. | `cmd/juju/model/show.go` |
| `destroy-model` | Destroys a model and its resources. | `cmd/juju/model/destroy.go` |
| `switch` | Selects or identifies the current controller and model. | `cmd/juju/commands/switch.go` |
| `migrate` | Migrates a model to another controller. | `cmd/juju/commands/migrate.go` |
| `model-config` | Gets or sets model configuration attributes. | `cmd/juju/model/config.go` |
| `model-defaults` | Gets or sets default model configuration. | `cmd/juju/model/defaults.go` |
| `model-constraints` | Shows model-level constraints. | `cmd/juju/model/constraints.go` (get) |
| `set-model-constraints` | Sets model-level constraints. | `cmd/juju/model/constraints.go` (set) |
| `export-bundle` | Exports the current model as a reusable bundle. | `cmd/juju/model/exportbundle.go` |
| `retry-provisioning` | Retries provisioning of failed machines. | `cmd/juju/model/retryprovisioning.go` |
| `set-credential` | Sets the credential for a model. | `cmd/juju/model/setcredential.go` |
| `upgrade-model` | Upgrades a model to a newer Juju version. | `cmd/juju/commands/upgrademodel.go` |
| `dump-model` | Dumps model database (developer mode). | `cmd/juju/model/dump.go` |
| `dump-db` | Dumps the controller database (developer mode). | `cmd/juju/model/dumpdb.go` |

### User & Access Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-user` | Adds a new user to a controller. | `cmd/juju/user/add.go` |
| `show-user` | Shows information about a user. | `cmd/juju/user/info.go` |
| `users` | Lists all users known to the controller. | `cmd/juju/user/list.go` |
| `change-user-password` | Changes the password for a user. | `cmd/juju/user/change_password.go` |
| `disable-user` | Disables a user account. | `cmd/juju/user/disenable.go` (disable) |
| `enable-user` | Re-enables a disabled user account. | `cmd/juju/user/disenable.go` (enable) |
| `remove-user` | Removes a user from the controller. | `cmd/juju/user/remove.go` |
| `login` | Logs in to a controller. | `cmd/juju/user/login.go` |
| `logout` | Logs out from a controller. | `cmd/juju/user/logout.go` |
| `whoami` | Displays the current controller, model, and logged-in user name. | `cmd/juju/user/whoami.go` |

### Permission Management (grant/revoke)

| Command | Description | Code Path |
|---------|-------------|-----------|
| `grant` | Grants access to a user for a model. | `cmd/juju/model/grantrevoke.go` |
| `revoke` | Revokes access from a user for a model. | `cmd/juju/model/grantrevoke.go` |
| `grant-cloud` | Grants access to a user for a cloud. | `cmd/juju/model/grantrevokecloud.go` |
| `revoke-cloud` | Revokes access from a user for a cloud. | `cmd/juju/model/grantrevokecloud.go` |

### Application Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `deploy` | Deploys a new application or bundle from Charmhub or local charm. | `cmd/juju/application/deploy.go` → `deployer.Deployer` |
| `add-unit` | Adds extra units of a deployed application. | `cmd/juju/application/addunit.go` |
| `remove-unit` | Removes application units. | `cmd/juju/application/removeunit.go` |
| `remove-application` | Removes a deployed application entirely. | `cmd/juju/application/removeapplication.go` |
| `config` | Gets or sets application configuration. | `cmd/juju/application/config.go` |
| `show-application` | Shows details of a deployed application. | `cmd/juju/application/show.go` |
| `show-unit` | Shows details of a specific unit. | `cmd/juju/application/showunit.go` |
| `expose` | Makes an application publicly available over the network. | `cmd/juju/application/expose.go` |
| `unexpose` | Removes public availability for an application. | `cmd/juju/application/unexpose.go` |
| `refresh` | Refreshes an application to a new charm revision or channel. | `cmd/juju/application/refresh.go` → `refresher.Refresher` |
| `trust` | Grants an application access to its own cloud credentials. | `cmd/juju/application/trust.go` |
| `bind` | Binds an application's endpoints to specific spaces. | `cmd/juju/application/bind.go` |
| `scale-application` | Scales a Kubernetes application. | `cmd/juju/application/scaleapplication.go` |
| `diff-bundle` | Compares a bundle to the current model state. | `cmd/juju/application/diffbundle.go` |
| `resolved` | Marks unit errors as resolved, optionally re-executing failed hooks. | `cmd/juju/application/resolved.go` |
| `constraints` | Shows application-level constraints. | `cmd/juju/application/constraints.go` (get) |
| `set-constraints` | Sets application-level constraints. | `cmd/juju/application/constraints.go` (set) |

### Relation/Integration Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `integrate` | Creates a relation between two applications (alias: `relate`). | `cmd/juju/application/integrate.go` |
| `remove-relation` | Removes a relation between applications. | `cmd/juju/application/removerelation.go` |
| `suspend-relation` | Suspends a relation. | `cmd/juju/application/suspendrelation.go` |
| `resume-relation` | Resumes a suspended relation. | `cmd/juju/application/resumerelation.go` |

### Cross-Model Relations (CMR)

| Command | Description | Code Path |
|---------|-------------|-----------|
| `offer` | Offers an application endpoint for cross-model consumption. | `cmd/juju/crossmodel/offer.go` |
| `remove-offer` | Removes a cross-model offer. | `cmd/juju/crossmodel/remove.go` |
| `show-offer` | Shows details of a cross-model offer. | `cmd/juju/crossmodel/show.go` |
| `offers` | Lists all offers in the model. | `cmd/juju/crossmodel/list.go` |
| `find-offers` | Searches for offers across controllers. | `cmd/juju/crossmodel/find.go` |
| `consume` | Consumes a remote application offer (creates SAAS). | `cmd/juju/application/consume.go` |
| `remove-saas` | Removes a consumed SAAS application. | `cmd/juju/application/removesaas.go` |

### Machine Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-machine` | Adds a new machine to the model. | `cmd/juju/machine/add.go` |
| `remove-machine` | Removes a machine from the model. | `cmd/juju/machine/remove.go` |
| `machines` | Lists machines in the model. | `cmd/juju/machine/list.go` |
| `show-machine` | Shows details of a specific machine. | `cmd/juju/machine/show.go` |

### Storage Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-storage` | Adds storage to a unit. | `cmd/juju/storage/add.go` |
| `attach-storage` | Attaches existing storage to a unit. | `cmd/juju/storage/attach.go` |
| `detach-storage` | Detaches storage from a unit. | `cmd/juju/storage/detach.go` |
| `remove-storage` | Removes storage from the model. | `cmd/juju/storage/remove.go` |
| `show-storage` | Shows storage details. | `cmd/juju/storage/show.go` |
| `storage` | Lists all storage in the model. | `cmd/juju/storage/list.go` |
| `import-filesystem` | Imports an existing filesystem into the model. | `cmd/juju/storage/import.go` |
| `create-storage-pool` | Creates a storage pool. | `cmd/juju/storage/poolcreate.go` |
| `update-storage-pool` | Updates a storage pool configuration. | `cmd/juju/storage/poolupdate.go` |
| `remove-storage-pool` | Removes a storage pool. | `cmd/juju/storage/pooldelete.go` |
| `storage-pools` | Lists storage pools. | `cmd/juju/storage/poollist.go` |

### Space & Network Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-space` | Adds a new network space. | `cmd/juju/space/add.go` |
| `remove-space` | Removes a network space. | `cmd/juju/space/remove.go` |
| `spaces` | Lists all spaces. | `cmd/juju/space/list.go` |
| `show-space` | Shows details of a space. | `cmd/juju/space/show.go` |
| `rename-space` | Renames a space. | `cmd/juju/space/rename.go` |
| `move-to-space` | Moves existing subnets to a space. | `cmd/juju/space/move.go` |
| `reload-spaces` | Reloads space definitions from the substrate. | `cmd/juju/space/reload.go` |
| `subnets` | Lists subnets. | `cmd/juju/subnet/list.go` |

### Secret Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-secret` | Adds a new secret. | `cmd/juju/secrets/add.go` |
| `update-secret` | Updates an existing secret. | `cmd/juju/secrets/update.go` |
| `remove-secret` | Removes a secret. | `cmd/juju/secrets/remove.go` |
| `secrets` | Lists secrets in the model. | `cmd/juju/secrets/list.go` |
| `show-secret` | Shows secret details or content. | `cmd/juju/secrets/show.go` |
| `grant-secret` | Grants access to a secret. | `cmd/juju/secrets/grantrevoke.go` (grant) |
| `revoke-secret` | Revokes access to a secret. | `cmd/juju/secrets/grantrevoke.go` (revoke) |

### Secret Backend Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-secret-backend` | Adds a new secret backend. | `cmd/juju/secretbackends/add.go` |
| `update-secret-backend` | Updates a secret backend. | `cmd/juju/secretbackends/update.go` |
| `remove-secret-backend` | Removes a secret backend. | `cmd/juju/secretbackends/remove.go` |
| `secret-backends` | Lists secret backends. | `cmd/juju/secretbackends/list.go` |
| `show-secret-backend` | Shows details of a secret backend. | `cmd/juju/secretbackends/show.go` |
| `model-secret-backend` | Gets or sets the model's default secret backend. | `cmd/juju/secretbackends/modelsecretbackend.go` |

### Actions & Operations

| Command | Description | Code Path |
|---------|-------------|-----------|
| `run` | Runs a command on target units. | `cmd/juju/action/run.go` |
| `exec` | Executes a command inside a unit's workload. | `cmd/juju/action/exec.go` |
| `actions` | Lists available actions for an application. | `cmd/juju/action/list.go` |
| `show-action` | Shows details of an action definition. | `cmd/juju/action/show.go` |
| `operations` | Lists all action operations. | `cmd/juju/action/listoperations.go` |
| `show-operation` | Shows the results of an operation. | `cmd/juju/action/showoperation.go` |
| `show-task` | Shows the results of a single task within an operation. | `cmd/juju/action/showtask.go` |
| `cancel-task` | Cancels a pending or running task. | `cmd/juju/action/cancel.go` |

### SSH & Debugging

| Command | Description | Code Path |
|---------|-------------|-----------|
| `ssh` | Opens an SSH session to a machine or container. | `cmd/juju/ssh/ssh.go` |
| `scp` | Copies files to/from a machine via SCP. | `cmd/juju/ssh/scp.go` |
| `debug-hooks` | Launches a tmux debugging session for hook execution. | `cmd/juju/ssh/debughooks.go` |
| `debug-code` | Launches a debugging session for a charm. | `cmd/juju/ssh/debugcode.go` |
| `debug-log` | Streams the controller/model debug log. | `cmd/juju/commands/debuglog.go` |

### SSH Key Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `add-ssh-key` | Adds a public SSH key to a model. | `cmd/juju/sshkeys/add_sshkeys.go` |
| `remove-ssh-key` | Removes an SSH key from a model. | `cmd/juju/sshkeys/remove_sshkeys.go` |
| `import-ssh-key` | Imports SSH keys from Launchpad or GitHub. | `cmd/juju/sshkeys/import_sshkeys.go` |
| `ssh-keys` | Lists SSH keys for a model. | `cmd/juju/sshkeys/list_sshkeys.go` |

### Status & Monitoring

| Command | Description | Code Path |
|---------|-------------|-----------|
| `status` | Displays the current status of Juju, applications, and units. | `cmd/juju/status/status.go` |
| `show-status-log` | Shows the status history for an entity. | `cmd/juju/status/history.go` |

### Firewall Rules

| Command | Description | Code Path |
|---------|-------------|-----------|
| `set-firewall-rule` | Sets a firewall rule for a service. | `cmd/juju/firewall/setrule.go` |
| `firewall-rules` | Lists firewall rules. | `cmd/juju/firewall/listrules.go` |

### Block / Operation Protection

| Command | Description | Code Path |
|---------|-------------|-----------|
| `disable-command` | Disables certain commands for the model. | `cmd/juju/block/disable.go` |
| `enable-command` | Re-enables disabled commands. | `cmd/juju/block/enable.go` |
| `disabled-commands` | Lists currently disabled commands. | `cmd/juju/block/list.go` |

### Backups

| Command | Description | Code Path |
|---------|-------------|-----------|
| `create-backup` | Creates a backup of the controller. | `cmd/juju/backups/create.go` |
| `download-backup` | Downloads a backup from the controller. | `cmd/juju/backups/download.go` |

### CharmHub Integration

| Command | Description | Code Path |
|---------|-------------|-----------|
| `find` | Queries Charmhub for available charms or bundles. | `cmd/juju/charmhub/find.go` |
| `info` | Shows detailed information about a charm or bundle. | `cmd/juju/charmhub/info.go` |
| `download` | Downloads a charm or bundle from Charmhub. | `cmd/juju/charmhub/download.go` |

### Resource Management

| Command | Description | Code Path |
|---------|-------------|-----------|
| `attach-resource` | Uploads a resource (file, OCI image) to an application. | `cmd/juju/resource/upload.go` |
| `resources` | Shows resources for an application or unit. | `cmd/juju/resource/list.go` |
| `charm-resources` | Shows resources needed by a charm. | `cmd/juju/resource/charmresources.go` |

### Dashboard

| Command | Description | Code Path |
|---------|-------------|-----------|
| `dashboard` | Opens the Juju dashboard in a browser. | `cmd/juju/dashboard/dashboard.go` |

### Tooling & Help

| Command | Description | Code Path |
|---------|-------------|-----------|
| `version` | Prints the Juju CLI client version. | `cmd/juju/commands/version.go` |
| `help-action-commands` | Lists all action commands available via `juju run`. | `cmd/juju/commands/help_action_commands.go` |
| `help-hook-commands` | Lists all hook commands available in charm hooks. | `cmd/juju/commands/help_hook_commands.go` |
| `sync-agent-binary` | Copies Juju agent binaries to the controller. | `cmd/juju/commands/synctools.go` |

### List-Only Commands (Observation)

Commands whose primary purpose is listing/observing a resource:

`actions`, `clouds`, `controllers`, `credentials`, `disabled-commands`, `firewall-rules`, `machines`, `models`, `offers`, `operations`, `regions`, `resources`, `secret-backends`, `secrets`, `spaces`, `ssh-keys`, `status`, `storage`, `storage-pools`, `subnets`, `users`

### Informational Commands

`info`, `show-action`, `show-application`, `show-cloud`, `show-controller`, `show-credential`, `show-machine`, `show-model`, `show-offer`, `show-operation`, `show-secret`, `show-secret-backend`, `show-space`, `show-status-log`, `show-storage`, `show-task`, `show-unit`, `show-user`, `whoami`, `version`

### Non-Standard / Orphan Commands

Commands that do not follow a simple verb-noun or listing pattern:

- `bootstrap` — self-initializing operation, no direct "unbootstrap" / "de-bootstrap"
- `integrate` — bidirectional verb, not a standard CRUD, alias: `relate`
- `consume` — creates a local representation of a remote offer, not a "create-saas"
- `resolved` — single verb for a state change (marks hooks as resolved)
- `whoami` — identity introspection
- `switch` — changes context (model/controller), not a noun operation
- `migrate` — moves a model between controllers
- `bind` — modifies application endpoint-to-space bindings
- `trust` — grants an application credential access, not a typical resource operation
- `dashboard` — opens a browser, not a resource operation
- `diff-bundle` — comparison/analysis, not a resource operation
- `sync-agent-binary` — internal tool for uploading binaries
- `retry-provisioning` — operational recovery command
- `debug-log` — streaming diagnostic output
- `exec` / `run` — direct execution on units, not resource management
- `ssh` / `scp` — transport commands
- `help-action-commands` / `help-hook-commands` — meta/help commands
