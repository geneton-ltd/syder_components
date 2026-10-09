package sydercomponents

import (
	"fmt"
	"io"
	"maps"
	"slices"
	"strings"
)

// element is an internal [Node] of [KindElement] kind representing an HTML element.
// Its tag name can be accessed with the [element.Name] method.
//
// Creating a new element with [newElement] recursively unpacks [KindGroup] nodes,
// saves [KindAttr] nodes into an internal attrs map, and saves all other nodes as children.
//
// Attributes are rendered in alphabetical order by their name.
// Void (self-closing) elements are rendered without children or a closing tag.
type element struct {
	name     string
	value    string
	children []Node
	attrs    map[string]Node
}

// voidElements returns list of elements that should not render any child content.
// It is a function to ensure immutability.
// Source: https://developer.mozilla.org/en-US/docs/Glossary/Void_element
func voidElements() []string {
	return []string{"area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr"}
}

func (e element) Kind() Kind       { return KindElement }
func (e element) Name() string     { return e.name }
func (e element) Value() string    { return e.value }
func (e element) Children() []Node { return e.children }

// Attr returns the value of an attribute by name and true if it exists.
// Otherwise, it returns an empty string and false.
func (e element) Attr(name string) (string, bool) {
	if n, ok := e.attrs[name]; ok {
		return n.Value(), true
	}

	return "", false
}

func (e element) Render(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "<%s", e.name); err != nil {
		return err
	}

	attrNames := slices.Sorted(maps.Keys(e.attrs))
	for _, attrName := range attrNames {
		if _, err := fmt.Fprint(w, " "); err != nil {
			return err
		}

		if err := e.attrs[attrName].Render(w); err != nil {
			return err
		}
	}

	if slices.Contains(voidElements(), e.name) {
		if _, err := fmt.Fprintf(w, "/>"); err != nil {
			return err
		}

		return nil
	}

	if _, err := fmt.Fprintf(w, ">"); err != nil {
		return err
	}

	for _, child := range e.children {
		if err := child.Render(w); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(w, "</%s>", e.name); err != nil {
		return err
	}

	return nil
}

func (e *element) unpackChild(child Node) {
	switch child.Kind() {
	case KindAttr:
		name := child.Name()
		if name != "class" {
			e.attrs[name] = child
		} else {
			var cls string
			if n, ok := e.attrs["class"]; ok {
				cls = n.Value() + " "
			}
			e.attrs["class"] = Attr(name, cls+strings.TrimSpace(child.Value()))
		}
	case KindGroup:
		for _, groupMember := range child.Children() {
			e.unpackChild(groupMember)
		}
	default:
		e.children = append(e.children, child)
	}
}

// newElement creates a new HTML [element] node, automatically categorizing
// passed items into attributes and children while unpacking any groups recursively.
func newElement(name string, children ...Node) Node {
	e := &element{
		name:     name,
		attrs:    map[string]Node{},
		children: []Node{},
	}

	for _, child := range children {
		e.unpackChild(child)
	}

	return *e
}
