# Output Contracts

## Summary

The qwen36 surface mixes three output styles:

- human-interactive output for `chat` and engine-selection flows
- strict machine-readable scalar/YAML output for `get`, `show-engine`, and completion generation
- daemon logs and HTTP API responses for `qwen36.server`

Only two contracts are clearly load-bearing in checked-in code: `qwen36 show-engine` YAML and `qwen36 get` scalar output. Those are consumed directly by scripts and should be treated as operationally stable even though they are undocumented.

## Command output matrix

| Command | Primary output form | Human-readable behavior | Machine-readable behavior | Parseability guidance | Stability expectation |
|---|---|---|---|---|---|
| `qwen36 chat` | interactive terminal session | User-facing conversation via `go-chat-client` | none observed on stdout beyond the interactive client flow | Do not script against terminal transcript output | Low stability for terminal text; user UX surface only |
| `qwen36 use-engine` | likely confirmation + success/error text | README and install-hook usage imply an interactive or affirmative mutation flow | none observed | Use exit status, then verify with `show-engine` or `get`; do not parse success text | Undocumented and should be treated as unstable text |
| `qwen36 show-engine` | YAML document | May be readable by humans, but clearly intended for structured inspection | yes: parsed by `yq` in scripts | Parse with tolerant YAML tooling; fields `name` and `components[]` are required by runtime scripts | Operationally high stability for `name` and `components[]`; broader schema undocumented |
| `qwen36 get <key>` | plain scalar text | Minimal, shell-friendly value output | yes: shell scripts consume raw stdout | Treat stdout as a bare value with trailing newline; handle empty values and missing-key failures explicitly | Operationally high stability because scripts depend on it |
| `qwen36 set <key>=<value>` | likely silent success or short status text | mutation acknowledgement is possible but not documented | none observed | Prefer exit status plus follow-up `get` for automation | Undocumented; assume text is unstable |
| `qwen36 completion bash` | whitespace-delimited completion candidates | not intended as a direct user-facing report | yes: consumed by bash completion script | Treat stdout as a completion word list; stderr is suppressed by the completer | Operationally medium-to-high stability for bash integration |
| `qwen36.server` | daemon logs plus local HTTP API | startup, wait, and failure messages are human-readable | yes, but over network rather than stdout: OpenAI-compatible API responses | Parse API responses from the local HTTP endpoint, not process stdout | Logs are human-facing; HTTP API shape is expected to follow OpenAI-compatible conventions |

## Per-command details

### `qwen36 chat`

Observable contract:

- `chat.sh` exports `OPENAI_BASE_URL` and `MODEL_NAME`
- then execs `go-chat-client`

Implications:

- The visible terminal output belongs to `go-chat-client`, not the shell wrapper.
- The wrapper itself is effectively silent on success.
- Any scripted use should target the underlying HTTP API instead of the interactive CLI.

### `qwen36 use-engine`

Observable contract:

- no stdout format is documented in the repository
- `--assume-yes` implies the default flow may prompt or otherwise request confirmation

Guidance:

- treat stdout/stderr text as human-facing only
- automation should use exit status plus a postcondition check such as `qwen36 show-engine`

### `qwen36 show-engine`

Observed machine-readable guarantees from code usage:

- `apps/server.sh` requires `.components[]`
- `apps/server.sh` and `apps/wait-for-server.sh` require `.name`

Minimum safe schema:

```yaml
name: <engine-name>
components:
  - <component-name>
  - <component-name>
```

Additional fields likely mirror the engine manifest and may include description, vendor, grade, device requirements, memory, disk space, and configuration-derived values. Only `name` and `components[]` are proven load-bearing from checked-in code.

### `qwen36 get <key>`

Observable contract:

- output must be a raw scalar suitable for command substitution, for example:
  - `port="$(qwen36 get http.port)"`
  - `verbose="$(qwen36 get verbose)"`
- some callers tolerate missing values by suppressing stderr and non-zero exit (`model-name`)

Implications:

- the command must not add framing text like `http.port: 8326`
- the command should preserve exact value content except for the trailing newline

### `qwen36 set <key>=<value>`

Observable contract:

- no checked-in script depends on stdout text from `set`
- hooks and users rely on mutation side effects only

Guidance:

- for automation, treat exit status as authoritative
- if the CLI ever emits success text, it should remain human-readable rather than machine-formatted unless a new format flag is added

### `qwen36 completion bash`

Observable contract:

- `apps/completion.bash` uses:

```bash
compgen -W "$($SNAP/bin/qwen36 completion bash 2>/dev/null)" -- "$cur"
```

Implications:

- stdout must contain completion words separated in a way `compgen -W` can consume
- stderr noise breaks completion UX and is therefore intentionally suppressed
- bash completion is effectively an API surface between the CLI and the shell script

### `qwen36.server`

There are two output channels:

1. Process/log output
   - `server.sh` prints wait messages such as `Waiting for required snap components: [...]`
   - on timeout it prints a remediation hint about `snap changes`
   - engine launch scripts print concrete setup errors like `Missing component: ...` and `gpu-layers snap option is not set`

2. Local HTTP API output
   - `check-server-llamacpp.sh` posts to `/v1/completions` (or another configured base path)
   - it expects JSON and checks for `.error` and `.choices[0].text`
   - this strongly implies an OpenAI-compatible JSON response schema

## Known load-bearing output fields

| Producer | Field / shape | Consumer |
|---|---|---|
| `show-engine` | `.name` | `apps/server.sh`, `apps/wait-for-server.sh` |
| `show-engine` | `.components[]` | `apps/server.sh` |
| `get` | raw scalar stdout | `apps/chat.sh`, `apps/check-server-llamacpp.sh`, engine launch scripts |
| HTTP completions endpoint | `.error.message` | `apps/check-server-llamacpp.sh` |
| HTTP completions endpoint | `.choices[0].text` | `apps/check-server-llamacpp.sh` |

## Contract risks

- `show-engine` is machine-readable but undocumented; changing field names would break daemon startup.
- `get` appears simple, but adding labels or color would break shell consumers immediately.
- `completion bash` is an implicit API and should be versioned carefully if more shells are introduced.
- `use-engine` and `set` lack a documented automation contract, so scripts should avoid parsing their stdout.

## Assessment

The output model is workable because the truly script-critical paths are narrow and simple. The main gap is documentation rather than design: the repository already treats `show-engine`, `get`, and the local HTTP API as contracts, but those contracts are not surfaced to users or contributors.