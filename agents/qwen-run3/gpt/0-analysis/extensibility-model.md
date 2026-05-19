# qwen36 Extensibility Model

## Main Extension Boundary

The primary extension mechanism is not a plugin command API. It is the runtime engine manifest system.

## Adding New Commands

Command extensibility is compile-time only.

To add a new command:

1. add a constructor under `cli/cmd/cli/commands` or `cli/cmd/cli/commands/debug`
2. return a `*cobra.Command`
3. register it in `cli/cmd/cli/main.go`
4. rebuild the snap

There is no dynamic discovery of external command binaries, shared objects, or Go plugins.

## Engine Extensibility

New engines are discovered at runtime from the engines directory.

Each engine currently consists of:

- `engines/<name>/engine.yaml`
- `engines/<name>/server`
- `engines/<name>/check-server`
- shared shell logic via `common.sh`
- matching component definitions in `snap/snapcraft.yaml`

The manifest schema includes at least:

- identity: `name`, `description`, `vendor`, `grade`
- compatibility constraints: devices, memory, disk-space
- required snap components
- engine-owned configuration values

This means engine packagers can extend the runtime without changing the Cobra command tree, as long as the manifest schema and runtime scripts are respected.

## Configuration Extensibility

The config system has two extension paths:

### standard keys

- package hooks can add defaults
- engine manifests can add engine-scoped keys
- the CLI can expose or consume those keys later

### passthrough keys

The open-ended namespace `passthrough.environment.*` lets users inject arbitrary environment variables for hidden `run`.

That is powerful, but it is also loosely controlled because it bypasses normal key-existence validation.

## Feature Gating As A Soft Extension Mechanism

`chat` and `webui` are conditionally registered at startup based on `ADDITIONAL_FEATURES`.

That gives maintainers a way to expose or hide parts of the command set without recompiling. It is a soft extension boundary rather than a plugin API.

## Hidden Debug Namespace

The `debug` subtree acts as a developer extension lane:

- `validate-engines`
- `select-engine`
- `chat`
- `serve-webui`

These commands are useful for engine authors, CI, and packaging work, but are intentionally separated from the normal user help surface.

## Middleware And Hooks

The command host uses minimal cross-cutting hooks:

- one `PersistentPreRunE` to turn on verbose logging through an env var
- normal Cobra completion hooks for dynamic engine-name completion

There is no richer middleware model for auth, audit, tracing, or pre/post command interception.

## Extension Boundaries Summary

| Area | Extensible? | Mechanism | Constraints |
|---|---|---|---|
| top-level commands | yes | compile-time Cobra registration | requires code change and rebuild |
| debug commands | yes | compile-time Cobra registration | still hidden unless surfaced |
| engines | yes | runtime manifest discovery | must match manifest schema and component packaging |
| engine defaults | yes | manifest `configurations` | no user-facing schema docs |
| config escape hatch | yes | `passthrough.environment.*` | weak validation |
| output formats | yes | code change in each command | no centralized formatter layer |
| chat/webui exposure | yes | `ADDITIONAL_FEATURES` | currently inconsistent with shipped snap manifest |

## Assessment

Strengths:

- engine addition is relatively cheap compared with command addition
- command and engine concerns are separated cleanly
- hidden debug commands already serve as a staging area for advanced behavior

Weaknesses:

- no formal plugin model for commands
- no published manifest schema or compatibility contract for third parties
- feature gating is operationally useful but currently creates product-surface drift
