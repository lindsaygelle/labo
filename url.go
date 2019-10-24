package labo

import "net/url"

// URL is a snapshot of a URL that relates to a Nintendo Labo resource.
//
// URL's are built from a normal Go *url.URL but expresses
// the internal properties that are normally returned from
// the provided struct methods.
type URL struct {
	Address  string     `json:"address"`
	Fragment string     `json:"fragment"`
	Host     string     `json:"host"`
	Hostname string     `json:"hostname"`
	Path     string     `json:"path"`
	Port     string     `json:"port"`
	Scheme   string     `json:"scheme"`
	Query    url.Values `json:"query"`
}

func newURL(u *url.URL) *URL {
	return &URL{
		Address:  u.String(),
		Fragment: u.Fragment,
		Host:     u.Host,
		Hostname: u.Hostname(),
		Path:     u.Path,
		Port:     u.Port(),
		Query:    u.Query()}
}
