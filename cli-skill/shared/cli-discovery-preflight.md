# Internal Module: cli-discovery-preflight

Run this module before any analysis command.

## Purpose

- Identify whether the current repository contains a CLI
- Determine which files should be parsed for command discovery and behavior analysis
- Produce interim discovery artifacts for downstream commands

## Fixed Output Directory

Write all preflight outputs to:

`cli-review/0-cli-discovery-preflight/`

This directory may contain multiple files.

## Required Output Files

Create these files at minimum:

1. `cli-review/0-cli-discovery-preflight/00-summary.md`
- Result: `yes`, `no`, or `uncertain`
- Confidence: `high`, `medium`, or `low`
- Short rationale

2. `cli-review/0-cli-discovery-preflight/01-candidate-entrypoints.md`
- Candidate CLI entrypoints and registration files
- Language/framework hints (Go Cobra, Python argparse/Typer/Click, Rust clap, Node Commander, etc.)

3. `cli-review/0-cli-discovery-preflight/02-parse-targets.md`
- Prioritized list of files to parse in subsequent analysis
- Why each file matters (registration, help text, flags, output, errors)

4. `cli-review/0-cli-discovery-preflight/03-gaps-and-assumptions.md`
- Missing context and blocked areas
- Explicit assumptions for downstream analysis

## Heuristics for CLI Detection

Use multiple signals. Positive matches increase confidence:

- Executable names and wrappers in repo scripts
- Command-registration patterns in source files
- Help text strings (`--help`, `Usage:`, `Commands:`)
- Parser library imports
- Shell completion definitions
- README usage blocks and examples
- Packaging metadata referencing command binaries

## Failure Policy

- If no CLI is found, do not write any files, and return the result "No CLI present. No CLI review necessary".
