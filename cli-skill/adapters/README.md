# Cross-Agent Adapter Scaffolding

This directory contains thin adapter scaffolds for:

- Copilot
- Claude Code
- Pi Coding Agent
- OpenCode

All adapters point to the same canonical command implementation under `cli-skill/`.

## Canonical Source

- Manifest: `cli-skill/schemas/commands.manifest.yaml`
- Shared preflight: `cli-skill/shared/cli-discovery-preflight.md`
- Active commands: `cli-skill/commands/`
- Future command stubs: `cli-skill/future-commands/`

## Fixed Preflight Output

The preflight module writes to:

- `cli-review/0-cli-discovery-preflight/`

## Notes

These are scaffolds for portability. If a specific runtime needs strict schema fields, adapt the corresponding file under `adapters/` while keeping command targets unchanged.

## Regeneration

When command mappings change, regenerate all adapter files from the canonical manifest:

```bash
node /project/scripts/sync-cli-skill-adapters.js
```

Source of truth:

- `cli-skill/schemas/commands.manifest.yaml`

## Git Hook Enforcement

A pre-commit hook is available at:

- `.githooks/pre-commit`

The hook:

1. Runs `node /project/scripts/sync-cli-skill-adapters.js`
2. Fails the commit if adapter files changed during sync
3. Prompts you to stage regenerated adapter files before retrying commit

To enable hooks in this repository:

```bash
git config core.hooksPath .githooks
chmod +x .githooks/pre-commit
```
