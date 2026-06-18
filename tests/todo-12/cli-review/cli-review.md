# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-12/cmd/todo/main.go` and `/project/tests/todo-12/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Flags and arguments documentation |
| Medium | 0 | — |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 95.5 💚 **Excellent**

---

## CLI changes in this PR

The latest variant introduces a non-compliance issue:
* **Misleading Flag Help Documentation:** The required `--todo` flag under the `add-schedule` command is documented as `(optional)` in the CLI help text. However, it is functionally required at runtime, leading to a direct contradiction that violates clarity and option documentation standards.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-todo-flag-in-add-schedule-is-marked-optional-in-help-but-is-functionally-required) | Usage MUST distinguish required from optional arguments. | `addScheduleCmd.Flags().String("todo", "", "Todo id (optional)")` but is functionally required at runtime. | The CLI standard requires that usage text clearly indicates whether options/flags are required or optional. Documenting a required flag as optional violates this. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] --todo flag in add-schedule is marked optional in help but is functionally required
**CLI Standard citation:** [Flags and arguments documentation](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-help.md#5-flags-and-arguments-documentation) — *"Usage MUST distinguish required from optional arguments."*
**Evidence:**
In `cmd/todo/main.go`, the `--todo` flag is registered as:
```go
	addScheduleCmd.Flags().String("todo", "", "Todo id (optional)")  // VIOLATION: missing required marker
```
However, the command's `RunE` method enforces this flag as strictly required at runtime:
```go
			if strings.TrimSpace(todoIDStr) == "" {
				return fmt.Errorf("--todo is required")
			}
```
If `--todo` is omitted, the command fails and prints `--todo is required`. Since the flag is required, documenting it as `(optional)` is incorrect and misleading.
**Remediation:** Update the help description to clarify that it is required, and use Cobra's `MarkFlagRequired` method to register it as required so that help documentation reflects this.
```go
	addScheduleCmd.Flags().String("todo", "", "Todo id")
	_ = addScheduleCmd.MarkFlagRequired("todo")
```

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** Standard verbs `add` and `remove` are used appropriately for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** The formatting helpers (`common.Bold`, `common.ColorSection`) and help outputs use `RenderInlineTags` to safely render bold and colors only when termenv detects terminal capabilities and no environment overrides (`NO_COLOR`) are set.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
