# qwen36 Argument Structure

## Introduction

The observable argument design follows three patterns:

1. **Verb-only commands** with no additional arguments: `chat`
2. **Verb-noun commands** with either positional values or flags: `use-engine`, `show-engine`
3. **Config access commands** that take either a positional key or a `key=value` assignment: `get`, `set`

The command set is small, but the argument model is only partially documented because the parser implementation is not available. The table below records everything that can be evidenced from README examples, snap hooks, and shell wrappers.

## Command Map

| Leaf Command | Positional Arguments | Flags | Required | Defaults | Type / Accepted Values | Aliases | Env Var Mapping |
|--------------|----------------------|-------|----------|----------|------------------------|---------|-----------------|
| `qwen36 chat` | None observed | None observed | No positional args observed | None | Starts interactive session | None observed | Snap app exposes `CHAT=$SNAP/bin/chat.sh` for chat runtime integration |
| `qwen36 use-engine` | Optional engine name when not using `--auto` | `--auto`, `--assume-yes` | Requires either an engine selection mode or `--auto` | No documented default on the command line; install hook uses `--auto` | Positional engine values evidenced: `cpu`, `cuda`; `--auto` selects automatically | None observed | None observed |
| `qwen36 show-engine` | None observed | None observed | No | Returns current engine | YAML output with at least `name` and `components` | None observed | None observed |
| `qwen36 get` | `<key>` | None observed | Yes | None | Keys evidenced: `http.port`, `http.host`, `http.base-path`, `model-name`, `verbose`, `server`, `model`, `multimodel-projector`, `gpu-layers` | None observed | None observed |
| `qwen36 set` | `<key>=<value>` | `--package` | Yes | None | Same key space as `get`; values are strings from the shell's point of view | None observed | None observed |
| `qwen36 completion bash` | Subcommand target `bash` | None observed | Yes, `bash` is the only evidenced target | None | Emits whitespace-delimited completion candidates for bash | None observed | None observed |

## Per-Command Detail

### `qwen36 chat`

- **Observed syntax**: `qwen36 chat`
- **Positional arguments**: none observed
- **Flags**: none observed
- **Output mode**: interactive terminal UI through `go-chat-client`
- **Runtime dependencies**: local server expected on `http://localhost:<port>/<base-path>`

### `qwen36 use-engine`

- **Observed syntax**:
  - `qwen36 use-engine --auto`
  - `qwen36 use-engine --auto --assume-yes`
  - `qwen36 use-engine cpu`
  - `qwen36 use-engine cuda`
- **Positional argument**:
  - `engine`: selects a named engine
- **Flags**:
  - `--auto`: auto-detect best engine from available hardware
  - `--assume-yes`: skip confirmation prompts during engine switching
- **Accepted engine names evidenced in repo**: `cpu`, `cuda`
- **Mutability**: yes; this command changes selected engine and likely derived config keys

### `qwen36 show-engine`

- **Observed syntax**: `qwen36 show-engine`
- **Positional arguments**: none observed
- **Flags**: none observed
- **Output contract relied upon by scripts**:
  - `.name` must exist
  - `.components[]` must exist

### `qwen36 get`

- **Observed syntax**: `qwen36 get <key>`
- **Positional argument**:
  - `key`: one config key from the store
- **Flags**: none observed
- **Keys evidenced**:
  - `http.port`
  - `http.host`
  - `http.base-path`
  - `model-name`
  - `verbose`
  - `server`
  - `model`
  - `multimodel-projector`
  - `gpu-layers`

### `qwen36 set`

- **Observed syntax**:
  - `qwen36 set http.port=8326`
  - `qwen36 set --package http.port="8326"`
- **Positional argument**:
  - `key=value`: assignment expression
- **Flags**:
  - `--package`: write package-scoped defaults or package-owned config layer
- **Keys evidenced**: same as `get`
- **Value types as observed by consumers**:
  - strings: `http.host`, `http.base-path`, `server`, `model`, `multimodel-projector`, `model-name`
  - integer-like strings: `http.port`, `gpu-layers`
  - boolean-like strings: `verbose`

### `qwen36 completion bash`

- **Observed syntax**: `qwen36 completion bash`
- **Positional argument**:
  - `bash`: shell target
- **Flags**: none observed
- **Output**: consumed by `compgen -W`, so words must be whitespace-safe and shell completion oriented

## Special Arguments

This CLI has several structural exceptions and non-standard patterns:

1. **`set` uses a positional `key=value` assignment** instead of a flag-based assignment model. That is acceptable for compact config mutation, but it should be documented as the only accepted form.

2. **`--package` appears only on `set`** and introduces a hidden configuration layer. This is a meaningful semantic modifier but is undocumented in the public README.

3. **`completion bash` is a second-level command where the top-level token is noun-led (`completion`)** rather than verb-led. Per DE013 grammar, this is a standards mismatch unless treated as a documented exception.

4. **`use-engine` mixes positional selection (`cpu`, `cuda`) with a mode flag (`--auto`)**. Mutual exclusivity and precedence are not documented.

## Gaps

- No help output was available to confirm optional versus required parser rules.
- No aliases or short flags are evidenced.
- No environment-variable overrides are evidenced for config keys.
- No explicit value validation rules are documented for `set`.
- No documentation on what happens when mutually exclusive arguments are provided together.
