# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-04/cmd/todo/main.go` and `/project/tests/todo-04/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Command Naming |
| Medium | 2 | Grammar + Vocabulary |
| Low | 4 | Tabular Data, Tone of Voice |
| Unrated | 1 | Logging Output |
| **Total** | **8** | |

**Overall rating:** 87.27 🟡 **Fair**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by W; Medium violation by 0.5*W; Low violation by 0.2*W. Clamp to 0-100%.

---

## CLI changes in this PR

The state-display command for pending reminders was renamed from `reminder-status` to `reminders` in `cmd/todo/main.go`. This change introduces a high-severity compliance violation under the Canonical CLI standard, which specifies that state-display commands for specific secondary objects must use the status suffix (e.g., `reminder-status` instead of `reminders`).

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-state-display-command-for-pending-reminders-is-named-reminders-instead-of-reminder-status) | State-display commands for specific secondary objects must use the status suffix. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) | Rename `reminders` back to `reminder-status` to comply with the standard suffix rule. |
| [MEDIUM-1](#medium-1-inconsistent-verbs-for-secondary-object-creation) | Inconsistent verbs are used to create secondary objects. Sinks use `create-sink` whereas schedules use `add-schedule`. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) | Rename `add-schedule` to `create-schedule` for consistency. |
| [MEDIUM-2](#medium-2-inconsistent-verbs-for-secondary-object-deletion) | Inconsistent verbs are used to delete secondary objects. Sinks use `delete-sink` whereas schedules use `remove-schedule`. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) | Rename `remove-schedule` to `delete-schedule` for consistency. |
| [LOW-1](#low-1-default-formatting-for-listing-commands-sinks-and-schedules-does-not-support-table-display) | Listing commands (`sinks`, `schedules`) output JSON formatting by default instead of a tabular structure. | [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) | Implement dedicated table print functions for these models. |
| [LOW-2](#low-2-table-column-headers-are-not-bolded) | Table column headers printed by `todo list` are not rendered in bold. | [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) | Style column labels with `common.Bold` or ANSI tags. |
| [LOW-3](#low-3-table-column-headers-cannot-be-hidden-via-no-headers-flag) | Command options do not support the `--no-headers` flag. | [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) | Add `--no-headers` to table output options. |
| [LOW-4](#low-4-non-standard-tone-of-voice-in-daemon-error-messages) | Daemon status and stop commands use `"failed"` error messages instead of `"cannot"`. | [CLI Copy and Tone of Voice](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#cli-copy-and-tone-of-voice) | Change error formatting to use standard `"cannot"` phrase. |
| [UNRATED-1](#unrated-1-non-standard-exact-time-formatting) | Exact time is formatted using a custom Go layout rather than ISO 8601 when `--rfc3339` is absent. | [Logging Output: Timestamps](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#timestamps) | Use RFC3339 format by default for exact timestamps. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] State-display command for pending reminders is named 'reminders' instead of 'reminder-status'
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"For showing state without changing it, use the shorthand status over show-status. For specific secondary objects, use foobar-status over show-foobar-status (e.g. snapcraft release-status), and foobar over show-foobar."*
**Evidence:**
```go
	reminderStatusCmd := &cobra.Command{
		Use:   "reminders",
		Short: "Print pending MOTD reminder messages",
```
The state-display command for pending reminders is named `reminders` instead of `reminder-status`, violating the status suffix rule for specific secondary objects.
**Remediation:** Rename the `reminders` command back to `reminder-status` to comply with the standard status suffix rule.

### [MEDIUM-1] Inconsistent verbs for secondary object creation
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) — *"tool create-foo <id> \| create a new instance of a secondary object"*
**Evidence:**
```go
	createSinkCmd := &cobra.Command{
		Use:   "create-sink <sink-id>",
...
	addScheduleCmd := &cobra.Command{
		Use:   "add-schedule <schedule-id>",
```
The client CLI uses inconsistent verbs for adding secondary objects. `sinks` are added with `create-sink`, while `schedules` are added with `add-schedule`.
**Remediation:** Rename the `add-schedule` command to `create-schedule` to align with the standard `create-foo` convention and match `create-sink`.

### [MEDIUM-2] Inconsistent verbs for secondary object deletion
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) — *"tool delete-foo <id> \| delete an instance of a secondary object"*
**Evidence:**
```go
	deleteSinkCmd := &cobra.Command{
		Use:   "delete-sink <sink-id>",
...
	removeScheduleCmd := &cobra.Command{
		Use:   "remove-schedule <schedule-id>",
```
The client CLI uses inconsistent verbs for deleting secondary objects. `sinks` are deleted with `delete-sink`, while `schedules` are removed with `remove-schedule`.
**Remediation:** Rename the `remove-schedule` command to `delete-schedule` to align with the standard `delete-foo` convention and match `delete-sink`.

### [LOW-1] Default formatting for listing commands sinks and schedules does not support table display
**CLI Standard citation:** [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) — *"Management of objects (machines, instances, packages, …) will often require processing of relational data. When rendering data to the output stream for users, tables are often used to structure the information."*
**Evidence:**
```go
func printJSONOrTable(v any, format string, rfc3339 bool) error {
	if format == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(v)
	}
	switch t := v.(type) {
	case model.Todo:
...
	default:
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(v)
	}
}
```
When listing `sinks` or `schedules`, the output format defaults to `"table"`, but because these arrays are not instances of `model.Todo`, they fall through to the JSON printer and output JSON by default rather than a tabular format.
**Remediation:** Add switch-case support in `printJSONOrTable` for `[]model.Sink` and `[]model.Schedule` and define formatting functions to display them as standard tables.

### [LOW-2] Table column headers are not bolded
**CLI Standard citation:** [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) — *"If present, column headers should use upper case (e.g. NAME, STATUS, etc) and bold font."*
**Evidence:**
```go
	// Print header with 2-space separator
	sep := "  "
	fmt.Printf("%-*s%s%-*s%s%-*s%s%s\n", idWidth, "ID", sep, stateWidth, "STATE", sep, dueWidth, "DUE", sep, "TITLE")
```
Table headers are printed in plain text without styling.
**Remediation:** Apply bold formatting tags (such as `<b>` or `common.Bold`) to header column text before outputting.

### [LOW-3] Table column headers cannot be hidden via `--no-headers` flag
**CLI Standard citation:** [Feedback: Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) — *"Show column headers by default but allow them to be hidden with --no-headers."*
**Evidence:**
```go
func addOutputFlags(cmd *cobra.Command) {
	cmd.Flags().String("format", "table", "Output format: table|json")
	cmd.Flags().Bool("rfc3339", false, "Show dates in RFC3339 format")
}
```
No `--no-headers` flag is defined on any list or output-related commands, preventing users from hiding headers.
**Remediation:** Add the `--no-headers` flag to output flags and conditionally skip printing headers in tabular printer functions.

### [LOW-4] Non-standard Tone of Voice in daemon error messages
**CLI Standard citation:** [CLI Copy and Tone of Voice](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#cli-copy-and-tone-of-voice) — *"Use 'cannot' instead of 'didnt / couldnt / wouldnt / failed to / unable to / etc'."*
**Evidence:**
```go
			if resp.StatusCode >= 300 {
				return fmt.Errorf("status request failed: %s", resp.Status)
			}
...
			if resp.StatusCode >= 300 {
				return fmt.Errorf("shutdown request failed: %s", resp.Status)
			}
```
Error messages use `"failed"` instead of `"cannot"`.
**Remediation:** Change error messages to use `"cannot request status"` and `"cannot shutdown"`.

### [UNRATED-1] Non-standard exact time formatting
**CLI Standard citation:** [Logging Output: Timestamps](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#timestamps) — *"For communicating exact time, use the date and time format defined in ISO 8601."*
**Evidence:**
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
The default exact time format (when `useRFC3339` is false) is `"2006-01-02 15:04 MST"`, which does not conform to ISO 8601.
**Remediation:** Use RFC3339 layout as the default format for exact timestamps, or fallback to it instead of custom layouts.

---

## Compliant Findings Summary

- **Verbs represent actions**: All subcommands that act on primary or secondary objects are named using standard verbs (e.g., `list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Logical command grouping**: Commands are organized into intuitive categories/groups (`Todos`, `Sinks`, `Schedules`, `Other`) in help outputs.
- **Single short/long flags avoided**: The tool does not duplicate flags with both short and long options for the same action, prioritizing clean and descriptive long flags.
- **Color auto-detection and NO_COLOR support**: ANSI escape styling dynamically disables itself when the standard output is redirected or the `NO_COLOR` environment variable is defined.
- **Concise messages and direct tone of voice**: System outputs and status information avoid chattiness and explain configuration and hints concisely.
- **Tabular spacing rules**: Spacing between printed columns conforms to the two-space delimiter requirement.
- **Empty state stderr output**: The `list` command successfully prints its empty state message (`No todos found`) to standard error rather than standard output while retaining a successful exit code (0).
