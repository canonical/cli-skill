# Juju CLI Configuration Model

## Overview

Juju CLI configuration follows a layered precedence model where values from different sources override each other in a defined order. This document describes the configuration sources, precedence rules, and per-command configuration behavior.

## Configuration Sources

### 1. Environment Variables

Environment variables provide the highest-priority configuration override mechanism.

| Variable | Purpose | Default | Example |
|----------|---------|---------|---------|
| `JUJU_DATA` | Configuration directory | `~/.local/share/juju` | `/custom/config/path` |
| `JUJU_CONTROLLER` | Default controller | (from accounts.yaml) | `mycontroller` |
| `JUJU_MODEL` | Default model | (current model) | `mymodel` |
| `JUJU_LOGGING_CONFIG` | Logging level | `<root>=WARNING` | `<root>=DEBUG` |
| `JUJU_DEV_FEATURE_FLAGS` | Developer features | (empty) | `developermode` |
| `JUJU_FEATURES` | User features | (empty) | `feature1,feature2` |
| `JUJU_STATUS_ISO_TIME` | ISO timestamps in status | `false` | `true` |
| `JUJU_STARTUP_LOGGING_CONFIG` | Startup logging | (empty) | `<root>=TRACE` |
| `JUJU_CONTAINER_TYPE` | Container type (MAAS) | (empty) | `lxd` |
| `XDG_DATA_HOME` | XDG data directory | `~/.local/share` | `/custom/xdg` |

### 2. Configuration Files

Configuration files store persistent settings in YAML format.

#### File Locations

| File | Location | Purpose |
|------|----------|---------|
| `controllers.yaml` | `$JUJU_DATA/` | Controller definitions |
| `accounts.yaml` | `$JUJU_DATA/` | User accounts per controller |
| `models.yaml` | `$JUJU_DATA/` | Model UUID mappings |
| `credentials.yaml` | `$JUJU_DATA/` | Cloud credentials |
| `aliases` | `$JUJU_DATA/` | User-defined command aliases |
| `bootstrap-params.yaml` | `$JUJU_DATA/` | Bootstrap parameter cache |

#### controllers.yaml Structure

```yaml
controllers:
  mycontroller:
    uuid: "12345678-1234-1234-1234-123456789012"
    api-endpoints:
      - "192.168.1.1:17070"
      - "192.168.1.2:17070"
    ca-cert: |
      -----BEGIN CERTIFICATE-----
      ...
      -----END CERTIFICATE-----
    cloud: aws
    region: us-east-1
    type: iaas
```

#### accounts.yaml Structure

```yaml
controllers:
  mycontroller:
    user: admin
    password: ""
    last-known-access: "2024-01-01T00:00:00Z"
```

#### credentials.yaml Structure

```yaml
credentials:
  aws:
    default-credential: my-credential
    default-region: us-east-1
    my-credential:
      auth-type: access-key
      access-key: AKIAXXXXX
      secret-key: xxxxxx
```

### 3. Command-Line Flags

Flags provide per-invocation configuration overrides.

#### Global Flags (All Commands)

| Flag | Environment Variable | Purpose |
|------|---------------------|---------|
| `-m`, `--model` | `JUJU_MODEL` | Target model |
| `-c`, `--controller` | `JUJU_CONTROLLER` | Target controller |
| `-B`, `--no-browser-login` | - | Disable browser auth |
| `-q`, `--quiet` | - | Suppress output |
| `-v`, `--verbose` | - | Verbose output |
| `--debug` | - | Debug logging |

#### Output Flags (Most Commands)

| Flag | Purpose |
|------|---------|
| `--format` | Output format (yaml, json, tabular, etc.) |
| `-o`, `--output` | Write output to file |
| `--color` | Enable ANSI colors |
| `--no-color` | Disable ANSI colors |

## Configuration Precedence

Values are resolved in the following order (highest precedence first):

```
┌─────────────────────────────────────────────────────────────┐
│  1. Command-line flags                                        │
│     --model=mymodel                                          │
├─────────────────────────────────────────────────────────────┤
│  2. Environment variables                                    │
│     JUJU_MODEL=mymodel                                       │
├─────────────────────────────────────────────────────────────┤
│  3. User aliases                                             │
│     alias "deploy-test" = "deploy --model test"             │
├─────────────────────────────────────────────────────────────┤
│  4. Configuration files                                      │
│     accounts.yaml: current-model                             │
├─────────────────────────────────────────────────────────────┤
│  5. Built-in defaults                                        │
│     default-model, default-controller                        │
└─────────────────────────────────────────────────────────────┘
```

## Per-Command Configuration

### Model-Targeted Commands

Commands operating on models accept model specification via:

1. **Flag**: `--model=<spec>`
2. **Environment**: `JUJU_MODEL=<spec>`
3. **Current model**: From `accounts.yaml`

Model specification formats:
- `<name>` - Model in current controller
- `<controller>:<name>` - Model in specific controller
- `<UUID>` - Direct model UUID

### Controller-Targeted Commands

Commands operating on controllers accept controller specification via:

1. **Flag**: `--controller=<name>`
2. **Environment**: `JUJU_CONTROLLER=<name>`
3. **Current controller**: From `accounts.yaml`

### Cloud Configuration Commands

Cloud-specific commands use both client and controller targets:

```bash
# Client-side operation
juju add-cloud --client mycloud

# Controller-side operation
juju add-cloud --controller mycontroller mycloud

# Both (default for some commands)
juju update-public-clouds
```

## Configuration Inheritance

### Bootstrap Configuration

Bootstrap inherits from multiple sources:

```
bootstrap configuration sources:
├── Command-line flags (--config, --constraints)
├── Environment (JUJU_DATA for credentials)
├── Cloud credentials (credentials.yaml)
├── Model defaults (model-defaults.yaml if exists)
└── Built-in defaults
```

### Deployment Configuration

Deployment commands inherit configuration:

```bash
juju deploy postgresql
├── Model: from --model or current
├── Constraints: from --constraints or model
├── Configuration: from --config or charm defaults
├── Channel: from --channel or "stable"
└── Credential: from model or user default
```

## Configuration Validation

### Type Validation

Configuration values are validated against expected types:

| Configuration | Type | Validation |
|--------------|------|------------|
| Constraints | string | Parsed via constraints parser |
| Config key=value | varies | Validated against charm schema |
| Model name | string | Must match `^[a-z][a-z0-9]*(?:-[a-z0-9]+)*$` |
| Controller name | string | Must match `^[a-z][a-z0-9-]*[a-z0-9]$` |

### Validation Errors

Invalid configuration produces descriptive errors:

```
ERROR invalid model name "MyModel": model names may only contain lowercase letters, digits and hyphens
ERROR invalid constraint "mem=8GB": value must be a number with optional suffix (M, G, T)
ERROR unknown configuration key "invalid-key"
```

## Configuration Caching

### API Connection Caching

The CLI caches API connections per controller:

```yaml
# models.yaml
current-model: mymodel
controllers:
  mycontroller:
    current-model: production
    models:
      production:
        uuid: "..."
      development:
        uuid: "..."
```

### Credential Caching

Credentials are cached after successful authentication:

```yaml
# accounts.yaml
controllers:
  mycontroller:
    user: admin
    password: (cached)
    last-known-access: "2024-01-01T00:00:00Z"
```

## Configuration Migration

### Version Compatibility

Configuration files are versioned for compatibility:

```yaml
# controllers.yaml
controllers:
  mycontroller:
    uuid: "..."
    agent-version: "3.4.0"  # Tracked for compatibility
```

### Upgrade Behavior

When upgrading Juju:
1. New fields are added with defaults
2. Deprecated fields are preserved but ignored
3. Breaking changes require explicit migration

## Configuration Debugging

### Inspecting Current Configuration

```bash
# Show current controller and model
juju whoami

# Show all controllers
juju controllers

# Show all models
juju models

# Show model configuration
juju model-config

# Show controller configuration
juju controller-config
```

### Debug Logging

Enable debug logging for configuration resolution:

```bash
JUJU_LOGGING_CONFIG="<root>=DEBUG" juju deploy mysql
```

### Configuration File Inspection

```bash
# List configuration files
ls -la ~/.local/share/juju/

# Validate YAML syntax
cat ~/.local/share/juju/controllers.yaml | python3 -m yaml
```

## Surprising Precedence Behaviors

### 1. Model Selection Priority

When both `-m` and `JUJU_MODEL` are set, `-m` takes precedence:

```bash
export JUJU_MODEL=model-a
juju status -m model-b  # Uses model-b, not model-a
```

### 2. Client vs Controller Cloud Operations

Cloud operations default to "both client and controller":

```bash
# Without flags, operates on both
juju update-public-clouds

# Explicit single target
juju update-public-clouds --client
juju update-public-clouds --controller mycontroller
```

### 3. Credential Resolution

Credentials are resolved per-cloud, not per-model:

```bash
# Model uses cloud's default credential
juju add-model mymodel aws

# Override per-model
juju add-model mymodel aws --credential my-credential
```

### 4. Alias Expansion

Aliases are expanded before flag parsing:

```yaml
# aliases file
deploy-test = deploy --model test --channel edge
```

```bash
juju deploy-test mysql  # Equivalent to: juju deploy --model test --channel edge mysql
```

### 5. Configuration File Location Override

`JUJU_DATA` affects all file operations:

```bash
export JUJU_DATA=/tmp/juju-config
juju bootstrap aws  # Uses controllers.yaml from /tmp/juju-config/
```
