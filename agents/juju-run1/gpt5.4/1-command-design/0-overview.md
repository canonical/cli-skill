# Juju CLI Command Design Overview

## Executive summary

Juju's CLI is broad, capable, and operationally mature. The command host, wrappers, and registration model are solid. The command surface itself is where most of the design debt sits.

The current shape is a large flat namespace of mostly top-level commands with three simultaneous grammars:
- bare verbs for major workflows: `deploy`, `consume`, `offer`, `refresh`, `switch`
- verb-noun pairs for explicit lifecycle operations: `add-model`, `remove-unit`, `show-controller`
- plural noun listings and singleton detail commands: `models`, `users`, `spaces`, `show-model`, `show-user`, `show-space`

That shape works, but it has predictable costs:
- discoverability depends on naming discipline because hierarchy is shallow
- complement pairs are not always complete or lexically aligned
- output and safety semantics drift across families
- usage metadata quality matters a lot because help is carrying the burden of structural clarity

## Key findings

- The source architecture is coherent; the command vocabulary is less so.
- Application, model, controller, cloud, and storage families are the main structural domains.
- The strongest design pattern is `add/remove/show/list/config/grant/revoke` across explicit nouns.
- The weakest pattern is the mix of bare verbs, plural nouns, and verb-noun forms for conceptually similar actions.
- Several commands are easy to confuse because the same operation family is expressed with different lexical strategies.
- The CLI already contains enough symmetry to standardize further without a radical redesign.

## Most important design tensions

### 1. Flatness vs discoverability

Juju exposes nearly everything at the top level. That keeps command typing short, but pushes a lot of burden onto verb choice and help quality.

### 2. Workflow verbs vs explicit nouns

Commands like `deploy`, `offer`, `consume`, `refresh`, and `trust` are short and memorable. Commands like `add-model`, `remove-space`, and `show-controller` are explicit and systematic. Juju uses both styles heavily, which creates lexical friction.

### 3. Human-first ergonomics vs machine-contract consistency

The CLI has meaningful JSON/YAML support, but defaults and schema expectations vary by family. This affects the command-surface design because reporting verbs are not fully normalized.

### 4. Strong safety in destroy flows vs weaker safety elsewhere

Destructive lifecycle commands are carefully designed. Broader mutation commands have less consistent preview, confirmation, and `--force` semantics.

## Recommended design direction

Juju does not need a wholesale move to nested subcommands. The more pragmatic direction is:
- preserve the flat top-level shape
- tighten verb semantics
- normalize complement pairs
- add aliases for migration rather than breaking old commands immediately
- standardize output and safety policies across command families

## Guide to the rest of this section

- `1-verb-noun-matrix.md`: current surface mapped by verbs and nouns
- `2-verb-taxonomy.md`: what each verb family means today, and where meanings collide
- `3-semantic-domains.md`: command clusters by operator mental model
- `4-symmetry-audit.md`: missing complement operations and uneven families
- `5-confusion-pairs.md`: likely operator confusion hotspots
- `6-recommendations.md`: concrete renames, aliases, deprecations, and grouping proposals
