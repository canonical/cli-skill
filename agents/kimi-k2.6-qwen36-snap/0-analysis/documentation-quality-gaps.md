# Documentation Quality Gaps

## Help Text vs. Code Behavior

### 1. `set` Long Description is Minimal
- **Help text**: "Set a configuration"
- **Actual behavior**: Supports multiple `key=value` pairs, rejects unknown keys (except passthrough), applies to three precedence layers (user default, but hidden flags for package/engine), triggers snap restart, and validates atomically.
- **Gap**: Users cannot discover the key-value syntax rules, the passthrough exemption, or the restart side effect from `--help`.

### 2. `unset` Long Description is Partially Misleading
- **Help text**: "Unset a user configuration, reverting to package or engine default. If no default exists for the key, it will be removed entirely."
- **Actual behavior**: Only ever unsets from the `UserConfig` layer. The key may still exist in `EngineConfig` or `PackageConfig`, in which case it is not "removed entirely" from the snap's config store—only the user override is removed.
- **Gap**: The phrase "removed entirely" implies deletion from all layers, which is incorrect.

### 3. `use-engine` Lacks Mode Documentation
- **Help text**: "Select an engine" with no further explanation.
- **Actual behavior**: Has three mutually exclusive invocation modes (explicit name, `--auto`, `--fix`) with different semantics and confirmation flows.
- **Gap**: `--help` does not explain what `--fix` does, that `--auto` evaluates hardware compatibility, or that providing a name plus a flag is illegal.

### 4. `run` Example Uses Deprecated Flag
- **Help text example**: `cli run --wait-for-components -- python3 -m http.server`
- **Actual behavior**: `--wait-for-components` is deprecated and ignored; the command always waits.
- **Gap**: The embedded example advertises a deprecated flag, teaching users an obsolete pattern.

### 5. `prune-cache` Has No Description of Default vs. Scoped Behavior
- **Help text**: "Remove cached data"
- **Actual behavior**: Without `--engine`, removes all components not used by the active engine. With `--engine`, removes components specific to that engine. Cannot prune active engine.
- **Gap**: The default (prune everything inactive) is surprising and not explained.

### 6. `status` Missing `--wait-for-components` Explanation
- **Help text**: "Show the status of the inference snap"
- **Actual behavior**: `--wait-for-components` blocks up to 3600s polling for snap components before reporting.
- **Gap**: Users may not understand that status can hang for a long duration if components are pending.

## Missing Examples

| Command | Missing Example |
|---|---|
| `set` | No example of setting `passthrough.environment.*` |
| `unset` | No example at all |
| `get` | No example of retrieving a multi-value object key |
| `list-engines` | No example of JSON usage for scripting |
| `show-engine` | No example of querying an inactive engine |
| `use-engine` | No example of `--fix` or `--auto` |
| `prune-cache` | No example of `--engine` scoped pruning |
| `show-machine` | No example of piping to `debug select-engine` |
| `debug select-engine` | No example of piped stdin invocation in help text |

## Outdated Guidance

1. **README refers to `stack-utils`**: The CLI README still describes installation as `snap install stack-utils` and `snap alias stack-utils inference-snaps-cli`. The actual product snap is `qwen36`, and the CLI binary is renamed to `qwen36` during the snap build (`bin/cli` → `bin/qwen36`). This creates confusion about the real command name.
2. **Snapcraft note about NVIDIA drivers**: The README recommends `nvidia-driver-550-server` and `nvidia-utils-550-server`, but the `snapcraft.yaml` uses `nvidia-utils-580` in stage-packages. The version mismatch could mislead users on clean 24.04 installs.
3. **Cobra `Use` field says `cli` in examples**: The `run` command's embedded example string hardcodes `cli run env` instead of `<snap-name> run env`, which is incorrect once the snap is installed.

## Ambiguity in Error Messages and Suggestions

1. **"key is not found" during `set`**: The suggestion says `Use "qwen36 get" to view available keys`, but `get` lists all keys without indicating which are valid for mutation or distinguishing user vs. read-only keys.
2. **No engine name specified for `use-engine`**: Returns `engine name not specified` with no further guidance. It should suggest `list-engines`, `show-engine`, or `--auto`.
3. **Server error suggestions**: `SuggestServerLogs()` assumes the service name is always `<snap>.server`. This is true today but not dynamically discovered, leaving a TODO in the code that affects user-facing advice.

## README / Online Docs Coverage

- **Online CLI reference**: The README links to `https://documentation.ubuntu.com/inference-snaps/reference/models-cli/`. No evidence in-repo of that documentation source; it may drift from actual CLI behavior.
- **No local man page or markdown CLI reference**: All documentation is either the README or external web docs.
- **Completion script is minimal**: `completion.bash` only wraps `qwen36 completion bash`, which itself only exposes command names. No per-flag completion (e.g., `--format` values) is documented or implemented.

## Suggestion Synthesis

| Priority | Fix |
|---|---|
| High | Rewrite `set`, `unset`, `use-engine` long descriptions and add examples |
| High | Remove deprecated `--wait-for-components` from `run` example text |
| High | Update README to reference `qwen36` instead of `stack-utils` |
| Medium | Add `--help` examples to `prune-cache`, `debug select-engine`, and `status` |
| Medium | Document the three config layers (`package`, `engine`, `user`) in help text or a top-level `config` group description |
| Low | Generate per-flag completion documentation or switch to Cobra native completion |
| Low | Add a local `docs/reference.md` that can be versioned alongside code changes |
