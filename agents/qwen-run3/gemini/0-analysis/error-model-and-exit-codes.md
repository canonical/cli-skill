# Error Model and Exit Codes

While explicit CLI error strings aren't fully outlined in documentation, we can infer standard shell integration from the scripts utilizing the CLI.

- Missing Configuration: Handled as non-fatal with silent fallback (`qwen36 get model-name 2>/dev/null || true` from wrapper `chat.sh`).
- Missing Snap Components: Fails fast if timeout occurs. Server execution has a robust wait phase, dropping with `"Error: timed out after..."` and standard POSIX exit code `1`. It coordinates with `snapctl` to halt systemd crash looping (`snapctl stop qwen36`).
- Standard POSIX Exit Codes are observed on operational failures (`0` strictly translates to success, non-zero reflects an issue reading/writing state).