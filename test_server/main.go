package main

import (
	"log"
	"net/http"

	"github.com/koron/go-github-webhook"
)

func main() {
	http.HandleFunc("/", webhook.PushHandler(func(raw *webhook.Event) {
		ev, err := raw.PushEvent()
		if err != nil {
			return
		}
		log.Printf("%#v", ev)
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
