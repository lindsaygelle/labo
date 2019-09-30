package labo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

// KitVR is an extension of a Nintendo Labo Kit that supports VR functionality.
type KitVR struct {
	*Kit
	Plaza interface{}
}

// NewKitVR is a constructor function that instantiates and returns a new KitVR struct pointer.
func NewKitVR(d *goquery.Document) (*KitVR, error) {
	var (
		ok bool
	)
	ok = (d != nil)
	if !ok {
		return nil, fmt.Errorf(errorGoQueryDocumentNil)
	}
	s := d.Find("html")
	ok = (s != nil)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionNil)
	}
	ok = (s.Length() > 0)
	if !ok {
		return nil, fmt.Errorf(errorGoQuerySelectionEmptyHTMLNodes, s)
	}
	kitVR := KitVR{
		Kit: &Kit{}}
	return &kitVR, nil
}
