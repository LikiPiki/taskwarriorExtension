---
title: "Plugins"
weight: 2
bookFlatSection: false
bookToc: true
---

# Plugins

## Create new plugin
Plugins provide support to add new functionality. The first step, create a new plugin file in `internal/plugin/` directory and declare a new plugin. This is a simple plugin code example.

```go
package plugin

import (
	"github.com/likipiki/tj/internal/task"
)

type MyPlugin struct{
        // All task from taskwarrior
        task.TaskController
}

// Initialize plugin
func (p *MyPlugin) Init(tasks []taskwarrior.Task) {
        p.Tasks = tasks
}

// Get plugin name
func (p *MyPlugin) GetCommandName() string {
        panic("not implemented") // TODO: Implement
}

// Get short plugin description with params
func (p *MyPlugin) GetDescription() string {
        panic("not implemented") // TODO: Implement
}

// Parse plugin arguments
func (p *MyPlugin) ParseArguments(args []string) bool {
        panic("not implemented") // TODO: Implement
}

// Main plugin logic
func (p *MyPlugin) Command() error {
        panic("not implemented") // TODO: Implement
}
```

You can manually generate plugin interface by following command:
```console
impl 'p *MyPlugin' github.com/likipiki/tj/internal/plugin.Plugin
```

First install `impl` tool from [here](https://github.com/josharian/impl).

## Plugin Registration

Go to `internal/controller/main.go` file and register plugin like other plugin:

This is simple `treePlugin` registration code snippet.
```go
treePlugin := plugin.TreePlugin{}

plugins := plugin.Plugins{
	&treePlugin,
}
``` 