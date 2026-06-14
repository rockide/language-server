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

type node struct {
	kind     NodeKind
	parent   Node
	index    int
	children []Node
	start    uint32
	end      uint32
}

func (n *node) addChild(child Node) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *node) setParent(parent Node) {
	n.parent = parent
}

func (n *node) setIndex(index int) {
	n.index = index
}

func (n *node) Kind() NodeKind {
	return n.kind
}

func (n *node) Range() (start, end uint32) {
	return n.start, n.end
}

func (n *node) Text(src []rune) string {
	return string(src[n.start:n.end])
}

func (n *node) PrevSibling() Node {
	if n.parent == nil || n.index == 0 {
		return nil
	}
	i := n.index - 1
	if i < 0 || i >= len(n.parent.Children()) {
		return nil
	}
	return n.parent.Children()[i]
}

func (n *node) NextSibling() Node {
	if n.parent == nil || n.index+1 >= len(n.parent.Children()) {
		return nil
	}
	return n.parent.Children()[n.index+1]
}

func (n *node) Parent() Node {
	return n.parent
}

func (n *node) Index() int {
	return n.index
}

func (n *node) Children() []Node {
	return n.children
}

func (n *node) IsInside(pos uint32) bool {
	return n.start <= pos && pos < n.end
}
