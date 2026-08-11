package sydercomponents

import h "maragu.dev/gomponents/html"

// Dialog returns a html dialog element. For attachment to happen, pass
// WithID option and then WithDialogTrigger with the same id to the trigger
// element (e.g. a button). Attachment is performed automatically.
func Dialog(opts ...Option) Component {
	opts = append([]Option{WithAttribute("closedBy", "any")}, opts...)
	dialog := newComponent("shadow-xl bg-white backdrop:bg-black/40 rounded-md", opts...)
	dialog.content = h.Dialog(dialog.manuallySetAttributes())
	return dialog
}

// DialogRelative returns dialog that is relative to its parent. The parent
// must have the "relative" class
func DialogRelative(opts ...Option) Component {
	opts = append([]Option{WithClassAppend("absolute m-0")}, opts...)
	return Dialog(opts...)
}

// DialogFixedCentered returns dialog fixed at the center of the screen (horizontally and vertically)
func DialogFixedCentered(opts ...Option) Component {
	opts = append([]Option{WithClassAppend("fixed inset-0 m-auto")}, opts...)
	return Dialog(opts...)
}
