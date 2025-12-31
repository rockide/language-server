package mcfunction

type INodeArg interface {
	INode
	ParamKind() ParameterKind
	ArgParamSpec() (ParameterSpec, bool)
}

type NodeArg struct {
	*Node
	paramKind ParameterKind
}

func (n *NodeArg) addChild(child INode) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *NodeArg) setParent(parent INode) {
	n.parent = parent
}

func (n *NodeArg) setIndex(index int) {
	n.index = index
}

func (n *NodeArg) Kind() NodeKind {
	return NodeKindCommandArg
}

func (n *NodeArg) Range() (start, end uint32) {
	return n.start, n.end
}

func (n *NodeArg) Text(src []rune) string {
	return string(src[n.start:n.end])
}

func (n *NodeArg) PrevSibling() INode {
	if n.parent == nil || n.index == 0 {
		return nil
	}
	i := n.index - 1
	if i < 0 || i >= len(n.parent.Children()) {
		return nil
	}
	return n.parent.Children()[i]
}

func (n *NodeArg) NextSibling() INode {
	if n.parent == nil {
		return nil
	}
	i := n.index + 1
	if i < 0 || i >= len(n.parent.Children()) {
		return nil
	}
	return n.parent.Children()[i]
}

func (n *NodeArg) Parent() INode {
	return n.parent
}

func (n *NodeArg) Index() int {
	return n.index
}

func (n *NodeArg) Children() []INode {
	return n.children
}

func (n *NodeArg) ParamKind() ParameterKind {
	return n.paramKind
}

func (n *NodeArg) IsInside(pos uint32) bool {
	return n.start <= pos && pos < n.end
}

func (n *NodeArg) ArgParamSpec() (ParameterSpec, bool) {
	commandNode, ok := n.parent.(INodeCommand)
	if !ok {
		return ParameterSpec{}, false
	}
	argIndex := n.index
	return commandNode.ParamSpec(argIndex)
}
