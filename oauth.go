package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)

	code := queryParts["code"][0]
	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("failed to exchage code: %s", err.Error())
	}

	if err := saveToken(credentialFileName, token); err != nil {
		log.Fatalf("failed to save token: %s", err.Error())
	}

	msg := `
<p>
	<strong>Success!</strong>
</p>"
<p>
	You are authenticated and can now return to the CLI.
</p>
`

	fmt.Fprintf(w, msg)
}
