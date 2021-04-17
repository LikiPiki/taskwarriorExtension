package controller

import (
	"log"
	"os"

	"github.com/jubnzv/go-taskwarrior"

	"github.com/likipiki/tj/internal/plugin"
)

type Controller struct {
	Plugins []plugin.Plugin
}

func Control() {
	tw, err := taskwarrior.NewTaskWarrior("~/.taskrc")

	if err != nil {
		panic("Error parsing TaskWarrior data")
	}

	tw.FetchAllTasks()

	args := os.Args[1:]

	treePlugin := plugin.TreePlugin{}

	plugins := []plugin.Plugin{
		&treePlugin,
	}

	for _, plugin := range plugins {
		plugin.Init(tw.Tasks)
	}

	if len(args) > 0 {
		command := args[0]

		for _, plugin := range plugins {
			if plugin.GetCommandName() == command {
				if plugin.ParseArguments(args[1:]) {
					err := plugin.Command()

					if err != nil {
						log.Println(err)
					}
				}
			}
		}

	}
}
