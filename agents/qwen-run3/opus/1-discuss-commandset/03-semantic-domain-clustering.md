# 03–05 — Semantic Domain Clustering, Symmetry Audit, and Confusion-Pair Audit

> **Scale note**: With < 15 commands, these three analyses are combined per the compact mode rule.

---

## Section 03: Semantic Domain Clustering

| Domain | Count | Commands | Naming Consistent? | Notes |
|--------|-------|----------|-------------------|-------|
| Engine management | 3 | `list-engines`, `show-engine`, `use-engine` | ✓ Yes — all use `-engine(s)` noun suffix | Complete observation + selection set. No create/delete needed (engines are built-in). |
| Configuration | 3 | `get`, `set`, `unset` | ✓ Yes — bare verbs per DE013 convention | Complete CRUD for config values. |
| System observation | 3 | `status`, `show-machine`, `version` | Partial — `status` and `version` use bare nouns, `show-machine` uses verb-noun | `status` follows DE013 shorthand rule. `version` follows DE013 recommended command. |
| User interaction | 2 | `chat`, `webui` | ✗ No — `chat` is a verb, `webui` is a noun | Both start interactive experiences but use different grammar. |
| Maintenance | 1 | `prune-cache` | n/a (single command) | Isolated; no related commands. |
| Hidden/internal | 3 | `run`, `serve-webui`, `debug` | Partial — mixed patterns | Internal use; consistency less critical. |

**Total commands accounted for**: 3 + 3 + 3 + 2 + 1 + 3 = **15** ✓

### Domain Observations

1. **Engine management** is the most consistent domain — a clean `list/show/use` trio.
2. **Configuration** perfectly follows Canonical conventions with `get/set/unset`.
3. **User interaction** is the weakest domain — `chat` and `webui` have no shared naming convention. If the pattern were consistent, it might be `start-chat` and `start-webui`, or `open-chat` and `open-webui`.

---

## Section 04: Symmetry Audit

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|-----------|----------------|-----------------|-------------------|--------------------|-|
| Set config | `set key=value` | `unset key` | ✓ Yes | ✓ Yes — unset reverts to default | Clean pair |
| Select engine | `use-engine <name>` | — | n/a | n/a | No reverse; selection is replaced, not undone |
| Install components | `use-engine` (triggers install) | `prune-cache` (removes components) | ✗ No — different commands, different verbs | Partial — prune only removes inactive engine components | Asymmetric naming: "use" installs, "prune" removes |
| Start server | `sudo snap start qwen36.server` | `sudo snap stop qwen36.server` | ✓ Yes (snap level) | ✓ Yes | Not CLI commands — delegated to snap |
| Open webui | `webui` | — | n/a | n/a | No "close" — browser responsibility |
| Start chat | `chat` | (Ctrl+C) | n/a | n/a | No programmatic "stop chat" |

### Missing Reverse Operations

| Forward | Expected Reverse | Status |
|---------|-----------------|--------|
| `use-engine` | `disable-engine` or revert | Not needed — `use-engine` is a selector, not a toggle |
| `prune-cache` | re-install components | Exists via `use-engine --fix` (partial reverse) |

### Asymmetries Flagged

1. **`use-engine` implicitly installs; `prune-cache` explicitly removes**: The install is a side-effect of selection; removal is a separate explicit action. This is appropriate but could be confusing — users may not realize `use-engine` downloads components.

---

## Section 05: Confusion-Pair Audit

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|-------------|----------------|----------------|
| `status` | `show-engine` | scope ambiguity | **medium** | `status` shows overall snap health; `show-engine` shows engine manifest/requirements details |
| `show-engine` | `list-engines` | functional overlap | **low** | `show-engine` = details of one; `list-engines` = summary of all. Standard list/show pattern. |
| `get` | `show-engine` | scope ambiguity | **medium** | `get` reads config keys (arbitrary key-value); `show-engine` reads engine manifest (structured). Some config values come from engine (e.g., `server`, `model`). |
| `chat` | `webui` | functional overlap | **low** | Both start interactive inference sessions but via different interfaces (terminal vs. browser). |
| `prune-cache` | `use-engine --fix` | functional overlap | **medium** | Both manage engine components but with opposite intent: prune removes unused, fix ensures required are present. |
| `version` | `status` | scope ambiguity | **low** | `version` shows CLI/snap versions only; `status` shows engine, service, and component state. |
| `set` | `use-engine` | functional overlap | **low** | `use-engine` sets engine config layer as a side-effect. Users might try `set server=llamacpp-cuda` instead of `use-engine cuda`. |

### Medium-Risk Pairs — Detailed Disambiguation

1. **`status` vs `show-engine`**: A user wanting to know "is my system working?" should use `status`. A user wanting to know "what are my engine's hardware requirements?" should use `show-engine`. The overlap is that both report on the active engine, but at different levels of detail.

2. **`get` vs `show-engine`**: Config keys like `server`, `model`, and `gpu-layers` originate from the engine manifest. Running `qwen36 get server` returns `llamacpp` — the same information visible in `show-engine` output. Users may not understand which command to use for which keys.

3. **`prune-cache` vs `use-engine --fix`**: Both manipulate component installation state. `prune-cache` removes components of engines you're NOT using. `use-engine --fix` ensures components of the engine you ARE using are present. The naming makes this clear if the user reads help, but a hurried user might confuse them.
