# Output Contracts

## Overview

Juju commands produce output in multiple formats: human-readable tabular output, machine-readable JSON/YAML, and structured status output. The default format is `smart`, which adapts based on the value type.

## Supported Output Formats

| Format | Type | Description | When Used |
|--------|------|-------------|-----------|
| `smart` | Human | Context-dependent heuristic (string → plain, complex → YAML) | Default for most commands |
| `tabular` | Human | Column-aligned table with bold headers | `status`, `machines`, `controllers`, `models`, `storage`, etc. |
| `yaml` | Machine | YAML serialization of the result struct | `--format yaml` |
| `json` | Machine | JSON serialization of the result struct | `--format json` |
| `oneline` | Human | Single-line condensed output | `status --format oneline` |
| `summary` | Human | Abbreviated summary | `status --format summary` |

## Format Selection

Commands that embed `cmd.Output` expose the `--format` flag. The default formatters are registered in `cmd.DefaultFormatters`:

```go
var DefaultFormatters = formatters{
    "smart": TypeFormatter{Formatter: FormatSmart, Serialisable: false},
    "yaml":  TypeFormatter{Formatter: FormatYaml, Serialisable: true},
    "json":  TypeFormatter{Formatter: FormatJson, Serialisable: true},
}
```

The `Serialisable` flag indicates whether the format is stable for scripting (JSON/YAML are serialisable; smart is not).

## Per-Command Output Details

### status

- **Default**: `tabular`
- **Alternatives**: `json`, `yaml`, `oneline`, `summary`
- **Tabular output**: Sections for model, applications, units, machines, relations. Sections toggleable via `--relations`, `--storage`
- **Stability**: High. The status output is the primary user-facing output, widely scripted against.
- **Notes**: `summary` format gives a compact one-line-per-unit view

### List Commands (models, controllers, users, machines, spaces, storage, etc.)

- **Default**: `tabular`
- **Alternatives**: `json`, `yaml`
- **Tabular output**: Column headers in uppercase bold. Columns: NAME, STATUS, etc.
- **Empty state**: When no items exist, some commands print column headers with no rows (violating DE013 guidance), while others print a message to stderr. Behavior is inconsistent.

### Show Commands (show-model, show-controller, etc.)

- **Default**: `yaml` (implicit via smart)
- **Alternatives**: `json`, `yaml`
- **Output**: Full detail for a single resource. YAML block with fields.
- **Notes**: `--show-secrets` or `--show-password` flags control sensitive field visibility.

### config / model-config / controller-config

- **Default**: `tabular` for listing all keys; `smart` (YAML) for specific keys
- **Outputs**: Tabular format with KEY, VALUE columns. When retrieving a single key, outputs plain value.
- **Machine-readable**: `--format yaml` or `--format json`
- **Notes**: Sensitive values are obscured by default unless `--show-secrets` is used (controller-config).

### bootstrap

- **Default**: Human-readable progress with ephemeral and non-ephemeral feedback.
- **Machine-readable**: Not supported. Bootstrap output is designed for interactive use.
- **Notes**: Progress reporting uses carriage-return overwriting for ephemeral lines and newlines for permanent information.

### deploy / refresh

- **Default**: Human-readable output showing deployment progress, resource downloads.
- **Machine-readable**: Not directly. `--dry-run` provides preview without side effects.
- **Notes**: Deploy progress includes charm download progress, machine allocation, and application status. Refresh shows the new revision or channel.

### Destroy Commands

- **Default**: Human-readable confirmation prompts and progress output.
- **Machine-readable**: Output goes to stderr (prompts) and stdout (progress).
- **Notes**: These are interactive by default. `--no-wait` skips progress display.

### run / exec

- **Default**: Colored output showing unit name, output, and exit code.
- **Machine-readable**: `--format json` or `--format yaml`
- **Notes**: Output interleaves stdout/stderr from the remote execution.

### backups

- **Default**: Human-readable output for `create-backup` (filename, ID). `download-backup` writes file bytes to `--output` path.
- **Machine-readable**: `--format yaml` or `--format json`

### charmhub (find, info, download)

- **Default**: `tabular` for `find` and `info`. Binary output for `download`.
- **Machine-readable**: `--format yaml` or `--format json`

## Output Stability Guarantees

1. **JSON/YAML output is considered stable** (Serialisable = true). Fields may be added but not removed or renamed in minor versions.
2. **Tabular output** column ordering and header names may change between versions.
3. **Smart format** has no stability guarantees.
4. **Progress output** (bootstrap, deploy) has no stability guarantees and is designed for interactive terminals only.
5. **Empty states**: Not consistent. Some commands return no output, others return headers. Per DE013, empty tables should print a message to stderr and return exit code 0.

## Output Destination

- **stdout**: Command results, table output, JSON/YAML output
- **stderr**: Errors, warnings, deprecation notices, prompts, progress, empty state messages
- The `-o, --output <path>` flag redirects stdout to a file for commands supporting it

## ANSI Color

- Colors are disabled by default (opt-in via `--color` or opt-out via `--no-color`)
- When `NO_COLOR` environment variable is set, colors are suppressed
- When output is piped (not a terminal), colors are automatically suppressed
- Uses basic ANSI colors (not 256-color)
