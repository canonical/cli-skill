# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/todo/cmd/todo/main.go` and `/project/todo/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 3 | Command Naming, Positional Parameters |
| Medium | 0 | — |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **3** | |

**Overall rating:** 86.36 🟡 **Fair**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by W; Medium violation by 0.5*W; Low violation by 0.2*W. Clamp to 0-100.

---

## CLI changes in this PR

* **Decreased Primary-Object Listing Clarity:** The primary-object listing command was renamed from `list` to `list-todos`, introducing a high-severity violation under the standard primary-object listing shorthand (which requires the verb `list`).
* **Non-Standard State-Display Command Naming:** The state-display command for pending reminders was renamed from `reminder-status` to `reminders`, introducing a high-severity violation under the status suffix rule (requiring specific secondary object state-display commands to use the `foobar-status` pattern).
* **Reduced Positional Argument Clarity:** The positional argument placeholder for `todo show` was modified from `<todo-id>` to `<id>`. This violates positional argument clarity and creates inconsistency with other todo commands that use `<todo-id>`.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-list-todos-command-violates-primary-object-listing-standard) | The primary-object list command must be named 'list' following standard shorthand for primary objects. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) | Rename the command back to `list` to restore compliance. |
| [HIGH-2](#high-2-reminders-command-violates-status-suffix-rule) | State-display commands for specific secondary objects must use the status suffix pattern (`foobar-status`). | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) | Rename the command to `reminder-status` to conform with the standard status suffix rules. |
| [HIGH-3](#high-3-todo-show-positional-argument-naming-violates-clarity-rules) | The positional argument placeholder for `todo show` is named `<id>` instead of `<todo-id>`, causing naming inconsistency and reducing clarity. | [Positional Parameters](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#positional-parameters) | Rename the argument placeholder from `<id>` to `<todo-id>` to match the other commands. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] list-todos command violates primary-object listing standard
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"Use foobars as a shorthand for listing information about all instances of a specific type of secondary object instead of list-foobar (e.g. snap services over snap list-services for listing services of a snap, but snap list for listing snaps)."* and *"tool list | overview of all instances of primary type"*

**Evidence:**
In `cmd/todo/main.go`:
```go
	listCmd := &cobra.Command{
		Use:   "list-todos",
		Short: "List todos",
```
The command is configured with the `Use` string `"list-todos"`, which violates the standard pattern. For primary objects, the list command must simply be named `"list"`.

**Remediation:** Change the `Use` field of `listCmd` to `"list"` in `cmd/todo/main.go` to restore compliance:
```go
	listCmd := &cobra.Command{
		Use:   "list",
```

### [HIGH-2] reminders command violates status suffix rule
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"For showing state without changing it, use the shorthand status over show-status. For specific secondary objects, use foobar-status over show-foobar-status (e.g. snapcraft release-status)"*

**Evidence:**
In `cmd/todo/main.go`:
```go
	reminderStatusCmd := &cobra.Command{
		Use:   "reminders",
		Short: "Print pending MOTD reminder messages",
```
The command is configured with the `Use` string `"reminders"`, which violates the status suffix rule. Because it is a state-display command for pending MOTD reminders, it must use the standard status suffix pattern (e.g., `"reminder-status"`).

**Remediation:** Change the `Use` field of `reminderStatusCmd` to `"reminder-status"` in `cmd/todo/main.go` to restore compliance:
```go
	reminderStatusCmd := &cobra.Command{
		Use:   "reminder-status",
```

### [HIGH-3] todo show positional argument naming violates clarity rules
**CLI Standard citation:** [Positional Parameters](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#positional-parameters) — *"Positional parameters are difficult for people to remember and use correctly, especially if they could be used interchangeably; do not use positional parameters unless the order is natural and easily memorizable"*

**Evidence:**
In `cmd/todo/main.go`:
```go
	showCmd := &cobra.Command{
		Use:   "show <id>",
		Short: "Show a todo",
```
The `show` command retrieves details for a single `todo` primary object. However, its positional argument placeholder is specified as `<id>` instead of `<todo-id>`. This creates inconsistency and ambiguity across the CLI, as other commands operating on individual todos consistently use the placeholder `<todo-id>` (e.g., `update <todo-id>`, `close <todo-id>`, `reopen <todo-id>`, `reject <todo-id>`).

**Remediation:** Change the `Use` field of `showCmd` to `"show <todo-id>"` in `cmd/todo/main.go` to restore compliance and maintain consistency across todo commands:
```go
	showCmd := &cobra.Command{
		Use:   "show <todo-id>",
```

---

## Compliant Findings Summary

- **Standard Verbs (add/remove) for Secondary Objects:** Under the updated CLI guidelines, `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** The formatting helpers (`common.Bold`, `common.ColorSection`) and help outputs (`colorizedHelp`, `rootHelp`) use `RenderInlineTags` to safely render bold and colors only when termenv detects terminal capabilities and no environment overrides (`NO_COLOR`) are set.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
