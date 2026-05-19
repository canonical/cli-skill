# Safety Model

Destructive or sensitive operations:
- `prune-cache`: removes snap components (potential data/model removal). It prompts for confirmation when interactive and refuses to prune the active engine.
- `set`/`unset`: change persistent configuration and may restart services; `set` can write engine/package config when hidden flags used.
- `use-engine` can install large components and unset engine-specific configuration, prompting for restart. It requires root.
- `run` executes a user-specified subprocess within engine environment — it exposes the host to arbitrary code execution (hidden, but powerful).

Safeguards observed:
- Interactive confirmations for component removal and component installation (`PromptYN`, `PromptRestartToApplyChanges`).
- Non-interactive use supported via `--assume-yes` or `--no-restart` flags.
- `use-engine` checks root (`utils.IsRootUser`) and refuses otherwise.

Missing/weak safeguards:
- `run` is hidden but has only a deprecated `--wait-for-components` and no sandboxing; consider explicit warnings in help and clearer isolation guidance.
- `prune-cache` could be more explicit about what is removed (component list already shown) and could offer a `--dry-run` flag.

Recommendations:
- Add `--dry-run` to `prune-cache` to preview removal without changes.
- Improve `run` help text to emphasize it runs arbitrary commands and is not sandboxed.
- Ensure `prune-cache` outputs component sizes (already present) and checks for user confirmation when stdout is a terminal.
