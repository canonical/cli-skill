# Confusion Pair Audit

- **set** vs **use-engine**: Modifying system behavior handles target domains differently. `set` covers implicit config vars, while `use-engine` uses verb-noun for `engine` but acts identically to a configuration setter. They essentially overlap in mutate-state operations.