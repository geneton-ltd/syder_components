package components

import sc "github.com/geneton-ltd/syder_components"

// TableContainer returns a HTML5 <div> element. It is meant
// to function as a container for the [Table] component, while
// adding horizontal scroll.
//
// Example:
//
//	co.TableContainer(
//		co.Table(
//			co.THead(
//				co.Tr(
//					co.Th(sc.Text("Name")),
//					co.Th(sc.Text("Email")),
//				),
//			),
//			co.TBody(
//				co.Tr(
//					co.Td(sc.Text("John")),
//					co.Td(sc.Text("john.doe@google.com")),
//				),
//				co.Tr(
//					co.Td(sc.Text("Mat")),
//					co.Td(sc.Text("mat.pat@youtube.com")),
//				),
//			),
//		),
//	)
func TableContainer(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("overflow-x-auto rounded border border-gray-300 shadow-sm")}, children...)
	return sc.Div(children...)
}

// Table returns a HTML5 <table> element with head and body separator.
func Table(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("min-w-full divide-y-2 divide-gray-200")}, children...)
	return sc.Table(children...)
}

// THead returns a HTML5 <thead> element with no default style.
func THead(children ...sc.Node) sc.Node {
	return sc.THead(children...)
}

// TBody returns a HTML5 <tbody> element with a thin row separator.
func TBody(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("divide-y divide-gray-200")}, children...)
	return sc.TBody(children...)
}

// Tr returns a HTML5 <tr> element with simple hover style.
func Tr(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("hover:bg-gray-100")}, children...)
	return sc.Tr(children...)
}

// Th returns a HTML5 <th> element with bold font and default spacing.
func Th(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("px-3 py-2 whitespace-nowrap font-medium text-gray-900")}, children...)
	return sc.Th(children...)
}

// Td returns a HTML5 <td> element with default spacing.
func Td(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("px-3 py-2 whitespace-nowrap text-gray-900")}, children...)
	return sc.Td(children...)
}
