# Juju Verb Taxonomy

## Overview

Juju uses a relatively small number of high-frequency verbs, but several of them carry multiple meanings depending on domain. This taxonomy groups verbs by operator intent rather than by source package.

## 1. CRUD-like resource management verbs

### `add`

Meaning: create or register a new named resource.

Examples:
- `add-model`
- `add-cloud`
- `add-credential`
- `add-space`
- `add-user`
- `add-secret`

Assessment:
- strong and predictable
- one of the best normalized verbs in the CLI

### `remove`

Meaning: delete, detach permanently, or unregister a resource from Juju control.

Examples:
- `remove-application`
- `remove-unit`
- `remove-space`
- `remove-cloud`
- `remove-secret`

Assessment:
- mostly strong
- semantics are broad, ranging from resource deletion to relationship teardown to client-local removal

### `show`

Meaning: display detailed information for one named resource.

Examples:
- `show-model`
- `show-user`
- `show-space`
- `show-offer`
- `show-storage`

Assessment:
- strong and understandable
- not used everywhere it could be, which creates list/show asymmetry

### plural noun listings

Meaning: list a class of resources.

Examples:
- `models`
- `controllers`
- `users`
- `spaces`
- `subnets`
- `offers`

Assessment:
- workable, but this is a separate grammar from `show-*`
- Juju alternates between plural-noun and `list-*` styles

## 2. Lifecycle and orchestration verbs

### `deploy`

Meaning: instantiate an application or bundle.

Assessment:
- core Juju verb, highly appropriate
- intentionally special rather than normalized into `add-application`

### `refresh`

Meaning: update an application's charm or associated deployment metadata.

Assessment:
- good domain verb, but overlaps conceptually with `upgrade`

### `upgrade`

Meaning: upgrade Juju agents/controller/model, not application charm refresh.

Examples:
- `upgrade-model`
- `upgrade-controller`

Assessment:
- meaningful distinction exists, but users must learn the split between app `refresh` and Juju `upgrade`

### `scale`

Meaning: set desired replica/unit count, currently explicit in `scale-application`.

Assessment:
- clear
- the explicit noun helps because scaling is currently restricted in scope

### `bootstrap`

Meaning: initialize a controller environment.

Assessment:
- canonical Juju term, not replaceable with a generic CRUD verb

### `migrate`

Meaning: move a model to another controller.

Assessment:
- clear and specialized

## 3. Relationship and connectivity verbs

### `integrate`

Meaning: create a relation between applications.

Assessment:
- good modern verb, less infrastructure-flavored than `relate`
- counterpart is still `remove-relation`, not `disintegrate` or `separate`, which is sensible but asymmetrical

### `offer` / `consume`

Meaning:
- `offer`: publish endpoints for other models
- `consume`: import a remote offer

Assessment:
- conceptually strong domain pair
- good example of semantically complementary verbs that are not simple lexical opposites

### `bind`

Meaning: assign application endpoints to spaces.

Assessment:
- accurate but niche; discoverability depends on docs

### `expose` / `unexpose`

Meaning: make an application reachable from outside the model / revoke that public reachability.

Assessment:
- strong paired verbs
- better than generic `enable` / `disable` here because they describe networking intent

## 4. Access-control verbs

### `grant` / `revoke`

Meaning: assign or remove permissions.

Assessment:
- strong and conventional
- somewhat overloaded because they apply to models, controllers, clouds, offers, and secrets via multiple commands

### `trust`

Meaning: grant elevated trust capability to an application.

Assessment:
- good domain term
- lacks an equally visible complement command name because untrust is folded into flags rather than a peer command

### `enable` / `disable`

Meaning: turn a capability or user state on/off.

Examples:
- `enable-user`, `disable-user`
- `enable-command`, `disable-command`
- `enable-destroy-controller`

Assessment:
- conventional, but semantically broad
- `enable-destroy-controller` is especially awkward because it names a workflow exception rather than a clean domain object

## 5. Query and reporting verbs

### `status`

Meaning: synthesized operational state report over model resources.

Assessment:
- a cornerstone command
- broader than `show-*`; effectively a dashboard/report verb

### `info`

Meaning: detailed repository information, currently Charmhub-focused.

Assessment:
- understandable but generic
- can be confused with `show-*` and with command help in large surfaces

### `find`

Meaning: search a catalog or external index.

Assessment:
- conventional
- better than `search` only because it is entrenched and paired with Charmhub mental models

### `debug-*`

Meaning: open debugging or log-inspection workflows.

Assessment:
- appropriate namespace prefix
- `debug-log` is a report stream, while `debug-hooks` and `debug-code` are interactive sessions

## 6. Data-transfer and remote-execution verbs

### `ssh` / `scp`

Meaning: pass through secure shell idioms inside Juju's model context.

Assessment:
- extremely discoverable for operators
- intentionally borrowed rather than normalized into Juju-specific verbs

### `exec`

Meaning: execute shell commands across remote targets.

Assessment:
- understandable, but adjacent to `ssh` and `run`

### `run`

Meaning: run an action on units.

Assessment:
- overloaded in a CLI ecosystem where `run` often means remote shell execution
- one of the more confusing verbs in the surface

## 7. Configuration verbs

### `config`

Meaning: app config get/set/reset.

Assessment:
- concise and established
- but creates a triad with `model-config` and `controller-config`, which are explicit while app config keeps the noun implicit

### `default-*`

Meaning: get/set/unset local defaults for a cloud dimension.

Assessment:
- acceptable but unusual compared with the rest of the CLI

### `set-*`

Meaning: imperative assignment of a specific config dimension.

Examples:
- `set-constraints`
- `set-model-constraints`
- `set-credential`
- `set-firewall-rule`

Assessment:
- useful where the target noun is a property rather than a standalone resource
- but creates overlap with `config`, `model-config`, and `controller-config`

## Verb collisions and ambiguities

### `run` vs `exec` vs `ssh`

- `run` means run a Juju action
- `exec` means run arbitrary shell commands remotely
- `ssh` means open a session or run a command over SSH

This is the clearest execution-verb collision in the CLI.

### `show` vs `info` vs `status`

- `show-*` means detailed resource inspection
- `info` means Charmhub package information
- `status` means synthesized live system report

The semantics are distinct, but users have to learn them by family.

### `config` vs `set-*`

Juju uses both:
- generalized config verbs: `config`, `model-config`, `controller-config`
- targeted set verbs: `set-constraints`, `set-credential`, `set-firewall-rule`

That is practical, but not lexically pure.

### `remove` vs `destroy` vs `kill`

- `remove-*`: delete or detach a scoped resource
- `destroy-*`: controlled teardown of model/controller scope
- `kill-controller`: emergency, forceful teardown

This is one of the better verb hierarchies in Juju because the severity gradient is visible in the verb choice.

## Assessment

The strongest Juju verbs are:
- `add`
- `remove`
- `show`
- `grant`
- `revoke`
- `deploy`
- `integrate`
- `expose` / `unexpose`
- `destroy` / `kill`

The weakest or most overloaded are:
- `run`
- `exec`
- `config` versus `set-*`
- `info` in a CLI that already has `show-*`
- `enable-destroy-controller` as a special-case exception name
