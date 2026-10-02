package sydercomponents

// If returns the given [Component], if condition is true. Otherwise, returns [Empty].
// This helper function is good for inlining elements conditionally.
// If it's important that the given [Component] is only evaluated if condition is true
// (for example, when using nilable variables), use [Iff] instead.
func If(condition bool, comp Component) Component {
	if condition {
		return comp
	}
	return Empty()
}

// Iff calls the given function and returns its result, if condition is true. Otherwise, returns [Empty].
// This helper function is good for inlining elements conditionally when the node depends on nilable data,
// or some other code that could potentially panic.
// If you just need simple conditional rendering, see [If].
func Iff(condition bool, f func() Component) Component {
	if condition {
		return f()
	}
	return Empty()
}
