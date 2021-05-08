package plugin

import (
	"github.com/jubnzv/go-taskwarrior"
	"github.com/likipiki/tj/internal/task"
)

type SplitPlugin struct {
	task.TaskController
	taskNumber int
	taskText   string
}

// Initialize plugin
func (s *SplitPlugin) Init(tasks []taskwarrior.Task) {
	s.InitTasks(tasks)
}

// Get plugin name
func (s *SplitPlugin) GetCommandName() string {
	return "split"
}

// Get short plugin description with params
func (s *SplitPlugin) GetDescription() string {
	return "Split task to many small tasks"
}

// Get short plugin usage string
func (s *SplitPlugin) GetUsage() string {
	return "tj split [task-number] [new-task]"
}

// Parse plugin arguments
func (s *SplitPlugin) ParseArguments(args []string) bool {
	if len(args) > 0 {
		return true
	}
	return false
}

// Main plugin logic
func (s *SplitPlugin) Command() error {
	return nil
}
