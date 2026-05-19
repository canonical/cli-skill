# 02 — Verb Taxonomy

List of verbs used and classification:

|verb|count|classification|notes|
|---|---:|---|---|
|status|1|informational|shows overall state|
|chat|1|interactive|starts chat client|
|webui|1|interactive|opens browser UI|
|list|1|read|lists engines|
|show|2|read|show-engine, show-machine (details)|
|use|1|mutating|selects engine, installs components|
|get|1|read|configuration read|
|set|1|mutating|configuration write|
|unset|1|mutating|configuration delete|
|prune|1|mutating/destructive|removes components|
|version|1|informational|shows versions|
|run|1|execution|runs subprocess in engine env|
|serve|1|server|serves web static files|
|debug|1|diagnostic|debug helper|

Observations:
- The verbs are consistent with DE013: most commands are verbs and map to single responsibilities.
- `show` is reused for both `engine` and `machine` which is consistent with the standard (show/info semantics).
- `prune` is a destructive verb; code prompts for confirmation.
