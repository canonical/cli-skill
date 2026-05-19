# Configuration Model

## Summary

qwen36 uses snapd configuration as its single persistent control plane. The observable runtime never reads a config file and never accepts environment variables as authoritative user input. Instead, install hooks and user commands write snap config, and every runtime script reads effective values through `qwen36 get` immediately before use.

## Configuration sources

| Source | How values enter the system | Scope | Evidence |
|---|---|---|---|
| Install hook defaults | `qwen36 set --package <key>=<value>` | package-provided baseline | `snap/hooks/install` sets `http.port`, `http.host`, and `verbose` |
| Engine manifest defaults | `engines/*/engine.yaml` `configurations:` block, applied by `use-engine` | engine-specific persisted settings | CPU and CUDA manifests define `server`, `model`, `multimodel-projector`, `http.base-path`, and CUDA-only `gpu-layers` |
| User config writes | `qwen36 set <key>=<value>` | persisted effective config | README documents `qwen36 set http.port=8326` |
| Engine selection writes | `qwen36 use-engine ...` | persisted effective config | install hook uses `use-engine --auto --assume-yes`; runtime depends on resulting `show-engine`/`get` values |
| Runtime fallbacks in scripts | shell-level defaults after `get` | process-local only, not persisted | `chat.sh` and `check-server-llamacpp.sh` substitute `v1` when `http.base-path` is empty |

## Effective precedence

The exact internal layering inside the missing Go submodule is not directly visible, but the effective behavior from the checked-in scripts is:

1. Persist defaults during install.
2. Persist engine-specific settings when `use-engine` runs.
3. Allow later writes through `qwen36 set` or another `use-engine` invocation to overwrite relevant keys.
4. Resolve the final value at runtime through `qwen36 get`.
5. Apply limited script-local fallback logic only when a retrieved value is empty.

In operational terms, this is a last-writer-wins model over snap config, followed by script-local fallback for a small number of keys.

## Known configuration keys

| Key | Purpose | Observed readers | Observed writers / source |
|---|---|---|---|
| `http.port` | TCP port for local inference API | `chat.sh`, `check-server-llamacpp.sh`, `engines/cpu/server`, `engines/cuda/server` | install hook default; user `set` |
| `http.host` | bind address for `llama-server` | `engines/cpu/server`, `engines/cuda/server` | install hook default; user `set` |
| `http.base-path` | API base path segment | `chat.sh`, `check-server-llamacpp.sh` | engine manifest `configurations:` via `use-engine`; user `set` likely possible |
| `verbose` | toggles `--verbose` for engine launch | `engines/cpu/server`, `engines/cuda/server` | install hook default; user `set` |
| `server` | selected engine component name | `engines/*/common.sh` | engine manifest via `use-engine` |
| `model` | selected model component name | `engines/*/common.sh` | engine manifest via `use-engine` |
| `multimodel-projector` | selected MM projector component name | `engines/*/common.sh` | engine manifest via `use-engine` |
| `gpu-layers` | number of GPU layers for CUDA launch | `engines/cuda/server` | CUDA manifest via `use-engine`; user `set` likely possible |
| `model-name` | optional model identifier for chat/OpenAI request | `chat.sh`, `check-server-llamacpp.sh` | not observed being written in this repository |

## Command-specific behavior

### `qwen36 use-engine`

`use-engine` is both a selection command and a config mutator. Its observable effect is broader than just remembering an engine name:

- it determines which engine manifest is considered current
- it indirectly decides which components must be installed before the daemon can start
- it seeds engine-dependent config keys such as `server`, `model`, `multimodel-projector`, `http.base-path`, and for CUDA `gpu-layers`

This means `use-engine` acts like a higher-level config macro over several lower-level keys.

### `qwen36 set`

`set` is the direct low-level escape hatch. It appears able to override runtime keys individually, even when they are also managed indirectly by `use-engine`.

Examples of realistic overrides:

- `http.port`
- `http.host`
- `verbose`
- `gpu-layers`

## Runtime read model

The runtime is strictly late-bound:

- `chat.sh` reads current values when the user starts chat
- `server.sh` reads current engine YAML when the daemon starts
- engine launchers read current config keys when they exec `llama-server`
- health checks read current endpoint settings on each probe

There is no evidence of hot reload after launch. Changing config affects the next process start or restart, not an already-running `llama-server` process.

## Fallbacks and surprising behavior

### Empty `http.base-path` becomes `v1`

Both `chat.sh` and `check-server-llamacpp.sh` do this:

- call `qwen36 get http.base-path`
- if the value is empty, substitute `v1`

That fallback is not stored back into config. It is a command-local runtime default.

### `model-name` is optional and error-tolerant

Scripts read `model-name` like this:

```bash
model_name="$(qwen36 get model-name 2>/dev/null || true)"
```

That implies one of two things:

- the key may legitimately be unset
- `qwen36 get` may emit an error for missing keys

Either way, `model-name` is treated as optional config, unlike `http.port`.

### Component paths are config, not discovery

The runtime does not discover model or engine artifacts dynamically. It reads component names from persisted config (`server`, `model`, `multimodel-projector`) and constructs paths under `$SNAP_COMPONENTS`.

## What is not part of the configuration model

- no user config file such as `$HOME/.config/qwen36/...`
- no documented environment variables for overriding config values at invocation time
- no command-specific flag overrides for host, port, verbosity, or model
- no observed profile, workspace, or per-project config scopes

## Assessment

The configuration model is simple and operationally coherent: one persisted store, one read command, one write command, and one higher-level engine-selection command. Its main downside is discoverability. Users must know internal key names, and some critical behavior such as `http.base-path` fallback and the optional nature of `model-name` is only apparent from scripts rather than user docs.