# 03 — Semantic Domain Clustering

## Domain Assignment (6 commands total)

| Domain | Count | Commands | Naming Consistent? | Notes |
|---|---|---|---|---|
| **Engine** | 2 | `use-engine`, `show-engine` | Yes (both use `-engine` suffix) | Missing list/discovery command. No CRUD completeness — only "select" and "show". |
| **Configuration** | 2 | `get`, `set` | Yes (standard get/set pair) | Missing `unset`/`reset`. No "list all" command. Noun is implicit. |
| **Inference** | 1 | `chat` | N/A (single command) | Primary user-facing command. Bare verb, no noun qualification needed. |
| **Shell Integration** | 1 | `completion` | N/A (single command) | Functions as a noun/utility. Only `bash` shell supported. |

## Self-Check

- Total commands: 6
- Sum of Count column: 2 + 2 + 1 + 1 = **6** ✓

## Domain Analysis

### Engine Domain (2 commands)

The engine domain handles hardware-specific inference backend selection. Two commands cover the core "inspect" and "select" operations.

**CRUD coverage**:
- Create: N/A (engines are bundled with the snap)
- Read: `show-engine` ✓
- Update: `use-engine` ✓ (selects/activates an engine)
- Delete: N/A (engines cannot be removed by users)
- List: ✗ Missing — no way to enumerate available engines

**Verb consistency**: Mixed — `show` is an observation verb, `use` is a non-standard mutation verb. Per DE013, the activation pattern would more consistently use `enable-engine` or even fold into `set engine=cpu`.

### Configuration Domain (2 commands)

The configuration domain follows the standard `get`/`set` pattern mandated by DE013 for configuration access.

**CRUD coverage**:
- Create: `set` (creates if not exists) ✓
- Read: `get` ✓
- Update: `set` (overwrites) ✓
- Delete: ✗ Missing — no `unset` or `reset` to restore defaults
- List: ✗ Missing — no way to show all current configuration

**Verb consistency**: Perfect — uses the canonical `get`/`set` pair.

### Inference Domain (1 command)

Single command that is the primary user workflow — starting an interactive chat session.

**Notes**: This is the "reason the tool exists" command. It's appropriately simple and prominent.

### Shell Integration Domain (1 command)

Utility command for shell completion generation.

**Notes**: `completion` acts as a noun rather than a verb. The pattern `completion <shell>` is common across Go CLI tools (cobra-style). Per strict DE013 grammar, this could be `generate-completion` but the current form is idiomatic in the ecosystem.

## Cross-Domain Observations

1. **No lifecycle commands**: There are no create/destroy/deploy/remove operations. This is appropriate — the snap manages lifecycle via `snap install`/`snap remove`.
2. **No access commands**: No login/logout/auth. Appropriate — the local server has no authentication.
3. **Small, focused command set**: 6 commands covering 4 domains is minimal and easy to learn.
4. **Missing glue**: No `status` command that bridges domains (e.g., "engine: cpu, server: running, components: installed").
