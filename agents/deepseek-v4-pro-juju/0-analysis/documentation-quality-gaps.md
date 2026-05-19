# Documentation Quality Gaps

## Overview

This analysis compares the CLI help output and code behavior against documentation expectations. Gaps are identified across several dimensions: help text accuracy, missing examples, outdated references, and ambiguity.

## 1. Help Text Gaps

### 1.1 Missing `See Also` Cross-References

Many commands have incomplete `SeeAlso` sections. Examples found via code inspection:

- `bootstrap`: SeeAlso lists are empty (no related commands linked)
- `switch`: No SeeAlso (should reference `controllers`, `models`, `whoami`)
- `debug-log`: No SeeAlso for `status`, `show-status-log`
- `add-unit`: No SeeAlso for `remove-unit`, `deploy`
- `run`: SeeAlso could reference `exec`, `operations`, `show-operation`
- `ssh`: SeeAlso could reference `scp`, `debug-hooks`, `debug-code`
- `status`: SeeAlso is comprehensive (✅)
- `deploy`: SeeAlso is comprehensive (✅)

### 1.2 Missing Examples

Commands lacking examples in their Info() struct:

- `migrate` — No `Examples` field in Info()
- `switch` — No `Examples` field
- `sync-agent-binary` — No examples
- `retry-provisioning` — No examples
- `dump-model` — No examples (but developer-mode gated)
- `dump-db` — No examples (but developer-mode gated)
- `enable-destroy-controller` — Minimal documentation
- `set-credential` — No examples
- `charm-resources` — Internal command, possibly intentionally undocumented

### 1.3 Flag Descriptions with Gaps

| Command | Flag | Issue |
|---------|------|-------|
| `bootstrap` | `--bootstrap-base` | Does not document valid values or format |
| `deploy` | `--base` | Does not list valid base identifiers |
| `status` | `--watch` | Duration format not explained |
| `debug-log` | `--level` | Does not list valid log levels |
| `refresh` | `--switch` | Format of charm URL not explained |
| `add-credential` | `-f`/`--file` | File format not described |
| `model-defaults` | `--cloud` | Default behavior not documented |
| `bind` | `--bind` | Format not described in flag help, only in doc string |
| `expose` | `--to-spaces`, `--to-cidrs` | Format not described in flag help |

### 1.4 Missing Flag Descriptions

Commands where flags exist but have no help description:

- `destroy-model`: `--t` has no description (only `--timeout` has one)
- `kill-controller`: `--t` has no description (see above)
- Multiple commands: Short flag aliases inherit help from the long form, but the list of short aliases is not documented in `--help` output.

## 2. Outdated Guidance

### 2.1 Deprecated `--series` Flag

Many commands still accept `--series` alongside `--base`. The flags are functionally equivalent but `--series` is the legacy name. Help text in some places still references `--series` as primary.

- Recommendation: Standardize on `--base` and add deprecation notice for `--series`.

### 2.2 `integrate` vs `relate` Alias

The `integrate` command has `relate` as an alias. Both appear in documentation. The `relate` alias reflects historical naming and may confuse users who see both in search results or older documentation.

### 2.3 `--integrations` vs `--relations` (status)

Both `--integrations` and `--relations` exist as identical flags for `status`. The `--integrations` flag description says "Same as `--relations`". This creates confusion and violates DE013 guidance against providing both.

### 2.4 `-c` Flag Ambiguity

On commands inheriting from `ControllerCommandBase`, `-c` maps to `--controller`. On `bootstrap`, `-c` is not used for anything. This inconsistency can confuse users.

## 3. Examples Quality

### 3.1 Good Examples (Comprehensive)

`deploy`, `status`, `add-credential`, `bootstrap`, `config`, `destroy-model`, `remove-application`, `remove-unit`, `run` — These commands include multiple examples covering common and edge cases with explanatory comments.

### 3.2 Minimal Examples

`add-cloud`, `remove-cloud`, `machines`, `spaces`, `subnets`, `whoami`, `controllers`, `models` — These have single-line examples only. They cover the happy path but do not show common failure scenarios.

### 3.3 No Examples

`migrate`, `switch`, `sync-agent-binary`, `retry-provisioning`, `enable-destroy-controller`, `set-credential` — These completely lack examples. Users must infer usage from argument descriptions or doc strings.

## 4. Code ↔ Help Mismatches

### 4.1 Missing Commands from Help Output

No major mismatches found. All registered commands appear in `juju help commands`. However, the help listing is alphabetical and grouped by package only in source, not in output. Users see a flat, undifferentiated list of ~130 commands with no grouping.

### 4.2 `--help` vs Markdown Docs

The Markdown generation (`cmd.PrintMarkdown`) can produce documentation from command Info() structs. However, there's no automated verification that the generated Markdown matches the live help. Potential drift between static docs and live help.

### 4.3 Feature-Gated Commands

`dump-model` and `dump-db` appear in `juju help commands` only when the `DeveloperMode` feature flag is enabled. Users without the flag enabled will not see these commands, which is acceptable but could cause confusion when documentation references them.

## 5. Ambiguity & Confusion

### 5.1 Scope Confusion: `config` vs `model-config` vs `controller-config`

Three config commands exist at different scopes. A new user might:
- Run `config` expecting to set model-level config
- Be confused about which applies where
- Not discover `controller-config` or `model-defaults`

The SeeAlso cross-references exist but the naming pattern is inconsistent: `config` (application), `model-config` (model), `controller-config` (controller), `model-defaults` (defaults).

### 5.2 `add-unit` vs `deploy -n`

Both add units to an application. `deploy -n 5` deploys with 5 units initially; `add-unit -n 3` adds 3 units to an existing deployment. The boundary and best practice guidance is not clearly explained.

### 5.3 `run` vs `exec`

`run` runs hook commands (on the unit agent's context); `exec` runs arbitrary commands inside the workload. The distinction is documented but subtle for new users. They share similar flags but different execution contexts.

### 5.4 `destroy-controller` vs `kill-controller`

- `destroy-controller` requires the controller to be reachable
- `kill-controller` works even when unreachable
- `destroy-controller --force` overlaps behaviorally with `kill-controller`
- Users may not know which to use when a controller is partially reachable

### 5.5 `migrate` Position

`migrate` is registered under `commands/` but operates on models. It's neither a purely controller operation nor a purely model operation. Its location in the help listing doesn't signal its dual nature.

## 6. Documentation Quality Summary

| Dimension | Rating | Notes |
|-----------|--------|-------|
| Help text completeness | Medium | Core commands well-documented; edge commands lack detail |
| Examples coverage | Medium | ~60% commands have examples; critical commands covered |
| Flag documentation | Medium | Most flags documented, but some descriptions are terse |
| SeeAlso cross-references | Low | Many commands don't link to related commands |
| Deprecation notices | Low | No active deprecation warnings despite legacy flags |
| Command grouping | Low | Flat list of 130 commands with no categorization |
| Machine-readable output docs | Medium | Formats documented but fields not explicitly listed |
| Error message docs | Low | No error catalog, no per-command error documentation |

## 7. Recommendations

1. **Add SeeAlso links** across all commands to improve discoverability
2. **Add examples** to all commands currently missing them
3. **Improve flag descriptions** with format specifications and valid values
4. **Standardize `--t` shorthand** description across destroy commands
5. **Add grouping** to `juju help commands` output (by domain: Cloud, Model, Application, etc.)
6. **Document common error scenarios** in per-command help
7. **Deprecate `--series`** formally with deprecation warning
8. **Rename `--integrations`** to a single flag (keep `--relations` or vice versa)
9. **Document `--dry-run` availability** in help for destructive commands
10. **Add migration commands** SeeAlso to both controller and model command groups
