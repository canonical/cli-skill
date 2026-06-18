# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-15/cmd/todo/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Flag Handling |
| Medium | 0 | — |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 95.45 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-mutually-exclusive-flags-should-be-documented-or-handled) | Mutually exclusive flags `--due` and `--clear-due` can be specified together without any validation or handling. | `--due` and `--clear-due` (see [Flags](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#flags)) | Provide validation code in the `RunE` method of the `update` command to make these flags mutually exclusive. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] Mutually exclusive flags should be documented or handled

**CLI Standard citation:** [Flags](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#flags) — *"Flags modify the performed action... Flags must not be dependent on ordering... as a policy, do not offer both short and long flags for the same action."* and [Parameters, Flags and Options](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#parameters-flags-and-options) — *"Arguments can be added to a command line to specify the object(s) that an action will be performed on, and to provide context to, or modify the action to be performed."*

**Evidence:**
```go
	updateCmd.Flags().String("title", "", "New title")
	updateCmd.Flags().String("due", "", "New due date (RFC3339 or human-readable)")
	updateCmd.Flags().Bool("clear-due", false, "Clear due date  // VIOLATION: no check for mutual exclusion with --due")
	addOutputFlags(updateCmd)
```
In `cmd/todo/main.go`, the `--due` flag and `--clear-due` flag are defined together under the `update` command. However, both flags can be passed at the same time, which is contradictory (you cannot set a new due date and clear the due date at the same time). There is no validation or error handling to prevent the user from passing both flags simultaneously.

**Remediation:** Add mutual exclusion validation to the command’s execution function (`RunE`) using `cmd.Flags().Changed()` to check if both are supplied, and return a clear user error if they are:
```go
			if cmd.Flags().Changed("due") && cmd.Flags().Changed("clear-due") {
				return fmt.Errorf("cannot specify both --due and --clear-due")
			}
```

---

## Compliant Findings Summary

- **Commands are verbs**: Subcommands representing actions are named using standard, intuitive verbs (e.g., `list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Shorthand for listing**: Secondary object lists use shorthand plurals like `sinks` and `schedules` in compliance with the standard.
- **Shorthand for showing details**: Singular nouns are used for secondary object detail commands (e.g., `sink <sink-id>` and `schedule <schedule-id>`).
- **Single short/long flags avoided**: Standard long flags are utilized cleanly without duplicating them with unnecessary single-character short options.
- **Color and font styling auto-detection**: ANSI color formatting is dynamically disabled when stdout is redirected or `NO_COLOR` is present in the environment.
- **Standardized tabular output**: The `list` command output strictly follows tabular specifications with left-aligned headers, a 2-space column delimiter, and no ASCII decorations.
- **Routing of empty states**: When listing produces no results, the empty state message (`No todos found`) is written to `stderr` with a successful exit code (0).
