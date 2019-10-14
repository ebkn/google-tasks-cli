package main

import (
	"context"
	"log"

	"google.golang.org/api/tasks/v1"
)

const credentialFileName = "credentials.json"

func getTasksClient() *tasks.Service {
	token, err := loadToken(credentialFileName)
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
