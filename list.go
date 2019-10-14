package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

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
