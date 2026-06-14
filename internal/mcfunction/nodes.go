package mcfunction

type Node interface {
	addChild(child Node)
	setParent(parent Node)
	setIndex(index int)

	Kind() NodeKind
	Range() (start, end uint32)
	Text([]rune) string
	PrevSibling() Node
	NextSibling() Node
	Parent() Node
	Index() int
	Children() []Node
	IsInside(pos uint32) bool
}

type ParamNode interface {
	ParamSpec() (ParameterSpec, bool)
}

type ArgNode interface {
	Node
	ParamNode
	ParamKind() ParameterKind
	CommandNode() CommandNode
}

type MapNode interface {
	ArgNode
	MapSpec() *MapSpec
}

type PairChildNode interface {
	ArgNode
	PairKind() PairKind
	Keys() []string
	KeySpec() (ParameterSpec, bool)
	ValueSpec() (ParameterSpec, bool)
}

type CommandNode interface {
	Node
	CommandName() string
	Args() []Node
	Spec() *Spec
	OverloadStates() []*overloadState
	ParamSpecAt(index int) (ParameterSpec, bool)
	IsValid() bool
}
