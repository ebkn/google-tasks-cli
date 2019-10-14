package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/skratchdot/open-golang/open"
	"github.com/urfave/cli"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	REDIRECT_URL = "http://127.0.0.1:9999/oauth/callback"
)

var (
	config *oauth2.Config
)

func authorize(c *cli.Context) error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %s", err.Error())
	}

	config = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		RedirectURL:  REDIRECT_URL,
		Scopes:       []string{"https://www.googleapis.com/auth/tasks"},
		Endpoint:     google.Endpoint,
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sslcli := &http.Client{Transport: tr}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	time.Sleep(1 * time.Second)
	open.Run(url)
	time.Sleep(1 * time.Second)
	log.Printf("Authentication URL: %s\n", url)

	http.HandleFunc("/oauth/callback", callbackHandler)

	server := http.Server{Addr: ":9999"}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)

	code := queryParts["code"][0]
	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("failed to exchage code: %s", err.Error())
	}

	if err := saveToken(credentialFile, token); err != nil {
		log.Fatalf("failed to save token: %s", err.Error())
	}

	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
	fmt.Fprintf(w, msg)
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var tok oauth2.Token
	err = json.NewDecoder(f).Decode(&tok)
	return &tok, err
}

func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(token)
}
