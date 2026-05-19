# Extensibility Model

The extensibility model leverages Canonical's Snap component model for backend execution dynamically, but features a limited surface for typical CLI plugins.

- **Inference Engines**: Exposes support for adding hardware-specific backends without bloating the primary snap heavily. Handled in `snapcraft.yaml` (e.g. `llamacpp`, `llamacpp-cuda`). If a user wishes to use custom extensions, they are confined to what the CLI's `use-engine` permits.
- **Chat Interface**: Decoupled (`go-chat-client`), which technically could be replaced as a binary if a user overrode snap boundaries, but it essentially acts as a fixed endpoint mapping in `qwen36.` No dynamic CLI subcommand registration pathways exist.