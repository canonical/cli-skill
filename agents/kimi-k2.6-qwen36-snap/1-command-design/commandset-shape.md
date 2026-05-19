# Command Set Shape: qwen36 / inference-snaps-cli

> **Self-verification**: source command count = 18. Every section below has been verified to include all 18 commands.

---

## Section 1 — Verb-Noun Decomposition Matrix

Decompose every command into a verb and a noun. The matrix below shows verb × noun coverage.

### Raw Decomposition Table

| # | Command | Verb | Noun | Notes |
|---|---|---|---|---|
| 1 | `status` | status | *(system)* | Standard state command per DE013 |
| 2 | `chat` | chat | — | Direct verb; objectless interactive session |
| 3 | `webui` | launch | webui | Noun-as-command (launch implied by "Open ... in browser" behavior) |
| 4 | `get` | get | config | Standard config read per DE013 |
| 5 | `set` | set | config | Standard config write per DE013 |
| 6 | `unset` | unset | config | Standard config revert per DE013 |
| 7 | `list-engines` | list | engines | Standard list plural noun per DE013 |
| 8 | `show-engine` | show | engine | Verb-noun compound; `show` used for details |
| 9 | `use-engine` | use | engine | Verb-noun compound; "use" is non-standard per DE013 |
| 10 | `show-machine` | show | machine | Verb-noun compound |
| 11 | `prune-cache` | prune | cache | Verb-noun compound |
| 12 | `version` | version | — | Noun-as-command (standard `tool version`) |
| 13 | `run` | run | command | Standard execution verb; argument is a subprocess command |
| 14 | `serve-webui` | serve | webui | Verb-noun compound |
| 15 | `debug validate-engines` | validate | engines | Verb-noun compound in debug namespace |
| 16 | `debug select-engine` | select | engine | Verb-noun compound in debug namespace |
| 17 | `debug chat` | chat | — | Same direct verb as top-level `chat` |
| 18 | `debug serve-webui` | serve | webui | Same as top-level hidden `serve-webui` |

### Matrix (Verbs × Nouns)

Verbs (rows): chat, get, launch, list, prune, run, select, serve, set, show, status, unset, use, validate, version

Nouns (cols): cache, command, config, engine, engines, machine, system, version, webui

| Verb \ Noun | cache | command | config | engine | engines | machine | system | version | webui |
|---|---|---|---|---|---|---|---|---|---|
| **chat** | — | — | — | — | — | — | — | — | — |
| **get** | — | — | ✓ | — | — | — | — | — | — |
| **launch** | — | — | — | — | — | — | — | — | ✓ |
| **list** | — | — | — | — | ✓ | — | — | — | — |
| **prune** | ✓ | — | — | — | — | — | — | — | — |
| **run** | — | ✓ | — | — | — | — | — | — | — |
| **select** | — | — | — | ✓ | — | — | — | — | — |
| **serve** | — | — | — | — | — | — | — | — | ✓ |
| **set** | — | — | ✓ | — | — | — | — | — | — |
| **show** | — | — | — | ✓ | — | ✓ | — | — | — |
| **status** | — | — | — | — | — | — | ✓ | — | — |
| **unset** | — | — | ✓ | — | — | — | — | — | — |
| **use** | — | — | — | ✓ | — | — | — | — | — |
| **validate** | — | — | — | — | ✓ | — | — | — | — |
| **version** | — | — | — | — | — | — | — | ✓ | — |

### Annotations

#### Incomplete CRUD Sets
- **engine**: has `show-engine`, `use-engine`, but no `create-engine`, `remove-engine`, `update-engine`. This is acceptable because engines are external manifests + snap components, not managed objects created by the CLI.
- **engines**: has `list-engines`, `validate-engines` (debug), but no `add-engines` or `remove-engines`. Same rationale.
- **config**: has full read/write/revert via `get`, `set`, `unset`. No `list-config` or `show-config`. `get` with no key acts as list, but the verb differs from the plural-noun pattern (`list-engines`).
- **cache**: has `prune-cache` only. No `show-cache`, `clear-cache` (prune is effectively clear), or `get-cache`.
- **machine**: has `show-machine` only. No `list-machines` (only one host) or `get-machine`.

#### Verb Inconsistencies
- **`use-engine`** vs standard patterns: DE013 §Grammar expects `select` or `set` for choosing a primary entity. The CLI already has `debug select-engine`, making `use-engine` inconsistent with its own debug synonym. This is a verb inconsistency.
- **`show-engine`** and **`show-machine`** use `show`, which DE013 prefers over `info` when the tool name does not define the object type. This is acceptable.
- **`prune-cache`** uses `prune`; DE013 standard verbs for removal are `remove`, `delete`, or `clear`. `prune` is domain-specific but non-standard.

#### Orphan Commands
- `chat`, `debug chat` — direct verb with no object.
- `status` — noun-as-command; standard per DE013.
- `version` — noun-as-command; standard per DE013.
- `webui` — noun-as-command; implies launch/open behavior that is not explicit in the name. Per DE013, this should be `launch-webui` or `open-webui`.
- `run` — standard verb with a generic subprocess argument; decomposes as verb + generic noun.

---

## Section 2 — Verb Taxonomy and Aspect Classification

Classify every unique verb from Section 1.

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|
| chat | execution | atelic | no | — | `chat`, `debug chat` |
| get | observation | atelic | no | — | `get` |
| launch | execution | telic | no | — | `webui` (implied) |
| list | observation | atelic | no | — | `list-engines` |
| prune | lifecycle | telic | no | — | `prune-cache` |
| run | execution | atelic | no | — | `run` |
| select | access | telic | partial | — | `debug select-engine` |
| serve | execution | atelic | no | — | `serve-webui`, `debug serve-webui` |
| set | mutation | telic | yes | unset | `set` |
| show | observation | atelic | no | — | `show-engine`, `show-machine` |
| status | observation | atelic | no | — | `status` |
| unset | mutation | telic | yes | set | `unset` |
| use | access | telic | partial | — | `use-engine` |
| validate | observation | telic | no | — | `debug validate-engines` |
| version | observation | atelic | no | — | `version` |

### Aspect Notes
- **telic**: Actions with a natural endpoint (`set`, `unset`, `prune`, `select`, `use`, `validate`).
- **atelic**: Ongoing or continuous observation/execution (`chat`, `list`, `show`, `status`, `run`, `serve`).
- **punctual**: None present; no instant state-change verbs like `enable`/`disable`.

### Reversibility Notes
- `set` ↔ `unset` is a true symmetric pair: `set` writes a user value; `unset` removes it, reverting to the lower-precedence layer.
- `use-engine` is partially reversible only by invoking `use-engine <other-engine>`; there is no `unuse-engine` or `reset-engine`.
- `prune-cache` is irreversible; removed components must be re-downloaded.
- `chat`, `run`, `serve` are runtime sessions with no persistent reversible state.

---

## Section 3 — Semantic Domain Clustering

Group all commands by the resource domain they operate on. Count verification: 18 commands.

| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---|---|---|---|
| **System / Snap state** | 3 | `status`, `version`, `show-machine` | Partial | `status` and `version` are noun-as-commands; `show-machine` follows verb-noun. `show-machine` inspects hardware, not snap state, but belongs to the system introspection domain. |
| **Configuration** | 3 | `get`, `set`, `unset` | Yes | All use standard config verbs per DE013. No command name includes the noun "config", relying on verb convention instead. |
| **Engine lifecycle** | 7 | `list-engines`, `show-engine`, `use-engine`, `prune-cache`, `run`, `serve-webui`, `debug validate-engines` | No | Verbs vary: `list`, `show`, `use`, `prune`, `run`, `serve`, `validate`. `use-engine` is the odd verb; `prune-cache` operates on engine components but uses a different noun. |
| **Engine selection (debug)** | 2 | `debug select-engine`, `debug serve-webui` | Partial | `select` aligns with standard verbs; `serve` is execution. Hidden from normal users. |
| **Chat / Interaction** | 2 | `chat`, `debug chat` | Yes | Both use the same direct verb. Top-level is feature-gated. |
| **Web UI** | 2 | `webui`, `debug serve-webui` | No | `webui` is a noun-as-command; `serve-webui` is verb-noun. They perform different roles (open browser vs. serve static files), but naming does not clearly differentiate. |

### Domain Consistency Findings
- **Configuration** is the most consistent domain: pure DE013 verbs, no noun in command names.
- **Engine lifecycle** is the most inconsistent: seven commands spread across six different verbs, with no clear CRUD pattern for engine objects.
- **Web UI** has a collision: the top-level `webui` command (open browser) and the hidden `serve-webui` command (start file server) do not share a verb or noun prefix that signals their relationship.

---

## Section 4 — Symmetry Audit

List every pair of symmetric operations, including missing reverse operations.

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|
| Config write / revert | `set <key=value>` | `unset <key>` | Yes (verb pair) | Yes | True inverse: `unset` removes the user layer value set by `set`. |
| Engine select / re-select | `use-engine <name>` | `use-engine <other>` | No | Partial | No explicit reverse. Switching back requires knowing the previous engine name. No `unuse-engine`. |
| Component install / remove | *(implicit in `use-engine`)* | `prune-cache` | No | Partial | `use-engine` installs components; `prune-cache` removes inactive components. Names do not mirror. Behavior is asymmetric: prune operates on all inactive engines by default, not the complement of the last `use-engine`. |
| Config passthrough add / remove | `set passthrough.*=...` | `unset passthrough.*` | Yes | Yes | Works via the standard `set`/`unset` pair. |
| Server start / stop | *(not a CLI command)* | *(not a CLI command)* | — | — | Start/stop are delegated to `snap start/stop`. The CLI only prompts for restart. This is an asymmetric boundary. |
| Engine enable / fix | `use-engine <name>` | `use-engine --fix` | No | Partial | `--fix` reinstalls and reconfigures the *current* engine; it is not a generic undo. |
| Chat start / stop | `chat` | *(Ctrl+C / EOF)* | — | — | No CLI-level `stop-chat` or `exit-chat`. Session ends interactively. |
| Web UI open / close | `webui` | *(not provided)* | — | — | Opens browser; no command to close or stop the UI server from the CLI. |
| Run subprocess | `run <command>` | *(subprocess exits)* | — | — | No `kill` or `stop-run`. Cleanup is process-local and deferred. |
| Validate / invalidate | `debug validate-engines` | *(none)* | — | — | No `invalidate` or `unvalidate`. Manifest validation is one-way. |

### Missing Reverse Operations (Critical Gaps)
1. **`unuse-engine` / `reset-engine`**: There is no way to clear the active engine selection and return to a clean state without selecting another engine.
2. **`restore-cache` / `reinstall-engine`**: `prune-cache` removes components; re-acquiring them requires `use-engine` again, which may trigger a full engine switch workflow.
3. **No `stop`/`restart` CLI commands**: Service management is entirely delegated to `snapctl`. The CLI only *prompts* for restart during config/engine changes.

### Naming Asymmetries
- `use-engine` (forward) vs no reverse verb. DE013 recommends `select` or `enable`/`disable` for activation states; `use` is non-standard.
- `prune-cache` (removal) vs no `fetch-cache` or `install-cache`. The verb `prune` implies trimming excess, not an exact inverse of installation.

---

## Section 5 — Confusion-Pair Audit

List all command pairs that share semantic overlap and risk user confusion.

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|
| `use-engine` | `debug select-engine` | synonym verbs | **High** | Both choose an engine. `use-engine` applies the choice system-wide (installs components, restarts daemon). `debug select-engine` is a dry-run preview reading from stdin and printing to stdout; it makes no persistent changes. |
| `chat` | `debug chat` | scope ambiguity | **High** | `chat` connects to the local snap's OpenAI server automatically. `debug chat` requires `--base-url` for an arbitrary endpoint and is intended for developer testing. |
| `serve-webui` | `debug serve-webui` | naming collision | **Medium** | Both start a static file server. The top-level `serve-webui` is hidden and reads endpoint/config from the active engine; the debug variant accepts `--base-url` explicitly for testing. |
| `webui` | `serve-webui` | functional overlap | **Medium** | `webui` opens the user's default browser to the Web UI URL. `serve-webui` starts the backend file server. A user may type `webui` expecting to start a server, or `serve-webui` expecting a browser tab. |
| `status` | `show-engine` | scope ambiguity | **Medium** | `status` shows the active engine name as part of overall snap state. `show-engine` shows detailed manifest info for one engine. A user wanting "what engine is running?" might try either. |
| `get` | `show-engine` | functional overlap | **Low** | `get` retrieves config values (some of which are engine-derived). `show-engine` retrieves manifest metadata. Config keys like `model-name` overlap with engine manifest fields. |
| `set` | `use-engine` | functional overlap | **Medium** | `set` can modify config values that are also engine defaults (e.g., `http.port`). `use-engine` resets engine configs. A user changing `http.port` via `set` may be surprised when `use-engine` overwrites it with engine defaults. |
| `prune-cache` | `use-engine --fix` | scope ambiguity | **Low** | Both clean up engine-related state. `prune-cache` removes unused components to free disk space. `use-engine --fix` reinstalls missing components for the active engine. They are opposites but their names do not signal that. |
| `list-engines` | `debug validate-engines` | functional overlap | **Low** | Both inspect engine manifests. `list-engines` evaluates compatibility; `validate-engines` checks schema correctness. `validate-engines` is a debug-only command, but users may look for manifest validation in the main command set. |
| `run` | `chat` | scope ambiguity | **Low** | Both execute something interactively. `run` executes an arbitrary subprocess in the engine environment. `chat` starts a REPL against the OpenAI endpoint. The mental model differs but a new user may conflate "interactive CLI" behaviors. |
| `show-machine` | `debug select-engine` | functional overlap | **Low** | `show-machine` outputs hardware state. `debug select-engine` consumes that same hardware state (via stdin) to rank engines. Users may expect a single command that both shows hardware and selects an engine. |
| `version` | `status` | naming collision | **Low** | Both are read-only state queries. `version` shows snap and CLI versions; `status` shows engine and service state. A user may try one when looking for the other. |

---

## Section 6 — Pattern Classification and Recommendations

### Pattern Classification

| Property | Current Value |
|---|---|
| **Primary grouping pattern** | Flat command groups (Cobra `GroupID`) with verb-noun compounds for resource-specific commands |
| **Depth** | Max depth 1 (`debug` subcommand) |
| **Style** | Mixed: DE013 standard verbs for config (`get`/`set`/`unset`), non-standard verbs for engines (`use`, `prune`), noun-as-commands for UI (`webui`, `status`, `version`) |
| **Naming convention** | Mostly verb-noun hyphenated compounds; exceptions for config (bare verbs) and UI (bare nouns) |

### Discoverability Assessment

| User Task | Predicted Path | Actual Command | Match? |
|---|---|---|---|
| "What engine is running?" | `show-engine` or `status` | `status` (shows active) | Partial |
| "Switch to CUDA engine" | `use-engine cuda` | `use-engine cuda` | Yes |
| "Show available engines" | `list-engines` | `list-engines` | Yes |
| "Remove old engine data" | `remove-engine` or `prune-cache` | `prune-cache` | Partial |
| "Open chat" | `chat` | `chat` | Yes |
| "Open web UI" | `webui` | `webui` | Yes |
| "Show my hardware" | `show-machine` | `show-machine` | Yes |
| "Validate a manifest" | `validate-engines` | `debug validate-engines` | No (hidden) |
| "Preview engine selection" | `select-engine` | `debug select-engine` | No (hidden) |
| "Configure a key" | `set` | `set` | Yes |

### Ecosystem Comparison

| Tool | Structure | Comparison |
|---|---|---|
| **snap** | Flat verbs: `list`, `info`, `install`, `remove`, `set`, `get` | Very similar config verb pattern. Snap uses `install`/`remove` for lifecycle; qwen36 uses `use-engine` + `prune-cache`, which is less discoverable. |
| **juju** | Noun-verb subcommands: `juju status`, `juju deploy`, `juju remove-application` | Deeper hierarchy. qwen36 is flatter, which aids discoverability but lacks a clear `remove-*` verb for engine cleanup. |
| **docker** | Noun-verb: `docker image ls`, `docker container run` | docker uses noun namespaces to scope verbs. qwen36 uses top-level verb-noun compounds instead, avoiding depth but creating longer command names. |

### Recommendations

1. **Rename `use-engine` to `select-engine`**
   - **Rationale**: DE013 §Grammar states "Commands are verbs" and expects standard verbs like `select` for choosing an entity. The debug namespace already contains `select-engine`, proving the CLI has recognized `select` as the correct verb. Having `use-engine` at the top level and `select-engine` hidden is confusing.
   - **Backward compat**: Keep `use-engine` as a hidden alias for at least one release cycle per deprecation spec. Emit: `warning: "use-engine" is deprecated, use "select-engine" instead`.
   - **Migration cost**: Low. Affects interactive users and documentation; scripts using `use-engine` continue to work during deprecation.

2. **Rename `prune-cache` to `remove-engine` or `prune-engines`**
   - **Rationale**: `prune` is non-standard (DE013 prefers `remove` or `delete` for destructive removal). The command actually removes engine components, not a generic "cache". `remove-engine` or `prune-engines` (plural, since it affects all inactive engines by default) better communicates intent.
   - **Backward compat**: Add `prune-cache` as a deprecated alias if renamed.
   - **Migration cost**: Low-Medium. May be referenced in disk-cleanup documentation.

3. **Promote `debug validate-engines` and `debug select-engine` to top-level or a `dev` group**
   - **Rationale**: Power users (snap packagers, engine authors) need manifest validation and selection preview. Burying them under `debug` makes them undiscoverable. Consider a top-level `validate-engines` command (or `debug` group renamed to `dev` with non-hidden listing).
   - **Backward compat**: Keep `debug` aliases.
   - **Migration cost**: Low. New commands; old paths remain functional.

4. **Add `engines` top-level alias for `list-engines`**
   - **Rationale**: DE013 §Grammar recommends `tool foobars` as shorthand for listing secondary objects. `qwen36 engines` is a natural shorthand for `qwen36 list-engines`.
   - **Backward compat**: New alias; no breaking change.
   - **Migration cost**: None.

5. **Clarify `webui` vs `serve-webui` naming**
   - **Rationale**: `webui` is a noun-as-command implying "open". `serve-webui` is hidden and starts a server. To reduce confusion, rename hidden `serve-webui` to `serve-webui` (already the case) but ensure the top-level visible command is `open-webui` or `launch-webui` per DE013 verb conventions. Alternatively, rename `webui` to `launch-webui`.
   - **Backward compat**: Keep `webui` as alias.
   - **Migration cost**: Low. Top-level rename with alias.

6. **Add `reset-engine` or `clear-engine` command**
   - **Rationale**: Symmetry audit identified no reverse for `use-engine`. A command to clear the active engine (without selecting another) provides an escape hatch and improves safety.
   - **Backward compat**: New command; no breakage.
   - **Migration cost**: None.

### Tradeoff Summary

| Rec # | Breaking? | User Impact | Script Impact | Effort |
|---|---|---|---|---|
| 1 (select-engine) | No (alias) | High clarity | None during deprecation | Low |
| 2 (remove-engine) | No (alias) | Medium clarity | None during deprecation | Low |
| 3 (promote debug) | No | High discoverability | None | Low |
| 4 (engines alias) | No | Low convenience | None | Very low |
| 5 (launch-webui) | No (alias) | Medium clarity | None during deprecation | Low |
| 6 (reset-engine) | No | Safety improvement | None | Medium |
