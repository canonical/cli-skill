# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/todo/cmd/todo/main.go` and `/project/todo/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 1 | Command Naming |
| Low | 1 | Command Naming |
| Unrated | 0 | — |
| **Total** | **2** | |

**Overall rating:** 96.82 💚 **Excellent**

---

## CLI changes in this PR

* **Incorrect Command Verb:** The primary-object detail display command was renamed from `show` to `info <todo-id>`, introducing a medium-severity violation under the standards (requiring the `show` verb for primary-object details).
* **Missing Shorthand Pattern:** The secondary-object detail display command was renamed from `sink` to `show-sink <sink-id>`, introducing a low-severity violation under the shorthand naming standard (requiring the `sink <sink-id>` pattern).

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-todo-show-command-must-use-verb) | Commands showing primary-object details must use the standard verb form 'show'. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) | Rename the command back to `show` to conform with the standard verb rules. |
| [LOW-1](#low-1-sink-show-command-must-use-shorthand) | Displaying details of a secondary object must use the singular noun shorthand pattern instead of a verb-noun compound like `show-foobar`. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) | Rename the command back to `sink` to conform with the standard shorthand rules. |

---

## Non-compliance Findings (with citations)

### [MEDIUM-1] todo show command must use verb
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"For actions performed on primary objects of commands, they must use standard verbs... Use show instead of info, display, inspect, or describe."*

**Evidence:**
In `/project/tests/todo-10/cmd/todo/main.go`:
```go
	showCmd := &cobra.Command{
		Use:   "info <todo-id>",
		Short: "Show a todo",
```
The command is configured with the `Use` string `"info <todo-id>"`, violating the standard rule which requires the verb `"show"` for retrieving details of a primary object.

**Remediation:** Rename the `Use` field of `showCmd` to `"show <todo-id>"` in `cmd/todo/main.go` to restore compliance:
```go
	showCmd := &cobra.Command{
		Use:   "show <todo-id>",
```

### [LOW-1] sink show command must use shorthand
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"Use foobar <id> as a shorthand for showing information about a specific secondary object instance instead of show-foobar (e.g. snap service snapd over snap show-service snapd)."*

**Evidence:**
In `/project/tests/todo-10/cmd/todo/main.go`:
```go
	sinkCmd := &cobra.Command{
		Use:   "show-sink <sink-id>",
		Short: "Show a sink",
```
The command is configured with the `Use` string `"show-sink <sink-id>"`, violating the shorthand pattern rule which requires `"sink <sink-id>"` instead of `"show-sink"`.

**Remediation:** Rename the `Use` field of `sinkCmd` to `"sink <sink-id>"` in `cmd/todo/main.go` to restore compliance:
```go
	sinkCmd := &cobra.Command{
		Use:   "sink <sink-id>",
```

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** Formatting helpers and help outputs use dynamic color/styling capability detection via `github.com/muesli/termenv`. Colors/formatting are disabled when stdout/stderr is redirected or `NO_COLOR` is present.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing:** Shorthand patterns for secondary objects listing are adhered to (`sinks`, `schedules`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
