---
title: "Usage"
weight: 1
bookToc: true
draft: false
---
# Usage

# Help

If yue run `tj` without arguments, you can find a little documentation about all commands and simple usage.

```console
tj -- is simple command line tool, provide new functionality for taskwarrior.

Usage:
  tj [command]

Avalible commands:
  tree  Show project task as tree with summary
```

# Task tree view
Command `tj tree my-project`. Show all tasks in `my-project` project with subprojects like a tree.
Ignoring deleted tasks, show only pending and completed tasks with progress bar.

```console
Project: my-project [############------------------] 40.00%
     1 Task 1 -3.183560
     2 Another task 1.816440
     3 Task with big urgency 10.716400
     4 [DONE] This task is done 11.398100

   my-project.subproject
     1 Foo task 1.816440
     2 [DONE] One 1.816440
     3 [DONE] Two 1.816440
     4 [DONE] Three 1.821920
   my-project.subproject2
     1 New task here 10.880600
     2 Bla bla bla task 10.880600
```