# Safety Model

## Destructive Operations
The following commands can cause data loss or infrastructure teardown:
- `destroy-controller` — tears down an entire controller and all its models.
- `kill-controller` — forcible termination of controller resources.
- `destroy-model` — removes all machines, applications, and storage in a model.
- `remove-application`, `remove-unit`, `remove-machine` — deletes running workloads or infrastructure.
- `remove-cloud`, `remove-credential`, `remove-user` — deletes persistent metadata.
- `remove-storage`, `detach-storage` — can lead to data loss if storage is not backed up.
- `remove-relation`, `remove-saas`, `remove-offer` — breaks cross-model integrations.

## Confirmations and Force Flags
- Most destructive commands prompt for confirmation unless:
  - `--no-prompt` is passed (explicit bypass).
  - `--force` is passed (some commands require both).
- `destroy-controller` uses `--destroy-all-models`, `--destroy-storage`, `--release-storage`, `--model-timeout` to control teardown scope.
- `kill-controller` uses `--no-wait` to skip graceful shutdown.

## Dry-run Support
- `deploy --dry-run` simulates deployment without creating resources.
- `diff-bundle` compares desired state with actual state, acting as a read-only safety check before applying changes.
- No global `--dry-run` flag exists; it is implemented per-command in the deployment path.

## Recovery Behavior
- `resolved` marks unit errors resolved and retries failed hooks.
- `retry-provisioning` retries failed machine provisioning.
- `enable-destroy-controller` removes disable-command blocks that may prevent teardown.
- Backups (`create-backup`, `download-backup`) allow point-in-time recovery of controller state.

## Recommendations
- Add `--dry-run` to `destroy-model`, `remove-application`, and `remove-unit` for pre-flight validation.
- Standardize confirmation wording across `destroy-*` and `remove-*` commands.
