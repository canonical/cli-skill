# 03. Semantic Domain Clustering

## Command accounting

Total commands analyzed: 7

Every command appears in exactly one domain below.

## Domain table

| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---|---|---|---|
| conversation | 1 | `qwen36 chat` | yes | Single-command domain for the interactive user workflow |
| engine | 2 | `qwen36 use-engine`, `qwen36 show-engine` | partial | Both refer to the same noun, but the verbs are not an ideal DE013 pair |
| configuration | 2 | `qwen36 get`, `qwen36 set` | mostly | Clear inspect/mutate pair, but missing `unset` and key enumeration |
| shell integration | 1 | `qwen36 completion bash` | yes | Single-purpose meta-command for shell completion generation |
| service runtime | 1 | `qwen36.server` | no | Operationally separate from the `qwen36 <verb>` hierarchy |

Count check: 1 + 2 + 2 + 1 + 1 = 7 commands.

## Domain notes

### Conversation

- Noun form: implicit rather than explicit
- CRUD coverage: not applicable
- Verb consistency: acceptable for a user-facing activity command

### Engine

- Noun form: consistently `engine`
- CRUD/lifecycle coverage: incomplete
- Present operations: inspect and select
- Missing likely operations: status alias, list/query of available engines, reset/unset of engine-managed state
- Verb consistency: mixed quality because `show` and `use` are understandable but not the strongest DE013-aligned pair

### Configuration

- Noun form: implicit config object rather than explicit `config` noun in command names
- CRUD coverage: partial
- Present operations: read and write
- Missing likely operations: `unset`, enumerate keys, inspect all effective config
- Verb consistency: strong within the tiny domain because `get` and `set` are standard verbs

### Shell integration

- Noun form: shell target is explicit (`bash`)
- CRUD coverage: not applicable
- Verb consistency: acceptable for a shell-completion meta-command

### Service runtime

- Noun form: explicit service name rather than verb grammar
- CRUD/lifecycle coverage: externally managed by snap, not by qwen36 command verbs
- Verb consistency: not applicable inside the user CLI because this is a service identifier, not a normal command

## Cross-domain observations

1. The CLI has one strong noun domain (`engine`) and one strong verb pair (`get`/`set`), but the two patterns are not unified.
2. The most discoverable domain is engine management because both commands share the same noun.
3. The least discoverable domain is configuration because users must know internal keys even though the commands themselves are familiar.
4. `qwen36.server` should be understood as packaging/runtime surface, not as precedent for expanding the user command hierarchy.

## Assessment

The command set clusters into five intuitive domains, but only two of them have more than one command. That small size argues against a major hierarchy refactor. The better path is to tighten the engine and configuration vocabularies while keeping the surface shallow.