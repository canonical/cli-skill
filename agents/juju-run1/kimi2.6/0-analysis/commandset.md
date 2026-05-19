# Juju CLI Command Set

## Source

- Generated from `docs/reference/juju-cli/list-of-juju-cli-commands/*.md`.
- Supplemented by command registration in `cmd/juju/commands/main.go`.

## Hierarchy

- Primary hierarchy is flat top-level verbs/noun-phrases under `juju <command>`.
- Additional help surfaces exist via `juju help`, `juju help-hook-commands`, and `juju help-action-commands`.
- Charm hook/action helper commands are documented separately and are not regular top-level Juju management commands.

## Full command inventory

| Command | Summary | Usage (from docs) |
|---|---|---|
| `actions` | List actions defined for an application. | `juju actions [options] <application>` |
| `add-cloud` | Add a cloud definition to Juju. | `juju add-cloud [options] <cloud name> [<cloud definition file>]` |
| `add-credential` | Adds a credential for a cloud to a local client and uploads it to a controller. | `juju add-credential [options] <cloud name>` |
| `add-k8s` | Adds a Kubernetes endpoint and credential to Juju. | `juju add-k8s [options] <k8s name>` |
| `add-machine` | Provision a new machine or assign one to the model. | `juju add-machine [options] [lxd[:<machine-id>] \| ssh:[<user>@]<host> \| <placement>] \| <private-key> \| <public-key>` |
| `add-model` | Adds a workload model. | `juju add-model [options] <model name> [cloud\|region\|(cloud/region)]` |
| `add-secret-backend` | Add a new secret backend to the controller. | `juju add-secret-backend [options] <backend-name> <backend-type>` |
| `add-secret` | Add a new secret. | `juju add-secret [options] <name> [key[#base64\|#file]=value...]` |
| `add-space` | Add a new network space. | `juju add-space [options] <name> [<CIDR1> <CIDR2> ...]` |
| `add-ssh-key` | Adds a public SSH key to a model. | `juju add-ssh-key [options] <ssh key> ...` |
| `add-storage` | Adds storage to a unit after it has been deployed. | `juju add-storage [options] <unit> <storage-directive>` |
| `add-unit` | Adds one or more units to a deployed application. | `juju add-unit [options] <application name>` |
| `add-user` | Adds a Juju user to a controller. | `juju add-user [options] <user name> [<display name>]` |
| `attach-resource` | Update a resource for an application. | `juju attach-resource [options] application <resource name>=<resource>` |
| `attach-storage` | Attaches existing storage to a unit. | `juju attach-storage [options] <unit> <storage> [<storage> ...]` |
| `autoload-credentials` | Attempts to automatically detect and add credentials for a cloud. | `juju autoload-credentials [options] [<cloud-type>]` |
| `bind` | Change bindings for a deployed application. | `juju bind [options] <application> [<default-space>] [<endpoint-name>=<space> ...]` |
| `bootstrap` | Initializes a cloud environment. | `juju bootstrap [options] [<cloud name>[/region] [<controller name>]]` |
| `cancel-task` | Cancel pending or running tasks. | `juju cancel-task [options] (<task-id>\|<task-id-prefix>) [...]` |
| `change-user-password` | Changes the password for the current or specified Juju user. | `juju change-user-password [options] [username]` |
| `charm-resources` | Display the resources for a charm in a repository. | `juju charm-resources [options] <charm>` |
| `clouds` | Lists all clouds available to Juju. | `N/A` |
| `config` | Get, set, or reset configuration for a deployed application. | `juju config [options] <application name> [--reset <key[,key]>] [<attribute-key>][=<value>] ...]` |
| `constraints` | Displays machine constraints for an application. | `juju constraints [options] <application>` |
| `consume` | Add a remote offer to the model. | `juju consume [options] <remote offer path> [<local application name>]` |
| `controller-config` | Displays or sets configuration settings for a controller. | `juju controller-config [options] [<attribute key>[=<value>] ...]` |
| `controllers` | Lists all controllers. | `N/A` |
| `create-backup` | Create a backup. | `juju create-backup [options] [<notes>]` |
| `create-storage-pool` | Create or define a storage pool. | `juju create-storage-pool [options] <name> <storage provider> [<key>=<value> [<key>=<value>...]]` |
| `credentials` | Lists Juju credentials for a cloud. | `juju credentials [options] [<cloud name>]` |
| `dashboard` | Print the Juju Dashboard URL, or open the Juju Dashboard in the default browser. | `N/A` |
| `debug-code` | Launch a tmux session to debug hooks and/or actions. | `juju debug-code [options] <unit name> [hook or action names]` |
| `debug-hooks` | Launch a tmux session to debug hooks and/or actions. | `juju debug-hooks [options] <unit name> [hook or action names]` |
| `debug-log` | Displays log messages for a model. | `N/A` |
| `default-credential` | Gets, sets, or unsets the default credential for a cloud on this client. | `juju default-credential [options] <cloud name> [<credential name>]` |
| `default-region` | Gets, sets, or unsets the default region for a cloud on this client. | `juju default-region [options] <cloud name> [<region>]` |
| `deploy` | Deploys a new application or bundle. | `juju deploy [options] <charm or bundle> [<application name>]` |
| `destroy-controller` | Destroys a controller. | `juju destroy-controller [options] <controller name>` |
| `destroy-model` | Terminate all machines/containers and resources for a non-controller model. | `juju destroy-model [options] [<controller name>:]<model name>` |
| `detach-storage` | Detaches storage from units. | `juju detach-storage [options] <storage> [<storage> ...]` |
| `diff-bundle` | Compares a bundle with a model and reports any differences. | `juju diff-bundle [options] <bundle file or name>` |
| `disable-command` | Disables commands for the model. | `juju disable-command [options] <command set> [message...]` |
| `disable-user` | Disables a Juju user. | `juju disable-user [options] <user name>` |
| `disabled-commands` | Lists disabled commands. | `N/A` |
| `documentation` | Generate the documentation for all commands | `juju documentation [options] --out <target-folder> --no-index --split --url <base-url> --discourse-ids <filepath>` |
| `download-backup` | Download a backup archive file. | `juju download-backup [options] /full/path/to/backup/on/controller` |
| `download` | Locates and then downloads a Charmhub charm. | `juju download [options] [options] <charm>` |
| `enable-command` | Enable commands that had been previously disabled. | `juju enable-command [options] <command set>` |
| `enable-destroy-controller` | Enable destroy-controller by removing disabled commands in the controller. | `N/A` |
| `enable-user` | Re-enables a previously disabled Juju user. | `juju enable-user [options] <user name>` |
| `exec` | Run the commands on the remote targets specified. | `juju exec [options] <commands>` |
| `export-bundle` | Exports the current model configuration as a reusable bundle. | `N/A` |
| `expose` | Makes an application publicly available over the network. | `juju expose [options] <application name>` |
| `find-offers` | Find offered application endpoints. | `N/A` |
| `find` | Queries the Charmhub store for available charms or bundles. | `juju find [options] [options] <query>` |
| `firewall-rules` | Prints the firewall rules. | `N/A` |
| `grant-cloud` | Grants access level to a Juju user for a cloud. | `juju grant-cloud [options] <user name> <permission> <cloud name> ...` |
| `grant-secret` | Grant access to a secret. | `juju grant-secret [options] <ID>\|<name> <application>[,<application>...]` |
| `grant` | Grants access level to a Juju user for a model, controller, or application offer. | `juju grant [options] <user name> <permission> [<model name> ... \| <offer url> ...]` |
| `help-action-commands` | Show help on a Juju charm action command. | `juju help-action-commands [options] [action]` |
| `help-hook-commands` | Show help on a Juju charm hook command. | `juju help-hook-commands [options] [hook]` |
| `help` | Show help on a command or other topic. | `juju help [options] [topic]` |
| `import-filesystem` | Imports a filesystem into the model. | `juju import-filesystem [options]` |
| `import-ssh-key` | Adds a public SSH key from a trusted identity source to a model. | `juju import-ssh-key [options] <lp\|gh>:<user identity> ...` |
| `info` | Displays detailed information about CharmHub charms. | `juju info [options] [options] <charm>` |
| `integrate` | Integrate two applications. | `juju integrate [options] <application>[:<endpoint>] <application>[:<endpoint>]` |
| `kill-controller` | Forcibly terminate all machines and other associated resources for a Juju controller. | `juju kill-controller [options] <controller name>` |
| `login` | Logs a user in to a controller. | `juju login [options] [controller host name or alias]` |
| `logout` | Logs a Juju user out of a controller. | `N/A` |
| `machines` | Lists machines in a model. | `N/A` |
| `migrate` | Migrate a workload model to another controller. | `juju migrate [options] <model-name> <target-controller-name>` |
| `model-config` | Displays or sets configuration values on a model. | `juju model-config [options] [<model-key>[=<value>] ...]` |
| `model-constraints` | Displays machine constraints for a model. | `N/A` |
| `model-defaults` | Displays or sets default configuration settings for new models. | `juju model-defaults [options] [<model-key>[<=value>] ...]` |
| `model-secret-backend` | Displays or sets the secret backend for a model. | `juju model-secret-backend [options] [<secret-backend-name>]` |
| `models` | Lists models a user can access on a controller. | `N/A` |
| `move-to-space` | Update a network space's CIDR. | `juju move-to-space [options] [--format yaml\|json] <name> <CIDR1> [ <CIDR2> ...]` |
| `offers` | Lists shared endpoints. | `juju offers [options] [<offer-name>]` |
| `offer` | Offer application endpoints for use in other models. | `juju offer [options] [model-name.]<application-name>:<endpoint-name>[,...] [offer-name]` |
| `operations` | Lists pending, running, or completed operations for specified application, units, machines, or all. | `N/A` |
| `refresh` | Refresh an application's charm. | `juju refresh [options] <application>` |
| `regions` | Lists regions for a given cloud. | `juju regions [options] <cloud>` |
| `register` | Registers a controller. | `juju register [options] <registration string>\|<controller host name>` |
| `reload-spaces` | Reloads spaces and subnets from substrate. | `N/A` |
| `remove-application` | Remove applications from the model. | `juju remove-application [options] <application> [<application>...]` |
| `remove-cloud` | Removes a cloud from Juju. | `juju remove-cloud [options] <cloud name>` |
| `remove-credential` | Removes Juju credentials for a cloud. | `juju remove-credential [options] <cloud name> <credential name>` |
| `remove-k8s` | Removes a k8s cloud from Juju. | `juju remove-k8s [options] <k8s name>` |
| `remove-machine` | Removes one or more machines from a model. | `juju remove-machine [options] <machine number> ...` |
| `remove-offer` | Removes one or more offers specified by their URL. | `juju remove-offer [options] <offer-url> ...` |
| `remove-relation` | Removes an existing relation between two applications. | `juju remove-relation [options] <application1>[:<relation name1>] <application2>[:<relation name2>] \| <relation-id>` |
| `remove-saas` | Remove consumed applications (SAAS) from the model. | `juju remove-saas [options] <saas-application-name> [<saas-application-name>...]` |
| `remove-secret-backend` | Removes a secret backend from the controller. | `juju remove-secret-backend [options] <backend-name>` |
| `remove-secret` | Remove a existing secret. | `juju remove-secret [options] <ID>\|<name>` |
| `remove-space` | Remove a network space. | `juju remove-space [options] <name>` |
| `remove-ssh-key` | Removes a public SSH key (or keys) from a model. | `juju remove-ssh-key [options] <ssh key id> ...` |
| `remove-storage-pool` | Remove an existing storage pool. | `juju remove-storage-pool [options] <name>` |
| `remove-storage` | Removes storage from the model. | `juju remove-storage [options] <storage> [<storage> ...]` |
| `remove-unit` | Remove application units from the model. | `juju remove-unit [options] <unit> [...] \| <application>` |
| `remove-user` | Deletes a Juju user from a controller. | `juju remove-user [options] <user name>` |
| `rename-space` | Rename a network space. | `juju rename-space [options] <old-name> <new-name>` |
| `resolved` | Marks unit errors resolved and re-executes failed hooks. | `juju resolved [options] [<unit> ...]` |
| `resources` | Show the resources for an application or unit. | `juju resources [options] <application or unit>` |
| `resume-relation` | Resumes a suspended relation to an application offer. | `juju resume-relation [options] <relation-id>[,<relation-id>]` |
| `retry-provisioning` | Retries provisioning for failed machines. | `juju retry-provisioning [options] <machine> [...]` |
| `revoke-cloud` | Revokes access from a Juju user for a cloud. | `juju revoke-cloud [options] <user name> <permission> <cloud name> ...` |
| `revoke-secret` | Revoke access to a secret. | `juju revoke-secret [options] <ID>\|<name> <application>[,<application>...]` |
| `revoke` | Revokes access from a Juju user for a model, controller, or application offer. | `juju revoke [options] <user name> <permission> [<model name> ... \| <offer url> ...]` |
| `run` | Run an action on a specified unit. | `juju run [options] <unit> [<unit> ...] <action-name> [<key>=<value> [<key>[.<key> ...]=<value>]]` |
| `scale-application` | Set the desired number of k8s application units. | `juju scale-application [options] <application> <scale>` |
| `scp` | Securely transfer files within a model. | `juju scp [options] <source> <destination>` |
| `secret-backends` | Lists secret backends available in the controller. | `N/A` |
| `secrets` | Lists secrets available in the model. | `N/A` |
| `set-constraints` | Sets machine constraints for an application. | `juju set-constraints [options] <application> <constraint>=<value> ...` |
| `set-credential` | Relates a remote credential to a model. | `juju set-credential [options] <cloud name> <credential name>` |
| `set-firewall-rule` | Sets a firewall rule. | `juju set-firewall-rule [options] <service-name>, --allowlist <cidr>[,<cidr>...]` |
| `set-model-constraints` | Sets machine constraints on a model. | `juju set-model-constraints [options] <constraint>=<value> ...` |
| `show-action` | Shows detailed information about an action. | `juju show-action [options] <application> <action>` |
| `show-application` | Displays information about an application. | `juju show-application [options] <application name or alias>` |
| `show-cloud` | Shows detailed information for a cloud. | `juju show-cloud [options] <cloud name>` |
| `show-controller` | Shows detailed information of a controller. | `juju show-controller [options] [<controller name> ...]` |
| `show-credential` | Shows credential information stored either on this client or on a controller. | `juju show-credential [options] [<cloud name> <credential name>]` |
| `show-machine` | Show a machine's status. | `juju show-machine [options] <machineID> ...` |
| `show-model` | Shows information about the current or specified model. | `juju show-model [options] <model name>` |
| `show-offer` | Shows extended information about the offered application. | `juju show-offer [options] [<controller>:]<offer url>` |
| `show-operation` | Show results of an operation. | `juju show-operation [options] <operation-id>` |
| `show-secret-backend` | Displays the specified secret backend. | `juju show-secret-backend [options] <backend-name>` |
| `show-secret` | Shows details for a specific secret. | `juju show-secret [options] <ID>\|<name>` |
| `show-space` | Shows information about the network space. | `juju show-space [options] <name>` |
| `show-status-log` | Output past statuses for the specified entity. | `juju show-status-log [options] <entity name>` |
| `show-storage` | Shows storage instance information. | `juju show-storage [options] <storage ID> [...]` |
| `show-task` | Show results of a task by ID. | `juju show-task [options] <task ID>` |
| `show-unit` | Displays information about a unit. | `juju show-unit [options] <unit name>` |
| `show-user` | Show information about a user. | `juju show-user [options] [<user name>]` |
| `spaces` | List known spaces, including associated subnets. | `juju spaces [options] [--short] [--format yaml\|json] [--output <path>]` |
| `ssh-keys` | Lists the currently known SSH keys for the current (or specified) model. | `N/A` |
| `ssh` | Initiates an SSH session or executes a command on a Juju machine or container. | `juju ssh [options] <[user@]target> [openssh options] [command]` |
| `status` | Report the status of the model, its machines, applications and units. | `juju status [options] [<selector> [...]]` |
| `storage-pools` | List storage pools. | `N/A` |
| `storage` | Lists storage details. | `juju storage [options] <filesystem\|volume> ...` |
| `subnets` | List subnets known to Juju. | `juju subnets [options] [--space <name>] [--zone <name>] [--format yaml\|json] [--output <path>]` |
| `suspend-relation` | Suspends a relation to an application offer. | `juju suspend-relation [options] <relation-id>[ <relation-id>...]` |
| `switch` | Selects or identifies the current controller and model. | `juju switch [options] [<controller>\|<model>\|<controller>:\|:<model>\|<controller>:<model>]` |
| `sync-agent-binary` | Copy agent binaries from the official agent store into a local controller. | `N/A` |
| `trust` | Sets the trust status of a deployed application to true. | `juju trust [options] <application name>` |
| `unexpose` | Removes public availability over the network for an application. | `juju unexpose [options] <application name>` |
| `unregister` | Unregisters a Juju controller. | `juju unregister [options] <controller name>` |
| `update-cloud` | Updates cloud information available to Juju. | `juju update-cloud [options] <cloud name>` |
| `update-credential` | Updates a controller credential for a cloud. | `juju update-credential [options] [<cloud-name> [<credential-name>]]` |
| `update-k8s` | Updates an existing Kubernetes endpoint used by Juju. | `juju update-k8s [options] <k8s name>` |
| `update-public-clouds` | Updates public cloud information available to Juju. | `N/A` |
| `update-secret-backend` | Update an existing secret backend on the controller. | `juju update-secret-backend [options] <backend-name>` |
| `update-secret` | Update an existing secret. | `juju update-secret [options] <ID>\|<name> [key[#base64\|#file]=value...]` |
| `update-storage-pool` | Update storage pool attributes. | `juju update-storage-pool [options] <name> [<key>=<value> [<key>=<value>...]]` |
| `upgrade-controller` | Upgrades Juju on a controller. | `N/A` |
| `upgrade-model` | Upgrades Juju on all machines in a model. | `N/A` |
| `users` | Lists Juju users allowed to connect to a controller or model. | `juju users [options] [model-name]` |
| `version` | Print the Juju CLI client version. | `N/A` |
| `whoami` | Print current login details. | `N/A` |
