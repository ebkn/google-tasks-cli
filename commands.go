package main

import "github.com/urfave/cli"

var commands = []cli.Command{
	commandAuthorize,
	commandAddTask,
	commandListTasks,
}

var commandAuthorize = cli.Command{
	Name:        "authorize",
	Description: "authorize with Google",
	Action:      doAuthorize,
}

var commandAddTask = cli.Command{
	Name:        "add",
	Usage:       "google-tasks-cli add xxx",
	Description: "add task",
	Aliases:     []string{"a"},
	Action:      doAddTask,
}

var commandListTasks = cli.Command{
	Name:        "list",
	Description: "list tasks",
	Aliases:     []string{"l"},
	Action:      doListTasks,
}
