package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Toycon struct {
	About       string
	Description string
	Icon        *Image
	Image       *Image
	Name        string
}

var (
	toyconFn = [](func(s *goquery.Selection, t *Toycon)){
		getToyconAbout,
		getToyconDescription,
		getToyconIcon,
		getToyconImage,
		getToyconName}
)

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

func getToyconIcon(s *goquery.Selection, t *Toycon) {}

func getToyconImage(s *goquery.Selection, t *Toycon) {}

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

func newToycon(s *goquery.Selection) *Toycon {
	var (
		t = &Toycon{}
	)
	for _, fn := range toyconFn {
		fn(s, t)
	}
	return t
}

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
