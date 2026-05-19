# 02 — Verb Taxonomy and Aspect Classification

## Verb Classification (6 unique verbs)

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|
| chat | execution | atelic | no | — | `qwen36 chat` |
| completion | observation | punctual | no | — | `qwen36 completion bash` |
| get | observation | punctual | no | set | `qwen36 get http.port` |
| set | mutation | punctual | yes | get (read), unset (missing) | `qwen36 set http.port=8326` |
| show | observation | punctual | no | — | `qwen36 show-engine` |
| use | mutation | punctual | partial | — | `qwen36 use-engine cpu`, `qwen36 use-engine --auto` |

## Self-Check

- Verbs in Section 1: chat, completion, get, set, show, use → **6 verbs**
- Verbs in this table: chat, completion, get, set, show, use → **6 verbs**
- ✓ Match confirmed

## Analysis

### Intent Group Distribution

| Intent Group | Count | Verbs |
|---|---|---|
| observation | 3 | completion, get, show |
| mutation | 2 | set, use |
| execution | 1 | chat |
| lifecycle | 0 | — |
| access | 0 | — |
| transfer | 0 | — |
| migration | 0 | — |

**Observation**: The CLI is heavily observation-oriented (50% of verbs), reflecting its role as a configuration manager for an inference server rather than a lifecycle management tool.

### Reversibility Gaps

- **`set`** is logically reversible via `get` (read) + `set` (overwrite), but there's no `unset`/`reset` command to restore defaults.
- **`use`** (engine selection) is partially reversible — you can switch to a different engine, but there's no "previous engine" or undo mechanism.
- **`chat`** is not reversible (interactive session, no state change to undo).

### Aspect Distribution

All verbs except `chat` are **punctual** (instant state change or instant query). `chat` is the only **atelic** verb (ongoing interactive session). This is appropriate for a configuration-focused CLI.

### Missing Verb Patterns

Per DE013 standards, the following patterns are notably absent:

| Expected Pattern | Status | Notes |
|---|---|---|
| `list` / `<nouns>` | Missing | No way to list available engines or all config keys |
| `status` | Missing | No command to show overall system status (server running? engine selected? components installed?) |
| `enable` / `disable` | Missing | Could replace `use-engine` for toggling engine activation |
| `version` | Missing | No version command (snap version visible via `snap info qwen36` but not the CLI itself) |
| `help` | Unknown | May exist but undocumented |
