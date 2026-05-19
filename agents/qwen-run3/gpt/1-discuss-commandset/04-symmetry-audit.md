# Symmetry Audit

## Scope

This audit covers the 12-command public surface. Because the CLI is small, many findings are about missing reverse operations rather than true forward/reverse pairs.

## Symmetry Table

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|
| Config override | `set <key=value>` | `unset <key>` | yes | yes | Strongest true inverse pair in the CLI. |
| Engine selection | `use-engine <engine>` | — | no | no | No `reset-engine`, `clear-engine`, or explicit reverse operation. |
| Engine enumeration | `list-engines` | — | no | no | Pure observation command; no inverse expected. |
| Engine detail | `show-engine` | — | no | no | Pure observation command; no inverse expected. |
| Host detail | `show-machine` | — | no | no | Pure observation command; no inverse expected. |
| Overall state | `status` | — | no | no | Canonical noun exception. |
| Interactive chat | `chat` | — | no | no | Session ends by EOF or interrupt, not by CLI reverse command. |
| Web UI launch | `webui` | — | no | no | Opens the UI but does not stop or close it. |
| Cache cleanup | `prune-cache` | — | no | no | One-way removal; reinstallation requires a separate `use-engine` path. |
| Version info | `version` | — | no | no | Canonical noun exception. |
| Config read | `get [<key>]` | — | no | no | Complementary to `set`, but not a reverse action. |
| Config list-all behavior | `get` with no key | — | no | no | Uses the same command for both targeted read and whole-map inspection. |

Self-check: all 12 public commands appear directly in the table.

## Missing Reverse Operations

| Forward Command | Missing Reverse | Why it matters |
|---|---|---|
| `use-engine` | `reset-engine` or `clear-engine` | users cannot clear engine state without choosing another engine |
| `prune-cache` | `restore-cache` or `reinstall-engine` | recovery exists only indirectly through `use-engine` |
| `webui` | `close-webui` or service-level control command | service lifecycle remains outside the CLI |
| `chat` | `stop-chat` | acceptable for an interactive tool, but still asymmetric |

## Findings

1. The CLI does not suffer from too many forced inverse pairs; that is good.
2. The absence of an explicit reverse for `use-engine` is the most important asymmetry because engine choice is a high-impact state change.
3. `set`/`unset` is a strong pair and should remain the model for future reversible mutations.
