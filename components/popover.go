package components

import sc "github.com/geneton-ltd/syder_components"

// Popover returns a HTML5 <div> element that floats natively using the Popover API.
// For attachment to happen, pass [ID] option and then [PopoverTrigger] with
// the same id to the trigger element (e.g. a button).
//
// Example:
//
//	components.Popover(sc.ID("settingsMenu"), ...)
//	components.Button(sc.PopoverTrigger("settingsMenu"), sc.Text("Settings"))
func Popover(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Attr("popover", "auto"), sc.Class("absolute shadow-xl bg-white rounded-md border p-4")}, children...)
	return sc.Div(children...)
}
