# Internal Module: cli-discovery-preflight

Run this module before any analysis command.

## Purpose

- Identify whether the current repository contains a CLI - the skill and simple scripts are NOT considered CLIs.
- If there is no CLI present, stop and return "No CLI found." This is an expected, valid, useful output for the user triggering this skill.
- Determine which files should be parsed for command discovery and behavior analysis
- Produce interim discovery artifacts for downstream commands

## Fixed Output Directory

Write all preflight outputs to:

`cli-review/0-cli-discovery-preflight/`

## Initial Pre-flight analysis: CLI Detection

### Heuristics for CLI Detection

Use multiple signals. Positive matches increase confidence:

- Executable names and wrappers in repo scripts
- Command-registration patterns in source files
- Help text strings (`--help`, `Usage:`, `Commands:`)
- Parser library imports
- Shell completion definitions
- README usage blocks and examples
- Packaging metadata referencing command binaries

## Phase 0: Structure Discovery
**Goal: Map the CLI surface area — architecture, commands, and arguments.**

Files to produce (in order):

1. architecture.md

* Short summary of the tech stack.
* Architecture style used by the CLI. Include one primary style and optional secondary style.
* Typical styles to classify against:
  - Client-server CLI
  - Monolith CLI
  - Library-interface CLI
  - Layered CLI application
  - Plugin-based architecture
  - Microkernel command host
  - Event-driven pipeline
  - Hexagonal (ports/adapters)
  - Command bus architecture

2. commandset.md

* Full list of CLI commands and hierarchy (top-level and subcommands).
* For each command include:
  - Name
  - Short description of what it does (based on docs/help)
  - Description of how it works (based on code path and key functions/modules)

3. argument-structure.md

* Complete map of all commands and all possible arguments.
* Include argument metadata when available: required/optional, default, type, accepted values, aliases, env var mapping.
* Start with an introduction that highlights common argument patterns.
* Add a dedicated section titled Special arguments describing structural exceptions and non-standard patterns.

**Phase 0 checkpoint**: Verify that cli-review/0-cli-discovery-preflight/architecture.md, cli-review/0-cli-discovery-preflight/commandset.md, and cli-review/0-cli-discovery-preflight/argument-structure.md all exist and are non-empty. Do not proceed until all three files are written.
