package main

import (
	"log"
	"net/http"

	"github.com/koron/go-github-webhook"
)

func main() {
	webhook.SetSecret([]byte("abcdefgh"))
	http.HandleFunc("/", webhook.HandlePush(func(ev *webhook.Event) {
		push := ev.PushEvent()
		if push == nil {
			return
		}
		log.Printf("push: verified=%v %#v", ev.Verified, push)
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
