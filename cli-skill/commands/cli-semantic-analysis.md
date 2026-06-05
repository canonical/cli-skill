# Command: /cli-semantic-analysis

Deep semantic and structural analysis of a CLI command set.
Use this workflow when the user says `discuss-commandset` or asks to review the command set shape, hierarchy, or naming patterns.

## Execution Order

1. Run `shared/cli-discovery-preflight.md`
2. Read all files in `cli-review/0-cli-discovery-preflight/`
3. Run semantic analysis phases and produce structured outputs

#### Required Inputs

- The CLI name and command list (from `cli-review/0-cli-discovery-preflight/commandset.md` or user-provided)
- Optionally: ecosystem comparisons (similar tools to benchmark against)

#### Analysis Dimensions

Evaluate the command set across these dimensions:

1. **Grouping and hierarchy**
	- What pattern is used? (noun-verb, verb-noun, flat, nested)
	- Is nesting depth consistent and appropriate?
	- Are related commands grouped logically?

2. **Naming consistency**
	- Do command names follow a single convention (e.g., all verbs, all noun-verb pairs)?
	- Are abbreviations used consistently or avoided?
	- Are there synonyms or near-duplicates that create confusion?

3. **Discoverability**
	- Can a new user predict where to find functionality?
	- Does the hierarchy guide exploration (e.g., `help`, tab completion)?
	- Are there orphan commands that don't fit the mental model?

4. **Ecosystem alignment**
	- How does the structure compare to similar CLI tools?
	- Does it match conventions users already know?

#### Output Directory

Create a `cli-review/3-discuss-commandset/` directory. Each section below produces its own numbered output file. When a section contains large tables (roughly >15 rows or >5 columns), also produce an `.html` version using clean typography: `Ubuntu Sans` via Google Fonts `@import`, falling back to `Arial, sans-serif`. Use dark headers (`#2b2b2b`), alternating row striping, sticky `<th>`, 14px base / 13px tables, `max-width: 1200–1400px`.

You **must** use these exact filenames:

```
cli-review/3-discuss-commandset/01-verb-noun-decomposition.md  (+.html +.json)
cli-review/3-discuss-commandset/02-verb-taxonomy.md             (+.html +.json)
cli-review/3-discuss-commandset/03-semantic-domain-clustering.md (+.html +.json)
cli-review/3-discuss-commandset/04-symmetry-audit.md             (+.html +.json)
cli-review/3-discuss-commandset/05-confusion-pair-audit.md       (+.html +.json)
cli-review/3-discuss-commandset/06-pattern-classification.md     (+.html +.json)
```

For every `.md` file, also produce a `.json` file with the same base name containing the table data as JSON arrays/objects. Use the helper script `mdtables_to_json.py` or write the JSON manually.

##### Section 1 → `01-verb-noun-decomposition.md`: Verb-Noun Decomposition Matrix

Decompose **every** command into a verb and a noun (e.g., `add-cloud` → `add` × `cloud`). The decomposition table must have one row per command — no command may be omitted.

Render as a grid:
- Rows = verbs (sorted alphabetically)
- Columns = nouns/resource types (sorted by frequency)
- Cells = `✓` if the command exists, `—` if the combination is absent

After the grid, annotate:
- **Incomplete CRUD sets**: nouns missing expected lifecycle verbs (e.g., has `add-` but no `remove-`)
- **Verb inconsistencies**: nouns using different verbs for equivalent operations (e.g., `destroy-controller` vs `remove-application`)
- **Orphan commands**: commands that do not decompose cleanly into verb-noun (e.g., `bootstrap`, `integrate`, `resolved`, `whoami`)

##### Section 2 → `02-verb-taxonomy.md`: Verb Taxonomy and Aspect Classification

Classify **every** unique verb from the matrix into the following table. Every verb that appears in Section 1 must appear here — verify by comparing verb lists before finishing:

| Verb | Intent Group | Aspect | Reversible | Paired Verb | CLI Examples |
|---|---|---|---|---|---|

Intent groups:
- **lifecycle**: create, add, deploy, remove, destroy, kill
- **mutation**: update, refresh, upgrade, config, set, bind
- **access**: grant, revoke, enable, disable, login, logout, register, unregister
- **observation**: show, list (plural noun commands), status, log, find, info
- **transfer**: attach, detach, expose, unexpose, consume, offer, integrate
- **execution**: run, exec, cancel, resolved, retry
- **migration**: migrate, import, export, download, scp, sync

Aspect (Aktionsart) values:
- **telic**: action has a natural endpoint (create, destroy, deploy)
- **atelic**: action is ongoing or continuous (run, debug, monitor)
- **punctual**: instant state change (switch, login, trust)

Reversibility values:
- **yes**: paired with a named inverse (add/remove, expose/unexpose)
- **no**: no inverse operation (destroy, kill, bootstrap)
- **partial**: can be undone but not via a single symmetric command (deploy → remove-application)

##### Section 3 → `03-semantic-domain-clustering.md`: Semantic Domain Clustering

Group **all** commands by the resource domain they operate on. Every command must appear in exactly one domain. After building the table, sum the Count column and verify it equals the total command count:

| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---|---|---|---|

Domains include: cloud, model, controller, application, unit, machine, user, secret, storage, space/network, relation/integration, credential, offer/SAAS, charm/resource.

For each domain, note:
- Whether all commands use the same noun form
- Whether the CRUD coverage is complete
- Whether the verb choices are consistent within the domain

##### Section 4 → `04-symmetry-audit.md`: Symmetry Audit

For **every** pair of symmetric operations (including missing reverse operations), list them side by side. Do not limit to a representative sample — list all pairs:

| Operation | Forward Command | Reverse Command | Naming Symmetric? | Behavior Symmetric? | Notes |
|---|---|---|---|---|---|

Check:
- `add-*` / `remove-*`
- `enable-*` / `disable-*`
- `expose` / `unexpose`
- `grant-*` / `revoke-*`
- `suspend-relation` / `resume-relation`
- `attach-*` / `detach-*`
- `register` / `unregister`
- `login` / `logout`
- `consume` / `remove-saas`
- `deploy` / `remove-application`
- `destroy-*` (does it have a creation counterpart?)

Flag:
- Missing reverse operations
- Naming asymmetries (e.g., `destroy-controller` is not reversed by `add-controller`)
- Behavioral asymmetries (e.g., reverse operation requires `--force` but forward does not)

##### Section 5 → `05-confusion-pair-audit.md`: Confusion-Pair Audit

List **all** command pairs that share semantic overlap and risk user confusion. Err on the side of inclusion — it is better to list a low-risk pair than to miss a real confusion source.

| Command A | Command B | Overlap Type | Confusion Risk | Disambiguation |
|---|---|---|---|---|

Overlap types:
- **synonym verbs**: different verbs, same operation (e.g., `remove-*` vs `destroy-*`)
- **scope ambiguity**: same verb, unclear which scope applies (e.g., `config` vs `model-config` vs `controller-config`)
- **functional overlap**: different commands that achieve similar outcomes (e.g., `exec` vs `run`)
- **naming collision**: names too similar, different purposes (e.g., `resources` vs `charm-resources`)

For each pair, rate confusion risk as `high`, `medium`, or `low` and provide a one-sentence disambiguation.

##### Section 6 → `06-pattern-classification.md`: Pattern Classification and Recommendations

- **Pattern classification**: primary grouping pattern, depth, and style
- **Discoverability assessment**: predicted user paths vs actual command locations
- **Ecosystem comparison**: how structure compares to 2-3 similar tools
- **Recommendations**: ordered list of structural improvements with rationale, each annotated with backward compat and migration cost

#### Build Order

Generate sections in this order (each builds on the previous):

1. Verb-Noun Decomposition Matrix
2. Verb Taxonomy and Aspect Classification
3. Semantic Domain Clustering
4. Symmetry Audit
5. Confusion-Pair Audit (uses insights from all above)
6. Pattern Classification and Recommendations

#### Response Format

1. Shape summary — one paragraph describing the current pattern and its strengths
2. Key findings — top 5 issues surfaced by the analysis sections
3. Recommendations — concrete structural changes, ordered by impact
4. Tradeoffs — for each recommendation, note backward compat and migration cost
