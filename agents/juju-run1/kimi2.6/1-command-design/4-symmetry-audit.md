## Section 4: Symmetry Audit

### Explicit Symmetric Pairs

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|
| Add/Remove cloud | `add-cloud` | `remove-cloud` | Yes | Yes | Standard CRUD pair |
| Add/Remove credential | `add-credential` | `remove-credential` | Yes | Yes | Standard CRUD pair |
| Add/Remove k8s | `add-k8s` | `remove-k8s` | Yes | Yes | Standard CRUD pair |
| Add/Remove machine | `add-machine` | `remove-machine` | Yes | Yes | Standard CRUD pair |
| Add/Remove model | `add-model` | `destroy-model` | **No** | Partial | `destroy` used instead of `remove` |
| Add/Remove secret | `add-secret` | `remove-secret` | Yes | Yes | Standard CRUD pair |
| Add/Remove secret-backend | `add-secret-backend` | `remove-secret-backend` | Yes | Yes | Standard CRUD pair |
| Add/Remove space | `add-space` | `remove-space` | Yes | Yes | Standard CRUD pair |
| Add/Remove SSH key | `add-ssh-key` | `remove-ssh-key` | Yes | Yes | Standard CRUD pair |
| Add/Remove storage | `add-storage` | `remove-storage` | Yes | Yes | Standard CRUD pair |
| Add/Remove storage-pool | `create-storage-pool` | `remove-storage-pool` | **No** | Yes | `create` vs `remove` |
| Add/Remove unit | `add-unit` | `remove-unit` | Yes | Yes | Standard CRUD pair |
| Add/Remove user | `add-user` | `remove-user` | Yes | Yes | Standard CRUD pair |
| Attach/Detach storage | `attach-storage` | `detach-storage` | Yes | Yes | Well paired |
| Enable/Disable command | `enable-command` | `disable-command` | Yes | Yes | Well paired |
| Enable/Disable user | `enable-user` | `disable-user` | Yes | Yes | Well paired |
| Expose/Unexpose app | `expose` | `unexpose` | Yes | Yes | Simple boolean toggle |
| Grant/Revoke (general) | `grant` | `revoke` | Yes | Yes | Well paired |
| Grant/Revoke cloud | `grant-cloud` | `revoke-cloud` | Yes | Yes | Well paired |
| Grant/Revoke secret | `grant-secret` | `revoke-secret` | Yes | Yes | Well paired |
| Login/Logout | `login` | `logout` | Yes | Yes | Standard session pair |
| Offer/Remove offer | `offer` | `remove-offer` | Yes | Yes | Well paired |
| Register/Unregister | `register` | `unregister` | Yes | Yes | Standard pair |
| Suspend/Resume relation | `suspend-relation` | `resume-relation` | Yes | Yes | Well paired |

### Asymmetric / Problematic Pairs

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|
| Create/Destroy controller | `bootstrap` | `destroy-controller` | **No** | Partial | `bootstrap` is domain-specific; `kill-controller` is an unpaired force variant |
| Deploy/Remove application | `deploy` | `remove-application` | **No** | Partial | `deploy` is the canonical verb; no `add-application` |
| Integrate/Remove relation | `integrate` | `remove-relation` | **No** | Partial | No `add-relation`; `integrate` is the create verb |
| Consume/Remove SAAS | `consume` | `remove-saas` | **No** | Partial | `consume` is the create verb; no `add-saas` |
| Trust/Untrust | `trust` | — | **No** | — | No `untrust`; presumably set `trust=false` or re-run with flag |
| Add/Remove application | — | `remove-application` | — | — | No symmetric `add-application` at all |
| Add/Remove relation | — | `remove-relation` | — | — | No symmetric `add-relation` at all |
| Create/Remove backup | `create-backup` | — | — | — | No `remove-backup` or `delete-backup` |
| Set/Remove firewall rule | `set-firewall-rule` | — | — | — | No `remove-firewall-rule` |
| Update/Remove cloud | `update-cloud` | `remove-cloud` | Yes | Yes | Pair exists but create uses `add` not `create` |
| Upgrade/Downgrade controller | `upgrade-controller` | — | — | — | No `downgrade-controller` |
| Upgrade/Downgrade model | `upgrade-model` | — | — | — | No `downgrade-model` |

---

