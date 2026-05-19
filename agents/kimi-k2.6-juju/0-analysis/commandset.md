# Command Set
## Hierarchy
The `juju` CLI is a flat super-command with ~150 canonical top-level commands (plus ~40 aliases). There is no nested sub-command hierarchy beyond the top level; every command is registered directly on the `juju` super-command.
## Command List
### actions
- **Summary**: List actions defined for an application.
- **Usage**: `juju actions [options] <application>`
- **Source package**: `cmd/juju/action`
- **Aliases**: list-actions
- **See also**: run, show-action

### add-cloud
- **Summary**: Add a cloud definition to Juju.
- **Usage**: `juju add-cloud [options] <cloud name> [<cloud definition file>]`
- **Source package**: `cmd/juju/cloud`
- **See also**: clouds, update-cloud, remove-cloud, update-credential

### add-credential
- **Summary**: Adds a credential for a cloud to a local client and uploads it to a controller.
- **Usage**: `juju add-credential [options] <cloud name>`
- **Source package**: `cmd/juju/cloud`
- **See also**: credentials, remove-credential, update-credential, default-credential, default-region, autoload-credentials

### add-k8s
- **Summary**: Adds a Kubernetes endpoint and credential to Juju.
- **Usage**: `juju add-k8s [options] <k8s name>`
- **Source package**: `cmd/juju/caas`
- **See also**: remove-k8s

### add-machine
- **Summary**: Provision a new machine or assign one to the model.
- **Usage**: `juju add-machine [options] [lxd[:<machine-id>] | ssh:[<user>@]<host> | <placement>] | <private-key> | <public-key>`
- **Source package**: `cmd/juju/machine`
- **See also**: remove-machine, model-constraints, set-model-constraints

### add-model
- **Summary**: Adds a workload model.
- **Usage**: `juju add-model [options] <model name> [cloud|region|(cloud/region)]`
- **Source package**: `cmd/juju/controller`
- **See also**: model-config, model-defaults, add-credential, autoload-credentials

### add-secret
- **Summary**: Add a new secret.
- **Usage**: `juju add-secret [options] <name> [key[#base64|#file]=value...]`
- **Source package**: `cmd/juju/secrets`

### add-secret-backend
- **Summary**: Add a new secret backend to the controller.
- **Usage**: `juju add-secret-backend [options] <backend-name> <backend-type>`
- **Source package**: `cmd/juju/secretbackends`
- **See also**: secret-backends, remove-secret-backend, show-secret-backend, update-secret-backend

### add-space
- **Summary**: Add a new network space.
- **Usage**: `juju add-space [options] <name> [<CIDR1> <CIDR2> ...]`
- **Source package**: `cmd/juju/space`
- **See also**: spaces, remove-space

### add-ssh-key
- **Summary**: Adds a public SSH key to a model.
- **Usage**: `juju add-ssh-key [options] <ssh key> ...`
- **Source package**: `cmd/juju/sshkeys`
- **See also**: ssh-keys, remove-ssh-key, import-ssh-key

### add-storage
- **Summary**: Adds storage to a unit after it has been deployed.
- **Usage**: `juju add-storage [options] <unit> <storage-directive>`
- **Source package**: `cmd/juju/storage`
- **See also**: import-filesystem, storage, storage-pools

### add-unit
- **Summary**: Adds one or more units to a deployed application.
- **Usage**: `juju add-unit [options] <application name>`
- **Source package**: `cmd/juju/application`
- **See also**: remove-unit

### add-user
- **Summary**: Adds a Juju user to a controller.
- **Usage**: `juju add-user [options] <user name> [<display name>]`
- **Source package**: `cmd/juju/user`
- **See also**: register, grant, users, show-user, disable-user, enable-user, change-user-password, remove-user

### attach-resource
- **Summary**: Update a resource for an application.
- **Usage**: `juju attach-resource [options] application <resource name>=<resource>`
- **Source package**: `cmd/juju/resource`
- **See also**: resources, charm-resources

### attach-storage
- **Summary**: Attaches existing storage to a unit.
- **Usage**: `juju attach-storage [options] <unit> <storage> [<storage> ...]`
- **Source package**: `cmd/juju/storage`

### autoload-credentials
- **Summary**: Attempts to automatically detect and add credentials for a cloud.
- **Usage**: `juju autoload-credentials [options] [<cloud-type>]`
- **Source package**: `cmd/juju/cloud`
- **See also**: add-credential, credentials, default-credential, remove-credential

### bind
- **Summary**: Change bindings for a deployed application.
- **Usage**: `juju bind [options] <application> [<default-space>] [<endpoint-name>=<space> ...]`
- **Source package**: `cmd/juju/application`
- **See also**: spaces, show-space, show-application

### bootstrap
- **Summary**: Initializes a cloud environment.
- **Usage**: `juju bootstrap [options] [<cloud name>[/region] [<controller name>]]`
- **Source package**: `cmd/juju/commands`
- **See also**: add-credential, autoload-credentials, add-model, controller-config, model-config, set-constraints, show-cloud

### cancel-task
- **Summary**: Cancel pending or running tasks.
- **Usage**: `juju cancel-task [options] (<task-id>|<task-id-prefix>) [...]`
- **Source package**: `cmd/juju/action`
- **See also**: show-task

### change-user-password
- **Summary**: Changes the password for the current or specified Juju user.
- **Usage**: `juju change-user-password [options] [username]`
- **Source package**: `cmd/juju/user`
- **See also**: add-user, register

### charm-resources
- **Summary**: Display the resources for a charm in a repository.
- **Usage**: `juju charm-resources [options] <charm>`
- **Source package**: `cmd/juju/resource`
- **Aliases**: list-charm-resources
- **See also**: resources, attach-resource

### clouds
- **Summary**: Lists all clouds available to Juju.
- **Usage**: `juju clouds [options]`
- **Source package**: `cmd/juju/cloud`
- **Aliases**: list-clouds
- **See also**: add-cloud, credentials, controllers, regions, default-credential, default-region, show-cloud, update-cloud, update-public-clouds

### config
- **Summary**: Get, set, or reset configuration for a deployed application.
- **Usage**: `juju config [options] <application name> [--reset <key[,key]>] [<attribute-key>][=<value>] ...]`
- **Source package**: `cmd/juju/application`
- **See also**: deploy, status, model-config, controller-config

### constraints
- **Summary**: Displays machine constraints for an application.
- **Usage**: `juju constraints [options] <application>`
- **Source package**: `cmd/juju/application`
- **See also**: set-constraints, model-constraints, set-model-constraints

### consume
- **Summary**: Add a remote offer to the model.
- **Usage**: `juju consume [options] <remote offer path> [<local application name>]`
- **Source package**: `cmd/juju/application`
- **See also**: integrate, offer, remove-saas

### controller-config
- **Summary**: Displays or sets configuration settings for a controller.
- **Usage**: `juju controller-config [options] [<attribute key>[=<value>] ...]`
- **Source package**: `cmd/juju/controller`
- **See also**: controllers, model-config, show-cloud

### controllers
- **Summary**: Lists all controllers.
- **Usage**: `juju controllers [options]`
- **Source package**: `cmd/juju/controller`
- **Aliases**: list-controllers
- **See also**: models, show-controller

### create-backup
- **Summary**: Create a backup.
- **Usage**: `juju create-backup [options] [<notes>]`
- **Source package**: `cmd/juju/backups`
- **See also**: download-backup

### create-storage-pool
- **Summary**: Create or define a storage pool.
- **Usage**: `juju create-storage-pool [options] <name> <storage provider> [<key>=<value> [<key>=<value>...]]`
- **Source package**: `cmd/juju/storage`
- **See also**: remove-storage-pool, update-storage-pool, storage-pools

### credentials
- **Summary**: Lists Juju credentials for a cloud.
- **Usage**: `juju credentials [options] [<cloud name>]`
- **Source package**: `cmd/juju/cloud`
- **Aliases**: list-credentials
- **See also**: add-credential, update-credential, remove-credential, default-credential, autoload-credentials, show-credential

### dashboard
- **Summary**: Print the Juju Dashboard URL, or open the Juju Dashboard in the default browser.
- **Usage**: `juju dashboard [options]`
- **Source package**: `cmd/juju/dashboard`

### debug-code
- **Summary**: Launch a tmux session to debug hooks and/or actions.
- **Usage**: `juju debug-code [options] <unit name> [hook or action names]`
- **Source package**: `cmd/juju/ssh`
- **See also**: ssh, debug-hooks

### debug-hooks
- **Summary**: Launch a tmux session to debug hooks and/or actions.
- **Usage**: `juju debug-hooks [options] <unit name> [hook or action names]`
- **Source package**: `cmd/juju/ssh`
- **Aliases**: debug-hook
- **See also**: ssh, debug-code

### debug-log
- **Summary**: Displays log messages for a model.
- **Usage**: `juju debug-log [options]`
- **Source package**: `cmd/juju/commands`
- **See also**: status, ssh

### default-credential
- **Summary**: Gets, sets, or unsets the default credential for a cloud on this client.
- **Usage**: `juju default-credential [options] <cloud name> [<credential name>]`
- **Source package**: `cmd/juju/cloud`
- **Aliases**: set-default-credentials
- **See also**: credentials, add-credential, remove-credential, autoload-credentials

### default-region
- **Summary**: Gets, sets, or unsets the default region for a cloud on this client.
- **Usage**: `juju default-region [options] <cloud name> [<region>]`
- **Source package**: `cmd/juju/cloud`
- **Aliases**: set-default-region
- **See also**: add-credential

### deploy
- **Summary**: Deploys a new application or bundle.
- **Usage**: `juju deploy [options] <charm or bundle> [<application name>]`
- **Source package**: `cmd/juju/application`
- **See also**: integrate, add-unit, config, expose, constraints, refresh, set-constraints, spaces, charm-resources

### destroy-controller
- **Summary**: Destroys a controller.
- **Usage**: `juju destroy-controller [options] <controller name>`
- **Source package**: `cmd/juju/controller`
- **See also**: kill-controller, unregister

### destroy-model
- **Summary**: Terminate all machines/containers and resources for a non-controller model.
- **Usage**: `juju destroy-model [options] [<controller name>:]<model name>`
- **Source package**: `cmd/juju/model`
- **See also**: destroy-controller

### detach-storage
- **Summary**: Detaches storage from units.
- **Usage**: `juju detach-storage [options] <storage> [<storage> ...]`
- **Source package**: `cmd/juju/storage`
- **See also**: storage, attach-storage

### diff-bundle
- **Summary**: Compares a bundle with a model and reports any differences.
- **Usage**: `juju diff-bundle [options] <bundle file or name>`
- **Source package**: `cmd/juju/application`
- **See also**: deploy

### disable-command
- **Summary**: Disables commands for the model.
- **Usage**: `juju disable-command [options] <command set> [message...]`
- **Source package**: `cmd/juju/block`
- **See also**: disabled-commands, enable-command

### disable-user
- **Summary**: Disables a Juju user.
- **Usage**: `juju disable-user [options] <user name>`
- **Source package**: `cmd/juju/user`
- **See also**: users, enable-user, login

### disabled-commands
- **Summary**: Lists disabled commands.
- **Usage**: `juju disabled-commands [options]`
- **Source package**: `cmd/juju/block`
- **Aliases**: list-disabled-commands
- **See also**: disable-command, enable-command

### documentation
- **Summary**: Generate the documentation for all commands
- **Usage**: `juju documentation [options] --out <target-folder> --no-index --split --url <base-url> --discourse-ids <filepath>`
- **Source package**: `cmd/juju/unknown`

### download
- **Summary**: Locates and then downloads a Charmhub charm.
- **Usage**: `juju download [options] [options] <charm>`
- **Source package**: `cmd/juju/charmhub`
- **See also**: info, find

### download-backup
- **Summary**: Download a backup archive file.
- **Usage**: `juju download-backup [options] /full/path/to/backup/on/controller`
- **Source package**: `cmd/juju/backups`
- **See also**: create-backup

### enable-command
- **Summary**: Enable commands that had been previously disabled.
- **Usage**: `juju enable-command [options] <command set>`
- **Source package**: `cmd/juju/block`
- **See also**: disable-command, disabled-commands

### enable-destroy-controller
- **Summary**: Enable destroy-controller by removing disabled commands in the controller.
- **Usage**: `juju enable-destroy-controller [options]`
- **Source package**: `cmd/juju/controller`
- **See also**: disable-command, disabled-commands, enable-command

### enable-user
- **Summary**: Re-enables a previously disabled Juju user.
- **Usage**: `juju enable-user [options] <user name>`
- **Source package**: `cmd/juju/user`
- **See also**: users, disable-user, login

### exec
- **Summary**: Run the commands on the remote targets specified.
- **Usage**: `juju exec [options] <commands>`
- **Source package**: `cmd/juju/action`
- **See also**: run, ssh

### export-bundle
- **Summary**: Exports the current model configuration as a reusable bundle.
- **Usage**: `juju export-bundle [options]`
- **Source package**: `cmd/juju/model`

### expose
- **Summary**: Makes an application publicly available over the network.
- **Usage**: `juju expose [options] <application name>`
- **Source package**: `cmd/juju/application`
- **See also**: unexpose

### find
- **Summary**: Queries the Charmhub store for available charms or bundles.
- **Usage**: `juju find [options] [options] <query>`
- **Source package**: `cmd/juju/charmhub`
- **See also**: info, download

### find-offers
- **Summary**: Find offered application endpoints.
- **Usage**: `juju find-offers [options]`
- **Source package**: `cmd/juju/crossmodel`
- **See also**: show-offer

### firewall-rules
- **Summary**: Prints the firewall rules.
- **Usage**: `juju firewall-rules [options]`
- **Source package**: `cmd/juju/firewall`
- **Aliases**: list-firewall-rules
- **See also**: set-firewall-rule

### grant
- **Summary**: Grants access level to a Juju user for a model, controller, or application offer.
- **Usage**: `juju grant [options] <user name> <permission> [<model name> ... | <offer url> ...]`
- **Source package**: `cmd/juju/model`
- **See also**: revoke, add-user, grant-cloud

### grant-cloud
- **Summary**: Grants access level to a Juju user for a cloud.
- **Usage**: `juju grant-cloud [options] <user name> <permission> <cloud name> ...`
- **Source package**: `cmd/juju/model`
- **See also**: grant, revoke-cloud, add-user

### grant-secret
- **Summary**: Grant access to a secret.
- **Usage**: `juju grant-secret [options] <ID>|<name> <application>[,<application>...]`
- **Source package**: `cmd/juju/secrets`

### help
- **Summary**: 
- **Usage**: ``
- **Source package**: `cmd/juju/unknown`

### help-action-commands
- **Summary**: Show help on a Juju charm action command.
- **Usage**: `juju help-action-commands [action]`
- **Source package**: `cmd/juju/commands`
- **See also**: help, help-hook-commands

### help-hook-commands
- **Summary**: Show help on a Juju charm hook command.
- **Usage**: `juju help-hook-commands [hook]`
- **Source package**: `cmd/juju/commands`
- **See also**: help, help-action-commands

### import-filesystem
- **Summary**: Imports a filesystem into the model.
- **Usage**: `juju import-filesystem [options]`
- **Source package**: `cmd/juju/storage`
- **See also**: storage

### import-ssh-key
- **Summary**: Adds a public SSH key from a trusted identity source to a model.
- **Usage**: `juju import-ssh-key [options] <lp|gh>:<user identity> ...`
- **Source package**: `cmd/juju/sshkeys`
- **See also**: add-ssh-key, ssh-keys

### info
- **Summary**: Displays detailed information about CharmHub charms.
- **Usage**: `juju info [options] [options] <charm>`
- **Source package**: `cmd/juju/charmhub`
- **See also**: find, download

### integrate
- **Summary**: Integrate two applications.
- **Usage**: `juju integrate [options] <application>[:<endpoint>] <application>[:<endpoint>]`
- **Source package**: `cmd/juju/application`
- **Aliases**: relate
- **See also**: consume, find-offers, set-firewall-rule, suspend-relation

### kill-controller
- **Summary**: Forcibly terminate all machines and other associated resources for a Juju controller.
- **Usage**: `juju kill-controller [options] <controller name>`
- **Source package**: `cmd/juju/controller`
- **See also**: destroy-controller, unregister

### login
- **Summary**: Logs a user in to a controller.
- **Usage**: `juju login [options] [controller host name or alias]`
- **Source package**: `cmd/juju/user`
- **See also**: disable-user, enable-user, logout, register, unregister

### logout
- **Summary**: Logs a Juju user out of a controller.
- **Usage**: `juju logout [options]`
- **Source package**: `cmd/juju/user`
- **See also**: change-user-password, login

### machines
- **Summary**: Lists machines in a model.
- **Usage**: `juju machines [options]`
- **Source package**: `cmd/juju/machine`
- **Aliases**: list-machines
- **See also**: status

### migrate
- **Summary**: Migrate a workload model to another controller.
- **Usage**: `juju migrate [options] <model-name> <target-controller-name>`
- **Source package**: `cmd/juju/commands`
- **See also**: login, controllers, status

### model-config
- **Summary**: Displays or sets configuration values on a model.
- **Usage**: `juju model-config [options] [<model-key>[=<value>] ...]`
- **Source package**: `cmd/juju/model`
- **See also**: models, model-defaults, show-cloud, controller-config

### model-constraints
- **Summary**: Displays machine constraints for a model.
- **Usage**: `juju model-constraints [options]`
- **Source package**: `cmd/juju/model`
- **See also**: models, constraints, set-constraints, set-model-constraints

### model-defaults
- **Summary**: Displays or sets default configuration settings for new models.
- **Usage**: `juju model-defaults [options] [<model-key>[<=value>] ...]`
- **Source package**: `cmd/juju/model`
- **Aliases**: model-default
- **See also**: models, model-config

### model-secret-backend
- **Summary**: Displays or sets the secret backend for a model.
- **Usage**: `juju model-secret-backend [options] [<secret-backend-name>]`
- **Source package**: `cmd/juju/secretbackends`
- **See also**: add-secret-backend, secret-backends, remove-secret-backend, show-secret-backend, update-secret-backend

### models
- **Summary**: Lists models a user can access on a controller.
- **Usage**: `juju models [options]`
- **Source package**: `cmd/juju/controller`
- **Aliases**: list-models
- **See also**: add-model

### move-to-space
- **Summary**: Update a network space's CIDR.
- **Usage**: `juju move-to-space [options] [--format yaml|json] <name> <CIDR1> [ <CIDR2> ...]`
- **Source package**: `cmd/juju/space`
- **See also**: add-space, spaces, reload-spaces, rename-space, show-space, remove-space

### offer
- **Summary**: Offer application endpoints for use in other models.
- **Usage**: `juju offer [options] [model-name.]<application-name>:<endpoint-name>[,...] [offer-name]`
- **Source package**: `cmd/juju/crossmodel`
- **See also**: consume, integrate, remove-saas

### offers
- **Summary**: Lists shared endpoints.
- **Usage**: `juju offers [options] [<offer-name>]`
- **Source package**: `cmd/juju/crossmodel`
- **Aliases**: list-offers
- **See also**: find-offers, show-offer

### operations
- **Summary**: Lists pending, running, or completed operations for specified application, units, machines, or all.
- **Usage**: `juju operations [options]`
- **Source package**: `cmd/juju/action`
- **Aliases**: list-operations
- **See also**: run, show-operation, show-task

### refresh
- **Summary**: Refresh an application's charm.
- **Usage**: `juju refresh [options] <application>`
- **Source package**: `cmd/juju/application`
- **See also**: deploy

### regions
- **Summary**: Lists regions for a given cloud.
- **Usage**: `juju regions [options] <cloud>`
- **Source package**: `cmd/juju/cloud`
- **Aliases**: list-regions
- **See also**: add-cloud, clouds, show-cloud, update-cloud, update-public-clouds

### register
- **Summary**: Registers a controller.
- **Usage**: `juju register [options] <registration string>|<controller host name>`
- **Source package**: `cmd/juju/controller`
- **See also**: add-user, change-user-password, unregister

### reload-spaces
- **Summary**: Reloads spaces and subnets from substrate.
- **Usage**: `juju reload-spaces [options]`
- **Source package**: `cmd/juju/space`
- **See also**: spaces, add-space, show-space, move-to-space

### remove-application
- **Summary**: Remove applications from the model.
- **Usage**: `juju remove-application [options] <application> [<application>...]`
- **Source package**: `cmd/juju/application`
- **See also**: scale-application, show-application

### remove-cloud
- **Summary**: Removes a cloud from Juju.
- **Usage**: `juju remove-cloud [options] <cloud name>`
- **Source package**: `cmd/juju/cloud`
- **See also**: add-cloud, update-cloud, clouds

### remove-credential
- **Summary**: Removes Juju credentials for a cloud.
- **Usage**: `juju remove-credential [options] <cloud name> <credential name>`
- **Source package**: `cmd/juju/cloud`
- **See also**: add-credential, autoload-credentials, credentials, default-credential, set-credential, update-credential

### remove-k8s
- **Summary**: Removes a k8s cloud from Juju.
- **Usage**: `juju remove-k8s [options] <k8s name>`
- **Source package**: `cmd/juju/caas`
- **See also**: add-k8s

### remove-machine
- **Summary**: Removes one or more machines from a model.
- **Usage**: `juju remove-machine [options] <machine number> ...`
- **Source package**: `cmd/juju/machine`
- **See also**: add-machine

### remove-offer
- **Summary**: Removes one or more offers specified by their URL.
- **Usage**: `juju remove-offer [options] <offer-url> ...`
- **Source package**: `cmd/juju/crossmodel`
- **See also**: find-offers, offer

### remove-relation
- **Summary**: Removes an existing relation between two applications.
- **Usage**: `juju remove-relation [options] <application1>[:<relation name1>] <application2>[:<relation name2>] | <relation-id>`
- **Source package**: `cmd/juju/application`
- **See also**: integrate, remove-application

### remove-saas
- **Summary**: Remove consumed applications (SAAS) from the model.
- **Usage**: `juju remove-saas [options] <saas-application-name> [<saas-application-name>...]`
- **Source package**: `cmd/juju/application`
- **See also**: consume, offer

### remove-secret
- **Summary**: Remove a existing secret.
- **Usage**: `juju remove-secret [options] <ID>|<name>`
- **Source package**: `cmd/juju/secrets`

### remove-secret-backend
- **Summary**: Removes a secret backend from the controller.
- **Usage**: `juju remove-secret-backend [options] <backend-name>`
- **Source package**: `cmd/juju/secretbackends`
- **See also**: add-secret-backend, secret-backends, show-secret-backend, update-secret-backend

### remove-space
- **Summary**: Remove a network space.
- **Usage**: `juju remove-space [options] <name>`
- **Source package**: `cmd/juju/space`
- **See also**: add-space, spaces, reload-spaces, rename-space, show-space

### remove-ssh-key
- **Summary**: Removes a public SSH key (or keys) from a model.
- **Usage**: `juju remove-ssh-key [options] <ssh key id> ...`
- **Source package**: `cmd/juju/sshkeys`
- **See also**: ssh-keys, add-ssh-key, import-ssh-key

### remove-storage
- **Summary**: Removes storage from the model.
- **Usage**: `juju remove-storage [options] <storage> [<storage> ...]`
- **Source package**: `cmd/juju/storage`
- **See also**: add-storage, attach-storage, detach-storage, list-storage, show-storage, storage

### remove-storage-pool
- **Summary**: Remove an existing storage pool.
- **Usage**: `juju remove-storage-pool [options] <name>`
- **Source package**: `cmd/juju/storage`
- **See also**: create-storage-pool, update-storage-pool, storage-pools

### remove-unit
- **Summary**: Remove application units from the model.
- **Usage**: `juju remove-unit [options] <unit> [...] | <application>`
- **Source package**: `cmd/juju/application`
- **See also**: remove-application, scale-application

### remove-user
- **Summary**: Deletes a Juju user from a controller.
- **Usage**: `juju remove-user [options] <user name>`
- **Source package**: `cmd/juju/user`
- **See also**: unregister, revoke, show-user, users, disable-user, enable-user, change-user-password

### rename-space
- **Summary**: Rename a network space.
- **Usage**: `juju rename-space [options] <old-name> <new-name>`
- **Source package**: `cmd/juju/space`
- **See also**: add-space, spaces, reload-spaces, remove-space, show-space

### resolved
- **Summary**: Marks unit errors resolved and re-executes failed hooks.
- **Usage**: `juju resolved [options] [<unit> ...]`
- **Source package**: `cmd/juju/application`
- **Aliases**: resolve

### resources
- **Summary**: Show the resources for an application or unit.
- **Usage**: `juju resources [options] <application or unit>`
- **Source package**: `cmd/juju/resource`
- **Aliases**: list-resources
- **See also**: attach-resource, charm-resources

### resume-relation
- **Summary**: Resumes a suspended relation to an application offer.
- **Usage**: `juju resume-relation [options] <relation-id>[,<relation-id>]`
- **Source package**: `cmd/juju/application`
- **See also**: integrate, offers, remove-relation, suspend-relation

### retry-provisioning
- **Summary**: Retries provisioning for failed machines.
- **Usage**: `juju retry-provisioning [options] <machine> [...]`
- **Source package**: `cmd/juju/model`

### revoke
- **Summary**: Revokes access from a Juju user for a model, controller, or application offer.
- **Usage**: `juju revoke [options] <user name> <permission> [<model name> ... | <offer url> ...]`
- **Source package**: `cmd/juju/model`
- **See also**: grant

### revoke-cloud
- **Summary**: Revokes access from a Juju user for a cloud.
- **Usage**: `juju revoke-cloud [options] <user name> <permission> <cloud name> ...`
- **Source package**: `cmd/juju/model`
- **See also**: grant-cloud

### revoke-secret
- **Summary**: Revoke access to a secret.
- **Usage**: `juju revoke-secret [options] <ID>|<name> <application>[,<application>...]`
- **Source package**: `cmd/juju/secrets`

### run
- **Summary**: Run an action on a specified unit.
- **Usage**: `juju run [options] <unit> [<unit> ...] <action-name> [<key>=<value> [<key>[.<key> ...]=<value>]]`
- **Source package**: `cmd/juju/action`
- **See also**: operations, show-operation, show-task

### scale-application
- **Summary**: Set the desired number of k8s application units.
- **Usage**: `juju scale-application [options] <application> <scale>`
- **Source package**: `cmd/juju/application`
- **See also**: remove-application, add-unit, remove-unit

### scp
- **Summary**: Securely transfer files within a model.
- **Usage**: `juju scp [options] <source> <destination>`
- **Source package**: `cmd/juju/ssh`
- **See also**: ssh

### secret-backends
- **Summary**: Lists secret backends available in the controller.
- **Usage**: `juju secret-backends [options]`
- **Source package**: `cmd/juju/secretbackends`
- **Aliases**: list-secret-backends
- **See also**: add-secret-backend, remove-secret-backend, show-secret-backend, update-secret-backend

### secrets
- **Summary**: Lists secrets available in the model.
- **Usage**: `juju secrets [options]`
- **Source package**: `cmd/juju/secrets`
- **Aliases**: list-secrets
- **See also**: add-secret, remove-secret, show-secret, update-secret

### set-constraints
- **Summary**: Sets machine constraints for an application.
- **Usage**: `juju set-constraints [options] <application> <constraint>=<value> ...`
- **Source package**: `cmd/juju/application`
- **See also**: constraints, model-constraints, set-model-constraints

### set-credential
- **Summary**: Relates a remote credential to a model.
- **Usage**: `juju set-credential [options] <cloud name> <credential name>`
- **Source package**: `cmd/juju/model`
- **See also**: credentials, show-credential, update-credential

### set-firewall-rule
- **Summary**: Sets a firewall rule.
- **Usage**: `juju set-firewall-rule [options] <service-name>, --allowlist <cidr>[,<cidr>...]`
- **Source package**: `cmd/juju/firewall`
- **See also**: firewall-rules

### set-model-constraints
- **Summary**: Sets machine constraints on a model.
- **Usage**: `juju set-model-constraints [options] <constraint>=<value> ...`
- **Source package**: `cmd/juju/model`
- **See also**: models, model-constraints, constraints, set-constraints

### show-action
- **Summary**: Shows detailed information about an action.
- **Usage**: `juju show-action [options] <application> <action>`
- **Source package**: `cmd/juju/action`
- **See also**: actions, run

### show-application
- **Summary**: Displays information about an application.
- **Usage**: `juju show-application [options] <application name or alias>`
- **Source package**: `cmd/juju/application`

### show-cloud
- **Summary**: Shows detailed information for a cloud.
- **Usage**: `juju show-cloud [options] <cloud name>`
- **Source package**: `cmd/juju/cloud`
- **See also**: clouds, add-cloud, update-cloud

### show-controller
- **Summary**: Shows detailed information of a controller.
- **Usage**: `juju show-controller [options] [<controller name> ...]`
- **Source package**: `cmd/juju/controller`
- **See also**: controllers

### show-credential
- **Summary**: Shows credential information stored either on this client or on a controller.
- **Usage**: `juju show-credential [options] [<cloud name> <credential name>]`
- **Source package**: `cmd/juju/cloud`
- **Aliases**: show-credentials
- **See also**: credentials, add-credential, update-credential, remove-credential, autoload-credentials

### show-machine
- **Summary**: Show a machine's status.
- **Usage**: `juju show-machine [options] <machineID> ...`
- **Source package**: `cmd/juju/machine`
- **See also**: add-machine

### show-model
- **Summary**: Shows information about the current or specified model.
- **Usage**: `juju show-model [options] <model name>`
- **Source package**: `cmd/juju/model`
- **See also**: add-model

### show-offer
- **Summary**: Shows extended information about the offered application.
- **Usage**: `juju show-offer [options] [<controller>:]<offer url>`
- **Source package**: `cmd/juju/crossmodel`
- **See also**: find-offers

### show-operation
- **Summary**: Show results of an operation.
- **Usage**: `juju show-operation [options] <operation-id>`
- **Source package**: `cmd/juju/action`
- **See also**: run, operations, show-task

### show-secret
- **Summary**: Shows details for a specific secret.
- **Usage**: `juju show-secret [options] <ID>|<name>`
- **Source package**: `cmd/juju/secrets`
- **See also**: add-secret, update-secret, remove-secret

### show-secret-backend
- **Summary**: Displays the specified secret backend.
- **Usage**: `juju show-secret-backend [options] <backend-name>`
- **Source package**: `cmd/juju/secretbackends`
- **See also**: add-secret-backend, secret-backends, remove-secret-backend, update-secret-backend

### show-space
- **Summary**: Shows information about the network space.
- **Usage**: `juju show-space [options] <name>`
- **Source package**: `cmd/juju/space`
- **See also**: add-space, spaces, reload-spaces, rename-space, remove-space

### show-status-log
- **Summary**: Output past statuses for the specified entity.
- **Usage**: `juju show-status-log [options] <entity name>`
- **Source package**: `cmd/juju/status`
- **See also**: status

### show-storage
- **Summary**: Shows storage instance information.
- **Usage**: `juju show-storage [options] <storage ID> [...]`
- **Source package**: `cmd/juju/storage`
- **See also**: storage, attach-storage, detach-storage, remove-storage

### show-task
- **Summary**: Show results of a task by ID.
- **Usage**: `juju show-task [options] <task ID>`
- **Source package**: `cmd/juju/action`
- **See also**: cancel-task, run, operations, show-operation

### show-unit
- **Summary**: Displays information about a unit.
- **Usage**: `juju show-unit [options] <unit name>`
- **Source package**: `cmd/juju/application`
- **See also**: add-unit, remove-unit

### show-user
- **Summary**: Show information about a user.
- **Usage**: `juju show-user [options] [<user name>]`
- **Source package**: `cmd/juju/user`
- **See also**: add-user, register, users

### spaces
- **Summary**: List known spaces, including associated subnets.
- **Usage**: `juju spaces [options] [--short] [--format yaml|json] [--output <path>]`
- **Source package**: `cmd/juju/space`
- **Aliases**: list-spaces
- **See also**: add-space, reload-spaces

### ssh
- **Summary**: Initiates an SSH session or executes a command on a Juju machine or container.
- **Usage**: `juju ssh [options] <[user@]target> [openssh options] [command]`
- **Source package**: `cmd/juju/ssh`
- **See also**: scp

### ssh-keys
- **Summary**: Lists the currently known SSH keys for the current (or specified) model.
- **Usage**: `juju ssh-keys [options]`
- **Source package**: `cmd/juju/sshkeys`
- **Aliases**: list-ssh-keys
- **See also**: add-ssh-key, remove-ssh-key

### status
- **Summary**: Report the status of the model, its machines, applications and units.
- **Usage**: `juju status [options] [<selector> [...]]`
- **Source package**: `cmd/juju/status`
- **See also**: machines, show-model, show-status-log, storage

### storage
- **Summary**: Lists storage details.
- **Usage**: `juju storage [options] <filesystem|volume> ...`
- **Source package**: `cmd/juju/storage`
- **Aliases**: list-storage
- **See also**: show-storage, add-storage, remove-storage

### storage-pools
- **Summary**: List storage pools.
- **Usage**: `juju storage-pools [options]`
- **Source package**: `cmd/juju/storage`
- **Aliases**: list-storage-pools
- **See also**: create-storage-pool, remove-storage-pool

### subnets
- **Summary**: List subnets known to Juju.
- **Usage**: `juju subnets [options] [--space <name>] [--zone <name>] [--format yaml|json] [--output <path>]`
- **Source package**: `cmd/juju/subnet`
- **Aliases**: list-subnets

### suspend-relation
- **Summary**: Suspends a relation to an application offer.
- **Usage**: `juju suspend-relation [options] <relation-id>[ <relation-id>...]`
- **Source package**: `cmd/juju/application`
- **See also**: integrate, offers, remove-relation, resume-relation

### switch
- **Summary**: Selects or identifies the current controller and model.
- **Usage**: `juju switch [options] [<controller>|<model>|<controller>:|:<model>|<controller>:<model>]`
- **Source package**: `cmd/juju/commands`
- **See also**: controllers, models, show-controller

### sync-agent-binary
- **Summary**: Copy agent binaries from the official agent store into a local controller.
- **Usage**: `juju sync-agent-binary [options]`
- **Source package**: `cmd/juju/commands`
- **See also**: upgrade-controller

### trust
- **Summary**: Sets the trust status of a deployed application to true.
- **Usage**: `juju trust [options] <application name>`
- **Source package**: `cmd/juju/application`
- **See also**: config

### unexpose
- **Summary**: Removes public availability over the network for an application.
- **Usage**: `juju unexpose [options] <application name>`
- **Source package**: `cmd/juju/application`
- **See also**: expose

### unregister
- **Summary**: Unregisters a Juju controller.
- **Usage**: `juju unregister [options] <controller name>`
- **Source package**: `cmd/juju/controller`
- **See also**: destroy-controller, kill-controller, register

### update-cloud
- **Summary**: Updates cloud information available to Juju.
- **Usage**: `juju update-cloud [options] <cloud name>`
- **Source package**: `cmd/juju/cloud`
- **See also**: add-cloud, remove-cloud, clouds

### update-credential
- **Summary**: Updates a controller credential for a cloud.
- **Usage**: `juju update-credential [options] [<cloud-name> [<credential-name>]]`
- **Source package**: `cmd/juju/cloud`
- **Aliases**: update-credentials
- **See also**: add-credential, credentials, remove-credential, set-credential

### update-k8s
- **Summary**: Updates an existing Kubernetes endpoint used by Juju.
- **Usage**: `juju update-k8s [options] <k8s name>`
- **Source package**: `cmd/juju/caas`
- **See also**: add-k8s, remove-k8s

### update-public-clouds
- **Summary**: Updates public cloud information available to Juju.
- **Usage**: `juju update-public-clouds [options]`
- **Source package**: `cmd/juju/cloud`
- **See also**: clouds

### update-secret
- **Summary**: Update an existing secret.
- **Usage**: `juju update-secret [options] <ID>|<name> [key[#base64|#file]=value...]`
- **Source package**: `cmd/juju/secrets`

### update-secret-backend
- **Summary**: Update an existing secret backend on the controller.
- **Usage**: `juju update-secret-backend [options] <backend-name>`
- **Source package**: `cmd/juju/secretbackends`
- **See also**: add-secret-backend, secret-backends, remove-secret-backend, show-secret-backend

### update-storage-pool
- **Summary**: Update storage pool attributes.
- **Usage**: `juju update-storage-pool [options] <name> [<key>=<value> [<key>=<value>...]]`
- **Source package**: `cmd/juju/storage`
- **See also**: create-storage-pool, remove-storage-pool, storage-pools

### upgrade-controller
- **Summary**: Upgrades Juju on a controller.
- **Usage**: `juju upgrade-controller [options]`
- **Source package**: `cmd/juju/commands`
- **See also**: upgrade-model

### upgrade-model
- **Summary**: Upgrades Juju on all machines in a model.
- **Usage**: `juju upgrade-model [options]`
- **Source package**: `cmd/juju/commands`
- **See also**: sync-agent-binary

### users
- **Summary**: Lists Juju users allowed to connect to a controller or model.
- **Usage**: `juju users [options] [model-name]`
- **Source package**: `cmd/juju/user`
- **Aliases**: list-users
- **See also**: add-user, register, show-user, disable-user, enable-user

### version
- **Summary**: Print the Juju CLI client version.
- **Usage**: `juju version [options]`
- **Source package**: `cmd/juju/commands`
- **See also**: show-controller, show-model

### whoami
- **Summary**: Print current login details.
- **Usage**: `juju whoami [options]`
- **Source package**: `cmd/juju/user`
- **See also**: controllers, login, logout, models, users

