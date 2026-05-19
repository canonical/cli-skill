# Documentation Quality Gaps

## Summary

The README covers the happy path well enough to install the snap, switch engines, and start chat. The main gaps are around contract details that the implementation already depends on: undocumented flags, undocumented config keys, undocumented YAML output, and missing operational guidance for the daemon and component-install wait path.

## Gap inventory

| Severity | Gap | Evidence in docs | Evidence in code | Why it matters |
|---|---|---|---|---|
| High | `show-engine` output shape is undocumented | README says “View current engine” but gives no sample output | `server.sh` parses `.name` and `.components[]` from YAML | Users and contributors cannot know this is a machine-readable contract |
| High | `use-engine --auto --assume-yes` is undocumented | README mentions `--auto` only | install hook uses `qwen36 use-engine --auto --assume-yes` | A visible command flag exists in implementation but not in user docs |
| High | `set --package` is undocumented | README documents only plain `set` | install hook uses `qwen36 set --package ...` | Hidden write mode affects config provenance and future maintenance |
| High | Valid config key space is undocumented | README shows only `http.port` | scripts also read `http.host`, `http.base-path`, `verbose`, `server`, `model`, `multimodel-projector`, `gpu-layers`, `model-name` | Users must guess keys and cannot tell which are safe to change |
| Medium | Auto engine selection behavior is overstated in the happy path | README says server auto-starts with the best engine | install hook skips auto-selection when `hardware-observe` is unavailable and only explains this in hook logs | Users may expect auto-selection even in environments where it is skipped |
| Medium | Daemon/component wait behavior is undocumented | README says check `snap logs qwen36.server` | `server.sh` may wait up to 3600s for components, then instructs users to use `snap changes` | Long install/start waits are surprising without explanation |
| Medium | Chat readiness errors are undocumented | README says `qwen36 chat` | `wait-for-server.sh` can time out or fail with explicit messages | Users need to know chat depends on a healthy local service |
| Medium | CUDA-specific config requirement is undocumented | README says CUDA requires GPU with 10GB+ VRAM | `engines/cuda/server` also requires `gpu-layers` config to be set | Direct config edits can break CUDA startup |
| Medium | `model-name` optionality is undocumented | no mention in README | chat and health check suppress missing-key failure for `model-name` | Users cannot tell which keys are required versus optional |
| Low | `completion bash` is undocumented | no completion section in README | `apps/completion.bash` depends on `qwen36 completion bash` | Shell integration exists but is undiscoverable |
| Low | API base-path fallback is undocumented | no mention in README | scripts default empty `http.base-path` to `v1` | Troubleshooting endpoint issues is harder than necessary |
| Low | No explicit service-management guidance beyond logs | README mentions `snap logs` only | actual runtime often needs `snap changes`, `snap restart`, or component awareness | Operational support path is incomplete |

## Doc-to-code mismatches

### `show-engine`

README:

```bash
qwen36 show-engine
```

Implementation reality:

- returns YAML
- runtime scripts parse it programmatically
- fields `name` and `components[]` are required in practice

Mismatch: the docs present it as a casual inspection command, while the implementation treats it as a machine-readable contract.

### `use-engine`

README examples:

- `qwen36 use-engine --auto`
- `qwen36 use-engine cpu`
- `qwen36 use-engine cuda`

Implementation reality:

- install hook also uses `--assume-yes`
- command likely performs multi-key config mutation, not just a simple symbolic switch

Mismatch: the docs do not explain confirmation behavior, overwrite scope, or engine-managed config.

### `set`

README example:

```bash
qwen36 set http.port=8326
```

Implementation reality:

- install hook uses a privileged-looking `--package` mode
- other keys materially affect server startup, not just a port number

Mismatch: docs make config mutation look narrow and simple, while the actual config surface is broader.

## Missing examples that would help most

1. `qwen36 show-engine` with sample YAML output.
2. `qwen36 get http.host` and `qwen36 get verbose`.
3. `qwen36 set verbose=true` and `qwen36 set gpu-layers=99` with caveats.
4. `qwen36 completion bash` with shell-install instructions.
5. Recovery examples for chat/server failures, including `snap changes` and `snap logs qwen36.server`.

## Missing conceptual documentation

- what `qwen36.server` is and how it relates to `qwen36 chat`
- which config keys are user-tunable versus engine-managed
- what `use-engine` actually changes under the hood
- when auto engine selection happens and when it does not
- which outputs are intended for machine parsing (`show-engine`, `get`, completion)

## Recommended doc fixes

### README additions

- Add a `Configuration keys` section with a table of observed keys and cautions.
- Add a `Engine inspection` section with sample `show-engine` YAML.
- Add a `Shell completion` section.
- Add a `Troubleshooting` section for startup waits, component downloads, and chat timeouts.

### Help-text additions

When the Go CLI is available, add or verify help text for:

- `use-engine --auto`
- `use-engine --assume-yes`
- `set --package` if it remains user-visible
- `completion bash`

### Contract documentation

Document that the following outputs are relied upon programmatically:

- `show-engine` YAML `name` and `components[]`
- `get` raw scalar stdout
- bash completion word list output

## Assessment

The docs are not wrong so much as incomplete. The repo already contains a stronger operational and machine-readable contract than the README suggests. Closing that gap would materially improve discoverability, scriptability, and contributor safety.