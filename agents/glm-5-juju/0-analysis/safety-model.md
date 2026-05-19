# Juju CLI Safety Model

## Overview

Juju manages critical cloud infrastructure including compute instances, storage volumes, network configurations, and application data. The CLI implements multiple safety mechanisms to prevent accidental data loss and infrastructure destruction.

## Destructive Operations

### High-Risk Commands

These commands can cause irreversible changes:

| Command | Risk Level | Impact |
|---------|------------|--------|
| `destroy-controller` | Critical | Terminates all models, machines, applications |
| `kill-controller` | Critical | Force destroys without cleanup |
| `destroy-model` | High | Terminates all model resources |
| `remove-application` | High | Deletes application and unit data |
| `remove-machine` | High | Terminates compute instance |
| `remove-unit` | Medium | Removes unit and its data |
| `remove-storage` | Medium | Deletes storage volumes |
| `remove-relation` | Medium | Breaks application connections |

### Moderate-Risk Commands

| Command | Risk Level | Impact |
|---------|------------|--------|
| `config` | Medium | Changes application configuration |
| `model-config` | Medium | Changes model behavior |
| `integrate` | Medium | Modifies application topology |
| `upgrade-model` | Medium | Upgrades agent versions |
| `refresh` | Medium | Upgrades charm version |

## Confirmation Mechanisms

### Interactive Prompts

For destructive operations, Juji prompts for confirmation:

```bash
juju destroy-model production
# WARNING: This command will destroy the model "production"
# including all applications, machines, and data.
# 
# Continue? [y/N] 
```

**Implementation:**
```go
// cmd/modelcmd/confirmation.go
func confirmationString() string {
    return fmt.Sprintf("This command will destroy the %q model including all applications, machines, and data. Continue? [y/N] ", modelName)
}

func (c *DestroyCommand) Run(ctx *cmd.Context) error {
    if !c.noPrompt {
        if !c.confirm(ctx) {
            return nil
        }
    }
    // ... proceed with destruction
}
```

### Skip Confirmation

Use `--no-prompt` or `-y` to bypass:

```bash
juju destroy-model production --no-prompt
juju destroy-model production -y
```

**Warning:** Automation scripts should only use this when consequences are fully understood.

## Force Flags

### Force Semantics

The `--force` flag bypasses safety checks:

| Check | Normal | With --force |
|-------|--------|--------------|
| Model not empty | Blocked | Proceed |
| Units in error | Blocked | Proceed |
| Relations exist | Blocked | Proceed |
| Storage attached | Blocked | Proceed |
| Hooks running | Blocked | Proceed |

**Example:**
```bash
# Normal: fails if relations exist
juju remove-application postgresql
ERROR cannot remove application: has active relations

# Force: removes despite relations
juju remove-application postgresql --force
WARNING: forcing removal of "postgresql" with 2 active relations
```

### Force Cascade Behavior

For model destruction:

```bash
juju destroy-model production --force
```

This will:
1. Remove all applications (with --force)
2. Remove all machines (with --force)
3. Destroy storage
4. Delete model

### Force Risks

Using `--force` can cause:

- **Orphaned resources** - Cloud instances may remain running
- **Data loss** - Storage destroyed without backup
- **Relation corruption** - Broken cross-model relations
- **State inconsistency** - Controller may not reflect actual state

## Dry-Run Mode

### Preview Operations

Several commands support `--dry-run`:

```bash
juju deploy postgresql --dry-run
Changes to be applied:
  - add application postgresql
  - add unit postgresql/0 to new machine
  - add machine in us-east-1a

Summary:
  applications: 1
  machines: 1
  units: 1

No changes were applied.
```

### Supported Commands

| Command | Dry-Run Support |
|---------|----------------|
| `deploy` | Yes |
| `bundle` (deploy) | Yes |
| `integrate` | No |
| `remove-*` | No |
| `upgrade-*` | No |

## Safe Defaults

### Read-Only by Default

Informational commands are safe:

```bash
juju status        # Safe: read-only
juju models        # Safe: read-only
juju config app    # Safe: read-only
juju ssh --dry-run # Preview only
```

### Non-Destructive Defaults

Commands default to safe behavior:

| Command | Default | Destructive Alternative |
|---------|---------|------------------------|
| `remove-storage` | Detach only | `--destroy-storage` |
| `remove-application` | Prompt | `--no-prompt` |
| `destroy-model` | Prompt | `--no-prompt` |

## Block Protection

### Command Blocking

Users can block destructive operations:

```bash
# Block all destructive operations
juju enable-command destroy-model
juju enable-command remove-object

# Check blocked commands
juju disabled-commands
```

### Blocked Commands

| Block Type | Commands Affected |
|------------|------------------|
| `destroy-model` | `destroy-model`, `destroy-controller` |
| `remove-object` | `remove-application`, `remove-machine`, `remove-unit`, `remove-relation` |
| `all-commands` | All destructive operations |

**Example:**
```bash
juju disable-command destroy-model "Production model - contact ops team"

juju destroy-model production
ERROR command is disabled: Production model - contact ops team
```

## HA and Leadership Safety

### Leader Unit Protection

Leader units have special protection:

```bash
juju remove-unit mysql/0
ERROR cannot remove unit "mysql/0": unit is the leader
Hint: remove all units or use --force
```

### HA Controller Safety

Multiple controllers require special handling:

```bash
juju destroy-controller aws-prod
ERROR controller has 3 enabled machines (HA)
Hint: use --force or reduce HA before destroying
```

## Storage Safety

### Detach vs Destroy

```bash
# Safe: detaches storage, data preserved
juju remove-storage pg-data/0

# Destructive: destroys storage volume
juju remove-storage pg-data/0 --destroy-storage
```

### Storage Confirmation

```bash
juju remove-storage pg-data/0 --destroy-storage
WARNING: This will permanently destroy storage "pg-data/0" 
including all data on the volume.

Continue? [y/N]
```

## Operation Timeouts

### Safety Timeouts

Operations have default timeouts:

| Operation | Default Timeout | Override |
|-----------|-----------------|----------|
| Bootstrap | 20 minutes | `--config bootstrap-timeout` |
| SSH | 30 seconds | `--timeout` |
| Action run | No timeout | `--wait <duration>` |

### Timeout Behavior

On timeout:
- Operation may still be in progress
- Check status with `juju status`
- May require manual intervention

## Recovery Mechanisms

### Undo Operations

Some operations can be reversed:

| Operation | Reversal |
|-----------|----------|
| `config key=value` | `config --reset key` |
| `integrate` | `remove-relation` |
| `add-unit` | `remove-unit` |
| `deploy` | `remove-application` |

### Irreversible Operations

These cannot be undone:

| Operation | Impact |
|-----------|--------|
| `destroy-model` | All data lost |
| `destroy-controller` | All models lost |
| `remove-storage --destroy` | Volume deleted |
| `kill-controller` | No cleanup |

### Backup Recovery

Before destructive operations:

```bash
# Create backup
juju create-backup

# Download for safekeeping
juju download-backup <backup-id>
```

## Safe Scripting Patterns

### Pre-Flight Checks

```bash
#!/bin/bash
set -euo pipefail

# Verify model exists
model="production"
if ! juju show-model "$model" &>/dev/null; then
    echo "Model $model not found"
    exit 1
fi

# Check for errors
errors=$(juju status --format json | jq '.applications | to_entries[] | select(.value.status.current == "error") | .key' | wc -l)
if [ "$errors" -gt 0 ]; then
    echo "Model has $errors applications in error state"
    exit 1
fi

# Proceed with operation
juju deploy postgresql
```

### Idempotent Operations

```bash
#!/bin/bash
set -euo pipefail

app="postgresql"

# Check if application exists
if juju status "$app" --format json 2>/dev/null | jq -e ".applications.$app" > /dev/null; then
    echo "Application $app already exists"
    exit 0
fi

juju deploy "$app"
```

### Safe Cleanup

```bash
#!/bin/bash
# Safe script with confirmation

echo "This script will:"
echo "  - Remove all applications"
echo "  - Destroy all machines"
echo "  - Destroy the model"
echo ""
read -p "Continue? [y/N] " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    exit 1
fi

# Proceed with cleanup
juju destroy-model --no-prompt
```

## Best Practices

### For Operators

1. **Always review** destructive command output before confirming
2. **Use dry-run** when available to preview changes
3. **Create backups** before major operations
4. **Test in staging** before production changes
5. **Document** the purpose of blocked commands

### For Automation

1. **Prefer explicit flags** over defaults (`--no-prompt` clearly shows intent)
2. **Check status first** - verify model state before operations
3. **Implement rollback** logic for failures
4. **Log all operations** for audit trails
5. **Use least-privilege** accounts for automated tasks

### For Development

1. **Default to safe** - destructive operations should require explicit flags
2. **Provide clear warnings** - explain consequences before action
3. **Support dry-run** for preview capability
4. **Enable blocking** for production protection
