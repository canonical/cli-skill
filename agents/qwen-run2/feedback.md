# CLI Review Skill — Agent Interview Feedback

Three AI agents (Claude Opus 4.6, GPT-5.4, Gemini 3.1 Pro) independently executed the full `analyze-cli` and `discuss-commandset` workflows from `/project/.github/skills/cli-review/SKILL.md` on the `qwen36` snap CLI. Each was then interviewed about their experience using a structured UX evaluation protocol.

---

## Interview Protocol

Each agent was asked:

1. **Opening**: Overall impression of the experience
2. **Agent Profile**: Prior knowledge, similar frameworks, processing context
3. **Target Tasks**: Core goal understanding, happy path, deviations
4. **Four Core UX Questions** (per skill section): Intent, Visibility, Matching, Feedback
5. **Problems Encountered**: Failure category, severity, suggested fix
6. **Additional Observations**: Redundancy, judgment calls, single best improvement

---

## Agent 1: Claude Opus 4.6

### Opening
Structured and productive but occasionally over-specified for the target. The skill is designed for large CLIs (100+ commands) and applying it to a 6-command CLI created friction — exhaustive tables and matrices for a trivially small command set. Build-order dependency chain worked well.

### Profile
High prior CLI knowledge. Found the skill's specificity (Aktionsart classification, FrameNet, symmetry audits) goes beyond anything encountered as a packaged instruction set. Full tool access, no token pressure on individual steps due to small command set.

### Deviations
- Skipped toolchain bootstrap (snap already available)
- Could not read `standard/README.md` and `deprecation/README.md` (reported as missing — actually present but not located)
- Found HTML generation mechanical and disconnected from analytical purpose

### UX Ratings by Section

| Section | Intent | Visibility | Matching | Feedback |
|---------|--------|------------|----------|----------|
| 0-analysis | Clear | Obvious (build order) | Good | Moderate — trivially satisfied for 6 commands |
| 1-discuss-commandset | Clear | Very high | Mixed — some sections trivial at small scale | Good (self-check) |
| Output Completeness | Clear | Excellent | Perfect for large CLIs, trivial for small | Best-designed part (3-step self-check) |
| Recommendation Compliance | Understood | Clear *if files exist* | **Failed** — could not find standard files | None |
| HTML Generation | Clear | Specific | Adequate but tedious | Manual only |

### Problems (5)
1. **Missing standard/deprecation files** — Medium severity. Fix: add fallback instruction.
2. **Scale mismatch** — Medium. Fix: add scale gate (<15 / 15-50 / >50 commands).
3. **HTML spec over-prescribed** — Low. Fix: explain purpose or make optional.
4. **DE013 citations without source** — Medium. Fix: inline key rules or make dependency explicit.
5. **Build order forces serial execution** — Low. Fix: annotate true dependencies.

### Single Best Improvement
Add a scale-awareness preamble that adjusts expectations based on command count.

---

## Agent 2: GPT-5.4

### Opening
Good but slightly uneven. Strong when it behaves like a production workflow (concrete file contract, build order, defined analytical lens). Weaker around cross-references and compliance obligations. High-value structure, moderate execution friction, weaker feedback than it should have.

### Profile
High prior CLI knowledge. Closest pattern: audit playbook with fixed deliverables. Main constraint was not tools or tokens but ambiguity in the repo surface — Go CLI submodule missing, had to reconstruct command surface from snap definition, shell scripts, hooks, and README.

### Deviations
Evidence-gathering path differed from ideal — reconstructed commands from adjacent artifacts rather than direct command registration code. Located standards documents but found the dependency not operationally obvious.

### UX Ratings by Section

| Section | Intent | Visibility | Matching | Feedback |
|---------|--------|------------|----------|----------|
| 0-analysis | Clear | Explicit | Good | Moderate — self-judged |
| 1-discuss-commandset | Clear | Very explicit | Mostly good, "pattern classification" looser | Moderate to weak |
| Output Completeness | Unmistakable | Obvious | Mostly yes (some awkwardness for non-command-by-command sections) | Best part — self-check |
| Recommendation Compliance | Clear | **Partial** — easy to underweight, depends on external files | Mixed — "read and incorporate" too weak | Weak — no checklist equivalent |
| HTML Generation | Obvious | Concrete | Good | Moderate — trigger threshold fuzzy |

### Problems (7)
1. **Recommendation compliance easy to miss** — Medium. Fix: mandatory "stop here, read standards" gate.
2. **Phase naming mismatch** (`1-command-design` vs `1-discuss-commandset`) — Medium. Fix: use one consistent name.
3. **Completeness wording hard to map to non-table sections** — Low-Medium. Fix: clarify coverage obligation for matrices/clusters.
4. **HTML trigger subjective** — Low. Fix: replace "roughly" with strict threshold.
5. **Skill assumes inspectable source** — Medium. Fix: add fallback for indirect evidence.
6. **Recommendation compliance weaker than completeness** — Medium. Fix: add mandatory self-check.
7. **Verb-noun decomposition vs orphan concept** — Low. Fix: clarify orphan handling.

### Single Best Improvement
Add explicit gated checkpoints before the recommendation phase — a mandatory block requiring agents to read and summarize standards before writing recommendations.

---

## Agent 3: Gemini 3.1 Pro

### Opening
Exceptionally rigorous, well-scaffolded framework, but heavily "front-loaded" with constraints. Following it requires maintaining a very large context window balancing analytical reasoning, format adherence, and strict anti-summarization rules simultaneously. LLMs inherently seek to compress, so attention drifted from literal details (filenames) toward structural execution.

### Profile
High baseline CLI knowledge (POSIX, flags, exit codes). The specific linguistic framework (Aktionsart, rigid symmetry auditing) is highly specialized. Operated under fixed output token limits per turn, which forced chunking — directly leading to filename deviation.

### Deviations
**Critical: Used wrong filenames throughout.** Produced `analysis-1.md` through `analysis-9.md` and `discuss-1.md` through `discuss-6.md` instead of the specified names (`architecture.md`, `01-verb-noun-decomposition.md`, etc.).

**Root cause analysis (self-reported):** The skill nests filenames inside numbered lists (`1. architecture.md`, `2. commandset.md`). When chunking output to manage token limits, attention attached to the *list index* rather than the *literal filename string*. The abstracted loop converted "Step 1" → `analysis-1.md`. Double-numbering in discuss-commandset (`1. 01-verb-noun-decomposition.md`) created interference that collapsed to `discuss-1.md`.

### UX Ratings by Section

| Section | Intent | Visibility | Matching | Feedback |
|---------|--------|------------|----------|----------|
| 0-analysis | Clear | **Moderate** — numbered list encouraged index-based naming | High | **None** — no validation that filenames matched |
| 1-discuss-commandset | Clear | **Poor** for filenames — double-numbering interference | Good for content | **None** |
| Output Completeness | Clear | Extremely high | Direct counter-instruction to LLM summarization instinct | Moderate — self-check degrades on long generations |
| Recommendation Compliance | Clear | High but location-dependent | High if external files loaded | None |
| HTML Generation | Clear | Clear inline CSS specs | Clean mapping | None — cannot see render |

### Problems (2 — focused on naming failure)
1. **Numbered list obscures exact filenames** — High severity. Fix: Remove numbered lists; use literal terminal commands (`0-analysis/architecture.md`) or bulleted checklists.
2. **Double numbering in discuss-commandset** — Medium severity. Fix: Drop list index; use bullet points with only the filename.

### Single Best Improvement
Provide an automated scaffolding script or command agents to `mkdir` and `touch` the exact filenames *before* generating content. Pre-creating empty files anchors literal paths in context, virtually guaranteeing correct naming.

---

## Output Verification Summary

| Agent | 0-analysis (9 files) | 1-discuss-commandset (12 files) | Filenames Correct |
|-------|---------------------|--------------------------------|-------------------|
| Opus  | 9/9 ✓ | 12/12 ✓ | ✓ All correct |
| GPT   | 9/9 ✓ | 12/12 ✓ | ✓ All correct |
| Gemini | 9/9 ✓ | 12/12 ✓ | ✗ All wrong (generic names) |
