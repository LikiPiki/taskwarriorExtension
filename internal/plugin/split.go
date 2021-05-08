package plugin

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/jubnzv/go-taskwarrior"
	"github.com/likipiki/tj/internal/task"
	"github.com/likipiki/tj/internal/utilities"
)

type SplitPlugin struct {
	task.TaskController
	TaskNumber int
	TaskText   string
}

// Initialize plugin
func (s *SplitPlugin) Init(tasks []taskwarrior.Task) {
	s.InitTasks(tasks)
}

// Get plugin name
func (s *SplitPlugin) GetCommandName() string {
	return "sp"
}

// Get short plugin description with params
func (s *SplitPlugin) GetDescription() string {
	return "Split task to many small tasks"
}

// Get short plugin usage string
func (s *SplitPlugin) GetUsage() string {
	return "tj sp [task-number] [new-task]"
}

// Parse plugin arguments
func (s *SplitPlugin) ParseArguments(args []string) bool {
	if len(args) > 1 {
		taskNumber, err := strconv.Atoi(args[0])
		s.TaskText = strings.Join(args[1:], " ")

		if err != nil {
			fmt.Println(utilities.ColorString("First argument must be a string", "red", ""))
			return false
		}

		s.TaskNumber = taskNumber
		return true
	}
	return false
}

// Main plugin logic
func (s *SplitPlugin) Command() error {
	fmt.Println(utilities.ColorString("Creating new task: ", "green", "") + s.TaskText)
	output, errOutput, err := utilities.RunConsoleCommand("task add " + s.TaskText)

	if err != nil {
		fmt.Print(utilities.ColorString(errOutput, "red", ""))
		return nil
	}

	fmt.Print(output)

	// Parsing output and find created task number
	createdTaskStr := ""
	for _, ch := range output {
		if unicode.IsDigit(ch) {
			createdTaskStr += string(ch)
		}
	}

	createdTask, err := strconv.Atoi(createdTaskStr)
	if err != nil {
		fmt.Println(
			utilities.ColorString(
				"Converting "+createdTaskStr+" failed!",
				"red",
				"",
			),
		)
		return nil
	}

	fmt.Println(
		utilities.ColorString(
			fmt.Sprintf(
				"Updating task %d", s.TaskNumber,
			),
			"green",
			"",
		),
	)
	output, errOutput, err = utilities.RunConsoleCommand(
		fmt.Sprintf(
			"task %d modify dep:%d",
			s.TaskNumber,
			createdTask,
		),
	)

	if err != nil {
		fmt.Print(utilities.ColorString(errOutput, "red", ""))
		return nil
	}

	fmt.Print(output)

	return nil
}
