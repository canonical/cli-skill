# Juju CLI Output Contracts

## Output Formats

Juju supports multiple output formats via the `--format` flag. Not all commands support all formats.

| Format | Typical Use | Stability |
|---|---|---|
| `yaml` | Human-readable structured data | Stable for scripting |
| `json` | Machine-readable structured data | Stable for scripting |
| `tabular` | Default for list/status commands | Not stable for parsing |
| `summary` | Condensed status view | Human-only |
| `oneline` | Single-line status | Human-only |
| `smart` | Adaptive formatting | Human-only |

### Format Support by Command Category

| Category | Default Format | Supported Formats |
|---|---|---|
| Status | tabular | yaml, json, tabular, summary, oneline |
| List commands (clouds, models, machines, etc.) | tabular | yaml, json, tabular |
| Show commands (show-cloud, show-model, etc.) | yaml | yaml, json, tabular |
| Config commands | yaml | yaml, json |
| Action output | yaml | yaml, json |
| Backup/Restore | tabular/human | tabular |

## Output Stability

### Machine-Readable Output (yaml, json)

- **Stable field names**: Top-level keys in yaml/json output are considered stable within a major version.
- **Additive changes only**: New fields may be added in minor versions; existing fields are not removed or renamed within a major version.
- **No structural guarantees for nested data**: The structure of nested maps and lists may evolve.
- **Empty collections**: Empty lists are rendered as `[]` in yaml/json, not omitted.

### Tabular Output

- **Not parseable**: Tabular output uses space-delimited columns with variable widths. The column order and presence of headers may change.
- **Headers**: Uppercase, bold when color is enabled. Can be suppressed with `--no-headers`.
- **Column width**: Responsive; columns may be truncated or omitted in narrow terminals.
- **Color**: Enabled when stdout is a TTY and `NO_COLOR` is not set.

### Empty States

| Context | Behavior |
|---|---|
| Tabular list, zero items | Message to stderr: "No <things> found." or similar. Exit code 0. |
| Machine format, zero items | Empty document (`{}`, `[]`, or key with empty list) |
| Show command, not found | Error to stderr, exit code 1 |

## Notable Output Behaviors

### Status Command

`juju status` is the most complex output command. It produces a deeply nested structure:

```yaml
model:
  name: default
  type: iaas
  controller: mycontroller
  cloud: aws
  region: us-east-1
  version: 3.5.0
machines:
  "0":
    juju-status:
      current: started
    dns-name: 10.0.0.1
applications:
  mysql:
    charm: mysql
    units:
      mysql/0:
        workload-status:
          current: active
```

The structure varies by:
- Model type (IAAS vs CAAS)
- Presence of containers
- Cross-model relations
- Storage attachments

### Debug Log

`juju debug-log` streams log lines in a fixed text format:
```
<timestamp> <entity> <level>: <message>
```

This is not structured and not suitable for machine parsing.

### Action Output

`juju run` and `juju exec` return structured output per unit. The output schema depends on the action being run. For `exec`, stdout/stderr are captured as strings.

### Bundle Export

`juju export-bundle` produces a bundle yaml that can be fed back into `juju deploy`. The output is a stable contract for bundle interchange.

## Redirection and Piping

When stdout is not a TTY:
- Color is automatically disabled
- Tabular output may still include headers unless `--no-headers` is used
- Progress spinners and ephemeral output are suppressed

## Output File (`-o` / `--output`)

Some commands support writing output directly to a file instead of stdout. This is most common with:
- `juju create-backup --filename`
- `juju download-backup --filename`
- `juju download --filepath`

## Warnings and Progress

- Warnings go to stderr
- Progress messages (e.g., during deploy) may be ephemeral (overwritten) when running in a TTY
- Non-TTY mode: progress is printed as sequential lines
