# Error Model and Exit Codes

Observed behavior:
- The CLI generally returns exit code `1` on error via `cobra` command propagation (`RunE` errors cause exit 1 in `main`). Specific error codes are not abundantly used.
- `use-engine`, `prune-cache`, `set`, `unset` require root; they return permission errors (from `common.ErrPermissionDenied`) which map to generic non-zero exit.
- The deprecation policy in `deprecation.md` prescribes `exit code 2` for removed commands in a major version; this CLI does not currently implement that specific signaling.

Error categories:
- Configuration and permission errors: missing privileges, missing keys (set/get/unset)
- Missing resources: no active engine (`common.ErrNoActiveEngine`), engine manifest not found
- Component management errors: component install failures and timeouts (snapctl errors)
- Validation errors: malformed args (e.g., `set` without key=value)
- Runtime errors: subprocess `run` returns the underlying program's exit code

Recommendations:
- Document exit codes explicitly in `--help` or `README` and adopt `2` for removed/renamed commands per `deprecation.md` when implementing deprecation flows.
- When returning known error categories, use mapped exit codes (e.g., 3 = permission, 4 = resource-not-found) for easier scripting.
