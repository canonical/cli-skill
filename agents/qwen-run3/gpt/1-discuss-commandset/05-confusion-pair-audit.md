# Confusion-Pair Audit

## Scope

This file errs on the side of inclusion. Every public command appears directly at least once.

## Confusion Table

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|
| `use-engine` | `show-engine` | scope ambiguity | high | `use-engine` mutates the active engine; `show-engine` inspects the active or named engine. |
| `list-engines` | `show-engine` | functional overlap | medium | `list-engines` is discovery; `show-engine` is detail for one engine. |
| `use-engine` | `list-engines` | functional overlap | medium | one chooses, the other enumerates. Users often need both in sequence. |
| `set` | `unset` | synonym-adjacent lifecycle | medium | `set` writes a user override; `unset` removes only the user-layer override. |
| `get` | `set` | scope ambiguity | medium | both operate on config keys, but one reads and one writes. |
| `get` | `status` | functional overlap | low | both can reveal effective runtime information, but `get` is key-level and `status` is state-level. |
| `set` | `use-engine` | functional overlap | high | both change effective runtime behavior; `use-engine` also rewrites engine defaults and may overwrite expectations around config provenance. |
| `status` | `show-engine` | scope ambiguity | medium | both answer some form of “what engine is active?” but at different detail levels. |
| `status` | `show-machine` | scope ambiguity | medium | both are inspection commands; one is snap/service state, the other is hardware state. |
| `chat` | `webui` | functional overlap | medium | both are end-user interaction entrypoints to the same backend server, but one is terminal-based and the other is browser-based. |
| `webui` | `status` | scope ambiguity | low | users may expect `status` to tell them whether the web UI is available, but `webui` performs the actual readiness path. |
| `prune-cache` | `use-engine` | functional overlap | medium | `prune-cache` removes inactive-engine components; `use-engine` may later reinstall them. They are related but not named as complements. |
| `version` | `status` | functional overlap | low | both are top-level read-only health or metadata checks. |

Self-check: distinct public commands referenced = 12.

## Findings

1. The highest-risk confusion cluster is the engine domain: `list-engines`, `show-engine`, and `use-engine` are adjacent and often used together.
2. The second highest-risk confusion cluster is `set` versus `use-engine`, because both change behavior at different abstraction levels.
3. `chat` versus `webui` is understandable but still requires clearer onboarding because both are user-entry commands to the same system.

## Low-Risk Outliers

- `version` is low risk because its name is conventional.
- `prune-cache` is unusual, but users generally encounter it only when intentionally doing maintenance.
