# 04 — Symmetry Audit

> **Scale note**: With < 15 commands (compact mode), this section is combined with sections 03 and 05 in [03-semantic-domain-clustering.md](03-semantic-domain-clustering.md). See the "Section 04: Symmetry Audit" heading in that file for the full analysis.

## Summary

The qwen36 CLI has one clean symmetric pair (`set`/`unset`) and otherwise relies on asymmetric operations appropriate for its domain:

- `use-engine` is a selector with no inverse (appropriate — you select a different engine, not "un-select")
- `prune-cache` serves as an informal reverse of `use-engine`'s component installation side-effect
- Server start/stop is delegated to snap daemon commands, not the CLI itself

No missing reverse operations represent a usability gap for a CLI of this scale.
