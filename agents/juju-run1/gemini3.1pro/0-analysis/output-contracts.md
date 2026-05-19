# Output Contracts

## Command Output Formats

### `qwen36 show-engine`

**Format**: YAML (machine-readable)

**Example output** (CPU engine):
```yaml
name: cpu
description: Use CPU with a good balance between speed and accuracy (35B total / 3B activated)
vendor: Canonical Ltd
grade: stable
devices:
  anyof:
    - type: cpu
      architecture: amd64
      flags: [sse4_2, f16c, fma, avx, avx2]
    - type: cpu
      architecture: arm64
      features: [asimd]
memory: 32G
disk-space: 25G
components:
  - llamacpp
  - model-qwen36-35b-a3b-ud-q4-k-xl
  - mmproj-qwen36-35b-a3b-f16
configurations:
  server: llamacpp
  model: model-qwen36-35b-a3b-ud-q4-k-xl
  multimodel-projector: mmproj-qwen36-35b-a3b-f16
  http.base-path: v1
```

**Stability**: The YAML structure is consumed by `yq` in shell scripts (e.g., `qwen36 show-engine | yq .name`, `qwen36 show-engine | yq .components[]`). Fields `name` and `components[]` are load-bearing for the server daemon startup.

**Parseability**: Pipe through `yq` for field extraction. No `--format` flag observed.

### `qwen36 get <key>`

**Format**: Plain text, single value on stdout (no trailing decoration observed)

**Example outputs**:
```
8326
127.0.0.1
v1
false
llamacpp
```

**Stability**: Consumed directly by shell scripts via command substitution (`port="$(qwen36 get http.port)"`). Output must be a bare value with no framing text.

**Edge cases**:
- `model-name` may produce empty output or write to stderr when unset (scripts use `2>/dev/null || true`)
- `gpu-layers` may be empty for CPU engines

### `qwen36 set <key>=<value>`

**Format**: No output on success (silent operation, exit code 0)

**Stability**: Write-only command. Success is indicated by exit code.

### `qwen36 use-engine`

**Format**: Human-readable text (interactive)

**Expected behavior**:
- With `--auto`: prints detected hardware capabilities and selected engine
- Without `--auto`: selects the named engine
- May prompt for confirmation unless `--assume-yes` is given

**Stability**: Output is informational; scripts rely only on the exit code and side effects (config changes).

### `qwen36 chat`

**Format**: Interactive terminal session

**Behavior**: Launches `go-chat-client` which provides an interactive chat interface. Not designed for scripting.

### `qwen36 completion bash`

**Format**: Whitespace-separated word list on stdout

**Example output**: A list of valid completion words for the current context.

**Stability**: Consumed by bash's `compgen -W` in the completion script. Must be a flat word list.

## Machine-Readable vs Human-Readable Summary

| Command | Format | Machine-Readable | Scriptable |
|---------|--------|-----------------|------------|
| `show-engine` | YAML | Yes | Yes (via yq) |
| `get` | Plain text value | Yes | Yes |
| `set` | Silent | N/A | Yes (exit code) |
| `use-engine` | Human text | No | Partially (exit code only) |
| `chat` | Interactive | No | No |
| `completion bash` | Word list | Yes | Yes |

## No `--format` or `--output` Flag

No commands offer alternative output formats (e.g., `--format json`). The output format is fixed per command.
