# Juju CLI Command Set Shape Analysis

Generated from `0-analysis/commandset.md` and `0-analysis/argument-structure.md`.

---

## Shape Summary

Juju's CLI is a large, mature, flat command surface with roughly 150 top-level commands. The dominant pattern is `verb-noun` hyphenation (`add-model`, `remove-unit`, `show-cloud`), but substantial noise comes from legacy noun-as-verb shorthands (`config`, `status`, `machines`), bare verbs (`deploy`, `bootstrap`, `refresh`), and a small set of noun-first configuration commands (`model-config`, `controller-config`). The command set has grown organically over many years, absorbing domain-specific vocabulary (charms, bundles, SAAS, CMR) alongside generic resource-management verbs. The result is a mostly consistent CRUD surface with notable asymmetries, orphan commands, and verb collisions that increase cognitive load for new users.

---

---

## Sections

- [Verb-Noun Decomposition Matrix](1-verb-noun-matrix.md)
- [Verb Taxonomy and Aspect Classification](2-verb-taxonomy.md)
- [Semantic Domain Clustering](3-semantic-domains.md)
- [Symmetry Audit](4-symmetry-audit.md)
- [Confusion-Pair Audit](5-confusion-pairs.md)
- [Pattern Classification and Recommendations](6-recommendations.md)
