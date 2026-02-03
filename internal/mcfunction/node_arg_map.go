package mcfunction

type nodeArgMap struct {
	*NodeArg
	mapSpec *MapSpec
}

func (n *nodeArgMap) addChild(child INode) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *nodeArgMap) setParent(parent INode) {
	n.parent = parent
}

func (n *nodeArgMap) setIndex(index int) {
	n.index = index
}

func (n *nodeArgMap) MapSpec() *MapSpec {
	return n.mapSpec
}
