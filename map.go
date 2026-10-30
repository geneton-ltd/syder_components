package sydercomponents

// Map a slice of anything to a [Group] (which is just a slice of [Node]-s).
func Map[T any](ts []T, cb func(T) Node) Node {
	nodes := make([]Node, 0, len(ts))
	for _, t := range ts {
		nodes = append(nodes, cb(t))
	}
	return Group(nodes...)
}
