package sydercomponents

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

// Button returns an HTML5 <button> element with default style without colors.
func Button(opts ...Option) Component {
	n := newComponent("rounded rounded-sm shadow-xs p-1 hover:cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed", opts...)
	n.content = h.Button(n.manuallySetAttributes())
	return n
}

// IconButton returns an HTML5 <button> element containing an icon wrapped
// in a <span> element.
//
// Example:
//
//	sc.IconButton(
//	    solid.ArrowLeft(),
//	    sc.Children(sc.Text("Previous")),
//	)
func IconButton(icon g.Node, opts ...Option) Component {
	opts = append(
		[]Option{
			ClassAppend("whitespace-nowrap flex gap-1 items-center"),
			Children(Any("span", Class("h-4 w-4 shrink-0"), NodeAppend(icon))),
		},
		opts...,
	)
	return Button(opts...)
}

// IconOnlyButton returns a [IconButton] with special styles to make sure the
// icon placement is accurate.
//
// Example:
//
//	sc.IconOnlyButton(
//	    solid.Trash(),
//	)
func IconOnlyButton(icon g.Node, opts ...Option) Component {
	opts = append([]Option{ClassAppend("w-fit self-center"), Type("button")}, opts...)
	return IconButton(icon, opts...)
}
