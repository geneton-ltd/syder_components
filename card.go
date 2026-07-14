package sydercomponents

import (
	h "maragu.dev/gomponents/html"
)

func Card(opts ...Option) Component {
	n := newComponent("w-full rounded-md shadow-xs p-2 border border-1 border-gray-600", opts...)
	n.content = h.Div(n.manuallySetAttributes())
	return n
}

func SubCard(opts ...Option) Component {
	options := append(opts, WithClassAppend("border-dashed hover:bg-white"))
	return Card(options...)
}
