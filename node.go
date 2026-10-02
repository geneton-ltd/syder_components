package sydercomponents

import g "maragu.dev/gomponents"

// Node allows injecting gomponents.Node into the tree.
// It cannot be inspected using the [Component.GetClass] or [Component.GetChildren].
// It is meant for foreign nodes that are not present in this library,
// not for gomponents' own html package helpers.
func Node(node g.Node) Component {
	return Component{
		content: node,
		kind:    KindNode,
	}
}
