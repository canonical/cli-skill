# 05 — Confusion-Pair Audit

## All Confusion Pairs

| # | Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|---|
| 1 | `show-engine` | `get server` | scope ambiguity | **medium** | `show-engine` outputs the full engine YAML (name, hardware requirements, components, configurations); `get server` returns only the server component name string. Users may not know which to use when debugging engine state. |
| 2 | `use-engine cpu` | `set server=llamacpp` | functional overlap | **medium** | `use-engine` is the high-level operation that sets multiple config keys atomically (server, model, mmproj, base-path, gpu-layers); `set server=llamacpp` changes only one key and could leave configuration in an inconsistent state. The relationship between these is undocumented. |
| 3 | `use-engine --auto` | `use-engine cpu` | scope ambiguity | **low** | Both are the same command with different invocation styles (automatic vs manual selection). Low risk because `--auto` clearly signals automated behavior, but users may wonder if `--auto` does additional validation beyond selection. |
| 4 | `get model-name` | `get model` | naming collision | **high** | Two config keys with near-identical names but different semantics: `model-name` is the human-readable model identifier for API responses; `model` is the snap component directory name. Users will frequently confuse these. |
| 5 | `set http.port=8326` | `set --package http.port=8326` | scope ambiguity | **medium** | Same command with and without `--package` writes to different precedence levels. Users unaware of the flag may inadvertently create user-level overrides that mask package defaults, or fail to understand why their `set` doesn't persist across updates. |
| 6 | `completion bash` | `completion` (bare) | scope ambiguity | **low** | If `completion` is invoked without the `bash` argument, behavior is unclear. Users may not know what shells are supported or what happens with an invalid/missing shell argument. |

## Self-Check

- Total commands: 6 (`chat`, `use-engine`, `show-engine`, `get`, `set`, `completion`)
- Commands referenced in pairs: `show-engine`, `get`, `use-engine`, `set`, `completion` → 5 of 6 commands appear
- `chat` does not appear in any confusion pair (it's unique in function and naming)
- ✓ Complete — all potential confusions identified

## Risk Summary

| Risk Level | Count | Pairs |
|---|---|---|
| High | 1 | `get model-name` vs `get model` |
| Medium | 3 | `show-engine` vs `get server`; `use-engine` vs `set server=...`; `set` vs `set --package` |
| Low | 2 | `use-engine --auto` vs `use-engine cpu`; `completion bash` vs `completion` |

## Recommendations

1. **Rename `model-name` to `model-label` or `api-model-name`** to clearly distinguish from the `model` component key. This is the highest-risk confusion pair.
2. **Document the relationship between `use-engine` and individual `set` commands** — make clear that `use-engine` is the safe, atomic way to switch engines while direct `set` on engine-related keys is an advanced/dangerous operation.
3. **Add a warning when users `set` engine-related keys directly** (server, model, multimodel-projector) outside of `use-engine`, noting that partial changes may leave configuration inconsistent.
