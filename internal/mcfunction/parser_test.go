package mcfunction_test

import (
	"testing"

	"github.com/rockide/language-server/internal/mcfunction"
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
	cmdNode, ok := root.Children()[0].(*mcfunction.NodeCommand)
	if !ok {
		t.Fatalf("Expected child node to be NodeCommand, got %T", root.Children()[0])
	}
	if cmdNode.Kind() != mcfunction.NodeKindCommand {
		t.Fatalf("Expected NodeKindCommand, got %v", cmdNode.Kind())
	}
}
