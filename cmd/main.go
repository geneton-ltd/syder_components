package main

import (
	"embed"
	"net/http"

	sc "github.com/geneton-ltd/syder_components"
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

//go:embed static/*
var static embed.FS

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(static)))

	simpleNavbar := sc.Navbar(
		sc.WithID("navbarID"),
		sc.WithChildren(
			sc.Any("div",
				sc.WithClass("flex"),
				sc.WithChildren(
					sc.NavLink(
						sc.WithID("link1"),
						sc.WithHref("/page1"),
						sc.WithNodes(
							g.Text("Page1"),
						),
					),
					sc.NavLink(
						sc.WithID("link2"),
						sc.WithNodes(
							g.Text("Page2"),
							h.Href("/page2"),
						),
					),
				),
			),
		),
	)

	content := sc.Any("div",
		sc.WithClass("w-full flex justify-center p-4"),
		sc.WithChildren(
			sc.Any("div",
				sc.WithClass("max-w-6xl"),
				sc.WithChildren(
					sc.Card(
						sc.WithClassAppend("min-w-100"),
						sc.WithNodes(g.Text("nice")),
					),
				),
			),
		),
	)

	p := c.HTML5(c.HTML5Props{
		Head: []g.Node{
			h.Link(h.Href("/static/output.css"), h.Rel("stylesheet")),
		},
		Body: []g.Node{
			simpleNavbar,
			content,
		},
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p.Render(w)
	})

	http.ListenAndServe(":8000", mux)
}
