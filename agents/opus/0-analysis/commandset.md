# Juju Command Set

This document maps the full CLI surface area of the Juju client. All commands are top-level subcommands of the `juju` binary; there is no nested subcommand hierarchy.

## Command Count

**Total commands: ~120** (excluding aliases)

## Command Inventory by Domain

### Action & Operations (`cmd/juju/action/`)

| Command | Description |
|---|---|
| `actions` | List actions defined for an application. |
| `cancel-task` | Cancel pending or running tasks. |
| `exec` | Run commands on remote targets. |
| `operations` | Lists pending, running, or completed operations. |
| `run` | Run an action on a specified unit. |
| `show-action` | Shows detailed information about an action. |
| `show-operation` | Show results of an operation. |
| `show-task` | Show results of a task by ID. |

### Application Lifecycle (`cmd/juju/application/`)

| Command | Description |
|---|---|
| `add-unit` | Adds extra units of a deployed application. |
| `bind` | Change bindings for a deployed application. |
| `config` | Get or set configuration options for an application. |
| `constraints` | Displays machine constraints for an application. |
| `consume` | Adds a remote offer to the model. |
| `deploy` | Deploys a new application or bundle. |
| `diff-bundle` | Compares a bundle with a model. |
| `expose` | Makes an application publicly available. |
| `integrate` | Integrate two applications (add a relation). |
| `refresh` | Refresh an application's charm. |
| `remove-application` | Remove an application from the model. |
| `remove-relation` | Remove a relation between applications. |
| `remove-saas` | Remove a consumed offer (SAAS). |
| `remove-unit` | Remove application units from the model. |
| `resolved` | Marks unit errors resolved. |
| `resume-relation` | Resume a suspended relation. |
| `scale-application` | Set the desired number of k8s application units. |
| `set-constraints` | Sets machine constraints for an application. |
| `show-application` | Displays information about an application. |
| `show-unit` | Displays information about a unit. |
| `suspend-relation` | Suspend a relation. |
| `trust` | Grant an application trusted access to cloud credentials. |
| `unexpose` | Remove public access over the network for an application. |

### Backups (`cmd/juju/backups/`)

| Command | Description |
|---|---|
| `create-backup` | Create a backup. |
| `download-backup` | Download a backup archive file. |

### Block / Operation Protection (`cmd/juju/block/`)

| Command | Description |
|---|---|
| `disable-command` | Disables commands for the model. |
| `disabled-commands` | Lists disabled commands. |
| `enable-command` | Enable commands that had been previously disabled. |

### CAAS / Kubernetes (`cmd/juju/caas/`)

| Command | Description |
|---|---|
| `add-k8s` | Adds a k8s endpoint and credential to Juju. |
| `remove-k8s` | Removes a k8s endpoint from Juju. |
| `update-k8s` | Updates a k8s endpoint in Juju. |

### CharmHub (`cmd/juju/charmhub/`)

| Command | Description |
|---|---|
| `download` | Download a charm or bundle from CharmHub. |
| `find` | Find charms and bundles on CharmHub. |
| `info` | Display information about a charm or bundle. |

### Cloud & Credentials (`cmd/juju/cloud/`)

| Command | Description |
|---|---|
| `add-cloud` | Adds a user-defined cloud to Juju. |
| `add-credential` | Adds or replaces credentials for a cloud. |
| `autoload-credentials` | Detects and loads credentials from environment. |
| `clouds` | Lists all clouds available to Juju. |
| `credentials` | Lists credentials available to Juju. |
| `default-credential` | Sets the default credential for a cloud. |
| `default-region` | Sets the default region for a cloud. |
| `regions` | Lists regions for a given cloud. |
| `remove-cloud` | Removes a cloud from Juju. |
| `remove-credential` | Removes credentials for a cloud. |
| `show-cloud` | Shows detailed information for a cloud. |
| `show-credential` | Shows credential information. |
| `update-cloud` | Updates cloud information available to Juju. |
| `update-credential` | Updates a credential for a cloud. |
| `update-public-clouds` | Updates public cloud information. |

### Controller Management (`cmd/juju/controller/`)

| Command | Description |
|---|---|
| `add-model` | Adds a workload model. |
| `controller-config` | Displays or sets configuration settings for a controller. |
| `controllers` | Lists all controllers. |
| `destroy-controller` | Destroys a controller. |
| `enable-destroy-controller` | Enable destroy-controller by removing disabled commands. |
| `kill-controller` | Forcibly terminate all resources for a controller. |
| `list-models` | Lists models a user can access on a controller. |
| `models` | Lists models a user can access on a controller. |
| `register` | Registers a Juju controller. |
| `show-controller` | Shows detailed information about a controller. |
| `unregister` | Unregisters a Juju controller. |

### Cross-Model Relations (`cmd/juju/crossmodel/`)

| Command | Description |
|---|---|
| `consume` | Adds a remote offer to the model. |
| `find-offers` | Find offered application endpoints. |
| `offer` | Offer application endpoints for use in other models. |
| `offers` | Lists offers made by this model. |
| `remove-offer` | Removes one or more offers. |
| `show-offer` | Shows extended information about an offered application. |

### Dashboard (`cmd/juju/dashboard/`)

| Command | Description |
|---|---|
| `dashboard` | Print the Juju Dashboard URL, or open it in the default browser. |

### Firewall (`cmd/juju/firewall/`)

| Command | Description |
|---|---|
| `firewall-rules` | Lists firewall rules. |
| `set-firewall-rule` | Sets a firewall rule. |

### Machine Management (`cmd/juju/machine/`)

| Command | Description |
|---|---|
| `add-machine` | Provision a new machine or assign one to the model. |
| `machines` | Lists machines in a model. |
| `remove-machine` | Removes one or more machines from a model. |
| `show-machine` | Show a machine's status. |

### Model Management (`cmd/juju/model/`)

| Command | Description |
|---|---|
| `destroy-model` | Terminate all machines/containers and resources for a model. |
| `dump-db` | Displays the mongo documents for the model. |
| `dump-model` | Displays the database-agnostic representation of the model. |
| `export-bundle` | Exports the current model configuration as a reusable bundle. |
| `grant` | Grants access to a Juju user for a model. |
| `grant-cloud` | Grants access to a cloud. |
| `model-config` | Displays or sets configuration settings for a model. |
| `model-constraints` | Displays machine constraints for a model. |
| `model-defaults` | Displays or sets default configuration for new models. |
| `model-secret-backend` | Displays or sets the secret backend for a model. |
| `retry-provisioning` | Retries provisioning for failed machines. |
| `revoke` | Revokes access from a Juju user for a model. |
| `revoke-cloud` | Revokes access to a cloud. |
| `set-credential` | Relates a remote credential to a model. |
| `set-model-constraints` | Sets machine constraints for a model. |
| `show-model` | Shows information about the current or specified model. |

### Resources (`cmd/juju/resource/`)

| Command | Description |
|---|---|
| `attach-resource` | Update a resource for an application. |
| `resources` | Show the resources for an application or unit. |
| `charm-resources` | Display the resources for a charm in a repository. |

### Secret Backends (`cmd/juju/secretbackends/`)

| Command | Description |
|---|---|
| `add-secret-backend` | Add a new secret backend to the controller. |
| `model-secret-backend` | Displays or sets the secret backend for a model. |
| `remove-secret-backend` | Removes a secret backend from the controller. |
| `secret-backends` | Lists secret backends available in the controller. |
| `show-secret-backend` | Displays the specified secret backend. |
| `update-secret-backend` | Update an existing secret backend on the controller. |

### Secrets (`cmd/juju/secrets/`)

| Command | Description |
|---|---|
| `add-secret` | Add a new secret. |
| `grant-secret` | Grant access to a secret. |
| `remove-secret` | Remove an existing secret. |
| `revoke-secret` | Revoke access to a secret. |
| `secrets` | Lists secrets available in the model. |
| `show-secret` | Shows details for a specific secret. |
| `update-secret` | Update an existing secret. |

### Spaces & Networking (`cmd/juju/space/`)

| Command | Description |
|---|---|
| `add-space` | Add a new network space. |
| `move-to-space` | Update a network space's CIDR. |
| `reload-spaces` | Reloads spaces and subnets from substrate. |
| `remove-space` | Remove a network space. |
| `rename-space` | Rename a network space. |
| `show-space` | Shows information about the network space. |
| `spaces` | List known spaces, including associated subnets. |

### SSH & Debugging (`cmd/juju/ssh/`)

| Command | Description |
|---|---|
| `debug-code` | Launch a tmux session to debug hooks and/or actions. |
| `debug-hooks` | Launch a tmux session to debug hooks and/or actions. |
| `scp` | Securely copy files to/from units. |
| `ssh` | Open an SSH session to a unit or machine. |

### SSH Keys (`cmd/juju/sshkeys/`)

| Command | Description |
|---|---|
| `add-ssh-key` | Adds a public SSH key to a model. |
| `import-ssh-key` | Imports a public SSH key from a provider. |
| `remove-ssh-key` | Removes a public SSH key from a model. |
| `ssh-keys` | Lists the public SSH keys in a model. |

### Status & History (`cmd/juju/status/`)

| Command | Description |
|---|---|
| `show-status-log` | Output past statuses for the specified entity. |
| `status` | Displays the current status of Juju, applications, and units. |

### Storage (`cmd/juju/storage/`)

| Command | Description |
|---|---|
| `add-storage` | Adds storage to a unit after it has been deployed. |
| `attach-storage` | Attaches existing storage to a unit. |
| `create-storage-pool` | Create or define a storage pool. |
| `detach-storage` | Detaches storage from units. |
| `import-filesystem` | Imports a filesystem into the model. |
| `remove-storage` | Removes storage from the model. |
| `remove-storage-pool` | Remove an existing storage pool. |
| `show-storage` | Shows storage instance information. |
| `storage` | Lists storage details. |
| `storage-pools` | List storage pools. |
| `update-storage-pool` | Update storage pool attributes. |

### Subnets (`cmd/juju/subnet/`)

| Command | Description |
|---|---|
| `subnets` | List subnets known to Juju. |

### User Management (`cmd/juju/user/`)

| Command | Description |
|---|---|
| `add-user` | Adds a Juju user to a controller. |
| `change-user-password` | Changes the password for a user. |
| `disable-user` | Disables a Juju user. |
| `enable-user` | Re-enables a disabled Juju user. |
| `login` | Logs a user in to a controller. |
| `logout` | Logs a Juju user out of a controller. |
| `remove-user` | Removes a Juju user from a controller. |
| `show-user` | Show information about a user. |
| `users` | Lists users in a controller. |
| `whoami` | Print current login details. |

### Top-Level / Bootstrap (`cmd/juju/commands/`)

| Command | Description |
|---|---|
| `bootstrap` | Initializes a cloud environment. |
| `debug-log` | Displays the Juju debug log. |
| `help-action-commands` | Show help on a Juju charm action command. |
| `help-hook-commands` | Show help on a Juju charm hook command. |
| `juju` | Enter an interactive shell for running Juju commands. |
| `migrate` | Migrate a workload model to another controller. |
| `switch` | Selects or identifies the current controller and model. |
| `sync-agent-binary` | Copy agent binaries into a local controller. |
| `upgrade-controller` | Upgrades Juju on a controller. |
| `upgrade-model` | Upgrades Juju on a model. |
| `version` | Print the Juju CLI client version. |

## Command Aliases

| Alias | Primary Command |
|---|---|
| `list-actions` | `actions` |
| `list-controllers` | `controllers` |
| `list-disabled-commands` | `disabled-commands` |
| `list-firewall-rules` | `firewall-rules` |
| `list-machines` | `machines` |
| `list-regions` | `regions` |
| `list-secret-backends` | `secret-backends` |
| `list-secrets` | `secrets` |
| `list-resources` | `resources` |
| `list-subnets` | `subnets` |
| `relate` | `integrate` |
| `resolve` | `resolved` |
| `debug-hook` | `debug-hooks` |
| `set-default-credentials` | `default-credential` |
| `set-default-region` | `default-region` |
| `update-credentials` | `update-credential` |

## How Commands Work

Each command implements the `cmd.Command` interface:

1. **Info()** returns metadata: Name, Purpose, Doc, Aliases, Examples
2. **SetFlags()** registers gnuflag flags
3. **Init()** parses positional arguments
4. **Run()** executes the command logic, typically:
   - Resolves controller/model context
   - Opens an API connection
   - Performs RPC calls
   - Formats and writes output

Commands that need model context embed `modelcmd.ModelCommandBase`, which automatically handles controller/model resolution from the client store or command-line flags (`-c`, `-m`).
