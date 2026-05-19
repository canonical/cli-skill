# 04. Symmetry Audit

## Scope note

This audit includes explicit pairs where they exist and missing reverse operations where they do not. Every command in the qwen36 surface appears at least once below.

## Symmetry table

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|
| inspect engine state | `qwen36 show-engine` | missing | no | no | DE013 would prefer `engine` or `engine-status`; there is no paired engine-state mutation verb in the same naming family |
| select engine | `qwen36 use-engine <engine>` | missing dedicated inverse | no | partial | Reversing is possible only by selecting a different engine or rerunning `--auto` |
| read config | `qwen36 get <key>` | `qwen36 set <key>=<value>` | partial | partial | These are paired inspect/mutate operations, not true inverses |
| write config | `qwen36 set <key>=<value>` | missing `qwen36 unset <key>` | no | no | DE013 explicitly names `unset` as the expected partner for `set` |
| generate bash completion | `qwen36 completion bash` | missing | no | no | Completion generation has no command-level removal or disable counterpart |
| start interactive conversation | `qwen36 chat` | missing | no | partial | The user exits the session externally; there is no named `end-chat` or `exit-chat` command |
| run daemon service | `qwen36.server` | external `snap stop qwen36.server` | no | partial | Service lifecycle symmetry exists only in snap tooling, not in qwen36 command grammar |

## Strongest symmetry present today

| Pair | Strength | Notes |
|---|---|---|
| `get` / `set` | moderate | Familiar read/write pair, but still incomplete without `unset` |
| `use-engine` / `show-engine` | weak-to-moderate | Shared noun helps, but the verbs are not an ideal pair |

## Missing or asymmetric reversals

| Current command | Expected reverse or partner | Why it matters |
|---|---|---|
| `qwen36 set` | `qwen36 unset` | Restoring defaults is currently manual and undocumented |
| `qwen36 use-engine` | `qwen36 set-engine <previous>` or reset alias | Engine selection is overwritable, but not explicitly reversible |
| `qwen36 show-engine` | `qwen36 engine` or `qwen36 engine-status` as canonical inspect form | Improves naming symmetry with other engine verbs |
| `qwen36 completion bash` | documented install/remove guidance | Shell UX lacks a full lifecycle story |
| `qwen36.server` | user-level `status` or `stop` wrapper, if desired | Today the lifecycle crosses tool boundaries into `snap` |

## Assessment

The qwen36 command set is intentionally small, so perfect symmetry is not required everywhere. The important asymmetry is config management: `set` has no `unset`, and engine management uses understandable but weaker verb choices than the Canonical standard recommends. The service surface is also structurally asymmetric, but that is mostly a consequence of snap packaging rather than poor command design.