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

**Overall rating:** <rating badge>
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 3*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## Compliance Matrix

| Standard Clause | Rule Summary | Evidence | Severity | Notes |
|-----------------|--------------|----------|----------|-------|
<findings>
Example:
${| Grammar + Vocabulary — Commands are verbs | Every command acting on a primary object must be a verb. | `registration-csv` is a noun phrase, not a verb. | High | Should use a verb such as `generate-registration-csv`. |}

---

## Non-compliance Findings (with citations)

Example:
${### [HIGH-1] `registration-csv` is not a verb
**Standard citation:** Grammar + Vocabulary — *"Commands are verbs. Every command that acts on a primary object of a command must be a verb."*
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
