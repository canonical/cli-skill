# Juju CLI Documentation Quality Gaps

## Methodology

This analysis compares the help output (`juju help <command>`) and generated documentation against the actual command behavior observed in the codebase. Gaps are categorized as mismatches, missing examples, outdated guidance, or ambiguity.

## Findings

### 1. Inconsistent Command Purpose Strings

Many commands lack `Purpose` strings in their `Info()` method, causing sparse help output:

| Command | Has Purpose? |
|---|---|
| `add-cloud` | No |
| `add-credential` | No |
| `add-k8s` | No |
| `add-unit` | No |
| `add-user` | No |
| `bootstrap` | No |
| `config` | No |
| `constraints` | No |
| `consume` | No |
| `controllers` | No |
| `credentials` | No |
| `debug-log` | No |
| `destroy-controller` | No |
| `disable-user` | No |
| `download` | No |
| `expose` | No |
| `find` | No |
| `grant` | No |
| `grant-cloud` | No |
| `remove-application` | No |
| `remove-cloud` | No |
| `remove-credential` | No |
| `remove-relation` | No |
| `remove-saas` | No |
| `remove-unit` | No |
| `resume-relation` | No |
| `suspend-relation` | No |
| `trust` | No |
| `unexpose` | No |
| `update-credential` | No |
| `upgrade-controller` | No |
| `upgrade-model` | No |

**Gap**: ~30 commands (25% of the command set) show only the command name in `juju help commands` with no one-line description.

### 2. Alias Documentation

Aliases are inconsistently documented:
- Some aliases appear in generated docs but not in `juju help commands`
- `relate` → `integrate` alias is not prominently advertised
- `list-*` aliases are hidden from help but documented in generated markdown

**Gap**: Users discovering commands via `juju help` may not find aliases.

### 3. Constraint Syntax Documentation

The constraints syntax (especially `spaces=dmz,^cms`) is documented in long-form help for some commands but not consistently across all commands that accept constraints. New users must discover the negative-match `^` syntax through trial and error or external docs.

**Gap**: No centralized, discoverable constraint syntax reference in CLI help.

### 4. Storage Directive Syntax

Storage directives (`pool,size,count`) are complex but only documented inline in `deploy` and `add-storage` help. The comma-separated syntax and optional fields are not intuitive.

**Gap**: Storage syntax lacks a dedicated help topic or examples.

### 5. Cross-Model Relation URL Format

The offer URL format (`<user>/<model>.<application>` or `controller:user/model.application`) is critical for `consume`, `offer`, and `remove-offer` but is buried in long help text.

**Gap**: No quick-reference for offer URL format in command help.

### 6. Feature Flag Gating

Commands gated by `featureflag.DeveloperMode` (`dump-model`, `dump-db`) are completely invisible unless the flag is set. There is no indication in standard help that these commands exist.

**Gap**: Hidden commands create discoverability issues for developers.

### 7. Model vs Application Config Confusion

`juju config`, `juju model-config`, and `juju controller-config` operate on different scopes but have similar interfaces. The help text does not clearly cross-reference the differences.

**Gap**: Users frequently confuse which config command to use.

### 8. Destroy vs Kill Controller

The distinction between `destroy-controller` and `kill-controller` is critical but subtle. `kill-controller` is a "last resort" that bypasses graceful shutdown, but this is only clear in the long help.

**Gap**: The safety implications are not prominent enough.

### 9. Status Format Stability

`juju status --format json` output is used heavily for scripting, but the help text does not document which fields are stable vs. experimental.

**Gap**: No machine-readable output stability contract in help.

### 10. Block Commands Documentation

The `disable-command` and `enable-command` commands use command set names (`destroy-model`, `remove-object`, `all`) that are not enumerated in help.

**Gap**: Users cannot discover valid block names from the CLI alone.

### 11. Interactive Shell

The `juju` interactive shell is mentioned in the top-level help but has no dedicated documentation topic. Its capabilities (completion, history) are undocumented.

**Gap**: REPL functionality is under-documented.

### 12. Plugin Discovery

The plugin system (`juju-<command>` binaries) is not documented in CLI help at all. Users must discover it from external documentation.

**Gap**: No mention of plugins in `juju help` or `juju help commands`.

## Recommendations

1. **Add Purpose strings to all commands**: Ensure every command has a one-line description.
2. **Create help topics for complex syntax**: constraints, storage directives, offer URLs.
3. **Cross-reference related commands**: `config` family, `destroy`/`kill` family.
4. **Document aliases in help**: Show aliases in `juju help <command>` output.
5. **Add a plugin help topic**: Explain how external plugins work.
6. **Stability annotations**: Mark yaml/json fields as stable or experimental in generated docs.
7. **Expose developer commands in help**: List them with a "(developer mode)" annotation even when gated.
