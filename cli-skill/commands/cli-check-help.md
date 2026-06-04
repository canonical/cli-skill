# Command: /cli-check-help

Analyze the CLI help system only.

## Execution Order

1. Run `shared/cli-discovery-preflight.md`
2. Use parse targets relevant to help generation and command registration
3. Evaluate help output quality and consistency

## Scope

- Top-level help command quality
- Subcommand help coverage and consistency
- Usage lines and grammar quality
- Flag documentation completeness (required/default/type/values)
- Example quality (happy path + failure path)
- Help command discoverability (`--help`, `help`, aliases)

## Required Output

Write:

- `cli-review/4-cli-check-help/help-audit.md`

Recommended sections:

1. Coverage map of help surfaces
2. Gaps and ambiguities
3. Inconsistencies and formatting issues
4. Scriptability concerns for help parsing
5. Concrete fixes and test suggestions
