# Visual Summary

Data source: [tests/rule_detection_summary.json](tests/rule_detection_summary.json)

## Coverage By Severity (Injected Rule Match)

```mermaid
xychart-beta
    title "Injected Rule Coverage by Severity"
    x-axis [HIGH, MEDIUM, LOW]
    y-axis "Count" 0 --> 10
    bar [8, 8, 8]
    bar [8, 8, 8]
```

Legend: first bar = injected, second bar = detected.

## Variation-Level Signal (Injected vs Reported Findings)

```mermaid
xychart-beta
    title "Per Variation: Injected vs Reported"
    x-axis [v01, v02, v03, v04, v05, v06, v07, v08, v09, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20]
    y-axis "Count" 0 --> 10
    line [1, 1, 1, 1, 1, 1, 1, 0, 0, 2, 1, 1, 1, 2, 1, 3, 1, 1, 1, 3]
    line [1, 1, 1, 8, 8, 6, 8, 7, 7, 9, 8, 8, 8, 9, 1, 10, 8, 8, 8, 3]
```

## Quick Tables

| Severity | Injected | Detected | Coverage |
|---|---:|---:|---:|
| HIGH | 8 | 8 | 100.0% |
| MEDIUM | 8 | 8 | 100.0% |
| LOW | 8 | 8 | 100.0% |