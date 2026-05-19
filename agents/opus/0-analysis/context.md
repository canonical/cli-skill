# Juju CLI Analysis Context

## Architecture Overview
- Language: Go
- CLI Framework: Custom `cmd` package (`github.com/juju/juju/cmd/cmd`) built on `github.com/juju/gnuflag`
- Pattern: SuperCommand (flat top-level command registry)
- Plugin system: Missing callback delegates to external `juju-<command>` binaries
- Client-server: CLI is a thin API client over WebSocket to Juju controller
- Command registration: `cmd/juju/commands/main.go` → `registerCommands()`
- Model context: Most commands embed `modelcmd.ModelCommandBase` for controller/model resolution
- Output: Custom `cmd.Output` with formatters (yaml, json, tabular)

## Complete Command Inventory

| Command | Package | Purpose | Aliases |
|---|---|---|---|
| actions | action | List actions defined for an application. | list-actions |
| add-cloud | cloud | | |
| add-credential | cloud | | |
| add-k8s | caas | | |
| add-machine | machine | Provision a new machine or assign one to the model. | |
| add-model | controller | Adds a workload model. | |
| add-secret-backend | secretbackends | Add a new secret backend to the controller. | |
| add-secret | secrets | Add a new secret. | |
| add-space | space | Add a new network space. | |
| add-ssh-key | sshkeys | | |
| add-storage | storage | Adds storage to a unit after it has been deployed. | |
| add-unit | application | | |
| add-user | user | | |
| attach-resource | resource | Update a resource for an application. | |
| attach-storage | storage | Attaches existing storage to a unit. | |
| autoload-credentials | cloud | | |
| bind | application | Change bindings for a deployed application. | |
| bootstrap | commands | | |
| cancel-task | action | Cancel pending or running tasks. | |
| change-user-password | user | Changes the password for the current or specified Juju user. | |
| clouds | cloud | Lists all clouds available to Juju. | |
| config | application | | |
| constraints | application | | |
| consume | application | | |
| controller-config | controller | Displays or sets configuration settings for a controller. | |
| controllers | controller | | list-controllers |
| create-backup | backups | Create a backup. | |
| create-storage-pool | storage | Create or define a storage pool. | |
| credentials | cloud | | |
| dashboard | dashboard | Print the Juju Dashboard URL, or open the Juju Dashboard in the default browser. | |
| debug-code | ssh | Launch a tmux session to debug hooks and/or actions. | |
| debug-hooks | ssh | Launch a tmux session to debug hooks and/or actions. | debug-hook |
| debug-log | commands | | |
| default-credential | cloud | | set-default-credentials |
| default-region | cloud | | set-default-region |
| deploy | application | Deploys a new application or bundle. | |
| destroy-controller | controller | | |
| destroy-model | model | Terminate all machines/containers and resources for a non-controller model. | |
| detach-storage | storage | Detaches storage from units. | |
| diff-bundle | application | Compares a bundle with a model and reports any differences. | |
| disable-command | block | Disables commands for the model. | |
| disable-user | user | | |
| disabled-commands | block | Lists disabled commands. | list-disabled-commands |
| download-backup | backups | Download a backup archive file. | |
| download | charmhub | | |
| dump-db | model | Displays the mongo documents for of the model. | |
| dump-model | model | Displays the database agnostic representation of the model. | |
| enable-command | block | Enable commands that had been previously disabled. | |
| enable-destroy-controller | controller | Enable destroy-controller by removing disabled commands in the controller. | |
| exec | action | Run the commands on the remote targets specified. | |
| export-bundle | model | Exports the current model configuration as a reusable bundle. | |
| expose | application | | |
| find-offers | crossmodel | Find offered application endpoints. | |
| find | charmhub | | |
| firewall-rules | firewall | | list-firewall-rules |
| grant-cloud | model | | |
| grant-secret | secrets | Grant access to a secret. | |
| grant | model | | |
| help-action-commands | commands | Show help on a Juju charm action command. | |
| help-hook-commands | commands | Show help on a Juju charm hook command. | |
| import-filesystem | storage | Imports a filesystem into the model. | |
| import-ssh-key | sshkeys | | |
| info | charmhub | | |
| integrate | application | | relate |
| juju | commands | Enter an interactive shell for running Juju commands | |
| kill-controller | controller | Forcibly terminate all machines and other associated resources for a Juju controller. | |
| login | user | Logs a user in to a controller. | |
| logout | user | Logs a Juju user out of a controller. | |
| machines | machine | | list-machines |
| migrate | commands | Migrate a workload model to another controller. | |
| model-constraints | model | Displays machine constraints for a model. | |
| model-secret-backend | secretbackends | Displays or sets the secret backend for a model. | |
| models | controller | Lists models a user can access on a controller. | |
| move-to-space | space | Update a network space's CIDR. | |
| offer | crossmodel | Offer application endpoints for use in other models. | |
| offers | crossmodel | | |
| operations | action | Lists pending, running, or completed operations for specified application, units, machines, or all. | |
| refresh | application | Refresh an application's charm. | |
| regions | cloud | Lists regions for a given cloud. | list-regions |
| register | controller | | |
| reload-spaces | space | Reloads spaces and subnets from substrate. | |
| remove-application | application | | |
| remove-cloud | cloud | | |
| remove-credential | cloud | | |
| remove-k8s | caas | | |
| remove-machine | machine | Removes one or more machines from a model. | |
| remove-offer | crossmodel | Removes one or more offers specified by their URL. | |
| remove-relation | application | | |
| remove-saas | application | | |
| remove-secret-backend | secretbackends | Removes a secret backend from the controller. | |
| remove-secret | secrets | Remove a existing secret. | |
| remove-space | space | Remove a network space. | |
| remove-ssh-key | sshkeys | | |
| remove-storage-pool | storage | Remove an existing storage pool. | |
| remove-storage | storage | Removes storage from the model. | |
| remove-unit | application | Remove application units from the model. | |
| remove-user | user | | |
| rename-space | space | Rename a network space. | |
| resolved | application | Marks unit errors resolved and re-executes failed hooks. | resolve |
| resources | resource | | list-resources |
| resume-relation | application | | |
| retry-provisioning | model | Retries provisioning for failed machines. | |
| run | action | Run an action on a specified unit. | |
| scale-application | application | Set the desired number of k8s application units. | |
| scp | ssh | | |
| secret-backends | secretbackends | Lists secret backends available in the controller. | list-secret-backends |
| secrets | secrets | Lists secrets available in the model. | list-secrets |
| set-credential | model | Relates a remote credential to a model. | |
| set-constraints | application | | |
| set-firewall-rule | firewall | | |
| set-model-constraints | model | | |
| show-action | action | Shows detailed information about an action. | |
| show-application | application | Displays information about an application. | |
| show-cloud | cloud | Shows detailed information for a cloud. | |
| show-controller | controller | | |
| show-credential | cloud | Shows credential information stored either on this client or on a controller. | |
| show-machine | machine | Show a machine's status. | |
| show-model | model | Shows information about the current or specified model. | |
| show-offer | crossmodel | Shows extended information about the offered application. | |
| show-operation | action | Show results of an operation. | |
| show-secret-backend | secretbackends | Displays the specified secret backend. | |
| show-secret | secrets | Shows details for a specific secret. | |
| show-space | space | Shows information about the network space. | |
| show-status-log | status | Output past statuses for the specified entity. | |
| show-storage | storage | Shows storage instance information. | |
| show-task | action | Show results of a task by ID. | |
| show-unit | application | Displays information about a unit. | |
| show-user | user | | |
| spaces | space | List known spaces, including associated subnets. | |
| ssh | ssh | | |
| ssh-keys | sshkeys | | |
| status | status | | |
| storage | storage | Lists storage details. | |
| storage-pools | storage | List storage pools. | |
| subnets | subnet | List subnets known to Juju. | list-subnets |
| suspend-relation | application | | |
| switch | commands | | |
| sync-agent-binary | commands | Copy agent binaries from the official agent store into a local controller. | |
| trust | application | | |
| unexpose | application | | |
| unregister | controller | Unregisters a Juju controller. | |
| update-cloud | cloud | Updates cloud information available to Juju. | |
| update-credential | cloud | | update-credentials |
| update-k8s | caas | | |
| update-public-clouds | cloud | Updates public cloud information available to Juju. | |
| update-secret-backend | secretbackends | Update an existing secret backend on the controller. | |
| update-secret | secrets | Update an existing secret. | |
| update-storage-pool | storage | Update storage pool attributes. | |
| upgrade-controller | commands | | |
| upgrade-model | commands | | |
| users | user | | |
| version | commands | Print the Juju CLI client version. | |
| whoami | user | Print current login details. | |
