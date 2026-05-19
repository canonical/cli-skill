# Documentation Quality Gaps

Based on reconstructing behavior from the README and internal wrappers:
- **Available Configuration Matrix Missing**: The CLI exposes `get`/`set` but doesn't map all allowable keys in the primary user documentation. Mentioned keys like `http.base-path` or string schemas meant for `model-name` lack explicit documentation contexts.
- **Error Condition Opacity**: What happens when `use-engine --auto` detects no compatible GPU and CPU memory is insufficient? Recovery paths for failed hardware evaluation are unclear.
- **Detailed Flag Reference**: There is no exhaustive `--help` manual output text listed showing advanced `go-chat-client` parameter pass-through capabilities or hidden subroutines.