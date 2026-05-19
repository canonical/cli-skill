# Juju CLI Argument Structure

## Introduction and common patterns

- Most commands follow `juju <command> [options] ...`.
- Global/common options are added through the super-command (`-h/--help`, `--description`; and in many commands output options like `--format`, `-o/--output`).
- Model/controller scoping is commonly expressed through global flags and environment mapping (`-m/--model`, `-c/--controller`, `JUJU_MODEL`, `JUJU_CONTROLLER`).
- Many operational commands use positional identifiers (`<application>`, `<unit>`, `<model>`, `<cloud>`) plus repeated variadic forms (`...`).

## Complete command to argument map (usage lines)

| Command | Usage signature |
|---|---|
| `actions` | `juju actions [options] <application>` |
| `add-cloud` | `juju add-cloud [options] <cloud name> [<cloud definition file>]` |
| `add-credential` | `juju add-credential [options] <cloud name>` |
| `add-k8s` | `juju add-k8s [options] <k8s name>` |
| `add-machine` | `juju add-machine [options] [lxd[:<machine-id>] \| ssh:[<user>@]<host> \| <placement>] \| <private-key> \| <public-key>` |
| `add-model` | `juju add-model [options] <model name> [cloud\|region\|(cloud/region)]` |
| `add-secret-backend` | `juju add-secret-backend [options] <backend-name> <backend-type>` |
| `add-secret` | `juju add-secret [options] <name> [key[#base64\|#file]=value...]` |
| `add-space` | `juju add-space [options] <name> [<CIDR1> <CIDR2> ...]` |
| `add-ssh-key` | `juju add-ssh-key [options] <ssh key> ...` |
| `add-storage` | `juju add-storage [options] <unit> <storage-directive>` |
| `add-unit` | `juju add-unit [options] <application name>` |
| `add-user` | `juju add-user [options] <user name> [<display name>]` |
| `attach-resource` | `juju attach-resource [options] application <resource name>=<resource>` |
| `attach-storage` | `juju attach-storage [options] <unit> <storage> [<storage> ...]` |
| `autoload-credentials` | `juju autoload-credentials [options] [<cloud-type>]` |
| `bind` | `juju bind [options] <application> [<default-space>] [<endpoint-name>=<space> ...]` |
| `bootstrap` | `juju bootstrap [options] [<cloud name>[/region] [<controller name>]]` |
| `cancel-task` | `juju cancel-task [options] (<task-id>\|<task-id-prefix>) [...]` |
| `change-user-password` | `juju change-user-password [options] [username]` |
| `charm-resources` | `juju charm-resources [options] <charm>` |
| `clouds` | `N/A` |
| `config` | `juju config [options] <application name> [--reset <key[,key]>] [<attribute-key>][=<value>] ...]` |
| `constraints` | `juju constraints [options] <application>` |
| `consume` | `juju consume [options] <remote offer path> [<local application name>]` |
| `controller-config` | `juju controller-config [options] [<attribute key>[=<value>] ...]` |
| `controllers` | `N/A` |
| `create-backup` | `juju create-backup [options] [<notes>]` |
| `create-storage-pool` | `juju create-storage-pool [options] <name> <storage provider> [<key>=<value> [<key>=<value>...]]` |
| `credentials` | `juju credentials [options] [<cloud name>]` |
| `dashboard` | `N/A` |
| `debug-code` | `juju debug-code [options] <unit name> [hook or action names]` |
| `debug-hooks` | `juju debug-hooks [options] <unit name> [hook or action names]` |
| `debug-log` | `N/A` |
| `default-credential` | `juju default-credential [options] <cloud name> [<credential name>]` |
| `default-region` | `juju default-region [options] <cloud name> [<region>]` |
| `deploy` | `juju deploy [options] <charm or bundle> [<application name>]` |
| `destroy-controller` | `juju destroy-controller [options] <controller name>` |
| `destroy-model` | `juju destroy-model [options] [<controller name>:]<model name>` |
| `detach-storage` | `juju detach-storage [options] <storage> [<storage> ...]` |
| `diff-bundle` | `juju diff-bundle [options] <bundle file or name>` |
| `disable-command` | `juju disable-command [options] <command set> [message...]` |
| `disable-user` | `juju disable-user [options] <user name>` |
| `disabled-commands` | `N/A` |
| `documentation` | `juju documentation [options] --out <target-folder> --no-index --split --url <base-url> --discourse-ids <filepath>` |
| `download-backup` | `juju download-backup [options] /full/path/to/backup/on/controller` |
| `download` | `juju download [options] [options] <charm>` |
| `enable-command` | `juju enable-command [options] <command set>` |
| `enable-destroy-controller` | `N/A` |
| `enable-user` | `juju enable-user [options] <user name>` |
| `exec` | `juju exec [options] <commands>` |
| `export-bundle` | `N/A` |
| `expose` | `juju expose [options] <application name>` |
| `find-offers` | `N/A` |
| `find` | `juju find [options] [options] <query>` |
| `firewall-rules` | `N/A` |
| `grant-cloud` | `juju grant-cloud [options] <user name> <permission> <cloud name> ...` |
| `grant-secret` | `juju grant-secret [options] <ID>\|<name> <application>[,<application>...]` |
| `grant` | `juju grant [options] <user name> <permission> [<model name> ... \| <offer url> ...]` |
| `help-action-commands` | `juju help-action-commands [options] [action]` |
| `help-hook-commands` | `juju help-hook-commands [options] [hook]` |
| `help` | `juju help [options] [topic]` |
| `import-filesystem` | `juju import-filesystem [options]` |
| `import-ssh-key` | `juju import-ssh-key [options] <lp\|gh>:<user identity> ...` |
| `info` | `juju info [options] [options] <charm>` |
| `integrate` | `juju integrate [options] <application>[:<endpoint>] <application>[:<endpoint>]` |
| `kill-controller` | `juju kill-controller [options] <controller name>` |
| `login` | `juju login [options] [controller host name or alias]` |
| `logout` | `N/A` |
| `machines` | `N/A` |
| `migrate` | `juju migrate [options] <model-name> <target-controller-name>` |
| `model-config` | `juju model-config [options] [<model-key>[=<value>] ...]` |
| `model-constraints` | `N/A` |
| `model-defaults` | `juju model-defaults [options] [<model-key>[<=value>] ...]` |
| `model-secret-backend` | `juju model-secret-backend [options] [<secret-backend-name>]` |
| `models` | `N/A` |
| `move-to-space` | `juju move-to-space [options] [--format yaml\|json] <name> <CIDR1> [ <CIDR2> ...]` |
| `offers` | `juju offers [options] [<offer-name>]` |
| `offer` | `juju offer [options] [model-name.]<application-name>:<endpoint-name>[,...] [offer-name]` |
| `operations` | `N/A` |
| `refresh` | `juju refresh [options] <application>` |
| `regions` | `juju regions [options] <cloud>` |
| `register` | `juju register [options] <registration string>\|<controller host name>` |
| `reload-spaces` | `N/A` |
| `remove-application` | `juju remove-application [options] <application> [<application>...]` |
| `remove-cloud` | `juju remove-cloud [options] <cloud name>` |
| `remove-credential` | `juju remove-credential [options] <cloud name> <credential name>` |
| `remove-k8s` | `juju remove-k8s [options] <k8s name>` |
| `remove-machine` | `juju remove-machine [options] <machine number> ...` |
| `remove-offer` | `juju remove-offer [options] <offer-url> ...` |
| `remove-relation` | `juju remove-relation [options] <application1>[:<relation name1>] <application2>[:<relation name2>] \| <relation-id>` |
| `remove-saas` | `juju remove-saas [options] <saas-application-name> [<saas-application-name>...]` |
| `remove-secret-backend` | `juju remove-secret-backend [options] <backend-name>` |
| `remove-secret` | `juju remove-secret [options] <ID>\|<name>` |
| `remove-space` | `juju remove-space [options] <name>` |
| `remove-ssh-key` | `juju remove-ssh-key [options] <ssh key id> ...` |
| `remove-storage-pool` | `juju remove-storage-pool [options] <name>` |
| `remove-storage` | `juju remove-storage [options] <storage> [<storage> ...]` |
| `remove-unit` | `juju remove-unit [options] <unit> [...] \| <application>` |
| `remove-user` | `juju remove-user [options] <user name>` |
| `rename-space` | `juju rename-space [options] <old-name> <new-name>` |
| `resolved` | `juju resolved [options] [<unit> ...]` |
| `resources` | `juju resources [options] <application or unit>` |
| `resume-relation` | `juju resume-relation [options] <relation-id>[,<relation-id>]` |
| `retry-provisioning` | `juju retry-provisioning [options] <machine> [...]` |
| `revoke-cloud` | `juju revoke-cloud [options] <user name> <permission> <cloud name> ...` |
| `revoke-secret` | `juju revoke-secret [options] <ID>\|<name> <application>[,<application>...]` |
| `revoke` | `juju revoke [options] <user name> <permission> [<model name> ... \| <offer url> ...]` |
| `run` | `juju run [options] <unit> [<unit> ...] <action-name> [<key>=<value> [<key>[.<key> ...]=<value>]]` |
| `scale-application` | `juju scale-application [options] <application> <scale>` |
| `scp` | `juju scp [options] <source> <destination>` |
| `secret-backends` | `N/A` |
| `secrets` | `N/A` |
| `set-constraints` | `juju set-constraints [options] <application> <constraint>=<value> ...` |
| `set-credential` | `juju set-credential [options] <cloud name> <credential name>` |
| `set-firewall-rule` | `juju set-firewall-rule [options] <service-name>, --allowlist <cidr>[,<cidr>...]` |
| `set-model-constraints` | `juju set-model-constraints [options] <constraint>=<value> ...` |
| `show-action` | `juju show-action [options] <application> <action>` |
| `show-application` | `juju show-application [options] <application name or alias>` |
| `show-cloud` | `juju show-cloud [options] <cloud name>` |
| `show-controller` | `juju show-controller [options] [<controller name> ...]` |
| `show-credential` | `juju show-credential [options] [<cloud name> <credential name>]` |
| `show-machine` | `juju show-machine [options] <machineID> ...` |
| `show-model` | `juju show-model [options] <model name>` |
| `show-offer` | `juju show-offer [options] [<controller>:]<offer url>` |
| `show-operation` | `juju show-operation [options] <operation-id>` |
| `show-secret-backend` | `juju show-secret-backend [options] <backend-name>` |
| `show-secret` | `juju show-secret [options] <ID>\|<name>` |
| `show-space` | `juju show-space [options] <name>` |
| `show-status-log` | `juju show-status-log [options] <entity name>` |
| `show-storage` | `juju show-storage [options] <storage ID> [...]` |
| `show-task` | `juju show-task [options] <task ID>` |
| `show-unit` | `juju show-unit [options] <unit name>` |
| `show-user` | `juju show-user [options] [<user name>]` |
| `spaces` | `juju spaces [options] [--short] [--format yaml\|json] [--output <path>]` |
| `ssh-keys` | `N/A` |
| `ssh` | `juju ssh [options] <[user@]target> [openssh options] [command]` |
| `status` | `juju status [options] [<selector> [...]]` |
| `storage-pools` | `N/A` |
| `storage` | `juju storage [options] <filesystem\|volume> ...` |
| `subnets` | `juju subnets [options] [--space <name>] [--zone <name>] [--format yaml\|json] [--output <path>]` |
| `suspend-relation` | `juju suspend-relation [options] <relation-id>[ <relation-id>...]` |
| `switch` | `juju switch [options] [<controller>\|<model>\|<controller>:\|:<model>\|<controller>:<model>]` |
| `sync-agent-binary` | `N/A` |
| `trust` | `juju trust [options] <application name>` |
| `unexpose` | `juju unexpose [options] <application name>` |
| `unregister` | `juju unregister [options] <controller name>` |
| `update-cloud` | `juju update-cloud [options] <cloud name>` |
| `update-credential` | `juju update-credential [options] [<cloud-name> [<credential-name>]]` |
| `update-k8s` | `juju update-k8s [options] <k8s name>` |
| `update-public-clouds` | `N/A` |
| `update-secret-backend` | `juju update-secret-backend [options] <backend-name>` |
| `update-secret` | `juju update-secret [options] <ID>\|<name> [key[#base64\|#file]=value...]` |
| `update-storage-pool` | `juju update-storage-pool [options] <name> [<key>=<value> [<key>=<value>...]]` |
| `upgrade-controller` | `N/A` |
| `upgrade-model` | `N/A` |
| `users` | `juju users [options] [model-name]` |
| `version` | `N/A` |
| `whoami` | `N/A` |

## Argument metadata model

- Required vs optional:
  Required positionals are represented as `<name>`, optional values as `[<name>]`, optional groups as `[ ... ]`.
- Repetition:
  Repeated arguments are represented with `...` and comma-separated forms in some commands.
- Types:
  Common types include identifiers (application/unit/model), names, URLs/paths, CIDRs, key/value settings, and action payloads.
- Aliases:
  Command aliases are supported by the framework and surfaced in docs/help where defined.
- Env-var mapping:
  Controller/model selection supports environment variables (`JUJU_MODEL`, `JUJU_CONTROLLER`), including conflict detection.

## Special arguments

- Embedded selector syntax:
  Selectors combine namespace and identity, e.g. `[<controller>:]<model>` and endpoint forms like `<app>:<endpoint>`.
- Key-value payload syntax:
  Several commands accept dynamic maps such as `<key>=<value>`, plus specialized forms like `key[#base64|#file]=value`.
- Variadic hybrid positional modes:
  Commands such as remove/show/run families allow lists, prefixes, or polymorphic target selectors in one signature.
- Plugin argument passthrough:
  Unknown subcommands are attempted as `juju-<name>` plugins; model/controller flags are extracted and propagated via environment.

## Source notes for flag defaults and per-command option metadata

- Per-command option tables: `docs/reference/juju-cli/list-of-juju-cli-commands/*.md`.
- Shared output option implementation: `cmd/cmd/output.go`.
- Global dispatch/help/flag behavior: `cmd/cmd/supercommand.go` and `cmd/cmd/cmd.go`.
- Model/controller env and precedence logic: `cmd/modelcmd/controller.go`, `cmd/modelcmd/modelcommand.go`.
