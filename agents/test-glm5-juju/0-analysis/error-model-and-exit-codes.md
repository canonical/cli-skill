# Juju CLI Error Model and Exit Codes

## Overview

The Juju CLI uses a structured error handling approach with specific exit codes and error message patterns. This document maps error categories, exit codes, and provides guidance for error handling in scripts.

## Exit Codes

### Standard Exit Codes

| Code | Meaning | Source |
|------|---------|--------|
| 0 | Success | Command completed successfully |
| 1 | Command error | Runtime error, displayed to user |
| 2 | Initialization error | Parse error, flag error, config error |
| N | Passthrough | Plugin exit code from `RcPassthroughError` |

### Exit Code Usage in Code

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

### Error Types

| Type | Exit Code | Message Shown | Use Case |
|------|-----------|---------------|----------|
| `ErrSilent` | 1 | No | Error already displayed |
| `RcPassthroughError` | N | No | Plugin exit code |
| `UnrecognizedCommand` | 2 | Yes | Unknown command |
| `ErrHelp` | 0 | Help text | User requested help |

## Error Categories

### 1. Initialization Errors (Exit 2)

#### Flag Parsing Errors

```
ERROR flag provided but not defined: --unknown-flag
ERROR bad flag syntax: -
ERROR invalid argument "abc" for "--num-units" flag: parse error
```

#### Missing Required Arguments

```
ERROR no charm name specified
ERROR model name is required
ERROR missing <application> argument
```

#### Configuration Errors

```
ERROR cannot read config file: open /path/to/config.yaml: no such file or directory
ERROR invalid YAML in config file: yaml: line 5: could not find expected ':'
```

### 2. Connection Errors (Exit 1)

#### Controller Connection

```
ERROR cannot connect to controller: controller "myctl" not found
ERROR cannot connect to controller: connection refused
ERROR cannot connect to controller: certificate signed by unknown authority
```

#### Authentication

```
ERROR cannot log in: invalid credentials
ERROR cannot log in: user "bob" not found
ERROR cannot log in: user "bob" is disabled
```

#### API Timeout

```
ERROR API connection timed out after 30s
ERROR controller not responding
```

### 3. Resource Errors (Exit 1)

#### Not Found

```
ERROR application "mysql" not found
ERROR model "my-model" not found
ERROR controller "myctl" not found
ERROR unit "postgresql/5" not found
ERROR machine "10" not found
ERROR charm "mysql" not found
```

#### Already Exists

```
ERROR application "mysql" already exists
ERROR model "my-model" already exists
ERROR user "bob" already exists
```

#### Invalid State

```
ERROR model "my-model" is currently busy
ERROR application "mysql" has hook failures
ERROR unit "mysql/0" is in error state
```

### 4. Validation Errors (Exit 1)

#### Invalid Names

```
ERROR invalid model name "MyModel": must be lowercase letters, digits, and hyphens
ERROR invalid application name "my_app": underscores not allowed
ERROR invalid controller name "-myctl": cannot start with hyphen
```

#### Invalid Values

```
ERROR invalid constraint "mem=8GB": value must be a number with optional suffix
ERROR invalid base "ubuntu@12.10": not supported
ERROR invalid channel "invalid": must be stable, candidate, beta, or edge
```

#### Schema Violations

```
ERROR config key "invalid-key" not defined by charm
ERROR unknown configuration option: ssl-cert
ERROR charm does not support action "invalid-action"
```

### 5. Permission Errors (Exit 1)

```
ERROR permission denied: user "bob" cannot access model "admin/my-model"
ERROR access denied: user "bob" does not have admin access
ERROR operation not allowed: requires admin permission
```

### 6. Dependency Errors (Exit 1)

#### Relations

```
ERROR cannot remove application "mysql": has active relations
ERROR relation between "wordpress:db" and "mysql:server" already exists
ERROR no compatible endpoints between "wordpress" and "postgresql"
```

#### Units/Machines

```
ERROR cannot remove machine "0": has running units
ERROR cannot remove unit "mysql/0": has attached storage
```

### 7. Cloud/Provider Errors (Exit 1)

```
ERROR cannot validate cloud "aws": credentials not found
ERROR cannot launch instance: quota exceeded
ERROR cannot provision machine: no matching instance type
ERROR cloud "aws" does not support region "invalid-region"
```

### 8. Charm Errors (Exit 1)

```
ERROR charm "mysql" not found in charmhub
ERROR charm "mysql" revision 99 not found in channel "stable"
ERROR invalid charm: metadata.yaml missing required field "name"
ERROR cannot download charm: network error
```

### 9. Bootstrap Errors (Exit 1)

```
ERROR bootstrap failed: cannot allocate public IP
ERROR bootstrap failed: timeout waiting for agent to start
ERROR bootstrap failed: cannot SSH to bootstrap machine
```

### 10. Upgrade Errors (Exit 1)

```
ERROR cannot upgrade model: already at latest version
ERROR cannot upgrade controller: model "my-model" not healthy
ERROR upgrade failed: agent binary not found
```

## Error Message Patterns

### Prefix Patterns

| Prefix | Category | Example |
|--------|----------|---------|
| `ERROR cannot` | Operation failure | `ERROR cannot deploy application` |
| `ERROR invalid` | Validation failure | `ERROR invalid model name` |
| `ERROR missing` | Required argument | `ERROR missing model name` |
| `ERROR unknown` | Not recognized | `ERROR unknown command "xyz"` |
| `ERROR permission denied` | Authorization | `ERROR permission denied` |
| `ERROR timeout` | Timeout | `ERROR timeout waiting for response` |

### Detail Patterns

```
ERROR <operation> failed: <reason>
ERROR cannot <operation>: <reason>
ERROR invalid <type> "<value>": <constraint>
```

### Wrapped Errors

Errors often include underlying causes:

```
ERROR cannot deploy "mysql": cannot download charm: network error: connection refused
```

## Error Handling in Scripts

### Basic Exit Code Handling

```bash
#!/bin/bash
set -e

juju deploy mysql
# Script exits on error with exit code 1

# Alternative: capture and handle
if ! juju deploy mysql; then
    echo "Deployment failed"
    exit 1
fi
```

### JSON Error Parsing

```python
import subprocess
import json

result = subprocess.run(
    ["juju", "status", "--format", "json"],
    capture_output=True,
    text=True
)

if result.returncode != 0:
    # Parse error from stderr
    error_msg = result.stderr.strip()
    if error_msg.startswith("ERROR "):
        error_msg = error_msg[6:]  # Remove "ERROR " prefix
    print(f"Command failed: {error_msg}")
    exit(result.returncode)

status = json.loads(result.stdout)
```

### Specific Error Detection

```bash
#!/bin/bash

output=$(juju deploy mysql 2>&1)
exit_code=$?

if [[ $exit_code -eq 0 ]]; then
    echo "Success"
elif [[ $output == *"not found"* ]]; then
    echo "Resource not found"
elif [[ $output == *"permission denied"* ]]; then
    echo "Permission error"
elif [[ $output == *"already exists"* ]]; then
    echo "Already exists, continuing..."
else
    echo "Unknown error: $output"
    exit $exit_code
fi
```

### Retry Logic

```bash
#!/bin/bash

max_retries=3
retry_count=0

while [[ $retry_count -lt $max_retries ]]; do
    if juju deploy mysql; then
        break
    fi
    ((retry_count++))
    echo "Retry $retry_count/$max_retries..."
    sleep 5
done

if [[ $retry_count -eq $max_retries ]]; then
    echo "Max retries exceeded"
    exit 1
fi
```

## Error Categories by Command Group

### Application Commands

| Error | Commands | Message Pattern |
|-------|----------|-----------------|
| Not found | `remove-application`, `config`, `expose` | `application "X" not found` |
| Already exists | `deploy` | `application "X" already exists` |
| Invalid name | `deploy` | `invalid application name` |
| Hook failure | `remove-application` | `application has hook failures` |

### Model Commands

| Error | Commands | Message Pattern |
|-------|----------|-----------------|
| Not found | `destroy-model`, `show-model` | `model "X" not found` |
| Already exists | `add-model` | `model "X" already exists` |
| Has resources | `destroy-model` | `model has applications/machines` |
| Busy | `destroy-model`, `migrate` | `model is currently busy` |

### Controller Commands

| Error | Commands | Message Pattern |
|-------|----------|-----------------|
| Not found | `destroy-controller`, `show-controller` | `controller "X" not found` |
| Already exists | `bootstrap`, `register` | `controller "X" already exists` |
| Has models | `destroy-controller` | `controller has models` |
| Connection | `login`, `register` | `cannot connect to controller` |

### Cloud Commands

| Error | Commands | Message Pattern |
|-------|----------|-----------------|
| Not found | `remove-cloud`, `show-cloud` | `cloud "X" not found` |
| Already exists | `add-cloud` | `cloud "X" already exists` |
| Invalid credentials | `add-credential`, `update-credential` | `invalid credentials` |
| Region not found | `regions` | `region "X" not found in cloud` |

### User Commands

| Error | Commands | Message Pattern |
|-------|----------|-----------------|
| Not found | `show-user`, `disable-user` | `user "X" not found` |
| Already exists | `add-user` | `user "X" already exists` |
| Invalid credentials | `login` | `invalid credentials` |
| Disabled | `login` | `user "X" is disabled` |

## Silent Errors

Some errors use `ErrSilent` to avoid duplicate messaging:

```go
// Error was already displayed, don't print again
return cmd.ErrSilent
```

This is used when:
- Error details were already logged
- Error is expected and handled internally
- Partial output was already written

## Error Logging

### Debug Mode

Enable detailed error logging:

```bash
juju deploy mysql --debug
```

Output includes:
- Full error stack trace
- API call details
- Internal state dumps

### Log File Location

```
~/.local/share/juju/log/juju.log
```

### Log Levels

| Level | Output |
|-------|--------|
| ERROR | Always visible |
| WARNING | Always visible |
| INFO | With `--verbose` |
| DEBUG | With `--debug` |
| TRACE | With `JUJU_LOGGING_CONFIG=<root>=TRACE` |

## Plugin Error Handling

### Plugin Exit Code Passthrough

When a plugin exits with code N, the CLI exits with the same code:

```go
// cmd/juju/commands/plugin.go
if exitError, ok := err.(*exec.ExitError); ok && exitError != nil {
    status := exitError.ProcessState.Sys().(syscall.WaitStatus)
    if status.Exited() {
        return utils.NewRcPassthroughError(status.ExitStatus())
    }
}
```

### Plugin Not Found

If a plugin executable is not found:

```
ERROR juju "unknown-command" is not a juju command. See "juju --help".

Did you mean:
    similar-command
```

## Error Recovery

### Transient Errors

Some errors are transient and can be retried:

```bash
# Network timeouts
juju status --retry-count 3 --retry-delay 1s
```

### Force Operations

For destructive operations, errors can be bypassed with `--force`:

```bash
juju remove-application mysql --force
juju destroy-controller myctl --force
```

**Warning:** Force operations skip validation and cleanup.
