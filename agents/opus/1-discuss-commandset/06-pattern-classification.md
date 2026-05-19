# Pattern Classification and Recommendations

## Pattern Classification

### Primary Grouping Pattern: Flat Verb-Noun with Noun-Only List Exceptions

Juju uses a **flat namespace** with ~157 top-level commands. There are no nested subcommands (e.g., `juju cloud add` is `juju add-cloud`). The hierarchy exists only in source code organization (`cmd/juju/<domain>/`), not in the CLI surface.

| Pattern | Count | Examples |
|---|---|---|
| Verb-noun compound | ~90 | `add-cloud`, `remove-unit`, `show-secret`, `update-credential` |
| Noun-only (list/show shorthand) | ~20 | `machines`, `controllers`, `secrets`, `status` |
| Verb-only (implicit noun) | ~25 | `deploy`, `expose`, `refresh`, `trust`, `migrate` |
| Orphan/special | ~15 | `bootstrap`, `config`, `juju`, `switch`, `whoami` |

### Naming Convention Depth

| Depth | Example | Frequency |
|---|---|---|
| Single word | `deploy`, `status`, `config` | Common |
| Two words (verb-noun) | `add-cloud`, `show-model` | Most common |
| Three words | `show-secret-backend`, `update-storage-pool`, `help-action-commands` | Moderate |
| Four words | `enable-destroy-controller` | Rare |

### DE013 Compliance Assessment

| DE013 Rule | Juju Compliance | Notes |
|---|---|---|
| Commands are verbs | Partial | ~45 commands are verb-only or verb-noun; ~20 are noun-only lists; ~15 are orphans |
| Use plural nouns for listing | Yes | `machines`, `controllers`, `secrets`, `users`, `clouds` follow this |
| Use `status` over `show-status` | Yes | `status` is correctly used |
| Use `show` for details | Partial | Many `show-*` commands exist, but some use noun-only (`info`, `dashboard`) |
| Verb-noun when verbs alone insufficient | Partial | Most commands follow this, but many verb-only commands have implicit nouns |
| At most one sublevel | N/A | Juju uses flat namespace; no sublevels at all |

### Comparison to DE013 Common Commands

| DE013 Standard | Juju Equivalent | Deviation? |
|---|---|---|
| `tool init` | `bootstrap` | Minor — `bootstrap` is domain-specific |
| `tool list` | `models`, `controllers`, `machines` | Yes — Juju uses plural nouns instead of `list-*` |
| `tool show <id>` | `show-application`, `show-model` | Yes — Juju prefixes with `show-` |
| `tool status` | `status` | No |
| `tool foobar-status` | N/A | N/A |
| `tool start/stop` | N/A | Juju manages lifecycle via deploy/remove/scale |
| `tool enable/disable` | `enable-command`, `disable-user` | Partial — inconsistent noun pairing |
| `tool get/set/unset` | `config`, `model-config`, `set-constraints` | Yes — no `unset`, `get` is implicit |
| `tool create-foo` | `create-backup`, `create-storage-pool` | Partial — Juju mostly uses `add-*` |
| `tool delete-foo` | `remove-*`, `destroy-*`, `kill-*` | Yes — three different deletion verbs |
| `tool update-foo` | `update-*`, `refresh` | Partial — `refresh` is domain-specific |

## Discoverability Assessment

### Predicted vs Actual Paths

| User Intent | Predicted Command | Actual Command | Discoverability Issue |
|---|---|---|---|
| List all clouds | `juju list-clouds` | `juju clouds` | User must know noun-only convention |
| Show cloud details | `juju cloud <name>` | `juju show-cloud <name>` | DE013 would prefer `juju cloud <name>` |
| Add a relation | `juju add-relation` | `juju integrate` | `add-relation` was removed in favor of `integrate` |
| Remove a relation | `juju remove-relation` | `juju remove-relation` | OK, but inconsistent with `integrate` |
| Create a model | `juju create-model` | `juju add-model` | `add` is used instead of `create` |
| Delete a model | `juju delete-model` | `juju destroy-model` | `destroy` instead of `delete` or `remove` |
| Delete a controller | `juju delete-controller` | `juju destroy-controller` | Same issue |
| Force delete controller | `juju destroy-controller --force` | `juju kill-controller` | Separate command instead of flag |
| Update application charm | `juju update-application` | `juju refresh` | Domain-specific verb |
| List firewall rules | `juju list-firewall-rules` | `juju firewall-rules` | Noun-only list |
| Show firewall rules | `juju show-firewall-rules` | `juju firewall-rules` | Same command for list and show |
| List disabled commands | `juju list-disabled-commands` | `juju disabled-commands` | Noun-only list |
| Show backup details | `juju show-backup` | N/A | No show-backup command exists |
| Remove a backup | `juju remove-backup` | N/A | No remove-backup command exists |
| List operations | `juju list-operations` | `juju operations` | Noun-only list |
| Show operation details | `juju show-operation` | `juju show-operation` | OK |
| List tasks | `juju list-tasks` | N/A | Tasks shown via `show-task` or `operations` |
| Cancel an operation | `juju cancel-operation` | `juju cancel-task` | Operation vs task terminology mismatch |

### Cross-Model Relation Confusion

Cross-model relation commands are scattered across domains:
- `offer`, `remove-offer`, `show-offer`, `find-offers`, `offers` → `crossmodel` package
- `consume`, `remove-saas`, `integrate`, `remove-relation`, `suspend-relation`, `resume-relation` → `application` package

Users looking for cross-model functionality may not discover `consume` or `integrate` because they are not grouped under an obvious `offer` or `relation` namespace.

## Ecosystem Comparison

### vs kubectl

| Aspect | kubectl | Juju |
|---|---|---|
| Hierarchy | `verb resource` (e.g., `kubectl get pods`) | Flat (`juju show-unit`) |
| Resource types | Explicit (`pods`, `services`, `deployments`) | Implicit or embedded in command name |
| CRUD verbs | `get`, `create`, `apply`, `delete`, `patch` | Fragmented (`show`, `add`, `deploy`, `remove`, `destroy`, `update`, `refresh`) |
| Namespace scoping | `-n namespace` | `-m model` (similar) |
| Output formats | `-o yaml/json/wide` | `--format yaml/json/tabular` (similar) |

**Observation**: kubectl's explicit resource-type model is more scalable and discoverable. Juju's flat namespace becomes unwieldy at ~157 commands.

### vs AWS CLI

| Aspect | AWS CLI | Juju |
|---|---|---|
| Hierarchy | `service command` (e.g., `aws ec2 describe-instances`) | Flat |
| Verb consistency | `describe`, `create`, `delete`, `update` across services | Inconsistent verb choice per domain |
| Discovery | Service-oriented help | Flat alphabetical list |

**Observation**: AWS CLI uses service prefixes to group commands. Juju has no equivalent grouping mechanism, making discovery harder.

### vs snap

| Aspect | snap | Juju |
|---|---|---|
| Primary objects | Snaps | Applications, models, controllers, clouds |
| Grammar | `snap <verb> [<snap>]` | `juju <verb-noun> [<args>]` |
| List command | `snap list` | `juju models`, `juju controllers` (plural nouns) |
| Info command | `snap info <snap>` | `juju show-application <app>` |

**Observation**: snap follows DE013 closely with clear verb-primary-object grammar. Juju deviates with verb-only commands and fragmented verb choices.

## Recommendations

### 1. Consolidate Deletion Verbs (High Impact)

**Problem**: `remove`, `destroy`, and `kill` are used inconsistently for deletion.

**Recommendation**: Standardize on `remove` for all deletions per DE013 §Grammar.
- `destroy-controller` → `remove-controller` (deprecate `destroy-controller`)
- `destroy-model` → `remove-model` (deprecate `destroy-model`)
- `kill-controller` → `remove-controller --force` (deprecate `kill-controller`)

**Deprecation plan** (per `.github/skills/cli-review/resources/deprecation.md`):
- Minor release: Add `remove-controller` and `remove-model` as aliases. Emit deprecation warning for `destroy-controller` and `destroy-model`.
- Major release: Remove `destroy-*` aliases, fail with error message.
- Next major: Clean up error messaging.

**Backward compat**: High breaking risk for scripts. Migration cost is high due to widespread use.

### 2. Converge on `create` for Creation (High Impact)

**Problem**: `add`, `create`, and `deploy` all create resources.

**Recommendation**:
- Keep `deploy` for applications (well-established ecosystem term)
- `add-cloud` → `create-cloud` (deprecate `add-cloud`)
- `add-model` → `create-model` (deprecate `add-model`)
- `add-unit` → `scale-application` or keep `add-unit` (both exist already)
- Standardize new commands on `create-*`

**Backward compat**: Medium risk. `add-*` is deeply ingrained.

### 3. Add Noun to Verb-Only Commands (Medium Impact)

**Problem**: `deploy`, `expose`, `refresh`, `trust`, `consume`, `integrate`, `migrate`, `offer` lack explicit nouns.

**Recommendation**: Add aliases with explicit nouns:
- `deploy-application` → alias for `deploy`
- `expose-application` → alias for `expose`
- `refresh-charm` → alias for `refresh`
- `trust-application` → alias for `trust`
- `consume-offer` → alias for `consume`
- `integrate-applications` → alias for `integrate`
- `migrate-model` → alias for `migrate`
- `offer-endpoint` → alias for `offer`

These are non-breaking additions.

### 4. Introduce `unset` Commands (Medium Impact)

**Problem**: `set-constraints`, `set-credential`, `set-firewall-rule`, `set-model-constraints` have no reversal.

**Recommendation**: Add `unset-*` commands:
- `unset-constraints`
- `unset-model-constraints`
- `unset-credential`
- `unset-firewall-rule`

These are new commands with no breaking changes.

### 5. Rename `resolved` to `resolve` (Low Impact)

**Problem**: `resolved` is grammatically awkward (past participle as imperative).

**Recommendation**: Make `resolve` the primary name and `resolved` the alias (currently the reverse). This is a documentation and help-text change with minimal code impact.

### 6. Improve Config Command Discoverability (Medium Impact)

**Problem**: `config`, `model-config`, `controller-config`, `model-defaults` are confusingly similar.

**Recommendation**: Add scoped aliases:
- `application-config` → alias for `config`
- `model-config` already exists
- `controller-config` already exists

Update help text to cross-reference related commands.

### 7. Group Cross-Model Commands (Medium Impact)

**Problem**: Cross-model relation commands are scattered.

**Recommendation**: Either:
- (a) Create aliases under a `crossmodel-` prefix: `crossmodel-offer`, `crossmodel-consume`
- (b) Add a help topic "cross-model-relations" that lists all related commands
- (c) Add `relate` as a primary alias for `integrate` (already exists) and improve documentation

Option (b) is non-breaking and recommended.

### 8. Add `list` Prefix Aliases (Low Impact)

**Problem**: Noun-only list commands violate user expectations from other CLIs.

**Recommendation**: Add `list-*` aliases for all noun-only list commands:
- `list-machines` → already exists
- `list-controllers` → already exists
- `list-secrets` → already exists
- `list-clouds` → add alias for `clouds`
- `list-models` → already exists
- etc.

These are already partially done. Complete the set.

### 9. Standardize `show` Over Noun-Only Details (Low Impact)

**Problem**: `info`, `dashboard` use noun-only instead of `show-*`.

**Recommendation**: Add aliases:
- `show-dashboard` → alias for `dashboard`
- `show-info` → less useful since `info` requires charm argument

### 10. Merge `debug-hooks` and `debug-code` Documentation (Low Impact)

**Problem**: These commands are nearly identical in function but have different names.

**Recommendation**: Update help text for both to explicitly state the difference (hooks vs actions). Consider deprecating one if they truly overlap.

## Tradeoff Summary

| Recommendation | Breaking Change? | Migration Cost | User Benefit |
|---|---|---|---|
| 1. Consolidate deletion verbs | Yes | High | High |
| 2. Converge on `create` | Yes | Medium | High |
| 3. Add explicit noun aliases | No | None | Medium |
| 4. Add `unset` commands | No | None | Medium |
| 5. Rename `resolved` | Minor | Low | Low |
| 6. Config cross-references | No | None | Medium |
| 7. Cross-model help topic | No | None | Medium |
| 8. `list-*` aliases | No | None | Low |
| 9. `show-*` aliases | No | None | Low |
| 10. Debug command clarity | No | None | Low |
