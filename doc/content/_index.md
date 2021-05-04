---
title: "Documentation"
weight: 1
bookFlatSection: true
bookToc: true
draft: false
---

# Taskwarrior Extension 

Taskwarrior commands extension. See [Wiki](https://github.com/LikiPiki/taskwarriorExtension/wiki) for more details!

## Features
- fast, written on golang language 
- simple customizable and extensible
- zero configuration
- colorfull output

## Installation

If you have got Golang installed on your machine, use this command. Else, see instructions in [Wiki](https://github.com/LikiPiki/taskwarriorExtension/wiki#installing-from-binary-file)

## Installing from binary file
1. Download binary file from [Releases](https://github.com/LikiPiki/taskwarriorExtension/releases) section
2. Make the file executable `chmod +x tj`
3. Add binary file to your `/usr/local/bin/`

```console
go install github.com/likipiki/taskwarriorExtension ./cmd/tj
```
## Avaliable commands
- `tj tree [project-name]` -- show project and subprojects tasks like tree

## References
-  [TaskWarrior](https://taskwarrior.org)
-  [go-taskwarrior library](https://github.com/jubnzv/go-taskwarrior)