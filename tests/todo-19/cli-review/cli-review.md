# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `cmd/todo/main.go` and `cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 0 | — |
| Low | 1 | Parameters, Flags and Options |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 99.09 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

* **Changed `--state` flag to `--filter-state` on the `list` command:** This change decreases compliance by introducing an unnecessarily wordy and inconsistent flag naming pattern. Filter flags should reflect the property being filtered directly (e.g., `--state`) rather than carrying generic prefixes like `filter-`.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [LOW-1](#low-1---filter-state-violates-consistent-naming-pattern) | Filter flags must follow consistent naming patterns (e.g., `--state` instead of `--filter-state`). | `Parameters, Flags and Options` section in `cmd/todo/main.go` | Rename `--filter-state` back to `--state`. |

---

## Non-compliance Findings (with citations)

Each finding heading must follow this exact format: `### [SEVERITY-N] <short description>`, where `SEVERITY` is `HIGH`, `MEDIUM`, `LOW`, or `UNRATED` (uppercase) and `N` is a per-severity counter starting at 1 (HIGH-1, HIGH-2, MEDIUM-1, LOW-1, …).

### [LOW-1] `--filter-state` violates consistent naming pattern
**CLI Standard citation:** [Parameters, Flags and Options](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#parameters-flags-and-options) — *"Arguments can be added to a command line to specify the object(s) that an action will be performed on, and to provide context to, or modify the action to be performed."*

**Evidence:**
- **command_path:** `todo list`
- **rule_clause:** Parameters, Flags and Options
- **evidence:** In `cmd/todo/main.go`, the `list` command registers `--filter-state` instead of `--state`:
```go
listCmd.Flags().String("filter-state", "", "Filter state: open|closed|reopened|rejected  // VIOLATION")
```
This is an unnecessarily wordy flag name. Consistent naming patterns for filter flags should use names that directly reflect what is being filtered (e.g., `--state`, `--kind`) rather than prefixing them with `filter-`.

**Remediation:** Change the flag registration in `cmd/todo/main.go` back to `--state`:
```go
listCmd.Flags().String("state", "", "Filter state: open|closed|reopened|rejected")
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
