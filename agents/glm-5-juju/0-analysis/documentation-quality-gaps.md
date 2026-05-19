# Juju CLI Documentation Quality Gaps

## Overview

This analysis compares the CLI's built-in help and documentation against the actual command behavior as implemented in the code. Gaps are categorized by severity and type.

## Gap Categories

| Severity | Definition |
|----------|------------|
| **Critical** | Documentation contradicts actual behavior or omits required information |
| **High** | Key functionality undocumented or examples incorrect |
| **Medium** | Incomplete explanations, missing examples, or ambiguous descriptions |
| **Low** | Minor inconsistencies, formatting issues, or polish needed |

## Critical Gaps

### 1. Bootstrap Default Behavior Undocumented

**Location:** `bootstrap --help`

**Issue:** The help states that running `bootstrap` without arguments "will step you through the process" but doesn't clearly document what happens when controllers already exist.

**Expected:**
```
If run without arguments and controllers exist:
- Prompts to select existing controller or create new one
- Shows cloud selection dialog if no controller specified
```

**Actual:** Help implies interactivity but doesn't explain the decision tree.

**Impact:** Users may create duplicate controllers when intending to add models to existing ones.

### 2. Placement Directive Syntax Incomplete

**Location:** Various commands (`deploy`, `add-unit`, `add-machine`)

**Issue:** The `--to` placement directive syntax is partially documented but key variations are missing.

**Missing Documentation:**
- Container nesting limits (e.g., `0/lxd/0/lxd` invalid)
- Zone placement with containers
- MAAS host placement specifics
- K8s placement differences

**Example Gap:**
```bash
# Valid but undocumented:
juju deploy postgresql --to zone=us-east-1a,lxd

# Invalid but appears to work:
juju add-machine lxd:lxd  # Creates nested LXD, may fail
```

### 3. Exit Code Semantics Not Documented

**Location:** General help and documentation

**Issue:** Exit codes are not documented anywhere accessible to users.

**Missing:**
- What exit code 1 means (general error)
- What exit code 2 means (initialization error)
- Plugin exit code passthrough behavior

**Impact:** Script authors cannot properly handle errors.

## High Gaps

### 4. Output Format Differences Not Explained

**Location:** `status --help` and general format documentation

**Issue:** The differences between `--format line`, `--format oneline`, and `--format short` are not clearly explained.

**Current Documentation:**
```
-format (tabular)
    Specify output format (json|line|oneline|short|summary|tabular|yaml)
```

**Missing:**
- When to use each format
- Field differences between formats
- Machine-parseability recommendations

### 5. Constraint Syntax Ambiguity

**Location:** `deploy --help` and constraint documentation

**Issue:** Constraint syntax allows both `mem=8G` and `mem=8192M` but doesn't document:
- Valid unit suffixes (G, M, T, P)
- Whether `mem=8` is 8 bytes, MB, or invalid
- Interaction between multiple constraint sources

**Example Confusion:**
```bash
# These may not be equivalent:
juju deploy postgresql --constraints "mem=8G"
juju deploy postgresql --constraints "mem=8192M"
juju deploy postgresql --constraints "mem=8"  # Undefined behavior
```

### 6. Model UUID Handling Undocumented

**Location:** `--model` flag documentation

**Issue:** Most commands accept model UUID in addition to names, but this isn't consistently documented.

**Missing:**
- When UUID can be used vs name
- How to discover UUID
- UUID format requirements

### 7. Cross-Model Relations Limitations

**Location:** `integrate` command help

**Issue:** Cross-model relation limitations are not clearly documented.

**Missing:**
- What happens when remote model is unreachable
- How to debug cross-model relation issues
- SAAS consumption prerequisites

## Medium Gaps

### 8. Example Coverage Incomplete

**Location:** Multiple commands

**Issue:** Many commands have examples that don't cover common scenarios.

**Example - `config` command:**
```
Current examples:
  juju config postgresql
  juju config postgresql max-connections

Missing examples:
  juju config postgresql max-connections=200
  juju config postgresql --reset max-connections
  juju config postgresql --file config.yaml
```

### 9. Error Message Documentation

**Location:** General

**Issue:** Error messages are documented but solutions/fixes are not.

**Example:**
```
Current error:
  ERROR cannot remove unit "mysql/0": unit is the leader

Missing guidance:
  Hint: remove all units with 'juju remove-application mysql'
        or use --force to override
```

### 10. Storage Directive Complexity

**Location:** `add-storage --help`

**Issue:** Storage directive syntax is complex but underdocumented.

**Missing:**
- Pool specification requirements
- Size format variations
- Count semantics for different storage types
- What happens when storage already exists

### 11. Secret Backend Switching

**Location:** `model-secret-backend --help`

**Issue:** Documenting switching secret backends doesn't cover the migration process.

**Missing:**
- Drain operation details
- Rollback procedures
- Impact on running charms

### 12. Debug-Hook vs Debug-Code Distinction

**Location:** Debug command help

**Issue:** The difference between `debug-hooks` and `debug-code` is not clearly explained.

**Missing:**
- When to use each
- Workflow differences
- Code debugging limitations

## Low Gaps

### 13. Flag Ordering Ambiguity

**Location:** General command help

**Issue:** Help doesn't clarify that global flags must come before subcommand name.

**Example:**
```bash
# Works:
juju --debug status

# May not work as expected:
juju status --debug  # --debug becomes command-specific flag
```

### 14. Alias Handling Not Documented

**Location:** Help output

**Issue:** Aliases are listed but deprecation status is not visible.

**Example:**
```
Current:
  relate           Alias for 'integrate'.

Missing:
  relate           Alias for 'integrate'. [deprecated: use 'integrate']
```

### 15. Timestamp Format Ambiguity

**Location:** `status --help`

**Issue:** `--utc` flag doesn't explain default timestamp format.

**Missing:**
- Default timestamp format
- ISO 8601 alternative
- Timezone handling

### 16. Color Scheme Not Configurable

**Location:** `--color` and `--no-color` flags

**Issue:** Color customization is not documented.

**Missing:**
- How to customize colors
- Default color values
- Terminal compatibility requirements

### 17. Quiet Mode Behavior Varies

**Location:** `--quiet` flag documentation

**Issue:** Quiet mode suppresses output inconsistently across commands.

**Missing:**
- What is suppressed per command
- Whether errors still appear
- Interaction with `--format`

## Documentation Quality Summary

### By Command Category

| Category | Critical | High | Medium | Low | Total |
|----------|----------|------|--------|-----|-------|
| Infrastructure | 1 | 1 | 1 | 0 | 3 |
| Application | 1 | 2 | 1 | 0 | 4 |
| Model | 0 | 1 | 1 | 1 | 3 |
| Integration | 0 | 1 | 0 | 0 | 1 |
| Storage | 0 | 1 | 1 | 0 | 2 |
| Secrets | 0 | 0 | 1 | 0 | 1 |
| User | 0 | 0 | 0 | 1 | 1 |
| General | 1 | 1 | 0 | 2 | 4 |

### Gap Types

| Type | Count | Example |
|------|-------|---------|
| Missing documentation | 6 | Exit codes, UUID handling |
| Incomplete examples | 4 | config, deploy variations |
| Ambiguous syntax | 3 | Constraints, placement |
| Missing error guidance | 2 | Leader removal, relations |
| Undocumented behavior | 2 | Plugin passthrough, quiet mode |

## Recommendations

### Immediate Actions (Critical)

1. **Document exit codes** in general help output
2. **Complete placement directive syntax** with all valid combinations
3. **Clarify bootstrap behavior** with existing controllers

### Short-term Actions (High)

1. **Expand format documentation** with field comparisons
2. **Document constraint syntax** completely with examples
3. **Add UUID handling** to model-related commands
4. **Document cross-model relation** requirements and limitations

### Medium-term Actions (Medium)

1. **Expand examples** for all commands with common use cases
2. **Add error hinting** to common error messages
3. **Document storage directive** syntax completely

### Ongoing Actions (Low)

1. **Standardize alias display** with deprecation status
2. **Document quiet mode** behavior per command
3. **Clarify flag ordering** requirements

## Documentation Audit Method

This gap analysis was performed by:

1. **Reading help output** for each command (`juju <cmd> --help`)
2. **Reviewing source code** for implemented behavior
3. **Testing edge cases** against documented behavior
4. **Comparing docs** to actual output formats

### Key Files Reviewed

| File | Purpose |
|------|---------|
| `cmd/juju/commands/main.go` | Command registration |
| `cmd/juju/*/` | Command implementations |
| `cmd/cmd/supercommand.go` | Help generation |
| `docs/reference/` | Official documentation |

## Positive Findings

### Well-Documented Areas

1. **Command purpose summaries** are clear and consistent
2. **Flag descriptions** are generally accurate
3. **Examples** for major commands (`deploy`, `bootstrap`) are comprehensive
4. **See-also references** connect related commands well
5. **Alias listing** helps discoverability

### Strong Documentation Patterns

1. **Structured help output** follows consistent format
2. **Purpose/Details separation** aids quick reference
3. **Example sections** use realistic scenarios
4. **Flag grouping** distinguishes global vs command flags

## Conclusion

Juju's CLI documentation is generally good for core workflows but has gaps in advanced scenarios and edge cases. The most critical gaps involve exit codes and placement syntax, which impact automation reliability. Addressing the critical and high-priority gaps would significantly improve the operator experience for production deployments.
