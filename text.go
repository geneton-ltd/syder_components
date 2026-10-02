package sydercomponents

import (
	g "maragu.dev/gomponents"
)

// Text creates a text Component that renders the escaped string t.
// For unescaped text, see [Raw].
func Text(t string) Component {
	return Component{
		content: g.Text(t),
		text:    t,
		kind:    KindText,
	}
}

// Raw creates a text Component that renders the unescaped string t.
func Raw(t string) Component {
	return Component{
		content: g.Raw(t),
		text:    t,
		kind:    KindText,
	}
}
