# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-15/cmd/todo/main.go` and `/project/tests/todo-15/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Parameters, Flags and Options |
| Medium | 0 | — |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 95.45 💚 **Excellent**

---

## CLI changes in this PR

* **Missing Mutual Exclusion for Update Command:** A high-severity issue has been identified in `cmd/todo/main.go` where the `update` command accepts both `--due` and `--clear-due` without enforcing mutual exclusion or documenting their conflicting behavior in the CLI.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-missing-mutual-exclusion-between---due-and---clear-due) | Mutually exclusive flags should be documented or handled | `--due` and `--clear-due` flags on `update` command are not mutually exclusive | Add validation logic to return an error when both flags are supplied. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] Missing mutual exclusion between `--due` and `--clear-due`

**CLI Standard citation:** [Parameters, Flags and Options](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md) — *"Arguments can be added to a command line to specify the object(s) that an action will be performed on, and to provide context to, or modify the action to be performed."*

**Evidence:**
In `/project/tests/todo-15/cmd/todo/main.go`:
```go
	updateCmd.Flags().String("title", "", "New title")
	updateCmd.Flags().String("due", "", "New due date (RFC3339 or human-readable)")
	updateCmd.Flags().Bool("clear-due", false, "Clear due date  // VIOLATION: no check for mutual exclusion with --due")
```
The CLI accepts both `--due <val>` and `--clear-due` simultaneously. This is a logical conflict because they represent opposing actions for the same attribute (setting a new due date vs. clearing it).

**Remediation:**
Add a check in the `RunE` block of the `updateCmd` to verify that both are not set at the same time:
```go
if clearDue && strings.TrimSpace(dueInput) != "" {
    return fmt.Errorf("cannot use both --due and --clear-due")
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
