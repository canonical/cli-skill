# Command: /cli-semantic-analysis

Deep semantic and structural analysis of a CLI command set.

## Execution Order

1. Run `shared/cli-discovery-preflight.md`
2. Read all files in `cli-review/0-cli-discovery-preflight/`
3. Run semantic analysis phases and produce structured outputs

## Scope

This command owns the "later steps" previously bundled in `cli-review/SKILL.md`, including:

- Verb-noun decomposition matrix
- Verb taxonomy and aspect classification
- Semantic domain clustering
- Symmetry audit
- Confusion-pair audit
- Pattern classification and recommendations
- Frame-based semantic analysis (where available)

## Standards and Change Safety

Before recommendations that rename/remove commands, read:

- `cli-review/resources/standard.md`
- `cli-review/resources/deprecation.md`

Any rename/remove recommendation must include a deprecation path.

## Required Output Directory

Write outputs to:

- `cli-review/2-cli-semantic-analysis/`

Suggested files:

1. `01-verb-noun-decomposition.md`
2. `02-verb-taxonomy.md`
3. `03-semantic-domain-clustering.md`
4. `04-symmetry-audit.md`
5. `05-confusion-pair-audit.md`
6. `06-pattern-classification.md`
7. `07-frame-analysis.md`
8. `summary.md`
