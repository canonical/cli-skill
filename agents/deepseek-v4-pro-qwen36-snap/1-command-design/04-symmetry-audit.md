# 04 — Symmetry Audit

All symmetric operation pairs (including missing reverse operations), listed exhaustively.

## Present Symmetric Pairs

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|-----------|----------------|-----------------|-------------------|---------------------|-------|
| Config set/unset | `set <key=value>` | `unset <key>` | ✅ Yes | ✅ Yes | Perfect symmetry. Both require root. Both prompt for restart. Both support `--assume-yes` and `--no-restart`. `set` accepts multiple key=value pairs; `unset` accepts exactly one key. |

---

## Missing Reverse Operations

| Forward Operation | Forward Command | Missing Reverse | Notes |
|-------------------|----------------|-----------------|-------|
| Activate engine | `use-engine <name>` | No `deactivate-engine` / `unuse-engine` | Switching from engine A to engine B deactivates A as a side effect, but there is no standalone "stop using any engine" command. Users who want to clean up must either `prune-cache` or switch to another engine. |
| Select engine (test) | `debug select-engine` | No reverse | `select-engine` is read-only (tests selection, doesn't apply it). No reverse needed. |
| Start chat (auto) | `chat` | No `stop-chat` / `exit-chat` | Chat is an interactive session exited by the user (Ctrl+C or `/exit`). No reverse command needed — the session is self-terminating. |
| Launch web UI | `webui` | No `close-webui` | The browser is opened as a separate process. No reverse — user closes the browser directly. |
| Serve web UI | `serve-webui` | No `stop-webui` | The server runs until killed (Ctrl+C or SIGTERM). No reverse command needed — it's a foreground process. |
| Run subprocess | `run <command>` | No `kill` / `stop-run` | The subprocess runs to completion or until killed. No reverse command. |
| List engines | `list-engines` | No reverse | Read-only operation. No reverse needed. |
| Show engine | `show-engine` | No reverse | Read-only. No reverse needed. |
| Show machine | `show-machine` | No reverse | Read-only. No reverse needed. |
| Get config | `get` | No reverse | Read-only. No reverse needed. |
| Status | `status` | No reverse | Read-only. No reverse needed. |
| Version | `version` | No reverse | Read-only. No reverse needed. |
| Validate engines | `debug validate-engines` | No reverse | Read-only validation. No reverse needed. |
| Prune cache | `prune-cache` | No `restore-cache` / `undo-prune` | Pruning removes components permanently. The only way to "reverse" is to reinstall components (e.g., via `use-engine` which reinstalls missing components). There is no dedicated restore command. |

---

## Potential Naming Asymmetries (without clear pairs)

| Forward Command | Would-be Reverse (if added) | Asymmetry |
|----------------|---------------------------|-----------|
| `use-engine` | `unuse-engine` | `use`/`unuse` is not a standard verb pair. Better: `activate-engine` / `deactivate-engine`, or `enable-engine` / `disable-engine`. |
| `prune-cache` | `restore-cache` | `prune` has no natural antonym in CLI vocabulary. Better: `remove-cache` / `add-cache` or `clean-cache` with no reverse expected. |
| `webui` | `close-webui` | `webui` is a noun-as-command. A proper verb-noun pair would be `open-webui` / `close-webui`, but `close-webui` doesn't map cleanly to closing a browser the CLI doesn't own. |

---

## Behavioral Asymmetries (within existing pairs)

### `set` vs `unset`

| Aspect | `set` | `unset` |
|--------|-------|---------|
| Args | `set <key=value>...` (≥1, repeatable) | `unset <key>` (exactly 1) |
| Hidden flags | `--package`, `--engine` | None |
| Change detection | Checks prev value for any change | Checks if effective value actually changed (after precedence) |
| Restart prompt | Yes (`--assume-yes`, `--no-restart`) | Yes (`--assume-yes`, `--no-restart`) |
| Root required | Yes | Yes |

**Finding**: `set` can handle multiple key=value pairs; `unset` handles exactly one. This is a minor behavioral asymmetry — `unset` could reasonably support multiple keys.

### `chat` vs `debug chat`

| Aspect | `chat` | `debug chat` |
|--------|--------|--------------|
| Base URL source | Auto-detected from active engine | `--base-url` flag (required) |
| Model name | Auto-detected | `--model` flag (optional) |
| Server check | Yes (checks snap service is active) | No |
| Visibility | Visible (conditional) | Hidden under `debug` |

**Finding**: These are not a symmetric pair — they serve different purposes (production vs debugging). However, the naming makes them appear symmetric. A better approach: make `chat` accept optional `--base-url` and `--model` flags, rendering `debug chat` redundant.

### `serve-webui` vs `debug serve-webui`

| Aspect | `serve-webui` (top-level) | `debug serve-webui` |
|--------|--------------------------|---------------------|
| Base URL | Auto-detected from active engine | `--base-url` flag (default: `http://localhost:8080/v1`) |
| Capabilities | `--capabilities` flag (default: `text`) | All capabilities enabled (hardcoded) |
| Host | Configurable via `--host` | Fixed to `localhost` |
| Port | Configurable via `--port` | Configurable via `--port` |
| Instance name | From snap env | Hardcoded `"debug"` |
| Engine name | From cache | Hardcoded `"unset"` |
| Visibility | Hidden | Hidden (under debug) |

**Finding**: Same pattern as chat — two commands with similar names and different contracts. Consolidation would reduce maintenance burden and user confusion.

---

## Summary

| Category | Count |
|----------|-------|
| Present symmetric pairs | 1 (`set`/`unset`) |
| Missing reverse operations (justified) | 12 (read-only or self-terminating) |
| Missing reverse operations (actionable) | 2 (`use-engine` deactivation, `prune-cache` restoration) |
| Naming asymmetries | 0 (the single pair is perfectly named) |
| Behavioral asymmetries in named pairs | 1 (multi-value `set` vs single-value `unset`) |

**Key finding**: The CLI has minimal symmetry issues because most commands are observation-only (no reverse needed) or self-terminating (chat, run, serve). The only actionable gap is the lack of a `deactivate-engine` command. The behavioral asymmetry between `set` and `unset` (multi-value vs single-value) is minor and could be addressed by supporting `unset <key>...`.
