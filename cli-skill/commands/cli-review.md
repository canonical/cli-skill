# Command: /cli-review

CLI standard compliance review only.

## Execution Order

1. Check for the existence of `cli-review/0-cli-discovery-preflight/`. If it does not exist, run `shared/cli-discovery-preflight.md`
2. Use preflight outputs from `cli-review/0-cli-discovery-preflight/`
3. Read `cli-skill/references/cli-standard.md` in full
4. Evaluate the CLI only against that standard

## Scope

- Check compliance with `cli-skill/references/cli-standard.md` only, do not use semantic criteria, or heuristics
- Findings must map to explicit rules from the standard
- Non-compliance and compliant evidence based on observed CLI behavior/docs/code
- Use these rules to determine severity:
  * High: violations of command structure and naming, use of positional parameters, accessibility/color violations
  * Medium: use if non-standard verbs for commands, inconsistent flag names or usage, extremely high complexity (eg. created by >20 commands)
  * Low: formatting violations, duplicate short/long flags

## Out Of Scope

- Any check not defined by `cli-skill/references/cli-standard.md`
- Heuristic UX analysis not grounded in the standard text
- General quality commentary without a standards citation

If an issue is not covered by the standard, do not include it in `/cli-review` findings.

## Required Output

Write:

- `cli-review/cli-review.md`

Required sections:

1. Summary
2. Compliance Matrix
3. Non-compliance Findings (with citations)
4. Remediation Actions (standards-mapped)
5. Compliant Findings Summary (concise, without citations)

## Summary Requirements
The summary should list the number of violations, and their severity. It shall include these in a table. It shall give an overall score (Excellent = >95%, Very Good = >90%, Good = >80%, Room for Improvement = >60%, Need for Action = <=60%) 
The score is calculated based on violations and set into relation with the size of the command set. Start with a score of 100%, the number of commands N, and the weight of a single command W=100/N. First, make a list of all the violations sorted by command. For each High violation, reduce the score by 5*W, for each Medium violation by 2*W, and for each Small violation by 0.5*W. USE THIS ALGORITHM, DO NOT REASON ABOUT IT, OR FIND ALTERNATIVE WAYS TO CALCULATE A SCORE.

## Compliance Matrix Requirements

The `Compliance Matrix` section must include a table with these columns:

- `Standard Clause`
- `Rule Summary`
- `Evidence`
- `Status` (`compliant`, `non-compliant`, or `not-assessable`)
- `Severity` (`High`, `Medium`, or `Low`)
- `Notes`

Every non-compliance finding 

1. Must include the violated clause from `cli-skill/references/cli-standard.md`
2. Must include concrete CLI evidence (command/help/code reference)
3. Should include whenever possible a remediation action that restores compliance
4. Cite the code block in markdown code format where the violation is detected

**Command completion checkpoint**: Verify that the review file exists in `cli-review/cli-review.md` and is non-empty. Do not proceed until confirmed.