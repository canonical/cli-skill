# qwen36 Command Set

## Scope

This analysis is limited to commands that are directly evidenced by the public repository. The leaf command inventory is:

1. `qwen36 chat`
2. `qwen36 use-engine`
3. `qwen36 show-engine`
4. `qwen36 get`
5. `qwen36 set`
6. `qwen36 completion bash`

There is also an intermediate grouping token, `qwen36 completion`, but the only evidenced leaf target is `bash`.

## Hierarchy

```text
qwen36
├── chat
├── use-engine
├── show-engine
├── get
├── set
└── completion
    └── bash
```

## Command Inventory

| Command | What It Does | How It Works |
|---------|--------------|--------------|
| `qwen36 chat` | Starts an interactive chat session against the local inference server. | The snap app exports `CHAT=$SNAP/bin/chat.sh`. That wrapper resolves `http.port`, `http.base-path`, and optional `model-name` using `qwen36 get`, exports `OPENAI_BASE_URL` and `MODEL_NAME`, and execs `go-chat-client`. The CLI likely dispatches to that wrapper or mirrors its behavior. |
| `qwen36 use-engine` | Selects the inference engine, either explicitly (`cpu`, `cuda`) or automatically (`--auto`). | The install hook invokes `qwen36 use-engine --auto --assume-yes`. The selected engine later drives `qwen36 show-engine`, which `apps/server.sh` parses to locate the engine name and required components. The command persists engine choice and engine-derived config into the CLI config store. |
| `qwen36 show-engine` | Prints the currently selected engine as YAML. | `apps/server.sh` parses `qwen36 show-engine | yq .components[]` and `qwen36 show-engine | yq .name`, so the command emits a YAML document containing at least `name` and `components`. The returned data corresponds to the selected `engines/*/engine.yaml` manifest plus resolved defaults. |
| `qwen36 get <key>` | Reads one config value from the CLI-managed configuration store. | All runtime wrappers depend on it. `apps/chat.sh` reads `http.port`, `http.base-path`, and `model-name`; engine launch scripts read `server`, `model`, `multimodel-projector`, `http.host`, `http.port`, `verbose`, and `gpu-layers`. The command outputs a scalar string suitable for shell substitution. |
| `qwen36 set <key>=<value>` | Writes one config value. | The install hook uses `qwen36 set --package` to seed package defaults. User-facing examples use plain `qwen36 set http.port=8326`. The command updates the same config store that `get` resolves, with multiple precedence layers. |
| `qwen36 completion bash` | Generates bash completion candidates for the CLI. | `apps/completion.bash` runs `$SNAP/bin/qwen36 completion bash` and feeds the resulting whitespace-delimited tokens into `compgen -W`, so the command behaves as a completion word generator rather than a static script printer. |

## Command Shape Assessment

- `chat`, `get`, and `set` fit common CLI verbs reasonably well.
- `use-engine` and `show-engine` are verb-noun forms, which aligns with DE013 when the noun must be explicit.
- `completion bash` is the least standard surface. `completion` is a noun-led grouping token rather than a verb, and the command's behavior is underdocumented.

## Missing Observable Commands

No public evidence was found for these common affordances:

- `qwen36 --help` (not documented in README)
- `qwen36 version` (not documented in README)
- `qwen36 status` (not documented)
- `qwen36 unset` (logical complement to `set`)
- `qwen36 list-engines` or equivalent discovery command (users cannot discover supported engines from the CLI)

That does not prove they do not exist in the private binary, only that they are not evidenced in the public packaging repository.

## Engine Support

Two engines are evidenced in the repository:

| Engine | Description | Requirements |
|--------|-------------|--------------|
| `cpu` | Uses CPU with llama.cpp (good balance of speed and accuracy) | 32GB RAM, 25GB disk, SSE4_2/F16C/FMA/AVX/AVX2 on amd64 or ASIMD on arm64 |
| `cuda` | Uses NVIDIA GPU via CUDA | 48GB RAM, 25GB disk, NVIDIA GPU with 24GB+ VRAM |

Both engines use the same model (`model-qwen36-35b-a3b-ud-q4-k-xl`) and multimodal projector (`mmproj-qwen36-35b-a3b-f16`).
