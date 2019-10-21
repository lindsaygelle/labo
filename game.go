package labo

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Game is a snapshot of a Nintendo game that is compatible with the Nintendo Labo VR Kit.
type Game struct {
	Description string `json:"description"`
	Href        *Href  `json:"href"`
	Image       *Image `json:"image"`
	Logo        *Image `json:"logo"`
	Name        string `json:"name"`
	Title       string `json:"title"`
}

var (
	gameFn = [](func(*goquery.Selection, *Game)){
		getGameDescription,
		getGameHref,
		getGameImage,
		getGameLogo,
		getGameName,
		getGameTitle}
)

func getGameDescription(s *goquery.Selection, g *Game) {
	const (
		CSS string = ".games__item__content .games__item__body p"
	)
	var (
		description = stringNil
		ok          bool
		substring   string
	)
	s = s.Find(CSS)
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if ok {
		description = substring
	}
	g.Description = description
}

func getGameHref(s *goquery.Selection, g *Game) {
	const (
		CSS string = ".games__item__content .games__item__body a"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	g.Href = newHref(s)
}

func getGameImage(s *goquery.Selection, g *Game) {
	const (
		CSS string = ".games__item__video button .video-button__image picture img"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	g.Image = newImage(s)
}

func getGameLogo(s *goquery.Selection, g *Game) {
	const (
		CSS string = ".games__item__content .games__item__logo picture img"
	)
	var (
		ok bool
	)
	s = s.Find(CSS)
	ok = (s.Length() > 0)
	if !ok {
		return
	}
	g.Logo = newImage(s)
}

func getGameName(s *goquery.Selection, g *Game) {
	const (
		CSS string = ".games__item__content .games__item__body a"
	)
	var (
		name      = stringNil
		ok        bool
		substring string
	)
	s = s.Find(CSS)
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if ok {
		name = strings.ReplaceAll(substring, "Learn more about", "")
		name = strings.TrimSpace(name)
	}
	g.Name = name
}

func getGameTitle(s *goquery.Selection, g *Game) {
	const (
		CSS string = ".games__item__content .games__item__body h2"
	)
	var (
		ok        bool
		substring string
		title     = stringNil
	)
	s = s.Find(CSS)
	substring = strings.TrimSpace(s.Text())
	ok = (len(substring) > 0)
	if ok {
		title = substring
	}
	g.Title = title
}

func newGame(s *goquery.Selection) *Game {
	var (
		game = &Game{}
	)
	for _, fn := range gameFn {
		fn(s, game)
	}
	return game
}

func newGames(s *goquery.Selection) []*Game {
	var (
		game  *Game
		games []*Game
		ok    bool
	)
	s.Each(func(i int, s *goquery.Selection) {
		game = newGame(s)
		ok = (game != nil)
		if !ok {
			return
		}
		games = append(games, game)
	})
	return games
}
