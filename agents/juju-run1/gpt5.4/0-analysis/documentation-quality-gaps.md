# Juju Documentation Quality Gaps

## Evidence base

This analysis compares:
- command metadata and generated markdown from `juju documentation --split`
- command help output from `juju help <command>`
- repo documentation under `docs/`

The goal here is not to assess prose quality alone. It is to identify coverage gaps, stale metadata, and places where the documented command contract diverges from the implemented one.

## High-signal gaps in generated command docs

### 1. Usage signatures are missing for zero-arg commands in generated markdown

The markdown generator only emits a Usage section when `Info.Args` is non-empty. As a result, commands such as:
- `clouds`
- `controllers`
- `dashboard`
- `whoami`
- `secret-backends`
- `disabled-commands`

lose an explicit usage block in generated docs, even though `juju help <command>` shows one.

Impact:
- generated reference docs are less complete than runtime help
- simple commands appear under-documented relative to more complex ones

### 2. Duplicated `[options]` in several help signatures

Verified examples:
- `find`: `juju find [options] [options] <query>`
- `download`: `juju download [options] [options] <charm>`
- `info`: `juju info [options] [options] <charm>`

Cause:
- those commands embed `[options]` in `Info.Args`, and the framework adds its own options marker

Impact:
- generated docs look unpolished
- suggests weak validation around usage metadata

### 3. Flags leaked into positional usage strings

Verified examples:
- `move-to-space` includes `[--format yaml|json]` in `Args`
- `spaces` usage effectively bakes format/output details into the signature
- `subnets` does the same

Impact:
- options table and usage string duplicate responsibility
- makes the usage contract less machine-normalizable

### 4. `help` is poorly represented as a normal command in generated inventories

`help` is a framework feature with special paths, and its generated usage metadata is not as cleanly inventoryable as normal commands. In metadata sweeps it can surface the long usage/help banner rather than a compact usage signature.

Impact:
- inventory generation from command metadata is slightly messy
- reference generation around help topics is less uniform than around normal commands

## Command-doc content quality gaps

### 5. Not all commands expose equal depth of examples and details

Large, high-impact commands such as `deploy`, `destroy-model`, and `destroy-controller` have strong detail sections.
Many smaller commands are much thinner, with minimal examples or sparse explanation.

Impact:
- users get uneven guidance depending on which family they are in
- advanced commands are often better documented than medium-complexity commands

### 6. Output-contract guidance is distributed, not standardized

Commands that support `--format` often explain it locally, but there is no centralized statement about:
- machine-readable output guarantees
- schema stability
- when tabular output is only for humans
- how stderr behaves under JSON/YAML modes

Impact:
- scripting guidance is fragmented

### 7. Safety semantics are documented mostly in help text, not as a shared policy

Destroy commands do a good job, but there is no broad documentation pattern for:
- what `--force` means across command families
- which commands support preview/dry-run
- when confirmation is expected
- how command blocking interacts with force flags

Impact:
- users learn safety behavior one command at a time

## Repo-doc coverage observations

### 8. The docs tree is task-oriented, not a complete command reference by itself

The repo documentation under `docs/` is strong on tutorials, how-to flows, and conceptual reference. It references commands throughout the task guides, but it is not a one-to-one replacement for command reference generated from the CLI.

That is acceptable as a docs strategy, but it means the generated command docs need to be especially clean. Right now they are useful, but not polished enough to fully carry that load.

### 9. Cross-linking to command reference exists, but contract details still live in code-generated help

The how-to guides frequently point readers to refs such as `command-juju-deploy`, `command-juju-config`, and similar command pages. That is good for discoverability, but it also means problems in generated metadata propagate into the reference experience.

### 10. Documentation coverage follows major workflows better than edge-case commands

The docs emphasize:
- deploy
- configure
- expose/unexpose
- storage lifecycle
- destroy flows
- status

Smaller or more operationally niche commands get less narrative support.
Examples include:
- `disabled-commands`
- `move-to-space`
- `enable-destroy-controller`
- some secret-backend operations
- some action/task inspection commands

## Specific verified inconsistencies worth fixing

- `find`, `download`, and `info` usage signatures contain duplicated `[options]`.
- `move-to-space`, `spaces`, and `subnets` embed flags in positional usage metadata.
- zero-arg commands lose usage sections in generated markdown.
- `documentation` itself exposes a usage signature that reads more like an implementation syntax block than a user-friendly command synopsis.

## Why these gaps matter

These are not only cosmetic issues.
They affect:
- inventory generation
- command discoverability
- scriptability expectations
- operator trust in help output
- the ability to evolve the command surface cleanly

When the CLI is as broad and flat as Juju's, high-quality command metadata is part of the product, not just documentation polish.

## Assessment

Strengths:
- generated command docs exist and are broadly useful
- major workflows are well represented in task docs
- high-risk commands often have strong examples and warnings
- docs and code are connected closely enough that drift is detectable

Weaknesses:
- generated metadata has visible structural defects
- reference completeness varies by command complexity
- output and safety semantics are not documented as cross-cutting policies
- the docs system depends on command metadata being clean, but there are too many handwritten inconsistencies in `Info.Args`

## Priority fixes

1. Stop encoding `[options]` and format hints inside `Info.Args`.
2. Make the markdown generator always emit a Usage section, even for zero-arg commands.
3. Establish a style guide for `Info.Args` and examples so new commands do not repeat the same issues.
4. Add centralized docs for output guarantees and safety semantics, then keep per-command docs focused on local specifics.
