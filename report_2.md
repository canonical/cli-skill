# Detailed Detection Report

Primary source: [tests/ANALYSIS.json](tests/ANALYSIS.json)
Supplemental source: [tests/rule_detection_summary.json](tests/rule_detection_summary.json)

## Overall

- Variations: 20
- Injected violations: 24
- Reported findings: 127
- Reported false negatives: 0

## Rule-Level Detection (Injected Rules)

| Rule ID | Severity | Injected | Detected | Coverage |
|---|---:|---:|---:|---:|
| mutually-exclusive-flags | HIGH | 1 | 1 | 100.0% |
| positional-arg-clarity | HIGH | 2 | 2 | 100.0% |
| required-flag-clarity | HIGH | 1 | 1 | 100.0% |
| state-status-suffix | HIGH | 2 | 2 | 100.0% |
| todo-list-command-verb | HIGH | 2 | 2 | 100.0% |
| error-message-consistency | LOW | 1 | 1 | 100.0% |
| filter-flag-naming | LOW | 2 | 2 | 100.0% |
| group-labels-consistency | LOW | 1 | 1 | 100.0% |
| schedule-list-shorthand | LOW | 1 | 1 | 100.0% |
| short-description-clarity | LOW | 1 | 1 | 100.0% |
| sink-list-shorthand | LOW | 1 | 1 | 100.0% |
| sink-show-shorthand | LOW | 1 | 1 | 100.0% |
| action-verb-alignment | MEDIUM | 2 | 2 | 100.0% |
| flag-plural-for-arrays | MEDIUM | 2 | 2 | 100.0% |
| output-format-flag-consistency | MEDIUM | 2 | 2 | 100.0% |
| todo-show-command-verb | MEDIUM | 1 | 1 | 100.0% |
| todo-todo-topic-clarity | MEDIUM | 1 | 1 | 100.0% |

## Per-Variation Outcome

| Variation | Injected | Reported Findings | Estimated Extra Findings | Report |
|---|---:|---:|---:|---|
| todo-01 | 1 | 1 | 0 | [tests/todo-01/cli-review/cli-review.md](tests/todo-01/cli-review/cli-review.md) |
| todo-02 | 1 | 1 | 0 | [tests/todo-02/cli-review/cli-review.md](tests/todo-02/cli-review/cli-review.md) |
| todo-03 | 1 | 1 | 0 | [tests/todo-03/cli-review/cli-review.md](tests/todo-03/cli-review/cli-review.md) |
| todo-04 | 1 | 8 | 7 | [tests/todo-04/cli-review/cli-review.md](tests/todo-04/cli-review/cli-review.md) |
| todo-05 | 1 | 8 | 7 | [tests/todo-05/cli-review/cli-review.md](tests/todo-05/cli-review/cli-review.md) |
| todo-06 | 1 | 6 | 5 | [tests/todo-06/cli-review/cli-review.md](tests/todo-06/cli-review/cli-review.md) |
| todo-07 | 1 | 8 | 7 | [tests/todo-07/cli-review/cli-review.md](tests/todo-07/cli-review/cli-review.md) |
| todo-08 | 0 | 7 | 7 | [tests/todo-08/cli-review/cli-review.md](tests/todo-08/cli-review/cli-review.md) |
| todo-09 | 0 | 7 | 7 | [tests/todo-09/cli-review/cli-review.md](tests/todo-09/cli-review/cli-review.md) |
| todo-10 | 2 | 9 | 7 | [tests/todo-10/cli-review/cli-review.md](tests/todo-10/cli-review/cli-review.md) |
| todo-11 | 1 | 8 | 7 | [tests/todo-11/cli-review/cli-review.md](tests/todo-11/cli-review/cli-review.md) |
| todo-12 | 1 | 8 | 7 | [tests/todo-12/cli-review/cli-review.md](tests/todo-12/cli-review/cli-review.md) |
| todo-13 | 1 | 8 | 7 | [tests/todo-13/cli-review/cli-review.md](tests/todo-13/cli-review/cli-review.md) |
| todo-14 | 2 | 9 | 7 | [tests/todo-14/cli-review/cli-review.md](tests/todo-14/cli-review/cli-review.md) |
| todo-15 | 1 | 1 | 0 | [tests/todo-15/cli-review/cli-review.md](tests/todo-15/cli-review/cli-review.md) |
| todo-16 | 3 | 10 | 7 | [tests/todo-16/cli-review/cli-review.md](tests/todo-16/cli-review/cli-review.md) |
| todo-17 | 1 | 8 | 7 | [tests/todo-17/cli-review/cli-review.md](tests/todo-17/cli-review/cli-review.md) |
| todo-18 | 1 | 8 | 7 | [tests/todo-18/cli-review/cli-review.md](tests/todo-18/cli-review/cli-review.md) |
| todo-19 | 1 | 8 | 7 | [tests/todo-19/cli-review/cli-review.md](tests/todo-19/cli-review/cli-review.md) |
| todo-20 | 3 | 3 | 0 | [tests/todo-20/cli-review/cli-review.md](tests/todo-20/cli-review/cli-review.md) |