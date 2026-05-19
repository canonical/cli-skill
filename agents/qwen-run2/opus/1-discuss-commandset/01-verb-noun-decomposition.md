# Verb-Noun Decomposition Matrix

## Decomposition Table

| Command | Verb | Noun |
|---------|------|------|
| `chat` | chat | *(implicit: model/session)* |
| `use-engine` | use | engine |
| `show-engine` | show | engine |
| `get` | get | *(config key)* |
| `set` | set | *(config key)* |
| `completion` | completion | *(shell)* |

## Verb × Noun Grid

|  | engine | config | shell | session |
|--|--------|--------|-------|---------|
| **chat** | — | — | — | ✓ |
| **use** | ✓ | — | — | — |
| **show** | ✓ | — | — | — |
| **get** | — | ✓ | — | — |
| **set** | — | ✓ | — | — |
| **completion** | — | — | ✓ | — |

## Annotations

### Incomplete CRUD Sets

| Noun | Available Verbs | Missing Verbs | Notes |
|------|----------------|---------------|-------|
| engine | use, show | list, add, remove | No way to list available engines; no add/remove (engines are snap components) |
| config | get, set | unset, list | No `unset` to restore defaults; no `list` to show all keys — per DE013, `unset` is recommended for restoring-to-default |
| shell (completion) | completion | — | Single-purpose, no lifecycle needed |
| session (chat) | chat | — | Single-purpose, interactive only |

### Verb Inconsistencies

| Issue | Detail |
|-------|--------|
| `use` vs `set` | `use-engine` selects the engine (a kind of "set"), but `set` is reserved for config keys. The verb `use` implies ongoing state rather than a point-in-time write. This is a reasonable distinction. |
| `show` vs `get` | `show-engine` displays engine info (YAML dump), while `get` retrieves a single config value. The semantic difference is "display structured object" vs "read scalar value". Per DE013 §Commonly Used Commands, `show` is for instance details while `get/set` is for configuration. |
| `completion` is a noun | The `completion` command uses a noun as the command name rather than a verb (e.g., `generate-completion` or simply exposing it via the completer mechanism only). |

### Orphan Commands

| Command | Issue |
|---------|-------|
| `completion` | Does not decompose into verb-noun; it is a noun used as a command. Per DE013, commands should be verbs. However, this is a common pattern in CLI tools (e.g., `kubectl completion`). |
| `chat` | Semi-orphan — `chat` is both a verb and a noun. As a verb, it means "start chatting." It doesn't act on a named resource type in the same way other commands do. |
