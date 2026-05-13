---
name: cli-review
description: "Review and discuss command-line interfaces (CLI), including command behavior, UX, flags, output clarity, errors, and docs. Use when the user asks to analyze-cli, review a command, critique CLI UX, or improve command design."
---

# CLI Review Skill

This skill helps review and discuss CLI commands from a user and developer perspective.
It is designed for iterative work on one command at a time.

## Primary Command: analyze-cli

Use this workflow when the user says `analyze-cli` or asks to analyze a CLI command.
The standard output of this flow is a directory named `0-analysis` with multiple markdown documents.

### Required Inputs

Collect as many of these as available:

- Command name and purpose
- Current syntax and examples
- Flags/options and defaults
- Actual outputs (success and error)
- Exit codes
- Target users and primary workflows
- Shell/platform constraints

If context is missing, ask focused follow-up questions before final analysis.

### Toolchain Bootstrap (Ubuntu, snap-first)

Before running analysis or command-design workflows, ensure required toolchains are present when execution depends on them.

1. Detect platform:
- If `/etc/os-release` indicates Ubuntu (`ID=ubuntu` or `ID_LIKE=ubuntu`), continue with this flow.
- If not Ubuntu, skip this section and use platform-appropriate package management.

2. Identify required tools:
- Infer from the requested work and repo context.
- Typical examples: `go`, `java`, `uv`.

3. Check whether each tool exists:
- Use `command -v <tool>`.
- Only install tools that are missing.

4. Install missing tools with Ubuntu preference order:
- Prefer Snap first.
- If Snap installation fails or Snap is unavailable, fallback to apt.
- If apt does not provide the tool, fallback to the official installer.

Recommended commands:

```bash
# Go
command -v go >/dev/null 2>&1 || {
	sudo snap install go --classic ||
	(sudo apt-get update && sudo apt-get install -y golang-go)
}

# Java
command -v java >/dev/null 2>&1 || {
	sudo snap install openjdk ||
	(sudo apt-get update && sudo apt-get install -y default-jdk)
}

# uv
command -v uv >/dev/null 2>&1 || {
	sudo snap install astral-uv --classic ||
	(sudo apt-get update && sudo apt-get install -y uv) ||
	curl -LsSf https://astral.sh/uv/install.sh | sh
}
```

5. Verify installs and continue:
- `go version`
- `java -version`
- `uv --version`

If installation requires elevated privileges or interactive approval, ask concise permission questions and proceed once approved.

### Output Directory And Files

Create `0-analysis` and write these files.

#### Required Core Files

- `architecture.md`
	- Short summary of the tech stack.
	- Architecture style used by the CLI. Include one primary style and optional secondary style.
	- Typical styles to classify against:
		- Client-server CLI
		- Monolith CLI
		- Library-interface CLI
		- Layered CLI application
		- Plugin-based architecture
		- Microkernel command host
		- Event-driven pipeline
		- Hexagonal (ports/adapters)
		- Command bus architecture

- `commandset.md`
	- Full list of CLI commands and hierarchy (top-level and subcommands).
	- For each command include:
		- Name
		- Short description of what it does (based on docs/help)
		- Description of how it works (based on code path and key functions/modules)

- `argument-structure.md`
	- Complete map of all commands and all possible arguments.
	- Include argument metadata when available: required/optional, default, type, accepted values, aliases, env var mapping.
	- Start with an introduction that highlights common argument patterns.
	- Add a dedicated section titled `Special arguments` describing structural exceptions and non-standard patterns.

#### Additional Analysis Files

- `configuration-model.md`
	- Describe config sources and precedence: flags, env vars, config files, defaults.
	- Note command-specific overrides and any surprising precedence behavior.

- `output-contracts.md`
	- Describe output formats by command (human-readable and machine-readable).
	- Document stability expectations for output fields and parseability guidance.

- `error-model-and-exit-codes.md`
	- Map error categories to representative messages and exit codes.
	- Include per-command or per-command-group differences.

- `safety-model.md`
	- Describe destructive operations, confirmations, dry-run support, force flags, and recovery behavior.

- `extensibility-model.md`
	- Explain how new commands or plugins are added.
	- Document registration paths, command discovery, middleware/hooks, and extension boundaries.

- `documentation-quality-gaps.md`
	- Compare docs/help output with code behavior.
	- List mismatches, missing examples, outdated guidance, and ambiguity.

### Build Order

Generate analysis files in this order to maximize reuse:

1. `architecture.md`
2. `commandset.md`
3. `argument-structure.md`
4. `configuration-model.md`
5. `output-contracts.md`
6. `error-model-and-exit-codes.md`
7. `safety-model.md`
8. `extensibility-model.md`
9. `documentation-quality-gaps.md`

### Analysis Checklist

1. Intent and mental model
- Is the command name clear and aligned with user intent?
- Does syntax follow common CLI conventions?

2. Discoverability and help
- Does `--help` explain usage, arguments, examples, and edge cases?
- Are defaults and required inputs explicit?

3. Flags and argument design
- Are short and long flags consistent and predictable?
- Are mutually exclusive and dependent flags handled clearly?
- Are positional arguments minimal and intuitive?

4. Output UX
- Is success output concise and actionable?
- Are errors specific, with next-step guidance?
- Is machine-readable output supported when useful (for example, json)?

5. Safety and reliability
- Are destructive operations gated (confirmation, dry-run, force semantics)?
- Are retries, timeouts, and network failures handled gracefully?
- Are exit codes stable and documented?

6. Consistency and ecosystem fit
- Does behavior match common patterns in similar tools?
- Are naming, formatting, and status messages consistent across commands?

7. Documentation quality
- Are README and examples task-oriented?
- Do examples include both common and failure scenarios?

### Response Format

Return feedback in this structure:

1. Summary
- One paragraph describing command quality and biggest gap.

2. Findings (ordered by severity)
- `Critical`: issues that can cause data loss, unsafe behavior, or workflow failure.
- `High`: major usability or reliability problems.
- `Medium`: clarity, consistency, and learnability issues.
- `Low`: polish and style improvements.

3. Proposed improvements
- Concrete changes to syntax, flags, messaging, and docs.
- Include before/after command examples where possible.

4. Open questions
- Unknowns that block precise recommendations.

5. Suggested tests
- Behavioral checks for parsing, errors, exit codes, and docs examples.

## Collaboration Mode

When discussing tradeoffs, explicitly compare alternatives and call out:

- Backward compatibility impact
- Scriptability impact
- Human readability vs machine readability
- Migration strategy for existing users

## Example Trigger Phrases

- analyze-cli <command>
- review this CLI command
- improve the UX of this command
- critique flags and help text
- propose better error messages for this CLI
- discuss-commandset
- review the command set shape
- analyze command hierarchy
- propose-command <name>
- add a new command
- design a command for <purpose>
- rename-command <old>
- rename this command
- better name for <command>
- frame-analysis
- analyze verb semantics
- compare verb meanings using FrameNet
- semantic difference between verbs

---

## Command Design Phase: 1-command-design

This phase provides structured workflows for discussing and evolving the shape of a CLI's command set.
It produces documents in a `1-command-design/` directory.

### Context Resolution

Before running any command in this phase:

1. **If `0-analysis/` exists**: read `commandset.md` and `argument-structure.md` as primary context. Reference naming patterns, hierarchy, and argument conventions already documented.
2. **If `0-analysis/` does not exist**: ask the user for minimal context:
	- Full command list (top-level and subcommands)
	- Current hierarchy/grouping pattern
	- Naming conventions in use
	- Target user personas

Proceed once you have enough context to reason about the command set structure.

---

### Command: discuss-commandset

Use this workflow when the user says `discuss-commandset` or asks to review the command set shape, hierarchy, or naming patterns.

#### Required Inputs

- The CLI name and command list (from `0-analysis/commandset.md` or user-provided)
- Optionally: ecosystem comparisons (similar tools to benchmark against)

#### Analysis Dimensions

Evaluate the command set across these dimensions:

1. **Grouping and hierarchy**
	- What pattern is used? (noun-verb, verb-noun, flat, nested)
	- Is nesting depth consistent and appropriate?
	- Are related commands grouped logically?

2. **Naming consistency**
	- Do command names follow a single convention (e.g., all verbs, all noun-verb pairs)?
	- Are abbreviations used consistently or avoided?
	- Are there synonyms or near-duplicates that create confusion?

3. **Discoverability**
	- Can a new user predict where to find functionality?
	- Does the hierarchy guide exploration (e.g., `help`, tab completion)?
	- Are there orphan commands that don't fit the mental model?

4. **Ecosystem alignment**
	- How does the structure compare to similar CLI tools?
	- Does it match conventions users already know?

#### Output File

Create `1-command-design/commandset-shape.md` with the following sections in order:

##### Section 1: Verb-Noun Decomposition Matrix

Decompose every command into a verb and a noun (e.g., `add-cloud` → `add` × `cloud`).

Render as a grid:
- Rows = verbs (sorted alphabetically)
- Columns = nouns/resource types (sorted by frequency)
- Cells = `✓` if the command exists, `—` if the combination is absent

After the grid, annotate:
- **Incomplete CRUD sets**: nouns missing expected lifecycle verbs (e.g., has `add-` but no `remove-`)
- **Verb inconsistencies**: nouns using different verbs for equivalent operations (e.g., `destroy-controller` vs `remove-application`)
- **Orphan commands**: commands that do not decompose cleanly into verb-noun (e.g., `bootstrap`, `integrate`, `resolved`, `whoami`)

##### Section 2: Verb Taxonomy and Aspect Classification

Classify every unique verb from the matrix into the following table:

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|

Intent groups:
- **lifecycle**: create, add, deploy, remove, destroy, kill
- **mutation**: update, refresh, upgrade, config, set, bind
- **access**: grant, revoke, enable, disable, login, logout, register, unregister
- **observation**: show, list (plural noun commands), status, log, find, info
- **transfer**: attach, detach, expose, unexpose, consume, offer, integrate
- **execution**: run, exec, cancel, resolved, retry
- **migration**: migrate, import, export, download, scp, sync

Aspect (Aktionsart) values:
- **telic**: action has a natural endpoint (create, destroy, deploy)
- **atelic**: action is ongoing or continuous (run, debug, monitor)
- **punctual**: instant state change (switch, login, trust)

Reversibility values:
- **yes**: paired with a named inverse (add/remove, expose/unexpose)
- **no**: no inverse operation (destroy, kill, bootstrap)
- **partial**: can be undone but not via a single symmetric command (deploy → remove-application)

##### Section 3: Confusion-Pair Audit

List command pairs that share semantic overlap and risk user confusion.

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|

Overlap types:
- **synonym verbs**: different verbs, same operation (e.g., `remove-*` vs `destroy-*`)
- **scope ambiguity**: same verb, unclear which scope applies (e.g., `config` vs `model-config` vs `controller-config`)
- **functional overlap**: different commands that achieve similar outcomes (e.g., `exec` vs `run`)
- **naming collision**: names too similar, different purposes (e.g., `resources` vs `charm-resources`)

For each pair, rate confusion risk as `high`, `medium`, or `low` and provide a one-sentence disambiguation.

##### Section 4: Semantic Domain Clustering

Group all commands by the resource domain they operate on:

| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---|---|---|---|

Domains include: cloud, model, controller, application, unit, machine, user, secret, storage, space/network, relation/integration, credential, offer/SAAS, charm/resource.

For each domain, note:
- Whether all commands use the same noun form
- Whether the CRUD coverage is complete
- Whether the verb choices are consistent within the domain

##### Section 5: Symmetry Audit

For each pair of symmetric operations, list them side by side:

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|

Check:
- `add-*` / `remove-*`
- `enable-*` / `disable-*`
- `expose` / `unexpose`
- `grant-*` / `revoke-*`
- `suspend-relation` / `resume-relation`
- `attach-*` / `detach-*`
- `register` / `unregister`
- `login` / `logout`
- `consume` / `remove-saas`
- `deploy` / `remove-application`
- `destroy-*` (does it have a creation counterpart?)

Flag:
- Missing reverse operations
- Naming asymmetries (e.g., `destroy-controller` is not reversed by `add-controller`)
- Behavioral asymmetries (e.g., reverse operation requires `--force` but forward does not)

##### Section 6: Pattern Classification and Recommendations

- **Pattern classification**: primary grouping pattern, depth, and style
- **Discoverability assessment**: predicted user paths vs actual command locations
- **Ecosystem comparison**: how structure compares to 2-3 similar tools
- **Recommendations**: ordered list of structural improvements with rationale, each annotated with backward compat and migration cost

#### Build Order

Generate sections in this order (each builds on the previous):

1. Verb-Noun Decomposition Matrix
2. Verb Taxonomy and Aspect Classification
3. Semantic Domain Clustering
4. Symmetry Audit
5. Confusion-Pair Audit (uses insights from all above)
6. Pattern Classification and Recommendations

#### Response Format

1. Shape summary — one paragraph describing the current pattern and its strengths
2. Key findings — top 5 issues surfaced by the analysis sections
3. Recommendations — concrete structural changes, ordered by impact
4. Tradeoffs — for each recommendation, note backward compat and migration cost

---

### Command: propose-command

Use this workflow when the user says `propose-command <name>` or asks to design a new command, add a command, or explore where new functionality should live.

#### Required Inputs

- **Purpose**: what the new command should do (user-provided)
- **Context**: existing command set (from `0-analysis/commandset.md` or user-provided)
- Optionally: target users, related existing commands, expected flags

#### Analysis Steps

1. **Naming options** — generate 3-5 candidate names. For each:
	- Name and syntax example
	- Pros (clarity, consistency, discoverability)
	- Cons (ambiguity, collision with existing commands, verbosity)
	- Ecosystem precedent (how similar tools name equivalent functionality)

2. **Hierarchy placement** — determine where the command fits:
	- Under which group/namespace?
	- At what depth?
	- Alongside which sibling commands?
	- Does it require a new group?

3. **Flag design** — propose initial flags:
	- Required vs optional
	- Short and long forms
	- Defaults
	- Consistency with existing flag patterns in the CLI

4. **Impact analysis** — assess effects on the existing command set:
	- Overlap with existing commands (functional duplication)
	- Confusion risk (similar names, adjacent functionality)
	- Commands that may need updating (shared flags, cross-references in help)
	- Whether existing commands should be deprecated in favor of the new one

#### Output File

Create `1-command-design/proposal-<name>.md` (use the proposed command name, slugified) with:

- **Purpose statement**: one sentence describing what the command does
- **Naming options table**: candidates with pros/cons/precedent
- **Recommended name**: chosen name with justification
- **Hierarchy placement**: where it lives and why
- **Proposed syntax**: full usage line with flags
- **Flag specification**: table of flags with type, default, description
- **Impact assessment**: overlap, confusion risks, and required updates
- **Open questions**: decisions that need user input

#### Response Format

1. Recommendation — the proposed name, placement, and syntax
2. Alternatives — other naming/placement options that were considered
3. Impact — what changes in the existing command set
4. Next steps — what needs deciding before implementation

---

### Command: rename-command

Use this workflow when the user says `rename-command <old>` or asks to rename an existing command, find a better name, or restructure command naming.

#### Required Inputs

- **Current command name**: the command to rename (user-provided)
- **Reason for rename**: why the current name is problematic (user-provided or inferred)
- **Context**: existing command set (from `0-analysis/commandset.md` or user-provided)

#### Analysis Steps

1. **Problem diagnosis** — articulate why the current name fails:
	- Misleading (does something different than the name suggests)
	- Inconsistent (breaks the naming pattern of sibling commands)
	- Ambiguous (could mean multiple things)
	- Verbose or too terse

2. **Alternative candidates** — generate 3-5 replacement names. For each:
	- Name and updated syntax
	- How it resolves the diagnosed problem
	- Consistency with the rest of the command set
	- Ecosystem precedent
	- Risk of new confusion

3. **Migration and deprecation strategy**:
	- Alias approach: keep old name as alias, introduce new name
	- Warning period: how long to show deprecation notices
	- Removal timeline: when to drop the old name
	- Documentation updates required
	- Impact on scripts and automation (breaking change assessment)

4. **Ripple effects**:
	- Subcommands that inherit the name
	- Related commands that reference it in help text or docs
	- Flag names or argument names derived from the command name
	- Configuration keys tied to the command name

#### Output File

Create `1-command-design/rename-<old>-to-<new>.md` (using slugified old and recommended new name) with:

- **Problem statement**: why the rename is needed
- **Candidates table**: alternative names with rationale and risk
- **Recommended name**: chosen name with justification
- **Migration plan**: phased approach with timeline
- **Ripple effects**: all downstream impacts
- **Communication plan**: how to inform users (changelog, deprecation warnings, docs)

#### Response Format

1. Recommendation — the new name and one-sentence rationale
2. Migration plan — phased steps from alias to full rename
3. Ripple effects — everything that changes
4. Risk assessment — likelihood of user confusion or breakage during transition

---

### Command: frame-analysis

Use this workflow when the user says `frame-analysis` or asks to analyze verb semantics using FrameNet, compare verb meanings, or understand semantic differences between CLI verbs.

This is a standalone analysis that produces a detailed semantic classification of CLI verbs using Berkeley FrameNet as the primary linguistic resource. It feeds into `discuss-commandset` (confusion-pair audit) and `rename-command` (candidate evaluation) but can be run independently.

#### Required Inputs

- **Verb list**: extracted from the verb-noun decomposition matrix (Section 1 of `commandset-shape.md`) or user-provided
- **FrameNet access**: via NLTK (`nltk.corpus.framenet`) or [Berkeley FrameNet](https://framenet.icsi.berkeley.edu/)

#### Analysis Steps

##### Step 1: Extract CLI Verbs

From the command set, extract all unique verbs. For compound commands, use the leading verb (e.g., `add-cloud` → `add`, `scale-application` → `scale`). List orphan commands that have no clear verb separately.

##### Step 2: FrameNet Frame Lookup

For each verb, query FrameNet for the semantic frame(s) it evokes as a lexical unit.

Record in a table:

| Verb | FrameNet Frame | Frame ID | Frame Definition | Lexical Unit Match |
|---|---|---|---|---|

Lexical unit match values:
- **primary**: the verb's first/default sense matches the CLI usage
- **extended**: the CLI usage is a metaphorical or domain-specific extension of the frame
- **no match**: FrameNet has no frame for this verb in its CLI sense

For verbs with multiple candidate frames, list all candidates and document the disambiguation choice with rationale.

##### Step 3: Frame Element Comparison

For each frame found, list its core frame elements (the semantic roles). Compare role structure between verbs that might be confused:

| Verb | Frame | Agent | Patient/Theme | Source | Goal | Result | Manner |
|---|---|---|---|---|---|---|---|

Use this to explain semantic differences:
- `remove` (Removing): Agent removes Theme *from* Source → implies the thing still exists elsewhere
- `destroy` (Destroying): Destroyer destroys Patient → implies permanent elimination, no Source
- `delete` (Removing): same frame as `remove` → linguistically equivalent, convention difference only
- `kill` (Killing): Killer kills Victim → implies the target was active/alive, strongest irreversibility

These role differences have direct UX implications: if a verb's frame has no Source role, the operation implies the resource cannot be recovered.

##### Step 4: Frame Relation Mapping

For each pair of frames found in Step 2, look up FrameNet's frame-to-frame relations:

| Frame A | Relation | Frame B | Distance |
|---|---|---|---|

Relation types:
- **Inheritance** (Is-a): child frame inherits parent structure
- **Using**: one frame presupposes another
- **Subframe**: one frame is a phase of another
- **Perspective_on**: two frames describe the same situation from different viewpoints
- **Precedes / Is_preceded_by**: temporal ordering

Frame distance = shortest path in the frame relation graph:
- Distance 0: same frame (true synonyms in FrameNet)
- Distance 1: directly related frames
- Distance 2+: indirectly related
- Unrelated: no path exists

Verbs at distance 0 are candidates for consolidation. Verbs at distance 1 may cause confusion. Verbs at distance 2+ are safe to coexist.

##### Step 5: CLI-Specific Frame Annotations

For verbs with no FrameNet match or where the CLI sense diverges from the linguistic sense, create manual frame definitions:

| Verb | Proposed Frame Name | Definition | Core Elements | Closest FrameNet Frame | Gap Notes |
|---|---|---|---|---|---|

Common CLI verbs requiring manual frames:
- `bootstrap`: closest to "Creating" but implies self-initializing infrastructure
- `deploy`: closest to "Placing" but implies provisioning + configuration + activation
- `integrate`: closest to "Attaching" but implies bidirectional relationship creation
- `refresh`: closest to "Cause_change" but implies re-pulling from an external source

##### Step 6: Insights Summary

Synthesize findings into actionable categories:

1. **True synonyms** (distance 0): verbs sharing the same frame that should use a single consistent verb across the CLI
2. **Near synonyms** (distance 1): verbs in related frames where naming should signal the semantic difference
3. **False friends**: verbs that look similar but evoke unrelated frames
4. **Safe coexistence**: verbs in unrelated frames that can coexist without confusion
5. **Frame-informed rename candidates**: verbs whose frame element structure suggests a better name

#### Output File

Create `1-command-design/frame-analysis.md` with:

- **Verb-to-frame mapping table** (Step 2)
- **Frame element comparison** (Step 3) — only for verb groups with potential confusion
- **Frame relation matrix** (Step 4) — only for related frames, not a full grid
- **CLI-specific extensions** (Step 5)
- **Insights and recommendations** (Step 6)

#### Precision Requirements

- Always cite FrameNet frame IDs (e.g., "Removing [Frame #123]")
- Distinguish primary lexical unit match from extended/metaphorical match
- Flag verbs with multiple candidate frames and document disambiguation
- Keep "FrameNet says" observations separate from "CLI domain interpretation"
- When a verb has no FrameNet coverage, explicitly state this rather than guessing

#### Response Format

1. Overview — count of verbs analyzed, frames found, coverage gaps
2. Key findings — verb groups that share frames (synonym risk), verb pairs in related frames (confusion risk)
3. Recommendations — which verbs to consolidate, rename, or keep as-is, grounded in frame semantics
4. CLI-specific gaps — verbs requiring manual frame definitions and their implications
