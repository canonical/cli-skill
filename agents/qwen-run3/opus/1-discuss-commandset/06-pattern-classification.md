# 06 — Pattern Classification and Recommendations

## Pattern Classification

| Dimension | Assessment |
|-----------|-----------|
| **Primary pattern** | Flat verb / verb-noun hierarchy |
| **Grouping** | Three logical groups (Basic, Configuration, Management) via cobra command groups |
| **Depth** | Single level (no subcommands visible to users) |
| **Style** | Mixed: bare verbs for common actions (`get`, `set`, `chat`), verb-noun for secondary objects (`show-engine`, `use-engine`, `prune-cache`) |
| **Scale** | 12 visible commands — compact and learnable |

## Discoverability Assessment

| User Goal | Predicted Path | Actual Command | Match? |
|-----------|---------------|----------------|--------|
| "What engine is running?" | `show-engine` or `status` | `show-engine` (active) or `status` (service state) | Partial — two commands serve overlapping needs |
| "Change the port" | `set port=8080` or `config port=8080` | `sudo qwen36 set http.port=8080` | ✗ — requires knowing the `http.` prefix and sudo requirement |
| "List what I can configure" | `get --help` or `config --list` | `qwen36 get` (dumps all values) | Partial — dumps values but not schema/types |
| "Start the server" | `start` or `serve` | `sudo snap start qwen36.server` | ✗ — not a CLI command; delegated to snap |
| "Clean up disk space" | `clean` or `remove` | `sudo qwen36 prune-cache` | ✗ — "prune-cache" is not intuitive for component removal |
| "Check system compatibility" | `check` or `requirements` | `qwen36 show-machine` + `qwen36 list-engines` | Partial — requires combining two commands |

## Ecosystem Comparison

| Feature | qwen36 | ollama | localai | Canonical standard |
|---------|--------|--------|---------|-------------------|
| Grammar | verb-noun flat | verb-noun flat | verb-noun flat | verb-noun (DE013) |
| Engine selection | `use-engine` | n/a (single backend) | n/a | — |
| Config | `get/set/unset` | `OLLAMA_*` env vars | YAML files | `get/set/unset` ✓ |
| Model management | snap components | `pull/push/rm` | download | — |
| Start interaction | `chat` | `run <model>` | API only | — |
| Server management | `snap start/stop` | built-in daemon | `run` | — |
| Output format flag | `--format` | `--format` (JSON only in API) | n/a | — |
| Root requirement | yes (write ops) | no (user-level) | no | — |

### Key Ecosystem Differences

1. **ollama** merges model selection into the interaction command (`ollama run llama3`) — simpler mental model but less separation of concerns.
2. **qwen36** delegates server lifecycle to snap — correct for snap architecture but makes "start/stop" invisible in the CLI's own help.
3. **qwen36**'s `get/set/unset` precisely follows DE013 — stronger Canonical compliance than most tools.

## Recommendations

### 1. Add `open-` prefix to `webui` (Low cost, High clarity)

**Current**: `qwen36 webui`
**Proposed**: `qwen36 open-webui` (with `webui` as deprecated alias)

**Rationale**: Per DE013 §Grammar, "Commands are verbs." The noun-only `webui` breaks this rule. The verb `open` makes the action explicit and matches `xdg-open` semantics.

**Standard citation**: DE013 §Grammar: "Every command that acts on a primary object must be a verb."

**Deprecation plan** (per deprecation spec):
- Minor version N+1: Add `open-webui`, keep `webui` as alias with warning: `warning: "webui" is deprecated, use "open-webui" instead`
- Major version N+1: Remove `webui` alias with error: `error: "webui" was removed, use "open-webui" instead`

**Backward compatibility**: Low impact — `webui` was recently added and likely has few script dependents.

---

### 2. Rename `prune-cache` to `remove-components` or add clarifying help (Medium cost, Medium clarity)

**Current**: `qwen36 prune-cache`
**Proposed**: Either rename to `remove-components` or keep name but improve help text and add `--list` flag

**Rationale**: "Cache" implies regenerable temporary data. The command actually removes installed snap components (engine binaries, model files) that are 22+ GB and must be re-downloaded. The noun is misleading.

**Standard citation**: DE013 §Grammar: verb choice "needs to imply or trigger recall of the object type it refers to."

**Alternative**: Keep `prune-cache` if the team's mental model treats inactive engine components as "cached" (stored but not actively needed). In this case, add explicit help text: "Removes snap components for inactive engines. Removed components must be re-downloaded from the store."

**Deprecation plan** (if renaming):
- Minor version: Add `remove-components`, keep `prune-cache` with warning
- Major version: Remove `prune-cache`

---

### 3. Add `--dry-run` to `use-engine` and `prune-cache` (Low cost, High safety)

**Current**: No preview mode for destructive operations.
**Proposed**: `--dry-run` flag that shows what would happen without executing.

**Rationale**: `use-engine --auto` may download gigabytes of components. `prune-cache` removes them irreversibly. Users should be able to preview these actions.

**Standard citation**: DE013 §Safety: "Are destructive operations gated (confirmation, dry-run, force semantics)?"

**Backward compatibility**: Additive change — no breaking impact.

---

### 4. Standardize `--format` flag across all observation commands (Low cost, High consistency)

**Current**: `list-engines` supports `table`/`json`; others support `json`/`yaml`; `get` has no format flag.
**Proposed**: All observation commands support `table`, `json`, `yaml` where applicable.

**Rationale**: Inconsistent format options force users to remember per-command differences.

**Standard citation**: DE013 §Consistency: "Are naming, formatting, and status messages consistent across commands?"

**Backward compatibility**: Additive — existing formats continue to work, new formats are additions.

---

### 5. Document server lifecycle integration (Low cost, High discoverability)

**Current**: `--help` shows no information about starting/stopping the server.
**Proposed**: Add a "Service Management" section to root `--help` (already partially implemented via `common.SuggestServiceManagement()`).

**Rationale**: Users of the CLI expect it to be self-contained. Delegating to `snap start/stop` without clear documentation creates a discoverability gap.

---

## Compliance Self-Check

| Recommendation | Cites Standard? | Deprecation Plan? | Existing Violation Flagged? |
|---------------|----------------|-------------------|---------------------------|
| 1. `open-webui` | ✓ DE013 §Grammar | ✓ | ✓ `webui` is noun-only |
| 2. `prune-cache` | ✓ DE013 §Grammar | ✓ | ✓ misleading noun |
| 3. `--dry-run` | ✓ DE013 §Safety | n/a (additive) | ✓ no preview mode |
| 4. `--format` | ✓ DE013 §Consistency | n/a (additive) | ✓ inconsistent formats |
| 5. Service docs | ✓ DE013 §Discoverability | n/a (docs only) | ✓ gap flagged |
