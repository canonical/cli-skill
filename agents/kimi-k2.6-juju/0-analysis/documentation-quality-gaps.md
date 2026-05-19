# Documentation Quality Gaps

## Help vs. Code Mismatches
1. **Alias Visibility**: Aliases are listed in the canonical command help (`Aliases: list-actions`) but `juju help commands` does not visually distinguish aliases from primary commands, causing discoverability friction.
2. **Flag Defaults**: Some defaults shown in help are Go zero-values (`""`, `[]`, `false`) without semantic explanation (e.g., `--constraints (= [])` does not explain that empty means "no constraints").
3. **Mode-gated Commands**: `dump` and `dumpdb` are only registered in `DeveloperMode`; help does not indicate this, and attempting to use them on a normal build yields `unrecognized command` without explaining why.
4. **Deprecated Aliases**: `relate` is an alias for `integrate`, but `relate` is not marked as deprecated in help text; the underlying code path may emit a warning, but the alias help output does not mention it.

## Missing Examples
- `grant-cloud`, `revoke-cloud` lack concrete role examples in help text.
- `model-secret-backend` has no usage examples.
- `consume` does not show a cross-controller URL example.
- `set-credential` lacks an example of referencing a remote credential.

## Outdated Guidance
- `bootstrap` help still mentions `ubuntu@22.04` and `3.6.x` version examples that may not reflect the current release.
- `sync-agent-binary` references an "official agent store" without a URL.
- `upgrade-controller` and `upgrade-model` do not clearly document supported upgrade paths (e.g., skipping minor versions).

## Ambiguity
- `config` vs `model-config` vs `controller-config`: the help summaries are similar ("Get, set, or reset configuration"). New users cannot tell which scope applies without reading the full Details section.
- `remove-application` vs `remove-unit` vs `destroy-model`: no guidance on when to use which for scaling down vs. total removal.
- `trust` command summary ("Sets the trust status of a deployed application to true") is vague about what "trust" means in terms of cloud credential access.

## Inconsistencies
- Some commands use "Get, set, or reset..." while others use "Displays or sets..."; a uniform verb tense would improve scanability.
- `--no-prompt` vs `--force` vs `--yes` (none exist): the confirmation bypass flag is not consistently named (`destroy-controller` uses `--no-prompt`, but there is no `--yes` shorthand).
