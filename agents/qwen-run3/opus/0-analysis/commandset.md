# Command Set

## Summary

The `qwen36` CLI exposes 12 visible commands (plus 3 hidden) organized into three groups: Basic, Configuration, and Management. The command set is flat (no subcommands beyond the hidden `debug` group).

## Command List

### Basic Commands

| Command | Description (docs/help) | How it works (code path) |
|---------|------------------------|--------------------------|
| `status` | Show the status of the inference snap | Queries active engine, service state, component installation status. Outputs in YAML or JSON format via `--format`. Optionally waits for components with `--wait-for-components`. |
| `chat` | Start the chat CLI | Reads `http.port` and `http.base-path` from config, checks server service is active, then launches the external `go-chat-client` binary with `OPENAI_BASE_URL` set. |
| `webui` | Launch web UI | Waits for components, resolves the web UI HTTP URL, checks both `server` and `server-webui` services are active, then opens the URL in the default browser via `xdg-open`. |

### Configuration Commands

| Command | Description (docs/help) | How it works (code path) |
|---------|------------------------|--------------------------|
| `get [<key>]` | Print configurations | Reads from the snap config store. Without args, dumps all config as YAML. With a key, prints that value (single value plain, multiple as YAML). |
| `set <key=value>...` | Set configurations | Requires root. Parses key=value pairs, validates against known keys, prompts for confirmation if value differs from current, applies via `snapctl set`. Supports hidden `--package` and `--engine` layers. Triggers service restart unless `--no-restart`. |
| `unset <key>` | Unset configurations | Requires root. Removes the user-layer value for a key, reverting to package/engine default. Triggers service restart unless `--no-restart`. |

### Management Commands

| Command | Description (docs/help) | How it works (code path) |
|---------|------------------------|--------------------------|
| `list-engines` | List available engines | Loads engine manifests from `$SNAP/engines/`, scores each against detected hardware, outputs as a table (with color-coded compatibility) or JSON via `--format`. |
| `show-engine [<engine>]` | Print information about an engine | Shows the active engine (or named engine) manifest with hardware requirements and compatibility score. Outputs in YAML or JSON via `--format`. |
| `use-engine [<engine>]` | Select an engine | Requires root. Three modes: (1) named engine switch, (2) `--auto` for hardware-based auto-selection, (3) `--fix` to repair active engine's components. Installs required snap components from the store, applies engine config, restarts service. |

### Ungrouped Commands

| Command | Description (docs/help) | How it works (code path) |
|---------|------------------------|--------------------------|
| `show-machine` | Print information about the host machine | Detects CPU (flags, architecture), GPUs (vendor, VRAM via PCI/clinfo/nvidia-smi), memory, disk. Outputs in YAML or JSON. |
| `prune-cache` | Remove cached data | Requires root. Identifies snap components from inactive engines, prompts for confirmation, then removes them via `snap remove-component`. Optionally targets a specific engine with `--engine`. |
| `version` | Show version information | Reads snap version from `$SNAP_VERSION` env and CLI version from Go build info. Outputs in YAML or JSON. |

### Hidden Commands

| Command | Description (docs/help) | How it works (code path) |
|---------|------------------------|--------------------------|
| `run <command>` | Run a subprocess in the engine's environment | Waits for components, loads engine environment (LD_LIBRARY_PATH, model paths), then exec's the given command. Uses `--` separator for subprocess args. |
| `serve-webui <static-files-dir>` | Serve static files and configurations for the web UI | Internal daemon command. Starts an HTTP server serving static files with injected config (OpenAI endpoint, model name, capabilities). |
| `debug` | Developer/debugging commands | Hidden subcommand group containing: `validate`, `select`, `chat`, `serve-webui`. |

### Snap Hooks (non-CLI, lifecycle)

| Hook | Purpose |
|------|---------|
| `install` | Sets default config (port 8326, host 127.0.0.1, verbose false). Runs `use-engine --auto --assume-yes` if hardware-observe is connected. |
| `post-refresh` | Redirects output to syslog. (Currently minimal — placeholder for future refresh logic.) |

## Total Command Count

- **Visible**: 12
- **Hidden**: 3 (run, serve-webui, debug)
- **Total**: 15 (including debug subgroup as one)

**Scale classification**: < 15 visible commands → **Compact mode**
