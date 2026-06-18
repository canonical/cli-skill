# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/tests/todo-02/cmd/todo/main.go` and `/project/tests/todo-02/cmd/todod/main.go`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | — |
| Medium | 0 | — |
| Low | 1 | Structure and Naming |
| Unrated | 0 | — |
| **Total** | **1** | |

**Overall rating:** 99.1 💚 **Excellent**

---

## CLI changes in this PR

The latest changes introduce a non-compliance issue:
* **Renamed Sinks Command to Non-Standard Verb:** The `todo` CLI's secondary listing command for sinks was changed from the compliant shorthand `sinks` to `list-sinks`. This is a naming violation of the CLI standard, which specifies that the listing command for secondary objects should use the plural name of the object as a shorthand instead of `list-foobar`.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [LOW-1](#low-1-list-sinks-violates-the-secondary-object-listing-shorthand-rule) | The listing command for secondary objects must use the plural name of the object as a shorthand instead of `list-foobar`. | `list-sinks` command in `cmd/todo/main.go` violates the secondary-object shorthand rule. | Rename `list-sinks` to `sinks` to restore compliance. |

---

## Non-compliance Findings (with citations)

### [LOW-1] `list-sinks` violates the secondary-object listing shorthand rule

**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#grammar--vocabulary) — *"Use foobars as a shorthand for listing information about all instances of a specific type of secondary object instead of list-foobar (e.g. snap services over snap list-services for listing services of a snap...)"* and standard commonly used commands mapping `tool foobars` to `overview of all instances of a secondary type`.

**Evidence:**
```go
	// Sink commands
	sinksCmd := &cobra.Command{
		Use:   "list-sinks",
		Short: "List sinks",
```

The command for listing the secondary object `sink` has been named `list-sinks`. It must be named `sinks` to adhere to the standard's shorthand for secondary-object listing.

**Remediation:**
Modify `cmd/todo/main.go` to rename `list-sinks` to `sinks`:
```go
	// Sink commands
	sinksCmd := &cobra.Command{
		Use:   "sinks",
		Short: "List sinks",
```

---

## Compliant Findings Summary

- **State-Display Shorthand Pattern:** The `reminder-status` command conforms perfectly with the standard's state-display shorthand (`foobar-status` pattern) for secondary objects.
- **Standard Verbs (add/remove) for Secondary Objects:** Under the updated CLI guidelines, `add` and `remove` are standard verbs permitted for secondary-object state mutations (e.g., `add-schedule` and `remove-schedule` are fully compliant).
- **TTY-aware Color and Formatting:** The formatting helpers (`common.Bold`, `common.ColorSection`) and help outputs (`colorizedHelp`, `rootHelp`) use `RenderInlineTags` to safely render bold and colors only when termenv detects terminal capabilities and no environment overrides (`NO_COLOR`) are set.
- **Primary-Object Command Structure:** Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Secondary-Object Listing/Details:** Shorthand patterns for secondary objects are adhered to (`sink`, `schedules`, `schedule`).
- **Flat Secondary Mutation Hierarchy:** Verb-noun naming structure (`create-sink`, `delete-sink`) is used for sinks, and flat configuration is handled with flags (`--todo`) rather than deep subcommands.
- **No Dual Flags:** Short and long flags are not duplicated for the same action.
- **Help/Version Support:** The CLI supports `--help` via Cobra defaults and exposes standard `version` and `status` commands.
- **Table Formatting Standards:** Tabular data output in `list` (and other list commands) strictly conforms to the 2-space column separator width, capitalized/left-aligned headers, and contains zero ASCII line decorations.
- **Proper Empty State Handling:** The empty state for list output is human-readable, routed to `stderr`, and terminates with a success exit code of 0.
