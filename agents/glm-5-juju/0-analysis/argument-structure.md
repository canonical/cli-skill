# Juju CLI Argument Structure

## Overview

Juju commands accept arguments through three mechanisms:
1. **Positional arguments** - Required or optional positional parameters
2. **Flags (options)** - Named parameters with short (`-x`) and long (`--xxx`) forms
3. **Environment variables** - Contextual overrides via `JUJU_*` variables

## Common Argument Patterns

### Global Flags

All commands inherit these flags from the SuperCommand:

| Flag | Short | Type | Description |
|------|-------|------|-------------|
| `--debug` | | bool | Enable debug logging (`--show-log --logging-config=<root>=DEBUG`) |
| `--help` | `-h` | bool | Show help for command |
| `--logging-config` | | string | Specify log levels for modules |
| `--quiet` | | bool | Suppress informational output |
| `--show-log` | | bool | Write log file to stderr |
| `--verbose` | | bool | Enable verbose output |

### Model-Scoped Command Flags

Commands that operate on a model (`ModelCommand`) include:

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--model` | `-m` | string | "" | Model to operate in (`[<controller>:]<model>\|<model UUID>`) |
| `--no-browser-login` | `-B` | bool | false | Do not use web browser for authentication |

### Controller-Scoped Command Flags

Commands that operate on a controller (`ControllerCommand`) include:

| Flag | Short | Type | Default | Description |
|------|-------|------|---------|-------------|
| `--controller` | `-c` | string | "" | Controller to operate in |
| `--no-browser-login` | `-B` | bool | false | Do not use web browser for authentication |

## Command Arguments by Category

### Infrastructure Commands

#### bootstrap
```
juju bootstrap [options] [<cloud name>[/region] [<controller name>]]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| cloud name/region | No | string | Target cloud and region (interactive if omitted) |
| controller name | No | string | Name for new controller |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--agent-version` | string | "" | Agent binary version |
| `--auto-upgrade` | bool | false | Upgrade to latest patch release after bootstrap |
| `--bootstrap-base` | string | "" | Base for bootstrap machine |
| `--bootstrap-constraints` | []string | [] | Bootstrap machine constraints |
| `--bootstrap-image` | string | "" | Image for bootstrap machine |
| `--build-agent` | bool | false | Build local agent binary |
| `--clouds` | bool | false | List available clouds |
| `--config` | | | Controller config (file or key=value) |
| `--constraints` | []string | [] | Model constraints |
| `--controller-charm-channel` | string | "4.1/stable" | Channel for controller charm |
| `--controller-charm-path` | string | "" | Path to local controller charm |
| `--credential` | string | "" | Credential to use |
| `--force` | bool | false | Bypass supported base checks |
| `--keep-broken` | bool | false | Don't destroy on failure |
| `--metadata-source` | string | "" | Local metadata path |
| `--model-default` | | | Model default config |
| `--no-switch` | bool | false | Don't switch to new controller |
| `--regions` | string | "" | List regions for a cloud |
| `--storage-pool` | | | Initial storage pool config |
| `--to` | string | "" | Placement directive |

#### add-cloud
```
juju add-cloud [<cloud name>] [<cloud definition file>]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| cloud name | No | string | Name for the cloud |
| cloud definition file | No | string | YAML file with cloud definition |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-f` | string | "" | Cloud definition file (alternative position) |
| `--force` | bool | false | Overwrite existing cloud |

#### add-k8s
```
juju add-k8s [<cloud name>]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| cloud name | No | string | Name for the Kubernetes cloud |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--controller` | `-c` | string | "" | Controller to add to |
| `--cloud` | | string | "" | Existing cloud to update |
| `--skip-storage` | bool | false | Skip storage discovery |

### Application Commands

#### deploy
```
juju deploy [options] <charm or bundle> [<application name>]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| charm or bundle | Yes | string | Charm name, URL, or local path |
| application name | No | string | Custom application name |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--attach-storage` | string | "" | Existing storage to attach |
| `--base` | string | "" | Deployment base (e.g., ubuntu@22.04) |
| `--bind` | string | "" | Endpoint-space bindings |
| `--channel` | string | "" | CharmHub channel |
| `--config` | | | Config file or key=value pairs |
| `--constraints` | []string | [] | Machine constraints |
| `--device` | | | Device constraints (K8s) |
| `--dry-run` | bool | false | Preview without deploying |
| `--force` | bool | false | Bypass checks |
| `--map-machines` | string | "" | Bundle machine mapping |
| `--num-units` | `-n` | int | 1 | Number of units |
| `--overlay` | []string | [] | Bundle overlay files |
| `--resource` | | | Resource name=file pairs |
| `--revision` | int | -1 | Charm revision |
| `--storage` | | | Storage directives |
| `--to` | string | "" | Placement directive |
| `--trust` | bool | false | Grant credential access |

#### add-unit
```
juju add-unit [options] <application name>
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| application name | Yes | string | Target application |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `-n` | `--num-units` | int | 1 | Units to add |
| `--attach-storage` | string | "" | Storage to attach |
| `--to` | string | "" | Placement directive |

#### remove-application
```
juju remove-application [options] <application name> [...]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| application name | Yes | string | Application to remove |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--force` | bool | false | Force removal without cleanup |
| `--destroy-storage` | bool | false | Destroy associated storage |
| `--no-prompt` | bool | false | Skip confirmation |

### Integration Commands

#### integrate / relate
```
juju integrate [options] <application1>[<endpoint name>] <application2>[<endpoint name>]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| application1[endpoint] | Yes | string | First application with optional endpoint |
| application2[endpoint] | Yes | string | Second application with optional endpoint |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--via` | string | "" | Cross-model relation via CIDR |

### Status Commands

#### status
```
juju status [options] [<selector> [...]]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| selector | No | string | Machine, application, or unit filter |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--color` | bool | false | ANSI color codes |
| `--format` | string | tabular | Output format (json\|line\|oneline\|short\|summary\|tabular\|yaml) |
| `--integrations` | bool | false | Show relations (alias: --relations) |
| `--no-color` | bool | false | Disable colors |
| `--output` | `-o` | string | "" | Output file |
| `--relations` | bool | false | Show relations section |
| `--retry-count` | int | 3 | API failure retries |
| `--retry-delay` | duration | 100ms | Retry delay |
| `--storage` | bool | false | Show storage section |
| `--utc` | bool | false | UTC timestamps |

### SSH Commands

#### ssh
```
juju ssh [options] <target> [<command> ...]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| target | Yes | string | Machine, unit, or container (e.g., 0, mysql/0, 0/lxd/1) |
| command | No | []string | Command to execute |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--container` | string | "" | Target container |
| `--no-host-key-checks` | bool | false | Disable host key verification |
| `--proxy` | bool | false | Proxy through controller |
| `--pty` | bool | true | Allocate pseudo-TTY |
| `--strict-host-key-checking` | bool | false | Strict host verification |

#### scp
```
juju scp [options] <source> <destination>
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| source | Yes | string | Source file (local or remote) |
| destination | Yes | string | Destination (local or remote) |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--container` | string | "" | Target container |
| `--no-host-key-checks` | bool | false | Disable verification |
| `--recursive` | `-r` | bool | false | Copy directories |

### Action Commands

#### run (action)
```
juju run [options] <action name> <target>
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| action name | Yes | string | Action to execute |
| target | Yes | string | Unit or application |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--params` | string | "" | Action parameters (YAML file or key=value) |
| `--wait` | duration | 0 | Wait timeout (0 = no wait) |

#### exec
```
juju exec [options] <command>
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| command | Yes | string | Command to execute |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--application` | `-a` | []string | Target applications |
| `--unit` | `-u` | []string | Target units |
| `--machine` | `-m` | []string | Target machines |
| `--format` | string | yaml | Output format |
| `--parallel` | bool | false | Execute in parallel |

### Secret Commands

#### add-secret
```
juju add-secret [options] <secret name> <key>=<value>...
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| secret name | Yes | string | Name for the secret |
| key=value | Yes | []string | Secret key-value pairs |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--owner` | string | "application" | Secret owner (application\|unit) |
| `--file` | string | "" | Load from YAML file |
| `--info` | string | "" | Description |
| `--expire` | duration | 0 | Expiration time |

### Storage Commands

#### add-storage
```
juju add-storage [options] <unit> <storage directive>
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| unit | Yes | string | Target unit |
| storage directive | Yes | string | Storage spec (name=size,count) |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--pool` | string | "" | Storage pool |
| `--size` | string | "" | Size in GB |

### User Management Commands

#### add-user
```
juju add-user [options] <user name> [...]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| user name | Yes | string | Username to add |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--display-name` | string | "" | Display name |
| `--model` | | | Model access grants |
| `--controller` | `-c` | "" | Controller access grant |

#### login
```
juju login [options] [<controller name>]
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| controller name | No | string | Controller to log into |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--user` | `-u` | string | "" | Username |
| `--password` | string | "" | Password (interactive if empty) |
| `--no-browser-login` | `-B` | bool | false | Skip browser login |
| `--token` | string | "" | JWT token |

### Config Commands

#### config
```
juju config [options] <application> [<key>[=<value>]]...
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| application | Yes | string | Target application |
| key=value | No | []string | Config key-value pairs |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--file` | `-f` | string | "" | YAML config file |
| `--reset` | []string | [] | Keys to reset |

#### model-config
```
juju model-config [options] [<key>[=<value>]]...
```

**Positional Arguments:**
| Argument | Required | Type | Description |
|----------|----------|------|-------------|
| key=value | No | []string | Config key-value pairs |

**Flags:**
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--file` | string | "" | YAML config file |
| `--reset` | []string | [] | Keys to reset |

## Special Arguments

### Placement Directives

Used with `--to` flags and in bundle YAML:

| Syntax | Description | Example |
|--------|-------------|---------|
| `<number>` | Existing machine | `23` |
| `lxd` | New LXD container on new machine | `lxd` |
| `lxd:<machine>` | New LXD container on machine | `lxd:25` |
| `<machine>/lxd/<n>` | Existing container | `24/lxd/3` |
| `zone=<name>` | Availability zone | `zone=us-east-1a` |
| `<host>.maas` | MAAS node | `node01.maas` |

### Storage Directives

Syntax: `[<count>,]<size>[,<pool>]`

| Component | Description | Example |
|-----------|-------------|---------|
| count | Number of instances | `2,10G` |
| size | Size specification | `10G`, `100M` |
| pool | Storage pool name | `10G,ebs-ssd` |

### Constraints Syntax

Space-separated key=value pairs:

| Key | Values | Example |
|-----|--------|---------|
| `arch` | amd64, arm64, armhf | `arch=amd64` |
| `cores` | Integer | `cores=4` |
| `mem` | Memory with unit | `mem=8G` |
| `root-disk` | Disk size | `root-disk=100G` |
| `spaces` | Space names | `spaces=dmz,^cms` |
| `tags` | Instance tags | `tags=production,ssd` |
| `virt-type` | Virtualization | `virt-type=kvm` |
| `zones` | Availability zones | `zones=us-east-1a` |

## Environment Variables

| Variable | Description | Affects |
|----------|-------------|---------|
| `JUJU_CONTROLLER` | Default controller | All model/controller commands |
| `JUJU_MODEL` | Default model | Model-scoped commands |
| `JUJU_DATA` | Config directory location | All commands |
| `JUJU_LOGGING_CONFIG` | Log configuration | All commands |
| `JUJU_DEV_FEATURE_FLAGS` | Developer features | All commands |
| `JUJU_FEATURES` | Production features | All commands |
| `JUJU_STATUS_ISO_TIME` | ISO timestamps in status | status command |

## Argument Precedence

1. **Command