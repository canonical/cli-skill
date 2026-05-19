# 06 — Pattern Classification and Recommendations

## Pattern Classification

### Primary Grouping Pattern

Juju uses a **flat command hierarchy** with a single level of commands beneath the `juju` root. There is no subcommand nesting. All ~130 commands are peers.

- **Pattern type**: Flat verb-noun (mixed with verb-only for core operations)
- **Depth**: 1 (root + commands, no nested groups)
- **Style**: Mixed: ~55% verb-noun (`add-cloud`, `show-model`), ~30% verb-only (`deploy`, `status`, `expose`), ~12% plural-noun as list shorthand (`models`, `controllers`), ~3% other (`whoami`, `resolved`)

### Discoverability Assessment

**Current discovery paths for a new user:**

| User Intent | Expected Command | Actual Command | Discoverable? |
|-------------|-----------------|----------------|---------------|
| "Create a controller" | `add-controller` or `create-controller` | `bootstrap` | ❌ No — `bootstrap` is jargon |
| "List my controllers" | `list-controllers` | `controllers` | ✅ Yes — plural noun is standard |
| "Deploy an app" | `deploy` | `deploy` | ✅ Yes |
| "Create an app" | `add-application` | `deploy` | ❌ No — different verb family |
| "Remove an app" | `remove-application` | `remove-application` | ✅ Yes |
| "Set app config" | `config` or `app-config` or `set-config` | `config` | ⚠️ Partial — bare `config` is ambiguous |
| "Set model config" | `model-config` | `model-config` | ✅ Yes |
| "Set controller config" | `controller-config` | `controller-config` | ✅ Yes |
| "List models" | `models` or `list-models` | `models` | ✅ Yes |
| "Destroy a model" | `destroy-model` or `remove-model` | `destroy-model` | ⚠️ Partial — might look for `remove-model` |
| "Add a relation" | `add-relation` or `link` | `integrate` | ❌ No — `integrate` is non-obvious |
| "Remove a relation" | `remove-relation` | `remove-relation` | ✅ Yes |
| "Run a command on a unit" | `run` or `exec` | `run` or `exec` | ⚠️ Partial — two commands |
| "Show app details" | `show-application` or `info` | `show-application` | ✅ Yes |
| "List secrets" | `secrets` or `list-secrets` | `secrets` | ✅ Yes |
| "Show current context" | `whoami` or `context` or `info` | `whoami` | ⚠️ Partial — UNIX convention |
| "Switch context" | `switch` or `use` | `switch` | ✅ Yes |
| "Offer an endpoint" | `offer` | `offer` | ✅ Yes |
| "Consume an offer" | `consume` or `import-offer` | `consume` | ⚠️ Partial — `consume` is metaphorical |
| "Add a user" | `add-user` | `add-user` | ✅ Yes |
| "View firewall rules" | `firewall-rules` or `list-firewall-rules` | `firewall-rules` | ✅ Yes |
| "Create a backup" | `create-backup` or `backup` | `create-backup` | ✅ Yes |
| "Open dashboard" | `dashboard` or `ui` or `gui` | `dashboard` | ✅ Yes |

**Discovery score**: ~65% of common intents map directly to discoverable command names.

### Help & Completion Support

- `juju help commands` — flat alphabetical listing of all 130 commands, no grouping
- `juju help <cmd>` — detailed per-command help with examples, flags, SeeAlso
- Tab completion: supported but must be configured by shell plugin
- Fuzzy matching: `FindClosestSubCommand` provides "Did you mean X?" suggestions
- `juju help topics` — only 1 topic registered (`basics`)

## Ecosystem Comparison

### Comparison: `kubectl`

| Dimension | juju | kubectl |
|-----------|------|---------|
| Pattern | Flat verb-noun | Flat verb-noun |
| Grouping | None | Resource group filtering (`kubectl get pods,services`) |
| Creation verbs | `add`, `create`, `deploy`, `bootstrap` | `create`, `apply`, `run` |
| Observation | `get`, `show`, `list-*` nouns | `get`, `describe` |
| Plugins | PATH-based `juju-*` | PATH-based `kubectl-*` |
| Config | Scoped commands | `config set-context`, `config use-context` |
| Depth | 1 | 1 (plus sub-resources like `rollout`) |

**Key difference**: kubectl groups resources (pods, services, deployments) with a uniform `get`/`describe`/`delete` pattern. Juju uses different verbs per resource type, requiring users to learn domain-specific vocabulary.

### Comparison: `snap`

| Dimension | juju | snap |
|-----------|------|------|
| Pattern | Flat verb-noun | Flat verb/verb-noun |
| Creation | `deploy`, `add-*`, `bootstrap` | `install` |
| Removal | `remove-*`, `destroy-*` | `remove` |
| Config | Scoped per level | `set`, `get`, `unset` |
| Grouping | None | None (but simpler: ~30 commands) |

**Key difference**: With ~30 commands, snap can afford a flat structure. With ~130 commands, juju has outgrown the flat approach but hasn't introduced grouping.

### Comparison: `lxc` (LXD)

| Dimension | juju | lxc |
|-----------|------|------|
| Pattern | Flat | Some nesting (`lxc config`, `lxc network`, `lxc storage`) |
| Grouping | None | Domain-based grouping |
| Depth | 1 | 2 (group + subcommand) |

**Key difference**: LXD groups commands by resource domain (config, network, storage, image, etc.), making discovery easier than Juju's flat list. LXD's grouping is a closer fit for Juju's domain organization.

## Recommendations

### Recommendation 1: Introduce Command Grouping in Help Output

**Rationale**: With 130 flat commands, `juju help commands` is overwhelming. Grouping by domain would dramatically improve discoverability without changing any command names.

**Proposed grouping** (matching current source organization):

```
juju help commands    →   Shows:
  Cloud & Credentials      add-cloud, update-cloud, remove-cloud, clouds, ...
  Controllers              bootstrap, controllers, show-controller, ...
  Models                   add-model, models, show-model, ...
  Applications             deploy, add-unit, remove-unit, config, ...
  Relations                integrate, remove-relation, suspend-relation, ...
  Cross-Model              offer, remove-offer, consume, find-offers, ...
  Machines                 add-machine, remove-machine, machines, ...
  Storage                  add-storage, attach-storage, storage, create-storage-pool, ...
  Spaces & Networks        add-space, remove-space, spaces, subnets, ...
  Users & Access           add-user, remove-user, grant, revoke, login, ...
  Secrets                  add-secret, remove-secret, secrets, grant-secret, ...
  Actions & Execution      run, exec, actions, operations, cancel-task, ...
  SSH & Debugging          ssh, scp, debug-hooks, debug-code, debug-log, ...
  Security & Firewall      set-firewall-rule, firewall-rules, ...
  Monitoring & Status      status, show-status-log, ...
  Backups                  create-backup, download-backup
  CharmHub                 find, info, download
  Resources                attach-resource, resources, charm-resources
  Dashboard                dashboard
  Tooling                  version, help-action-commands, help-hook-commands, sync-agent-binary, migrate, switch
```

**Backward compatibility**: ✅ Full. No command names change. Only help output format changes. **Migration cost**: Low — update `juju help commands` formatting logic. No user scripts affected.

### Recommendation 2: Rename `config` → `application-config`

**Rationale**: Per DE013 §Grammar: "When verbs alone are not sufficient to distinguish between objects, use the verb-noun form." The bare `config` is ambiguous with `model-config` and `controller-config`. Renaming to `application-config` makes the scope explicit and aligns with the `model-config`/`controller-config` pattern.

**Deprecation plan** (per deprecation spec):
1. **Minor version**: Add `application-config` as primary name. Keep `config` as deprecated alias with warning: `warning: "config" is deprecated, use "application-config" instead`
2. **Next major**: Remove `config`, fail with: `error: "config" was removed in 4.0, use "application-config" instead`
3. **N+1 major**: Remove error message, show: `error: invalid command "config"`

**Backward compatibility**: ⚠️ Breaking (mitigated by deprecation period). Scripts using `juju config <app>` will need updating during deprecation window. **Migration cost**: Medium — scripts must update command name. One full release cycle of deprecation period.

### Recommendation 3: Rename `constraints` → `application-constraints` and `set-constraints` → `set-application-constraints`

**Rationale**: Same scope ambiguity issue as `config`. `constraints` and `set-constraints` operate on applications but bare names look like model-level operations.

**Deprecation plan**: Same phased approach as Recommendation 2.

**Backward compatibility**: ⚠️ Breaking (mitigated). **Migration cost**: Medium.

### Recommendation 4: Add Missing Reverse Operations

**Rationale**: Several domains lack symmetric reverse commands (see Symmetry Audit). Adding these would complete the CRUD coverage.

| Priority | New Command | Reverses |
|----------|------------|----------|
| **High** | `remove-firewall-rule` | `set-firewall-rule` |
| **High** | `untrust` | `trust` |
| **Medium** | `remove-backup` | `create-backup` |
| **Medium** | `detach-resource` | `attach-resource` |
| **Low** | `export-ssh-key` | `import-ssh-key` |

**Backward compatibility**: ✅ Full. Adding new commands does not break existing ones. **Migration cost**: None — purely additive.

### Recommendation 5: Consider Renaming `integrate` → `relate` (or standardize)

**Rationale**: `integrate` is a technical term; `relate` is more intuitive and already exists as an alias. Per DE013: "Commands are verbs" — `relate` is clearer semantically. The inverse (`remove-relation`) uses the noun `relation`. Using `relate` would create the pair `relate`/`remove-relation` which is imperfect.

**Alternative**: Rename `integrate` → `add-relation` to match the CRUD pattern used across the rest of the CLI, creating the clean pair `add-relation`/`remove-relation`. This aligns with DE013's "verb-noun when verbs alone are insufficient."

**Deprecation plan**: Same phased approach.

**Backward compatibility**: ⚠️ Breaking. `integrate` is widely used in documentation and tutorials. **Migration cost**: High — extensive ecosystem impact (docs, tutorials, blog posts, community knowledge).

### Recommendation 6: Resolve `--integrations` / `--relations` Duplication (status)

**Rationale**: Per DE013 §Flags: "Do not provide both forms." Having both `--integrations` and `--relations` as identical flags for `status` creates unnecessary cognitive load. Keep one.

**Recommendation**: Keep `--relations` and deprecate `--integrations` (or vice versa, but pick one). `--relations` is the more common term in the ecosystem.

**Backward compatibility**: ⚠️ Breaking (minor). **Migration cost**: Low — one flag removal with deprecation warning.

### Recommendation 7: Add Help Topics for Key Concepts

**Rationale**: Currently only 1 help topic (`basics`). Adding domain-specific topics (`juju help models`, `juju help relations`, `juju help secrets`) would guide users through key concepts.

**Proposed topics**:
- `models` — Model lifecycle and management
- `relations` — Integration patterns and CMR
- `secrets` — Secret management and backends
- `spaces` — Network spaces and bindings
- `storage` — Storage pools and filesystems
- `security` — Firewall rules, access control, trust
- `kubernetes` — K8s-specific operations

**Backward compatibility**: ✅ Full. **Migration cost**: None — additive.

### Recommendation 8: Standardize Create Verb (`add` vs `create`)

**Rationale**: Juju uses `add` for most creation (`add-cloud`, `add-model`) but `create` for some (`create-backup`, `create-storage-pool`). Per DE013, `create-foo` is the recommended form for secondary objects, but consistency matters more.

**Recommendation**: Standardize on `add-*` for consistency with the majority of commands. Rename `create-backup` → `add-backup`, `create-storage-pool` → `add-storage-pool`.

**Deprecation plan**: Phased approach with aliases.

**Backward compatibility**: ⚠️ Breaking. **Migration cost**: Medium — fewer commands affected than recommend 2/3.

## Summary of Recommendations by Impact

| # | Recommendation | Impact | Backward Compat? | Migration Cost |
|---|---------------|--------|-----------------|---------------|
| 1 | Command grouping in help | High - discoverability | ✅ Yes | Low |
| 2 | Rename `config` → `application-config` | High - clarity | ⚠️ Breaking | Medium |
| 3 | Rename `constraints` → `application-constraints` | Medium - clarity | ⚠️ Breaking | Medium |
| 4 | Add missing reverse operations | Medium - completeness | ✅ Yes | None |
| 5 | Standardize `integrate` → `add-relation` | High - consistency | ⚠️ Breaking | High |
| 6 | Resolve `--integrations`/`--relations` | Low - deduplication | ⚠️ Minor | Low |
| 7 | Add help topics | Medium - education | ✅ Yes | None |
| 8 | Standardize `create` → `add` | Low - consistency | ⚠️ Breaking | Medium |

## Tradeoffs

### Grouping vs Flat

**Pro**: Grouping dramatically improves discoverability for new users and narrows the command list for targeted help. **Con**: Existing users who have memorized the flat list may need a transition period. Mitigation: keep `juju help commands --all` for the flat list.

### Verb Consistency vs Ecosystem Familiarity

**Pro**: Standardizing on verb patterns (add/remove, show/list) makes the CLI predictable. **Con**: Some non-standard verbs (`deploy`, `bootstrap`, `integrate`) have strong ecosystem recognition. Changing them would disrupt established workflows. Recommendation: prioritize pattern consistency for new commands; deprecate legacy verbs slowly.

### DE013 Compliance vs Practical Constraints

The DE013 standard provides excellent guidance, but some Juju-specific conventions (`bootstrap`, `deploy`) are deeply embedded. Full compliance would require renaming 15+ commands, which is a major breaking change. A pragmatic approach: apply DE013 to new commands and incremental renames over multiple major versions.