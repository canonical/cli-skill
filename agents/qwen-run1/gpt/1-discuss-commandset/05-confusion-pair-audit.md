# 05 Confusion Pair Audit

## Scope

This audit lists the most plausible command confusions in the six-command surface. Every leaf command appears in the table at least once.

| Pair | Why A User Could Confuse Them | Actual Difference | Risk Level | Recommendation |
|---|---|---|---|---|
| `qwen36 use-engine` vs `qwen36 show-engine` | Same noun, adjacent verbs, both about engine state. | `use-engine` mutates state; `show-engine` reports current state as YAML. | High | Improve docs and help with explicit "select" versus "inspect" language. If a rename is ever pursued, `show-engine` is the better candidate for normalization per DE013. |
| `qwen36 get` vs `qwen36 set` | Minimal pair differing by one verb. | `get` reads a single config value; `set` writes a single config value. | Medium | Keep as-is. This is standard and productive confusion, not harmful confusion. |
| `qwen36 get` vs `qwen36 show-engine` | Both can reveal engine-related information. | `get` returns one flattened config value; `show-engine` returns structured YAML describing the selected engine. | Medium | Document when to prefer key lookup versus engine inspection. |
| `qwen36 set` vs `qwen36 use-engine` | Both can change effective server behavior. | `set` mutates one explicit key; `use-engine` changes a bundle of engine-related values and likely component selection. | High | Document that `use-engine` is the supported high-level operation and `set` is lower-level configuration. |
| `qwen36 chat` vs `qwen36 use-engine` | New users may think engine choice happens inside chat startup. | `chat` starts the interactive client; `use-engine` prepares or changes the runtime backend. | Medium | README should place engine selection before chat and show verification via `show-engine`. |
| `qwen36 completion bash` vs `qwen36 chat` | Both appear terminal-facing and interactive-adjacent. | `completion bash` emits completion candidates for shell integration; `chat` opens an interaction loop. | Low | Add a README shell-completion section so the command's purpose is obvious. |

Self-check: output command count = 6 unique leaf commands represented.

## Findings

1. The only high-risk confusions are in the engine and config domains.
2. `show-engine` is understandable, but its `show-*` grammar makes it look like a generic display command rather than a domain-specific inspection command.
3. `set` versus `use-engine` needs explicit documentation because both change runtime behavior at different abstraction levels.

## Recommendation Compliance Notes

Per DE013, command names should make the user's mental model predictable. The safest immediate improvement is documentation, not renaming. If the team later decides to normalize `show-engine`, follow the deprecation spec exactly: add the replacement in a minor release, keep the old form working with a warning for at least one cycle, then remove it in a major release with exit code 2 and a migration hint.