package sydercomponents

import "io"

// Kind describes what kind of [Node] it is.
// This is useful for testing to determine the node's type.
type Kind string

const (
	// KindText is used by [Text] and [Raw] text components.
	// Its content is readable with the [Node.Value] method.
	KindText Kind = "text"

	// KindAttr is used by attribute nodes.
	// Its name and value can be accessed via [Node.Name] and [Node.Value] or [Node.Attr].
	KindAttr Kind = "attr"

	// KindElement is used by all HTML components.
	// Its children and attributes can be accessed with [Node.Attr] and [Node.Children].
	KindElement Kind = "element"

	// KindGroup is used by [Group] to bundle multiple nodes together.
	KindGroup Kind = "group"

	// KindGNode is used by [GNode] to wrap [maragu.dev/gomponents.Node] instances.
	KindGNode Kind = "gnode"
)

// Node represents a single node in the library's rendering tree.
// It provides methods for inspection in unit tests and HTML rendering.
type Node interface {
	Kind() Kind
	Name() string
	Value() string
	Children() []Node
	Attr(name string) (string, bool)
	Render(w io.Writer) error
}
