# qwen36 Output Contracts

## Output Formats By Command

| Command | Default Output | Supported Formats | Intended Audience |
|---|---|---|---|
| `status` | YAML | `yaml`, `json` | human and scripts |
| `chat` | interactive terminal stream | none | human only |
| `webui` | human-readable messages | none | human only |
| `get` | scalar text or YAML object | none | human and shell scripts |
| `set` | silent success unless prompts/errors occur | none | human and scripts by exit status |
| `unset` | silent success unless prompts/errors occur | none | human and scripts by exit status |
| `list-engines` | table | `table`, `json` | human and scripts |
| `show-engine` | YAML | `yaml`, `json` | human and scripts |
| `use-engine` | human-readable evaluation and confirmation text | none | human first |
| `show-machine` | YAML | `yaml`, `json` | human and scripts |
| `prune-cache` | human-readable component list and prompt | none | human first |
| `version` | YAML | `yaml`, `json` | human and scripts |

## Structured Output Contracts

### `status`

`status` emits a struct with these top-level fields:

- `engine`
- `services`
- `endpoints`
- `model`

Stability expectation: medium

- field names are code-defined and likely stable for current scripts
- nested service status values are locale-dependent because snapd service strings are passed through verbatim

### `show-engine`

`show-engine` emits YAML or JSON based on `common.EngineDetails`, which includes manifest data plus compatibility annotations.

Important fields observed or directly implied:

- `name`
- `description`
- `vendor`
- `grade`
- `components`
- `configurations`
- `score`
- `compatible`
- `compatibility-issues`

Operational stability: high for `name` and `components`

Reason: `apps/server.sh` depends on `.name` and `.components[]` when selecting and starting the server.

### `list-engines`

`list-engines --format=json` emits:

- `active-engine`
- `engines` array

Table mode is presentation-only:

- active engine is denoted with a trailing `*`
- compatibility column collapses state into `yes`, `devel`, or `no`
- widths are dynamic and should not be parsed

### `show-machine`

`show-machine` prints the hardware information struct from `pkg/types.HwInfo`.

Stability expectation: medium

- field names are struct-based
- hardware warning strings are not stable and appear on stderr only when verbose

### `version`

`version` emits a two-field object:

- `snap`
- `cli`

Empty values are normalized to `unset`.

## Plain-Text Contracts

### `get`

`get` is unusual because it has no `--format` flag.

Observed behavior:

- with one scalar value, it prints only the scalar followed by newline
- with an object or subtree, it prints YAML
- with no key, it prints all merged config as YAML

This means machine users must handle both scalar and YAML cases.

### `set` and `unset`

- success is mostly silent except for optional restart prompts
- failure is reported on stderr and returns non-zero
- scripts should verify success by exit code and, if needed, a follow-up `get`

### `use-engine`

Typical stdout path includes:

- compatibility evaluation lines
- selected engine line
- component installation list and sizes
- optional cancellation or restart prompt text

This is not stable machine-readable output.

## Interactive And Streaming Output

### `chat`

- prints connection and readiness messages
- then starts a streaming REPL
- response tokens are written directly to stdout
- reasoning fragments are visually distinguished in the client implementation, which is presentation detail, not a script contract

### `webui`

- prints the URL
- may prompt the user to press Enter to open a browser
- is unsuitable for automation as-is

## Parseability Guidance

Recommended for scripts:

- prefer `status --format=json`
- prefer `show-engine --format=json` or YAML with `yq`
- prefer `list-engines --format=json`
- avoid table output
- treat `get` as a shell-scalar API, not a general structured API

## Stability Risks

1. changing `show-engine` field names would break server startup wrappers
2. changing `get` to print labels would break shell substitutions across wrappers
3. localized service states in `status.services` make strict parsing fragile
4. several mutation commands lack any structured success output, forcing automation to depend entirely on exit status
