# 03 Semantic Domain Clustering

## Scope

Every leaf command appears in exactly one domain. Source command count: 6. Sum of domain counts below: 6.

| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---:|---|---|---|
| conversation | 1 | `qwen36 chat` | Yes | Single interactive command. No CRUD expectation. |
| engine | 2 | `qwen36 use-engine`, `qwen36 show-engine` | Mostly | Noun is consistent, but observation uses `show-*`, which DE013 treats as second-best to noun shorthand or `*-status`. CRUD coverage is partial: choose and inspect exist; list/reset do not. |
| configuration | 2 | `qwen36 get`, `qwen36 set` | Yes | Strongest domain in the CLI. Standard read/write pair, but no `unset`. |
| shell integration | 1 | `qwen36 completion bash` | No | Domain is valid, but the command shape is noun-led and only one shell target is evidenced. |

## Domain Analysis

### conversation

- Noun form: implicit rather than explicit
- CRUD coverage: not applicable
- Verb consistency: one command only

### engine

- Noun form: explicit and stable as `engine`
- CRUD coverage: incomplete
- Verb consistency: acceptable but not ideal because `show-engine` conflicts with DE013 observation guidance

### configuration

- Noun form: implicit config domain derived from keys
- CRUD coverage: mostly useful, but no `unset`
- Verb consistency: high

### shell integration

- Noun form: explicit in semantics, but not in user-facing grammar
- CRUD coverage: not applicable
- Verb consistency: low because the top-level literal is a noun

## Findings

1. The command set clusters cleanly into four domains, which is appropriate for a small CLI.
2. Configuration and engine management are the core domains.
3. The shell-integration domain is underdeveloped and underdocumented.
4. No command currently helps users discover supported engine names from within the CLI itself.

## Recommendation Compliance Notes

Per DE013, the engine domain is the most obvious place to improve naming consistency, with `show-engine` as the candidate for eventual normalization. Per the deprecation spec, any rename should be additive first, with the old command retained for at least one cycle and removed only in a major release.