package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	OAUTH_REDIRECT_URL = "http://127.0.0.1:9999/oauth/callback"
)

var (
	config            *oauth2.Config
	googleOAuthScopes = []string{"https://www.googleapis.com/auth/tasks"}
)

var (
	ctx = context.Background()
)

func doAuthorize(c *cli.Context) error {
	config = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		RedirectURL:  OAUTH_REDIRECT_URL,
		Scopes:       googleOAuthScopes,
		Endpoint:     google.Endpoint,
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	log.Printf("Authentication URL: %s\n", url)

	time.Sleep(1 * time.Second)
	open.Run(url)
	time.Sleep(1 * time.Second)

	http.HandleFunc("/oauth/callback", callbackHandler)

	server := http.Server{Addr: ":9999"}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}
