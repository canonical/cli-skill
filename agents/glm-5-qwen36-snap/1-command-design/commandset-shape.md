# qwen36 Command Set Shape Analysis

This document provides a comprehensive analysis of the qwen36 CLI command set shape, including verb-noun decomposition, verb taxonomy, semantic domain clustering, symmetry audit, confusion-pair audit, and pattern classification.

---

# 01 Verb-Noun Decomposition

## Scope

Source command count: 6 leaf commands.

- `qwen36 chat`
- `qwen36 use-engine`
- `qwen36 show-engine`
- `qwen36 get`
- `qwen36 set`
- `qwen36 completion bash`

## Decomposition Matrix

Semantic decomposition is used where the literal surface is not already verb-noun. This keeps every command in one matrix while still flagging grammar mismatches.

| Verb | engine | config | completion | conversation |
|------|-------:|-------:|-----------:|-------------:|
| chat | — | — | — | ✓ |
| generate | — | — | ✓ | — |
| get | — | ✓ | — | — |
| set | — | ✓ | — | — |
| show | ✓ | — | — | — |
| use | ✓ | — | — | — |

## Command-Level Mapping

| Command | Literal Surface | Semantic Verb | Semantic Noun | Clean Verb-Noun? | Notes |
|---------|-----------------|---------------|---------------|------------------|-------|
| `qwen36 chat` | verb only | `chat` | `conversation` | Partial | Usable, but noun is implicit rather than explicit. |
| `qwen36 use-engine` | verb-noun | `use` | `engine` | Yes | Clear enough, though `use` is broader than `select`. |
| `qwen36 show-engine` | verb-noun | `show` | `engine` | Yes, but non-ideal | Per DE013, showing state usually prefers `engine` or `engine-status` over `show-engine`. |
| `qwen36 get` | verb only | `get` | `config` | Partial | Common config grammar, noun implied by argument key space. |
| `qwen36 set` | verb only | `set` | `config` | Partial | Common config grammar, noun implied by assignment. |
| `qwen36 completion bash` | noun + target | `generate` | `completion` | No | Literal command is noun-led and conflicts with DE013 verb-first guidance. |

Self-check: output command count = 6.

## Incomplete CRUD Sets

| Noun | Present Verbs | Expected Missing Verbs | Assessment |
|------|---------------|------------------------|------------|
| `engine` | `use`, `show` | `list`, `status` or `engine`, optional `unset` or `reset` | Discovery is weak. Users can select or inspect the current engine but cannot list supported engines from the documented surface. |
| `config` | `get`, `set` | `unset` | Config grammar is close to standard, but restoring defaults is not surfaced. |
| `completion` | `generate` | none required | Narrow feature. Only bash is evidenced. |
| `conversation` | `chat` | none required | Single-purpose interaction command. |

## Verb Inconsistencies

| Resource | Current Surface | Consistency Issue |
|----------|-----------------|-------------------|
| `engine` | `use-engine`, `show-engine` | Mutation and observation verbs are fine, but DE013 prefers noun shorthand or `*-status` over `show-*` for observation. |
| `completion` | `completion bash` | Uses a noun-led grouping token instead of a verb. |
| `config` | `get`, `set` | Internally consistent and aligned with DE013 common commands. |

## Orphan Or Exception Commands

| Command | Why It Is Exceptional |
|---------|----------------------|
| `qwen36 chat` | Verb-only, implicit noun. Acceptable, but distinct from the rest of the config-and-engine grammar. |
| `qwen36 completion bash` | Only command that is noun-led at the top level. |

## Recommendation Compliance Notes

Per DE013 grammar, the strongest naming issue is `show-engine`, followed by `completion bash`.

Recommended change only if the team is willing to pay migration cost:

1. Add a preferred observational alias such as `engine` or `engine-status` in the next minor release while keeping `show-engine` working.
2. Emit a stderr warning: `warning: "qwen36 show-engine" is deprecated, use "qwen36 engine" instead`.
3. Keep that aliasing for at least one release cycle.
4. In the next major release, remove `show-engine` and fail with exit code 2 plus: `error: "qwen36 show-engine" was removed in 4.0, use "qwen36 engine" instead`.

For `completion bash`, the standard conflict is real, but the benefit of renaming is smaller. It should only be changed as part of a broader command grammar cleanup.

---

# 02 Verb Taxonomy

## Scope

This section classifies the six verbs implied by the six leaf commands. Source verb count: 6. Output verb count: 6.

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|------|--------------|--------|------------|-------------|--------------|
| `chat` | execution | atelic | no | none | `qwen36 chat` |
| `generate` | execution | telic | no | none | `qwen36 completion bash` |
| `get` | observation | punctual | partial | `set` | `qwen36 get http.port` |
| `set` | mutation | punctual | partial | `get` | `qwen36 set http.port=8326` |
| `show` | observation | punctual | no | none | `qwen36 show-engine` |
| `use` | mutation | punctual | partial | none | `qwen36 use-engine cpu`, `qwen36 use-engine --auto` |

## Notes By Verb

| Verb | Assessment |
|------|------------|
| `chat` | Clear for interactive use, but it does not tell the user whether the command starts a local client, a local server session, or a remote session. |
| `generate` | Semantic, not literal. The literal surface is `completion`, which is noun-led. |
| `get` | Strong fit for config observation. Aligned with DE013 common commands. |
| `set` | Strong fit for config mutation. Aligned with DE013 common commands. |
| `show` | Understandable, but DE013 explicitly prefers noun shorthand or `*-status` over `show-*` for observation commands. |
| `use` | Understandable, but slightly vague. It selects an engine rather than merely "using" one transiently. |

## Taxonomy Findings

1. The command set leans heavily on observation and mutation, which matches the CLI's role as a control plane.
2. `get` and `set` are the best-aligned verbs in the set.
3. `show` and noun-led `completion` are the main standards mismatches.
4. `use` is workable, but if the command family grows, `select-engine` would be semantically sharper than `use-engine`. That said, the improvement is not large enough to justify breaking users without a broader redesign.

## Recommendation Compliance Notes

Per DE013, any future cleanup should prefer canonical verbs and avoid noun-led command groups. Per the deprecation specification, if `show-engine` or `use-engine` were ever renamed, the next minor release should keep the old form working with a warning, and the next major release should remove it with exit code 2 and a replacement hint.

---

# 03 Semantic Domain Clustering

## Scope

Every leaf command appears in exactly one domain. Source command count: 6. Sum of domain counts below: 6.

| Domain | Count | Commands | Naming Consistent? | Notes |
|--------|------:|----------|-------------------|-------|
| conversation | 1 | `qwen36 chat` | Yes | Single interactive command. No CRUD expectation. |
| engine | 2 | `qwen36 use-engine`, `qwen36 show-engine` | Mostly | Noun is consistent, but observation uses `show-*`, which DE013 treats as second-best to noun shorthand or `*-status`. CRUD coverage is partial: choose and inspect exist; list/reset do not. |
| configuration | 2 | `qwen36 get`, `qwen36 set` | Yes | Strongest domain in the CLI. Standard read/write pair, but no `unset`. |
| shell integration | 1 | `qwen36 completion bash` | No | Domain is valid, but the command shape is noun-led and only one shell target is evidenced. |

## Domain Analysis

### conversation

- Noun form: implicit rather than explicit
- CRUD coverage: not applicable
- Verb consistency: one command only

### engine

- Noun form: explicit and stable as `engine`
- CRUD coverage: incomplete
- Verb consistency: acceptable but not ideal because `show-engine` conflicts with DE013 observation guidance

### configuration

- Noun form: implicit config domain derived from keys
- CRUD coverage: mostly useful, but no `unset`
- Verb consistency: high

### shell integration

- Noun form: explicit in semantics, but not in user-facing grammar
- CRUD coverage: not applicable
- Verb consistency: low because the top-level literal is a noun

## Findings

1. The command set clusters cleanly into four domains, which is appropriate for a small CLI.
2. Configuration and engine management are the core domains.
3. The shell-integration domain is underdeveloped and underdocumented.
4. No command currently helps users discover supported engine names from within the CLI itself.

## Recommendation Compliance Notes

Per DE013, the engine domain is the most obvious place to improve naming consistency, with `show-engine` as the candidate for eventual normalization. Per the deprecation spec, any rename should be additive first, with the old command retained for at least one cycle and removed only in a major release.

---

# 04 Symmetry Audit

## Scope

This audit includes all six leaf commands. Because the command set is small, most symmetry findings are about missing reverse operations rather than mismatched existing pairs.

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|-----------|-----------------|-----------------|-------------------|---------------------|-------|
| Engine selection | `qwen36 use-engine <engine>` | none | No | No | There is no explicit `unset-engine`, `reset-engine`, or `auto-engine` inverse. `--auto` is a mode on the same command rather than a reverse operation. |
| Engine observation vs engine mutation | `qwen36 show-engine` | `qwen36 use-engine <engine>` | No | No | These commands are complementary, not inverse. One inspects state; the other changes it. |
| Config mutation vs config observation | `qwen36 set <key>=<value>` | `qwen36 get <key>` | Partially | Partially | Common and useful pair, but `get` is not a true inverse because it does not restore prior state. |
| Config restoration | `qwen36 set <key>=<value>` | none | No | No | No `unset` command is documented, so there is no direct route back to defaults. |
| Interactive session entry | `qwen36 chat` | none | No | No | Session termination is shell-driven rather than command-driven. |
| Shell completion generation | `qwen36 completion bash` | none | No | No | Generation is one-way. No install/remove/manage complement is exposed. |

Self-check: every leaf command appears at least once in the table.

## Findings

1. The only meaningful near-pair is `get` and `set`.
2. The engine domain is asymmetrical: users can set and inspect the current engine but cannot list supported engines or restore defaults explicitly.
3. The CLI is not overburdened with forced symmetry, which is good, but it does need one missing complement: `unset` for config.

## Recommendation Compliance Notes

Per DE013, `get`/`set` already match Canonical's standard command vocabulary; adding `unset` would improve symmetry without renaming anything. That is the safest additive improvement because it can land in a minor release without breaking scripts. Any attempt to rename `show-engine` or `use-engine` should follow the full one-cycle deprecation process described in the deprecation specification.

---

# 05 Confusion Pair Audit

## Scope

This audit lists the most plausible command confusions in the six-command surface. Every leaf command appears in the table at least once.

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|-----------|-----------|--------------|----------------|----------------|
| `qwen36 use-engine` | `qwen36 show-engine` | scope ambiguity | High | Same noun, adjacent verbs, both about engine state. `use-engine` mutates; `show-engine` inspects. Improve docs with explicit "select" vs "inspect" language. |
| `qwen36 get` | `qwen36 set` | synonym verbs | Medium | Minimal pair differing by one verb. `get` reads config; `set` writes config. Standard and productive confusion, not harmful. |
| `qwen36 get` | `qwen36 show-engine` | functional overlap | Medium | Both can reveal engine-related information. `get` returns one scalar config value; `show-engine` returns structured YAML. Document when to prefer key lookup vs engine inspection. |
| `qwen36 set` | `qwen36 use-engine` | functional overlap | High | Both can change effective server behavior. `set` mutates one explicit key; `use-engine` changes a bundle of engine-related values. Document that `use-engine` is the high-level operation. |
| `qwen36 chat` | `qwen36 use-engine` | scope ambiguity | Medium | New users may think engine choice happens inside chat. `chat` starts client; `use-engine` prepares backend. README should place engine selection before chat. |
| `qwen36 completion bash` | `qwen36 chat` | naming collision | Low | Both appear terminal-facing. `completion bash` emits completion tokens; `chat` opens interaction. Add README shell-completion section. |

Self-check: output command count = 6 unique leaf commands represented.

## Findings

1. The only high-risk confusions are in the engine and config domains.
2. `show-engine` is understandable, but its `show-*` grammar makes it look like a generic display command rather than a domain-specific inspection command.
3. `set` versus `use-engine` needs explicit documentation because both change runtime behavior at different abstraction levels.

## Recommendation Compliance Notes

Per DE013, command names should make the user's mental model predictable. The safest immediate improvement is documentation, not renaming. If the team later decides to normalize `show-engine`, follow the deprecation spec exactly: add the replacement in a minor release, keep the old form working with a warning for at least one cycle, then remove it in a major release with exit code 2 and a migration hint.

---

# 06 Pattern Classification

## Scope

This classification covers all six leaf commands.

| Command | Surface Pattern | Argument Pattern | Standards Fit | Scriptability | Recommendation |
|---------|-----------------|------------------|---------------|---------------|----------------|
| `qwen36 chat` | flat verb | no args | Good | Low | Keep. It is memorable and task-oriented. |
| `qwen36 use-engine` | flat verb-noun | positional noun value plus optional mode flags | Acceptable | Medium | Keep for now. If a broader redesign ever happens, `select-engine` is a possible sharper verb, but not enough benefit to justify churn alone. |
| `qwen36 show-engine` | flat verb-noun | no args | Weak | High | Best rename candidate because DE013 prefers noun shorthand or `*-status` over `show-*`. Do not break immediately; add alias first. |
| `qwen36 get` | flat verb | single positional key | Strong | High | Keep. Standard config read command. |
| `qwen36 set` | flat verb | single positional `key=value` plus `--package` | Strong | High | Keep. Add documentation for `--package` and value validation. |
| `qwen36 completion bash` | two-level noun namespace | second-level shell target | Weak | Medium | Keep short term for compatibility; document it. Only revisit if the team decides to standardize every top-level command around verb-first grammar. |

Self-check: output command count = 6.

## Pattern Families

| Family | Commands | Assessment |
|--------|----------|------------|
| Flat verbs | `qwen36 chat`, `qwen36 get`, `qwen36 set` | Strongest family. Short, predictable, and aligned with DE013 common verbs. |
| Flat verb-noun | `qwen36 use-engine`, `qwen36 show-engine` | Coherent domain family, but `show-engine` is the weakest member because of the standard's observation guidance. |
| Two-level namespace | `qwen36 completion bash` | Functional but stylistically inconsistent with the rest of the command set. |

## Overall Classification

The CLI is mostly a flat command set with one small exception namespace. That is a good size and shape for a six-command tool. The main issue is not excessive hierarchy; it is uneven grammar quality across commands.

## Discoverability Assessment

| User Goal | Predicted Path | Actual Command | Gap |
|-----------|----------------|----------------|-----|
| Start chatting | `qwen36 chat` | `qwen36 chat` | None |
| See current engine | `qwen36 engine` or `qwen36 status` | `qwen36 show-engine` | Minor naming friction |
| Change engine | `qwen36 use-engine <name>` | `qwen36 use-engine <name>` | None, but `list-engines` is missing |
| View config | `qwen36 config` or `qwen36 get` | `qwen36 get <key>` | Key discovery is hard |
| Change config | `qwen36 config <key>=<value>` or `qwen36 set` | `qwen36 set <key>=<value>` | None, but key list is missing |
| Shell completion | `qwen36 --help` or `qwen36 completion` | `qwen36 completion bash` | Undocumented in README |

## Ecosystem Comparison

| Tool | Command Pattern | Similarity to qwen36 |
|------|-----------------|----------------------|
| `snap` | Flat verbs (`snap list`, `snap install`, `snap get`, `snap set`) | High similarity for `get`/`set` |
| `lxc` | Flat verbs with noun objects (`lxc list`, `lxc launch`) | Similar flat structure |
| `juju` | Flat verbs (`juju deploy`, `juju status`, `juju config`) | Similar flat structure |
| `multipass` | Flat verbs (`multipass list`, `multipass start`, `multipass get`) | High similarity for `get`/`set` |

## Recommendations

| Priority | Change | Rationale | Backward Compat | Migration Cost |
|----------|--------|-----------|-----------------|----------------|
| 1 | Add `qwen36 engine` as alias for `show-engine` | DE013 compliance | Fully compatible | Low |
| 2 | Add `qwen36 unset <key>` command | Symmetry with `get`/`set` | Additive | Low |
| 3 | Add `qwen36 list-engines` command | Discoverability | Additive | Low |
| 4 | Document `--package` flag | Hidden feature | Compatible | None |
| 5 | Add `qwen36 status` command | Server health visibility | Additive | Low |
| 6 | Consider `select-engine` as alias for `use-engine` | Semantic precision | Compatible via alias | Low |

## Recommendation Compliance Notes

Per DE013, retain `get` and `set` as the anchor vocabulary and avoid introducing more noun-led top-level commands. Per the deprecation specification, if `show-engine` is ever normalized to `engine` or `engine-status`, the migration should be:

1. Minor release: add the replacement command and keep `show-engine` working.
2. Minor release stderr warning: `warning: "qwen36 show-engine" is deprecated, use "qwen36 engine" instead`.
3. Maintain both for at least one full cycle.
4. Major release: remove `show-engine`, return exit code 2, and print `error: "qwen36 show-engine" was removed in 4.0, use "qwen36 engine" instead`.
