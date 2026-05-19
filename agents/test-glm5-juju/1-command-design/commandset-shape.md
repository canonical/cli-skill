# Juju CLI Command Set Shape Analysis

## Overview

This document analyzes the Juju CLI command set shape across six dimensions: verb-noun decomposition, verb taxonomy, semantic domain clustering, symmetry, confusion pairs, and pattern classification.

## Section 1: Verb-Noun Decomposition Matrix

Commands decomposed into verb × noun combinations. Grid shows which verb-noun pairs exist.

### Verb Columns

| Verb | application | cloud | controller | credential | k8s | machine | model | offer | relation | saas | secret | secret-backend | space | ssh-key | storage | storage-pool | subnet | unit | user | backup | firewall-rule | filesystem | task | operation | action | bundle | charm | resource | credential-cloud |
|------|:-----------:|:-----:|:----------:|:----------:|:---:|:-------:|:-----:|:-----:|:-------:|:----:|:------:|:--------------:|:-----:|:-------:|:-------:|:------------:|:------:|:----:|:----:|:------:|:--------------:|:----------:|:----:|:---------:|:------:|:------:|:-----:|:--------:|:----------------|
| add | ✓ | ✓ | — | ✓ | ✓ | ✓ | ✓ | — | — | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — |
| remove | ✓ | ✓ | — | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — |
| show | ✓ | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | — | — | ✓ | ✓ | ✓ | — | ✓ | — | — | ✓ | ✓ | — | — | — | ✓ | ✓ | ✓ | — | — | — | — |
| list | — | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | — | — | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | — | ✓ | — | ✓ | — | — | ✓ | ✓ | — | — | ✓ | — |
| destroy | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| kill | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| deploy | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — |
| integrate | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| suspend | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| resume | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| expose | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| unexpose | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| refresh | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| scale | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| bind | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| trust | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| config | ✓ | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| constraints | ✓ | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| defaults | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| grant | — | — | — | — | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ |
| revoke | — | — | — | — | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ |
| update | — | ✓ | — | ✓ | ✓ | — | — | — | — | — | ✓ | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — |
| create | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | — |
| download | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | ✓ | — | — | — |
| find | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — |
| info | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — |
| run | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| exec | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| cancel | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — |
| offer | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| consume | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| attach | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — |
| detach | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| import | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — |
| bootstrap | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| register | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| unregister | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| login | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| logout | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| enable | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — |
| disable | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — |
| upgrade | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| sync | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| migrate | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| export | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| diff | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — |
| resolved | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — |
| switch | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| set | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — |
| move | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| rename | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| reload | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| autoload | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| default | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| retry | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| debug | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| status | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| whoami | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |

### Orphan Commands (No Clear Verb-Noun Decomposition)

| Command | Notes |
|---------|-------|
| `bootstrap` | Creates controller + initial model, unique operation |
| `whoami` | Shows current context, identity operation |
| `switch` | Changes context, navigation operation |
| `debug-log` | Continuous log stream, observation operation |
| `dashboard` | Opens UI, navigation operation |
| `ssh` | Interactive shell, access operation |
| `scp` | File transfer, access operation |
| `exec` | Remote execution, action operation |
| `run` | Action execution, operation |
| `help` | Documentation, meta-operation |

### Incomplete CRUD Sets

| Noun | Create (add/deploy) | Read (show/list) | Update (config/update) | Delete (remove/destroy) |
|------|:-------------------:|:----------------:|:----------------------:|:-----------------------:|
| application | ✓ (deploy) | ✓ (show-application) | ✓ (config, refresh) | ✓ (remove-application) |
| unit | ✓ (add-unit) | ✓ (show-unit) | — | ✓ (remove-unit) |
| machine | ✓ (add-machine) | ✓ (show-machine) | — | ✓ (remove-machine) |
| model | ✓ (add-model) | ✓ (show-model) | ✓ (model-config) | ✓ (destroy-model) |
| controller | ✓ (bootstrap) | ✓ (show-controller) | ✓ (controller-config) | ✓ (destroy-controller) |
| cloud | ✓ (add-cloud) | ✓ (show-cloud) | ✓ (update-cloud) | ✓ (remove-cloud) |
| credential | ✓ (add-credential) | ✓ (show-credential) | ✓ (update-credential) | ✓ (remove-credential) |
| space | ✓ (add-space) | ✓ (show-space) | ✓ (rename-space) | ✓ (remove-space) |
| secret | ✓ (add-secret) | ✓ (show-secret) | ✓ (update-secret) | ✓ (remove-secret) |
| secret-backend | ✓ (add-secret-backend) | ✓ (show-secret-backend) | ✓ (update-secret-backend) | ✓ (remove-secret-backend) |
| storage | ✓ (add-storage) | ✓ (show-storage) | — | ✓ (remove-storage) |
| storage-pool | ✓ (create-storage-pool) | ✓ (storage-pools) | ✓ (update-storage-pool) | ✓ (remove-storage-pool) |
| user | ✓ (add-user) | ✓ (show-user) | ✓ (change-user-password) | ✓ (remove-user) |
| offer | ✓ (offer) | ✓ (show-offer) | — | ✓ (remove-offer) |
| relation | ✓ (integrate) | — | ✓ (suspend/resume) | ✓ (remove-relation) |
| firewall-rule | — | ✓ (firewall-rules) | ✓ (set-firewall-rule) | — |

## Section 2: Verb Taxonomy and Aspect Classification

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|------|--------------|--------|------------|-------------|--------------|
| add | lifecycle | telic | yes | remove | add-unit, add-model, add-cloud |
| remove | lifecycle | telic | yes | add | remove-unit, remove-model, remove-cloud |
| deploy | lifecycle | telic | partial | remove-application | deploy mysql |
| destroy | lifecycle | telic | no | — | destroy-model, destroy-controller |
| kill | lifecycle | telic | no | — | kill-controller |
| bootstrap | lifecycle | telic | no | — | bootstrap aws |
| create | lifecycle | telic | yes | remove | create-backup, create-storage-pool |
| show | observation | atelic | no | — | show-application, show-model |
| list | observation | atelic | no | — | models, controllers, clouds |
| status | observation | atelic | no | — | status |
| config | mutation | atelic | yes | — | config mysql, model-config |
| update | mutation | atelic | yes | — | update-cloud, update-secret |
| refresh | mutation | atelic | no | — | refresh mysql |
| upgrade | mutation | telic | partial | — | upgrade-model, upgrade-controller |
| set | mutation | atelic | yes | — | set-constraints, set-firewall-rule |
| grant | access | telic | yes | revoke | grant admin my-model |
| revoke | access | telic | yes | grant | revoke admin my-model |
| enable | access | punctual | yes | disable | enable-user, enable-command |
| disable | access | punctual | yes | enable | disable-user, disable-command |
| login | access | punctual | yes | logout | login mycontroller |
| logout | access | punctual | yes | login | logout |
| register | access | punctual | yes | unregister | register mycontroller |
| unregister | access | punctual | yes | register | unregister mycontroller |
| trust | access | punctual | no | — | trust mysql |
| integrate | transfer | telic | yes | remove-relation | integrate wordpress mysql |
| suspend | transfer | punctual | yes | resume | suspend-relation 1 |
| resume | transfer | punctual | yes | suspend | resume-relation 1 |
| expose | transfer | telic | yes | unexpose | expose mysql |
| unexpose | transfer | telic | yes | expose | unexpose mysql |
| offer | transfer | telic | yes | remove-offer | offer mysql:db |
| consume | transfer | telic | yes | remove-saas | consume remote.mysql |
| attach | transfer | telic | yes | detach | attach-storage storage/0 |
| detach | transfer | telic | yes | attach | detach-storage storage/0 |
| bind | transfer | atelic | no | — | bind mysql dmz |
| run | execution | atelic | no | — | run backup |
| exec | execution | atelic | no | — | exec --application mysql |
| cancel | execution | punctual | no | — | cancel-task 1 |
| resolved | execution | punctual | no | — | resolved mysql/0 |
| migrate | migration | telic | no | — | migrate mymodel otherctl |
| import | migration | telic | no | — | import-ssh-key, import-filesystem |
| download | migration | telic | no | — | download mysql, download-backup |
| sync | migration | atelic | no | — | sync-agent-binary |
| export | migration | telic | no | — | export-bundle |
| scale | mutation | atelic | no | — | scale-application mysql 3 |
| switch | navigation | punctual | no | — | switch mymodel |
| move | mutation | atelic | no | — | move-to-space dmz 10.0.0.0/24 |
| rename | mutation | punctual | no | — | rename-space old new |
| reload | mutation | atelic | no | — | reload-spaces |
| retry | execution | punctual | no | — | retry-provisioning 0 |
| default | mutation | punctual | no | — | default-credential aws mycred |
| diff | observation | atelic | no | — | diff-bundle bundle.yaml |

## Section 3: Semantic Domain Clustering

| Domain | Count | Commands | Naming Consistent? | Notes |
|--------|-------|----------|:------------------:|-------|
| Application | 21 | deploy, remove-application, add-unit, remove-unit, config, refresh, expose, unexpose, integrate, remove-relation, scale-application, show-application, show-unit, bind, trust, constraints, set-constraints, diff-bundle, resolved, suspend-relation, resume-relation | Yes | Core workload management |
| Model | 15 | add-model, destroy-model, show-model, model-config, model-constraints, model-defaults, model-secret-backend, grant, revoke, migrate, export-bundle, switch, retry-provisioning, set-credential, set-model-constraints | Mostly | Some split with controller commands |
| Controller | 10 | bootstrap, destroy-controller, kill-controller, show-controller, controller-config, controllers, add-model, register, unregister, enable-destroy-controller | Yes | includes bootstrap (unique) |
| Cloud | 12 | add-cloud, remove-cloud, show-cloud, update-cloud, clouds, regions, add-credential, remove-credential, update-credential, credentials, default-credential, default-region | Yes | Credential commands mixed but consistent |
| Kubernetes (k8s) | 3 | add-k8s, remove-k8s, update-k8s | Yes | Subset of cloud commands |
| User | 10 | add-user, remove-user, show-user, users, login, logout, change-user-password, enable-user, disable-user, whoami | Yes | Complete lifecycle |
| Storage | 11 | add-storage, remove-storage, show-storage, storage, attach-storage, detach-storage, create-storage-pool, remove-storage-pool, update-storage-pool, storage-pools, import-filesystem | Yes | Pool vs instance distinction |
| Secret | 12 | add-secret, remove-secret, update-secret, show-secret, secrets, grant-secret, revoke-secret, add-secret-backend, remove-secret-backend, update-secret-backend, show-secret-backend, secret-backends, model-secret-backend | Yes | Clear secret/backend split |
| Space | 7 | add-space, remove-space, show-space, spaces, move-to-space, rename-space, reload-spaces | Yes | Complete CRUD + move/rename |
| SSH | 8 | ssh, scp, add-ssh-key, remove-ssh-key, ssh-keys, import-ssh-key, debug-hooks, debug-code | Yes | Keys separate from access |
| Machine | 4 | add-machine, remove-machine, show-machine, machines | Yes | Simple CRUD |
| Action | 8 | actions, run, exec, show-action, show-operation, show-task, operations, cancel-task | Mostly | exec vs run overlap |
| Cross-model (Offer/SAAS) | 8 | offer, remove-offer, show-offer, offers, consume, remove-saas, find-offers, integrate (with remote) | Mostly | offer/saas terminology split |
| Charm Hub | 3 | find, info, download | Yes | Discovery operations |
| Resource | 3 | resources, attach-resource, charm-resources | Yes | Charm resources |
| Backup | 2 | create-backup, download-backup | Yes | Controller backup |
| Block | 3 | disable-command, enable-command, disabled-commands | Yes | Protection operations |
| Firewall | 2 | firewall-rules, set-firewall-rule | Yes | Cross-model firewall |
| Subnet | 1 | subnets | Yes | Read-only network info |
| Upgrade | 3 | upgrade-model, upgrade-controller, sync-agent-binary | Yes | Version management |
| Help | 5 | help, version, documentation, help-action-commands, help-hook-commands | Yes | Meta-commands |
| **Total** | **131** | | | |

## Section 4: Symmetry Audit

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|-----------|-----------------|-----------------|:-----------------:|:-------------------:|-------|
| add/remove unit | add-unit | remove-unit | Yes | Yes | Direct inverse |
| add/remove application | deploy | remove-application | No | Partial | Different verbs, asymmetry in cleanup |
| add/remove model | add-model | destroy-model | No | No | remove vs destroy, different safety levels |
| add/remove controller | bootstrap | destroy-controller | No | No | Different verbs, bootstrap is unique |
| add/remove cloud | add-cloud | remove-cloud | Yes | Yes | Direct inverse |
| add/remove credential | add-credential | remove-credential | Yes | Yes | Direct inverse |
| add/remove k8s | add-k8s | remove-k8s | Yes | Yes | Direct inverse |
| add/remove machine | add-machine | remove-machine | Yes | Yes | Direct inverse |
| add/remove space | add-space | remove-space | Yes | Yes | Direct inverse |
| add/remove ssh-key | add-ssh-key | remove-ssh-key | Yes | Yes | Direct inverse |
| add/remove storage | add-storage | remove-storage | Yes | Partial | Attach state differs |
| add/remove storage-pool | create-storage-pool | remove-storage-pool | Partial | Yes | create vs add inconsistency |
| add/remove secret | add-secret | remove-secret | Yes | Yes | Direct inverse |
| add/remove secret-backend | add-secret-backend | remove-secret-backend | Yes | Yes | Direct inverse |
| add/remove user | add-user | remove-user | Yes | Yes | Direct inverse |
| add/remove offer | offer | remove-offer | No | Yes | Verb vs noun mismatch |
| add/remove saas | consume | remove-saas | No | Yes | Different verbs entirely |
| integrate/remove | integrate | remove-relation | No | Yes | integrate vs remove-relation |
| expose/unexpose | expose | unexpose | Yes | Yes | Direct inverse |
| suspend/resume | suspend-relation | resume-relation | Yes | Yes | Direct inverse |
| attach/detach | attach-storage | detach-storage | Yes | Yes | Direct inverse |
| grant/revoke | grant | revoke | Yes | Yes | Direct inverse |
| grant/revoke cloud | grant-cloud | revoke-cloud | Yes | Yes | Direct inverse |
| grant/revoke secret | grant-secret | revoke-secret | Yes | Yes | Direct inverse |
| enable/disable user | enable-user | disable-user | Yes | Yes | Direct inverse |
| enable/disable command | enable-command | disable-command | Yes | Yes | Direct inverse |
| login/logout | login | logout | Yes | Yes | Direct inverse |
| register/unregister | register | unregister | Yes | Yes | Direct inverse |
| default credential | default-credential | — | No | No | No reverse (unset?) |
| default region | default-region | — | No | No | No reverse (unset?) |
| bind | bind | — | No | No | No reverse operation |
| trust | trust | — | No | No | No untrust command |
| scale up/down | scale-application | — | No | No | Scale down same command |

### Symmetry Issues Summary

| Issue | Commands Affected | Severity |
|-------|-------------------|----------|
| Verb mismatch (add/remove vs deploy/remove) | application, model | High |
| Verb mismatch (add/remove vs create/remove) | storage-pool | Medium |
| Verb mismatch (offer/remove-offer) | crossmodel | Medium |
| Verb mismatch (consume/remove-saas) | saas | Medium |
| Missing reverse operation | default-credential, default-region, bind, trust | Medium |
| destroy vs remove inconsistency | model, controller | Medium |

## Section 5: Confusion-Pair Audit

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|--------------|:--------------:|----------------|
| `deploy` | `add-unit` | Functional overlap | Medium | deploy creates app, add-unit scales existing |
| `deploy` | `add-model` | Naming similarity | Low | deploy for apps, add-model for models |
| `remove-application` | `remove-unit` | Scope ambiguity | Medium | remove-application removes all, remove-unit specific |
| `remove-application` | `destroy-model` | Scope ambiguity | Medium | remove-application for one app, destroy-model for all |
| `destroy-model` | `destroy-controller` | Scope ambiguity | High | model vs controller scope difference |
| `destroy-controller` | `kill-controller` | Verb synonym | High | destroy is graceful, kill is forced |
| `remove-relation` | `remove-saas` | Scope ambiguity | Medium | remove-relation for local, remove-saas for remote |
| `integrate` | `offer` | Functional overlap | Medium | integrate connects, offer exposes for cross-model |
| `integrate` | `consume` | Functional overlap | Medium | integrate for local/consumed, consume for remote |
| `offer` | `consume` | Direction ambiguity | Medium | offer creates, consume accepts |
| `config` | `model-config` | Scope ambiguity | High | config for app, model-config for model |
| `config` | `controller-config` | Scope ambiguity | High | config for app, controller-config for controller |
| `constraints` | `model-constraints` | Scope ambiguity | High | constraints for app, model-constraints for model |
| `set-constraints` | `set-model-constraints` | Scope ambiguity | High | app vs model scope |
| `grant` | `grant-cloud` | Scope ambiguity | Medium | grant for model, grant-cloud for cloud |
| `grant` | `grant-secret` | Scope ambiguity | Medium | grant for model, grant-secret for secret |
| `secrets` | `secret-backends` | Scope ambiguity | Medium | secrets are values, backends are storage |
| `add-secret` | `add-secret-backend` | Scope ambiguity | Medium | secret is value, backend is storage |
| `storage` | `storage-pools` | Scope ambiguity | Medium | storage instances vs pool definitions |
| `create-storage-pool` | `add-storage` | Verb inconsistency | Medium | create for pool, add for instance |
| `add-ssh-key` | `import-ssh-key` | Functional overlap | Medium | add is direct, import is from source |
| `ssh` | `scp` | Functional overlap | Low | ssh for shell, scp for file transfer |
| `run` | `exec` | Synonym verbs | High | run for actions, exec for arbitrary commands |
| `run` | `actions` | Functional overlap | Low | actions lists, run executes |
| `show-operation` | `show-task` | Scope ambiguity | Medium | operation is batch, task is individual |
| `operations` | `actions` | Scope ambiguity | Medium | operations are executions, actions are definitions |
| `login` | `register` | Functional overlap | Medium | login to existing, register adds new controller |
| `enable-user` | `enable-command` | Scope ambiguity | Low | user vs command, distinct help text |
| `upgrade-model` | `upgrade-controller` | Scope ambiguity | Medium | model vs controller scope |
| `refresh` | `upgrade-model` | Functional overlap | Medium | refresh updates charm, upgrade updates Juju |
| `relate` | `integrate` | Alias | Low | relate is alias for integrate |
| `find` | `info` | Scope ambiguity | Low | find searches, info shows details |
| `find` | `find-offers` | Scope ambiguity | Medium | find for charms, find-offers for cross-model |
| `resources` | `charm-resources` | Scope ambiguity | Medium | resources for deployed app, charm-resources for charm |
| `default-credential` | `default-region` | Functional similarity | Low | Both set defaults, different targets |
| `expose` | `unexpose` | Direction confusion | Low | Clear inverse relationship |
| `suspend-relation` | `resume-relation` | Direction confusion | Low | Clear inverse relationship |
| `reload-spaces` | `add-space` | Functional overlap | Medium | reload refreshes, add creates new |
| `rename-space` | `move-to-space` | Functional overlap | Medium | rename changes name, move changes subnet membership |
| `whoami` | `show-user` | Functional overlap | Medium | whoami shows current context, show-user shows other users |

## Section 6: Pattern Classification and Recommendations

### Pattern Classification

**Primary Pattern: Flat Verb-Noun Structure**

- All commands are top-level (`juju <verb>-<noun>`)
- No nested subcommands beyond the primary command
- Grouping is logical (by domain) not structural

**Depth: 1 level**

- `juju deploy mysql` - Command + argument
- No: `juju application deploy mysql` - Would be 2 levels

**Style: Verb-Noun Composite**

- Commands combine verb and noun: `add-model`, `remove-unit`
- Orphan commands: `bootstrap`, `whoami`, `switch`, `status`

### Discoverability Assessment

**Predictable Paths:**
- User wants to add X → `add-X` (mostly works)
- User wants to remove X → `remove-X` (mostly works)
- User wants to show X → `show-X` (mostly works)

**Problematic Paths:**
- User wants to create application → `create-application`? No, `deploy`
- User wants to delete model → `remove-model`? No, `destroy-model`
- User wants to delete controller → `remove-controller`? No, `destroy-controller`
- User wants to connect apps → `connect`? No, `integrate`
- User wants to create storage pool → `add-storage-pool`? No, `create-storage-pool`

### Ecosystem Comparison

| Feature | Juju | kubectl | terraform | helm |
|---------|------|---------|-----------|------|
| Pattern | Flat verb-noun | Nested resource | Resource-action | Noun-verb |
| Depth | 1 | 2-3 | 1 | 1-2 |
| CRUD consistency | Partial | Yes (get/create/delete) | Yes (apply/destroy) | Partial |
| Global flags | Yes | Yes | Yes | Yes |
| Help integration | Good | Excellent | Good | Good |

### Recommendations

#### 1. Verb Standardization (High Impact)

| Current | Recommended | Affected Commands |
|---------|-------------|-------------------|
| `deploy` | `add-application` (alias) | deploy |
| `destroy-model` | `remove-model` (alias) | destroy-model |
| `destroy-controller` | `remove-controller` (alias) | destroy-controller |
| `create-storage-pool` | `add-storage-pool` (alias) | create-storage-pool |
| `kill-controller` | `destroy-controller --force` | kill-controller (deprecate) |

**Migration:** Add aliases first, deprecate old names in next major version.

#### 2. Missing Reverse Operations (Medium Impact)

Add commands for:
- `unset-default-credential` (or make default-credential accept empty value)
- `unset-default-region`
- `unbind` (reverse of bind)
- `untrust` (reverse of trust)

#### 3. Scope Disambiguation (Medium Impact)

For ambiguous pairs, improve help text:

```
juju config --help
...
This command configures APPLICATION settings.
For model configuration, use: juju model-config
For controller configuration, use: juju controller-config
```

#### 4. Alias Cleanup (Low Impact)

Current aliases:
- `relate` → `integrate` (keep)
- `model-config` → `config` (confusing - different scope)
- `model-constraints` → `constraints` (confusing)
- `model-defaults` → `defaults` (confusing)

**Recommendation:** Remove model-* aliases that create confusion, keep only `relate`.

#### 5. Documentation Improvements (Low Impact)

- Add "See Also" for all scope-ambiguous commands
- Document verb conventions in `juju help`
- Add `juju help verbs` topic

### Backward Compatibility Impact

| Recommendation | Breaking? | Migration Path |
|----------------|-----------|----------------|
| Add aliases | No | N/A |
| Deprecate old names | Yes (next major) | 2-version deprecation cycle |
| Add missing commands | No | N/A |
| Remove confusing aliases | No (aliases) | N/A |

### Migration Cost

| Change | Effort | Risk |
|--------|--------|------|
| Add aliases for standardization | Low | Low |
| Add missing reverse commands | Medium | Low |
| Improve help text | Low | None |
| Remove aliases | Low | Low |
| Deprecate commands | Medium | Medium (user scripts) |

## Summary

### Shape Summary

The Juju CLI uses a flat verb-noun command structure with 159 top-level commands organized into 22 functional domains. The primary pattern is `verb-noun` composition, with notable exceptions like `bootstrap`, `deploy`, and `destroy-*`. The command set shows partial CRUD consistency with some verb mismatches.

### Key Findings

1. **Verb inconsistency**: `deploy`/`remove-application`, `add`/`destroy`, `create`/`remove` create confusion
2. **Missing reverse operations**: `trust`, `bind`, `default-*` have no inverse
3. **Scope ambiguity**: `config` vs `model-config` vs `controller-config` causes user confusion
4. **Alias confusion**: `model-config` alias to `config` hides scope difference
5. **Good symmetry**: Most add/remove, enable/disable, grant/revoke pairs are consistent

### Priority Recommendations

1. Add aliases to standardize verbs (`add-application`, `remove-model`, `remove-controller`)
2. Improve help text for scope-ambiguous commands
3. Add missing reverse operations (`untrust`, `unbind`)
4. Remove confusing aliases (`model-config`, `model-constraints`, `model-defaults`)
5. Document verb conventions in help system
