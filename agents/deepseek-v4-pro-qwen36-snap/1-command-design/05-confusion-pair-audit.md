# 05 — Confusion-Pair Audit

All command pairs that share semantic overlap and risk user confusion. Listed exhaustively — every pair with potential overlap is included.

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|-------------|----------------|----------------|
| `list-engines` | `show-engine` | scope ambiguity | **high** | `list-engines` shows all engines in a table with compatibility scores; `show-engine` shows detailed YAML/JSON for one engine. Users may not know which to use when they want "information about engines." The verb distinction (`list` vs `show`) is correct per DE013 but is a learned convention. |
| `chat` | `debug chat` | naming collision | **high** | Both named `chat` but with different interfaces: top-level auto-detects engine and checks server status; debug requires explicit `--base-url`. A user discovering `debug chat` may assume it behaves exactly like `chat`. |
| `webui` | `serve-webui` | functional overlap | **high** | `webui` launches a browser to the web UI; `serve-webui` starts an HTTP server to serve the web UI. Both relate to "web UI" but one is a consumer (launch browser) and the other is a provider (serve files). The names do not make this distinction clear. |
| `serve-webui` | `debug serve-webui` | naming collision | **high** | Two `serve-webui` commands in different namespaces with different flag sets and defaults. Their relationship is unclear: is `debug serve-webui` "the same but with more debug output" (common expectation) or a completely different mode? |
| `use-engine` | `debug select-engine` | functional overlap | **medium** | Both involve engine selection. `use-engine` applies the change; `debug select-engine` only tests without applying. Users may not understand the distinction between "selecting" and "using" an engine. |
| `prune-cache` | `use-engine` (component install) | functional overlap | **medium** | `prune-cache` removes unused engine components; switching engines via `use-engine` installs missing components. Both manage components but in opposite directions. A user might run `prune-cache` to free space, then `use-engine` to switch back, reinstalling what they just removed. No cross-reference between the commands. |
| `get` | `show-engine` | functional overlap | **medium** | `get` shows config values; `show-engine` shows engine details including configurations. A user wanting to see engine config might use `show-engine` instead of `get`, or vice versa. |
| `status` | `show-engine` | functional overlap | **medium** | `status` shows the active engine name as one field; `show-engine` shows the full engine details. Users may use `status` when they want `show-engine` (seeing only the name), or `show-engine` when they want `status` (missing service and model status). |
| `status` | `get` | functional overlap | **low** | `status` shows endpoints and model info that are derived from engine config. A user wanting to see the OpenAI endpoint might use `get` or `status`. Both commands surface overlapping information in different formats. |
| `set` | `use-engine` | functional overlap | **low** | Both can change configuration. `set` modifies individual config keys; `use-engine` applies an engine's entire default configuration set. Users might manually `set` a config that `use-engine` later overwrites. |
| `run` | `chat` | functional overlap | **low** | Both execute something in the engine environment. `run` executes arbitrary subprocesses; `chat` starts a specific chat client. Users might try `run chat-client` expecting the same behavior as `chat`. |
| `show-engine` | `show-machine` | naming collision | **low** | Both start with `show-` followed by a singular noun. Users might momentarily confuse the two, though `engine` and `machine` are semantically distinct. |
| `list-engines` | `show-machine` | scope ambiguity | **low** | Both produce hardware-related output. `list-engines` shows engine compatibility with hardware; `show-machine` shows the hardware itself. Related but distinct. |
| `status` | `version` | functional overlap | **low** | Neither takes arguments. Both show read-only information about the system. Users wanting "what version am I on" might try `status` instead of `version`. |

---

## Confusion Risk Summary

| Risk Level | Count | Pairs |
|-----------|-------|-------|
| **High** | 4 | `list-engines`/`show-engine`, `chat`/`debug chat`, `webui`/`serve-webui`, `serve-webui`/`debug serve-webui` |
| **Medium** | 4 | `use-engine`/`debug select-engine`, `prune-cache`/`use-engine`, `get`/`show-engine`, `status`/`show-engine` |
| **Low** | 6 | `status`/`get`, `set`/`use-engine`, `run`/`chat`, `show-engine`/`show-machine`, `list-engines`/`show-machine`, `status`/`version` |

---

## Key Findings

### 1. The Debug Namespace Creates Confusion (2 high-risk pairs)

The `debug` namespace mirrors two top-level commands (`chat`, `serve-webui`) with different semantics. Users will naturally assume debug variants are "the same command with extra debugging" rather than fundamentally different interfaces. Per DE013, when a second command level is introduced, all commands must follow the same grammar rules. The debug namespace violates this by reusing command names with different contracts.

**Recommendation**: Eliminate `debug chat` and `debug serve-webui` by adding their flags (`--base-url`, `--model`) to the top-level commands, or rename them to clarify their distinct purpose (e.g., `debug chat-with-url`, `debug serve-webui-with-url`).

### 2. `list-engines` vs `show-engine` (high risk per convention)

The `list` vs `show` distinction is standard per DE013 but is not intuitive for novice CLI users. Many tools use `list` and `show` interchangeably. The CLI should make the distinction clearer in help text: `list-engines --help` should mention `show-engine` for detailed view, and vice versa.

### 3. `webui` vs `serve-webui` (high risk)

These are functionally complementary (launch browser vs serve files) but naming does not convey this. Users may ask: "What's the difference between `webui` and `serve-webui`?" The answer requires understanding the web UI architecture (separate server + browser client). Consider renaming to clarify: `open-webui` (or `launch-webui`) and `serve-webui`.

### 4. `prune-cache` and `use-engine` are unlinked (medium risk)

Removing components via `prune-cache` and reinstalling them via `use-engine` are inverse-ish operations but the commands don't cross-reference each other. If a user prunes cache and later wants to switch to a previously-used engine, they might not realize `use-engine` will reinstall components. A note in `prune-cache` help text or output could help.

### 5. Multiple ways to see overlapping information (3 medium-risk pairs)

`get`, `show-engine`, and `status` all surface configuration/engine information with different scopes and formats. Users need to develop a mental model of which command to use for which information. Cross-references in help text would aid discovery.
