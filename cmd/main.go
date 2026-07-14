package main

import (
	"embed"
	"net/http"

	sydercomponents "github.com/geneton-ltd/syder_components"
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

//go:embed static/*
var static embed.FS

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(static)))

	simpleNavbar := sydercomponents.Navbar(
		sydercomponents.WithID("navbarID"),
		sydercomponents.WithChildren(
			h.Div(
				h.Class("flex"),
				sydercomponents.NavLink(
					sydercomponents.WithID("link1"),
					sydercomponents.WithChildren(
						g.Text("Page1"),
						h.Href("/page1"),
					),
				),
				sydercomponents.NavLink(
					sydercomponents.WithID("link2"),
					sydercomponents.WithChildren(
						g.Text("Page2"),
						h.Href("/page2"),
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
		},
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p.Render(w)
	})

	http.ListenAndServe(":8000", mux)
}
