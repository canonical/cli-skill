# Verb Taxonomy And Aspect Classification

## Scope

Matrix verb count from Section 1: 7.

Verbs classified here:

- `get`
- `list`
- `prune`
- `set`
- `show`
- `unset`
- `use`

Orphan commands outside the matrix but still present in the command set: `status`, `chat`, `webui`, `version`.

## Taxonomy Table

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|
| `get` | observation | atelic | no | — | `get http.port` |
| `list` | observation | atelic | no | — | `list-engines` |
| `prune` | lifecycle | telic | no | — | `prune-cache` |
| `set` | mutation | telic | yes | `unset` | `set http.port=8326` |
| `show` | observation | atelic | no | — | `show-engine`, `show-machine` |
| `unset` | mutation | telic | yes | `set` | `unset http.port` |
| `use` | access | telic | partial | — | `use-engine cpu`, `use-engine --auto` |

Self-check: 7 unique matrix verbs in Section 1, 7 rows here.

## Notes

### Strong verbs

- `get`, `set`, and `unset` are aligned with DE013 common commands and form the strongest micro-grammar in the CLI.
- `list` is standard and predictable.

### Weak or non-standard verbs

- `use` is understandable but semantically weak for selection. Hidden `debug select-engine` shows that `select` is the sharper verb in this domain.
- `prune` is meaningful to maintainers, but it is less predictable than `remove` or `clear` for general users.

### Observation pattern

- `show` is acceptable, but DE013 prefers noun shorthand or `*-status` where practical for observation of secondary objects.
- This makes `show-engine` and `show-machine` acceptable but not ideal from a standards perspective.

## Orphan Command Notes

| Command | Practical classification | Note |
|---|---|---|
| `status` | observation | canonical noun exception |
| `chat` | execution | direct interaction verb, no explicit object |
| `webui` | execution | effectively `open` or `launch`, but the literal command is noun-led |
| `version` | observation | canonical noun exception |

## Findings

1. The command set is dominated by observation and mutation, which fits a local control-plane CLI.
2. The config verbs are more coherent than the engine verbs.
3. The engine domain should either standardize on `use` everywhere or adopt `select` consistently. The current split between public and debug vocabulary is a real naming inconsistency.
