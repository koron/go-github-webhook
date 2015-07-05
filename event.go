package webhook

import (
	"encoding/json"
)

// Repository represents a Github's repository.
type Repository struct {
	Name        string
	FullName    string
	Private     bool
	HTMLURL     string `json:"html_url"`
	Description string
	Fork        bool
	URL         string
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	PushedAt    int64  `json:"pushed_at"`
}

// Commit represents a Github's commit.
type Commit struct {
	ID        string
	Distinct  bool
	Message   string
	Timestamp string
	URL       string
	Added     []string
	Removed   []string
	Modified  []string
}

// PushEvent represents a Github's webhook push event.
type PushEvent struct {
	Ref        string
	Before     string
	After      string
	Created    bool
	Deleted    bool
	Forced     bool
	Compare    string
	Commits    []Commit
	Repository Repository
}

// PushEvent returns a PushEvent struct.
func (r *Event) PushEvent() *PushEvent {
	if r.Header.EventType != "push" {
		return nil
	}
	event := new(PushEvent)
	if err := json.Unmarshal(r.Body, event); err != nil {
		logErr(err)
		return nil
	}
	return event
}
