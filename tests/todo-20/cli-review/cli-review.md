# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-20/cmd/todo/main.go` and `/project/tests/todo-20/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 2 | Grammar + Vocabulary |
| Low | 1 | Tone of Voice |
| Unrated | 0 | — |
| **Total** | **3** | |

**Overall rating:** 94.78 🟢 **Good**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by W; Medium violation by 0.5*W; Low violation by 0.2*W. Clamp to 0-100.

---

## CLI changes in this PR

The latest changes introduced a few non-compliance issues:
* **Ambiguity between Help Topic and List Command:** Added a help topic for `todos` registered with `Use: "list"`, causing a direct naming collision and command-lookup conflict with the actual `list` command.
* **Non-Standard Action Verb for Closing:** Renamed the standard verb-led `close` command to `mark-closed`, violating the action verb naming consistency.
* **Inconsistent Error Message Format:** Added a title-cased and parenthesized error message in the `add-schedule` command when parsing invalid Todo IDs, departing from the consistent lowercase and direct style of the rest of the CLI.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-ambiguous-command-name-collision-between-help-topic-and-list-command) | Avoid ambiguity between help topics and commands. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) | Rename the help topic command `list` to `todos` to resolve the conflict. |
| [MEDIUM-2](#medium-2-close-command-uses-non-standard-verb-adjective-phrase) | Commands acting on primary objects must be verbs and align with their effect. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commands-are-verbs) | Rename `mark-closed` to the standard verb `close`. |
| [LOW-1](#low-1-inconsistent-error-message-formatting) | Error messages must be consistent in tone, capitalization, and formatting. | [Errors, warnings and success messages](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#errors-warnings-and-success-messages) | Change the error message to lowercase without parentheses and extra punctuation. |

---

## Non-compliance Findings (with citations)

### [MEDIUM-1] Ambiguous command name collision between help topic and list command
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"Good grammar is concise, and easy to learn and remember."* and *"Conversely, if you mix different grammar styles, or pick an uneven vocabulary, this will make it difficult to deduce functionality."*
**Evidence:**
```go
	todosTopicCmd := &cobra.Command{
		Use:   "list",  // VIOLATION: ambiguous with list command
		Short: "What todos mean in this app",
```
The help topic for `todos` is registered using the name `list`. This creates an ambiguous command name and collides with the actual `list` command.
**Remediation:** Change the help topic subcommand `Use` field from `"list"` to `"todos"`.

### [MEDIUM-2] Close command uses non-standard verb-adjective phrase
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commands-are-verbs) — *"Commands are verbs. Every command that acts on a primary object of a command must be a verb."* and [Commonly used commands](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) which lists standard lifecycle action commands.
**Evidence:**
```go
	closeCmd := todoActionCmd("mark-closed <todo-id>", "Close a todo", func(cli *client.Client, id string) (model.Todo, error) {
		return cli.CloseTodo(context.Background(), id)
	}, newClient)
```
The close command is registered as `mark-closed` instead of the standard verb `close`.
**Remediation:** Rename the `mark-closed` command to the standard verb `close`.

### [LOW-1] Inconsistent error message formatting
**CLI Standard citation:** [Errors, warnings and success messages](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#errors-warnings-and-success-messages) — *"All messages should be human-readable, and as short and succinct as possible while considering the task at hand."* and [CLI Copy and Tone of Voice](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#be-concise-precise-and-clear-at-all-times) — *"Be concise, precise, and clear at all times."*
**Evidence:**
```go
			todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
			if err != nil {
				return fmt.Errorf("Invalid Todo ID (error: %v)", err)  // VIOLATION: inconsistent format
			}
```
The error output when parsing an invalid Todo ID uses title-casing, parentheses, and punctuation, which is inconsistent with the other standard lowercase error outputs in the CLI (such as `"--todo is required"`).
**Remediation:** Rewrite the error message to be consistent with standard CLI error formatting (e.g., `invalid todo ID: %w`).

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
