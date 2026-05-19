# Output Contracts

## Output by Command

### qwen36 chat

| Format | Target | Stability |
|--------|--------|-----------|
| Interactive text stream | Human (terminal) | Unstable — controlled by go-chat-client |

- Output is a streaming chat conversation rendered in the terminal.
- No machine-readable output mode.
- No `--format` flag.
- Exit behavior: interactive session ends on user quit (Ctrl+C or exit command).

### qwen36 use-engine

| Format | Target | Stability |
|--------|--------|-----------|
| Human-readable text | Human (terminal) | Unstable |

- Prints status messages during engine selection (hardware detection results, selected engine name, confirmation prompt).
- No structured output mode (no `--format=json`).
- With `--assume-yes`, may suppress confirmation output.
- Expected output: engine name and success confirmation.

### qwen36 show-engine

| Format | Target | Stability |
|--------|--------|-----------|
| YAML | Machine-readable | Semi-stable (tied to engine.yaml schema) |

- Outputs the full engine.yaml content for the currently selected engine.
- Parsed by other scripts using `yq` (e.g., `qwen36 show-engine | yq .name`, `qwen36 show-engine | yq .components[]`).
- Fields include: `name`, `description`, `vendor`, `grade`, `devices`, `memory`, `disk-space`, `components`, `configurations`.
- **Parseability**: YAML output is the primary machine interface. Stable as long as engine.yaml schema doesn't change.

### qwen36 get

| Format | Target | Stability |
|--------|--------|-----------|
| Plain text (single value) | Machine-readable | Stable |

- Outputs the raw value of the requested configuration key, with a trailing newline.
- No formatting, no key echo, no decoration.
- Empty output (no newline) if the key is unset.
- **Parseability**: Trivially parseable — single line of text. Used extensively in shell scripts via command substitution.

### qwen36 set

| Format | Target | Stability |
|--------|--------|-----------|
| Silent on success | Both | Stable |

- Produces no output on success (exit code 0).
- Error messages on stderr for invalid keys or values.
- Matches `snapctl set` behavior.

### qwen36 completion bash

| Format | Target | Stability |
|--------|--------|-----------|
| Space-separated word list | Machine (shell completion system) | Semi-stable |

- Outputs a space-separated list of valid completion words for the current context.
- Consumed by `completion.bash` via `compgen -W`.
- The word list changes as commands/subcommands are added.

### qwen36.server (daemon)

| Format | Target | Stability |
|--------|--------|-----------|
| Log lines to stdout/stderr | Human (journalctl/snap logs) | Unstable |

- Server startup messages, component wait status, and llama-server output.
- Logs accessible via `snap logs qwen36.server` or `journalctl`.
- llama-server output format is upstream-controlled and not stable.

## Stability Expectations

| Command | Output Stability | Breaking Change Risk |
|---------|-----------------|---------------------|
| `get` | High — raw value, used in scripts | Low |
| `set` | High — silent success | Low |
| `show-engine` | Medium — YAML schema could evolve | Medium (scripts depend on field names) |
| `completion bash` | Medium — word list grows with new commands | Low |
| `use-engine` | Low — human-readable, informational | Low |
| `chat` | Low — interactive, upstream-dependent | Low |
| `server` | Low — llama-server logs, not a contract | Low |

## Machine-Readable Output Gaps

- No `--format=json` or `--format=yaml` option on any command.
- `show-engine` outputs YAML but this is incidental (it dumps the file) rather than a designed contract.
- No `--quiet` flag to suppress informational output.
- No structured error output (errors are free-form text on stderr).
