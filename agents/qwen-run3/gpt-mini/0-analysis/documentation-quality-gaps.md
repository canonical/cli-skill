# Documentation Quality Gaps

Observed gaps between code and documentation:

- README.md documents `qwen36 use-engine --auto`, `qwen36 chat`, `qwen36 show-engine`, and `qwen36 get/set` examples — these match implemented commands. Good coverage for basic workflows.

- Missing examples and details:
  - `prune-cache` lacks example usage and `--engine` behavior in README.
  - `run` is hidden and undocumented in README; since it can run arbitrary subprocesses, this is a notable omission.
  - `serve-webui` is hidden and undocumented; no instructions for static file layout or capabilities list in README.

- Help text inconsistencies:
  - Many commands default to `yaml` and support `json`; README does not document `--format` usage consistently for `status`, `show-engine`, `show-machine`, `version`.
  - `run` marks `--wait-for-components` as deprecated in code; README has no deprecation notes.

- Machine-readable output contract documentation:
  - JSON shapes (e.g., `list-engines` JSON object keys) are not documented in README — consumers have no stable contract reference.

- Completion scripts and shell integration:
  - `apps/completion.bash` is present and included in the snap, but README does not explain enabling shell completion.

Recommendations:
- Add a `CLI` section that lists all commands and the JSON schemas for programmatic outputs.
- Document `prune-cache` and `run` (or explicitly mark them as intentionally hidden advanced commands) and provide safety guidance.
- Add examples for `--format=json` usage and link to machine-readable contracts.
