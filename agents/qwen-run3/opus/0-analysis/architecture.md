# Architecture

## Tech Stack

- **Language**: Go 1.24
- **CLI Framework**: [cobra](https://github.com/spf13/cobra) (command registration, flag parsing, completions)
- **Packaging**: Snap (strict confinement, components for engines/models)
- **Inference Engine**: llama.cpp (C/C++, built via CMake, shipped as snap components)
- **Configuration Store**: snapctl (snap daemon configuration layer)
- **Chat Client**: External Go binary (`go-chat-client`)
- **Shell Scripts**: Bash wrappers for server lifecycle, health checks, and completion

## Architecture Style

**Primary**: Client-server CLI

The `qwen36` CLI acts as a management client for a local inference server (`llama-server`). The server runs as a snap daemon service, exposing an OpenAI-compatible HTTP API. The CLI configures, starts, and queries the server but does not perform inference itself.

**Secondary**: Layered CLI application

The CLI is internally layered:

1. **Command layer** (`cmd/cli/commands/`): Cobra command definitions with flags and validation
2. **Common layer** (`cmd/cli/common/`): Shared context, spinners, service helpers
3. **Package layer** (`pkg/`): Domain packages for engine management, hardware detection, snap store interaction, storage/config

## Component Architecture

The snap uses **snap components** to modularly deliver:

- **Engine components** (`llamacpp`, `llamacpp-cuda`): Server binaries specific to hardware
- **Model components** (`model-qwen36-35b-a3b-ud-q4-k-xl`): Quantized model weights (~22 GB)
- **Projector components** (`mmproj-qwen36-35b-a3b-f16`): Multimodal projectors

Engine selection is mediated by YAML manifests (`engines/<name>/engine.yaml`) that declare hardware requirements (CPU flags, GPU vendor, VRAM, memory, disk). The CLI scores engines against detected hardware to auto-select or validate manual selections.

## Data Flow

```
User → qwen36 CLI → snapctl (config) → snap daemon (llama-server) → OpenAI API (localhost:8326)
                  → engines/*.yaml (manifests)
                  → $SNAP_COMPONENTS (model/engine/projector files)
```
