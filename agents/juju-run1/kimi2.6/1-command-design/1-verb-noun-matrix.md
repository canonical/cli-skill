## Section 1: Verb-Noun Decomposition Matrix

Rows are verbs sorted alphabetically. Columns are primary resource nouns sorted roughly by frequency (most numerous on the left). `✓` = command exists in that slot. `—` = absent. Orphan commands (not fitting `verb-noun`) are listed after the matrix.

### Matrix

| Verb ↓ / Noun → | app | unit | machine | model | ctrl | cloud | cred | secret | sec-b | space | ssh-key | storage | st-pool | offer | relat | saas | k8s | bundle | backup | firewall | fs | action | task | oper | hook | charm | region | misc |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| **add** | — | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | sec |
| **attach** | resource | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **cancel** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — |
| **change** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | usr-pwd |
| **create** | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — |
| **debug** | — | code | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | log |
| **default** | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — |
| **destroy** | — | — | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **detach** | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **diff** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — |
| **disable** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | cmd, user |
| **download** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — |
| **enable** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | cmd, user, destr-ctrl |
| **export** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — |
| **find** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | offers, charm |
| **grant** | — | — | — | ✓* | — | ✓* | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | model/etc* |
| **help** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | ✓ | — | — | topics |
| **import** | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — |
| **integrate** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **kill** | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **login** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **logout** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **migrate** | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **move-to** | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **offer** | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **refresh** | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **register** | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **reload** | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **remove** | ✓ | ✓ | ✓ | — | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | saas |
| **rename** | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **resolved** | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **resume** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **retry** | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **revoke** | — | — | — | ✓* | — | ✓* | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | model/etc* |
| **run** | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — |
| **scale** | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **set** | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | constr, m-constr |
| **show** | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | — | ✓ | — | — | — | — | — | — | — | ✓ | ✓ | ✓ | — | — | — | st-log, sec-b |
| **suspend** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **switch** | — | — | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **sync** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | agent-bin |
| **trust** | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **unexpose** | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **unregister** | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **update** | — | — | — | — | — | ✓ | ✓ | ✓ | ✓ | — | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | pub-cld |
| **upgrade** | — | — | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| **whoami** | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |

---

### Orphan Commands (No Clear Verb-Noun Decomposition)

Orphan commands fall into several sub-categories. Some are noun-plural list shorthands, some are bare verbs with no object, and some are noun-first configuration/query commands.

#### Noun-plural list/show shorthands

| Command | Canonical Form | What it lists/shows |
|---|---|---|
| `actions` | list-actions | Actions for an application |
| `clouds` | list-clouds | Available clouds |
| `controllers` | list-controllers | Controllers |
| `credentials` | list-credentials | Credentials for a cloud |
| `machines` | list-machines | Machines in a model |
| `models` | list-models | Accessible models |
| `offers` | list-offers | Shared endpoints |
| `operations` | list-operations | Pending/running/completed operations |
| `regions` | list-regions | Regions for a cloud |
| `secret-backends` | list-secret-backends | Secret backends |
| `secrets` | list-secrets | Secrets in a model |
| `spaces` | list-spaces | Network spaces |
| `ssh-keys` | list-ssh-keys | SSH keys |
| `storage-pools` | list-storage-pools | Storage pools |
| `subnets` | list-subnets | Subnets |
| `users` | list-users | Juju users |

#### Noun-first/get-set configuration hybrids

| Command | Notes |
|---|---|
| `config` | Get/set/reset application configuration |
| `controller-config` | Get/set controller configuration |
| `model-config` | Get/set model configuration |
| `model-defaults` | Get/set defaults for new models |
| `model-constraints` | Display model constraints |
| `model-secret-backend` | Get/set secret backend for model |
| `constraints` | Display machine constraints |
| `firewall-rules` | Print firewall rules |
| `status` | Report model status |
| `storage` | List storage details |
| `resources` | Show resources for app/unit |
| `charm-resources` | Display resources for a charm in repo |

#### Bare verbs / participles

| Command | Notes |
|---|---|
| `bind` | Change application bindings |
| `bootstrap` | Initialize cloud / create controller |
| `consume` | Add a remote offer to the model |
| `dashboard` | Print/open dashboard URL |
| `debug-log` | Display log messages |
| `deploy` | Deploy application/bundle |
| `documentation` | Generate command docs |
| `exec` | Run commands on remote targets |
| `expose` | Make application publicly available |
| `find` | Query Charmhub |
| `find-offers` | Find offered endpoints |
| `help` | General help dispatcher |
| `info` | Display charm info |
| `integrate` | Integrate two applications |
| `login` | Log in to controller |
| `logout` | Log out |
| `migrate` | Migrate model to another controller |
| `offer` | Offer application endpoints |
| `refresh` | Refresh application charm |
| `register` | Register controller |
| `reload-spaces` | Reload spaces from substrate |
| `resolved` | Mark unit errors resolved |
| `run` | Run an action |
| `scp` | Secure copy |
| `ssh` | SSH session |
| `switch` | Switch controller/model |
| `sync-agent-binary` | Sync agent binaries |
| `trust` | Set trust true |
| `unexpose` | Reverse expose |
| `update-public-clouds` | Update public cloud metadata |
| `version` | Print version |
| `whoami` | Print current login |

#### Command-set management

| Command | Notes |
|---|---|
| `disabled-commands` | List disabled commands |
| `disable-command` | Disable a command set |
| `enable-command` | Re-enable a command set |
| `enable-destroy-controller` | Remove controller-level command disabling |

---

### Annotations

#### Incomplete CRUD Sets

| Resource | Has "Create" | Has "Read" | Has "Update" | Has "Delete" | Missing / Notes |
|---|---|---|---|---|---|
| cloud | `add-cloud` | `show-cloud`, `clouds` | `update-cloud` | `remove-cloud` | `create-cloud` missing (add used) |
| controller | `bootstrap` | `show-controller`, `controllers` | — | `destroy-controller`, `kill-controller` | `update-controller` missing; two delete verbs |
| model | `add-model` | `show-model`, `models`, `status` | `model-config`, `model-defaults` | `destroy-model` | `remove-model` missing |
| application | `deploy` | `show-application`, `status` | `config`, `refresh`, `scale-application` | `remove-application` | No symmetric `add-application` |
| unit | `add-unit` | `show-unit`, `status` | — | `remove-unit` | `update-unit` missing |
| machine | `add-machine` | `show-machine`, `machines` | — | `remove-machine` | `update-machine` missing |
| credential | `add-credential` | `show-credential`, `credentials` | `update-credential` | `remove-credential` | `create-credential` missing (add used) |
| secret | `add-secret` | `show-secret`, `secrets` | `update-secret` | `remove-secret` | `create-secret` missing (add used) |
| secret-backend | `add-secret-backend` | `show-secret-backend`, `secret-backends` | `update-secret-backend` | `remove-secret-backend` | `create-*` missing (add used) |
| space | `add-space` | `show-space`, `spaces` | `rename-space`, `move-to-space` | `remove-space` | `update-space` missing (rename/move instead) |
| ssh-key | `add-ssh-key`, `import-ssh-key` | `ssh-keys` | — | `remove-ssh-key` | No `show-ssh-key`; `update-ssh-key` missing |
| storage | `add-storage` | `show-storage`, `storage` | — | `remove-storage` | `update-storage` missing |
| storage-pool | `create-storage-pool` | `storage-pools` | `update-storage-pool` | `remove-storage-pool` | `show-storage-pool` missing |
| offer | `offer` | `show-offer`, `offers` | — | `remove-offer` | `update-offer` missing |
| relation | `integrate` | — | `suspend-relation`, `resume-relation` | `remove-relation` | `show-relation` missing |
| saas | `consume` | — | — | `remove-saas` | `show-saas` missing |
| k8s | `add-k8s` | — | `update-k8s` | `remove-k8s` | `show-k8s`, `k8s` list missing |
| bundle | `deploy` | — | — | — | `show-bundle`, `list-bundles` missing |
| backup | `create-backup` | — | — | — | `show-backup`, `list-backups`, `remove-backup` missing |
| firewall-rule | `set-firewall-rule` | `firewall-rules` | `set-firewall-rule` (overwrite) | — | No `remove-firewall-rule` |
| filesystem | `import-filesystem` | — | — | — | CRUD extremely thin |
| action | — | `show-action`, `actions` | — | — | No `add-action` (declared in charm) |
| task | — | `show-task` | — | `cancel-task` | `list-tasks` missing |
| operation | — | `show-operation`, `operations` | — | — | `cancel-operation` missing |
| charm | `download` | `info`, `find` | — | — | `show-charm` missing (info used) |
| region | — | `regions` | `default-region` | — | `add-region` / `remove-region` missing |

#### Verb Inconsistencies

1. **`destroy-*` vs `remove-*`**: Controller and model use `destroy` (`destroy-controller`, `destroy-model`), while almost every other resource uses `remove`. These are the only two resources with `destroy` as their delete verb.
2. **`deploy` vs `add-*`**: Applications and bundles are created with `deploy`, not `add-application`. This is the only significant resource that bypasses the `add` verb entirely.
3. **`kill-controller` vs `destroy-controller`**: Two different verbs for deleting a controller. `kill` implies force; `destroy` implies normal teardown. This is an exception path masquerading as a top-level command.
4. **`consume` vs `add-*`**: SAAS entries are created with `consume`, not `add-saas`.
5. **`download` vs `get-*` / `show-*`**: Charms use `download` for retrieval and `info` for detail — neither maps to `get` or `show` cleanly.
6. **`default-*` as a verb**: `default-credential` and `default-region` use an adjective as a pseudo-verb. No other command follows this pattern.
7. **`create-*` limited to storage-pool and backup**: Every other resource uses `add-*` for creation. Storage pools and backups are the exceptions.
8. **`enable-destroy-controller`**: Embeds another command name (`destroy-controller`) rather than a resource noun. This is unique.

---

