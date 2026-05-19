# Verb-Noun Decomposition Matrix

## Scope

Source command count for command-shape analysis: 12 public commands.

Commands counted:

1. `status`
2. `chat`
3. `webui`
4. `get`
5. `set`
6. `unset`
7. `list-engines`
8. `show-engine`
9. `use-engine`
10. `show-machine`
11. `prune-cache`
12. `version`

Note: the shipped qwen36 snap currently exposes only 10 commands by default because `chat` and `webui` are feature-gated but the root snap manifest does not set `ADDITIONAL_FEATURES`. This workflow still analyzes the intended public command set of 12 because the CLI source and README both treat `chat` and `webui` as public commands.

## Grid

Rows are verbs. Columns are nouns or resource types.

| Verb | config | engine | machine | cache |
|---|---|---|---|---|
| `get` | ✓ | — | — | — |
| `list` | — | ✓ | — | — |
| `prune` | — | — | — | ✓ |
| `set` | ✓ | — | — | — |
| `show` | — | ✓ | ✓ | — |
| `unset` | ✓ | — | — | — |
| `use` | — | ✓ | — | — |

## Command-Level Mapping

| Command | Verb | Noun | Clean Verb-Noun? | Notes |
|---|---|---|---|---|
| `get` | `get` | `config` | partial | Standard config verb; noun is implicit. |
| `set` | `set` | `config` | partial | Standard config verb; noun is implicit. |
| `unset` | `unset` | `config` | partial | Standard config revert verb; noun is implicit. |
| `list-engines` | `list` | `engines` | yes | Conventional verb-noun list command. |
| `show-engine` | `show` | `engine` | yes | Observation command for one engine. |
| `use-engine` | `use` | `engine` | yes | Selection command, though the verb is semantically weak. |
| `show-machine` | `show` | `machine` | yes | Observation command for host hardware. |
| `prune-cache` | `prune` | `cache` | yes | Non-standard but understandable maintenance verb. |
| `status` | `status` | — | orphan | Noun-as-command; standard status exception. |
| `chat` | `chat` | — | orphan | Bare interaction verb, no explicit resource noun. |
| `webui` | — | `webui` | orphan | Noun-led command; launch/open behavior is implied, not named. |
| `version` | `version` | — | orphan | Noun-as-command; standard version exception. |

Self-check: 8 commands are represented in the grid and 4 are explicitly accounted for as orphans. Total accounted commands = 12.

## Incomplete CRUD Sets

| Noun | Present Verbs | Missing Expected Verbs | Notes |
|---|---|---|---|
| `config` | `get`, `set`, `unset` | `list` or `show` | `get` with no key already acts as list-all, so the gap is discoverability rather than capability. |
| `engine` | `list`, `show`, `use` | `select`, `unset`, `status` | Public engine management lacks a canonical selection verb and a reset path. |
| `machine` | `show` | `list`, `status` | Reasonable for a single-host snap, but discoverability is shallow. |
| `cache` | `prune` | `show`, `restore` | Maintenance is one-way. |

## Verb Inconsistencies

- `use-engine` is inconsistent with hidden `debug select-engine`; the same product uses both `use` and `select` for engine choice.
- `webui` breaks the otherwise verb-oriented command shape by using a noun as the command.
- `prune-cache` uses domain jargon instead of a more predictable removal verb.

## Orphan Commands

| Command | Why it does not decompose cleanly |
|---|---|
| `status` | standard noun-as-command for state inspection |
| `chat` | direct interaction verb with no explicit object |
| `webui` | noun-led launcher command |
| `version` | standard noun-as-command for metadata |
