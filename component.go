package sydercomponents

import (
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type component struct {
	children []g.Node

	classOverride string
	classAppend   string

	// content after all elements are defined
	content g.Node
	class   string
	id      string
}

type opt func(c *component)

func WithChildren(value ...g.Node) opt {
	return func(c *component) {
		c.children = value
	}
}

func WithClassAppend(value string) opt {
	return func(c *component) {
		c.class += value
	}
}

func WithClass(value string) opt {
	return func(c *component) {
		c.class = value
	}
}

func WithID(value string) opt {
	return func(c *component) {
		c.id = value
	}
}

func (c component) Render(w io.Writer) error {
	return c.content.Render(w)
}

func newComponent(cls string, opts ...opt) component {
	c := component{class: cls}
	for _, o := range opts {
		o(&c)
	}

	return c
}

// Class contains the applied class to the element
func (c component) Class() string {
	return c.class
}

// ID contains th element id, but only if WithID was called
func (c component) ID() string {
	return c.id
}

// manuallySetAttributes is supposed to be called for every
// new component. It injects all manually set attributes, such as
// class (via the WithClass hook), ID (via the WithID hook)
// children (via the WithChildren hook) and others
func (c component) manuallySetAttributes() g.Group {
	return []g.Node{
		h.Class(c.class),
		g.If(c.id != "", h.ID(c.id)),
		g.Group(c.children),
	}
}
