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

func NewGroovyIdentListener() *GroovyIdentifierListener {
	return &GroovyIdentifierListener{}
}

// TODO :
// 1. delete groovy build code
// 2. use regex replace it
// 3. remove the features
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
		declare := child.GetChild(0).(antlr.ParseTree).GetText()
		var deps []string = nil

		for _, arg := range child.GetChild(1).(antlr.ParseTree).(*parser.ArgumentListContext).AllArgumentListElement() {
			if reflect.TypeOf(arg.(*parser.ArgumentListElementContext).GetChild(0)).String() == "*parser.ExpressionListElementContext" {
				listElementContext := arg.(*parser.ArgumentListElementContext).GetChild(0).(*parser.ExpressionListElementContext)
				literalPrmrAltContext := listElementContext.
					GetChild(0).
					GetChild(0).
					GetChild(0).
					GetChild(0).
				(*parser.LiteralPrmrAltContext)

				result := literalPrmrAltContext.Literal().GetChild(0).(*parser.StringLiteralContext).StringLiteral().GetText()
				deps = append(deps, result)
			}
		}

		fmt.Println(declare, deps)
	}
}
