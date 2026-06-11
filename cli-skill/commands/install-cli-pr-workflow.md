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
# CLI metadata

## Review workflow

- workflow_file: .github/workflows/cli-review.yml
- reusable_workflow: ./.github/workflows/cli-skill-review-reusable.yml
- trigger: pull_request (opened, synchronize, reopened, ready_for_review)
```

If the file already exists, append this section only when missing; do not remove unrelated existing sections.

## Required Output

Write or update:

- `knowledge_base/CLI.md`
- `.github/workflows/cli-review.yml`

## Constraints

- The workflow file content must match `cli-skill/references/cli-skill-review-pr.yml` exactly.
- Do not modify `cli-skill/references/cli-skill-review-pr.yml`.
- Do not add extra jobs or triggers to `.github/workflows/cli-review.yml`.
