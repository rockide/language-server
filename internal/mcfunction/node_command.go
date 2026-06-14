package mcfunction

type nodeCommand struct {
	*node
	name           string
	spec           *Spec
	overloadStates []*overloadState
}

func (n *nodeCommand) addChild(child Node) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *nodeCommand) setParent(parent Node) {
	n.parent = parent
}

func (n *nodeCommand) CommandName() string {
	return n.name
}

func (n *nodeCommand) Args() []Node {
	if len(n.children) == 0 {
		return nil
	}
	first := n.children[0]
	if first.Kind() != NodeKindCommandLit {
		return n.children
	}
	return n.children[1:]
}

func (n *nodeCommand) Spec() *Spec {
	return n.spec
}

func (n *nodeCommand) OverloadStates() []*overloadState {
	return n.overloadStates
}

func (n *nodeCommand) ParamSpecAt(index int) (ParameterSpec, bool) {
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

func (n *nodeCommand) IsValid() bool {
	if n.Kind() == NodeKindInvalidCommand {
		return false
	}
	for _, o := range n.overloadStates {
		if o.matched {
			return true
		}
	}
	return false
}
