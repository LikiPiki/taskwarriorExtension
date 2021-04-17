package controller

import (
	"fmt"
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

	plugins := plugin.Plugins{
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
					break
				}
			}
		}
	} else {
		ShowHelp(plugins)
	}
}

func ShowHelp(plugins plugin.Plugins) {
	fmt.Println("tj -- is simple command line tool, provide new functionality for taskwarrior.\n")
	fmt.Println("Usage:\n  tj [command]\n")
	fmt.Println("Avalible commands:")

	for _, plugin := range plugins {
		fmt.Printf("  %s  %s\n", plugin.GetCommandName(), plugin.GetDescription())
	}
}
