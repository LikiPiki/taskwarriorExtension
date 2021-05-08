package plugin

import "github.com/jubnzv/go-taskwarrior"

const (
	taskCompleted = "completed"
	taskDeleted   = "deleted"
	taskPending   = "pending"
)

type Plugin interface {
	// Initialize plugin
	Init(tasks []taskwarrior.Task)
	// Get plugin name
	GetCommandName() string
	// Get short plugin description with params
	GetDescription() string
	// Get short plugin usage string
	GetUsage() string
	// Parse plugin argumetns
	ParseArguments(args []string) bool
	// Main plugin logic
	Command() error
}

// Plugins is an array of Plugin interfaces
type Plugins []Plugin
