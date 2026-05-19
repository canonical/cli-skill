# Verb Taxonomy and Aspect Classification

## Verb Classification Table

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|------|-------------|--------|------------|-------------|--------------|
| chat | execution | atelic | N/A | — | `qwen36 chat` |
| use | mutation | punctual | yes | — (set back to previous engine) | `qwen36 use-engine cpu`, `qwen36 use-engine cuda`, `qwen36 use-engine --auto` |
| show | observation | punctual | N/A | — | `qwen36 show-engine` |
| get | observation | punctual | N/A | set | `qwen36 get http.port`, `qwen36 get verbose` |
| set | mutation | punctual | yes | get (read), unset (restore, missing) | `qwen36 set http.port=8326`, `qwen36 set verbose=true` |
| completion | observation | punctual | N/A | — | `qwen36 completion bash` |

## Analysis

### Intent Group Distribution

| Intent Group | Count | Verbs | Coverage |
|-------------|-------|-------|----------|
| observation | 3 | show, get, completion | 50% of commands |
| mutation | 2 | use, set | 33% of commands |
| execution | 1 | chat | 17% of commands |
| lifecycle | 0 | — | 0% |
| access | 0 | — | 0% |
| transfer | 0 | — | 0% |
| migration | 0 | — | 0% |

### Aspect Distribution

| Aspect | Count | Verbs |
|--------|-------|-------|
| punctual | 5 | use, show, get, set, completion |
| atelic | 1 | chat |
| telic | 0 | — |

### Observations

1. **Heavily observational**: 50% of commands are read-only observation. This reflects the snap's nature as a pre-configured inference server where most user interaction is reading state.

2. **No lifecycle verbs**: There are no create/destroy/deploy/remove verbs. The engine and model lifecycle is managed by snap component installation, not CLI commands.

3. **No telic verbs**: All commands are either instantaneous state changes (punctual) or ongoing interactions (atelic). There is no long-running operation initiated by the CLI that has a natural completion point (the server daemon handles that, but it's not a CLI command users invoke).

4. **`use` is an unusual verb for CLIs**: Per DE013 §Commonly Used Commands, the standard vocabulary doesn't include `use`. The closest standard patterns would be `set-engine` or `switch-engine`. However, `use` communicates intent clearly ("I want to use this engine") and is distinct from `set` (which is reserved for key=value configuration).

5. **Missing `unset` verb**: DE013 specifies `get/set/unset` as a triplet for configuration management. The qwen36 CLI has `get` and `set` but no `unset` for restoring defaults.
