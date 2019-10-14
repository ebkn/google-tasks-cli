package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
	"google.golang.org/api/tasks/v1"
)

func doAddTask(c *cli.Context) error {
	srv := getTasksClient()

	title := c.Args().First()
	if title == "" {
		return fmt.Errorf("no title found.")
	}

	taskLists, err := srv.Tasklists.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists. %v", err)
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

func doListTasks(c *cli.Context) error {
	srv := getTasksClient()

	taskLists, err := srv.Tasklists.List().Do()
	if err != nil {
		log.Fatalf("Unable to retrieve task lists. %v", err)
	}

	if len(taskLists.Items) == 0 {
		fmt.Print("No task lists found.")
		return nil
	}

	taskList := taskLists.Items[0]
	tasks, err := srv.Tasks.List(taskList.Id).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve tasks. %v", err)
	}
	for _, t := range tasks.Items {
		fmt.Printf("- %s (~%s)\n", t.Title, t.Due)
	}
	return nil
}
