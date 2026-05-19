# Pattern Classification and Recommendations

## Pattern Classification

| Aspect | Assessment |
|--------|-----------|
| **Primary grouping pattern** | Flat with verb-noun compound commands |
| **Depth** | Single level only (no subcommand nesting beyond `completion bash`) |
| **Style** | Mixed: verb-noun for engine commands (`use-engine`, `show-engine`), bare verbs for config (`get`, `set`), bare noun for utility (`completion`), bare verb for interaction (`chat`) |
| **Command count** | 6 user-facing commands |
| **Grammar compliance** | Partially DE013-compliant; `get`/`set` and `show-engine` follow the standard; `use-engine` and `completion` deviate slightly |

## Discoverability Assessment

| User Intent | Predicted Path | Actual Command | Discoverability |
|-------------|---------------|----------------|-----------------|
| "Start talking to the model" | `qwen36 chat` | `qwen36 chat` | Excellent — intuitive |
| "Which engine is active?" | `qwen36 show-engine` or `qwen36 status` | `qwen36 show-engine` | Good — `show-engine` is findable |
| "Switch to GPU" | `qwen36 use-engine cuda` | `qwen36 use-engine cuda` | Good — verb-noun is clear |
| "What port is the server on?" | `qwen36 get port` or `qwen36 config` | `qwen36 get http.port` | Medium — need to know the key name |
| "List available engines" | `qwen36 engines` or `qwen36 list-engines` | *(not available)* | Poor — no such command exists |
| "See all configuration" | `qwen36 config` or `qwen36 get --all` | *(not available)* | Poor — no enumeration |
| "What version is installed?" | `qwen36 version` or `qwen36 --version` | *(not available; use `snap info qwen36`)* | Poor — missing standard command |
| "Reset to defaults" | `qwen36 unset http.port` | *(not available)* | Poor — no unset command |
| "Enable completions" | `qwen36 completion bash` | `qwen36 completion bash` | Medium — need to know the command exists |

## Ecosystem Comparison

| Feature | qwen36 | ollama | localai | llm (Simon Willison) |
|---------|--------|--------|---------|---------------------|
| Chat command | `chat` | `run` | N/A (API only) | `chat` |
| Engine/model selection | `use-engine` | `pull`/`run <model>` | config file | `models` |
| Config get/set | `get`/`set` | Environment vars | config file | `keys` |
| List models/engines | *(missing)* | `list` | API endpoint | `models list` |
| Version command | *(missing)* | `--version` | `--version` | `--version` |
| Server management | daemon (snap) | `serve` | `run` | embedded |
| Completions | `completion bash` | shell integration | N/A | `install` |
| Help | `--help` (presumed) | `--help` | `--help` | `--help` |

### Key differences from ecosystem

1. **No `list` / `models` command**: Every comparable tool provides a way to enumerate available options.
2. **No `version`**: Standard in all comparable tools.
3. **Engine-centric rather than model-centric**: Other tools let users pick models; qwen36 has a fixed model with selectable engines.
4. **Configuration via snap**: Unusual — most tools use config files or env vars.

## Recommendations

### 1. Add `version` command (Low cost, High value)

**Rationale**: Per DE013 §Commonly Used Commands, `tool version` is a standard command. Every comparable CLI provides `--version` or a `version` subcommand.

**Deprecation impact**: None — purely additive.

**Implementation**: Add `version` subcommand to Go binary that prints the snap version string.

### 2. Add `unset` command (Low cost, Medium value)

**Rationale**: Per DE013 §Commonly Used Commands, `get/set/unset` is the standard triplet for configuration management. The current CLI has `get` and `set` but no way to restore defaults.

**Deprecation impact**: None — purely additive.

**Implementation**: `qwen36 unset <key>` wraps `snapctl unset <key>` to restore the package-level default.

### 3. Add `engines` command (Low cost, High value)

**Rationale**: Per DE013 §Commonly Used Commands, use `foobars` as a shorthand for listing secondary objects. Users currently have no way to discover available engines or check hardware requirements without reading source files.

**Deprecation impact**: None — purely additive.

**Implementation**: `qwen36 engines` lists available engines with name, description, and compatibility status. This is the DE013 plural-noun pattern for secondary object listing.

### 4. Rename `completion` to follow verb pattern (Medium cost, Low value)

**Rationale**: Per DE013 §Grammar, commands must be verbs. `completion` is a noun. However, this pattern is widespread in the CLI ecosystem (kubectl, gh, docker all use `completion`).

**Recommendation**: Accept the deviation. The ecosystem precedent is strong enough that changing to a verb form (e.g., `generate-completion`) would reduce discoverability.

**Deprecation impact**: If renamed, requires deprecation notice for ≥1 cycle per the deprecation spec. Not recommended.

### 5. Consider `status` command (Medium cost, Medium value)

**Rationale**: Per DE013 §Commonly Used Commands, `tool status` shows current tool state. A `qwen36 status` command could consolidate: current engine, server health, port, model loaded, component installation status.

**Deprecation impact**: None — purely additive.

**Implementation**: `qwen36 status` outputs a human-readable summary of system state.

### 6. Document `--help` completeness (Zero cost, High value)

**Rationale**: Every command's `--help` should explain usage, arguments, examples, and edge cases per DE013 §Feedback. Currently undocumented whether this is the case (Go CLI submodule not checked out).

**Action**: Audit `--help` output for all commands once CLI submodule is available.

## Recommendation Priority

| # | Recommendation | Backward Compat | Migration Cost | Priority |
|---|---------------|-----------------|----------------|----------|
| 1 | Add `version` | No break | Zero | P0 |
| 2 | Add `unset` | No break | Zero | P1 |
| 3 | Add `engines` | No break | Zero | P1 |
| 4 | Keep `completion` as-is | N/A | N/A | — |
| 5 | Add `status` | No break | Low | P2 |
| 6 | Audit `--help` | No break | Low | P1 |

## Tradeoffs

| Recommendation | Benefit | Cost | Risk |
|---------------|---------|------|------|
| Add `version` | Standard compliance, user convenience | Minimal code change | None |
| Add `unset` | Complete get/set/unset triplet, safer config management | Small code addition | Low — needs to handle "already default" gracefully |
| Add `engines` | Discoverability, self-documenting hardware requirements | Medium code (format engine list, check hardware) | Low — read-only command |
| Add `status` | Single command for troubleshooting | Medium code (aggregate multiple data sources) | Low — may need updating as features change |
