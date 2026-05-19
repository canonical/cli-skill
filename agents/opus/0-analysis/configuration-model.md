# Juju CLI Configuration Model

## Config Sources and Precedence

Juju configuration is resolved from multiple sources, with the following precedence (highest to lowest):

1. **Command-line flags** (e.g., `--config key=value`, `--constraints`)
2. **Environment variables** (e.g., `JUJU_MODEL`, `JUJU_CONTROLLER`)
3. **Client store files** (`~/.local/share/juju/`)
4. **Built-in defaults**

## Client Store

The client store is a YAML-based local configuration repository at `~/.local/share/juju/`:

| File | Purpose |
|---|---|
| `controllers.yaml` | Controller endpoints, CA certificates, API ports |
| `models.yaml` | Model metadata per controller (name, UUID, type, owner) |
| `accounts.yaml` | User credentials, macaroon data per controller |
| `cookies/` | HTTP cookie jar for macaroon authentication |
| `ssh/` | Cached SSH keys and known_hosts |
| `aliases` | User-defined command aliases |
| `public-clouds.yaml` | Cached public cloud definitions |

### Controller Store (`controllers.yaml`)

```yaml
controllers:
  mycontroller:
    uuid: <uuid>
    api-endpoints: ["10.0.0.1:17070"]
    ca-cert: <PEM>
    cloud: aws
    region: us-east-1
    agent-version: 3.5.0
```

### Models Store (`models.yaml`)

```yaml
controllers:
  mycontroller:
    current-model: admin/default
    models:
      admin/default:
        uuid: <uuid>
        type: iaas
        owner: admin
```

## Environment Variables

| Variable | Effect |
|---|---|
| `JUJU_CONTROLLER` | Default controller name |
| `JUJU_MODEL` | Default model name (format: `[controller:]<model>`) |
| `JUJU_DATA` | Overrides `~/.local/share/juju` path |
| `JUJU_LOGGING_CONFIG` | Default logging configuration |
| `JUJU_FEATURE_FLAGS` | Enables feature flags (e.g., `developer-mode`) |
| `JUJU_USER_DOMAIN` | Preferred domain for SSO/macaroon discharge |
| `NO_COLOR` | Disables colored output |

## Command-Specific Config Overrides

### Bootstrap Configuration

`juju bootstrap` has the most complex configuration path:

1. Cloud definition (public clouds.yaml, personal clouds.yaml)
2. Credential (from `credentials.yaml` or detected via `add-credential`)
3. Bootstrap config file (`--config` file or key=value pairs)
4. Bootstrap constraints (`--bootstrap-constraints`)
5. Model constraints (`--constraints`)
6. Agent version (`--agent-version`)

### Model Configuration

`juju model-config` and `juju model-defaults` operate on controller-stored config:

- `model-config` affects the current model
- `model-defaults` sets defaults for new models
- Both accept `key=value` pairs as positional arguments

### Application Configuration

`juju config <application>` gets/sets application-level config:

- Without key=value pairs: displays current config
- With key=value pairs: sets values
- Setting to empty string restores default (not unsets)

## Config Precedence Surprises

1. **Application constraints override model constraints**: When both are set, application constraints take precedence for that application's unit provisioning.

2. **Cloud credentials are merged**: `credentials.yaml` can contain both client-side and controller-side credentials. The `--client` and `--controller` flags on credential commands control which side is targeted.

3. **Feature flags are process-global**: `JUJU_FEATURE_FLAGS` is read once at process startup via `featureflag.SetFlagsFromEnvironment()`. It cannot be changed per-command.

4. **Model resolution is sticky**: If `-m` is not provided, the current model from the client store is used. There is no automatic model discovery.

5. **Bootstrap config file vs inline**: `--config file.yaml` and `--config key=value` can both be specified and are merged.

## Configuration File Formats

| Config Type | Format | Location |
|---|---|---|
| Cloud definitions | YAML | `~/.local/share/juju/clouds.yaml` |
| Credentials | YAML | `~/.local/share/juju/credentials.yaml` |
| Bootstrap config | YAML | User-provided file |
| Model config | Key=Value | Controller database (queried via API) |
| Controller config | Key=Value | Controller database |

## Validation

Configuration validation happens at multiple levels:
- **Client-side**: Basic key validation for known config keys
- **Controller-side**: Full validation including value ranges, provider-specific checks
- **Bootstrap-time**: Cloud/credential compatibility checks
