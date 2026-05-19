# Juju Confusion Pairs

## Overview

These are command pairs or families that are likely to confuse users because of similar names, overlapping outcomes, or inconsistent verb choices.

## Highest-risk confusion pairs

### 1. `run` vs `exec` vs `ssh`

Why it is confusing:
- all three imply remote execution
- `run` means run a Juju action
- `exec` means run shell commands remotely on selected targets
- `ssh` may also execute a command, or open an interactive session

User failure mode:
- operators reach for `run` expecting shell execution
- scripts choose `ssh` where `exec` would be semantically clearer

Severity: high

### 2. `integrate` vs `remove-relation`

Why it is confusing:
- forward operation uses modern workflow verb `integrate`
- reverse operation uses older noun-centric `relation`

User failure mode:
- users search for `disintegrate`, `unintegrate`, or `remove-integrate` style commands
- the pair is not lexically discoverable without help output

Severity: high

### 3. `config` vs `model-config` vs `controller-config`

Why it is confusing:
- app config keeps the noun implicit
- model/controller config make the noun explicit
- all three support related get/set/file/reset patterns

User failure mode:
- accidental edits at the wrong scope
- expectation that all support the exact same reset/import/output semantics

Severity: high

### 4. `refresh` vs `upgrade-model` vs `upgrade-controller`

Why it is confusing:
- all sound like “upgrade something”
- `refresh` upgrades application charm content
- `upgrade-model` and `upgrade-controller` upgrade Juju agents/infrastructure layer

User failure mode:
- users assume `refresh` covers all upgrade cases
- users expect an `upgrade-application` verb because `upgrade-*` exists elsewhere

Severity: high

## Medium-risk confusion pairs

### 5. `show-*` vs plural noun listings

Examples:
- `show-model` vs `models`
- `show-user` vs `users`
- `show-space` vs `spaces`
- `show-storage` vs `storage`

Why it is confusing:
- some nouns use plural listing plus `show-*`
- others use only plural list or only a show command
- the list verb is implicit, not explicit

Severity: medium

### 6. `offer` / `offers` / `find-offers` / `show-offer` / `remove-offer`

Why it is confusing:
- this family is internally coherent, but dense
- `offer` is both a noun and a verb
- `offers` is the listing form, but `find-offers` searches beyond the local listing

Severity: medium

### 7. `destroy-controller` vs `kill-controller` vs `unregister`

Why it is confusing:
- all can feel like “stop using this controller”
- but they mean radically different things:
  - graceful destroy
  - forceful destroy
  - local client unregister only

Severity: medium to high

### 8. `remove-saas` vs `consume`

Why it is confusing:
- `consume` creates a local view of a remote offered application
- `remove-saas` tears down that consumed view
- the noun vocabulary changes from “offer” to “saas” across the pair

Severity: medium

### 9. `resources` vs `charm-resources` vs `attach-resource`

Why it is confusing:
- all are resource-oriented
- one is app/unit scoped, one is charm catalog scoped, one mutates app resources

Severity: medium

### 10. `find` vs `info` vs `show-*`

Why it is confusing:
- `find` and `info` are Charmhub/catalog verbs
- `show-*` is a common Juju local-resource inspection verb
- generic `info` is semantically weaker than explicit `show-charm` would be

Severity: medium

## Lower-risk but notable pairs

### 11. `enable-command` / `disable-command` / `enable-destroy-controller`

Why it is confusing:
- one pair operates on abstract command-set blocks
- the third is a one-off workflow command with a much more specific semantic target

Severity: medium

### 12. `default-region` / `default-credential` vs `set-credential`

Why it is confusing:
- all are cloud/credential-affecting commands
- some affect local defaults, one relates credentials to a model

Severity: medium

### 13. `machines` vs `status`

Why it is confusing:
- both show machine information, but at different detail and aggregation levels
- there is also `show-machine`

Severity: low to medium

## Root causes

The confusion pairs come from four recurring causes:
- mixed lexical strategies: verb-noun, noun plural, bare verb
- domain-specific exceptions that are reasonable in isolation
- overloaded execution verbs
- partial complement pairs where the reverse operation uses a different noun system

## Mitigation direction

The least disruptive fixes are:
- add discoverability aliases rather than breaking existing commands immediately
- normalize help text with stronger “see also” pairings
- standardize family naming where new commands are added
- expose a few missing peer commands or aliases such as `untrust` and an alias for `remove-relation`
