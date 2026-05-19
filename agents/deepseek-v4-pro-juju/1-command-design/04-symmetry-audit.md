# 04 — Symmetry Audit

## Overview

Every pair of symmetric operations (forward and reverse) is listed here, including missing reverse operations.

## Exhaustive Symmetry Table

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|-----------|----------------|-----------------|-------------------|---------------------|-------|
| Cloud CRUD | `add-cloud` | `remove-cloud` | ✅ Yes | ✅ Yes | Both local operations with controller awareness |
| K8s Cloud CRUD | `add-k8s` | `remove-k8s` | ✅ Yes | ⚠️ No | `remove-k8s` supports `--force`; `add-k8s` has more flags |
| Credential CRUD | `add-credential` | `remove-credential` | ✅ Yes | ⚠️ No | `add-credential` has `-f`/`--file`; `remove-credential` has no file option |
| Credential Update | `update-credential` | — | — | — | No `revert-credential` or `restore-credential` |
| Model CRUD | `add-model` | `destroy-model` | ❌ No | ❌ No | Create is `add-model`; destroy is `destroy-model` (not `remove-model`). Different verb families. `destroy` implies total deletion of container. |
| Controller CRUD | `bootstrap` | `destroy-controller` / `kill-controller` | ❌ No | ❌ No | No `add-controller`. Creation is `bootstrap` or `register`. Destruction has two levels (`destroy`/`kill`). |
| Machine CRUD | `add-machine` | `remove-machine` | ✅ Yes | ⚠️ No | `remove-machine` has `--force`, `--no-wait`, `--keep-instance`; `add-machine` lacks `--force` |
| Unit CRUD | `add-unit` | `remove-unit` | ✅ Yes | ⚠️ No | `remove-unit` has `--force`, `--dry-run`, `--no-wait`, `--destroy-storage`; `add-unit` has none of these |
| Application CRUD | `deploy` | `remove-application` | ❌ No | ❌ No | Create is `deploy` (verb-only, charm-centric); remove is `remove-application` (verb-noun). Asymmetric. |
| User CRUD | `add-user` | `remove-user` | ✅ Yes | ⚠️ No | `remove-user` has `--yes`, `--force`; `add-user` has neither |
| User Access | `disable-user` | `enable-user` | ✅ Yes | ✅ Yes | Clean toggle |
| User Password | `change-user-password` | — | — | — | No `reset-user-password` or `revert-user-password` |
| Space CRUD | `add-space` | `remove-space` | ✅ Yes | ✅ Yes | Clean pair |
| Space Rename | `rename-space` | — | — | — | No `restore-space-name` |
| Storage CRUD | `add-storage` | `remove-storage` | ✅ Yes | ⚠️ No | `remove-storage` has `--force`, `--no-wait`, `--destroy-storage` |
| Storage Pool CRUD | `create-storage-pool` | `remove-storage-pool` | ❌ No | ⚠️ No | Create is `create-storage-pool` (not `add-storage-pool`). Different create verb. |
| Storage Attach | `attach-storage` | `detach-storage` | ✅ Yes | ⚠️ No | `detach-storage` has `--force` |
| File System Import | `import-filesystem` | — | — | — | No `export-filesystem` |
| Secret CRUD | `add-secret` | `remove-secret` | ✅ Yes | ✅ Yes | Clean pair |
| Secret Update | `update-secret` | — | — | — | No revert |
| Secret Permission | `grant-secret` | `revoke-secret` | ✅ Yes | ✅ Yes | Clean pair |
| Secret Backend CRUD | `add-secret-backend` | `remove-secret-backend` | ✅ Yes | ✅ Yes | Clean pair |
| Secret Backend Update | `update-secret-backend` | — | — | — | No revert |
| Model Permission | `grant` | `revoke` | ✅ Yes | ✅ Yes | Clean verb pair |
| Cloud Permission | `grant-cloud` | `revoke-cloud` | ✅ Yes | ✅ Yes | Clean pair |
| Command Block | `disable-command` | `enable-command` | ✅ Yes | ✅ Yes | Clean pair |
| Firewall Rule | `set-firewall-rule` | — | — | — | No `remove-firewall-rule` or `unset-firewall-rule` |
| Expose | `expose` | `unexpose` | ✅ Yes | ✅ Yes | Clean negation prefix pair |
| Trust | `trust` | — | — | — | `trust --remove` acts as reverse, but no `untrust` command |
| Register Controller | `register` | `unregister` | ✅ Yes | ✅ Yes | Clean negation prefix pair |
| Login | `login` | `logout` | ✅ Yes | ✅ Yes | Clean session pair |
| Offer | `offer` | `remove-offer` | ❌ No | ⚠️ No | Create is `offer` (verb-only); remove is `remove-offer` (verb-noun) |
| Consume SAAS | `consume` | `remove-saas` | ❌ No | ❌ No | Create is `consume` (verb-only); remove is `remove-saas` (noun mismatch). `consume` → `remove-saas` is asymmetric in verbs AND nouns. |
| Relation | `integrate` | `remove-relation` | ❌ No | ⚠️ No | Create is `integrate` (verb-only); remove is `remove-relation` (verb-noun). `relate` alias vs `remove-relation` also asymmetric. |
| Relation Suspend | `suspend-relation` | `resume-relation` | ✅ Yes | ✅ Yes | Clean pair |
| Backup | `create-backup` | — | — | — | No `remove-backup`. `download-backup` is read, not delete. |
| Upgrade | `upgrade-model` | — | — | — | No `downgrade-model`; not reversible |
| Upgrade Controller | `upgrade-controller` | — | — | — | No `downgrade-controller`; not reversible |
| Model Migrate | `migrate` | — | — | — | No `unmigrate`; migration is one-way |
| SSH Key | `add-ssh-key` | `remove-ssh-key` | ✅ Yes | ✅ Yes | Clean pair |
| SSH Key Import | `import-ssh-key` | — | — | — | No `export-ssh-key` |
| Bootstrap | `bootstrap` | — | — | — | No `debootstrap`; irrecoverable (controller must be destroyed) |
| Bundle Export | `export-bundle` | — | — | — | No `import-bundle`; `deploy` with a bundle file creates applications but is not the reverse |
| Resource Attach | `attach-resource` | — | — | — | No `detach-resource` or `remove-resource` standalone; `remove-application` cleans up |
| Scale | `scale-application` | — | — | — | Scaling is not a paired operation (can scale up or down, no `unscale`) |
| Refresh | `refresh` | — | — | — | No `revert` to previous charm version |
| Resolved | `resolved` | — | — | — | No `unresolve`; state is final once resolved |
| Autoload Credentials | `autoload-credentials` | — | — | — | No `unload-credentials` |
| Public Clouds Update | `update-public-clouds` | — | — | — | No `revert-public-clouds` |
| Sync Agent Binary | `sync-agent-binary` | — | — | — | No `unsync` |
| Retry Provisioning | `retry-provisioning` | — | — | — | No undo |
| Debug | `debug-log`, `debug-hooks`, `debug-code` | — | — | — | Debug commands are non-destructive observation |
| Run / Exec | `run`, `exec` | `cancel-task` | — | — | `cancel-task` can cancel an operation created by `run`, but not direct `exec` |
| Model Dump | `dump-model`, `dump-db` | — | — | — | Read-only observation |
| Status / Show | `status`, `show-*`, `list-*` (plurals) | — | — | — | Read-only observation, no reverse needed |
| Switch | `switch` | — | — | — | No reverse; switching to another model is the "undo" |
| Dashboard | `dashboard` | — | — | — | Browser-based, no CLI reverse |

## Missing Reverse Operations (Gaps)

| Missing Reverse | Forward | Recommended Name | Priority |
|----------------|---------|-----------------|----------|
| Remove firewall rule | `set-firewall-rule` | `remove-firewall-rule` | High |
| Remove trust | `trust` | `untrust` | Medium |
| Remove backup | `create-backup` | `remove-backup` | Medium |
| Detach resource | `attach-resource` | `detach-resource` | Medium |
| Unresolve | `resolved` | `unresolve` or `reopen-unit` | Low |
| Export SSH key | `import-ssh-key` | `export-ssh-key` | Low |

## Naming Asymmetries

| Asymmetry | Details |
|-----------|---------|
| `add-model` / `destroy-model` | Different verb families. `destroy` implies permanent, container-level deletion. `remove` would be misleading for models since models are containers. This is arguably correct. |
| `bootstrap` / `destroy-controller` | `bootstrap` has no verb-noun structure. Implies self-initializing creation. No `add-controller`. |
| `deploy` / `remove-application` | The fundamental creation verb for applications is `deploy`, not `add-application`. This is semantically rich but asymmetric. |
| `offer` / `remove-offer` | `offer` is verb-only; `remove-offer` uses the noun. Should be `create-offer` / `remove-offer` or `offer` / `unoffer`. |
| `consume` / `remove-saas` | Completely different verbs and nouns. `consume` → `remove-saas` is the most asymmetric pair. |
| `integrate` / `remove-relation` | `integrate` (or `relate`) → `remove-relation`. Different verb families, different noun presence. |
| `create-backup` / (none) | No `remove-backup`. Inconsistent with most domains that use `add`/`remove`. |
| `create-storage-pool` / `remove-storage-pool` | `create` instead of `add` for storage pools. Inconsistent with `add-storage`. |

## Behavioral Asymmetries

| Asymmetry | Details |
|-----------|---------|
| `remove-*` commands have `--force`, `--dry-run`, `--no-wait` | Forward `add-*` commands generally lack these. Destructive operations have richer safety controls. |
| `destroy-*` requires `--force` for teardown | `add-model` and `bootstrap` have no force variant |
| `kill-controller` is implicitly forced | `bootstrap` has no forced variant |
| `remove-user` has `--yes` confirmation | `add-user` does not (it's non-destructive) |
| `unregister` is local-only | `register` contacts the controller |
| `logout` is local | `login` contacts the controller |

## Summary

- **Clean symmetric pairs**: 20 (`add/remove-*` for cloud, k8s, credential, machine, unit, user, space, storage, secret, secret-backend, ssh-key; `enable/disable-*` for user and command; `expose/unexpose`; `grant/revoke` and variants; `login/logout`; `register/unregister`; `suspend/resume-relation`; `attach/detach-storage`)
- **Asymmetric pairs**: 8 (`deploy/remove-application`, `bootstrap/destroy-controller`, `add-model/destroy-model`, `offer/remove-offer`, `consume/remove-saas`, `integrate/remove-relation`, `create-storage-pool/remove-storage-pool`, `create-backup`/none)
- **Missing reverses**: 4 high-priority gaps
