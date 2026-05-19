# 06 Pattern Classification

## Scope

This classification covers all six leaf commands.

| Command | Surface Pattern | Argument Pattern | Standards Fit | Scriptability | Recommendation |
|---|---|---|---|---|---|
| `qwen36 chat` | flat verb | no args | Good | Low | Keep. It is memorable and task-oriented. |
| `qwen36 use-engine` | flat verb-noun | positional noun value plus optional mode flags | Acceptable | Medium | Keep for now. If a broader redesign ever happens, `select-engine` is a possible sharper verb, but not enough benefit to justify churn alone. |
| `qwen36 show-engine` | flat verb-noun | no args | Weak | High | Best rename candidate because DE013 prefers noun shorthand or `*-status` over `show-*`. Do not break immediately; add alias first. |
| `qwen36 get` | flat verb | single positional key | Strong | High | Keep. Standard config read command. |
| `qwen36 set` | flat verb | single positional `key=value` plus `--package` | Strong | High | Keep. Add documentation for `--package` and value validation. |
| `qwen36 completion bash` | two-level noun namespace | second-level shell target | Weak | Medium | Keep short term for compatibility; document it. Only revisit if the team decides to standardize every top-level command around verb-first grammar. |

Self-check: output command count = 6.

## Pattern Families

| Family | Commands | Assessment |
|---|---|---|
| Flat verbs | `qwen36 chat`, `qwen36 get`, `qwen36 set` | Strongest family. Short, predictable, and aligned with DE013 common verbs. |
| Flat verb-noun | `qwen36 use-engine`, `qwen36 show-engine` | Coherent domain family, but `show-engine` is the weakest member because of the standard's observation guidance. |
| Two-level namespace | `qwen36 completion bash` | Functional but stylistically inconsistent with the rest of the command set. |

## Overall Classification

The CLI is mostly a flat command set with one small exception namespace. That is a good size and shape for a six-command tool. The main issue is not excessive hierarchy; it is uneven grammar quality across commands.

## Recommendation Compliance Notes

Per DE013, retain `get` and `set` as the anchor vocabulary and avoid introducing more noun-led top-level commands. Per the deprecation specification, if `show-engine` is ever normalized to `engine` or `engine-status`, the migration should be:

1. Minor release: add the replacement command and keep `show-engine` working.
2. Minor release stderr warning: `warning: "qwen36 show-engine" is deprecated, use "qwen36 engine" instead`.
3. Maintain both for at least one full cycle.
4. Major release: remove `show-engine`, return exit code 2, and print `error: "qwen36 show-engine" was removed in 4.0, use "qwen36 engine" instead`.