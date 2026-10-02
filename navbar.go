package sydercomponents

import (
	h "maragu.dev/gomponents/html"
)

// Navbar returns a HTML5 <div> element with simple styles
// acting as a container for navigation elements.
func Navbar(opts ...Option) Component {
	n := newComponent("h-[50px] bg-gray-800 flex justify-between text-white", opts...)
	n.content = h.Div(n.manuallySetAttributes())
	return n
}

// NavLink returns a HTML5 <a> element, which is supposed to be put into [Navbar].
func NavLink(opts ...Option) Component {
	n := newComponent("hover:bg-gray-600 flex items-center px-2", opts...)
	n.content = h.A(n.manuallySetAttributes())
	return n
}
