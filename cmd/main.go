package main

import (
	"embed"
	"log"
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

	navbar := sc.Navbar(
		sc.WithID("main-nav"),
		sc.WithChildren(
			sc.Any("div",
				sc.WithClass("flex h-full"),
				sc.WithChildren(
					sc.NavLink(
						sc.WithClassAppend("h-full no-underline transition-colors px-4"),
						sc.WithHref("/"),
						sc.WithNodes(
							g.Text("Dashboard"),
						),
					),
					sc.NavLink(
						sc.WithClassAppend("h-full no-underline transition-colors px-4"),
						sc.WithHref("/settings"),
						sc.WithNodes(g.Text("Settings")),
					),
				),
			),
		),
	)

	users := []struct {
		name string
		role string
	}{
		{
			name: "Joey Chadwick",
			role: "Reviewer",
		},
		{
			name: "Maggie Peterson",
			role: "Evaluator",
		},
		{
			name: "Walter Kramer",
			role: "Admin",
		},
	}

	var rows []sc.Component
	for _, u := range users {
		rows = append(rows, sc.Tr(sc.WithChildren(
			sc.Td(sc.WithNodes(g.Text(u.name))),
			sc.Td(sc.WithNodes(g.Text(u.role))),
			sc.Td(sc.WithClassAppend("text-right"), sc.WithChildren(
				sc.Button(
					sc.WithClassAppend("bg-green-600 text-white hover:bg-green-700 px-2"),
					sc.WithNodes(g.Text("Edit")),
				),
			),
			))))
	}

	modalID := "modal-id"
	dashboardCard := sc.Card(
		sc.WithClassAppend("bg-white p-3"),
		sc.WithChildren(
			sc.Any("div",
				sc.WithClass("flex justify-between items-center mb-4"),
				sc.WithNodes(
					h.H2(h.Class("text-xl font-bold"), g.Text("Team Members")),
				),
				sc.WithChildren(
					sc.Button(
						sc.WithClassAppend("bg-blue-600 text-white hover:bg-blue-700 px-2"),
						sc.WithDialogTrigger(modalID),
						sc.WithNodes(g.Text("+ Add User")),
					),
				),
			),
			sc.SubCard(sc.WithChildren(
				sc.TableContainer(sc.WithChildren(
					sc.Table(sc.WithChildren(
						sc.THead(sc.WithChildren(
							sc.Tr(sc.WithChildren(
								sc.Th(sc.WithNodes(g.Text("Name"))),
								sc.Th(sc.WithNodes(g.Text("Role"))),
								sc.Th(sc.WithClassAppend("text-right"), sc.WithNodes(g.Text("Actions"))),
							)),
						)),
						sc.TBody(sc.WithChildren(rows...)),
					)),
				)),
			)),
		),
	)

	addUserModal := sc.DialogFixedCentered(
		sc.WithID(modalID),
		sc.WithClassAppend("w-full max-w-md p-6"),
		sc.WithChildren(
			sc.Any(
				"div",
				sc.WithClass("flex flex-col gap-4"),
				sc.WithNodes(
					h.H3(h.Class("text-lg font-bold border-b pb-2"), g.Text("Create New User")),
					h.P(h.Class("text-gray-600"), g.Text("Lorem ipsum does not actually mean anything...")),
				),
				sc.WithChildren(sc.Button(
					sc.WithClassAppend("bg-blue-600 text-white hover:bg-blue-700 px-2"),
					sc.WithNodes(g.Text("Close")),
					sc.WithDialogClose(modalID),
					sc.WithType("button"),
				)),
			),
		),
	)

	content := sc.Any("div",
		sc.WithClass("w-full flex justify-center p-8 bg-gray-50 min-h-screen"),
		sc.WithChildren(
			sc.Any("div",
				sc.WithClass("w-full max-w-5xl flex flex-col gap-4"),
				sc.WithChildren(
					dashboardCard,
					addUserModal,
				),
			),
		),
	)

	p := c.HTML5(c.HTML5Props{
		Title: "Syder Components - Example",
		Head: []g.Node{
			h.Link(h.Href("/static/output.css"), h.Rel("stylesheet")),
		},
		Body: []g.Node{
			navbar,
			content,
		},
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p.Render(w)
	})

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
