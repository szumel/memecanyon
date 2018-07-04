package main

import (
	"log"
	"net/http"

	"github.com/szumel/memecanyon/cmd/http/endpoint/meme"
)

func main() {
	http.HandleFunc("/v1/meme", meme.ListCollection)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
