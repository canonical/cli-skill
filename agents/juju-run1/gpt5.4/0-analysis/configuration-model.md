# Juju CLI Configuration Model

## Overview

Juju has several distinct configuration planes:
- local client configuration and context selection
- controller configuration
- model configuration
- application charm configuration
- cloud and credential defaults
- command-specific ephemeral flags

These planes are related, but not merged into one central configuration system. The CLI resolves them through command wrappers plus package-specific logic.

## 1. Local client configuration and context resolution

The first configuration layer is local client state, not server state.

Sources include:
- the Juju XDG data directory initialized during startup
- the client store for controllers, accounts, and current model/controller
- user alias file configured via `UserAliasesFilename`
- environment variables such as `JUJU_MODEL` and `JUJU_CONTROLLER`

For model-scoped commands in `modelcmd`, the effective target typically resolves as:
1. explicit command identifier or `-m/--model`
2. environment variable value if allowed
3. current controller/model from the local client store

For controller-scoped commands the equivalent pattern applies with `-c/--controller` and current-controller lookup.

This is a strong aspect of the CLI: the context-resolution model is consistent and shared by wrappers rather than duplicated per command.

## 2. Controller configuration

`juju controller-config` exposes controller-level settings.

Characteristics:
- scope: controller-wide
- read/write surface is schema-driven from controller config metadata
- supports `--file`, inline `key=value`, `--format`, and `--ignore-read-only-fields`
- default display format is `tabular`
- single-key retrieval with default tabular output uses `FormatSmart` instead of full table/YAML output

Precedence inside one invocation:
1. values loaded from `--file`
2. inline `key=value` arguments override file values
3. read-only fields can be ignored on import with `--ignore-read-only-fields`

Controller config is not resettable through the shared `ConfigCommandBase` path used here.

## 3. Model configuration

`juju model-config` exposes model-level settings.

Characteristics:
- scope: one model
- supports get-one, get-all, set, file import, and reset
- default display format is `tabular`
- can ignore read-only fields during YAML import
- provides many defaults that shape later operations, including deployment behavior

Documented precedence inside one invocation:
1. file input via `--file`
2. inline `key=value` arguments override file values
3. reset keys are applied as part of the same action set, but the same key cannot be both set and reset in one call

Important downstream effect:
- `default-base` from model config participates in deploy resolution when `--base` and bundle-level bases are absent

## 4. Application charm configuration

`juju config` manages application config exposed by the charm.

Characteristics:
- scope: one deployed application
- default format is `yaml`
- supports get-one, get-all, set, file import, and reset
- values can be set inline with `key=value`
- values can be set to file contents with the `@path` convention for some value flows

Documented precedence inside one invocation:
1. YAML loaded from `--file`
2. inline `key=value` arguments override file values
3. resets can be combined with sets, but not on the same key

This is one of the clearer and better-documented precedence models in the CLI.

## 5. Cloud and credential defaults

The cloud/credential family introduces another configuration plane stored locally and partly reflected to controllers.

Important commands:
- `default-region`
- `default-credential`
- `add-credential`
- `update-credential`
- `credentials`
- `clouds`, `regions`, `show-cloud`

Behaviorally, these defaults influence later commands such as bootstrap and cloud/model creation, but they are not expressed through one general-purpose precedence document. Operators need to know the family-specific defaults.

## 6. Deploy-time effective configuration

`deploy` composes several configuration sources.

Verified examples from help/docs:
- charm source and revision/channel/base from CLI flags
- application config from repeated `--config` values or YAML files
- constraints from `--constraints`
- storage and resource directives from dedicated flags
- trust state from `--trust`
- placement from `--to`

One documented precedence chain is explicit and important:

Final base selection, highest to lowest priority:
1. `--base`
2. bundle charm URL base
3. bundle top-level base
4. `default-base` model config
5. first base in the charm manifest

This is one of the few places where Juju documents a cross-plane precedence stack clearly.

## 7. Shared config parser behavior

The `cmd/juju/config.ConfigCommandBase` provides a common mini-language for config commands.

It determines actions in this order:
1. `SetFile` if `--file` is present
2. `Reset` if `--reset` is present
3. parse remaining args into `GetOne` or `SetArgs`
4. fall back to `GetAll`

Important invariants enforced centrally:
- get-one cannot be mixed with any other action
- multiple get keys are rejected
- the same key cannot be set and reset in one invocation
- file input is processed before CLI overrides, so CLI values win

This is a good design choice: it gives app config, model config, and controller config a shared mental model.

## 8. Non-config settings that still behave like configuration

Several features behave like configuration even though they use dedicated commands:
- block state via `disable-command` / `enable-command`
- model secret backend via `model-secret-backend`
- trust and access-control relationships via `trust`, `grant*`, `revoke*`
- cloud/client defaults via `default-region` and `default-credential`
- current focus via `switch`

From an operator perspective, these are all configuration. From a CLI surface perspective, they are split across distinct command families.

## Defaults and fallbacks

Important defaulting behavior verified in code and docs:
- current controller/model come from client store if not explicitly supplied
- `status` reads an environment variable for ISO time when `--utc` is not set
- output format defaults vary by command, not framework-wide
- app config defaults to YAML output, controller/model config default to tabular output
- some commands treat absence of a query or selector as a request for a full listing rather than an error

## Strengths

- Context resolution is robust and shared.
- Config command semantics are centrally reusable.
- File plus inline override behavior is consistent in the config family.
- Deployment docs clearly document the most important base precedence chain.

## Weaknesses and gaps

- There is no single user-facing model of all configuration planes together.
- Cloud defaults, client store state, and command flags are documented family by family rather than as one precedence system.
- Output defaults vary by command family and are not explained as a global policy.
- Some state that operators treat as configuration is expressed through one-off verbs rather than a coherent configuration grammar.

## Practical precedence summary

### Target selection
1. explicit command flags / qualified identifiers
2. relevant environment variable
3. local client store current controller/model

### App/model/controller config updates
1. values from `--file`
2. inline `key=value` overrides
3. resets applied where supported, but never on keys also being set

### Deploy base resolution
1. `--base`
2. bundle charm URL base
3. bundle top-level base
4. model `default-base`
5. charm manifest base

That is the effective configuration model Juju exposes today: several strong local precedence rules, but no single top-level abstraction stitching them together.
