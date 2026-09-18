package sydercomponents

import h "maragu.dev/gomponents/html"

// Popover returns a HTML5 <div> element that floats natively using the Popover API.
// For attachment to happen, pass WithID option and then WithPopoverTrigger with
// the same id to the trigger element (e.g. a button).
//
// Example:
//
//	sc.Popover(sc.WithID("settingsMenu"), ...)
//	sc.Button(sc.WithPopoverTrigger("settingsMenu"), sc.WithNodes(g.Text("Settings")))
func Popover(opts ...Option) Component {
	opts = append([]Option{WithAttribute("popover", "auto")}, opts...)
	n := newComponent("absolute shadow-xl bg-white rounded-md border p-4", opts...)
	n.content = h.Div(n.manuallySetAttributes())
	return n
}
