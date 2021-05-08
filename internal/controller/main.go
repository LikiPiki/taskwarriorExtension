package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/jubnzv/go-taskwarrior"

	"github.com/likipiki/tj/internal/plugin"
	"github.com/likipiki/tj/internal/utilities"
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
	contextPlugin := plugin.ContextPlugin{}
	addPlugin := plugin.AddPlugin{}
	splitPlugin := plugin.SplitPlugin{}

	plugins := plugin.Plugins{
		&treePlugin,
		&contextPlugin,
		&addPlugin,
		&splitPlugin,
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
				fmt.Println(utilities.ColorString("Plugin usage:\n  ", "orange", "") + plugin.GetUsage())
				break
			}
		}
	} else {
		ShowHelp(plugins)
	}
}

func ShowHelp(plugins plugin.Plugins) {
	fmt.Println("tj -- is simple command-line tool, which provide new functionality for taskwarrior.\n")
	fmt.Println("Usage:\n  tj [command]\n")
	fmt.Println("Avalible commands:")

	for _, plugin := range plugins {
		// %-4s - longest command name length
		fmt.Printf("  %-5s -- %s\n", plugin.GetCommandName(), plugin.GetDescription())
	}
}
