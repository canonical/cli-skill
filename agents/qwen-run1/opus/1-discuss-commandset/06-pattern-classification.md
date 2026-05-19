# 06 — Pattern Classification and Recommendations

## Pattern Classification

### Primary Grouping Pattern: **Flat verb-noun**

The `qwen36` CLI uses a flat command structure with no nesting:

```
qwen36 <command> [args] [flags]
```

Commands use a mix of:
- **Bare verbs**: `chat`, `get`, `set`
- **Verb-noun compounds**: `use-engine`, `show-engine`
- **Noun-like utility**: `completion`

### Depth: 1 level

All commands are at the top level. No subcommand nesting exists.

### Style: Configuration manager with embedded action

The CLI is primarily a configuration management wrapper around snap options, with one "action" command (`chat`) and one "utility" command (`completion`).

## Discoverability Assessment

### Predicted User Paths vs Reality

| User Intent | Predicted Command | Actual Command | Match? |
|---|---|---|---|
| "Start chatting with the model" | `qwen36 chat` | `qwen36 chat` | ✓ |
| "Which engine am I using?" | `qwen36 engine` or `qwen36 status` | `qwen36 show-engine` | Partial |
| "Switch to GPU" | `qwen36 engine cuda` or `qwen36 set engine cuda` | `qwen36 use-engine cuda` | Partial |
| "What port is the server on?" | `qwen36 config` or `qwen36 show port` | `qwen36 get http.port` | ✓ (if user knows the key) |
| "List all settings" | `qwen36 config` or `qwen36 get` (no arg) | *(not available)* | ✗ |
| "What engines are available?" | `qwen36 engines` or `qwen36 list-engines` | *(not available)* | ✗ |
| "Is the server running?" | `qwen36 status` | *(not available — must use `snap logs`)* | ✗ |
| "What version?" | `qwen36 version` or `qwen36 --version` | *(not available)* | ✗ |

### Discoverability Score: 4/8 (50%)

Half of common user intents map directly to existing commands. The other half require external knowledge (snap commands, config key names) or don't have CLI support.

## Ecosystem Comparison

### Compared Tools

| Tool | Domain | Commands | Pattern |
|---|---|---|---|
| `ollama` | LLM inference | `run`, `pull`, `list`, `show`, `serve`, `create`, `rm`, `cp` | Flat verbs, docker-inspired |
| `multipass` | VMs | `launch`, `list`, `info`, `start`, `stop`, `delete`, `shell` | Flat verbs, lifecycle-focused |
| `snap` | Package mgmt | `install`, `remove`, `list`, `info`, `get`, `set`, `services` | Flat verbs, DE013 compliant |

### Comparison Analysis

| Aspect | qwen36 | ollama | multipass | snap |
|---|---|---|---|---|
| List available items | ✗ | `ollama list` | `multipass list` | `snap list` |
| Show item details | `show-engine` | `ollama show` | `multipass info` | `snap info` |
| Status/health | ✗ | implicit (via `list`) | `multipass list` shows state | `snap services` |
| Config get/set | ✓ | ✗ (env vars only) | `multipass get/set` | `snap get/set` |
| Start action | `chat` | `ollama run` | `multipass shell` | N/A |
| Version | ✗ | `ollama --version` | `multipass version` | `snap version` |

**Key insight**: `qwen36` closely follows `snap get/set` patterns (appropriate since it's a snap), but lacks `list` and `status` commands that all comparison tools provide.

## Recommendations

### 1. Add `engines` command (list available engines)

**Rationale**: Per DE013, "use `foobars` as a shorthand for listing information about all instances of a specific type of secondary object." Users need to discover available engines before selecting one.

**Proposed**: `qwen36 engines` → lists available engines with name, description, and hardware match status.

**Backward compat**: Additive (minor version). No breaking change.

**Migration cost**: None — new command only.

### 2. Add `status` command

**Rationale**: Per DE013, "use the shorthand `status` over `show-status`." Users need a single command to understand system health: engine selected, server running, components installed.

**Proposed**: `qwen36 status` → shows engine, server state, component readiness.

**Backward compat**: Additive (minor version). No breaking change.

**Migration cost**: None — new command only.

### 3. Add `unset` command for configuration

**Rationale**: Per DE013, the standard config triple is `get/set/unset`. Without `unset`, users cannot restore defaults.

**Proposed**: `qwen36 unset <key>` → removes user-level override, restoring package default.

**Backward compat**: Additive (minor version). No breaking change.

**Migration cost**: None — new command only.

### 4. Rename `model-name` config key to `api-model-name`

**Rationale**: Confusion-pair audit identified `model-name` vs `model` as high-risk naming collision.

**Proposed**: Rename to `api-model-name`. Per deprecation spec: keep `model-name` working as alias for one cycle, emit deprecation warning, remove in next major.

**Backward compat**: Breaking change if done without transition. With alias + deprecation warning: minor version acceptable.

**Migration cost**: Medium — scripts using `get model-name` need updating within the deprecation window.

**Deprecation process** (per deprecation/README.md):
- Minor version N+1: Add `api-model-name`, make `model-name` an alias that prints a deprecation warning
- Major version N+1: Remove `model-name` alias with error message pointing to `api-model-name`
- Major version N+2: Remove error message

### 5. Consider renaming `use-engine` to `set-engine` or `switch-engine`

**Rationale**: Per DE013, `use` is not in the standard verb vocabulary. The operation is functionally a `set` (configuration mutation) or `switch` (toggling between options). `snap` itself doesn't have a `use` verb.

**Proposed**: `switch-engine` (clearer intent than `set-engine` which could confuse with `set`).

**Backward compat**: Breaking change. Per deprecation spec:
- Minor version: Add `switch-engine` as alias, deprecation warning on `use-engine`
- Major version: Remove `use-engine`

**Migration cost**: Medium — `use-engine` is used in install hooks and documentation. All references need updating.

**Alternative**: Keep `use-engine` as-is — the command set is small and `use` is intuitive even if non-standard. The cost of change may exceed the benefit per the deprecation spec's stability guidance.

### 6. Add `version` command or `--version` flag

**Rationale**: Per DE013 commonly used commands table, `tool version` should show the release version. All comparison tools support this.

**Proposed**: `qwen36 version` → outputs snap version.

**Backward compat**: Additive (minor version). No breaking change.

**Migration cost**: None.

## Summary

| # | Recommendation | Priority | Type | Breaking? |
|---|---|---|---|---|
| 1 | Add `engines` command | High | Addition | No |
| 2 | Add `status` command | High | Addition | No |
| 3 | Add `unset` command | Medium | Addition | No |
| 4 | Rename `model-name` key | Medium | Rename (with deprecation) | Yes (managed) |
| 5 | Consider `switch-engine` | Low | Rename (optional) | Yes (managed) |
| 6 | Add `version` command | Low | Addition | No |
