# Symmetry Audit

## Complete Symmetry Matrix

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|
| Add/Remove cloud | `add-cloud` | `remove-cloud` | Yes | Yes | |
| Add/Remove credential | `add-credential` | `remove-credential` | Yes | Yes | |
| Add/Remove k8s | `add-k8s` | `remove-k8s` | Yes | Yes | |
| Add/Remove machine | `add-machine` | `remove-machine` | Yes | Yes | |
| Add/Remove model | `add-model` | `destroy-model` | **No** | Partial | Reverse uses `destroy`, not `remove` |
| Add/Remove secret | `add-secret` | `remove-secret` | Yes | Yes | |
| Add/Remove secret-backend | `add-secret-backend` | `remove-secret-backend` | Yes | Yes | |
| Add/Remove space | `add-space` | `remove-space` | Yes | Yes | |
| Add/Remove SSH key | `add-ssh-key` | `remove-ssh-key` | Yes | Yes | |
| Add/Remove storage | `add-storage` | `remove-storage` | Yes | Yes | |
| Add/Remove unit | `add-unit` | `remove-unit` | Yes | Yes | |
| Add/Remove user | `add-user` | `remove-user` | Yes | Yes | |
| Attach/Detach storage | `attach-storage` | `detach-storage` | Yes | Yes | |
| Create/Remove backup | `create-backup` | — | — | — | No reverse; download-backup is retrieval, not destruction |
| Create/Remove storage pool | `create-storage-pool` | `remove-storage-pool` | Yes | Yes | |
| Deploy/Remove application | `deploy` | `remove-application` | **No** | Partial | Forward is verb-only; reverse is verb-noun. No `undeploy`. |
| Enable/Disable command | `enable-command` | `disable-command` | Yes | Yes | |
| Enable/Disable user | `enable-user` | `disable-user` | Yes | Yes | |
| Expose/Unexpose application | `expose` | `unexpose` | Yes | Yes | Toggle pair without noun |
| Grant/Revoke model access | `grant` | `revoke` | Yes | Yes | |
| Grant/Revoke cloud access | `grant-cloud` | `revoke-cloud` | Yes | Yes | |
| Grant/Revoke secret access | `grant-secret` | `revoke-secret` | Yes | Yes | |
| Integrate/Remove relation | `integrate` | `remove-relation` | **No** | Partial | `integrate` lacks noun; `add-relation` alias was removed in favor of `integrate`. |
| Login/Logout user | `login` | `logout` | Yes | Yes | |
| Offer/Remove offer | `offer` | `remove-offer` | **No** | Partial | Forward is verb-only; no `unoffer`. |
| Register/Unregister controller | `register` | `unregister` | Yes | Yes | |
| Resume/Suspend relation | `resume-relation` | `suspend-relation` | Yes | Yes | |
| Set/Show constraints (app) | `set-constraints` | `constraints` | **No** | Yes | Get uses noun-only; set uses verb-noun. No `unset-constraints`. |
| Set/Show constraints (model) | `set-model-constraints` | `model-constraints` | **No** | Yes | Same asymmetry as application constraints. |
| Set/Show model credential | `set-credential` | `show-credential` | **No** | Yes | Set is verb-noun; show is verb-noun but `set-credential` is model-scoped while `show-credential` is cloud-scoped. |
| Trust/Untrust application | `trust` | — | — | — | No reverse command. Use `trust --remove` (flag-based toggle). |
| Add/Remove resource | `attach-resource` | — | — | — | No detach/remove resource command. |
| Consume/Remove SAAS | `consume` | `remove-saas` | **No** | Partial | Forward is verb-only; reverse uses `saas` noun. No `unconsume`. |
| Bootstrap/Destroy controller | `bootstrap` | `destroy-controller` | **No** | Partial | `bootstrap` has no clear reverse verb. `kill-controller` is forceful alternative. |
| Add/Remove block | `enable-command` | `disable-command` | Yes | Yes | Enables/disables command blocks. |
| Import/Export bundle | `import-filesystem` | `export-bundle` | **No** | No | Not symmetric operations; different domains. |
| Upgrade controller | `upgrade-controller` | — | — | — | No downgrade command. |
| Upgrade model | `upgrade-model` | — | — | — | No downgrade command. |
| Add/Detach storage (unit) | `attach-storage` | `detach-storage` | Yes | Yes | |
| Import/Remove SSH key | `import-ssh-key` | `remove-ssh-key` | **No** | Yes | Forward is `import`; reverse is `remove`. Asymmetric verb choice. |
| Move/Rename space | `move-to-space` | `rename-space` | **No** | No | Different mutations, not inverses. |
| Reload/Remove space | `reload-spaces` | `remove-space` | **No** | No | Not symmetric. |
| Show/Update cloud | `show-cloud` | `update-cloud` | Yes | Yes | |
| Show/Update credential | `show-credential` | `update-credential` | Yes | Yes | |
| Show/Update secret | `show-secret` | `update-secret` | Yes | Yes | |
| Show/Update secret-backend | `show-secret-backend` | `update-secret-backend` | Yes | Yes | |
| Show/Update storage-pool | `show-storage` | `update-storage-pool` | **No** | Yes | `show-storage` vs `update-storage-pool` — noun mismatch. |
| Create/Download backup | `create-backup` | `download-backup` | **No** | No | Download is not the inverse of create. |
| Add/Remove relation | `integrate` | `remove-relation` | **No** | Partial | `integrate` replaced `add-relation`. The old `add-relation` alias was removed. |
| Change/Reset password | `change-user-password` | — | — | — | No reset-password command. Admin must use `change-user-password` on behalf of user. |
| Debug/Resolve unit | `debug-hooks` | `resolved` | **No** | No | Different operations; `resolved` marks errors as resolved, not inverse of debug. |

## Missing Reverse Operations Summary

| Forward Command | Missing Reverse | Impact |
|---|---|---|
| `deploy` | `undeploy` or `remove-application` (naming mismatch) | Medium |
| `bootstrap` | `unbootstrap` or `destroy-controller` (naming mismatch) | Low |
| `create-backup` | `remove-backup` or `delete-backup` | Low |
| `trust` | `untrust` | Low |
| `set-constraints` | `unset-constraints` | Medium |
| `set-model-constraints` | `unset-model-constraints` | Medium |
| `set-firewall-rule` | `unset-firewall-rule` or `remove-firewall-rule` | Low |
| `set-credential` | `unset-credential` | Low |
| `consume` | `unconsume` | Low |
| `offer` | `unoffer` | Low |
| `integrate` | `disintegrate` or `unintegrate` | Low |
| `upgrade-controller` | `downgrade-controller` | Low |
| `upgrade-model` | `downgrade-model` | Low |
| `enable-destroy-controller` | `disable-destroy-controller` | Low |
| `migrate` | `unmigrate` | Low |
| `exec` | — | N/A |
| `run` | — | N/A |
| `refresh` | `unrefresh` or `revert` | Low |
| `bind` | `unbind` | Low |
| `scale-application` | `unscale-application` (scale to 0 is partial) | Low |

## Naming Asymmetries by Severity

### Critical (breaks mental model)

| Pair | Issue |
|---|---|
| `add-model` / `destroy-model` | `destroy` is not the inverse of `add` |
| `bootstrap` / `destroy-controller` | `bootstrap` has no semantic relationship to `destroy` |
| `deploy` / `remove-application` | `deploy` is not the inverse of `remove` |

### High (consistent confusion)

| Pair | Issue |
|---|---|
| `integrate` / `remove-relation` | `integrate` replaced `add-relation`, breaking symmetry |
| `set-constraints` / `constraints` | Get command is noun-only, set is verb-noun |
| `consume` / `remove-saas` | `consume` has no noun; `saas` is a different concept from `offer` |

### Medium (minor friction)

| Pair | Issue |
|---|---|
| `import-ssh-key` / `remove-ssh-key` | `import` vs `remove` verb mismatch |
| `create-backup` / `download-backup` | Not inverses |
| `attach-resource` / — | No removal command |
| `login` / `logout` | Symmetric but `whoami` is a third command in the same domain |

## Behavioral Asymmetries

| Forward | Reverse | Asymmetry |
|---|---|---|
| `destroy-controller --force` | `bootstrap` | Reverse requires cloud setup; forward can be forced |
| `remove-application --force` | `deploy` | Deploy provisions resources; remove with force skips cleanup |
| `remove-unit --destroy-storage` | `add-unit` | Add does not auto-create storage; remove can destroy it |
| `kill-controller` | `destroy-controller` | `kill` is more forceful but both destroy; two commands for same outcome with different safety |
