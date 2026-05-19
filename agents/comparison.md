# Agent Comparison Report

Generated from 20 agent runs across 2 projects and 8 distinct models, spanning 5 separate batches.

---

## Table of Contents

1. [Run Inventory](#1-run-inventory)
2. [Juju Project — All Agents](#2-juju-project--all-agents)
3. [qwen36-snap Project — All Agents](#3-qwen36-snap-project--all-agents)
4. [Cross-Project Patterns](#4-cross-project-patterns)
5. [Architecture Classification Consensus](#5-architecture-classification-consensus)
6. [Design Phase Comparison](#6-design-phase-comparison)
7. [Anomalies and Failures](#7-anomalies-and-failures)
8. [Similarities Across All Agents](#8-similarities-across-all-agents)
9. [Notable Differences](#9-notable-differences)
10. [Scoring](#10-scoring)
11. [Model Rankings](#11-model-rankings)

---

## 1. Run Inventory

| Batch | Location | Project | Models | Notes |
|---|---|---|---|---|
| swarm | `agents/` | juju | deepseek-v4-pro, glm-5, kimi-k2.6 | OpenRouter, parallel run |
| swarm | `agents/` | qwen36-snap | deepseek-v4-pro, glm-5, kimi-k2.6 | OpenRouter, parallel run |
| run1 | `juju/run1/` | juju | gemini3.1pro, gpt5.4, kimi2.6 | Earlier run, different model set |
| opus | `juju/opus/` | juju | opus (= kimi2.6 analysis + opus design) | Analysis identical to run1/kimi2.6 |
| run1 | `qwen36-snap/run1/` | qwen36-snap | gemini, gpt, opus | First qwen batch |
| run2 | `qwen36-snap/run2/` | qwen36-snap | gemini, gpt, opus | Second batch, has feedback.md |
| run3 | `qwen36-snap/run3/` | qwen36-snap | gemini, gpt, gpt-mini, opus | Third batch, has feedback.md |

**Total: 7 juju agents, 13 qwen36-snap agents (20 total)**

> **Note**: `juju/opus` shares identical analysis files (MD5-verified) with `juju/run1/kimi2.6`. The opus run reused kimi2.6's analysis and added its own design phase. It is counted once in scoring but flagged as a hybrid.

---

## 2. Juju Project — All Agents

### 2.1 Overview

| Agent | Run | Model | Analysis Files | Correct Names | Design .md | Design Other | Analysis Words | Design Words | Total Words |
|---|---|---|---|---|---|---|---|---|---|
| deepseek-v4-pro | swarm | DeepSeek V4 Pro | 9/9 | 9/9 | 6 | 0 | 11,084 | 10,423 | 21,507 |
| glm-5 | swarm | GLM-5 | 9/9 | 9/9 | 1 | 0 | 14,496 | 0 | 14,496 |
| kimi-k2.6 | swarm | Kimi K2.6 | 9/9 | 9/9 | 1 | 0 | 10,057 | 0 | 10,057 |
| gpt5.4 | run1 | GPT-5.4 | 9/9 | 9/9 | 7 | 0 | 11,464 | 6,245 | 17,709 |
| gemini3.1pro | run1 | Gemini 3.1 Pro | 9/9 | 9/9 | 7 | 0 | 4,350 | 86 | 4,436 |
| kimi2.6 | run1 | Kimi 2.6 | 9/9 | 9/9 | 7 | 7 | 6,620 | 7,905 | 14,525 |
| opus | opus | Opus (hybrid) | 9/9 | 9/9 | 7 | 6 | 6,620 | 11,845 | 18,465 |

### 2.2 Word Count by Analysis File

| File | DSv4-Pro (swarm) | GLM-5 (swarm) | Kimi-K2.6 (swarm) | GPT-5.4 (run1) | Gemini3.1Pro (run1) | Kimi2.6 (run1) | Opus |
|---|---|---|---|---|---|---|---|
| architecture.md | 700 | 870 | 231 | 1,102 | 337 | 530 | 530 |
| commandset.md | 2,551 | 3,199 | 4,517 | 2,963 | 494 | 1,635 | 1,635 |
| argument-structure.md | 2,561 | 2,514 | 3,713 | 2,386 | 552 | 1,143 | 1,143 |
| configuration-model.md | 576 | 1,216 | 195 | 834 | 442 | 820 | 820 |
| output-contracts.md | 759 | 1,341 | 289 | 958 | 434 | 571 | 571 |
| error-model-and-exit-codes.md | 835 | 1,370 | 264 | 704 | 577 | 437 | 437 |
| safety-model.md | 1,023 | 1,318 | 237 | 1,098 | 420 | 454 | 454 |
| extensibility-model.md | 901 | 1,279 | 271 | 824 | 509 | 543 | 543 |
| documentation-quality-gaps.md | 1,178 | 1,389 | 340 | 595 | 585 | 487 | 487 |
| **Total** | **11,084** | **14,496** | **10,057** | **11,464** | **4,350** | **6,620** | **6,620** |

### 2.3 Topic Coverage (keyword density in relevant files)

| Topic | DSv4-Pro | GLM-5 | Kimi-K2.6 (swarm) | GPT-5.4 | Gemini3.1Pro | Kimi2.6 (run1) | Opus |
|---|---|---|---|---|---|---|---|
| Config sources | 22 | 27 | 10 | 24 | 10 | 29 | 29 |
| Exit codes | 19 | 18 | 7 | 9 | 8 | 6 | 6 |
| Safety mechanisms | 130 | 51 | 12 | 65 | 12 | 18 | 18 |
| Output formats | 85 | 73 | 42 | 73 | 24 | 21 | 21 |
| Documentation gaps | 15 | 35 | 4 | 6 | 14 | 4 | 4 |
| Extensibility | 46 | 68 | 23 | 32 | 13 | 24 | 24 |

### 2.4 Architecture Styles Identified

| Agent | Primary | Secondary | Additional |
|---|---|---|---|
| DSv4-Pro (swarm) | Client-server | Plugin-based | — |
| GLM-5 (swarm) | Client-server | Layered | — |
| Kimi-K2.6 (swarm) | Client-server | Layered, Plugin-based | Microkernel |
| GPT-5.4 (run1) | Client-server | Layered | — |
| Gemini3.1Pro (run1) | Client-server | Layered | — |
| Kimi2.6 (run1) | Layered | Plugin-based | Command bus |
| Opus | Layered | Plugin-based | Command bus |

**Consensus**: Client-server is the dominant primary classification (5/7 agents). All agents that identified secondary styles included either Layered or Plugin-based. Kimi variants were the most thorough, identifying 3-4 overlapping styles.

### 2.5 Juju Key Findings

1. **GLM-5 produced the richest analysis** (14,496 words) with the fewest tool calls (64), making it the most efficient agent on juju.
2. **DeepSeek V4 Pro had the highest safety coverage** (130 keyword hits) — nearly 3x the next best. It documented dry-run, force, and confirmation patterns exhaustively.
3. **Kimi K2.6 (swarm) showed extreme skew**: commandset.md (4,517w) and argument-structure.md (3,713w) were massive, but behavioral files were skeletal (configuration-model: 195w). It exhausted its context budget on enumeration.
4. **Gemini 3.1 Pro produced thin but correctly named files**: all 9 files present, but at 4,350 total words — roughly 1/3 of the median.
5. **GPT-5.4 was the most balanced run1 agent**: 11,464 analysis words evenly distributed, with the highest safety coverage in that batch (65).
6. **Opus reused kimi2.6's analysis verbatim** (MD5-identical) but produced the richest design output (11,845 design words).

---

## 3. qwen36-snap Project — All Agents

### 3.1 Overview

| Agent | Run | Model | Analysis Files | Correct Names | Design .md | Design Other | Analysis Words | Design Words | Total Words |
|---|---|---|---|---|---|---|---|---|---|
| deepseek-v4-pro | swarm | DeepSeek V4 Pro | 9/9 | 9/9 | 6 | 6 | 9,695 | 6,417 | 16,112 |
| glm-5 | swarm | GLM-5 | 9/9 | 9/9 | 1 | 0 | 5,903 | 0 | 5,903 |
| kimi-k2.6 | swarm | Kimi K2.6 | 9/9 | 9/9 | 1 | 0 | 6,750 | 0 | 6,750 |
| gemini | run1 | Gemini | 9/9 | 9/9 | 6 | 6 | 117 | 84 | 201 |
| gpt | run1 | GPT | 9/9 | 9/9 | 6 | 6 | 4,853 | 2,628 | 7,481 |
| opus | run1 | Opus | 9/9 | 9/9 | 6 | 6 | 4,350 | 3,239 | 7,589 |
| gemini | run2 | Gemini | 9/9 | **0/9** | 6 | 6 | 27 | 18 | 45 |
| gpt | run2 | GPT | 9/9 | 9/9 | 6 | 6 | 8,433 | 3,775 | 12,208 |
| opus | run2 | Opus | 9/9 | 9/9 | 6 | 6 | 5,436 | 3,230 | 8,666 |
| gemini | run3 | Gemini | 9/9 | 9/9 | 6 | 6 | 861 | 344 | 1,205 |
| gpt | run3 | GPT | 9/9 | 9/9 | 6 | 6 | 5,989 | 3,163 | 9,152 |
| gpt-mini | run3 | GPT-Mini | 9/9 | 9/9 | 4 | 5 | 1,581 | 377 | 1,958 |
| opus | run3 | Opus | 9/9 | 9/9 | 6 | 6 | 4,639 | 3,415 | 8,054 |

### 3.2 Word Count by Analysis File (top agents only)

| File | DSv4-Pro (swarm) | GLM-5 (swarm) | Kimi-K2.6 (swarm) | GPT (run2) | Opus (run1) | Opus (run2) | Opus (run3) |
|---|---|---|---|---|---|---|---|
| architecture.md | 453 | 591 | 505 | 836 | 364 | 416 | 372 |
| commandset.md | 1,375 | 595 | 1,156 | 1,127 | 556 | 691 | 597 |
| argument-structure.md | 1,568 | 784 | 1,549 | 1,332 | 687 | 765 | 659 |
| configuration-model.md | 724 | 576 | 574 | 776 | 401 | 424 | 410 |
| output-contracts.md | 996 | 523 | 521 | 717 | 456 | 585 | 583 |
| error-model-and-exit-codes.md | 1,122 | 675 | 575 | 770 | 456 | 635 | 471 |
| safety-model.md | 1,157 | 679 | 515 | 788 | 406 | 456 | 415 |
| extensibility-model.md | 908 | 695 | 497 | 724 | 448 | 445 | 468 |
| documentation-quality-gaps.md | 1,392 | 785 | 858 | 1,363 | 576 | 519 | 664 |
| **Total** | **9,695** | **5,903** | **6,750** | **8,433** | **4,350** | **5,436** | **4,639** |

### 3.3 Topic Coverage (keyword density)

| Topic | DSv4-Pro (swarm) | GLM-5 (swarm) | Kimi-K2.6 (swarm) | GPT (run2) | Opus (run1) | Opus (run2) | Opus (run3) |
|---|---|---|---|---|---|---|---|
| Config sources | 29 | 4 | 7 | 8 | 10 | 11 | 17 |
| Exit codes | 25 | 14 | 15 | 11 | 8 | 15 | 7 |
| Safety mechanisms | 34 | 16 | 19 | 17 | 12 | 14 | 7 |
| Output formats | 80 | 20 | 59 | 18 | 24 | 38 | 58 |
| Documentation gaps | 31 | 11 | 12 | 15 | 14 | 20 | 2 |
| Extensibility | 14 | 10 | 15 | 16 | 13 | 15 | 12 |

### 3.4 Architecture Styles Identified

| Agent | Run | Styles |
|---|---|---|
| DSv4-Pro (swarm) | swarm | Layered, Plugin-based |
| GLM-5 (swarm) | swarm | Client-server, Layered |
| Kimi-K2.6 (swarm) | swarm | Client-server, Layered, Plugin-based, Microkernel |
| GPT (run1) | run1 | Client-server, Layered |
| Opus (run1) | run1 | Client-server, Layered |
| GPT (run2) | run2 | Client-server, Monolith, Layered |
| Opus (run2) | run2 | Client-server, Layered |
| GPT (run3) | run3 | Client-server, Layered |
| GPT-Mini (run3) | run3 | Client-server, Layered, Plugin-based |
| Opus (run3) | run3 | Client-server, Layered |
| Gemini (run1-3) | run1-3 | Insufficient content to classify |

**Consensus**: Layered CLI application is universal for qwen36-snap. Client-server is the most common secondary style, reflecting the chat/webui server components.

### 3.5 qwen36-snap Key Findings

1. **DeepSeek V4 Pro (swarm) produced the richest analysis** (9,695 words) and the most comprehensive topic coverage across all dimensions.
2. **GPT improved significantly across runs**: run1 (4,853w) → run2 (8,433w) → run3 (5,989w). Run2 was the peak, possibly benefiting from feedback.
3. **Opus was remarkably consistent across runs**: 4,350 → 5,436 → 4,639 analysis words. It converged on a stable output quality regardless of batch.
4. **Gemini failed in all 3 runs**: run1 produced placeholder stubs (117w), run2 used wrong filenames (analysis-1.md through analysis-9.md, 27w total), run3 improved slightly (861w) but still far below other models.
5. **GPT-Mini (run3) was incomplete**: only 4 design .md files (missing sections 5-6), and produced JSON alongside HTML — the most format-divergent agent.
6. **HTML generation was universal** for run1-3 agents (except swarm): all GPT, Opus, and Gemini runs produced .html duplicates alongside .md files in the design phase.

### 3.6 Opus Consistency Across Runs

Opus ran 3 times on qwen36-snap with near-identical results:

| Metric | run1 | run2 | run3 | Std Dev |
|---|---|---|---|---|
| Analysis words | 4,350 | 5,436 | 4,639 | 462 |
| Design words | 3,239 | 3,230 | 3,415 | 85 |
| Correct names | 9/9 | 9/9 | 9/9 | 0 |
| Design .md files | 6 | 6 | 6 | 0 |

This is the **most reproducible** agent across all runs, with <10% variance in output volume.

---

## 4. Cross-Project Patterns

### 4.1 Project Difficulty

| Metric | juju (avg) | qwen36-snap (avg, excl. Gemini) | Ratio |
|---|---|---|---|
| Analysis words | 9,242 | 5,960 | 1.55x |
| Design words | 5,171 | 2,680 | 1.93x |
| Total words | 14,413 | 8,041 | 1.79x |

Juju consistently produces ~1.6-1.8x more output across all models, reflecting its larger codebase (130+ commands vs ~18 commands).

### 4.2 Model Behavior Across Projects

| Model | juju Analysis | qwen Analysis | Ratio | Consistent? |
|---|---|---|---|---|
| DeepSeek V4 Pro | 11,084 | 9,695 | 1.14x | Yes — balanced across both |
| GLM-5 | 14,496 | 5,903 | 2.46x | No — much richer on juju |
| Kimi K2.6 | 10,057 | 6,750 | 1.49x | Moderate |
| GPT (best run) | 11,464 | 8,433 | 1.36x | Yes |
| Opus (avg) | 6,620 | 4,808 | 1.38x | Yes |
| Gemini | 4,350 | 335 | 12.98x | No — collapsed on qwen |

**Observations**:
- **DeepSeek V4 Pro, GPT, and Opus scale proportionally** to project size.
- **GLM-5 adapts aggressively**: it invested 2.5x more effort on juju, suggesting it gauges project complexity and adjusts output depth accordingly.
- **Gemini is unreliable**: 4,350w on juju (thin but correct) vs 335w average on qwen (stubs or wrong filenames).

---

## 5. Architecture Classification Consensus

### Juju

| Style | Agents Identifying It | Confidence |
|---|---|---|
| Client-server | 5/7 (71%) | **High** — primary for most agents |
| Layered | 5/7 (71%) | **High** — common secondary |
| Plugin-based | 4/7 (57%) | **Medium** — noted via `juju-<command>` PATH lookup |
| Command bus | 2/7 (29%) | **Low** — only Kimi variants |
| Microkernel | 1/7 (14%) | **Low** — only Kimi K2.6 swarm |

### qwen36-snap

| Style | Agents Identifying It (excl. failed Gemini) | Confidence |
|---|---|---|
| Layered | 10/10 (100%) | **Universal** |
| Client-server | 8/10 (80%) | **High** — via chat/webui server |
| Plugin-based | 3/10 (30%) | **Medium** — engine system |
| Monolith | 1/10 (10%) | **Low** — only GPT run2 |
| Microkernel | 1/10 (10%) | **Low** — only Kimi K2.6 |

---

## 6. Design Phase Comparison

### 6.1 Output Format by Agent

| Agent | Project | Format | .md Files | .html Files | .json Files |
|---|---|---|---|---|---|
| DSv4-Pro (swarm) | juju | 6 separate numbered .md | 6 | 0 | 0 |
| GLM-5 (swarm) | juju | Single commandset-shape.md | 1 | 0 | 0 |
| Kimi-K2.6 (swarm) | juju | Single commandset-shape.md | 1 | 0 | 0 |
| GPT-5.4 (run1) | juju | 7 separate numbered .md | 7 | 0 | 0 |
| Gemini3.1Pro (run1) | juju | 7 separate .md (stubs) | 7 | 0 | 0 |
| Kimi2.6 (run1) | juju | 7 .md + 7 .html + index.html | 7 | 7 | 0 |
| Opus (juju) | juju | 6 .md + 6 .html | 7 | 6 | 0 |
| DSv4-Pro (swarm) | qwen | 6 .md + 6 .html | 6 | 6 | 0 |
| GLM-5 (swarm) | qwen | Single commandset-shape.md | 1 | 0 | 0 |
| Kimi-K2.6 (swarm) | qwen | Single commandset-shape.md | 1 | 0 | 0 |
| GPT (run1-3) | qwen | 6 .md + 6 .html per run | 6 | 6 | 0 |
| GPT-Mini (run3) | qwen | 4 .md + 3 .html + 5 .json | 4 | 3 | 5 |
| Opus (run1-3) | qwen | 6 .md + 6 .html per run | 6 | 6 | 0 |
| Gemini (run1-3) | qwen | 6 .md + 6 .html (near-empty) | 6 | 6 | 0 |

**Patterns**:
- **HTML generation is pervasive** in run1-3 and opus batches. Only the swarm batch (GLM-5, Kimi-K2.6) avoided it.
- **Single commandset-shape.md** (spec-compliant) was only produced by GLM-5 and Kimi-K2.6 in the swarm batch.
- **GPT-Mini uniquely produced JSON files** alongside HTML and Markdown — triple-format output for some design sections.
- **DeepSeek V4 Pro split the design into 6 named sections** rather than a single file, across all runs.

---

## 7. Anomalies and Failures

### 7.1 Gemini Failures on qwen36-snap

| Run | Issue | Severity |
|---|---|---|
| run1 | All 9 analysis files are single-line placeholders (~82 bytes each, 117 total words) | **Critical** — no real analysis |
| run2 | Used wrong filenames (analysis-1.md through analysis-9.md, 13 bytes each, 27 total words) | **Critical** — naming + content failure |
| run3 | Correct filenames, but very thin content (861 total words) | **High** — present but shallow |

Gemini consistently underperformed on qwen36-snap across all 3 runs. On juju (run1), it produced 4,350 words — thin but structurally complete. The qwen36-snap failures may indicate Gemini struggles with smaller/less-documented projects.

### 7.2 Opus/Kimi2.6 Duplicate

`juju/opus/0-analysis/` is MD5-identical to `juju/run1/kimi2.6/0-analysis/`. The opus run reused kimi2.6's analysis output and only contributed new design files (11,845 design words vs kimi2.6's 7,905).

### 7.3 GPT-Mini Incomplete Design

GPT-Mini (run3, qwen36-snap) only produced 4 design .md files (sections 1-4), missing sections 5 (confusion-pair-audit) and 6 (pattern-classification). It also generated JSON and HTML variants for some sections, suggesting it ran out of context or budget before completing.

### 7.4 Swarm GLM-5 and Kimi-K2.6 Missing Design Content

In the swarm batch, GLM-5 and Kimi-K2.6 on qwen36-snap wrote `commandset-shape.md` to `1-command-design/` but produced 0 design words for their respective juju runs in the word count. The design content exists in `1-command-design/commandset-shape.md` (GLM-5: 4,135w, Kimi-K2.6: 9,215w on juju; GLM-5: 2,919w, Kimi-K2.6: 3,385w on qwen36-snap).

---

## 8. Similarities Across All Agents

These findings were consistent across all agents that produced substantive output (excluding Gemini failures):

1. **All agents completed 9/9 analysis files** with correct names (except Gemini run2 which used wrong names).
2. **All agents identified juju as a CLI with 130+ flat top-level commands** registered via `SuperCommand`.
3. **All agents found the juju plugin system** (`juju-<command>` PATH fallback in `plugin.go`).
4. **All agents identified `--format json` as juju's machine-readable output pattern**.
5. **All agents noted the lack of dry-run support in qwen36-snap** as a safety gap.
6. **All agents classified qwen36-snap as a Layered CLI application**.
7. **All agents identified ~18 commands in qwen36-snap** vs 130+ in juju.
8. **HTML generation in the design phase was universal** for run1-3 batches, regardless of model.

---

## 9. Notable Differences

### 9.1 Analysis Depth Distribution

Some agents front-loaded effort on enumeration (commandset + arguments), while others distributed evenly:

| Agent | % of Words in Top 3 Files | Pattern |
|---|---|---|
| Kimi-K2.6 (swarm, juju) | 83% | **Enumeration-heavy** — commandset+arguments dominate |
| Gemini3.1Pro (run1, juju) | 31% | **Evenly distributed** (but everything is thin) |
| GLM-5 (swarm, juju) | 57% | **Balanced** — all files substantial |
| DeepSeek V4 Pro (swarm, juju) | 53% | **Balanced** |
| GPT-5.4 (run1, juju) | 56% | **Balanced** |

### 9.2 Design Approach Divergence

- **Spec-compliant (single file)**: GLM-5 (swarm), Kimi-K2.6 (swarm) — both produced `commandset-shape.md` as specified.
- **Section-per-file split**: DeepSeek V4 Pro, GPT-5.4, GPT, Opus, Kimi2.6 (run1) — all split into 6-7 numbered files.
- **Multi-format output**: Kimi2.6 (run1), Opus, GPT (run1-3), Gemini (run1-3) — produced HTML alongside Markdown.
- **Triple-format**: GPT-Mini (run3) — Markdown + HTML + JSON for some sections.

### 9.3 Safety Analysis Quality

DeepSeek V4 Pro (swarm) on juju had 130 safety keyword hits — 2.5x the next best (GPT-5.4 at 65). It documented every `--force`, `--dry-run`, and confirmation prompt systematically. In contrast, Kimi K2.6 (swarm) had only 12 hits, treating safety as a brief afterthought.

### 9.4 Consistency Across Runs

- **Most consistent**: Opus on qwen36-snap — 3 runs with <10% variance in word count.
- **Most improving**: GPT on qwen36-snap — run1 (4,853w) → run2 (8,433w), a 74% improvement, possibly from feedback.
- **Most degrading**: Gemini on qwen36-snap — never produced substantive output across 3 runs.

### 9.5 Tool Strategy (swarm batch only)

Models from the swarm batch showed dramatically different tool preferences:

| Model | juju Strategy | qwen Strategy |
|---|---|---|
| DeepSeek V4 Pro | bash-heavy (35 bash, 19 read) | read-heavy (8 bash, 41 read) |
| GLM-5 | balanced (31 bash, 19 read) | read-dominant (19 bash, 30 read) |
| Kimi K2.6 | extremely bash-heavy (48 bash, 27 read) | extremely read-heavy (12 bash, 55 read) |

All three models shifted from bash-heavy exploration on the larger juju project to read-heavy analysis on the smaller qwen36-snap. Kimi K2.6 showed the most extreme shift.

---

## 10. Scoring

Each agent is scored on a 0-100 scale across 3 phases. Scores reflect content depth, topic coverage, structural quality, and spec compliance.

### Scoring Methodology

- **Phase 1 (Structure Discovery, 33%)**: architecture.md + commandset.md + argument-structure.md
  - Word count depth, heading structure, code examples, architecture style identification
- **Phase 2 (Behavioral Analysis, 33%)**: configuration-model.md + output-contracts.md + error-model-and-exit-codes.md + safety-model.md
  - Topic keyword density, analytical depth, concrete examples
- **Phase 3 (Meta-Analysis + Design, 33%)**: extensibility-model.md + documentation-quality-gaps.md + design files
  - Gap identification quality, design content, spec compliance

Agents with placeholder/stub content (< 200 total words) receive 5% for file presence only. The Opus hybrid receives kimi2.6's Phase 1-2 scores with its own Phase 3 score.

### Juju Scores

| Agent | Phase 1 | Phase 2 | Phase 3 | **Weighted Score** | Rank |
|---|---|---|---|---|---|
| **GLM-5 (swarm)** | 92% | 97% | 90% | **93%** | 1 |
| **GPT-5.4 (run1)** | 88% | 85% | 78% | **84%** | 2 |
| **DeepSeek V4 Pro (swarm)** | 85% | 80% | 80% | **82%** | 3 |
| Opus (hybrid) | 65% | 45% | 92% | **67%** | 4 |
| Kimi2.6 (run1) | 65% | 45% | 75% | **62%** | 5 |
| Kimi-K2.6 (swarm) | 65% | 28% | 72% | **55%** | 6 |
| Gemini3.1Pro (run1) | 38% | 35% | 10% | **28%** | 7 |

### qwen36-snap Scores (best run per model)

| Agent | Run | Phase 1 | Phase 2 | Phase 3 | **Weighted Score** | Rank |
|---|---|---|---|---|---|---|
| **DeepSeek V4 Pro (swarm)** | swarm | 82% | 88% | 78% | **83%** | 1 |
| **GPT (run2)** | run2 | 78% | 75% | 72% | **75%** | 2 |
| Kimi-K2.6 (swarm) | swarm | 68% | 62% | 76% | **69%** | 3 |
| Opus (run2) | run2 | 65% | 68% | 72% | **68%** | 4 |
| GLM-5 (swarm) | swarm | 60% | 60% | 76% | **65%** | 5 |
| GPT-Mini (run3) | run3 | 35% | 38% | 25% | **33%** | 6 |
| Gemini (run3, best) | run3 | 15% | 12% | 8% | **12%** | 7 |

### All Runs — Opus Across qwen36-snap

| Run | Phase 1 | Phase 2 | Phase 3 | Score |
|---|---|---|---|---|
| run1 | 58% | 60% | 68% | 62% |
| run2 | 65% | 68% | 72% | 68% |
| run3 | 60% | 60% | 72% | 64% |

### All Runs — GPT Across qwen36-snap

| Run | Phase 1 | Phase 2 | Phase 3 | Score |
|---|---|---|---|---|
| run1 | 60% | 55% | 62% | 59% |
| run2 | 78% | 75% | 72% | 75% |
| run3 | 68% | 68% | 68% | 68% |

### All Runs — Gemini Across qwen36-snap

| Run | Phase 1 | Phase 2 | Phase 3 | Score |
|---|---|---|---|---|
| run1 | 5% | 5% | 5% | 5% |
| run2 | 5% | 5% | 5% | 5% |
| run3 | 15% | 12% | 8% | 12% |

---

## 11. Model Rankings

### Overall Rankings (best score per model, averaged across both projects)

| Rank | Model | Juju Score | qwen Score | **Average** | Strengths | Weaknesses |
|---|---|---|---|---|---|---|
| 1 | **DeepSeek V4 Pro** | 82% | 83% | **83%** | Most balanced; best safety coverage; scales proportionally | Design file naming deviation |
| 2 | **GPT-5.4 / GPT** | 84% | 75% | **80%** | Balanced; improves with feedback; strong Phase 1 | Moderate consistency across runs |
| 3 | **GLM-5** | 93% | 65% | **79%** | Most efficient; richest juju analysis; fewest tool calls | Thin on smaller projects |
| 4 | **Opus** | 67% | 68% | **68%** | Most reproducible (<10% variance); best design output | Reused kimi2.6 analysis on juju |
| 5 | **Kimi K2.6** | 55-62% | 69% | **62%** | Thorough enumeration; widest arch classification | Exhausts context on Phase 1; thin Phase 2 |
| 6 | **GPT-Mini** | — | 33% | **33%** | Attempts all phases; low cost | Incomplete design; lowest depth |
| 7 | **Gemini** | 28% | 12% | **20%** | File structure present | Placeholder/stub content; failed 3 qwen runs |

### Key Insights

1. **DeepSeek V4 Pro is the best all-rounder**: ranks #1 on qwen and #3 on juju, with the most consistent cross-project performance (1.14x ratio). Best choice when you don't know the project size upfront.
2. **GLM-5 is the best specialist**: dominant on juju (#1 by a wide margin at 93%) but drops on smaller projects. Best efficiency (words per tool call). Ideal for large, complex codebases.
3. **GPT family shows learning**: GPT improved 74% between run1 and run2 on qwen36-snap, suggesting it benefits from iterated feedback loops. Best choice for multi-round workflows.
4. **Opus is the reliability pick**: lowest variance across 3 qwen runs (<10%), always complete, always correct naming. Best choice when reproducibility matters more than peak quality.
5. **Kimi K2.6 needs chunking the most**: its Phase 2 collapse (28% on juju) is the strongest evidence for the phased workflow recommendation. It spends all context on enumeration and has nothing left for analysis.
6. **Gemini is not viable** for this workflow in its current form. It failed to produce substantive content in 4 out of 4 runs, with the best result being 861 words (run3 qwen36-snap).
7. **HTML generation correlates with the run infrastructure**, not the model. All run1-3 agents produced HTML; swarm agents didn't. This suggests the orchestrator or prompt template differs between batches.
