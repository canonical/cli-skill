# Extensibility Model

## Overview

The `qwen36` CLI is designed with a **data-driven extensibility model** centered on engine manifests. New hardware backends and configurations can be added without modifying the CLI binary. However, adding entirely new **commands** to the CLI requires modifying the Go source code — there is no dynamic command registration or external plugin system.

## Extension Points

### 1. Engine Manifests (Primary Extension Mechanism)

**What**: YAML files in `$SNAP/engines/` defining hardware requirements, required components, and default configurations.

**How to add a new engine**:
1. Create a new YAML file in the engines directory.
2. Define: `name`, `description`, `vendor`, `grade`, `devices` (CPU/PCI/USB requirements), `memory`, `disk-space`, `components`, `configurations`.
3. Package it with the snap or add it as a snap component.

**Discovery**: Engines are discovered at runtime via `engines.LoadManifests(enginesDir)`, which reads all files in the engines directory. No explicit registration step.

**Example engine manifest structure**:
```yaml
name: cpu
description: CPU-based inference
vendor: llama.cpp
grade: stable
devices:
  anyof:
    - type: cpu
      architecture: amd64
      flags: [avx2]
memory: "16GiB"
disk-space: "25GiB"
components:
  - llamacpp
  - model-qwen36-35b-a3b-ud-q4-k-xl
  - mmproj-qwen36-35b-a3b-f16
configurations:
  http.port: "8080"
```

**Boundaries**:
- Engines can declare hardware requirements via device descriptors (CPU architecture, flags, PCI vendor/device IDs, GPU microarchitecture, compute capability, VRAM).
- Engines can reference snap components by name and set configuration values.
- Engines **cannot** add new CLI commands, modify command behavior, or inject middleware.

### 2. Snap Components

**What**: Independently installable snap components (e.g., `llamacpp`, `llamacpp-cuda`, model weights, multimodal projector).

**How to add**: Define a new component in `snap/snapcraft.yaml` under `components:` and provide content. Components are installed/removed via `snapctl InstallComponents`/`RemoveComponents`.

**Discovery**: Component availability is checked at runtime via `common.ComponentInstalled()`, which looks in `/snap/$SNAP_INSTANCE_NAME/components/$SNAP_REVISION/<component-name>/`.

**Component settings**: Each component can include a `component.yaml` file with:
- `servers`: Map of server names to settings (`protocol`, `base-path`).
- `environment`: Array of `KEY=VALUE` environment variable declarations (supports `$COMPONENT` expansion).
- `layout`: Symlink definitions for filesystem layout adjustments.

**Boundaries**: Components declare what they provide; the CLI consumes component settings via `common.EngineComponentSettings()`. Components cannot modify CLI behavior directly.

### 3. Configuration Passthrough

**What**: User-defined environment variables passed through to subprocesses via `qwen36 set passthrough.environment.<key>=<value>`.

**How it works**: The `run` command reads all `passthrough.environment.*` keys, converts them to `UPPER_SNAKE_CASE` environment variables, and injects them into the subprocess environment.

**Boundaries**: Only affects `run` command. Passthrough keys are not validated beyond the key=value format.

### 4. Feature Gating via Environment Variable

**What**: The `ADDITIONAL_FEATURES` environment variable controls whether `chat` and `webui` commands are available.

**Values**: Comma-separated list — `"chat"` and/or `"webui"`. Read once at startup via `common.AdditionalFeatures()`.

**Use case**: Different inference snaps can ship the same CLI but enable different features based on their packaging. The `qwen36` snap ships with both features enabled; other snaps may enable only one or neither.

**Boundaries**: Binary toggle only — no per-feature configuration, no runtime reconfiguration.

### 5. Chat Client (Go External Dependency)

**What**: The `go-chat-client` package (`github.com/jpm-canonical/go-chat-client`) provides the interactive chat interface.

**Integration**: Called from `common.ChatClient(baseUrl, modelName, verbose)` which wraps the external library.

**Extensibility**: The chat client itself is not extensible from the CLI side. Changing chat behavior requires modifying the external library or replacing it.

### 6. Web UI Configuration (`serve-webui`)

**What**: The `serve-webui` command serves static files with a `/config.json` endpoint.

**Configuration endpoint** provides:
- `openaiBaseURL`: Resolved from active engine.
- `capabilities`: Comma-separated list (default: `text`; supported capabilities from `webui.SupportedCapabilities()`).
- `instanceName`: Snap instance name.
- `engineName`: Active engine name.

**Extensibility**: New capabilities can be added to `webui.SupportedCapabilities()`. The web UI consumes these capabilities to enable/disable features.

## What Cannot Be Extended Without Code Changes

| Extension Goal | Currently Possible? | What Would Need to Change |
|---------------|---------------------|--------------------------|
| Add new top-level command | No | Modify `cmd/cli/main.go` to register new Cobra commands |
| Add new debug subcommand | No | Modify `cmd/cli/commands/debug/debug.go` |
| Add new config source/tier | No | Modify `pkg/storage/config.go` — precedence order is hardcoded |
| Add new output format | No | Per-command format handling is explicit (`switch cmd.format`) |
| Add new hardware probe | No | Modify `pkg/hardware_info/` |
| Add middleware/hooks for commands | No | No hook infrastructure exists |
| Customize help text per engine | No | Help text is static Go strings |
| Auto-discover commands from files | No | All commands are compiled-in |

## Extension Registration Path Summary

```
New Engine:
  YAML file → $SNAP/engines/ → engines.LoadManifests() → selector.ScoreEngines() → CLI

New Component:
  snapcraft.yaml → snap component → snapctl InstallComponents → $SNAP_COMPONENTS/<name>/ → CLI reads component.yaml

New Passthrough Config:
  qwen36 set passthrough.environment.X=Y → snapctl storage → Config.Get("passthrough") → run command

New Feature (chat/webui):
  env ADDITIONAL_FEATURES=chat,webui → common.AdditionalFeatures() → main.go conditional registration
```

## Architectural Notes

### Engine Manifest Validation

`debug validate-engines` provides a way to check engine manifests without installing them. This supports the engine author workflow: create manifest → validate → test with `debug select-engine`.

### Testing Pipeline

The `debug select-engine` command reads machine info from stdin (produced by `show-machine --format=yaml`), enabling offline testing:
```bash
qwen36 show-machine --format=yaml > machine.yaml
cat machine.yaml | qwen36 debug select-engine --engines test_data/engines/
```

This is a key extensibility feature for engine developers — they can test engine selection against real hardware data without running on that hardware.

### No Plugin Binary Loading

The CLI uses only data-driven extensibility (YAML manifests). There is no dynamic library loading, no scripting, no IPC-based plugin protocol, and no external process spawning for plugins. This is appropriate for a snap-packaged tool where confinement makes binary plugins difficult.
