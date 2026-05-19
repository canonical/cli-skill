# Juju CLI Argument Structure

## Common Argument Patterns

Juju commands share a set of common flag patterns due to the layered command framework. Most commands inherit flags from `modelcmd.ModelCommandBase` or `modelcmd.CommandBase`.

### Global Flags (Inherited by Most Commands)

| Flag | Short | Type | Default | Description | Env Var |
|---|---|---|---|---|---|
| `--model` | `-m` | string | current model | Model to operate in | |
| `--no-browser-login` | `-B` | bool | false | Do not use web browser for authentication | |
| `--format` | | string | varies | Output format (yaml, json, tabular) | |
| `--output` | `-o` | string | stdout | Specify an output file | |
| `--debug` | | bool | false | Enable debug logging | |
| `--verbose` | | bool | false | Show more verbose output | |
| `--quiet` | | bool | false | Suppress all non-error output | |
| `--show-log` | | bool | false | Show log messages | |
| `--logging-config` | | string | "" | Specify log levels for modules | |

### Controller-Agnostic Commands

Commands that do not require a model context (e.g., `bootstrap`, `clouds`, `login`) may skip model flags. These typically inherit from `modelcmd.OptionalControllerCommand` or `modelcmd.ControllerCommandBase`.

### Positional Arguments

Juju uses positional arguments sparingly and primarily for object identifiers:

- **Single identifier**: `juju show-application <application>`
- **Multiple identifiers**: `juju remove-machine <machine> [<machine>...]`
- **Key-value pairs**: `juju config <application> <key>=<value> [<key>=<value>...]`
- **No positional args**: `juju status`, `juju controllers`

### Key=Value Flags and Arguments

A pervasive pattern across Juju is the use of `key=value` syntax:

| Context | Example |
|---|---|
| Application config | `juju config mysql max-connections=100` |
| Model config | `juju model-config logging-config='<root>=INFO'` |
| Constraints | `juju deploy mysql --constraints "mem=4G cores=2"` |
| Storage | `juju deploy mysql --storage data=100G` |
| Bundle overlays | `juju deploy bundle.yaml --overlay overlay.yaml` |

## Command Argument Map

The following table maps argument patterns by command category. It is not exhaustive but covers the most significant structural patterns.

| Command | Positional Args | Required Flags | Optional Flags | Special Patterns |
|---|---|---|---|---|
| `bootstrap` | `<cloud>[/<region>]` | | `--constraints`, `--config`, `--bootstrap-constraints`, `--build-agent`, `--agent-version` | Key=value config pairs |
| `add-model` | `<model>` | | `--config`, `--owner`, `--credential`, `--cloud`, `--region` | |
| `deploy` | `<charm/bundle>` | | `--config`, `--constraints`, `--storage`, `--bind`, `--num-units`, `--to`, `--channel`, `--revision`, `--resource`, `--attach-storage`, `--force` | Storage directives, resource refs |
| `config` | `<application>` | | | key[=value] pairs as positional args |
| `model-config` | | | | key[=value] pairs as positional args |
| `controller-config` | | | | key[=value] pairs as positional args |
| `constraints` | `<application>` | | | |
| `set-constraints` | `<application>` | | | `key=value` constraints as positional args |
| `model-constraints` | | | | |
| `set-model-constraints` | | | | `key=value` constraints as positional args |
| `status` | | | `--format`, `--relations`, `--storage`, `--units`, `--color`, `--no-color` | `--format` supports yaml, json, tabular, summary, oneline |
| `show-status-log` | `<entity>` | | `--format`, `--limit` | |
| `add-unit` | `<application>` | | `--num-units`, `--to`, `--attach-storage` | |
| `remove-unit` | `<unit> [<unit>...]` | | `--destroy-storage`, `--force`, `--no-wait` | |
| `remove-application` | `<application> [<application>...]` | | `--destroy-storage`, `--force`, `--no-wait` | |
| `integrate` | `<application> <application>` | | `--via`, `--allow-cross-model` | |
| `remove-relation` | `<application>[:<endpoint>] <application>[:<endpoint>]` | | `--force` | |
| `expose` | `<application>` | | `--endpoints`, `--to-spaces`, `--to-cidrs` | |
| `unexpose` | `<application>` | | `--endpoints` | |
| `trust` | `<application>` | | `--scope`, `--remove` | Boolean toggle via flag |
| `refresh` | `<application>` | | `--channel`, `--revision`, `--switch`, `--force`, `--no-wait`, `--resource` | |
| `scale-application` | `<application> <scale>` | | `--attach-storage` | Integer scale as positional |
| `add-machine` | | | `--constraints`, `--disks`, `--series`, `--base`, `--to`, `--model` | Zone/constraint directives |
| `remove-machine` | `<machine> [<machine>...]` | | `--force`, `--no-keep-instance`, `--no-wait` | |
| `ssh` | `<target>` | | `--proxy`, `--no-host-key-checks`, `--ssh-option`, `--container` | Remainder passed to ssh |
| `scp` | `<source> <dest>` | | `--proxy`, `--no-host-key-checks`, `--ssh-option`, `--container` | Machine/unit paths with `:` prefix |
| `run` | `<unit> [<unit>...] <action-name>` | | `--wait`, `--params`, `--format` | Action params as JSON/YAML file |
| `exec` | `<command>` | `--all`, `--machine`, `--unit` | `--model`, `--timeout` | Double-dash separation for remote command |
| `add-cloud` | `<cloud>` | | `--file`, `--replace`, `--force`, `--credential` | Interactive or file-based input |
| `add-credential` | `<cloud>` | | `--file`, `--replace`, `--client`, `--controller` | Interactive or file-based input |
| `clouds` | | | `--format`, `--all`, `--client`, `--controller` | |
| `credentials` | | | `--format`, `--show-secrets`, `--client`, `--controller` | |
| `add-user` | `<user>` | | `--models`, `--acl`, `--controller` | |
| `grant` | `<user> <access>` | | `--model`, `--controller`, `--acl` | Access level positional |
| `revoke` | `<user> <access>` | | `--model`, `--controller` | |
| `add-secret` | `<name>` | | `--info`, `--file`, `--owner`, `--rotate-policy`, `--expire`, `--description` | Key=value data pairs |
| `update-secret` | `<name>` | | `--info`, `--file`, `--rotate-policy`, `--expire`, `--description` | Key=value data pairs |
| `add-storage` | `<unit> <storage-directive>` | | | Storage directive syntax: `name=pool,size,count` |
| `attach-storage` | `<unit> <storage>` | | | |
| `create-storage-pool` | `<name> <provider>` | | `--config` | Key=value config pairs |
| `add-space` | `<name> [<CIDR>...]` | | | CIDR blocks as positional |
| `move-to-space` | `<space> <CIDR>` | | `--force` | |
| `offer` | `<application>[:<endpoints>]` | | `--controller`, `--model` | |
| `consume` | `<offer-url>` | | `--application-alias` | |
| `register` | `<controller-name> [<api-endpoint>]` | | `--replace` | |
| `destroy-model` | `[<model>]` | | `--force`, `--no-wait`, `--destroy-storage`, `--release-storage`, `--timeout`, `-y` | Confirmation prompt |
| `destroy-controller` | `<controller>` | | `--force`, `--no-wait`, `--destroy-all-models`, `--destroy-storage`, `--release-storage`, `--timeout`, `-y` | Confirmation prompt |
| `kill-controller` | `<controller>` | | `--timeout` | Forceful termination |
| `create-backup` | | | `--no-download`, `--filename` | |
| `download-backup` | `<id>` | | `--filename` | |
| `enable-command` | `<command-set>` | | | Block command set names |
| `disable-command` | `<command-set>` | | `--message` | |
| `set-firewall-rule` | `<rule-name>` | | `--allowlist`, `--well-known` | |

## Special Arguments

### Constraints Syntax

Constraints appear as space-separated `key=value` pairs or as a single quoted string:

```
juju deploy mysql --constraints "mem=4G cores=2 arch=amd64"
juju set-constraints mysql mem=4G cores=2
```

Supported constraint keys include: `arch`, `container`, `cores`, `cpu-power`, `mem`, `disk`, `spaces`, `tags`, `virt-type`, `zones`, `root-disk`, `root-disk-source`, `instance-type`.

The `spaces` constraint uses a special syntax with `^` for negative matching:
```
juju deploy mysql --constraints spaces=dmz,^cms,^database
```

### Storage Directives

Storage directives use comma-separated pool, size, and count:
```
juju deploy mysql --storage data=ebs,100G,1
juju add-storage mysql/0 data=ebs,50G
```

### Resource References

Resources are referenced by name and revision or path:
```
juju deploy mysql --resource mysql-image=7
juju attach-resource mysql mysql-image=./local-image.tar
```

### Endpoint Syntax for Relations

Relations can be specified with endpoint names:
```
juju integrate mysql:db wordpress:mysql
juju remove-relation mysql:db wordpress
```

### Machine/Unit Targets for SSH

SSH targets use `machine:` or `unit:` prefix, or bare identifiers:
```
juju ssh 0
juju ssh mysql/0
juju ssh mysql/0 -- ls /var/lib/mysql
```

### Application Aliases for Cross-Model Relations

When consuming offers, an alias can be assigned:
```
juju consume admin/controller.mysql mysql-alias
```

### Double-Dash Separator for `exec`

The `exec` command uses `--` to separate juju flags from the remote command:
```
juju exec -u mysql/0 -- df -h
```

### Bundle Overlays

The `deploy` command accepts overlay files to modify bundle behavior:
```
juju deploy bundle.yaml --overlay production.yaml --overlay secrets.yaml
```
