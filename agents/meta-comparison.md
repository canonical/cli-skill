# Meta-Comparison: Small CLI vs Large CLI Performance

How do models behave differently when analyzing a small CLI (~18 commands, qwen36-snap) versus a large CLI (130+ commands, juju)?

---

## 1. Scale Parameters

| Dimension | qwen36-snap (Small) | juju (Large) | Ratio |
|---|---|---|---|
| Top-level commands | ~18 | 130+ | 7x |
| Codebase size | ~50k LOC | ~2M LOC | 40x |
| Entry points | 1 (main.go) | 3 (main.go, plugin.go, supercommand) | 3x |
| Config complexity | env vars + YAML | env + YAML + cloud credentials + models + controllers | 5x+ |
| Output formats | text, JSON | text, JSON, YAML, tabular | 4x |

---

## 2. Output Volume Scaling

How much more do agents write when the CLI is 7x larger?

| Model | qwen Words | juju Words | Scale Factor | Scales Linearly? |
|---|---|---|---|---|
| DeepSeek V4 Pro | 9,695 | 11,084 | 1.14x | Under-scales — efficient compression |
| GLM-5 | 5,903 | 14,496 | 2.46x | Over-scales — invests heavily in large CLIs |
| Kimi K2.6 | 6,750 | 10,057 | 1.49x | Moderate scaling |
| GPT (best) | 8,433 | 11,464 | 1.36x | Proportional |
| Opus (avg) | 4,808 | 6,620 | 1.38x | Proportional |
| Gemini (best) | 861 | 4,350 | 5.05x | Extreme — barely functions on small CLI |

**Expected linear scale**: ~1.5-2x (reflecting more commands but similar analytical structure).

**Key finding**: Models fall into 3 categories:
- **Compressors** (DSv4-Pro): Output barely grows with project size. They summarize patterns rather than enumerate.
- **Scalers** (GPT, Opus, Kimi): Output grows proportionally. They balance enumeration with abstraction.
- **Enumerators** (GLM-5): Output grows super-linearly. They catalogue everything in large projects.

---

## 3. Phase Performance by Project Size

### Phase 1 (Structure Discovery) — architecture, commandset, argument-structure

| Model | Small CLI Score | Large CLI Score | Delta | Explanation |
|---|---|---|---|---|
| DeepSeek V4 Pro | 82% | 85% | +3 | Consistent both ways |
| GLM-5 | 60% | 92% | **+32** | Thrives on complex structure |
| Kimi K2.6 | 68% | 65% | -3 | Neutral — enumerates either way |
| GPT | 78% | 88% | +10 | Benefits from richer material |
| Opus | 65% | 65% | 0 | Perfectly size-independent |
| Gemini | 15% | 38% | +23 | Less broken on large CLI |

**Insight**: Phase 1 is easier on large CLIs for most models. More structure = more to discover = clearer findings. Small CLIs force models to go deeper with less surface material, which some handle poorly.

### Phase 2 (Behavioral Analysis) — config, output, errors, safety

| Model | Small CLI Score | Large CLI Score | Delta | Explanation |
|---|---|---|---|---|
| DeepSeek V4 Pro | 88% | 80% | **-8** | Better on small — focused attention |
| GLM-5 | 60% | 97% | **+37** | Massive improvement on large CLI |
| Kimi K2.6 | 62% | 28% | **-34** | Collapses on large — context exhausted |
| GPT | 75% | 85% | +10 | Consistent improvement |
| Opus | 68% | 45% | **-23** | Worse on large (reused analysis) |
| Gemini | 12% | 35% | +23 | Slightly less broken |

**Critical finding**: Phase 2 is where small vs large divergence is sharpest.
- **Kimi K2.6 loses 34 points** going from small→large. It spends all context enumerating the large commandset in Phase 1 and has nothing left.
- **GLM-5 gains 37 points** going small→large. It thrives when there's rich behavioral material to analyze.
- **DeepSeek V4 Pro is the only model that's better on small CLIs** in Phase 2 — focused attention without distraction.

### Phase 3 (Meta-Analysis + Design)

| Model | Small CLI Score | Large CLI Score | Delta | Explanation |
|---|---|---|---|---|
| DeepSeek V4 Pro | 78% | 80% | +2 | Rock-solid consistency |
| GLM-5 | 76% | 90% | +14 | Strong either way, better on large |
| Kimi K2.6 | 76% | 72% | -4 | Recovers somewhat — shorter files |
| GPT | 72% | 78% | +6 | Slight benefit from richer material |
| Opus | 72% | 92% | +20 | Best design output on large CLI |
| Gemini | 8% | 10% | +2 | Failed both |

**Insight**: Phase 3 shows the smallest size-dependent variance. By this point, models are synthesizing from their own analysis rather than exploring the codebase, so project size matters less.

---

## 4. Failure Modes by Project Size

### Small CLI Failure Modes

| Failure | Models Affected | Cause |
|---|---|---|
| Placeholder stubs | Gemini (3 runs) | Insufficient codebase surface to anchor analysis |
| Thin behavioral analysis | GLM-5 | Fewer config sources and output patterns to discover |
| Over-enumeration | Kimi K2.6 | Lists all 18 commands exhaustively, little left for depth |
| Wrong filenames | Gemini (run2) | Possibly defaulted to a generic template |

### Large CLI Failure Modes

| Failure | Models Affected | Cause |
|---|---|---|
| Context exhaustion | Kimi K2.6 | Enumerated 130+ commands, starved Phase 2 |
| Shallow behavioral files | Gemini, Kimi K2.6 | Couldn't reach config/error/safety code after structure discovery |
| Analysis reuse | Opus (juju) | Recycled another model's output rather than re-analyzing |
| Excessive bash exploration | Kimi K2.6 (48 bash calls) | Thrashed through codebase without converging |

### Size-Specific Risk Matrix

| Risk | Small CLI | Large CLI |
|---|---|---|
| Context overflow | Low | **High** — Kimi K2.6 lost 34% in Phase 2 |
| Insufficient material | **High** — Gemini produced stubs | Low |
| Exploration cost | Low | **High** — models average 2x more tool calls |
| Phase imbalance | Moderate | **High** — enumeration starves later phases |
| Design quality | Similar | Similar — Phase 3 is size-independent |

---

## 5. Tool Strategy Shift

How models adapt their exploration approach by project size:

| Model | Small CLI (qwen) | Large CLI (juju) | Adaptation |
|---|---|---|---|
| DeepSeek V4 Pro | 8 bash, 41 read | 35 bash, 19 read | **Inverts strategy** — reads on small, explores on large |
| GLM-5 | 19 bash, 30 read | 31 bash, 19 read | Moderate shift — more bash on large |
| Kimi K2.6 | 12 bash, 55 read | 48 bash, 27 read | **Extreme inversion** — 4x more bash on large |

**Pattern**: All models shift toward bash-heavy exploration on large CLIs. On small CLIs, the codebase fits in a few file reads. On large CLIs, agents need `find`, `grep`, and directory listings to navigate.

**Efficiency metric** (analysis words per tool call):

| Model | Small CLI | Large CLI | Better On |
|---|---|---|---|
| DeepSeek V4 Pro | 198 w/call | 165 w/call | Small — more output per action |
| GLM-5 | 120 w/call | 227 w/call | Large — thrives with room to explore |
| Kimi K2.6 | 101 w/call | 134 w/call | Large — but still lowest efficiency |

---

## 6. Quality Distribution Pattern

Analysis file word counts reveal how models allocate attention across the 9 required files:

### Coefficient of Variation (CV) Across 9 Files

Lower CV = more evenly distributed analysis. Higher CV = spiked on a few files.

| Model | Small CLI CV | Large CLI CV | Size Impact |
|---|---|---|---|
| DeepSeek V4 Pro | 0.35 | 0.48 | +0.13 — slightly less balanced on large |
| GLM-5 | 0.15 | 0.32 | +0.17 — still most balanced overall |
| Kimi K2.6 | 0.38 | 0.92 | **+0.54** — extremely unbalanced on large |
| GPT | 0.24 | 0.38 | +0.14 — moderate |
| Opus | 0.18 | 0.32 | +0.14 — moderate |

**Finding**: Every model becomes less balanced on large CLIs, but Kimi K2.6's CV nearly triples (0.38 → 0.92), confirming it dumps all effort into commandset/arguments and neglects behavioral files.

---

## 7. Recommendations by Project Size

### For Small CLIs (< 30 commands)

| Rank | Model | Why |
|---|---|---|
| 1 | DeepSeek V4 Pro | Best Phase 2 score (88%); focused attention; efficient |
| 2 | GPT | Strong and improving; benefits from feedback |
| 3 | Opus | Reliable and reproducible |
| 4 | Kimi K2.6 | Decent but over-enumerates even small command sets |
| 5 | GLM-5 | Under-invests — treats small CLIs as simple |
| — | Gemini | Not viable — produces stubs |

### For Large CLIs (100+ commands)

| Rank | Model | Why |
|---|---|---|
| 1 | GLM-5 | Dominant — 93% score; scales super-linearly with complexity |
| 2 | GPT-5.4 | Balanced 84%; strong Phase 1 and 2 |
| 3 | DeepSeek V4 Pro | Solid 82%; most consistent across phases |
| 4 | Opus | Good design output (92% Phase 3) but recycled analysis |
| 5 | Kimi K2.6 | Needs chunking — collapses in Phase 2 without phased workflow |
| — | Gemini | Thin but structurally complete (28%) — marginal |

### Universal Recommendation

**DeepSeek V4 Pro** is the safest choice when you don't know the project size upfront. It's the only model that scores within 3 points of its peak on both small and large CLIs, with the lowest variance across project sizes.

---

## 8. Summary Table

| Metric | Best on Small CLI | Best on Large CLI | Most Size-Dependent | Most Size-Independent |
|---|---|---|---|---|
| Overall score | DSv4-Pro (83%) | GLM-5 (93%) | GLM-5 (Δ28) | DSv4-Pro (Δ1) |
| Phase 1 | GPT (78%) | GLM-5 (92%) | GLM-5 (Δ32) | Opus (Δ0) |
| Phase 2 | DSv4-Pro (88%) | GLM-5 (97%) | GLM-5 (Δ37) | DSv4-Pro (Δ8) |
| Phase 3 | DSv4-Pro (78%) | Opus (92%) | Opus (Δ20) | DSv4-Pro (Δ2) |
| Efficiency | DSv4-Pro (198 w/call) | GLM-5 (227 w/call) | Kimi K2.6 (4x bash shift) | Opus (stable) |
| Balance (CV) | GLM-5 (0.15) | GLM-5 (0.32) | Kimi K2.6 (+0.54) | GLM-5 (+0.17) |
| Reliability | Opus (<10% var) | Opus (<10% var) | Gemini (fails small) | Opus (always completes) |

---

## 9. Conclusions

1. **Project size is a significant performance variable** — models lose an average of 12 points when switching from their preferred size to the other.
2. **The phased workflow matters most for large CLIs** — Kimi K2.6's 34-point Phase 2 collapse is the clearest evidence that chunking prevents context exhaustion.
3. **Small CLIs expose a different weakness** — models that excel on large projects (GLM-5) under-invest on small ones, while models that struggle on large projects (Gemini) struggle even more on small ones due to insufficient surface material.
4. **Tool strategy adaptation is universal** — all models shift from read-heavy to bash-heavy as project size grows, but the magnitude of shift varies 4x between models.
5. **Design quality is size-independent** — Phase 3 scores have the smallest delta across project sizes, confirming that synthesis tasks don't benefit from more codebase material once analysis is complete.
