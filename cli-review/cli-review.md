# Canonical CLI automated review report
*This report is AI-generated. Please [report issues with the cli-skill](https://github.com/canonical/cli-skill/issues/new/choose) so we can improve this report.*

**Scope:** CLI standard compliance review of /project/todo/cmd/todo/main.go and /project/todo/cmd/todod/main.go

## Summary

| **Severity** | **Count** | **Guideline Categories** |
| --- | --- | --- |
| High | 1 | Positional parameter usage |
| Medium | 1 | Command-set complexity |
| Low | 0 | — |
| Unrated | 1 | Non-standard verb choice (SHOULD consistency guidance) |
| **Total** | **3** | |

**Overall rating:** 94.0 🟢 **Good**

---

## CLI changes in this PR

Not applicable (review executed on local workspace state without PR diff context).

---

## Compliance matrix

| Finding | Rule Summary | Evidence | Notes |
|---------|--------------|----------|-------|
| [HIGH-1](#high-1-create-uses-two-positional-parameters-for-primary-object-creation) | Positional parameters should only be used when order is natural and unambiguous. | `create <todo-id> <title>` and `cobra.MinimumNArgs(2)` in `todo/cmd/todo/main.go` | Title semantics are content-like and not naturally positional. |
| [MEDIUM-1](#medium-1-command-surface-complexity-is-high-21-commands-in-todo) | Commands should converge to concise grammar and remain easy to learn/remember. | `todo` command set contains 21 top-level commands in `todo/cmd/todo/main.go` | Scope policy maps very high complexity to medium severity. |
| [UNRATED-1](#unrated-1-schedule-secondary-object-uses-addremove-verbs-instead-of-standard-createupdatedelete-family) | Common command consistency guidance lists `create-foo`, `delete-foo`, `update-foo` for secondary objects. | `add-schedule`, `remove-schedule` in `todo/cmd/todo/main.go` | SHOULD-style consistency issue; reported as Unrated per scope. |

---

## Non-compliance Findings (with citations)

### [HIGH-1] Create uses two positional parameters for primary-object creation
**CLI Standard citation:** [Positional Parameters](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#positional-parameters) — positional parameters should only be used when there is no doubt about positional meaning and order is naturally memorizable.

**Evidence:**
```go
createCmd := &cobra.Command{
    Use:   "create <todo-id> <title>",
    Short: "Create a todo",
    Args:  cobra.MinimumNArgs(2),
```

Using two positional parameters for creation requires users to memorize ordering (`id`, then free-form title) for a non-directional action.

**Remediation:** Prefer a single positional identifier with explicit content flag, e.g. `create <todo-id> --title "..."`.

### [MEDIUM-1] Command surface complexity is high (21 commands in `todo`)
**CLI Standard citation:** [Conclusion: Keep it Simple (for your Users)](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#conclusion-keep-it-simple-for-your-users) and [All commands must converge](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#rule-all-commands-must-converge).

**Evidence:**
```go
root.AddCommand(listCmd, showCmd, createCmd, updateCmd, closeCmd, reopenCmd, rejectCmd)
root.AddCommand(sinksCmd, sinkCmd, createSinkCmd, updateSinkCmd, deleteSinkCmd, enableSinkCmd, disableSinkCmd)
root.AddCommand(schedulesCmd, scheduleCmd, addScheduleCmd, removeScheduleCmd)
root.AddCommand(motdMessageCmd, statusCmd, versionCmd)
```

The command set is broad, increasing memorization burden and grammar complexity.

**Remediation:** Consider reducing exposed verb surface where possible (e.g. merge lifecycle variants or hide less-frequent operations behind clearer grouping boundaries).

### [UNRATED-1] Schedule secondary object uses add/remove verbs instead of standard create/update/delete family
**CLI Standard citation:** [Commonly used commands](https://github.com/canonical/cli-skill/blob/main/cli-skill/references/cli-standard.md#commonly-used-commands) (`tool create-foo <id>`, `tool delete-foo <id>`, `tool update-foo <id>`).

**Evidence:**
```go
addScheduleCmd := &cobra.Command{ Use: "add-schedule <schedule-id>" }
removeScheduleCmd := &cobra.Command{ Use: "remove-schedule <schedule-id>" }
```

The schedule secondary-object mutation verbs diverge from the common consistency family.

**Remediation:** Consider `create-schedule` and `delete-schedule` (and optional `update-schedule` if mutability is introduced later).

---

## Compliant Findings Summary

- Primary-object actions are verb-led (`list`, `show`, `create`, `update`, `close`, `reopen`, `reject`).
- Secondary-object listing/detail shorthand follows standard patterns (`sinks`, `sink`, `schedules`, `schedule`).
- Verb-noun form is used for secondary-object mutation commands (`create-sink`, `update-sink`, `delete-sink`, etc.).
- Hierarchy is expressed via flags for schedule scoping (`--todo`) instead of deeper nested command levels.
- No dual short/long flag pairs were introduced for the same action.
- Help flags are supported through Cobra defaults.
