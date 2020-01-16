package ast_util

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/domain/core_domain"
)

func AddPosition(m *domain.JMethod, ctx *antlr.BaseParserRuleContext) {
	m.StartLine = ctx.GetStart().GetLine()
	m.StartLinePosition = ctx.GetStart().GetColumn()
	m.StopLine = ctx.GetStop().GetLine()
	m.StopLinePosition = ctx.GetStop().GetColumn()
}


func AddFunctionPosition(m *core_domain.CodeFunction, ctx *antlr.BaseParserRuleContext) {
	m.Position.StartLine = ctx.GetStart().GetLine()
	m.Position.StartLinePosition = ctx.GetStart().GetColumn()
	m.Position.StopLine = ctx.GetStop().GetLine()
	m.Position.StopLinePosition = ctx.GetStop().GetColumn()
}

