# 01 — Verb-Noun Decomposition Matrix

## Overview

Every Juju command is decomposed into a verb and noun (resource type). The matrix below shows which verb-noun combinations exist in the command set. A `✓` indicates the command exists; `—` indicates the combination is absent.

## Decomposition Table

### Legend

Rows = Verbs (alphabetical), Columns = Nouns (by frequency)

### Full Decomposition of All Commands

| Command | Verb | Noun | Combined? |
|---------|------|------|-----------|
| `actions` | list | action | ✓ (plural noun) |
| `add-cloud` | add | cloud | ✓ |
| `add-credential` | add | credential | ✓ |
| `add-k8s` | add | k8s (cloud) | ✓ |
| `add-machine` | add | machine | ✓ |
| `add-model` | add | model | ✓ |
| `add-secret` | add | secret | ✓ |
| `add-secret-backend` | add | secret-backend | ✓ |
| `add-space` | add | space | ✓ |
| `add-ssh-key` | add | ssh-key | ✓ |
| `add-storage` | add | storage | ✓ |
| `add-unit` | add | unit | ✓ |
| `add-user` | add | user | ✓ |
| `attach-resource` | attach | resource | ✓ |
| `attach-storage` | attach | storage | ✓ |
| `autoload-credentials` | autoload | credential | — (compound verb) |
| `bind` | bind | (endpoint) | — (verb-only, implicit object) |
| `bootstrap` | bootstrap | (controller) | — (verb-only, implicit object) |
| `cancel-task` | cancel | task | ✓ |
| `change-user-password` | change | user-password | ✓ |
| `clouds` | list | cloud | ✓ (plural noun) |
| `config` | config | application | — (verb-only, implicit object) |
| `consume` | consume | offer | — (verb-only, implicit object) |
| `constraints` | get | constraint | — (implicitly application) |
| `controller-config` | config | controller | ✓ (noun-verb) |
| `controllers` | list | controller | ✓ (plural noun) |
| `create-backup` | create | backup | ✓ |
| `create-storage-pool` | create | storage-pool | ✓ |
| `credentials` | list | credential | ✓ (plural noun) |
| `dashboard` | open | dashboard | — (verb-only, implicit) |
| `debug-code` | debug | code | ✓ |
| `debug-hooks` | debug | hook | ✓ |
| `debug-log` | debug | log | ✓ |
| `default-credential` | set-default | credential | ✓ |
| `default-region` | set-default | region | ✓ |
| `deploy` | deploy | charm/application | — (verb-only, implicit object) |
| `destroy-controller` | destroy | controller | ✓ |
| `destroy-model` | destroy | model | ✓ |
| `detach-storage` | detach | storage | ✓ |
| `diff-bundle` | diff | bundle | ✓ |
| `disable-command` | disable | command (block) | ✓ |
| `disable-user` | disable | user | ✓ |
| `disabled-commands` | list | disabled-command | ✓ (plural noun) |
| `download` | download | charm | — (verb-only, implicit object) |
| `download-backup` | download | backup | ✓ |
| `dump-db` | dump | db | ✓ |
| `dump-model` | dump | model | ✓ |
| `enable-command` | enable | command (block) | ✓ |
| `enable-destroy-controller` | enable | destroy-controller | ✓ (meta) |
| `enable-user` | enable | user | ✓ |
| `exec` | exec | (command) | — (verb-only) |
| `export-bundle` | export | bundle | ✓ |
| `expose` | expose | (application) | — (verb-only, implicit object) |
| `find` | find | charm | — (verb-only, implicit object) |
| `find-offers` | find | offer | ✓ |
| `firewall-rules` | list | firewall-rule | ✓ (plural noun) |
| `grant` | grant | permission | — (verb-only, model scope implicit) |
| `grant-cloud` | grant | cloud-permission | ✓ |
| `grant-secret` | grant | secret-permission | ✓ |
| `help-action-commands` | help | action-command | ✓ |
| `help-hook-commands` | help | hook-command | ✓ |
| `import-filesystem` | import | filesystem | ✓ |
| `import-ssh-key` | import | ssh-key | ✓ |
| `info` | info | charm | — (verb-only, implicit object) |
| `integrate` | integrate | relation | — (verb-only) |
| `kill-controller` | kill | controller | ✓ |
| `login` | login | (session) | — (verb-only) |
| `logout` | logout | (session) | — (verb-only) |
| `machines` | list | machine | ✓ (plural noun) |
| `migrate` | migrate | model | — (verb-only, implicit object) |
| `model-config` | config | model | ✓ (noun-verb) |
| `model-constraints` | get | model-constraint | ✓ (noun-verb) |
| `model-defaults` | get/set | model-default | ✓ (noun-verb) |
| `model-secret-backend` | get/set | model-secret-backend | ✓ (noun-verb) |
| `models` | list | model | ✓ (plural noun) |
| `move-to-space` | move | subnet | ✓ (verb-phrase) |
| `offer` | offer | endpoint | — (verb-only, implicit object) |
| `offers` | list | offer | ✓ (plural noun) |
| `operations` | list | operation | ✓ (plural noun) |
| `refresh` | refresh | charm | — (verb-only, implicit object) |
| `regions` | list | region | ✓ (plural noun) |
| `register` | register | controller | — (verb-only, implicit object) |
| `reload-spaces` | reload | space | ✓ |
| `remove-application` | remove | application | ✓ |
| `remove-cloud` | remove | cloud | ✓ |
| `remove-credential` | remove | credential | ✓ |
| `remove-k8s` | remove | k8s (cloud) | ✓ |
| `remove-machine` | remove | machine | ✓ |
| `remove-offer` | remove | offer | ✓ |
| `remove-relation` | remove | relation | ✓ |
| `remove-saas` | remove | saas | ✓ |
| `remove-secret` | remove | secret | ✓ |
| `remove-secret-backend` | remove | secret-backend | ✓ |
| `remove-space` | remove | space | ✓ |
| `remove-ssh-key` | remove | ssh-key | ✓ |
| `remove-storage` | remove | storage | ✓ |
| `remove-storage-pool` | remove | storage-pool | ✓ |
| `remove-unit` | remove | unit | ✓ |
| `remove-user` | remove | user | ✓ |
| `rename-space` | rename | space | ✓ |
| `resolved` | resolve | (unit-error) | — (verb-only) |
| `resources` | list | resource | ✓ (plural noun) |
| `resume-relation` | resume | relation | ✓ |
| `retry-provisioning` | retry | provisioning | ✓ |
| `revoke` | revoke | permission | — (verb-only) |
| `revoke-cloud` | revoke | cloud-permission | ✓ |
| `revoke-secret` | revoke | secret-permission | ✓ |
| `run` | run | (command) | — (verb-only) |
| `scale-application` | scale | application | ✓ |
| `scp` | scp/copy | (file) | — (verb-only) |
| `secret-backends` | list | secret-backend | ✓ (plural noun) |
| `secrets` | list | secret | ✓ (plural noun) |
| `set-constraints` | set | constraint | — (application implicit) |
| `set-credential` | set | credential | — (model implicit) |
| `set-firewall-rule` | set | firewall-rule | ✓ |
| `set-model-constraints` | set | model-constraint | ✓ |
| `show-action` | show | action | ✓ |
| `show-application` | show | application | ✓ |
| `show-cloud` | show | cloud | ✓ |
| `show-controller` | show | controller | ✓ |
| `show-credential` | show | credential | ✓ |
| `show-machine` | show | machine | ✓ |
| `show-model` | show | model | ✓ |
| `show-offer` | show | offer | ✓ |
| `show-operation` | show | operation | ✓ |
| `show-secret` | show | secret | ✓ |
| `show-secret-backend` | show | secret-backend | ✓ |
| `show-space` | show | space | ✓ |
| `show-status-log` | show | status-log | ✓ |
| `show-storage` | show | storage | ✓ |
| `show-task` | show | task | ✓ |
| `show-unit` | show | unit | ✓ |
| `show-user` | show | user | ✓ |
| `spaces` | list | space | ✓ (plural noun) |
| `ssh` | ssh | (machine) | — (verb-only) |
| `ssh-keys` | list | ssh-key | ✓ (plural noun) |
| `status` | status | (all) | — (verb-only, state report) |
| `storage` | list | storage | ✓ (plural noun) |
| `storage-pools` | list | storage-pool | ✓ (plural noun) |
| `subnets` | list | subnet | ✓ (plural noun) |
| `suspend-relation` | suspend | relation | ✓ |
| `switch` | switch | (context) | — (verb-only) |
| `sync-agent-binary` | sync | agent-binary | ✓ |
| `trust` | trust | (application) | — (verb-only, implicit object) |
| `unexpose` | unexpose | (application) | ✓ (negation of expose) |
| `unregister` | unregister | controller | ✓ (negation of register) |
| `update-cloud` | update | cloud | ✓ |
| `update-credential` | update | credential | ✓ |
| `update-k8s` | update | k8s | ✓ |
| `update-public-clouds` | update | public-cloud | ✓ |
| `update-secret` | update | secret | ✓ |
| `update-secret-backend` | update | secret-backend | ✓ |
| `update-storage-pool` | update | storage-pool | ✓ |
| `upgrade-controller` | upgrade | controller | ✓ |
| `upgrade-model` | upgrade | model | ✓ |
| `users` | list | user | ✓ (plural noun) |
| `version` | version | (client) | — (verb-only, info) |
| `whoami` | whoami | (identity) | — (verb-only) |

## Verb-Noun Grid

Key nouns and their verb coverage:

| Noun | add | remove | update | show | list | set | Other |
|------|-----|--------|--------|------|------|-----|-------|
| cloud | ✓ | ✓ | ✓ | ✓ | ✓ (clouds) | — | default-region, default-credential, register, unregister |
| credential | ✓ | ✓ | ✓ | ✓ | ✓ (credentials) | — | autoload, default |
| controller | — | destroy | — | ✓ | ✓ | — | kill, enable-destroy, bootstrap, register, unregister, upgrade, config |
| model | ✓ | destroy | — | ✓ | ✓ (models) | — | migrate, dump, upgrade, config, defaults, constraints, export-bundle |
| application | — | ✓ (remove-app) | — | ✓ | — | — | deploy, config, expose, unexpose, refresh, trust, bind, scale, constraints, diff-bundle, resolved |
| unit | ✓ | ✓ | — | ✓ | — | — | resolved |
| machine | ✓ | ✓ | — | ✓ | ✓ | — | — |
| user | ✓ | ✓ | — | ✓ | ✓ | — | enable, disable, change-password, login, logout, whoami |
| space | ✓ | ✓ | rename | ✓ | ✓ | — | move-to, reload |
| storage | ✓ | ✓ | ✓ | ✓ | ✓ | — | attach, detach, import-filesystem |
| storage-pool | — | ✓ | ✓ | — | ✓ | — | create |
| secret | ✓ | ✓ | ✓ | ✓ | ✓ | — | grant, revoke |
| secret-backend | ✓ | ✓ | ✓ | ✓ | ✓ | — | model-secret-backend |
| relation | — | ✓ | — | — | — | — | integrate, suspend, resume |
| offer | — | ✓ | — | ✓ | ✓ | — | offer, consume, find-offers |
| ssh-key | ✓ | ✓ | — | — | ✓ | — | import |
| backup | — | — | — | — | — | — | create, download |
| charm | — | — | — | — | — | — | deploy, refresh, find, info, download |
| resource | — | — | — | — | ✓ | — | attach |
| action | — | — | — | ✓ | ✓ | — | run, exec, cancel-task |

## Annotations

### Incomplete CRUD Sets

Nouns missing expected lifecycle verbs:

| Noun | Missing | Notes |
|------|---------|-------|
| `controller` | `add-controller` | Created via `bootstrap` or `register`, no dedicated `add-controller` |
| `application` | `add-application` | Created via `deploy`, not `add-application` |
| `application` | `show-application` exists; `list` is `status` | No standalone `applications` list command |
| `backup` | `remove-backup`, `list-backups`, `show-backup` | Only `create` and `download` |
| `relation` | `add-relation` | Named `integrate` (with `relate` alias), not `add-relation` |
| `relation` | `show-relation`, `list-relations` | No observation commands |
| `offer` | `update-offer` | No update after creation |
| `machine` | `update-machine` | No update command |
| `space` | `update-space` | Only `rename-space`, no general update |
| `storage` | `update-storage` | No standalone update (pool has update) |
| `region` | `add-region`, `remove-region`, `update-region` | Only `default-region` and `regions` (list) |
| `saas` | `add-saas` | Created via `consume`, not `add-saas` |
| `status-log` | Only `show-status-log` | No `list-status-logs` |
| `subnet` | Only `subnets` (list) | No add/remove/update |
| `task` | Only `show-task`, `cancel-task` | Operations provide the list |
| `firewall-rule` | `remove-firewall-rule` | Only `set-firewall-rule` and `firewall-rules` (list) |
| `agent-binary` | Only `sync-agent-binary` | No `list-agent-binaries` |

### Verb Inconsistencies

| Issue | Commands | Recommendation |
|-------|----------|----------------|
| `destroy` vs `remove` for controllers/models | `destroy-controller`, `destroy-model` vs `remove-application`, `remove-machine` | Models and controllers use `destroy`; applications and machines use `remove`. These are semantically different: `destroy` implies total deletion of the container; `remove` implies removal of the entity. However, `remove-application` destroys the application completely — same as `destroy`. Inconsistent. |
| `kill` vs `destroy` for controllers | `kill-controller` vs `destroy-controller` | `kill` is a forced `destroy`. FrameNet: `kill` (Killing) implies the target was alive; `destroy` (Destroying) operates on objects. |
| `deploy` vs `add-application` | `deploy <charm>` vs... nothing | `deploy` does create an application, but the verb reflects the charm-deployment action rather than the resource creation. |
| `integrate` vs `add-relation` | `integrate` (alias `relate`) | Per DE013 §Grammar: "Commands are verbs." `integrate` is a pure verb without noun. The alias `relate` is even less descriptive. |
| `consume` vs `add-saas` | `consume` creates a SAAS | The verb `consume` is metaphorical (consume an offer → create local reference). Not obvious to new users. |
| `config` (app) vs `model-config` vs `controller-config` | Three config commands | While logically distinct, the naming pattern flips: application uses bare `config`, model/controller use noun-verb. |
| `default-credential` vs `set-default-credential` | Command named `default-credential` but conceptually sets a default | Per DE013: should be `set-default-credential` or just accept `default` as a verb. |

### Orphan Commands

Commands that do not decompose cleanly into verb-noun:

| Command | Issue |
|---------|-------|
| `bootstrap` | Self-initializing operation. Implicitly creates a controller. No reverse operation. |
| `integrate` | Verb-only with bidirectional semantics. Creates (not adds) a relation. |
| `resolved` | Adjective used as a command. Marks unit errors as resolved. |
| `whoami` | UNIX convention verb used as introspection command. |
| `switch` | Changes model/controller context. No noun target. |
| `migrate` | Moves a model between controllers. The noun (model) is implicit. |
| `dashboard` | Opens a browser. Not a resource operation. |
| `version` | Informational. No noun. |
| `status` | Global state report. No noun. |
| `ssh` / `scp` | Transport verbs. Target is a machine, but implicit. |
| `run` / `exec` | Execution verbs. Target is a unit, but implicit via positional arg. |
| `trust` | Grants credential access. Verb-only. |
| `bind` | Modifies application endpoint-space bindings. No clear noun. |
| `diff-bundle` | Comparison/analysis operation. |
| `export-bundle` | Export operation on the entire model. |
| `help-action-commands` / `help-hook-commands` | Meta/help commands. |
| `sync-agent-binary` | Internal tooling operation. |
| `enable-destroy-controller` | Meta-command that enables another operation. |
| `retry-provisioning` | Operational recovery. Compound verb-phrase. |
| `debug-log` | Streaming diagnostic output. |
| `autoload-credentials` | Compound verb. Unusual prefix. |
| `move-to-space` | Verb-phrase instead of verb-noun (`move subnet` would be more standard). |
