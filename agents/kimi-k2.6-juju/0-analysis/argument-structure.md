# Argument Structure
## Common Patterns
Almost every command shares a standard global option block and a model-targeting pattern:
- **Global options**: `--debug`, `--help` (`-h`), `--logging-config`, `--quiet`, `--show-log`, `--verbose`
- **Model context**: `-m, --model` (accepts `[<controller>:]<model>`) is present on most workload commands.
- **Output control**: `--format` (json|yaml|tabular|default|smart|oneline|summary) and `-o, --output` appear on read/reporting commands.
- **Browser auth**: `-B, --no-browser-login` appears on most commands that may need controller authentication.
- **Boolean force/dry-run**: `--force`, `--dry-run`, `--no-prompt` are common on destructive or mutating commands.
- **File input**: `--file`, `--config` accept YAML/JSON file paths or stdin (`-`).
- **Positional patterns**: commands typically take 0–2 positional args: a target entity name and optionally a value or file.

## Command Argument Map
| Command | Required Positional | Optional Positional | Global Flags | Command Flags |
|---|---|---|---|---|
| actions | <application> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output, schema |
| add-cloud | <cloud name> [<cloud definition file>] | <cloud name> [<cloud definition file>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, credential, file, force, target-controller |
| add-credential | <cloud name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, file, region |
| add-k8s | <k8s name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, cloud, cluster-name, context-name, credential, region... |
| add-machine | [lxd[:<machine-id>] | ssh:[<user>@]<host> | <placement>] | <private-key> | <public-key> | [lxd[:<machine-id>] | ssh:[<user>@]<host> | <placement>] | <private-key> | <public-key> | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, base, constraints, disks, model, private-key, public-key |
| add-model | <model name> [cloud|region|(cloud/region)] | <model name> [cloud|region|(cloud/region)] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, config, credential, no-switch, owner, target-controller |
| add-secret | <name> [key[#base64|#file]=value...] | <name> [key[#base64|#file]=value...] | debug, help, logging-config, quiet, show-log, verbose | file, info, model |
| add-secret-backend | <backend-name> <backend-type> |  | debug, help, logging-config, quiet, show-log, verbose | controller, config, import-id |
| add-space | <name> [<CIDR1> <CIDR2> ...] | <name> [<CIDR1> <CIDR2> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| add-ssh-key | <ssh key> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| add-storage | <unit> <storage-directive> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| add-unit | <application name> |  | debug, help, logging-config, quiet, show-log, verbose | attach-storage, model, num-units, to |
| add-user | <user name> [<display name>] | <user name> [<display name>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| attach-resource | application <resource name>=<resource> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| attach-storage | <unit> <storage> [<storage> ...] | <unit> <storage> [<storage> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| autoload-credentials | [<cloud-type>] | [<cloud-type>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client |
| bind | <application> [<default-space>] [<endpoint-name>=<space> ...] | <application> [<default-space>] [<endpoint-name>=<space> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, force, model |
| bootstrap | [<cloud name>[/region] [<controller name>]] | [<cloud name>[/region] [<controller name>]] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, agent-version, auto-upgrade, bootstrap-base, bootstrap-constraints, bootstrap-image, build-agent, clouds... |
| cancel-task | (<task-id>|<task-id-prefix>) [...] | (<task-id>|<task-id-prefix>) [...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output |
| change-user-password |  | [username] | debug, help, logging-config, quiet, show-log, verbose | controller, no-prompt, reset |
| charm-resources | <charm> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, channel, format, model, output |
| clouds |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, all, controller, client, format, output |
| config | <application name> [--reset <key[,key]>] [<attribute-key>][=<value>] ...] | <application name> [--reset <key[,key]>] [<attribute-key>][=<value>] ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, color, file, format, model, no-color, output, reset |
| constraints | <application> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output |
| consume | <remote offer path> [<local application name>] | <remote offer path> [<local application name>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| controller-config | [<attribute key>[=<value>] ...] | [<attribute key>[=<value>] ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, color, file, format, ignore-read-only-fields, no-color, output |
| controllers |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, managed, output, refresh |
| create-backup | [<notes>] | [<notes>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, filename, model, no-download |
| create-storage-pool | <name> <storage provider> [<key>=<value> [<key>=<value>...]] | <name> <storage provider> [<key>=<value> [<key>=<value>...]] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| credentials | [<cloud name>] | [<cloud name>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, format, output, show-secrets |
| dashboard |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, browser, hide-credential, model, port |
| debug-code | <unit name> [hook or action names] | <unit name> [hook or action names] | debug, help, logging-config, quiet, show-log, verbose | at, container, model, no-host-key-checks, proxy, pty |
| debug-hooks | <unit name> [hook or action names] | <unit name> [hook or action names] | debug, help, logging-config, quiet, show-log, verbose | container, model, no-host-key-checks, proxy, pty |
| debug-log |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, color, date, exclude-labels, exclude-module, firehose, format, include... |
| default-credential | <cloud name> [<credential name>] | <cloud name> [<credential name>] | debug, help, logging-config, quiet, show-log, verbose | reset |
| default-region | <cloud name> [<region>] | <cloud name> [<region>] | debug, help, logging-config, quiet, show-log, verbose | reset |
| deploy | <charm or bundle> [<application name>] | <charm or bundle> [<application name>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, attach-storage, base, bind, channel, config, constraints, device... |
| destroy-controller | <controller name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, destroy-all-models, destroy-storage, force, model-timeout, no-prompt, no-wait, release-storage |
| destroy-model | [<controller name>:]<model name> | [<controller name>:]<model name> | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, destroy-storage, force, no-prompt, no-wait, release-storage, timeout |
| detach-storage | <storage> [<storage> ...] | <storage> [<storage> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, force, model |
| diff-bundle | <bundle file or name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, annotations, arch, base, channel, model, map-machines, overlay |
| disable-command | <command set> [message...] | <command set> [message...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| disable-user | <user name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| disabled-commands |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, all, format, model, output |
| documentation | --out <target-folder> --no-index --split --url <base-url> --discourse-ids <filepath> |  |  |  |
| download | [options] <charm> | [options] <charm> | debug, help, logging-config, quiet, show-log, verbose | arch, base, channel, charmhub-url, filepath, no-progress, resources, revision |
| download-backup |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, filename, model |
| enable-command | <command set> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| enable-destroy-controller |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| enable-user | <user name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| exec | <commands> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, all, background, color, execution-group, format, model, machine... |
| export-bundle |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, filename, include-charm-defaults, model |
| expose | <application name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, endpoints, model, to-cidrs, to-spaces |
| find | [options] <query> | [options] <query> | debug, help, logging-config, quiet, show-log, verbose | category, channel, charmhub-url, columns, format, output, publisher, type |
| find-offers |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, interface, model, output, offer, url |
| firewall-rules |  |  | debug, help, logging-config, quiet, show-log, verbose | format, model, output |
| grant | <user name> <permission> [<model name> ... | <offer url> ...] | <user name> <permission> [<model name> ... | <offer url> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| grant-cloud | <user name> <permission> <cloud name> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| grant-secret | <ID>|<name> <application>[,<application>...] | <ID>|<name> <application>[,<application>...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| help |  |  |  |  |
| help-action-commands |  |  | debug, help, logging-config, quiet, show-log, verbose |  |
| help-hook-commands |  |  | debug, help, logging-config, quiet, show-log, verbose |  |
| import-filesystem |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, force, model |
| import-ssh-key | <lp|gh>:<user identity> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| info | [options] <charm> | [options] <charm> | debug, help, logging-config, quiet, show-log, verbose | arch, base, channel, charmhub-url, config, format, output, revision... |
| integrate | <application>[:<endpoint>] <application>[:<endpoint>] | <application>[:<endpoint>] <application>[:<endpoint>] | debug, help, logging-config, quiet, show-log, verbose | model, via |
| kill-controller | <controller name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, no-prompt, timeout |
| login |  | [controller host name or alias] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, no-prompt, trust, user |
| logout |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, force |
| machines |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, color, format, model, output, utc |
| migrate | <model-name> <target-controller-name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, dry-run |
| model-config | [<model-key>[=<value>] ...] | [<model-key>[=<value>] ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, color, file, format, ignore-read-only-fields, model, no-color, output... |
| model-constraints |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output |
| model-defaults | [<model-key>[<=value>] ...] | [<model-key>[<=value>] ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, cloud, color, file, format, ignore-read-only-fields, no-color... |
| model-secret-backend | [<secret-backend-name>] | [<secret-backend-name>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| models |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, all, controller, exact-time, format, output, user, uuid |
| move-to-space | [--format yaml|json] <name> <CIDR1> [ <CIDR2> ...] | [--format yaml|json] <name> <CIDR1> [ <CIDR2> ...] | debug, help, logging-config, quiet, show-log, verbose | force, format, model, output |
| offer | [model-name.]<application-name>:<endpoint-name>[,...] [offer-name] | [model-name.]<application-name>:<endpoint-name>[,...] [offer-name] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| offers | [<offer-name>] | [<offer-name>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, active-only, allowed-consumer, application, connected-user, format, interface, model... |
| operations |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, actions, format, limit, model, machines, output, offset... |
| refresh | <application> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, base, bind, channel, config, force, force-base, force-units... |
| regions | <cloud> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, format, output |
| register | <registration string>|<controller host name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, replace |
| reload-spaces |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| remove-application | <application> [<application>...] | <application> [<application>...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, destroy-storage, dry-run, force, model, no-prompt, no-wait |
| remove-cloud | <cloud name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, target-controller |
| remove-credential | <cloud name> <credential name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, force |
| remove-k8s | <k8s name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client |
| remove-machine | <machine number> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, dry-run, force, keep-instance, model, no-prompt, no-wait |
| remove-offer | <offer-url> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, force, yes |
| remove-relation | <application1>[:<relation name1>] <application2>[:<relation name2>] | <relation-id> | <application1>[:<relation name1>] <application2>[:<relation name2>] | <relation-id> | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, force, model |
| remove-saas | <saas-application-name> [<saas-application-name>...] | <saas-application-name> [<saas-application-name>...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, force, model, no-wait |
| remove-secret | <ID>|<name> |  | debug, help, logging-config, quiet, show-log, verbose | model, revision |
| remove-secret-backend | <backend-name> |  | debug, help, logging-config, quiet, show-log, verbose | controller, force |
| remove-space | <name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, force, model, yes |
| remove-ssh-key | <ssh key id> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| remove-storage | <storage> [<storage> ...] | <storage> [<storage> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, force, model, no-destroy |
| remove-storage-pool | <name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| remove-unit | <unit> [...] | <application> | <unit> [...] | <application> | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, destroy-storage, dry-run, force, model, no-prompt, no-wait, num-units |
| remove-user | <user name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, yes |
| rename-space | <old-name> <new-name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model, rename |
| resolved | [<unit> ...] | [<unit> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, all, model, no-retry |
| resources | <application or unit> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, details, format, model, output |
| resume-relation | <relation-id>[,<relation-id>] | <relation-id>[,<relation-id>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| retry-provisioning | <machine> [...] | <machine> [...] | debug, help, logging-config, quiet, show-log, verbose | all, model |
| revoke | <user name> <permission> [<model name> ... | <offer url> ...] | <user name> <permission> [<model name> ... | <offer url> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| revoke-cloud | <user name> <permission> <cloud name> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller |
| revoke-secret | <ID>|<name> <application>[,<application>...] | <ID>|<name> <application>[,<application>...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| run | <unit> [<unit> ...] <action-name> [<key>=<value> [<key>[.<key> ...]=<value>]] | <unit> [<unit> ...] <action-name> [<key>=<value> [<key>[.<key> ...]=<value>]] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, background, color, format, model, no-color, output, params... |
| scale-application | <application> <scale> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| scp | <source> <destination> |  | debug, help, logging-config, quiet, show-log, verbose | container, model, no-host-key-checks, proxy |
| secret-backends |  |  | debug, help, logging-config, quiet, show-log, verbose | controller, format, output, reveal |
| secrets |  |  | debug, help, logging-config, quiet, show-log, verbose | format, model, output, owner, revisions |
| set-constraints | <application> <constraint>=<value> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| set-credential | <cloud name> <credential name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| set-firewall-rule | <service-name>, --allowlist <cidr>[,<cidr>...] | <service-name>, --allowlist <cidr>[,<cidr>...] | debug, help, logging-config, quiet, show-log, verbose | allowlist, model, whitelist |
| set-model-constraints | <constraint>=<value> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| show-action | <application> <action> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| show-application | <application name or alias> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output |
| show-cloud | <cloud name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, format, include-config, output |
| show-controller | [<controller name> ...] | [<controller name> ...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, output, show-password |
| show-credential | [<cloud name> <credential name>] | [<cloud name> <credential name>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, format, output, show-secrets |
| show-machine | <machineID> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, color, format, model, output, utc |
| show-model | <model name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, output |
| show-offer | [<controller>:]<offer url> | [<controller>:]<offer url> | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output |
| show-operation | <operation-id> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output, utc, wait, watch |
| show-secret | <ID>|<name> |  | debug, help, logging-config, quiet, show-log, verbose | format, model, output, revision, reveal, revisions |
| show-secret-backend | <backend-name> |  | debug, help, logging-config, quiet, show-log, verbose | controller, format, output, reveal |
| show-space | <name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output |
| show-status-log | <entity name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, days, format, from-date, model, output, type, utc |
| show-storage | <storage ID> [...] | <storage ID> [...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output |
| show-task | <task ID> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output, utc, wait, watch |
| show-unit | <unit name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, app, endpoint, format, model, output, related-unit |
| show-user | [<user name>] | [<user name>] | debug, help, logging-config, quiet, show-log, verbose | controller, exact-time, format, output |
| spaces | [--short] [--format yaml|json] [--output <path>] | [--short] [--format yaml|json] [--output <path>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output, short |
| ssh | <[user@]target> [openssh options] [command] | <[user@]target> [openssh options] [command] | debug, help, logging-config, quiet, show-log, verbose | container, model, no-host-key-checks, proxy, pty |
| ssh-keys |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, full, model |
| status | [<selector> [...]] | [<selector> [...]] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, color, format, integrations, model, no-color, output, relations... |
| storage | <filesystem|volume> ... |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, filesystem, format, model, output, volume |
| storage-pools |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, name, output, provider |
| subnets | [--space <name>] [--zone <name>] [--format yaml|json] [--output <path>] | [--space <name>] [--zone <name>] [--format yaml|json] [--output <path>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, model, output, space, zone |
| suspend-relation | <relation-id>[ <relation-id>...] | <relation-id>[ <relation-id>...] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model, message |
| switch | [<controller>|<model>|<controller>:|:<model>|<controller>:<model>] | [<controller>|<model>|<controller>:|:<model>|<controller>:<model>] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, model |
| sync-agent-binary |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, agent-version, dry-run, local-dir, model, public, source, stream |
| trust | <application name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model, remove, scope |
| unexpose | <application name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, endpoints, model |
| unregister | <controller name> |  | debug, help, logging-config, quiet, show-log, verbose | no-prompt |
| update-cloud | <cloud name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client |
| update-credential | [<cloud-name> [<credential-name>]] | [<cloud-name> [<credential-name>]] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client, file, force, region |
| update-k8s | <k8s name> |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client |
| update-public-clouds |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, controller, client |
| update-secret | <ID>|<name> [key[#base64|#file]=value...] | <ID>|<name> [key[#base64|#file]=value...] | debug, help, logging-config, quiet, show-log, verbose | auto-prune, file, info, model, name |
| update-secret-backend | <backend-name> |  | debug, help, logging-config, quiet, show-log, verbose | controller, config, force, reset |
| update-storage-pool | <name> [<key>=<value> [<key>=<value>...]] | <name> [<key>=<value> [<key>=<value>...]] | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, model |
| upgrade-controller |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, agent-stream, agent-version, build-agent, controller, dry-run, ignore-agent-versions, timeout... |
| upgrade-model |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, agent-stream, agent-version, dry-run, ignore-agent-versions, model, timeout, yes |
| users |  | [model-name] | debug, help, logging-config, quiet, show-log, verbose | all, controller, exact-time, format, output |
| version |  |  |  |  |
| whoami |  |  | debug, help, logging-config, quiet, show-log, verbose | no-browser-login, format, output |

## Special Arguments
- **Multi-value key=value flags**: `--config key=value`, `--model-default key=value`, `--storage-pool name=type ...` accept repeated key=value pairs and are parsed into maps.
- **Comma-separated lists**: `--reset key1,key2` (config), `--to machine1,machine2` (add-unit), `--bootstrap-constraints 'cores=2 mem=4G'` use constraint DSL.
- **Placement directives**: `--to` accepts provider-specific placement (e.g., `lxd:0`, `zone=us-east-1a`).
- **Stdin pipe**: Many file-reading commands accept `-` to read from stdin (e.g., `juju config app --file -`).
- **Model/controller scoping**: Positional arguments like `controller:model` or UUIDs are accepted by `--model`.
- **Resource attachment**: `attach-resource app file=path` uses `name=path` positional pairs.

## Aliases
The following aliases share the exact same argument structure as their canonical commands and are omitted from the main table above for brevity:

| Alias | Canonical Command |
|---|---|
| debug-hook | debug-hooks |
| list-actions | actions |
| list-charm-resources | charm-resources |
| list-clouds | clouds |
| list-controllers | controllers |
| list-credentials | credentials |
| list-disabled-commands | disabled-commands |
| list-firewall-rules | firewall-rules |
| list-machines | machines |
| list-models | models |
| list-offers | offers |
| list-operations | operations |
| list-regions | regions |
| list-resources | resources |
| list-secret-backends | secret-backends |
| list-secrets | secrets |
| list-spaces | spaces |
| list-ssh-keys | ssh-keys |
| list-storage | storage |
| list-storage-pools | storage-pools |
| list-subnets | subnets |
| list-users | users |
| model-default | model-defaults |
| relate | integrate |
| resolve | resolved |
| set-default-credentials | default-credential |
| set-default-region | default-region |
| show-credentials | show-credential |
| update-credentials | update-credential |
