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

	settingsModalID := "settingsMenu"
	settingsModal := sc.Popover(
		sc.ID(settingsModalID),
		sc.ClassAppend("!p-0 top-[50px] left-[150px] w-[150px] bg-white z-50 m-0 border border-gray-300"),
		sc.Children(
			sc.Div(
				sc.Class("flex flex-col h-full"),
				sc.Children(
					sc.NavLink(
						sc.ClassAppend("h-full no-underline transition-colors p-2 hover:!bg-gray-300"),
						sc.Href("/settings/general"),
						sc.Nodes(g.Text("General")),
					),
					sc.NavLink(
						sc.ClassAppend("h-full no-underline transition-colors p-2 hover:!bg-gray-300"),
						sc.Href("/settings/passwords"),
						sc.Nodes(g.Text("Passwords")),
					),
				),
			),
		),
	)

	navbar := sc.Navbar(
		sc.ID("main-nav"),
		sc.Children(
			sc.Div(
				sc.Class("flex h-full"),
				sc.Children(
					sc.NavLink(
						sc.ClassAppend("h-full no-underline transition-colors px-4"),
						sc.Href("/"),
						sc.Nodes(
							g.Text("Dashboard"),
						),
					),
					sc.Button(
						sc.ClassAppend("h-full hover:bg-gray-600 !rounded-none transition-colors px-4"),
						sc.Nodes(g.Text("Settings")),
						sc.PopoverTrigger(settingsModalID),
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
		rows = append(rows, sc.Tr(sc.Children(
			sc.Td(sc.Nodes(g.Text(u.name))),
			sc.Td(sc.Nodes(g.Text(u.role))),
			sc.Td(sc.ClassAppend("text-right"), sc.Children(
				sc.Button(
					sc.ClassAppend("bg-green-600 text-white hover:bg-green-700 px-2"),
					sc.Nodes(g.Text("Edit")),
				),
			),
			))))
	}

	modalID := "modal-id"
	dashboardCard := sc.Card(
		sc.ClassAppend("bg-white p-3"),
		sc.Children(
			sc.Div(
				sc.Class("flex justify-between items-center mb-4"),
				sc.Nodes(
					h.H2(h.Class("text-xl font-bold"), g.Text("Team Members")),
				),
				sc.Children(
					sc.Button(
						sc.ClassAppend("bg-blue-600 text-white hover:bg-blue-700 px-2"),
						sc.DialogTrigger(modalID),
						sc.Nodes(g.Text("+ Add User")),
					),
				),
			),
			sc.SubCard(sc.Children(
				sc.TableContainer(sc.Children(
					sc.Table(sc.Children(
						sc.THead(sc.Children(
							sc.Tr(sc.Children(
								sc.Th(sc.Nodes(g.Text("Name"))),
								sc.Th(sc.Nodes(g.Text("Role"))),
								sc.Th(sc.ClassAppend("text-right"), sc.Nodes(g.Text("Actions"))),
							)),
						)),
						sc.TBody(sc.Children(rows...)),
					)),
				)),
			)),
		),
	)

	addUserModal := sc.DialogFixedCentered(
		sc.ID(modalID),
		sc.ClassAppend("w-full max-w-md p-6"),
		sc.Children(
			sc.Div(
				sc.Class("flex flex-col gap-4"),
				sc.Nodes(
					h.H3(h.Class("text-lg font-bold border-b pb-2"), g.Text("Create New User")),
					h.P(h.Class("text-gray-600"), g.Text("Lorem ipsum does not actually mean anything...")),
				),
				sc.Children(sc.Button(
					sc.ClassAppend("bg-blue-600 text-white hover:bg-blue-700 px-2"),
					sc.Nodes(g.Text("Close")),
					sc.DialogClose(modalID),
					sc.Type("button"),
				)),
			),
		),
	)

	content := sc.Div(
		sc.Class("w-full flex justify-center p-8 bg-gray-50 min-h-screen"),
		sc.Children(
			sc.Div(
				sc.Class("w-full max-w-5xl flex flex-col gap-4"),
				sc.Children(
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
			settingsModal,
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
