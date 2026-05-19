# Section 1: Verb-Noun Decomposition Matrix
Commands that decompose cleanly into `verb-noun` are shown below. Orphan commands (no clean decomposition) are listed after the matrix.

## Matrix
| Verb | cloud | credential | secret | storage | user | controller | credentials | model | secret-backend | space | application | k8s | machine | ssh-key | storage-pool | unit | backup | bundle | command | offer | offers | relation | spaces | task | action | action-commands | actions | agent-binary | charm-resources | clouds | code | constraints | controllers | default-credentials | default-region | destroy-controller | disabled-commands | filesystem | firewall-rule | firewall-rules | hook | hook-commands | hooks | keys | log | machines | model-constraints | models | operation | operations | provisioning | public-clouds | region | regions | resource | resources | saas | secret-backends | secrets | ssh-keys | status-log | storage-pools | subnets | to-space | user-password | users |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| add | ✓ | ✓ | ✓ | ✓ | ✓ | — | — | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| attach | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — |
| autoload | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| cancel | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| change | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — |
| create | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| debug | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | ✓ | — | ✓ | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| default | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — |
| destroy | — | — | — | — | — | ✓ | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| detach | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| diff | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| disable | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| download | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| enable | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| export | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| find | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| grant | ✓ | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| help | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| import | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| kill | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| list | — | — | — | ✓ | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | ✓ | — | — | — | ✓ | — | ✓ | ✓ | — | — | ✓ | — | — | — | ✓ | — | — | ✓ | — | — | — | — | — | ✓ | — | ✓ | — | ✓ | — | — | — | ✓ | — | ✓ | — | ✓ | ✓ | ✓ | — | ✓ | ✓ | — | — | ✓ |
| move | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — |
| reload | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| remove | ✓ | ✓ | ✓ | ✓ | ✓ | — | — | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | — | — | ✓ | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — |
| rename | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| retry | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| revoke | ✓ | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| scale | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| set | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | ✓ | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| show | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | — | — | ✓ | — | — | — | ✓ | — | — | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — |
| ssh | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| suspend | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| sync | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| update | ✓ | ✓ | ✓ | — | — | — | ✓ | — | ✓ | — | — | ✓ | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| upgrade | — | — | — | — | — | ✓ | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |

## Orphan Commands
These commands do not decompose cleanly into `verb-noun`:
- `actions`
- `bind`
- `bootstrap`
- `charm-resources`
- `clouds`
- `config`
- `constraints`
- `consume`
- `controller-config`
- `controllers`
- `credentials`
- `dashboard`
- `deploy`
- `disabled-commands`
- `documentation`
- `download`
- `exec`
- `expose`
- `find`
- `firewall-rules`
- `grant`
- `help`
- `info`
- `integrate`
- `login`
- `logout`
- `machines`
- `migrate`
- `model-config`
- `model-constraints`
- `model-default`
- `model-defaults`
- `model-secret-backend`
- `models`
- `offer`
- `offers`
- `operations`
- `refresh`
- `regions`
- `register`
- `relate`
- `resolve`
- `resolved`
- `resources`
- `resume-relation`
- `revoke`
- `run`
- `scp`
- `secret-backends`
- `secrets`
- `spaces`
- `ssh`
- `status`
- `storage`
- `storage-pools`
- `subnets`
- `switch`
- `trust`
- `unexpose`
- `unregister`
- `users`
- `version`
- `whoami`

## Decomposition Reference
| Command | Verb | Noun |
|---|---|---|
| `add-cloud` | add | cloud |
| `add-credential` | add | credential |
| `add-k8s` | add | k8s |
| `add-machine` | add | machine |
| `add-model` | add | model |
| `add-secret` | add | secret |
| `add-secret-backend` | add | secret-backend |
| `add-space` | add | space |
| `add-ssh-key` | add | ssh-key |
| `add-storage` | add | storage |
| `add-unit` | add | unit |
| `add-user` | add | user |
| `attach-resource` | attach | resource |
| `attach-storage` | attach | storage |
| `autoload-credentials` | autoload | credentials |
| `cancel-task` | cancel | task |
| `change-user-password` | change | user-password |
| `create-backup` | create | backup |
| `create-storage-pool` | create | storage-pool |
| `debug-code` | debug | code |
| `debug-hook` | debug | hook |
| `debug-hooks` | debug | hooks |
| `debug-log` | debug | log |
| `default-credential` | default | credential |
| `default-region` | default | region |
| `destroy-controller` | destroy | controller |
| `destroy-model` | destroy | model |
| `detach-storage` | detach | storage |
| `diff-bundle` | diff | bundle |
| `disable-command` | disable | command |
| `disable-user` | disable | user |
| `download-backup` | download | backup |
| `enable-command` | enable | command |
| `enable-destroy-controller` | enable | destroy-controller |
| `enable-user` | enable | user |
| `export-bundle` | export | bundle |
| `find-offers` | find | offers |
| `grant-cloud` | grant | cloud |
| `grant-secret` | grant | secret |
| `help-action-commands` | help | action-commands |
| `help-hook-commands` | help | hook-commands |
| `import-filesystem` | import | filesystem |
| `import-ssh-key` | import | ssh-key |
| `kill-controller` | kill | controller |
| `list-actions` | list | actions |
| `list-charm-resources` | list | charm-resources |
| `list-clouds` | list | clouds |
| `list-controllers` | list | controllers |
| `list-credentials` | list | credentials |
| `list-disabled-commands` | list | disabled-commands |
| `list-firewall-rules` | list | firewall-rules |
| `list-machines` | list | machines |
| `list-models` | list | models |
| `list-offers` | list | offers |
| `list-operations` | list | operations |
| `list-regions` | list | regions |
| `list-resources` | list | resources |
| `list-secret-backends` | list | secret-backends |
| `list-secrets` | list | secrets |
| `list-spaces` | list | spaces |
| `list-ssh-keys` | list | ssh-keys |
| `list-storage` | list | storage |
| `list-storage-pools` | list | storage-pools |
| `list-subnets` | list | subnets |
| `list-users` | list | users |
| `move-to-space` | move | to-space |
| `reload-spaces` | reload | spaces |
| `remove-application` | remove | application |
| `remove-cloud` | remove | cloud |
| `remove-credential` | remove | credential |
| `remove-k8s` | remove | k8s |
| `remove-machine` | remove | machine |
| `remove-offer` | remove | offer |
| `remove-relation` | remove | relation |
| `remove-saas` | remove | saas |
| `remove-secret` | remove | secret |
| `remove-secret-backend` | remove | secret-backend |
| `remove-space` | remove | space |
| `remove-ssh-key` | remove | ssh-key |
| `remove-storage` | remove | storage |
| `remove-storage-pool` | remove | storage-pool |
| `remove-unit` | remove | unit |
| `remove-user` | remove | user |
| `rename-space` | rename | space |
| `retry-provisioning` | retry | provisioning |
| `revoke-cloud` | revoke | cloud |
| `revoke-secret` | revoke | secret |
| `scale-application` | scale | application |
| `set-constraints` | set | constraints |
| `set-credential` | set | credential |
| `set-default-credentials` | set | default-credentials |
| `set-default-region` | set | default-region |
| `set-firewall-rule` | set | firewall-rule |
| `set-model-constraints` | set | model-constraints |
| `show-action` | show | action |
| `show-application` | show | application |
| `show-cloud` | show | cloud |
| `show-controller` | show | controller |
| `show-credential` | show | credential |
| `show-credentials` | show | credentials |
| `show-machine` | show | machine |
| `show-model` | show | model |
| `show-offer` | show | offer |
| `show-operation` | show | operation |
| `show-secret` | show | secret |
| `show-secret-backend` | show | secret-backend |
| `show-space` | show | space |
| `show-status-log` | show | status-log |
| `show-storage` | show | storage |
| `show-task` | show | task |
| `show-unit` | show | unit |
| `show-user` | show | user |
| `ssh-keys` | ssh | keys |
| `suspend-relation` | suspend | relation |
| `sync-agent-binary` | sync | agent-binary |
| `update-cloud` | update | cloud |
| `update-credential` | update | credential |
| `update-credentials` | update | credentials |
| `update-k8s` | update | k8s |
| `update-public-clouds` | update | public-clouds |
| `update-secret` | update | secret |
| `update-secret-backend` | update | secret-backend |
| `update-storage-pool` | update | storage-pool |
| `upgrade-controller` | upgrade | controller |
| `upgrade-model` | upgrade | model |

## Annotations
### Incomplete CRUD Sets
- **cloud**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **credential**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **secret**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **storage**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **user**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **model**: has {'add', 'destroy'} but missing {'remove', 'create', 'deploy', 'kill'}
- **secret-backend**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **space**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **k8s**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **machine**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **ssh-key**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **storage-pool**: has {'remove', 'create'} but missing {'add', 'deploy', 'destroy', 'kill'}
- **unit**: has {'remove', 'add'} but missing {'create', 'deploy', 'destroy', 'kill'}
- **backup**: has {'create'} but missing {'kill', 'remove', 'destroy', 'add', 'deploy'}

### Verb Inconsistencies
- `controller`: uses `destroy` (`destroy-controller`) and `kill` (`kill-controller`) for teardown; `kill` implies forceful termination whereas `destroy` is graceful.
- `application`: uses `remove` (`remove-application`) and `scale` (`scale-application`) for reduction, but no `destroy-application`.
- `model`: uses `destroy` (`destroy-model`) and `add` (`add-model`) but no `remove-model`.
- `saas`: `remove-saas` is used instead of `destroy-saas` or `unregister-saas`, inconsistent with `consume`.

### Orphan Notes
- `bootstrap` is a self-contained initialization verb with no explicit noun.
- `integrate` and `relate` are pure verbs; the relation endpoints are positional arguments.
- `resolved` / `resolve` are past-participle verbs used as standalone commands.
- `whoami`, `status`, `version`, `documentation` are nouns used as commands.
- `trust`, `consume`, `expose`, `unexpose`, `bind`, `refresh`, `migrate`, `switch`, `run`, `exec`, `ssh`, `scp`, `offer`, `find`, `info`, `download`, `help`, `login`, `logout`, `register`, `unregister`, `grant`, `revoke` are pure verbs without a noun in the command name.
- `model-config`, `controller-config`, `application-constraints`, `model-constraints`, `model-defaults`, `model-secret-backend`, `storage-pools`, `firewall-rules`, `secret-backends`, `charm-resources`, `ssh-keys`, `disabled-commands`, `pool-create`, `pool-delete`, `pool-update`, `pool-list`, `filesystem`, `volume`, `actions`, `operations`, `resources`, `offers`, `spaces`, `subnets`, `storage`, `secrets`, `users`, `machines`, `models`, `controllers`, `clouds`, `credentials`, `regions`, `backups`, `blocks`, `caas` are noun-only or noun-noun compounds.


---

# Section 2: Verb Taxonomy and Aspect Classification
| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|
| add | lifecycle | telic | yes | remove | add-cloud, add-model, add-unit |
| attach | transfer | telic | yes | detach | attach-storage, attach-resource |
| autoload | access | telic | partial | remove | autoload-credentials |
| cancel | execution | punctual | partial | run | cancel-task |
| change | mutation | telic | partial | change | change-user-password |
| create | lifecycle | telic | yes | remove | create-backup, create-storage-pool |
| debug | execution | atelic | no | — | debug-log, debug-hooks, debug-code |
| default | mutation | punctual | yes | unset | default-region, default-credential |
| destroy | lifecycle | telic | no | — | destroy-controller, destroy-model |
| detach | transfer | telic | yes | attach | detach-storage |
| diff | observation | telic | no | — | diff-bundle |
| disable | access | punctual | yes | enable | disable-command, disable-user |
| download | transfer | telic | no | — | download, download-backup |
| enable | access | punctual | yes | disable | enable-command, enable-user, enable-destroy-controller |
| export | transfer | telic | no | — | export-bundle |
| find | observation | telic | no | — | find, find-offers |
| grant | access | punctual | yes | revoke | grant, grant-cloud, grant-secret |
| help | observation | telic | no | — | help, help-action-commands, help-hook-commands |
| import | transfer | telic | partial | remove | import-filesystem |
| kill | lifecycle | telic | no | — | kill-controller |
| list | observation | atelic | no | — | list-actions, list-clouds, list-machines, list-models, ... |
| move | mutation | telic | partial | move | move-to-space |
| reload | mutation | telic | partial | reload | reload-spaces |
| remove | lifecycle | telic | yes | add | remove-application, remove-unit, remove-machine, remove-cloud, remove-credential, remove-k8s, remove-model, remove-offer, remove-relation, remove-saas, remove-secret, remove-secret-backend, remove-space, remove-ssh-key, remove-storage, remove-user |
| rename | mutation | telic | partial | rename | rename-space |
| retry | execution | punctual | partial | retry | retry-provisioning |
| revoke | access | punctual | yes | grant | revoke, revoke-cloud, revoke-secret |
| scale | mutation | telic | partial | scale | scale-application |
| set | mutation | punctual | yes | unset | set-constraints, set-credential, set-default-region, set-default-credentials, set-firewall-rule, set-model-constraints |
| show | observation | telic | no | — | show-action, show-application, show-cloud, show-controller, show-credential, show-machine, show-model, show-offer, show-secret, show-space, show-status-log, show-storage, show-task, show-unit, show-user, show-operation, show-secret-backend |
| ssh | execution | atelic | no | — | ssh, scp |
| suspend | mutation | punctual | yes | resume | suspend-relation |
| sync | transfer | telic | no | — | sync-agent-binary |
| update | mutation | telic | partial | update | update-cloud, update-credential, update-k8s, update-public-clouds, update-secret, update-secret-backend, update-storage-pool |
| upgrade | mutation | telic | partial | upgrade | upgrade-controller, upgrade-model |


---

# Section 3: Semantic Domain Clustering
| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---|---|---|---|
| action/execution | 16 | `actions`, `cancel-task`, `debug-code`, `debug-hook`, `debug-hooks`, `debug-log`, `exec`, `list-actions`, `list-operations`, `operations`, `run`, `scp`, `show-action`, `show-operation`, `show-task`, `ssh` | Mixed |  |
| application | 38 | `add-unit`, `attach-resource`, `bind`, `charm-resources`, `config`, `constraints`, `consume`, `deploy`, `diff-bundle`, `export-bundle`, `expose`, `find-offers`, `integrate`, `list-charm-resources`, `list-offers`, `list-resources`, `offer`, `offers`, `refresh`, `relate`, `remove-application`, `remove-offer`, `remove-relation`, `remove-saas`, `remove-unit`, `resolve`, `resolved`, `resources`, `resume-relation`, `scale-application`, `set-constraints`, `show-application`, `show-offer`, `show-unit`, `status`, `suspend-relation`, `trust`, `unexpose` | Mixed | CRUD mostly complete; `deploy` lacks a symmetric `destroy-application`. |
| backup | 2 | `create-backup`, `download-backup` | Mixed |  |
| caas/k8s | 3 | `add-k8s`, `remove-k8s`, `update-k8s` | Mixed |  |
| charmhub | 3 | `download`, `find`, `info` | Yes |  |
| cloud/credential | 23 | `add-cloud`, `add-credential`, `autoload-credentials`, `clouds`, `credentials`, `default-credential`, `default-region`, `grant-cloud`, `list-clouds`, `list-credentials`, `list-regions`, `regions`, `remove-cloud`, `remove-credential`, `revoke-cloud`, `set-credential`, `show-cloud`, `show-credential`, `show-credentials`, `update-cloud`, `update-credential`, `update-credentials`, `update-public-clouds` | Mixed | `default-region` and `default-credential` are noun-first orphans. |
| command-block | 4 | `disable-command`, `disabled-commands`, `enable-command`, `list-disabled-commands` | Mixed |  |
| controller | 8 | `controller-config`, `controllers`, `destroy-controller`, `kill-controller`, `register`, `show-controller`, `unregister`, `upgrade-controller` | Mixed | `kill-controller` vs `destroy-controller` is an inconsistency. |
| dashboard | 1 | `dashboard` | N/A |  |
| firewall | 3 | `firewall-rules`, `list-firewall-rules`, `set-firewall-rule` | Mixed |  |
| machine | 5 | `add-machine`, `list-machines`, `machines`, `remove-machine`, `show-machine` | Mixed |  |
| meta | 7 | `bootstrap`, `documentation`, `help`, `show-status-log`, `switch`, `sync-agent-binary`, `version` | Mixed |  |
| model | 14 | `add-model`, `destroy-model`, `grant`, `list-models`, `migrate`, `model-config`, `model-constraints`, `model-default`, `model-defaults`, `model-secret-backend`, `models`, `revoke`, `show-model`, `upgrade-model` | Mixed | `destroy-model` exists but no `remove-model`; `migrate` is an orphan. |
| other | 9 | `enable-destroy-controller`, `help-action-commands`, `help-hook-commands`, `list-controllers`, `retry-provisioning`, `set-default-credentials`, `set-default-region`, `set-model-constraints`, `ssh-keys` | Mixed |  |
| secret | 14 | `add-secret`, `add-secret-backend`, `grant-secret`, `list-secret-backends`, `list-secrets`, `remove-secret`, `remove-secret-backend`, `revoke-secret`, `secret-backends`, `secrets`, `show-secret`, `show-secret-backend`, `update-secret`, `update-secret-backend` | Mixed |  |
| space/network | 8 | `add-space`, `list-spaces`, `move-to-space`, `reload-spaces`, `remove-space`, `rename-space`, `show-space`, `spaces` | Mixed |  |
| ssh-key | 4 | `add-ssh-key`, `import-ssh-key`, `list-ssh-keys`, `remove-ssh-key` | Mixed |  |
| storage | 13 | `add-storage`, `attach-storage`, `create-storage-pool`, `detach-storage`, `import-filesystem`, `list-storage`, `list-storage-pools`, `remove-storage`, `remove-storage-pool`, `show-storage`, `storage`, `storage-pools`, `update-storage-pool` | Mixed | Pool commands use noun-verb (`pool-create`) instead of verb-noun. |
| subnet | 2 | `list-subnets`, `subnets` | Yes |  |
| user | 11 | `add-user`, `change-user-password`, `disable-user`, `enable-user`, `list-users`, `login`, `logout`, `remove-user`, `show-user`, `users`, `whoami` | Mixed |  |

**Total commands accounted for: 188** (expected 188)


---

# Section 4: Symmetry Audit
| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|
| add cloud | add-cloud | remove-cloud | Yes | Yes |  |
| add credential | add-credential | remove-credential | Yes | Yes |  |
| add k8s | add-k8s | remove-k8s | Yes | Yes |  |
| add machine | add-machine | remove-machine | Yes | Yes |  |
| add model | add-model | destroy-model | No | Partial | `destroy-model` is not `remove-model`; semantic mismatch per DE013 §Grammar |
| add secret | add-secret | remove-secret | Yes | Yes |  |
| add secret-backend | add-secret-backend | remove-secret-backend | Yes | Yes |  |
| add space | add-space | remove-space | Yes | Yes |  |
| add ssh-key | add-ssh-key | remove-ssh-key | Yes | Yes |  |
| add storage | add-storage | remove-storage | Yes | Yes |  |
| add unit | add-unit | remove-unit | Yes | Yes |  |
| add user | add-user | remove-user | Yes | Yes |  |
| attach storage | attach-storage | detach-storage | Yes | Yes |  |
| consume offer | consume | remove-saas | No | No | Forward is verb-only, reverse is `remove-saas` not `unconsume` |
| create backup | create-backup | download-backup | No | No | No deletion counterpart; `download-backup` is retrieval |
| create storage-pool | create-storage-pool | remove-storage-pool | Yes | Yes |  |
| default credential | default-credential | — | No | No | No `unset-default-credential`; uses positional args to unset |
| default region | default-region | — | No | No | No `unset-default-region` |
| deploy application | deploy | remove-application | No | Partial | Per DE013 §Grammar, `deploy` should be paired with `delete-application` or `destroy-application` |
| enable command | enable-command | disable-command | Yes | Yes |  |
| enable destroy-controller | enable-destroy-controller | — | No | No | No disable counterpart; removes blocks |
| enable user | enable-user | disable-user | Yes | Yes |  |
| expose application | expose | unexpose | Yes | Yes |  |
| grant access | grant | revoke | Yes | Yes |  |
| grant cloud | grant-cloud | revoke-cloud | Yes | Yes |  |
| grant secret | grant-secret | revoke-secret | Yes | Yes |  |
| integrate applications | integrate | remove-relation | No | No | `integrate` is orphan verb; reverse is `remove-relation` |
| integrate applications (alias) | relate | remove-relation | No | No | `relate` alias also lacks symmetric reverse |
| login | login | logout | Yes | Yes |  |
| offer endpoint | offer | remove-offer | No | No | Forward is orphan verb; reverse uses `remove` |
| register controller | register | unregister | Yes | Yes |  |
| resume relation | resume-relation | suspend-relation | Yes | Yes |  |
| set constraints | set-constraints | — | No | No | No `unset-constraints`; reset via empty value |
| set firewall-rule | set-firewall-rule | — | No | No | No unset counterpart; override with empty value |
| set model-constraints | set-model-constraints | — | No | No | No unset counterpart |
| suspend relation | suspend-relation | resume-relation | Yes | Yes |  |
| trust application | trust | — | No | No | No `untrust` command; `trust` toggles or sets to true |

### Missing Reverse Operations
- `deploy` → no `destroy-application` or `delete-application`
- `add-model` → no `remove-model` (only `destroy-model`)
- `consume` → no `unconsume`
- `offer` → no `unoffer`
- `trust` → no `untrust`
- `set-constraints` → no `unset-constraints`
- `set-firewall-rule` → no `unset-firewall-rule`
- `default-credential` / `default-region` → no `unset-default-*`


---

# Section 5: Confusion-Pair Audit
| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|
| destroy-controller | kill-controller | synonym verbs | high | `destroy` attempts graceful teardown; `kill` forcefully terminates immediately. |
| remove-application | destroy-model | scope ambiguity | medium | `remove-application` deletes one app; `destroy-model` deletes the entire model including all apps. |
| remove-unit | remove-application | scope ambiguity | medium | `remove-unit` scales down; `remove-application` deletes the whole app. |
| remove-machine | remove-unit | scope ambiguity | medium | `remove-machine` removes infrastructure; `remove-unit` removes workload instances. |
| config | model-config | scope ambiguity | high | `config` targets an application; `model-config` targets the model. |
| model-config | controller-config | scope ambiguity | high | `model-config` is per-model; `controller-config` is per-controller. |
| constraints | model-constraints | scope ambiguity | medium | `constraints` is per-application; `model-constraints` is per-model. |
| set-constraints | set-model-constraints | scope ambiguity | medium | Same scope split as above. |
| show-cloud | show-credential | functional overlap | low | Both display cloud metadata, but credentials are a sub-resource. |
| show-controller | show-model | functional overlap | low | Both show metadata, but at different hierarchy levels. |
| status | show-model | functional overlap | medium | `status` gives runtime state; `show-model` gives static config. |
| integrate | relate | synonym verbs | high | `relate` is an alias for `integrate`; users may not know which is canonical. |
| expose | unexpose | naming collision | low | Direct opposites; naming is actually clear. |
| grant | grant-cloud | scope ambiguity | medium | `grant` is for model/controller/offer; `grant-cloud` is specifically for clouds. |
| revoke | revoke-cloud | scope ambiguity | medium | Same scope split as above. |
| add-cloud | add-k8s | functional overlap | medium | `add-k8s` is a special case of adding a Kubernetes cloud. |
| update-cloud | update-k8s | functional overlap | low | Similar overlap as add. |
| remove-cloud | remove-k8s | functional overlap | low | Similar overlap as add. |
| create-storage-pool | pool-create | synonym verbs | high | `pool-create` is an orphan command with the same semantics as `create-storage-pool`. |
| remove-storage-pool | pool-delete | synonym verbs | high | `pool-delete` is an orphan command equivalent to `remove-storage-pool`. |
| update-storage-pool | pool-update | synonym verbs | high | `pool-update` is an orphan command equivalent to `update-storage-pool`. |
| list-storage-pools | pool-list | synonym verbs | high | `pool-list` is an orphan command equivalent to `list-storage-pools`. |
| list-actions | actions | naming collision | low | `actions` is canonical, `list-actions` is its alias. |
| list-clouds | clouds | naming collision | low | Canonical vs alias. |
| list-controllers | controllers | naming collision | low | Canonical vs alias. |
| list-credentials | credentials | naming collision | low | Canonical vs alias. |
| list-machines | machines | naming collision | low | Canonical vs alias. |
| list-models | models | naming collision | low | Canonical vs alias. |
| list-offers | offers | naming collision | low | Canonical vs alias. |
| list-operations | operations | naming collision | low | Canonical vs alias. |
| list-regions | regions | naming collision | low | Canonical vs alias. |
| list-resources | resources | naming collision | low | Canonical vs alias. |
| list-secret-backends | secret-backends | naming collision | low | Canonical vs alias. |
| list-secrets | secrets | naming collision | low | Canonical vs alias. |
| list-spaces | spaces | naming collision | low | Canonical vs alias. |
| list-ssh-keys | ssh-keys | naming collision | low | Canonical vs alias. |
| list-storage | storage | naming collision | low | Canonical vs alias. |
| list-storage-pools | storage-pools | naming collision | low | Canonical vs alias. |
| list-subnets | subnets | naming collision | low | Canonical vs alias. |
| list-users | users | naming collision | low | Canonical vs alias. |
| show-credentials | show-credential | naming collision | low | Singular canonical vs plural alias. |
| update-credentials | update-credential | naming collision | low | Singular canonical vs plural alias. |
| set-default-credentials | default-credential | naming collision | low | Alias vs canonical. |
| set-default-region | default-region | naming collision | low | Alias vs canonical. |
| resolved | resolve | synonym verbs | medium | `resolve` is an alias for `resolved`; both mark errors resolved. |
| debug-hooks | debug-hook | naming collision | low | Plural canonical vs singular alias. |


---

# Section 6: Pattern Classification and Recommendations
## Pattern Classification
- **Primary grouping pattern**: **Flat verb-noun** with a large number of top-level commands.
- **Depth**: Strictly depth-1 (all commands are direct children of `juju`).
- **Style**: Mixed — most commands follow `verb-noun` (e.g., `add-cloud`), but a significant minority are noun-only (`models`, `status`), noun-noun (`model-config`), or verb-only orphans (`deploy`, `integrate`).
- **Alias strategy**: Plural nouns are aliases for `list-*` variants (e.g., `clouds` → `list-clouds`), and some past-tense verbs (`resolved`) have base-form aliases (`resolve`).

## Discoverability Assessment
- **Predicted path**: A new user looking to 'see all machines' will likely try `juju machines` or `juju list machines`. The CLI supports `juju machines` (alias), which is good.
- **Friction points**:
  - `model-config` vs `controller-config` vs `config` are not namespaced; users must know the scope from docs.
  - `destroy-model` vs `remove-application` vs `kill-controller` use different teardown verbs, making it hard to guess the right one.
  - Pool management uses `pool-create` instead of `create-storage-pool`, breaking the verb-noun convention.
  - `trust` and `consume` are verb-only; there is no `application-trust` or `model-consume` to signal the target.

## Ecosystem Comparison
| Tool | Pattern | Depth | Notes |
|---|---|---|---|
| `kubectl` | verb-noun | 2+ | `kubectl get pods`, clear CRUD verbs |
| `aws` CLI | noun-verb | 2 | `aws ec2 describe-instances`, grouped by service |
| `snap` | verb-primary | 1-2 | `snap install`, `snap info`; secondary objects use `snap services` |
| `juju` | mixed flat | 1 | No service grouping; all commands at top level |

## Recommendations
1. **Consolidate storage-pool verbs** — `pool-create`, `pool-delete`, `pool-update`, `pool-list` should be deprecated in favor of `create-storage-pool`, `remove-storage-pool`, `update-storage-pool`, `list-storage-pools` to restore verb-noun consistency (per DE013 §Grammar).
   - *Deprecation plan*: Minor release adds deprecation warnings to `pool-*` commands; next major release removes them.
   - *Backward compat*: High — aliases can redirect `pool-create` → `create-storage-pool` for one cycle.
2. **Introduce `delete-application` as alias for `remove-application`** — `deploy` → `delete-application` mirrors `create` → `delete` better than `remove`. Per DE013, `delete` is the standard teardown verb for secondary objects.
   - *Deprecation plan*: Keep `remove-application` as canonical for now; add `delete-application` alias with no warning. Evaluate switching canonical name in a future major release.
3. **Add `untrust` command** — `trust` currently has no inverse. Add `untrust` that sets trust to false, symmetric with `expose`/`unexpose`.
   - *Backward compat*: None — new command only.
4. **Rename `kill-controller` to `force-destroy-controller`** — `kill` is a strong, irreversible verb without a linguistic inverse in this domain. `force-destroy-controller` makes its relationship to `destroy-controller` explicit.
   - *Deprecation plan*: Minor release emits warning on `kill-controller`; next major release removes it. Alias `kill-controller` → `force-destroy-controller` for one cycle.
   - *Migration cost*: Medium — scripts using `kill-controller` need updating.
5. **Namespace config commands under `config`** — Instead of `model-config`, `controller-config`, and `config`, consider `config model`, `config controller`, `config application`. This reduces top-level clutter and aligns with DE013's allowance for one sublevel.
   - *Backward compat*: Breaking. Keep top-level aliases for at least one major cycle.
   - *Migration cost*: High — documentation, scripts, and muscle memory all affected.
6. **Standardize confirmation bypass flags** — Replace `--no-prompt` with `--yes` or `-y` (per DE013 Flags) and ensure all destructive commands support it consistently.
   - *Backward compat*: Add `--yes` as alias for `--no-prompt`; deprecate `--no-prompt` over one cycle.
