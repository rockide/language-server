package mcfunction

func WalkNodeTree(root INode, fn func(INode) bool) {
	if !fn(root) {
		return
	}
	for _, child := range root.Children() {
		WalkNodeTree(child, fn)
	}
}

func IsInside(n INode, pos uint32) bool {
	start, end := n.Range()
	return pos >= start && pos <= end
}

func NodeAt(root INode, pos uint32) INode {
	var result INode
	WalkNodeTree(root, func(n INode) bool {
		matched := IsInside(n, pos)
		if matched {
			result = n
		}
		return true
	})
	return result
}
