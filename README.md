# Syder Components

[![CI](https://github.com/geneton-ltd/syder_components/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/geneton-ltd/syder_components/actions/workflows/test.yaml)

Reusable, composable and testable frontend components built on top of the [gomponents](https://github.com/maragudk/gomponents) library.

## 1. Installation

Add the library to your project with:
```shell
GOPRIVATE=github.com/geneton-ltd/* go get github.com/geneton-ltd/syder_components
```

Copy the library's `static/output.css` file into your project's `static/css/syder_components.css`.

> [!WARNING]
> If you use SSH login for Github (you can check with `ssh -T git@github.com`), you need to tell git to use SSH instead of HTTPS with `git config --global url."git@github.com:".insteadOf "https://github.com/"`.

> [!NOTE] 
> Some top level components are yet to be added to the library. For that you will sometimes import the `gomponents` library. 

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
			sc.ID("helloWorldBtnID"),
			sc.Children(sc.Text("Hello, world!")),
		)

		page := co.HTML5(
			co.HTML5Props{
				Title:    "Syder Components - Example",
				Language: "en",
				Head: []g.Node{
					h.Link(h.Rel("stylesheet"), h.Href("/static/css/syder_components.css")),
				},
				Body: []g.Node{btn},
			},
		)
		page.Render(w)
	})

	http.ListenAndServe(":8000", mux)
}
```

## 3. Components Overview
We provide a number of styled components available in the `syder_components/components` package. They are listed below:

| Component | Description |
| --------- | ----------- |
| Any | Renders given HTML5 tag. It does not apply any styles. |
| Button | HTML5 `<button>` element with default style without colors. |
| IconButton | HTML5 `<button>` element containing an icon wrapped in a `<span>` element. |
| IconOnlyButton | `IconButton` with special styles to make sure the icon placement is accurate. |
| Card | A styled `<div>` for grouping content. |
| Dialog | HTML5 `<dialog>` element. For attachment to happen, pass `ID` option and then `DialogTrigger` with the same id to the trigger element (e.g. a button). Attachment is performed automatically. |
| DialogFixedCentered | `Dialog` fixed at the center of the screen (horizontally and vertically). |
| DialogRelative | `Dialog` that is relative to its parent. The parent must have the "relative" class. |
| Navbar | HTML5 `<div>` element with simple styles acting as a container for navigation elements. |
| NavLink | HTML5 `<a>` element, which is supposed to be put into `Navbar`. |
| Popover | HTML5 `<div>` element that floats natively using the Popover API. For attachment to happen, pass `ID` option and then `PopoverTrigger` with the same id to the trigger element (e.g. a button). |
| Raw | Raw creates a text Component that renders the unescaped string t. |
| SubCard | A `Card` designed for grouping nested content. It applies a dashed border to visually separate it from primary cards. |
| Table | HTML5 `<table>` element with head and body separator. |
| TableContainer | HTML5 `<div>` element meant to function as a container for the `Table` component, while adding horizontal scroll. |
| Text | Text creates a text Component that renders the escaped string t. |
| THead | HTML5 `<thead>` element with no default style. |
| TBody | HTML5 `<tbody>` element with a thin row separator. |
| Tr | HTML5 `<tr>` element with simple hover style. |
| Th | HTML5 `<th>` element with bold font and default spacing. |
| Td | HTML5 `<td>` element with default spacing. |

All HTML5 elements are also available with no default style in the `syder_components` package.

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

## 5. Core Concepts
- The main concept of the library is the functional-options pattern (every component takes variadic options). We provide numerous attribute options to set various HTML attributes. If some attribute is missing, simply use the `Attribute` option.
- Child content can be added with the `Children` option. The `Children` option together with the `Class` option overwrite the children/class of the component. This can cause problems with nested default components or default components with predefined styles. For this the `ChildrenAppend` and `ClassAppend` options are provided.
- Escaped text nodes can be created with the `Text` component, which can be tested with the `GetText` method. For unescaped text, use the `Raw` component.
- To add `gomponents.Node` to children, wrap it inside the `Node` component. Note that this does not mean that it will be testable. It is in place just to allow adding `gomponents.Node` as children to `Component`.
- For conditional rendering, the `If(condition bool, component Component)` function is available. If the condition is `false`, the function returns the `Empty` component. For lazy evaluation, the `Iff(condition bool, f func() Component) Component` function is available.
- To track attributes of raw HTML5 you can use the `Any` component.

The advantage of this library is that one can access the attributes of a component (with the `Get*` methods), which allows easier testing without bloating the testing code with the HTML parsing library.

## 6. Tailwind CSS Setup
Each component internally uses Tailwind CSS classes. Therefore, to actually apply these styles:

1. Copy `static/output.css` into your project directory, for example to `static/css/syder_components.css`.
2. Import the stylesheet in your website's `<head>`.

## 7. Edge Cases
Here is the list of some things to have in mind when using some of the components:

- `DialogRelative` requires using `relative` class on its parent element.
- `Dialog` (and `Popover` component for closing) requires passing the element id to the trigger element with the `DialogTrigger` (or `PopoverTrigger`) options.
- `Dialog` and `Popover` elements rely on the `DialogTrigger/PopoverTrigger/DialogClose/PopoverClose` options, which add the `command/commandfor/popovertarget/popovertargetaction` attributes, which rely on the [Invoker Commands API](https://developer.mozilla.org/en-US/docs/Web/API/Invoker_Commands_API). This API is supported in all modern browsers, but it is quite new and therefore might not be supported in all old browsers.

## License
Syder Components is available under the MIT license. See the LICENSE file for more info.
