# 01 — Verb-Noun Decomposition Matrix

## Decomposition Grid

Each command is decomposed into a **verb** (the action) and a **noun** (the resource or object acted upon).

**Verbs** (rows, sorted alphabetically) × **Nouns** (columns, sorted by frequency)

| Verb        | config | engine | machine | cache | version | status | ui | chat | model | subprocess |
|-------------|--------|--------|---------|-------|---------|--------|-----|------|-------|------------|
| chat        | —      | —      | —       | —     | —       | —      | —   | ✓    | —     | —          |
| get         | ✓      | —      | —       | —     | —       | —      | —   | —    | —     | —          |
| list        | —      | ✓      | —       | —     | —       | —      | —   | —    | —     | —          |
| prune       | —      | —      | —       | ✓     | —       | —      | —   | —    | —     | —          |
| run         | —      | —      | —       | —     | —       | —      | —   | —    | —     | ✓          |
| serve       | —      | —      | —       | —     | —       | —      | ✓   | —    | —     | —          |
| set         | ✓      | —      | —       | —     | —       | —      | —   | —    | —     | —          |
| show        | —      | ✓      | ✓       | —     | —       | —      | —   | —    | —     | —          |
| status      | —      | —      | —       | —     | —       | ✓      | —   | —    | —     | —          |
| unset       | ✓      | —      | —       | —     | —     | —      | —   | —    | —     | —          |
| use         | —      | ✓      | —       | —     | —       | —      | —   | —    | —     | —          |
| validate    | —      | ✓      | —       | —     | —       | —      | —   | —    | —     | —          |
| webui       | —      | —      | —       | —     | —       | —      | —   | —    | —     | —          |
| version     | —      | —      | —       | —     | ✓       | —      | —   | —    | —     | —          |
| select      | —      | ✓      | —       | —     | —       | —      | —   | —    | —     | —          |

---

## Annotations

### Incomplete CRUD Sets

- **config**: Has `get` (read), `set` (create/update), `unset` (delete). **CRUD complete** ✅.
- **engine**: Has `list` (read-all), `show` (read-one), `use` (select/activate), `validate` (check), `select` (test-select). **Missing**: `remove-engine` (deactivate/uninstall engine). Engine lifecycle has no explicit deactivation command — `use-engine` auto-clears the previous engine, but there is no standalone "deactivate the current engine" command.
- **cache**: Has `prune` (delete). **Missing**: `show-cache` (list cache contents), no read counterpart.
- **machine**: Has `show` (read). **Missing**: No mutation commands (machine info is inherently read-only, so this is acceptable).
- **chat**: Has two chat commands: `chat` (uses active engine) and `debug chat` (explicit URL). Inconsistent namespace.
- **ui**: Has `webui` (launch browser) and two `serve-webui` commands (top-level hidden + debug). Unclear distinction between production and debug modes.

### Verb Inconsistencies

- **`use-engine`** uses `use` as the verb for selecting/activating an engine. The standard verb for this concept would typically be `select` (used in `debug select-engine`!), `activate`, or `set`. The `use` verb is informal.
- **`prune-cache`** uses `prune` for deletion. The standard verb would be `remove`, `delete`, or `clean`. `prune` is a horticultural metaphor that is uncommon in CLI tools.
- **`show-machine`** and **`show-engine`** use `show` consistently, but `list-engines` uses `list` instead of `show-engines`. Per DE013, `foobars` (plural noun) is preferred over `list-foobar` for secondary objects — but `engines` alone would conflict with the engine directory convention.

### Orphan Commands

Commands that do not decompose cleanly into verb-noun:

| Command | Reason |
|---------|--------|
| `status` | No clear noun — `status` is both verb and noun. It shows the state of the entire system (engine + services + endpoints + model). Per DE013, `status` is the correct shorthand for showing current state without a noun. |
| `version` | No clear noun — `version` is a noun used as a command. Per DE013, `version` is a standard command name. |
| `webui` | No clear verb-noun decomposition. It is a noun used as a command meaning "launch the web UI." Per DE013, this would be better as something like `open-webui` or `launch-webui`. |
| `chat` | No clear noun — `chat` is both verb and noun. It means "start a chat session." |

### Verb Count Summary

| Verb | Occurrences | Commands |
|------|-------------|----------|
| `show` | 2 | `show-engine`, `show-machine` |
| `chat` | 2 | `chat`, `debug chat` |
| `serve` | 2 | `serve-webui`, `debug serve-webui` |
| `get` | 1 | `get` |
| `set` | 1 | `set` |
| `unset` | 1 | `unset` |
| `list` | 1 | `list-engines` |
| `use` | 1 | `use-engine` |
| `prune` | 1 | `prune-cache` |
| `run` | 1 | `run` |
| `validate` | 1 | `debug validate-engines` |
| `select` | 1 | `debug select-engine` |
| `status` | 1 | `status` |
| `version` | 1 | `version` |
| `webui` | 1 | `webui` |

**15 unique verbs for 18 commands** — the verb set is moderately fragmented. Five verbs appear only once, which is expected for a CLI of this size. The duplication comes from the `debug` namespace (chat, serve-webui) mirroring top-level commands, and from `show` being used for two resource types.
