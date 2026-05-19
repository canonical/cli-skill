# Extensibility Model

## Summary

The qwen36 snap is extensible mainly through packaging and engine metadata, not through a plugin API. The most mature extension seam is engine addition: each engine lives in its own directory with a manifest, launcher, and health check, and the selected engine is exposed through `show-engine` and `get`. Extending the command set itself requires changes in the private Go CLI submodule, which is not present here.

## Extension boundaries

| Area | How it extends today | Boundary |
|---|---|---|
| Engine variants | add a new `engines/<name>/` directory and manifest, plus any required snap parts/components | bounded by engine manifest schema and CLI support for selection |
| Runtime components | add or replace snap components for model, projector, or engine binaries | bounded by snapcraft parts/components and component init scripts |
| Config surface | add new persisted keys consumed by scripts or engine launchers | bounded by Go CLI `get`/`set` support and script adoption |
| User commands | register new commands/subcommands in the Go CLI | bounded by the private Go submodule |
| Shell completion | update CLI completion generator output consumed by `apps/completion.bash` | bounded by command registration and completion implementation in Go |

## How a new engine is added

A new engine extension would need to provide all of these pieces:

1. `engines/<name>/engine.yaml`
   - declares engine name
   - hardware requirements
   - required components
   - engine-specific config defaults under `configurations:`

2. `engines/<name>/common.sh`
   - resolves `server`, `model`, and `multimodel-projector` from persisted config
   - sources component init scripts and exports runtime paths

3. `engines/<name>/server`
   - reads runtime config through `qwen36 get`
   - translates config into concrete `llama-server` flags or another engine executable

4. `engines/<name>/check-server`
   - conforms to the implicit health-check protocol consumed by `wait-for-server.sh`

5. snapcraft additions
   - new parts/components in `snap/snapcraft.yaml`
   - any required stage packages or plugs

6. CLI integration
   - `qwen36 use-engine` must recognize the new engine name or auto-detection case
   - `show-engine` must surface the manifest in the same YAML contract

## Registration and discovery paths

### Engine discovery

The runtime discovers only the selected engine, not all engines dynamically. Discovery path:

1. `use-engine` selects/persists an engine.
2. `show-engine` exposes the selected manifest.
3. `server.sh` reads `show-engine` output to find the active engine name and required components.
4. The daemon execs `engines/<name>/server` based on that name.

This is a manifest-driven dispatch model, not a plugin scan of arbitrary directories at runtime.

### Command discovery

Command discovery is handled by the Go CLI and exposed indirectly through:

- the user-facing command parser
- the `completion bash` generator
- the `bin/qwen36` snap app entrypoint

Because the command registration code is missing from this checkout, command extension cannot be implemented or audited fully here.

## Middleware and hooks

There is no evidence of a formal middleware stack inside the CLI from the checked-in sources. The closest analogues are:

- snap hooks (`install`, `post-refresh`) that run before normal user workflows
- wrapper scripts (`chat.sh`, `server.sh`, `wait-for-server.sh`) that sit between persisted config and runtime behavior
- health check protocol (`check-server`) that standardizes engine readiness behavior

This is closer to script orchestration than a plugin or interceptor framework.

## Stable internal contracts that extensions must preserve

| Contract | Why it matters |
|---|---|
| `show-engine` returns YAML with `name` and `components[]` | daemon startup depends on it |
| `get` returns raw scalar stdout | all shell wrappers depend on it |
| `check-server` returns `0/1/2` with ready/starting/failed semantics | `wait-for-server.sh` depends on it |
| component init scripts export `MODEL_FILE` / `MMPROJ_FILE` | engine launchers depend on it |
| engine config keys use the existing snap-config path | launchers and hooks depend on persisted config |

## What is easy to extend

- add or tune engine launch flags in engine-specific server scripts
- add new engine config keys and consume them in scripts
- change model or projector artifacts by updating components and init scripts
- add documentation and examples around existing commands

## What is hard to extend

- add new user commands without the private Go submodule
- change command grammar without updating completion generation and docs in lockstep
- change `show-engine` or `get` output shape without breaking shell consumers
- add a new runtime health-check model without preserving the `0/1/2` protocol

## Recommended extension approach

### For runtime features

Prefer extending the manifest-and-script path:

- add a config key
- read it in the relevant engine launcher
- keep output contracts unchanged

### For command-surface features

Prefer adding commands in a backward-compatible way through the Go CLI:

- keep existing verbs working
- add new help/completion support
- update README and examples together

### For engine growth

Keep the existing per-engine directory convention. It is the clearest extension seam in the repository and already encodes the correct startup responsibilities.

## Assessment

The qwen36 snap is extensible, but mostly as a packaged runtime rather than as a general-purpose CLI framework. The cleanest seam is engine addition through `engines/<name>/...` plus snapcraft metadata. The weakest seam is command evolution, because the public repo does not include the registration logic for the actual `qwen36` binary.