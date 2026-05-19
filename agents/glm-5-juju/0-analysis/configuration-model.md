# Juju CLI Configuration Model

## Overview

Juju's configuration system manages multiple sources of settings with a well-defined precedence hierarchy. Configuration influences command behavior, authentication, and connection parameters.

## Configuration Sources

### 1. Command-Line Flags (Highest Priority)

Flags directly override all other configuration sources:

```bash
juju status --model mycontroller:production --format json
juju deploy postgresql --config myconfig.yaml --channel 14/stable
```

### 2. Environment Variables

Juju-specific environment variables provide context without modifying files:

| Variable | Purpose | Example |
|----------|---------|---------|
| `JUJU_CONTROLLER` | Default controller name | `JUJU_CONTROLLER=aws-prod` |
| `JUJU_MODEL` | Default model name or UUID | `JUJU_MODEL=production` |
| `JUJU_DATA` | Configuration directory | `JUJU_DATA=~/.config/juju` |
| `JUJU_LOGGING_CONFIG` | Log level configuration | `JUJU_LOGGING_CONFIG="<root>=DEBUG"` |
| `JUJU_DEV_FEATURE_FLAGS` | Developer feature flags | `JUJU_DEV_FEATURE_FLAGS="developer-mode"` |
| `JUJU_FEATURES` | Production features | `JUJU_FEATURES="secrets"` |
| `JUJU_STATUS_ISO_TIME` | ISO timestamp format | `JUJU_STATUS_ISO_TIME=true` |

### 3. Configuration Files

Juju stores persistent configuration in YAML files within the data directory:

```
~/.local/share/juju/          # Linux default
~/AppData/Roaming/Juju/        # Windows default
~/Library/Application Support/Juju/  # macOS default
```

#### Directory Structure

```
~/.local/share/juju/
├── accounts.yaml              # Controller account details
├── controllers.yaml           # Controller definitions
├── credentials.yaml           # Cloud credentials
├── models.yaml                # Model mappings
├── clouds.yaml                # User-defined clouds
├── bootstrap-config.yaml      # Bootstrap configuration
├── aliases                    # Command aliases
├── ssh/
│   ├── juju_id_rsa           # Juju-generated SSH key
│   └── juju_id_rsa.pub
└── jujuc/
    └── server.pem             # Controller certificate
```

### 4. Defaults (Lowest Priority)

Built-in defaults apply when no other source specifies a value.

## Configuration File Formats

### accounts.yaml

Stores per-controller account information:

```yaml
accounts:
  aws-prod:
    user: admin
    last-known-access: "2024-01-15"
    bootstrap-timeout: 1200
```

### controllers.yaml

Defines known controllers:

```yaml
controllers:
  aws-prod:
    uuid: "abc123-def456"
    api-endpoints:
      - "10.0.0.1:17070"
      - "10.0.0.2:17070"
    ca-cert: |
      -----BEGIN CERTIFICATE-----
      ...
      -----END CERTIFICATE-----
    cloud: aws
    region: us-east-1
    type: iaas
  lxd-local:
    uuid: "xyz789"
    api-endpoints:
      - "10.0.0.10:17070"
    ca-cert: |
      -----BEGIN CERTIFICATE-----
      ...
      -----END CERTIFICATE-----
    cloud: lxd
    region: localhost
    type: iaas

current-controller: aws-prod
```

### credentials.yaml

Cloud authentication credentials:

```yaml
credentials:
  aws:
    default-credential: default
    default-region: us-east-1
    credentials:
      default:
        auth-type: access-key
        access-key: AKIAIOSFODNN7EXAMPLE
        secret-key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
  gce:
    default-credential: production
    credentials:
      production:
        auth-type: jsonfile
        path: /home/user/.config/gcloud/application_default_credentials.json
```

### models.yaml

Model-to-controller mappings and current context:

```yaml
models:
  aws-prod:production:
    uuid: "model-uuid-1"
    controller: aws-prod
    user: admin
  aws-prod:staging:
    uuid: "model-uuid-2"
    controller: aws-prod
    user: admin

current-model: "aws-prod:production"
```

### clouds.yaml

User-defined cloud definitions:

```yaml
clouds:
  my-openstack:
    type: openstack
    auth-types:
      - userpass
      - access-key
    endpoint: "https://openstack.example.com:5000/v3"
    regions:
      region1:
        endpoint: "https://openstack.example.com:5000/v3"
```

## Configuration Precedence

The complete precedence order (highest to lowest):

1. **Command-line flags** - Direct override
2. **Environment variables** - Session context
3. **Configuration files** - Persistent settings
4. **Defaults** - Built-in fallback

### Example: Model Resolution

```bash
# Highest: Explicit flag
juju status -m other-controller:other-model

# High: Environment variable
export JUJU_MODEL=staging
juju status  # Uses staging model

# Medium: Configuration file current-model
# (stored in models.yaml)
juju status  # Uses current-model setting

# Low: Default/error
juju status  # Error if no model specified
```

### Example: Controller Resolution

```bash
# Highest: Explicit flag
juju models -c specific-controller

# High: Environment variable
export JUJU_CONTROLLER=aws-prod
juju models  # Uses aws-prod

# Medium: Configuration file
# (current-controller in controllers.yaml)
juju models  # Uses current-controller

# Low: Error
juju models  # Error if no controllers registered
```

## Application Configuration

### Runtime Configuration

Applications deployed with Juju have configuration managed through the `config` command:

```bash
# View all config
juju config postgresql

# View specific key
juju config postgresql max-connections

# Set configuration
juju config postgresql max-connections=200

# Set from file
juju config postgresql --file=postgres-config.yaml

# Reset to default
juju config postgresql --reset max-connections
```

### Configuration Files (YAML format)

```yaml
postgresql:
  max-connections: 200
  shared-buffers: "256MB"
  log-destination: "stderr"
```

### Charm-Defined Configuration

Charms define their configuration in `config.yaml`:

```yaml
options:
  max-connections:
    type: int
    default: 100
    description: "Maximum number of database connections"
  shared-buffers:
    type: string
    default: "128MB"
    description: "Amount of memory for shared buffers"
```

## Model Configuration

### Model-Level Settings

```bash
# View all model config
juju model-config

# Set model config
juju model-config logging-config="<root>=INFO"

# Set model defaults (for new models)
juju model-defaults default-base=ubuntu@22.04

# Set per-controller model defaults
juju model-defaults -c aws-prod default-base=ubuntu@22.04
```

### Key Model Configuration Options

| Option | Type | Description |
|--------|------|-------------|
| `default-base` | string | Default OS for new machines |
| `logging-config` | string | Log level configuration |
| `http-proxy` | string | HTTP proxy for outgoing connections |
| `https-proxy` | string | HTTPS proxy for outgoing connections |
| `no-proxy` | string | Hosts to bypass proxy |
| `apt-http-proxy` | string | APT-specific proxy |
| `apt-https-proxy` | string | APT-specific HTTPS proxy |
| `enable-os-refresh-update` | bool | Run OS updates on provisioning |
| `enable-os-upgrade` | bool | Run OS upgrades on provisioning |
| `firewall-mode` | string | Firewall implementation mode |

## Controller Configuration

### Controller-Level Settings

```bash
# View controller config
juju controller-config

# Set controller config
juju controller-config auditing-enabled=true
```

### Key Controller Configuration Options

| Option | Type | Description |
|--------|------|-------------|
| `auditing-enabled` | bool | Enable audit logging |
| `api-port` | int | API server port |
| `autocert-url` | string | Auto-cert URL for TLS |
| `login-token-refresh-url` | string | JWT refresh endpoint |
| `max-charm-state-size` | int | Max charm state size (bytes) |
| `max-debug-log-duration` | duration | Max debug-log session time |

## Bootstrap Configuration

### Pre-Bootstrap Options

```bash
juju bootstrap aws production \
  --agent-version=3.6.0 \
  --config logging-config="<root>=DEBUG" \
  --model-default default-base=ubuntu@22.04 \
  --constraints="cores=4 mem=16G"
```

### Bootstrap-Time Configuration Keys

| Key | Type | Description |
|-----|------|-------------|
| `admin-secret` | string | Initial admin password |
| `authorized-keys` | string | SSH keys for controllers |
| `ca-cert` | string | Custom CA certificate |
| `ca-private-key` | string | CA private key |
| `bootstrap-timeout` | int | Bootstrap timeout (seconds) |
| `controller-service-type` | string | K8s service type |

## Feature Flags

### Developer Features (`JUJU_DEV_FEATURE_FLAGS`)

Enable experimental features during development:

```bash
export JUJU_DEV_FEATURE_FLAGS="developer-mode"
juju model-dump  # Only available with developer-mode
juju model-dump-db
```

### Production Features (`JUJU_FEATURES`)

Enable polished features that may have compatibility implications:

```bash
export JUJU_FEATURES="secrets"
```

## Configuration API Integration

### Command Override Patterns

```go
// ModelCommandBase provides model resolution
type ModelCommandBase struct {
    CommandBase
    store           jujuclient.ClientStore
    _modelIdentifier string
    _controllerName  string
    allowDefaultModel bool
}

// Resolution order in SetModelIdentifier:
// 1. Explicit parameter (from flag parsing)
// 2. JUJU_MODEL environment variable
// 3. Current model from models.yaml
// 4. Error if no model available
```

### Client Store Interface

```go
type ClientStore interface {
    ControllerByName(name string) (*ControllerDetails, error)
    CurrentController() (string, error)
    CurrentModel(controller string) (string, error)
    ModelByName(controller, name string) (*ModelDetails, error)
    SetCurrentModel(controller, name string) error
    // ... additional methods
}
```

## Configuration Validation

### Type Validation

Configuration values are validated against schema definitions:

- **int**: Integer values with optional min/max bounds
- **string**: String values with optional regex patterns
- **bool**: Boolean values
- **float**: Floating-point numbers
- **attrs**: Key-value attribute maps

### Secret Handling

Sensitive configuration (passwords, keys) is handled specially:

- Not displayed in output by default
- Masked in logs
- Stored with restricted file permissions (0600)

## Configuration Migration

### Version Upgrades

During upgrades, configuration is preserved:

- Controller configuration migrates to new schema
- Model configuration is preserved
- Credential formats are converted if needed
- Deprecated options are removed with warnings
