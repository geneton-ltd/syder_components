package sydercomponents

// If returns the given [Node] if condition is true. Otherwise, returns nil.
// This helper function is good for inlining elements conditionally.
// If it's important that the given [Node] is only evaluated if condition is true
// (for example, when using nilable variables), use [Iff] instead.
func If(condition bool, n Node) Node {
	if condition {
		return n
	}
	return nil
}

// Iff calls the given function and returns its result if condition is true. Otherwise, returns nil.
// This helper function is good for inlining elements conditionally when the node depends on nilable data,
// or some other code that could potentially panic.
// If you just need simple conditional rendering, see [If].
func Iff(condition bool, f func() Node) Node {
	if condition {
		return f()
	}
	return nil
}
