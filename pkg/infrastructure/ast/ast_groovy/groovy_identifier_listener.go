package ast_groovy

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/languages/groovy"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"reflect"
	"strings"
)

var nodeDeps []core_domain.CodeDependency

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

func (s *GroovyIdentifierListener) GetDepsInfo() []core_domain.CodeDependency {
	return nodeDeps
}

func buildGroovyMap(pathExprCtx *parser.PathExpressionContext) []core_domain.CodeDependency {
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

func buildBlockStatements(closureContext *parser.ClosureContext) []core_domain.CodeDependency {
	var results []core_domain.CodeDependency
	statementsContext := closureContext.BlockStatementsOpt().(*parser.BlockStatementsOptContext).BlockStatements().(*parser.BlockStatementsContext)
	for _, blockStatement := range statementsContext.AllBlockStatement() {
		var result *core_domain.CodeDependency = nil

		commandExprCtx := blockStatement.GetChild(0).GetChild(0).GetChild(0).(*parser.CommandExpressionContext)
		pathExpression := commandExprCtx.GetChild(0).(*parser.PostfixExprAltForExprContext).GetChild(0).(*parser.PostfixExpressionContext).PathExpression()
		scope := pathExpression.GetChild(0).(antlr.ParseTree).GetText()

		//  with quote testImplementation('org.springframework.boot:spring-boot-starter-test')
		isWithQuote := pathExpression.GetChildCount() >= 2
		if isWithQuote {
			argumentsContext := pathExpression.GetChild(1).(*parser.PathElementContext).GetChild(0).(*parser.ArgumentsContext)
			argListCtx := argumentsContext.GetChild(1).(*parser.EnhancedArgumentListContext)
			for _, argElement := range argListCtx.AllEnhancedArgumentListElement() {
				result = ConvertToJDep(argElement.GetText())
			}
		}

		// normal: developmentOnly 'org.springframework.boot:spring-boot-devtools'
		if commandExprCtx.GetChildCount() >= 2 {
			argumentListContext := commandExprCtx.GetChild(1).(*parser.ArgumentListContext)
			result = BuildDependency(argumentListContext)
		}

		if result != nil {
			result.Scope = scope
			results = append(results, *result)
		}
	}

	return results
}

func BuildDependency(argumentListContext *parser.ArgumentListContext) *core_domain.CodeDependency {
	var result *core_domain.CodeDependency = nil
	for _, arg := range argumentListContext.AllArgumentListElement() {
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
	return result
}

func ConvertToJDep(result string) *core_domain.CodeDependency {
	withQuote := strings.ReplaceAll(result, "'", "")
	split := strings.Split(withQuote, ":")
	return core_domain.NewCodeDependency(split[0], split[1])
}
