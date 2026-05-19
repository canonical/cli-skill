## Section 3: Semantic Domain Clustering

| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---|---|---|---|
| cloud | 11 | `add-cloud`, `remove-cloud`, `update-cloud`, `show-cloud`, `clouds`, `regions`, `add-k8s`, `remove-k8s`, `update-k8s`, `update-public-clouds`, `autoload-credentials` | Mixed | Contains k8s sub-domain; `update-public-clouds` is an orphan; `regions` is list-only |
| controller | 12 | `bootstrap`, `destroy-controller`, `kill-controller`, `show-controller`, `controllers`, `controller-config`, `upgrade-controller`, `register`, `unregister`, `login`, `logout`, `whoami`, `dashboard`, `enable-destroy-controller` | Mixed | `bootstrap` is create verb; `kill` is exceptional; `enable-destroy-controller` is unique |
| model | 13 | `add-model`, `destroy-model`, `show-model`, `models`, `migrate`, `model-config`, `model-constraints`, `model-defaults`, `model-secret-backend`, `upgrade-model`, `switch`, `status`, `set-model-constraints` | Mixed | `destroy-model` vs `remove-*` elsewhere; `status` is a model-scoped orphan |
| application | 16 | `deploy`, `remove-application`, `show-application`, `refresh`, `scale-application`, `actions`, `config`, `constraints`, `resources`, `expose`, `unexpose`, `trust`, `charm-resources`, `attach-resource`, `download`, `info`, `find` | Mixed | `deploy` is the primary create verb (no `add-application`); multiple noun orphans |
| unit | 13 | `add-unit`, `remove-unit`, `show-unit`, `exec`, `run`, `debug-code`, `debug-hooks`, `ssh`, `scp`, `add-storage`, `attach-storage`, `detach-storage`, `resolved` | Mixed | `resolved` is participle-only; `exec`/`run`/`ssh` overlap in execution |
| machine | 6 | `add-machine`, `remove-machine`, `show-machine`, `machines`, `retry-provisioning` | Yes | Clean `add/remove/show` + list; `retry-provisioning` is an outlier |
| credential | 9 | `add-credential`, `remove-credential`, `show-credential`, `update-credential`, `default-credential`, `set-credential`, `credentials`, `autoload-credentials` | Mixed | `default-credential` uses adjective-as-verb; `autoload-credentials` is orphan verb + plural |
| secret | 10 | `add-secret`, `remove-secret`, `show-secret`, `update-secret`, `grant-secret`, `revoke-secret`, `secrets`, `add-secret-backend`, `remove-secret-backend`, `show-secret-backend`, `update-secret-backend`, `secret-backends`, `model-secret-backend` | Yes | Strong `add/remove/show/update` + backend sub-domain |
| space/network | 10 | `add-space`, `remove-space`, `show-space`, `rename-space`, `move-to-space`, `spaces`, `reload-spaces`, `subnets`, `bind` | Mixed | `reload-spaces` is a bare verb; `subnets` is list-only; `bind` is an orphan |
| storage | 10 | `add-storage`, `remove-storage`, `show-storage`, `storage`, `attach-storage`, `detach-storage`, `create-storage-pool`, `remove-storage-pool`, `update-storage-pool`, `storage-pools` | Mixed | `create-*` for pool vs `add-*` for storage instance; `storage` is noun-only list |
| offer/CMR | 10 | `offer`, `remove-offer`, `show-offer`, `offers`, `find-offers`, `consume`, `remove-saas`, `grant`, `revoke`, `integrate` | Mixed | `consume`/`remove-saas` are asymmetric; `integrate` is relation-level |
| relation | 6 | `integrate`, `remove-relation`, `suspend-relation`, `resume-relation` | Mixed | No `add-relation` (integrate used); `suspend/resume` are well paired |
| user | 9 | `add-user`, `remove-user`, `show-user`, `disable-user`, `enable-user`, `change-user-password`, `users` | Yes | Clean lifecycle; `change-user-password` is the only compound verb |
| charm/resource | 8 | `download`, `info`, `find`, `charm-resources`, `resources`, `attach-resource`, `deploy` | Mixed | `deploy` pulls in from application domain; `info`/`find` are orphan observation verbs |
| bundle | 3 | `deploy`, `diff-bundle`, `export-bundle` | Mixed | No `show-bundle` or `list-bundles`; `deploy` is shared with app domain |
| backup | 2 | `create-backup`, `download-backup` | Mixed | Missing show/list/remove; `create` used instead of `add` |
| task/operation | 5 | `cancel-task`, `show-task`, `operations`, `show-operation`, `run` | Mixed | No `list-tasks`; `run` overlaps with action execution |
| firewall | 2 | `set-firewall-rule`, `firewall-rules` | Partial | `set` is used; no `remove-firewall-rule` |
| config | 5 | `config`, `controller-config`, `model-config`, `model-defaults` | Mixed | Noun-first hybrids; inconsistent with `get/set/unset` standard |
| action/hook help | 3 | `help-action-commands`, `help-hook-commands`, `show-action` | Yes | Well-scoped help surface |

---

