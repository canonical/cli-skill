# Architecture

## Summary

Juju is a **client-server CLI application** where the `juju` client communicates with Juju controllers over a WebSocket-based API (via `github.com/juju/juju/api`). The CLI is structured as a **flat command hierarchy** with approximately 130 top-level commands registered into a single `SuperCommand` (the root `juju` command). There is no nested subcommand grouping — all commands are peers at the same level, though source code is organized into domain packages (`cloud`, `controller`, `model`, `application`, etc.) under `cmd/juju/`.

## Architecture Style

- **Primary: Client-server CLI** — The `juju` binary connects to remote Juju controllers via API calls. Most commands open an API connection, issue one or more requests, and display results. The controller and its agents (the server side) reside in `cmd/jujud`.
- **Secondary: Plugin-based architecture** — Unknown subcommands are resolved via the plugin system: if `juju foo` does not match a registered command, the CLI searches `$PATH` for an executable named `juju-foo` and delegates to it (see `cmd/juju/commands/plugin.go`).

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Language | Go 1.x |
| CLI Framework | Custom `cmd` package (`cmd/cmd/`) atop `github.com/juju/gnuflag` |
| Flag parsing | `gnuflag` (GNU-style getopt clone) with `FlagKnownAs: "option"` |
| Config storage | Local filesystem under `$XDG_DATA_HOME/juju` (~/.local/share/juju) via `jujuclient.ClientStore` |
| API transport | WebSocket + HTTP (macaroon-based authentication via `httpbakery`) |
| Output formats | JSON, YAML, tabular (human-readable), smart (default heuristic) |
| Controller | `jujud` daemon with domain services + DQLite |
| Plugins | `PATH`-based discovery of `juju-*` executables |

## Command Structure

All commands are top-level children of the `juju` SuperCommand. There is no deep nesting. The source organizes commands into these domain packages:

- `cmd/juju/action` — run, exec, operations, actions, cancel-task, show-action, show-operation, show-task
- `cmd/juju/application` — deploy, add-unit, remove-unit, config, expose, unexpose, integrate, refresh, trust, scale-application, bind, resolved, consume, suspend-relation, resume-relation, remove-relation, remove-application, remove-saas, show-application, show-unit, constraints, diff-bundle
- `cmd/juju/backups` — create-backup, download-backup
- `cmd/juju/block` — disable-command, enable-command, disabled-commands
- `cmd/juju/caas` — add-k8s, update-k8s, remove-k8s
- `cmd/juju/charmhub` — info, find, download
- `cmd/juju/cloud` — add-cloud, update-cloud, remove-cloud, clouds, show-cloud, regions, add-credential, update-credential, remove-credential, credentials, show-credential, default-region, default-credential, autoload-credentials, update-public-clouds
- `cmd/juju/commands` — bootstrap, debug-log, migrate, switch, sync-agent-binary, upgrade-model, upgrade-controller, version
- `cmd/juju/controller` — add-model, destroy-controller, kill-controller, enable-destroy-controller, controllers, show-controller, register, unregister, controller-config, models
- `cmd/juju/crossmodel` — offer, remove-offer, show-offer, offers, find-offers
- `cmd/juju/dashboard` — dashboard
- `cmd/juju/firewall` — set-firewall-rule, firewall-rules
- `cmd/juju/machine` — add-machine, remove-machine, machines, show-machine
- `cmd/juju/model` — config, destroy-model, show-model, grant, revoke, grant-cloud, revoke-cloud, model-constraints, model-defaults, model-config, set-model-constraints, set-credential, retry-provisioning, export-bundle, dump-model, dump-db
- `cmd/juju/resource` — attach-resource, resources, charm-resources (internal)
- `cmd/juju/secretbackends` — secret-backends, add-secret-backend, update-secret-backend, remove-secret-backend, show-secret-backend, model-secret-backend
- `cmd/juju/secrets` — secrets, add-secret, update-secret, remove-secret, show-secret, grant-secret, revoke-secret
- `cmd/juju/space` — add-space, remove-space, spaces, show-space, rename-space, move-to-space, reload-spaces
- `cmd/juju/ssh` — ssh, scp, debug-code, debug-hooks
- `cmd/juju/sshkeys` — add-ssh-key, remove-ssh-key, import-ssh-key, ssh-keys
- `cmd/juju/status` — status, show-status-log
- `cmd/juju/storage` — add-storage, attach-storage, detach-storage, import-filesystem, remove-storage, show-storage, storage, create-storage-pool, remove-storage-pool, update-storage-pool, storage-pools
- `cmd/juju/subnet` — subnets
- `cmd/juju/user` — add-user, change-user-password, disable-user, enable-user, remove-user, show-user, users, login, logout, whoami

## Command Base Types (Embedding Hierarchy)

```
cmd.Command (interface)
  └── cmd.CommandBase
        └── modelcmd.CommandBase       // adds model/controller selection (+ controller flags)
              ├── modelcmd.ModelCommandBase   // requires model selection
              │     └── modelcmd.OptionalModelCommandBase  // optional model
              └── modelcmd.ControllerCommandBase           // controller only (no model)
                    └── modelcmd.OptionalControllerCommandBase
```

Common flags injected by the base types:
- `-m, --model <name>` — target model
- `-c, --controller <name>` — target controller
- `--no-color` — disable ANSI color output

## Configuration Sources & Precedence

1. **Command-line flags** (highest)
2. **Environment variables** (e.g., `JUJU_DATA`, `JUJU_MODEL`, `JUJU_CONTROLLER`)
3. **Local client store** — `$XDG_DATA_HOME/juju/` (~/`.local/share/juju/`)
   - `accounts.yaml` — per-controller user credentials
   - `controllers.yaml` — known controllers and endpoints
   - `models.yaml` — known models per controller
   - `credentials.yaml` — cloud credentials
   - `clouds.yaml` — user-defined clouds
4. **Public cloud metadata** — embedded public cloud definitions
5. **Defaults** (lowest)

## Plugin System

If a subcommand is not recognized, the `MissingCallback` is invoked, which calls `RunPlugin()`. This searches `$PATH` for `juju-<subcommand>` executables. Common Juju flags (`-m`, `--model`, `-c`, `--controller`) are extracted and passed to the plugin.

## Interactive Shell

When `juju` is run with no arguments in a terminal, it enters an interactive REPL shell (`cmd/juju/commands/repl.go` and `cmd/juju/interact/`).
