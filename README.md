# Components

Reusable, composable and testable frontend components to be used alongside the [gomponents](https://github.com/maragudk/gomponents) library.

## 1. Installation

Add the library to your project with:
```shell
GOPRIVATE=github.com/geneton-ltd/* go get github.com/geneton-ltd/syder_components
```

Copy the library's `static/output.css` file into your project's `static/css/syder_components.css`.

> [!WARNING]
> If you use SSH login for Github (you can check with `ssh -T git@github.com`), you need to tell git to use SSH instead of HTTPS with `git config --global url."git@github.com:".insteadOf "https://github.com/"`.

> [!NOTE] 
> This library is supposed to be used alongside gomponents. If you do not have them installed, you can do so with:
> `go get maragu.dev/gomponents`

## 2. Quickstart

You can get started with this simple `main.go` file (most of it is just boilerplate to get you started):
```go
package main

import (
	"embed"
	"net/http"

	sc "github.com/geneton-ltd/syder_components"
	g "maragu.dev/gomponents"
	co "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

//go:embed static/*
var static embed.FS

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(static)))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		btn := sc.Button(
			sc.WithID("helloWorldBtnID"),
			sc.WithNodes(g.Text("Hello, world!")),
		)

		page := co.HTML5(
			co.HTML5Props{
				Title:    "Syder Components - Example",
				Language: "en",
				Head: []g.Node{
					h.Link(h.Rel("stylesheet"), h.Href("/static/css/output.css")),
				},
				Body: []g.Node{btn},
			},
		)
		page.Render(w)
	})

	http.ListenAndServe(":8000", mux)
}
```

## 3. Components overview
The following is the list of all available components:
| Component | Description |
| --------- | ----------- |
| Any | Renders given HTML5 tag. It does not apply any styles. |
| Button | HTML5 `<button>` element with default style without colors. | 
| IconButton | HTML5 `<button>` element containing an icon wrapped in a `<span>` element. | 
| IconOnlyButton | *IconButton* with special styles to make sure the icon placement is accurate. |
| Card | A styled `<div>` for grouping content. |
| SubCard | A *Card* designed for grouping nested content. It applies a dashed border to visually separate it from primary cards. |
| Dialog | HTML5 `<dialog>` element. For attachment to happen, pass `WithID` option and then `WithDialogTrigger` with the same id to the trigger element (e.g. a button). Attachment is performed automatically. |
| DialogRelative | *Dialog* that is relative to its parent. The parent must have the "relative" class. |
| DialogFixedCentered | *Dialog* fixed at the center of the screen (horizontally and vertically). |
| Navbar | HTML5 `<div>` element with simple styles acting as a container for navigation elements. |
| NavLink | HTML5 `<a>` element, which is supposed to be put into *Navbar*. |
| Popover | HTML5 `<div>` element that floats natively using the Popover API. For attachment to happen, pass WithID option and then WithPopoverTrigger with the same id to the trigger element (e.g. a button). |
| TableContainer | HTML5 `<div>` element meant to function as a container for the *Table* component, while adding horizontal scroll. |
| Table | HTML5 `<table>` element with head and body separator. |
| THead | HTML5 `<thead>` element with no default style. |
| TBody | HTML5 `<tbody>` element with a thin row separator. |
| Tr | HTML5 `<tr>` element with simple hover style. |
| Th | HTML5 `<th>` element with bold font and default spacing. |
| Td | HTML5 `<td>` element with default spacing. |


## 4. Development
To start example code on port 8000 with live-reloading install [air](https://github.com/air-verse/air) and run:
```shell
make dev
```

Run all tests with:
```shell
make test
```

Finally, to rebuild Tailwind CSS install the [Tailwind CLI](https://tailwindcss.com/docs/installation/tailwind-cli) and start the watcher with:
```shell
make tw
```

## 5. Core concepts
The main concept of the library is the functional-options pattern (every component takes variadic options). We provide numerous `With*` options to set various HTML attributes. If some attribute is missing, simply use the `WithAttribute` option. Child content can be added with `WithNodes` (for *gomponents* nodes) or with `WithChildren` (for our `Component` nodes). Both of those options together with the `WithClass` option overwrite the nodes/children/class of the component. This can cause problems with nested default components or default components with predefined styles. For this the `WithNodeAppend`, `WithChildrenAppend` and `WithClassAppend` options are provided. Finally, one can track attributes of raw HTML5 by creating a `Any(tag)` `Component`. 

The advantage of this library is that one can access the attributes of a component (with the `Get*` methods), which allows easier testing without bloating the testing code with the HTML parsing library.

## 6. Customizing styles
Each component internally uses Tailwind CSS classes. Therefore, to actually apply these styles, copy `static/output.css` into your project directory and import the stylesheet in your website's `<head>`.

## 7. Edge cases
Here is the list of some things to have in mind when using some of the components:

- `DialogRelative` requires using `relative` class on its parent element.
- `Dialog` (and `Popover` component for closing) requires passing the element id to the trigger element with the `WithDialogTrigger` (or `WithPopoverTrigger`) options.
- `Dialog` and `Popover` elements rely on the `With*Trigger` options, which add the `command/commandfor/popovertarget/popovertargetaction` attributes, which rely on the [Invoker Commands API](https://developer.mozilla.org/en-US/docs/Web/API/Invoker_Commands_API). This API is supported in all modern browsers, but it is quite new and therefore might not be supported in all old browsers.

## License
Syder Components is available under the MIT license. See the LICENSE file for more info.
