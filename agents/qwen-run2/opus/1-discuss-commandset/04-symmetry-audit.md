# Symmetry Audit

## Symmetric Operation Pairs

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|-----------|----------------|-----------------|-------------------|--------------------|----|
| Engine selection | `use-engine cpu` | `use-engine cuda` | Yes (same command, different arg) | Yes | Not a true forward/reverse pair; rather the same command with different targets. Reversal is "use the other engine." |
| Engine observation | `show-engine` | — | N/A | N/A | Read-only, no reverse needed. |
| Config write | `set <key>=<value>` | `unset <key>` (MISSING) | N/A — reverse doesn't exist | N/A | Per DE013 §get/set/unset, `unset` should exist for restoring-to-default. Currently the only way to "undo" a set is to `set` to the previous value manually. |
| Config read | `get <key>` | `set <key>=<value>` | Yes (get/set is standard pair) | Yes (read/write symmetry) | Naming follows DE013 standard. |
| Start chat | `chat` | Ctrl+C / exit | N/A | N/A | Interactive session has no CLI-level "stop" command; user terminates via signal. |
| Generate completions | `completion bash` | — | N/A | N/A | Utility, no reverse needed. |

## Missing Reverse Operations

| Forward Command | Expected Reverse | Impact | Recommendation |
|----------------|-----------------|--------|----------------|
| `set <key>=<value>` | `unset <key>` | Users cannot restore package defaults without knowing the original value | Add `unset` command per DE013 §get/set/unset |
| `use-engine <name>` | — (no formal reverse) | Acceptable — `use-engine` with a different name is the natural reverse | No change needed |

## Naming Asymmetries

| Pair | Asymmetry | Severity |
|------|-----------|----------|
| `use-engine` / `show-engine` | `use` is a mutation verb, `show` is observation — these aren't forward/reverse but they share the `-engine` noun consistently | None (not a pair) |
| `get` / `set` | Perfectly symmetric naming | None |

## Behavioral Asymmetries

| Pair | Asymmetry | Impact |
|------|-----------|--------|
| `use-engine` → `use-engine` (switch back) | The forward switch prompts for confirmation; the reverse switch also prompts. Symmetric. | Low |
| `set` → (manual restore) | Setting a value is instant and silent; "undoing" requires knowing the previous value | Medium — mitigated if `unset` is added |

## Summary

The qwen36 CLI has very few symmetric pairs due to its small command set. The only significant gap is the missing `unset` command, which breaks the DE013-standard `get`/`set`/`unset` triplet. All other operations either have natural symmetry (get/set, use-engine with different args) or are inherently one-directional (show, chat, completion).
