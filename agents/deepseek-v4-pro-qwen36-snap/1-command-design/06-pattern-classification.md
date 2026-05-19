# 06 — Pattern Classification and Recommendations

## Pattern Classification

### Primary Grouping Pattern: Flat with Cosmetic Groups

The `qwen36` CLI uses a **flat command hierarchy** with cosmetic grouping for help display only. Commands are registered into Cobra `Group` objects (`basic`, `config`, `engine`) purely to organize `--help` output — there is no structural nesting, no separate command namespaces, and no subcommand routing. Ungrouped commands appear under "Additional Commands" in the help footer.

**Structural characteristics**:
- **Depth**: 1 level for most commands (root → command). Exception: `debug` is a true subcommand group with depth 2 (root → debug → subcommand).
- **Style**: **Verb-noun** for most commands (`show-engine`, `list-engines`, `use-engine`, `prune-cache`, `serve-webui`). Exceptions: `status`, `version`, `chat`, `webui`, `get`, `set`, `unset` use standalone verb or noun forms.
- **Command sorting**: Disabled (`cobra.EnableCommandSorting = false`) — commands appear in registration order.

### Command Count: 18 total (12 visible top-level, 2 hidden top-level, 4 debug subcommands)

### DE013 Compliance Summary

Per DE013 §Grammar, commands must be verbs. Assessment:

| DE013 Requirement | Compliance | Notes |
|------------------|------------|-------|
| Commands are verbs | ⚠️ Partial | `webui` is a noun used as a command (should be `open-webui` or `launch-webui`). `status` and `version` are standard exceptions. `chat` is both verb and noun — acceptable. |
| Verb-noun for ambiguous objects | ✅ Good | `show-engine`, `show-machine`, `list-engines`, `use-engine`, `prune-cache` all use verb-noun. |
| `foobars` plural for listing | ⚠️ Partial | `list-engines` uses `list-` prefix instead of bare `engines`. Per DE013, plural noun alone is preferred (`snap services` not `snap list-services`). But `list-engines` is clearer for discoverability. |
| `status` shorthand | ✅ Compliant | Uses `status` not `show-status`. |
| `version` command | ✅ Compliant | Standard command. |
| `get`/`set`/`unset` for config | ✅ Compliant | Standard triad per DE013. |
| At most one sublevel | ✅ Compliant | Only `debug` has subcommands; the rest are top-level. |
| No deeply nested compounds | ✅ Compliant | No verb-noun-noun compounds. |

---

## Discoverability Assessment

### Predicted User Paths vs Actual Command Locations

| User Goal | Intuitive Command | Actual Command | Discoverability Issue |
|-----------|------------------|----------------|----------------------|
| "What's my system status?" | `status` | `status` ✅ | — |
| "What engines are available?" | `engines`, `list engines` | `list-engines` | Plural noun `engines` is not a command. User must know to prefix with `list-`. |
| "Tell me about the cpu engine" | `show-engine cpu`, `info engine cpu` | `show-engine cpu` ✅ | — |
| "Switch to cuda engine" | `use-engine cuda`, `switch-engine cuda`, `set-engine cuda` | `use-engine cuda` | `use` is slightly informal; `switch` or `select` might be more discoverable. |
| "Change the port" | `set http.port=9000`, `config set http.port=9000` | `set http.port=9000` ✅ | — |
| "Clean up disk space" | `clean`, `prune`, `remove-cache` | `prune-cache` | `prune` is uncommon. Users might try `clean` or `remove` first. |
| "Chat with the model" | `chat` | `chat` ✅ | — |
| "Open the web interface" | `webui`, `ui`, `open`, `dashboard` | `webui` | Noun-as-command. Users might try `open`, `launch`, or `ui`. |
| "What version?" | `version`, `--version` | `version` ✅ | — |
| "Show my hardware" | `show-machine`, `info`, `hardware` | `show-machine` ✅ | — |
| "Validate an engine file" | `validate engine.yaml` | `debug validate-engines engine.yaml` | Buried under `debug` namespace. Users may not discover it. |
| "Test which engine I'd get" | `select-engine`, `test-engine` | `debug select-engine` | Buried under `debug`. |

**Discovery score**: 7/12 user goals map to intuitive commands. 5/12 require knowledge of naming conventions or `debug` namespace.

### Help Text Discoverability

- `--help` at top level shows grouped commands: Basic, Configuration, Management, Additional.
- Hidden commands (`run`, `serve-webui`, `debug`) are not shown.
- No `--help` output shows the `debug` subcommands until the user runs `qwen36 debug --help`.
- Tab completion is available but not documented.

---

## Ecosystem Comparison

### Compared to: `snap` (snapd CLI)

| Aspect | `snap` | `qwen36` |
|--------|--------|----------|
| Hierarchy | Flat with some subcommands (`snap changes`, `snap connections`) | Flat with one subcommand group (`debug`) |
| Verb usage | Mixed: `install`, `remove`, `list`, `info`, `enable`, `disable` | Verb-noun: `use-engine`, `list-engines`, `show-engine` |
| Config commands | `snap get/set/unset <snap> <key>` | `qwen36 get/set/unset <key>` |
| Output formats | Table (default), no `--format` flag | YAML/JSON/table with `--format` |
| Hidden commands | None | `run`, `serve-webui`, `debug` |

**Alignment**: Good. The `get`/`set`/`unset` pattern matches `snap`. Verb-noun commands align with `snap` conventions like `snap connections`.

### Compared to: `docker` CLI

| Aspect | `docker` | `qwen36` |
|--------|----------|----------|
| Hierarchy | Management commands (`docker container`, `docker image`) + aliases (`docker run` = `docker container run`) | Flat, no management command groups |
| Naming | Object-verb format for management: `docker container ls` | Verb-object: `list-engines` |
| Hidden commands | None | 3 hidden commands |

**Alignment**: Different paradigm. Docker uses noun-verb hierarchy; qwen36 uses flat verb-noun. Docker's approach scales better to many object types; qwen36's approach is simpler for ~5 resource types.

### Compared to: `systemctl`

| Aspect | `systemctl` | `qwen36` |
|--------|-------------|----------|
| Hierarchy | Flat, verb-noun: `systemctl start <unit>`, `systemctl status <unit>` | Flat, verb-noun |
| Status | `status` shows unit state | `status` shows engine + services + endpoints |
| Config | No config command (uses unit files) | `get`/`set`/`unset` |

**Alignment**: Good. Flat verb-first pattern is similar. `status` semantics differ (systemctl's `status` is per-unit; qwen36's is system-wide).

---

## Recommendations

Ordered by impact, each annotated with backward compatibility and migration cost.

### 1. Consolidate `chat` and `debug chat` into a single `chat` command

**Current**: Top-level `chat` (auto-detects engine) and `debug chat` (requires `--base-url`).

**Proposed**: Add `--base-url` and `--model` flags to top-level `chat`. If `--base-url` is provided, skip engine auto-detection. Deprecate `debug chat`, redirecting to `chat --base-url <url>`.

**Per DE013 §Grammar and deprecation spec**:
- Minor version: Add flags to `chat`, mark `debug chat` as deprecated with warning: `"debug chat" is deprecated, use "chat --base-url <url>" instead`.
- Major version: Remove `debug chat`, return error pointing to `chat --base-url`.

**Backward compat**: High — `debug chat` users must change their invocations. Low-impact since `debug chat` is hidden.
**Migration cost**: Low — single flag addition, deprecation warning, eventual removal.

### 2. Consolidate `serve-webui` and `debug serve-webui`

**Current**: Two `serve-webui` commands in different namespaces with different flag defaults.

**Proposed**: Unify into a single `serve-webui` (keep hidden). The top-level version already accepts `--port`, `--host`, `--capabilities`. Add `--base-url` flag to override auto-detection (matching the consolidated `chat` pattern). Remove `debug serve-webui`.

**Backward compat**: Low — both commands are hidden. `debug serve-webui` users must use `serve-webui --base-url ...`.
**Migration cost**: Low.

### 3. Rename `webui` to `open-webui` or `launch-webui`

**Current**: `webui` is a noun used as a command.

**Per DE013 §Grammar**: Commands must be verbs. `webui` is not a verb.

**Proposed**: Rename to `open-webui` (or `launch-webui`). Per deprecation spec, keep `webui` as an alias in the minor version with deprecation warning: `"webui" is deprecated, use "open-webui" instead`.

**Backward compat**: Medium — `webui` is a visible command mentioned in README. Scripts may reference it.
**Migration cost**: Low — alias + deprecation warning. The old name `webui` is short and memorable; `open-webui` is clearer about intent.
**Alternative**: `launch-webui` (more specific but longer).

### 4. Add `deactivate-engine` command

**Current**: No way to deactivate an engine without switching to another one or manually pruning.

**Proposed**: Add `deactivate-engine` that unsets the active engine, clears engine configs, and optionally prunes components.

**Backward compat**: None — new command.
**Migration cost**: Low — new command, no existing behavior changed.
**Design note**: Per DE013, use verb-noun form. `deactivate-engine` fits the pattern. Alternatives: `unuse-engine` (follows `use` verb), `disable-engine`, `unset-engine`. Recommended: `deactivate-engine` — it's unambiguous and follows the `activate`/`deactivate` pair convention.

### 5. Add `--no-headers` flag to `list-engines` and uppercase table headers

**Current**: Table headers are lowercase, no `--no-headers` flag.

**Per DE013 §Tabular Data**: Column headers should be UPPERCASE and bold. `--no-headers` should be supported.

**Proposed**: Change headers to `ENGINE`, `VENDOR`, `DESCRIPTION`, `COMPAT`. Add `--no-headers` flag for script-friendly output.

**Backward compat**: Low — table rendering change only. Scripts that parse table output should use `--no-headers`.
**Migration cost**: Low.

### 6. Add `Example` field to all visible commands

**Current**: Only `run` (hidden) has examples.

**Per DE013**: Help text should include examples.

**Proposed**: Add `Example` to every command's Cobra definition. Examples should cover common use cases and edge cases.

**Backward compat**: None — additive only.
**Migration cost**: Low — documentation work only.

### 7. Normalize `--format` defaults

**Current**: `list-engines` defaults to `table`; all others default to `yaml`.

**Proposed**: Either make `list-engines` default to `yaml` for consistency, or document the rationale for the `table` default. If keeping `table` default, ensure it's clearly labeled as "human-readable only" in help text.

**Backward compat**: Medium — changing `list-engines` default would affect scripts and user expectations. **Recommendation**: Keep `table` default but add `--format` documentation explaining the difference.
**Migration cost**: None (keep current defaults, improve docs).

### 8. Move `debug validate-engines` to top level as `validate-engine`

**Current**: Engine validation is hidden under `debug`.

**Proposed**: Add `validate-engine <manifest>` as a visible top-level command (or keep it hidden but accessible directly). This is useful for engine developers and doesn't require root.

**Backward compat**: None — new command.
**Migration cost**: Low.

### 9. Support `unset <key>...` (multiple keys)

**Current**: `unset` accepts exactly one key; `set` accepts multiple key=value pairs.

**Proposed**: Allow `unset` to accept multiple keys for symmetry with `set`.

**Backward compat**: None — additive. Single-key `unset` still works.
**Migration cost**: Low.

---

## Recommendations Summary Table

| # | Recommendation | Type | Backward Compat | Migration Cost | Priority |
|---|---------------|------|-----------------|---------------|----------|
| 1 | Consolidate `chat` + `debug chat` | Consolidation | High (hidden cmd) | Low | **High** |
| 2 | Consolidate `serve-webui` + `debug serve-webui` | Consolidation | Low (both hidden) | Low | **High** |
| 3 | Rename `webui` → `open-webui` | Rename | Medium | Low (alias) | **Medium** |
| 4 | Add `deactivate-engine` | Addition | None | Low | **Medium** |
| 5 | Uppercase table headers + `--no-headers` | Standardization | Low | Low | **Low** |
| 6 | Add `Example` to all commands | Documentation | None | Low | **High** |
| 7 | Document `--format` defaults | Documentation | None | None | **Low** |
| 8 | Expose `validate-engine` at top level | Addition | None | Low | **Low** |
| 9 | Multi-key `unset` | Enhancement | None | Low | **Low** |

---

## Tradeoffs

### Recommendation 1 & 2 (Consolidation)

**Pro**: Reduces command count, eliminates confusion pairs, simplifies maintenance.
**Con**: Breaks existing workflows for users of `debug chat` and `debug serve-webui` (both hidden, so impact minimal). Requires careful flag design to avoid ambiguity between auto-detected and explicit `--base-url` modes.

### Recommendation 3 (Rename `webui`)

**Pro**: Complies with DE013 verb-noun standard, improves discoverability (user knows it "opens" something).
**Con**: Short memorable name `webui` is lost. Backward compat requires alias support for at least one major version. `open-webui` is longer to type. README and external docs must be updated.

### Recommendation 4 (Add `deactivate-engine`)

**Pro**: Completes engine lifecycle management. Users can cleanly deactivate without side effects.
**Con**: Adds another command to an already engine-heavy CLI. The use case is narrow — most users switch engines rather than deactivate entirely. May be YAGNI.

### Recommendation 5 (Table headers)

**Pro**: Aligns with DE013 and `snap list` conventions.
**Con**: Users who have learned the lowercase headers will need to adjust. Breaking change for any scripts that parse table headers (though those scripts should switch to `--format=json`).
