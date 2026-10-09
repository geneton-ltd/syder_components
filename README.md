# Syder Components

[![CI](https://github.com/geneton-ltd/syder_components/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/geneton-ltd/syder_components/actions/workflows/test.yaml)

Reusable, composable and testable frontend components compatible with the [gomponents](https://github.com/maragudk/gomponents) library.

## 1. Installation

Add the library to your project with:
```shell
GOPRIVATE=github.com/geneton-ltd/* go get github.com/geneton-ltd/syder_components
```

Copy the library's `static/output.css` file into your project's `static/css/syder_components.css`.

> [!WARNING]
> If you use SSH login for Github (you can check with `ssh -T git@github.com`), configure git to use SSH instead of HTTPS with `git config --global url."git@github.com:".insteadOf "https://github.com/"`.

## 2. Quickstart

Get started with this simple `main.go` file:
```go
package main

import (
	"embed"
	"net/http"

	sc "github.com/geneton-ltd/syder_components"
	co "github.com/geneton-ltd/syder_components/components"
)

//go:embed static/*
var static embed.FS

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.FileServer(http.FS(static)))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // pre-styled button from the components package
		btn := co.Button(sc.ID("helloWorldBtnID"), sc.Text("Hello, world!"))

		page := co.Page(
            co.PageProps{
				Title:    "Syder Components - Example",
				Language: "en",
				Head: []sc.Node{
					sc.Link(sc.Rel("stylesheet"), sc.Href("/static/css/syder_components.css")),
				},
				Body: []sc.Node{btn},
			},
		)
		page.Render(w)
	})

	http.ListenAndServe(":8000", mux)
}
```

## 3. Package Structure & Components 
The library is organized into two packages: 

### 1. `syder_components` (`sc`)
Contains base HTML primitives, attributes, text helpers, and control structures.
- **Attributes:** `Class`, `ID`, `Attr`, `Href`, `Src`, `Type`, `Disabled`, `Required`, `Data`, `Aria`, etc.
- **Text:** `Text(string)` (escaped HTML) and `Raw(string)` (raw unescaped HTML).
- **Control & Grouping:** `If`, `Iff`, `Group`, `GNode` (wrapper for `gomponents.Node` instances).
- **Unstyled HTML Elements:** Every standard non-deprecated HTML5 tag is available as an unstyled element (e.g., `sc.Div()`, `sc.Span()`, `sc.H1()`, `sc.Img()`). Obscure tags can be created with `sc.Any("tag", ...)`.

### 2. `syder_components/components` (`co`)
Contains pre-styled Tailwind CSS components:

| Component | Description |
| --------- | ----------- |
| Button | HTML5 `<button>` with pre-configured padding, round corners and hover states. |
| IconButton | `<button>` containing an icon wrapped in a fixed-size `<span>`. |
| IconOnlyButton | `IconButton` configured for standalone icon buttons. |
| Card | Styled `<div>` container with border, round corners and shadow. |
| SubCard | `Card` with a dashed border for nesting sub-content. |
| Dialog | Native HTML5 `<dialog>` modal with backdrop and shadow. |
| DialogRelative | `Dialog` positioned relative to its parent container. |
| DialogFixedCentered | `Dialog` fixed at the center of the viewport. |
| Navbar | Top navigation container `<div>`. |
| NavLink | Navigation element `<a>` designed for use inside `Navbar`. |
| Popover | Floating non-modal container using the native HTML5 Popover API. |
| TableContainer | Wrapper `<div>` adding horizontal scroll and borders for data tables. |
| Table | `<table>` element with pre-configured width and borders. |
| THead/TBody | Pre-styled table section containers. |
| Tr/Th/Td | Pre-styled table rows, header cells and standard cells. |

All HTML5 elements are also available with no default style in the `syder_components` package.

## 4. Core Concepts

### Everything is a `Node`
Every element, attribute, text fragment and group implements the `sc.Node` interface. Children and attributes are passed directly as variadic arguments:
```go
sc.Div(
    sc.ID("main"),
    sc.Class("p-4 bg-white"),
    sc.H1(sc.Text("Title")),
)
```

### Automatic Class Merging
You do not need a special `ClassAppend` option. Passing `sc.Class` multiple times automatically joins the class names with a space:
```go
// Renders class="p-4 bg-red-500 text-white"
sc.Div(
    sc.Class("p-4"),
    sc.Class("bg-red-500 text-white"),
)
```
For non-class attributes (like `id` or `type`), the last passed attribute overwrites previous ones.

### Groups and Unpacking
Multi-attributes or sibling elements can be bundled using `sc.Group(...)`. When passed to an element, groups are recursively unpacked:
```go
// DialogTrigger automatically unpacks into commandfor and command attributes
co.Button(
    sc.DialogTrigger("my-modal"),
    sc.Text("Open Modal"),
)
```

### Conditional Rendering (`If` and `Iff`)
- `sc.If(condition bool, node Node)` returns the node if `true` or `nil` if `false` (nil nodes are safely omitted during rendering).
- `sc.Iff(condition bool, fn func() Node)` lazily evaluates the function only if the condition is `true`, preventing nil-pointer panics when referencing nullable data.

### Testing without HTML Parsers
All `sc.Node` instances provide inspection methods for unit testing without parsing HTML strings:
- `node.Kind()`: returns `KindElement`, `KindText`, `KindAttr`, `KindGroup` or `KindGNode`.
- `node.Name()`: returns the HTML tag name (e.g. `"div"`) or attribute name.
- `node.Value()`: returns text nodes content or attribute value.
- `node.Attr("class")`: retrieves an attribute value by name.
- `node.Children()`: returns the slice of child elements.

## 5. Tailwind CSS Setup
Each component internally uses Tailwind CSS classes. To actually apply these styles in your project:

1. Copy `static/output.css` into your project directory (e.g., `static/css/syder_components.css`).
2. Import the stylesheet in your website's `<head>`.

## 6. Edge Cases & Native APIs

- `DialogRelative` requires parent container to have the `relative` CSS class.
- `Dialog` and `Popover` trigger options (`DialogTrigger`, `PopoverTrigger`, `DialogClose`, `PopoverClose`) use [native browser APIs](https://developer.mozilla.org/en-US/docs/Web/API/Invoker_Commands_API) (`commandfor`, `command`, `popovertarget`) and therefore require zero Javascript in modern browsers.

## 7. Development
Run the example server with live-reloading (requires [air](https://github.com/air-verse/air)):
```shell
make dev
```

Run test suite:
```shell
make test
```

Start Tailwind CSS rebuild watcher for the library (requires [Tailwind CLI](https://tailwindcss.com/docs/installation/tailwind-cli)):
```shell
make tw
```

Start Tailwind CSS rebuild watched for the example server (requires [Tailwind CLI](https://tailwindcss.com/docs/installation/tailwind-cli)):
```shell
make tw-dev
```

## License
Syder Components is available under the MIT license. See the LICENSE file for details.
