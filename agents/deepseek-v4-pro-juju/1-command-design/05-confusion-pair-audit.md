# 05 — Confusion-Pair Audit

## Overview

Every pair of commands that share semantic overlap and risk user confusion. Pairs are listed exhaustively with overlap type, risk level, and disambiguation.

## Confusion Pairs

### Group 1: Synonym Verbs (Different verbs, same or similar operation)

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|-------------|----------------|---------------|
| `destroy-controller` | `kill-controller` | Synonym verbs | **High** | `destroy` requires controller to be reachable; `kill` works even when controller is unreachable. `kill` is last-resort forced destruction. |
| `destroy-model` | `remove-application` | Synonym verbs (different scope) | **Medium** | `destroy-model` deletes the entire model including all apps; `remove-application` deletes a single app. Different verbs signal different impact. |
| `remove-relation` | `remove-saas` | Synonym verbs (different nouns) | **Medium** | `remove-relation` removes an integration between two applications; `remove-saas` removes a consumed cross-model SAAS reference. |
| `remove-application` | `remove-machine` | Synonym verbs (different nouns) | **Low** | Clear noun distinction. Users rarely confuse these. |
| `remove-unit` | `remove-machine` | Synonym verbs (related nouns) | **Medium** | `remove-unit` removes a unit (which may remove the machine if it's the last unit). `remove-machine` removes the machine directly. |
| `run` | `exec` | Synonym verbs | **High** | `run` executes in the charm/hook context (actions); `exec` executes directly in the workload container. Different security and environment contexts. |
| `integrate` | `relate` (alias) | True synonyms (aliases) | **Low** | `relate` is a deprecated alias for `integrate`. Both do the same thing. Alias should eventually be deprecated per DE013. |
| `refresh` | `upgrade-model` | Synonym verbs (different scope) | **Medium** | `refresh` updates a charm; `upgrade-model` upgrades the Juju agent version for the model. |
| `upgrade-model` | `upgrade-controller` | Synonym verbs (different nouns) | **Medium** | Similar name but `upgrade-model` affects model agents; `upgrade-controller` affects controller agents. |
| `deploy` | `add-unit` | Functional overlap | **Medium** | `deploy -n 5` deploys with 5 units initially; `add-unit -n 3` adds 3 more units to an existing deployment. Both add units, but `deploy` also creates the application. |
| `deploy` | `refresh` | Functional overlap | **Low** | `deploy` creates a new application; `refresh` updates an existing one. Different lifecycle stages. |

### Group 2: Scope Ambiguity (Same verb, unclear which scope)

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|-------------|----------------|---------------|
| `config` | `model-config` | Scope ambiguity | **High** | `config` operates on an application; `model-config` operates on a model. Almost identical names but different scopes. |
| `config` | `controller-config` | Scope ambiguity | **Medium** | `config` (application) vs `controller-config`. The bare `config` is at application level; controller config is explicitly scoped. |
| `model-config` | `model-defaults` | Scope ambiguity | **High** | `model-config` changes an existing model; `model-defaults` sets default config for new models. Very similar names with different effects. |
| `model-config` | `controller-config` | Scope ambiguity | **Medium** | Both set configuration but at different hierarchy levels. |
| `constraints` | `model-constraints` | Scope ambiguity | **High** | Bare `constraints` shows application constraints; `model-constraints` shows model constraints. |
| `set-constraints` | `set-model-constraints` | Scope ambiguity | **High** | `set-constraints` sets application constraints; `set-model-constraints` sets model constraints. |
| `machines` | `status` | Scope ambiguity | **Low** | `machines` lists only machines; `status` lists everything including machines. |
| `models` | `status` | Scope ambiguity | **Low** | `models` lists all models on a controller; `status` shows the current model's state. |
| `grant` | `grant-cloud` | Scope ambiguity | **Medium** | `grant` grants model access; `grant-cloud` grants cloud access. |
| `revoke` | `revoke-cloud` | Scope ambiguity | **Medium** | Same as grant pair. |

### Group 3: Functional Overlap (Different commands that achieve similar outcomes)

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|-------------|----------------|---------------|
| `destroy-model --force` | `destroy-model` + manual cleanup | Functional overlap | **Medium** | `--force` skips error handling and waiting; manual cleanup is piecemeal. Users may not know when `--force` is appropriate. |
| `destroy-controller` | `kill-controller` | Functional overlap | **High** | Both destroy the controller. `kill-controller` is for unreachable controllers. `destroy-controller --force` overlaps behaviorally with `kill-controller`. |
| `remove-application --force --no-wait` | `remove-application` then `remove-machine` | Functional overlap | **Medium** | `--force --no-wait` does rapid teardown; manual step-by-step gives more control. |
| `consume` | `deploy` (with SAAS) | Functional overlap | **Low** | `consume` creates a local reference to a remote app; `deploy` creates a new app from a charm. Different targets. |
| `export-bundle` | `show-model` + manual reconstruction | Functional overlap | **Low** | `export-bundle` produces a reusable YAML bundle; `show-model` shows the model in detail. |

### Group 4: Naming Collision (Names too similar, different purposes)

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|-------------|----------------|---------------|
| `resources` | `charm-resources` | Naming collision | **High** | `resources` shows resources for an application/unit; `charm-resources` shows what resources a charm requires before deployment. |
| `secrets` | `secret-backends` | Naming collision | **Medium** | `secrets` lists secrets in the model; `secret-backends` lists secret storage backends. |
| `model-secret-backend` | `secret-backends` | Naming collision | **Medium** | `model-secret-backend` sets the default backend; `secret-backends` lists all backends. |
| `storage` | `storage-pools` | Naming collision | **Medium** | `storage` lists storage instances; `storage-pools` lists storage pool definitions. |
| `add-storage` | `attach-storage` | Naming collision | **Medium** | `add-storage` creates and attaches new storage; `attach-storage` attaches existing storage to a different unit. |
| `remove-storage` | `detach-storage` | Naming collision | **Medium** | `remove-storage` destroys storage; `detach-storage` disconnects but preserves it. |
| `regions` | `default-region` | Naming collision | **Low** | `regions` lists; `default-region` sets a default. |
| `credentials` | `default-credential` | Naming collision | **Low** | Similar to regions pair. |
| `offer` | `offers` | Naming collision | **Medium** | `offer` creates; `offers` lists. Singular/plural distinction is subtle for non-native speakers. |
| `find-offers` | `offers` | Naming collision | **Medium** | `find-offers` searches across controllers; `offers` lists in the current model. |
| `show-offer` | `show-operation` | Naming collision | **Low** | Similar `show-*` prefix but different domains. Users unlikely to confuse these specific pairs. |
| `dump-model` | `dump-db` | Naming collision | **Medium** | Both dump but at different levels. `dump-model` dumps model data; `dump-db` dumps the full controller DB. Both developer-mode gated. |
| `create-backup` | `download-backup` | Naming collision | **Low** | `create` makes a backup on the controller; `download` fetches it locally. |
| `add-ssh-key` | `import-ssh-key` | Naming collision | **Medium** | `add` adds a raw key; `import` fetches from GitHub/Launchpad. Both add SSH keys but through different mechanisms. |
| `update-credential` | `set-credential` | Naming collision | **Medium** | `update-credential` changes credential values; `set-credential` applies a credential to a model. Different targets. |
| `add-credential` | `autoload-credentials` | Naming collision | **Medium** | `add-credential` adds manually; `autoload-credentials` detects and imports from environment. |

### Group 5: Verb Peculiarities (Verbs that may mislead)

| Command | Potential Confusion | Risk |
|---------|-------------------|------|
| `bootstrap` | No obvious reverse. Users may expect `debootstrap` or `remove-controller`. | Medium |
| `consume` | Metaphorical. Users may think it "uses up" the remote offer (it doesn't). | Medium |
| `resolved` | Adjective, not a verb. Users may not recognize it as a command. | Low |
| `trust` | With `--remove`, reverses. New users may expect `untrust`. | Medium |
| `switch` | Changes context. Users may think it "switches" model state or toggles something. | Low |
| `integrate` | Vague verb. `relate` alias is clearer to English speakers. | Medium |
| `whoami` | UNIX convention. Expected by power users; non-obvious to newcomers. | Low |
| `disabled-commands` | Awkward plural: `disabled-commands` is a list of disabled commands, not a list-of-disabled-commands. It actually lists which commands ARE disabled. | Low |
| `default-credential` | Sets the default, doesn't show it. Same for `default-region`. | Medium |

## Risk Summary

| Risk Level | Count | Pairs |
|-----------|-------|-------|
| **High** | 8 | `destroy-controller`/`kill-controller`, `run`/`exec`, `config`/`model-config`, `model-config`/`model-defaults`, `constraints`/`model-constraints`, `set-constraints`/`set-model-constraints`, `destroy-controller`/`kill-controller` (functional), `resources`/`charm-resources` |
| **Medium** | 23 | Various scope and naming collisions |
| **Low** | 10 | Minor naming similarities, low practical confusion |

## Top 5 Most Critical Confusion Pairs

1. **`destroy-controller` vs `kill-controller`** (High) — Users may not know which to use. Wrong choice can leave orphan resources or fail unnecessarily.
2. **`run` vs `exec`** (High) — Similar names, similar purpose, but completely different execution environments. Choosing wrong one gives confusing errors.
3. **`config` vs `model-config` vs `controller-config`** (High) — Three config commands at different scopes. `config` is application scope but not named `application-config`.
4. **`model-config` vs `model-defaults`** (High) — Very similar names. One changes current model, one changes future defaults. Users may accidentally set defaults when intending to change their model.
5. **`resources` vs `charm-resources`** (High) — `resources` shows deployed app resources; `charm-resources` shows charm requirements. Different stages of the lifecycle.

## DE013 Compliance Notes

Per DE013 §Vocabulary: "Choosing the right verb is not trivial: it needs to imply or trigger recall of the object type it refers to. And when a command acts on different object types, it needs to help the user differentiate between these types."

- `config` fails this: it doesn't signal it acts on applications, unlike `model-config` and `controller-config` which explicitly name their object.
- `constraints`/`set-constraints` similarly fail: the object is not named, making them ambiguous with `model-constraints`/`set-model-constraints`.
- Recommendation: Rename `config` → `application-config` (per DE013: "When verbs alone are not sufficient... use the verb-noun form").
