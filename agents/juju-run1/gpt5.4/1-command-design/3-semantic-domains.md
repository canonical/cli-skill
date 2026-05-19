# Juju Semantic Domains

## Overview

The Juju command surface clusters into a handful of strong operator domains. These domains matter more than source packages when evaluating discoverability, because users think in terms of tasks and system areas, not Go package names.

## 1. Controller and model administration

Primary commands:
- `bootstrap`
- `add-model`
- `models`
- `show-model`
- `model-config`
- `model-defaults`
- `model-constraints`
- `set-model-constraints`
- `destroy-model`
- `migrate`
- `controllers`
- `show-controller`
- `controller-config`
- `destroy-controller`
- `kill-controller`
- `enable-destroy-controller`
- `register`
- `unregister`
- `switch`
- `whoami`

Mental model:
- where am I connected
- what controller/model am I operating on
- how is that scope configured
- can I move, destroy, or recover that scope

Comments:
- this is one of Juju's strongest and most mature domains
- it is also one of the most terminology-heavy

## 2. Application lifecycle and topology

Primary commands:
- `deploy`
- `refresh`
- `config`
- `add-unit`
- `remove-unit`
- `remove-application`
- `scale-application`
- `show-application`
- `show-unit`
- `status`
- `resolved`
- `retry-provisioning`
- `trust`

Mental model:
- deploy something
- configure it
- scale it
- inspect it
- repair or remove it

Comments:
- this is the dominant operator workflow domain
- naming is fairly good, but execution and mutation semantics overlap with other domains

## 3. Integration and cross-model connectivity

Primary commands:
- `integrate`
- `remove-relation`
- `suspend-relation`
- `resume-relation`
- `offer`
- `offers`
- `show-offer`
- `find-offers`
- `remove-offer`
- `consume`
- `remove-saas`

Mental model:
- connect applications locally or across models
- publish or consume shared endpoints
- suspend or resume connectivity

Comments:
- this domain is semantically coherent
- `offer` / `consume` is one of the better domain-specific verb pairs in the CLI

## 4. Cloud, credential, and substrate management

Primary commands:
- `clouds`
- `show-cloud`
- `add-cloud`
- `update-cloud`
- `remove-cloud`
- `regions`
- `credentials`
- `add-credential`
- `update-credential`
- `remove-credential`
- `show-credential`
- `default-region`
- `default-credential`
- `autoload-credentials`
- `add-k8s`
- `update-k8s`
- `remove-k8s`
- `set-credential`

Mental model:
- what cloud inventory is known
- what credentials exist locally and remotely
- what defaults should later commands use

Comments:
- functionally rich but lexically mixed
- some client-local concepts and controller-side concepts are adjacent in the same domain

## 5. Compute, machine, and remote access

Primary commands:
- `add-machine`
- `machines`
- `show-machine`
- `remove-machine`
- `ssh`
- `scp`
- `exec`
- `debug-hooks`
- `debug-code`

Mental model:
- provision or inspect compute targets
- reach them interactively or non-interactively
- debug workloads on those targets

Comments:
- access and execution concepts are powerful but partially overlapping
- this domain contains some of the biggest confusion-pair risks

## 6. Storage management

Primary commands:
- `add-storage`
- `storage`
- `show-storage`
- `remove-storage`
- `detach-storage`
- `attach-storage`
- `import-filesystem`
- `create-storage-pool`
- `storage-pools`
- `remove-storage-pool`
- `update-storage-pool`

Mental model:
- provision, inspect, attach, detach, preserve, and reuse storage resources

Comments:
- this is a comparatively coherent domain
- explicit destroy-vs-release semantics are a strength
- missing detail commands for some nouns weaken symmetry slightly

## 7. Networking and placement

Primary commands:
- `spaces`
- `show-space`
- `add-space`
- `move-to-space`
- `rename-space`
- `remove-space`
- `reload-spaces`
- `subnets`
- `set-firewall-rule`
- `firewall-rules`
- `bind`
- `expose`
- `unexpose`

Mental model:
- network segmentation and endpoint placement
- access exposure and firewall policy
- mapping app endpoints to network topology

Comments:
- semantically real, but spread across several lexical styles
- users likely encounter it as part of app management rather than as a standalone domain

## 8. Identity, access, and operator context

Primary commands:
- `add-user`
- `remove-user`
- `enable-user`
- `disable-user`
- `show-user`
- `users`
- `change-user-password`
- `login`
- `logout`
- `whoami`
- `grant`
- `revoke`
- `grant-cloud`
- `revoke-cloud`
- `add-ssh-key`
- `remove-ssh-key`
- `import-ssh-key`
- `ssh-keys`

Mental model:
- who can access what
- who am I logged in as
- what keys and credentials are trusted

Comments:
- broad but understandable domain
- mixes authentication, authorization, and SSH identity management

## 9. Secrets and secret backends

Primary commands:
- `secrets`
- `show-secret`
- `add-secret`
- `update-secret`
- `remove-secret`
- `grant-secret`
- `revoke-secret`
- `secret-backends`
- `show-secret-backend`
- `add-secret-backend`
- `update-secret-backend`
- `remove-secret-backend`
- `model-secret-backend`

Mental model:
- manage secret values
- manage storage backends for secrets
- control app access to secrets

Comments:
- one of the more self-contained newer domains
- naming quality is relatively strong

## 10. Charm and artifact catalog access

Primary commands:
- `find`
- `info`
- `download`
- `resources`
- `charm-resources`
- `attach-resource`
- `export-bundle`
- `diff-bundle`

Mental model:
- discover charms and bundles
- inspect artifact metadata
- compare bundle intent with live state

Comments:
- catalog and live-resource operations sit close together, which is efficient but semantically mixed

## 11. Actions, tasks, and observability

Primary commands:
- `actions`
- `run`
- `show-action`
- `operations`
- `show-operation`
- `show-task`
- `cancel-task`
- `show-status-log`
- `debug-log`
- `status`

Mental model:
- observe current system state
- invoke or inspect action execution
- review logs and past statuses

Comments:
- internally coherent for advanced users
- externally confusing because `run`, `exec`, `ssh`, and action/task/operation nouns are not obvious to newcomers

## 12. Meta and self-description

Primary commands:
- `help`
- `documentation`
- `version`
- `dashboard`
- `help-action-commands`
- `help-hook-commands`
- `sync-agent-binary`

Mental model:
- learn the CLI
- get out-of-band tooling support
- open companion interfaces

Comments:
- these commands are necessary, but they make the top-level namespace feel even broader

## Domain-level assessment

The best-defined semantic domains are:
- controller/model administration
- application lifecycle
- storage
- secrets
- cross-model offers

The least tidy domains are:
- networking and placement
- compute/access/execution
- action/task/observability
- cloud defaults and credentials

Those are the areas where naming, grouping, and documentation improvements would buy the most operator clarity.
