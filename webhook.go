package webhook

import (
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var BodyMaxLen int64 = 1024 * 1024

// Header represents webhook's delivery headers.
type Header struct {
	Event     string
	Signature string
	Deliverty string
}

// Payload represents webhook's payload.
type Payload struct {
}

func parseHeader(r *http.Request) *Header {
	return &Header{
		Event:     r.Header.Get("X-Github-Event"),
		Signature: r.Header.Get("X-Hub-Signature"),
		Deliverty: r.Header.Get("X-Github-Delivery"),
	}
}

func Split(r *http.Request) (*Header, []byte, error) {
	if r.ContentLength == 0 || r.Body == nil {
		return nil, nil, errors.New("no body")
	} else if r.ContentLength > BodyMaxLen {
		return nil, nil, fmt.Errorf("too big body: %d > %d",
			r.ContentLength, BodyMaxLen)
	}
	defer r.Body.Close()
	h := parseHeader(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, nil, err
	}
	// TODO:
	return h, body, nil
}

func verify(key, signature, message []byte) bool {
	hash := hmac.New(sha1.New, key)
	hash.Write(message)
	mac := hash.Sum(nil)
	return hmac.Equal(signature, mac)
}

// Parse parses a HTTP request as Github's webhook.
func Parse(r *http.Request) {
	//h := parseHeader(r)
}
