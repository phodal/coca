// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JavaScriptParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseJavaScriptParserListener is a complete listener for a parse tree produced by JavaScriptParser.
type BaseJavaScriptParserListener struct{}

var _ JavaScriptParserListener = &BaseJavaScriptParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJavaScriptParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJavaScriptParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJavaScriptParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJavaScriptParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseJavaScriptParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseJavaScriptParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterSourceElement is called when production sourceElement is entered.
func (s *BaseJavaScriptParserListener) EnterSourceElement(ctx *SourceElementContext) {}

// ExitSourceElement is called when production sourceElement is exited.
func (s *BaseJavaScriptParserListener) ExitSourceElement(ctx *SourceElementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJavaScriptParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJavaScriptParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseJavaScriptParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseJavaScriptParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseJavaScriptParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseJavaScriptParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseJavaScriptParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseJavaScriptParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterImportFromBlock is called when production importFromBlock is entered.
func (s *BaseJavaScriptParserListener) EnterImportFromBlock(ctx *ImportFromBlockContext) {}

// ExitImportFromBlock is called when production importFromBlock is exited.
func (s *BaseJavaScriptParserListener) ExitImportFromBlock(ctx *ImportFromBlockContext) {}

// EnterModuleItems is called when production moduleItems is entered.
func (s *BaseJavaScriptParserListener) EnterModuleItems(ctx *ModuleItemsContext) {}

// ExitModuleItems is called when production moduleItems is exited.
func (s *BaseJavaScriptParserListener) ExitModuleItems(ctx *ModuleItemsContext) {}

// EnterImportDefault is called when production importDefault is entered.
func (s *BaseJavaScriptParserListener) EnterImportDefault(ctx *ImportDefaultContext) {}

// ExitImportDefault is called when production importDefault is exited.
func (s *BaseJavaScriptParserListener) ExitImportDefault(ctx *ImportDefaultContext) {}

// EnterImportNamespace is called when production importNamespace is entered.
func (s *BaseJavaScriptParserListener) EnterImportNamespace(ctx *ImportNamespaceContext) {}

// ExitImportNamespace is called when production importNamespace is exited.
func (s *BaseJavaScriptParserListener) ExitImportNamespace(ctx *ImportNamespaceContext) {}

// EnterImportFrom is called when production importFrom is entered.
func (s *BaseJavaScriptParserListener) EnterImportFrom(ctx *ImportFromContext) {}

// ExitImportFrom is called when production importFrom is exited.
func (s *BaseJavaScriptParserListener) ExitImportFrom(ctx *ImportFromContext) {}

// EnterAliasName is called when production aliasName is entered.
func (s *BaseJavaScriptParserListener) EnterAliasName(ctx *AliasNameContext) {}

// ExitAliasName is called when production aliasName is exited.
func (s *BaseJavaScriptParserListener) ExitAliasName(ctx *AliasNameContext) {}

// EnterExportDeclaration is called when production ExportDeclaration is entered.
func (s *BaseJavaScriptParserListener) EnterExportDeclaration(ctx *ExportDeclarationContext) {}

// ExitExportDeclaration is called when production ExportDeclaration is exited.
func (s *BaseJavaScriptParserListener) ExitExportDeclaration(ctx *ExportDeclarationContext) {}

// EnterExportDefaultDeclaration is called when production ExportDefaultDeclaration is entered.
func (s *BaseJavaScriptParserListener) EnterExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) {
}

// ExitExportDefaultDeclaration is called when production ExportDefaultDeclaration is exited.
func (s *BaseJavaScriptParserListener) ExitExportDefaultDeclaration(ctx *ExportDefaultDeclarationContext) {
}

// EnterExportFromBlock is called when production exportFromBlock is entered.
func (s *BaseJavaScriptParserListener) EnterExportFromBlock(ctx *ExportFromBlockContext) {}

// ExitExportFromBlock is called when production exportFromBlock is exited.
func (s *BaseJavaScriptParserListener) ExitExportFromBlock(ctx *ExportFromBlockContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseJavaScriptParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseJavaScriptParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseJavaScriptParserListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseJavaScriptParserListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseJavaScriptParserListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {
}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseJavaScriptParserListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {
}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseJavaScriptParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseJavaScriptParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseJavaScriptParserListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseJavaScriptParserListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseJavaScriptParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseJavaScriptParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseJavaScriptParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseJavaScriptParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterDoStatement is called when production DoStatement is entered.
func (s *BaseJavaScriptParserListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production DoStatement is exited.
func (s *BaseJavaScriptParserListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterWhileStatement is called when production WhileStatement is entered.
func (s *BaseJavaScriptParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production WhileStatement is exited.
func (s *BaseJavaScriptParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production ForStatement is entered.
func (s *BaseJavaScriptParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production ForStatement is exited.
func (s *BaseJavaScriptParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForInStatement is called when production ForInStatement is entered.
func (s *BaseJavaScriptParserListener) EnterForInStatement(ctx *ForInStatementContext) {}

// ExitForInStatement is called when production ForInStatement is exited.
func (s *BaseJavaScriptParserListener) ExitForInStatement(ctx *ForInStatementContext) {}

// EnterForOfStatement is called when production ForOfStatement is entered.
func (s *BaseJavaScriptParserListener) EnterForOfStatement(ctx *ForOfStatementContext) {}

// ExitForOfStatement is called when production ForOfStatement is exited.
func (s *BaseJavaScriptParserListener) ExitForOfStatement(ctx *ForOfStatementContext) {}

// EnterVarModifier is called when production varModifier is entered.
func (s *BaseJavaScriptParserListener) EnterVarModifier(ctx *VarModifierContext) {}

// ExitVarModifier is called when production varModifier is exited.
func (s *BaseJavaScriptParserListener) ExitVarModifier(ctx *VarModifierContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseJavaScriptParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseJavaScriptParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseJavaScriptParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseJavaScriptParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseJavaScriptParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseJavaScriptParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterYieldStatement is called when production yieldStatement is entered.
func (s *BaseJavaScriptParserListener) EnterYieldStatement(ctx *YieldStatementContext) {}

// ExitYieldStatement is called when production yieldStatement is exited.
func (s *BaseJavaScriptParserListener) ExitYieldStatement(ctx *YieldStatementContext) {}

// EnterWithStatement is called when production withStatement is entered.
func (s *BaseJavaScriptParserListener) EnterWithStatement(ctx *WithStatementContext) {}

// ExitWithStatement is called when production withStatement is exited.
func (s *BaseJavaScriptParserListener) ExitWithStatement(ctx *WithStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseJavaScriptParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseJavaScriptParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseJavaScriptParserListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseJavaScriptParserListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterCaseClauses is called when production caseClauses is entered.
func (s *BaseJavaScriptParserListener) EnterCaseClauses(ctx *CaseClausesContext) {}

// ExitCaseClauses is called when production caseClauses is exited.
func (s *BaseJavaScriptParserListener) ExitCaseClauses(ctx *CaseClausesContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseJavaScriptParserListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseJavaScriptParserListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterDefaultClause is called when production defaultClause is entered.
func (s *BaseJavaScriptParserListener) EnterDefaultClause(ctx *DefaultClauseContext) {}

// ExitDefaultClause is called when production defaultClause is exited.
func (s *BaseJavaScriptParserListener) ExitDefaultClause(ctx *DefaultClauseContext) {}

// EnterLabelledStatement is called when production labelledStatement is entered.
func (s *BaseJavaScriptParserListener) EnterLabelledStatement(ctx *LabelledStatementContext) {}

// ExitLabelledStatement is called when production labelledStatement is exited.
func (s *BaseJavaScriptParserListener) ExitLabelledStatement(ctx *LabelledStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseJavaScriptParserListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseJavaScriptParserListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseJavaScriptParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseJavaScriptParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatchProduction is called when production catchProduction is entered.
func (s *BaseJavaScriptParserListener) EnterCatchProduction(ctx *CatchProductionContext) {}

// ExitCatchProduction is called when production catchProduction is exited.
func (s *BaseJavaScriptParserListener) ExitCatchProduction(ctx *CatchProductionContext) {}

// EnterFinallyProduction is called when production finallyProduction is entered.
func (s *BaseJavaScriptParserListener) EnterFinallyProduction(ctx *FinallyProductionContext) {}

// ExitFinallyProduction is called when production finallyProduction is exited.
func (s *BaseJavaScriptParserListener) ExitFinallyProduction(ctx *FinallyProductionContext) {}

// EnterDebuggerStatement is called when production debuggerStatement is entered.
func (s *BaseJavaScriptParserListener) EnterDebuggerStatement(ctx *DebuggerStatementContext) {}

// ExitDebuggerStatement is called when production debuggerStatement is exited.
func (s *BaseJavaScriptParserListener) ExitDebuggerStatement(ctx *DebuggerStatementContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseJavaScriptParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseJavaScriptParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseJavaScriptParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseJavaScriptParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassTail is called when production classTail is entered.
func (s *BaseJavaScriptParserListener) EnterClassTail(ctx *ClassTailContext) {}

// ExitClassTail is called when production classTail is exited.
func (s *BaseJavaScriptParserListener) ExitClassTail(ctx *ClassTailContext) {}

// EnterClassElement is called when production classElement is entered.
func (s *BaseJavaScriptParserListener) EnterClassElement(ctx *ClassElementContext) {}

// ExitClassElement is called when production classElement is exited.
func (s *BaseJavaScriptParserListener) ExitClassElement(ctx *ClassElementContext) {}

// EnterMethodDefinition is called when production methodDefinition is entered.
func (s *BaseJavaScriptParserListener) EnterMethodDefinition(ctx *MethodDefinitionContext) {}

// ExitMethodDefinition is called when production methodDefinition is exited.
func (s *BaseJavaScriptParserListener) ExitMethodDefinition(ctx *MethodDefinitionContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseJavaScriptParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseJavaScriptParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameterArg is called when production formalParameterArg is entered.
func (s *BaseJavaScriptParserListener) EnterFormalParameterArg(ctx *FormalParameterArgContext) {}

// ExitFormalParameterArg is called when production formalParameterArg is exited.
func (s *BaseJavaScriptParserListener) ExitFormalParameterArg(ctx *FormalParameterArgContext) {}

// EnterLastFormalParameterArg is called when production lastFormalParameterArg is entered.
func (s *BaseJavaScriptParserListener) EnterLastFormalParameterArg(ctx *LastFormalParameterArgContext) {
}

// ExitLastFormalParameterArg is called when production lastFormalParameterArg is exited.
func (s *BaseJavaScriptParserListener) ExitLastFormalParameterArg(ctx *LastFormalParameterArgContext) {
}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseJavaScriptParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseJavaScriptParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterSourceElements is called when production sourceElements is entered.
func (s *BaseJavaScriptParserListener) EnterSourceElements(ctx *SourceElementsContext) {}

// ExitSourceElements is called when production sourceElements is exited.
func (s *BaseJavaScriptParserListener) ExitSourceElements(ctx *SourceElementsContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseJavaScriptParserListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseJavaScriptParserListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseJavaScriptParserListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseJavaScriptParserListener) ExitElementList(ctx *ElementListContext) {}

// EnterArrayElement is called when production arrayElement is entered.
func (s *BaseJavaScriptParserListener) EnterArrayElement(ctx *ArrayElementContext) {}

// ExitArrayElement is called when production arrayElement is exited.
func (s *BaseJavaScriptParserListener) ExitArrayElement(ctx *ArrayElementContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseJavaScriptParserListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseJavaScriptParserListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterPropertyExpressionAssignment is called when production PropertyExpressionAssignment is entered.
func (s *BaseJavaScriptParserListener) EnterPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// ExitPropertyExpressionAssignment is called when production PropertyExpressionAssignment is exited.
func (s *BaseJavaScriptParserListener) ExitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// EnterComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is entered.
func (s *BaseJavaScriptParserListener) EnterComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// ExitComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is exited.
func (s *BaseJavaScriptParserListener) ExitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// EnterFunctionProperty is called when production FunctionProperty is entered.
func (s *BaseJavaScriptParserListener) EnterFunctionProperty(ctx *FunctionPropertyContext) {}

// ExitFunctionProperty is called when production FunctionProperty is exited.
func (s *BaseJavaScriptParserListener) ExitFunctionProperty(ctx *FunctionPropertyContext) {}

// EnterPropertyGetter is called when production PropertyGetter is entered.
func (s *BaseJavaScriptParserListener) EnterPropertyGetter(ctx *PropertyGetterContext) {}

// ExitPropertyGetter is called when production PropertyGetter is exited.
func (s *BaseJavaScriptParserListener) ExitPropertyGetter(ctx *PropertyGetterContext) {}

// EnterPropertySetter is called when production PropertySetter is entered.
func (s *BaseJavaScriptParserListener) EnterPropertySetter(ctx *PropertySetterContext) {}

// ExitPropertySetter is called when production PropertySetter is exited.
func (s *BaseJavaScriptParserListener) ExitPropertySetter(ctx *PropertySetterContext) {}

// EnterPropertyShorthand is called when production PropertyShorthand is entered.
func (s *BaseJavaScriptParserListener) EnterPropertyShorthand(ctx *PropertyShorthandContext) {}

// ExitPropertyShorthand is called when production PropertyShorthand is exited.
func (s *BaseJavaScriptParserListener) ExitPropertyShorthand(ctx *PropertyShorthandContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseJavaScriptParserListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseJavaScriptParserListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseJavaScriptParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseJavaScriptParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseJavaScriptParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseJavaScriptParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterExpressionSequence is called when production expressionSequence is entered.
func (s *BaseJavaScriptParserListener) EnterExpressionSequence(ctx *ExpressionSequenceContext) {}

// ExitExpressionSequence is called when production expressionSequence is exited.
func (s *BaseJavaScriptParserListener) ExitExpressionSequence(ctx *ExpressionSequenceContext) {}

// EnterTemplateStringExpression is called when production TemplateStringExpression is entered.
func (s *BaseJavaScriptParserListener) EnterTemplateStringExpression(ctx *TemplateStringExpressionContext) {
}

// ExitTemplateStringExpression is called when production TemplateStringExpression is exited.
func (s *BaseJavaScriptParserListener) ExitTemplateStringExpression(ctx *TemplateStringExpressionContext) {
}

// EnterTernaryExpression is called when production TernaryExpression is entered.
func (s *BaseJavaScriptParserListener) EnterTernaryExpression(ctx *TernaryExpressionContext) {}

// ExitTernaryExpression is called when production TernaryExpression is exited.
func (s *BaseJavaScriptParserListener) ExitTernaryExpression(ctx *TernaryExpressionContext) {}

// EnterLogicalAndExpression is called when production LogicalAndExpression is entered.
func (s *BaseJavaScriptParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production LogicalAndExpression is exited.
func (s *BaseJavaScriptParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterPowerExpression is called when production PowerExpression is entered.
func (s *BaseJavaScriptParserListener) EnterPowerExpression(ctx *PowerExpressionContext) {}

// ExitPowerExpression is called when production PowerExpression is exited.
func (s *BaseJavaScriptParserListener) ExitPowerExpression(ctx *PowerExpressionContext) {}

// EnterPreIncrementExpression is called when production PreIncrementExpression is entered.
func (s *BaseJavaScriptParserListener) EnterPreIncrementExpression(ctx *PreIncrementExpressionContext) {
}

// ExitPreIncrementExpression is called when production PreIncrementExpression is exited.
func (s *BaseJavaScriptParserListener) ExitPreIncrementExpression(ctx *PreIncrementExpressionContext) {
}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *BaseJavaScriptParserListener) EnterObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {
}

// ExitObjectLiteralExpression is called when production ObjectLiteralExpression is exited.
func (s *BaseJavaScriptParserListener) ExitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {
}

// EnterMetaExpression is called when production MetaExpression is entered.
func (s *BaseJavaScriptParserListener) EnterMetaExpression(ctx *MetaExpressionContext) {}

// ExitMetaExpression is called when production MetaExpression is exited.
func (s *BaseJavaScriptParserListener) ExitMetaExpression(ctx *MetaExpressionContext) {}

// EnterInExpression is called when production InExpression is entered.
func (s *BaseJavaScriptParserListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production InExpression is exited.
func (s *BaseJavaScriptParserListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterLogicalOrExpression is called when production LogicalOrExpression is entered.
func (s *BaseJavaScriptParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production LogicalOrExpression is exited.
func (s *BaseJavaScriptParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterNotExpression is called when production NotExpression is entered.
func (s *BaseJavaScriptParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production NotExpression is exited.
func (s *BaseJavaScriptParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterPreDecreaseExpression is called when production PreDecreaseExpression is entered.
func (s *BaseJavaScriptParserListener) EnterPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {
}

// ExitPreDecreaseExpression is called when production PreDecreaseExpression is exited.
func (s *BaseJavaScriptParserListener) ExitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {}

// EnterArgumentsExpression is called when production ArgumentsExpression is entered.
func (s *BaseJavaScriptParserListener) EnterArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *BaseJavaScriptParserListener) ExitArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// EnterAwaitExpression is called when production AwaitExpression is entered.
func (s *BaseJavaScriptParserListener) EnterAwaitExpression(ctx *AwaitExpressionContext) {}

// ExitAwaitExpression is called when production AwaitExpression is exited.
func (s *BaseJavaScriptParserListener) ExitAwaitExpression(ctx *AwaitExpressionContext) {}

// EnterThisExpression is called when production ThisExpression is entered.
func (s *BaseJavaScriptParserListener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production ThisExpression is exited.
func (s *BaseJavaScriptParserListener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseJavaScriptParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseJavaScriptParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterUnaryMinusExpression is called when production UnaryMinusExpression is entered.
func (s *BaseJavaScriptParserListener) EnterUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// ExitUnaryMinusExpression is called when production UnaryMinusExpression is exited.
func (s *BaseJavaScriptParserListener) ExitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// EnterAssignmentExpression is called when production AssignmentExpression is entered.
func (s *BaseJavaScriptParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production AssignmentExpression is exited.
func (s *BaseJavaScriptParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterPostDecreaseExpression is called when production PostDecreaseExpression is entered.
func (s *BaseJavaScriptParserListener) EnterPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {
}

// ExitPostDecreaseExpression is called when production PostDecreaseExpression is exited.
func (s *BaseJavaScriptParserListener) ExitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {
}

// EnterTypeofExpression is called when production TypeofExpression is entered.
func (s *BaseJavaScriptParserListener) EnterTypeofExpression(ctx *TypeofExpressionContext) {}

// ExitTypeofExpression is called when production TypeofExpression is exited.
func (s *BaseJavaScriptParserListener) ExitTypeofExpression(ctx *TypeofExpressionContext) {}

// EnterInstanceofExpression is called when production InstanceofExpression is entered.
func (s *BaseJavaScriptParserListener) EnterInstanceofExpression(ctx *InstanceofExpressionContext) {}

// ExitInstanceofExpression is called when production InstanceofExpression is exited.
func (s *BaseJavaScriptParserListener) ExitInstanceofExpression(ctx *InstanceofExpressionContext) {}

// EnterUnaryPlusExpression is called when production UnaryPlusExpression is entered.
func (s *BaseJavaScriptParserListener) EnterUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// ExitUnaryPlusExpression is called when production UnaryPlusExpression is exited.
func (s *BaseJavaScriptParserListener) ExitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// EnterDeleteExpression is called when production DeleteExpression is entered.
func (s *BaseJavaScriptParserListener) EnterDeleteExpression(ctx *DeleteExpressionContext) {}

// ExitDeleteExpression is called when production DeleteExpression is exited.
func (s *BaseJavaScriptParserListener) ExitDeleteExpression(ctx *DeleteExpressionContext) {}

// EnterImportExpression is called when production ImportExpression is entered.
func (s *BaseJavaScriptParserListener) EnterImportExpression(ctx *ImportExpressionContext) {}

// ExitImportExpression is called when production ImportExpression is exited.
func (s *BaseJavaScriptParserListener) ExitImportExpression(ctx *ImportExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseJavaScriptParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseJavaScriptParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterBitXOrExpression is called when production BitXOrExpression is entered.
func (s *BaseJavaScriptParserListener) EnterBitXOrExpression(ctx *BitXOrExpressionContext) {}

// ExitBitXOrExpression is called when production BitXOrExpression is exited.
func (s *BaseJavaScriptParserListener) ExitBitXOrExpression(ctx *BitXOrExpressionContext) {}

// EnterSuperExpression is called when production SuperExpression is entered.
func (s *BaseJavaScriptParserListener) EnterSuperExpression(ctx *SuperExpressionContext) {}

// ExitSuperExpression is called when production SuperExpression is exited.
func (s *BaseJavaScriptParserListener) ExitSuperExpression(ctx *SuperExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseJavaScriptParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseJavaScriptParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterBitShiftExpression is called when production BitShiftExpression is entered.
func (s *BaseJavaScriptParserListener) EnterBitShiftExpression(ctx *BitShiftExpressionContext) {}

// ExitBitShiftExpression is called when production BitShiftExpression is exited.
func (s *BaseJavaScriptParserListener) ExitBitShiftExpression(ctx *BitShiftExpressionContext) {}

// EnterParenthesizedExpression is called when production ParenthesizedExpression is entered.
func (s *BaseJavaScriptParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production ParenthesizedExpression is exited.
func (s *BaseJavaScriptParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// EnterAdditiveExpression is called when production AdditiveExpression is entered.
func (s *BaseJavaScriptParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production AdditiveExpression is exited.
func (s *BaseJavaScriptParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production RelationalExpression is entered.
func (s *BaseJavaScriptParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production RelationalExpression is exited.
func (s *BaseJavaScriptParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterPostIncrementExpression is called when production PostIncrementExpression is entered.
func (s *BaseJavaScriptParserListener) EnterPostIncrementExpression(ctx *PostIncrementExpressionContext) {
}

// ExitPostIncrementExpression is called when production PostIncrementExpression is exited.
func (s *BaseJavaScriptParserListener) ExitPostIncrementExpression(ctx *PostIncrementExpressionContext) {
}

// EnterYieldExpression is called when production YieldExpression is entered.
func (s *BaseJavaScriptParserListener) EnterYieldExpression(ctx *YieldExpressionContext) {}

// ExitYieldExpression is called when production YieldExpression is exited.
func (s *BaseJavaScriptParserListener) ExitYieldExpression(ctx *YieldExpressionContext) {}

// EnterBitNotExpression is called when production BitNotExpression is entered.
func (s *BaseJavaScriptParserListener) EnterBitNotExpression(ctx *BitNotExpressionContext) {}

// ExitBitNotExpression is called when production BitNotExpression is exited.
func (s *BaseJavaScriptParserListener) ExitBitNotExpression(ctx *BitNotExpressionContext) {}

// EnterNewExpression is called when production NewExpression is entered.
func (s *BaseJavaScriptParserListener) EnterNewExpression(ctx *NewExpressionContext) {}

// ExitNewExpression is called when production NewExpression is exited.
func (s *BaseJavaScriptParserListener) ExitNewExpression(ctx *NewExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseJavaScriptParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseJavaScriptParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterArrayLiteralExpression is called when production ArrayLiteralExpression is entered.
func (s *BaseJavaScriptParserListener) EnterArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {
}

// ExitArrayLiteralExpression is called when production ArrayLiteralExpression is exited.
func (s *BaseJavaScriptParserListener) ExitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {
}

// EnterMemberDotExpression is called when production MemberDotExpression is entered.
func (s *BaseJavaScriptParserListener) EnterMemberDotExpression(ctx *MemberDotExpressionContext) {}

// ExitMemberDotExpression is called when production MemberDotExpression is exited.
func (s *BaseJavaScriptParserListener) ExitMemberDotExpression(ctx *MemberDotExpressionContext) {}

// EnterClassExpression is called when production ClassExpression is entered.
func (s *BaseJavaScriptParserListener) EnterClassExpression(ctx *ClassExpressionContext) {}

// ExitClassExpression is called when production ClassExpression is exited.
func (s *BaseJavaScriptParserListener) ExitClassExpression(ctx *ClassExpressionContext) {}

// EnterMemberIndexExpression is called when production MemberIndexExpression is entered.
func (s *BaseJavaScriptParserListener) EnterMemberIndexExpression(ctx *MemberIndexExpressionContext) {
}

// ExitMemberIndexExpression is called when production MemberIndexExpression is exited.
func (s *BaseJavaScriptParserListener) ExitMemberIndexExpression(ctx *MemberIndexExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseJavaScriptParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseJavaScriptParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterBitAndExpression is called when production BitAndExpression is entered.
func (s *BaseJavaScriptParserListener) EnterBitAndExpression(ctx *BitAndExpressionContext) {}

// ExitBitAndExpression is called when production BitAndExpression is exited.
func (s *BaseJavaScriptParserListener) ExitBitAndExpression(ctx *BitAndExpressionContext) {}

// EnterBitOrExpression is called when production BitOrExpression is entered.
func (s *BaseJavaScriptParserListener) EnterBitOrExpression(ctx *BitOrExpressionContext) {}

// ExitBitOrExpression is called when production BitOrExpression is exited.
func (s *BaseJavaScriptParserListener) ExitBitOrExpression(ctx *BitOrExpressionContext) {}

// EnterAssignmentOperatorExpression is called when production AssignmentOperatorExpression is entered.
func (s *BaseJavaScriptParserListener) EnterAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// ExitAssignmentOperatorExpression is called when production AssignmentOperatorExpression is exited.
func (s *BaseJavaScriptParserListener) ExitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// EnterVoidExpression is called when production VoidExpression is entered.
func (s *BaseJavaScriptParserListener) EnterVoidExpression(ctx *VoidExpressionContext) {}

// ExitVoidExpression is called when production VoidExpression is exited.
func (s *BaseJavaScriptParserListener) ExitVoidExpression(ctx *VoidExpressionContext) {}

// EnterCoalesceExpression is called when production CoalesceExpression is entered.
func (s *BaseJavaScriptParserListener) EnterCoalesceExpression(ctx *CoalesceExpressionContext) {}

// ExitCoalesceExpression is called when production CoalesceExpression is exited.
func (s *BaseJavaScriptParserListener) ExitCoalesceExpression(ctx *CoalesceExpressionContext) {}

// EnterAssignable is called when production assignable is entered.
func (s *BaseJavaScriptParserListener) EnterAssignable(ctx *AssignableContext) {}

// ExitAssignable is called when production assignable is exited.
func (s *BaseJavaScriptParserListener) ExitAssignable(ctx *AssignableContext) {}

// EnterFunctionDecl is called when production FunctionDecl is entered.
func (s *BaseJavaScriptParserListener) EnterFunctionDecl(ctx *FunctionDeclContext) {}

// ExitFunctionDecl is called when production FunctionDecl is exited.
func (s *BaseJavaScriptParserListener) ExitFunctionDecl(ctx *FunctionDeclContext) {}

// EnterAnoymousFunctionDecl is called when production AnoymousFunctionDecl is entered.
func (s *BaseJavaScriptParserListener) EnterAnoymousFunctionDecl(ctx *AnoymousFunctionDeclContext) {}

// ExitAnoymousFunctionDecl is called when production AnoymousFunctionDecl is exited.
func (s *BaseJavaScriptParserListener) ExitAnoymousFunctionDecl(ctx *AnoymousFunctionDeclContext) {}

// EnterArrowFunction is called when production ArrowFunction is entered.
func (s *BaseJavaScriptParserListener) EnterArrowFunction(ctx *ArrowFunctionContext) {}

// ExitArrowFunction is called when production ArrowFunction is exited.
func (s *BaseJavaScriptParserListener) ExitArrowFunction(ctx *ArrowFunctionContext) {}

// EnterArrowFunctionParameters is called when production arrowFunctionParameters is entered.
func (s *BaseJavaScriptParserListener) EnterArrowFunctionParameters(ctx *ArrowFunctionParametersContext) {
}

// ExitArrowFunctionParameters is called when production arrowFunctionParameters is exited.
func (s *BaseJavaScriptParserListener) ExitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) {
}

// EnterArrowFunctionBody is called when production arrowFunctionBody is entered.
func (s *BaseJavaScriptParserListener) EnterArrowFunctionBody(ctx *ArrowFunctionBodyContext) {}

// ExitArrowFunctionBody is called when production arrowFunctionBody is exited.
func (s *BaseJavaScriptParserListener) ExitArrowFunctionBody(ctx *ArrowFunctionBodyContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseJavaScriptParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseJavaScriptParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseJavaScriptParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseJavaScriptParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseJavaScriptParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseJavaScriptParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterBigintLiteral is called when production bigintLiteral is entered.
func (s *BaseJavaScriptParserListener) EnterBigintLiteral(ctx *BigintLiteralContext) {}

// ExitBigintLiteral is called when production bigintLiteral is exited.
func (s *BaseJavaScriptParserListener) ExitBigintLiteral(ctx *BigintLiteralContext) {}

// EnterIdentifierName is called when production identifierName is entered.
func (s *BaseJavaScriptParserListener) EnterIdentifierName(ctx *IdentifierNameContext) {}

// ExitIdentifierName is called when production identifierName is exited.
func (s *BaseJavaScriptParserListener) ExitIdentifierName(ctx *IdentifierNameContext) {}

// EnterReservedWord is called when production reservedWord is entered.
func (s *BaseJavaScriptParserListener) EnterReservedWord(ctx *ReservedWordContext) {}

// ExitReservedWord is called when production reservedWord is exited.
func (s *BaseJavaScriptParserListener) ExitReservedWord(ctx *ReservedWordContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseJavaScriptParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseJavaScriptParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterGetter is called when production getter is entered.
func (s *BaseJavaScriptParserListener) EnterGetter(ctx *GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *BaseJavaScriptParserListener) ExitGetter(ctx *GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *BaseJavaScriptParserListener) EnterSetter(ctx *SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *BaseJavaScriptParserListener) ExitSetter(ctx *SetterContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseJavaScriptParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseJavaScriptParserListener) ExitEos(ctx *EosContext) {}
