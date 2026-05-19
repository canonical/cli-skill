# qwen36 Safety Model

## Overview

The observable CLI is low on destructive operations in the data-loss sense, but several commands can change service behavior in ways that affect availability, hardware usage, and script assumptions.

## Mutating Commands

| Command | Mutates State | Safety Mechanisms | Recovery Path |
|---------|---------------|-------------------|---------------|
| `qwen36 chat` | No persistent mutation observed | None needed | Exit the interactive client |
| `qwen36 use-engine` | Yes, changes selected engine and likely derived configuration | `--assume-yes` implies an interactive confirmation step exists by default | Re-run `use-engine` with a different selection or `--auto` |
| `qwen36 show-engine` | No | None needed | N/A |
| `qwen36 get` | No | None needed | N/A |
| `qwen36 set` | Yes, mutates configuration | No dry-run or force flag observed; `--package` appears restricted to packaging flows | Re-run `set` with a corrected value |
| `qwen36 completion bash` | No | None needed | N/A |

## Main Safety Concerns

### Engine Switching Affects Service Readiness

Selecting an engine changes which snap components and runtime parameters are required. A bad selection can lead to:

- Server startup failure
- Missing component errors
- Incompatible hardware usage
- Mis-sized GPU offload settings

This is the main operational risk in the CLI.

### Config Mutation Has No Documented Validation Contract

Because `set` is undocumented beyond one README example, users do not know:

- Which keys are safe to change
- Which values are validated immediately
- Which changes require a service restart
- Whether invalid values are rejected at write time or only at runtime

### Hidden Package Layer

`set --package` introduces a privileged or package-owned config path that users are not taught about. That is not unsafe by itself, but it makes debugging harder because effective values may come from an invisible source.

### No Rollback Mechanism

There is no documented command to:
- Reset configuration to defaults
- Undo a configuration change
- Revert to a previous engine selection

## Confirmations, Dry-Run, And Force

| Mechanism | Present? | Notes |
|-----------|----------|-------|
| Confirmation | Implied for `use-engine` | `--assume-yes` flag exists, suggesting default behavior prompts |
| Dry-run | No | No evidence for any dry-run mode |
| Force flags | No | No evidence for force semantics |
| Rollback | No | No dedicated rollback command; rollback is manual through another `use-engine` or `set` |

## Recovery Behavior

Operational recovery is delegated to wrappers rather than the CLI:

- `apps/server.sh` waits for required components and stops the service on timeout
- `apps/wait-for-server.sh` polls server health before interactive usage
- `apps/check-server-llamacpp.sh` distinguishes healthy, starting, and failed states

That is useful for service robustness, but it does not replace explicit CLI guidance about safe mutations.

## Destructive Operations

| Operation | Destructive? | Confirmation | Recovery |
|-----------|--------------|--------------|----------|
| Change engine | Service availability only | Implied via `--assume-yes` flag | Run `use-engine` again |
| Change config | Service behavior only | None | Run `set` again with correct value |
| Start chat session | No | None | Exit client |
| View engine | No | None | N/A |
| Get config | No | None | N/A |

## Safety Documentation Gaps

- No warning text is documented for engine switches.
- No list of restart-sensitive config keys is published.
- No statement explains whether config writes are transactional.
- No recovery command exists for restoring package defaults or engine defaults.
- No validation errors are documented for `set` operations.
- No documentation on what happens if user switches to incompatible engine (e.g., CUDA without GPU).

## Recommendations

1. Add explicit confirmation prompts for `use-engine` with clear warnings about service disruption.
2. Document which config keys require a service restart.
3. Add an `unset` command to restore individual keys to defaults.
4. Add a `reset` command to restore all config to defaults.
5. Add validation for `set` operations with clear error messages.
6. Document the behavior of `use-engine --auto` when no suitable engine is found.
