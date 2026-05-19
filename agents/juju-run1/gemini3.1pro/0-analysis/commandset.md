# Command Set

## CLI Binary: `qwen36`

| Command | Description | How It Works |
|---------|-------------|--------------|
| `chat` | Start an interactive chat session with the model | Invokes `$SNAP/bin/chat.sh` which reads config (`http.port`, `model-name`, `http.base-path`), sets `OPENAI_BASE_URL` and `MODEL_NAME` env vars, then launches `go-chat-client` which connects to the running llama-server's OpenAI-compatible API |
| `use-engine` | Select the inference engine for the snap | Accepts a positional argument (`cpu`, `cuda`) or the `--auto` flag for automatic hardware detection. Reads engine YAML files from `$SNAP/engines/*/engine.yaml`, evaluates hardware requirements (CPU flags, GPU vendor/VRAM, RAM, disk), and writes selected engine configuration to snap options. Flags: `--auto` (detect best engine), `--assume-yes` (skip confirmation) |
| `show-engine` | Display the current engine configuration | Outputs the engine YAML content for the currently selected engine (structured YAML with name, description, vendor, grade, devices, memory, disk-space, components, configurations) |
| `get <key>` | Retrieve a configuration value | Reads a single snap option by key. Known keys: `http.port`, `http.host`, `http.base-path`, `model-name`, `verbose`, `server`, `model`, `multimodel-projector`, `gpu-layers`. Returns the raw value to stdout |
| `set <key>=<value>` | Set a configuration value | Writes a snap option. Accepts `key=value` format as positional argument. Flag: `--package` (sets a package-level default rather than user-level override). Used in install hook to set defaults |
| `completion bash` | Generate bash shell completions | Outputs a list of completion words for the current input context. Used by the bash completion script (`completion.bash`) which registers via `complete -F` |

## Daemon App: `server`

The snap also exposes a `server` daemon app (defined in `snapcraft.yaml`), which is not a CLI command but a systemd service:

| App | Description | How It Works |
|-----|-------------|--------------|
| `qwen36.server` | Long-running inference server daemon | Runs `bin/server.sh` which: (1) waits for required snap components to be installed, (2) reads current engine from `qwen36 show-engine`, (3) execs the engine-specific server script (`engines/<engine>/server`) which launches `llama-server` with appropriate flags |

## Shell Scripts (Supporting)

These are not user-facing commands but support the daemon:

| Script | Purpose |
|--------|---------|
| `server.sh` | Entry point for the server daemon; waits for components, then delegates to engine-specific server |
| `chat.sh` | Entry point for the chat command; configures OpenAI API URL and launches go-chat-client |
| `wait-for-server.sh` | Polls the engine's check-server script until the server is ready (60s timeout) |
| `check-server-llamacpp.sh` | Health check for llama-server: verifies process running, port open, completions API responding |
| `completion.bash` | Bash completion registration script |
| `engines/cpu/server` | Launches llama-server with CPU-specific flags |
| `engines/cuda/server` | Launches llama-server with CUDA-specific flags (--n-gpu-layers, --no-mmap, --fit) |
| `engines/*/common.sh` | Shared setup: resolves component paths, validates components exist, sets LD_LIBRARY_PATH |

## Snap Hooks

| Hook | Purpose |
|------|---------|
| `install` | Sets default config (`http.port=8326`, `http.host=127.0.0.1`, `verbose=false`), runs `use-engine --auto --assume-yes` if hardware-observe is connected |
| `post-refresh` | Minimal: only sets up syslog redirection (placeholder for future logic) |
