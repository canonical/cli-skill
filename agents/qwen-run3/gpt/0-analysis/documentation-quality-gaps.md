# qwen36 Documentation Quality Gaps

## Summary

The codebase is better documented by source than by user-facing docs. The biggest gap is not missing prose in one command help page; it is disagreement between the README, the reusable CLI source, and the shipped qwen36 snap manifest about what the public command surface actually is.

## High-Severity Gaps

| Gap | Evidence | Impact |
|---|---|---|
| `chat` and `webui` are documented as public commands but are not enabled by the shipped qwen36 snap manifest. | README shows `qwen36 chat`; source conditionally registers both commands; root `snap/snapcraft.yaml` does not set `ADDITIONAL_FEATURES`. | users can follow docs exactly and still not see the commands in `--help` |
| `set` help text is too short for a stateful command. | Help says only `Set a configuration`. | users cannot discover multi-key syntax, restart behavior, or hidden config layers |
| `use-engine` help omits the three operating modes. | Public help lists flags but does not explain `--auto` vs `--fix` vs explicit engine selection. | users must learn by failure or source reading |
| `run` help still teaches a deprecated flag. | Example block still shows `--wait-for-components`. | docs actively train obsolete usage |
| No user-facing config key reference exists. | Code and wrappers use many keys; README shows only a few. | users cannot know which settings are supported or safe |

## Medium-Severity Gaps

| Gap | Evidence | Impact |
|---|---|---|
| `unset` wording overpromises deletion. | Help says the key is removed entirely if no default exists, but it only acts on user config. | users may misunderstand layer behavior |
| `prune-cache` does not explain default behavior. | Short help is just `Remove cached data`. | users may not realize it prunes all inactive-engine components by default |
| `status --wait-for-components` is underexplained. | Flag exists, but long waits are not described clearly. | a status command can appear hung without warning |
| No examples for `unset`, `prune-cache`, or `show-machine`. | command help is sparse or example-free | slower learnability |
| No docs explain `passthrough.environment.*`. | behavior exists in code only | advanced users cannot discover a powerful extension path |

## Standards Violations To Flag

Per DE013 grammar and vocabulary:

- `webui` is noun-led, but user-facing commands should be verbs unless there is a strong exception.
- `show-engine` and `show-machine` are understandable, but DE013 prefers noun shorthand or `*-status` for observation where possible.
- `list-engines` is acceptable, but DE013 would also allow the shorthand `engines` for listing a secondary object type.

These are not all severe enough to justify immediate renames, but they are real standards mismatches and should be called out.

## Missing Examples

Docs should add examples for:

- `qwen36 use-engine --auto`
- `qwen36 use-engine --fix`
- `qwen36 show-engine --format json`
- `qwen36 list-engines --format json`
- `qwen36 unset http.port`
- `qwen36 set passthrough.environment.foo=bar`
- `qwen36 prune-cache --engine cuda`
- `qwen36 debug select-engine < hardware.yaml`

## Missing Failure Guidance

The current docs do not explain common failure cases such as:

- no active engine selected
- missing or not-yet-installed components
- server inactive for `chat`
- port or value misconfiguration after `set`
- why `chat`/`webui` may be absent from help despite being mentioned in the README

## Recommendation Compliance Notes

If the project chooses to normalize names to match DE013, changes should follow the deprecation spec:

- next minor release: add the new name while keeping the old one working
- warn on old usage via stderr
- keep both for at least one release cycle
- remove the old form only in the next major release, returning exit code 2 with a replacement hint

## Most Important Fixes

1. align the shipped qwen36 snap manifest with the documented public feature set, or update the README immediately
2. expand `set`, `unset`, and `use-engine` help text and examples
3. remove deprecated usage from `run --help`
4. publish a config key reference and a short output contract for `show-engine`
