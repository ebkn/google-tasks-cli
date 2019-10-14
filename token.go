package main

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

const credentialFileName = "credentials.json"

func loadToken() (*oauth2.Token, error) {
	file, err := os.Open(credentialFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tok oauth2.Token
	err = json.NewDecoder(file).Decode(&tok)
	return &tok, err
}

func saveToken(token *oauth2.Token) error {
	file, err := os.OpenFile(credentialFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(token)
}
