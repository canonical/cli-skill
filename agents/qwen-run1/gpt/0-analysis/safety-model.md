# qwen36 Safety Model

## Overview

The observable CLI is low on destructive operations in the data-loss sense, but several commands can change service behavior in ways that affect availability, hardware usage, and script assumptions.

## Mutating Commands

| Command | Mutates State | Safety Mechanisms | Recovery Path |
|---|---|---|---|
| `qwen36 chat` | No persistent mutation observed | None needed | Exit the interactive client |
| `qwen36 use-engine` | Yes, changes selected engine and likely derived configuration | `--assume-yes` implies an interactive confirmation step exists by default | Re-run `use-engine` with a different selection or `--auto` |
| `qwen36 show-engine` | No | None needed | N/A |
| `qwen36 get` | No | None needed | N/A |
| `qwen36 set` | Yes, mutates configuration | No dry-run or force flag observed; `--package` appears restricted to packaging flows | Re-run `set` with a corrected value |
| `qwen36 completion bash` | No | None needed | N/A |

## Main Safety Concerns

### Engine switching affects service readiness

Selecting an engine changes which snap components and runtime parameters are required. A bad selection can lead to:

- server startup failure
- missing component errors
- incompatible hardware usage
- mis-sized GPU offload settings

This is the main operational risk in the CLI.

### Config mutation has no documented validation contract

Because `set` is undocumented beyond one README example, users do not know:

- which keys are safe to change
- which values are validated immediately
- which changes require a service restart
- whether invalid values are rejected at write time or only at runtime

### Hidden package layer

`set --package` introduces a privileged or package-owned config path that users are not taught about. That is not unsafe by itself, but it makes debugging harder because effective values may come from an invisible source.

## Confirmations, Dry-Run, And Force

- Confirmation: implied for `use-engine` because `--assume-yes` exists
- Dry-run: no evidence for any dry-run mode
- Force flags: no evidence for force semantics
- Rollback: no dedicated rollback command; rollback is manual through another `use-engine` or `set`

## Recovery Behavior

Operational recovery is delegated to wrappers rather than the CLI:

- `apps/server.sh` waits for required components and stops the service on timeout
- `apps/wait-for-server.sh` polls server health before interactive usage
- `apps/check-server-llamacpp.sh` distinguishes healthy, starting, and failed states

That is useful for service robustness, but it does not replace explicit CLI guidance about safe mutations.

## Safety Documentation Gaps

- No warning text is documented for engine switches.
- No list of restart-sensitive config keys is published.
- No statement explains whether config writes are transactional.
- No recovery command exists for restoring package defaults or engine defaults.