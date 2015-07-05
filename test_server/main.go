package main

import (
	"log"
	"net/http"

	"github.com/koron/go-github-webhook"
)

func main() {
	// To verify webhook's payload, set secret by SetSecret().
	webhook.SetSecret([]byte("abcdefgh"))

	// Add a HandlerFunc to process webhook.
	http.HandleFunc("/", webhook.HandlePush(func(ev *webhook.Event) {
		push := ev.PushEvent()
		if push == nil {
			return
		}
		log.Printf("push: verified=%v %#v", ev.Verified, push)
	}))

	// Start web server.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
