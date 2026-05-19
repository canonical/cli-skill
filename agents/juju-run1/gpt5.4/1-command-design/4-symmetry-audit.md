# Juju Symmetry Audit

## Overview

This audit checks whether established verb-noun pairs have their expected complements, and whether similar nouns are handled consistently across domains.

## Strong symmetry pairs

### Add / remove

Well represented across many nouns:
- `add-model` / no exact `remove-model`, but `destroy-model` exists as the destructive complement
- `add-cloud` / `remove-cloud`
- `add-credential` / `remove-credential`
- `add-space` / `remove-space`
- `add-user` / `remove-user`
- `add-secret` / `remove-secret`
- `add-secret-backend` / `remove-secret-backend`
- `add-ssh-key` / `remove-ssh-key`

Assessment:
- strong overall
- model uses `destroy` instead of `remove`, which is defensible because teardown is much heavier than simple deletion

### Grant / revoke

Very strong across access-control surfaces:
- `grant` / `revoke`
- `grant-cloud` / `revoke-cloud`
- `grant-secret` / `revoke-secret`

Assessment:
- one of the cleanest symmetry families in the CLI

### Expose / unexpose

- `expose` / `unexpose`

Assessment:
- excellent naming symmetry

### Suspend / resume

- `suspend-relation` / `resume-relation`

Assessment:
- strong and explicit

### Enable / disable

- `enable-user` / `disable-user`
- `enable-command` / `disable-command`

Assessment:
- good, although the command-block noun is abstract

## Acceptable asymmetry

### Destroy vs kill

- `destroy-controller` / `kill-controller`

This is not a complement pair in the normal sense. `kill-controller` is a harsher escape hatch, not an undo or inverse.

Assessment:
- good asymmetry; the stronger verb intentionally communicates danger

### Offer / consume

- `offer` / `consume`
- `remove-offer` / `remove-saas`

This is semantically paired rather than lexically paired.

Assessment:
- good domain language, though less mechanically discoverable than simple inverses

### Register / unregister

- `register` / `unregister`

Assessment:
- clean and conventional

## Partial symmetry and gaps

### Show / list coverage gaps

For many nouns Juju offers both list and detail commands:
- `models` / `show-model`
- `users` / `show-user`
- `spaces` / `show-space`
- `offers` / `show-offer`

But some nouns are only partially covered:
- `subnets` exists, but no `show-subnet`
- `storage-pools` exists, but no `show-storage-pool`
- `firewall-rules` exists, but no `show-firewall-rule`
- `secret-backends` and `show-secret-backend` are good; `resources` has no `show-resource`

Assessment:
- list/show symmetry is present as a pattern, but not completed consistently across all nouns

### Create / update / remove families

Storage pools are strong:
- `create-storage-pool`
- `update-storage-pool`
- `remove-storage-pool`
- but no `show-storage-pool`

Secrets backends are also strong:
- `add-secret-backend`
- `update-secret-backend`
- `remove-secret-backend`
- `show-secret-backend`
- `secret-backends`

Clouds are partial:
- `add-cloud`
- `update-cloud`
- `remove-cloud`
- `show-cloud`
- `clouds`

Assessment:
- some families are impressively complete; others stop just short of full symmetry

### Config triad symmetry

There is a coherent config family:
- `config` for apps
- `model-config`
- `controller-config`

There is also a coherent constraints family:
- `constraints`
- `set-constraints`
- `model-constraints`
- `set-model-constraints`

Assessment:
- strong family resemblance
- app config keeps the noun implicit while model/controller config make it explicit, which is slightly asymmetrical but still understandable

## Notable asymmetries that likely hurt discoverability

### Trust has no peer `untrust` command

Today trust revocation is folded into `trust` flags rather than a visible complement command.

Effect:
- the operator must know trust is both set and unset through one verb
- this breaks the otherwise strong paired-verb style present elsewhere

### Integrate is paired with remove-relation, not a direct lexical inverse

`integrate` is a good forward verb, but the reverse operation uses the older resource noun `relation` instead of something like `disconnect` or `unintegrate`.

Effect:
- the pair is semantically correct but lexically non-obvious

### Cloud defaults are special-cased instead of joining a broader config grammar

`default-region` and `default-credential` do not live under a more systematic `cloud-config` or `cloud-defaults` family.

Effect:
- local-default management feels more ad hoc than model/controller config management

### Bundle and Charmhub artifact verbs are mixed

- `find`, `info`, `download`
- `export-bundle`, `diff-bundle`
- `resources`, `charm-resources`, `attach-resource`

These operations are meaningful, but the verb system is not very symmetrical across the broader artifact domain.

### Access execution verbs are asymmetric

- `ssh`, `scp`, `exec`, `run`

These are all valid commands, but there is no coherent complementary system among them because they target different layers.
That is acceptable from a capability standpoint and still problematic from a naming standpoint.

## Complement pairs that are effectively missing

- `show-storage-pool`
- `show-subnet`
- `show-firewall-rule`
- visible `untrust`
- visible inverse or alias for `integrate` / `remove-relation`
- some kind of `plan` or `preview` command family for high-impact mutators besides `deploy --dry-run`

## Assessment

The Juju CLI already contains enough symmetry to be teachable:
- add/remove
- grant/revoke
- enable/disable
- expose/unexpose
- suspend/resume
- list/show in many domains

The main issue is incompleteness, not absence. A focused alias and naming pass could make the command surface feel much more systematic without changing its underlying capability model.
