package components

import (
	sc "github.com/geneton-ltd/syder_components"
)

type PageProps struct {
	Title       string
	Description string
	Language    string
	HtmlAttrs   []sc.Node
	Head        []sc.Node
	Body        []sc.Node
}

func Page(pageProps PageProps) sc.Node {
	return sc.Group(
		sc.Raw("<!doctype html>"),
		sc.Html(
			sc.If(pageProps.Language != "", sc.Lang(pageProps.Language)),
			sc.Group(pageProps.HtmlAttrs...),
			sc.Head(
				sc.Meta(sc.Charset("utf-8")),
				sc.Meta(sc.Name("viewport"), sc.Content("width=device-width, initial-scale=1")),
				sc.Title(sc.Text(pageProps.Title)),
				sc.If(pageProps.Description != "", sc.Meta(sc.Name("description"), sc.Content(pageProps.Description))),
				sc.Group(pageProps.Head...),
			),
			sc.Body(sc.Group(pageProps.Body...)),
		),
	)
}
