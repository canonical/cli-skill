# Command: /install-cli-pr-workflow

Install or refresh a pull-request workflow that runs the CLI review reusable workflow.

## Execution Order

1. Ensure `knowledge_base/` exists.
2. If `knowledge_base/CLI.md` exists, update it with minimal required metadata. Keep existing content where possible.
3. If `knowledge_base/CLI.md` does not exist, create a minimal version.
4. Check whether `.github/workflows/cli-review.yml` already exists. If it exists, do not create or overwrite it.
5. If `.github/workflows/cli-review.yml` does not exist, read `cli-skill/references/cli-skill-review-pr.yml`.
6. If `.github/workflows/cli-review.yml` does not exist, copy `cli-skill/references/cli-skill-review-pr.yml` to `.github/workflows/cli-review.yml` literally (exact content, no edits).
7. Ensure `.github/workflows/` exists before writing.

## Required Metadata in `knowledge_base/CLI.md`

The file must contain, at minimum:

```markdown
# CLI scope

This repository <describe the role of the CLI here>.
paths: [<paths where CLI source code resides>]
```