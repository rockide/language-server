package mcfunction

import (
	"testing"
)

func TestNodeInterfacesSatisfied(t *testing.T) {
	var (
		_ Node             = (*node)(nil)
		_ Node             = (*nodeArg)(nil)
		_ Node             = (*nodeCommand)(nil)
		_ Node             = (*nodeArgCommand)(nil)
		_ Node             = (*nodeArgPair)(nil)
		_ Node             = (*nodeArgPairChild)(nil)
		_ Node             = (*nodeArgMap)(nil)
		_ ArgNode          = (*nodeArg)(nil)
		_ ArgNode          = (*nodeArgCommand)(nil)
		_ ArgNode          = (*nodeArgPair)(nil)
		_ ArgNode          = (*nodeArgMap)(nil)
		_ ArgNode          = (*nodeArgPairChild)(nil)
		_ MapNode       = (*nodeArgMap)(nil)
		_ PairChildNode = (*nodeArgPairChild)(nil)
		_ CommandNode      = (*nodeCommand)(nil)
		_ CommandNode      = (*nodeArgCommand)(nil)
		_ ParamNode        = (*nodeArg)(nil)
		_ ParamNode        = (*nodeArgCommand)(nil)
		_ ParamNode        = (*nodeArgPair)(nil)
		_ ParamNode        = (*nodeArgPairChild)(nil)
	)

	src := []rune("gamemode")

	keySpec := &ParameterSpec{Kind: ParameterKindString, Name: "key"}
	valueSpec := &ParameterSpec{Kind: ParameterKindInteger, Name: "value"}
	pair := &nodeArgPair{
		nodeArg:   &nodeArg{node: &node{}},
		keySpec:   keySpec,
		valueSpec: valueSpec,
	}
	pairChild := &nodeArgPairChild{
		nodeArg:  &nodeArg{node: &node{}},
		pairKind: PairKindKey,
	}
	pair.addChild(pairChild)
	if got := pairChild.PairKind(); got != PairKindKey {
		t.Errorf("PairKind: got %v, want %v", got, PairKindKey)
	}
	ks, ok := pairChild.KeySpec()
	if !ok || ks.Name != "key" {
		t.Errorf("KeySpec: got %v, %v", ks, ok)
	}
	vs, ok := pairChild.ValueSpec()
	if !ok || vs.Name != "value" {
		t.Errorf("ValueSpec: got %v, %v", vs, ok)
	}

	arg := &nodeArg{
		node:      &node{kind: NodeKindCommandArg, start: 4, end: 5},
		paramKind: ParameterKindSelector,
	}
	if got := arg.Kind(); got != NodeKindCommandArg {
		t.Errorf("nodeArg.Kind promoted: got %v, want %v", got, NodeKindCommandArg)
	}
	if got := arg.ParamKind(); got != ParameterKindSelector {
		t.Errorf("nodeArg.ParamKind: got %v, want %v", got, ParameterKindSelector)
	}
	if s, e := arg.Range(); s != 4 || e != 5 {
		t.Errorf("nodeArg.Range promoted: got (%d,%d), want (4,5)", s, e)
	}
	if !arg.IsInside(4) || arg.IsInside(5) {
		t.Errorf("nodeArg.IsInside promoted: failed at 4 (true) and 5 (false)")
	}
	if got := arg.Text(src); got != "m" {
		t.Errorf("nodeArg.Text promoted: got %q, want %q", got, "m")
	}

	cmd := &nodeCommand{
		node: &node{kind: NodeKindCommand},
		name: "gamemode",
	}
	cmd.addChild(arg)
	if got := cmd.CommandName(); got != "gamemode" {
		t.Errorf("nodeCommand.CommandName: got %v, want %v", got, "gamemode")
	}
	if got := cmd.Args(); len(got) != 1 || got[0] != arg {
		t.Errorf("nodeCommand.Args: got %v, want [%v]", got, arg)
	}
	if n := arg.PrevSibling(); n != nil {
		t.Errorf("nodeArg.PrevSibling promoted: got %T, want nil", n)
	}
}

func TestNodeTreeShape(t *testing.T) {
	src := []rune("say hello world")

	root := &node{kind: NodeKindFile, start: 0, end: uint32(len(src))}

	cmd := &nodeCommand{
		node: &node{kind: NodeKindCommand, start: 0, end: uint32(len(src))},
		name: "say",
	}
	root.addChild(cmd)

	lit := &node{kind: NodeKindCommandLit, start: 0, end: 3}
	cmd.addChild(lit)

	arg1 := &nodeArg{
		node:      &node{kind: NodeKindCommandArg, start: 4, end: 9},
		paramKind: ParameterKindString,
	}
	cmd.addChild(arg1)

	arg2 := &nodeArg{
		node:      &node{kind: NodeKindCommandArg, start: 10, end: 15},
		paramKind: ParameterKindString,
	}
	cmd.addChild(arg2)

	if got := root.Text(src); got != "say hello world" {
		t.Errorf("root.Text: got %q, want %q", got, "say hello world")
	}
	if got := lit.Text(src); got != "say" {
		t.Errorf("lit.Text: got %q, want %q", got, "say")
	}
	if got := arg1.Text(src); got != "hello" {
		t.Errorf("arg1.Text: got %q, want %q", got, "hello")
	}
	if got := arg2.Text(src); got != "world" {
		t.Errorf("arg2.Text: got %q, want %q", got, "world")
	}

	if got := cmd.Args(); len(got) != 2 {
		t.Fatalf("cmd.Args: got %d, want 2", len(got))
	}
	if n := arg1.PrevSibling(); n != lit {
		t.Errorf("arg1.PrevSibling: got %T, want lit", n)
	}
	if n := arg1.NextSibling(); n != arg2 {
		t.Errorf("arg1.NextSibling: got %T, want arg2", n)
	}
	if n := lit.Parent(); n != cmd {
		t.Errorf("lit.Parent: got %T, want cmd", n)
	}
	if n := arg1.Parent(); n != cmd {
		t.Errorf("arg1.Parent: got %T, want cmd", n)
	}
	if n := cmd.Parent(); n != root {
		t.Errorf("cmd.Parent: got %T, want root", n)
	}
}

func TestNodeArgPairAndMap(t *testing.T) {
	keySpec := &ParameterSpec{Kind: ParameterKindString, Name: "k"}
	pair := &nodeArgPair{nodeArg: &nodeArg{node: &node{}}, keySpec: keySpec}
	child := &nodeArgPairChild{nodeArg: &nodeArg{node: &node{}}, pairKind: PairKindValue}
	pair.addChild(child)
	ks, ok := child.KeySpec()
	if !ok || ks.Name != "k" {
		t.Errorf("child.KeySpec: got %v, %v", ks, ok)
	}
	_, ok = child.ValueSpec()
	if ok {
		t.Errorf("child.ValueSpec: expected false when no value spec")
	}

	arg := &nodeArg{node: &node{kind: NodeKindCommandArg}, paramKind: ParameterKindMap}
	m := &nodeArgMap{nodeArg: arg, mapSpec: selectorArg}
	if m.MapSpec() != selectorArg {
		t.Errorf("MapSpec: got %v, want selectorArg", m.MapSpec())
	}
	if m.Kind() != NodeKindCommandArg {
		t.Errorf("nodeArgMap.Kind: got %v, want NodeKindCommandArg", m.Kind())
	}
}
