# Output Contracts

### Human-Readable Output
- `use-engine`: Success/failure status.
- `chat`: Interactive streaming responses in the terminal (powered by `go-chat-client`).
- `get` / `set`: Prints raw extracted values or confirmation dialogs.

### Machine-Readable Output
- `show-engine`: Produces structured output designed for programmatic parsing. Inferred through internal shell scripts which pipe `qwen36 show-engine | yq .components[]` and `yq .name` to inspect properties. This implies a rigid stability contract for the emitted JSON/YAML format.