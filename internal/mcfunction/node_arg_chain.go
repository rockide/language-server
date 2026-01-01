package mcfunction

type NodeArgCommand struct {
	*NodeCommand
	paramKind ParameterKind
}

func (n *NodeArgCommand) addChild(child INode) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *NodeArgCommand) setParent(parent INode) {
	n.parent = parent
}

func (n *NodeArgCommand) setIndex(index int) {
	n.index = index
}

func (n *NodeArgCommand) Kind() NodeKind {
	return NodeKindCommandArg
}

func (n *NodeArgCommand) Range() (start, end uint32) {
	return n.start, n.end
}

func (n *NodeArgCommand) Text(src []rune) string {
	return string(src[n.start:n.end])
}

func (n *NodeArgCommand) PrevSibling() INode {
	if n.parent == nil || n.index == 0 {
		return nil
	}
	i := n.index - 1
	if i < 0 || i >= len(n.parent.Children()) {
		return nil
	}
	return n.parent.Children()[i]
}

func (n *NodeArgCommand) NextSibling() INode {
	if n.parent == nil {
		return nil
	}
	i := n.index + 1
	if i < 0 || i >= len(n.parent.Children()) {
		return nil
	}
	return n.parent.Children()[i]
}

func (n *NodeArgCommand) Parent() INode {
	return n.parent
}

func (n *NodeArgCommand) Index() int {
	return n.index
}

func (n *NodeArgCommand) Children() []INode {
	return n.children
}

func (n *NodeArgCommand) ParamKind() ParameterKind {
	return n.paramKind
}

func (n *NodeArgCommand) ParamSpec() (ParameterSpec, bool) {
	commandNode, ok := n.parent.(INodeCommand)
	if !ok {
		return ParameterSpec{}, false
	}
	argIndex := n.index
	return commandNode.ParamSpecAt(argIndex)
}

func (n *NodeArgCommand) IsInside(pos uint32) bool {
	return n.start <= pos && pos < n.end
}

func (n *NodeArgCommand) CommandName() string {
	return n.name
}

func (n *NodeArgCommand) Args() []INode {
	return n.children
}

func (n *NodeArgCommand) Spec() *Spec {
	return n.spec
}

func (n *NodeArgCommand) OverloadStates() []*overloadState {
	return n.overloadStates
}

func (n *NodeArgCommand) ParamSpecAt(index int) (ParameterSpec, bool) {
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

func (n *NodeArgCommand) IsValid() bool {
	for _, o := range n.overloadStates {
		if o.matched {
			return true
		}
	}
	return false
}
