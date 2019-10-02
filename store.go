package labo

type Store struct {
	Href   *Href       `json:"href"`
	Images []*Image    `json:"images"`
	Parts  []*Part     `json:"parts"`
	Video  interface{} `json:"video"`
}
