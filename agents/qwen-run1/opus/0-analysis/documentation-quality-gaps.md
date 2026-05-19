# Documentation Quality Gaps

## README vs Actual Behavior

### Gap 1: `use-engine --auto` presented as setup step, not explained

**README says**:
```bash
qwen36 use-engine --auto
```

**Reality**: The install hook already runs `use-engine --auto --assume-yes` automatically. The README presents this as a manual step in "Quick Start" without explaining that it may have already been done during installation.

**Impact**: User confusion — they may wonder if this is needed or what it does if the engine is already selected.

### Gap 2: `--assume-yes` flag not documented

The `--assume-yes` flag on `use-engine` is used in the install hook but never mentioned in the README or any user-facing documentation.

### Gap 3: `--package` flag not documented

The `--package` flag on `set` is used in the install hook but not mentioned in user documentation. Its semantics (package-level vs user-level) are nowhere explained.

### Gap 4: No `--help` documentation

There is no documentation of what `qwen36 --help` or `qwen36 <command> --help` outputs. It's unknown whether help text exists, what it contains, or how comprehensive it is.

### Gap 5: Configuration keys not fully listed

The README shows `http.port` as an example but doesn't list all available configuration keys. Users must guess or experiment to discover `http.host`, `http.base-path`, `model-name`, `verbose`, `server`, `model`, `multimodel-projector`, `gpu-layers`.

### Gap 6: Server management not documented

The README mentions `snap logs qwen36.server` but doesn't document:
- How to restart the server after config changes
- How to stop/start the server manually
- That config changes require a server restart
- The relationship between `use-engine` and server restart

### Gap 7: No error scenario documentation

No documentation covers:
- What happens if no engine matches hardware
- What happens if components aren't installed yet
- How to troubleshoot a failed server
- What error messages mean

### Gap 8: `chat` command prerequisites not documented

The `chat` command requires the server to be running. This dependency is not stated in the README. If the server is down, the user gets no guidance from the README about what to do.

### Gap 9: Engine hardware requirements not user-visible

The README lists "cpu" and "cuda" engines with brief descriptions but doesn't explain:
- Minimum hardware requirements for each (32G RAM, specific CPU flags)
- How `--auto` detection works
- What happens if requirements aren't met

### Gap 10: `show-engine` output format not documented

The README says "View current engine" with `qwen36 show-engine` but doesn't show example output or explain that it outputs YAML. Users don't know what to expect.

## Missing Examples

| Scenario | Status |
|----------|--------|
| Basic chat session | Not shown |
| Setting multiple config values | Not shown |
| Checking server health | Not shown |
| Switching engines manually | Shown (minimal) |
| Using completion | Not shown |
| Troubleshooting failed start | Not shown |
| Checking current config | Shown (one key only) |

## Ambiguities

1. **`model-name` vs `model`**: Two config keys with similar names but different purposes — `model-name` is the API model identifier, `model` is the snap component name. This distinction is never explained.
2. **What "engine" means**: The README doesn't clearly define what an "engine" is — is it the backend? The hardware profile? The full configuration set?
3. **Component installation**: The README mentions "snap components" under "Components" but doesn't explain how they're installed or that the server waits for them.

## Outdated or Incorrect Content

No clearly outdated content identified, as the project appears to be in early development (`grade: devel`).
