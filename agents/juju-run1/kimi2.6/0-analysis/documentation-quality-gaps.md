# Juju CLI Documentation Quality Gaps

## Scope

Comparison between generated command docs under `docs/reference/juju-cli/list-of-juju-cli-commands/` and framework/code behavior in `cmd/`.

## Findings (ordered by severity)

### High

1. Incomplete usage lines for multiple commands
- Several command pages show `N/A` for usage in extracted inventory.
- This weakens scriptability and argument discoverability for those commands.

2. Inconsistent option representation style
- Some usage lines duplicate markers (`[options] [options]`) or include unusual punctuation/spacing.
- Inconsistency can confuse copy/paste usage and automation wrappers.

### Medium

3. Global-vs-command option separation is not always explicit
- Framework provides global behavior (`help`, description, formatter plumbing), but individual pages may not clearly distinguish global options from command-local options.

4. Plugin behavior is underrepresented in command reference
- Runtime plugin fallback (`juju-<name>`) is a critical extension behavior but is typically absent from command reference pages.

5. Error and exit-code contracts are not consistently documented per command
- Framework behavior is clear in code, but user-facing docs do not systematically specify expected non-zero exit behavior.

### Low

6. Safety semantics are distributed
- Confirmation/no-prompt and disable-command controls are documented in individual commands, but not consolidated into a unified safety reference section for operators.

## Mismatch examples

- Command docs are generated from command metadata, but additional runtime behavior (plugin fallback, env conflict detection, serializable error handling) lives in framework code and is easy to miss from per-command pages.

## Recommended improvements

1. Add CI checks to fail docs generation when usage block is missing.
2. Normalize usage formatter to avoid duplicated `[options]` tokens.
3. Add a shared "Global behavior" page covering:
- global flags and formatting
- plugin fallback semantics
- environment precedence and conflict rules
- exit code model
4. Add per-command "Exit codes" and "Safety" mini-sections for destructive commands.

## Evidence pointers

- Command doc generation: `cmd/cmd/documentation.go`, `scripts/md-gen/juju-cli-commands/main.go`
- Existing command docs: `docs/reference/juju-cli/list-of-juju-cli-commands/*.md`
- Framework runtime behavior: `cmd/cmd/supercommand.go`, `cmd/cmd/cmd.go`, `cmd/juju/commands/plugin.go`, `cmd/modelcmd/controller.go`
