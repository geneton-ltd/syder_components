package sydercomponents

// Any renders given HTML5 tag.
// It does not apply any styles.
//
// Example:
//
//	sc.Any(
//		"div",
//		sc.Class("flex flex-col gap-2")
//	)
func Any(tag string, children ...Node) Node {
	return newElement(tag, children...)
}
