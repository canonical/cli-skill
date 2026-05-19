# qwen36 Configuration Model

## Observed Configuration Sources

The repository shows three effective sources of configuration:

1. Engine manifests in `engines/*/engine.yaml`
2. CLI-managed persisted values written by `qwen36 set`
3. Package-scoped defaults written by `qwen36 set --package` during snap install

The CLI likely resolves these into a single value returned by `qwen36 get <key>`.

## Known Keys

The public repository evidences these keys:

| Key | Purpose | Observed Producers | Observed Consumers |
|---|---|---|---|
| `http.port` | HTTP listen port for local server | install hook, user `set` | chat wrapper, engine launchers, health check |
| `http.host` | HTTP bind host | install hook, user `set` | engine launchers |
| `http.base-path` | API prefix such as `v1` | engine manifest defaults | chat wrapper, health check |
| `model-name` | Optional model name reported to clients | unknown CLI source | chat wrapper, health check |
| `verbose` | Enables verbose engine logs | install hook, user `set` | engine launchers |
| `server` | Selected server component name | engine selection | engine `common.sh` |
| `model` | Selected model component name | engine selection | engine `common.sh` |
| `multimodel-projector` | Selected multimodal projector component name | engine selection | engine `common.sh` |
| `gpu-layers` | Number of layers offloaded to GPU | cuda engine manifest default, user override possible | cuda engine launcher |

## Engine-Derived Defaults

The engine manifests define configuration values that appear to be written or projected into the config store when an engine is selected.

### CPU engine defaults

- `server=llamacpp`
- `model=model-qwen36-35b-a3b-ud-q4-k-xl`
- `multimodel-projector=mmproj-qwen36-35b-a3b-f16`
- `http.base-path=v1`

### CUDA engine defaults

- `server=llamacpp-cuda`
- `model=model-qwen36-35b-a3b-ud-q4-k-xl`
- `multimodel-projector=mmproj-qwen36-35b-a3b-f16`
- `http.base-path=v1`
- `gpu-layers=99`

## Package Defaults

The install hook seeds package-scoped defaults with:

- `http.port=8326`
- `http.host=127.0.0.1`
- `verbose=false`

These are written using `qwen36 set --package`, which strongly suggests a separate config layer from ordinary user writes.

## Inferred Precedence

The exact precedence is not documented, but the most defensible interpretation from the artifacts is:

1. User overrides written with plain `qwen36 set`
2. Engine-selected values or engine-derived config
3. Package defaults written with `qwen36 set --package`
4. Hard-coded fallbacks in shell wrappers

Evidence for step 4:

- `apps/chat.sh` treats an empty `http.base-path` as `v1`
- `apps/check-server-llamacpp.sh` does the same

Because the CLI source is private, the exact merge behavior should be treated as inferred rather than confirmed.

## Command-Specific Override Behavior

- `use-engine` appears to be the main command that changes multiple related keys together.
- `set` changes one key at a time.
- `get` returns the resolved value used by runtime wrappers.
- `show-engine` exposes the selected engine model in YAML form rather than exposing raw config values.

## Surprising or Risky Precedence Behavior

- The existence of `--package` is not documented in the user README, so users may not understand why a value exists or where it came from.
- Engine manifests provide defaults, but there is no documented way to inspect per-key provenance.
- `model-name` is treated as optional by scripts, which suggests config completeness can vary by engine or install state.

## Documentation Need

This CLI needs one authoritative config reference that documents:

- every key
- value types
- allowed ranges
- which commands mutate the key
- the precedence between package defaults, engine selection, and user overrides