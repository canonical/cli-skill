# qwen36 Command Set

## Scope

This file documents the full command hierarchy present in the checked-out CLI source and distinguishes between:

- public user-facing commands
- conditional public commands gated by `ADDITIONAL_FEATURES`
- hidden internal and debug commands

For the command-shape workflow, the relevant public command list is 12 commands:

1. `status`
2. `chat` (conditional)
3. `webui` (conditional)
4. `get`
5. `set`
6. `unset`
7. `list-engines`
8. `show-engine`
9. `use-engine`
10. `show-machine`
11. `prune-cache`
12. `version`

The shipped qwen36 snap manifest does not currently set `ADDITIONAL_FEATURES`, so the actually exposed help surface defaults to 10 commands, not 12. That mismatch is a documentation and packaging issue, not a parser issue.

## Hierarchy

```text
qwen36
├── status
├── chat                    (conditional public)
├── webui                   (conditional public)
├── get
├── set
├── unset
├── list-engines
├── show-engine
├── use-engine
├── show-machine
├── prune-cache
├── version
├── run                     (hidden)
├── serve-webui             (hidden)
└── debug                   (hidden namespace)
    ├── validate-engines
    ├── select-engine
    ├── chat
    └── serve-webui
```

## Public And Conditional Commands

| Command | Description | How it works |
|---|---|---|
| `status` | Show overall snap status. | Reads the active engine from cache, snap service states from `snapctl`, server endpoints from component metadata, and model information from engine environment settings; prints YAML or JSON. |
| `chat` | Start the chat CLI. | Resolves the OpenAI endpoint, checks that the `server` service is active, waits for the model to become ready, then starts an interactive readline chat session. Present only when `ADDITIONAL_FEATURES` includes `chat`. |
| `webui` | Launch the built-in web UI. | Waits for components, checks both `server` and `server-webui` services, probes server readiness, prints the URL, and optionally opens it via `xdg-open`. Present only when `ADDITIONAL_FEATURES` includes `webui`. |
| `get [<key>]` | Print one or more resolved configuration values. | With no key, flattens and prints the merged config map as YAML. With a key, returns either a scalar or a YAML object for that subtree. |
| `set <key=value>...` | Set one or more configuration values. | Parses key/value pairs, validates keys, writes to the user config layer by default, and optionally restarts the snap unless `--no-restart` is used. |
| `unset <key>` | Remove a user configuration override. | Deletes the key from the user layer only, allowing engine or package defaults to reappear, then optionally restarts the snap if the effective value changed. |
| `list-engines` | List available engines. | Loads all manifests, scores them against the host hardware, marks the active engine, and prints either a table or JSON. |
| `show-engine [<engine>]` | Print details for the active or named engine. | Loads and scores manifests, resolves the target engine, and prints the manifest-derived detail record as YAML or JSON. |
| `use-engine [<engine>]` | Select or repair an engine. | Supports explicit selection, `--auto`, and `--fix`; installs missing components, rewrites engine config, updates the active-engine cache entry, and may restart the daemon. |
| `show-machine` | Print host hardware information. | Collects CPU, memory, disk, and PCI/GPU data through the hardware info packages and prints YAML or JSON. |
| `prune-cache` | Remove components used only by inactive engines. | Computes removable components from manifest overlap, prompts for confirmation, and removes components through `snapctl remove-components`. |
| `version` | Show snap and CLI version information. | Combines `SNAP_VERSION` with Go build info and prints YAML or JSON. |

## Hidden Commands

| Command | Description | How it works |
|---|---|---|
| `run <command>` | Run a subprocess in the active engine environment. | Waits for components, loads env vars and temporary symlink layout from component manifests, applies passthrough environment keys, and executes the child process. |
| `serve-webui <static-files-dir>` | Serve web UI files for the active snap configuration. | Resolves the active engine and OpenAI endpoint, builds a web UI config object, and starts an HTTP server for static assets. |
| `debug validate-engines <manifest>...` | Validate one or more engine manifest files. | Runs manifest schema validation and prints pass/fail per file. |
| `debug select-engine` | Preview engine selection for supplied hardware data. | Reads YAML hardware info from stdin, scores engines from a directory, and prints the selected engine plus compatibility detail. |
| `debug chat` | Start the chat client against an arbitrary base URL. | Uses the shared chat client without snap service checks; requires `--base-url`. |
| `debug serve-webui <static-files-dir>` | Serve debug web UI content against an arbitrary endpoint. | Starts the same file server as the hidden top-level command but takes explicit endpoint input. |

## Command Shape Observations

- The public surface is compact and flat.
- `get`/`set`/`unset` are the strongest grammar cluster.
- Engine-related commands are split across `list`, `show`, and `use`, plus hidden `select` in debug.
- `webui` is the only clearly noun-led public command, which conflicts with DE013 grammar guidance.
