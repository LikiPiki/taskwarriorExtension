package plugin

import (
	"fmt"
	"strings"

	"github.com/jubnzv/go-taskwarrior"
	"github.com/likipiki/tj/internal/task"
	"github.com/likipiki/tj/internal/utilities"
)

type ContextPlugin struct {
	task.TaskController
	Config  *utilities.Config
	Context string
}

// Initialize plugin
func (c *ContextPlugin) Init(tasks []taskwarrior.Task) {
	c.InitTasks(tasks)
	c.Config = c.Config.GetConfig()
}

// Get plugin name
func (c *ContextPlugin) GetCommandName() string {
	return "ctx"
}

// Get short plugin description with params
func (c *ContextPlugin) GetDescription() string {
	return "Add context to adding task"
}

// Get short plugin usage string
func (c *ContextPlugin) GetUsage() string {
	return "tj ctx [new-context]"
}

// Parse plugin argumetns
func (c *ContextPlugin) ParseArguments(args []string) bool {
	if len(args) > 0 {
		c.Context = strings.Join(args, " ")
	}
	return true
}

// Main plugin logic
func (c *ContextPlugin) Command() error {
	if len(c.Context) > 0 {
		c.Config.Context = c.Context
		c.Config.UpdateConfig()
		fmt.Println(utilities.ColorString("Set up context: ", "green", "") + c.Context)
		return nil
	}
	fmt.Println(utilities.ColorString("Current context: ", "green", "") + c.Config.Context)
	return nil
}
