package components

import sc "github.com/geneton-ltd/syder_components"

// Button returns an HTML5 <button> element with default style without colors.
func Button(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("rounded rounded-sm shadow-xs p-1 hover:cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed")}, children...)
	return sc.Button(children...)
}

// IconButton returns an HTML5 <button> element containing an icon wrapped
// in a <span> element.
//
// Example:
//
//	components.IconButton(
//		sc.GNode(solid.ArrowLeft()),
//		sc.Children(sc.Text("Previous")),
//	)
func IconButton(icon sc.Node, children ...sc.Node) sc.Node {
	children = append(
		[]sc.Node{
			sc.Class("whitespace-nowrap flex gap-1 items-center"),
			sc.Span(sc.Class("h-4 w-4 shrink-0"), icon),
		},
		children...,
	)
	return Button(children...)
}

// IconOnlyButton returns a [IconButton] with special styles to make sure the
// icon placement is accurate.
//
// Example:
//
//	components.IconOnlyButton(
//		sc.GNode(solid.Trash())
//	)
func IconOnlyButton(icon sc.Node, children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("w-fit self-center"), sc.Type("button")}, children...)
	return IconButton(icon, children...)
}
