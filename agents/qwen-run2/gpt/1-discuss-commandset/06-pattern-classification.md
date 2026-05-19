# 06. Pattern Classification And Recommendations

## Pattern classification

Primary pattern: flat verb-led CLI with a small number of verb-noun compounds.

Secondary patterns:

- one nested meta-command: `completion bash`
- one service/runtime outlier: `qwen36.server`

Depth assessment:

- main CLI depth is shallow: `qwen36 <command>`
- one secondary level exists for completion: `qwen36 completion bash`
- the service entrypoint is parallel to the main CLI rather than nested within it

## Command-by-command classification

| Command | Pattern | Depth | Fits current style? | Notes |
|---|---|---|---|---|
| `qwen36 chat` | bare verb | 1 | yes | Friendly user task verb, but outside the administrative vocabulary of the other commands |
| `qwen36 use-engine` | verb-noun compound | 1 | partial | Shared noun is good; `use` is semantically weaker than `set` |
| `qwen36 show-engine` | verb-noun compound | 1 | partial | Shared noun is good; `show-*` is weaker than DE013-preferred forms |
| `qwen36 get` | bare verb | 1 | yes | Strong fit with standard config vocabulary |
| `qwen36 set` | bare verb | 1 | yes | Strong fit with standard config vocabulary |
| `qwen36 completion bash` | meta-command with shell target | 2 | yes | Common shell-integration pattern |
| `qwen36.server` | service entrypoint | parallel surface | no | Snap operational surface, not ideal precedent for end-user grammar |

## Discoverability assessment

| User intent | Predicted command | Actual command | Assessment |
|---|---|---|---|
| start chatting | `qwen36 chat` | `qwen36 chat` | good |
| switch to GPU | `qwen36 set-engine cuda` or `qwen36 use-engine cuda` | `qwen36 use-engine cuda` | good once learned, but the verb is not the most obvious |
| inspect current engine | `qwen36 engine`, `qwen36 engine-status`, or `qwen36 show-engine` | `qwen36 show-engine` | acceptable, but DE013 would prefer a tighter inspect form |
| read one config value | `qwen36 get http.port` | `qwen36 get http.port` | good if the user already knows the key |
| restore a config default | `qwen36 unset http.port` | missing | poor |
| inspect daemon health | `qwen36 status` | external `snap logs` or service tools | poor |
| enable shell completion | `qwen36 completion bash` | `qwen36 completion bash` | acceptable, but undocumented |

## Ecosystem comparison

| Tool | Dominant pattern | Relevant comparison to qwen36 |
|---|---|---|
| `snap` | flat verbs with some noun-oriented inspection commands | qwen36 already inherits the snap mental model for config and service operations, but only partially exposes it |
| `ollama` | flat verbs such as `run`, `serve`, `show`, `list` | qwen36 is similarly small, but it splits client, config, and daemon concepts more sharply |
| `multipass` | flat verbs with precise admin vocabulary (`launch`, `list`, `info`, `shell`) | multipass shows how a small CLI can keep verbs precise without adding hierarchy |

## Key findings

1. The command set is small enough that a full hierarchy redesign would be more expensive than useful.
2. The weakest names are `use-engine` and `show-engine`, not because they are unclear, but because DE013 provides stronger alternatives.
3. The biggest discoverability gap is missing config reversal (`unset`) and missing runtime inspection (`status`).
4. `qwen36.server` should remain treated as packaging/runtime surface, not as the template for future user commands.
5. Documentation is currently doing more work than command grammar to hold the model together.

## Recommendations

| Priority | Recommendation | Rationale | Backward compatibility | Migration cost | Deprecation-compliant plan |
|---|---|---|---|---|---|
| 1 | Add `qwen36 unset <key>` | Per DE013 common commands, `get/set/unset` is the standard config trio. This is the highest-value non-breaking addition. | safe, additive | low | no deprecation needed; add help and docs in the next minor release |
| 2 | Introduce `qwen36 set-engine` as the canonical alias for `qwen36 use-engine` | Per DE013 Grammar + Vocabulary, commands should use precise verbs, and `set-engine` aligns with the existing `set` vocabulary better than `use-engine`. | old command can keep working during transition | medium | next minor: add `set-engine`, keep `use-engine`, emit `warning: "qwen36 use-engine" is deprecated, use "qwen36 set-engine" instead`; next major: remove old name with exit code 2 and message `error: "qwen36 use-engine" was removed in X.0, use "qwen36 set-engine" instead`; next+1 major: remove the special-case message |
| 3 | Introduce `qwen36 engine` as the canonical inspect alias for `qwen36 show-engine` | DE013 prefers `foobar` over `show-foobar` for specific secondary objects. `engine` is shorter and matches the current machine-readable role. | old command can keep working during transition | medium | next minor: add `engine`, keep `show-engine`, warn on old form; next major: remove old form with exit code 2; next+1 major: clean up fallback messaging |
| 4 | Add `qwen36 status` for daemon and engine readiness | DE013 prefers `status` for current tool state. This reduces confusion between `chat`, `show-engine`, and `qwen36.server`. | safe, additive | medium | no deprecation needed; add in a minor release and document as the preferred runtime inspection path |
| 5 | Keep the overall flat hierarchy and avoid introducing an `engine ...` subgroup for now | The command set only has seven commands. A deeper hierarchy would add migration cost without enough scale benefit. | preserves current scripts and mental model | low | no change needed; revisit only if the command count grows materially |

## Tradeoffs

### `set-engine` alias

- Human readability: better than `use-engine` because it signals persisted state mutation.
- Scriptability: neutral if introduced as an alias first.
- Migration risk: moderate but manageable because the surface is tiny.
- Reason not to do more: a full `engine set` subgroup would cost more and add depth with little gain.

### `engine` alias for `show-engine`

- Human readability: shorter and closer to DE013.
- Machine readability: unchanged if output stays YAML-compatible.
- Migration risk: moderate because existing scripts may call `show-engine`.
- Constraint: do not change the YAML schema while renaming the command path.

### `unset` and `status`

- Human readability: materially better, especially for troubleshooting.
- Scriptability: improves automation because reversal and health checks become first-class verbs.
- Migration risk: low because both are additive.

## Final assessment

qwen36 does not need a dramatic command-shape rewrite. The right move is a standards-aligned vocabulary cleanup around engine management, plus two additive commands that close the biggest functional gaps. Those changes are small enough to justify, but any rename should still follow the deprecation spec because even a tiny CLI accumulates scripts quickly.