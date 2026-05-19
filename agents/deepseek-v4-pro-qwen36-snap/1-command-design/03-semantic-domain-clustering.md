# 03 — Semantic Domain Clustering

All 18 commands grouped by the resource domain they operate on. Every command appears in exactly one domain.

| Domain | Count | Commands | Naming Consistent? | Notes |
|--------|-------|----------|--------------------|-------|
| **Engine** | 6 | `list-engines`, `show-engine`, `use-engine`, `debug validate-engines`, `debug select-engine`, `run` | ⚠️ Partial | Core engine management: `list`, `show`, `use` are consistent. `run` executes subprocesses in engine environment — it's engine-adjacent but grouped here as it depends on active engine. Debug commands use different verbs (`validate`, `select`) vs top-level verbs. |
| **Configuration** | 3 | `get`, `set`, `unset` | ✅ Consistent | Clean `get`/`set`/`unset` triad. No noun suffix — operates on the implicit "config" namespace. Per DE013, `get`/`set`/`unset` is the standard for configuration management. |
| **Machine** | 1 | `show-machine` | ✅ N/A | Single command. Noun form `machine` is consistent with `show-engine`. |
| **Status** | 1 | `status` | ✅ N/A | Singleton. Per DE013, `status` is the standard command for current tool state. |
| **Version** | 1 | `version` | ✅ N/A | Singleton. Per DE013 standard. |
| **Cache/Storage** | 1 | `prune-cache` | ✅ N/A | Singleton. Verb `prune` is uncommon but consistent within the domain. |
| **Chat** | 2 | `chat`, `debug chat` | ⚠️ Duplicate | Two chat commands in different namespaces with different semantics: top-level uses active engine, debug requires explicit `--base-url`. |
| **Web UI** | 3 | `webui`, `serve-webui`, `debug serve-webui` | ❌ Inconsistent | `webui` is a noun-as-command (launch). `serve-webui` is verb-noun (serve). Two `serve-webui` variants (production hidden, debug hidden) with different flag structures. The domain is over-represented for a single feature. |

**Total command count verification**: 6 + 3 + 1 + 1 + 1 + 1 + 2 + 3 = **18** ✅

---

## Domain Analysis

### Engine Domain (6 commands)

**CRUD coverage**:
- **Create**: ❌ No `add-engine` / `create-engine` / `install-engine`. Engines are added by packaging (snap components), not by CLI command.
- **Read**: ✅ `list-engines` (read-all), `show-engine` (read-one).
- **Update**: ✅ `use-engine` (activate), `debug select-engine` (test selection).
- **Delete**: ❌ No `remove-engine` / `uninstall-engine`. Engine components can be pruned via `prune-cache` but there's no dedicated engine removal.

**Verb consistency within domain**:
- `list`, `show`, `use` follow the DE013 pattern for secondary objects.
- `validate` and `select` are in debug namespace — appropriate as developer tools.
- `run` is the odd one out — it's hidden and engine-dependent but uses a verb unlike all other engine commands.

### Configuration Domain (3 commands)

**CRUD coverage**: Complete ✅ — `get` (read), `set` (create/update), `unset` (delete).
**Verb consistency**: Perfect. The `get`/`set`/`unset` triad follows DE013 exactly.
**Note**: No noun is needed — per DE013, config operates on the implicit configuration namespace of the tool.

### Web UI Domain (3 commands)

**Issue**: Three commands for what is essentially one feature:
1. `webui` — launches the browser to the web UI.
2. `serve-webui` (hidden, top-level) — serves the production web UI.
3. `debug serve-webui` (hidden, debug) — serves a debug web UI with explicit `--base-url`.

**Recommendation**: Consolidate. The two `serve-webui` variants could be unified with a `--debug` flag, or the debug variant could be removed since `debug chat` already provides a way to test with explicit URLs. The `webui` and `serve-webui` commands have different purposes (launch vs serve) — the naming does not make this distinction clear.

### Chat Domain (2 commands)

**Issue**: `chat` (top-level, auto-detects engine) and `debug chat` (requires `--base-url`). Both use the same verb in different namespaces with different contracts. This is confusing: a user who discovers `debug chat` might assume it behaves like `chat`.

### Singletons (3 domains, 3 commands)

`show-machine`, `status`, `version`, and `prune-cache` are each the sole command in their domain. This is acceptable for a CLI of this size — not every domain needs multiple commands. `show-machine` could conceptually fit under a machine management domain with commands like `probe-machine` or `benchmark-machine` if the CLI grows.

---

## Domain Naming Consistency

| Domain | Noun Form Used | Plural/Issues |
|--------|---------------|---------------|
| Engine | `engine` (singular in `show-engine`, `use-engine`), `engines` (plural in `list-engines`) | ✅ Correct: plural for listing, singular for operations on one engine |
| Configuration | Implicit (no noun) | ✅ DE013 convention |
| Machine | `machine` (singular) | ✅ Consistent |
| Cache | `cache` (singular) | ✅ Consistent |
| Chat | No noun (verb-as-command: `chat`) | ⚠️ Inconsistent with rest of CLI — should it be `chat` (verb) or have a noun? |
| Web UI | `webui` (noun-as-command) | ❌ Inconsistent — should be verb-noun (e.g., `open-webui`) |

---

## Key Findings

1. **Over-fragmentation in Web UI**: Three commands for one feature creates discoverability problems. Users won't know which to use.
2. **Missing engine lifecycle commands**: No way to add or remove engines via CLI. This is delegated to the snap packaging system, which is reasonable but limits the CLI's self-sufficiency.
3. **`run` is a domain anomaly**: It's the only engine-adjacent command that doesn't have "engine" in its name, yet it requires an active engine.
4. **Chat domain split**: Two chat commands in different namespaces with different semantics is confusing. Consider making the top-level `chat` support `--base-url` and `--model` flags to subsume `debug chat`.
5. **Imbalanced domain sizes**: Engine (6) and Web UI (3) dominate while most domains are singletons. This reflects the CLI's primary purpose (engine management + chat/web UI access).