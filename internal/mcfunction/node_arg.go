package mcfunction

type nodeArg struct {
	*node
	paramKind ParameterKind
}

func (n *nodeArg) addChild(child Node) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *nodeArg) setParent(parent Node) {
	n.parent = parent
}

func (n *nodeArg) ParamKind() ParameterKind {
	return n.paramKind
}

func (n *nodeArg) ParamSpec() (ParameterSpec, bool) {
	commandNode, ok := n.parent.(CommandNode)
	if !ok {
		return ParameterSpec{}, false
	}
	argIndex := n.index
	return commandNode.ParamSpecAt(argIndex)
}

func (n *nodeArg) CommandNode() CommandNode {
	commandNode, ok := n.parent.(CommandNode)
	if !ok {
		return nil
	}
	return commandNode
}
