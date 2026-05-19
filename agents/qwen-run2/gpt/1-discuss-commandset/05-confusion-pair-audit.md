# 05. Confusion-Pair Audit

## Scope note

The qwen36 surface is small, so true confusion risk is moderate rather than severe. The table below errs on the side of inclusion and keeps every command represented in at least one pair.

## Confusion pairs

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|
| `qwen36 chat` | `qwen36.server` | functional overlap | high | `chat` is the user client workflow; `qwen36.server` is the backing daemon and should usually be managed indirectly or via snap service commands |
| `qwen36 chat` | `qwen36 show-engine` | scope ambiguity | low | `chat` uses the current engine but does not inspect it; `show-engine` only reports engine metadata |
| `qwen36 chat` | `qwen36 get` | functional overlap | low | `chat` consumes current config indirectly, while `get` exposes one config value directly |
| `qwen36 chat` | `qwen36 completion bash` | naming collision | low | Both affect the terminal, but `chat` is runtime interaction and `completion bash` is shell setup support |
| `qwen36 use-engine` | `qwen36 show-engine` | scope ambiguity | medium | One changes the selected engine and the other reports it; shared noun helps, but users still need to learn the split |
| `qwen36 use-engine` | `qwen36 set` | functional overlap | medium | Both can change runtime behavior, but `use-engine` is a high-level engine macro while `set` changes individual config keys |
| `qwen36 use-engine` | `qwen36.server` | functional overlap | medium | `use-engine` influences which daemon runtime will launch, but it does not start the daemon itself |
| `qwen36 show-engine` | `qwen36 get` | scope ambiguity | medium | Both inspect state, but `show-engine` returns structured engine metadata while `get` returns one scalar config value |
| `qwen36 show-engine` | `qwen36.server` | functional overlap | low | `show-engine` describes selected runtime intent; `qwen36.server` executes the runtime |
| `qwen36 get` | `qwen36 set` | synonym verbs | medium | The pair is familiar, but users still need to know hidden key names and the missing `unset` counterpart |
| `qwen36 get` | `qwen36 completion bash` | scope ambiguity | low | Both emit machine-consumable stdout, but one is config inspection and the other is shell candidate generation |
| `qwen36 set` | `qwen36.server` | functional overlap | medium | `set` can break or change daemon startup behavior without being the service control command itself |
| `qwen36 completion bash` | `qwen36.server` | naming collision | low | Both are operational/terminal-adjacent commands, but they serve unrelated purposes |

## Risk summary

| Risk | Count | Pairs |
|---|---|---|
| high | 1 | `chat` / `qwen36.server` |
| medium | 6 | `use-engine` / `show-engine`, `use-engine` / `set`, `use-engine` / `qwen36.server`, `show-engine` / `get`, `get` / `set`, `set` / `qwen36.server` |
| low | 6 | all remaining listed pairs |

## Main confusion themes

1. Client versus daemon: `chat` and `qwen36.server` are the most likely pair to confuse new users.
2. Macro versus primitive config mutation: `use-engine` and `set` both change runtime behavior at different levels.
3. Structured inspect versus scalar inspect: `show-engine` and `get` are both read commands with different abstraction levels.

## Assessment

The command set is small enough that users can learn it quickly, but the biggest confusion source is conceptual rather than lexical: the tool bundles a client command, a configuration API, and a daemon service surface into one snap. Documentation and a `status` command would likely reduce confusion more than a wholesale rename campaign.