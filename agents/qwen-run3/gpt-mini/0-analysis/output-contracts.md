# Output Contracts

Many commands support both machine-readable and human-readable outputs. Common formats: `yaml`, `json`, and `table`.

Commands and output expectations:
- `status`: `yaml` (default) or `json`. Structure: `{engine, services, endpoints?, model?}`. Stable keys: `engine`, `services` map.
- `list-engines`: `table` (default) or `json`. JSON shape: `{"active-engine": string, "engines": [..]}`. Table layout is for human consumption; JSON is for automation.
- `show-engine`: `yaml` (default) or `json`. Exposes engine metadata.
- `show-machine`: `yaml` (default) or `json`. Hardware information schema from `pkg/types`.
- `version`: `yaml` (default) or `json`. Fields: `snap`, `cli`.

Stability expectations:
- Human-readable `table` and YAML outputs may change form but should remain parseable for humans; JSON outputs are machine contracts and should be considered stable across minor releases. Any breaking change to JSON must follow deprecation rules in `deprecation.md`.

Parseability guidance:
- Commands that accept `--format=json` emit a single JSON document or an indented JSON structure (e.g., `list-engines` prints a JSON object). Consumers should parse full JSON output, not rely on whitespace.
