# 02 — Verb Taxonomy and Aspect Classification

## Verb Table

Every unique verb from the decomposition matrix is classified below.

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|------|-------------|--------|------------|-------------|-------------|
| show | observation | punctual | n/a | — | `show-engine`, `show-machine`, (implicit in `status`, `version`) |
| list | observation | punctual | n/a | — | `list-engines` |
| get | observation | punctual | n/a | — | `get [key]` |
| set | mutation | punctual | yes | unset | `set key=value` |
| unset | mutation | punctual | yes | set | `unset key` |
| use | mutation | punctual | partial | — | `use-engine cpu` |
| prune | lifecycle (destruction) | telic | no | — | `prune-cache` |
| chat | execution | atelic | n/a | — | `chat` |
| launch/open | execution | punctual | n/a | — | `webui` (implicit verb) |
| run | execution | atelic | n/a | — | `run <command>` (hidden) |
| serve | execution | atelic | n/a | — | `serve-webui` (hidden) |
| debug | observation | atelic | n/a | — | `debug` (hidden, namespace) |

## Intent Group Distribution

| Intent Group | Count | Verbs |
|-------------|-------|-------|
| observation | 4 | show, list, get, debug |
| mutation | 3 | set, unset, use |
| execution | 3 | chat, launch, run, serve |
| lifecycle | 1 | prune |

**Observation**: The CLI is heavily weighted toward observation (4 verbs) and light on lifecycle operations (1 verb). This is appropriate for an inference snap where the user's primary workflow is: select engine → configure → observe status → interact (chat).

## Aspect (Aktionsart) Analysis

| Aspect | Verbs | Count |
|--------|-------|-------|
| punctual | show, list, get, set, unset, use, launch | 7 |
| telic | prune | 1 |
| atelic | chat, run, serve, debug | 4 |

**Observation**: The CLI favors punctual (instant) operations. The atelic verbs (`chat`, `run`, `serve`) represent long-running interactive or server processes — a natural split between "configure the system" (punctual) and "use the system" (atelic).

## Reversibility Assessment

| Reversible | Verbs | Notes |
|-----------|-------|-------|
| yes | set ↔ unset | Clean symmetric pair |
| partial | use | Can switch to another engine, but no "un-use" or "previous engine" command |
| no | prune | Components must be re-downloaded from the snap store |
| n/a | show, list, get, chat, run, serve, debug | Read-only or interactive |

## Key Findings

1. **`set`/`unset` is a well-designed symmetric pair** — correct per DE013 get/set/unset convention.

2. **`use` has no inverse** — switching engines is one-directional. There is no `qwen36 previous-engine` or undo. Per DE013, this is acceptable since "use" is a selection verb, not a creation verb.

3. **`prune` is irreversible without explicit warning** — the command does prompt for confirmation, but does not state "this cannot be undone" or "components will need to be re-downloaded." Per safety best practices, irreversible actions should state the consequence explicitly.

4. **Verb vocabulary is minimal and well-chosen** — 11 unique verbs for 12-15 commands indicates low redundancy and good verb economy. The Canonical standard's recommended vocabulary (`get/set/unset`, `show`, `list`, `status`, `version`) is followed closely.
