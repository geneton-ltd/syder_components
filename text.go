package sydercomponents

import (
	"fmt"
	"html/template"
	"io"
)

type text struct {
	value string
}

func (t text) Kind() Kind                      { return KindText }
func (t text) Name() string                    { return "" }
func (t text) Value() string                   { return t.value }
func (t text) Children() []Node                { return []Node{} }
func (t text) Attr(name string) (string, bool) { return "", false }
func (t text) Render(w io.Writer) error {
	_, err := fmt.Fprint(w, t.value)
	return err
}

// Text creates a text [Node] that renders the escaped string t.
// For unescaped text, see [Raw].
func Text(t string) Node {
	return text{value: template.HTMLEscapeString(t)}
}

// Raw creates a text [Node] that renders the unescaped string t.
func Raw(t string) Node {
	return text{value: t}
}
