# 01 — Verb-Noun Decomposition

Decomposition of every command into `verb × noun` where possible.

|command|verb|noun|
|---|---:|---|
|status|status|(tool)|
|chat|chat|(tool)
|webui|webui|(tool)
|list-engines|list|engines
|show-engine|show|engine
|use-engine|use|engine
|get|get|configuration
|set|set|configuration
|unset|unset|configuration
|show-machine|show|machine
|prune-cache|prune|cache
|version|version|(tool)
|run|run|subprocess
|serve-webui|serve|webui
|debug|debug|(tool)

Notes:
- CRUD completeness: `engines` has `list`, `show`, `use` (covers discovery and selection but lacks explicit `install`/`remove` commands since component management is done via snap components). `configuration` has `get/set/unset` covering read/write/delete.
- Orphans: `status`, `chat`, `webui`, `version`, `debug` do not map to resource CRUD and are top-level verbs representing actions on the tool itself.
