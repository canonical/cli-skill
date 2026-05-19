# Command Set

## Top-Level Commands

The `qwen36` CLI exposes a **flat top-level command hierarchy** with one exception: the `debug` subcommand group. Commands are organized into groups for help display purposes only — the Cobra `Group` feature is used purely for cosmetic grouping in `--help` output, not for structural nesting.

### Command Groups (Help Display)

| Group ID | Title | Commands |
|----------|-------|----------|
| `basic` | Basic Commands | `status`, `chat` (conditional), `webui` (conditional) |
| `config` | Configuration Commands | `get`, `set`, `unset` |
| `engine` | Management Commands | `list-engines`, `show-engine`, `use-engine` |
| _(ungrouped)_ | Additional Commands | `show-machine`, `prune-cache`, `version` |
| _(hidden)_ | — | `run`, `serve-webui`, `debug` |

---

## Full Command Roster

### 1. `status`

- **Description**: Show the status of the inference snap, including active engine name, service statuses, server endpoints, and model information.
- **How it works**: `status.Run` retrieves the active engine from cache, queries snapd for service statuses (`snapctl.Services()`), resolves server endpoints from active engine component settings, and extracts model metadata from component environment variables. Output is rendered as YAML (default) or JSON via `--format`.
- **Key code**: `cmd/cli/commands/status.go`, `common/services.go`, `common/endpoints.go`, `common/model_status.go`

### 2. `chat`

- **Description**: Start an interactive text chat with the inference server via its OpenAI-compatible API.
- **How it works**: `chat.Run` resolves the OpenAI endpoint from the active engine, checks that the server snap service is active, then launches the `go-chat-client` with the resolved base URL. If the server is inactive, it prints a suggestion to start it. **Conditional**: only available when `ADDITIONAL_FEATURES` env var includes `"chat"`.
- **Key code**: `cmd/cli/commands/chat.go`, `common/endpoints.go`, `common/chat.go`

### 3. `webui`

- **Description**: Open the snap's built-in web UI in the default browser.
- **How it works**: Waits for engine components, resolves the web UI HTTP URL, checks that both `server` and `server-webui` services are active, waits for the OpenAI server to be ready (via chat client health check), then offers to open the URL in the default browser using `xdg-open`. **Conditional**: only available when `ADDITIONAL_FEATURES` env var includes `"webui"`.
- **Key code**: `cmd/cli/commands/webui.go`, `common/endpoints.go`

### 4. `get [<key>]`

- **Description**: Print one or all configuration values. With no argument, prints all configs as YAML. With a key, prints the value of that key.
- **How it works**: `get.Run` delegates to `Config.Get(key)` or `Config.GetAll()`, which reads from snapctl storage with precedence rules (UserConfig > EngineConfig > PackageConfig). Single values are printed as plain text; multiple values or full output as YAML.
- **Key code**: `cmd/cli/commands/get.go`, `pkg/storage/config.go`

### 5. `set <key=value>...`

- **Description**: Set one or more configuration values using `key=value` pairs. Supports `--package` (hidden) and `--engine` (hidden) flags to target different config layers.
- **How it works**: Requires root privileges. Parses key=value pairs from args, validates keys exist (or are `passthrough.*` keys), writes to `UserConfig` layer via `Config.Set()`. After changes, prompts to restart the snap (unless `--no-restart` or `--assume-yes`).
- **Key code**: `cmd/cli/commands/set.go`, `pkg/storage/config.go`

### 6. `unset <key>`

- **Description**: Unset a user configuration, reverting to the package or engine default. If no default exists, the key is removed entirely.
- **How it works**: Requires root privileges. Checks the key exists, calls `Config.Unset(key, UserConfig)`, compares old and new values. If the value changed, prompts to restart (unless `--no-restart` or `--assume-yes`).
- **Key code**: `cmd/cli/commands/unset.go`, `pkg/storage/config.go`

### 7. `list-engines`

- **Description**: List all available engines with compatibility status.
- **How it works**: `listEngines.Run` scores all engines against host hardware, retrieves the active engine from cache, then renders a table (default) or JSON (`--format json`). The table shows engine name, vendor, description, and compatibility (yes/devel/no). Active engine is marked with `*`. Engines sorted by descending score.
- **Key code**: `cmd/cli/commands/list-engines.go`, `common/engine.go`, `common/engine_details.go`

### 8. `show-engine [<engine>]`

- **Description**: Print detailed information about the active engine, or a specified engine.
- **How it works**: Without args, shows the active engine (from cache). With an engine name, shows that specific engine after scoring it against host hardware. Output is YAML (default) or JSON via `--format`. Supports shell completion for engine names.
- **Key code**: `cmd/cli/commands/show-engine.go`, `common/engine_details.go`

### 9. `use-engine [<engine>]`

- **Description**: Select an engine. Supports `--auto` for automatic selection, `--fix` to repair the active engine, or specifying an engine name directly.
- **How it works**: Requires root privileges. For manual selection: loads the named engine manifest, installs missing components (with confirmation prompt), unsets previous engine config, sets the new engine as active in cache, applies engine configs, prompts to restart. For `--auto`: scores all engines, displays compatibility summary, selects the top-scoring engine. For `--fix`: re-installs missing components and re-applies configs for the already-active engine.
- **Key code**: `cmd/cli/commands/use-engine.go`, `common/engine.go`, `pkg/selector/`

### 10. `show-machine`

- **Description**: Print detailed information about the host machine (CPU, memory, disk, PCI devices, GPU).
- **How it works**: Probes hardware via `/proc/cpuinfo`, `/proc/meminfo`, `lspci`, `clinfo`, `statfs`, and kernel sysfs. Displays a spinner during probing. Output is YAML (default) or JSON via `--format`. Warnings (e.g., missing `clinfo`) go to stderr when `--verbose`.
- **Key code**: `cmd/cli/commands/show-machine.go`, `pkg/hardware_info/`

### 11. `prune-cache`

- **Description**: Remove cached data (engine components that are not needed by the active engine).
- **How it works**: Requires root privileges. With no `--engine` flag: finds all components from inactive engines not shared with the active engine, lists them with sizes, prompts for confirmation, then removes them via `snapctl.RemoveComponents()`. With `--engine <name>`: removes components specific to that named engine only. Refuses to prune the active engine.
- **Key code**: `cmd/cli/commands/prune-cache.go`, `pkg/snap_store/`

### 12. `version`

- **Description**: Show snap and CLI version information.
- **How it works**: Reads snap version from `env.SnapVersion()` and CLI version from Go build info (`debug.ReadBuildInfo()`). Output is YAML (default) or JSON via `--format`.
- **Key code**: `cmd/cli/commands/version.go`

### 13. `run <command>` (hidden)

- **Description**: Run a subprocess in the active engine's environment.
- **How it works**: Waits for engine components, loads the engine environment (env vars + symlinks from component settings), processes `passthrough.environment.*` config keys as environment variables, then executes the specified command via `exec.Command`. The `--wait-for-components` flag is deprecated (always waits now). Supports `--` separator for passing arguments to the subprocess.
- **Key code**: `cmd/cli/commands/run.go`, `common/engine.go`

### 14. `serve-webui <static-files-dir>` (hidden)

- **Description**: Serve static files and configurations for the web UI on a local HTTP server.
- **How it works**: Waits for engine components, resolves the OpenAI endpoint, reads capabilities from the `--capabilities` flag (default: "text"), and starts an HTTP server serving the static directory at the configured port (default 8081) and host (default localhost). Provides a `/config.json` endpoint with OpenAI base URL and capabilities.
- **Key code**: `cmd/cli/commands/serve-webui.go`, `pkg/webui/`

### 15. `debug` (hidden command group)

Developer/debugging command group containing four subcommands:

#### 15a. `debug validate-engines <manifest>...`

- **Description**: Validate one or more engine manifest YAML files.
- **How it works**: Loads and validates each manifest file path using `engines.Validate()`. Prints ✅ for valid, ❌ for invalid with error details. Returns non-zero exit if any manifest is invalid.
- **Key code**: `cmd/cli/commands/debug/validate.go`, `pkg/engines/validate.go`

#### 15b. `debug select-engine`

- **Description**: Test which engine would be chosen for a given machine, using hardware info piped via stdin.
- **How it works**: Reads YAML hardware info from stdin (produced by `show-machine --format=yaml`), loads engine manifests from `--engines` directory, scores them, prints compatibility summary to stderr (with color: ✅/🟠/❌), and outputs the full selection result (JSON or YAML) to stdout.
- **Key code**: `cmd/cli/commands/debug/select.go`, `pkg/selector/`

#### 15c. `debug chat`

- **Description**: Start a chat session with explicit connection parameters (`--base-url` and `--model`) rather than auto-detecting from the active engine.
- **How it works**: Requires `--base-url`. Starts the `go-chat-client` with the provided parameters. Useful for testing against arbitrary OpenAI-compatible servers.
- **Key code**: `cmd/cli/commands/debug/chat.go`, `common/chat.go`

#### 15d. `debug serve-webui <static-files-dir>`

- **Description**: Serve web UI static files for debugging with explicit `--base-url` instead of auto-detection.
- **How it works**: Similar to the top-level `serve-webui` but accepts `--base-url` explicitly (default `http://localhost:8080/v1`) rather than resolving from the active engine. Enables all capabilities. Fixed to `localhost` only.
- **Key code**: `cmd/cli/commands/debug/serve-webui.go`, `pkg/webui/`

---

## Command Count Summary

| Category | Count |
|----------|-------|
| Top-level visible commands | 12 (`status`, `chat`, `webui`, `get`, `set`, `unset`, `list-engines`, `show-engine`, `use-engine`, `show-machine`, `prune-cache`, `version`) |
| Top-level hidden commands | 2 (`run`, `serve-webui`) |
| Debug subcommands | 4 (`validate-engines`, `select-engine`, `chat`, `serve-webui`) |
| **Total** | **18** |
