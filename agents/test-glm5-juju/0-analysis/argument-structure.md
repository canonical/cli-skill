# Juju CLI Argument Structure Analysis

## Introduction

The Juju CLI uses a structured argument pattern based on the Go `gnuflag` library. This document maps all commands and their argument patterns, including flags, positional arguments, and environment variable mappings.

### Common Argument Patterns

**1. Model Selection** (`-m`, `--model`):
- Format: `[<controller name>:]<model name>|<model UUID>`
- Most commands operating on a model accept this flag
- Environment variable: `JUJU_MODEL`

**2. Controller Selection** (`-c`, `--controller`):
- Controller-targeted commands
- Environment variable: `JUJU_CONTROLLER`

**3. Output Formatting** (`--format`, `-o`, `--output`):
- Formats: `default`, `json`, `yaml`, `tabular`, `line`, `oneline`, `summary`
- `-o` / `--output`: Write output to file

**4. Authentication** (`-B`, `--no-browser-login`):
- Disable web browser authentication flow

**5. Force/Confirmation** (`--force`, `--no-prompt`):
- Override safety checks and confirmations

## Complete Argument Matrix

### Action Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `actions` | `<application>` | - | `--format`, `--model`, `--output`, `--schema` | `JUJU_MODEL` |
| `cancel-task` | `<task-id>...` | - | `--model` | `JUJU_MODEL` |
| `exec` | - | - | `--application`, `--format`, `--model`, `--num-tasks`, `--output`, `--timeout`, `--unit` | `JUJU_MODEL` |
| `operations` | - | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `run` | `<unit>` | - | `--model`, `--params`, `--string` | `JUJU_MODEL` |
| `show-action` | `<application> <action>` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `show-operation` | `<operation-id>` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `show-task` | `<task-id>` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |

### Application Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-unit` | `<application>` | - | `--attach-storage`, `--model`, `--num-units`, `--to` | `JUJU_MODEL` |
| `bind` | `<application> [<default-space>]` | - | `--model`, `<endpoint>=<space>...` | `JUJU_MODEL` |
| `config` | `<application>` | - | `--color`, `--file`, `--format`, `--model`, `--no-color`, `--output`, `--reset`, `<key>=<value>...` | `JUJU_MODEL` |
| `consume` | `<remote-offer-path> [<local-name>]` | - | `--controller`, `--model` | `JUJU_MODEL` |
| `deploy` | `<charm> [<name>]` | - | `--attach-storage`, `--base`, `--bind`, `--channel`, `--config`, `--constraints`, `--device`, `--dry-run`, `--force`, `--map-machines`, `--model`, `--num-units`, `--overlay`, `--resource`, `--revision`, `--storage`, `--to`, `--trust` | `JUJU_MODEL` |
| `diff-bundle` | `<bundle>` | - | `--annotation`, `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `expose` | `<application>` | - | `--endpoints`, `--model` | `JUJU_MODEL` |
| `integrate` | `<app>[:<endpoint>] <app>[:<endpoint>]` | - | `--model`, `--via` | `JUJU_MODEL` |
| `refresh` | `<application>` | - | `--base`, `--channel`, `--force`, `--model`, `--resource`, `--revision`, `--trust` | `JUJU_MODEL` |
| `remove-application` | `<application>...` | - | `--destroy-storage`, `--dry-run`, `--force`, `--model`, `--no-prompt`, `--no-wait` | `JUJU_MODEL` |
| `remove-relation` | `<endpoint1> <endpoint2>` | - | `--force`, `--model` | `JUJU_MODEL` |
| `remove-saas` | `<saas>...` | - | `--model` | `JUJU_MODEL` |
| `remove-unit` | `<unit>...` | - | `--destroy-storage`, `--force`, `--model`, `--no-prompt`, `--no-wait` | `JUJU_MODEL` |
| `resolved` | `<unit>` | - | `--model`, `--retry` | `JUJU_MODEL` |
| `resume-relation` | `<relation-id>...` | - | `--model` | `JUJU_MODEL` |
| `scale-application` | `<application> <scale>` | - | `--model` | `JUJU_MODEL` |
| `show-application` | `<application>` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `show-unit` | `<unit>` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `suspend-relation` | `<relation-id>...` | - | `--model` | `JUJU_MODEL` |
| `trust` | `<application>` | - | `--model` | `JUJU_MODEL` |
| `unexpose` | `<application>` | - | `--endpoints`, `--model` | `JUJU_MODEL` |

### Backup Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `create-backup` | `[<notes>]` | - | `--controller`, `--filename`, `--keep-copy`, `--no-download` | `JUJU_CONTROLLER` |
| `download-backup` | `<path>` | - | `--controller`, `--filename` | `JUJU_CONTROLLER` |

### Block Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `disable-command` | `<command-set> [message]` | - | `--controller` | `JUJU_CONTROLLER` |
| `disabled-commands` | - | - | `--controller`, `--format`, `--output` | `JUJU_CONTROLLER` |
| `enable-command` | `<command-set>` | - | `--controller` | `JUJU_CONTROLLER` |

### Cloud Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-cloud` | `<cloud-name> [<file>]` | - | `--client`, `--controller`, `--credential`, `--file`, `--force`, `--target-controller` | - |
| `add-credential` | `<cloud-name>` | - | `--client`, `--controller`, `--file`, `--region` | - |
| `add-k8s` | `<k8s-name>` | - | `--client`, `--cloud`, `--cluster-name`, `--context-name`, `--controller`, `--credential`, `--region`, `--skip-storage`, `--storage` | - |
| `autoload-credentials` | `[<cloud-type>]` | - | - | - |
| `clouds` | - | - | `--client`, `--controller`, `--format`, `--output` | - |
| `credentials` | `[<cloud>]` | - | `--client`, `--controller`, `--format`, `--output`, `--show-secrets` | - |
| `default-credential` | `<cloud> [<credential>]` | - | `--client`, `--controller` | - |
| `default-region` | `<cloud> [<region>]` | - | `--client`, `--controller` | - |
| `regions` | `<cloud>` | - | `--client`, `--controller`, `--format`, `--output` | - |
| `remove-cloud` | `<cloud>` | - | `--client`, `--controller` | - |
| `remove-credential` | `<cloud> <credential>` | - | `--client`, `--controller` | - |
| `remove-k8s` | `<k8s>` | - | `--client`, `--controller` | - |
| `show-cloud` | `<cloud>` | - | `--client`, `--controller`, `--format`, `--include-config-defaults`, `--output` | - |
| `show-credential` | `[<cloud> <credential>]` | - | `--client`, `--controller`, `--format`, `--output`, `--show-secrets` | - |
| `update-cloud` | `<cloud>` | - | `--client`, `--controller`, `--file` | - |
| `update-credential` | `[<cloud> [<credential>]]` | - | `--client`, `--controller`, `--file` | - |
| `update-k8s` | `<k8s>` | - | `--client`, `--controller`, `--context-name`, `--credential`, `--file` | - |
| `update-public-clouds` | - | - | `--client`, `--controller` | - |

### Controller Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-model` | `<model> [<cloud-region>]` | - | `--config`, `--controller`, `--credential`, `--no-switch`, `--owner`, `--target-controller` | `JUJU_CONTROLLER` |
| `controller-config` | `[<key>=<value>...]` | - | `--color`, `--controller`, `--file`, `--format`, `--ignore-read-only-fields`, `--no-color`, `--output` | `JUJU_CONTROLLER` |
| `controllers` | - | - | `--client`, `--format`, `--output` | - |
| `destroy-controller` | `<controller>` | - | `--destroy-all-models`, `--destroy-storage`, `--force`, `--model-timeout`, `--no-prompt`, `--no-wait`, `--release-storage` | - |
| `enable-destroy-controller` | - | - | `--controller` | `JUJU_CONTROLLER` |
| `kill-controller` | `<controller>` | - | `--no-prompt` | - |
| `list-models` | - | - | `--controller`, `--format`, `--output` | `JUJU_CONTROLLER` |
| `register` | `<string>|<host>` | - | `--controller-name` | - |
| `show-controller` | `[<controller>...]` | - | `--format`, `--output` | - |
| `unregister` | `<controller>` | - | - | - |

### Cross-Model Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `find-offers` | `[<pattern>]` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `offer` | `[<model>.]<app>:<endpoint>... [<name>]` | - | `--controller` | `JUJU_CONTROLLER` |
| `offers` | `[<name>]` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `remove-offer` | `<offer>` | - | `--force`, `--model` | `JUJU_MODEL` |
| `show-offer` | `[<controller>:]<offer>` | - | `--format`, `--output` | - |

### Charm Hub Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `download` | `<charm>` | - | `--channel`, `--config`, `--filepath`, `--revision` | - |
| `find` | `[<query>]` | - | `--channel`, `--format`, `--output`, `--type` | - |
| `info` | `<charm>` | - | `--channel`, `--format`, `--output` | - |

### Dashboard Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `dashboard` | - | - | `--controller`, `--hide-banner`, `--no-browser`, `--no-webserver`, `--open` | `JUJU_CONTROLLER` |

### Firewall Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `firewall-rules` | - | - | `--controller`, `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `set-firewall-rule` | `<service>` | `--allowlist` | `--model` | `JUJU_MODEL` |

### Machine Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-machine` | `[<placement>]` | - | `--base`, `--constraints`, `--disks`, `--model`, `--num-machines`, `--private-key`, `--public-key` | `JUJU_MODEL` |
| `machines` | - | - | `--color`, `--format`, `--model`, `--no-color`, `--output`, `--physical` | `JUJU_MODEL` |
| `remove-machine` | `<machine>...` | - | `--force`, `--model`, `--no-prompt`, `--no-wait` | `JUJU_MODEL` |
| `show-machine` | `<machine-id>...` | - | `--color`, `--format`, `--model`, `--no-color`, `--output` | `JUJU_MODEL` |

### Model Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `config` | `[<application>]` | - | `--color`, `--file`, `--format`, `--model`, `--no-color`, `--output`, `--reset`, `<key>=<value>...` | `JUJU_MODEL` |
| `constraints` | `[<application>]` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `defaults` | `[<key>=<value>...]` | - | `--cloud`, `--controller`, `--file`, `--format`, `--ignore-read-only-fields`, `--output`, `--region` | - |
| `destroy-model` | `[<controller>:]<model>` | - | `--destroy-storage`, `--force`, `--model-timeout`, `--no-prompt`, `--no-wait`, `--release-storage` | - |
| `export-bundle` | - | - | `--filename`, `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `grant` | `<user> <permission> [<model>...]` | - | `--controller` | - |
| `grant-cloud` | `<user> <permission> <cloud>...` | - | `--controller` | - |
| `migrate` | `<model> <controller>` | - | `--force` | - |
| `model-config` | `[<key>=<value>...]` | - | `--color`, `--file`, `--format`, `--model`, `--no-color`, `--output`, `--reset` | `JUJU_MODEL` |
| `model-constraints` | - | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `model-defaults` | `[<key>=<value>...]` | - | `--cloud`, `--controller`, `--file`, `--format`, `--ignore-read-only-fields`, `--output`, `--region` | - |
| `model-secret-backend` | `[<backend>]` | - | `--model` | `JUJU_MODEL` |
| `retry-provisioning` | `<machine>...` | - | `--model` | `JUJU_MODEL` |
| `revoke` | `<user> <permission> [<model>...]` | - | `--controller` | - |
| `revoke-cloud` | `<user> <permission> <cloud>...` | - | `--controller` | - |
| `set-constraints` | `<application> <constraint>=<value>...` | - | `--model` | `JUJU_MODEL` |
| `set-credential` | `<cloud> <credential>` | - | `--model` | `JUJU_MODEL` |
| `set-model-constraints` | `<constraint>=<value>...` | - | `--model` | `JUJU_MODEL` |
| `show-model` | `<model>` | - | `--format`, `--output` | - |

### Resource Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `attach-resource` | `<application> <name>=<value>` | - | `--model` | `JUJU_MODEL` |
| `charm-resources` | `<charm>` | - | `--channel`, `--format`, `--output` | - |
| `resources` | `<application>` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |

### Secret Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-secret` | `<name> [key=value...]` | - | `--file`, `--info`, `--model` | `JUJU_MODEL` |
| `add-secret-backend` | `<name> <type>` | - | `--config`, `--controller`, `--import-id` | - |
| `grant-secret` | `<id|name> <application>...` | - | `--model` | `JUJU_MODEL` |
| `remove-secret` | `<id|name>` | - | `--model` | `JUJU_MODEL` |
| `remove-secret-backend` | `<name>` | - | `--controller` | - |
| `revoke-secret` | `<id|name> <application>...` | - | `--model` | `JUJU_MODEL` |
| `secret-backends` | - | - | `--controller`, `--format`, `--output` | - |
| `secrets` | - | - | `--format`, `--model`, `--output`, `--show-secrets` | `JUJU_MODEL` |
| `show-secret` | `<id|name>` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `show-secret-backend` | `<name>` | - | `--controller`, `--format`, `--output` | - |
| `update-secret` | `<id|name> [key=value...]` | - | `--file`, `--info`, `--model` | `JUJU_MODEL` |
| `update-secret-backend` | `<name>` | - | `--config`, `--controller` | - |

### Space Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-space` | `<name> [<CIDR>...]` | - | `--model` | `JUJU_MODEL` |
| `move-to-space` | `<name> <CIDR>...` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `reload-spaces` | - | - | `--model` | `JUJU_MODEL` |
| `remove-space` | `<name>` | - | `--force`, `--model` | `JUJU_MODEL` |
| `rename-space` | `<old-name> <new-name>` | - | `--model` | `JUJU_MODEL` |
| `show-space` | `<name>` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `spaces` | - | - | `--format`, `--model`, `--output`, `--short` | `JUJU_MODEL` |

### SSH Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-ssh-key` | `<key>...` | - | `--model` | `JUJU_MODEL` |
| `debug-code` | `<unit> [hooks...]` | - | `--model` | `JUJU_MODEL` |
| `debug-hooks` | `<unit> [hooks...]` | - | `--model` | `JUJU_MODEL` |
| `import-ssh-key` | `<lp|gh>:<user>...` | - | `--model` | `JUJU_MODEL` |
| `remove-ssh-key` | `<key-id>...` | - | `--model` | `JUJU_MODEL` |
| `scp` | `<source> <dest>` | - | `--model`, `--no-preserve` | `JUJU_MODEL` |
| `ssh` | `<target> [command]` | - | `--container`, `--model`, `--no-pseudo-tty`, `--pty` | `JUJU_MODEL` |
| `ssh-keys` | - | - | `--full`, `--model` | `JUJU_MODEL` |

### Storage Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-storage` | `<unit> <directive>` | - | `--model` | `JUJU_MODEL` |
| `attach-storage` | `<unit> <storage>...` | - | `--model` | `JUJU_MODEL` |
| `create-storage-pool` | `<name> <type> [<key=value>...]` | - | `--model` | `JUJU_MODEL` |
| `detach-storage` | `<storage>...` | - | `--force`, `--model`, `--no-prompt` | `JUJU_MODEL` |
| `import-filesystem` | `<provider> <id> <name>` | - | `--model` | `JUJU_MODEL` |
| `remove-storage` | `<storage>...` | - | `--destroy-storage`, `--force`, `--model`, `--no-prompt` | `JUJU_MODEL` |
| `remove-storage-pool` | `<name>` | - | `--model` | `JUJU_MODEL` |
| `show-storage` | `<storage>...` | - | `--format`, `--model`, `--output` | `JUJU_MODEL` |
| `storage` | `[<id>...]` | - | `--filesystem`, `--format`, `--model`, `--output`, `--volume` | `JUJU_MODEL` |
| `storage-pools` | - | - | `--format`, `--model`, `--names`, `--output`, `--providers` | `JUJU_MODEL` |
| `update-storage-pool` | `<name> [<key=value>...]` | - | `--model` | `JUJU_MODEL` |

### Subnet Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `subnets` | - | - | `--format`, `--model`, `--output`, `--space`, `--zone` | `JUJU_MODEL` |

### Status/Debug Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `bootstrap` | `[<cloud>[/<region>] [<name>]]` | - | `--agent-version`, `--auto-upgrade`, `--bootstrap-base`, `--bootstrap-constraints`, `--bootstrap-image`, `--build-agent`, `--clouds`, `--config`, `--constraints`, `--controller-charm-channel`, `--controller-charm-path`, `--credential`, `--force`, `--keep-broken`, `--metadata-source`, `--model-default`, `--no-switch`, `--regions`, `--storage-pool`, `--to` | - |
| `debug-log` | - | - | `--color`, `--combine`, `--exclude-labels`, `--exclude-modules`, `--format`, `--include-labels`, `--include-modules`, `--level`, `--limit`, `--model`, `--no-color`, `--replay`, `--replay-from-end`, `--time`, `--type`, `--utc` | `JUJU_MODEL` |
| `show-status-log` | `<entity>` | - | `--color`, `--from`, `--model`, `--no-color`, `--output`, `--replay`, `--to`, `--type` | `JUJU_MODEL` |
| `status` | `[<selector>...]` | - | `--color`, `--format`, `--integrations`, `--model`, `--no-color`, `--output`, `--relations`, `--retry-count`, `--retry-delay`, `--storage`, `--utc` | `JUJU_MODEL` |
| `switch` | `[<target>]` | - | - | - |
| `whoami` | - | - | `--format`, `--output` | - |

### User Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `add-user` | `<user> [<display-name>]` | - | `--controller` | - |
| `change-user-password` | `[<user>]` | - | `--controller` | - |
| `disable-user` | `<user>` | - | `--controller` | - |
| `enable-user` | `<user>` | - | `--controller` | - |
| `login` | `[<host>]` | - | `--controller`, `--no-prompt`, `--password`, `--refresh`, `--skip-login`, `--user` | - |
| `logout` | - | - | `--controller` | - |
| `remove-user` | `<user>` | - | `--controller`, `--no-prompt` | - |
| `show-user` | `[<user>]` | - | `--controller`, `--format`, `--output` | - |
| `users` | `[<model>]` | - | `--controller`, `--format`, `--output` | - |

### Upgrade Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `sync-agent-binary` | - | - | `--agent-version`, `--destination`, `--dry-run`, `--local`, `--source`, `--stream` | - |
| `upgrade-controller` | - | - | `--agent-version`, `--controller`, `--switch-agent-version` | - |
| `upgrade-model` | - | - | `--agent-version`, `--model`, `--switch-agent-version` | - |

### Help/Documentation Commands

| Command | Positional Args | Required Flags | Optional Flags | Env Vars |
|---------|----------------|---------------|----------------|----------|
| `documentation` | - | `--out` | `--discourse-ids`, `--no-index`, `--split`, `--url` | - |
| `help` | `[<topic>]` | - | - | - |
| `help-action-commands` | `[<action>]` | - | - | - |
| `help-hook-commands` | `[<hook>]` | - | - | - |
| `version` | - | - | `--all`, `--format`, `--output` | - |

## Special Arguments

### Model Specification

The `<model>` argument accepts multiple formats:
- `<name>` - Model name in current controller
- `<controller>:<name>` - Model in specific controller
- `<UUID>` - Model UUID

**Examples:**
```bash
juju status -m mymodel
juju status -m mycontroller:mymodel
juju status -m 1234-abcd-...
```

### Placement Directives

The `--to` flag accepts placement specifications:
- `<machine-id>` - Deploy to specific machine
- `lxd` - New LXD container on new machine
- `lxd:<machine-id>` - LXD container on specific machine
- `<machine-id>/lxd/<container-id>` - Specific container
- `zone=<zone>` - Availability zone (provider-dependent)
- `<host>.maas` - MAAS node

**Examples:**
```bash
juju deploy mysql --to 23
juju deploy mysql --to lxd:5
juju add-unit mysql -n 2 --to 3,lxd:5
```

### Constraints Syntax

Constraints are specified as comma-separated key=value pairs:
```
arch=<arch>
mem=<size>
cores=<n>
storage=<size>
virt-type=<type>
spaces=<space1>,<space2>
tags=<tag1>,<tag2>
```

**Examples:**
```bash
juju deploy postgresql --constraints mem=8G
juju deploy haproxy --constraints spaces=dmz,^cms
```

### Configuration Syntax

Application/model configuration via `--config`:
- `--config <key>=<value>` - Single key-value
- `--config <file.yaml>` - YAML configuration file
- `--config <key>=@<file>` - Read value from file

**Examples:**
```bash
juju deploy mediawiki --config name='my wiki'
juju model-config image-stream=daily
juju config mysql --file=mycfg.yaml
```

### Resource Specification

Resources for charms via `--resource`:
- `--resource <name>=<path>` - Local file
- `--resource <name>=<revision>` - Specific revision

**Examples:**
```bash
juju deploy foo --resource bar=/some/file.tgz
```

### Endpoint Specification

For integrate/relate commands:
```
<application>[:<endpoint>]
```

**Examples:**
```bash
juju integrate wordpress:db mysql:server
juju integrate wordpress percona-cluster
```

### Cross-Model Offers

Remote offer URLs:
```
[<controller>:][<qualifier>/]<model>.<application>
```

**Examples:**
```bash
juju consume secrets.easyrsa
juju consume other-controller:prod.mysql
```

## Environment Variables

| Variable | Purpose | Used By |
|----------|---------|--------|
| `JUJU_MODEL` | Default model | Most model commands |
| `JUJU_CONTROLLER` | Default controller | Controller commands |
| `JUJU_DATA` | Configuration directory | All commands |
| `JUJU_LOGGING_CONFIG` | Logging configuration | All commands |
| `JUJU_DEV_FEATURE_FLAGS` | Developer features | All commands |
| `JUJU_FEATURES` | User features | All commands |
| `JUJU_STARTUP_LOGGING_CONFIG` | Startup logging | All commands |
| `JUJU_STATUS_ISO_TIME` | ISO timestamp format | `status` command |
| `JUJU_CONTAINER_TYPE` | Container type | MAAS provider |
| `XDG_DATA_HOME` | XDG data directory | Configuration path |

## Flag Naming Conventions

### Short vs Long Forms
- Short flags: Single letter, single hyphen (`-m`, `-c`, `-n`)
- Long flags: Word, double hyphen (`--model`, `--controller`, `--num-units`)

### Common Short Flag Assignments
| Short | Long | Purpose |
|-------|------|--------|
| `-m` | `--model` | Model selection |
| `-c` | `--controller` | Controller selection |
| `-n` | `--num-units` | Unit count |
| `-o` | `--output` | Output file |
| `-f` | `--file` | Input file |
| `-B` | `--no-browser-login` | Disable browser auth |

### Boolean Flags
Boolean flags use `--flag` for true, `--no-flag` for false:
- `--color` / `--no-color`
- `--prompt` / `--no-prompt`
- `--wait` / `--no-wait`

## Argument Validation

### Required Arguments
Commands validate required positional arguments and return error:
```
ERROR no charm name specified
```

### Argument Types
| Type | Validation |
|------|------------|
| `<application>` | Must exist in model |
| `<unit>` | Format: `<app>/<n>` |
| `<machine>` | Integer or `<machine>/<container>/<n>` |
| `<model>` | Model name or UUID |
| `<charm>` | Charm name or path |
| `<bundle>` | Bundle name or path |

### Multi-Value Arguments
Commands accepting multiple values use `...` notation:
```bash
juju remove-application app1 app2 app3
juju cancel-task task1 task2
```

### Key=Value Arguments
Used for configuration and constraints:
```bash
juju config app key=value
juju model-config key=value
```

Key value can be read from file using `@`:
```bash
juju config app key=@filename
```