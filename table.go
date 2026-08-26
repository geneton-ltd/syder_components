package sydercomponents

import (
	h "maragu.dev/gomponents/html"
)

// TableContainer returns a HTML <div> element. It is meant
// to function as a container for the [Table] component, while
// adding horizontal scroll.
//
// Example:
//
//	sc.TableContainer(
//	  sc.Table(
//	    sc.THead(
//	      sc.Tr(sc.Th("Name"), sc.Th("Email"))
//	    ),
//	    sc.TBody(
//	      sc.Tr(sc.Td("John"), sc.Td("john.doe@google.com")),
//	      sc.Tr(sc.Td("Mat"), sc.Td("mat.pat@youtube.com")),
//	    ),
//	  )
//	)
func TableContainer(opts ...Option) Component {
	n := newComponent("overflow-x-auto rounded border border-gray-300 shadow-sm", opts...)
	n.content = h.Div(n.manuallySetAttributes())
	return n
}

// Table returns a HTML <table> element with head and body separator.
func Table(opts ...Option) Component {
	n := newComponent("min-w-full divide-y-2 divide-gray-200", opts...)
	n.content = h.Table(n.manuallySetAttributes())
	return n
}

// THead returns a HTML <thead> element with no default style.
func THead(opts ...Option) Component {
	n := newComponent("", opts...)
	n.content = h.THead(n.manuallySetAttributes())
	return n
}

// TBody returns a HTML <tbody> element with a thin row separator.
func TBody(opts ...Option) Component {
	n := newComponent("divide-y divide-gray-200", opts...)
	n.content = h.TBody(n.manuallySetAttributes())
	return n
}

// Tr returns a HTML <tr> element with simple hover style.
func Tr(opts ...Option) Component {
	n := newComponent("hover:bg-gray-100", opts...)
	n.content = h.Tr(n.manuallySetAttributes())
	return n
}

// Th returns a HTML <th> element with bold font and default spacing.
func Th(opts ...Option) Component {
	n := newComponent("px-3 py-2 whitespace-nowrap font-medium text-gray-900", opts...)
	n.content = h.Th(n.manuallySetAttributes())
	return n
}

// Td returns a HTML <td> element with defaut spacing.
func Td(opts ...Option) Component {
	n := newComponent("px-3 py-2 whitespace-nowrap text-gray-900", opts...)
	n.content = h.Td(n.manuallySetAttributes())
	return n
}
