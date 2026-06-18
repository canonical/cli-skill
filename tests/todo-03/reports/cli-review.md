# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-03/cmd/todo/main.go` and `/project/tests/todo-03/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Mutually Exclusive Flags |
| Medium | 0 | — |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 95.45 💚 **Excellent**

---

## CLI changes in this PR

* **Lack of Mutually Exclusive Flags Enforcement:** The `update` command was modified to add `--due` and `--clear-due` flags without enforcing mutual exclusion or documenting their relationship, resulting in a **missing mutual exclusion between --due and --clear-due** (violates **mutually-exclusive-flags** standard rule). Users can specify both flags concurrently (e.g., `todo update 1 --due 2025-12-31 --clear-due`), which leads to conflicting and ambiguous instructions to the underlying application.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-update-command-allows-specifying-both-due-and-clear-due-without-enforcing-mutual-exclusion) | Mutually exclusive options and flags should be clearly documented and handled to avoid conflicting actions. | `--due` and `--clear-due` flags in `update` (see [Parameters, Flags and Options](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#parameters-flags-and-options)) | Allowing both conflicting options simultaneously creates user ambiguity. Rule ID: `mutually-exclusive-flags`. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] update command allows specifying both --due and --clear-due without enforcing mutual exclusion

**CLI Standard citation:** [Parameters, Flags and Options](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#parameters-flags-and-options) — *"Arguments can be added to a command line to specify the object(s) that an action will be performed on, and to provide context to, or modify the action to be performed..."*

**Evidence:**
In `/project/tests/todo-03/cmd/todo/main.go`:
```go
	updateCmd.Flags().String("due", "", "New due date (RFC3339 or human-readable)")
	updateCmd.Flags().Bool("clear-due", false, "Clear due date  // VIOLATION: no check for mutual exclusion with --due")
```
And in the command execution:
```go
			title, _ := cmd.Flags().GetString("title")
			dueInput, _ := cmd.Flags().GetString("due")
			clearDue, _ := cmd.Flags().GetBool("clear-due")
```
Both flags are defined and read without checking if they are both provided. When both are specified, the application behaves ambiguously (setting a due date while trying to clear it), constituting a missing mutual exclusion between --due and --clear-due.

**Remediation:** Enforce mutual exclusion at runtime using Cobra's built-in `MarkFlagsMutuallyExclusive` or by checking the flags in the `RunE` block:
```go
			if dueInput != "" && clearDue {
				return fmt.Errorf("cannot specify both --due and --clear-due")
			}
```
Or register them as mutually exclusive in the command's setup:
```go
	updateCmd.Flags().String("due", "", "New due date (RFC3339 or human-readable)")
	updateCmd.Flags().Bool("clear-due", false, "Clear due date")
	updateCmd.MarkFlagsMutuallyExclusive("due", "clear-due")
```

---

## Compliant Findings Summary

- **Descriptive Positional Argument:** The `show` command uses the descriptive positional argument `<todo-id>` instead of `<id>`, which clarifies what entity's ID is expected.
- **State-Display Shorthand Pattern:** The `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** Formatting helpers and help outputs use dynamic color/styling capability detection via `github.com/muesli/termenv`. Colors/formatting are disabled when stdout/stderr is redirected or `NO_COLOR` is present.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
