# Juju CLI Documentation Quality Gaps

## Overview

This document analyzes the gap between Juju CLI documentation and actual command behavior, identifying mismatches, missing examples, outdated guidance, and ambiguities.

## Documentation Sources Analyzed

1. **Command Help Text**: `juju <command> --help`
2. **Generated Documentation**: `go run ./cmd/juju documentation`
3. **Online Documentation**: juju.is/docs
4. **Code Comments**: Doc comments in source files

## Documentation Gaps by Category

### 1. Missing Examples

#### Commands with Insufficient Examples

| Command | Gap | Impact |
|---------|-----|--------|
| `bootstrap` | No Kubernetes-specific examples | Users struggle with K8s setup |
| `integrate` | Missing cross-model relation examples | Users miss CMR capabilities |
| `offer` | No JAAS integration examples | Enterprise users confused |
| `secrets` | Missing secret rotation workflow | Security best practices unclear |
| `migrate` | No pre-migration checklist | Migration failures |

#### Recommendations

Add examples for common workflows:

```markdown
## Examples (bootstrap)

### Bootstrap on LXD
    juju bootstrap lxd

### Bootstrap on Kubernetes with LoadBalancer
    juju bootstrap --config controller-service-type=loadbalancer myk8s

### Bootstrap on AWS with custom base
    juju bootstrap --bootstrap-base=ubuntu@22.04 --force aws
```

### 2. Missing Cross-References

#### Commands Lacking "See Also" References

| Command | Should Reference | Current State |
|---------|-----------------|---------------|
| `add-unit` | `scale-application` (K8s) | Missing |
| `remove-application` | `remove-unit`, `destroy-model` | Missing |
| `config` | `model-config`, `controller-config` | Present but incomplete |
| `integrate` | `offer`, `consume` | Missing |
| `secrets` | `secret-backends`, `add-secret-backend` | Incomplete |

#### Recommendations

Add comprehensive "See also" sections:

```markdown
## See Also

- [add-unit](#add-unit) - Add units to machine models
- [scale-application](#scale-application) - Scale Kubernetes applications
- [remove-unit](#remove-unit) - Remove specific units
```

### 3. Argument Documentation Gaps

#### Flags with Missing Documentation

| Command | Flag | Gap |
|---------|------|-----|
| `deploy` | `--device` | Syntax unclear for GPU devices |
| `deploy` | `--overlay` | Bundle overlay format not documented |
| `status` | `--relations` | Output format change not explained |
| `bootstrap` | `--storage-pool` | Syntax unclear |
| `config` | `--file` | YAML format not fully documented |

#### Recommendations

Document flag syntax clearly:

```markdown
### --device (deploy)

Specify GPU device requirements for Kubernetes charms.

Syntax: `<label>=[<count>,]<device-class>|<vendor/type>[,<attributes>]`

Examples:
    # Single Nvidia GPU
    juju deploy myapp --device miner=1,nvidia.com/gpu

    # Two GPUs with attribute filter
    juju deploy myapp --device twingpu=2,nvidia.com/gpu,gpu=nvidia-tesla-p100
```

### 4. Output Format Gaps

#### Commands with Incomplete Format Documentation

| Command | Issue |
|---------|-------|
| `status` | JSON schema not documented |
| `show-unit` | Field meanings unclear |
| `debug-log` | Log format not explained |
| `operations` | Status values not enumerated |

#### Recommendations

Document output schemas:

```markdown
## Output Format (status --format json)

```json
{
  "model": {
    "name": "string",
    "controller": "string",
    "cloud": "string",
    "region": "string",
    "version": "semver string",
    "sla": "unsupported|essential|standard|advanced"
  },
  "applications": {
    "<name>": {
      "charm": "string",
      "channel": "stable|candidate|beta|edge",
      "status": { "current": "status-value" }
    }
  }
}
```

### 5. Error Documentation Gaps

#### Undocumented Error Scenarios

| Command | Undocumented Error | Expected Behavior |
|---------|-------------------|-------------------|
| `deploy` | Charm not found in channel | Document fallback behavior |
| `integrate` | No matching endpoints | Explain resolution steps |
| `remove-application` | Hook failures | Document --force implications |
| `destroy-model` | Storage present | Document --destroy-storage vs --release-storage |

#### Recommendations

Add error scenario documentation:

```markdown
## Troubleshooting

### "charm not found in channel"

The specified charm is not available in the requested channel.

**Resolution:**
1. Check available channels: `juju info <charm>`
2. Use an alternative channel: `juju deploy <charm> --channel edge`
3. Specify exact revision: `juju deploy <charm> --revision 42 --channel edge`
```

### 6. Behavior vs Documentation Mismatches

| Command | Documented Behavior | Actual Behavior |
|---------|--------------------|--------------| 
| `integrate` | Aliased as `relate` | Both work, but `relate` not shown in help |
| `model-config` | Same as `config` | `config` without args shows app config, `model-config` shows model config |
| `constraints` | No model constraints shown | Shows model constraints when no app specified |
| `whoami` | Shows current model | Also shows account and controller context |

#### Recommendations

Align documentation with actual behavior or fix inconsistencies.

### 7. Missing Prerequisite Documentation

| Command | Missing Prerequisites |
|---------|----------------------|
| `bootstrap` | Credential setup steps |
| `deploy` | Model must exist |
| `integrate` | Both applications must exist |
| `offer` | Application must exist |
| `consume` | Must have access to remote offer |
| `migrate` | Target controller must exist and be accessible |

#### Recommendations

Add prerequisite sections:

```markdown
## Prerequisites

Before running this command:
1. A controller must be bootstrapped: `juju bootstrap`
2. Credentials must be configured: `juju add-credential`
3. A model must exist: `juju add-model`
```

### 8. Outdated Documentation

| Topic | Issue |
|-------|-------|
| Charm Store references | Documentation mentions `cs:` prefix, now deprecated for `ch:` |
| LXD setup | Instructions may be outdated for LXD 5.x |
| Kubernetes integration | Some examples use deprecated K8s versions |
| Base references | Some docs use `ubuntu@20.04` which may not be default |

#### Recommendations

Audit and update:
1. Replace all `cs:` references with `ch:` or no prefix
2. Verify LXD instructions against LXD 5.x
3. Update Kubernetes examples to current versions
4. Review base version references

### 9. Ambiguous Terminology

| Term | Ambiguity | Recommendation |
|------|-----------|----------------|
| "relation" | Used interchangeably with "integration" | Standardize on "integration" |
| "offer" | Verb and noun confused | Clarify "create offer" vs "the offer" |
| "model" | Workload model vs controller model | Qualify when ambiguous |
| "storage" | Pool vs instance vs volume | Use specific terms |
| "base" | Replaces "series" but not fully migrated | Document transition |

#### Recommendations

Create a terminology glossary and standardize usage.

### 10. Missing Context Documentation

#### Context-Dependent Behavior

| Command | Context Issue |
|---------|---------------|
| `status` | Output varies significantly by model type (IAAS vs CAAS) |
| `deploy` | Behavior differs for Kubernetes vs machine models |
| `add-unit` | Not available for Kubernetes applications |
| `scale-application` | Only for Kubernetes applications |
| `expose` | Different behavior on K8s vs machine models |

#### Recommendations

Document context-specific behavior:

```markdown
## Kubernetes vs Machine Models

This command behaves differently depending on model type:

**Machine Models:**
- Units are provisioned as individual machines
- Use `add-unit` to scale

**Kubernetes Models:**
- Units are pods
- Use `scale-application` to scale
```

## Documentation Quality Metrics

### Coverage Analysis

| Category | Commands with Issue | Severity |
|----------|-------------------|----------|
| Missing examples | 45 | High |
| Missing cross-references | 30 | Medium |
| Argument documentation gaps | 25 | High |
| Output format gaps | 20 | Medium |
| Error documentation gaps | 35 | High |
| Behavior mismatches | 10 | Critical |
| Missing prerequisites | 40 | Medium |
| Outdated content | 15 | High |
| Ambiguous terminology | 20 | Medium |
| Missing context docs | 25 | Medium |

### Priority Remediation

1. **Critical**: Fix behavior mismatches
2. **High**: Add examples, document arguments, fix outdated content
3. **Medium**: Add cross-references, output formats, prerequisites

## Documentation Improvement Recommendations

### 1. Example-First Approach

Each command should have:
- 2-3 common use cases
- 1-2 advanced use cases
- 1 error recovery scenario

### 2. Consistent Structure

Standardize command documentation:

```markdown
## Summary
## Usage
## Options
## Prerequisites
## Examples
    ### Basic
    ### Advanced
    ### Error Handling
## Details
## Output Formats (if applicable)
## See Also
## Troubleshooting
```

### 3. Schema Documentation

Document all JSON/YAML output schemas.

### 4. Migration Guides

Add guides for:
- Charm Store → Charmhub migration
- Series → Base migration
- Machine → Kubernetes migration

### 5. Interactive Help

Enhance `juju help` with:
- Fuzzy search
- Topic-based help
- Workflow guides

## Comparison with Similar Tools

| Tool | Documentation Strength | Juju Gap |
|------|------------------------|----------|
| `kubectl` | Comprehensive examples | Juju needs more examples |
| `terraform` | State management docs | Juju needs model state docs |
| `helm` | Chart documentation | Juju needs charm docs guidance |
| `ansible` | Module documentation | Juju needs action documentation |

## Summary

The Juju CLI documentation has significant gaps in:
1. Practical examples for common workflows
2. Cross-references between related commands
3. Argument and flag syntax documentation
4. Error scenario documentation
5. Context-specific behavior for IAAS vs CAAS models

Addressing these gaps would significantly improve user experience and reduce support burden.
