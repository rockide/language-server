package mcfunction

type NodeCommand struct {
	*Node
	name           string
	spec           *Spec
	overloadStates []*overloadState
}

func (n *NodeCommand) addChild(child INode) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *NodeCommand) setParent(parent INode) {
	n.parent = parent
}

func (n *NodeCommand) setIndex(index int) {
	n.index = index
}

func (n *NodeCommand) CommandName() string {
	return n.name
}

func (n *NodeCommand) Args() []INode {
	if len(n.children) == 0 {
		return nil
	}
	first := n.children[0]
	if first.Kind() != NodeKindCommandLit {
		return n.children
	}
	return n.children[1:]
}

func (n *NodeCommand) Spec() *Spec {
	return n.spec
}

func (n *NodeCommand) OverloadStates() []*overloadState {
	return n.overloadStates
}

func (n *NodeCommand) ParamSpecAt(index int) (ParameterSpec, bool) {
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

func (n *NodeCommand) IsValid() bool {
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
