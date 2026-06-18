# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of `/project/todo`

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 0 | None |
| Medium | 0 | None |
| Low | 1 | Grammar + Vocabulary |
| Unrated | 0 | None |
| **Total** | **1** | |

**Overall rating:** 99.09 💚 **Excellent**
> The scoring algorithm starts with 100%, number of commands N, weight W=100/N. For each High violation, reduce by 2*W; Medium violation by 1*W; Low violation by 0.5*W. Clamp to 0-100.

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [LOW-1](#low-1-use-shorthand-sinks-for-listing-secondary-objects-instead-of-list-sinks) | Use plural secondary object name `sinks` as a shorthand for listing information instead of `list-sinks`. | [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-shorthand-for-listing) | Rename `list-sinks` to `sinks` to comply with the standard. |

---

## Non-compliance Findings (with citations)

### [LOW-1] Use shorthand 'sinks' for listing secondary objects instead of 'list-sinks'
**CLI Standard citation:** [Grammar + Vocabulary](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-shorthand-for-listing) — *"Use foobars as a shorthand for listing information about all instances of a specific type of secondary object instead of list-foobar"*
**Evidence:**
```go
	sinksCmd := &cobra.Command{
		Use:   "list-sinks",
		Short: "List sinks",
```
The client CLI defines `list-sinks` as the command to retrieve all sinks.
**Remediation:** Rename the `list-sinks` command to `sinks` to follow the standard shorthand for listing guidelines.

---

## Compliant Findings Summary

- **Verbs represent actions**: All subcommands that act on primary or secondary objects are named using standard verbs (e.g., `list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- **Logical command grouping**: Commands are organized into intuitive categories/groups (`Todos`, `Sinks`, `Schedules`, `Other`) in help outputs.
- **Single short/long flags avoided**: The tool does not duplicate flags with both short and long options for the same action, prioritizing clean and descriptive long flags.
- **Color auto-detection and NO_COLOR support**: ANSI escape styling dynamically disables itself when the standard output is redirected or the `NO_COLOR` environment variable is defined.
- **Concise messages and direct tone of voice**: System outputs and status information avoid chattiness and explain configuration and hints concisely.
- **Tabular spacing rules**: Spacing between printed columns conforms to the two-space delimiter requirement.
- **Empty state stderr output**: The `list` command successfully prints its empty state message (`No todos found`) to standard error rather than standard output while retaining a successful exit code (0).
