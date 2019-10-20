package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Project is a snapshot of a Nintendo Labo Kit project that can be built with the contents of a Nintendo Labo kit.
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

// getProjectIcon searches the *goquery.Selection for the *labo.Image required for a labo.Project
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

// getProjectIcon searches the *goquery.Selection for the *labo.Image required for a labo.Project
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

// getProjectName searches the *goquery.Selection for the name of the project required for a labo.Project
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

// getProjectScreenshots searches the *goquery.Selection for the slice of *labo.Image required for a labo.Project
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

// newProject is a constructor function that instantiates and returns a new *labo.Project.
func newProject(s *goquery.Selection) *Project {
	var (
		p = &Project{}
	)
	for _, fn := range projectFn {
		fn(s, p)
	}
	return p
}

// newProjects is a constructor function that instantiates and returns a new slice of *labo.Project.
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
