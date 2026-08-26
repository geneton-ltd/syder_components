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
	c := newComponent("", opts...)
	c.content = g.El(tag, c.manuallySetAttributes())
	return c
}
