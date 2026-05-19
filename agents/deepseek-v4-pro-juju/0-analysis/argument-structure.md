# Argument Structure — Juju CLI

## Introduction

Juju follows the **flat command hierarchy** pattern with position-dependent arguments. Commands are verbs (e.g., `deploy`, `add-model`) or verb-noun compounds (e.g., `add-cloud`, `remove-relation`, `show-status-log`). Listing commands use the bare noun plural form (e.g., `models`, `controllers`, `secrets`).

### Common Argument Patterns

1. **Prefix pattern**: Most commands use verb-noun (`add-*`, `remove-*`, `show-*`, `update-*`, `list-*` → converted to noun plurals for listing)
2. **Positional arguments**: Typically one or two positional parameters for resource identifiers (e.g., `juju deploy <charm> [<app-name>]`)
3. **Flag-based options**: Additional configuration via long flags with GNU-style `--flag` syntax
4. **Global injectors**: `-m/--model` and `-c/--controller` flags are added by the base type, not individual commands
5. **Key=value flags**: `--config key=value`, `--resource name=path`, `--storage name=size,type` use key=value pairs
6. **Repeatable accumulators**: `--config`, `--storage`, `--resource` can be specified multiple times
7. **Comma-separated list flags**: `--endpoints`, `--to-spaces`, `--to-cidrs` use comma-separated values

### Flag Style

- All flags use **long** form (`--flag-name`). Short single-character flags exist but are rare.
- The flag prefix is referred to as `"option"` internally (gnuflag `FlagKnownAs`).
- Flag names use **kebab-case** (e.g., `--bootstrap-base`, `--agent-version`, `--no-host-key-checks`).
- Boolean flags are toggle-only; no `--no-` prefix convention is enforced (some use `--no-wait`, others use `--no-retry`).
- Per DE013, Juju **offers both short and long flags** for many options (e.g., `-f`/`--file`, `-t`/`--timeout`, `-m`/`--model`), violating the standard that states "do not offer both short and long flags for the same action."

### Base Type Injected Arguments

These are added by the command base class hierarchy and available on almost every command:

| Flag | Type | Description |
|------|------|-------------|
| `-m, --model <name>` | string | Target model identifier (`[<controller>:]<model>`) |
| `-c, --controller <name>` | string | Target controller name |
| `--no-color` | bool | Disable ANSI color output |

### Output Format Argument

Most commands that produce structured output support the `--format` flag via the embedded `cmd.Output` struct:

| Flag | Type | Default | Values |
|------|------|---------|--------|
| `--format <fmt>` | string | `smart` | `smart`, `yaml`, `json` |
| `-o, --output <path>` | string | stdout | File path for output |

---

## Per-Command Argument Map

### bootstrap

```
juju bootstrap [<cloud>[/<region>] [<controller-name>]] [flags]
```

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<cloud>` | positional | No | Cloud to bootstrap into (defaults to configured) |
| `<region>` | positional (in cloud) | No | Cloud region (e.g., `aws/us-east-1`) |
| `<controller-name>` | positional | No | Name for the new controller |
| `--constraints` | key=value | No | Model constraints |
| `--bootstrap-constraints` | key=value | No | Bootstrap machine constraints |
| `--bootstrap-base` | string | No | Base for bootstrap machine |
| `--bootstrap-image` | string | No | Image ID for bootstrap machine |
| `--build-agent` | bool | No | Build local agent binary |
| `--metadata-source` | string | No | Path to agent/image metadata |
| `--to` | string | No | Placement directive |
| `--keep-broken` | bool | No | Keep failed bootstrap environment |
| `--auto-upgrade` | bool | No | Auto-upgrade to latest patch |
| `--agent-version` | string | No | Agent binary version |
| `--credential` | string | No | Credential name |
| `--config` | repeatable key=value | No | Model config values |
| `--model-default` | repeatable key=value | No | Model default config |
| `--storage-pool` | repeatable key=value | No | Storage pool definitions |
| `--clouds` | bool | No | Print available clouds |
| `--regions` | string | No | Print regions for named cloud |
| `--no-switch` | bool | No | Don't switch to new controller |
| `--force` | bool | No | Bypass checks |
| `--controller-charm-path` | string | No | Path to local controller charm |
| `--controller-charm-channel` | string | No | Charmhub channel for controller charm |
| `--db-snap` / `--db-snap-channel` | string | No | DQLite snap channel |

### deploy

```
juju deploy <charm-or-bundle> [<application-name>] [flags]
```

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<charm-or-bundle>` | positional | Yes | Charm/bundle URL or path |
| `<application-name>` | positional | No | Custom application name |
| `--channel` | string | No (`stable`) | Charmhub channel |
| `--config` | repeatable key=value | No | Application config |
| `--n` | int | No (`1`) | Number of units |
| `--trust` | bool | No | Grant credential access |
| `--overlay` | repeatable string | No | Overlay bundle files |
| `--constraints` | key=value | No | Application constraints |
| `--base` | string | No | Target OS base |
| `--revision` | int | No (`-1` = latest) | Charm revision |
| `--dry-run` | bool | No | Preview without deploying |
| `--force` | bool | No | Bypass validation |
| `--storage` | repeatable key=value | No | Storage directives |
| `--device` | repeatable key=value | No | Device constraints |
| `--resource` | repeatable key=value | No | Resource upload specs |
| `--bind` | string | No | Endpoint-to-space bindings |
| `--map-machines` | string | No | Use existing machines (bundles) |
| `--series` / `--base` | string | No | (deprecated vs current) OS selection |

### add-model

```
juju add-model <name> [<cloud>[/<region>]] [flags]
```

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<name>` | positional | Yes | Model name |
| `[<cloud>[/<region>]]` | positional | No | Cloud and region |
| `--credential` | string | No | Cloud credential |
| `--config` | repeatable key=value | No | Model configuration |

### destroy-model / destroy-controller

```
juju destroy-model [<model-name>] [flags]
juju destroy-controller <controller-name> [flags]
```

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `--destroy-storage` | bool | No | Remove attached storage |
| `--release-storage` | bool | No | Release storage without destroying |
| `--force` | bool | No | Force destruction |
| `--no-wait` | bool | No | Don't wait for completion |
| `--t / --timeout` | duration | No | Timeout per step |
| `--model-timeout` | duration | No | Timeout per model step |

### migrate

```
juju migrate <model-name> <target-controller> [flags]
```

| Argument | Type | Required | Description |
|----------|------|----------|-------------|
| `<model-name>` | positional | Yes | Model to migrate |
| `<target-controller>` | positional | Yes | Destination controller |

### List Commands (models, controllers, users, spaces, etc.)

Pattern: `juju <plural-noun> [flags]`

| Command | Positional Args | Notable Flags |
|---------|-----------------|---------------|
| `controllers` | None | `--refresh` |
| `models` | None | `--user`, `--uuid` |
| `users` | None | None |
| `machines` | None | None |
| `spaces` | [name] | `--format` |
| `subnets` | `[--space] [--zone]` | `--space`, `--zone` |
| `credentials` | None | `--format` |
| `clouds` | None | `--all` |
| `storage` | None | `--filesystem`, `--volume` |
| `storage-pools` | [name] | `--name`, `--provider`, `--format` |
| `operations` | None | `--format`, `--limit`, `--offset` |
| `actions` | `<application>` | `--format`, `--schema` |
| `resources` | `<app-or-unit>` | `--details` |
| `secrets` | None | `--format` |
| `secret-backends` | None | `--format` |
| `offers` | None | `--format`, `--interface`, `--application` |
| `firewall-rules` | None | `--format` |
| `disabled-commands` | None | None |
| `ssh-keys` | None | `--full` |
| `regions` | `<cloud>` | `--format` |

### Show Commands (show-model, show-controller, etc.)

Pattern: `juju show-<noun> [<identifier>] [flags]`

| Command | Positional Args | Notable Flags |
|---------|-----------------|---------------|
| `show-model` | [name] | `--format` |
| `show-controller` | [name] | `--format`, `--show-password` |
| `show-machine` | [machine-id] | `--format` |
| `show-cloud` | `<cloud>` | `--format`, `--include-config` |
| `show-credential` | [cloud credential] | `--format`, `--show-secrets` |
| `show-application` | [app] | `--format` |
| `show-unit` | `<unit>` | `--format` |
| `show-storage` | `<storage-id>` | `--format` |
| `show-space` | [name] | `--format` |
| `show-offer` | [endpoint] | `--format` |
| `show-action` | `<app> <action>` | `--format` |
| `show-operation` | `<operation-id>` | `--format` |
| `show-task` | `<task-id>` | `--format` |
| `show-secret` | [id] | `--format`, `--reveal` |
| `show-secret-backend` | [name] | `--format` |
| `show-user` | [user] | `--format` |
| `show-status-log` | `<entity>` | `--type`, `-n`, `--days`, `--from-date`, `--utc` |

### Add / Remove Commands

Pattern: `juju <verb>-<noun> <identifier> [flags]`

| Command | Positional Args | Notable Flags |
|---------|-----------------|---------------|
| `add-cloud` | `<name> [file]` | `--force`, `-f/--file`, `--credential` |
| `remove-cloud` | `<name>` | None |
| `add-credential` | `[cloud [name]]` | `-f/--file` (YAML) |
| `remove-credential` | `<cloud> <name>` | None |
| `update-credential` | `[[cloud] [name]]` | `-f/--file`, `--force`, `--region` |
| `add-k8s` | `<name>` | `--cluster-name`, `--region`, `--storage`, `--credential` |
| `remove-k8s` | `<name>` | `--force` |
| `add-machine` | None | `--constraints`, `--base`, `--series`, `--disks`, `-n/--num` |
| `remove-machine` | `<machine-id...>` | `--force`, `--no-wait`, `--keep-instance` |
| `add-unit` | `[<app>]` | `-n/--num-units`, `--to`, `--attach-storage` |
| `remove-unit` | `<unit...>` | `--force`, `--no-wait`, `--dry-run`, `--destroy-storage`, `--num-units` |
| `remove-application` | `<app...>` | `--force`, `--no-wait`, `--dry-run`, `--destroy-storage` |
| `remove-relation` | `<app1> <app2>` | `--force` |
| `remove-saas` | `<saas-name...>` | `--force`, `--no-wait` |
| `add-space` | `<name> [cidr...]` | None |
| `remove-space` | `<name>` | None |
| `add-storage` | `<unit> <storage-name>[=<constraints>]` | None |
| `attach-storage` | `<storage> <unit>` | None |
| `detach-storage` | `<storage> [unit...]` | `--force` |
| `remove-storage` | `<storage...>` | `--force`, `--no-wait`, `--destroy-storage` |
| `add-secret` | `<name>` | `--file`, `--secret`, `--label`, `--description` |
| `remove-secret` | `[id]` | `--revision` |
| `add-secret-backend` | `<name> <type> [key=value...]` | None |
| `remove-secret-backend` | `<name>` | None |
| `add-user` | `<username> [display-name]` | `--controller` |
| `remove-user` | `<username>` | `--yes`, `--force` |

### Config Commands

Pattern: `juju config <app> [key[=value]]`

| Command | Args | Flags |
|---------|------|-------|
| `config` | `<app> [key[=value]]` | `--file`, `--color`, `--no-color`, `--reset` |
| `model-config` | `[key[=value]]` | `--file`, `--color`, `--no-color`, `--reset` |
| `controller-config` | `[key[=value]]` | `--file`, `--color`, `--no-color`, `--reset` |
| `model-defaults` | `[key[=value]]` | `--file`, `--color`, `--no-color`, `--reset`, `--cloud` |

### Constraints Commands

| Command | Positional Args | Flags |
|---------|-----------------|-------|
| `constraints` | `<app>` | `--format` |
| `set-constraints` | `<app> <constraints>` | None |
| `model-constraints` | None | `--format` |
| `set-model-constraints` | `<constraints>` | None |

### Grant / Revoke

| Command | Positional Args |
|---------|-----------------|
| `grant` | `<user> <permission> [<model>]` |
| `revoke` | `<user> <permission> [<model>]` |
| `grant-cloud` | `<user> <permission> <cloud>` |
| `revoke-cloud` | `<user> <permission> <cloud>` |
| `grant-secret` | `<secret-id> [app-or-unit...]` |
| `revoke-secret` | `<secret-id> [app-or-unit...]` |

### SSH / Debug Commands

| Command | Args | Notable Flags |
|---------|------|---------------|
| `ssh` | `[<target>] [<cmd>]` | `--proxy`, `--no-host-key-checks`, `--pty`, `--container` |
| `scp` | `<source...> <dest>` | `--proxy`, `--no-host-key-checks`, `-r/--recursive` |
| `debug-hooks` | `<unit> [<hook>...]` | None |
| `debug-code` | `<unit>` | `--at` (`all`, `hook`, etc) |
| `debug-log` | None | `--replay`, `--level`, `--include`, `--exclude`, `--include-module`, `--exclude-module`, `--limit`, `--no-tail`, `--retry`, `--retry-delay` |

### status

```
juju status [<selector>] [flags]
```

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `[<selector>]` | positional | No | Pattern to filter (e.g., `app/0`) |
| `--format` | string | `tabular` | `tabular`, `yaml`, `json`, `oneline`, `summary` |
| `--color` | bool | false | Enable ANSI color |
| `--no-color` | bool | false | Disable ANSI color |
| `--utc` | bool | false | UTC timestamps |
| `--relations` | bool | false | Show relations section |
| `--integrations` | bool | false | Same as `--relations` |
| `--storage` | bool | false | Show storage section |
| `--watch` | duration | No | Watch mode refresh interval |
| `--retry-count` | int | 3 | Retry count |
| `--retry-delay` | duration | 100ms | Retry delay |

### run / exec

| Command | Args | Flags |
|---------|------|-------|
| `run` | `<unit...> <cmd> [args...]` | `--operator`, `--timeout`, `--format`, `--background`, `--wait` |
| `exec` | `<unit...> <cmd> [args...]` | `--operator`, `--timeout`, `--format`, `--background`, `--wait` (similar to `run`) |

### secrets / secret-backends

These follow CRUD patterns:

| Command | Args | Flags |
|---------|------|-------|
| `add-secret` | `<name>` | `--file`, `--secret`, `--label`, `--description`, `--rotate`, `--expire` |
| `update-secret` | `<id>` | `--file`, `--secret`, `--label`, `--description`, `--rotate`, `--expire` |
| `remove-secret` | `<id>` | `--revision` |
| `secrets` | None | `--format`, `--owner`, `--reveal` |
| `show-secret` | `[id]` | `--format`, `--reveal`, `--revision` |

### integrate (relation)

```
juju integrate <app>[:<endpoint>] <app>[:<endpoint>] [flags]
```

| Flag | Type | Description |
|------|------|-------------|
| `--via` | string | CIDRs for cross-model egress |

### suspend-relation / resume-relation

```
juju suspend-relation <app1> <app2> [--message <reason>]
juju resume-relation <app1> <app2>
```

### crossmodel (CMR) Commands

| Command | Args | Flags |
|---------|------|-------|
| `offer` | `<app>:<endpoint> <name>` | `--controller` |
| `remove-offer` | `<offer-name>` | `--force`, `--yes` |
| `show-offer` | `[endpoint]` | `--format` |
| `offers` | None | `--format`, `--interface`, `--application`, `--connected-user` |
| `find-offers` | None | `--format`, `--interface`, `--url` |
| `consume` | `[<controller>:]<offer>` [alias] | `--controller-alias` |
| `remove-saas` | `<saas...>` | `--force`, `--no-wait` |

---

## Special Arguments

### `--` Separator

Juju does **not** use `--` as a separator between juju flags and command arguments. Commands like `run`/`exec` pass arguments after the command verbatim without requiring a separator token. This is unlike tools such as `docker exec`, `kubectl exec`, or `multipass exec` which use `--` to separate the CLI's own flags from the payload command's flags.

### `=` in Flag Values

Flagname=value is supported (gnuflag default). Both `--flag value` and `--flag=value` are valid. Commands like `--config key=value` embed `=` within the value itself, but the `=` separating the flag name from its value is the first `=`.

### Repeatable Multi-Value Flags

The standard supports two patterns (repeat the flag, or comma-separated values in one flag). Juju uses **both** patterns inconsistently:

- **Repeat for singular**: `--config key1=val1 --config key2=val2` (singular `--config`)
- **Comma in values**: `--endpoints ep1,ep2` (list of endpoints)
- **Repeat for singular resource**: `--resource name=path --resource name2=path2` (singular `--resource`)

Per DE013, singular flag name means repeatable; plural means comma-separated. Juju's usage is largely consistent with this: singular `--config`, `--storage`, `--resource` are repeatable.

### Deprecated / Legacy Flags

- `--series` is largely deprecated in favor of `--base`
- `--channel` accepts `latest/stable` style (channels as paths)
- `-m` / `--model` and `-c` / `--controller` are injected by the framework
- `-f` and `--file` both exist as duplicate short/long forms (DE013 violation)

### Boolean Flag Behavior

Boolean flags default to `false` and are set to `true` when present. There is no `--no-<flag>` built-in convention; negation is done via separate flags like `--no-wait`, `--no-retry`, `--no-host-key-checks`, `--no-switch`, `--no-color`, `--no-tail`. This is verbose and inconsistent with some DE013 guidance.

### Env Var Mappings

| Env Var | Maps To |
|---------|---------|
| `JUJU_DATA` | XDG data home directory override |
| `JUJU_MODEL` | Default model name |
| `JUJU_CONTROLLER` | Default controller name |
| `JUJU_FEATURES` | Feature flag toggles |
| `JUJU_LOGGING_CONFIG` | Logging configuration |
| `NO_COLOR` | Disable ANSI color output |
| `HTTP_PROXY`/`HTTPS_PROXY`/`NO_PROXY` | Proxy configuration |
