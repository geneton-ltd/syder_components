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
	opts = append(
		[]Option{
			WithClassAppend("whitespace-nowrap flex gap-1 items-center"),
			WithChildren(Any("span", WithClass("h-4 w-4 shrink-0"), WithNodeAppend(icon))),
		},
		opts...,
	)
	return Button(opts...)
}

func IconOnlyButton(icon g.Node, opts ...Option) Component {
	opts = append([]Option{WithClassAppend("w-fit self-center"), WithType("button")}, opts...)
	return IconButton(icon, opts...)
}
