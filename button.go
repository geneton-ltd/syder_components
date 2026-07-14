package sydercomponents

import h "maragu.dev/gomponents/html"

func Button(opts ...Option) Component {
	n := newComponent("rounded rounded-sm shadow-xs p-1 hover:cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed", opts...)
	n.content = h.Button(n.manuallySetAttributes())
	return n
}

func IconButton(icon Component, opts ...Option) Component {
	options := append(opts, WithClassAppend("whitespace-nowrap flex gap-1 items-center"))
	options = append(options, WithChildren(Any("span", WithClass("h-4 w-4 shrink-0"), WithChildren(icon))))
	return Button(options...)
}

func IconOnlyButton(icon Component, opts ...Option) Component {
	options := append(opts, WithClassAppend("w-fit self-center"), WithType("button"))
	return IconButton(icon, options...)
}
