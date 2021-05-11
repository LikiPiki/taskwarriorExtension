package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/jubnzv/go-taskwarrior"

	"github.com/likipiki/tj/internal/plugin"
	"github.com/likipiki/tj/internal/utilities"
)

const (
	version = "0.1"
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
					return
				}
				fmt.Println(utilities.ColorString("Plugin usage:\n  ", "orange", "") + plugin.GetUsage())
				return
			}
		}
	}

	ShowHelp(plugins)
}

func ShowHelp(plugins plugin.Plugins) {
	fmt.Printf(
		"tj -- is simple command-line tool, which provide new functionality for taskwarrior. Version %s\n",
		version,
	)
	fmt.Print("Usage:\n  tj [command]\n\n")
	fmt.Println("Avalible commands:")

	for _, plugin := range plugins {
		// %-4s - longest command name length
		fmt.Printf("  %-4s -- %s\n", plugin.GetCommandName(), plugin.GetDescription())
	}
}
