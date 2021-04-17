package task

import "github.com/jubnzv/go-taskwarrior"

type Tasks []taskwarrior.Task

type TaskController struct {
	Tasks
}

func (t *TaskController) InitTasks(tasks []taskwarrior.Task) {
	t.Tasks = tasks
}
