# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-01/cmd/todo/main.go` and `/project/tests/todo-01/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Command Naming |
| Medium | 0 | — |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 95.45 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## CLI changes in this PR

* **Decreased Primary Command Shorthand Compliance:** The `list` command configuration was modified to change its command name from `list` to `list-todos`. This violates the CLI standard's primary-object shorthand rule which specifies using `list` to overview all instances of the primary object type (todos).

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-list-todos-violates-the-primary-object-shorthand-naming-rule) | The overview command for the primary object (todos) must be named `list` as a shorthand, rather than using `list-todos`. | `list-todos` (see [The standard grammar for Canonical command-line interfaces](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#the-standard-grammar-for-canonical-command-line-interfaces)) | Rename `list-todos` to `list` to conform to the standard primary-object shorthand pattern. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] list-todos violates the primary-object shorthand naming rule

**CLI Standard citation:** [Commonly Used Commands](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) — *"tool list | overview of all instances of primary type"*

**Evidence:**
```go
	listCmd := &cobra.Command{
		Use:   "list-todos",
		Short: "List todos",
```
The command overviewing the primary object (`todo`) is named `list-todos` instead of using the shorthand `list`.

**Remediation:** Rename the `list-todos` command Use string to `list` in `cmd/todo/main.go` to conform to the primary-object shorthand naming pattern:
```go
	listCmd := &cobra.Command{
		Use:   "list",
```

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** Formatting helpers and help outputs use dynamic color/styling capability detection via `github.com/muesli/termenv`. Colors/formatting are disabled when stdout/stderr is redirected or `NO_COLOR` is present.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for `list` is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
