# Output Contracts

## Format Flag Convention

Most inspection commands support a `--format` flag with values `json` or `yaml`. One command (`list-engines`) also supports `table`. The default format varies:

| Command | Default | Supported |
|---|---|---|
| `status` | `yaml` | `json`, `yaml` |
| `show-engine` | `yaml` | `json`, `yaml` |
| `show-machine` | `yaml` | `json`, `yaml` |
| `version` | `yaml` | `json`, `yaml` |
| `list-engines` | `table` | `table`, `json` |
| `debug select-engine` | `yaml` | `json`, `yaml` |

Human-readable defaults (`yaml`, `table`) are used for direct terminal inspection; json is provided for scripting.

## Machine-Readable Output Stability

### JSON/YAML Schemas

The CLI does **not** publish formal schemas. However, struct tags (`json`, `yaml`) in the source code define field names. Stability expectations:

- **Stable**: top-level keys like `engine`, `services`, `endpoints`, `model` (in `status`), `snap`, `cli` (in `version`), `active-engine`, `engines` (in `list-engines`).
- **Semi-stable**: nested engine manifest fields (e.g., `compatibility-report`, `score`, `grade`, `components`) are bound to the manifest schema, which may evolve.
- **Unstable**: `compatibility-issues` list contents (subject to wording changes).

### Table Output (`list-engines`)

- Not parseable by standard tools.
- Column widths are dynamically computed based on terminal width (max 80 chars).
- Active engine is denoted by a trailing `*` on the engine name.
- Compatible status renders as `yes`, `devel`, or `no`.
- **This format is explicitly not suitable for scripting.**

## Standard Error vs Standard Output

- **stdout**: primary command output (configs, status, engine lists, version data, chat responses).
- **stderr**: warnings, progress spinners, verbose logs, confirmation prompts, completion error messages, `list-engines` empty-state notice.
- `show-machine` and `debug select-engine` print warnings to stderr when `--verbose` is set.
- Error messages from command handlers are returned as `error` from `RunE`; Cobra prints them to stderr and exits non-zero.

## Streaming / Interactive Output

- `chat` streams tokens directly to stdout inline (no buffering per line). Reasoning content is colored blue. Main content is uncolored.
- `webui` prints a static URL and a prompt line to stdout; `xdg-open` is forked silently.
- `use-engine --auto` prints evaluation lines to stdout (✔ / • / ✘ markers) before restart prompts.

## Parseability Guidance for Scripts

Recommended patterns:
- Query config: `qwen36 get http.port --format=json | jq -r '.["http.port"]'`
- Query active engine: `qwen36 status --format=json | jq -r '.engine'`
- List engines: `qwen36 list-engines --format=json | jq '.engines[].name'`
- Check version: `qwen36 version --format=json | jq -r '.cli'`

Avoid parsing table output or colored stdout. Use `--format=json` for all automation.

## Known Output Hazards

1. **Service status localization**: `snapctl.Services()` returns status strings in the host OS locale (e.g., "active", "inactivo"). The CLI passes these through verbatim, making locale-dependent parsing unreliable. Documented bug: [LP#2137543](https://bugs.launchpad.net/snapd/+bug/2137543).
2. **Empty JSON/YAML key behavior**: `get` on a missing key returns an error, not empty JSON. `get` on a found scalar prints just the scalar value (not a JSON object) when not using `--format`, which is inconsistent with the machine-readable flag pattern.
3. **No `--format` on `get`/`set`/`unset`**: These config commands only emit YAML scalar or YAML document output. There is no JSON option, making them harder to use in shell pipelines.
