# Juju Command-Surface Recommendations

## Design goal

Improve clarity and symmetry without abandoning the flat top-level CLI or breaking established operator workflows abruptly.

## 1. Normalize usage metadata first

This is the highest-leverage change because Juju depends heavily on generated help and docs.

Recommendations:
- ban `[options]` inside `Info.Args`
- ban embedding flags inside `Info.Args`
- make the markdown generator always emit a Usage section, even when `Args` is empty
- add lint or tests for duplicated option markers and malformed usage strings

Why first:
- it improves help, docs, inventories, and downstream command-design work immediately
- it does not require semantic changes to the CLI surface

## 2. Add discoverability aliases for asymmetric pairs

Recommended aliases:
- alias `remove-relation` from `disconnect` or `separate`, while keeping the current command
- alias a trust-removal path as `untrust` if the underlying `trust` command already supports revocation flags
- consider aliasing `info` with a more explicit Charmhub-oriented name in docs, even if the command stays `info`

Why:
- aliases improve discoverability without forcing a disruptive rename

## 3. Tighten execution verb semantics

Current risk area:
- `run`, `exec`, and `ssh` overlap in user mental models

Recommendations:
- keep `run` for actions for backward compatibility, but position `exec` more prominently as remote shell execution in docs and help cross-links
- add stronger cross-references among `run`, `exec`, and `ssh`
- consider a long-term alias such as `run-action` for clarity, even if `run` remains canonical

## 4. Complete list/show symmetry for missing nouns

High-value additions or aliases:
- `show-storage-pool`
- `show-subnet`
- `show-firewall-rule` if the data model supports it cleanly

Why:
- Juju already teaches a list/show pattern in many domains
- completing that pattern improves predictability more than inventing new verbs

## 5. Standardize safety semantics across mutators

Recommendations:
- document a CLI-wide interpretation guide for `--force`, `--no-prompt`, and dry-run behavior
- add dry-run or preview support to more high-impact mutators where practical: networking changes, relation changes, config import, and refresh-like operations
- reserve `kill-*` for emergency teardown semantics and avoid overloading `--force` indefinitely for similar escape hatches

## 6. Clarify scope in configuration commands

Recommendations:
- keep `config`, `model-config`, and `controller-config`, because they are already well established
- strengthen the help text to emphasize scope first line, before details
- standardize examples so each config command starts with the same pattern: inspect all, inspect one key, set inline, set from file, reset if supported
- consider a future `cloud-defaults` style grouping or topic doc for `default-region` and `default-credential`

## 7. Improve family-level grouping in documentation without changing the CLI shape yet

Recommendations:
- generate a categorized `help commands` topic grouped by semantic domain rather than one flat alphabetical block only
- keep alphabetical output available, but add grouped help for domains such as `applications`, `models`, `clouds`, `identity`, `storage`, `networking`, `secrets`
- improve “See also” links so forward and reverse operations are always paired

Why:
- Juju's biggest discoverability issue is flatness; grouped help mitigates that without forcing nested subcommands

## 8. Regularize output contracts

Recommendations:
- declare that list/show/report commands should support JSON and YAML unless there is a strong reason not to
- standardize default human-readable formats by family
- document JSON/YAML stability expectations explicitly
- keep `status` rich, but document format-specific behavioral differences centrally

## 9. Tidy the odd one-off names

Candidates for review:
- `enable-destroy-controller`
- `help-action-commands`
- `help-hook-commands`

Potential direction:
- preserve them, but improve topic-based discoverability so these do not have to carry as much of the UX burden as standalone names

## 10. Prefer additive migration paths

For any naming cleanup:
1. add aliases first
2. update docs and generated help
3. mark older names deprecated only after the alias has been discoverable for at least one release cycle
4. avoid removing short entrenched verbs unless the confusion cost is severe

## Proposed near-term roadmap

### Near term

- fix usage metadata defects
- add grouped command help
- add cross-links for confusion pairs
- add a few discoverability aliases

### Mid term

- standardize list/show and output support across more nouns
- add more dry-run coverage
- clarify execution verb guidance in help/docs

### Long term

- consider whether a topic-oriented subcommand layer should exist in parallel with the flat surface, not necessarily as a replacement
- formalize command metadata linting so new commands cannot regress usage or docs quality

## Net recommendation

Do not redesign Juju into a totally different CLI grammar. The current flat command surface is workable and already institutionalized. The right move is a disciplined normalization pass:
- better metadata
- stronger symmetry
- a few carefully chosen aliases
- clearer grouped discoverability
- standardized output and safety policies

That would materially improve operator experience without paying the compatibility cost of a ground-up surface rewrite.
