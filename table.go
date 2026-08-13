package sydercomponents

import (
	h "maragu.dev/gomponents/html"
)

func TableContainer(opts ...Option) Component {
	n := newComponent("overflow-x-auto rounded border border-gray-300 shadow-sm", opts...)
	n.content = h.Div(n.manuallySetAttributes())
	return n
}

func Table(opts ...Option) Component {
	n := newComponent("min-w-full divide-y-2 divide-gray-200", opts...)
	n.content = h.Table(n.manuallySetAttributes())
	return n
}

func THead(opts ...Option) Component {
	n := newComponent("", opts...)
	n.content = h.THead(n.manuallySetAttributes())
	return n
}

func TBody(opts ...Option) Component {
	n := newComponent("divide-y divide-gray-200", opts...)
	n.content = h.TBody(n.manuallySetAttributes())
	return n
}

func Tr(opts ...Option) Component {
	n := newComponent("hover:bg-gray-100", opts...)
	n.content = h.Tr(n.manuallySetAttributes())
	return n
}

func Th(opts ...Option) Component {
	n := newComponent("px-3 py-2 whitespace-nowrap font-medium text-gray-900", opts...)
	n.content = h.Th(n.manuallySetAttributes())
	return n
}

func Td(opts ...Option) Component {
	n := newComponent("px-3 py-2 whitespace-nowrap text-gray-900", opts...)
	n.content = h.Td(n.manuallySetAttributes())
	return n
}
