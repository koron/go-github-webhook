package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/koron/go-github-webhook"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	raw, err := webhook.Parse(r, []byte("abcdefgh"))
	if err != nil {
		fmt.Printf("  error=%#v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("  raw.Header=%#v\n", raw.Header)
	fmt.Printf("  raw.Verified=%#v\n", raw.Verified)
	w.WriteHeader(http.StatusOK)
}

func main() {
	err := http.ListenAndServe(":8080", &handler{})
	if err != nil {
		log.Fatal(err)
	}
}
