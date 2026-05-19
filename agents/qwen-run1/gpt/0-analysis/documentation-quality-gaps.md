# qwen36 Documentation Quality Gaps

## Summary

The README is sufficient for a happy-path demo, but it is not sufficient as command reference documentation. The largest gap is that the private CLI is treated as self-explanatory even though the public repository does not expose help text, output contracts, flag semantics, exit codes, or configuration precedence.

## Findings

| Severity | Gap | Evidence |
|---|---|---|
| High | No authoritative command reference exists for the six observable leaf commands. | README only shows a few examples and no per-command syntax blocks. |
| High | `set --package` is operationally important but undocumented. | Snap install hook uses it; README never mentions it. |
| High | `show-engine` returns YAML consumed by scripts, but the schema is undocumented. | `apps/server.sh` parses `.name` and `.components[]`. |
| High | `use-engine` flag semantics are incomplete. | README shows `cpu`, `cuda`, and `--auto`; only hooks reveal `--assume-yes`. |
| High | No exit code documentation exists. | Public docs contain no exit code table; shell wrappers rely on code-based behavior. |
| Medium | Config key reference is incomplete. | User prompt listed keys, but README only shows `http.port` and engine changes. |
| Medium | No explanation of config precedence or where values originate. | Engine manifests, package defaults, and user writes all participate. |
| Medium | Completion behavior is undocumented. | `completion bash` exists for the completer but is absent from README. |
| Medium | `chat` does not explain its dependency on the local server or expected server state. | README says `qwen36 chat`; wrappers show server URL construction and health assumptions. |
| Medium | No failure examples are provided. | README has no unsupported-engine, missing-component, or server-timeout examples. |
| Low | The server daemon relationship is implied, not clearly explained. | README mentions auto-start and `snap logs qwen36.server` but not the control-plane/data-plane split. |

## Mismatches Between Docs And Observed Behavior

1. README documents `use-engine cpu` and `use-engine cuda`, but not `--assume-yes` even though packaging depends on it.
2. README documents `show-engine` only as "View current engine" and does not say that the output is YAML.
3. README documents `get` and `set` without listing the full key space.
4. README does not mention `completion bash`, even though the snap app exposes bash completion through it.

## Missing Examples

The docs should add examples for:

- `qwen36 use-engine --auto`
- `qwen36 use-engine --auto --assume-yes`
- `qwen36 show-engine` with sample YAML
- `qwen36 get gpu-layers`
- `qwen36 set verbose=true`
- `qwen36 completion bash`

## Ambiguities

- Are `cpu` and `cuda` the only valid engines?
- Which keys are user-supported versus internal?
- Does changing config require restarting the daemon?
- What happens when `get` is called for an unknown key?
- Is `model-name` expected to be unset on some installs?

## Documentation Improvements To Prioritize

1. Publish a concise command reference page with syntax, flags, and examples.
2. Publish a config key table with types, defaults, and provenance.
3. Document the `show-engine` YAML schema and its intended stability.
4. Document failure modes and exit codes for `get`, `set`, and `use-engine`.
5. Add shell completion documentation, including what `completion bash` emits.