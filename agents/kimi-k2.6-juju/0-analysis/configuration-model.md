# Configuration Model

## Sources (highest to lowest precedence)
1. **Command-line flags** (e.g., `--config key=value`, `--model-default`, `--constraints`).
2. **Environment variables** (e.g., `JUJU_CONTROLLER`, `JUJU_MODEL`, `JUJU_LOGGING_CONFIG`, `JUJU_DATA`, `JUJU_FEATURE_FLAGS`).
3. **Client configuration files** in `~/.local/share/juju/`:
   - `controllers.yaml` — controller endpoints, CA certs.
   - `models.yaml` — model metadata, active model.
   - `accounts.yaml` — user credentials, macaroons.
   - `cookies` — authentication cookies.
   - `public-clouds.yaml` — cloud metadata fetched on first run.
4. **Controller/model configuration** persisted on the server and retrieved via API.
5. **Built-in defaults** defined in charm metadata, provider code, and Juju source constants.

## Command-specific Overrides
- `bootstrap` can set controller config keys via `--config` that become defaults for the controller model.
- `model-config` and `controller-config` mutate server-side configuration; local file precedence only matters when bootstrapping.
- `deploy --config` overrides model defaults for the new application.
- `set-constraints` overrides model constraints for an application.

## Surprising Precedence
- The `--model` flag overrides the `JUJU_MODEL` env var, which overrides the `models.yaml` current-model setting.
- User aliases (`aliases` file in Juju data dir) are processed *after* flag parsing but *before* subcommand lookup, allowing transparent argument injection.
- Bootstrap-time `--config` values override public cloud metadata defaults but not provider-level hardcoded defaults.
