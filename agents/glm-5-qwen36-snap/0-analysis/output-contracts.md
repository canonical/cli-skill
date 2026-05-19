# qwen36 Output Contracts

## Overview

Only two command outputs are strongly evidenced by consuming scripts:

1. `show-engine` returns YAML
2. `get` returns scalar text

The remaining commands are either interactive (`chat`) or mutation-oriented (`set`, `use-engine`) and are not documented with stable output examples.

## Per-Command Contracts

| Command | Output Mode | Known Contract | Parseability Guidance | Stability Assessment |
|---------|-------------|----------------|----------------------|---------------------|
| `qwen36 chat` | Interactive terminal session | No public transcript available. Likely terminal UI from `go-chat-client`. | Not suitable for machine parsing. Treat as human-only. | Undocumented |
| `qwen36 use-engine` | Human-readable success / prompt / error output | Not documented. Presence of `--assume-yes` implies the default path may prompt interactively. | Avoid scripting stdout text; use `show-engine` after mutation to verify state. | Undocumented |
| `qwen36 show-engine` | YAML | Must contain at least `name` and `components[]`; likely mirrors selected engine manifest and resolved configuration. | Safe for machine parsing only for fields actually consumed by scripts. Prefer tolerant YAML parsing with `yq`. | Partially stable, but undocumented |
| `qwen36 get <key>` | Plain scalar text | Returns the resolved config value for one key. Scripts rely on shell substitution semantics. | Treat stdout as a raw string with trailing newline. Validate emptiness explicitly because some missing values are tolerated. | Operationally stable, not documented |
| `qwen36 set <key>=<value>` | Likely human-readable success or silent success | No public success output documented. | Script callers should rely on exit status and optionally re-read via `get`. | Undocumented |
| `qwen36 completion bash` | Whitespace-delimited completion tokens | Output is passed to `compgen -W`, so tokenization is word-based rather than line-oriented script generation. | Suitable only for bash completion consumption. Do not assume one token per line. | Operationally stable for current completer, undocumented |

## YAML Contract for `show-engine`

The minimum schema evidenced by scripts is:

```yaml
name: <engine-name>
components:
  - <component-name>
  - <component-name>
```

The engine manifests suggest additional fields may appear, such as:

- `description`
- `vendor`
- `grade`
- `devices`
- `memory`
- `disk-space`
- `configurations`

Because only `name` and `components` are consumed by public scripts, those are the only defensible stable fields.

## Human vs Machine Readability

- `show-engine` is the only clearly machine-readable command.
- `get` is machine-usable but only as a scalar value API, not as structured output.
- `completion bash` is machine-consumable in a narrow shell-completion context.
- `chat`, `set`, and `use-engine` should be treated as human-oriented until help text or examples define otherwise.

## Output Risks

- Changing YAML field names in `show-engine` would break `apps/server.sh`.
- Changing `get` to emit labels or formatting would break shell command substitution throughout the snap.
- Changing `completion bash` from token output to full shell script text would break `apps/completion.bash`.

Per the deprecation guidance, all of those would be breaking changes and should be versioned or transitioned carefully.

## Recommendations for Output Stability

1. Document the `show-engine` YAML schema explicitly.
2. Consider adding `--format json` for machine-readable output on relevant commands.
3. Document exit codes for success/failure on all commands.
4. Establish stability guarantees for output format in documentation.
