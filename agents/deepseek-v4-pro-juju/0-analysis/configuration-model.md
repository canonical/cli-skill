# Configuration Model

## Config Sources and Precedence

Juju's configuration model follows a layered approach with explicit precedence:

### 1. Precedence Order (Highest to Lowest)

1. **Command-line flags** — Direct overrides via `--config key=value`, `--file config.yaml`, etc.
2. **Environment variables** — `JUJU_DATA`, `JUJU_MODEL`, `JUJU_CONTROLLER`, `JUJU_FEATURES`, `JUJU_LOGGING_CONFIG`, `NO_COLOR`, proxy variables
3. **Local client store** (`$XDG_DATA_HOME/juju/`) — Persistent state for controllers, models, accounts, credentials, clouds
4. **Embedded public cloud metadata** — Compiled-in list of public clouds (AWS, Azure, GCP, etc.)
5. **Application/model/controller defaults** — Default configuration values set by the system or administrator

### 2. Client Store Layout

```
~/.local/share/juju/
├── accounts.yaml          # Per-controller user account details
├── controllers.yaml       # Known controllers (endpoints, CA certs)
├── models.yaml            # Known models per controller
├── credentials.yaml       # Cloud credentials
├── clouds.yaml            # User-defined clouds
├── ssh/                   # SSH keys for Juju-managed machines
├── aliases                # User-defined command aliases
└── jimm-kube-config.yaml  # Kubernetes config for JAAS
```

### 3. Config File Format

Configuration files are in **YAML** format. The `--file` flag (also `-f`) accepts a YAML file path:

```yaml
# Example: model config file
default-space: dmz
development: true
test-mode: true
```

For application config, the `--config` flag accepts either:
- A path to a YAML file
- Inline `key=value` pairs (repeatable)

### 4. Command-Specific Config Commands

Three distinct config commands exist for different scopes:

| Command | Scope | Storage |
|---------|-------|---------|
| `config` | Application-level | Controller-side application config |
| `model-config` | Model-level | Controller-side model config |
| `controller-config` | Controller-level | Controller-side controller config |
| `model-defaults` | Default model config per cloud | Controller-side defaults |

### 5. Configuration Operations

All config commands support:

- **Get**: `juju config <app> [key]` — reads specific or all keys
- **Set**: `juju config <app> key=value` — sets individual keys
- **Reset**: `juju config <app> --reset key1,key2` — reverts to defaults
- **File input**: `juju config <app> --file config.yaml` — bulk set from file

### 6. Surprising Precedence Behaviors

- **Controller config vs model config**: Some settings exist at multiple levels. Controller config may cascade to models, but explicit model config overrides controller config.
- **Model defaults vs model config**: `model-defaults` sets default values for new models; they do not affect existing models. Existing models use `model-config`.
- **Cloud region config**: Region-specific configuration can override cloud-level configuration.
- **Bootstrap config**: `bootstrap --config` and `bootstrap --model-default` set configuration for the initial controller model and defaults for future models respectively.
- **Constraints**: Model constraints, application constraints, and machine constraints cascade (application overrides model, machine overrides application).
- **`--force` bypass**: The `--force` flag allows commands to proceed despite configuration validation failures, but does not change the stored configuration.

### 7. Credential Configuration

Credentials are stored in `credentials.yaml` and managed by:

- `add-credential` — adds/replaces credentials for a cloud
- `update-credential` — updates credential values (can read from YAML file via `-f`)
- `remove-credential` — removes stored credentials
- `default-credential` — sets which credential is used by default for a cloud
- `autoload-credentials` — detects and imports credentials from the local environment

### 8. Proxy Configuration

Proxies are configured via standard environment variables (`HTTP_PROXY`, `HTTPS_PROXY`, `NO_PROXY`) and are detected at startup. The proxy configuration is installed into the HTTP transport automatically during `juju` initialization.

### 9. Feature Flags

Feature flags are controlled via the `JUJU_FEATURES` environment variable, which is read at init time:

```bash
JUJU_FEATURES=developer-mode juju dump-model
```

Feature flags affect which commands are registered (e.g., `dump-model` and `dump-db` require `DeveloperMode` feature flag enabled).
