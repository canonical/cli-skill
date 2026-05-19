# 02 Verb Taxonomy

## Scope

This section classifies the six verbs implied by the six leaf commands. Source verb count: 6. Output verb count: 6.

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|
| `chat` | execution | atelic | no | none | `qwen36 chat` |
| `generate` | execution | telic | no | none | `qwen36 completion bash` |
| `get` | observation | punctual | partial | `set` | `qwen36 get http.port` |
| `set` | mutation | punctual | partial | `get` | `qwen36 set http.port=8326` |
| `show` | observation | punctual | no | none | `qwen36 show-engine` |
| `use` | mutation | punctual | partial | none | `qwen36 use-engine cpu`, `qwen36 use-engine --auto` |

## Notes By Verb

| Verb | Assessment |
|---|---|
| `chat` | Clear for interactive use, but it does not tell the user whether the command starts a local client, a local server session, or a remote session. |
| `generate` | Semantic, not literal. The literal surface is `completion`, which is noun-led. |
| `get` | Strong fit for config observation. Aligned with DE013 common commands. |
| `set` | Strong fit for config mutation. Aligned with DE013 common commands. |
| `show` | Understandable, but DE013 explicitly prefers noun shorthand or `*-status` over `show-*` for observation commands. |
| `use` | Understandable, but slightly vague. It selects an engine rather than merely "using" one transiently. |

## Taxonomy Findings

1. The command set leans heavily on observation and mutation, which matches the CLI's role as a control plane.
2. `get` and `set` are the best-aligned verbs in the set.
3. `show` and noun-led `completion` are the main standards mismatches.
4. `use` is workable, but if the command family grows, `select-engine` would be semantically sharper than `use-engine`. That said, the improvement is not large enough to justify breaking users without a broader redesign.

## Recommendation Compliance Notes

Per DE013, any future cleanup should prefer canonical verbs and avoid noun-led command groups. Per the deprecation specification, if `show-engine` or `use-engine` were ever renamed, the next minor release should keep the old form working with a warning, and the next major release should remove it with exit code 2 and a replacement hint.