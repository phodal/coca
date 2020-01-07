package groovy

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/domain"
	parser "github.com/phodal/coca/languages/groovy"
	"reflect"
	"strings"
)

var nodeDeps []domain.JDependency

type GroovyIdentifierListener struct {
	parser.BaseGroovyParserListener
}

func NewGroovyIdentListener() *GroovyIdentifierListener {
	nodeDeps = nil
	return &GroovyIdentifierListener{}
}

// TODO :
// 1. delete groovy build code
// 2. use regex replace it
// 3. remove the features
func (s *GroovyIdentifierListener) EnterScriptStatement(ctx *parser.ScriptStatementContext) {
	if reflect.TypeOf(ctx.GetChild(0)).String() == "*parser.ExpressionStmtAltContext" {
		cmdExpr := ctx.GetChild(0).(*parser.ExpressionStmtAltContext).StatementExpression().GetChild(0).(*parser.CommandExpressionContext).Expression()
		if cmdExpr != nil {
			if reflect.TypeOf(cmdExpr.GetChild(0)).String() == "*parser.PostfixExpressionContext" {
				pathExprCtx := cmdExpr.GetChild(0).(*parser.PostfixExpressionContext).GetChild(0).(*parser.PathExpressionContext)
				buildGroovyMap(pathExprCtx)
			}
		}
	}
}

func (s *GroovyIdentifierListener) GetDepsInfo() []domain.JDependency {
	return nodeDeps
}

func buildGroovyMap(pathExprCtx *parser.PathExpressionContext) []domain.JDependency {
	if reflect.TypeOf(pathExprCtx.GetChild(0)).String() == "*parser.IdentifierPrmrAltContext" {
		if pathExprCtx.GetChild(0).(antlr.ParseTree).GetText() != "dependencies" {
			return nil
		}
	}
	pathChild := pathExprCtx.GetChild(1)
	if pathChild != nil {
		pathElement := pathChild.(*parser.PathElementContext)

		if pathElement.ClosureOrLambdaExpression() != nil {
			expressionContext := pathElement.ClosureOrLambdaExpression().(*parser.ClosureOrLambdaExpressionContext)
			if reflect.TypeOf(expressionContext.GetChild(0)).String() == "*parser.ClosureContext" {
				closureContext := expressionContext.GetChild(0).(*parser.ClosureContext)
				nodeDeps = buildBlockStatements(closureContext)
				return nodeDeps
			}
		}
	}
	return nil
}

func buildBlockStatements(closureContext *parser.ClosureContext) []domain.JDependency {
	var results []domain.JDependency
	statementsContext := closureContext.BlockStatementsOpt().(*parser.BlockStatementsOptContext).BlockStatements().(*parser.BlockStatementsContext)
	for _, blockStatement := range statementsContext.AllBlockStatement() {
		child := blockStatement.GetChild(0).GetChild(0).GetChild(0).(*parser.CommandExpressionContext)
		declare := child.GetChild(0).(antlr.ParseTree).GetText()

		if child.GetChildCount() < 2 {
			continue
		}

		var result *domain.JDependency = nil
		for _, arg := range child.GetChild(1).(antlr.ParseTree).(*parser.ArgumentListContext).AllArgumentListElement() {
			if reflect.TypeOf(arg.(*parser.ArgumentListElementContext).GetChild(0)).String() == "*parser.ExpressionListElementContext" {
				listElementContext := arg.(*parser.ArgumentListElementContext).GetChild(0).(*parser.ExpressionListElementContext)
				literalPrmrAltContext := listElementContext.
					GetChild(0).
					GetChild(0).
					GetChild(0).
					GetChild(0).
				(*parser.LiteralPrmrAltContext)

				resultStr := literalPrmrAltContext.Literal().GetChild(0).(*parser.StringLiteralContext).StringLiteral().GetText()
				result = ConvertToJDep(resultStr)
			}
		}

		if result != nil {
			result.Scope = declare
			results = append(results, *result)
		}
	}

	return results
}

func ConvertToJDep(result string) *domain.JDependency {
	withQuote := strings.ReplaceAll(result, "'", "")
	split := strings.Split(withQuote, ":")
	return domain.NewJDependency(split[0], split[1])
}
