# 01. Verb-Noun Decomposition Matrix

## Command accounting

Total commands analyzed: 7

| Command | Decomposed verb | Decomposed noun | Status | Notes |
|---|---|---|---|---|
| `qwen36 chat` | `chat` | none | orphan | Bare conversational verb; no explicit object noun is exposed in the command name |
| `qwen36 use-engine` | `use` | `engine` | decomposed | Verb-noun compound |
| `qwen36 show-engine` | `show` | `engine` | decomposed | Verb-noun compound |
| `qwen36 get` | `get` | `config` | decomposed | Bare verb, but the direct object is clearly config state via `<key>` |
| `qwen36 set` | `set` | `config` | decomposed | Bare verb, but the direct object is clearly config state via `<key>=<value>` |
| `qwen36 completion bash` | `completion` | `shell` | decomposed | Meta-command generating shell-specific completion words |
| `qwen36.server` | none | none | orphan | Snap service entrypoint outside the `qwen36 ...` grammar |

Commands covered by matrix: 5

Commands covered as orphans: 2

Coverage check: 5 + 2 = 7 commands accounted for.

## Verb x noun grid

Rows are verbs sorted alphabetically. Columns are nouns sorted by frequency.

| Verb \ Noun | engine | config | shell |
|---|---|---|---|
| `completion` | - | - | yes |
| `get` | - | yes | - |
| `set` | - | yes | - |
| `show` | yes | - | - |
| `use` | yes | - | - |

## Incomplete CRUD and lifecycle sets

| Noun | Present commands | Missing expected operations | Notes |
|---|---|---|---|
| `engine` | `use-engine`, `show-engine` | `list-engines`, `engine` or `engine-status`, `set-engine`, reset/unset form | The engine domain has selection and inspection, but no explicit status/reset/list symmetry |
| `config` | `get`, `set` | `unset`, `get --all` or equivalent enumeration | DE013 explicitly lists `get/set/unset` as the common trio |
| `shell` | `completion bash` | installation/removal guidance, other shell targets | The shell-integration domain is single-command and undocumented |

## Verb inconsistencies

| Domain | Current pattern | Inconsistency | Effect |
|---|---|---|---|
| engine | `use-engine`, `show-engine` | One mutation verb is underspecified (`use`), while the inspection verb uses the non-preferred `show-*` form | The pair works, but it does not align tightly with DE013 vocabulary |
| config | `get`, `set` | Bare verbs omit the object noun entirely, while engine commands include it | The CLI mixes bare verbs and verb-noun compounds for adjacent admin tasks |
| shell integration | `completion bash` | Uses a noun-like meta-command rather than a task verb | Reasonable for shell completion, but grammatically distinct from the rest of the set |
| service runtime | `qwen36.server` | Uses a service name outside the normal command hierarchy | Operationally necessary for snap, but conceptually outside the user grammar |

## Orphan commands

| Command | Why it is orphaned | Closest conceptual role |
|---|---|---|
| `qwen36 chat` | The command is a bare activity verb with no explicit object noun, so it does not fit the matrix cleanly | start an interactive conversation session |
| `qwen36.server` | It is a snap service/app identifier, not a normal `qwen36 <verb>` command | serve the inference API |

## Observations

1. The command set is small enough that the mixed grammar is learnable, but the mix is still visible: bare verbs, verb-noun compounds, and a meta-command all coexist.
2. `engine` is the only noun with more than one dedicated command, which makes its naming quality disproportionately important.
3. The strongest gap is config symmetry: `get` and `set` exist, but the DE013-aligned `unset` counterpart does not.
4. The service surface should be treated as an operational outlier rather than a model for new end-user commands.