package astutil

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/domain/core_domain"
)

func AddFunctionPosition(m *core_domain.CodeFunction, ctx *antlr.BaseParserRuleContext) {
	m.Position.StartLine = ctx.GetStart().GetLine()
	m.Position.StartLinePosition = ctx.GetStart().GetColumn()
	m.Position.StopLine = ctx.GetStop().GetLine()
	m.Position.StopLinePosition = ctx.GetStop().GetColumn()
}

