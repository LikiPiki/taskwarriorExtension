# Taskwarrior Extension
![build](https://github.com/likipiki/taskwarriorExtension/actions/workflows/build.yaml/badge.svg) ![deploy](https://github.com/likipiki/taskwarriorExtension/actions/workflows/hugo.yaml/badge.svg)

Taskwarrior commands extension. See [Documentation](https://likipiki.github.io/taskwarriorExtension/) for more details!

## Features
- fast, written on golang language 
- simple customizable and extensible
- zero configuration
- colorfull output

## Installation

If you have got Golang installed on your machine, use this command. Else, see instructions in [Documentation](https://likipiki.github.io/taskwarriorExtension/#installation)

```console
go install github.com/likipiki/taskwarriorExtension ./cmd/tj
```
## Avaliable commands
- `tj tree [project-name]` -- show project and subprojects tasks like tree

## Development
Branching model, plugin creation and others your can find in the [Development](https://likipiki.github.io/taskwarriorExtension/docs/development/) page.

## References
-  [TaskWarrior](https://taskwarrior.org)
-  [go-taskwarrior library](https://github.com/jubnzv/go-taskwarrior)