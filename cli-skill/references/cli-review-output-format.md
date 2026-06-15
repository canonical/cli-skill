# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of <paths for examined files>

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | <# high> | <main categories for high issues> |
| Medium | <# medium> | <main categories for medium issues> |
| Low | <# low> | <main categories for low issues> |
| Unrated | <# unrated> | <main categories for unrated issues> |
| **Total** | **<# total>** | |

**Overall rating:** <score> <rating badge>
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

<describe findings: what code changes improve or decrease the CLI>

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
<findings>
Example:
${| HIGH-1 | Every command acting on a primary object must be a verb. | `registration-csv` is a noun phrase, not a verb. | Should use a verb such as `generate-registration-csv`. |}

---

## Non-compliance Findings (with citations)

Each finding heading must follow this exact format: `### [SEVERITY-N] <short description>`, where `SEVERITY` is `HIGH`, `MEDIUM`, `LOW`, or `UNRATED` (uppercase) and `N` is a per-severity counter starting at 1 (HIGH-1, HIGH-2, MEDIUM-1, LOW-1, …).

Example:
${### [HIGH-1] `registration-csv` is not a verb
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commands-are-verbs) — *"Commands are verbs. Every command that acts on a primary object of a command must be a verb."*
**Evidence:**
```python
app.command(
    name="registration-csv",
)(_registration_csv_cmd)
```
The command name is a noun phrase. It should be a verb-noun compound (e.g. `generate-registration-csv`).
**Remediation:** Rename to `generate-registration-csv` or `export-registration-csv`.}

---

## Compliant Findings Summary

<list of short bullet points of compliant findings>
