---
description: "Explore renaming a CLI command: alternative names, migration strategy, and ripple effects."
---

Use the `cli-review` skill. Run the `rename-command` workflow from the "Command Design Phase: 1-command-design" section.

Follow the Context Resolution rules: use `0-analysis/commandset.md` if it exists, otherwise ask the user for the existing command set.

Ask the user for the command to rename and the reason if not already provided.

Generate `1-command-design/rename-<old>-to-<new>.md` covering:
- Problem diagnosis
- 3-5 alternative name candidates with rationale and risk
- Migration and deprecation strategy (aliases, warning period, removal timeline)
- Ripple effects on subcommands, help text, flags, and configuration
