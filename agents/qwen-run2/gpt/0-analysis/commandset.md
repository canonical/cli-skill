# Command Set

## Command hierarchy

```text
qwen36
├── chat
├── use-engine [cpu|cuda|--auto] [--assume-yes]
├── show-engine
├── get <key>
├── set [--package] <key>=<value>
└── completion bash

qwen36.server
└── daemon service entrypoint
```

## Scope note

The private Go CLI submodule is not present in this checkout, so the command inventory is reconstructed from the snap definition, README examples, shell integration, and hook usage. The commands below are the ones directly evidenced in the repository and in the task brief.

## Full command list

| Command | Kind | Short description | How it works |
|---|---|---|---|
| `qwen36 chat` | user command | Starts an interactive chat session against the local inference server. | The snap exposes `CHAT=$SNAP/bin/chat.sh` in the `qwen36` app environment. `apps/chat.sh` resolves `http.port`, optional `model-name`, and `http.base-path` with `qwen36 get`, falls back to `v1` when `http.base-path` is empty, exports `OPENAI_BASE_URL` and `MODEL_NAME`, and execs `go-chat-client`. The chat path depends on the daemon being available; `apps/wait-for-server.sh` polls the selected engine health check for up to 60 seconds before allowing the chat chain to proceed. |
| `qwen36 use-engine` | user command | Selects the inference engine explicitly (`cpu`, `cuda`) or automatically (`--auto`). | README examples show `qwen36 use-engine cpu`, `qwen36 use-engine cuda`, and `qwen36 use-engine --auto`. The install hook uses `qwen36 use-engine --auto --assume-yes`. The selected engine later drives `qwen36 show-engine`, and its manifest-derived config controls which components the daemon waits for and which runtime options are read through `qwen36 get`. |
| `qwen36 show-engine` | user command | Prints the current engine description as YAML. | `apps/server.sh` parses `qwen36 show-engine | yq .components[]` and `qwen36 show-engine | yq .name`, so the command must emit YAML with at least `name` and `components[]`. In practice it represents the currently selected engine manifest plus resolved engine-specific configuration. |
| `qwen36 get <key>` | user command | Reads one effective snap configuration value. | Shell scripts rely on it for `http.port`, `http.host`, `http.base-path`, `verbose`, `server`, `model`, `multimodel-projector`, `gpu-layers`, and optional `model-name`. The command behaves as the read side of the snap config control plane and must output a bare scalar suitable for shell command substitution. |
| `qwen36 set <key>=<value>` | user command | Writes one persisted configuration value. | README documents `qwen36 set http.port=8326`. The install hook uses `qwen36 set --package` to seed defaults for `http.port`, `http.host`, and `verbose`. The command writes the configuration that later drives chat setup and engine launch scripts. |
| `qwen36 completion bash` | user command | Emits completion words for bash completion. | `apps/completion.bash` invokes `$SNAP/bin/qwen36 completion bash` and feeds the result to `compgen -W`. That means the command is a generator for completion candidates rather than a user-facing workflow command. Only the `bash` subcommand is evidenced in this repository. |
| `qwen36.server` | daemon service | Runs the long-lived inference service. | Declared as a `daemon: simple` snap app named `server`. It runs `apps/server.sh`, which waits for required components listed by `show-engine`, then execs `engines/<name>/server`. The engine launcher resolves component paths through `qwen36 get`, sources component init scripts, and finally starts `llama-server` with CPU- or CUDA-specific flags. |

## Command groups

### Conversation

- `qwen36 chat`

### Engine management

- `qwen36 use-engine`
- `qwen36 show-engine`
- `qwen36.server`

### Configuration management

- `qwen36 get`
- `qwen36 set`

### Shell integration

- `qwen36 completion bash`

## Command relationships

| Relationship | Commands | Notes |
|---|---|---|
| Mutate and inspect engine selection | `use-engine` ↔ `show-engine` | `use-engine` changes the selected engine; `show-engine` is the machine-readable inspection path used by scripts. |
| Mutate and inspect config | `set` ↔ `get` | `set` persists config; `get` reads effective values consumed by scripts. |
| User action and backing service | `chat` ↔ `qwen36.server` | `chat` is a client workflow; `qwen36.server` is the runtime that serves the local API. |
| Shell UX support | `completion bash` ↔ all `qwen36` commands | Completion generation depends on the Go CLI’s command inventory and syntax. |

## Observed design pattern

The command set is mostly a flat verb-first surface with one verb-noun compound (`use-engine`, `show-engine`) and one nested meta-command (`completion bash`). That keeps the surface small, but it mixes three shapes:

- bare verbs: `chat`, `get`, `set`
- verb-noun compounds: `use-engine`, `show-engine`
- meta-command + shell target: `completion bash`

The service entrypoint `qwen36.server` sits outside that grammar and behaves as a packaging/runtime interface rather than a normal user command.

## Gaps caused by the missing Go submodule

The repository does not expose the Go command registration or help text, so the following remain unverified from source:

- whether additional hidden commands or aliases exist
- exact help output and usage strings
- argument validation and error wording for invalid command lines
- whether `set` accepts multiple assignments in one invocation
- whether `completion` supports shells other than bash

Based on the checked-in snap surface, the seven commands above are the full observable command set.