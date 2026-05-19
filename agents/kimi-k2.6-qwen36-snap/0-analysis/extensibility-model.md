# Extensibility Model

## Adding New Commands

The CLI uses Cobra's imperative command registration in `main.go`. To add a new command:

1. Create a new Go file under `cmd/cli/commands/` (or `cmd/cli/commands/debug/` for debug commands).
2. Implement a constructor function returning `*cobra.Command`.
3. Register it in `main.go` inside `addCommandGroup()`, `appendCommandToGroup()`, or `addCommands()` depending on desired visibility.
4. Rebuild and redeploy the snap.

There is **no plugin loading at runtime**. The binary is monolithic; all commands are compiled in.

## Engine Extension (Plugin Boundary)

The primary extensibility point is the **engine manifest** system:

- Manifests are YAML files in `$SNAP/engines/<name>.yaml` (or `.json`).
- Each manifest declares:
  - `name`, `description`, `vendor`, `grade`
  - `components`: list of snap component names required
  - `configurations`: map of config keys/values applied to the engine config layer
  - `compatibility` rules (memory, disk, device PCI IDs, CPU features)
- The CLI discovers manifests at runtime via `engines.LoadManifests()`.

**Implication**: Adding a new engine does not require CLI code changes. A snap packager can drop a new manifest and corresponding snap components, and the CLI will automatically discover, score, and manage it.

## Configuration Extensibility

- The `passthrough.` key namespace allows arbitrary config injection without CLI schema changes.
- `passthrough.environment.*` values are injected as environment variables during `run`.
- No validation is applied to passthrough keys, making this an open extension channel with associated safety tradeoffs.

## Hidden / Debug Namespace

The `debug` subcommand is a conventional extension container for non-user-facing functionality:
- `validate-engines`
- `select-engine`
- `chat` (endpoint-agnostic variant)
- `serve-webui` (endpoint-agnostic variant)

These commands are hidden from help but available for developers and CI pipelines.

## Feature Gating

Commands can be gated by the `ADDITIONAL_FEATURES` environment variable at startup:
- `chat` and `webui` are only registered if `ADDITIONAL_FEATURES` contains `chat` or `webui` respectively.
- This is compile-time flag behavior driven by runtime environment, not by config files.

## Middleware / Hooks

- **PersistentPreRunE**: sets the `VERBOSE` environment variable from the global `--verbose` flag.
- No auth middleware, no request logging middleware, no tracing.
- Cobra's native pre/post run hooks are available but unused beyond verbosity.

## Extension Boundaries

| Area | Extensible? | Mechanism |
|---|---|---|
| New CLI command | Compile-time only | Go code + `main.go` registration |
| New engine | Runtime | Manifest file in `$SNAP/engines/` |
| New config key | Runtime | `passthrough.*` or engine manifest |
| New output format | Compile-time only | Add case to `--format` switch in command handler |
| New hardware detector | Compile-time only | Extend `pkg/hardware_info` and `pkg/selector` |
| New chat capability | Partial | `webui.Config.Capabilities` is a string slice; backend must support it |

## Limits of Runtime Extensibility

- The CLI cannot load external Go plugins (no `plugin` package usage).
- The Cobra command tree is built once at startup; no hot-reload.
- Engine manifests are read on each invocation; there is no caching of parsed manifests across CLI runs (other than the active engine name in snapd config).
