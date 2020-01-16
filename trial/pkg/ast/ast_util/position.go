package ast_util

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/domain"
	"github.com/phodal/coca/pkg/domain/trial"
)

func AddPosition(m *domain.JMethod, ctx *antlr.BaseParserRuleContext) {
	m.StartLine = ctx.GetStart().GetLine()
	m.StartLinePosition = ctx.GetStart().GetColumn()
	m.StopLine = ctx.GetStop().GetLine()
	m.StopLinePosition = ctx.GetStop().GetColumn()
}


func AddFunctionPosition(m *trial.CodeFunction, ctx *antlr.BaseParserRuleContext) {
	m.CodePosition.StartLine = ctx.GetStart().GetLine()
	m.CodePosition.StartLinePosition = ctx.GetStart().GetColumn()
	m.CodePosition.StopLine = ctx.GetStop().GetLine()
	m.CodePosition.StopLinePosition = ctx.GetStop().GetColumn()
}

