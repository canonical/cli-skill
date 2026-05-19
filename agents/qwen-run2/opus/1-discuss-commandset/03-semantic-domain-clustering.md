# Semantic Domain Clustering

## Domain Table

| Domain | Count | Commands | Naming Consistent? | Notes |
|--------|-------|----------|-------------------|-------|
| Engine | 2 | `use-engine`, `show-engine` | Yes — both use `-engine` suffix | Missing `list-engines` or `engines` (per DE013 plural-noun pattern for listing secondary objects). No add/remove since engines are snap components. |
| Configuration | 2 | `get`, `set` | Yes — standard get/set pair | Missing `unset` per DE013. No `list` to show all keys. Bare verbs without noun suffix because config is the primary operation domain. |
| Interaction | 1 | `chat` | N/A (single command) | The primary user-facing command. Could be expanded (e.g., `chat --image`, `chat --file`) but currently takes no arguments. |
| Shell Integration | 1 | `completion` | N/A (single command) | Noun-as-command pattern. Per DE013, this is acceptable for utility/meta commands. |

**Total commands: 6** (use-engine + show-engine + get + set + chat + completion = 6 ✓)

## Domain Analysis

### Engine Domain

- **CRUD coverage**: show (read) and use (update/select) only. No create, delete, or list.
- **Verb consistency**: `use` and `show` are semantically appropriate — `show` for observation, `use` for selection.
- **Gap**: No way to list available engines or their hardware requirements from the CLI. Users must know engine names a priori or rely on `--auto`.

### Configuration Domain

- **CRUD coverage**: get (read) and set (write). Missing unset (delete/restore) and list (enumerate).
- **Verb consistency**: Perfect — `get`/`set` is the DE013-standard pair.
- **Gap**: No way to enumerate valid keys. No way to restore a key to its default value. No way to distinguish package-level vs user-level settings in output.

### Interaction Domain

- **CRUD coverage**: N/A — single-purpose interactive command.
- **Verb consistency**: `chat` is clear and unambiguous.
- **Gap**: No non-interactive mode (e.g., `qwen36 chat --prompt "question"` for scripting). No way to specify model or parameters at invocation time.

### Shell Integration Domain

- **CRUD coverage**: N/A — utility command.
- **Verb consistency**: `completion` is a noun, which is a minor DE013 deviation but a widely-accepted pattern.
- **Gap**: Only `bash` is supported. No `zsh`, `fish`, or `powershell`.

## Cross-Domain Observations

1. **Small, focused command set**: 6 commands across 4 domains is minimal and learnable.
2. **Engine domain is the most complete**: 2 commands with clear verb differentiation.
3. **Configuration domain has the biggest functional gaps**: Missing `unset` and key enumeration.
4. **No overlap between domains**: Each command belongs to exactly one domain with no ambiguity.
