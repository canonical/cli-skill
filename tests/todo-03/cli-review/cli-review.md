# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-03/cmd/todo/main.go` and `/project/tests/todo-03/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 0 | — |
| Low | 1 | Grammar + Vocabulary |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 99.09 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

The command for listing configured schedules was renamed from `schedules` to `list-schedules` in `cmd/todo/main.go`. This change introduces a low-severity compliance violation under the Canonical CLI standard, which specifies that the plural form of the secondary object name (e.g., `schedules`) should be used as a shorthand for listing information about those objects, rather than using `list-foobar` (e.g., `list-schedules`).

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [LOW-1](#low-1-schedule-listing-command-must-use-shorthand-not-list-schedules) | Listing command for a secondary object type must use the plural shorthand rather than the `list-foobar` verb-noun form. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) | Rename the 'list-schedules' command back to 'schedules' to follow the standard shorthand pattern. |

---

## Non-compliance Findings (with citations)

### [LOW-1] Schedule listing command must use shorthand (not 'list-schedules')
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"Use foobars as a shorthand for listing information about all instances of a specific type of secondary object instead of list-foobar (e.g. snap services over snap list-services for listing services of a snap, but snap list for listing snaps)."*

**Evidence:**
In `/project/tests/todo-03/cmd/todo/main.go`:
```go
	schedulesCmd := &cobra.Command{
		Use:   "list-schedules",
		Short: "List schedules",
```
The command for listing schedules (a secondary object type) is defined with the use name `"list-schedules"`. Following the standard grammar of Canonical CLI, listing secondary objects must use the plural shorthand `"schedules"` instead of a verb-noun compound like `list-schedules`.

**Remediation:** Rename the `Use` field of `schedulesCmd` to `"schedules"` in `cmd/todo/main.go` to conform to the secondary-object shorthand naming pattern:
```go
	schedulesCmd := &cobra.Command{
		Use:   "schedules",
```

---

## Compliant Findings Summary

- **Primary show command naming**: The command for viewing individual todo details uses the correct standard pattern `show <todo-id>`.
- **Secondary object listing shorthand (sinks)**: Secondary sinks use the correct plural form shorthand `sinks`.
- **Secondary object detail shorthand**: Individual secondary object details use the correct singular form shorthand `sink <sink-id>` and `schedule <schedule-id>`.
- **Status state-display commands**: System-wide status is mapped to `status` and `reminder-status` which is the correct pattern.
- **Unified output formats**: Formatting flags and time display structures are consistently integrated via `addOutputFlags` and `outputFlags` helpers.
