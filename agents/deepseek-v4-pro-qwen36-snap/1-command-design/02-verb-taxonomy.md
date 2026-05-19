# 02 — Verb Taxonomy and Aspect Classification

All 15 unique verbs from the decomposition matrix, classified.

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|------|-------------|--------|------------|-------------|--------------|
| `chat` | execution | atelic | no | — | `chat`, `debug chat` |
| `get` | observation | punctual | partial | `set` | `get` |
| `list` | observation | punctual | no | — | `list-engines` |
| `prune` | lifecycle | telic | no | — | `prune-cache` |
| `run` | execution | atelic | no | — | `run` |
| `select` | lifecycle | telic | partial | — | `debug select-engine` |
| `serve` | execution | atelic | no | — | `serve-webui`, `debug serve-webui` |
| `set` | mutation | punctual | partial | `unset` | `set` |
| `show` | observation | punctual | no | — | `show-engine`, `show-machine` |
| `status` | observation | punctual | no | — | `status` |
| `unset` | mutation | punctual | partial | `set` | `unset` |
| `use` | lifecycle | punctual | partial | — | `use-engine` |
| `validate` | observation | punctual | no | — | `debug validate-engines` |
| `version` | observation | punctual | no | — | `version` |
| `webui` | execution | punctual | no | — | `webui` |

---

## Intent Group Distribution

| Intent Group | Count | Verbs |
|-------------|-------|-------|
| **observation** | 6 | `get`, `list`, `show`, `status`, `validate`, `version` |
| **execution** | 4 | `chat`, `run`, `serve`, `webui` |
| **lifecycle** | 3 | `prune`, `select`, `use` |
| **mutation** | 2 | `set`, `unset` |
| **access** | 0 | — |
| **transfer** | 0 | — |
| **migration** | 0 | — |

---

## Aspect Distribution

| Aspect | Count | Verbs |
|--------|-------|-------|
| **punctual** | 10 | `get`, `list`, `set`, `show`, `status`, `unset`, `use`, `validate`, `version`, `webui` |
| **atelic** | 4 | `chat`, `run`, `serve`, `prune` |
| **telic** | 1 | `select` |

---

## Reversibility Analysis

### Paired verbs:
- **`set` ↔ `unset`**: Properly symmetric. Both are punctual, both in the mutation group.

### Verbs with no inverse:
- **`prune`**: Destructive and irreversible (cannot "un-prune"). This is expected — deleting components is permanent.
- **`chat`**, **`run`**, **`serve`**, **`show`**, **`list`**, **`status`**, **`validate`**, **`version`**, **`webui`**: These are observation or execution verbs that don't have meaningful inverses.

### Verbs with partial reversibility:
- **`get` → `set`**: `get` reads, `set` writes. Not technically an inverse pair but complementary. The true inverse of `set` is `unset`.
- **`use`**: Partially reversible — `use-engine` to switch to engine A undoes the activation of engine B, but there is no standalone "deactivate engine" command.
- **`select`**: Partially reversible — only through a subsequent `use-engine` call or by manual manifest modification. No `deselect` command.

---

## Key Findings

1. **Observation-heavy**: 6 of 15 verbs (40%) are observation. This reflects the CLI's primary role as a management/monitoring tool — users check status far more often than they change state.

2. **No access verbs**: There are no `grant`, `revoke`, `enable`, `disable`, `login`, `logout` commands. The CLI has no authentication or authorization model beyond the `sudo` requirement for mutation commands.

3. **No transfer/migration verbs**: The CLI doesn't handle copying, syncing, or migrating models or data between systems. All model management is via the snap component system.

4. **Punctual-dominant**: Most verbs are punctual (instant state reads or writes). Only 4 are atelic (ongoing: `chat`, `run`, `serve`, `prune`). This is appropriate for a management CLI — long-running operations are delegated to the inference server, not the CLI.

5. **The `webui` verb misclassification**: `webui` is classified as execution/punctual, but it's actually a noun used as a command. It should be decomposed as `open` × `webui` or `launch` × `webui` for consistency with the verb-noun pattern.

6. **Verb fragmentation in the debug namespace**: `debug` duplicates the `chat` and `serve` verbs from the top level, creating ambiguity. `debug chat` and `chat` have different semantics (one uses active engine, the other requires `--base-url`).