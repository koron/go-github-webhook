package main

import (
	"fmt"
	"log"
	"net/http"
	".."
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL.Path)
	head, _, err := webhook.Split(r)
	fmt.Printf("  header=%#v\n", head)
	fmt.Printf("  error=%#v\n", err)
	w.WriteHeader(http.StatusOK)
}

func main() {
	err := http.ListenAndServe(":8080", &handler{})
	if err != nil {
		log.Fatal(err)
	}
}
