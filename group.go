package sydercomponents

import (
	"io"
)

// group is an internal [Node] of [KindGroup] kind.
// Its content is accessible with the [group.Children] method.
// Each child is rendered using its own [Node.Render] method.
// It has no name, value or attributes.
type group struct {
	children []Node
}

func (g group) Kind() Kind                      { return KindGroup }
func (g group) Name() string                    { return "" }
func (g group) Value() string                   { return "" }
func (g group) Children() []Node                { return g.children }
func (g group) Attr(name string) (string, bool) { return "", false }

func (g group) Render(w io.Writer) error {
	for _, child := range g.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	return nil
}

// Group groups multiple [Node] instances into a single [Node].
// When passed to an element constructor, its children are unpacked automatically.
func Group(nodes ...Node) Node {
	return group{children: nodes}
}
