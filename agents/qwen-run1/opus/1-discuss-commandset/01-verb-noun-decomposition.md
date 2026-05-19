# 01 — Verb-Noun Decomposition Matrix

## Command Inventory (6 commands)

| # | Full Command | Verb | Noun |
|---|---|---|---|
| 1 | `chat` | chat | *(implicit: model/session)* |
| 2 | `use-engine` | use | engine |
| 3 | `show-engine` | show | engine |
| 4 | `get` | get | *(configuration)* |
| 5 | `set` | set | *(configuration)* |
| 6 | `completion` | completion | *(shell)* |

## Verb × Noun Grid

|  | engine | configuration | shell | session |
|---|---|---|---|---|
| **chat** | — | — | — | ✓ |
| **completion** | — | — | ✓ | — |
| **get** | — | ✓ | — | — |
| **set** | — | ✓ | — | — |
| **show** | ✓ | — | — | — |
| **use** | ✓ | — | — | — |

## Annotations

### Incomplete CRUD Sets

**Engine domain**:
- Has `show-engine` (read) and `use-engine` (update/select)
- Missing: no `list-engines` or `engines` command to discover available engines without knowing names
- Missing: no `add-engine` / `remove-engine` (though engines are snap components, not user-manageable)

**Configuration domain**:
- Has `get` (read) and `set` (write)
- Missing: no `unset` or `reset` command to restore defaults
- Missing: no `list` or `config` command to show all current configuration

### Verb Inconsistencies

- **`use-engine`** uses `use` as the verb for selection/activation. This is non-standard in Canonical CLI vocabulary — the standard would be `set` (for config-like changes) or `switch` (for toggling between options). Per DE013, `enable/disable` is the standard for feature toggles.
- **`show-engine`** uses `show` correctly per DE013 standards for displaying state.

### Orphan Commands

- **`chat`**: Does not decompose into verb-noun cleanly. It's a bare verb that implicitly operates on a "session" or "model conversation." Per DE013, this is acceptable for primary-type actions on single-purpose tools.
- **`completion`**: Functions as a noun (generates shell completions), not a verb. More accurately, this would be `generate-completion` or follow the pattern `completion <shell>`.
- **`get`** / **`set`**: These are verbs without explicit nouns — the noun "configuration" is implied. Per DE013, `get/set` is the standard pattern for configuration access.
