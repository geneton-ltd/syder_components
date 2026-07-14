package sydercomponents

import (
	g "maragu.dev/gomponents"
)

func Any(tag string, opts ...Option) Component {
	c := newComponent("", opts...)
	c.content = g.El(tag, c.manuallySetAttributes())
	return c
}
