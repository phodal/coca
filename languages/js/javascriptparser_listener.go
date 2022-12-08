// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // JavaScriptParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// JavaScriptParserListener is a complete listener for a parse tree produced by JavaScriptParser.
type JavaScriptParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSourceElement is called when entering the sourceElement production.
	EnterSourceElement(c *SourceElementContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterImportFromBlock is called when entering the importFromBlock production.
	EnterImportFromBlock(c *ImportFromBlockContext)

	// EnterModuleItems is called when entering the moduleItems production.
	EnterModuleItems(c *ModuleItemsContext)

	// EnterImportDefault is called when entering the importDefault production.
	EnterImportDefault(c *ImportDefaultContext)

	// EnterImportNamespace is called when entering the importNamespace production.
	EnterImportNamespace(c *ImportNamespaceContext)

	// EnterImportFrom is called when entering the importFrom production.
	EnterImportFrom(c *ImportFromContext)

	// EnterAliasName is called when entering the aliasName production.
	EnterAliasName(c *AliasNameContext)

	// EnterExportDeclaration is called when entering the ExportDeclaration production.
	EnterExportDeclaration(c *ExportDeclarationContext)

	// EnterExportDefaultDeclaration is called when entering the ExportDefaultDeclaration production.
	EnterExportDefaultDeclaration(c *ExportDefaultDeclarationContext)

	// EnterExportFromBlock is called when entering the exportFromBlock production.
	EnterExportFromBlock(c *ExportFromBlockContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterVariableStatement is called when entering the variableStatement production.
	EnterVariableStatement(c *VariableStatementContext)

	// EnterVariableDeclarationList is called when entering the variableDeclarationList production.
	EnterVariableDeclarationList(c *VariableDeclarationListContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterEmptyStatement_ is called when entering the emptyStatement_ production.
	EnterEmptyStatement_(c *EmptyStatement_Context)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterDoStatement is called when entering the DoStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterWhileStatement is called when entering the WhileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the ForStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForInStatement is called when entering the ForInStatement production.
	EnterForInStatement(c *ForInStatementContext)

	// EnterForOfStatement is called when entering the ForOfStatement production.
	EnterForOfStatement(c *ForOfStatementContext)

	// EnterVarModifier is called when entering the varModifier production.
	EnterVarModifier(c *VarModifierContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterYieldStatement is called when entering the yieldStatement production.
	EnterYieldStatement(c *YieldStatementContext)

	// EnterWithStatement is called when entering the withStatement production.
	EnterWithStatement(c *WithStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterCaseClauses is called when entering the caseClauses production.
	EnterCaseClauses(c *CaseClausesContext)

	// EnterCaseClause is called when entering the caseClause production.
	EnterCaseClause(c *CaseClauseContext)

	// EnterDefaultClause is called when entering the defaultClause production.
	EnterDefaultClause(c *DefaultClauseContext)

	// EnterLabelledStatement is called when entering the labelledStatement production.
	EnterLabelledStatement(c *LabelledStatementContext)

	// EnterThrowStatement is called when entering the throwStatement production.
	EnterThrowStatement(c *ThrowStatementContext)

	// EnterTryStatement is called when entering the tryStatement production.
	EnterTryStatement(c *TryStatementContext)

	// EnterCatchProduction is called when entering the catchProduction production.
	EnterCatchProduction(c *CatchProductionContext)

	// EnterFinallyProduction is called when entering the finallyProduction production.
	EnterFinallyProduction(c *FinallyProductionContext)

	// EnterDebuggerStatement is called when entering the debuggerStatement production.
	EnterDebuggerStatement(c *DebuggerStatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassTail is called when entering the classTail production.
	EnterClassTail(c *ClassTailContext)

	// EnterClassElement is called when entering the classElement production.
	EnterClassElement(c *ClassElementContext)

	// EnterMethodDefinition is called when entering the methodDefinition production.
	EnterMethodDefinition(c *MethodDefinitionContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameterArg is called when entering the formalParameterArg production.
	EnterFormalParameterArg(c *FormalParameterArgContext)

	// EnterLastFormalParameterArg is called when entering the lastFormalParameterArg production.
	EnterLastFormalParameterArg(c *LastFormalParameterArgContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterSourceElements is called when entering the sourceElements production.
	EnterSourceElements(c *SourceElementsContext)

	// EnterArrayLiteral is called when entering the arrayLiteral production.
	EnterArrayLiteral(c *ArrayLiteralContext)

	// EnterElementList is called when entering the elementList production.
	EnterElementList(c *ElementListContext)

	// EnterArrayElement is called when entering the arrayElement production.
	EnterArrayElement(c *ArrayElementContext)

	// EnterObjectLiteral is called when entering the objectLiteral production.
	EnterObjectLiteral(c *ObjectLiteralContext)

	// EnterPropertyExpressionAssignment is called when entering the PropertyExpressionAssignment production.
	EnterPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// EnterComputedPropertyExpressionAssignment is called when entering the ComputedPropertyExpressionAssignment production.
	EnterComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext)

	// EnterFunctionProperty is called when entering the FunctionProperty production.
	EnterFunctionProperty(c *FunctionPropertyContext)

	// EnterPropertyGetter is called when entering the PropertyGetter production.
	EnterPropertyGetter(c *PropertyGetterContext)

	// EnterPropertySetter is called when entering the PropertySetter production.
	EnterPropertySetter(c *PropertySetterContext)

	// EnterPropertyShorthand is called when entering the PropertyShorthand production.
	EnterPropertyShorthand(c *PropertyShorthandContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterExpressionSequence is called when entering the expressionSequence production.
	EnterExpressionSequence(c *ExpressionSequenceContext)

	// EnterTemplateStringExpression is called when entering the TemplateStringExpression production.
	EnterTemplateStringExpression(c *TemplateStringExpressionContext)

	// EnterTernaryExpression is called when entering the TernaryExpression production.
	EnterTernaryExpression(c *TernaryExpressionContext)

	// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterPowerExpression is called when entering the PowerExpression production.
	EnterPowerExpression(c *PowerExpressionContext)

	// EnterPreIncrementExpression is called when entering the PreIncrementExpression production.
	EnterPreIncrementExpression(c *PreIncrementExpressionContext)

	// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
	EnterObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// EnterMetaExpression is called when entering the MetaExpression production.
	EnterMetaExpression(c *MetaExpressionContext)

	// EnterInExpression is called when entering the InExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterNotExpression is called when entering the NotExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterPreDecreaseExpression is called when entering the PreDecreaseExpression production.
	EnterPreDecreaseExpression(c *PreDecreaseExpressionContext)

	// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
	EnterArgumentsExpression(c *ArgumentsExpressionContext)

	// EnterAwaitExpression is called when entering the AwaitExpression production.
	EnterAwaitExpression(c *AwaitExpressionContext)

	// EnterThisExpression is called when entering the ThisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterFunctionExpression is called when entering the FunctionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
	EnterUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// EnterAssignmentExpression is called when entering the AssignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterPostDecreaseExpression is called when entering the PostDecreaseExpression production.
	EnterPostDecreaseExpression(c *PostDecreaseExpressionContext)

	// EnterTypeofExpression is called when entering the TypeofExpression production.
	EnterTypeofExpression(c *TypeofExpressionContext)

	// EnterInstanceofExpression is called when entering the InstanceofExpression production.
	EnterInstanceofExpression(c *InstanceofExpressionContext)

	// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
	EnterUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// EnterDeleteExpression is called when entering the DeleteExpression production.
	EnterDeleteExpression(c *DeleteExpressionContext)

	// EnterImportExpression is called when entering the ImportExpression production.
	EnterImportExpression(c *ImportExpressionContext)

	// EnterEqualityExpression is called when entering the EqualityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterBitXOrExpression is called when entering the BitXOrExpression production.
	EnterBitXOrExpression(c *BitXOrExpressionContext)

	// EnterSuperExpression is called when entering the SuperExpression production.
	EnterSuperExpression(c *SuperExpressionContext)

	// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterBitShiftExpression is called when entering the BitShiftExpression production.
	EnterBitShiftExpression(c *BitShiftExpressionContext)

	// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterAdditiveExpression is called when entering the AdditiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterRelationalExpression is called when entering the RelationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterPostIncrementExpression is called when entering the PostIncrementExpression production.
	EnterPostIncrementExpression(c *PostIncrementExpressionContext)

	// EnterYieldExpression is called when entering the YieldExpression production.
	EnterYieldExpression(c *YieldExpressionContext)

	// EnterBitNotExpression is called when entering the BitNotExpression production.
	EnterBitNotExpression(c *BitNotExpressionContext)

	// EnterNewExpression is called when entering the NewExpression production.
	EnterNewExpression(c *NewExpressionContext)

	// EnterLiteralExpression is called when entering the LiteralExpression production.
	EnterLiteralExpression(c *LiteralExpressionContext)

	// EnterArrayLiteralExpression is called when entering the ArrayLiteralExpression production.
	EnterArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// EnterMemberDotExpression is called when entering the MemberDotExpression production.
	EnterMemberDotExpression(c *MemberDotExpressionContext)

	// EnterClassExpression is called when entering the ClassExpression production.
	EnterClassExpression(c *ClassExpressionContext)

	// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
	EnterMemberIndexExpression(c *MemberIndexExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterBitAndExpression is called when entering the BitAndExpression production.
	EnterBitAndExpression(c *BitAndExpressionContext)

	// EnterBitOrExpression is called when entering the BitOrExpression production.
	EnterBitOrExpression(c *BitOrExpressionContext)

	// EnterAssignmentOperatorExpression is called when entering the AssignmentOperatorExpression production.
	EnterAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext)

	// EnterVoidExpression is called when entering the VoidExpression production.
	EnterVoidExpression(c *VoidExpressionContext)

	// EnterCoalesceExpression is called when entering the CoalesceExpression production.
	EnterCoalesceExpression(c *CoalesceExpressionContext)

	// EnterAssignable is called when entering the assignable production.
	EnterAssignable(c *AssignableContext)

	// EnterFunctionDecl is called when entering the FunctionDecl production.
	EnterFunctionDecl(c *FunctionDeclContext)

	// EnterAnoymousFunctionDecl is called when entering the AnoymousFunctionDecl production.
	EnterAnoymousFunctionDecl(c *AnoymousFunctionDeclContext)

	// EnterArrowFunction is called when entering the ArrowFunction production.
	EnterArrowFunction(c *ArrowFunctionContext)

	// EnterArrowFunctionParameters is called when entering the arrowFunctionParameters production.
	EnterArrowFunctionParameters(c *ArrowFunctionParametersContext)

	// EnterArrowFunctionBody is called when entering the arrowFunctionBody production.
	EnterArrowFunctionBody(c *ArrowFunctionBodyContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterNumericLiteral is called when entering the numericLiteral production.
	EnterNumericLiteral(c *NumericLiteralContext)

	// EnterBigintLiteral is called when entering the bigintLiteral production.
	EnterBigintLiteral(c *BigintLiteralContext)

	// EnterIdentifierName is called when entering the identifierName production.
	EnterIdentifierName(c *IdentifierNameContext)

	// EnterReservedWord is called when entering the reservedWord production.
	EnterReservedWord(c *ReservedWordContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterGetter is called when entering the getter production.
	EnterGetter(c *GetterContext)

	// EnterSetter is called when entering the setter production.
	EnterSetter(c *SetterContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSourceElement is called when exiting the sourceElement production.
	ExitSourceElement(c *SourceElementContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitImportFromBlock is called when exiting the importFromBlock production.
	ExitImportFromBlock(c *ImportFromBlockContext)

	// ExitModuleItems is called when exiting the moduleItems production.
	ExitModuleItems(c *ModuleItemsContext)

	// ExitImportDefault is called when exiting the importDefault production.
	ExitImportDefault(c *ImportDefaultContext)

	// ExitImportNamespace is called when exiting the importNamespace production.
	ExitImportNamespace(c *ImportNamespaceContext)

	// ExitImportFrom is called when exiting the importFrom production.
	ExitImportFrom(c *ImportFromContext)

	// ExitAliasName is called when exiting the aliasName production.
	ExitAliasName(c *AliasNameContext)

	// ExitExportDeclaration is called when exiting the ExportDeclaration production.
	ExitExportDeclaration(c *ExportDeclarationContext)

	// ExitExportDefaultDeclaration is called when exiting the ExportDefaultDeclaration production.
	ExitExportDefaultDeclaration(c *ExportDefaultDeclarationContext)

	// ExitExportFromBlock is called when exiting the exportFromBlock production.
	ExitExportFromBlock(c *ExportFromBlockContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitVariableStatement is called when exiting the variableStatement production.
	ExitVariableStatement(c *VariableStatementContext)

	// ExitVariableDeclarationList is called when exiting the variableDeclarationList production.
	ExitVariableDeclarationList(c *VariableDeclarationListContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitEmptyStatement_ is called when exiting the emptyStatement_ production.
	ExitEmptyStatement_(c *EmptyStatement_Context)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitDoStatement is called when exiting the DoStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitWhileStatement is called when exiting the WhileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the ForStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForInStatement is called when exiting the ForInStatement production.
	ExitForInStatement(c *ForInStatementContext)

	// ExitForOfStatement is called when exiting the ForOfStatement production.
	ExitForOfStatement(c *ForOfStatementContext)

	// ExitVarModifier is called when exiting the varModifier production.
	ExitVarModifier(c *VarModifierContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitYieldStatement is called when exiting the yieldStatement production.
	ExitYieldStatement(c *YieldStatementContext)

	// ExitWithStatement is called when exiting the withStatement production.
	ExitWithStatement(c *WithStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitCaseClauses is called when exiting the caseClauses production.
	ExitCaseClauses(c *CaseClausesContext)

	// ExitCaseClause is called when exiting the caseClause production.
	ExitCaseClause(c *CaseClauseContext)

	// ExitDefaultClause is called when exiting the defaultClause production.
	ExitDefaultClause(c *DefaultClauseContext)

	// ExitLabelledStatement is called when exiting the labelledStatement production.
	ExitLabelledStatement(c *LabelledStatementContext)

	// ExitThrowStatement is called when exiting the throwStatement production.
	ExitThrowStatement(c *ThrowStatementContext)

	// ExitTryStatement is called when exiting the tryStatement production.
	ExitTryStatement(c *TryStatementContext)

	// ExitCatchProduction is called when exiting the catchProduction production.
	ExitCatchProduction(c *CatchProductionContext)

	// ExitFinallyProduction is called when exiting the finallyProduction production.
	ExitFinallyProduction(c *FinallyProductionContext)

	// ExitDebuggerStatement is called when exiting the debuggerStatement production.
	ExitDebuggerStatement(c *DebuggerStatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassTail is called when exiting the classTail production.
	ExitClassTail(c *ClassTailContext)

	// ExitClassElement is called when exiting the classElement production.
	ExitClassElement(c *ClassElementContext)

	// ExitMethodDefinition is called when exiting the methodDefinition production.
	ExitMethodDefinition(c *MethodDefinitionContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameterArg is called when exiting the formalParameterArg production.
	ExitFormalParameterArg(c *FormalParameterArgContext)

	// ExitLastFormalParameterArg is called when exiting the lastFormalParameterArg production.
	ExitLastFormalParameterArg(c *LastFormalParameterArgContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitSourceElements is called when exiting the sourceElements production.
	ExitSourceElements(c *SourceElementsContext)

	// ExitArrayLiteral is called when exiting the arrayLiteral production.
	ExitArrayLiteral(c *ArrayLiteralContext)

	// ExitElementList is called when exiting the elementList production.
	ExitElementList(c *ElementListContext)

	// ExitArrayElement is called when exiting the arrayElement production.
	ExitArrayElement(c *ArrayElementContext)

	// ExitObjectLiteral is called when exiting the objectLiteral production.
	ExitObjectLiteral(c *ObjectLiteralContext)

	// ExitPropertyExpressionAssignment is called when exiting the PropertyExpressionAssignment production.
	ExitPropertyExpressionAssignment(c *PropertyExpressionAssignmentContext)

	// ExitComputedPropertyExpressionAssignment is called when exiting the ComputedPropertyExpressionAssignment production.
	ExitComputedPropertyExpressionAssignment(c *ComputedPropertyExpressionAssignmentContext)

	// ExitFunctionProperty is called when exiting the FunctionProperty production.
	ExitFunctionProperty(c *FunctionPropertyContext)

	// ExitPropertyGetter is called when exiting the PropertyGetter production.
	ExitPropertyGetter(c *PropertyGetterContext)

	// ExitPropertySetter is called when exiting the PropertySetter production.
	ExitPropertySetter(c *PropertySetterContext)

	// ExitPropertyShorthand is called when exiting the PropertyShorthand production.
	ExitPropertyShorthand(c *PropertyShorthandContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitExpressionSequence is called when exiting the expressionSequence production.
	ExitExpressionSequence(c *ExpressionSequenceContext)

	// ExitTemplateStringExpression is called when exiting the TemplateStringExpression production.
	ExitTemplateStringExpression(c *TemplateStringExpressionContext)

	// ExitTernaryExpression is called when exiting the TernaryExpression production.
	ExitTernaryExpression(c *TernaryExpressionContext)

	// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitPowerExpression is called when exiting the PowerExpression production.
	ExitPowerExpression(c *PowerExpressionContext)

	// ExitPreIncrementExpression is called when exiting the PreIncrementExpression production.
	ExitPreIncrementExpression(c *PreIncrementExpressionContext)

	// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
	ExitObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// ExitMetaExpression is called when exiting the MetaExpression production.
	ExitMetaExpression(c *MetaExpressionContext)

	// ExitInExpression is called when exiting the InExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitNotExpression is called when exiting the NotExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitPreDecreaseExpression is called when exiting the PreDecreaseExpression production.
	ExitPreDecreaseExpression(c *PreDecreaseExpressionContext)

	// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
	ExitArgumentsExpression(c *ArgumentsExpressionContext)

	// ExitAwaitExpression is called when exiting the AwaitExpression production.
	ExitAwaitExpression(c *AwaitExpressionContext)

	// ExitThisExpression is called when exiting the ThisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitFunctionExpression is called when exiting the FunctionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
	ExitUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitPostDecreaseExpression is called when exiting the PostDecreaseExpression production.
	ExitPostDecreaseExpression(c *PostDecreaseExpressionContext)

	// ExitTypeofExpression is called when exiting the TypeofExpression production.
	ExitTypeofExpression(c *TypeofExpressionContext)

	// ExitInstanceofExpression is called when exiting the InstanceofExpression production.
	ExitInstanceofExpression(c *InstanceofExpressionContext)

	// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
	ExitUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// ExitDeleteExpression is called when exiting the DeleteExpression production.
	ExitDeleteExpression(c *DeleteExpressionContext)

	// ExitImportExpression is called when exiting the ImportExpression production.
	ExitImportExpression(c *ImportExpressionContext)

	// ExitEqualityExpression is called when exiting the EqualityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitBitXOrExpression is called when exiting the BitXOrExpression production.
	ExitBitXOrExpression(c *BitXOrExpressionContext)

	// ExitSuperExpression is called when exiting the SuperExpression production.
	ExitSuperExpression(c *SuperExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitBitShiftExpression is called when exiting the BitShiftExpression production.
	ExitBitShiftExpression(c *BitShiftExpressionContext)

	// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitRelationalExpression is called when exiting the RelationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitPostIncrementExpression is called when exiting the PostIncrementExpression production.
	ExitPostIncrementExpression(c *PostIncrementExpressionContext)

	// ExitYieldExpression is called when exiting the YieldExpression production.
	ExitYieldExpression(c *YieldExpressionContext)

	// ExitBitNotExpression is called when exiting the BitNotExpression production.
	ExitBitNotExpression(c *BitNotExpressionContext)

	// ExitNewExpression is called when exiting the NewExpression production.
	ExitNewExpression(c *NewExpressionContext)

	// ExitLiteralExpression is called when exiting the LiteralExpression production.
	ExitLiteralExpression(c *LiteralExpressionContext)

	// ExitArrayLiteralExpression is called when exiting the ArrayLiteralExpression production.
	ExitArrayLiteralExpression(c *ArrayLiteralExpressionContext)

	// ExitMemberDotExpression is called when exiting the MemberDotExpression production.
	ExitMemberDotExpression(c *MemberDotExpressionContext)

	// ExitClassExpression is called when exiting the ClassExpression production.
	ExitClassExpression(c *ClassExpressionContext)

	// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
	ExitMemberIndexExpression(c *MemberIndexExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitBitAndExpression is called when exiting the BitAndExpression production.
	ExitBitAndExpression(c *BitAndExpressionContext)

	// ExitBitOrExpression is called when exiting the BitOrExpression production.
	ExitBitOrExpression(c *BitOrExpressionContext)

	// ExitAssignmentOperatorExpression is called when exiting the AssignmentOperatorExpression production.
	ExitAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext)

	// ExitVoidExpression is called when exiting the VoidExpression production.
	ExitVoidExpression(c *VoidExpressionContext)

	// ExitCoalesceExpression is called when exiting the CoalesceExpression production.
	ExitCoalesceExpression(c *CoalesceExpressionContext)

	// ExitAssignable is called when exiting the assignable production.
	ExitAssignable(c *AssignableContext)

	// ExitFunctionDecl is called when exiting the FunctionDecl production.
	ExitFunctionDecl(c *FunctionDeclContext)

	// ExitAnoymousFunctionDecl is called when exiting the AnoymousFunctionDecl production.
	ExitAnoymousFunctionDecl(c *AnoymousFunctionDeclContext)

	// ExitArrowFunction is called when exiting the ArrowFunction production.
	ExitArrowFunction(c *ArrowFunctionContext)

	// ExitArrowFunctionParameters is called when exiting the arrowFunctionParameters production.
	ExitArrowFunctionParameters(c *ArrowFunctionParametersContext)

	// ExitArrowFunctionBody is called when exiting the arrowFunctionBody production.
	ExitArrowFunctionBody(c *ArrowFunctionBodyContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitNumericLiteral is called when exiting the numericLiteral production.
	ExitNumericLiteral(c *NumericLiteralContext)

	// ExitBigintLiteral is called when exiting the bigintLiteral production.
	ExitBigintLiteral(c *BigintLiteralContext)

	// ExitIdentifierName is called when exiting the identifierName production.
	ExitIdentifierName(c *IdentifierNameContext)

	// ExitReservedWord is called when exiting the reservedWord production.
	ExitReservedWord(c *ReservedWordContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitGetter is called when exiting the getter production.
	ExitGetter(c *GetterContext)

	// ExitSetter is called when exiting the setter production.
	ExitSetter(c *SetterContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
