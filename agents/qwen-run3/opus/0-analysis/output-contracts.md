# Output Contracts

## Output Formats by Command

| Command | Human-Readable | Machine-Readable | Default |
|---------|---------------|-----------------|---------|
| `status` | YAML | JSON, YAML | YAML |
| `chat` | Interactive TUI | — | Interactive |
| `webui` | Opens browser | — | — |
| `get` | Plain value (single key), YAML (multiple) | YAML | Plain/YAML |
| `set` | Confirmation prompts, success message | — | Human |
| `unset` | Confirmation, revert message | — | Human |
| `list-engines` | Colored table | JSON | Table |
| `show-engine` | YAML | JSON, YAML | YAML |
| `use-engine` | Progress spinner, confirmation | — | Human |
| `show-machine` | YAML | JSON, YAML | YAML |
| `prune-cache` | Confirmation, component removal log | — | Human |
| `version` | YAML | JSON, YAML | YAML |

## Format Flag Behavior

The `--format` flag is available on observation commands but with inconsistent value sets:

- `status`, `show-engine`, `show-machine`, `version`: support `json` and `yaml`
- `list-engines`: supports `table` and `json` (no YAML)
- `get`: no `--format` flag; output format is implicit (single value = plain, multiple = YAML)

## Stability Expectations

### Stable outputs (machine-parseable)

- **JSON output** from `--format json` on any command: fields should be considered stable within a major snap track
- **YAML output** from `--format yaml`: same stability as JSON
- **`get` single-key output**: prints raw value with a trailing newline — stable contract used by shell scripts (e.g., `$(qwen36 get http.port)`)

### Unstable outputs

- **Table output** from `list-engines --format table`: colored, human-oriented, column widths may change
- **Progress spinners and confirmation prompts**: interactive text, not parseable
- **Chat output**: delegated to `go-chat-client`, not controlled by this CLI

## Parseability Guidance

1. For scripting, always use `--format json` where available
2. For single config values, use `qwen36 get <key>` which outputs a bare value
3. The `show-engine` YAML output is consumed by other snap scripts (`server.sh` pipes through `yq`)
4. Error messages go to stderr; exit codes are documented in the error model

## Known Gaps

1. **`get` has no `--format` flag**: Cannot request JSON output for config values. Scripts using `get` must parse YAML if requesting all values.
2. **`list-engines` has no YAML format**: Inconsistent with other commands that offer both JSON and YAML.
3. **No `--quiet` / `--silent` flag**: Commands with spinners have no way to suppress interactive output for CI/scripting contexts.
