# 05 — Confusion-Pair Audit

> **Scale note**: With < 15 commands (compact mode), this section is combined with sections 03 and 04 in [03-semantic-domain-clustering.md](03-semantic-domain-clustering.md). See the "Section 05: Confusion-Pair Audit" heading in that file for the full analysis.

## Summary

Three medium-risk confusion pairs were identified:

| Pair | Risk | Core Issue |
|------|------|-----------|
| `status` vs `show-engine` | medium | Both show engine info at different granularities |
| `get` vs `show-engine` | medium | Engine config keys overlap with engine manifest data |
| `prune-cache` vs `use-engine --fix` | medium | Both manipulate component state with opposite intents |

No high-risk pairs exist. The CLI's small scale and consistent naming keep confusion low overall.
