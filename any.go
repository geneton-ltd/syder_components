package sydercomponents

import (
	g "maragu.dev/gomponents"
)

// Any renders given HTML5 tag.
// It does not apply any styles.
//
// Example:
//
//	sc.Any(
//	    "div",
//	    sc.WithClassAppend("flex flex-col gap-2")
//	)
func Any(tag string, opts ...Option) Component {
	n := newComponent("", opts...)
	n.content = g.El(tag, n.manuallySetAttributes())
	return n
}
