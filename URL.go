package labo

import "net/url"

// URL is a string that points to a Nintendo Labo Kit URL.
type URL string

func (URL URL) String() string {
	return string(URL)
}

// URL is a function that parses the Nintendo Labo Kit URL into a *url.URL.
func (URL URL) URL() (*url.URL, error) {
	return url.Parse(URL.String())
}
