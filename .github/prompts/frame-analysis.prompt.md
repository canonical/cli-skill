---
description: "Analyze CLI verb semantics using FrameNet: frame lookup, frame element comparison, relation mapping, and CLI-specific annotations."
---

Use the `cli-review` skill. Run the `frame-analysis` workflow from the "Command Design Phase: 1-command-design" section.

Extract the verb list from `1-command-design/commandset-shape.md` (Section 1: Verb-Noun Decomposition Matrix) if it exists. Otherwise, extract verbs from `0-analysis/commandset.md` or ask the user for a verb list.

Generate `1-command-design/frame-analysis.md` following all six steps:
1. Extract CLI Verbs
2. FrameNet Frame Lookup (with frame IDs)
3. Frame Element Comparison (for confusable verb groups)
4. Frame Relation Mapping (with distance)
5. CLI-Specific Frame Annotations (for uncovered verbs)
6. Insights Summary (synonyms, near-synonyms, false friends, safe coexistence)

Follow the Precision Requirements: cite frame IDs, distinguish primary vs extended match, separate FrameNet observations from CLI interpretation.
