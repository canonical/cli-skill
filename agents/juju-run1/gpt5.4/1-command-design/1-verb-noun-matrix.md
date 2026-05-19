# Juju Verb-Noun Matrix

## Reading the matrix

This matrix groups the current surface by dominant resource noun. Cells contain representative built-in commands. Empty cells indicate either a genuine gap or a place where Juju uses a different lexical strategy.

Columns:
- `app`: application / unit / relation surface
- `model`: model surface
- `controller`: controller surface
- `cloud`: cloud / region / k8s / credential surface
- `machine`: machine surface
- `storage`: storage / pool / filesystem surface
- `network`: space / subnet / firewall surface
- `identity`: user / access / ssh-key surface
- `secret`: secret / secret-backend surface
- `offer`: cross-model offer / saas surface
- `charm`: charm / resource surface
- `meta`: diagnostic or host-level commands

## Matrix

| Verb | app | model | controller | cloud | machine | storage | network | identity | secret | offer | charm | meta |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `add` | `add-unit` | `add-model` |  | `add-cloud`, `add-credential`, `add-k8s` | `add-machine` | `add-storage` | `add-space` | `add-user`, `add-ssh-key` | `add-secret`, `add-secret-backend` |  |  |  |
| `attach` | `attach-resource` |  |  |  |  | `attach-storage` |  |  |  |  | `attach-resource` |  |
| `autoload` |  |  |  | `autoload-credentials` |  |  |  |  |  |  |  |  |
| `bind` | `bind` |  |  |  |  |  |  |  |  |  |  |  |
| `bootstrap` |  |  | `bootstrap` |  |  |  |  |  |  |  |  |  |
| `cancel` |  |  |  |  |  |  |  |  |  |  |  | `cancel-task` |
| `change` |  |  |  |  |  |  |  | `change-user-password` |  |  |  |  |
| `config` | `config` | `model-config`, `model-defaults`, `model-secret-backend` | `controller-config` |  |  |  |  |  |  |  |  |  |
| `consume` |  |  |  |  |  |  |  |  |  | `consume` |  |  |
| `create` |  |  |  |  |  | `create-storage-pool` |  |  |  |  |  | `create-backup` |
| `debug` |  |  |  |  |  |  |  |  |  |  |  | `debug-code`, `debug-hooks`, `debug-log` |
| `default` |  |  |  | `default-region`, `default-credential` |  |  |  |  |  |  |  |  |
| `deploy` | `deploy` |  |  |  |  |  |  |  |  |  |  |  |
| `destroy` | `destroy-model` affects app resources | `destroy-model` | `destroy-controller` |  |  | destroy/release storage as flags |  |  |  |  |  |  |
| `detach` |  |  |  |  |  | `detach-storage` |  |  |  |  |  |  |
| `diff` | `diff-bundle` |  |  |  |  |  |  |  |  |  |  |  |
| `disable` |  |  |  |  |  |  |  | `disable-user` |  |  |  | `disable-command` |
| `download` |  |  |  |  |  |  |  |  |  |  | `download`, `download-backup` |  |
| `enable` |  |  | `enable-destroy-controller` |  |  |  |  | `enable-user` |  |  |  | `enable-command` |
| `exec` | app-targeted remote execution |  |  |  | machine-targeted transitively |  |  |  |  |  |  | `exec` |
| `export` |  | `export-bundle` |  |  |  |  |  |  |  |  |  |  |
| `expose` | `expose`, `unexpose` |  |  |  |  |  | network side effects |  |  |  |  |  |
| `find` |  |  |  |  |  |  |  |  |  | `find-offers` | `find`, `info` |  |
| `grant` |  | model access through `grant` | controller access through `grant` | `grant-cloud` |  |  |  | user/access target | `grant-secret` | offer access through `grant` |  |  |
| `help` |  |  |  |  |  |  |  |  |  |  |  | `help`, `help-action-commands`, `help-hook-commands` |
| `import` |  |  |  |  |  | `import-filesystem` |  | `import-ssh-key` |  |  |  |  |
| `integrate` | `integrate`, `remove-relation`, `suspend-relation`, `resume-relation` |  |  |  |  |  |  |  |  |  |  |  |
| `kill` |  |  | `kill-controller` |  |  |  |  |  |  |  |  |  |
| `list` / plural noun |  | `models` | `controllers` | `clouds`, `regions`, `credentials` | `machines` | `storage`, `storage-pools` | `spaces`, `subnets`, `firewall-rules` | `users`, `ssh-keys` | `secrets`, `secret-backends` | `offers` | `resources`, `charm-resources` | `actions`, `operations`, `disabled-commands` |
| `login/logout` |  |  | controller session focus |  |  |  |  | `login`, `logout`, `whoami` |  |  |  |  |
| `migrate` |  | `migrate` |  |  |  |  |  |  |  |  |  |  |
| `move` |  |  |  |  |  |  | `move-to-space` |  |  |  |  |  |
| `offer` |  |  |  |  |  |  |  |  |  | `offer`, `remove-offer`, `show-offer` |  |  |
| `refresh` | `refresh` |  |  |  |  |  |  |  |  |  | charm refresh semantics |  |
| `register/unregister` |  |  | `register`, `unregister` |  |  |  |  |  |  |  |  |  |
| `reload` |  |  |  |  |  |  | `reload-spaces` |  |  |  |  |  |
| `remove` | `remove-application`, `remove-unit`, `remove-relation`, `remove-saas` |  |  | `remove-cloud`, `remove-credential`, `remove-k8s` | `remove-machine` | `remove-storage`, `remove-storage-pool` | `remove-space` | `remove-user`, `remove-ssh-key` | `remove-secret`, `remove-secret-backend` | `remove-offer` |  |  |
| `rename` |  |  |  |  |  |  | `rename-space` |  |  |  |  |  |
| `resolved` | `resolved` |  |  |  |  |  |  |  |  |  |  |  |
| `resume/suspend` | relation lifecycle |  |  |  |  |  |  |  |  | offer relation lifecycle |  |  |
| `retry` |  | `retry-provisioning` |  |  | machine provisioning indirectly |  |  |  |  |  |  |  |
| `revoke` |  | model access through `revoke` | controller access through `revoke` | `revoke-cloud` |  |  |  | user/access target | `revoke-secret` | offer access through `revoke` |  |  |
| `run` | `run` action |  |  |  |  |  |  |  |  |  |  | `run` |
| `scale` | `scale-application` |  |  |  |  |  |  |  |  |  |  |  |
| `scp/ssh` | app/unit/machine access |  |  |  | `ssh`, `scp` |  |  |  |  |  |  |  |
| `set` | `set-constraints` | `set-model-constraints`, `set-credential` |  |  |  | `set-firewall-rule` partly network | `set-firewall-rule` |  |  |  |  |  |
| `show` | `show-application`, `show-unit`, `show-action`, `show-operation`, `show-task` | `show-model` | `show-controller` | `show-cloud`, `show-credential` | `show-machine` | `show-storage` | `show-space`, `show-status-log` | `show-user` | `show-secret`, `show-secret-backend` | `show-offer` |  |  |
| `status` | application/unit/model status | model-wide | controller indirectly via `show-controller` not `status` |  | machine status via model report | storage section in status | relation/network sections |  |  |  |  | `status` |
| `switch` |  | model/controller focus | controller/model focus |  |  |  |  |  |  |  |  | `switch` |
| `sync` |  |  | controller agent maintenance |  |  |  |  |  |  |  |  | `sync-agent-binary` |
| `trust` | `trust` |  |  |  |  |  |  |  |  |  |  |  |
| `update` | `update-secret` |  |  | `update-cloud`, `update-credential`, `update-k8s`, `update-public-clouds` |  | `update-storage-pool` |  |  | `update-secret-backend` |  |  |  |
| `upgrade` | application refresh-adjacent but separate | `upgrade-model` | `upgrade-controller` |  |  |  |  |  |  |  |  |  |

## Observations from the matrix

- The densest nouns are `application`, `model`, `controller`, `cloud`, `storage`, and `identity`.
- `offer` and `secret` families are relatively coherent despite being smaller.
- `network` operations are present but lexically uneven: `spaces`, `subnets`, `move-to-space`, `reload-spaces`, `set-firewall-rule`.
- Several commands are meta-operations with no stable noun slot, which is acceptable but increases namespace pressure.

## Empty slots that stand out

- No `show-storage-pool`; only `storage-pools` plus create/update/remove.
- No `show-subnet`; only plural listing.
- No `show-firewall-rule`; only `firewall-rules` and `set-firewall-rule`.
- No explicit `untrust`; trust removal is folded into `trust` flags rather than a complement command.
- No explicit `validate-bundle` or `plan-deploy`; deploy only has `--dry-run`.
- No `rename-model`, `rename-controller`, or `rename-cloud` class of operations.

## Orphan or weakly attached commands

These do not fit cleanly into the verb-noun grid and should be treated as special-surface commands:
- `status`
- `switch`
- `dashboard`
- `documentation`
- `version`
- `help-action-commands`
- `help-hook-commands`
- `debug-log`
- `exec`
- `run`
- `ssh`
- `scp`

These commands are not necessarily bad. They just depend more on documentation because the naming system alone does less explanatory work.
