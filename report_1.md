# Detailed Detection Report

Primary source: completed harness run artifacts in `/tmp/cli-skill-harness-daon7Z/tests/`
Supplemental source: `ANALYSIS.json` from the completed run

## Overall

- Variations: 20
- Injected violations: 24
- Detected violations: 20
- Reported false negatives: 4
- Average detection coverage: 89.2%

## Rule-Level Detection (Injected Rules)

| Rule ID | Severity | Injected | Detected | Coverage |
|---|---:|---:|---:|---:|
| action-verb-alignment | MEDIUM | 2 | 2 | 100.0% |
| error-message-consistency | LOW | 1 | 1 | 100.0% |
| filter-flag-naming | LOW | 2 | 2 | 100.0% |
| flag-plural-for-arrays | MEDIUM | 2 | 2 | 100.0% |
| group-labels-consistency | LOW | 1 | 0 | 0.0% |
| mutually-exclusive-flags | HIGH | 1 | 1 | 100.0% |
| output-format-flag-consistency | MEDIUM | 2 | 2 | 100.0% |
| positional-arg-clarity | HIGH | 2 | 2 | 100.0% |
| required-flag-clarity | HIGH | 1 | 1 | 100.0% |
| schedule-list-shorthand | LOW | 1 | 1 | 100.0% |
| short-description-clarity | LOW | 1 | 1 | 100.0% |
| sink-list-shorthand | LOW | 1 | 1 | 100.0% |
| sink-show-shorthand | LOW | 1 | 1 | 100.0% |
| state-status-suffix | HIGH | 2 | 2 | 100.0% |
| todo-list-command-verb | HIGH | 2 | 0 | 0.0% |
| todo-show-command-verb | MEDIUM | 1 | 1 | 100.0% |
| todo-todo-topic-clarity | MEDIUM | 1 | 0 | 0.0% |

## Per-Variation Outcome

| Variation | Injected | Detected | False Negatives | Coverage | Artifact Notes |
|---|---:|---:|---:|---:|---|
| todo-01 | 1 | 0 | 1 | 0.0% | temporary run artifact |
| todo-02 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-03 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-04 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-05 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-06 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-07 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-08 | 0 | 0 | 0 | 100.0% | temporary run artifact |
| todo-09 | 0 | 0 | 0 | 100.0% | temporary run artifact |
| todo-10 | 2 | 2 | 0 | 100.0% | temporary run artifact |
| todo-11 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-12 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-13 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-14 | 2 | 1 | 1 | 50.0% | temporary run artifact |
| todo-15 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-16 | 3 | 2 | 1 | 66.7% | temporary run artifact |
| todo-17 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-18 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-19 | 1 | 1 | 0 | 100.0% | temporary run artifact |
| todo-20 | 3 | 2 | 1 | 66.7% | temporary run artifact |

## Notes

- The completed harness run reported `89.2%` average coverage across all 20 variations.
- `Detected` and `False Negatives` are taken from `ANALYSIS.json` produced by the harness.
- The temporary workspace preserved the per-variation artifacts under `/tmp/cli-skill-harness-daon7Z/tests/` for later inspection.