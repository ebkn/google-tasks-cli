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
	ctx = context.Background()
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
	if err := newApp().Run(os.Args); err != nil {
		exitCode := 1
		if excoder, ok := err.(cli.ExitCoder); ok {
			exitCode = excoder.ExitCode()
		}
		logger.Log("error", err.Error())
		os.Exit(exitCode)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "google-tasks-cli (gtc)"
	app.Usage = "Manage Google Tasks API"
	app.Version = "0.0.1"
	app.Author = "ebkn"
	app.Email = "ktennis.mqekr12@gmail.com"
	app.Commands = commands
	return &app
}
