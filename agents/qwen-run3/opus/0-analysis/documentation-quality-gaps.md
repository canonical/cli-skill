# Documentation Quality Gaps

## Help Text vs. Code Behavior

### 1. `show-engine` help mentions wrong command

The code comment says "cli use-engine" for Args documentation:
```go
// Args
// cli use-engine <engine> requires 1 argument
// cli use-engine --auto does not support any arguments
```

This comment is copy-pasted from `use-engine` and appears in `show-engine.go`. It does not affect user-facing help but indicates maintenance sloppiness.

### 2. `run` examples use `cli` instead of snap name

```go
Example: "  cli run env\n" +
    "  cli run -- echo \"Hello World!\"\n" +
    "  cli run --wait-for-components -- python3 -m http.server",
```

The examples hardcode `cli` instead of using the dynamic snap instance name. Users would see `cli` in help text instead of `qwen36`.

### 3. No examples in most command help

Only `run` provides example usage in its cobra `Example` field. All other commands lack examples, reducing discoverability for new users.

### 4. `--wait-for-components` on `run` is deprecated but not explained

The flag is marked deprecated with message `"run" always waits for components.` but this behavioral change is not documented anywhere accessible.

## README vs. Actual Commands

### 5. README shows `qwen36 use-engine --auto` without `sudo`

The README Quick Start section shows:
```bash
qwen36 use-engine --auto
```

But this command requires root and will fail without `sudo`. The README should show `sudo qwen36 use-engine --auto`.

### 6. README shows `snap logs qwen36.server` for status

The README suggests `snap logs qwen36.server` for checking status, but does not mention the `qwen36 status` command which provides structured output.

### 7. README does not mention `list-engines` or `show-machine`

The README Configuration section jumps straight to `show-engine` / `use-engine` / `get` / `set` without mentioning the discovery commands.

## Missing Documentation

### 8. No man page or `--help` long-form documentation

The CLI provides short help via `--help` but no man page, no `help <topic>` system, and no reference documentation beyond the README.

### 9. No documentation of config keys

There is no `qwen36 get --help` output or documentation listing all valid configuration keys, their types, and valid values. Users must discover keys through `qwen36 get` (dump all) or trial and error.

### 10. No documentation of engine manifest schema

The engine YAML format is undocumented. Third-party engine contributors have no specification to follow.

### 11. Hidden commands have no discoverability path

`run`, `serve-webui`, and `debug` are hidden. There is no `qwen36 --show-hidden` or documentation explaining these exist and when to use them.

## Ambiguity

### 12. `prune-cache` name is misleading

The command removes snap *components* (engine binaries, model files), not a traditional cache. The name implies removing temporary/regenerable data, but the actual operation removes installed software components that must be re-downloaded.

### 13. `--format` default inconsistency

- `list-engines` defaults to `table`
- All other commands default to `yaml`

This inconsistency is undocumented and may surprise users scripting multiple commands.
