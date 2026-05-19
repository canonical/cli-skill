# Juju CLI Command Set

## Scope

This inventory is built from:
- top-level registration in `cmd/juju/commands/main.go`
- framework-provided built-ins from `cmd/cmd`
- generated command docs from `juju documentation --split`
- help signatures from `juju help <command>`

The source package column names the package that owns the command constructor or implementation surface.

## Built-in framework commands

| Command | Summary | Usage signature | Source package |
| --- | --- | --- | --- |
| `help` | Show help on a command or other topic. | Special-cased by the supercommand; effectively `juju help [topic]` | `cmd/cmd` |
| `documentation` | Generate the documentation for all commands | `juju documentation [options] --out <target-folder> --no-index --split --url <base-url> --discourse-ids <filepath>` | `cmd/cmd` |
| `version` | Print the Juju CLI client version. | `juju version [options]` | `cmd/juju/commands` |

## Registered command inventory

| Command | Summary | Usage signature | Source package |
| --- | --- | --- | --- |
| `actions` | List actions defined for an application. | `juju actions [options] <application>` | `cmd/juju/action` |
| `add-cloud` | Add a cloud definition to Juju. | `juju add-cloud [options] <cloud name> [<cloud definition file>]` | `cmd/juju/cloud` |
| `add-credential` | Adds a credential for a cloud to a local client and uploads it to a controller. | `juju add-credential [options] <cloud name>` | `cmd/juju/cloud` |
| `add-k8s` | Adds a Kubernetes endpoint and credential to Juju. | `juju add-k8s [options] <k8s name>` | `cmd/juju/caas` |
| `add-machine` | Provision a new machine or assign one to the model. | `juju add-machine [options] [lxd[:<machine-id>] \| ssh:[<user>@]<host> \| <placement>] \| <private-key> \| <public-key>` | `cmd/juju/machine` |
| `add-model` | Adds a workload model. | `juju add-model [options] <model name> [cloud\|region\|(cloud/region)]` | `cmd/juju/controller` |
| `add-secret` | Add a new secret. | `juju add-secret [options] <name> [key[#base64\|#file]=value...]` | `cmd/juju/secrets` |
| `add-secret-backend` | Add a new secret backend to the controller. | `juju add-secret-backend [options] <backend-name> <backend-type>` | `cmd/juju/secretbackends` |
| `add-space` | Add a new network space. | `juju add-space [options] <name> [<CIDR1> <CIDR2> ...]` | `cmd/juju/space` |
| `add-ssh-key` | Adds a public SSH key to a model. | `juju add-ssh-key [options] <ssh key> ...` | `cmd/juju/sshkeys` |
| `add-storage` | Adds storage to a unit after it has been deployed. | `juju add-storage [options] <unit> <storage-directive>` | `cmd/juju/storage` |
| `add-unit` | Adds one or more units to a deployed application. | `juju add-unit [options] <application name>` | `cmd/juju/application` |
| `add-user` | Adds a Juju user to a controller. | `juju add-user [options] <user name> [<display name>]` | `cmd/juju/user` |
| `attach-resource` | Update a resource for an application. | `juju attach-resource [options] application <resource name>=<resource>` | `cmd/juju/resource` |
| `attach-storage` | Attaches existing storage to a unit. | `juju attach-storage [options] <unit> <storage> [<storage> ...]` | `cmd/juju/storage` |
| `autoload-credentials` | Attempts to automatically detect and add credentials for a cloud. | `juju autoload-credentials [options] [<cloud-type>]` | `cmd/juju/cloud` |
| `bind` | Change bindings for a deployed application. | `juju bind [options] <application> [<default-space>] [<endpoint-name>=<space> ...]` | `cmd/juju/application` |
| `bootstrap` | Initializes a cloud environment. | `juju bootstrap [options] [<cloud name>[/region] [<controller name>]]` | `cmd/juju/commands` |
| `cancel-task` | Cancel pending or running tasks. | `juju cancel-task [options] (<task-id>\|<task-id-prefix>) [...]` | `cmd/juju/action` |
| `change-user-password` | Changes the password for the current or specified Juju user. | `juju change-user-password [options] [username]` | `cmd/juju/user` |
| `charm-resources` | Display the resources for a charm in a repository. | `juju charm-resources [options] <charm>` | `cmd/juju/resource` |
| `clouds` | Lists all clouds available to Juju. | `juju clouds [options]` | `cmd/juju/cloud` |
| `config` | Get, set, or reset configuration for a deployed application. | `juju config [options] <application name> [--reset <key[,key]>] [<attribute-key>][=<value>] ...]` | `cmd/juju/application` |
| `constraints` | Displays machine constraints for an application. | `juju constraints [options] <application>` | `cmd/juju/application` |
| `consume` | Add a remote offer to the model. | `juju consume [options] <remote offer path> [<local application name>]` | `cmd/juju/application` |
| `controller-config` | Displays or sets configuration settings for a controller. | `juju controller-config [options] [<attribute key>[=<value>] ...]` | `cmd/juju/controller` |
| `controllers` | Lists all controllers. | `juju controllers [options]` | `cmd/juju/controller` |
| `create-backup` | Create a backup. | `juju create-backup [options] [<notes>]` | `cmd/juju/backups` |
| `create-storage-pool` | Create or define a storage pool. | `juju create-storage-pool [options] <name> <storage provider> [<key>=<value> [<key>=<value>...]]` | `cmd/juju/storage` |
| `credentials` | Lists Juju credentials for a cloud. | `juju credentials [options] [<cloud name>]` | `cmd/juju/cloud` |
| `dashboard` | Print the Juju Dashboard URL, or open the Juju Dashboard in the default browser. | `juju dashboard [options]` | `cmd/juju/dashboard` |
| `debug-code` | Launch a tmux session to debug hooks and/or actions. | `juju debug-code [options] <unit name> [hook or action names]` | `cmd/juju/ssh` |
| `debug-hooks` | Launch a tmux session to debug hooks and/or actions. | `juju debug-hooks [options] <unit name> [hook or action names]` | `cmd/juju/ssh` |
| `debug-log` | Displays log messages for a model. | `juju debug-log [options]` | `cmd/juju/commands` |
| `default-credential` | Gets, sets, or unsets the default credential for a cloud on this client. | `juju default-credential [options] <cloud name> [<credential name>]` | `cmd/juju/cloud` |
| `default-region` | Gets, sets, or unsets the default region for a cloud on this client. | `juju default-region [options] <cloud name> [<region>]` | `cmd/juju/cloud` |
| `deploy` | Deploys a new application or bundle. | `juju deploy [options] <charm or bundle> [<application name>]` | `cmd/juju/application` |
| `destroy-controller` | Destroys a controller. | `juju destroy-controller [options] <controller name>` | `cmd/juju/controller` |
| `destroy-model` | Terminate all machines/containers and resources for a non-controller model. | `juju destroy-model [options] [<controller name>:]<model name>` | `cmd/juju/model` |
| `detach-storage` | Detaches storage from units. | `juju detach-storage [options] <storage> [<storage> ...]` | `cmd/juju/storage` |
| `diff-bundle` | Compares a bundle with a model and reports any differences. | `juju diff-bundle [options] <bundle file or name>` | `cmd/juju/application` |
| `disable-command` | Disables commands for the model. | `juju disable-command [options] <command set> [message...]` | `cmd/juju/block` |
| `disable-user` | Disables a Juju user. | `juju disable-user [options] <user name>` | `cmd/juju/user` |
| `disabled-commands` | Lists disabled commands. | `juju disabled-commands [options]` | `cmd/juju/block` |
| `download` | Locates and then downloads a Charmhub charm. | `juju download [options] [options] <charm>` | `cmd/juju/charmhub` |
| `download-backup` | Download a backup archive file. | `juju download-backup [options] /full/path/to/backup/on/controller` | `cmd/juju/backups` |
| `enable-command` | Enable commands that had been previously disabled. | `juju enable-command [options] <command set>` | `cmd/juju/block` |
| `enable-destroy-controller` | Enable destroy-controller by removing disabled commands in the controller. | `juju enable-destroy-controller [options]` | `cmd/juju/controller` |
| `enable-user` | Re-enables a previously disabled Juju user. | `juju enable-user [options] <user name>` | `cmd/juju/user` |
| `exec` | Run the commands on the remote targets specified. | `juju exec [options] <commands>` | `cmd/juju/action` |
| `export-bundle` | Exports the current model configuration as a reusable bundle. | `juju export-bundle [options]` | `cmd/juju/model` |
| `expose` | Makes an application publicly available over the network. | `juju expose [options] <application name>` | `cmd/juju/application` |
| `find` | Queries the Charmhub store for available charms or bundles. | `juju find [options] [options] <query>` | `cmd/juju/charmhub` |
| `find-offers` | Find offered application endpoints. | `juju find-offers [options]` | `cmd/juju/crossmodel` |
| `firewall-rules` | Prints the firewall rules. | `juju firewall-rules [options]` | `cmd/juju/firewall` |
| `grant` | Grants access level to a Juju user for a model, controller, or application offer. | `juju grant [options] <user name> <permission> [<model name> ... \| <offer url> ...]` | `cmd/juju/model` |
| `grant-cloud` | Grants access level to a Juju user for a cloud. | `juju grant-cloud [options] <user name> <permission> <cloud name> ...` | `cmd/juju/model` |
| `grant-secret` | Grant access to a secret. | `juju grant-secret [options] <ID>\|<name> <application>[,<application>...]` | `cmd/juju/secrets` |
| `help-action-commands` | Show help on a Juju charm action command. | `juju help-action-commands [action]` | `cmd/juju/commands` |
| `help-hook-commands` | Show help on a Juju charm hook command. | `juju help-hook-commands [hook]` | `cmd/juju/commands` |
| `import-filesystem` | Imports a filesystem into the model. | `juju import-filesystem [options] <storage-provider> <provider-id> <storage-name>` | `cmd/juju/storage` |
| `import-ssh-key` | Adds a public SSH key from a trusted identity source to a model. | `juju import-ssh-key [options] <lp\|gh>:<user identity> ...` | `cmd/juju/sshkeys` |
| `info` | Displays detailed information about CharmHub charms. | `juju info [options] [options] <charm>` | `cmd/juju/charmhub` |
| `integrate` | Integrate two applications. | `juju integrate [options] <application>[:<endpoint>] <application>[:<endpoint>]` | `cmd/juju/application` |
| `kill-controller` | Forcibly terminate all machines and other associated resources for a Juju controller. | `juju kill-controller [options] <controller name>` | `cmd/juju/controller` |
| `login` | Logs a user in to a controller. | `juju login [options] [controller host name or alias]` | `cmd/juju/user` |
| `logout` | Logs a Juju user out of a controller. | `juju logout [options]` | `cmd/juju/user` |
| `machines` | Lists machines in a model. | `juju machines [options]` | `cmd/juju/machine` |
| `migrate` | Migrate a workload model to another controller. | `juju migrate [options] <model-name> <target-controller-name>` | `cmd/juju/commands` |
| `model-config` | Displays or sets configuration values on a model. | `juju model-config [options] [<model-key>[=<value>] ...]` | `cmd/juju/model` |
| `model-constraints` | Displays machine constraints for a model. | `juju model-constraints [options]` | `cmd/juju/model` |
| `model-defaults` | Displays or sets default configuration settings for new models. | `juju model-defaults [options] [<model-key>[<=value>] ...]` | `cmd/juju/model` |
| `model-secret-backend` | Displays or sets the secret backend for a model. | `juju model-secret-backend [options] [<secret-backend-name>]` | `cmd/juju/secretbackends` |
| `models` | Lists models a user can access on a controller. | `juju models [options]` | `cmd/juju/controller` |
| `move-to-space` | Update a network space's CIDR. | `juju move-to-space [options] [--format yaml\|json] <name> <CIDR1> [ <CIDR2> ...]` | `cmd/juju/space` |
| `offer` | Offer application endpoints for use in other models. | `juju offer [options] [model-name.]<application-name>:<endpoint-name>[,...] [offer-name]` | `cmd/juju/crossmodel` |
| `offers` | Lists shared endpoints. | `juju offers [options] [<offer-name>]` | `cmd/juju/crossmodel` |
| `operations` | Lists pending, running, or completed operations for specified application, units, machines, or all. | `juju operations [options]` | `cmd/juju/action` |
| `refresh` | Refresh an application's charm. | `juju refresh [options] <application>` | `cmd/juju/application` |
| `regions` | Lists regions for a given cloud. | `juju regions [options] <cloud>` | `cmd/juju/cloud` |
| `register` | Registers a controller. | `juju register [options] <registration string>\|<controller host name>` | `cmd/juju/controller` |
| `reload-spaces` | Reloads spaces and subnets from substrate. | `juju reload-spaces [options]` | `cmd/juju/space` |
| `remove-application` | Remove applications from the model. | `juju remove-application [options] <application> [<application>...]` | `cmd/juju/application` |
| `remove-cloud` | Removes a cloud from Juju. | `juju remove-cloud [options] <cloud name>` | `cmd/juju/cloud` |
| `remove-credential` | Removes Juju credentials for a cloud. | `juju remove-credential [options] <cloud name> <credential name>` | `cmd/juju/cloud` |
| `remove-k8s` | Removes a k8s cloud from Juju. | `juju remove-k8s [options] <k8s name>` | `cmd/juju/caas` |
| `remove-machine` | Removes one or more machines from a model. | `juju remove-machine [options] <machine number> ...` | `cmd/juju/machine` |
| `remove-offer` | Removes one or more offers specified by their URL. | `juju remove-offer [options] <offer-url> ...` | `cmd/juju/crossmodel` |
| `remove-relation` | Removes an existing relation between two applications. | `juju remove-relation [options] <application1>[:<relation name1>] <application2>[:<relation name2>] \| <relation-id>` | `cmd/juju/application` |
| `remove-saas` | Remove consumed applications (SAAS) from the model. | `juju remove-saas [options] <saas-application-name> [<saas-application-name>...]` | `cmd/juju/application` |
| `remove-secret` | Remove a existing secret. | `juju remove-secret [options] <ID>\|<name>` | `cmd/juju/secrets` |
| `remove-secret-backend` | Removes a secret backend from the controller. | `juju remove-secret-backend [options] <backend-name>` | `cmd/juju/secretbackends` |
| `remove-space` | Remove a network space. | `juju remove-space [options] <name>` | `cmd/juju/space` |
| `remove-ssh-key` | Removes a public SSH key (or keys) from a model. | `juju remove-ssh-key [options] <ssh key id> ...` | `cmd/juju/sshkeys` |
| `remove-storage` | Removes storage from the model. | `juju remove-storage [options] <storage> [<storage> ...]` | `cmd/juju/storage` |
| `remove-storage-pool` | Remove an existing storage pool. | `juju remove-storage-pool [options] <name>` | `cmd/juju/storage` |
| `remove-unit` | Remove application units from the model. | `juju remove-unit [options] <unit> [...] \| <application>` | `cmd/juju/application` |
| `remove-user` | Deletes a Juju user from a controller. | `juju remove-user [options] <user name>` | `cmd/juju/user` |
| `rename-space` | Rename a network space. | `juju rename-space [options] <old-name> <new-name>` | `cmd/juju/space` |
| `resolved` | Marks unit errors resolved and re-executes failed hooks. | `juju resolved [options] [<unit> ...]` | `cmd/juju/application` |
| `resources` | Show the resources for an application or unit. | `juju resources [options] <application or unit>` | `cmd/juju/resource` |
| `resume-relation` | Resumes a suspended relation to an application offer. | `juju resume-relation [options] <relation-id>[,<relation-id>]` | `cmd/juju/application` |
| `retry-provisioning` | Retries provisioning for failed machines. | `juju retry-provisioning [options] <machine> [...]` | `cmd/juju/model` |
| `revoke` | Revokes access from a Juju user for a model, controller, or application offer. | `juju revoke [options] <user name> <permission> [<model name> ... \| <offer url> ...]` | `cmd/juju/model` |
| `revoke-cloud` | Revokes access from a Juju user for a cloud. | `juju revoke-cloud [options] <user name> <permission> <cloud name> ...` | `cmd/juju/model` |
| `revoke-secret` | Revoke access to a secret. | `juju revoke-secret [options] <ID>\|<name> <application>[,<application>...]` | `cmd/juju/secrets` |
| `run` | Run an action on a specified unit. | `juju run [options] <unit> [<unit> ...] <action-name> [<key>=<value> [<key>[.<key> ...]=<value>]]` | `cmd/juju/action` |
| `scale-application` | Set the desired number of k8s application units. | `juju scale-application [options] <application> <scale>` | `cmd/juju/application` |
| `scp` | Securely transfer files within a model. | `juju scp [options] <source> <destination>` | `cmd/juju/ssh` |
| `secret-backends` | Lists secret backends available in the controller. | `juju secret-backends [options]` | `cmd/juju/secretbackends` |
| `secrets` | Lists secrets available in the model. | `juju secrets [options]` | `cmd/juju/secrets` |
| `set-constraints` | Sets machine constraints for an application. | `juju set-constraints [options] <application> <constraint>=<value> ...` | `cmd/juju/application` |
| `set-credential` | Relates a remote credential to a model. | `juju set-credential [options] <cloud name> <credential name>` | `cmd/juju/model` |
| `set-firewall-rule` | Sets a firewall rule. | `juju set-firewall-rule [options] <service-name>, --allowlist <cidr>[,<cidr>...]` | `cmd/juju/firewall` |
| `set-model-constraints` | Sets machine constraints on a model. | `juju set-model-constraints [options] <constraint>=<value> ...` | `cmd/juju/model` |
| `show-action` | Shows detailed information about an action. | `juju show-action [options] <application> <action>` | `cmd/juju/action` |
| `show-application` | Displays information about an application. | `juju show-application [options] <application name or alias>` | `cmd/juju/application` |
| `show-cloud` | Shows detailed information for a cloud. | `juju show-cloud [options] <cloud name>` | `cmd/juju/cloud` |
| `show-controller` | Shows detailed information of a controller. | `juju show-controller [options] [<controller name> ...]` | `cmd/juju/controller` |
| `show-credential` | Shows credential information stored either on this client or on a controller. | `juju show-credential [options] [<cloud name> <credential name>]` | `cmd/juju/cloud` |
| `show-machine` | Show a machine's status. | `juju show-machine [options] <machineID> ...` | `cmd/juju/machine` |
| `show-model` | Shows information about the current or specified model. | `juju show-model [options] <model name>` | `cmd/juju/model` |
| `show-offer` | Shows extended information about the offered application. | `juju show-offer [options] [<controller>:]<offer url>` | `cmd/juju/crossmodel` |
| `show-operation` | Show results of an operation. | `juju show-operation [options] <operation-id>` | `cmd/juju/action` |
| `show-secret` | Shows details for a specific secret. | `juju show-secret [options] <ID>\|<name>` | `cmd/juju/secrets` |
| `show-secret-backend` | Displays the specified secret backend. | `juju show-secret-backend [options] <backend-name>` | `cmd/juju/secretbackends` |
| `show-space` | Shows information about the network space. | `juju show-space [options] <name>` | `cmd/juju/space` |
| `show-status-log` | Output past statuses for the specified entity. | `juju show-status-log [options] <entity name>` | `cmd/juju/status` |
| `show-storage` | Shows storage instance information. | `juju show-storage [options] <storage ID> [...]` | `cmd/juju/storage` |
| `show-task` | Show results of a task by ID. | `juju show-task [options] <task ID>` | `cmd/juju/action` |
| `show-unit` | Displays information about a unit. | `juju show-unit [options] <unit name>` | `cmd/juju/application` |
| `show-user` | Show information about a user. | `juju show-user [options] [<user name>]` | `cmd/juju/user` |
| `spaces` | List known spaces, including associated subnets. | `juju spaces [options] [--short] [--format yaml\|json] [--output <path>]` | `cmd/juju/space` |
| `ssh` | Initiates an SSH session or executes a command on a Juju machine or container. | `juju ssh [options] <[user@]target> [openssh options] [command]` | `cmd/juju/ssh` |
| `ssh-keys` | Lists the currently known SSH keys for the current (or specified) model. | `juju ssh-keys [options]` | `cmd/juju/sshkeys` |
| `status` | Report the status of the model, its machines, applications and units. | `juju status [options] [<selector> [...]]` | `cmd/juju/status` |
| `storage` | Lists storage details. | `juju storage [options] <filesystem\|volume> ...` | `cmd/juju/storage` |
| `storage-pools` | List storage pools. | `juju storage-pools [options]` | `cmd/juju/storage` |
| `subnets` | List subnets known to Juju. | `juju subnets [options] [--space <name>] [--zone <name>] [--format yaml\|json] [--output <path>]` | `cmd/juju/subnet` |
| `suspend-relation` | Suspends a relation to an application offer. | `juju suspend-relation [options] <relation-id>[ <relation-id>...]` | `cmd/juju/application` |
| `switch` | Selects or identifies the current controller and model. | `juju switch [options] [<controller>\|<model>\|<controller>:\|:<model>\|<controller>:<model>]` | `cmd/juju/commands` |
| `sync-agent-binary` | Copy agent binaries from the official agent store into a local controller. | `juju sync-agent-binary [options]` | `cmd/juju/commands` |
| `trust` | Sets the trust status of a deployed application to true. | `juju trust [options] <application name>` | `cmd/juju/application` |
| `unexpose` | Removes public availability over the network for an application. | `juju unexpose [options] <application name>` | `cmd/juju/application` |
| `unregister` | Unregisters a Juju controller. | `juju unregister [options] <controller name>` | `cmd/juju/controller` |
| `update-cloud` | Updates cloud information available to Juju. | `juju update-cloud [options] <cloud name>` | `cmd/juju/cloud` |
| `update-credential` | Updates a controller credential for a cloud. | `juju update-credential [options] [<cloud-name> [<credential-name>]]` | `cmd/juju/cloud` |
| `update-k8s` | Updates an existing Kubernetes endpoint used by Juju. | `juju update-k8s [options] <k8s name>` | `cmd/juju/caas` |
| `update-public-clouds` | Updates public cloud information available to Juju. | `juju update-public-clouds [options]` | `cmd/juju/cloud` |
| `update-secret` | Update an existing secret. | `juju update-secret [options] <ID>\|<name> [key[#base64\|#file]=value...]` | `cmd/juju/secrets` |
| `update-secret-backend` | Update an existing secret backend on the controller. | `juju update-secret-backend [options] <backend-name>` | `cmd/juju/secretbackends` |
| `update-storage-pool` | Update storage pool attributes. | `juju update-storage-pool [options] <name> [<key>=<value> [<key>=<value>...]]` | `cmd/juju/storage` |
| `upgrade-controller` | Upgrades Juju on a controller. | `juju upgrade-controller [options]` | `cmd/juju/commands` |
| `upgrade-model` | Upgrades Juju on all machines in a model. | `juju upgrade-model [options]` | `cmd/juju/commands` |
| `users` | Lists Juju users allowed to connect to a controller or model. | `juju users [options] [model-name]` | `cmd/juju/user` |
| `whoami` | Print current login details. | `juju whoami [options]` | `cmd/juju/user` |

## Dynamic commands not covered by the static registry

### PATH plugins

If the user enters an unrecognized subcommand, Juju attempts to run an executable named `juju-<subcommand>` from `PATH`. Plugins are not part of the static inventory because they are discovered dynamically at runtime.

### User aliases

The top-level supercommand also loads user aliases from the Juju XDG data directory aliases file. Those aliases can change the effective command surface on a given client, but they are local client configuration rather than built-in commands.

## Inventory observations

- The built-in surface is large and almost entirely flat.
- Command families are implemented in coherent source packages, but exposed as top-level commands rather than nested groups.
- A few help signatures reveal documentation drift or metadata issues, notably duplicated `[options]` in `find`, `download`, and `info`, plus argument strings that embed flags in `move-to-space`, `spaces`, and `subnets`.
- The resource/cross-model/secret families are structurally strong and relatively symmetric; application and controller lifecycle commands are broader but more semantically mixed.
