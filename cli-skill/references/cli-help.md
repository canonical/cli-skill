# Canonical CLI help rules

This document defines normative rules for checking CLI help quality.

## Normative language

- MUST: required for compliance
- SHOULD: strongly recommended; deviations need clear justification
- CAN: optional enhancement

## 1. Help entrypoints

- A CLI MUST provide help via both command and flag forms.
- A CLI MUST accept, at minimum: `tool help`, `tool --help`, `tool -h`, `tool help <command>`, `tool <command> --help`.
- A CLI SHOULD support `tool help --all` for full command overview.

## 2. Top-level help content

- Top-level help MUST print to stdout on successful help requests.
- Top-level help MUST include:
  - Usage
  - Summary (purpose sentence or short paragraph)
  - Global options
  - Commands or command groups
  - Instruction for command-level or topic-level help
- Top-level help SHOULD include version in header or first block.
- Top-level help SHOULD include topic grouping when command count is large.
- Top-level help CAN include setup/config suggestions when environment is uninitialized.

## 3. Command help content

- Command help MUST include command-specific Usage.
- Command help MUST describe command intent and default behavior.
- Command help MUST list command options.
- Command help SHOULD group options by function (for example: target, - Command help SHOULD include `Examples` for non-trivial commands.
- Command help CAN include `Related commands`.

## 4. Topic help content

- If topical help exists, it MUST include:
  - topic summary
  - commands in the topic
  - instruction to access command-specific help
- Topic help SHOULD include brief context explaining the domain.

## 5. Flags and arguments documentation

- Every flag MUST have a one-line description.
- Flags with constrained values MUST list accepted values.
- Flags with defaults MUST show defaults.
- Usage MUST distinguish required from optional arguments.
- If a flag has implications (for example `--devmode` implies weaker validation), help SHOULD state the implication.
- Flag names and vocabulary MUST be consistent with CLI grammar and existing commands.

## 6. Guidance, recovery, and next steps

- Help MUST include a clear path to more detailed help (`help <command>`, `help <topic>`, or equivalent).
- In incomplete or invalid invocation paths, the CLI MUST print concise error feedback to stderr and MUST include corrected usage.
- Help SHOULD provide actionable next steps for common setup failures (for example missing config file).
- Help CAN include documentation URLs.

## 7. Visual hierarchy and scannability

- Help output MUST have stable section ordering across commands.
- Help output MUST use clear section labels (for example `Usage`, `Global options`, `Examples`, `Related commands`).
- Help output MUST preserve alignment/indentation that allows fast column scanning for command and option lists.
- Help output MUST separate major sections with whitespace.
- Help output SHOULD keep line widths readable in standard terminals; wrapped descriptions SHOULD use hanging indentation.
- Help output SHOULD prioritize key actions near the top (usage, primary commands, high-value options).
- If color or text styling is used, it MUST not be required to understand structure; monochrome output MUST remain readable.
- Commands/topics in grouped lists SHOULD be visually distinct from descriptions (alignment, spacing, or consistent delimiter).
- Visual emphasis CAN be used for warnings/suggestions, but SHOULD be sparse and consistent.

## 8. Consistency rules

- Section naming SHOULD be consistent across top-level and command help.
- Terminology SHOULD be consistent across help pages (`command`, `option`, `topic`, `summary`).
- Equivalent option classes SHOULD be documented in the same format across commands.

## 9. Compliance scoring guidance

- A checker MUST report evidence per violated rule.
- A checker MUST treat missing required sections as non-compliant.
- A checker SHOULD classify severity as:
  - High: missing help entrypoints, missing Usage, missing section structure, unreadable hierarchy
  - Medium: missing defaults/accepted values, weak grouping, poor discoverability
  - Low: wording, minor formatting, weak but present hierarchy cues
- A checker CAN compute an aggregate score, but MUST include rule-level findings.
