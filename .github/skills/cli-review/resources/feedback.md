# Feedback Interview Protocol

This document defines the structured interview process used by the `give-feedback` command. It captures UX feedback from the agent's perspective on how well the skill instructions worked in practice.

## Prerequisites

- At least one prior workflow must have been completed (e.g., `analyze-cli`, `discuss-commandset`)
- The agent must have its own output files available for reference

## Interview Protocol

Answer each section honestly and in detail, reflecting on the experience of following this skill's instructions as an AI agent.

### Opening Question

How did you feel about the experience of following this skill? What was your overall impression?

### Step 1: Agent Profile

- What was your prior knowledge level regarding CLI review before reading this skill?
- Had you encountered similar structured analysis frameworks before?
- In what context did you process these instructions (token limits, tool availability, source code access, etc.)?

### Step 2: Target Tasks

- What was the core goal of the skill as you understood it?
- What was the "happy path" — the exact correct sequence of steps?
- Did you follow that sequence, or did you deviate? If so, where and why?

### Step 3: Four Core UX Questions

For each major section of the skill that you executed — answer all four:

- **Intent**: Did you understand what the skill wanted you to achieve at this step?
- **Visibility**: Was the correct action obvious from reading the instructions?
- **Matching**: Did the skill's language match the actual work you needed to do?
- **Feedback**: After completing each step, did you know whether you had succeeded?

Apply these questions to whichever sections you executed (e.g., 0-analysis, 1-discuss-commandset, Output Completeness, Recommendation Compliance, HTML generation, Scale Awareness, etc.)

### Step 4: Record and Analyze Problems

For each problem encountered:

| Problem | Failed Question | Severity | Suggested Fix |
|---------|----------------|----------|---------------|
| Description | Intent / Visibility / Matching / Feedback | low / medium / high | How to rewrite the instruction |

Severity definitions:
- **high**: Caused incorrect output, missing files, wrong filenames, or compliance failure
- **medium**: Caused confusion, extra effort, or suboptimal output quality
- **low**: Minor friction that did not affect output correctness

### Step 5: Additional Observations

- Were there any instructions that were redundant or contradictory?
- Were there sections where you had to make judgment calls not covered by the skill?
- What single change would most improve this skill for an AI agent following it?

## GitHub Issue Filing

After completing the interview, file a GitHub issue for **every problem rated medium or high severity**.

**Repository**: `canonical/cli-skill`

**Issue format**:

- **Title**: `[feedback] <short problem description>`
- **Labels**: `feedback`, and either `severity/high` or `severity/medium`
- **Body**:
  ```
  ## Problem
  <description of what went wrong>

  ## Failed UX Question
  <Intent | Visibility | Matching | Feedback>

  ## Severity
  <high | medium> — <one-sentence justification>

  ## Context
  - CLI analyzed: <name>
  - Skill section: <which part of the skill>
  - Agent model: <model name if known>

  ## Suggested Fix
  <concrete rewrite or structural change to the skill>

  ## Evidence
  <reference to the specific output file or behavior that demonstrates the problem>
  ```

File each issue separately — do not batch multiple problems into one issue.
