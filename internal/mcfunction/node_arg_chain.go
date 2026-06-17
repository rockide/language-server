package mcfunction

type nodeArgCommand struct {
	*nodeCommand
	paramKind ParameterKind
}

func (n *nodeArgCommand) addChild(child Node) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *nodeArgCommand) setParent(parent Node) {
	n.parent = parent
}

func (n *nodeArgCommand) ParamKind() ParameterKind {
	return n.paramKind
}

func (n *nodeArgCommand) ParamSpec() (ParameterSpec, bool) {
	commandNode, ok := n.parent.(CommandNode)
	if !ok {
		return ParameterSpec{}, false
	}
	argIndex := n.index
	return commandNode.ParamSpecAt(argIndex)
}

func (n *nodeArgCommand) CommandName() string {
	return n.name
}

func (n *nodeArgCommand) Args() []Node {
	return n.children
}

func (n *nodeArgCommand) Spec() *Spec {
	return n.spec
}

func (n *nodeArgCommand) OverloadStates() []*overloadState {
	return n.overloadStates
}

func (n *nodeArgCommand) ParamSpecAt(index int) (ParameterSpec, bool) {
	for _, o := range n.overloadStates {
		if !o.matched {
			continue
		}
		if index >= 0 && index < len(o.ov.Parameters) {
			return o.ov.Parameters[index], true
		}
	}
	return ParameterSpec{}, false
}

func (n *nodeArgCommand) IsValid() bool {
	for _, o := range n.overloadStates {
		if o.matched {
			return true
		}
	}
	return false
}

func (n *nodeArgCommand) CommandNode() CommandNode {
	return n
}
