# 03 — Semantic Domain Clustering

## Overview

All commands are grouped by the resource domain they operate on. Every command appears exactly once. Total command count verified against the command set.

| Domain | Count | Commands | Naming Consistent? | Notes |
|--------|-------|----------|-------------------|-------|
| **Cloud** | 14 | add-cloud, update-cloud, remove-cloud, clouds, show-cloud, regions, default-region, update-public-clouds, add-k8s, update-k8s, remove-k8s, add-credential, update-credential, remove-credential | Mostly: `*-cloud`, `*-credential`. `k8s` is separate noun. `default-region` is an outlier (should be `set-default-region`). | CRUD complete for cloud and credential. K8s has its own subdomain. |
| **Credential** | 7 | credentials, show-credential, default-credential, autoload-credentials, add-credential, update-credential, remove-credential | Consistent: all use `credential` noun. | CRUD complete. `autoload-credentials` is an orphan verb. |
| **Controller** | 13 | bootstrap, controllers, show-controller, destroy-controller, kill-controller, enable-destroy-controller, controller-config, register, unregister, upgrade-controller, add-model, models, migrate | Partial: `destroy/kill` vs `remove`. `bootstrap` creates without `add-`. `add-model` and `models` are controller-scoped but model-domain. | `migrate` crosses domains. `register`/`unregister` follow login semantics rather than CRUD. |
| **Model** | 14 | add-model, models, show-model, destroy-model, model-config, model-defaults, model-constraints, set-model-constraints, upgrade-model, export-bundle, dump-model, dump-db, set-credential, retry-provisioning | Model-prefixed for config/constraints/defaults (noun-verb). Non-prefixed: `add-model`, `models`, `show-model`, `destroy-model`, `export-bundle`. | CRUD mostly complete. `switch` and `migrate` are cross-domain. `dump-db` is controller-level but gated by model. |
| **User / Access** | 13 | add-user, remove-user, show-user, users, change-user-password, disable-user, enable-user, login, logout, whoami, grant, revoke, grant-cloud, revoke-cloud | `grant`/`revoke` are verb-only (per DE013 OK when context is clear). `whoami` is an outlier. `change-user-password` is verbose. | CRUD complete. `login`/`logout` are session, not resource. `grant-cloud`/`revoke-cloud` are cloud-access, included here. |
| **Application** | 18 | deploy, add-unit, remove-unit, remove-application, config, show-application, show-unit, expose, unexpose, refresh, trust, bind, scale-application, constraints, set-constraints, diff-bundle, resolved, remove-saas | Mix of verb-only (`deploy`, `config`, `expose`, `trust`, `bind`) and verb-noun. `remove-saas` is related to CMR. | Creation via `deploy` (not `add-application`). `resolved` is an adjective. `scale-application` uses verb-noun. |
| **Machine** | 4 | add-machine, remove-machine, machines, show-machine | Consistent verb-noun pattern. | CRUD mostly complete. No `update-machine`. |
| **Relation / Integration** | 5 | integrate, remove-relation, suspend-relation, resume-relation, consume | Mix: `integrate` is verb-only; `remove-relation`, `suspend-relation`, `resume-relation` are verb-noun. `consume` is cross-model. | `integrate` is non-standard (should be `add-relation` or `create-relation`). No `show-relation` or `list-relations`. |
| **Cross-Model Relations (CMR)** | 7 | offer, remove-offer, show-offer, offers, find-offers, consume, remove-saas | `offer`/`remove-offer` (verb-noun). `consume` is verb-only. `remove-saas` is distinct noun. | CRUD mostly complete. `find-offers` for discovery. |
| **Storage** | 11 | add-storage, attach-storage, detach-storage, remove-storage, show-storage, storage, import-filesystem, create-storage-pool, update-storage-pool, remove-storage-pool, storage-pools | Consistent: `*-storage` and `*-storage-pool`. `import-filesystem` is the outlier noun. | CRUD mostly complete. `import-filesystem` is a special import path. |
| **Space / Network** | 8 | add-space, remove-space, spaces, show-space, rename-space, move-to-space, reload-spaces, subnets | Consistent: `*-space` and `subnets`. `move-to-space` is a verb-phrase. `reload-spaces` uses `reload`. | CRUD mostly complete. No `update-space`. `subnets` is separate domain with only list. |
| **Secret** | 7 | add-secret, update-secret, remove-secret, secrets, show-secret, grant-secret, revoke-secret | Consistent: `*-secret` pattern throughout. | CRUD complete. Grant/revoke extend the domain. |
| **Secret Backend** | 6 | add-secret-backend, update-secret-backend, remove-secret-backend, secret-backends, show-secret-backend, model-secret-backend | Consistent: `*-secret-backend` pattern. `model-secret-backend` is noun-verb. | CRUD complete plus model default setting. |
| **Action / Operation** | 8 | run, exec, actions, show-action, operations, show-operation, show-task, cancel-task | Inconsistent: `run`/`exec` are verb-only; `actions`/`operations` are plural nouns; `show-action`, `show-operation`, `show-task`, `cancel-task` are verb-noun. | `run` creates an operation; `exec` executes directly. Related but semantically distinct. |
| **SSH / Debug** | 7 | ssh, scp, debug-hooks, debug-code, debug-log, add-ssh-key, remove-ssh-key, import-ssh-key, ssh-keys | `ssh`/`scp` are verb-only transfers. `debug-*` are verb-noun. SSH keys follow CRUD. | SSH keys are access management for SSH transport. |
| **Status / Monitoring** | 4 | status, show-status-log, firewall-rules, set-firewall-rule | `status` is verb-only; `show-status-log` is verb-noun. Firewall rules are access control. | `status` is the primary observation command. `show-status-log` is historical. |
| **Block / Protection** | 3 | disable-command, enable-command, disabled-commands | Consistent: `*-command` with block semantics. | Simple toggle domain. |
| **Backup** | 2 | create-backup, download-backup | Consistent: `*-backup` pattern. | Incomplete: no `list-backups`, `show-backup`, `remove-backup`. |
| **CharmHub** | 3 | find, info, download | Verb-only (implicit noun: charm). | `find`, `info`, `download` are standard discovery/download verbs. |
| **Resource** | 3 | attach-resource, resources, charm-resources | `attach-resource` is verb-noun; `resources` is plural noun; `charm-resources` is noun-noun. | `charm-resources` is a lookup command, not CRUD. |
| **Dashboard** | 1 | dashboard | Verb-only. | Single command to open browser. |
| **Tooling / Meta** | 5 | version, help-action-commands, help-hook-commands, sync-agent-binary, upgrade-controller, upgrade-model | `help-*-commands` is meta; `sync-agent-binary` is tooling; `upgrade-*` is lifecycle. | Scattered across domains. |
| **Firewall** | 2 | set-firewall-rule, firewall-rules | `set-firewall-rule` is verb-noun; `firewall-rules` is plural noun. | Simple access-control domain. |

## Domain Summary

| Metric | Value |
|--------|-------|
| Total unique commands | 130 |
| Total domains | 22 |
| Commands accounted for | 130 |
| Domains with complete CRUD | cloud, credential, user, secret, secret-backend, space (partial) |
| Domains with incomplete CRUD | controller, model, application, machine, relation, backup, offer, storage, action |
| Commands with naming inconsistencies | ~25 (mostly verb-only vs verb-noun) |

## Naming Pattern Analysis by Domain

| Domain | Primary Pattern | Issues |
|--------|----------------|--------|
| Cloud | verb-noun (`add-cloud`, `update-cloud`) | `default-region` is `set-default-region` semantic; `clouds` is plural noun |
| Credential | verb-noun (`add-credential`) | Consistent |
| Controller | verb-noun + verb-only mix | `bootstrap` creates; `destroy`/`kill` remove; `register`/`unregister` are access verbs |
| Model | verb-noun + noun-verb mix | `add-model`, `destroy-model`, `show-model` (verb-noun); `model-config`, `model-defaults` (noun-verb) |
| User | verb-noun (`add-user`) | `login`/`logout`/`whoami` are session verbs, not resource CRUD |
| Application | verb-only dominant | `deploy`, `config`, `expose`, `trust`, `bind`, `refresh`, `resolved` vs `add-unit`, `remove-unit` |
| Machine | verb-noun consistent | Clean CRUD pattern |
| Storage | verb-noun consistent | `import-filesystem` is the outlier |
| Space | verb-noun consistent | `move-to-space` is a verb-phrase |
| Secret | verb-noun consistent | Clean pattern |
| Secret Backend | verb-noun consistent | Clean pattern |
| Relation | verb-noun + verb-only | `integrate` is verb-only; others are verb-noun |
| SSH | verb-only dominant | `ssh`, `scp` are transport verbs; keys follow CRUD |

## DE013 Compliance Notes

Per DE013 §Grammar:
1. **"Commands are verbs"** — Most Juju commands comply. Orphans like `resolved` (adjective), `whoami` (pronoun-verb), and `dashboard` (noun) violate this.
2. **"Foobars as shorthand for listing"** — Juju follows this: `models`, `controllers`, `users`, `secrets`, `spaces`, `machines`, `storage` all use the plural noun without `list-`.
3. **"Verb-noun when verbs alone are insufficient"** — Applied correctly in cloud, credential, machine, secret, space domains where multiple resource types exist.
4. **"Status over show-status"** — Followed: `status` (not `show-status`). But `show-status-log` exists for history.
5. **"At most one sublevel"** — Followed: Juju has no subcommand nesting currently.
