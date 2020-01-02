package wadapter

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/adapter/coca_file"
	"github.com/phodal/coca/core/domain/bs_domain"
	"github.com/phodal/coca/core/infrastructure/ast/bs"
)

type WBadSmell struct {
}

func (w *WBadSmell) Analysis(code string) bs_domain.BsJClass {
	parser := coca_file.ProcessString(code)
	context := parser.CompilationUnit()

	listener := bs.NewBadSmellListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetNodeInfo()
}
