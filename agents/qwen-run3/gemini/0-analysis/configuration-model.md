# Configuration Model

The system uses an underlying snap data store for configuration, which the CLI reads from and writes to. There is a direct translation between CLI settings and runtime env-vars for components.

### Sources and Precedence
- **Snap Config**: The primary source of truth, manipulated via `qwen36 get` and `qwen36 set`.
- **Environment Variables**: Dynamically derived at runtime using wrapper scripts (`chat.sh` pulls state dynamically to synthesize `OPENAI_BASE_URL`).
- **Defaults**: Hardcoded defaults live in execution wrappers (e.g., `api_base_path` defaults to `v1` if unset).

### Key Configurations
- `http.port`
- `model-name`
- `http.base-path`