# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-13/cmd/todo/main.go` and `/project/tests/todo-13/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 1 | Grammar + Vocabulary |
| Low | 0 | — |
| Unrated | 2 | Tabular Data |
| **Total** | **3** | |

**Overall rating:** 97.73 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

* **Action Verb Alignment Inconsistency:** The command for closing a todo was updated to `mark-closed`. While other state-changing actions use standard single active verbs (`reopen` and `reject`), the use of `mark-closed` introduces a naming and structural inconsistency in the transition command-family.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-inconsistent-verb-mark-closed-for-todo-state-transition) | Action command uses inconsistent/non-standard verb `mark-closed` instead of the standard verb `close`. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-commands-are-verbs) | Rename `mark-closed` to `close` for verb consistency and clarity. |
| [UNRATED-1](#unrated-1-table-column-headers-are-not-rendered-in-bold) | Table column headers printed by `todo list` are not rendered in bold. | [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format) | Apply bold formatting tags to header column text before outputting. |
| [UNRATED-2](#unrated-2-table-column-headers-cannot-be-hidden-via-no-headers-flag) | Table column headers cannot be hidden via a `--no-headers` flag. | [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format) | Add `--no-headers` to table output options. |

---

## Non-compliance Findings (with citations)

### [MEDIUM-1] Inconsistent verb `mark-closed` for todo state transition
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-commands-are-verbs) — *"Commands are verbs. Every command that acts on a primary object of a command must be a verb."* and [rule-all-commands-must-converge](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-all-commands-must-converge) — *"All commands must converge on a grammar based on the above rules."*
**Evidence:**
```go
	closeCmd := todoActionCmd("mark-closed <todo-id>", "Close a todo", func(cli *client.Client, id string) (model.Todo, error) {
		return cli.CloseTodo(context.Background(), id)
	}, newClient)
```
The CLI uses the compound/verb-adjective phrase `mark-closed` to change a todo's state, while using standard single active verbs (`reopen` and `reject`) for other state transitions. This introduces command-family verb inconsistency on the primary object actions where the standard verb `close` is perfectly fitting.
**Remediation:** Rename the `mark-closed` command to the standard, consistent verb form `close`.

### [UNRATED-1] Table column headers are not rendered in bold
**CLI Standard citation:** [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format) — *"If present, column headers should use upper case (e.g. NAME, STATUS, etc) and bold font."*
**Evidence:**
```go
	// Print header with 2-space separator
	sep := "  "
	fmt.Printf("%-*s%s%-*s%s%-*s%s%s\n", idWidth, "ID", sep, stateWidth, "STATE", sep, dueWidth, "DUE", sep, "TITLE")
```
Table headers are printed in plain text without styling.
**Remediation:** Apply bold formatting tags (such as `<b>` or `common.Bold`) to header column text before outputting.

### [UNRATED-2] Table column headers cannot be hidden via `--no-headers` flag
**CLI Standard citation:** [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format) — *"Show column headers by default but allow them to be hidden with --no-headers."*
**Evidence:**
```go
func addOutputFlags(cmd *cobra.Command) {
	cmd.Flags().String("format", "table", "Output format: table|json")
	cmd.Flags().Bool("rfc3339", false, "Show dates in RFC3339 format")
}
```
No `--no-headers` flag is defined on listing or output-related commands, preventing users from hiding headers.
**Remediation:** Add the `--no-headers` flag to output flags and conditionally skip printing headers in tabular printer functions.

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The updated `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** Under the updated CLI guidelines, `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** The formatting helpers (`common.Bold`, `common.ColorSection`) and help outputs (`colorizedHelp`, `rootHelp`) use `RenderInlineTags` to safely render bold and colors only when termenv detects terminal capabilities and no environment overrides (`NO_COLOR`) are set.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
