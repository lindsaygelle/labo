package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Toycon is a snapshot of a Nintendo Labo Toycon.
//
// Toycon's are the cardboard products built from a Nintendo Labo kit. Each Toycon
// contains a series of features that are unique to the Nintendo Labo Kit.
type Toycon struct {
	About       string     `json:"about"`
	Description string     `json:"description"`
	Features    []*Feature `json:"features"`
	Icon        *Image     `json:"icon"`
	Image       *Image     `json:"image"`
	Name        string     `json:"name"`
}

var (
	toyconFn = [](func(s *goquery.Selection, t *Toycon)){
		getToyconAbout,
		getToyconDescription,
		getToyconFeatures,
		getToyconIcon,
		getToyconImage,
		getToyconName}
)

// getToyconAbout searches the *goquery.Selection for the about string required for a labo.Toycon
func getToyconAbout(s *goquery.Selection, t *Toycon) {
	const (
		CSS string = ".toy-con-sub-header p"
	)
	var (
		ok        bool
		substring string
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	t.About = substring
}

// getToyconDescription searches the *goquery.Selection for the description required for a labo.Toycon
func getToyconDescription(s *goquery.Selection, t *Toycon) {
	const (
		CSS string = ".right-column .toy-con-info .copy p"
	)
	var (
		ok        bool
		substring string
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	t.Description = substring
}

// getToyconFeatures searches the *goquery.Selection for the *labo.Feature slice required for a labo.Toycon
func getToyconFeatures(s *goquery.Selection, t *Toycon) {
	const (
		CSS string = ".left-column .toy-con-slider"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	t.Features = newFeatures(s)
}

// getToyconIcon searches the *goquery.Selection for the icon *labo.Image struct required for a labo.Toycon
func getToyconIcon(s *goquery.Selection, t *Toycon) {
	const (
		CSS string = ".right-column .toy-con-info .icon > img:nth-child(1)"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	t.Icon = newImage(s)
}

// getToyconImage searches the *goquery.Selection for the *labo.Image struct required for a labo.Toycon
func getToyconImage(s *goquery.Selection, t *Toycon) {
	const (
		CSS string = ".right-column .main-image picture:nth-child(1) > img:nth-child(1)"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	t.Image = newImage(s)
}

// getToyconName searches the *goquery.Selection for the name of the Toycon required for a labo.Toycon
func getToyconName(s *goquery.Selection, t *Toycon) {
	const (
		CSS string = ".toy-con-header h3"
	)
	var (
		ok        bool
		substring string
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if !ok {
		return
	}
	t.Name = substring
}

// newToycon is a constructor function that instantiates and returns a new *labo.Toycon.
func newToycon(s *goquery.Selection) *Toycon {
	var (
		t = &Toycon{}
	)
	for _, fn := range toyconFn {
		fn(s, t)
	}
	return t
}

// newToycons is a constructor function that instantiates and returns a new slice of *labo.Toycon.
func newToycons(s *goquery.Selection) []*Toycon {
	var (
		toycon  *Toycon
		toycons []*Toycon
		ok      bool
	)
	s.Each(func(i int, s *goquery.Selection) {
		toycon = newToycon(s)
		ok = (toycon != nil)
		if !ok {
			return
		}
		toycons = append(toycons, toycon)
	})
	return toycons
}
