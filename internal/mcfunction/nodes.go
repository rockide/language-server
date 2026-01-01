package mcfunction

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

type NodeParam interface {
	ParamSpec() (ParameterSpec, bool)
}

type INodeArg interface {
	INode
	NodeParam
	ParamKind() ParameterKind
}

type INodeArgMap interface {
	INodeArg
	MapSpec() *MapSpec
}

type INodeArgPairChild interface {
	INodeArg
	PairKind() PairKind
	Keys() []string
	KeySpec() (ParameterSpec, bool)
	ValueSpec() (ParameterSpec, bool)
}

type INodeCommand interface {
	INode
	CommandName() string
	Args() []INode
	Spec() *Spec
	OverloadStates() []*overloadState
	ParamSpecAt(index int) (ParameterSpec, bool)
	IsValid() bool
}
