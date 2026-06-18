# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-14`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | None |
| Medium | 2 | Grammar + Vocabulary |
| Low | 6 | Tabular Data, Tone of Voice, Help Organization, Description Clarity |
| Unrated | 1 | Logging Output |
| **Total** | **9** | |

**Overall rating:** 90.0 🟡 **Fair**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by W; Medium violation by 0.5*W; Low violation by 0.2*W. Clamp to 0-100.

---

## CLI changes in this PR

The review was executed directly on the `/project/tests/todo-14` codebase. The examined CLI supports robust features for task management, alert/webhook sinks, and schedules. Key details regarding the compliance and design of the current CLI:
- Separated client (`todo`) and background daemon (`todod`) communicating over UNIX domain sockets.
- Integrated Cobra framework for command routing and termenv for terminal formatting with NO_COLOR awareness.
- Uses repeatable singular flags consistently (e.g., `--event`, `--sink`).
- Specific opportunities for improved standards compliance are outlined below, such as verb consistency for secondary objects and default tabular formatting for lists.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-inconsistent-verbs-for-creating-secondary-objects) | Inconsistent verbs are used to create secondary objects. | Sinks use `create-sink` whereas schedules use `add-schedule` (violates [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands)). | Change `add-schedule` to `create-schedule` to keep creation verbs consistent. |
| [MEDIUM-2](#medium-2-inconsistent-verbs-for-deleting-secondary-objects) | Inconsistent verbs are used to delete secondary objects. | Sinks use `delete-sink` whereas schedules use `remove-schedule` (violates [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands)). | Change `remove-schedule` to `delete-schedule` to match standard `delete-foo` naming. |
| [LOW-1](#low-1-listing-commands-output-json-by-default-instead-of-tables) | Listing commands should output in tabular format by default. | `sinks` and `schedules` output JSON by default instead of a tabular structure when `--format table` is defaulted or specified (violates [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format)). | Implement dedicated tabular layout rendering for these list outputs. |
| [LOW-2](#low-2-table-column-headers-are-not-rendered-in-bold-font) | Table column headers must use upper case and bold font. | Column headers `ID`, `STATE`, `DUE`, `TITLE` printed by `todo list` are not rendered in bold font (violates [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format)). | Format header strings with a bold utility like `common.Bold`. |
| [LOW-3](#low-3-table-column-headers-cannot-be-hidden-via-no-headers) | Table column headers must allow being hidden with `--no-headers`. | The `todo list` command does not provide a `--no-headers` flag to hide headers (violates [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format)). | Introduce the `--no-headers` flag to output controls and omit headers when set. |
| [LOW-4](#low-4-daemon-error-messages-use-failed-instead-of-cannot) | Do not use "failed" or "failed to" in user-facing error messages. | Daemon status and stop commands use `"status request failed: %s"` and `"shutdown request failed: %s"` rather than `"cannot"` (violates [CLI Copy and Tone of Voice](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#cli-copy-and-tone-of-voice)). | Replace "failed" with "cannot" in CLI error formatting. |
| [LOW-5](#low-5-inconsistent-command-group-titles) | Command group labels in metadata must use consistent formatting. | Group title `"Todo Commands"` lacks a colon suffix, whereas other group titles like `"Sinks:"`, `"Schedules:"`, and `"Other:"` have them (violates [Feedback](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#feedback)). | Change group title `"Todo Commands"` to `"Todos:"` to match colon formatting. |
| [LOW-6](#low-6-command-short-description-is-noun-first) | Command short descriptions should be verb-first. | The short description of `todo list` is `"Todo list"`, which is noun-first rather than verb-first (violates [The Language of the Command Line](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#the-language-of-the-command-line)). | Rename the description to `"List todos"`. |
| [UNRATED-1](#unrated-1-exact-time-uses-non-iso-8601-format-by-default) | Exact time format must follow ISO 8601 (should rule). | Default exact date/time formatting without `--rfc3339` uses non-ISO-8601 format `"2006-01-02 15:04 MST"` (violates [Logging Output](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#timestamps)). | Use ISO 8601/RFC3339 formatting as the default for exact timestamps. |

---

## Non-compliance Findings (with citations)

### [MEDIUM-1] Inconsistent verbs for creating secondary objects
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) — *"tool create-foo <id> - create a new instance of a secondary object, use flags to specify hierarchy"*
**Evidence:**
In `cmd/todo/main.go`:
```go
	createSinkCmd := &cobra.Command{
		Use:   "create-sink <sink-id>",
		Short: "Create a sink",
...
	addScheduleCmd := &cobra.Command{
		Use:   "add-schedule <schedule-id>",
		Short: "Add an immutable schedule",
```
Creating a sink (secondary object) uses the verb `create` (`create-sink`), whereas creating a schedule (secondary object) uses the verb `add` (`add-schedule`). This inconsistency violates the standard recommendation to use consistent verbs across similar resource-management tasks.
**Remediation:** Standardize the command naming to either use `create-schedule` instead of `add-schedule`, or align the resource creation verbs consistently.

### [MEDIUM-2] Inconsistent verbs for deleting secondary objects
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) — *"tool delete-foo <id> - delete an instance of a secondary object"*
**Evidence:**
In `cmd/todo/main.go`:
```go
	deleteSinkCmd := &cobra.Command{
		Use:   "delete-sink <sink-id>",
		Short: "Delete a sink",
...
	removeScheduleCmd := &cobra.Command{
		Use:   "remove-schedule <schedule-id>",
		Short: "Remove a schedule",
```
Deleting a sink uses the verb `delete` (`delete-sink`), whereas deleting a schedule uses the verb `remove` (`remove-schedule`). Under the standard, `delete-foo` is the prescribed verb for deleting secondary objects.
**Remediation:** Rename `remove-schedule` to `delete-schedule` to bring it into compliance with the standard delete-foo naming convention and ensure consistency with `delete-sink`.

### [LOW-1] Listing commands output JSON by default instead of tables
**CLI Standard citation:** [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format) — *"Format of tables: All tables should follow a standard format..."*
**Evidence:**
In `cmd/todo/main.go`:
```go
func printJSONOrTable(v any, format string, rfc3339 bool) error {
	if format == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(v)
	}
	switch t := v.(type) {
	case model.Todo:
		fmt.Printf("id: %d\n", t.ID)
		fmt.Printf("title: %s\n", t.Title)
		fmt.Printf("state: %s\n", t.State)
		fmt.Printf("due: %s\n", common.FormatTime(t.DueAt, rfc3339))
		fmt.Printf("created: %s\n", common.FormatTime(&t.CreatedAt, rfc3339))
		fmt.Printf("updated: %s\n", common.FormatTime(&t.UpdatedAt, rfc3339))
		return nil
	default:
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(v)
	}
}
```
The commands `sinks` and `schedules` retrieve slices of `model.Sink` and `model.Schedule` respectively. Because these types fall through to the `default` case in `printJSONOrTable`, they are outputted as JSON even when `--format table` is specified or left as default.
**Remediation:** Implement tabular output formatting logic in `printJSONOrTable` for slices of `model.Sink` and `model.Schedule` to ensure they render in a proper table by default.

### [LOW-2] Table column headers are not rendered in bold font
**CLI Standard citation:** [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format) — *"If present, column headers should use upper case (e.g. NAME, STATUS, etc) and bold font."*
**Evidence:**
In `cmd/todo/main.go` (in function `printTodos`):
```go
	// Print header with 2-space separator
	sep := "  "
	fmt.Printf("%-*s%s%-*s%s%-*s%s%s\n", idWidth, "ID", sep, stateWidth, "STATE", sep, dueWidth, "DUE", sep, "TITLE")
```
The column headers `ID`, `STATE`, `DUE`, and `TITLE` are printed as plain text without any bold escape sequence or tag.
**Remediation:** Format the header cells using a bold utility such as `common.Bold` or wrap the headers with `<b>` and `</b>` tags prior to rendering.

### [LOW-3] Table column headers cannot be hidden via `--no-headers`
**CLI Standard citation:** [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-table-format) — *"Show column headers by default but allow them to be hidden with --no-headers."*
**Evidence:**
In `cmd/todo/main.go` (where `listCmd` is registered):
```go
	listCmd.Flags().String("state", "", "Filter state: open|closed|reopened|rejected")
	addOutputFlags(listCmd)
```
The CLI table commands (such as `todo list`) do not define or support a `--no-headers` flag.
**Remediation:** Add a `--no-headers` boolean flag globally or to tabular list commands and respect it in the table rendering code.

### [LOW-4] Daemon error messages use "failed" instead of "cannot"
**CLI Standard citation:** [CLI Copy and Tone of Voice](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#cli-copy-and-tone-of-voice) — *"Use “cannot” instead of “didn’t / couldn’t / wouldn’t / failed to / unable to / etc”."*
**Evidence:**
In `cmd/todod/main.go`:
```go
			if resp.StatusCode >= 300 {
				return fmt.Errorf("status request failed: %s", resp.Status)
			}
...
			if resp.StatusCode >= 300 {
				return fmt.Errorf("shutdown request failed: %s", resp.Status)
			}
```
The error messages returned by `todod status` and `todod stop` use the phrase "request failed" instead of "cannot".
**Remediation:** Rewrite these error messages to use "cannot", for example: `cannot retrieve status: %s` and `cannot shutdown daemon: %s`.

### [LOW-5] Inconsistent command group titles
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-commands-are-logically-grouped) — *"Commands are logically grouped. Not all commands are equal... We differentiate between these domains by grouping commands..."* / General help formatting consistency.
**Evidence:**
In `cmd/todo/main.go`:
```go
	// Add command groups for help organization
	root.AddGroup(&cobra.Group{ID: "todos", Title: "Todo Commands"})  // VIOLATION: inconsistent format
	root.AddGroup(&cobra.Group{ID: "sinks", Title: "Sinks:"})
	root.AddGroup(&cobra.Group{ID: "schedules", Title: "Schedules:"})
	root.AddGroup(&cobra.Group{ID: "other", Title: "Other:"})
```
The group title for `"todos"` is `"Todo Commands"` (no colon suffix), whereas `"sinks"`, `"schedules"`, and `"other"` all use colons in their titles (`"Sinks:"`, `"Schedules:"`, `"Other:"`).
**Remediation:** Update the `"todos"` group title to be `"Todos:"` to match the colon-terminated naming pattern of the other groups.

### [LOW-6] Command short description is noun-first
**CLI Standard citation:** [The Language of the Command Line](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#the-language-of-the-command-line) — *"Your user experience can be greatly improved by a grammar that is adequate to the scope of the product (use the verb paradigm for simple cases...)"*
**Evidence:**
In `cmd/todo/main.go`:
```go
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "Todo list",
```
The short description for the `todo list` command is `"Todo list"`, which is noun-first/a noun phrase rather than a verb-first action.
**Remediation:** Change the command description to a verb-first action phrase such as `"List todos"`.

### [UNRATED-1] Exact time uses non-ISO-8601 format by default
**CLI Standard citation:** [Logging Output](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#timestamps) — *"For communicating exact time, use the date and time format defined in ISO 8601."*
**Evidence:**
In `internal/common/timeutil.go`:
```go
func FormatTime(t *time.Time, useRFC3339 bool) string {
	if t == nil {
		return ""
	}
	if useRFC3339 {
		return t.Format(time.RFC3339)
	}
	return t.Local().Format("2006-01-02 15:04 MST")
}
```
Unless the `--rfc3339` flag is explicitly set, exact times are formatted as `"2006-01-02 15:04 MST"` which does not follow ISO 8601 format. (Since the standard states "you should use / for communicating... use...", and this applies to default exact times, this is an unrated finding reflecting standard preferences).
**Remediation:** Default to rendering exact times in ISO 8601/RFC3339 format.

---

## Compliant Findings Summary

- **Verbs represent actions**: All subcommands that act on primary or secondary objects are named using standard verbs (e.g., `list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Logical command grouping**: Commands are organized into intuitive categories/groups (`Todos`, `Sinks`, `Schedules`, `Other`) in help outputs.
- **Single short/long flags avoided**: The tool does not duplicate flags with both short and long options for the same action, prioritizing clean and descriptive long flags.
- **Color auto-detection and NO_COLOR support**: ANSI escape styling dynamically disables itself when the standard output is redirected or the `NO_COLOR` environment variable is defined.
- **Concise messages and direct tone of voice**: System outputs and status information avoid chattiness and explain configuration and hints concisely.
- **Tabular spacing rules**: Spacing between printed columns conforms to the two-space delimiter requirement.
- **Empty state stderr output**: The `list` command successfully prints its empty state message (`No todos found`) to standard error rather than standard output while retaining a successful exit code (0).
