# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-04/cmd/todo/main.go` and `/project/tests/todo-04/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Structure and Naming |
| Medium | 0 | — |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 95.45 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by W; Medium violation by 0.5*W; Low violation by 0.2*W. Clamp to 0-100%.

---

## CLI changes in this PR

The command for printing pending MOTD reminder messages has been renamed from `reminder-status` to `reminders` in `cmd/todo/main.go`. This change introduces a high-severity compliance violation under the Canonical CLI standard, which specifies that commands displaying state for specific secondary objects must use the status suffix (e.g., `reminder-status` or `reminders-status`), rather than just the plural object name (e.g., `reminders`).

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-state-display-command-for-reminders-must-use-the-status-suffix) | State-display commands for specific secondary objects must use the status suffix. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#showing-state-shorthand) | Rename the command to `reminder-status` to follow the standard state-display naming pattern. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] State-display command for reminders must use the status suffix
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#showing-state-shorthand) — *"For showing state without changing it, use the shorthand status over show-status. For specific secondary objects, use foobar-status over show-foobar-status (e.g. snapcraft release-status), and foobar over show-foobar."*

**Evidence:**
In `/project/tests/todo-04/cmd/todo/main.go`:
```go
	reminderStatusCmd := &cobra.Command{
		Use:   "reminders",
		Short: "Print pending MOTD reminder messages",
```
The command for printing pending MOTD reminder messages (a state-display action for specific secondary objects) is defined with the use name `"reminders"`. Under the Canonical CLI standard, specific secondary state-display commands must use the status suffix (`reminder-status` or `reminders-status`) to clearly distinguish state-display from listing secondary objects (which uses the plural form shorthand like `sinks` or `schedules`).

**Remediation:** Rename the `Use` field of `reminderStatusCmd` back to `"reminder-status"` (or `"reminders-status"`) in `cmd/todo/main.go` to conform to the state-display suffix rule:
```go
	reminderStatusCmd := &cobra.Command{
		Use:   "reminder-status",
```

---

## Compliant Findings Summary

- **Primary show command naming**: The command for viewing individual todo details uses the correct standard pattern `show <todo-id>`.
- **Secondary object listing shorthand (sinks)**: Secondary sinks use the correct plural form shorthand `sinks`.
- **Secondary object detail shorthand**: Individual secondary object details use the correct singular form shorthand `sink <sink-id>` and `schedule <schedule-id>`.
- **Status state-display commands**: System-wide status is mapped to `status`, which is the correct pattern.
- **Unified output formats**: Formatting flags and time display structures are consistently integrated via `addOutputFlags` and `outputFlags` helpers.
- **Primary-Object Command Structure**: Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Flat Secondary Mutation Hierarchy**: Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags**: Short and long flags are not duplicated for the same action.
- **Help/Version Support**: The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards**: Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling**: The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
