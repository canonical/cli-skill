# qwen36 CLI Architecture

## Summary

`qwen36` is a snap-packaged Go CLI built with Cobra that controls a local inference service for the Qwen3.6 model. The shipped system is split across:

- a Go command binary under `cli/cmd/cli`
- snap packaging and lifecycle hooks in `snap/`
- shell wrappers in `apps/` that start, probe, and bridge the server
- engine manifests and launchers in `engines/cpu` and `engines/cuda`
- snap components for model weights, projector weights, and inference binaries

## Primary Architecture Style

Primary style: Layered CLI application

The main layers are:

1. Command layer: Cobra commands in `cli/cmd/cli/commands` and `cli/cmd/cli/commands/debug`
2. Policy and shared behavior layer: helpers in `cli/cmd/cli/common`
3. Configuration and cache layer: `cli/pkg/storage` using `snapctl`
4. Engine abstraction layer: `cli/pkg/engines`, `cli/pkg/selector`, and manifest-driven engine metadata
5. Runtime layer: shell scripts in `apps/` and `engines/*/server`

## Secondary Architecture Styles

Secondary style: Client-server CLI

- `chat` and `webui` are clients for a local OpenAI-compatible HTTP server.
- The actual model server runs as the snap daemon app `server`.
- The CLI configures and interrogates that daemon rather than embedding inference in the command process.

Secondary style: Manifest-driven component host

- Engine behavior is not hardcoded in one place.
- Engine manifests declare hardware requirements, required components, and default configuration.
- `use-engine` and `list-engines` discover manifests at runtime and score them against the host.

## Key Control Flow

### Install and first-run path

1. `snap/hooks/install` seeds package defaults with `qwen36 set --package`.
2. The same hook tries `qwen36 use-engine --auto --assume-yes` when hardware inspection is available.
3. That writes the active engine cache entry and engine-level config.

### Server startup path

1. The `server` daemon runs `apps/server.sh`.
2. That script calls `qwen36 show-engine`, waits for required components, and resolves the selected engine name.
3. It then execs `engines/<engine>/server`.
4. The engine launcher sources component init scripts and starts `llama-server` with config-derived host, port, and model parameters.

### Chat path

1. `chat` resolves the OpenAI endpoint from engine component settings plus config.
2. It verifies the `server` service is active.
3. It uses the shared chat client to wait for model readiness and then starts an interactive readline session.

## Runtime Boundaries

### Snap-specific boundary

The CLI is tightly bound to snapd:

- config is persisted with `snapctl set/get/unset`
- service status comes from `snapctl services`
- component installation and removal go through `snapctl install-components` and `snapctl remove-components`
- environment discovery depends on `SNAP`, `SNAP_INSTANCE_NAME`, `SNAP_REVISION`, and `SNAP_COMPONENTS`

### Feature-gated boundary

Top-level `chat` and `webui` are added only when `ADDITIONAL_FEATURES` contains `chat` or `webui`. The reusable CLI submodule supports that gating, but the shipped qwen36 snap manifest does not currently set `ADDITIONAL_FEATURES`, so the built snap defaults to a 10-command public surface even though the source and README imply a 12-command surface.

## Architectural Strengths

- Engine capabilities are separated cleanly from command parsing.
- Config precedence is explicit in code: package < engine < user.
- Snap hooks reuse the CLI instead of reimplementing config logic in shell.
- Hidden debug commands provide a useful developer lane without cluttering normal help.

## Architectural Risks

- The product depends heavily on snapd semantics, reducing portability and testability outside a snap.
- Service management is partly outside the CLI surface, so users must switch between `qwen36` and `snap` commands.
- `chat` and `webui` feature exposure is inconsistent between source, README, and shipped snapcraft config.
- Shell wrappers depend on undocumented output contracts from `show-engine` and `get`.
