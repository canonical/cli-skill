# Command Set

This document lists every top-level command and hidden commands discovered in the CLI source under `cli/cmd/cli`.

|command|description|hidden|
|---|---:|---:|
|status|Show the status|false|
|chat|Start the chat CLI|false (conditional on build/runtime)|
|webui|Launch web UI|false (conditional on build/runtime)|
|list-engines|List available engines|false|
|show-engine|Print information about an engine|false|
|use-engine|Select an engine|false|
|get|Print configurations|false|
|set|Set configurations|false|
|unset|Unset configurations|false|
|show-machine|Print information about the host machine|false|
|prune-cache|Remove cached data|false|
|version|Show version information|false|
|run|Run a subprocess|true (hidden)
|serve-webui|Serve static files and configurations (webui) |true (hidden)
|debug|Debug command (hidden)|true (hidden)

Sources: `cli/cmd/cli/main.go` and all files in `cli/cmd/cli/commands/`.
