# Juju CLI Extensibility Model

## Overview

Juju provides two primary extensibility mechanisms:
1. **Plugin Architecture** - External binaries for custom commands
2. **Command Registration** - Internal command registration for forks

## Plugin Architecture

### Plugin Discovery

Plugins are discovered automatically from the system PATH:

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

### Plugin Naming Convention

Plugins must follow the naming pattern:

| Pattern | Valid | Example |
|---------|-------|---------|
| `juju-<name>` | Yes | `juju-myplugin` |
| `juju-<Name>` | No | `juju-MyPlugin` |
| `juju-` | No | Missing name |
| `juju-123` | No | Must start with letter |

### Plugin Execution

When a command is not found, Juju attempts plugin execution:

```go
func RunPlugin(callback cmd.MissingCallback) cmd.MissingCallback {
    return func(ctx *cmd.Context, subcommand string, args []string) error {
        cmdName := JujuPluginPrefix + subcommand
        plugin := &PluginCommand{name: cmdName}
        // ...
        err := plugin.Run(ctx)
        // If plugin not found, fall through to callback
        if isExecutableNotFound(err) {
            return callback(ctx, subcommand, args)
        }
        return err
    }
}
```

### Plugin Command Implementation

```go
type PluginCommand struct {
    cmd.CommandBase
    name           string
    controllerName string
    modelName      string
    args           []string
}

func (c *PluginCommand) Run(ctx *cmd.Context) error {
    command := exec.Command(c.name, c.args...)
    
    // Inject context environment
    env := os.Environ()
    if c.controllerName != "" {
        env = utils.Setenv(env, osenv.JujuControllerEnvKey+"="+c.controllerName)
    }
    if c.modelName != "" {
        env = utils.Setenv(env, osenv.JujuModelEnvKey+"="+c.modelName)
    }
    command.Env = env
    
    // Connect I/O
    command.Stdin = ctx.Stdin
    command.Stdout = ctx.Stdout
    command.Stderr = ctx.Stderr
    
    return command.Run()
}
```

### Plugin Environment Variables

Plugins receive context through environment variables:

| Variable | Description |
|----------|-------------|
| `JUJU_CONTROLLER` | Current controller name |
| `JUJU_MODEL` | Current model name |
| `JUJU_DATA` | Configuration directory path |

### Plugin Description Protocol

Plugins can provide descriptions for help output:

```bash
juju-myplugin --description
```

Output (first line only):
```
My custom plugin for Juju
```

### Plugin Exit Code Passthrough

Plugin exit codes are preserved:

```go
if exitError, ok := err.(*exec.ExitError); ok {
    status := exitError.ProcessState.Sys().(syscall.WaitStatus)
    if status.Exited() {
        return utils.NewRcPassthroughError(status.ExitStatus())
    }
}
```

This allows plugins to return arbitrary exit codes that scripts can handle.

## Internal Command Registration

### Command Interface

All commands implement the `Command` interface:

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

### Registration Pattern

Commands are registered in `registerCommands()`:

```go
// cmd/juju/commands/main.go
func registerCommands(r commandRegistry) {
    // Infrastructure commands
    r.Register(newBootstrapCommand())
    r.Register(cloud.NewAddCloudCommand(&cloudToCommandAdaptor{}))
    
    // Application commands
    r.Register(application.NewDeployCommand())
    r.Register(application.NewRemoveApplicationCommand())
    
    // Model commands
    r.Register(model.NewAddModelCommand())
    r.Register(model.NewDestroyCommand())
    // ...
}
```

### Command Base Types

#### CommandBase

Basic command implementation:

```go
type CommandBase struct{}

func (c *CommandBase) IsSuperCommand() bool { return false }
func (c *CommandBase) SetFlags(f *gnuflag.FlagSet) {}
func (c *CommandBase) Init(args []string) error {
    return CheckEmpty(args)
}
func (c *CommandBase) AllowInterspersedFlags() bool { return true }
```

#### ModelCommandBase

Model-scoped commands:

```go
type ModelCommandBase struct {
    CommandBase
    store            jujuclient.ClientStore
    _modelIdentifier string
    _controllerName  string
    allowDefaultModel bool
}
```

#### ControllerCommandBase

Controller-scoped commands:

```go
type ControllerCommandBase struct {
    CommandBase
    store            jujuclient.ClientStore
    _controllerName  string
    allowDefaultController bool
}
```

### SuperCommand

The root command router:

```go
type SuperCommand struct {
    CommandBase
    Name     string
    Purpose  string
    Doc      string
    subcmds  map[string]commandReference
    // ...
}

func (c *SuperCommand) Register(subcmd Command) {
    info := subcmd.Info()
    c.insert(commandReference{
        name:    info.Name, 
        command: subcmd,
    })
}
```

## Command Aliases

### Alias Registration

Aliases provide alternative names for commands:

```go
// cmd/cmd/supercommand.go
func (c *SuperCommand) RegisterAlias(name, forName string, check DeprecationCheck) {
    action, found := c.subcmds[forName]
    if !found {
        panic(fmt.Sprintf("%q not found when registering alias", forName))
    }
    c.insert(commandReference{
        name:    name,
        command: action.command,
        alias:   forName,
        check:   check,
    })
}
```

### Deprecated Aliases

Aliases can have deprecation warnings:

```go
type DeprecationCheck interface {
    Deprecated() (bool, string)  // Is deprecated, replacement message
    Obsolete() bool              // Should be hidden/removed
}

r.RegisterDeprecated(subcmd, check)
```

### User Aliases File

Users can define custom aliases in `~/.local/share/juju/aliases`:

```
# Format: name = command [args...]
prod = switch aws-prod:production
stg = switch aws-prod:staging
```

## Extension Points

### 1. Flag Extensions

Commands can add global flags:

```go
type GlobalFlags interface {
    AddFlags(*gnuflag.FlagSet)
}

type MyFlags struct{}

func (f *MyFlags) AddFlags(fs *gnuflag.FlagSet) {
    fs.BoolVar(&f.myFlag, "my-flag", false, "My custom flag")
}

// In NewSuperCommand:
jcmd = jujucmd.NewSuperCommand(cmd.SuperCommandParams{
    GlobalFlags: &MyFlags{},
    // ...
})
```

### 2. Missing Callback

Handle unknown commands:

```go
type MissingCallback func(ctx *Context, subcommand string, args []string) error

// Register custom handler
jcmd = jujucmd.NewSuperCommand(cmd.SuperCommandParams{
    MissingCallback: func(ctx *cmd.Context, subcommand string, args []string) error {
        // Custom handling
        return fmt.Errorf("unknown command: %s", subcommand)
    },
})
```

### 3. Notification Hooks

Receive run notifications:

```go
jcmd = jujucmd.NewSuperCommand(cmd.SuperCommandParams{
    NotifyRun: func(cmdName string) {
        log.Printf("Running command: %s", cmdName)
    },
})
```

### 4. Help Topics

Add custom help topics:

```go
c.AddHelpTopic("mytopic", "My Topic", `
This is help content for my custom topic.
It can be multi-line.
`)
```

## Creating a Plugin

### Shell Script Plugin

```bash
#!/bin/bash
# /usr/local/bin/juju-hello

case "$1" in
    --description)
        echo "Say hello from Juju"
        exit 0
        ;;
    -h|--help)
        echo "Usage: juju hello [name]"
        exit 0
        ;;
esac

name="${1:-World}"
echo "Hello, $name! (Controller: $JUJU_CONTROLLER, Model: $JUJU_MODEL)"
exit 0
```

### Python Plugin

```python
#!/usr/bin/env python3
# /usr/local/bin/juju-pystatus

import os
import sys
import json

if "--description" in sys.argv:
    print("Show model status in Python")
    sys.exit(0)

# Access environment
controller = os.environ.get("JUJU_CONTROLLER", "unknown")
model = os.environ.get("JUJU_MODEL", "unknown")

# Call juju status
import subprocess
result = subprocess.run(
    ["juju", "status", "--format", "json"],
    capture_output=True,
    text=True
)

if result.returncode == 0:
    status = json.loads(result.stdout)
    print(f"Model: {status['model']['name']}")
    print(f"Applications: {len(status.get('applications', {}))}")
else:
    print(f"Error: {result.stderr}", file=sys.stderr)
    sys.exit(result.returncode)
```

### Go Plugin

```go
// /usr/local/bin/juju-goplugin
package main

import (
    "fmt"
    "os"
    "os/exec"
)

func main() {
    if len(os.Args) > 1 && os.Args[1] == "--description" {
        fmt.Println("Go-based Juju plugin")
        os.Exit(0)
    }
    
    // Access context
    controller := os.Getenv("JUJU_CONTROLLER")
    model := os.Getenv("JUJU_MODEL")
    
    // Execute juju commands
    cmd := exec.Command("juju", "status", "--format", "json")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    if err := cmd.Run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
```

## Best Practices

### For Plugin Authors

1. **Handle `--description`** - Return a single-line description
2. **Handle `-h/--help`** - Show usage information
3. **Preserve exit codes** - Use meaningful exit codes
4. **Respect context** - Use JUJU_CONTROLLER and JUJU_MODEL
5. **Handle missing context** - Gracefully handle empty environment

### For Fork Authors

1. **Follow naming conventions** - Match Juju's verb-noun pattern
2. **Use base types** - Inherit from ModelCommandBase or ControllerCommandBase
3. **Register in groups** - Keep command registration organized
4. **Add tests** - Use the tc testing framework
5. **Update docs** - Keep documentation synchronized

## Limitations

### Current Constraints

| Aspect | Limitation |
|--------|------------|
| Plugin I/O | No direct API access |
| Authentication | Must handle separately |
| State | No shared state with Juju |
| Completion | Manual implementation |
| Help | Custom formatting required |

### Workarounds

1. **API Access** - Call `juju` commands with `--format json`
2. **Authentication** - Use `juju login` in automation
3. **State** - Store in temp files or environment
4. **Completion** - Implement shell completion separately
5. **Help** - Use consistent help format

## Future Extensions

### Potential Enhancements

1. **Native Plugin API** - Direct Go plugin support
2. **Completion Integration** - Automatic tab completion
3. **Context Injection** - Richer environment variables
4. **Error Handling** - Structured error output
5. **Metrics** - Plugin usage tracking
