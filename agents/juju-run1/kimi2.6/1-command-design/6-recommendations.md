## Section 6: Pattern Classification and Recommendations

### Pattern Classification

- **Primary grouping pattern**: Flat top-level `verb-noun` hyphenation covering ~55% of commands.
- **Secondary patterns**:
  - Noun-plural list shorthands (~16 commands)
  - Bare verbs for high-frequency or domain-specific operations (~25 commands)
  - Noun-first configuration hybrids (~10 commands)
  - Embedded compound commands (`enable-destroy-controller`, `change-user-password`)
- **Nesting depth**: Strictly flat. No command groups or subcommand namespaces.
- **Style**: Mixed legacy and modern. Modern additions follow `verb-noun` more consistently; older commands retain noun-first and bare-verb forms.

### Discoverability Assessment

- **Predicted user paths**: A user wanting to "delete an application" will likely try `juju remove-application` (works) or `juju delete-application` (fails). They may also try `juju destroy-application` (fails) because `destroy` is used for model/controller.
- **Tab-completion burden**: 150+ flat commands make tab-completion noisy. Discoverability is reduced because related commands are not namespaced.
- **Help surface**: `juju help` lists all commands. Domain clustering is not explicit in help output.
- **Orphan risk**: Commands like `config`, `status`, `machines`, `integrate` do not follow predictable patterns, increasing learning curve.

### Ecosystem Comparison

| Tool | Pattern | Depth | Key Difference from Juju |
|---|---|---|---|
| **kubectl** | verb-noun | 1-2 levels (resource types) | Uses `get/describe/delete/apply` uniformly; resource types are nouns after the verb (`kubectl get pods`). Juju mixes verbs.
| **awscli** | noun-verb | 2 levels (`aws <service> <action>`) | Services are grouped; actions inside services. Juju is flat.
| **lxc** | verb-noun | 1-2 levels | Recent redesign toward uniform verb-noun; deprecating mixed flags. Juju has more historical baggage.
| **snap** | verb-noun | flat | Very consistent `install/remove/list/info` verbs; noun-plural lists. Juju is similar but less disciplined.
| **docker** | verb-noun | 1-2 levels (with object type) | `docker container ls`, `docker image rm`. Grouped by object. Juju lacks this grouping.

### Recommendations

| # | Recommendation | Rationale | Backward Compat | Migration Cost |
|---|---|---|---|---|
| 1 | **Standardize delete verb to `remove`** | `destroy-controller` and `destroy-model` are the only `destroy-*` commands. Rename to `remove-controller` and `remove-model`, keep `destroy-*` as deprecation aliases. | Breaks scripts using `destroy-*` until aliases are retired | Low-medium; one deprecation cycle |
| 2 | **Introduce `add-application` as alias to `deploy`** | `deploy` is the outlier for application creation. Add `add-application` as a supported alias to align with the rest of the CRUD surface. | Fully backward compatible | Very low |
| 3 | **Rename `kill-controller` to `--force` on `destroy-controller`** | `kill` is a separate top-level command for a rare force path. Fold it into `destroy-controller --force` to reduce verb proliferation. | `kill-controller` must be deprecated | Medium; requires flag additions |
| 4 | **Add `list-*` aliases for all noun-plural orphans** | `machines` → `list-machines`, `models` → `list-models`, etc. Support both forms; `list-*` becomes the canonical documented form. | Fully backward compatible | Low |
| 5 | **Create a `config` namespace or consistent `--scope` pattern** | `config`, `model-config`, `controller-config`, and `model-defaults` are confusingly similar. Consider `juju config --scope=model` or a subcommand pattern, while keeping aliases. | Breaking if old forms removed; safe if kept as aliases | Medium |
| 6 | **Add `show-relation` command** | Relations have `integrate` and `remove-relation` but no `show-relation`. This is a pure gap. | Fully backward compatible | Very low |
| 7 | **Add `--help` topics by domain** | Group help output by domain (cloud, model, application, etc.) rather than a single alphabetical list. Improve discoverability without changing commands. | Fully backward compatible | Low |
| 8 | **Deprecate `enable-destroy-controller` in favor of `enable command destroy-controller`** or `reenable-destroy` | Embedding a command name inside another command is unprecedented and confusing. | Requires deprecation cycle | Medium |
| 9 | **Add `trust`/`untrust` pair** | `trust` exists but no `untrust`. Users must set `trust=false` via flags or config. Add an explicit symmetric command. | Fully backward compatible | Very low |
| 10 | **Audit and align `download` / `get` / `show` for Charmhub** | `download`, `info`, and `find` all query Charmhub but use inconsistent base verbs. Consider grouping under a `charm` or `hub` namespace. | Breaking if changed; safe as aliases | Medium |

---

## Key Findings (Top 5)

1. **Verb asymmetry in deletion**: `destroy-*` is used only for controller and model, while `remove-*` is used for ~15 other resources. This is the most glaring inconsistency.
2. **`deploy` is a singleton create verb**: Applications are created with `deploy`, breaking the `add-*` pattern universally used elsewhere.
3. **~35 orphan commands** do not follow `verb-noun`, including critical ones like `config`, `status`, `machines`, and `integrate`. These increase learning friction.
4. **High confusion risk between `exec` and `run`**: Both execute on units but have different semantics (shell vs action). This is a documented support pain point.
5. **No command grouping/namespacing**: With ~150 flat commands, tab completion and help output are unwieldy. Resources like `cloud`, `model`, and `application` have no structural grouping.

---

## Tradeoffs Summary

| Change | Backward Compatible? | Script Impact | Human UX Impact | Effort |
|---|---|---|---|---|
| `destroy-*` → `remove-*` aliases | Yes (with deprecation) | Medium | High clarity gain | Medium |
| `add-application` alias | Yes | None | Medium discoverability gain | Low |
| `kill-controller` → `--force` | Requires deprecation | Low | Reduces dangerous verb collision | Medium |
| `list-*` aliases | Yes | None | High for newcomers | Low |
| Config scope consolidation | Yes (with aliases) | Medium | High clarity gain | Medium |
| Domain-grouped help | Yes | None | High discoverability gain | Low |
