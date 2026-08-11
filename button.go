package sydercomponents

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Button(opts ...Option) Component {
	n := newComponent("rounded rounded-sm shadow-xs p-1 hover:cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed", opts...)
	n.content = h.Button(n.manuallySetAttributes())
	return n
}

func IconButton(icon g.Node, opts ...Option) Component {
	options := []Option{WithClassAppend("whitespace-nowrap flex gap-1 items-center")}
	options = append(options, WithChildren(Any("span", WithClass("h-4 w-4 shrink-0"), WithNodeAppend(icon))))
	options = append(options, opts...)
	return Button(options...)
}

func IconOnlyButton(icon g.Node, opts ...Option) Component {
	options := append([]Option{WithClassAppend("w-fit self-center"), WithType("button")}, opts...)
	return IconButton(icon, options...)
}
