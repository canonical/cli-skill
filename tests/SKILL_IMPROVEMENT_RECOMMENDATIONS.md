# Skill Improvement Recommendations

## Summary

Resumed-session interrogation across 20 variations shows high recall on injected violations but low precision in mixed/baseline-heavy runs.

- Variations analyzed: 20
- Variations with at least one false positive: 15
- Total false positives identified by resumed analysis: 103
- Most frequent recurring false-positive categories: low-2-table-column-headers-are-not-bolded (13), low-4-non-standard-tone-of-voice-in-daemon-error-messages (13), unrated-1-non-standard-exact-time-formatting (13)

The strongest pattern is repeated baseline findings being emitted in many variations, even when the injected change is narrow or when no violation was injected.

## Individual Suggestions

1. Add a strict two-lane finding model in harness mode
- Emit `Delta Findings` only for changed commands/flags/args in the variant under review.
- Emit `Baseline Findings` separately and mark them non-scoring by default.
- In harness score calculation, include only `Delta Findings`.

2. Add explicit false-positive guardrails for the top recurring categories
- Do not emit `table-column-headers-are-not-bolded` in `Delta Findings` unless table rendering code or formatting flags changed in the variant.
- Do not emit `non-standard-tone-of-voice-in-daemon-error-messages` in `Delta Findings` unless error strings or error-formatting helpers changed in the variant.
- Do not emit `non-standard-exact-time-formatting` in `Delta Findings` unless date/time formatting paths or RFC3339 flags changed in the variant.

3. Add command-family consistency scoping
- For verb-consistency findings (for example `create/add`, `delete/remove`), only emit in `Delta Findings` when at least one command in that object family changed.
- If no family member changed, classify as baseline informational only.

4. Require concrete evidence tuples for every scored finding
- Every scored finding must include: `command_path`, `token_target`, `observed_text`, `expected_text`, `rule_clause`.
- Drop findings missing any of these fields from scoring.

5. Add deterministic deduplication before final output
- Dedup key format: `rule_clause|command_path|token_target|observed_text`.
- Keep only one scored finding per dedup key.

6. Add mutation/injection verification gate
- Before scoring, verify each planned mutation actually exists in the variant code.
- If mutation check fails, mark the case `invalid-variation` and exclude from precision/recall aggregates.

7. Add a hard output contract for benchmark comparability
- Enforce machine-readable metrics block with counts for scored and non-scored findings.
- Include both totals and per-severity breakdown for scored findings only.

## Rules To Add To skill.md

- Rule CLI-AMEND-001: In harness mode, report output MUST contain exactly two sections: `Delta Findings` and `Baseline Findings`.
- Rule CLI-AMEND-002: Harness scoring MUST use `Delta Findings` only; `Baseline Findings` MUST be excluded from precision and recall metrics.
- Rule CLI-AMEND-003: A finding MAY be placed in `Delta Findings` only if the finding maps to a changed command, changed flag, changed positional argument, or changed formatter path in the variant.
- Rule CLI-AMEND-004: The finding categories `table-column-headers-are-not-bolded`, `non-standard-tone-of-voice-in-daemon-error-messages`, and `non-standard-exact-time-formatting` MUST default to `Baseline Findings` unless their owning code path changed.
- Rule CLI-AMEND-005: Command-family verb consistency findings (for example add/create, remove/delete) MUST be scored only when at least one command in that specific family was changed in the variant.
- Rule CLI-AMEND-006: Every scored finding MUST include all evidence fields: `command_path`, `token_target`, `observed_text`, `expected_text`, `rule_clause`.
- Rule CLI-AMEND-007: Findings missing any required evidence field MUST be dropped from scored output.
- Rule CLI-AMEND-008: Reviewer MUST deduplicate scored findings using key `rule_clause|command_path|token_target|observed_text` before severity assignment.
- Rule CLI-AMEND-009: If planned mutation/injection is not present in the variant, reviewer MUST mark the variation `invalid-variation` and MUST exclude it from aggregate harness metrics.
- Rule CLI-AMEND-010: Final report MUST include metrics block with fields: `delta_findings_count`, `baseline_findings_count`, `scored_high_count`, `scored_medium_count`, `scored_low_count`, `invalid_variation`.
