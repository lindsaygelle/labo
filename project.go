package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Project struct {
	Icon        *Image   `json:"icon"`
	Image       *Image   `json:"image"`
	Name        string   `json:"name"`
	Screenshots []*Image `json:"screenshots"`
}

var (
	projectFn = [](func(s *goquery.Selection, p *Project)){
		getProjectIcon,
		getProjectImage,
		getProjectName,
		getProjectScreenshots}
)

func getProjectIcon(s *goquery.Selection, p *Project) {
	const (
		CSS string = ".toycon-icon:nth-child(1) > .icon:nth-child(1) > img:nth-child(1)"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	p.Icon = newImage(s)
}

func getProjectImage(s *goquery.Selection, p *Project) {
	const (
		CSS string = ".toycon-image:nth-child(1) > picture:nth-child(1) > img:nth-child(1)"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	p.Image = newImage(s)
}

func getProjectName(s *goquery.Selection, p *Project) {
	const (
		CSS string = ".toycon-icon:nth-child(1) > p:nth-child(1)"
	)
	var (
		name string
		ok   bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if ok {
		name = strings.TrimSpace(s.Text())
	}
	p.Name = name
}

func getProjectScreenshots(s *goquery.Selection, p *Project) {
	const (
		CSS string = ".toycon-icon:nth-child(1) > .screenshot"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	p.Screenshots = newImages(s)
}

func newProject(s *goquery.Selection) *Project {
	var (
		p = &Project{}
	)
	for _, fn := range projectFn {
		fn(s, p)
	}
	return p
}

func newProjects(s *goquery.Selection) []*Project {
	var (
		ok       bool
		project  *Project
		projects []*Project
	)
	s.Each(func(i int, s *goquery.Selection) {
		project = newProject(s)
		ok = (project != nil)
		if !ok {
			return
		}
		projects = append(projects, project)
	})
	return projects
}
