package components

import sc "github.com/geneton-ltd/syder_components"

// Dialog returns a HTML5 <dialog> element. For attachment
// to happen, pass [sc.ID] and then [sc.DialogTrigger] with
// the same id to the trigger element (e.g. a button).
//
// Example:
//
//	components.Dialog(sc.ID("id"), sc.Data("name", "john"))
//	components.Button(sc.DialogTrigger("id"))
func Dialog(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("shadow-xl bg-white backdrop:bg-black/40 rounded-md")}, children...)
	return sc.Dialog(children...)
}

// DialogRelative returns [Dialog] that is relative to its parent. The parent
// must have the "relative" class.
func DialogRelative(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("absolute m-0")}, children...)
	return Dialog(children...)
}

// DialogFixedCentered returns [Dialog] fixed at the center of the screen (horizontally and vertically).
func DialogFixedCentered(children ...sc.Node) sc.Node {
	children = append([]sc.Node{sc.Class("fixed inset-0 m-auto")}, children...)
	return Dialog(children...)
}
