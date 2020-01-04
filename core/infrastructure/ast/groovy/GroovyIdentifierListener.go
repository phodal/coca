package groovy

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/groovy"
	"reflect"
)

type GroovyIdentifierListener struct {
	parser.BaseGroovyParserListener
}

func (s *GroovyIdentifierListener) EnterScriptStatement(ctx *parser.ScriptStatementContext) {
	fmt.Println("EnterScriptStatement")
	if reflect.TypeOf(ctx.GetChild(0)).String() == "*parser.ExpressionStmtAltContext" {
		cmdExpr := ctx.GetChild(0).(*parser.ExpressionStmtAltContext).StatementExpression().GetChild(0).(*parser.CommandExpressionContext).Expression()
		if cmdExpr != nil {
			postfixExpressionContext := cmdExpr.GetChild(0).(*parser.PostfixExpressionContext)
			if postfixExpressionContext != nil {
				pathExprCtx := postfixExpressionContext.GetChild(0).(*parser.PathExpressionContext)
				buildGroovyMap(pathExprCtx)
			}
		}
	}
}

func buildGroovyMap(pathExprCtx *parser.PathExpressionContext) {
	if reflect.TypeOf(pathExprCtx.GetChild(0)).String() == "*parser.IdentifierPrmrAltContext" {
		fmt.Println(pathExprCtx.GetChild(0).(antlr.ParseTree).GetText())
	}
	pathChild := pathExprCtx.GetChild(1)
	if pathChild != nil {
		pathElement := pathChild.(*parser.PathElementContext)

		if pathElement.ClosureOrLambdaExpression() != nil {
			expressionContext := pathElement.ClosureOrLambdaExpression().(*parser.ClosureOrLambdaExpressionContext)
			if reflect.TypeOf(expressionContext.GetChild(0)).String() == "*parser.ClosureContext" {
				closureContext := expressionContext.GetChild(0).(*parser.ClosureContext)
				buildBlockStatements(closureContext)
			}
		}
	}
}

func buildBlockStatements(closureContext *parser.ClosureContext) {
	statementsContext := closureContext.BlockStatementsOpt().(*parser.BlockStatementsOptContext).BlockStatements().(*parser.BlockStatementsContext)
	for _, blockStatement := range statementsContext.AllBlockStatement() {
		child := blockStatement.GetChild(0).GetChild(0).GetChild(0).(*parser.CommandExpressionContext)

		fmt.Println(child.GetChild(0).(antlr.ParseTree).GetText())
		fmt.Println(reflect.TypeOf(child.GetChild(1).(antlr.ParseTree).GetText()))
	}
}
