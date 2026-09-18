package sydercomponents

import h "maragu.dev/gomponents/html"

// Dialog returns a HTML5 <dialog> element. For attachment to happen, pass
// WithID option and then WithDialogTrigger with the same id to the trigger
// element (e.g. a button). Attachment is performed automatically.
//
// Example:
//
//	sc.Dialog(sc.WithID("id"), sc.WithData("name", "john"))
//	sc.Button(sc.WithDialogTrigger("id"))
func Dialog(opts ...Option) Component {
	n := newComponent("shadow-xl bg-white backdrop:bg-black/40 rounded-md", opts...)
	n.content = h.Dialog(n.manuallySetAttributes())
	return n
}

// DialogRelative returns [Dialog] that is relative to its parent. The parent
// must have the "relative" class.
func DialogRelative(opts ...Option) Component {
	opts = append([]Option{WithClassAppend("absolute m-0")}, opts...)
	return Dialog(opts...)
}

// DialogFixedCentered returns [Dialog] fixed at the center of the screen (horizontally and vertically).
func DialogFixedCentered(opts ...Option) Component {
	opts = append([]Option{WithClassAppend("fixed inset-0 m-auto")}, opts...)
	return Dialog(opts...)
}
