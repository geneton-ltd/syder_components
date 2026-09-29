package sydercomponents

import (
	g "maragu.dev/gomponents"
)

// Any renders given HTML5 tag.
// It does not apply any styles.
//
// Example:
//
//	sc.Any(
//	    "div",
//	    sc.ClassAppend("flex flex-col gap-2")
//	)
func Any(tag string, opts ...Option) Component {
	n := newComponent(tag, "", opts...)
	n.content = g.El(tag, n.manuallySetAttributes())
	return n
}

// Empty returns a Component that renders absolutely nothing.
// It is useful for conditional rendering.
func Empty() Component {
	return Component{}
}

// Text creates a text DOM [Node] that renders the escaped string t.
func Text(t string) Component {
	return Component{
		content: g.Text(t),
		text:    t,
		tag:     "",
	}
}

// Raw creates a text DOM [Node] that renders the unescaped string t.
func Raw(t string) Component {
	return Component{
		content: g.Raw(t),
		text:    t,
		tag:     "",
	}
}

// Wrap allows injecting raw gomponents.Node into the tree.
// They cannot be inspected using the [GetClass] or [GetChildren].
func Wrap(node g.Node) Component {
	return Component{
		content: node,
	}
}
