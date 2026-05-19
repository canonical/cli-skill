# Pattern Classification And Recommendations

## Shape Summary

The qwen36 CLI is a compact flat command set with one hidden debug namespace. Its strongest pattern is a small cluster of standard config verbs (`get`, `set`, `unset`). Its weakest pattern is the interaction and engine surface, where noun-led `webui`, public `use-engine`, and hidden `select-engine` pull the grammar in different directions.

## Pattern Classification

| Property | Assessment |
|---|---|
| Primary grouping pattern | flat top-level command set |
| Secondary grouping pattern | one hidden namespace: `debug` |
| Dominant grammar | verb-led commands with a few noun exceptions |
| Strongest family | `get` / `set` / `unset` |
| Weakest family | `chat` / `webui` plus `use-engine` vs hidden `select-engine` |

## Key Findings

1. The public command set is small and mostly easy to scan.
2. `get`, `set`, and `unset` are aligned with DE013 common command vocabulary.
3. `webui` violates DE013 guidance that commands should be verbs.
4. `use-engine` conflicts semantically with hidden `debug select-engine`, revealing a split vocabulary for the same action.
5. The shipped snap surface and documented/source surface are inconsistent because `chat` and `webui` are feature-gated but not enabled in the product manifest.

## Current Standards Violations To Flag

- Per DE013 §Grammar, commands should be verbs. `webui` is noun-led.
- Per DE013 §Grammar, command vocabulary should be even and predictable. `use-engine` versus hidden `select-engine` is uneven vocabulary for the same semantic action.
- Per DE013 §Grammar and §Commonly used commands, observation commands may prefer noun shorthand or explicit status forms over `show-*` when the shorthand is viable. `show-engine` and `show-machine` are acceptable, but they are second-best choices rather than ideal ones.
- Per DE013 §Grammar, plural-noun list shorthands are preferred for secondary objects. `list-engines` is valid but less aligned than `engines` would be.

## Recommendations

### 1. Fix the product-surface mismatch before renaming anything

Recommendation: either enable `chat` and `webui` in the shipped qwen36 snap or remove them from public docs until they are enabled.

Rationale: this is the highest-impact usability issue because it breaks discoverability and user trust before grammar refinements even matter.

Backward compatibility impact: none if the commands are enabled; documentation-only if the README is corrected.

Migration cost: low.

### 2. Add `engines` as a public alias for `list-engines`

Recommendation: add `engines` as a preferred alias, while keeping `list-engines` for compatibility.

Rationale: per DE013 §Grammar, plural-noun shorthand is preferred for listing secondary objects. `engines` is shorter and more discoverable.

Backward compatibility impact: none, because this is additive.

Migration cost: very low.

### 3. Introduce `select-engine` as the preferred public name for engine choice

Recommendation: add `select-engine` and begin deprecating `use-engine`.

Rationale: per DE013 §Grammar, command verbs should be clear and even. `select` is the sharper verb and is already used in the hidden debug surface, so the project has effectively validated it already.

Required deprecation plan per the deprecation specification:

- next minor release: add `select-engine`, keep `use-engine` working
- deprecation warning text on old command: `warning: "qwen36 use-engine" is deprecated, use "qwen36 select-engine" instead`
- keep both for at least one full release cycle
- next major release: remove `use-engine`, return exit code 2, and emit `error: "qwen36 use-engine" was removed in 4.0, use "qwen36 select-engine" instead`

Backward compatibility impact: medium without the alias, low with the alias-and-warning path.

Migration cost: low to medium.

### 4. Rename `webui` to a verb-led launcher such as `open-webui`

Recommendation: introduce `open-webui` or `launch-webui` and deprecate `webui`.

Rationale: per DE013 §Grammar, commands should be verbs. `webui` is the clearest standards violation in the public surface.

Required deprecation plan per the deprecation specification:

- next minor release: add `open-webui`, keep `webui`
- deprecation warning text: `warning: "qwen36 webui" is deprecated, use "qwen36 open-webui" instead`
- keep both for at least one cycle
- next major release: remove `webui`, return exit code 2, and emit `error: "qwen36 webui" was removed in 4.0, use "qwen36 open-webui" instead`

Backward compatibility impact: medium without aliasing, low with the required transition path.

Migration cost: low.

### 5. Add an explicit reverse operation for engine choice

Recommendation: add `reset-engine` or `clear-engine` as a public command.

Rationale: the symmetry audit shows that engine selection is the most important high-impact state change without a reverse command.

Standards alignment: this is compatible with DE013 §Grammar because it adds a clear verb-led command rather than renaming a stable existing one.

Backward compatibility impact: none.

Migration cost: medium because behavior must be defined carefully.

## Tradeoffs

| Recommendation | Backward Compat | Scriptability | Human Readability | Migration Cost |
|---|---|---|---|---|
| fix feature exposure | excellent | improves because docs match reality | major improvement | low |
| add `engines` alias | excellent | neutral to positive | positive | very low |
| add `select-engine` and deprecate `use-engine` | good if alias kept | positive after transition | positive | low-medium |
| add `open-webui` and deprecate `webui` | good if alias kept | neutral | positive | low |
| add `reset-engine` | excellent | positive | positive | medium |

## Recommended Order

1. align shipped features with docs
2. add `engines`
3. add `select-engine` with deprecation warning on `use-engine`
4. add `open-webui` with deprecation warning on `webui`
5. design and add `reset-engine`
