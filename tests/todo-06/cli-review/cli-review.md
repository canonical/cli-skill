# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-06/cmd/todo/main.go` and `/project/tests/todo-06/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 1 | Flags and Options |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 97.73 💚 **Excellent**

---

## CLI changes in this PR

The latest changes introduce a non-compliance issue:
* **Plural Flag Name for Repeatable Array Flag:** The repeatable array flag on `add-schedule` has been renamed from `--sink` to `--sinks`. Under the CLI standard, repeated flags must be singular. This change introduces a new Medium-severity violation.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-plural-flag-name-sinks-for-repeatable-flag) | Flag names must be singular when accepting multiple values via repetition. | `--sinks` flag in `cmd/todo/main.go` violates the repeatable flag naming rule. | Rename `--sinks` to `--sink` to restore compliance. |

---

## Non-compliance Findings (with citations)

### [MEDIUM-1] Plural flag name `--sinks` for repeatable flag

**CLI Standard citation:** [Parameters, Flags and Options](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#flags-accepting-multiple-arguments) — *"To enable a flag to accept several values, there are two accepted patterns for the time being: it can be repeated, in which case the flag name must be singular..."*

**Evidence:**
```go
	addScheduleCmd.Flags().StringArray("sinks", nil, "Sink ids (repeatable)  // VIOLATION: plural flag for array")
```

The repeatable flag on the `add-schedule` command is named `--sinks`, which is plural. Repeated flags must be named in their singular form.

**Remediation:**
Modify `cmd/todo/main.go` to rename `--sinks` to `--sink`:
```go
	addScheduleCmd.Flags().StringArray("sink", nil, "Sink ids (repeatable)")
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
