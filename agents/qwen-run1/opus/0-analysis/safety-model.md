# Safety Model

## Destructive Operations

### Engine Selection (`use-engine`)

**Risk level**: Medium — overwrites current engine configuration (multiple snap options changed simultaneously)

**Safeguards**:
- Interactive confirmation prompt (default behavior)
- `--assume-yes` flag to bypass confirmation (used only in automated contexts like install hooks)

**Recovery**: Re-run `use-engine` with the previous engine name or `--auto` to restore.

### Configuration Changes (`set`)

**Risk level**: Low — individual key-value changes

**Safeguards**: None observed. Changes take effect immediately.

**Recovery**: Re-run `set` with the previous value. No undo/history mechanism.

### Server Restart

**Risk level**: Low-Medium — changing config while server is running requires restart to take effect

**Safeguards**: The server daemon reads config at startup. Config changes don't hot-reload; a service restart is needed (`snap restart qwen36.server`).

## Confirmation Prompts

| Command | Confirmation | Bypass Flag |
|---------|-------------|-------------|
| `use-engine <name>` | Yes (interactive) | `--assume-yes` |
| `use-engine --auto` | Yes (interactive) | `--assume-yes` |
| `set` | No | — |
| `get` | No (read-only) | — |
| `show-engine` | No (read-only) | — |
| `chat` | No | — |

## Dry-Run Support

**Not available.** No `--dry-run` flag is observed on any command.

## Force Flags

**Not available.** No `--force` flag is observed. The `--assume-yes` flag on `use-engine` is the closest equivalent (skips confirmation, does not force past errors).

## Network Safety

- **Server binds to localhost by default** (`http.host=127.0.0.1`) — not exposed to network
- **No TLS**: The server uses plain HTTP on localhost
- **No authentication**: The OpenAI-compatible API has no auth mechanism within the snap

## Resource Safety

- **Memory**: Engine YAML declares minimum memory requirements (32G for CPU, 48G for CUDA) — these are validated during engine selection, not at runtime
- **Disk space**: Engine YAML declares minimum disk requirements (25G)
- **GPU VRAM**: CUDA engine requires 24G VRAM — validated during `--auto` detection
- **Component timeout**: Server startup waits up to 3600s for components, then stops the service rather than running indefinitely

## Data Safety

- **No user data stored**: The CLI doesn't manage user data; it only manages inference configuration
- **Model files are read-only**: Snap components containing model weights are immutable
- **Configuration is snap-managed**: Stored in snapd's internal database, not in user-writable files

## Privilege Model

- **Snap confinement**: strict — limits file system and network access
- **Install hook**: Runs with snap's default hook privileges
- **Hardware access**: Requires explicit interface connections (`hardware-observe`, `opengl`, `network-bind`)
- **No sudo required**: CLI commands run as the calling user within snap's sandbox
