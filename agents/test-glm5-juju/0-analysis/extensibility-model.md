# Juju CLI Extensibility Model

## Overview

The Juju CLI provides multiple extension mechanisms for adding new functionality, customizing behavior, and integrating with external tools. This document describes the plugin system, command registration, user aliases, and extension boundaries.

## Extension Mechanisms

### 1. External Plugins

The primary extensibility mechanism is the external plugin system.

#### Plugin Discovery

Plugins are discovered by scanning `PATH` for executables matching the pattern `juju-*`:

```go
// cmd/juju/commands/plugin.go
const JujuPluginPrefix = "juju-"
const JujuPluginPattern = "^juju-[a-zA-Z]"

func findPlugins() []string {
    re := regexp.MustCompile(JujuPluginPattern)
    path := os.Getenv("PATH")
    plugins := []string{}
    for _, name := range filepath.SplitList(path) {
        entries, err := os.ReadDir(name)
        if err != nil {
            continue
        }
        for _, entry := range entries {
            fi, err := entry.Info()
            if err != nil {
                continue
            }
            if re.Match([]byte(fi.Name())) && (fi.Mode()&0111) != 0 {
                plugins = append(plugins, entry.Name())
            }
        }
    }
    sort.Strings(plugins)
    return plugins
}
```

#### Plugin Execution Flow

```
User runs: juju my-plugin arg1 arg2
         │
         ▼
SuperCommand checks registered commands
         │
         ├─ Not found? → MissingCallback
         │                    │
         ▼                    ▼
    RunPlugin searches PATH for juju-my-plugin
                              │
                              ├─ Found? → Execute juju-my-plugin arg1 arg2
                              │
                              └─ Not found? → Original MissingCallback
                                                    │
                                                    ▼
                                            "did you mean?" suggestion
```

#### Plugin Environment

Plugins inherit environment variables from the CLI:

```go
// cmd/juju/commands/plugin.go
env := os.Environ()
if c.controllerName != "" {
    env = utils.Setenv(env, osenv.JujuControllerEnvKey+"="+c.controllerName)
}
if c.modelName != "" {
    env = utils.Setenv(env, osenv.JujuModelEnvKey+"="+c.modelName)
}
command.Env = env
```

#### Plugin Arguments

Plugins receive common Juju flags:
- `-m`, `--model`: Model context
- `-c`, `--controller`: Controller context

```bash
# Example: juju-my-plugin receives --model flag
juju my-plugin --model mymodel arg1 arg2
# Plugin receives: arg1 arg2 with JUJU_MODEL=mymodel in environment
```

#### Plugin Description

Plugins can provide a description via `--description`:

```bash
$ juju-my-plugin --description
My custom Juju plugin for X
```

#### Creating a Plugin

```bash
#!/bin/bash
# /usr/local/bin/juju-hello

if [[ "$1" == "--description" ]]; then
    echo "Say hello to Juju"
    exit 0
fi

echo "Hello from Juju plugin!"
echo "Model: $JUJU_MODEL"
echo "Controller: $JUJU_CONTROLLER"
```

Make executable:
```bash
chmod +x /usr/local/bin/juju-hello
juju hello
```

### 2. User Aliases

Aliases provide shortcuts for common command sequences.

#### Alias File Location

```
~/.local/share/juju/aliases
```

#### Alias File Format

```
name = command [args...]
```

#### Example Aliases

```
# Shortcuts
d = deploy
rm = remove-application
st = status --relations

# Common patterns
deploy-test = deploy --model test
prod-status = status -m production
db-backup = exec postgresql/0 -- backup
```

#### Using Aliases

```bash
$ juju d mysql
# Equivalent to: juju deploy mysql

$ juju deploy-test postgresql
# Equivalent to: juju deploy --model test postgresql

$ juju prod-status
# Equivalent to: juju status -m production
```

#### Alias Resolution

```go
// cmd/cmd/supercommand.go
if userAlias, found := c.userAliases[args[0]]; found && !c.noAlias {
    logger.Debugf(context.TODO(), "using alias %q=%q", args[0], strings.Join(userAlias, " "))
    args = append(userAlias, args[1:]...)
}
```

#### Disabling Aliases

```bash
juju --no-alias d mysql
# Runs "d" command directly, not alias
```

### 3. Built-in Command Registration

New commands can be added to the Juju CLI source code.

#### Registration Point

```go
// cmd/juju/commands/main.go
func registerCommands(r commandRegistry) {
    // ... existing commands ...
    
    // Add new command
    r.Register(myplugin.NewMyCommand())
}
```

#### Command Interface

```go
// cmd/cmd/cmd.go
type Command interface {
    IsSuperCommand() bool
    Info() *Info
    SetFlags(f *gnuflag.FlagSet)
    Init(args []string) error
    Run(ctx *Context) error
    AllowInterspersedFlags() bool
}
```

#### Command Implementation Template

```go
package mycommand

import (
    "github.com/juju/gnuflag"
    "github.com/juju/juju/cmd/cmd"
)

type myCommand struct {
    cmd.CommandBase
    name string
    force bool
}

func NewMyCommand() cmd.Command {
    return &myCommand{}
}

func (c *myCommand) Info() *cmd.Info {
    return &cmd.Info{
        Name:    "my-command",
        Purpose: "Brief description",
        Doc:     "Detailed documentation...",
        Aliases: []string{"mc"},
    }
}

func (c *myCommand) SetFlags(f *gnuflag.FlagSet) {
    f.BoolVar(&c.force, "force", false, "Force operation")
}

func (c *myCommand) Init(args []string) error {
    if len(args) == 0 {
        return errors.New("name is required")
    }
    c.name = args[0]
    return cmd.CheckEmpty(args[1:])
}

func (c *myCommand) Run(ctx *cmd.Context) error {
    // Implementation
    ctx.Infof("Running my-command with name: %s", c.name)
    return nil
}
```

### 4. ModelCommand and ControllerCommand

Specialized command wrappers provide context.

#### ModelCommand

Commands operating on a model:

```go
type ModelCommand struct {
    cmd.CommandBase
    model API
    ModelName string
}

func (c *ModelCommand) SetFlags(f *gnuflag.FlagSet) {
    f.StringVar(&c.ModelName, "m", "", "Model to operate in")
    f.StringVar(&c.ModelName, "model", "", "")
}
```

#### ControllerCommand

Commands operating on a controller:

```go
type ControllerCommand struct {
    cmd.CommandBase
    controller API
    ControllerName string
}

func (c *ControllerCommand) SetFlags(f *gnuflag.FlagSet) {
    f.StringVar(&c.ControllerName, "c", "", "Controller to operate in")
    f.StringVar(&c.ControllerName, "controller", "", "")
}
```

### 5. Feature Flags

Experimental features can be enabled via environment variables.

#### Developer Features

```bash
export JUJU_DEV_FEATURE_FLAGS=developermode,feature1,feature2
```

#### User Features

```bash
export JUJU_FEATURES=feature1,feature2
```

#### Feature Flag Usage

```go
// cmd/juju/commands/main.go
func init() {
    featureflag.SetFlagsFromEnvironment(osenv.JujuFeatureFlagEnvKey, osenv.JujuFeatures)
}

// In command code
if featureflag.Enabled(featureflag.DeveloperMode) {
    // Developer-only functionality
    r.Register(model.NewDumpCommand())
}
```

### 6. Embedded Commands

Commands can be exposed via the controller API.

#### Embedded Command Whitelist

```go
// apiserver/apiserver.go
var allowedEmbeddedCommands = []string{
    "status",
    "show-unit",
    "show-application",
    // ...
}
```

#### Dashboard Integration

The Juju Dashboard can invoke embedded commands directly through the API.

## Extension Boundaries

### Internal Extension (Requires Source Modification)

| Extension Point | Location | Difficulty |
|----------------|----------|------------|
| Add new command | `cmd/juju/*/` | Medium |
| Add new API client | `api/` | Medium |
| Add new formatter | `cmd/cmd/output.go` | Easy |
| Add new flag type | `cmd/cmd/args.go` | Easy |

### External Extension (No Source Modification)

| Extension Point | Location | Difficulty |
|----------------|----------|------------|
| External plugin | `PATH` (juju-*) | Easy |
| User aliases | `~/.local/share/juju/aliases` | Easy |
| Environment config | Environment variables | Easy |
| Feature flags | `JUJU_DEV_FEATURE_FLAGS` | Easy |

## Plugin Best Practices

### 1. Naming Convention

Use descriptive names following the `juju-<verb>-<noun>` pattern:

```bash
juju-backup-db
juju-scale-app
juju-show-metrics
```

### 2. Description Support

Always implement `--description`:

```bash
if [[ "$1" == "--description" ]]; then
    echo "Brief description of what the plugin does"
    exit 0
fi
```

### 3. Environment Usage

Use Juju environment variables:

```bash
# Respect JUJU_MODEL and JUJU_CONTROLLER
if [[ -n "$JUJU_MODEL" ]]; then
    model_flag="--model $JUJU_MODEL"
fi
```

### 4. Exit Codes

Follow Juju exit code conventions:

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Error |
| 2 | Invalid arguments |

### 5. Error Formatting

Format errors consistently:

```bash
echo "ERROR $0: message" >&2
exit 1
```

## Extension Discovery

### Listing Plugins

```bash
juju help plugins
```

This outputs plugin names and descriptions:

```
Plugin         Description
juju-backup    Backup utilities for Juju
juju-myplugin  Custom plugin for X
```

### Help Integration

Plugins are included in the help system:

```bash
juju help my-plugin
```

## Extension Security

### Plugin Trust Model

1. Plugins are executed with user privileges
2. No sandboxing or isolation
3. Environment variables provide context, not authentication
4. Plugins can access Juju configuration files

### Security Recommendations

1. Only install plugins from trusted sources
2. Verify plugin signatures if available
3. Audit plugin code before installation
4. Use minimal permissions for plugin files

## Extension Registry

### Community Plugins

Common community plugins:
- `juju-wait` - Wait for model to settle
- `juju-crystal` - Crystal language integration
- `juju-replay` - Replay command sequences

### Distribution

Plugins can be distributed via:
- Snap Store (`snap install juju-myplugin`)
- Debian packages
- Homebrew
- Direct installation

## Summary

| Mechanism | Use Case | Complexity |
|-----------|----------|------------|
| External Plugin | Custom operations, integrations | Low |
| User Aliases | Shortcuts, common patterns | Very Low |
| Built-in Command | Core functionality | High |
| Feature Flags | Experimental features | Medium |
| Embedded Commands | Dashboard/API access | High |
