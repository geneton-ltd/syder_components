package sydercomponents

import (
	"io"
)

type group struct {
	children []Node
}

func (g group) Kind() Kind                  { return KindGroup }
func (g group) Name() string                { return "" }
func (g group) Value() string               { return "" }
func (g group) Children() []Node            { return g.children }
func (g group) Attributes() map[string]Node { return map[string]Node{} }
func (g group) Render(w io.Writer) error {
	for _, child := range g.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}
	return nil
}

func Group(nodes ...Node) Node {
	return group{children: nodes}
}
