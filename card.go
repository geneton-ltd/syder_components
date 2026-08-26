package sydercomponents

import (
	h "maragu.dev/gomponents/html"
)

// Card returns a styled <div> for grouping content.
func Card(opts ...Option) Component {
	n := newComponent("w-full rounded-md shadow-xs p-2 border border-1 border-gray-600", opts...)
	n.content = h.Div(n.manuallySetAttributes())
	return n
}

// Subcard returns a [Card] designed for grouping nested content.
// It applies a dashed border to visually separate it from primary cards.
func SubCard(opts ...Option) Component {
	options := append([]Option{WithClassAppend("border-dashed hover:bg-white")}, opts...)
	return Card(options...)
}
