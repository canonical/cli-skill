# Command: /cli-check-help

Analyze the CLI help system only, using canonical help rules.

## Execution Order

1. Check for the existence of `cli-review/0-cli-discovery-preflight/`. If it does not exist, run `shared/cli-discovery-preflight.md`
2. Use preflight outputs from `cli-review/0-cli-discovery-preflight/`
3. Read `cli-skill/references/cli-help.md` in full
4. Read `cli-skill/references/cli-help-examples.md` in full
5. Evaluate help against `cli-help.md` only

## Scope

- Top-level help page compliance
- Help for every specific command (complete command set coverage)
- Topic help compliance when topical help is available
- Visual hierarchy and scanability compliance
- Entry point compliance (`help`, `--help`, `-h`, `help <command>`, `<command> --help`)

## Required Output

Write:

- `cli-review/4-cli-check-help/help-audit.md`

Required sections:

1. Coverage map of help surfaces
2. Rule compliance matrix (from `cli-help.md`)
3. Non-compliant findings (all)
4. Diffs for proposed corrections
5. Suggested improved help output

## Coverage requirements

- The analysis MUST include:
	- one top-level help page
	- all command-specific help pages discoverable in the CLI command set
	- topical help pages when the CLI exposes help topics
- Do not provide representative sampling only. Coverage must be exhaustive.

## Non-compliance reporting requirements

- List all non-compliant aspects found under rules in `cli-help.md`.
- For each finding, include:
	- violated rule (quote the relevant MUST/SHOULD/CAN item)
	- evidence (exact snippet from actual help output)
	- severity (`High`, `Medium`, `Low`)
	- affected scope (`top-level`, `command:<name>`, or `topic:<name>`)
	- concrete remediation

## Diff requirements

- Provide a diff whenever a textual fix is appropriate.
- Use fenced `diff` blocks.
- Keep diffs minimal and targeted to the violated rule.
- If no direct diff is possible (for example missing grouping semantics), explain the structural change needed.

## Suggested help output requirements

- End the report with a proposed help output style.
- Base this proposal on patterns in `cli-help-examples.md`.
- Use `<pre>` blocks and HTML tags in the same style as the examples.
- Include at least:
	- one improved top-level help mock
	- one improved command help mock
- Proposed output must preserve CLI-specific command names and vocabulary.

## Out of scope

- General CLI design review not tied to help rules
- Semantic command taxonomy analysis
- Heuristic UX analysis outside `cli-help.md`
