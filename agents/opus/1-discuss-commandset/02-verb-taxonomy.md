# Verb Taxonomy and Aspect Classification

## Unique Verbs from Command Set

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|
| add | lifecycle | telic | partial | remove | add-cloud, add-model, add-unit |
| attach | transfer | telic | yes | detach | attach-storage, attach-resource |
| autoload | access | telic | no | — | autoload-credentials |
| bind | mutation | telic | partial | unbind (missing) | bind |
| bootstrap | lifecycle | telic | no | — | bootstrap |
| cancel | execution | punctual | no | — | cancel-task |
| change | mutation | telic | yes | — | change-user-password |
| create | lifecycle | telic | no | — | create-backup, create-storage-pool |
| debug | execution | atelic | no | — | debug-hooks, debug-code, debug-log |
| deploy | lifecycle | telic | partial | remove-application | deploy |
| destroy | lifecycle | telic | no | — | destroy-controller, destroy-model |
| detach | transfer | telic | yes | attach | detach-storage |
| diff | observation | telic | no | — | diff-bundle |
| disable | access | punctual | yes | enable | disable-command, disable-user |
| download | transfer | telic | no | — | download, download-backup |
| dump | observation | telic | no | — | dump-model, dump-db |
| enable | access | punctual | yes | disable | enable-command, enable-destroy-controller, enable-user |
| exec | execution | telic | no | — | exec |
| export | transfer | telic | no | import | export-bundle |
| expose | transfer | punctual | yes | unexpose | expose |
| find | observation | telic | no | — | find, find-offers |
| grant | access | punctual | yes | revoke | grant, grant-cloud, grant-secret |
| help | observation | telic | no | — | help-action-commands, help-hook-commands |
| import | transfer | telic | no | export | import-filesystem, import-ssh-key |
| integrate | transfer | telic | partial | remove-relation | integrate |
| kill | lifecycle | telic | no | — | kill-controller |
| login | access | punctual | yes | logout | login |
| logout | access | punctual | yes | login | logout |
| migrate | migration | telic | no | — | migrate |
| move | mutation | telic | no | — | move-to-space |
| offer | transfer | telic | partial | remove-offer | offer |
| refresh | mutation | telic | partial | — | refresh |
| register | access | punctual | yes | unregister | register |
| reload | mutation | punctual | no | — | reload-spaces |
| remove | lifecycle | telic | partial | add | remove-application, remove-cloud, remove-unit |
| rename | mutation | telic | no | — | rename-space |
| resolve | execution | punctual | no | — | resolved |
| retry | execution | atelic | no | — | retry-provisioning |
| revoke | access | punctual | yes | grant | revoke, revoke-cloud, revoke-secret |
| run | execution | telic | no | — | run |
| scale | mutation | telic | partial | scale (to 0) | scale-application |
| scp | transfer | telic | no | — | scp |
| set | mutation | punctual | yes | unset (missing) | set-constraints, set-credential, set-firewall-rule |
| show | observation | telic | no | — | show-action, show-application, show-cloud, show-controller, show-credential, show-machine, show-model, show-offer, show-operation, show-secret, show-secret-backend, show-space, show-status-log, show-storage, show-task, show-unit, show-user |
| ssh | execution | atelic | no | — | ssh |
| suspend | transfer | punctual | yes | resume | suspend-relation |
| resume | transfer | punctual | yes | suspend | resume-relation |
| switch | access | punctual | no | — | switch |
| sync | migration | telic | no | — | sync-agent-binary |
| trust | access | punctual | yes | untrust (missing) | trust |
| unexpose | transfer | punctual | yes | expose | unexpose |
| unregister | access | punctual | yes | register | unregister |
| update | mutation | telic | no | — | update-cloud, update-credential, update-k8s, update-public-clouds, update-secret, update-secret-backend, update-storage-pool |
| upgrade | mutation | telic | no | — | upgrade-controller, upgrade-model |

## Implicit Verbs for Noun-Only Commands

| Noun-Only Command | Implied Verb | Intent Group | Aspect |
|---|---|---|---|
| actions | list | observation | telic |
| clouds | list | observation | telic |
| controllers | list | observation | telic |
| credentials | list | observation | telic |
| disabled-commands | list | observation | telic |
| firewall-rules | list | observation | telic |
| machines | list | observation | telic |
| models | list | observation | telic |
| offers | list | observation | telic |
| operations | list | observation | telic |
| regions | list | observation | telic |
| resources | list | observation | telic |
| secret-backends | list | observation | telic |
| secrets | list | observation | telic |
| spaces | list | observation | telic |
| ssh-keys | list | observation | telic |
| storage | list | observation | telic |
| storage-pools | list | observation | telic |
| subnets | list | observation | telic |
| users | list | observation | telic |
| status | show | observation | telic |
| info | show | observation | telic |

## Aspect Distribution

| Aspect | Count | Verbs |
|---|---|---|
| Telic | 38 | add, attach, autoload, bind, bootstrap, cancel, change, create, debug, deploy, destroy, detach, diff, disable, download, dump, enable, exec, export, expose, find, grant, help, import, integrate, kill, login, logout, migrate, move, offer, refresh, register, reload, remove, rename, resolve, retry, revoke, run, scale, scp, set, show, suspend, resume, sync, trust, unexpose, unregister, update, upgrade |
| Atelic | 3 | debug, retry, ssh |
| Punctual | 10 | disable, enable, expose, grant, login, logout, register, revoke, set, suspend, resume, switch, trust, unexpose, unregister |

## Reversibility Analysis

| Reversibility | Count | Examples |
|---|---|---|
| Yes (named inverse) | 14 | attach/detach, disable/enable, expose/unexpose, grant/revoke, login/logout, register/unregister, suspend/resume |
| No | 18 | bootstrap, cancel, create, destroy, diff, download, dump, exec, find, help, kill, migrate, remove (no generic add), retry, run, scp, ssh, switch, sync, update, upgrade |
| Partial | 10 | add/remove (asymmetric granularity), deploy/remove-application, integrate/remove-relation, offer/remove-offer, scale (to 0), set/unset (missing), trust/untrust (missing), bind/unbind (missing), change/revert (missing), import/export (not round-trip) |

## Key Observations

1. **Observation dominance**: `show` is the most overloaded verb, appearing in 18 commands. This aligns with DE013 recommendations for `show` as the primary read verb.

2. **Lifecycle verb fragmentation**: Juju uses `add`, `create`, `deploy`, `bootstrap`, `kill`, `destroy`, and `remove` for creation/destruction. Per DE013 §Grammar, this should converge to a smaller set (ideally `create`/`delete` or `add`/`remove`).

3. **Missing `unset`**: `set-constraints`, `set-credential`, `set-firewall-rule`, and `set-model-constraints` have no corresponding `unset` command. Users must know default values to revert.

4. **Toggle verbs without nouns**: `expose`, `unexpose`, `trust`, `enable`, `disable` act on implicit objects, reducing clarity.

5. **`resolved` is grammatically odd**: A past participle used as an imperative. The alias `resolve` is more natural but secondary.

6. **`migrate` is a migration verb without a return path**: There is no `unmigrate` or `migrate-back` command.
