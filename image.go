package labo

type Image struct {
	Alt  string `json:"alt"`
	Href *Href  `json:"href"`
	Src  string `json:"src"`
}
