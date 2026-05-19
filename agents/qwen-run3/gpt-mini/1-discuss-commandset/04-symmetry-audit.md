# 04 — Symmetry Audit

Audit results:

- Configuration symmetry: `get` / `set` / `unset` provide a complete CRUD-like surface for configuration keys. Good symmetry.
- Engine symmetry: `list-engines`, `show-engine`, `use-engine` provide discovery and selection; there is no explicit `install-engine` or `remove-engine` command because component management is delegated to snap components. This design is symmetric with snap-based component model but should be documented.
- Show/Info symmetry: `show-engine` and `show-machine` use `show` consistently for detail commands.

Asymmetries / suggestions:
- Consider an explicit `install-engine`/`remove-engine` or expose that via `use-engine --install` variant documented clearly to avoid surprise; currently installation happens automatically during `use-engine` if components are missing.
- `prune-cache` is a specialized destructive command (no direct `restore-cache` counterpart). Provide `--dry-run` to preview.
