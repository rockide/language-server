package mcfunction_test

import (
	"testing"

	"github.com/rockide/language-server/internal/mcfunction"
	"github.com/rockide/language-server/internal/mcfunction/builtin"
)

func TestQuickEvent(t *testing.T) {
	parser := mcfunction.NewParser(mcfunction.ParserOptions{
		EventAlias: true,
	})
	root, err := parser.Parse([]rune("@s foo:bar"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if root == nil {
		t.Fatal("Expected non-nil root node")
	}
	if root.Kind() != mcfunction.NodeKindFile {
		t.Fatalf("Expected NodeKindCommand, got %v", root.Kind())
	}
	if len(root.Children()) != 1 {
		t.Fatalf("Expected 1 child node, got %d", len(root.Children()))
	}
	cmdNode, ok := root.Children()[0].(mcfunction.CommandNode)
	if !ok {
		t.Fatalf("Expected child node to be CommandNode, got %T", root.Children()[0])
	}
	if cmdNode.Kind() != mcfunction.NodeKindCommand {
		t.Fatalf("Expected NodeKindCommand, got %v", cmdNode.Kind())
	}
}

func argAs(t *testing.T, n mcfunction.Node) mcfunction.ArgNode {
	t.Helper()
	a, ok := n.(mcfunction.ArgNode)
	if !ok {
		t.Fatalf("expected ArgNode, got %T", n)
	}
	return a
}

func TestParseSay(t *testing.T) {
	parser := mcfunction.NewParser(mcfunction.ParserOptions{}, builtin.Commands...)
	src := []rune("say hello world")
	root, err := parser.Parse(src)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	cmd, ok := root.Children()[0].(mcfunction.CommandNode)
	if !ok {
		t.Fatalf("child type: got %T, want CommandNode", root.Children()[0])
	}
	if cmd.CommandName() != "say" {
		t.Errorf("CommandName: got %q, want %q", cmd.CommandName(), "say")
	}
	if !cmd.IsValid() {
		t.Errorf("IsValid: expected true")
	}
	if len(cmd.Args()) != 2 {
		t.Fatalf("Args: got %d, want 2", len(cmd.Args()))
	}
	first := argAs(t, cmd.Args()[0])
	if first.ParamKind() != mcfunction.ParameterKindString {
		t.Errorf("Args[0].ParamKind: got %v, want KindString", first.ParamKind())
	}
	second := argAs(t, cmd.Args()[1])
	if got := second.Text(src); got != "world" {
		t.Errorf("Args[1].Text: got %q, want %q", got, "world")
	}
	if n := first.NextSibling(); n != second {
		t.Errorf("first.NextSibling: got %T, want Args[1]", n)
	}
	if n := second.PrevSibling(); n != first {
		t.Errorf("second.PrevSibling: got %T, want Args[0]", n)
	}
}

func TestParseGamemode(t *testing.T) {
	parser := mcfunction.NewParser(mcfunction.ParserOptions{}, builtin.Commands...)
	root, err := parser.Parse([]rune("gamemode creative @a"))
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	cmd := root.Children()[0].(mcfunction.CommandNode)
	if cmd.CommandName() != "gamemode" {
		t.Fatalf("CommandName: got %q, want %q", cmd.CommandName(), "gamemode")
	}
	if !cmd.IsValid() {
		t.Errorf("IsValid: expected true")
	}
	args := cmd.Args()
	if len(args) != 2 {
		t.Fatalf("Args: got %d, want 2", len(args))
	}
	if a := argAs(t, args[0]); a.ParamKind() != mcfunction.ParameterKindLiteral {
		t.Errorf("Args[0].ParamKind: got %v, want KindLiteral", a.ParamKind())
	}
	if a := argAs(t, args[1]); a.ParamKind() != mcfunction.ParameterKindSelector {
		t.Errorf("Args[1].ParamKind: got %v, want KindSelector", a.ParamKind())
	}
}

func TestParseTellraw(t *testing.T) {
	parser := mcfunction.NewParser(mcfunction.ParserOptions{}, builtin.Commands...)
	root, err := parser.Parse([]rune(`tellraw @a {"rawtext":[{"text":"hi"}]}`))
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	cmd := root.Children()[0].(mcfunction.CommandNode)
	if cmd.CommandName() != "tellraw" {
		t.Fatalf("CommandName: got %q, want %q", cmd.CommandName(), "tellraw")
	}
	if !cmd.IsValid() {
		t.Errorf("IsValid: expected true")
	}
	args := cmd.Args()
	if len(args) != 2 {
		t.Fatalf("Args: got %d, want 2", len(args))
	}
	if a := argAs(t, args[0]); a.ParamKind() != mcfunction.ParameterKindSelector {
		t.Errorf("Args[0].ParamKind: got %v, want KindSelector", a.ParamKind())
	}
	if a := argAs(t, args[1]); a.ParamKind() != mcfunction.ParameterKindRawMessage {
		t.Errorf("Args[1].ParamKind: got %v, want KindRawMessage", a.ParamKind())
	}
}

func TestParseUnknownCommand(t *testing.T) {
	parser := mcfunction.NewParser(mcfunction.ParserOptions{}, builtin.Commands...)
	root, err := parser.Parse([]rune("xyz"))
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	cmd := root.Children()[0].(mcfunction.CommandNode)
	if cmd.Kind() != mcfunction.NodeKindInvalidCommand {
		t.Errorf("Kind: got %v, want NodeKindInvalidCommand", cmd.Kind())
	}
	if cmd.IsValid() {
		t.Errorf("IsValid: expected false for unknown command")
	}
}
