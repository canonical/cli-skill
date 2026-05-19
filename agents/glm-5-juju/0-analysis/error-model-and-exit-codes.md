# Juju CLI Error Model and Exit Codes

## Overview

Juju uses a structured error handling system that provides clear feedback for operators and enables reliable script error handling. Errors are classified into categories with appropriate exit codes.

## Exit Codes

### Standard Exit Codes

| Code | Name | Meaning | Script Handling |
|------|------|---------|-----------------|
| 0 | Success | Command completed successfully | Continue |
| 1 | General Error | Command failed (any reason) | Check stderr |
| 2 | Initialization Error | Context/command setup failed | Check environment |
| N | Passthrough | Plugin exit code preserved | Plugin-specific |

### Exit Code Implementation

```go
// cmd/cmd/cmd.go
func Main(c Command, ctx *Context, args []string) int {
    // ...
    if err := c.Run(ctx); err != nil {
        if utils.IsRcPassthroughError(err) {
            return err.(*utils.RcPassthroughError).Code
        }
        if err != ErrSilent {
            WriteError(ctx.Stderr, err)
        }
        return 1
    }
    return 0
}
```

## Error Categories

### 1. Initialization Errors (Exit Code 2)

Errors that occur before command execution begins.

| Error | Trigger | Example |
|-------|---------|---------|
| `ErrNoControllersDefined` | No controllers registered | Fresh installation, run `bootstrap` first |
| `ErrNoCurrentController` | Current controller not set | Corrupted config, use `juju switch` |
| `ErrNoModelSpecified` | Model required but not specified | Model-scoped command without `-m` |
| `ErrInvalidArgs` | Invalid command arguments | Wrong number of positional args |
| `ErrInvalidFlag` | Invalid flag value | Unknown flag, wrong type |

**Example output:**
```
ERROR No model in focus.

Please use "juju models" to see models available to you.
You can set current model by running "juju switch"
or specify any other model on the command line using the "-m" option.
```

### 2. Connection Errors (Exit Code 1)

Errors related to controller/model connectivity.

| Error | Condition | Resolution |
|-------|-----------|------------|
| Connection refused | Controller unreachable | Check network, controller status |
| Authentication failed | Invalid credentials | Re-login with `juju login` |
| Certificate error | Invalid/self-signed cert | Check `ca-cert` in controllers.yaml |
| Timeout | Operation timed out | Increase timeout, check network |

**Example output:**
```
ERROR cannot connect to controller "aws-prod": 
    connection refused (10.0.0.1:17070)
```

### 3. Validation Errors (Exit Code 1)

Errors from invalid input or state.

| Error Type | Example |
|------------|---------|
| Invalid application name | `ERROR invalid application name: "my-app!"` |
| Invalid charm URL | `ERROR invalid charm URL: "bad-url"` |
| Invalid constraint | `ERROR invalid constraint value: "mem=abc"` |
| Invalid config key | `ERROR unknown config option: "invalid-key"` |
| Invalid placement | `ERROR invalid placement directive: "xyz"` |

### 4. State Errors (Exit Code 1)

Errors from invalid model state or conflicts.

| Error | Condition |
|-------|-----------|
| Application exists | Deploy duplicate application name |
| Model not empty | Destroy model with applications |
| Unit not found | Operate on non-existent unit |
| Relation not found | Remove non-existent relation |
| Storage in use | Remove attached storage |

**Example output:**
```
ERROR cannot destroy model "production": model has 3 applications
Use --force to override this check, or remove applications first.
```

### 5. Permission Errors (Exit Code 1)

Access control errors.

| Error | Condition |
|-------|-----------|
| Unauthorized | User lacks permission |
| Access denied | Insufficient model access |
| Admin required | Admin-only operation |

**Example output:**
```
ERROR permission denied: user "dev" cannot add units to model "production"
    required access: admin
    current access: read
```

### 6. Operation Errors (Exit Code 1)

Runtime failures during command execution.

| Error | Condition |
|-------|-----------|
| Hook failed | Charm hook execution failure |
| Action failed | Action execution failure |
| Provisioning failed | Machine provisioning failure |
| Upgrade failed | Model upgrade failure |

**Example output:**
```
ERROR unit "postgresql/0" hook "config-changed" failed:
    exit status 1
    see juju debug-log for details
```

## Error Output Format

### Standard Error Format

```
ERROR <message>
```

### Detailed Error Format (with --debug)

```
ERROR <message>
    <file>:<line>: <error detail>
    <file>:<line>: <error detail>
```

**Example:**
```
ERROR cannot deploy application "postgresql"
    github.com/juju/juju/api/client/application/client.go:123: charm not found
    github.com/juju/juju/cmd/juju/application/deploy.go:456: deploy failed
```

### Error Verbosity Levels

| Flag | Output |
|------|--------|
| Default | ERROR line only |
| `--verbose` | ERROR + context hints |
| `--debug` | ERROR + full stack trace |

## Silent Errors

Some errors are suppressed to avoid noise:

```go
var ErrSilent = errors.New("cmd: error out silently")
```

When `ErrSilent` is returned:
- Exit code is 1
- No error message is written to stderr
- Useful for expected failure conditions

## Plugin Exit Codes

Plugins can return arbitrary exit codes:

```go
// Passthrough error for plugins
func NewRcPassthroughError(code int) error {
    return &RcPassthroughError{Code: code}
}
```

Plugin exit codes are preserved:
```bash
juju my-plugin arg1 arg2
echo $?  # Returns plugin's exit code directly
```

## Error Message Patterns

### Actionable Errors

Include guidance on resolution:

```
ERROR cannot destroy model "production": model contains applications
Hint: use --force to destroy, or remove applications with 'juju remove-application'
```

### Contextual Errors

Include relevant identifiers:

```
ERROR cannot remove unit "mysql/2": unit is the leader
Hint: remove all units or use --force to override
```

### Transient Errors

Indicate retry possibility:

```
ERROR connection to controller "aws-prod" timed out
Hint: check network connectivity and controller status, then retry
```

## Error Handling Best Practices

### For Operators

1. Read the full error message
2. Check hints for resolution steps
3. Use `--debug` for detailed diagnostics
4. Check `juju debug-log` for hook failures

### For Scripts

```bash
#!/bin/bash
set -e

# Capture both stdout and stderr
output=$(juju status --format json 2>&1)
exit_code=$?

if [ $exit_code -ne 0 ]; then
    echo "Command failed with exit code $exit_code"
    echo "$output"
    exit $exit_code
fi

# Parse output
echo "$output" | jq -r '.model.name'
```

### Retry Logic

```bash
#!/bin/bash
max_retries=3
retry_delay=5

for i in $(seq 1 $max_retries); do
    if juju status --format json > status.json 2>&1; then
        break
    fi
    
    if [ $i -lt $max_retries ]; then
        echo "Retry $i/$max_retries in ${retry_delay}s..."
        sleep $retry_delay
    else
        echo "Max retries reached"
        exit 1
    fi
done
```

## Specific Error Scenarios

### Bootstrap Failures

```
ERROR bootstrap failed: cannot start controller instance
    -- 
    Instance provision failed:
    - Instance "i-abc123" terminated unexpectedly
    
Hint: check cloud credentials and quotas
```

### Deploy Failures

```
ERROR cannot deploy "postgresql":
    charm "postgresql" not found in channel "stable"
    
Hint: check available channels with 'juju info postgresql'
```

### Integration Failures

```
ERROR cannot add relation between "postgresql:db" and "mysql:db":
    relations between these endpoints are not possible
    (interface mismatch: "pgsql" != "mysql")
```

### Upgrade Failures

```
ERROR model upgrade failed:
    unit "postgresql/0" agent is not in idle state
    current state: executing
    
Hint: wait for all units to become idle before upgrading
```

## Error Recovery

### Controller Connection Issues

```bash
# Check controller status
juju controllers

# Re-login
juju logout
juju login

# Verify connection
juju whoami
```

### Model State Issues

```bash
# Check for blocked units
juju status --color | grep -i error

# Resolve failed units
juju resolved postgresql/0

# Force removal if needed
juju remove-unit postgresql/0 --force
```

### Credential Issues

```bash
# Check credentials
juju credentials

# Add/update credentials
juju add-credential aws

# Update model credential
juju set-credential aws-cred
```

## Error Categories by Command

### Infrastructure Commands

| Command | Common Errors |
|---------|---------------|
| `bootstrap` | Cloud not found, credential invalid, quota exceeded |
| `add-cloud` | Cloud exists, invalid config |
| `add-credential` | Invalid format, cloud not found |

### Application Commands

| Command | Common Errors |
|---------|---------------|
| `deploy` | Charm not found, invalid name, insufficient resources |
| `add-unit` | Application not found, invalid placement |
| `remove-application` | Has relations, storage attached |

### Model Commands

| Command | Common Errors |
|---------|---------------|
| `add-model` | Name conflict, quota exceeded |
| `destroy-model` | Model not empty, has controllers |
| `model-config` | Invalid key, invalid value |

### Integration Commands

| Command | Common Errors |
|---------|---------------|
| `integrate` | Endpoints incompatible, relation exists |
| `remove-relation` | Relation not found, blocked by charm |
