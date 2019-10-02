package labo

import "net/url"

type Href struct {
	Link   string   `json:"link"`
	Target string   `json:"target"`
	URL    *url.URL `json:"URL"`
}
