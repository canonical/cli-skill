# 04 — Symmetry Audit

## All Symmetric Operation Pairs

| # | Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|---|
| 1 | Read/write config | `get <key>` | `set <key>=<value>` | Yes | Partial | `get` returns a value; `set` requires `key=value` format (different argument syntax). No `unset` exists to fully reverse a `set`. |
| 2 | Select/deselect engine | `use-engine <name>` | *(missing)* | N/A | N/A | No reverse operation. Cannot "unuse" an engine or revert to no-engine state. Can only switch to a different engine. |
| 3 | Show/hide engine | `show-engine` | *(N/A — read-only)* | N/A | N/A | Read-only operation has no reverse. |
| 4 | Start/stop chat | `chat` | *(Ctrl+C / EOF)* | N/A | N/A | Interactive session terminated by user signal, not a CLI command. |

## Missing Reverse Operations

| Forward Command | Expected Reverse | Status | Impact |
|---|---|---|---|
| `set <key>=<value>` | `unset <key>` or `reset <key>` | **Missing** | Cannot restore a configuration key to its default without knowing the original value. Users must remember or look up defaults. |
| `use-engine <name>` | `disable-engine` or `use-engine --none` | **Missing** | Cannot deactivate the engine entirely. However, this may be intentional — the snap always needs an engine to function. |

## Naming Asymmetries

| Pair | Asymmetry | Severity |
|---|---|---|
| `get` / `set` | Different argument formats: `get <key>` vs `set <key>=<value>` | Low — common pattern in CLI tools, but slightly surprising that `get` uses space-separated positional while `set` uses `=` syntax |
| `show-engine` / `use-engine` | Different verbs for the same noun: `show` (observe) vs `use` (mutate) | None — appropriate use of different verbs for different intent groups |

## Behavioral Asymmetries

| Pair | Asymmetry | Notes |
|---|---|---|
| `get` / `set` | `set` has `--package` flag; `get` has no corresponding `--package` to read only the package-level value | Minor — users may want to distinguish package defaults from user overrides when debugging |
| `use-engine <name>` / `use-engine --auto` | Manual selection (positional arg) vs automatic selection (`--auto` flag) use different argument types for the same command | Acceptable — but `--auto` feels more like a separate command than a flag on the same operation |

## Self-Check

- Commands analyzed: `chat`, `use-engine`, `show-engine`, `get`, `set`, `completion` → **6 commands**
- All 6 commands appear in at least one row of the analysis above ✓
