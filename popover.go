package sydercomponents

import h "maragu.dev/gomponents/html"

// Popover returns a HTML5 <div> element that floats natively using the Popover API.
// For attachment to happen, pass ID option and then PopoverTrigger with
// the same id to the trigger element (e.g. a button).
//
// Example:
//
//	sc.Popover(sc.ID("settingsMenu"), ...)
//	sc.Button(sc.PopoverTrigger("settingsMenu"), sc.Nodes(g.Text("Settings")))
func Popover(opts ...Option) Component {
	opts = append([]Option{Attribute("popover", "auto")}, opts...)
	n := newComponent("absolute shadow-xl bg-white rounded-md border p-4", opts...)
	n.content = h.Div(n.manuallySetAttributes())
	return n
}
