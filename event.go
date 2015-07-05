package webhook

import (
	"encoding/json"
)

type Repository struct {
	Name        string
	FullName    string
	Private     bool
	HTML_URL    string `json:"html_url"`
	Description string
	Fork        bool
	URL         string
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	PushedAt    int64  `json:"pushed_at"`
}

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

func (r *Event) PushEvent() (*PushEvent, error) {
	if r.Header.Event != "push" {
		return nil, nil
	}
	var event PushEvent
	if err := json.Unmarshal(r.Body, &event); err != nil {
		return nil, err
	}
	return &event, nil
}
