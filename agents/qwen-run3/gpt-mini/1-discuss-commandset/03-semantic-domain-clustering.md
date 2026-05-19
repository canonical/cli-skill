# 03 — Semantic Domain Clustering

Clusters discovered (based on `main.go` groups):

- Basic Commands
  - `status`, `chat`, `webui`
- Configuration Commands
  - `get`, `set`, `unset`
- Engine Management
  - `list-engines`, `show-engine`, `use-engine`
- Administrative / Info
  - `show-machine`, `prune-cache`, `version`
- Hidden / Advanced
  - `run`, `serve-webui`, `debug`

Observations:
- Clustering is consistent with `main.go` grouping: `basic`, `config`, `engine`, and additional commands.
- Engine management is cohesive; configuration commands form a clear CRUD trio.
