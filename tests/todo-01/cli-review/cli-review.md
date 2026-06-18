# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-01/cmd/todo/main.go` and `/project/tests/todo-01/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Structure and Naming |
| Medium | 0 | — |
| Low | 0 | — |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 95.5 💚 **Excellent**

---

## CLI changes in this PR

The latest changes introduce a non-compliance issue:
* **Renamed List Command to Non-Standard Verb:** The `todo` CLI's primary listing command was changed from the compliant shorthand `list` to `list-todos`. This is a naming violation of the CLI standard, which specifies that the listing command for primary objects should be named `list`.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-list-todos-violates-the-primary-object-listing-shorthand-rule) | The listing command for primary objects must use the shorthand `list` instead of `list-foobar`. | `list-todos` command in `cmd/todo/main.go` violates the primary-object shorthand rule. | Rename `list-todos` to `list` to restore compliance. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] `list-todos` violates the primary-object listing shorthand rule

**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"Use foobars as a shorthand for listing information about all instances of a specific type of secondary object instead of list-foobar ... but snap list for listing snaps."* and standard commonly used commands mapping `tool list` to `overview of all instances of primary type`.

**Evidence:**
```go
	listCmd := &cobra.Command{
		Use:   "list-todos",
		Short: "List todos",
```

The command for listing the primary object `todo` has been named `list-todos`. It must be named `list` to adhere to the standard's shorthand for primary-object listing.

**Remediation:**
Modify `cmd/todo/main.go` to rename `list-todos` to `list`:
```go
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List todos",
```

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** Under the updated CLI guidelines, `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** The formatting helpers (`common.Bold`, `common.ColorSection`) and help outputs (`colorizedHelp`, `rootHelp`) use `RenderInlineTags` to safely render bold and colors only when termenv detects terminal capabilities and no environment overrides (`NO_COLOR`) are set.
- **Primary-Object Command Structure:** Other primary-object actions are verb-led (`show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sinks`, `sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` (and other list commands) strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for list output is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
