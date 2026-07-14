package sydercomponents

import (
	h "maragu.dev/gomponents/html"
)

func Navbar(opts ...opt) component {
	n := newComponent("h-[50px] bg-gray-800 flex justify-between text-white ", opts...)
	n.content = h.Div(n.manuallySetAttributes())
	return n
}

func NavLink(opts ...opt) component {
	n := newComponent("underline hover:bg-gray-600 flex items-center px-2", opts...)
	n.content = h.A(n.manuallySetAttributes())

	return n
}
