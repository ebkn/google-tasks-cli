package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load .env file: %s", err.Error())
	}

	if err := newApp().Run(os.Args); err != nil {
		exitCode := 1
		if excoder, ok := err.(cli.ExitCoder); ok {
			exitCode = excoder.ExitCode()
		}
		fmt.Printf("error: %s", err.Error())
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
	return app
}
