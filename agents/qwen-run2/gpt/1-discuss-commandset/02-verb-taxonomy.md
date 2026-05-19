# 02. Verb Taxonomy And Aspect Classification

## Scope note

To keep every command represented, the taxonomy includes two special cases:

- `chat` is treated as a standalone verb even though it is an orphan in the decomposition matrix.
- `qwen36.server` is mapped to the inferred operational verb `serve`, because the surface itself is a service identifier rather than a user verb.

Total commands covered by examples: 7

## Verb taxonomy

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|
| `chat` | execution | atelic | partial | exit/end session is external, not a named CLI command | `qwen36 chat` |
| `completion` | observation | punctual | no | none | `qwen36 completion bash` |
| `get` | observation | punctual | partial | `set` | `qwen36 get http.port`, `qwen36 get verbose` |
| `serve` | execution | atelic | partial | stop/start are external via snap service controls | `qwen36.server` |
| `set` | mutation | punctual | partial | `get` for inspection, `unset` missing for reversal | `qwen36 set http.port=8326`, `qwen36 set verbose=true` |
| `show` | observation | punctual | no | none | `qwen36 show-engine` |
| `use` | mutation | punctual | partial | rerun `use-engine` with another value; no named inverse | `qwen36 use-engine cpu`, `qwen36 use-engine cuda`, `qwen36 use-engine --auto` |

## Classification notes

### Intent groups

- `observation`: `completion`, `get`, and `show` all reveal information rather than mutate persistent state.
- `mutation`: `set` and `use` both change persisted runtime behavior.
- `execution`: `chat` and `serve` initiate or sustain running processes.

### Aspect

- `punctual` fits the admin verbs because they represent discrete state changes or queries.
- `atelic` fits `chat` and `serve` because both describe continuing activities rather than one-shot actions.

### Reversibility

- None of the qwen36 verbs has a perfect named inverse.
- The closest paired relation is `get` / `set`, but that is inspect-versus-mutate, not true reversal.
- `set` lacks the DE013-standard `unset` partner.
- `use` is reversible only by overwriting the chosen engine with another value.

## Observed strengths

1. The verb inventory is very small, which keeps learnability high.
2. Most verbs are common CLI vocabulary and easy to understand in isolation.
3. The split between observation and mutation is generally clear.

## Observed weaknesses

1. `use` is semantically weaker than `set` for persistent engine selection.
2. `show-*` conflicts with the DE013 preference for `foobar` or `status` forms when showing state.
3. The surface lacks a full reversible mutation vocabulary, especially `unset`.
4. `completion` is functionally valid but grammatically distinct from the rest of the command set.

## DE013 alignment notes

- DE013 says commands should be verbs, which qwen36 mostly satisfies.
- DE013 also provides standard verbs such as `get`, `set`, and `unset`; qwen36 matches two of the three.
- For showing state, DE013 prefers `status` over `show-status` and `foobar` over `show-foobar` for specific secondary objects, which makes `show-engine` the clearest naming outlier.
- For state mutation of a secondary object, a more specific verb such as `set-engine` would align better than `use-engine`.