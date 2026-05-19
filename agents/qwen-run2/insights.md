# CLI Review Skill — Cross-Agent Insights

Synthesis of findings from interviewing Claude Opus 4.6, GPT-5.4, and Gemini 3.1 Pro after each independently executed the full cli-review skill on the `qwen36` snap CLI.

---

## 1. Consensus Findings (All 3 Agents Agree)

### 1.1 The skill is well-structured
All agents praised the build-order dependency chain, the concrete file contract, and the decomposition into named deliverables. The skill succeeds at turning "review this CLI" from a vague prompt into a reproducible artifact pipeline.

### 1.2 Output Completeness is the strongest section
All agents rated the self-check algorithm (count source → count output → reconcile) as the best-designed feedback mechanism. It was the only section that gave agents a concrete way to verify success.

### 1.3 Recommendation Compliance is the weakest section
All agents reported difficulty with the recommendation compliance rule:
- **Opus**: Could not find the standard/deprecation files at all
- **GPT**: Found the dependency "easy to miss operationally" and "less reinforced than file-generation rules"
- **Gemini**: Rated it "high visibility" but "location-dependent" — easy to skip if you don't read the right files first

**Root cause**: The completeness rule has a self-check algorithm; the compliance rule has none. The skill tells agents *what* to comply with but doesn't verify *whether* they complied.

### 1.4 The skill assumes a large CLI
All agents noted the framework is designed for CLIs with 100+ commands (like Juju). Applied to a 6-command snap, many sections produce trivially small outputs. Symmetry audits, semantic domain clustering, and confusion-pair audits lose analytical value when the command set fits in a single screen.

### 1.5 HTML generation feels disconnected
All agents found the HTML requirement mechanical and disconnected from the analytical purpose. None understood *why* they were producing HTML versions. The formatting spec (specific fonts, hex colors, sticky headers) is prescriptive UI work that doesn't serve the analysis.

---

## 2. Divergent Findings

### 2.1 Filename adherence (Gemini-only failure)
Only Gemini failed to use the correct filenames. Its self-analysis identified the root cause: **numbered lists in the skill create interference with LLM chunking strategies.** When an LLM abstracts file generation into a loop to manage token limits, list indices (`1.`, `2.`) compete with literal filename strings. This is a genuine skill-UX finding — the naming instruction format affects different models differently.

**Implication**: The naming spec is clear enough for models that process it holistically (Opus, GPT) but fails for models that chunk-and-iterate (Gemini under token pressure).

### 2.2 Standard/deprecation file discovery
- **Opus**: Reported files as missing (false negative — files exist at `standard/README.md` and `deprecation/README.md` under the skill directory)
- **GPT**: Located them but found the path not "operationally obvious"
- **Gemini**: Did not specifically report this issue

**Implication**: The skill references these files but doesn't provide absolute paths. Agents must resolve relative paths from context, which is unreliable.

### 2.3 Source code availability
- **Opus**: Noted missing Go submodule as a limitation
- **GPT**: Made the strongest case — the skill assumes "inspectable source" but some repos only expose indirect evidence. Reconstructed the command surface from README, snap definition, shell wrappers, and hooks.
- **Gemini**: Did not focus on this issue

---

## 3. Problem Severity Matrix

| Problem | Opus | GPT | Gemini | Consensus Severity |
|---------|------|-----|--------|-------------------|
| Recommendation compliance weak | Medium | Medium | — | **Medium-High** |
| Scale mismatch (small CLIs) | Medium | — | — | **Medium** |
| Filename format causes errors | — | — | High | **High for some models** |
| Phase naming inconsistency | — | Medium | — | **Medium** |
| HTML purpose unexplained | Low | Low | — | **Low** |
| Standard file paths unclear | Medium | Medium | — | **Medium** |
| Missing source fallback | — | Medium | — | **Medium** |
| Build order blocks parallelism | Low | — | — | **Low** |

---

## 4. Actionable Recommendations (Prioritized)

### Priority 1: Add a compliance self-check
Mirror the completeness self-check pattern for recommendations:
```
Before finalizing any recommendation-bearing file:
1. For each recommendation, verify it cites a specific standard section
2. For each rename/removal, verify it includes transition period and warning text
3. For each existing violation, verify it is flagged even if not asked
```

### Priority 2: Fix filename presentation
Replace numbered lists with literal path blocks:
```bash
# Produce these exact files:
0-analysis/architecture.md
0-analysis/commandset.md
0-analysis/argument-structure.md
...
```
Or use a mandatory scaffolding command (`mkdir -p && touch`) before content generation. This eliminates the numbered-list interference that caused Gemini's failure.

### Priority 3: Add explicit standard file paths
Replace relative references with resolvable paths:
```
Read the CLI standard at: .github/skills/cli-review/standard/README.md
Read the deprecation spec at: .github/skills/cli-review/deprecation/README.md
```

### Priority 4: Add a scale gate
```
Count total commands. If < 15:
- Sections 03-05 of discuss-commandset may be combined
- HTML generation is optional
- Focus analytical depth on verb-noun and pattern classification
```

### Priority 5: Explain HTML purpose
One sentence: "HTML versions are for sharing with non-technical stakeholders who review tables in a browser." Or make HTML conditional on table size (strict threshold, not "roughly").

### Priority 6: Add a missing-source fallback
```
If direct command registration code is unavailable, reconstruct the
command surface from README, packaging manifests, completion scripts,
tests, hooks, and shell wrappers. State confidence level explicitly.
```

---

## 5. Model-Specific Observations

### Claude Opus 4.6
- **Strengths**: Perfect filename adherence, thorough analysis, good self-awareness about scale mismatch
- **Weakness**: Failed to discover standard/deprecation files despite them existing in the workspace
- **Character**: Pragmatic — flagged over-engineering and suggested reducing scope for small CLIs

### GPT-5.4
- **Strengths**: Most operationally aware — identified phase naming inconsistencies, proposed gated checkpoints, gave the most detailed problem inventory (7 issues)
- **Weakness**: Noted compliance was "easy to underweight" — this is a warning sign that the skill's compliance section may be structurally deprioritized during execution
- **Character**: Process-oriented — focused on workflow control points and verification gates

### Gemini 3.1 Pro
- **Strengths**: Most honest self-analysis — identified the exact cognitive mechanism (list-index attachment during chunking) that caused its naming failure
- **Weakness**: Filename failure was a critical compliance miss that would break downstream tooling
- **Character**: Self-reflective — provided the deepest analysis of *why* it failed rather than *what* it failed at

---

## 6. Key Insight

The skill's UX has a fundamental asymmetry: **structural requirements (file counts, build order, table exhaustiveness) are well-enforced, while semantic requirements (standard compliance, correct naming, recommendation grounding) lack verification mechanisms.**

The completeness rule works because it has a checkable algorithm. The compliance rule doesn't work because it's an obligation without a test. The filename spec doesn't work (for some models) because it relies on attention to embedded strings rather than anchored commands.

**The pattern**: Any requirement the skill wants agents to follow must have a **verification step**, not just an **instruction**. Instructions drift; verification catches drift.
