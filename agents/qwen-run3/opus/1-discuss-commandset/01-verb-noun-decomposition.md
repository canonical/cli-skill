# 01 — Verb-Noun Decomposition

## Decomposition Table

| # | Command | Verb | Noun |
|---|---------|------|------|
| 1 | `status` | show (implicit) | status |
| 2 | `chat` | chat | — (action-object merged) |
| 3 | `webui` | launch (implicit) | webui |
| 4 | `get` | get | config (implicit) |
| 5 | `set` | set | config (implicit) |
| 6 | `unset` | unset | config (implicit) |
| 7 | `list-engines` | list | engines |
| 8 | `show-engine` | show | engine |
| 9 | `use-engine` | use | engine |
| 10 | `show-machine` | show | machine |
| 11 | `prune-cache` | prune | cache |
| 12 | `version` | show (implicit) | version |
| 13 | `run` (hidden) | run | subprocess |
| 14 | `serve-webui` (hidden) | serve | webui |
| 15 | `debug` (hidden) | debug | — (namespace) |

## Verb × Noun Grid

|  | engine | config | machine | cache | status | version | webui | subprocess |
|--|--------|--------|---------|-------|--------|---------|-------|-----------|
| **show** | ✓ | — | ✓ | — | ✓ (implicit) | ✓ (implicit) | — | — |
| **list** | ✓ | — | — | — | — | — | — | — |
| **use** | ✓ | — | — | — | — | — | — | — |
| **get** | — | ✓ | — | — | — | — | — | — |
| **set** | — | ✓ | — | — | — | — | — | — |
| **unset** | — | ✓ | — | — | — | — | — | — |
| **prune** | — | — | — | ✓ | — | — | — | — |
| **chat** | — | — | — | — | — | — | — | — |
| **launch** | — | — | — | — | — | — | ✓ (implicit) | — |
| **run** | — | — | — | — | — | — | — | ✓ |
| **serve** | — | — | — | — | — | — | ✓ | — |

## Analysis

### Incomplete CRUD Sets

| Noun | Available Verbs | Missing Expected |
|------|----------------|-----------------|
| engine | show, list, use | — (engines are not created/deleted by users; appropriate) |
| config | get, set, unset | — (complete for config lifecycle) |
| cache | prune | show/list (no way to inspect cache contents before pruning) |
| machine | show | — (read-only, appropriate) |

### Verb Inconsistencies

1. **`status` vs `show-engine` vs `show-machine`**: The `status` command shows system state without the `show-` prefix, breaking the pattern. Per DE013 §Grammar, "For showing state without changing it, use the shorthand `status`" — so `status` is correct here, but it creates an inconsistency where `show-*` is used for component details while `status` is used for overall state.

2. **`version` without `show-` prefix**: Follows the DE013 convention ("tool version: show release version") correctly.

3. **`get/set/unset` vs `show-engine/use-engine`**: Config uses bare verbs (`get`, `set`) while engine management uses verb-noun compounds (`show-engine`, `use-engine`). Per DE013, this is correct — `get/set` act on the primary config object type, while engine is a secondary object using the verb-noun pattern.

### Orphan Commands

| Command | Why it doesn't decompose cleanly |
|---------|--------------------------------|
| `chat` | Both verb and noun simultaneously — "start a chat session." Not a CRUD operation. |
| `webui` | Noun used as a command — implied verb is "open" or "launch." |
| `debug` | Namespace/group header, not an action. |

### Recommendations

1. **`webui` should be a verb-form**: Per DE013, "Every command that acts on a primary object must be a verb." Consider `open-webui` or `launch-webui` to make the action explicit. However, since the Canonical standard also says "use the shorthand status over show-status" for showing state, a shorthand for launching could be acceptable. **Recommendation**: Rename to `open-webui` with `webui` as a deprecated alias (per deprecation spec §Next minor version).

2. **`prune-cache` naming**: The command actually removes snap components (installed engine/model files), not a regenerable cache. The verb "prune" is correct (selective removal), but the noun "cache" is misleading. Consider `remove-components` or keep `prune-cache` if the mental model of components as "cached engines" is intentional.
