## Section 5: Confusion-Pair Audit

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|
| `remove-application` | `destroy-model` | synonym verbs | High | Both delete resources; `remove` is standard, `destroy` is reserved for model/controller. User may typo `destroy-application` expecting it to exist. |
| `remove-unit` | `remove-machine` | scope ambiguity | Medium | Both remove infrastructure; a unit runs on a machine. Users may be unsure whether to remove the unit or the underlying machine. |
| `kill-controller` | `destroy-controller` | synonym verbs | High | Both terminate controllers; `kill` is forceful, `destroy` is graceful. Risk of accidental forceful teardown. |
| `config` | `model-config` | scope ambiguity | High | Both get/set configuration; `config` targets applications, `model-config` targets the model. Easy to use wrong default. |
| `model-config` | `controller-config` | scope ambiguity | Medium | Both configuration surfaces at different scopes. New users may conflate them. |
| `constraints` | `set-constraints` | functional overlap | Medium | `constraints` shows, `set-constraints` sets — but the naming suggests they act on the same thing. Clear enough with usage. |
| `model-constraints` | `set-model-constraints` | functional overlap | Medium | Same pattern as above at model scope. |
| `resources` | `charm-resources` | naming collision | Medium | Both show resources; `resources` is for deployed app/unit, `charm-resources` is for a charm in a repo. Scope differs but names are close. |
| `exec` | `run` | functional overlap | High | Both execute commands on units. `exec` runs arbitrary shell commands; `run` executes named charm actions. Critical distinction for operators. |
| `ssh` | `exec` | functional overlap | Medium | Both access unit/machine. `ssh` is interactive shell; `exec` is command execution. Less confusing than `exec`/`run` because ssh is familiar. |
| `find` | `info` | functional overlap | Medium | Both query Charmhub. `find` searches; `info` shows details for a specific charm. Distinction is standard but worth noting. |
| `find-offers` | `offers` | naming collision | Low | Both relate to offers. `find-offers` searches across controllers; `offers` lists local model's offers. Different scopes. |
| `show-operation` | `show-task` | naming collision | Medium | Both show execution results. `operation` is the orchestration; `task` is the per-unit execution. Related but different granularity. |
| `operations` | `run` | functional overlap | Low | `operations` lists them; `run` creates them. Standard CRUD relationship. |
| `grant` | `grant-cloud` | scope ambiguity | Medium | `grant` acts on model/controller/offer; `grant-cloud` is cloud-specific. Subcommand naming helps but bare `grant` is ambiguous without context. |
| `revoke` | `revoke-cloud` | scope ambiguity | Medium | Same pattern as grant pair. |
| `add-secret` | `grant-secret` | functional overlap | Low | `add-secret` creates it; `grant-secret` shares access. Different lifecycle phases. |
| `consume` | `integrate` | functional overlap | Medium | Both connect applications. `consume` brings a remote offer into the model; `integrate` wires two local endpoints. Different CMR vs local flows. |
| `consume` | `offer` | functional overlap | Low | `offer` exports; `consume` imports. They are inverse directions of the same concept. |
| `suspend-relation` | `remove-relation` | functional overlap | Medium | Both stop relations. `suspend` pauses with resume possible; `remove` deletes. Users may choose wrong one if they want temporary vs permanent. |
| `refresh` | `upgrade-model` | functional overlap | Low | Both update software. `refresh` updates a charm; `upgrade-model` updates Juju agent on all machines. Different targets. |
| `refresh` | `update-secret` | naming collision | Low | Both use "update" semantics. `refresh` is charm-specific jargon; `update` is generic. No real confusion expected. |
| `download` | `download-backup` | naming collision | Low | Both download. `download` is for charms; `download-backup` is for controller backups. Context disambiguates. |
| `status` | `show-status-log` | naming collision | Low | Both show status. `status` is current snapshot; `show-status-log` is historical. Suffix helps. |
| `whoami` | `show-user` | functional overlap | Low | Both show user info. `whoami` is the current session; `show-user` is any user. Clear enough. |
| `switch` | `migrate` | functional overlap | Medium | Both change the active model context. `switch` changes local client pointer; `migrate` moves model data to another controller. Very different operations, same mental model phrase. |
| `trust` | `enable-user` | functional overlap | Low | Both grant permissions. `trust` is application-level cloud access; `enable-user` is account reactivation. Different domains. |
| `debug-code` | `debug-hooks` | naming collision | High | Names differ by one word but both launch tmux debug sessions. `debug-code` is for actions; `debug-hooks` is for hooks. Users may conflate them. |

---

