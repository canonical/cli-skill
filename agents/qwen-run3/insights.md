# CLI Review Skill — Run 3 Insights

Analysis of 4-agent evaluation (Opus, GPT, Gemini, GPT-mini) on the updated SKILL.md after implementing fixes 1-7 from run2.

---

## 1. Fix Validation — Did Run 2 Fixes Work?

| Fix | Status | Evidence |
|-----|--------|----------|
| 1. Compliance self-check | ✓ Working | Opus explicitly cited it as effective; GPT confirmed compliance verification worked |
| 2. Literal path blocks for filenames | ✓ Working | **Gemini now produces correct filenames** (was broken in run2). All 3 successful agents got filenames right |
| 3. Explicit standard/deprecation paths | ✓ Working | All agents found and read the resource files. No "file not found" reports |
| 4. Mandatory stop-and-read gate | ✓ Working | Opus praised the Recommendation Compliance gate as effective. GPT confirmed it improved output quality |
| 5. Scale gate | ⚠ Partially working | All agents recognized compact mode, but it created a NEW problem — conflict between "MAY combine" and exact filename requirements |
| 6. Phase naming fix | ✓ Working | No agent reported phase naming confusion |
| 7. Missing-source fallback | ⚠ Partially working | All agents used the fallback successfully, but Opus found contradictory "STOP" vs "reconstruct" language |

**Score: 5/7 fully working, 2/7 need refinement.**

---

## 2. New Issues Introduced by Fixes

### The Compact Mode Paradox (Fix 5 side-effect)
The scale gate introduced a new ambiguity that ALL agents flagged: compact mode says "MAY combine sections 03-05" but the filename spec demands all 6 files exist. Agents resolved this differently:
- **Opus**: Combined content in 03, created stub files for 04-05
- **GPT**: Created all 12 full files, ignoring compact mode
- **Gemini**: Created all 12 full files, explicitly choosing safety over optimization

This is now the **#1 skill issue** — it replaced the filename format problem that was #1 in run2.

---

## 3. Persistent Issues (Present in Both Run 2 and Run 3)

| Issue | Run 2 Status | Run 3 Status |
|-------|-------------|-------------|
| Scale mismatch (skill designed for large CLIs) | Reported by Opus | Reported by all — now manifests as compact mode confusion |
| HTML purpose unexplained | Reported by all | Reported by Gemini (tedious for small CLIs) |
| No guidance on source/shipped mismatch | Not reported | NEW from GPT — possibly because GPT read more deeply |
| Missing-source fallback ambiguity | Not specifically flagged | Flagged by Opus as contradictory |

---

## 4. Model Capability Observations

### Opus (Claude Opus 4.6)
- **Output**: 22 files, all correct names, 186-line feedback
- **Behavior**: Made the most sophisticated judgment calls (combined 03-05 with stubs, counted hidden vs visible commands). Most thorough self-interview.
- **Unique finding**: Analysis Checklist doesn't map to output files

### GPT (GPT-5.4)
- **Output**: 22 files, all correct names, 286-line feedback
- **Behavior**: Most operationally aware — identified the source/shipped surface mismatch that no other agent noticed. Longest and most detailed feedback.
- **Unique finding**: No execution model for reality mismatches; GitHub issue filing assumes tooling

### Gemini (Gemini 3.1 Pro)
- **Output**: 22 files, all correct names, 40-line feedback
- **Behavior**: Conservative executor — chose safety over optimization on every ambiguous rule. Shortest feedback but correct execution.
- **Key improvement from run2**: Filename fix completely resolved the naming failure

### GPT-mini (GPT-5 mini)
- **Output**: Incomplete — missing 6+ files, created .json files, no feedback
- **Behavior**: Appears to have hit capacity limits. Produced analysis files but couldn't complete discuss-commandset. Created unexpected .json format files.
- **Assessment**: Model is not suitable for this skill's output volume

---

## 5. Emotional Landscape

All agents reported similar emotional trajectories:
- **Satisfaction** from clear structure and concrete deliverables
- **Mild frustration** from compact mode ambiguity
- **Confidence** in output quality
- **Uncertainty** about interpretation correctness

The skill's emotional signature is: "I know what to do but not always how much to do."

---

## 6. Top 10 Issues (Ranked by Severity × Consensus)

| Rank | Issue | Severity | Agents | Fix |
|------|-------|----------|--------|-----|
| 1 | Compact mode spec is contradictory — "MAY combine" vs exact filenames | high | Opus, GPT, Gemini | Write explicit compact mode spec: which files exist, content policy, HTML policy |
| 2 | No guidance for source/shipped command surface mismatch | high | GPT | Add rule: validate built CLI, compare to source, document discrepancy, declare normative surface |
| 3 | GitHub issue filing in feedback assumes tooling exists | high | GPT | Add fallback: if no issue tool, write issue drafts in feedback.md using the template |
| 4 | Missing-source fallback has contradictory STOP vs reconstruct instructions | medium | Opus | Split into: (a) no source at all → STOP, (b) partial source → reconstruct with confidence |
| 5 | HTML fidelity between .md and .html undefined | medium | GPT | State: "HTML must mirror Markdown content exactly, with added visual formatting" |
| 6 | Output completeness validation is distributed across the document | medium | GPT | Add a final consolidated validation checklist before the feedback phase |
| 7 | No stopping rule for evidence gathering in large repos | medium | GPT | Add: "Read enough to identify public command surface, config model, output formats, mutation paths" |
| 8 | "Focus analytical depth" in compact mode is vague | medium | Opus | Quantify: "01 and 06 full depth; 02-05 abbreviated to key findings" |
| 9 | HTML trigger threshold conflicts with compact mode threshold | medium | Opus | Single rule: "compact mode → HTML only if table exceeds 15 rows; standard/full → always" |
| 10 | GPT-mini produced incomplete output and wrong formats | medium | GPT-mini | Note in skill: "This skill requires models with sufficient output capacity. Smaller models may not complete the full workflow." |

---

## 7. Key Insight

**The fundamental pattern from run2 holds and is confirmed**: any requirement the skill wants agents to follow must have a **verification step**, not just an instruction.

Run3 adds a corollary: **any flexibility the skill offers must have a decision procedure**, not just a permission. "MAY combine" without specifying when to combine, what the result looks like, and how to handle file existence requirements creates more confusion than having no flexibility at all.

The skill's strongest sections (Output Completeness, Recommendation Compliance) work because they pair instructions with self-checks. The weakest section (Scale Awareness) fails because it offers flexibility without a decision algorithm.
