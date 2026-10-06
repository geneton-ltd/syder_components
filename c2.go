package sydercomponents

import (
	"fmt"
	"io"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

type CK2 string

const (
	CK2Element CK2 = "element"
	CK2Tag     CK2 = "tag"
	CK2Text    CK2 = "text"
)

type C2 struct {
	kind  CK2
	name  string
	value string
	nodes []C2Node
}

func (c2 C2) Kind() CK2       { return c2.kind }
func (c2 C2) Name() string    { return c2.name }
func (c2 C2) Value() string   { return c2.value }
func (c2 C2) Nodes() []C2Node { return c2.nodes }

type C2Node interface {
	Render(w io.Writer) error
	Kind() CK2
	Name() string
	Value() string
	Nodes() []C2Node
}

func (c2 C2) Render(w io.Writer) error {
	switch c2.kind {
	case CK2Tag:
		if _, err := fmt.Fprintf(w, " %s=\"%s\" ", c2.name, c2.value); err != nil {
			return err
		}
	case CK2Text:
		if _, err := fmt.Fprint(w, c2.value); err != nil {
			return err
		}
	case CK2Element:
		if _, err := fmt.Fprintf(w, "<%s", c2.name); err != nil {
			return err
		}
		nonTagsIndices := []int{}
		for i, n := range c2.nodes {
			if n.Kind() != CK2Tag {
				nonTagsIndices = append(nonTagsIndices, i)
			} else {
				if err := n.Render(w); err != nil {
					return err
				}
			}
		}

		if _, err := fmt.Fprint(w, ">"); err != nil {
			return err
		}

		for _, idx := range nonTagsIndices {
			if err := c2.nodes[idx].Render(w); err != nil {
				return err
			}
		}

		if _, err := fmt.Fprintf(w, "</%s>", c2.name); err != nil {
			return err
		}
	}

	return nil
}

func newC2Element(name string, children ...C2Node) C2Node {
	return C2{kind: CK2Element, name: name, nodes: children}
}

func C2Div(children ...C2Node) C2Node {
	return newC2Element("div", children...)
}

func C2Text(v string) C2Node {
	return C2{
		kind:  CK2Text,
		value: v,
	}
}

func C2H1(children ...C2Node) C2Node {
	return newC2Element("h1", children...)
}

func C2Class(v string) C2Node {
	return C2{
		kind:  CK2Tag,
		name:  "class",
		value: v,
	}
}

func TestC2() C2Node {
	return C2Div(
		C2Class("h-full w-full"),
		C2H1(
			C2Text("this is heading"),
		),
	)
}

func TestC2InGNode() g.Node {
	return h.Div(
		h.Style("border:1px solid black; padding:5px"),
		h.Div(
			TestC2(),
		),
	)
}
