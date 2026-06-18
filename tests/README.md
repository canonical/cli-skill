# CLI Review Testing Harness

## Overview

This testing harness validates the **CLI Review skill** itself by:

1. **Generating 20 test variations** of `/todo` with deliberately injected CLI standard violations
2. **Running `cli-review` on each variation** using a specified provider/model
3. **Analyzing detection accuracy** to identify false positives and false negatives
4. **Interrogating the model** (in a continued session) about improvements
5. **Generating an amendment summary** using a separate Gemini session

## CLI Standard Rules Being Tested

The harness tests these rules extracted from `cli-skill/references/cli-standard.md` and applied specifically to the todo CLI:

| Rule ID | Title | Violation Pattern |
|---------|-------|-------------------|
| `todo-list-command-verb` | Shorthand for listing | Rename `list` → `list-todos` |
| `todo-show-command-verb` | Verb consistency | Change `show` → `info` |
| `sink-list-shorthand` | Shorthand for secondary list | Rename `sinks` → `list-sinks` |
| `sink-show-shorthand` | Shorthand for secondary show | Rename `sink` → `show-sink` |
| `schedule-list-shorthand` | Shorthand for listing | Rename `schedules` → `list-schedules` |
| `schedule-verb-consistency` | Consistent verb pairs | `create-schedule` instead of `add-schedule` |
| `sink-verb-consistency` | Consistent verb pairs | `add-sink` instead of `create-sink` |
| `state-status-suffix` | Status suffix rule | Rename `reminder-status` → `reminders` |
| `positional-arg-clarity` | Clear argument names | Use `<id>` instead of `<todo-id>` |
| `flag-plural-for-arrays` | Plural flags for arrays | Use `--sinks` for repeatable values |
| `output-format-flag-consistency` | Consistent flag names | Use `--output` instead of `--format` |
| `filter-flag-naming` | Filter flag pattern | Use `--filter-state` instead of `--state` |
| `required-flag-clarity` | Mark required flags | Hide required flag in help text |
| `mutually-exclusive-flags` | Flag exclusivity | Allow both `--due` and `--clear-due` |
| `action-verb-alignment` | Action clarity | Rename `close` → `mark-closed` |
| `group-labels-consistency` | Label formatting | Use `Todo Commands` instead of `Todos:` |
| `short-description-clarity` | Description verb-first | Use `Todo list` instead of `List todos` |
| `flag-help-completeness` | Help text completeness | Remove help from `--state` flag |
| `todo-todo-topic-clarity` | Avoid ambiguity | Make topic command clash with list |
| `error-message-consistency` | Error format | Inconsistent error message formatting |

## Test Variations

The harness creates 20 variations with todo-specific violations:

- **Variation 1**: `todo-list-command-verb` — Rename `list` to `list-todos`
- **Variation 2**: `sink-list-shorthand` — Rename `sinks` to `list-sinks`
- **Variation 3**: `schedule-list-shorthand` — Rename `schedules` to `list-schedules`
- **Variation 4**: `state-status-suffix` — Rename `reminder-status` to `reminders`
- **Variation 5**: `positional-arg-clarity` — Use `<id>` instead of `<todo-id>`
- **Variation 6**: `flag-plural-for-arrays` — Use `--sinks` for repeatable flag
- **Variation 7**: `output-format-flag-consistency` — Use `--output` instead of `--format`
- **Variation 8**: `schedule-verb-consistency` — Use `create-schedule` instead of `add-schedule`
- **Variation 9**: `sink-verb-consistency` — Use `add-sink` instead of `create-sink`
- **Variation 10**: Combined — `todo-show-command-verb` + `sink-show-shorthand`
- **Variation 11**: Combined — `filter-flag-naming` changes
- **Variation 12**: Combined — `required-flag-clarity` issues
- **Variation 13**: Combined — `action-verb-alignment` changes
- **Variation 14**: Combined — `group-labels-consistency` + `short-description-clarity`
- **Variation 15**: Combined — `mutually-exclusive-flags` issues
- **Variation 16**: Triple — `todo-list-command-verb` + `state-status-suffix` + `positional-arg-clarity`
- **Variation 17**: Combined — `sink-verb-consistency` + `flag-plural-for-arrays`
- **Variation 18**: Combined — `schedule-verb-consistency` + `output-format-flag-consistency`
- **Variation 19**: Combined — `filter-flag-naming` + `flag-help-completeness`
- **Variation 20**: Triple — `todo-todo-topic-clarity` + `error-message-consistency` + `action-verb-alignment`

Each variation:
- Lives in `/tests/todo-<variation>/`
- Contains a copy of the `/todo` project with injected violations
- Stores metadata in `violation_metadata.json`
- Receives a report and session from `cli-review`

## Directory Structure

```
/tests/
├── todo-01/
│   ├── cmd/
│   ├── internal/
│   ├── violation_metadata.json
│   └── reports/
│       ├── cli-review.md
│       └── session.jsonl
├── todo-02/
│   └── ...
├── ...
├── todo-20/
│   └── ...
└── ANALYSIS.json
```

## Usage

### Phase 1-3: Generate Variations, Run Reviews, Analyze

```bash
python3 scripts/testing_harness.py <provider> <model>
```

**Example:**
```bash
python3 scripts/testing_harness.py anthropic claude-3-5-sonnet
```

This will:
1. Create 20 variations in `/tests/todo-01/` through `/tests/todo-20/`
2. Run `pi cli-review` on each using your provider/model
3. Save reports and sessions
4. Analyze detection accuracy
5. Generate `/tests/ANALYSIS.json` with results

### Phase 4: Interrogate Model (Continued Session)

In the same session, after analysis:

```
Query: Based on the false negatives and false positives identified above,
what checks, rules, or logic should be added to the cli-review command to 
improve detection accuracy and reduce false positives?

Context:
- False negatives: [list of rules not detected]
- False positives: [list of incorrect detections]
- Common patterns in missed violations: [patterns that went undetected]
```

### Phase 5: Amendment Summary (New Gemini Session)

Create a final session with `pi` using Gemini:

```bash
pi --provider google --model gemini-3.1-pro-preview
```

Provide the model with:
- All recommendations from Phase 4
- Suggestions for rule refinements
- Proposed additions to cli-review logic
- Implementation priorities

## Output Files

### Per Variation

- **`violation_metadata.json`**: Metadata about injected violations
  ```json
  {
    "variation_id": 1,
    "violations": [
      {
        "rule_id": "rule-commands-are-verbs",
        "rule_title": "Commands are verbs",
        "description": "...",
        "file_path": "cmd/todo/main.go",
        "change_description": "Changed 'reminder-status' to 'reminders'"
      }
    ]
  }
  ```

- **`reports/cli-review.md`**: The generated report from `cli-review`

- **`reports/session.jsonl`**: The full session transcript

### Overall Summary

- **`ANALYSIS.json`**: Aggregate statistics
  ```json
  {
    "total_variations": 20,
    "total_violations_injected": 35,
    "average_detection_coverage": 0.85,
    "results": [
      {
        "variation": "todo-01",
        "violations_injected": 1,
        "violations_detected": 1,
        "false_negatives": [],
        "coverage_percent": 100.0
      }
    ]
  }
  ```

## Interpreting Results

### Coverage Metrics

- **Coverage %**: Percentage of injected violations detected (100% = all found)
- **False Negatives**: Violations that were NOT detected by cli-review
- **False Positives**: Detections that don't correspond to injected violations

### What to Look For

1. **Low coverage on specific rules**: E.g., if `rule-no-dual-flags` has 0% coverage, the cli-review logic doesn't check for dual flags
2. **Patterns in false negatives**: E.g., violations in comments vs. code, or violations in specific files
3. **False positives**: May indicate over-aggressive pattern matching or context misunderstanding

## Extending the Harness

### Add New Rules

Edit `CLI_RULES` array in `testing_harness.py`:

```python
CLI_RULES.append({
    "id": "rule-new-rule",
    "title": "My New Rule",
    "description": "Rule description",
    "violation_pattern": "How to violate it"
})
```

### Add New Variation Patterns

Edit `_inject_violations()` method to map variation IDs to rule combinations:

```python
violation_patterns = {
    # ... existing patterns ...
    21: ["rule-commands-are-verbs", "rule-new-rule"],
}
```

### Customize Violation Injection

Edit `_create_violation()` to inject your violation pattern:

```python
elif rule_id == "rule-new-rule":
    # Inject your violation
    if 'pattern' in content:
        content = content.replace('pattern', 'violation')
        change_desc = "Description of what was changed"
```

## Notes

- The harness uses `pi` to invoke the CLI review skill
- Sessions are continued (not restarted) to accumulate context
- Violations are injected directly into Go code files
- Analysis is simple pattern matching; can be enhanced with full parser
- Results depend on the provider/model's capability and the cli-review logic

## Future Enhancements

- [ ] Parse reports more intelligently (regex + AST)
- [ ] Generate variations programmatically from rule definitions
- [ ] Create mutation matrix (variations × rules)
- [ ] Track model performance across different providers
- [ ] Generate detailed HTML reports with diffs
- [ ] Integrate with CI/CD for regression testing
