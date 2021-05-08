package plugin

import (
	"fmt"
	"strings"

	"github.com/jubnzv/go-taskwarrior"

	"github.com/likipiki/tj/internal/task"
	"github.com/likipiki/tj/internal/utilities"
)

type TreePlugin struct {
	task.TaskController
	Project string
}

// Initialize plugin
func (t *TreePlugin) Init(tasks []taskwarrior.Task) {
	t.InitTasks(tasks)
}

// Get plugin name
func (t *TreePlugin) GetCommandName() string {
	return "tree"
}

// Get short plugin description with params
func (t *TreePlugin) GetDescription() string {
	return "Show project task as tree with summary"
}

func (t *TreePlugin) GetUsage() string {
	return "tj tree [project]"
}

// Parse plugin arguments
func (t *TreePlugin) ParseArguments(args []string) bool {
	if len(args) > 0 {
		t.Project = args[0]
		return true
	}
	return false
}

// Main plugin logic
func (t *TreePlugin) Command() error {
	projectTasks := make(map[string][]taskwarrior.Task)
	completed := 0
	all := 0

	// Sort tasks in current project and subprojects
	for _, task := range t.Tasks {
		if strings.HasPrefix(task.Project, t.Project) && task.Status != taskDeleted {
			projectTasks[task.Project] = append(projectTasks[task.Project], task)
			if task.Status == taskCompleted {
				completed++
			}
			all++
		}
	}
	if mainProject := projectTasks[t.Project]; len(mainProject) > 0 {
		fmt.Print(utilities.ColorString("Project: ", "orange", "bold") + t.Project + " ")
		fmt.Print(utilities.ProgressBarString(30, completed, all))

		for i, task := range mainProject {
			treePrintTask(i, task)
		}

		fmt.Println("")
		delete(projectTasks, t.Project)

		for key := range projectTasks {
			fmt.Println("  ", key)

			for i, task := range projectTasks[key] {
				treePrintTask(i, task)
			}
		}
	}

	return nil
}

func treePrintTask(i int, task taskwarrior.Task) {
	if task.Status == taskCompleted {
		task.Description = utilities.ColorString("[DONE] ", "green", "bold") + task.Description
	}

	fmt.Println(
		"    ",
		utilities.ColorString(fmt.Sprintf("%d", i+1), "orange", "bold"),
		task.Description,
		utilities.ColorString(fmt.Sprintf("%f", task.Urgency), "orange", "bold"),
	)
}
