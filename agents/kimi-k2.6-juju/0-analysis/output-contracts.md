# Output Contracts

## Human-readable Formats
- **default / smart**: Adapts output type (string, bool, list, YAML map) based on data shape. Used by simple getters.
- **tabular**: Column-aligned tables for lists (machines, models, users, spaces). May include ANSI color unless `--no-color` is set.
- **oneline**: Condensed single-line status, useful for shell pipelines.
- **summary**: High-level aggregation (status summary).

## Machine-readable Formats
- **yaml**: Stable structural output; default for many config/show commands.
- **json**: Explicitly serialisable; `SuperCommand.isSerialisableFormatDirective` suppresses extra stderr writes when JSON is selected.
- **format stability**: YAML/JSON field names are generally stable because they map to Go struct tags in API facades. However, tabular column ordering and human-readable text are not contractual.

## Per-command Output Behavior
| Command Category | Default Format | Supported Formats | Notes |
|---|---|---|---|
| List commands (machines, models, users) | tabular | tabular, yaml, json | `--no-color` disables ANSI |
| Show commands (show-cloud, show-model) | yaml | yaml, json | Often includes nested maps |
| Config commands | yaml | yaml, json | `config` also supports raw key output |
| Status | tabular | tabular, yaml, json, oneline, summary, short | `--relations`/`--integrations` toggle extra sections |
| Action output | yaml | yaml, json, smart | Task results are map-structured |
| Backups / storage / resources | tabular | tabular, yaml, json | Pool lists use custom formatters |

## Parseability Guidance
- Scripts should prefer `--format json` or `--format yaml` and avoid parsing tabular output.
- When `--output` is used, the file receives the formatted data; stderr still receives logs/errors unless `--quiet` or machine format is selected.
- Empty results in machine format produce `{}` (JSON) or `{}` (YAML) rather than nothing, to keep parsers stable.
