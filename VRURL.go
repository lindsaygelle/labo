package labo

import "net/url"

// VRURL is a string that points to a Nintendo Labo VR Kit.
type VRURL string

func (VRURL VRURL) String() string {
	return string(VRURL)
}

// URL is a function that parses the Nintendo Labo VR Kit URL into a *url.URL.
func (VRURL VRURL) URL() (*url.URL, error) {
	return url.Parse(VRURL.String())
}
