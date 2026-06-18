# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-14/cmd/todo/main.go` and `/project/tests/todo-14/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 0 | — |
| Low | 2 | Visual hierarchy and scannability, Command description clarity |
| Unrated | 0 | — |
| **Total** | **2** | |

**Overall rating:** 98.18 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

* **Inconsistent Command Group Labels:** The command group label for the "todos" category was changed to `"Todo Commands"` in `cmd/todo/main.go`. This introduced an inconsistency because the other categories (`"Sinks:"`, `"Schedules:"`, `"Other:"`) all end with a colon suffix, whereas `"Todo Commands"` does not.
* **Noun-first Command Short Description:** The short description of the `list` command was modified to `"Todo list"` instead of being verb-led (e.g. `"List todos"`). Action commands must have verb-first descriptions for clarity and consistency across help screens.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [LOW-1](#low-1-inconsistent-command-group-labels) | Command group labels must be consistently formatted. | In `cmd/todo/main.go`, `root.AddGroup(&cobra.Group{ID: "todos", Title: "Todo Commands"})` does not use a colon suffix, violating consistency with other groups under [Feedback](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#feedback). | Standardize the group label to `"Todos:"` or add colons consistently. |
| [LOW-2](#low-2-noun-first-command-short-description) | Action command short descriptions should be verb-led. | In `cmd/todo/main.go`, `listCmd.Short` is `"Todo list"`, violating [Commands are verbs](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-commands-are-verbs). | Update the short description to `"List todos"`. |

---

## Non-compliance Findings (with citations)

### [LOW-1] Inconsistent command group labels
**CLI Standard citation:** [Feedback](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#feedback) — *"Help output MUST use clear section labels... Terminology and section naming SHOULD be consistent across help pages."*
**Evidence:**
```go
	// Add command groups for help organization
	root.AddGroup(&cobra.Group{ID: "todos", Title: "Todo Commands"})  // VIOLATION: inconsistent format
	root.AddGroup(&cobra.Group{ID: "sinks", Title: "Sinks:"})
	root.AddGroup(&cobra.Group{ID: "schedules", Title: "Schedules:"})
	root.AddGroup(&cobra.Group{ID: "other", Title: "Other:"})
```
The category title `"Todo Commands"` is title-cased and lacks a trailing colon, whereas all other sibling group labels are pluralized and end with a colon (`"Sinks:"`, `"Schedules:"`, `"Other:"`).
**Remediation:** Change `"Todo Commands"` to `"Todos:"` to match the colon-suffixed pluralized style of other group labels.

### [LOW-2] Noun-first command short description
**CLI Standard citation:** [Commands are verbs](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-commands-are-verbs) — *"Commands are verbs. Every command that acts on a primary object of a command must be a verb."* Action commands should have verb-led descriptions for clarity and consistency.
**Evidence:**
```go
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Todo list",
```
The short description `"Todo list"` uses a noun-first phrase rather than a verb-led action statement.
**Remediation:** Change `Short: "Todo list"` to `Short: "List todos"`.

---

## Compliant Findings Summary

- **Primary-Object Command Structure:** All other core commands for managing todos are verb-led (e.g. `list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Commands:** Secondary objects adhere to standard patterns like `sinks`, `sink`, `create-sink`, `delete-sink`, `schedules`, `schedule`, `add-schedule`, `remove-schedule`.
- **Shorthand Pattern:** State/info commands conform to standard shorthand patterns, e.g. `reminder-status` instead of `show-reminder-status`.
- **Proper Color Handling:** Terminal color capability detection via `termenv` and the `NO_COLOR` override is correctly and safely handled.
- **Empty States and Table Formatting:** Tables are left-aligned, use bold uppercase headers, use a 2-space delimiter, and write empty states to stderr with a 0 exit code.
