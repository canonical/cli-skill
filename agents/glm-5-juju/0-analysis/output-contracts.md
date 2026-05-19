# Juju CLI Output Contracts

## Overview

Juju commands produce output in multiple formats to serve both human operators and programmatic consumers. The output contract defines the structure, stability expectations, and parseability guarantees for command output.

## Output Formats

### Available Formats

| Format | Machine-Readable | Human-Readable | Use Case |
|--------|-----------------|----------------|----------|
| `tabular` | No | Yes | Interactive use (default) |
| `yaml` | Yes | Yes | Scripting, CI/CD |
| `json` | Yes | No | Programmatic parsing |
| `line` | Partial | Yes | Compact one-line per entity |
| `oneline` | Partial | Yes | Alias for `line` |
| `short` | No | Yes | Condensed status |
| `summary` | No | Yes | Aggregated statistics |

### Format Selection

```bash
# Default tabular
juju status

# YAML output
juju status --format yaml

# JSON output
juju status --format json

# Line format
juju status --format line

# Write to file
juju status --format yaml --output status.yaml
```

## Output Contracts by Command Category

### Status Command

#### Tabular Format (Default)

```
Model       Controller  Cloud/Region  Version  SLA          Timestamp
production  aws-prod     aws/us-east-1 3.6.0    unsupported  12:34:56Z

App        Version  Status  Scale  Charm         Channel   Rev  Address  Exposed  Message
postgresql 15.4     active      3  postgresql    14/stable  123  10.0.0.1 no       3 units

Unit          Workload  Agent  Machine  Address    Open ports  Message
postgresql/0  active    idle   0        10.0.1.10  5432/tcp    ready
postgresql/1  active    idle   1        10.0.1.11  5432/tcp    ready
postgresql/2  active    idle   2        10.0.1.12  5432/tcp    ready

Machine  State    Address    InstId      AZ          Message
0        started  10.0.1.10  i-abc123    us-east-1a  running
1        started  10.0.1.11  i-def456    us-east-1b  running
2        started  10.0.1.12  i-ghi789    us-east-1c  running
```

#### JSON Format Structure

```json
{
  "model": {
    "name": "production",
    "type": "iaas",
    "controller": "aws-prod",
    "cloud": "aws",
    "region": "us-east-1",
    "version": "3.6.0",
    "sla": "unsupported"
  },
  "applications": {
    "postgresql": {
      "charm": "postgresql",
      "channel": "14/stable",
      "revision": 123,
      "status": {
        "current": "active",
        "message": "",
        "since": "2024-01-15T12:00:00Z"
      },
      "units": {
        "postgresql/0": {
          "machine": "0",
          "address": "10.0.1.10",
          "status": {
            "current": "active",
            "message": "ready"
          },
          "agent-status": {
            "current": "idle"
          }
        }
      }
    }
  },
  "machines": {
    "0": {
      "instance-id": "i-abc123",
      "address": "10.0.1.10",
      "status": {
        "current": "started"
      }
    }
  }
}
```

#### YAML Format Structure

```yaml
model:
  name: production
  type: iaas
  controller: aws-prod
  cloud: aws
  region: us-east-1
  version: "3.6.0"
  sla: unsupported
applications:
  postgresql:
    charm: postgresql
    channel: 14/stable
    revision: 123
    status:
      current: active
      message: ""
      since: "2024-01-15T12:00:00Z"
    units:
      postgresql/0:
        machine: "0"
        address: 10.0.1.10
        status:
          current: active
          message: ready
        agent-status:
          current: idle
machines:
  "0":
    instance-id: i-abc123
    address: 10.0.1.10
    status:
      current: started
```

### List Commands

#### Controllers/List

**Tabular:**
```
Controller  Model    User   Access   Cloud/Region  Models  Nodes  HA  Type
aws-prod    default  admin  superuser aws/us-east-1      5      3   3   iaas
lxd-local   default  admin  superuser lxd/localhost      2      1   0   iaas
```

**JSON:**
```json
{
  "controllers": [
    {
      "name": "aws-prod",
      "model": "default",
      "user": "admin",
      "access": "superuser",
      "cloud": "aws",
      "region": "us-east-1",
      "models": 5,
      "nodes": 3,
      "ha": 3,
      "type": "iaas"
    }
  ]
}
```

#### Models/List

**Tabular:**
```
Controller  Model       User   Access     Cloud/Region    Type  Status     Machines  Cores
aws-prod    production  admin  superuser  aws/us-east-1   iaas  available         5      8
aws-prod    staging     admin  superuser  aws/us-east-1   iaas  available         3      4
```

#### Users/List

**Tabular:**
```
User     Display name  Access      Date created   Last connection
admin    Administrator superuser   2024-01-01     just now
dev      Developer     admin       2024-01-15     2 hours ago
```

### Show Commands

#### show-application

**JSON:**
```json
{
  "name": "postgresql",
  "charm": "ch:postgresql-123",
  "channel": "14/stable",
  "config": {
    "max-connections": 100,
    "shared-buffers": "128MB"
  },
  "constraints": "cores=2 mem=4G",
  "units": {
    "postgresql/0": {
      "machine": "0",
      "address": "10.0.1.10",
      "workload-status": "active",
      "agent-status": "idle"
    }
  },
  "relations": {
    "db": ["mysql/db-client"]
  },
  "endpoint-bindings": {
    "db": "default"
  }
}
```

#### show-unit

**JSON:**
```json
{
  "name": "postgresql/0",
  "application": "postgresql",
  "machine": "0",
  "address": "10.0.1.10",
  "workload-status": {
    "current": "active",
    "message": "ready"
  },
  "agent-status": {
    "current": "idle"
  },
  "relations": {
    "db": {
      "related-applications": ["mysql"],
      "interface": "pgsql",
      "scope": "global"
    }
  },
  "opened-ports": [
    {"port": 5432, "protocol": "tcp"}
  ]
}
```

### Action Commands

#### run (action output)

**On success:**
```yaml
action-id: "1"
results:
  result: |
    Configuration updated successfully
    max-connections: 200
status: completed
timing:
  enqueue: "2024-01-15T12:00:00Z"
  start: "2024-01-15T12:00:01Z"
  complete: "2024-01-15T12:00:02Z"
```

**On failure:**
```yaml
action-id: "2"
results:
  return-code: 1
  stderr: "Error: connection refused"
status: failed
message: "Action execution failed"
timing:
  enqueue: "2024-01-15T12:01:00Z"
  start: "2024-01-15T12:01:01Z"
  complete: "2024-01-15T12:01:02Z"
```

### Debug Commands

#### debug-log output

Streaming log output (not structured):
```
12:34:56 unit-postgresql-0: postgresql/0: debug: Connection established
12:34:57 unit-postgresql-0: postgresql/0: info: Ready to accept connections
12:34:58 unit-postgresql-1: postgresql/1: debug: Connection established
```

### Deploy Commands

#### deploy output

**Tabular (default):**
```
Located charm "postgresql" in channel "14/stable", revision 123
Deploying charm "postgresql" as "postgresql"
```

**YAML (with --dry-run):**
```yaml
changes:
  - operation: add-application
    charm: ch:postgresql-123
    application: postgresql
    units: 3
  - operation: add-unit
    application: postgresql
    placement: new
  - operation: add-unit
    application: postgresql
    placement: new
  - operation: add-unit
    application: postgresql
    placement: new
summary:
  applications:
    - postgresql
  machines: 3
  units: 3
```

## Output Stability Guarantees

### Stable Fields (JSON/YAML)

These fields are guaranteed to maintain backward compatibility:

| Field | Stability | Notes |
|-------|-----------|-------|
| `model.name` | Stable | Model identifier |
| `model.type` | Stable | `iaas` or `caas` |
| `applications.<name>.status.current` | Stable | Status value |
| `applications.<name>.units.<name>.status.current` | Stable | Status value |
| `machines.<id>.status.current` | Stable | Status value |

### Unstable/Internal Fields

These fields may change between releases:

| Field | Stability | Notes |
|-------|-----------|-------|
| `*-status.since` | Unstable | Timestamp format may change |
| `*-status.message` | Unstable | Human-readable, may change |
| Internal metadata | Unstable | Implementation details |

### Deprecation Policy

When output structure changes:

1. **Additive changes**: New fields added without notice
2. **Field removal**: 2-release deprecation warning period
3. **Breaking changes**: Major version bump (rare)

## Parseability Guidelines

### Recommended Parsing Approach

**For scripts:**
```bash
# Get application status as JSON
status=$(juju status --format json)
app_status=$(echo "$status" | jq -r '.applications.postgresql.status.current')

# Get specific unit address
unit_address=$(echo "$status" | jq -r '.applications.postgresql.units."postgresql/0".address')
```

**For CI/CD pipelines:**
```yaml
# Example GitHub Actions step
- name: Check application status
  run: |
    status=$(juju status --format json)
    if [ "$(echo "$status" | jq -r '.applications.postgresql.status.current')" != "active" ]; then
      echo "Application not active"
      exit 1
    fi
```

### Field Access Patterns

| Information | JSON Path |
|-------------|-----------|
| Model name | `.model.name` |
| Application status | `.applications.<name>.status.current` |
| Unit address | `.applications.<name>.units.<unit>.address` |
| Machine address | `.machines.<id>.address` |
| Charm channel | `.applications.<name>.channel` |

### Filtering by Selector

When using selectors with status:
```bash
# Filter by machine
juju status 0

# Filter by application
juju status postgresql

# Filter by unit
juju status postgresql/0

# Multiple selectors
juju status postgresql mysql
```

## Error Output

### Error Format

Errors are written to stderr:

```
ERROR cannot connect to controller "aws-prod": connection refused
```

**Structured errors (with --debug):**
```
ERROR cannot connect to controller "aws-prod"
    github.com/juju/juju/api/connection.go:123: connection refused
    github.com/juju/juju/cmd/juju/commands/bootstrap.go:456: bootstrap failed
```

### Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Command not found or initialization error |
| N | Passthrough (from plugins) |

## Streaming Output

### Log Streaming

```bash
juju debug-log --format json
```

Each line is a JSON object:
```json
{"time":"12:34:56","entity":"unit-postgresql-0","level":"DEBUG","message":"Connection established"}
{"time":"12:34:57","entity":"unit-postgresql-0","level":"INFO","message":"Ready to accept connections"}
```

### Event Streaming

Some commands produce streaming output for long-running operations:

```bash
juju bootstrap aws production --debug
```

Progress messages are written to stdout/stderr as they occur.

## Quiet Mode

With `--quiet` flag:

- Suppresses informational messages
- Only errors are written to stderr
- Output (if any) is written to stdout

```bash
# Compare:
juju deploy postgresql
# Located charm "postgresql"...
# Deploying charm...

juju deploy postgresql --quiet
# (no output on success)

juju deploy postgresql --quiet --format json
{"application": "postgresql", "units": 1}
```

## Color and Formatting

### Color Control

```bash
# Enable color explicitly
juju status --color

# Disable color
juju status --no-color

# Auto-detection (default when terminal supports it)
juju status
```

### Table Formatting

Tabular output uses Unicode box-drawing characters when terminal supports it, ASCII fallback otherwise.

## Version Output

```bash
juju version
# 3.6.0

juju version --all
# version: 3.6.0
# git-hash: abc123def
# build-date: "2024-01-15"
# go-version: go1.22.0
```
