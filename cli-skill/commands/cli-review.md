# Command: /cli-review

CLI standard compliance review only.

## Execution Order

1. Check for the existence of `cli-review/0-cli-discovery-preflight/`. If it does not exist, run `cli-skill/shared/cli-discovery-preflight.md`
2. Use preflight outputs from `cli-review/0-cli-discovery-preflight/`
3. Checkpoint: Before going to Phase 1, make sure that preflight analysis is complete. Do not proceed before it is done.
4. Read `cli-skill/references/cli-standard.md` in full
5. **Phase 1 — Collect all findings (no severity yet).** Walk every rule in the standard. For each rule, check all CLI commands and flags. Record every violation as a plain list entry: `[problem description] [evidence] [reference to code] [reference to cli standard]`. Do not assign severity in this phase. Do not stop early. Complete the full standard before moving on.
6. Checkpoint: Before going to Phase 2, make sure that no duplicate findings are listed, DO THIS ONLY by analysing the `reference to code`. 
7. **Phase 2 — Assign severity.** For each finding collected in Phase 1, assign exactly one severity (`High`, `Medium`, `Low`, or `Unrated`) using the rules in the `## Scope` section. Do not add or remove findings in this phase.
8. Build the score JSON `{"commands": <int>, "issues": [...]}` from the complete, severity-annotated list.
9. Resolve the scoring script path in this exact order and use the first existing path:
  - `.cli-skill-infra/scripts/calculate_cli_score.py`
  - `scripts/calculate_cli_score.py`
  Execute: `python3 <resolved_script_path> <json_file>`.
10. Use the script JSON output values (`score`, `rating`, `rating_badge`, counts) in the summary. **Do not compute score manually.**

### Hard Constraints

- Forbidden: manual arithmetic, estimating score, discussing alternative formulas, or reconciling formula discrepancies.
- Forbidden: “let me think / let me recalculate / let me verify” style narration in final output.
- If scoring script execution fails, stop and report failure; do not produce a scored summary.

## Scope

- Check compliance with `cli-skill/references/cli-standard.md` only, do not use semantic criteria, or heuristics
- Findings must map to clauses from the standard, suggestions in the standard will result in unrated findings
- Non-compliance and compliant evidence based on observed CLI behavior/docs/code
- Use these rules to determine severity. Severity can only be [High|Medium|Low|Unrated]:
  * High: violations of command structure and naming (SHOULD RULES DO NOT COUNT, THEY ARE UNRATED but should be included), use of positional parameters, accessibility/color violations (IF NO COLOR IS USED, eg. only plain or boldbold, NO_COLOR need not be detected)
  * Medium: use if inconsistent verbs for commands (e.g. add-foo vs. kill-foo), inconsistent flag names or usage, extremely high complexity (eg. created by >20 commands)
  * Low: formatting violations, duplicate short/long flags
  * Unrated: use if no standard rule applies, or if the finding is not certain, there is a SHOULD rule, the rule applies to a debug/secret command, use of non-standard verbs when there is a standard verb that fits the action
- All SHOULD rules must be unrated violations - report them, but don't include them in the scoring. This can be a use of non-standard verbs, or in cases where the standard says "you should", or "prefer".

## Out Of Scope

- Any check not defined by `cli-skill/references/cli-standard.md`
- Heuristic UX analysis not grounded in the standard text
- General quality commentary without a standards citation

If an issue is not covered by the standard, do not include it in `/cli-review` findings.

## Required Output

Read `cli-skill/references/cli-review-output-format.md` in full. Follow the structure and formatting strictly. Generate the text for all <placeholders> and ${examples}.

The output **must** use exactly these top-level headings in this order:

```
# Canonical CLI automated review report
## Summary          (violation table + rating badge returned by the scoring script)
## CLI changes in this PR (include only if the workflow was triggered from a PR creation/update)
## Compliance matrix  (table with columns: Finding | Rule Summary | Evidence | Notes)
## Non-compliance Findings (with citations)
## Compliant Findings Summary (concise, without citations)
```

Write:
- `cli-review/cli-review.md`

### Summary Requirements
The summary should list the number of violations, and their severity. It shall include these in a table. It shall give an overall score. This score is calculated by a script, DO NOT REASON ABOUT IT, AND DO NOT CHANGE THE ALGORITHM. 

To calculate the score:
1. Create a JSON table with structure: `{"commands": <int>, "issues": [{"severity": "High|Medium|Low|Unrated", "category": <str>, "message": <str>}, ...]}`
2. Execute: `python3 <resolved_script_path> <json_file>` using the path resolution defined in `Execution Order`
3. This script returns JSON with `score` (0-100), `passed` (boolean), `rating badge`, and severity counts
4. Use this `score` and `rating badge` in the summary section of the markdown output to report the compliance score

The script implements the standard algorithm: Start with 100%, weight W=100/#commands. For each High violation, reduce by W; Medium violation by 0.5*W; Low violation by 0.2*W. Clamp to 0-100%.

### CLI Change Requirements
Analyze the files that have been changes as part of this PR. Create a detailed summary of how each change affects the compliance of the CLI.

### Compliance Matrix Requirements

The `Compliance Matrix` section must include a table with these columns:

- `Finding` - In the form of `[SEVERITY-N]`, link to non-compliance finding in the `Non-compliance findings` section
- `Rule Summary`
- `Evidence`, include the name of the relevant section of the cli standard by using `reference to cli standard`
- `Notes`

### Non-compliance findings requirements
Every non-compliance finding must use this exact heading format:

```
### [SEVERITY-N] <short description>
```

Where `SEVERITY` is one of `HIGH`, `MEDIUM`, `LOW`, or `UNRATED` (uppercase), and `N` is a counter that resets to 1 for each severity group (e.g. HIGH-1, HIGH-2, MEDIUM-1, LOW-1). Number findings within each group in the order they appear in the Compliance Matrix.

Each finding:

1. Must include the violated clause from `cli-skill/references/cli-standard.md`
2. Must include links to the relevant sections of the CLI standard by using `reference to cli standard` and `https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md` as a base URL
3. Must include concrete CLI evidence (command/help/code reference)
4. Should include whenever possible a remediation action that restores compliance
5. Cite the code block in markdown code format where the violation is detected

**Non-compliance findings checkpoint**: Verify that each non-compliance finding is linked in the compliance matrix, and that each row in the compliance matrix links to a finding. Do not proceed until confirmed.

**Command completion checkpoint**: Verify that the review file exists in `cli-review/cli-review.md` and is non-empty. Do not proceed until confirmed.