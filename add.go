package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"
	"google.golang.org/api/tasks/v1"
)

func doAddTask(c *cli.Context) error {
	srv := getTasksClient()

	title := c.Args().First()
	if title == "" {
		return errors.New("no title found")
	}

	taskLists, err := srv.Tasklists.List().Do()
	if err != nil {
		return fmt.Errorf("Unable to retrieve task lists. %s", err.Error())
	}
	if len(taskLists.Items) == 0 {
		fmt.Print("No task lists found.")
		return nil
	}

	taskList := taskLists.Items[0]
	task := tasks.Task{Title: title}
	_, err = srv.Tasks.Insert(taskList.Id, &task).Do()
	if err != nil {
		return fmt.Errorf("failed to add task: %s", err.Error())
	}
	return nil
}
