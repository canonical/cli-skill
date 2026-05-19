# Safety Model

## Destructive Operations Inventory

### `set` — Configuration changes with restart
- **Impact**: Modifies snap config (persistent). May trigger snap daemon restart, interrupting any active inference sessions.
- **Gating**: Interactive prompt (`Restart <snap> to apply the changes? [Y/n]`) unless `--no-restart` or `--assume-yes`.
- **Dry-run**: Not supported.
- **Recovery**: Re-run `set` with the old value, or `unset` to revert.

### `unset` — Configuration removal with restart
- **Impact**: Removes user config layer, potentially reverting to package/engine defaults. Restart side effect same as `set`.
- **Gating**: Interactive prompt on restart unless `--no-restart` or `--assume-yes`.
- **Dry-run**: Not supported.
- **Recovery**: Re-run `set` with desired value.

### `use-engine` — Engine switch + component install/remove
- **Impact**: Changes active engine, downloads/instals snap components (potentially multi-GB), applies engine configs, may restart daemon.
- **Gating**:
  - Component list printed with sizes.
  - Interactive `Do you want to continue? [Y/n]` unless `--assume-yes`.
  - Restart prompt unless `--no-restart`.
- **Dry-run**: Not supported. `--auto` evaluates but still applies.
- **Recovery**: `use-engine --fix` reinstalls missing components and re-applies config. Switch back to prior engine manually.

### `prune-cache` — Component removal
- **Impact**: Removes snap components used by inactive engines, freeing disk space. Cannot be undone without re-downloading.
- **Gating**:
  - Lists components and affected engines.
  - Interactive confirmation prompt (`Continue pruning ...? [y/N]`) unless non-terminal or already confirmed.
- **Dry-run**: Not supported.
- **Force**: Not supported (no `--force` flag).
- **Recovery**: Re-run `use-engine` for the removed engine; components will be re-installed.

## Non-Destructive but State-Changing Operations

### `run <command>`
- **Impact**: Creates **temporary symlinks** in the active engine's layout paths. These are cleaned up on normal process exit via `defer`, but **not on SIGTERM/SIGKILL**.
- **Risk**: Stale symlinks may remain after unclean shutdown, potentially confusing future engine environment loads or component lookups.
- **Mitigation**: Documented TODO in code to add signal handling.

## Force Flags

There is **no `--force` flag** anywhere in the CLI. Destructive actions require either interactive confirmation or explicit `--assume-yes`.

## Confirmation Semantics

| Command | Default Prompt Answer | Override |
|---|---|---|
| `set` restart | Yes | `--no-restart`, `--assume-yes` |
| `unset` restart | Yes | `--no-restart`, `--assume-yes` |
| `use-engine` component install | Yes | `--assume-yes` |
| `use-engine` restart | Yes | `--no-restart`, `--assume-yes` |
| `prune-cache` removal | No | (none; non-terminal skips prompt) |

## Safety Gaps and Risks

1. **No dry-run / preview mode**: Users cannot see exactly what `use-engine` or `prune-cache` will do before it happens. `list-engines` shows compatibility but not the install plan.
2. **No rollback**: If `use-engine` fails mid-way (e.g., component install succeeds but config application fails), the system may be in a half-configured state. There is no atomic transaction.
3. **Stale symlink risk**: `run` leaves temporary symlinks on signal death. Repeated crashes could accumulate dangling links.
4. **No backup before config change**: `set` and `unset` do not snapshot prior config. Misconfiguration may require manual reversion.
5. **`--assume-yes` is a blunt instrument**: It suppresses ALL prompts in a command, not just one. A user wanting to skip the restart prompt but still confirm component installation cannot express that preference.
