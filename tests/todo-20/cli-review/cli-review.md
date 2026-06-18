# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/todo/cmd/todo/main.go` and `/project/todo/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 2 | Grammar + Vocabulary, Command Structure |
| Low | 1 | Error Formatting |
| Unrated | 0 | — |
| **Total** | **3** | |

**Overall rating:** 94.55% 🟢 **Good**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

The latest changes introduced several regression issues affecting CLI usability, command clarity, and error output consistency:
* **Command Ambiguity:** A help topic command (`todosTopicCmd`) was added with the same trigger `list` as the actual `list` command, causing ambiguity and overriding/shadowing command behavior.
* **Non-standard Verb for Closure:** The standard `close` action on the primary `todo` object has been changed to `mark-closed`, violating the direct-verb paradigm of the standard and adding unnecessary verbosity.
* **Inconsistent Error Formatting:** A newly added error check for parsing `todo` ID uses capitalized terms and redundant error info, which conflicts with standard lowercase and unpunctuated error outputs in the codebase.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [MEDIUM-1](#medium-1-ambiguous-list-help-topic-creates-command-conflict) | Avoid ambiguity between help topics and commands. | `list` triggers both a help topic and the list command. | Help topics should have unique non-conflicting names (e.g. `todos`). |
| [MEDIUM-2](#medium-2-non-standard-mark-closed-verb-for-closure) | Commands acting on primary objects must use concise standard verbs. | `mark-closed` is used instead of standard `close` verb. | Rename command to `close` to restore compliance. |
| [LOW-1](#low-1-inconsistent-error-message-formatting) | Error messages must be lowercase, short, and consistent. | `"Invalid Todo ID (error: %v)"` is capitalized and has redundant format. | Change to a standard lowercase error message like `"invalid todo id: %w"`. |

---

## Non-compliance Findings (with citations)

### [MEDIUM-1] Ambiguous `list` help topic creates command conflict

**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md) — *"Good grammar is concise, and easy to learn and remember. ... conversely, if you mix different grammar styles, or pick an uneven vocabulary, this will make it difficult to deduce functionality ('what will this command do?' can be easily triggered by using vastly different verbs, or by mixing different objects without allowing the user to differentiate between them)."*

**Evidence:**
```go
todosTopicCmd := &cobra.Command{
	Use:   "list",  // VIOLATION: ambiguous with list command
	Short: "What todos mean in this app",
```
And:
```go
listCmd := &cobra.Command{
	Use:   "list",
	Short: "List todos",
```
Both commands are registered under `root` in `cmd/todo/main.go`. Specifying `list` as the trigger for an informational topic command directly conflicts with the primary `list` command.

**Remediation:**
Change the `Use` string of `todosTopicCmd` to a non-conflicting term or concept name (e.g., `"todos-topic"` or `"todos"`) or register it correctly under Cobra's custom help topics mechanism.

---

### [MEDIUM-2] Non-standard `mark-closed` verb for closure

**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commands-are-verbs) — *"Commands are verbs. Every command that acts on a primary object of a command must be a verb."* and *"At Canonical, most of the commands we build are complex... With this standard Canonical command grammar, we strive to create a precise, minimal interface."*

**Evidence:**
```go
closeCmd := todoActionCmd("mark-closed <todo-id>", "Close a todo", func(cli *client.Client, id string) (model.Todo, error) {
	return cli.CloseTodo(context.Background(), id)
}, newClient)
```
The command name `mark-closed` is a wordy verb-noun compound that acts on the primary object of the tool. It replaces the direct, standard, and concise verb `close`.

**Remediation:**
Rename the action command back to `close` (e.g., `close <todo-id>`), which is a clear, standard, and concise verb representing the closure state transition.

---

### [LOW-1] Inconsistent error message formatting

**CLI Standard citation:** [Errors, warnings and success messages](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-messages-human-readable) — *"All messages should be human-readable (almost always using natural language or tabularized data), and as short and succinct as possible while considering the task at hand."*

**Evidence:**
```go
todoID, err := strconv.ParseInt(todoIDStr, 10, 64)
if err != nil {
	return fmt.Errorf("Invalid Todo ID (error: %v)", err)  // VIOLATION: inconsistent format
}
```
The error output here violates the consistent Go-idiomatic pattern observed in other errors (such as `"--todo is required"` or `"invalid --schedule"`), which are lowercase and do not use capitalized terms or redundant `(error: %v)` formatting.

**Remediation:**
Refactor the error message to be lowercase and follow the standard format:
```go
return fmt.Errorf("invalid todo id: %w", err)
```

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The `reminder-status` command conforms to the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs for Secondary Objects:** Under the CLI guidelines, `add` and `remove` are standard verbs permitted for secondary-object state mutations (`add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** The formatting helpers (`common.Bold`, `common.ColorSection`) and help outputs (`colorizedHelp`, `rootHelp`) use `RenderInlineTags` to safely render bold and colors only when termenv detects terminal capabilities and no environment overrides (`NO_COLOR`) are set.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
