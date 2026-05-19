# Command Set: qwen36 / inference-snaps-cli

## Hierarchy

The CLI is a Cobra root command with **flat grouping** (no deep nesting beyond one hidden debug namespace). Commands are organized into three visual groups plus ungrouped and hidden commands.

```
<snap-name>                             (root)
├── status                              [basic]
├── chat                                [basic]*
├── webui                               [basic]*
├── get                                 [config]
├── set                                 [config]
├── unset                               [config]
├── list-engines                        [engine]
├── show-engine                         [engine]
├── use-engine                          [engine]
├── show-machine                        (ungrouped)
├── prune-cache                         (ungrouped)
├── version                             (ungrouped)
├── run                                 (hidden)
├── serve-webui                         (hidden)
└── debug                               (hidden)
    ├── validate-engines
    ├── select-engine
    ├── chat
    └── serve-webui
```

Commands marked `*` are conditionally registered based on the `ADDITIONAL_FEATURES` environment variable.

## Full Command Inventory

### Basic Commands (Group: `basic`)

#### `status`
- **Short**: Show the status
- **Long**: Show the status of the inference snap
- **How it works**: Reads `Cache.GetActiveEngine()`, then queries snapd service statuses via `snapctl.Services()`, OpenAI server endpoints via component manifests, and model status via HTTP probe. Outputs a struct with Engine, Services, Endpoints, and Model fields in yaml or json.

#### `chat`
- **Short**: Start the chat CLI
- **Long**: Chat with the server via its OpenAI API. Supports text-based prompting only.
- **How it works**: Resolves the OpenAI endpoint URL from config + component manifests. Checks that the snap daemon service `server` is active via `snapctl.Services()`. Then launches an interactive readline-based REPL that streams chat completions via the `openai-go` SDK.

#### `webui`
- **Short**: Launch web UI
- **Long**: Open the snap's builtin web user interface in the default browser
- **How it works**: Waits for components, resolves the Web UI HTTP URL, checks service statuses for both `server` and `server-webui`, probes the OpenAI endpoint readiness, then prompts to open the URL with `xdg-open`.

### Configuration Commands (Group: `config`)

#### `get [<key>]`
- **Short**: Print configurations
- **Long**: Print one or more configurations
- **How it works**: If no key is given, calls `Config.GetAll()` and prints all resolved configs as YAML. If a key is given, calls `Config.Get(key)` and prints the value (single scalar on stdout, or YAML for multi-field objects).

#### `set <key=value>...`
- **Short**: Set configurations
- **Long**: Set a configuration
- **How it works**: Parses `key=value` pairs, validates keys exist (unless prefixed with `passthrough.`), writes to `UserConfig` layer, then conditionally prompts to restart the snap daemon via `PromptRestartToApplyChanges`. Hidden flags `--package` and `--engine` allow writing to lower precedence layers.

#### `unset <key>`
- **Short**: Unset configurations
- **Long**: Unset a user configuration, reverting to package or engine default. If no default exists, the key is removed entirely.
- **How it works**: Looks up current value via `Config.Get()`, calls `Config.Unset(key, UserConfig)`, then conditionally prompts for restart if the effective value changes.

### Management Commands (Group: `engine`)

#### `list-engines`
- **Short**: List available engines
- **Long**: (default help text)
- **How it works**: Loads all engine manifests, scores them against the current host machine capabilities (CPU, GPU, memory, disk), reads the active engine from cache, and prints a table or JSON output showing each engine's name, vendor, description, compatibility, and active marker.

#### `show-engine [<engine>]`
- **Short**: Print information about an engine
- **Long**: Print information about the active engine, or the specified engine
- **How it works**: If no argument is provided, shows the currently active engine via `Cache.GetActiveEngine()`. Otherwise loads the named manifest and prints its scored manifest + compatibility details as yaml or json.

#### `use-engine [<engine>]`
- **Short**: Select an engine
- **Long**: (default help text)
- **How it works**: Three mutually exclusive modes:
  1. **Explicit**: `use-engine <name>` — validates the manifest, installs missing snap components (with confirmation prompt), unsets previous engine configs, sets active engine cache, applies new engine configs, and optionally restarts the snap.
  2. **Auto**: `use-engine --auto` — scores all engines, picks the highest-score stable engine, then runs explicit mode.
  3. **Fix**: `use-engine --fix` — reinstalls missing components and re-applies config for the currently active engine; falls back to auto-selection if the active engine manifest is missing.

### Ungrouped Commands

#### `show-machine`
- **Short**: Print information about the host machine
- **Long**: Print information about the host machine, including hardware and compute resources
- **How it works**: Gathers hardware information (CPU architecture, instruction sets, PCI devices, memory, disk) via `/proc`, `lspci`, `df`, `clinfo`, and `nvidia-smi`. Prints as JSON or YAML. Warnings printed to stderr when verbose.

#### `prune-cache`
- **Short**: Remove cached data
- **Long**: (default help text)
- **How it works**: Computes which snap components are used only by inactive engines, optionally scoped to a single engine via `--engine`. Prompts for confirmation, then calls `snapctl.RemoveComponents()` and unsets engine configs. Cannot prune the active engine.

#### `version`
- **Short**: Show version information
- **Long**: (default help text)
- **How it works**: Reads `SNAP_VERSION` from the snap environment and the Go build info for the CLI binary. Prints both as yaml or json. Replaces empty strings with `"unset"`.

### Hidden Commands

#### `run <command>`
- **Short**: Run a subprocess
- **Long**: Run a command in the engine's environment. Supports `--` separator to pass flags to the subprocess.
- **How it works**: Waits for required components, loads the active engine's environment variables and temporary symlinks from component manifests, processes passthrough config env vars, then `exec`s the child process with connected stdout/stderr. Cleans up symlinks on normal exit (not on SIGKILL).

#### `serve-webui <static-files-dir>`
- **Short**: Serve static files and configurations for the web UI
- **Long**: (default help text)
- **How it works**: Waits for components, resolves the OpenAI endpoint, reads the active engine name, builds a `webui.Config`, validates it, then starts an HTTP file server on the configured host/port serving the static directory and a generated config JSON endpoint.

### Debug Commands (Hidden group: `debug`)

#### `debug validate-engines <manifest>...`
- **Short**: Validate engine manifest files
- **Long**: (default help text)
- **How it works**: Calls `engines.Validate()` on each provided manifest path, printing pass/fail per file. Exits with error if any manifest is invalid.

#### `debug select-engine`
- **Short**: Test which engine will be chosen
- **Long**: Test which engine will be chosen from a directory of engines, given the machine information piped in via stdin
- **How it works**: Reads `types.HwInfo` from stdin YAML, loads engine manifests from the `--engines` dir, scores them, and prints compatibility summary to stderr and full selection results to stdout as JSON or YAML.

#### `debug chat`
- **Short**: Start the chat CLI providing connection parameters
- **Long**: Open a text-only chat session to the OpenAI-compatible server at the provided URL
- **How it works**: Unlike top-level `chat`, this does not depend on snapd services. Requires `--base-url` and optionally `--model`, then starts the same readline REPL client against the arbitrary endpoint.

#### `debug serve-webui <static-files-dir>`
- **Short**: Serve web UI static files and configurations for debugging
- **Long**: (default help text)
- **How it works**: Same as `serve-webui` but accepts `--base-url` and `--port` explicitly, sets all capabilities, and hardcodes debug instance/engine names. Prints the generated config and then starts the HTTP server.
