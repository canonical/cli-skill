# 02 — Verb Taxonomy and Aspect Classification

## Unique Verb Inventory

Every unique verb extracted from the decomposition matrix, classified by intent group and linguistic aspect.

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|------|-------------|--------|------------|-------------|--------------|
| `add` | lifecycle | telic | yes | `remove` | add-cloud, add-credential, add-k8s, add-machine, add-model, add-secret, add-secret-backend, add-space, add-ssh-key, add-storage, add-unit, add-user |
| `attach` | transfer | telic | yes | `detach` | attach-resource, attach-storage |
| `autoload` | lifecycle | telic | no | — | autoload-credentials |
| `bind` | mutation | telic | partial | — | bind |
| `bootstrap` | lifecycle | telic | no | — | bootstrap |
| `cancel` | execution | telic | no | — | cancel-task |
| `change` | mutation | telic | partial | — | change-user-password |
| `config` (app) | mutation | atelic | partial | — | config |
| `config` (controller) | mutation | atelic | partial | — | controller-config |
| `config` (model) | mutation | atelic | partial | — | model-config |
| `consume` | transfer | telic | yes | `remove-saas` | consume |
| `create` | lifecycle | telic | yes | `remove`/`destroy` | create-backup, create-storage-pool |
| `dashboard` | execution | punctual | no | — | dashboard |
| `debug` | execution | atelic | no | — | debug-code, debug-hooks, debug-log |
| `default` (set) | mutation | punctual | partial | — | default-credential, default-region |
| `deploy` | lifecycle | telic | yes | `remove-application` | deploy |
| `destroy` | lifecycle | telic | no | — | destroy-controller, destroy-model |
| `detach` | transfer | telic | yes | `attach` | detach-storage |
| `diff` | observation | punctual | no | — | diff-bundle |
| `disable` | access | punctual | yes | `enable` | disable-command, disable-user |
| `download` | transfer | telic | no | — | download, download-backup |
| `dump` | observation | telic | no | — | dump-db, dump-model |
| `enable` | access | punctual | yes | `disable` | enable-command, enable-user, enable-destroy-controller |
| `exec` | execution | telic | no | — | exec |
| `export` | transfer | telic | yes | `import` | export-bundle |
| `expose` | access | punctual | yes | `unexpose` | expose |
| `find` | observation | telic | no | — | find, find-offers |
| `get` (constraints) | observation | punctual | no | — | constraints, model-constraints |
| `grant` | access | punctual | yes | `revoke` | grant, grant-cloud, grant-secret |
| `help` | observation | punctual | no | — | help-action-commands, help-hook-commands |
| `import` | transfer | telic | yes | `export` | import-filesystem, import-ssh-key |
| `info` | observation | punctual | no | — | info |
| `integrate` | transfer | telic | yes | `remove-relation` | integrate |
| `kill` | lifecycle | telic | no | — | kill-controller |
| `list` (plural nouns) | observation | punctual | no | — | actions, clouds, controllers, credentials, disabled-commands, firewall-rules, machines, models, offers, operations, regions, resources, secret-backends, secrets, spaces, ssh-keys, storage, storage-pools, subnets, users |
| `login` | access | punctual | yes | `logout` | login |
| `logout` | access | punctual | yes | `login` | logout |
| `migrate` | migration | telic | no | — | migrate |
| `model` (config/defaults) | mutation | atelic | partial | — | model-config, model-defaults, model-constraints, model-secret-backend |
| `move` | transfer | telic | no | — | move-to-space |
| `offer` | transfer | telic | yes | `remove-offer` | offer |
| `refresh` | mutation | telic | partial | — | refresh |
| `register` | access | punctual | yes | `unregister` | register |
| `reload` | mutation | telic | no | — | reload-spaces |
| `remove` | lifecycle | telic | yes | `add` | remove-application, remove-cloud, remove-credential, remove-k8s, remove-machine, remove-offer, remove-relation, remove-saas, remove-secret, remove-secret-backend, remove-space, remove-ssh-key, remove-storage, remove-storage-pool, remove-unit, remove-user |
| `rename` | mutation | punctual | no | — | rename-space |
| `resolve` | execution | telic | no | — | resolved |
| `resume` | execution | punctual | yes | `suspend` | resume-relation |
| `retry` | execution | telic | no | — | retry-provisioning |
| `revoke` | access | punctual | yes | `grant` | revoke, revoke-cloud, revoke-secret |
| `run` | execution | telic | no | — | run |
| `scale` | mutation | telic | partial | — | scale-application |
| `scp` | transfer | telic | no | — | scp |
| `set` | mutation | punctual | yes | `unset`/`get` | set-constraints, set-credential, set-firewall-rule, set-model-constraints |
| `show` | observation | punctual | no | — | show-action, show-application, show-cloud, show-controller, show-credential, show-machine, show-model, show-offer, show-operation, show-secret, show-secret-backend, show-space, show-status-log, show-storage, show-task, show-unit, show-user |
| `ssh` | execution | atelic | no | — | ssh |
| `status` | observation | atelic | no | — | status |
| `suspend` | execution | punctual | yes | `resume` | suspend-relation |
| `switch` | migration | punctual | no | — | switch |
| `sync` | migration | telic | no | — | sync-agent-binary |
| `trust` | access | punctual | partial | — | trust |
| `unexpose` | access | punctual | yes | `expose` | unexpose |
| `unregister` | access | punctual | yes | `register` | unregister |
| `update` | mutation | telic | partial | — | update-cloud, update-credential, update-k8s, update-public-clouds, update-secret, update-secret-backend, update-storage-pool |
| `upgrade` | mutation | telic | no | — | upgrade-controller, upgrade-model |
| `version` | observation | punctual | no | — | version |
| `whoami` | observation | punctual | no | — | whoami |

### Verb Count Summary

| Intent Group | Count | Verbs |
|-------------|-------|-------|
| lifecycle | 8 | add, autoload, bootstrap, create, deploy, destroy, kill, remove |
| mutation | 9 | bind, change, config (3 variants), default, refresh, reload, rename, scale, set, update, upgrade |
| access | 9 | disable, enable, expose, grant, login, logout, register, revoke, trust, unexpose, unregister |
| observation | 9 | diff, dump, find, get, help, info, list, show, status, version, whoami |
| transfer | 8 | attach, consume, detach, download, export, import, integrate, move, offer, scp |
| execution | 9 | cancel, dashboard, debug, exec, resolve, resume, retry, run, ssh, suspend |
| migration | 3 | migrate, switch, sync |

### Aspect Analysis

- **Telic (has natural endpoint)**: 32 verbs — Most lifecycle, transfer, and execution verbs. These operations complete.
- **Atelic (ongoing)**: 3 verbs — `debug`, `ssh`, `status`. These can run indefinitely.
- **Punctual (instant)**: 15 verbs — Access toggles and most observation verbs. State changes happen immediately.

### Reversibility Summary

| Category | Count | Verbs |
|----------|-------|-------|
| Fully reversible (paired) | 14 | add/remove, attach/detach, disable/enable, expose/unexpose, export/import, grant/revoke, login/logout, register/unregister, resume/suspend, deploy/remove-application, integrate/remove-relation, offer/remove-offer, consume/remove-saas, create/remove (backup, storage-pool) |
| Partially reversible | 6 | bind, change, refresh, scale, update, trust |
| Irreversible | 19 | autoload, bootstrap, cancel, dashboard, destroy, diff, download, dump, exec, find, help, info, kill, migrate, move, reload, retry, upgrade, version, whoami |

### Key Observations

1. **`destroy` is irreversible**: Unlike `remove`, `destroy` has no creation counterpart. `destroy-controller` cannot be reversed by any `add-controller` command; `bootstrap` is the only way to create a controller. Similarly, `destroy-model` has no simple undo. This is correct: destroying the container deletes everything inside it.

2. **`kill` is a special case of `destroy`**: `kill-controller` implies the controller is not cooperating (unreachable). FrameNet would classify `kill` as "Killing" (implies victim was alive) vs `destroy` as "Destroying" (patient is an object). The semantic difference justifies the separate command.

3. **`deploy` as creation**: `deploy` is the creation verb for applications, not `add-application`. This is a deviation from the add-X pattern but is semantically richer: deploying a charm is more than creating an application record.

4. **`resolved` as verb**: "Resolved" is an adjective. The active verb form should be `resolve`. This is a DE013 §Grammar violation: "Commands are verbs."

5. **`default` as verb**: `default-credential` and `default-region` use `default` as a verb, which is unusual. `set-default-credential` would be more standard per DE013's `set` semantics. The alias `set-default-credentials` exists for `default-credential`.

6. **`config` as verb**: Application config uses bare `config`. Model and controller use `model-config` and `controller-config`. The asymmetry could be confusing.
