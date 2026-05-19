# qwen36 Extensibility Model

## Summary

This snap is extensible at the packaging and engine-definition level, but not at the CLI parser level from the public repository because the `cli/` submodule is private and empty here.

## Observable Extension Boundaries

### 1. Engine extension boundary

The cleanest extension point is the `engines/` directory. Each engine currently provides:

- `engine.yaml`
- `common.sh`
- `server`
- `check-server`

The manifest supplies:

- engine name
- descriptive metadata
- hardware requirements
- required snap components
- default configuration values

To add a new engine, the packaging model strongly suggests you would need:

1. a new `engines/<name>/engine.yaml`
2. launcher and health-check scripts under `engines/<name>/`
3. matching component definitions in `snap/snapcraft.yaml`
4. CLI support so `use-engine` can discover or accept the new engine

### 2. Component extension boundary

Snap components already isolate:

- server binaries
- model weights
- multimodal projector weights

That makes it straightforward to swap artifacts without rewriting launcher logic, as long as the CLI and engine manifests agree on component names.

### 3. Configuration key boundary

The shell scripts are decoupled from storage details by using `qwen36 get`. New engine or runtime features can be exposed by:

- adding a new config key
- teaching the CLI to persist and read it
- consuming it in shell wrappers or engine launchers

The current example is `gpu-layers`, which only matters for the CUDA path.

## Non-Observable Extension Boundaries

These likely exist but cannot be audited:

- command registration in the Go CLI
- shared middleware for prompting, validation, and output formatting
- completion generation internals
- config backend abstraction

## Discovery Model

The public repository suggests that runtime discovery happens through selected engine state rather than directory scanning alone:

- `use-engine` chooses an engine
- `show-engine` returns the selected engine description
- `server.sh` trusts `show-engine` rather than enumerating `engines/*`

That means adding a new engine is not just a file-system operation. The CLI must know how to surface it.

## Constraints On Adding New Commands

Without the private `cli/` source, the project cannot publicly:

- add new top-level commands
- adjust parser rules
- improve help text
- expose new completion targets

So the current repo is packaging-extensible but CLI-surface-constrained.

## Recommendation

If this project expects public contributors to add engines or config keys, it should publish at least:

- the engine discovery contract
- the shape of `show-engine` YAML
- the config precedence model
- minimal help output or command grammar reference