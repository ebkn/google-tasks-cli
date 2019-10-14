package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli"
)

const credentialFile = "credentials.json"

var (
	ctx context.Context
)

func main() {
	ctx = context.Background()

	// token, err := tokenFromFile(credentialFile)
	// if err != nil {
	// 	log.Println("You should authorize with Google. Try authorize command.")
	// 	log.Fatalf("error: %s", err.Error())
	// }
	//
	// log.Println(token)
	//
	app := cli.NewApp()
	app.Name = "google-tasks-cli"

	app.Commands = []cli.Command{
		{
			Name:   "authorize",
			Usage:  "authorize with Google",
			Action: authorize,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
