package sydercomponents

import "io"

// Kind describes what kind of [Component] it is.
// This is useful for testing to determine the [Component] kind.
// It defaults to being [KindElement].
type Kind string

const (
	// KindText is used by [Text] and [Raw] text components.
	// Its content is readable with [Component.GetText].
	KindText Kind = "text"

	// KindAttr is used by .
	// No getters work for this kind (except [Component.GetKind]).
	KindAttr Kind = "attr"

	// KindElement is used by all HTML components.
	// Its attributes can be accessed with Get* methods.
	KindElement Kind = "element"

	// KindGroup
	KindGroup Kind = "group"
)

// Node interface defines
type Node interface {
	Kind() Kind
	Name() string
	Value() string
	Children() []Node
	Attr(name string) (string, bool)
	Render(w io.Writer) error
}
