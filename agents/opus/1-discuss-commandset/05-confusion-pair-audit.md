# Confusion-Pair Audit

## High-Risk Pairs

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|
| `config` | `model-config` | scope ambiguity | **high** | `config` operates on applications; `model-config` operates on the entire model. |
| `config` | `controller-config` | scope ambiguity | **high** | `config` is application-scoped; `controller-config` is controller-scoped. |
| `model-config` | `controller-config` | scope ambiguity | **high** | Both set key=value configurations but for different scopes (model vs controller). |
| `constraints` | `model-constraints` | scope ambiguity | **high** | `constraints` shows application constraints; `model-constraints` shows model constraints. |
| `set-constraints` | `set-model-constraints` | scope ambiguity | **high** | Both set constraints but at different scopes (application vs model). |
| `remove-application` | `destroy-model` | synonym verbs | **high** | Both destroy applications; `destroy-model` destroys everything in the model including the app. |
| `destroy-controller` | `kill-controller` | functional overlap | **high** | Both terminate controllers; `kill` bypasses graceful shutdown. |
| `destroy-model` | `kill-controller` | functional overlap | **high** | `destroy-model` removes one model; `kill-controller` destroys the entire controller and all its models. |
| `integrate` | `consume` | functional overlap | **high** | Both establish cross-model relations; `integrate` is intra-model, `consume` is inter-model. |
| `add-model` | `create-backup` | synonym verbs | medium | Both "create" things but in completely different domains; users may conflate creation commands. |
| `deploy` | `add-unit` | functional overlap | medium | `deploy` creates the first unit implicitly; `add-unit` adds more. Users sometimes use `add-unit` when they meant `deploy`. |
| `show-unit` | `show-application` | scope ambiguity | medium | Both show information about application workloads; `show-unit` is per-unit, `show-application` aggregates. |
| `show-machine` | `show-unit` | scope ambiguity | medium | Both show workload status; `show-machine` is infrastructure-focused, `show-unit` is application-focused. |
| `status` | `show-status-log` | naming collision | medium | Both contain "status" but `status` is a snapshot while `show-status-log` is a historical stream. |
| `run` | `exec` | functional overlap | **high** | Both execute commands remotely; `run` invokes charm actions, `exec` runs arbitrary shell commands. |
| `ssh` | `exec` | functional overlap | medium | Both provide remote execution; `ssh` is interactive shell, `exec` is non-interactive command dispatch. |
| `debug-hooks` | `debug-code` | naming collision | **high** | Names differ by one word but do nearly the same thing (launch tmux for debugging). |
| `add-cloud` | `update-cloud` | functional overlap | medium | Both modify cloud metadata; `add` is for new clouds, `update` for existing. |
| `add-credential` | `update-credential` | functional overlap | medium | Both modify credentials; `add` creates, `update` modifies existing. |
| `update-credential` | `set-credential` | functional overlap | medium | Both change which credential a model uses; `update-credential` updates the credential definition, `set-credential` assigns it to a model. |
| `grant` | `grant-cloud` | scope ambiguity | medium | `grant` gives model access; `grant-cloud` gives cloud access. |
| `grant` | `trust` | functional overlap | medium | Both grant permissions; `grant` is user-to-model, `trust` is application-to-cloud-credential. |
| `revoke` | `revoke-cloud` | scope ambiguity | medium | `revoke` removes model access; `revoke-cloud` removes cloud access. |
| `remove-relation` | `suspend-relation` | functional overlap | medium | Both stop relations; `remove` is permanent, `suspend` is temporary. |
| `refresh` | `upgrade-model` | functional overlap | medium | Both update software; `refresh` updates a charm, `upgrade-model` updates Juju agent version. |
| `refresh` | `update-secret` | synonym verbs | medium | Both mean "update" but operate on different objects (charm vs secret). |
| `upgrade-controller` | `upgrade-model` | scope ambiguity | medium | Both upgrade Juju version but at different scopes (controller vs model). |
| `sync-agent-binary` | `upgrade-controller` | functional overlap | medium | Both update controller binaries; `sync` copies from store, `upgrade` performs a rolling upgrade. |
| `download` | `download-backup` | naming collision | medium | Both are `download` commands but for different objects (charm vs backup). |
| `find` | `info` | functional overlap | medium | Both query CharmHub; `find` searches, `info` shows details for a specific item. |
| `find` | `find-offers` | naming collision | medium | Both `find` commands but in different domains (CharmHub vs cross-model offers). |
| `show-cloud` | `clouds` | functional overlap | medium | `clouds` lists all clouds; `show-cloud` shows one cloud in detail. |
| `show-credential` | `credentials` | functional overlap | medium | `credentials` lists all; `show-credential` shows one in detail. |
| `show-controller` | `controllers` | functional overlap | medium | `controllers` lists all; `show-controller` shows one in detail. |
| `show-model` | `models` | functional overlap | medium | `models` lists all; `show-model` shows one in detail. |
| `show-user` | `users` | functional overlap | medium | `users` lists all; `show-user` shows one in detail. |
| `show-secret` | `secrets` | functional overlap | medium | `secrets` lists all; `show-secret` shows one in detail. |
| `show-space` | `spaces` | functional overlap | medium | `spaces` lists all; `show-space` shows one in detail. |
| `show-storage` | `storage` | functional overlap | medium | `storage` lists all; `show-storage` shows one in detail. |
| `users` | `whoami` | functional overlap | medium | `users` lists all users; `whoami` shows only the current user. |
| `login` | `register` | functional overlap | medium | Both establish client-controller relationships; `login` authenticates, `register` imports controller metadata. |
| `login` | `switch` | functional overlap | low | Both change active context; `login` changes user, `switch` changes model/controller. |
| `logout` | `unregister` | functional overlap | medium | Both disconnect from a controller; `logout` removes auth, `unregister` removes metadata. |
| `dump-model` | `export-bundle` | functional overlap | medium | Both export model state; `dump-model` is diagnostic, `export-bundle` is reusable. |
| `dump-model` | `dump-db` | naming collision | medium | Both `dump` commands but at different abstraction levels (model vs raw MongoDB). |
| `disable-command` | `enable-command` | naming collision | low | Opposites by design; confusion only if user forgets which is which. |
| `firewall-rules` | `set-firewall-rule` | functional overlap | low | `firewall-rules` lists; `set-firewall-rule` modifies. Standard list/set pattern. |
| `storage-pools` | `create-storage-pool` | functional overlap | low | `storage-pools` lists; `create-storage-pool` creates. Standard pattern. |
| `disabled-commands` | `disable-command` | naming collision | low | Noun/verb forms; `disabled-commands` lists blocked commands. |
| `help-action-commands` | `help-hook-commands` | naming collision | low | Very similar names but for different charm concepts (actions vs hooks). |
| `import-filesystem` | `attach-storage` | functional overlap | medium | Both bring storage into the model; `import` is for existing filesystems, `attach` is for existing Juju storage. |
| `scale-application` | `add-unit` | functional overlap | medium | Both increase unit count; `scale-application` sets absolute count, `add-unit` adds relative count. |
| `add-k8s` | `update-k8s` | functional overlap | low | `add` creates new k8s cloud; `update` modifies existing. |
| `retry-provisioning` | `add-machine` | functional overlap | low | Both result in machines; `retry` is for failed attempts, `add` is for new requests. |
| `change-user-password` | `disable-user` | functional overlap | low | Both modify user state but for different purposes (password vs access). |
| `model-defaults` | `model-config` | scope ambiguity | medium | Both configure models; `model-defaults` applies to future models, `model-config` to current model. |
| `default-credential` | `set-credential` | functional overlap | medium | Both assign credentials; `default-credential` sets cloud default, `set-credential` assigns to specific model. |
| `default-region` | `regions` | functional overlap | low | `regions` lists; `default-region` sets default. Standard pattern. |
| `autoload-credentials` | `detect-credentials` | functional overlap | low | `autoload-credentials` (cloud command) and `detect-credentials` are similar; only `autoload-credentials` exists in CLI. |
| `migrate` | `register` | functional overlap | low | Both connect to new controllers; `migrate` moves models, `register` adds controller to client. |
| `consume` | `integrate` | functional overlap | **high** | Both create relations; `integrate` is for same-controller apps, `consume` is for remote offers. |
| `offer` | `expose` | functional overlap | medium | Both make applications accessible externally; `offer` is cross-model, `expose` is network-facing. |
| `remove-offer` | `remove-saas` | functional overlap | medium | Both remove cross-model relations; `remove-offer` deletes the offer, `remove-saas` deletes the local proxy. |
| `attach-resource` | `upload` (not a command) | naming collision | low | `attach-resource` updates resources; users may look for `upload-resource`. |
| `charm-resources` | `resources` | naming collision | medium | `charm-resources` shows charm repository resources; `resources` shows deployed application resources. |
| `status` | `machines` | functional overlap | low | `status` includes machines; `machines` is a filtered list. |
| `status` | `storage` | functional overlap | low | `status` can include storage; `storage` is a dedicated list. |
| `status` | `models` | functional overlap | low | `status` is per-model; `models` lists across controllers. |
| `scp` | `ssh` | functional overlap | low | Both use SSH transport; `scp` transfers files, `ssh` opens a shell. |
| `enable-destroy-controller` | `destroy-controller` | naming collision | medium | `enable-destroy-controller` removes blocks; `destroy-controller` performs destruction. |
| `kill-controller` | `enable-destroy-controller` | functional overlap | low | Both enable destruction; `kill` is direct, `enable-destroy` is preparatory. |
| `update-public-clouds` | `update-cloud` | functional overlap | medium | Both update cloud metadata; `update-public-clouds` is for built-in definitions, `update-cloud` is for personal clouds. |
| `cancel-task` | `remove-unit` | functional overlap | low | Both stop things; `cancel-task` stops actions, `remove-unit` removes workloads. |
| `show-operation` | `show-task` | scope ambiguity | medium | Both show action results; `show-operation` is aggregate, `show-task` is per-unit. |
| `operations` | `actions` | functional overlap | medium | Both list action-related items; `actions` lists definitions, `operations` lists executions. |
| `version` | `show-controller` | functional overlap | low | Both show version info; `version` is CLI client, `show-controller` shows controller agent version. |
| `juju` (interactive) | `ssh` | functional overlap | low | Both provide interactive shells; `juju` REPL runs Juju commands, `ssh` runs shell commands. |
| `export-bundle` | `deploy` | functional overlap | low | `export-bundle` output can be fed into `deploy`; they are inverse operations at the model level. |
| `diff-bundle` | `export-bundle` | functional overlap | low | Both compare model state to bundles; `diff-bundle` shows deltas, `export-bundle` produces the bundle. |

## Medium-Risk Pairs (Not Listed Above)

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|
| `add-storage` | `create-storage-pool` | naming collision | medium | `add-storage` assigns storage to a unit; `create-storage-pool` defines a pool type. |
| `remove-storage` | `detach-storage` | functional overlap | medium | `detach` unassigns; `remove` destroys. |
| `reload-spaces` | `move-to-space` | functional overlap | low | Both modify space state; `reload` re-imports from substrate, `move` reassigns CIDRs. |
| `move-to-space` | `rename-space` | functional overlap | low | Both modify spaces; `move` changes CIDRs, `rename` changes the name. |
| `set-firewall-rule` | `firewall-rules` | functional overlap | low | Standard set/list pattern. |
| `model-secret-backend` | `secret-backends` | scope ambiguity | medium | `model-secret-backend` sets per-model backend; `secret-backends` lists controller backends. |
| `update-secret` | `grant-secret` | functional overlap | low | Both modify secrets; `update` changes data, `grant` changes access. |
| `grant-secret` | `revoke-secret` | naming collision | low | Opposites by design. |
| `debug-log` | `show-status-log` | functional overlap | medium | Both show historical logs; `debug-log` is controller/agent logs, `show-status-log` is entity status history. |
| `sync-agent-binary` | `upgrade-controller` | functional overlap | medium | Both put binaries on controller; `sync` is manual copy, `upgrade` is managed rollout. |
