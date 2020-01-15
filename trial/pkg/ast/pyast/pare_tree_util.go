package pyast

import "github.com/antlr/antlr4/runtime/Go/antlr"

func GetNodeIndex(node antlr.ParseTree) int {
	if node == nil || node.GetParent() == nil {
		return -1
	}
	parent := node.GetParent()

	for i := 0; i < parent.GetChildCount(); i++ {
		if parent.GetChild(i) == node {
			return i
		}
	}
	return 0
}

func GetLeftSibling(ctx antlr.ParseTree) antlr.Tree {
	index := GetNodeIndex(ctx)
	if index < 1 {
		return nil
	}
	return ctx.GetParent().GetChild(index - 1)
}
