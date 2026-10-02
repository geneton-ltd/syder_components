package sydercomponents

import h "maragu.dev/gomponents/html"

// Dialog returns a HTML5 <dialog> element. For attachment to happen, pass
// [ID] option and then [DialogTrigger] with the same id to the trigger
// element (e.g. a button). Attachment is performed automatically.
//
// Example:
//
//	sc.Dialog(sc.ID("id"), sc.Data("name", "john"))
//	sc.Button(sc.DialogTrigger("id"))
func Dialog(opts ...Option) Component {
	n := newComponent("shadow-xl bg-white backdrop:bg-black/40 rounded-md", opts...)
	n.content = h.Dialog(n.manuallySetAttributes())
	return n
}

// DialogRelative returns [Dialog] that is relative to its parent. The parent
// must have the "relative" class.
func DialogRelative(opts ...Option) Component {
	opts = append([]Option{ClassAppend("absolute m-0")}, opts...)
	return Dialog(opts...)
}

// DialogFixedCentered returns [Dialog] fixed at the center of the screen (horizontally and vertically).
func DialogFixedCentered(opts ...Option) Component {
	opts = append([]Option{ClassAppend("fixed inset-0 m-auto")}, opts...)
	return Dialog(opts...)
}
