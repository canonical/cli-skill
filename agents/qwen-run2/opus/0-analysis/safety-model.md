# Safety Model

## Destructive Operations

| Command | Destructive? | What it Changes | Reversible? |
|---------|-------------|-----------------|-------------|
| `use-engine` | Moderate | Overwrites 4-5 snap config keys; may restart server daemon | Yes — run again with previous engine name |
| `set` | Low-Moderate | Overwrites a configuration value | Yes — set to previous value |
| `chat` | None | Read-only interaction | N/A |
| `show-engine` | None | Read-only | N/A |
| `get` | None | Read-only | N/A |
| `completion` | None | Read-only | N/A |
| `server` (daemon) | Moderate | Loads model into memory (~22GB RAM), binds network port | Yes — `snap stop qwen36.server` |

## Confirmation Mechanisms

| Command | Confirmation? | Mechanism | Bypass |
|---------|--------------|-----------|--------|
| `use-engine` | Yes (interactive) | Prompt before switching | `--assume-yes` flag |
| `set` | No | — | — |
| All others | No | — | — |

### `use-engine` confirmation behavior

- When run interactively, `use-engine` prompts the user to confirm the engine switch before writing configuration and restarting the server.
- The `--assume-yes` flag (used in the install hook) suppresses this prompt.
- There is no `--dry-run` flag to preview what would change.

## Dry-Run Support

**None.** No command in the qwen36 CLI supports `--dry-run` or equivalent preview functionality.

- `use-engine --auto` cannot show what it would select without actually selecting it.
- `set` cannot preview the effect of a configuration change.

## Force Flags

| Flag | Command | Effect |
|------|---------|--------|
| `--assume-yes` | `use-engine` | Skips interactive confirmation |

There is no `--force` flag on any command. The `--assume-yes` flag serves a similar role for non-interactive contexts (hooks, scripts).

## Recovery Behavior

### Server component timeout
- If required snap components are not installed within 3600 seconds, `server.sh` calls `snapctl stop qwen36` to prevent systemd restart loops.
- Recovery: install missing components, then `snap start qwen36.server`.

### Server health check failure
- `wait-for-server.sh` gives a 60-second window for the server to become healthy.
- On fatal errors (exit code 2), it fails immediately without waiting.
- Recovery: check `snap logs qwen36.server`, fix the issue, restart.

### Engine misconfiguration
- If `gpu-layers` is not set for the CUDA engine, the server exits immediately with exit code 1.
- Recovery: `qwen36 set gpu-layers=99` or `qwen36 use-engine cuda` (which sets it from engine.yaml).

## Safety Gaps

1. **No undo/rollback**: Switching engines with `use-engine` has no built-in rollback. If the new engine fails, the user must manually switch back.
2. **No dry-run anywhere**: Users cannot preview changes before applying them.
3. **No backup of previous config**: `use-engine` overwrites config keys without saving the previous state.
4. **Server auto-restart not gated**: The daemon's `daemon: simple` declaration means systemd will restart it automatically on crash, which could cause rapid restart loops if configuration is broken (mitigated by the component timeout logic in server.sh).
5. **Install hook auto-selects without user consent**: On fresh install, `use-engine --auto --assume-yes` selects and configures an engine without user interaction, which is appropriate for a snap but means users may not know what was selected.
