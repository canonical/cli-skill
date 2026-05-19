# qwen36 Safety Model

## Safety Profile

This CLI has relatively few destructive commands in the data-loss sense, but it has several commands that can disrupt service availability, alter hardware usage, or remove large downloaded components.

## High-Impact Commands

| Command | Risk | What it changes | Guardrails |
|---|---|---|---|
| `use-engine` | High | active engine, engine config, installed components, daemon restart | component list preview, confirmation prompt, optional restart prompt |
| `prune-cache` | High | removes inactive-engine components from disk | explicit confirmation prompt; refuses active engine |
| `set` | Medium | persistent user config, possible daemon restart | restart prompt; no dry-run |
| `unset` | Medium | removes user overrides, possible daemon restart | restart prompt; no dry-run |
| `run` | Medium | temporary engine env and symlink layout, subprocess execution | waits for components but no preview |

## Confirmation Model

| Command | Prompt | Default | Override |
|---|---|---|---|
| `use-engine` component install | `Do you want to continue?` | yes | `--assume-yes` |
| `use-engine` restart | `Restart <snap> to apply the changes?` | yes | `--assume-yes`, `--no-restart` |
| `set` restart | `Restart <snap> to apply the changes?` | yes | `--assume-yes`, `--no-restart` |
| `unset` restart | `Restart <snap> to apply the changes?` | yes | `--assume-yes`, `--no-restart` |
| `prune-cache` removal | `Continue pruning ...?` | no | no force flag; non-interactive use skips prompt |

## Dry-Run And Force Support

| Mechanism | Present? | Notes |
|---|---|---|
| Dry-run | no | no public command provides a preview-only mode |
| Force flag | no | the CLI uses confirmations rather than `--force` |
| Partial automation override | yes | `--assume-yes` and `--no-restart` suppress prompts, but only for certain commands |

## Recovery Paths

### Engine switch recovery

- rerun `use-engine <other-engine>` to switch back
- rerun `use-engine --fix` to reinstall missing components and reapply current engine config
- there is no explicit `reset-engine` or `clear-engine`

### Config recovery

- rerun `set` with the prior value
- use `unset` to fall back to lower-precedence defaults
- there is no config snapshot, transaction log, or bulk rollback command

### Cache pruning recovery

- rerun `use-engine <pruned-engine>` to reinstall components for that engine
- there is no dedicated restore command

## Safety Constraints And Gaps

### 1. no dry-run for engine or cache operations

Users cannot preview:

- which config keys will change during `use-engine`
- which components will be removed by `prune-cache` without already entering the destructive flow
- whether a restart will actually be necessary before they commit the change

### 2. config writes are weakly validated

`set` checks key existence but does not enforce type correctness. That means values like a non-numeric `http.port` may be accepted and fail only later at runtime.

### 3. `run` has cleanup risk on abnormal termination

`run` creates temporary symlinks for engine layout paths and relies on deferred cleanup. The source explicitly notes that cleanup does not run on SIGTERM or SIGKILL.

### 4. no service-management symmetry inside the CLI

The CLI prompts for restart, but actual start/stop/restart control is delegated to `snap start|stop|restart`. That split is workable, but it weakens the mental model because mutation commands affect services without providing a first-class service command set.

### 5. feature-gated user commands affect safety expectations

`chat` and `webui` are documented and coded as user-facing commands, but the shipped qwen36 snap does not enable them by default. That means users may turn to lower-level hidden commands or direct `snap` operations when the documented path is missing.

## Overall Assessment

Strengths:

- the CLI does prompt before expensive component installs and most restarts
- cache pruning refuses the active engine
- `use-engine --fix` provides a useful repair path

Weaknesses:

- no dry-run
- no rollback primitive
- no structured safety output for automation
- hidden or mismatched feature exposure complicates safe troubleshooting
