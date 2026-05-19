# Juju Command Set Shape Analysis

## Shape Summary

Juju's CLI follows a flat command structure with verb-noun naming as the dominant pattern. All 150+ commands are direct children of the `juju` root, with no nested subcommands. The naming primarily uses hyphenated verb-noun pairs (`add-model`, `remove-application`) with some standalone verbs (`bootstrap`, `deploy`, `integrate`). The structure prioritizes discoverability through consistent naming conventions and extensive alias support for backward compatibility.

## Key Findings

1. **Verb inconsistency in removal operations**: `remove-*`, `destroy-*`, and `kill-*` are all used for deletion operations
2. **Incomplete CRUD symmetry**: Many nouns have `add-*` but no matching `remove-*` (or vice versa)
3. **Semantic overlap**: `deploy` and `add-application` conceptually overlap
4. **Orphan commands**: Several commands don't decompose into verb-noun (bootstrap, integrate, ssh, whoami)
5. **Legacy aliases**: Heavy alias burden from historical naming changes

---

## Section 1: Verb-Noun Decomposition Matrix

| Verb | cloud | model | controller | application | unit | machine | user | secret | storage | space | relation | credential | offer | backend | resource |
|------|-------|-------|------------|-------------|------|---------|------|--------|---------|-------|----------|------------|-------|---------|----------|
| add | ✓ | ✓ | — | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | — | ✓ | — |
| remove | ✓ | ✓ | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — |
| show | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | — |
| list | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | ✓ |
| update | ✓ | — | — | — | — | — | — | ✓ | — | — | — | ✓ | — | ✓ | ✓ |
| create | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — |
| destroy | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | — | — |
| kill | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — |
| grant | — | ✓ | — | — | — | — | — | ✓ | — | — | — | ✓ | — | — | — |
| revoke | — | ✓ | — | — | — | — | — | ✓ | — | — | — | ✓ | — | — | — |
| enable | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — |
| disable | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — |

### Incomplete CRUD Sets

| Noun | add | remove | show | list | Status |
|------|-----|--------|------|------|--------|
| cloud | ✓ | ✓ | ✓ | ✓ | Complete |
| model | ✓ | ✓ | ✓ | ✓ | Complete |
| controller | ✗ | ✓ | ✓ | ✓ | Incomplete (use register) |
| application | ✗ | ✓ | ✓ | ✓ | Incomplete (use deploy) |
| unit | ✓ | ✓ | ✓ | ✓ | Complete |
| machine | ✓ | ✓ | ✓ | ✓ | Complete |
| user | ✓ | ✓ | ✓ | ✓ | Complete |
| secret | ✓ | ✓ | ✓ | ✓ | Complete |
| storage | ✓ | ✓ | ✓ | ✓ | Complete |
| space | ✓ | ✓ | ✓ | ✓ | Complete |
| secret-backend | ✓ | ✓ | ✓ | ✓ | Complete |

### Verb Inconsistencies for Deletion

| Verb | Use Case | Examples |
|------|----------|----------|
| `remove-*` | Standard deletion of resources | remove-application, remove-unit, remove-machine, remove-user, remove-secret, remove-storage, remove-cloud, remove-credential, remove-space, remove-offer, remove-relation, remove-saas, remove-ssh-key, remove-storage-pool |
| `destroy-*` | Destructive cascading deletion | destroy-model, destroy-controller |
| `kill-*` | Force termination without cleanup | kill-controller |

### Orphan Commands (No Clean Verb-Noun Decomposition)

| Command | Category | Reason |
|---------|----------|--------|
| `bootstrap` | Infrastructure | Creates controller (not add-controller) |
| `deploy` | Application | Creates application (not add-application) |
| `integrate` | Integration | Creates relation (not add-relation) |
| `expose` | Network | Sets network visibility |
| `unexpose` | Network | Removes network visibility |
| `trust` | Security | Grants credential access |
| `ssh` | Access | Remote shell |
| `scp` | Access | File transfer |
| `whoami` | Identity | Shows current context |
| `switch` | Context | Changes active model/controller |
| `login` | Auth | Authenticates to controller |
| `logout` | Auth | Ends session |
| `register` | Controller | Adds controller connection |
| `unregister` | Controller | Removes controller connection |
| `refresh` | Application | Upgrades charm |
| `scale-application` | Application | Scales K8s workloads |
| `exec` | Execution | Runs commands on targets |
| `run` | Action | Executes charm action |
| `status` | Observation | Shows model state |

---

## Section 2: Verb Taxonomy and Aspect Classification

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|------|--------------|--------|------------|-------------|--------------|
| add | lifecycle | telic | yes | remove | add-model, add-user, add-machine, add-unit, add-secret, add-storage, add-space, add-cloud, add-credential, add-k8s, add-secret-backend |
| remove | lifecycle | telic | yes | add | remove-model, remove-user, remove-machine, remove-unit, remove-secret, remove-storage, remove-space, remove-cloud, remove-credential, remove-k8s, remove-secret-backend, remove-application, remove-relation, remove-offer, remove-saas |
| destroy | lifecycle | telic | no | — | destroy-model, destroy-controller |
| kill | lifecycle | telic | no | — | kill-controller |
| deploy | lifecycle | telic | partial | remove-application | deploy |
| bootstrap | lifecycle | telic | no | — | bootstrap |
| show | observation | atelic | no | — | show-model, show-controller, show-application, show-unit, show-machine, show-secret, show-storage, show-space, show-cloud, show-credential, show-offer, show-secret-backend, show-user, show-action, show-operation, show-task |
| list | observation | atelic | no | — | models, controllers, applications, units, machines, users, secrets, storage, spaces, clouds, credentials, offers, secret-backends, actions, operations, resources, subnets, ssh-keys, storage-pools, firewall-rules, regions |
| status | observation | atelic | no | — | status |
| whoami | observation | punctual | no | — | whoami |
| find | observation | atelic | no | — | find, find-offers |
| info | observation | atelic | no | — | info |
| update | mutation | telic | yes | — | update-cloud, update-credential, update-secret, update-secret-backend, update-storage-pool |
| config | mutation | atelic | no | — | config, model-config, controller-config |
| set | mutation | punctual | yes | — | set-constraints, set-model-constraints, set-firewall-rule, set-credential |
| refresh | mutation | telic | no | — | refresh |
| scale | mutation | telic | no | — | scale-application |
| upgrade | mutation | telic | no | — | upgrade-model, upgrade-controller |
| grant | access | punctual | yes | revoke | grant, grant-cloud, grant-secret |
| revoke | access | punctual | yes | grant | revoke, revoke-cloud, revoke-secret |
| enable | access | punctual | yes | disable | enable-command, enable-user, enable-destroy-controller |
| disable | access | punctual | yes | enable | disable-command, disable-user |
| login | access | punctual | yes | logout | login |
| logout | access | punctual | yes | login | logout |
| register | access | punctual | yes | unregister | register |
| unregister | access | punctual | yes | register | unregister |
| expose | transfer | punctual | yes | unexpose | expose |
| unexpose | transfer | punctual | yes | expose | unexpose |
| trust | transfer | punctual | no | — | trust |
| bind | transfer | punctual | no | — | bind |
| integrate | transfer | telic | yes | remove-relation | integrate |
| suspend | transfer | punctual | yes | resume | suspend-relation |
| resume | transfer | punctual | yes | suspend | resume-relation |
| offer | transfer | telic | yes | remove-offer | offer |
| consume | transfer | telic | yes | remove-saas | consume |
| attach | transfer | punctual | yes | detach | attach-storage, attach-resource |
| detach | transfer | punctual | yes | attach | detach-storage |
| import | transfer | telic | no | — | import-filesystem, import-ssh-key |
| run | execution | atelic | no | — | run |
| exec | execution | atelic | no | — | exec |
| cancel | execution | punctual | no | — | cancel-task |
| resolve | execution | punctual | no | — | resolved |
| ssh | execution | atelic | no | — | ssh |
| scp | execution | telic | no | — | scp |
| migrate | migration | telic | no | — | migrate |
| download | migration | telic | no | — | download, download-backup |
| switch | execution | punctual | no | — | switch |
| create | lifecycle | telic | yes | remove | create-backup, create-storage-pool |
| rename | mutation | punctual | no | — | rename-space |
| move | mutation | punctual | no | — | move-to-space |
| reload | mutation | punctual | no | — | reload-spaces |
| retry | execution | punctual | no | — | retry-provisioning |
| sync | transfer | telic | no | — | sync-agent-binary |
| export | transfer | telic | no | — | export-bundle |
| diff | observation | atelic | no | — | diff-bundle |
| dashboard | observation | punctual | no | — | dashboard |
| version | observation | punctual | no | — | version |

---

## Section 3: Semantic Domain Clustering

| Domain | Count | Commands | Naming Consistent? | Notes |
|--------|-------|----------|-------------------|-------|
| controller | 12 | bootstrap, destroy-controller, kill-controller, show-controller, controllers, register, unregister, enable-destroy-controller, upgrade-controller, controller-config, add-model, models | Partial | `bootstrap` breaks verb-noun pattern; uses `controllers` plural noun |
| model | 16 | add-model, destroy-model, show-model, models, model-config, model-defaults, model-constraints, set-model-constraints, upgrade-model, migrate, switch, export-bundle, grant, revoke, retry-provisioning, set-credential | Partial | `destroy-model` vs `remove-*` inconsistency; `switch` is orphan |
| application | 17 | deploy, remove-application, show-application, applications, config, constraints, set-constraints, refresh, expose, unexpose, trust, bind, scale-application, integrate, remove-relation, show-unit, units | Partial | `deploy` not `add-application`; `integrate` creates relation not application |
| unit | 6 | add-unit, remove-unit, show-unit, units, resolved, exec | Yes | `resolved` is orphan but domain-consistent |
| machine | 6 | add-machine, remove-machine, show-machine, machines, ssh, retry-provisioning | Yes | `ssh` connects to machines; domain-consistent |
| user | 11 | add-user, remove-user, show-user, users, change-user-password, enable-user, disable-user, grant, revoke, login, logout | Yes | Auth commands logically grouped with users |
| secret | 8 | add-secret, remove-secret, show-secret, secrets, update-secret, grant-secret, revoke-secret, model-secret-backend | Yes | Complete CRUD with access control |
| secret-backend | 6 | add-secret-backend, remove-secret-backend, show-secret-backend, secret-backends, update-secret-backend, model-secret-backend | Yes | Complete CRUD; `model-secret-backend` is setter |
| storage | 10 | add-storage, remove-storage, show-storage, storage, attach-storage, detach-storage, import-filesystem, create-storage-pool, remove-storage-pool, storage-pools | Partial | Pool commands use `create/remove` not `add/remove`; `import-filesystem` is orphan |
| space | 7 | add-space, remove-space, show-space, spaces, rename-space, move-to-space, reload-spaces | Yes | Extra `rename` and `move` verbs are domain-appropriate |
| subnet | 2 | subnets, list-subnets | No | Only list, no CRUD |
| relation | 4 | integrate, remove-relation, suspend-relation, resume-relation | Partial | `integrate` is orphan; other verbs consistent |
| credential | 8 | add-credential, remove-credential, show-credential, credentials, update-credential, set-credential, default-credential, set-default-credential | Partial | `default-credential` is getter, `set-default-credential` is setter |
| cloud | 8 | add-cloud, remove-cloud, show-cloud, clouds, update-cloud, add-k8s, remove-k8s, update-k8s | Partial | K8s commands separate from generic cloud commands |
| offer/crossmodel | 6 | offer, remove-offer, show-offer, offers, find-offers, consume | Partial | `consume` is orphan; `remove-saas` pairs with `consume` |
| ssh-key | 4 | add-ssh-key, remove-ssh-key, ssh-keys, import-ssh-key | Yes | Complete CRUD for keys |
| backup | 3 | create-backup, download-backup, backups | Partial | `create` not `add`; no remove |
| resource | 4 | resources, charm-resources, attach-resource, list-resources | Partial | `attach-resource` is orphan; aliases overlap |
| action | 8 | actions, show-action, run, cancel-task, operations, show-operation, show-task, exec | Partial | `run` and `exec` are orphans in action domain |
| charm/charmhub | 4 | find, info, download, refresh | No | No verb-noun pattern; standalone verbs |

**Total Commands Analyzed: 150+**

---

## Section 4: Symmetry Audit

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|-----------|-----------------|-----------------|-------------------|---------------------|-------|
| Model creation | add-model | destroy-model | No | No | `add` vs `destroy`; removal is cascading |
| Controller creation | bootstrap | destroy-controller | No | No | `bootstrap` vs `destroy`; removal cascades |
| Controller force-delete | — | kill-controller | — | No | No forward equivalent |
| Application creation | deploy | remove-application | No | Yes | `deploy` vs `remove`; non-symmetric verbs |
| Unit addition | add-unit | remove-unit | Yes | Yes | Symmetric naming and behavior |
| Machine addition | add-machine | remove-machine | Yes | Yes | Symmetric naming and behavior |
| User creation | add-user | remove-user | Yes | Yes | Symmetric naming and behavior |
| Secret creation | add-secret | remove-secret | Yes | Yes | Symmetric naming and behavior |
| Secret backend | add-secret-backend | remove-secret-backend | Yes | Yes | Symmetric naming and behavior |
| Storage addition | add-storage | remove-storage | Yes | Yes | Symmetric naming and behavior |
| Space addition | add-space | remove-space | Yes | Yes | Symmetric naming and behavior |
| Cloud addition | add-cloud | remove-cloud | Yes | Yes | Symmetric naming and behavior |
| Credential addition | add-credential | remove-credential | Yes | Yes | Symmetric naming and behavior |
| Relation creation | integrate | remove-relation | No | Yes | `integrate` vs `remove-relation` |
| Relation suspension | suspend-relation | resume-relation | Yes | Yes | Symmetric naming and behavior |
| Offer creation | offer | remove-offer | No | Yes | `offer` (verb only) vs `remove-offer` |
| SAAS consumption | consume | remove-saas | No | Yes | `consume` vs `remove-saas` |
| Network exposure | expose | unexpose | Yes | Yes | Symmetric naming and behavior |
| Storage attachment | attach-storage | detach-storage | Yes | Yes | Symmetric naming and behavior |
| Access grant | grant | revoke | Yes | Yes | Symmetric naming and behavior |
| Cloud access | grant-cloud | revoke-cloud | Yes | Yes | Symmetric naming and behavior |
| Secret access | grant-secret | revoke-secret | Yes | Yes | Symmetric naming and behavior |
| User enable | enable-user | disable-user | Yes | Yes | Symmetric naming and behavior |
| Command block | enable-command | disable-command | Yes | Yes | Symmetric naming and behavior |
| Login | login | logout | Yes | Yes | Symmetric naming and behavior |
| Registration | register | unregister | Yes | Yes | Symmetric naming and behavior |
| SSH key import | import-ssh-key | remove-ssh-key | No | Yes | `import` vs `remove` |
| Storage pool | create-storage-pool | remove-storage-pool | No | Yes | `create` vs `remove` |

### Missing Reverse Operations

| Forward Command | Missing Reverse | Notes |
|-----------------|-----------------|-------|
| `bootstrap` | No `decommission` | Must use `destroy-controller` |
| `deploy` | No `undeploy` | Must use `remove-application` |
| `integrate` | No `disintegrate` | Must use `remove-relation` |
| `consume` | No `unconsume` | Must use `remove-saas` |
| `trust` | No `untrust` | Cannot revoke trust once granted |
| `import-ssh-key` | No `export-ssh-key` | Removal is the reverse |
| `create-backup` | No `delete-backup` | Backups auto-expire or manual deletion |
| `import-filesystem` | No `export-filesystem` | Storage remains in model |

---

## Section 5: Confusion-Pair Audit

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|--------------|----------------|----------------|
| `deploy` | `add-application` | synonym verbs | high | `deploy` creates from charm; `add-application` does not exist |
| `destroy-model` | `remove-model` | synonym verbs | high | `destroy-model` cascades; no `remove-model` exists |
| `destroy-controller` | `remove-controller` | synonym verbs | high | `destroy-controller` cascades; use `unregister` to disconnect |
| `kill-controller` | `destroy-controller` | force variants | high | `kill` is force; `destroy` attempts cleanup |
| `config` | `model-config` | scope ambiguity | high | `config` is application-level; `model-config` is model-level |
| `constraints` | `set-constraints` | verb inconsistency | medium | `constraints` shows; `set-constraints` modifies |
| `model-constraints` | `set-model-constraints` | verb inconsistency | medium | Same pattern as application constraints |
| `integrate` | `relate` | synonym alias | medium | `relate` is deprecated alias for `integrate` |
| `resolved` | `resolve` | synonym alias | low | `resolve` is deprecated alias for `resolved` |
| `ssh-keys` | `list-ssh-keys` | list-* alias | low | `list-*` is deprecated alias form |
| `models` | `list-models` | list-* alias | low | Same command, different names |
| `controllers` | `list-controllers` | list-* alias | low | Same command, different names |
| `users` | `list-users` | list-* alias | low | Same command, different names |
| `storages` | `list-storage` | list-* alias | low | `storage` is singular noun, `list-*` is alias |
| `spaces` | `list-spaces` | list-* alias | low | Same command, different names |
| `secrets` | `list-secrets` | list-* alias | low | Same command, different names |
| `run` | `exec` | functional overlap | medium | `run` executes actions; `exec` runs commands on targets |
| `run` | `action` | functional overlap | low | `run` executes a specific action; `actions` lists available actions |
| `ssh` | `exec` | functional overlap | low | `ssh` opens shell; `exec` runs commands on multiple targets |
| `refresh` | `upgrade-model` | functional overlap | low | `refresh` upgrades charm; `upgrade-model` upgrades agent version |
| `create-backup` | `download-backup` | functional overlap | low | `create` generates; `download` retrieves existing |
| `add-secret` | `update-secret` | CRUD overlap | low | `add` creates new; `update` modifies existing |
| `attach-resource` | `resources` | functional overlap | low | `attach-resource` uploads; `resources` lists |
| `add-storage` | `attach-storage` | functional overlap | high | `add-storage` creates new; `attach-storage` connects existing |
| `remove-storage` | `detach-storage` | functional overlap | high | `remove-storage` deletes; `detach-storage` disconnects without deleting |
| `show-cloud` | `clouds` | granularity | low | `show-cloud` shows one; `clouds` lists all |
| `show-model` | `models` | granularity | low | `show-model` shows one; `models` lists all |
| `show-user` | `users` | granularity | low | `show-user` shows one; `users` lists all |
| `show-controller` | `controllers` | granularity | low | `show-controller` shows one; `controllers` lists all |
| `show-application` | `applications` | granularity | low | `show-application` shows one; `applications` lists all (via status) |
| `show-unit` | `units` | granularity | low | `show-unit` shows one; `units` lists all (via status) |
| `show-machine` | `machines` | granularity | low | `show-machine` shows one; `machines` lists all (via status) |
| `show-secret` | `secrets` | granularity | low | `show-secret` shows one; `secrets` lists all |
| `show-offer` | `offers` | granularity | low | `show-offer` shows one; `offers` lists all |
| `show-storage` | `storage` | granularity | low | `show-storage` shows one; `storage` lists all |
| `show-space` | `spaces` | granularity | low | `show-space` shows one; `spaces` lists all |
| `show-credential` | `credentials` | granularity | low | `show-credential` shows one; `credentials` lists all |
| `show-secret-backend` | `secret-backends` | granularity | low | `show-secret-backend` shows one; `secret-backends` lists all |
| `show-action` | `actions` | granularity | low | `show-action` shows action spec; `actions` lists available actions |
| `show-operation` | `operations` | granularity | low | `show-operation` shows result; `operations` lists history |
| `show-task` | `operations` | granularity | low | `show-task` shows task result; `operations` lists operations |
| `find` | `info` | functional overlap | low | `find` searches charms; `info` shows details for one charm |
| `register` | `login` | functional overlap | medium | `register` adds controller connection; `login` authenticates to existing |
| `set-default-credential` | `default-credential` | verb inconsistency | medium | `set-*` vs bare noun; getter uses bare noun |
| `set-default-region` | `default-region` | verb inconsistency | medium | Same pattern as credential |
| `set-credential` | `update-credential` | functional overlap | high | `set-credential` sets model credential; `update-credential` updates stored credential |
| `model-secret-backend` | `update-secret-backend` | functional overlap | medium | `model-secret-backend` sets model default; `update-secret-backend` modifies backend definition |

---

## Section 6: Pattern Classification and Recommendations

### Pattern Classification

**Primary Pattern**: Verb-Noun with hyphen separator
- Examples: `add-model`, `remove-application`, `show-unit`, `update-cloud`
- Frequency: ~80% of commands
- Style: Consistent, discoverable

**Secondary Pattern**: Standalone Verbs
- Examples: `bootstrap`, `deploy`, `integrate`, `ssh`, `scp`
- Frequency: ~15% of commands
- Style: Action-oriented, memorable but less discoverable

**Tertiary Pattern**: Plural Nouns (list commands)
- Examples: `models`, `controllers`, `users`, `secrets`
- Frequency: ~5% of commands
- Style: Implicit `list` verb, concise

**Legacy Pattern**: `list-*` aliases
- Examples: `list-models`, `list-users`, `list-secrets`
- Status: Deprecated but retained for compatibility

### Discoverability Assessment

**Strengths**:
1. Consistent verb-noun pattern for most commands
2. Help output organized by functional category
3. Tab completion works for all commands
4. `juju help commands` shows complete list

**Weaknesses**:
1. Orphan verbs (`bootstrap`, `deploy`) break pattern
2. Verb synonyms for deletion (`remove`, `destroy`, `kill`)
3. Scope ambiguity between model-level and controller-level config
4. No subcommand grouping for related commands

### Ecosystem Comparison

| Aspect | Juju | kubectl | docker | aws |
|--------|------|---------|--------|-----|
| Structure | Flat | Nested (`get`, `describe`) | Nested | Nested |
| Naming | Verb-noun | Verb-noun | Verb-noun | Noun-verb |
| Aliases | Many | Few | Few | Few |
| Subcommands | None | Yes (`kubectl get pods`) | Yes (`docker image ls`) | Yes (`aws s3 ls`) |

### Recommendations

**High Priority**:
1. Standardize deletion verbs: Use `remove-*` for all resource deletion; deprecate `destroy-*` and `kill-*`
2. Add `undeploy` as alias for `remove-application` to complete deployment symmetry
3. Create `add-application` as alias for `deploy` to complete CRUD symmetry
4. Document verb semantics clearly in help output

**Medium Priority**:
1. Unify config command naming: `config` → `application-config` for clarity
2. Add `untrust` command to allow credential access revocation
3. Deprecate `list-*` aliases with clear timeline
4. Add `delete-backup` command for backup lifecycle management

**Low Priority**:
1. Consider nested structure for future major version (`juju model list`, `juju application deploy`)
2. Unify show/list commands into single pattern
3. Add completion hints for partial command matches

### Tradeoffs

**Backward Compatibility**: All naming changes must preserve aliases for existing scripts. Migration path: new name + deprecated alias for 2 releases, then remove alias.

**Scriptability**: Verb-noun pattern aids scripting; flat structure simplifies parsing. Nested structure would require script updates.

**Human vs Machine**: Current flat structure is optimized for humans. Nested structure better for programmatic discovery.

**Migration Cost**: High. Thousands of scripts depend on current command names. Any rename requires extensive deprecation period.