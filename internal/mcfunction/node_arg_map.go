package mcfunction

type nodeArgMap struct {
	*nodeArg
	mapSpec *MapSpec
}

func (n *nodeArgMap) addChild(child Node) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *nodeArgMap) setParent(parent Node) {
	n.parent = parent
}

func (n *nodeArgMap) MapSpec() *MapSpec {
	return n.mapSpec
}
