package plugin

import (
	"fmt"
	"strings"

	"github.com/jubnzv/go-taskwarrior"
	"github.com/likipiki/tj/internal/task"
	"github.com/likipiki/tj/internal/utilities"
)

type AddPlugin struct {
	task.TaskController
	Config *utilities.Config
	Args   string
}

// Initialize plugin
func (a *AddPlugin) Init(tasks []taskwarrior.Task) {
	a.InitTasks(tasks)
	a.Config = a.Config.GetConfig()
}

// Get plugin name
func (a *AddPlugin) GetCommandName() string {
	return "add"
}

// Get short plugin description with params
func (a *AddPlugin) GetDescription() string {
	return "Add new task, using context"
}

// Get short plugin usage string
func (a *AddPlugin) GetUsage() string {
	return "tj add Buy flowers +home"
}

// Parse plugin argumetns
func (a *AddPlugin) ParseArguments(args []string) bool {
	if len(args) > 0 {
		a.Args = strings.Join(args, " ")
		return true
	}
	return false
}

// Main plugin logic
func (a *AddPlugin) Command() error {
	fmt.Println("task add " + utilities.ColorString(a.Config.Context, "green", "") + " " + a.Args)
	output, errOutput, err := utilities.RunConsoleCommand("task add " + a.Config.Context + " " + a.Args)

	if err != nil {
		fmt.Print(utilities.ColorString(errOutput, "red", ""))
		return nil
	}

	fmt.Print(utilities.ColorString(output, "green", ""))
	return nil
}
