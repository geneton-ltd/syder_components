package components

import sc "github.com/geneton-ltd/syder_components"

// Card returns a styled <div> for grouping content.
func Card(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("w-full rounded-md shadow-xs p-2 border border-1 border-gray-600")}, children...)
	return sc.Div(children...)
}

// SubCard returns a [Card] designed for grouping nested content.
// It applies a dashed border to visually separate it from primary cards.
func SubCard(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("border-dashed hover:bg-white")}, children...)
	return Card(children...)
}
