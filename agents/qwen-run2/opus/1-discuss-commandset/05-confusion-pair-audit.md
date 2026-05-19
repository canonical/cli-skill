# Confusion-Pair Audit

## Confusion Pairs

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|-------------|----------------|----------------|
| `use-engine` | `set server=llamacpp` | functional overlap | medium | `use-engine` is the high-level command that sets multiple config keys atomically (server, model, mmproj, base-path, gpu-layers) based on engine.yaml; `set server=<value>` is a low-level override of a single key that could put the system in an inconsistent state. |
| `show-engine` | `get server` | scope ambiguity | medium | `show-engine` returns the full engine YAML (name, description, hardware reqs, components, configurations); `get server` returns only the server component name. Users wanting "what engine am I using?" might try either. |
| `use-engine --auto` | `use-engine cpu` | functional overlap | low | `--auto` detects hardware and picks the best engine; positional argument forces a specific engine. The distinction is clear from the flag name but a user might wonder if `--auto` is the default when no argument is given. |
| `set` (via `qwen36 set`) | `snap set qwen36` | functional overlap | medium | Both write to the same snapctl configuration store. `qwen36 set` is a wrapper that may have additional validation or the `--package` flag; `snap set qwen36` is the standard snap mechanism. Users familiar with snap may bypass the CLI wrapper. |
| `get` (via `qwen36 get`) | `snap get qwen36` | functional overlap | medium | Both read from snapctl. Identical behavior for basic usage, but `snap get qwen36` can retrieve nested objects as JSON while `qwen36 get` retrieves scalar values. |
| `chat` | `qwen36.server` | scope ambiguity | low | `chat` is the user-facing interactive command; `qwen36.server` is the daemon. New users might confuse "starting the server" with "starting a chat" — but the server auto-starts as a daemon, so this is rarely an issue in practice. |

## Risk Summary

| Risk Level | Count | Pairs |
|------------|-------|-------|
| High | 0 | — |
| Medium | 4 | use-engine/set, show-engine/get, qwen36 set/snap set, qwen36 get/snap get |
| Low | 2 | use-engine --auto/use-engine cpu, chat/server |

## Analysis

### Key Insight: CLI wrapper vs snap CLI overlap

The most systemic confusion source is that `qwen36 get`/`set` duplicate `snap get`/`set` functionality. This is intentional — the Go binary wraps snapctl to provide a consistent interface without requiring users to know snap internals. However, documentation should clarify the relationship.

### No high-risk pairs

The command set is small enough (6 commands) that confusion risk is inherently limited. The commands operate on clearly different resources (engine vs config key vs interactive session vs completions), reducing semantic overlap.

### Mitigation recommendations

1. Document that `qwen36 get`/`set` is the preferred interface over `snap get`/`set` (for validation and future-proofing).
2. Add help text to `show-engine` clarifying it shows the full engine spec, not just the name.
3. Consider whether `use-engine` with no argument should default to `--auto` or show an error with usage guidance.
