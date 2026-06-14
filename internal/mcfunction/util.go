package mcfunction

func WalkNodeTree(root Node, fn func(Node) bool) {
	if !fn(root) {
		return
	}
	for _, child := range root.Children() {
		WalkNodeTree(child, fn)
	}
}

func IsInside(n Node, pos uint32) bool {
	start, end := n.Range()
	return pos >= start && pos <= end
}

func NodeAt(root Node, pos uint32) Node {
	var result Node
	WalkNodeTree(root, func(n Node) bool {
		matched := IsInside(n, pos)
		if matched {
			result = n
		}
		return true
	})
	return result
}
