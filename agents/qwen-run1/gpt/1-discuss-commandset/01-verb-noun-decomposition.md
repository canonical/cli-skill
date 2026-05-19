# 01 Verb-Noun Decomposition

## Scope

Source command count: 6 leaf commands.

- `qwen36 chat`
- `qwen36 use-engine`
- `qwen36 show-engine`
- `qwen36 get`
- `qwen36 set`
- `qwen36 completion bash`

## Decomposition Matrix

Semantic decomposition is used where the literal surface is not already verb-noun. This keeps every command in one matrix while still flagging grammar mismatches.

| Verb | engine | config | completion | conversation |
|---|---:|---:|---:|---:|
| chat | - | - | - | ✓ |
| generate | - | - | ✓ | - |
| get | - | ✓ | - | - |
| set | - | ✓ | - | - |
| show | ✓ | - | - | - |
| use | ✓ | - | - | - |

## Command-Level Mapping

| Command | Literal Surface | Semantic Verb | Semantic Noun | Clean Verb-Noun? | Notes |
|---|---|---|---|---|---|
| `qwen36 chat` | verb only | `chat` | `conversation` | Partial | Usable, but noun is implicit rather than explicit. |
| `qwen36 use-engine` | verb-noun | `use` | `engine` | Yes | Clear enough, though `use` is broader than `select`. |
| `qwen36 show-engine` | verb-noun | `show` | `engine` | Yes, but non-ideal | Per DE013, showing state usually prefers `engine` or `engine-status` over `show-engine`. |
| `qwen36 get` | verb only | `get` | `config` | Partial | Common config grammar, noun implied by argument key space. |
| `qwen36 set` | verb only | `set` | `config` | Partial | Common config grammar, noun implied by assignment. |
| `qwen36 completion bash` | noun + target | `generate` | `completion` | No | Literal command is noun-led and conflicts with DE013 verb-first guidance. |

Self-check: output command count = 6.

## Incomplete CRUD Sets

| Noun | Present Verbs | Expected Missing Verbs | Assessment |
|---|---|---|---|
| `engine` | `use`, `show` | `list`, `status` or `engine`, optional `unset` or `reset` | Discovery is weak. Users can select or inspect the current engine but cannot list supported engines from the documented surface. |
| `config` | `get`, `set` | `unset` | Config grammar is close to standard, but restoring defaults is not surfaced. |
| `completion` | `generate` | none required | Narrow feature. Only bash is evidenced. |
| `conversation` | `chat` | none required | Single-purpose interaction command. |

## Verb Inconsistencies

| Resource | Current Surface | Consistency Issue |
|---|---|---|
| `engine` | `use-engine`, `show-engine` | Mutation and observation verbs are fine, but DE013 prefers noun shorthand or `*-status` over `show-*` for observation. |
| `completion` | `completion bash` | Uses a noun-led grouping token instead of a verb. |
| `config` | `get`, `set` | Internally consistent and aligned with DE013 common commands. |

## Orphan Or Exception Commands

| Command | Why It Is Exceptional |
|---|---|
| `qwen36 chat` | Verb-only, implicit noun. Acceptable, but distinct from the rest of the config-and-engine grammar. |
| `qwen36 completion bash` | Only command that is noun-led at the top level. |

## Recommendation Compliance Notes

Per DE013 grammar, the strongest naming issue is `show-engine`, followed by `completion bash`.

Recommended change only if the team is willing to pay migration cost:

1. Add a preferred observational alias such as `engine` or `engine-status` in the next minor release while keeping `show-engine` working.
2. Emit a stderr warning: `warning: "qwen36 show-engine" is deprecated, use "qwen36 engine" instead`.
3. Keep that aliasing for at least one release cycle.
4. In the next major release, remove `show-engine` and fail with exit code 2 plus: `error: "qwen36 show-engine" was removed in 4.0, use "qwen36 engine" instead`.

For `completion bash`, the standard conflict is real, but the benefit of renaming is smaller. It should only be changed as part of a broader command grammar cleanup.