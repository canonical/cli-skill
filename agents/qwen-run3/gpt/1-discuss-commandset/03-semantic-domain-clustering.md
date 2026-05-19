# Semantic Domain Clustering

## Scope

Every public command appears in exactly one domain. Source command count: 12.

## Domain Table

| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---:|---|---|---|
| Configuration | 3 | `get`, `set`, `unset` | yes | Strongest domain. Standard read/write/revert verbs. |
| Engine management | 3 | `list-engines`, `show-engine`, `use-engine` | partial | Good noun consistency, but `use` is weaker than `select`, and the listing form uses `list-*` rather than plural shorthand. |
| Interaction | 2 | `chat`, `webui` | no | `chat` is verb-led, `webui` is noun-led. They are both launch-oriented user entrypoints but follow different grammar. |
| System inspection | 3 | `status`, `show-machine`, `version` | partial | `status` and `version` are noun exceptions; `show-machine` is verb-noun. |
| Maintenance | 1 | `prune-cache` | yes | Clear single-purpose maintenance command, but the noun is implementation-oriented. |

Self-check: domain counts sum to 12.

## Domain Notes

### Configuration

- noun is implicit rather than named in the command
- behavior is symmetrical enough for everyday use
- the biggest issue is not naming but discoverability of valid keys

### Engine management

- all three commands act on the same conceptual resource
- users can enumerate, inspect, and choose engines
- hidden `debug select-engine` reveals a competing verb vocabulary for the same domain

### Interaction

- `chat` is easy to understand
- `webui` is less explicit because the verb is missing
- this is the weakest naming cluster in the visible public surface

### System inspection

- `status` and `version` fit conventional CLI exceptions
- `show-machine` is clear but could also have been shortened to `machine` if the project wanted to follow DE013 shorthand more aggressively

### Maintenance

- `prune-cache` is easy to read once the user already knows that engines are distributed as snap components
- for first-time users, `cache` is more implementation-oriented than task-oriented

## Findings

1. The command set clusters cleanly into five small domains, which is appropriate for a compact CLI.
2. Configuration is the best-shaped domain.
3. Interaction is the least consistent domain because it mixes verb-led and noun-led commands.
4. Engine management is mostly coherent but has a visible `use` versus `select` vocabulary split between public and debug surfaces.
