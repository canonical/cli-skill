# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `cmd/todo/main.go` and `cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 3 | Command Naming, Flag Naming |
| Low | 4 | Tabular Data, Tone of Voice |
| Unrated | 1 | Logging Output |
| **Total** | **8** | |

**Overall rating:** 89.55 🟡 **Fair**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

This manual run was executed outside of a pull request context on the `todo-17` codebase. No pull request diff is present, but a full review of the current code has been conducted.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-inconsistent-verbs-are-used-to-create-secondary-objects) | Create verb consistency for secondary objects. | [Commonly used commands](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) | `add-schedule` should be renamed to `create-schedule`. |
| [MEDIUM-2](#medium-2-inconsistent-verbs-are-used-to-delete-secondary-objects) | Delete verb consistency for secondary objects. | [Commonly used commands](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) | `remove-schedule` should be renamed to `delete-schedule`. |
| [MEDIUM-3](#medium-3-plural-flag-name---sinks-for-repeatable-array-flag) | Flag names must be singular when repeatable. | [Flags accepting multiple arguments](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#flags-accepting-multiple-arguments) | `--sinks` flag should be renamed to `--sink`. |
| [LOW-1](#low-1-listing-commands-output-json-by-default) | Tabular default formatting for list commands. | [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) | `sinks` and `schedules` output JSON instead of tables. |
| [LOW-2](#low-2-table-column-headers-are-not-rendered-in-bold) | Table column headers should be bold. | [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) | Column headers printed by `todo list` are not bold. |
| [LOW-3](#low-3-table-headers-cannot-be-hidden-via-a---no-headers-flag) | Headers should support being hidden. | [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) | Missing `--no-headers` option. |
| [LOW-4](#low-4-non-succinct-error-messages) | Error messages must use "cannot" instead of "failed". | [Errors, warnings and success messages](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#errors-warnings-and-success-messages) | Daemon commands use "status request failed" and "shutdown request failed". |
| [UNRATED-1](#unrated-1-default-exact-date-and-time-formatting-is-not-iso-8601) | Exact date/time formatting must use ISO-8601. | [Timestamps](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#timestamps) | Exact date/time format is "2006-01-02 15:04 MST". |

---

## Non-compliance Findings (with citations)

### [MEDIUM-1] Inconsistent verbs are used to create secondary objects

**CLI Standard citation:** [Commonly used commands](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) — *"tool create-foo <id> — create a new instance of a secondary object, use flags to specify hierarchy"*

**Evidence:**
```go
addScheduleCmd := &cobra.Command{
	Use:   "add-schedule <schedule-id>",
	Short: "Add an immutable schedule",
```
The CLI uses the verb `create` for creating sinks (`create-sink`) but uses the verb `add` for creating schedules (`add-schedule`). It should use consistent verb structures.

**Remediation:** Rename `add-schedule` to `create-schedule` to align with the standard verb pattern.


### [MEDIUM-2] Inconsistent verbs are used to delete secondary objects

**CLI Standard citation:** [Commonly used commands](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) — *"tool delete-foo <id> — delete an instance of a secondary object"*

**Evidence:**
```go
removeScheduleCmd := &cobra.Command{
	Use:   "remove-schedule <schedule-id>",
	Short: "Remove a schedule",
```
The CLI uses the verb `delete` for deleting sinks (`delete-sink`) but uses the verb `remove` for deleting schedules (`remove-schedule`). It should use consistent verb structures.

**Remediation:** Rename `remove-schedule` to `delete-schedule` to align with the standard verb pattern.


### [MEDIUM-3] Plural flag name `--sinks` for repeatable array flag

**CLI Standard citation:** [Flags accepting multiple arguments](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#flags-accepting-multiple-arguments) — *"To enable a flag to accept several values, there are two accepted patterns for the time being: it can be repeated, in which case the flag name must be singular; otherwise, it may accept a comma-separated list of values, in which case the flag name must be plural."*

**Evidence:**
```go
addScheduleCmd.Flags().StringArray("sinks", nil, "Sink ids (repeatable)  // VIOLATION: plural flag for array")
```
The flag `--sinks` is defined as a repeatable string array, but it has a plural name.

**Remediation:** Rename the flag from `--sinks` to `--sink` to comply with the standard's singular naming for repeatable flags.


### [LOW-1] Listing commands output JSON by default

**CLI Standard citation:** [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) — *"When rendering data to the output stream for users, tables are often used to structure the information."*

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
Secondary object list commands (`sinks` and `schedules`) fall into the `default` case of `printJSONOrTable`, rendering JSON output even when table format is the default.

**Remediation:** Implement proper tabular layout printing for `sinks` and `schedules` in `printJSONOrTable`.


### [LOW-2] Table column headers are not rendered in bold

**CLI Standard citation:** [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) — *"If present, column headers should use upper case (e.g. NAME, STATUS, etc) and bold font."*

**Evidence:**
```go
fmt.Printf("%-*s%s%-*s%s%-*s%s%s\n", idWidth, "ID", sep, stateWidth, "STATE", sep, dueWidth, "DUE", sep, "TITLE")
```
Table headers printed by `printTodos` are formatted as plain text rather than bold font.

**Remediation:** Wrap column headers with formatting to render them in bold, for example by using `common.Bold`.


### [LOW-3] Table headers cannot be hidden via a `--no-headers` flag

**CLI Standard citation:** [Tabular Data](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#tabular-data) — *"Show column headers by default but allow them to be hidden with --no-headers."*

**Evidence:**
```go
fmt.Printf("%-*s%s%-*s%s%-*s%s%s\n", idWidth, "ID", sep, stateWidth, "STATE", sep, dueWidth, "DUE", sep, "TITLE")
```
There is no option or flag supporting the hiding of table headers in listing outputs.

**Remediation:** Add a `--no-headers` flag to listing commands and conditionally skip printing the column headers.


### [LOW-4] Non-succinct error messages

**CLI Standard citation:** [Errors, warnings and success messages](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#errors-warnings-and-success-messages) — *"Use “cannot” instead of “didn’t / couldn’t / wouldn’t / failed to / unable to / etc”."*

**Evidence:**
```go
return fmt.Errorf("status request failed: %s", resp.Status)
...
return fmt.Errorf("shutdown request failed: %s", resp.Status)
```
Error messages in `cmd/todod/main.go` use the word "failed" instead of "cannot".

**Remediation:** Revise the error messages to use "cannot", e.g., "cannot request status" and "cannot shut down daemon".


### [UNRATED-1] Default exact date and time formatting is not ISO-8601

**CLI Standard citation:** [Timestamps](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#timestamps) — *"For communicating exact time, use the date and time format defined in ISO 8601."*

**Evidence:**
```go
const ExactLayout = "2006-01-02 15:04 MST"
```
The exact time formatting used without the `--rfc3339` flag is not ISO-8601 compliant.

**Remediation:** Update `ExactLayout` to conform to ISO-8601 format.

---

## Compliant Findings Summary

- **Primary Object Commands**: The primary object (`todo`) utilizes verb-based commands such as `list`, `show`, `create`, `update`, `close`, `reopen`, and `reject`.
- **Secondary Object Singular/Plural Patterns**: Secondary objects adhere to standard pluralization for listing (`sinks`, `schedules`) and singular form for details (`sink`, `schedule`).
- **Clear Tone of Voice**: The CLI copy uses friendly, active, and direct sentences without contractions, avoiding chatty formatting.
- **Empty States routed to Stderr**: Empty table states (e.g. when there are no todos) are correctly printed to `stderr` with a successful exit code of 0.
- **Single/Long Flag Policy**: The CLI does not duplicate short and long options for the same action, prioritizing long flags (`--format`, `--rfc3339`, `--clear-due`).
- **Color Capability Detection**: The `todo` CLI utilizes Termenv to dynamically check background color and support capability, ensuring `NO_COLOR` environment variable is fully respected.
