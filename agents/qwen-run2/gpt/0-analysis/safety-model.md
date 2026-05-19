# Safety Model

## Summary

qwen36 has very little destructive surface. None of the observed commands delete user data or remove snap components. The main safety concerns are operational rather than destructive:

- selecting an engine that the host cannot support
- mutating runtime config to values that prevent the daemon from starting
- starting chat before the server is ready

The command set therefore relies more on validation, readiness gating, and recoverability than on confirmations, force flags, or dry runs.

## Safety profile by command

| Command | Mutates state? | Potential risk | Confirmation / gating | Dry run | Recovery path |
|---|---|---|---|---|---|
| `qwen36 chat` | no persistent mutation observed | user confusion if server is not ready or unhealthy | gated by `wait-for-server.sh` health polling | none observed | fix server state, then rerun `chat` |
| `qwen36 use-engine` | yes | selecting an incompatible engine or overwriting engine-derived config | `--assume-yes` implies a confirmation path may exist by default | none observed | rerun `use-engine` with a different selection or `--auto` |
| `qwen36 show-engine` | no | low risk; inspection only | none needed | n/a | n/a |
| `qwen36 get` | no | low risk; inspection only | none needed | n/a | n/a |
| `qwen36 set` | yes | misconfiguring host, port, verbosity, or CUDA layer count | no confirmation or validation surface is visible from checked-in scripts | none observed | rerun `set` with corrected value or rerun `use-engine` to restore engine-managed keys |
| `qwen36 completion bash` | no | low risk; shell integration only | none needed | n/a | regenerate completion output |
| `qwen36.server` | no user mutation, but starts service process | failed startup or repeated service errors if config/components are wrong | gated by component wait loop and engine-specific preflight checks | none observed | install missing components, fix config, restart service |

## Concrete safety mechanisms

### Readiness gate before chat

`apps/wait-for-server.sh` prevents the chat path from connecting blindly. It waits for:

- engine selection to resolve
- `llama-server` process to exist
- port to be open
- completions API to respond successfully

This reduces false starts and gives explicit guidance on timeout or failure.

### Component availability gate before daemon startup

`apps/server.sh` waits up to 3600 seconds for required engine/model components. On timeout it:

- prints the missing component list
- instructs the user to inspect `snap changes`
- stops the service to avoid indefinite retries by systemd

This is a meaningful safety measure against noisy restart loops.

### Engine-specific preflight checks

The engine launchers verify that:

- model component directory exists
- mmproj component directory exists
- server component directory exists
- for CUDA, `gpu-layers` is set

Those checks fail early with explicit messages instead of launching a partially configured server.

## What the CLI does not provide

### No dry-run support

No observed command offers a preview mode such as `--dry-run`.

Practical consequence:

- users cannot preview which keys `use-engine` will overwrite
- users cannot validate a `set` operation without actually applying it

### No force semantics

There is no observed `--force` flag.

That is acceptable for the current surface because there are no destructive operations, but it means risky config writes depend entirely on validation rather than an explicit “I know what I’m doing” gate.

### No built-in rollback command

There is no `unset`, `reset`, `status`, or `rollback-engine` command.

Recovery is manual and currently relies on:

- running `qwen36 use-engine` again
- overriding values with `qwen36 set`
- using snap service tooling such as `snap logs` / `snap restart`

## Safety-relevant edge cases

### CUDA requires `gpu-layers`

If CUDA is selected but `gpu-layers` is absent, the server fails immediately with:

`gpu-layers snap option is not set`

This is safe failure rather than unsafe fallback, but it exposes the user to a config sharp edge.

### Optional `model-name`

`model-name` is read opportunistically and ignored when absent. That is safe in the sense that chat can proceed without it, but the behavior is implicit and not documented.

### Empty `http.base-path`

The wrapper scripts silently substitute `v1` when `http.base-path` is empty. That fallback is safe and pragmatic, though hidden.

## Safety assessment

The current safety model is appropriate for a non-destructive inference snap: it prefers hard preflight failures and readiness polling over prompts and rollback machinery. The weakest area is configuration mutation. `set` and `use-engine` can change important runtime behavior, but the user-visible safety contract around validation, confirmation, and recovery is not documented.

## Low-cost safety improvements

- Add a documented readback pattern after mutations: `use-engine`/`set` followed by `show-engine` or `get`.
- Add an `unset` or reset command for engine-managed config keys.
- Add a `status` command so users do not need to infer daemon health from chat failures.
- Document which keys are safe to override directly and which are engine-managed.