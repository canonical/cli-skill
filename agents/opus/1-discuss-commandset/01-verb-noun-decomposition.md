# Verb-Noun Decomposition Matrix

## Decomposition of Every Command

| Command | Verb | Noun | Type |
|---|---|---|---|
| actions | list | action | noun-shorthand |
| add-cloud | add | cloud | verb-noun |
| add-credential | add | credential | verb-noun |
| add-k8s | add | k8s | verb-noun |
| add-machine | add | machine | verb-noun |
| add-model | add | model | verb-noun |
| add-secret | add | secret | verb-noun |
| add-secret-backend | add | secret-backend | verb-noun |
| add-space | add | space | verb-noun |
| add-ssh-key | add | ssh-key | verb-noun |
| add-storage | add | storage | verb-noun |
| add-unit | add | unit | verb-noun |
| add-user | add | user | verb-noun |
| attach-resource | attach | resource | verb-noun |
| attach-storage | attach | storage | verb-noun |
| autoload-credentials | autoload | credential | verb-noun |
| bind | bind | application | verb-implied |
| bootstrap | bootstrap | controller | orphan |
| cancel-task | cancel | task | verb-noun |
| change-user-password | change | user-password | verb-noun |
| clouds | list | cloud | noun-shorthand |
| config | get/set | application-config | orphan |
| constraints | get | application-constraint | orphan |
| consume | consume | offer | verb-implied |
| controller-config | get/set | controller-config | orphan |
| controllers | list | controller | noun-shorthand |
| create-backup | create | backup | verb-noun |
| create-storage-pool | create | storage-pool | verb-noun |
| credentials | list | credential | noun-shorthand |
| dashboard | show | dashboard | orphan |
| debug-code | debug | code | verb-noun |
| debug-hooks | debug | hook | verb-noun |
| debug-log | debug | log | verb-noun |
| default-credential | set | default-credential | orphan |
| default-region | set | default-region | orphan |
| deploy | deploy | application | verb-implied |
| destroy-controller | destroy | controller | verb-noun |
| destroy-model | destroy | model | verb-noun |
| detach-storage | detach | storage | verb-noun |
| diff-bundle | diff | bundle | verb-noun |
| disable-command | disable | command | verb-noun |
| disable-user | disable | user | verb-noun |
| disabled-commands | list | disabled-command | noun-shorthand |
| download | download | charm | verb-implied |
| download-backup | download | backup | verb-noun |
| dump-db | dump | db | verb-noun |
| dump-model | dump | model | verb-noun |
| enable-command | enable | command | verb-noun |
| enable-destroy-controller | enable | destroy-controller | verb-noun |
| exec | exec | command | verb-implied |
| export-bundle | export | bundle | verb-noun |
| expose | expose | application | verb-implied |
| find | find | charm | verb-implied |
| find-offers | find | offer | verb-noun |
| firewall-rules | list | firewall-rule | noun-shorthand |
| grant | grant | model-access | verb-implied |
| grant-cloud | grant | cloud-access | verb-noun |
| grant-secret | grant | secret-access | verb-noun |
| help-action-commands | help | action-command | verb-noun |
| help-hook-commands | help | hook-command | verb-noun |
| import-filesystem | import | filesystem | verb-noun |
| import-ssh-key | import | ssh-key | verb-noun |
| info | show | charm-info | verb-implied |
| integrate | integrate | application | verb-implied |
| juju | — | — | orphan |
| kill-controller | kill | controller | verb-noun |
| login | login | user | verb-implied |
| logout | logout | user | verb-implied |
| machines | list | machine | noun-shorthand |
| migrate | migrate | model | verb-implied |
| model-constraints | get | model-constraint | orphan |
| model-defaults | get/set | model-default | orphan |
| model-secret-backend | get/set | model-secret-backend | orphan |
| models | list | model | noun-shorthand |
| move-to-space | move | space | verb-noun |
| offer | offer | application | verb-implied |
| offers | list | offer | noun-shorthand |
| operations | list | operation | noun-shorthand |
| refresh | refresh | application | verb-implied |
| regions | list | region | noun-shorthand |
| register | register | controller | verb-implied |
| reload-spaces | reload | space | verb-noun |
| remove-application | remove | application | verb-noun |
| remove-cloud | remove | cloud | verb-noun |
| remove-credential | remove | credential | verb-noun |
| remove-k8s | remove | k8s | verb-noun |
| remove-machine | remove | machine | verb-noun |
| remove-offer | remove | offer | verb-noun |
| remove-relation | remove | relation | verb-noun |
| remove-saas | remove | saas | verb-noun |
| remove-secret | remove | secret | verb-noun |
| remove-secret-backend | remove | secret-backend | verb-noun |
| remove-space | remove | space | verb-noun |
| remove-ssh-key | remove | ssh-key | verb-noun |
| remove-storage | remove | storage | verb-noun |
| remove-storage-pool | remove | storage-pool | verb-noun |
| remove-unit | remove | unit | verb-noun |
| remove-user | remove | user | verb-noun |
| rename-space | rename | space | verb-noun |
| resolved | resolve | unit | verb-implied |
| resources | list | resource | noun-shorthand |
| resume-relation | resume | relation | verb-noun |
| retry-provisioning | retry | provisioning | verb-noun |
| revoke | revoke | model-access | verb-implied |
| revoke-cloud | revoke | cloud-access | verb-noun |
| revoke-secret | revoke | secret-access | verb-noun |
| run | run | action | verb-implied |
| scale-application | scale | application | verb-noun |
| scp | scp | file | verb-implied |
| secret-backends | list | secret-backend | noun-shorthand |
| secrets | list | secret | noun-shorthand |
| set-constraints | set | application-constraint | verb-noun |
| set-credential | set | model-credential | verb-noun |
| set-firewall-rule | set | firewall-rule | verb-noun |
| set-model-constraints | set | model-constraint | verb-noun |
| show-action | show | action | verb-noun |
| show-application | show | application | verb-noun |
| show-cloud | show | cloud | verb-noun |
| show-controller | show | controller | verb-noun |
| show-credential | show | credential | verb-noun |
| show-machine | show | machine | verb-noun |
| show-model | show | model | verb-noun |
| show-offer | show | offer | verb-noun |
| show-operation | show | operation | verb-noun |
| show-secret | show | secret | verb-noun |
| show-secret-backend | show | secret-backend | verb-noun |
| show-space | show | space | verb-noun |
| show-status-log | show | status-log | verb-noun |
| show-storage | show | storage | verb-noun |
| show-task | show | task | verb-noun |
| show-unit | show | unit | verb-noun |
| show-user | show | user | verb-noun |
| spaces | list | space | noun-shorthand |
| ssh | ssh | session | verb-implied |
| ssh-keys | list | ssh-key | noun-shorthand |
| status | show | status | noun-shorthand |
| storage | list | storage | noun-shorthand |
| storage-pools | list | storage-pool | noun-shorthand |
| subnets | list | subnet | noun-shorthand |
| suspend-relation | suspend | relation | verb-noun |
| switch | switch | model | orphan |
| sync-agent-binary | sync | agent-binary | verb-noun |
| trust | trust | application | verb-implied |
| unexpose | unexpose | application | verb-implied |
| unregister | unregister | controller | verb-implied |
| update-cloud | update | cloud | verb-noun |
| update-credential | update | credential | verb-noun |
| update-k8s | update | k8s | verb-noun |
| update-public-clouds | update | public-cloud | verb-noun |
| update-secret | update | secret | verb-noun |
| update-secret-backend | update | secret-backend | verb-noun |
| update-storage-pool | update | storage-pool | verb-noun |
| upgrade-controller | upgrade | controller | verb-noun |
| upgrade-model | upgrade | model | verb-noun |
| users | list | user | noun-shorthand |
| version | show | version | orphan |
| whoami | show | identity | orphan |

## Matrix: Verbs × Nouns

Verbs (rows): add, attach, autoload, bind, bootstrap, cancel, change, create, debug, deploy, destroy, detach, diff, disable, download, dump, enable, exec, export, expose, find, grant, help, import, integrate, kill, login, logout, migrate, move, offer, refresh, register, reload, remove, rename, resolve, retry, revoke, run, scale, scp, set, show, ssh, suspend, switch, sync, trust, unexpose, unregister, update, upgrade

Nouns (columns): action, action-command, application, application-config, backup, charm, cloud, code, command, controller, credential, db, default-credential, default-region, disabled-command, filesystem, firewall-rule, hook, hook-command, k8s, log, machine, model, model-access, model-constraint, model-credential, model-default, offer, operation, provisioning, public-cloud, relation, resource, saas, secret, secret-backend, space, ssh-key, status, status-log, storage, storage-pool, subnet, task, unit, user, user-password, version

| | action | action-command | application | application-config | backup | charm | cloud | code | command | controller | credential | db | default-credential | default-region | disabled-command | filesystem | firewall-rule | hook | hook-command | k8s | log | machine | model | model-access | model-constraint | model-credential | model-default | offer | operation | provisioning | public-cloud | relation | resource | saas | secret | secret-backend | space | ssh-key | status | status-log | storage | storage-pool | subnet | task | unit | user | user-password | version |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| add | — | — | — | — | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | ✓ | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | ✓ | — | ✓ | ✓ | ✓ | ✓ | — | — | ✓ | — | — | — | ✓ | ✓ | — | — |
| attach | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — |
| autoload | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| bind | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| bootstrap | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| cancel | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — |
| change | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — |
| create | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — |
| debug | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | ✓ | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| deploy | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| destroy | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| detach | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — |
| diff | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| disable | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — |
| download | — | — | — | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| dump | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| enable | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — |
| exec | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| export | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| expose | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| find | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| grant | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — |
| help | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| import | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — |
| integrate | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| kill | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| login | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — |
| logout | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — |
| migrate | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| move | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — |
| offer | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| refresh | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| register | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| reload | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — |
| remove | — | — | ✓ | — | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | ✓ | — | ✓ | — | — | — | — | — | ✓ | — | — | — | ✓ | — | ✓ | ✓ | ✓ | ✓ | ✓ | — | — | ✓ | ✓ | — | — | ✓ | ✓ | — | — |
| rename | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — |
| resolve | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — |
| retry | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| revoke | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — |
| run | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| scale | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| scp | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| set | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | ✓ | — | — | ✓ | — | — | — | — | — | — | — | ✓ | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| show | ✓ | — | ✓ | — | — | — | ✓ | — | — | ✓ | ✓ | — | — | — | — | — | — | — | — | — | — | ✓ | ✓ | — | — | — | — | ✓ | ✓ | — | — | — | — | — | ✓ | ✓ | ✓ | — | ✓ | ✓ | ✓ | — | — | ✓ | ✓ | ✓ | — | — |
| ssh | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| suspend | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| switch | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| sync | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| trust | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| unexpose | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| unregister | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |
| update | — | — | — | — | — | — | ✓ | — | — | — | ✓ | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | ✓ | ✓ | — | — | — | — | — | ✓ | — | — | — | — | — | — |
| upgrade | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | ✓ | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — | — |

## Annotations

### Incomplete CRUD Sets

| Noun | Has Create | Has Read | Has Update | Has Delete | Missing |
|---|---|---|---|---|---|
| action | run | show-action | — | cancel-task | update |
| application | deploy | show-application | refresh, config | remove-application | — |
| backup | create-backup | — | — | download-backup (not delete) | list, delete |
| cloud | add-cloud | show-cloud | update-cloud | remove-cloud | — |
| controller | bootstrap | show-controller | upgrade-controller, controller-config | destroy-controller, kill-controller | — |
| credential | add-credential | show-credential | update-credential | remove-credential | — |
| k8s | add-k8s | — | update-k8s | remove-k8s | show |
| machine | add-machine | show-machine | — | remove-machine | update |
| model | add-model | show-model | model-config, upgrade-model | destroy-model, migrate | — |
| offer | offer | show-offer | — | remove-offer | update |
| relation | integrate | — | suspend-relation, resume-relation | remove-relation | show |
| resource | attach-resource | resources | — | — | detach, update |
| secret | add-secret | show-secret | update-secret | remove-secret | — |
| secret-backend | add-secret-backend | show-secret-backend | update-secret-backend | remove-secret-backend | — |
| space | add-space | show-space, spaces | move-to-space, rename-space | remove-space | update |
| ssh-key | add-ssh-key | ssh-keys | — | remove-ssh-key | update |
| storage | add-storage | show-storage | attach-storage | remove-storage, detach-storage | update |
| storage-pool | create-storage-pool | storage-pools | update-storage-pool | remove-storage-pool | — |
| subnet | — | subnets | — | — | add, remove, update |
| task | run | show-task | — | cancel-task | update |
| unit | add-unit | show-unit | — | remove-unit | update |
| user | add-user | show-user | change-user-password, enable-user, disable-user | remove-user | update |

### Verb Inconsistencies

| Noun | Inconsistent Verbs | Commands | Issue |
|---|---|---|---|
| application | deploy / add-unit | `deploy` creates app, `add-unit` adds units | Different lifecycle verbs for same resource family |
| controller | bootstrap / add-model / destroy / kill | `bootstrap`, `add-model`, `destroy-controller`, `kill-controller` | `bootstrap` is not `create-controller`; `kill` vs `destroy` inconsistency |
| credential | add / update / autoload / detect | `add-credential`, `update-credential`, `autoload-credentials` | `detect` is missing; `autoload` is irregular |
| model | add / destroy / migrate / dump / export | `add-model`, `destroy-model`, `migrate`, `dump-model`, `export-bundle` | No `update-model`; `migrate` is unusual |
| storage | add / attach / detach / remove / import | `add-storage`, `attach-storage`, `detach-storage`, `remove-storage`, `import-filesystem` | `import` uses different noun (`filesystem`) |
| config | config / model-config / controller-config / model-defaults | Overloaded term with scoped variants | Scope ambiguity (see confusion audit) |

### Orphan Commands

| Command | Why Orphan | Suggested Decomposition |
|---|---|---|
| `bootstrap` | Self-initializing infrastructure; no clear noun | create-controller (but breaks convention) |
| `config` | Get/set hybrid; no noun in name | application-config (but conflicts with `model-config`) |
| `consume` | Domain-specific verb; implicit noun is `offer` | consume-offer |
| `dashboard` | Noun-only; not a list command | show-dashboard |
| `deploy` | Domain-specific; implicit noun `application` | deploy-application (verbose) |
| `diff-bundle` | Verb-noun exists but `diff` is rare in CLI | compare-bundle |
| `exec` | Generic execution verb | exec-command (redundant) |
| `expose` | Toggle verb; paired with `unexpose` | expose-application |
| `integrate` | Domain-specific; means "create relation" | integrate-applications |
| `juju` | Interactive shell meta-command | — |
| `migrate` | Domain-specific; implicit noun `model` | migrate-model |
| `offer` | Domain-specific; implicit noun `application-endpoint` | offer-endpoint |
| `refresh` | Domain-specific; means "update charm" | refresh-charm |
| `resolved` | Past participle used as command | resolve-unit |
| `run` | Generic execution; implicit noun `action` | run-action |
| `scp` | Binary passthrough | — |
| `ssh` | Binary passthrough | — |
| `status` | Special noun-only status command | show-status (but redundant with `juju status`) |
| `switch` | Context-switching verb | switch-model |
| `sync-agent-binary` | Triple compound; unwieldy | sync-agent-binaries |
| `trust` | Toggle verb | trust-application |
| `unexpose` | Toggle verb; paired with `expose` | unexpose-application |
| `unregister` | Irregular verb; paired with `register` | unregister-controller |
| `version` | Meta-command | show-version |
| `whoami` | Compound pronoun command | show-identity |
