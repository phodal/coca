package astutil

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

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
