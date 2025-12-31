package mcfunction

type NodeKind uint8

// file
// ├─ comment
// ├─ command
// │  ├─ command lit
// │  ├─ command arg
// ├─ invalid command
// │  ├─ command lit
// │  ├─ command arg

const (
	NodeKindFile NodeKind = iota
	NodeKindComment
	NodeKindCommand
	NodeKindInvalidCommand
	NodeKindCommandLit
	NodeKindCommandArg
)

type INode interface {
	addChild(child INode)
	setParent(parent INode)
	setIndex(index int)

	Kind() NodeKind
	Range() (start, end uint32)
	Text([]rune) string
	PrevSibling() INode
	NextSibling() INode
	Parent() INode
	Index() int
	Children() []INode
	IsInside(pos uint32) bool
}

type Node struct {
	kind     NodeKind
	parent   INode
	index    int
	children []INode
	start    uint32
	end      uint32
}

func (n *Node) addChild(child INode) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *Node) setParent(parent INode) {
	n.parent = parent
}

func (n *Node) setIndex(index int) {
	n.index = index
}

func (n *Node) Kind() NodeKind {
	return n.kind
}

func (n *Node) Range() (start, end uint32) {
	return n.start, n.end
}

func (n *Node) Text(src []rune) string {
	return string(src[n.start:n.end])
}

func (n *Node) PrevSibling() INode {
	if n.parent == nil || n.index == 0 {
		return nil
	}
	i := n.index - 1
	if i < 0 || i >= len(n.parent.Children()) {
		return nil
	}
	return n.parent.Children()[i]
}

func (n *Node) NextSibling() INode {
	if n.parent == nil || n.index+1 >= len(n.parent.Children()) {
		return nil
	}
	return n.parent.Children()[n.index+1]
}

func (n *Node) Parent() INode {
	return n.parent
}

func (n *Node) Index() int {
	return n.index
}

func (n *Node) Children() []INode {
	return n.children
}

func (n *Node) IsInside(pos uint32) bool {
	return n.start <= pos && pos < n.end
}
