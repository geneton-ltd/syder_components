package sydercomponents

import (
	"io"
	"maragu.dev/gomponents"
)

// gnode is an internal [Node] of [KindGNode] kind representing a [g.Node].
// Its content is not accessible via getters. It has no name, value, attributes or children.
type gnode struct {
	content gomponents.Node
}

func (g gnode) Kind() Kind                      { return KindGNode }
func (g gnode) Name() string                    { return "" }
func (g gnode) Value() string                   { return "" }
func (g gnode) Children() []Node                { return []Node{} }
func (g gnode) Attr(name string) (string, bool) { return "", false }
func (g gnode) Render(w io.Writer) error {
	if g.content == nil {
		return nil
	}
	return g.content.Render(w)
}

// GNode wraps a [gomponents.Node], so that it can be passed into our component tree as a standard [Node].
func GNode(node gomponents.Node) Node {
	return gnode{content: node}
}
