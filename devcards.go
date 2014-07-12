package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/extemporalgenome/slug"
	"github.com/go-martini/martini"
)

type Devcards struct {
	Name     string
	children []Devcards
	cards    []Card
}

type Card interface {
	Title() string
}

func (d *Devcards) Serve() {
	m := martini.Classic()
	d.serve(m, "/")
	m.Run()
}

func (d Devcards) serve(m *martini.ClassicMartini, baseUrl string) {
	url := path.Join(baseUrl, slug.Slug(d.Name))
	m.Get(url, martiniHandler(&d, url))
	for _, child := range d.children {
		child.serve(m, url)
	}
}

func martiniHandler(d *Devcards, url string) func(http.ResponseWriter) {
	return func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Printf("Rendering %s for url %s\n", d.Name, url)
		fmt.Fprintf(w, d.render(url))
	}
}

func (d *Devcards) render(baseUrl string) string {
	s := "Devcards: " + d.Name
	for _, child := range d.children {
		s += "<br><a href=\"" + path.Join(baseUrl, slug.Slug(child.Name), "/") + "\">" + child.Name + "</a>"
	}
	for _, card := range d.cards {
		s += "<br>" + card.Title()
	}
	return s
}

func (d *Devcards) AddChild(child Devcards) {
	d.children = append(d.children, child)
}
