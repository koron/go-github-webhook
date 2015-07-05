package webhook

import (
	"log"
	"net/http"
	"os"
)

var (
	Secret []byte
	Logger *log.Logger = log.New(os.Stderr, "", log.LstdFlags)
)

func secret() []byte {
	return Secret
}

func logErr(err error) {
	if Logger == nil {
		return
	}
	Logger.Printf("webhoook error: %s", err)
	return
}

type HandlerFunc func(raw *Event)

func PushHandler(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := Parse(r, secret())
		if err != nil {
			logErr(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		if raw.Header.Event != "push" {
			return
		}
		f(raw)
	}
}
