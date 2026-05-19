# Safety Model

## Overview

The `qwen36` CLI manages a machine-learning inference stack that includes downloading large model files (tens of GB), switching inference engines, removing cached components, and restarting system services. The safety model balances convenience with safeguards against accidental data loss or service disruption.

## Destructive Operations

### 1. Engine Switching (`use-engine`)

**Destructive potential**: Switching engines uninstalls the old engine's components, changes system service configuration, and may require a snap restart.

**Safeguards**:
- **Root guard**: Requires `sudo` (`utils.IsRootUser()` check, returns `ErrPermissionDenied`).
- **Component installation confirmation**: When switching to a new engine, missing components are listed with sizes and the user is prompted: `"Do you want to continue? [Y/n]"` (default: yes).
- **`--assume-yes` flag**: Bypasses confirmation prompts.
- **Cancellation**: Responding "no" exits with `"Cancelled. No changes applied."` and exit code 0. No partial changes are left.
- **Restart prompt**: After switching, user is prompted: `"Restart <snap> to apply the changes? [Y/n]"` (default: yes).
- **`--no-restart` flag**: Suppresses restart prompt.

**Missing safeguard**: There is no **dry-run** mode. The user cannot see what components will be installed/removed without actually performing the operation.

### 2. Cached Component Pruning (`prune-cache`)

**Destructive potential**: Removes snap components (engine binaries, model weights) that are not needed by the active engine. This frees disk space but cannot be undone without re-downloading.

**Safeguards**:
- **Root guard**: Requires `sudo`.
- **Active engine protection**: Refuses to prune the active engine: `cannot prune the active engine "<name>"`.
- **Component listing with sizes**: Before pruning, lists all components to be removed with their installed sizes.
- **Confirmation prompt**: `"Continue pruning [<engines>] engines? [y/N]"` — **default is NO** (unlike `use-engine` component install which defaults to YES).
- **Piped/non-interactive safety**: If stdout is not a terminal (`!utils.IsTerminalOutput()`), the confirmation prompt is **skipped** and the command proceeds — this is a **safety gap** (see below).
- **Cancellation**: Responding "no" or Ctrl+C exits with `"Cancelled. No changes applied."` and exit code 0.

**Safety gap**: When output is piped or non-interactive, `prune-cache` skips the confirmation prompt and proceeds without any confirmation. This means `sudo qwen36 prune-cache | cat` will silently delete components. Compare with `set`/`unset`/`use-engine` which have `--assume-yes` to explicitly opt into non-interactive operation.

### 3. Configuration Changes (`set` / `unset`)

**Destructive potential**: Changing configuration values may affect running services. The CLI prompts to restart the snap to apply changes.

**Safeguards**:
- **Root guard**: Both commands require `sudo`.
- **Restart prompt**: After changes, user is prompted: `"Restart <snap> to apply the changes? [Y/n]"` (default: yes).
- **`--no-restart` flag**: Suppresses restart prompt.
- **`--assume-yes` flag**: Auto-confirms restart.
- **Change detection on `unset`**: If the unset operation does not actually change the effective value (because a lower-tier config provides the same value), no restart prompt is shown.

**Safety gap**: No dry-run. Users cannot preview what the effective config would be after a `set`/`unset` without persisting the change.

### 4. Snap Restart (via prompt)

**Destructive potential**: Restarting the snap stops and restarts the inference server daemon, interrupting any active inference requests.

**Safeguards**:
- **Prompted**: All commands that trigger a restart (`set`, `unset`, `use-engine`) prompt for confirmation.
- **Opt-out flags**: `--no-restart` suppresses entirely; `--assume-yes` auto-confirms.
- **No background restart**: Restart is always user-acknowledged (unless `--assume-yes`).

### 5. Subprocess Execution (`run`)

**Destructive potential**: The hidden `run` command executes an arbitrary subprocess in the engine's environment with passthrough environment variables from config.

**Safeguards**:
- **Hidden command**: Not shown in help, reducing accidental use.
- **Engine environment loading**: Sets up engine-specific environment and symlinks, then cleans up after (via deferred function).
- **No root guard**: Notably, `run` does **not** check for root. It can be run by any user who can access the CLI.
- **No argument sanitization**: The subprocess command and args are passed directly to `exec.Command`. No shell injection protection is needed since `exec.Command` does not invoke a shell, but there is no whitelist of allowed commands.

### 6. Chat Client (`chat` / `debug chat`)

**Destructive potential**: None (read-only interaction with inference server).

**Safeguards**: Server readiness check before starting chat. If server is inactive, suggests starting it.

### 7. Web UI (`webui` / `debug serve-webui`)

**Destructive potential**: `webui` opens a browser via `xdg-open` — this could open a malicious URL if the web UI port is hijacked. `serve-webui` binds an HTTP port.

**Safeguards**:
- **Service readiness check**: `webui` checks that both `server` and `server-webui` services are active before proceeding.
- **OpenAI server health check**: `webui` waits for the chat server to be ready before offering to open the browser.
- **`serve-webui` binds to localhost by default**: The top-level `serve-webui` binds to localhost only; `debug serve-webui` is also fixed to localhost.

## Force Flags

There is **no `--force` flag** on any command. The closest equivalents:
- `--assume-yes`: Bypasses confirmation prompts but does not override safety checks (e.g., cannot force-prune the active engine).
- `--no-restart`: Suppresses restart but does not force anything.

## Dry-Run Support

**No command supports dry-run**. There is no `--dry-run` flag. Users cannot preview the effects of `use-engine`, `prune-cache`, `set`, or `unset` without committing the change.

## Recovery Behavior

### If engine switching fails mid-operation:
- **All-or-nothing install**: Component installation is sequential with retry logic (60-minute timeout, 10-second retry for snapd timeouts/busy states). If one component fails, remaining components are not installed — but already-installed components are not rolled back.
- **Config cleanup**: Old engine config is unset **before** new engine config is set. If switching fails, the old engine's config is already gone but the old engine name is still in cache (the cache write happens after config operations). This means a failed switch can leave the system in a partially-configured state.

### If `prune-cache` is cancelled:
- No changes are applied (confirmation happens before any deletions).

### If snap restart fails:
- The restart error is returned to the user: `restarting snap: <err>`. Config changes are already persisted.

## Safety Summary Table

| Command | Requires Root | Confirmation Prompt | Default Answer | Dry-Run | Can Cancel | Auto-mode Safety |
|---------|---------------|---------------------|----------------|---------|-----------|-------------------|
| `set` | Yes | Restart | Y | No | Yes (`--no-restart`) | `--assume-yes` |
| `unset` | Yes | Restart | Y | No | Yes (`--no-restart`) | `--assume-yes` |
| `use-engine` | Yes | Components + Restart | Y (components), Y (restart) | No | Yes | `--assume-yes` |
| `prune-cache` | Yes | Confirmation | **N** | No | Yes | Non-TTY skips prompt (⚠️ risk) |
| `run` | **No** | None | N/A | No | N/A | N/A |
| `chat` | No | None | N/A | N/A | N/A | N/A |
| `webui` | No | Open browser | N/A | N/A | Ctrl+C | N/A |
| `serve-webui` | No | None | N/A | N/A | N/A | N/A |

**Key finding**: The `N` default for `prune-cache` vs `Y` default for `use-engine` component installation is an inconsistency. The justification is plausible (pruning is purely destructive, installing is preparatory) but may surprise users who expect destructive operations to consistently default to "no".
