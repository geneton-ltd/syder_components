package sydercomponents

import g "maragu.dev/gomponents"

// Text returns a raw text node.
// It is full testable via the GetText() method.
func Text(val string) Component {
	return Component{
		content: g.Text(val),
		text:    val,
	}
}
