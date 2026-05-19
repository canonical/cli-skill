# Juju CLI Output Contracts

## Overview

The Juju CLI provides multiple output formats for programmatic and human consumption. This document describes the output contracts for each command, format stability expectations, and parseability guidance.

## Output Format Types

### Available Formats

| Format | Machine-Readable | Use Case |
|--------|-----------------|----------|
| `smart` | Partial | Default, adapts to content type |
| `yaml` | Yes | Configuration, structured data |
| `json` | Yes | API integration, scripts |
| `tabular` | No | Human-readable tables |
| `line` | Partial | Quick unit scans |
| `oneline` | Partial | Compact unit info |
| `summary` | No | Aggregated status |
| `default` | Varies | Command-specific default |

### Format Selection

```bash
juju status --format yaml      # Machine-readable YAML
juju status --format json      # Machine-readable JSON
juju status --format tabular   # Human-readable table
juju status --format line      # One unit per line
juju status --format summary   # Aggregated summary
```

## Output Contract by Command Category

### Status Commands

#### `status`

**Tabular Format:**
```
Model       Controller  Cloud/Region      Version  SLA
my-model    myctl       aws/us-east-1     3.4.0    unsupported

App         Version  Status   Scale  Charm        Channel  Rev  Exposed  Message
postgresql  14.0     active       2  postgresql   stable    42  no       
mysql       8.0      active       1  mysql        stable    38  yes      

Unit           Workload  Agent  Machine  Public address  Ports     Message
postgresql/0*  active    idle   0        10.0.0.1       5432/tcp  
postgresql/1   active    idle   1        10.0.0.2       5432/tcp  
mysql/0*       active    idle   2        10.0.0.3       3306/tcp  

Machine  State    Address    InstId        Base          AZ
0        started  10.0.0.1   i-12345678    ubuntu@22.04   us-east-1a
1        started  10.0.0.2   i-87654321    ubuntu@22.04   us-east-1b
2        started  10.0.0.3   i-11111111    ubuntu@22.04   us-east-1c
```

**JSON Format Fields:**
```json
{
  "model": {
    "name": "my-model",
    "controller": "myctl",
    "cloud": "aws",
    "region": "us-east-1",
    "version": "3.4.0",
    "sla": "unsupported"
  },
  "applications": {
    "postgresql": {
      "charm": "postgresql",
      "channel": "stable",
      "version": "14.0",
      "status": {
        "current": "active",
        "message": ""
      },
      "scale": 2,
      "exposed": false,
      "units": {
        "postgresql/0": {
          "workload-status": {"current": "active"},
          "agent-status": {"current": "idle"},
          "machine": "0",
          "public-address": "10.0.0.1",
          "ports": ["5432/tcp"]
        }
      }
    }
  },
  "machines": {
    "0": {
      "instance-id": "i-12345678",
      "base": "ubuntu@22.04",
      "address": "10.0.0.1",
      "status": {"current": "started"}
    }
  }
}
```

**Stability:**
- JSON/YAML structure is stable across minor versions
- Field additions are backward-compatible
- Field removals require major version bump

#### `show-status-log`

**Tabular Format:**
```
Time                   Type      Status     Message
2024-01-01 00:00:00   juju      pending    deploying
2024-01-01 00:01:00   workload  active    Database ready
```

### Application Commands

#### `deploy`

**Success Output:**
```
Deployed "postgresql" from charmhub-charm/postgresql-42 to my-model on aws/us-east-1 as postgresql
```

**JSON Output:**
```json
{
  "application": "postgresql",
  "charm": "charmhub-charm/postgresql-42",
  "model": "my-model",
  "units": ["postgresql/0"]
}
```

#### `config`

**Tabular Format:**
```
Application  Option    Value     Default  
postgresql   database  mydb      postgresql
             port      5432      5432     
             ssl       true      false    
```

**YAML Format:**
```yaml
postgresql:
  database: mydb
  port: 5432
  ssl: true
```

#### `show-application`

**JSON Format:**
```json
{
  "name": "postgresql",
  "charm": "charmhub-charm/postgresql-42",
  "base": "ubuntu@22.04",
  "status": {
    "current": "active"
  },
  "units": ["postgresql/0", "postgresql/1"],
  "relations": {
    "provides": {
      "db": ["mysql"]
    },
    "requires": {}
  }
}
```

#### `show-unit`

**JSON Format:**
```json
{
  "name": "postgresql/0",
  "application": "postgresql",
  "machine": "0",
  "workload-status": {
    "current": "active",
    "message": ""
  },
  "agent-status": {
    "current": "idle"
  },
  "public-address": "10.0.0.1",
  "ports": ["5432/tcp"],
  "relations": {
    "db": ["mysql/0"]
  }
}
```

### Model Commands

#### `models`

**Tabular Format:**
```
Controller: myctl

Model       Cloud/Region      Status     Machines  Units  Access
admin/my-model  aws/us-east-1  available  2         3      admin
admin/test      aws/us-east-1  available  1         1      admin
```

**JSON Format:**
```json
{
  "controller": "myctl",
  "models": [
    {
      "name": "admin/my-model",
      "cloud": "aws",
      "region": "us-east-1",
      "status": "available",
      "machines": 2,
      "units": 3,
      "access": "admin"
    }
  ]
}
```

#### `model-config`

**Tabular Format:**
```
Attribute         Value        Default  
name              my-model     my-model
logging-config    <root>=WARN  <root>=WARN
image-stream      released     released
```

### Controller Commands

#### `controllers`

**Tabular Format:**
```
Controller  Model    Cloud/Region   Users     Machines  Units  
myctl*      default  aws/us-east-1 admin@myctl 2         3      
```

**JSON Format:**
```json
{
  "controllers": {
    "myctl": {
      "model-uuid": "...",
      "cloud": "aws",
      "region": "us-east-1",
      "current-model": "admin/my-model",
      "users": ["admin@myctl"]
    }
  }
}
```

### Cloud Commands

#### `clouds`

**Tabular Format:**
```
Clouds in this client:
Cloud     Regions  Default   Type        Description
aws       16       us-east-1 ec2         Amazon Web Services
google    22       us-east1  gce         Google Compute Engine
```

**JSON Format:**
```json
{
  "clouds": {
    "aws": {
      "type": "ec2",
      "regions": {"us-east-1": {"endpoint": "..."}},
      "default-region": "us-east-1",
      "description": "Amazon Web Services"
    }
  }
}
```

#### `credentials`

**Tabular Format:**
```
Client credentials for cloud "aws":
Cloud   Credentials
aws     my-credential*
```

### Secret Commands

#### `secrets`

**Tabular Format:**
```
ID                  Owner          Revision  Last Updated
secret-abc123       postgresql/0   1         2024-01-01
secret-def456       mysql/0        3         2024-01-02
```

**JSON Format:**
```json
{
  "secrets": [
    {
      "id": "secret-abc123",
      "owner": "postgresql/0",
      "revision": 1,
      "last-updated": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### Storage Commands

#### `storage`

**Tabular Format:**
```
Storage unit  ID        Status   Persistent  Location
data/0        vol-abc12 attached  true        postgresql/0
```

#### `storage-pools`

**Tabular Format:**
```
Name     Provider  Attrs
ebs      ebs       volume-type=gp2
```

## Output Stability Contracts

### Stability Levels

| Level | Meaning | Commands |
|-------|---------|----------|
| **stable** | Structure frozen, only additive changes | `status`, `models`, `controllers` |
| **beta** | Minor structural changes allowed | Most application commands |
| **alpha** | Structure may change | New commands, experimental features |

### Stable JSON Fields

For commands marked as **stable**:

1. Top-level object structure is frozen
2. Field names do not change
3. Field types do not change
4. New fields may be added (backward-compatible)
5. Deprecated fields marked but preserved

### Backward Compatibility Guarantees

```json
// Version 3.4.0
{
  "model": {"name": "my-model"},
  "applications": {...}
}

// Version 3.5.0 (adds field)
{
  "model": {"name": "my-model", "type": "iaas"},  // New field added
  "applications": {...}
}

// Version 4.0.0 (breaking change)
{
  "model-info": {...},  // Renamed from "model"
  "apps": {...}         // Renamed from "applications"
}
```

## Parseability Guidance

### JSON Parsing

```python
import json
import subprocess

result = subprocess.run(
    ["juju", "status", "--format", "json"],
    capture_output=True,
    text=True
)
status = json.loads(result.stdout)

# Access stable fields
model_name = status["model"]["name"]
apps = status["applications"].keys()
```

### YAML Parsing

```python
import yaml
import subprocess

result = subprocess.run(
    ["juju", "config", "postgresql", "--format", "yaml"],
    capture_output=True,
    text=True
)
config = yaml.safe_load(result.stdout)
port = config["postgresql"]["port"]
```

### Handling Missing Fields

```python
# Defensive field access
def get_unit_status(status, app, unit):
    try:
        return status["applications"][app]["units"][unit]["workload-status"]["current"]
    except KeyError:
        return "unknown"
```

### Handling Deprecated Fields

```python
# Support both old and new field names
def get_applications(status):
    if "applications" in status:
        return status["applications"]
    elif "apps" in status:  # Future version
        return status["apps"]
    else:
        return {}
```

## Output Verbosity Control

### Quiet Mode

```bash
juju deploy mysql --quiet  # Suppresses info output
```

### Verbose Mode

```bash
juju deploy mysql --verbose  # Additional details
```

### Debug Mode

```bash
juju deploy mysql --debug  # Full debug logging
```

## Error Output

### Error Format

Errors are always written to stderr:

```
ERROR cannot deploy "mysql": charm not found
```

### JSON Error Format

For machine-readable error handling, errors include structured data:

```json
{
  "error": "cannot deploy \"mysql\": charm not found",
  "error-code": "charm-not-found",
  "details": {...}
}
```

### Exit Codes (Non-Zero)

| Code | Meaning |
|------|---------|
| 1 | General error |
| 2 | Parse/initialization error |
| N | Plugin exit code passthrough |

## Output Redirection

### File Output

```bash
juju status --output /tmp/status.json --format json
```

### Pipe Integration

```bash
juju status --format json | jq '.applications | keys'
juju config postgresql --format yaml | yq '.postgresql.port'
```

### Stream Processing

```bash
juju debug-log --replay | grep -i error
```

## Tabular Output Customization

### Color Control

```bash
juju status --color     # Force ANSI colors
juju status --no-color  # Disable colors
```

### Timezone Handling

```bash
juju status --utc  # Display timestamps in UTC
```

### Section Filtering

```bash
juju status --relations    # Include relations section
juju status --storage      # Include storage section
```

## Output Format Summary Table

| Command | Default | Supports JSON | Supports YAML | Supports Tabular |
|---------|---------|---------------|---------------|------------------|
| `status` | tabular | Yes | Yes | Yes |
| `models` | tabular | Yes | Yes | Yes |
| `controllers` | tabular | Yes | Yes | Yes |
| `config` | tabular | Yes | Yes | Yes |
| `clouds` | tabular | Yes | Yes | Yes |
| `credentials` | tabular | Yes | Yes | Yes |
| `secrets` | tabular | Yes | Yes | Yes |
| `storage` | tabular | Yes | Yes | Yes |
| `deploy` | text | No | No | No |
| `add-model` | text | No | No | No |
| `bootstrap` | text | No | No | No |
