# Output Contracts

## Overview

The `qwen36` CLI produces two categories of output across its commands:

1. **Human-readable output**: Tables, YAML (default), and interactive prompts.
2. **Machine-readable output**: JSON (via `--format json`) and YAML (structured, parseable).

Commands that display data support `--format` with `yaml` (default for most) or `json`. The sole exception is `list-engines`, which defaults to `table` format and also supports `json`.

## Output by Command

### `status`

| Format | Output | Stability |
|--------|--------|-----------|
| `yaml` (default) | YAML document with keys: `engine` (string), `services` (map), `endpoints` (map), `model` (map) | **Unstable** — keys depend on engine component settings |
| `json` | JSON object with same structure | **Unstable** — same caveats |

The `endpoints` field is populated from active engine component YAML files. Its keys are server names defined within those components. The `model` field depends on `MODEL_NAME` env var in component settings. Both are inherently engine/component dependent.

### `show-machine`

| Format | Output | Stability |
|--------|--------|-----------|
| `yaml` (default) | YAML document describing CPU, memory, disk, PCI devices, GPU info | **Unstable** — structure reflects `types.HwInfo` struct which may evolve |
| `json` | JSON object with same structure | **Unstable** |

Warnings go to **stderr**. The main output goes to **stdout**, making it pipeable:
```bash
qwen36 show-machine --format=json | qwen36 debug select-engine --engines test_data/engines/
```

### `show-engine [<engine>]`

| Format | Output | Stability |
|--------|--------|-----------|
| `yaml` (default) | YAML document including engine manifest fields + `compatible` (bool) + `compatibility-issues` ([]string) | **Unstable** — wraps `engines.ScoredManifest` struct |
| `json` | JSON with same structure | **Unstable** |

### `list-engines`

| Format | Output | Stability |
|--------|--------|-----------|
| `table` (default) | Formatted table: columns `engine`, `vendor`, `description`, `compat` | **Human-only** — not designed for parsing |
| `json` | JSON object: `active-engine` (string), `engines` (array of `EngineDetails`) | **Unstable** — same struct as `show-engine` |

Table specifics:
- Active engine marked with `*` suffix on name.
- Compat column: `yes` (compatible + stable), `devel` (compatible + non-stable grade), `no` (incompatible).
- Max width: 80 chars. Description column truncated with `WrapTruncate`.
- Empty state: `"No engines found."` to stderr if only header row present.

### `version`

| Format | Output | Stability |
|--------|--------|-----------|
| `yaml` (default) | YAML: `snap:` and `cli:` keys | **Stable** |
| `json` | JSON: same keys | **Stable** |

`cli` version comes from Go build info (`debug.ReadBuildInfo().Main.Version`). Empty/unset versions display as `"unset"`.

### `get [<key>]`

| Scenario | Output |
|----------|--------|
| `get` (no args) | YAML document of all config key-value pairs |
| `get <key>` (single primitive) | Plain text value (no YAML wrapping) |
| `get <key>` (nested object) | YAML sub-document |

When a key is not found: error message with suggestion to run `get` to view available keys.

### `set <key=value>...`

| Scenario | Output |
|----------|--------|
| Success | Restart prompt: `"Restart <snap> to apply the changes? [Y/n]"` |
| `--no-restart` | No output on success |
| `--assume-yes` | Auto-confirms restart |
| Non-interactive (piped) | No prompt, no restart |

### `unset <key>`

Same output contract as `set` — prompts to restart if the value actually changed.

### `use-engine [<engine>]`

| Scenario | Output |
|----------|--------|
| `--auto` | Per-engine compatibility summary (✔/✘/•), selected engine name, component installation progress, restart prompt |
| `<engine>` | Component installation prompts, `"Engine changed to <name>."`, restart prompt |
| `--fix` | Component installation/repair progress |
| Cancelled | `"Cancelled. No changes applied."` |

### `prune-cache`

| Scenario | Output |
|----------|--------|
| Components to remove | List of components with sizes (`<name> (<size>)`), per-engine annotation in multi-engine mode |
| No components | `"No components to remove."` |
| Confirmation | `"Continue pruning [<engines>] engines?"` or `"Continue pruning <engine> engine?"`, default `N` |
| Cancelled | `"Cancelled. No changes applied."` |

### `run <command>` (hidden)

Subprocess stdout and stderr are passed through directly. No structured output from `run` itself beyond error messages.

### `serve-webui` (hidden)

| Output | Description |
|--------|-------------|
| `Serving "<dir>" on http://localhost:<port>` | Startup message |
| `/config.json` endpoint | JSON with `openaiBaseURL`, `capabilities`, `instanceName`, `engineName` |

### `webui`

| Output | Description |
|--------|-------------|
| `Web UI is available at <url>` | Info message |
| `Press [Enter] to open it in your browser, or [Ctrl+C] to abort.` | Interactive prompt |

### `chat`

No structured output. The `go-chat-client` handles all terminal I/O.

### `debug validate-engines`

| Output | Description |
|--------|-------------|
| `✅ <path>` | Valid manifest |
| `❌ <path>: <error>` | Invalid manifest |
| `not all manifests are valid` | Aggregate error message on stderr (exit code 1) |

### `debug select-engine`

| Stream | Output |
|--------|--------|
| **stderr** | Colorized per-engine summary: ✅/🟠/❌ with score/grade/incompatibility reasons |
| **stderr** | `Selected engine for your hardware configuration: <name>` (bold green) |
| **stdout** | JSON or YAML of `EngineSelection` struct (`engines` array + `top-engine` string) |

### `debug chat` / `debug serve-webui`

Same as their top-level counterparts but with different flag defaults.

## Parseability Guidance

### Do parse:
- `--format=json` output from `status`, `show-machine`, `show-engine`, `version`, `list-engines`, `debug select-engine`
- Exit codes: `0` = success, `1` = error

### Do NOT parse:
- Table output from `list-engines` (human-only, column widths vary)
- YAML output (default format) — structure may change across versions
- Interactive prompts and progress messages
- Spinner/progress bar output

### Stability expectations:
No formal output versioning or stability guarantees exist. The CLI is `grade: devel` and all output structures should be considered **unstable**. The `version` command is the most likely to remain stable since it has the simplest structure.

## Machine-Readable Format Conventions

- JSON output uses `json.MarshalIndent` with 2-space indentation.
- YAML output uses `yaml.Marshal` with default gopkg.in/yaml.v3 formatting.
- All machine-readable output goes to **stdout**. Errors, warnings, and human-facing messages go to **stderr**.
- `list-engines` empty state (`--format=json`): returns JSON with empty engines array and active engine (may be empty string).
