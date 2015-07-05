package webhook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"strings"
)

// BodyMaxLen limits max size of payload.
var BodyMaxLen int64 = 1024 * 1024

// Header represents webhook's delivery headers.
type Header struct {
	EventType string
	Signature string
	Deliverty string
}

// Event represents raw event of Github's webhook.
type Event struct {
	Header   Header
	Body     []byte
	Verified bool
}

func parseHeader(r *http.Request) *Header {
	return &Header{
		EventType: r.Header.Get("X-Github-Event"),
		Signature: r.Header.Get("X-Hub-Signature"),
		Deliverty: r.Header.Get("X-Github-Delivery"),
	}
}

func split(r *http.Request) (*Header, []byte, error) {
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
	return h, body, nil
}

func verify(h hash.Hash, body, signature []byte) (bool, error) {
	h.Write(body)
	mac := h.Sum(nil)
	return hmac.Equal(mac, signature), nil
}

func verifySignature(header *Header, body, secret []byte) (bool, error) {
	s := header.Signature
	if s == "" {
		return false, nil
	}
	if strings.HasPrefix(s, "sha1=") {
		signature, err := hex.DecodeString(s[5:])
		if err != nil {
			return false, err
		}
		return verify(hmac.New(sha1.New, secret), body, signature)
	}
	return false, errors.New("unknown signature type")
}

// Parse parses a HTTP request as Github's webhook.
func Parse(r *http.Request, secret []byte) (*Event, error) {
	head, body, err := split(r)
	if err != nil {
		return nil, err
	}
	verified, err := verifySignature(head, body, secret)
	if err != nil {
		return nil, err
	}
	return &Event{
		Header:   *head,
		Body:     body,
		Verified: verified,
	}, nil
}
