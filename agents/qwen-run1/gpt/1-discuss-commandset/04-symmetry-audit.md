# 04 Symmetry Audit

## Scope

This audit includes all six leaf commands. Because the command set is small, most symmetry findings are about missing reverse operations rather than mismatched existing pairs.

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|
| Engine selection | `qwen36 use-engine <engine>` | none | No | No | There is no explicit `unset-engine`, `reset-engine`, or `auto-engine` inverse. `--auto` is a mode on the same command rather than a reverse operation. |
| Engine observation vs engine mutation | `qwen36 show-engine` | `qwen36 use-engine <engine>` | No | No | These commands are complementary, not inverse. One inspects state; the other changes it. |
| Config mutation vs config observation | `qwen36 set <key>=<value>` | `qwen36 get <key>` | Partially | Partially | Common and useful pair, but `get` is not a true inverse because it does not restore prior state. |
| Config restoration | `qwen36 set <key>=<value>` | none | No | No | No `unset` command is documented, so there is no direct route back to defaults. |
| Interactive session entry | `qwen36 chat` | none | No | No | Session termination is shell-driven rather than command-driven. |
| Shell completion generation | `qwen36 completion bash` | none | No | No | Generation is one-way. No install/remove/manage complement is exposed. |

Self-check: every leaf command appears at least once in the table.

## Findings

1. The only meaningful near-pair is `get` and `set`.
2. The engine domain is asymmetrical: users can set and inspect the current engine but cannot list supported engines or restore defaults explicitly.
3. The CLI is not overburdened with forced symmetry, which is good, but it does need one missing complement: `unset` for config.

## Recommendation Compliance Notes

Per DE013, `get`/`set` already match Canonical's standard command vocabulary; adding `unset` would improve symmetry without renaming anything. That is the safest additive improvement because it can land in a minor release without breaking scripts. Any attempt to rename `show-engine` or `use-engine` should follow the full one-cycle deprecation process described in the deprecation specification.