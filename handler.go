package webhook

import "net/http"

var (
	secret []byte
)

func getSecret() []byte {
	return secret
}

func SetSecret(v []byte) {
	secret = v
}

type HandlerFunc func(ev *Event)

func Handle(f HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ev, err := Parse(r, getSecret())
		if err != nil {
			logErr(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		f(ev)
	}
}

func filterHandle(eventType string, f HandlerFunc) http.HandlerFunc {
	return Handle(func(ev *Event) {
		if ev.Header.EventType != eventType {
			return
		}
		f(ev)
	})
}

func HandlePush(f HandlerFunc) http.HandlerFunc {
	return filterHandle("push", f)
}
