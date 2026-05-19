# Symmetry Audit

| domain | expected symmetry | actual pairs | gap |
|--------|-------------------|--------------|-----|
| Config | get / set / unset | get / set | `unset` is missing. |
| Engine | show / update / set | show-engine / use-engine | Lacks consistent verbs. `set-engine` vs `show-engine` would be symmetric. |