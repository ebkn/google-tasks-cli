package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli"
	"google.golang.org/api/tasks/v1"
)

const credentialFile = "credentials.json"

var (
	ctx context.Context
)

func getTasksClient() *tasks.Service {
	token, err := tokenFromFile(credentialFile)
	if err != nil {
		log.Println("You should authorize with Google. Try authorize command.")
		log.Fatalf("error: %s", err.Error())
	}

	client := config.Client(context.Background(), token)

	srv, err := tasks.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}
	return srv
}

func main() {
	ctx = context.Background()

	app := cli.NewApp()
	app.Name = "google-tasks-cli"

	app.Commands = []cli.Command{
		{
			Name:        "authorize",
			Description: "authorize with Google",
			Action:      authorize,
		},
		{
			Name:        "add",
			Usage:       "google-tasks-cli add xxx",
			Description: "add task",
			Aliases:     []string{"a"},
			Action:      addTask,
		},
		{
			Name:        "list",
			Description: "list tasks",
			Aliases:     []string{"l"},
			Action:      getTasks,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
