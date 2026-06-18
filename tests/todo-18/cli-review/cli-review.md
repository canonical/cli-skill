# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `cmd/todo/main.go` and `cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 1 | Flag Naming |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 97.73 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

* **Changed `--format` to `--output`:** The output format flag on query/display commands was changed from `--format` to `--output`. This breaks the consistency of format selection flags across the suite and departs from standard CLI conventions (violates flag consistency).

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-inconsistent-output-format-flag-naming) | All commands must consistently support output format flags. | [Tabular Data / Empty states for tables](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#empty-state-machine-readable) | Change `--output` flag registration back to `--format`. |

---

## Non-compliance Findings (with citations)

Each finding heading must follow this exact format: `### [SEVERITY-N] <short description>`, where `SEVERITY` is `HIGH`, `MEDIUM`, `LOW`, or `UNRATED` (uppercase) and `N` is a per-severity counter starting at 1 (HIGH-1, HIGH-2, MEDIUM-1, LOW-1, …).

### [MEDIUM-1] Inconsistent output format flag naming

**CLI Standard citation:** [Tabular Data / Empty states for tables](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#empty-state-machine-readable) — *"If the output is in a machine-readable format, the output should simply output the zero value... $ foo list --format=json"*

**Evidence:**
```go
func addOutputFlags(cmd *cobra.Command) {
	cmd.Flags().String("output", "table", "Output format: table|json  // VIOLATION: inconsistent flag name")
	cmd.Flags().Bool("rfc3339", false, "Show dates in RFC3339 format")
}
```
And in `statusCmd` registration:
```go
	statusCmd.Flags().String("output", "table", "Output format: table|json  // VIOLATION: inconsistent flag name")
```
The command-line interface registers the `--output` flag on query commands, but the standard expects `--format` (e.g., `--format=json`) to specify output formatting for tabular data and machine-readable formats. In addition, the internal logic in `outputFlags` still queries `format` via `cmd.Flags().GetString("format")`, which means the flag value cannot be read properly, breaking functionality.

**Remediation:** Change `--output` back to `--format` in both `addOutputFlags` and `statusCmd` flag registrations:
```go
func addOutputFlags(cmd *cobra.Command) {
	cmd.Flags().String("format", "table", "Output format: table|json")
	cmd.Flags().Bool("rfc3339", false, "Show dates in RFC3339 format")
}
```

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** Under the updated CLI guidelines, `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** The formatting helpers (`common.Bold`, `common.ColorSection`) and help outputs (`colorizedHelp`, `rootHelp`) use `RenderInlineTags` to safely render bold and colors only when termenv detects terminal capabilities and no environment overrides (`NO_COLOR`) are set.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
