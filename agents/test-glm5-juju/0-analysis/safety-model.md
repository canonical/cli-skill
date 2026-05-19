# Juju CLI Safety Model

## Overview

The Juju CLI implements multiple safety mechanisms to prevent accidental data loss and destructive operations. This document describes confirmation prompts, force flags, dry-run support, and recovery behavior for dangerous operations.

## Safety Mechanisms Overview

| Mechanism | Purpose | Applicable Commands |
|-----------|---------|---------------------|
| Confirmation prompts | Prevent accidental destruction | `destroy-*`, `remove-*`, `kill-*` |
| Force flags | Override safety checks | Most destructive commands |
| Dry-run mode | Preview without execution | `deploy`, `remove-application` |
| Storage protection | Prevent storage data loss | `destroy-model`, `destroy-controller` |
| Block commands | Prevent accidental destruction | Controller-wide protection |

## Confirmation Prompts

### Default Behavior

Destructive operations require explicit confirmation:

```bash
$ juju destroy-model my-model
WARNING! This command will destroy the "my-model" model.
This includes all machines, applications, and data.

Continue? [y/N]
```

### Disabling Prompts

Prompts can be disabled with `--no-prompt` or by setting the `mode` model configuration:

```bash
# Per-command override
juju destroy-model my-model --no-prompt

# Model configuration
juju model-config mode=requires-prompts  # Always prompt
juju model-config mode=""                # Default behavior
```

### Confirmation for Specific Resources

```bash
$ juju destroy-controller myctl
WARNING! This command will destroy the "myctl" controller.
This includes:
 - 3 models will be destroyed
 - 5 machines will be destroyed
 - 2 storage instances will be destroyed

Continue? [y/N]
```

## Force Flags

### Commands with Force Support

| Command | Force Behavior | Warning |
|---------|---------------|---------|
| `bootstrap` | Skip supported base checks | May result in unstable system |
| `deploy` | Skip base/LXD profile checks | Charm may not function correctly |
| `remove-application` | Remove despite hook failures | Units may not clean up |
| `remove-unit` | Remove despite hook failures | Unit may not clean up |
| `remove-machine` | Remove despite running units | Applications may be incomplete |
| `remove-relation` | Remove despite errors | May leave inconsistent state |
| `remove-offer` | Remove despite active consumers | May break cross-model relations |
| `bind` | Bind to unavailable spaces | Units may not connect |
| `destroy-model` | Destroy despite errors | May leave orphaned resources |
| `destroy-controller` | Destroy despite errors | May leave orphaned resources |

### Force Behavior Details

#### Application Removal

```bash
# Normal removal (fails if hook errors)
juju remove-application mysql

# Force removal (ignores hook errors)
juju remove-application mysql --force

# Force with no-wait (skip cleanup steps)
juju remove-application mysql --force --no-wait
```

**Warning message:**
```
Using --force will remove all units without giving them the opportunity
to shutdown cleanly. This may result in data loss or inconsistent state.
```

#### Model Destruction

```bash
# Normal destruction (fails if errors)
juju destroy-model my-model

# Force destruction (ignores all errors)
juju destroy-model my-model --force

# Force with timeout
juju destroy-model my-model --force --model-timeout 30s
```

#### Controller Destruction

```bash
# Normal destruction (requires all models destroyed first)
juju destroy-controller myctl --destroy-all-models

# Force destruction (destroys everything, ignores errors)
juju destroy-controller myctl --destroy-all-models --force
```

**Warning message:**
```
WARNING: Passing --force with --model-timeout will continue the final
destruction without consideration or respect for clean shutdown or
resource cleanup. If --model-timeout elapses with --force, you may have
resources left behind that will require manual cleanup.
```

### Force + No-Wait Combination

The `--force --no-wait` combination is the most destructive:

```bash
juju remove-application mysql --force --no-wait
```

This:
1. Skips all validation checks
2. Skips cleanup steps
3. Proceeds without waiting for any operations to complete
4. May leave orphaned resources

## Dry-Run Mode

### Supported Commands

| Command | Dry-Run Output |
|---------|---------------|
| `deploy` | Shows what would be deployed |
| `remove-application` | Shows what would be removed |

### Deploy Dry-Run

```bash
$ juju deploy mysql --dry-run
Would deploy:
  Application: mysql
  Charm: mysql (channel: stable, revision: 42)
  Base: ubuntu@22.04
  Units: 1
  Constraints: (none specified)
  
No changes made (--dry-run specified)
```

### Remove-Application Dry-Run

```bash
$ juju remove-application mysql --dry-run
Would remove:
  Application: mysql
  Units: mysql/0, mysql/1
  Relations: mysql:db → wordpress:db
  
No changes made (--dry-run specified)
```

## Storage Protection

### Storage in Model Destruction

When destroying a model with persistent storage:

```bash
$ juju destroy-model my-model --destroy-all-models
This model has persistent storage that would be destroyed:
  - mysql-data/0 (100GB, attached to mysql/0)
  
You must specify one of:
  --destroy-storage    Destroy all storage instances
  --release-storage    Release storage from Juju management (keeps data)
  
ERROR model has persistent storage, must specify --destroy-storage or --release-storage
```

### Storage Options

| Option | Behavior | Data Retention |
|--------|----------|----------------|
| `--destroy-storage` | Destroy storage instances | Data is deleted |
| `--release-storage` | Release from Juju management | Data is preserved |

### Storage Attachment Warnings

```bash
$ juju remove-unit mysql/0
ERROR cannot remove unit mysql/0: has attached storage mysql-data/0
Use --destroy-storage to destroy storage or detach first with 'juju detach-storage'
```

## Block Commands

### Overview

Block commands prevent accidental destruction at the controller level:

```bash
# Prevent destroy operations
juju disable-command destroy-model

# Prevent all destructive operations
juju disable-command destroy-controller

# List blocked commands
juju disabled-commands
```

### Block Types

| Block Type | Prevents |
|------------|----------|
| `destroy-model` | `destroy-model` command |
| `destroy-controller` | `destroy-controller`, `kill-controller` |
| `remove-object` | `remove-application`, `remove-unit`, `remove-machine`, etc. |
| `all` | All destructive operations |

### Enabling Blocked Commands

```bash
# Enable specific operation
juju enable-command destroy-model

# Enable all operations
juju enable-command all
```

### Blocked Command Behavior

```bash
$ juju destroy-model my-model
ERROR cannot destroy model: "destroy-model" is disabled
To enable, run: juju enable-command destroy-model
```

## Kill Controller

### Difference from Destroy

| Aspect | `destroy-controller` | `kill-controller` |
|--------|---------------------|-------------------|
| Graceful | Yes (with --force) | No |
| Cleanup | Yes | No |
| Recovery | Possible | Not possible |
| Use case | Normal decommission | Emergency removal |

### Kill Controller Behavior

```bash
$ juju kill-controller myctl
WARNING! This command will forcibly kill the controller.
This is a destructive operation that may leave orphaned resources.

The controller will be killed without any cleanup.
No hosted models will be destroyed.
Cloud instances may remain running.

Continue? [y/N]
```

### Kill Controller + No-Prompt

```bash
juju kill-controller myctl --no-prompt
```

This immediately kills without confirmation.

## Recovery Behavior

### Failed Bootstrap

```bash
$ juju bootstrap aws myctl
ERROR bootstrap failed: timeout waiting for controller

# Options:
# 1. Check logs
juju debug-log --replay

# 2. Retry with increased timeout
juju bootstrap aws myctl --config bootstrap-timeout=1800

# 3. Keep broken instance for debugging
juju bootstrap aws myctl --keep-broken
```

### Failed Upgrade

```bash
$ juju upgrade-model
ERROR upgrade failed: agent binary not found

# Recovery:
# 1. Check available versions
juju upgrade-model --agent-version

# 2. Force specific version
juju upgrade-model --agent-version=3.4.0
```

### Failed Migration

```bash
$ juju migrate my-model other-controller
ERROR model migration failed

# Recovery:
# 1. Check model health
juju status

# 2. Force migration
juju migrate my-model other-controller --force
```

## Safety Checklist for Destructive Operations

### Before Destroying a Model

1. Check for applications you want to keep:
   ```bash
   juju status
   ```

2. Export bundle for backup:
   ```bash
   juju export-bundle > model-backup.yaml
   ```

3. Handle storage:
   ```bash
   juju storage
   # Decide: destroy or release
   ```

4. Confirm destruction:
   ```bash
   juju destroy-model my-model
   ```

### Before Destroying a Controller

1. List all models:
   ```bash
   juju models
   ```

2. Export each model:
   ```bash
   for model in $(juju models --format yaml | yq '.models[].name'); do
     juju export-bundle -m $model > ${model}-backup.yaml
   done
   ```

3. Check storage:
   ```bash
   juju storage --model <each-model>
   ```

4. Destroy with appropriate flags:
   ```bash
   juju destroy-controller myctl --destroy-all-models --destroy-storage
   ```

### Before Force-Removing

1. Check for dependent resources:
   ```bash
   juju status --relations
   juju storage
   ```

2. Consider graceful removal first:
   ```bash
   # Try normal removal
   juju remove-application mysql
   
   # If fails, check error
   juju resolved mysql/0  # Resolve any hook errors
   juju remove-application mysql
   ```

3. Use force only if necessary:
   ```bash
   juju remove-application mysql --force
   ```

## Safety Model Summary Table

| Command | Prompt | Force | Dry-Run | Storage Check |
|---------|--------|-------|---------|---------------|
| `bootstrap` | No | Yes | No | No |
| `deploy` | No | Yes | Yes | No |
| `add-model` | No | No | No | No |
| `destroy-model` | Yes | Yes | No | Yes |
| `destroy-controller` | Yes | Yes | No | Yes |
| `kill-controller` | Yes | No | No | No |
| `remove-application` | No | Yes | Yes | Yes |
| `remove-unit` | No | Yes | No | Yes |
| `remove-machine` | No | Yes | No | No |
| `remove-relation` | No | Yes | No | No |
| `remove-offer` | No | Yes | No | No |
| `migrate` | No | Yes | No | No |
| `upgrade-model` | No | No | No | No |
| `upgrade-controller` | No | No | No | No |

## Best Practices

### For Scripting

```bash
#!/bin/bash
# Always check for success
if ! juju destroy-model my-model --no-prompt --destroy-storage; then
    echo "Model destruction failed"
    # Attempt force destruction
    juju destroy-model my-model --no-prompt --force
fi
```

### For Interactive Use

1. Always read confirmation prompts completely
2. Use `--dry-run` when available
3. Export backups before destruction
4. Use `--force` only as a last resort
5. Check for storage before model/controller destruction

### For CI/CD

```yaml
# Example: Safe model destruction in CI
script: |
  # Export backup
  juju export-bundle > backup.yaml
  
  # Destroy with explicit storage handling
  juju destroy-model ${MODEL} --no-prompt --destroy-storage || \
  juju destroy-model ${MODEL} --no-prompt --force
  
  # Verify destruction
  if juju show-model ${MODEL} 2>/dev/null; then
    echo "Model still exists"
    exit 1
  fi
```
