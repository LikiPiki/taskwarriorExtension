---
title: "Development"
weight: 1
bookFlatSection: false
bookToc: true
---

# Development

# Development setup
1. Fork my repo
2. Clone forked repo to your `$GOPATH/src/github.com/username/repo-name`
3. Create branch and checkout to it
4. Create a new feature
5. Suggest pull request

## Building
I use simple [`Makefile`](https://github.com/LikiPiki/taskwarriorExtension/blob/master/Makefile) to build, install and testing project.
Following commands:
- `make build`
- `make install`
## Roadmap
- [x] Projects tasks tree view
- [x] Adding tasks with context
- [ ] Split task to smaller ones
- [ ] Sort task in tree view by urgency
- [ ] Create usage GIF

## Notes
- I use [`github.com/jubnzv/go-taskwarrior`](https://github.com/jubnzv/go-taskwarrior) library for extracting all tasks from taskwarrior.
- If you want write/edit documentation see [`README`](https://github.com/LikiPiki/taskwarriorExtension/tree/master/doc) here
