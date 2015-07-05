package webhook

import "net/http"

var (
	secret []byte
)

func getSecret() []byte {
	return secret
}

// SetSecret set secret to verify webhook payload.
func SetSecret(v []byte) {
	secret = v
}

// HandlerFunc handles webhook events.
type HandlerFunc func(ev *Event)

// Handle generates http.HandlerFunc to handle webhook events.
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

// HandlePush generates http.HandlerFunc to handle webhook push events.
func HandlePush(f HandlerFunc) http.HandlerFunc {
	return filterHandle("push", f)
}
