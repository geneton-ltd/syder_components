package main

import (
	"embed"
	"log"
	"net/http"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
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
	settingsModal := co.Popover(
		sc.ID(settingsModalID),
		sc.Class("!p-0 top-[50px] left-[150px] w-[150px] bg-white z-50 m-0 border border-gray-300"),
		sc.Div(
			sc.Class("flex flex-col h-full"),
			co.NavLink(
				sc.Class("h-full no-underline transition-colors p-2 hover:!bg-gray-300"),
				sc.Href("/settings/general"),
				sc.Text("General"),
			),
			co.NavLink(
				sc.Class("h-full no-underline transition-colors p-2 hover:!bg-gray-300"),
				sc.Href("/settings/passwords"),
				sc.Text("Passwords"),
			),
		),
	)

	navbar := co.Navbar(
		sc.ID("main-nav"),
		sc.Div(
			sc.Class("flex h-full"),
			co.NavLink(
				sc.Class("h-full no-underline transition-colors px-4"),
				sc.Href("/"),
				sc.Text("Dashboard"),
			),
			sc.Button(
				sc.Class("h-full hover:bg-gray-600 !rounded-none transition-colors px-4"),
				sc.Text("Settings"),
				sc.PopoverTrigger(settingsModalID),
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

	var rows []sc.Node
	for _, u := range users {
		rows = append(rows, sc.Tr(
			sc.Td(sc.Text(u.name)),
			sc.Td(sc.Text(u.role)),
			sc.Td(sc.Class("text-right"),
				sc.Button(
					sc.Class("bg-green-600 text-white hover:bg-green-700 px-2"),
					sc.Text("Edit"),
				),
			),
		))
	}

	modalID := "modal-id"
	dashboardCard := co.Card(
		sc.Class("bg-white p-3"),
		sc.Div(
			sc.Class("flex justify-between items-center mb-4"),
			sc.H2(sc.Class("text-xl font-bold"), sc.Text("Team Members")),
			sc.Button(
				sc.Class("bg-blue-600 text-white hover:bg-blue-700 px-2"),
				sc.DialogTrigger(modalID),
				sc.Text("+ Add User"),
			),
		),
		co.SubCard(
			co.TableContainer(
				sc.Table(
					sc.THead(
						sc.Tr(
							sc.Th(sc.Text("Name")),
							sc.Th(sc.Text("Role")),
							sc.Th(sc.Class("text-right"), sc.Text("Actions")),
						),
					),
					sc.TBody(rows...),
				),
			),
		),
	)

	addUserModal := co.DialogFixedCentered(
		sc.ID(modalID),
		sc.Class("w-full max-w-md p-6"),
		sc.Div(
			sc.Class("flex flex-col gap-4"),
			sc.H3(sc.Class("text-lg font-bold border-b pb-2"), sc.Text("Create New User")),
			sc.P(sc.Class("text-gray-600"), sc.Text("Lorem ipsum does not actually mean anything...")),
			sc.Button(
				sc.Class("bg-blue-600 text-white hover:bg-blue-700 px-2"),
				sc.Text("Close"),
				sc.DialogClose(modalID),
				sc.Type("button"),
			),
		),
	)

	content := sc.Div(
		sc.Class("w-full flex justify-center p-8 bg-gray-50 min-h-screen"),
		sc.Div(
			sc.Class("w-full max-w-5xl flex flex-col gap-4"),
			dashboardCard,
			addUserModal,
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
