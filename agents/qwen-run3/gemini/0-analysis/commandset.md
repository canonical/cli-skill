# Command Set

The CLI command hierarchy is mostly flat, focused on operation and configuration, without deep nested subcommand layers.

- `use-engine`
  - Reconfigures the active inference engine. Uses an auto-detection mode (`--auto`) or manual target (`cpu`, `cuda`).
  - Driven by central CLI logic which manipulates snap configuration state to reboot underlying daemon services.
- `show-engine`
  - Emits the currently active inference engine details (outputs metadata formatted in JSON/YAML).
- `get`
  - Retrieves snap-managed configuration values (e.g. `http.port`, `http.base-path`, `model-name`).
- `set`
  - Saves configuration values (e.g., `http.port=8326`).
- `chat`
  - Starts an interactive terminal chat session.
  - Acts as a wrapper (`chat.sh`) that provisions `OPENAI_BASE_URL` and `MODEL_NAME` relative to configuration, executing the bundled `go-chat-client`.