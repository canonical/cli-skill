# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-07/cmd/todo/main.go` and `/project/tests/todo-07/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 1 | Flags and Options |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 97.73 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by W; Medium violation by 0.5*W; Low violation by 0.2*W. Clamp to 0-100.

---

## CLI changes in this PR

* **Decreased Flag Consistency:** The `--format` flag was renamed to `--output` across various query and action commands. This creates inconsistency in output option vocabulary and breaks formatting functionality because helper functions like `outputFlags` still expect `--format` to be registered.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-inconsistent-flag-name-output-used-instead-of-format) | All query commands must consistently support standard output format options (`--format` or `-f`). | [Empty states for tables](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) | Rename the `--output` flag to `--format` in registration helper `addOutputFlags` and `statusCmd`. |

---

## Non-compliance Findings (with citations)

### [MEDIUM-1] Inconsistent flag name output used instead of format

**CLI Standard citation:** [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) — *"If the output is in a machine-readable format, the output should simply output the zero value. ... $ foo list --format=json"*

**Evidence:**
```go
// addOutputFlags registers --format and --rfc3339 on a command.
func addOutputFlags(cmd *cobra.Command) {
	cmd.Flags().String("output", "table", "Output format: table|json  // VIOLATION: inconsistent flag name")
	cmd.Flags().Bool("rfc3339", false, "Show dates in RFC3339 format")
}
```
And on `statusCmd` in `cmd/todo/main.go`:
```go
	statusCmd.Flags().String("output", "table", "Output format: table|json  // VIOLATION: inconsistent flag name")
```
While the query logic in `outputFlags` expects the `--format` flag:
```go
func outputFlags(cmd *cobra.Command) (format string, rfc3339 bool) {
	format, _ = cmd.Flags().GetString("format")
	rfc3339, _ = cmd.Flags().GetBool("rfc3339")
```

The `--output` flag was registered on query commands instead of `--format`, which deviates from standard CLI naming conventions. Furthermore, this naming discrepancy causes a runtime issue as the flag getter (`outputFlags`) attempts to retrieve `format` which is not registered on those commands.

**Remediation:** In `cmd/todo/main.go`, change `--output` to `--format` in the flag registration functions:
```go
func addOutputFlags(cmd *cobra.Command) {
	cmd.Flags().String("format", "table", "Output format: table|json")
	cmd.Flags().Bool("rfc3339", false, "Show dates in RFC3339 format")
}
```
And on `statusCmd`:
```go
	statusCmd.Flags().String("format", "table", "Output format: table|json")
```

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The updated `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** Under the updated CLI guidelines, `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** The formatting helpers (`common.Bold`, `common.ColorSection`) and help outputs (`colorizedHelp`, `rootHelp`) use `RenderInlineTags` to safely render bold and colors only when termenv detects terminal capabilities and no environment overrides (`NO_COLOR`) are set.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
