# Documentation Quality Gaps

## Help Output vs Code Behavior

Since the Go CLI submodule is not checked out, `--help` output cannot be directly verified. Analysis is based on the README.md, snapcraft.yaml, and shell scripts.

## Identified Gaps

### 1. Undocumented `--package` flag on `set`

| Aspect | Finding |
|--------|---------|
| **Gap** | The `--package` flag on `qwen36 set` is used in the install hook but not documented in the README |
| **Impact** | Users cannot understand the two-tier configuration model (package vs user settings) |
| **Source** | `snap/hooks/install` line: `qwen36 set --package http.port="8326"` |

### 2. Undocumented `--assume-yes` flag on `use-engine`

| Aspect | Finding |
|--------|---------|
| **Gap** | `--assume-yes` is used in the install hook but not mentioned in user documentation |
| **Impact** | Scripters don't know how to use `use-engine` non-interactively |
| **Source** | `snap/hooks/install` line: `qwen36 use-engine --auto --assume-yes` |

### 3. `model-name` config key purpose unclear

| Aspect | Finding |
|--------|---------|
| **Gap** | `model-name` is read by chat.sh and check-server but never set in any visible code path |
| **Impact** | Users don't know what this key does or how to set it; scripts silently ignore its absence |
| **Source** | `chat.sh`: `model_name="$(qwen36 get model-name 2>/dev/null || true)"` |

### 4. No documentation of all configuration keys

| Aspect | Finding |
|--------|---------|
| **Gap** | README shows `http.port` and `http.port` examples but doesn't list all valid keys |
| **Impact** | Users must reverse-engineer valid keys from hook scripts and engine.yaml files |
| **Missing keys** | `http.host`, `http.base-path`, `verbose`, `model-name`, `server`, `model`, `multimodel-projector`, `gpu-layers` |

### 5. Server health check behavior undocumented

| Aspect | Finding |
|--------|---------|
| **Gap** | The 3-tier exit code system (0/1/2) in check-server and the retry logic in wait-for-server are not explained anywhere |
| **Impact** | Users troubleshooting server issues have no guide for interpreting behavior |

### 6. No examples of error scenarios

| Aspect | Finding |
|--------|---------|
| **Gap** | README only shows happy-path examples |
| **Impact** | Users encountering "Missing component", timeouts, or hardware incompatibility have no reference |
| **Missing scenarios** | Component not installed, no compatible engine, server timeout, GPU not detected |

### 7. Engine hardware requirements not user-visible

| Aspect | Finding |
|--------|---------|
| **Gap** | Engine requirements (CPU flags, VRAM, RAM) are in engine.yaml but not surfaced in user docs or `--help` |
| **Impact** | Users cannot determine if their hardware will support an engine without reading source YAML |
| **Suggestion** | `show-engine` or `use-engine --list` should display requirements |

### 8. `http.base-path` fallback inconsistency

| Aspect | Finding |
|--------|---------|
| **Gap** | `chat.sh` has a hardcoded fallback (`v1`) for `http.base-path` that isn't in docs or help |
| **Impact** | Behavior diverges silently if config and code get out of sync |
| **Source** | `chat.sh`: `if [ -z "$api_base_path" ]; then api_base_path="v1"; fi` |

### 9. Missing `completion` documentation

| Aspect | Finding |
|--------|---------|
| **Gap** | README doesn't mention `qwen36 completion bash` or how to enable tab completion |
| **Impact** | Users miss the discoverability benefit of shell completions |
| **Note** | The `completer` field in snapcraft.yaml auto-registers completions, but users with non-standard setups may need manual setup |

### 10. No version command

| Aspect | Finding |
|--------|---------|
| **Gap** | There is no `qwen36 version` or `qwen36 --version` command documented or visible in the command set |
| **Impact** | Users cannot verify which version is installed via the CLI (must use `snap info qwen36`) |
| **DE013** | The Canonical CLI standard recommends `tool version` as a standard command |

## Summary

The README provides a good quick-start overview but is insufficient as a reference. The primary gaps are:
- Missing flag documentation (`--package`, `--assume-yes`)
- No comprehensive config key reference
- No error/troubleshooting guidance
- No documentation of the health check and recovery model
- Missing standard commands (`version`)
