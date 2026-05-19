# Documentation Quality Gaps

## Sources Analyzed

1. **README.md** (project root): End-user quick-start guide.
2. **cli/README.md**: Developer-focused CLI repo README.
3. **Code comments and help text** (`Short`, `Long`, `Example` fields in Cobra command definitions).
4. **CLI reference documentation** (linked from README at `https://documentation.ubuntu.com/inference-snaps/` — not directly available in repo).

## Summary

The documentation is **developer-oriented** and assumes familiarity with the snap ecosystem. End-user documentation is thin, with the primary reference being external (Ubuntu documentation site). The in-CLI help text is functional but lacks examples, edge-case guidance, and machine-readable output documentation on several commands.

---

## Gap 1: README Mentions Commands Not in Code / Missing Commands

### README shows `show-engine` and `use-engine`
✅ Present in code.

### README shows `get http.port` and `set http.port=8326`
✅ Present in code.

### README shows `qwen36 chat`
✅ Present (if `ADDITIONAL_FEATURES` includes `chat`).

### README does NOT mention:
- `status` — the primary "what's happening" command
- `list-engines` — the primary engine discovery command
- `show-machine` — mentioned in `cli/README.md` but not in root `README.md`
- `prune-cache` — disk cleanup command
- `version`
- `webui` — launch web UI
- `unset` — config removal
- Any `debug` subcommands

**Impact**: New users reading the root README will miss 7 out of 12 visible commands (58% of the visible command set undocumented at the top level).

---

## Gap 2: Conditional Commands Not Documented

The `chat` and `webui` commands are conditionally compiled based on the `ADDITIONAL_FEATURES` environment variable. Neither the README nor the CLI help text explains:
- That these commands may be absent.
- What environment variable controls them.
- How a user can check if they're available.

**Impact**: A user on a snap that doesn't enable chat will see no `chat` command with no explanation of why.

---

## Gap 3: No Examples in CLI Help Text

Per DE013, commands should provide examples. Current state:

| Command | Has `Example` Field? |
|---------|---------------------|
| `run` | ✅ Yes: `cli run env`, `cli run -- echo "Hello World!"`, `cli run --wait-for-components -- python3 -m http.server` |
| `status` | ❌ |
| `chat` | ❌ |
| `webui` | ❌ |
| `get` | ❌ |
| `set` | ❌ |
| `unset` | ❌ |
| `list-engines` | ❌ |
| `show-engine` | ❌ |
| `use-engine` | ❌ |
| `show-machine` | ❌ |
| `prune-cache` | ❌ |
| `version` | ❌ |
| `serve-webui` | ❌ |
| `debug validate-engines` | ❌ |
| `debug select-engine` | ❌ |
| `debug chat` | ❌ |
| `debug serve-webui` | ❌ |

Only **1 out of 18** commands has an `Example` field in its Cobra definition. The `run` command (hidden) has examples; the visible commands that users actually see do not.

**Impact**: Users must guess at syntax, especially for `set` (which uses `key=value` pairs) and `use-engine` (which has `--auto`/`--fix` modes).

---

## Gap 4: `--format` Flag Documentation Inconsistency

Commands that support `--format`:
- `status`: Has flag, documented. Default: `yaml`.
- `show-machine`: Has flag, documented. Default: `yaml`.
- `show-engine`: Has flag, documented. Default: `yaml`.
- `version`: Has flag, documented. Default: `yaml`.
- `list-engines`: Has flag, documented. Default: `table`.
- `debug select-engine`: Has flag, documented. Default: `yaml`.

**Issue**: The `--format` flag help text shows accepted values but does not explain what each format looks like. Users must trial-and-error to understand the difference between `yaml` and `json` output on a per-command basis.

**Issue**: `list-engines` defaults to `table` while all other commands default to `yaml`. This inconsistency is not explained in help text.

---

## Gap 5: Hidden Flags Not Documented

| Flag | Command | Visibility |
|------|---------|------------|
| `--package` | `set` | Hidden |
| `--engine` | `set` | Hidden |

These flags are intentionally hidden (used internally for engine switching). However, the existence of per-tier configuration is a user-facing concept (users can `get` merged values). The inability to set/view per-tier configs is a deliberate limitation that is not communicated.

---

## Gap 6: Error Message Guidance

The CLI provides helpful suggestions in error conditions:
- Key not found → suggests `qwen36 get` to view available keys.
- Server inactive → suggests `sudo snap start <service>` to start.
- Permission denied → suggests `sudo`.

But these suggestions are not documented in help text. Users must encounter errors to discover them. A `--help` section on "Common Issues" or "Troubleshooting" would improve discoverability.

---

## Gap 7: `prune-cache` Behavior Undocumented

The README does not mention `prune-cache` at all. The help text (`Short: "Remove cached data"`) is terse. Important behavioral details that should be documented:
- What "cached data" means (inactive engine components).
- That the active engine's components are protected.
- The confirmation prompt defaults to "no".
- The `--engine` flag for pruning a specific engine.
- That it requires root.
- That sizes are shown when available.

---

## Gap 8: `use-engine --fix` Semantics

The `--fix` flag on `use-engine` has nuanced behavior:
- If no active engine, exits with success (nothing to fix).
- If active engine manifest is gone, auto-selects a new engine.
- If active engine exists, reinstalls missing components and re-applies configs.

None of this is documented in help text (`Short: "Select an engine"` — the `--fix` flag is not mentioned in `Short`/`Long`). Users must use `--help` to discover `--fix`, and even then its behavior is underspecified.

---

## Gap 9: Configuration Model Not Explained

The three-tier config model (package → engine → user) is not explained in any user-facing documentation within the repo:
- The README shows `get`/`set` examples but doesn't explain precedence.
- `qwen36 set --help` doesn't explain the tiers.
- There's no documentation on what keys exist (e.g., `http.port`, `webui.http.port`, `passthrough.environment.*`).

The only way to discover available keys is to run `qwen36 get` and inspect the output.

---

## Gap 10: `status` Output Fields Undocumented

The `status` command's help text (`Short: "Show the status"`) doesn't list what fields are displayed. Users must run the command to discover it shows `engine`, `services`, `endpoints`, and `model` information. The `--wait-for-components` flag is also underexplained.

---

## Gap 11: External Documentation Link Rot Risk

The `cli/README.md` links to `https://documentation.ubuntu.com/inference-snaps/` and `https://documentation.ubuntu.com/inference-snaps/reference/models-cli/` for CLI reference. If these URLs change or the external docs become outdated, the repo has no self-contained CLI reference.

---

## Gap 12: No Changelog or Versioning Documentation

There is no CHANGELOG.md, no release notes, and no documentation of CLI stability guarantees. The snap `version` is `"3.6"` but there's no indication of what changed between versions. The `grade: devel` in `snapcraft.yaml` signals that the snap is not production-ready, but this is not communicated to CLI users.

---

## Gap 13: Shell Completion Not Documented

The CLI includes a bash completion script (`bin/completion.bash`), configured via `completer: bin/completion.bash` in `snapcraft.yaml`. Shell completion is not mentioned in any README or help text. Users must discover it by chance or by reading the snapcraft.yaml.

---

## Gap 14: Tabular Output Conventions

The `list-engines` table follows many DE013 table conventions (no ASCII decorations, bold headers, left-aligned, compact mode). But:
- Column headers are **lowercase** (`engine`, `vendor`, `description`, `compat`) — DE013 says they should be UPPERCASE.
- There's no `--no-headers` flag for machine processing of table output.
- The compact borderless style is not documented as intentional.

---

## Recommendations Priority

| Priority | Gap | Suggested Fix |
|----------|-----|---------------|
| **Critical** | Missing README commands (Gap 1) | Add all commands to README quick-start with brief examples |
| **Critical** | No examples in help (Gap 3) | Add `Example` field to every command's Cobra definition |
| **High** | Config model undocumented (Gap 9) | Add CONFIGURATION.md or extend README with config section |
| **High** | Conditional commands unexplained (Gap 2) | Document `ADDITIONAL_FEATURES` in README |
| **High** | `prune-cache` undocumented (Gap 7) | Add to README with examples and safety notes |
| **Medium** | Hidden flags unexplained (Gap 5) | Document config tier concept even if hidden flags stay hidden |
| **Medium** | `--format` inconsistency (Gap 4) | Normalize defaults or document the difference |
| **Medium** | `status` output fields (Gap 10) | Document output fields in help text |
| **Low** | Table header casing (Gap 14) | Uppercase table headers per DE013 |
| **Low** | Shell completion (Gap 13) | Document in README |
| **Low** | External docs link (Gap 11) | Consider a CLI_REFERENCE.md in-repo |
