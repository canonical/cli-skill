# Argument Structure

The `qwen36` CLI favors positional arguments for essential operands, generally avoiding flag-heavy designs.

- `use-engine`
  - `<engine>` (positional): The name of the engine to use (e.g. `cpu`, `cuda`).
  - `--auto` (flag): Automatically detect the ideal engine based on hardware.
- `show-engine`
  - No arguments required.
- `get`
  - `<key>` (positional required): The key to retrieve (e.g., `http.port`).
- `set`
  - `<key>=<value>` (positional required): Key-value pair configuration setter.
- `chat`
  - No arguments explicitly documented.

## Special arguments
The `use-engine` command features a mutually exclusive conceptual split: it either takes a positional semantic target (`cpu`, `cuda`) OR a flag (`--auto`) as the standalone modifier. This is slightly irregular compared to standard POSIX argument layouts where flags complement positionals rather than replacing them.