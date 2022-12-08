// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GroovyParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// GroovyParserListener is a complete listener for a parse tree produced by GroovyParser.
type GroovyParserListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterScriptStatements is called when entering the scriptStatements production.
	EnterScriptStatements(c *ScriptStatementsContext)

	// EnterScriptStatement is called when entering the scriptStatement production.
	EnterScriptStatement(c *ScriptStatementContext)

	// EnterPackageDeclaration is called when entering the packageDeclaration production.
	EnterPackageDeclaration(c *PackageDeclarationContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterModifiersOpt is called when entering the modifiersOpt production.
	EnterModifiersOpt(c *ModifiersOptContext)

	// EnterModifiers is called when entering the modifiers production.
	EnterModifiers(c *ModifiersContext)

	// EnterClassOrInterfaceModifiersOpt is called when entering the classOrInterfaceModifiersOpt production.
	EnterClassOrInterfaceModifiersOpt(c *ClassOrInterfaceModifiersOptContext)

	// EnterClassOrInterfaceModifiers is called when entering the classOrInterfaceModifiers production.
	EnterClassOrInterfaceModifiers(c *ClassOrInterfaceModifiersContext)

	// EnterClassOrInterfaceModifier is called when entering the classOrInterfaceModifier production.
	EnterClassOrInterfaceModifier(c *ClassOrInterfaceModifierContext)

	// EnterVariableModifier is called when entering the variableModifier production.
	EnterVariableModifier(c *VariableModifierContext)

	// EnterVariableModifiersOpt is called when entering the variableModifiersOpt production.
	EnterVariableModifiersOpt(c *VariableModifiersOptContext)

	// EnterVariableModifiers is called when entering the variableModifiers production.
	EnterVariableModifiers(c *VariableModifiersContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterTypeBound is called when entering the typeBound production.
	EnterTypeBound(c *TypeBoundContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterEnumConstants is called when entering the enumConstants production.
	EnterEnumConstants(c *EnumConstantsContext)

	// EnterEnumConstant is called when entering the enumConstant production.
	EnterEnumConstant(c *EnumConstantContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterMemberDeclaration is called when entering the memberDeclaration production.
	EnterMemberDeclaration(c *MemberDeclarationContext)

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterMethodName is called when entering the methodName production.
	EnterMethodName(c *MethodNameContext)

	// EnterReturnType is called when entering the returnType production.
	EnterReturnType(c *ReturnTypeContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableDeclaratorId is called when entering the variableDeclaratorId production.
	EnterVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterVariableInitializers is called when entering the variableInitializers production.
	EnterVariableInitializers(c *VariableInitializersContext)

	// EnterDims is called when entering the dims production.
	EnterDims(c *DimsContext)

	// EnterDimsOpt is called when entering the dimsOpt production.
	EnterDimsOpt(c *DimsOptContext)

	// EnterStandardType is called when entering the standardType production.
	EnterStandardType(c *StandardTypeContext)

	// EnterTypeType is called when entering the typeType production.
	EnterTypeType(c *TypeTypeContext)

	// EnterClassOrInterfaceType is called when entering the classOrInterfaceType production.
	EnterClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// EnterGeneralClassOrInterfaceType is called when entering the generalClassOrInterfaceType production.
	EnterGeneralClassOrInterfaceType(c *GeneralClassOrInterfaceTypeContext)

	// EnterStandardClassOrInterfaceType is called when entering the standardClassOrInterfaceType production.
	EnterStandardClassOrInterfaceType(c *StandardClassOrInterfaceTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeArgument is called when entering the typeArgument production.
	EnterTypeArgument(c *TypeArgumentContext)

	// EnterAnnotatedQualifiedClassName is called when entering the annotatedQualifiedClassName production.
	EnterAnnotatedQualifiedClassName(c *AnnotatedQualifiedClassNameContext)

	// EnterQualifiedClassNameList is called when entering the qualifiedClassNameList production.
	EnterQualifiedClassNameList(c *QualifiedClassNameListContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterThisFormalParameter is called when entering the thisFormalParameter production.
	EnterThisFormalParameter(c *ThisFormalParameterContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterQualifiedNameElement is called when entering the qualifiedNameElement production.
	EnterQualifiedNameElement(c *QualifiedNameElementContext)

	// EnterQualifiedNameElements is called when entering the qualifiedNameElements production.
	EnterQualifiedNameElements(c *QualifiedNameElementsContext)

	// EnterQualifiedClassName is called when entering the qualifiedClassName production.
	EnterQualifiedClassName(c *QualifiedClassNameContext)

	// EnterQualifiedStandardClassName is called when entering the qualifiedStandardClassName production.
	EnterQualifiedStandardClassName(c *QualifiedStandardClassNameContext)

	// EnterIntegerLiteralAlt is called when entering the integerLiteralAlt production.
	EnterIntegerLiteralAlt(c *IntegerLiteralAltContext)

	// EnterFloatingPointLiteralAlt is called when entering the floatingPointLiteralAlt production.
	EnterFloatingPointLiteralAlt(c *FloatingPointLiteralAltContext)

	// EnterStringLiteralAlt is called when entering the stringLiteralAlt production.
	EnterStringLiteralAlt(c *StringLiteralAltContext)

	// EnterBoolLiteralAlt is called when entering the boolLiteralAlt production.
	EnterBoolLiteralAlt(c *BoolLiteralAltContext)

	// EnterNilLiteralAlt is called when entering the nilLiteralAlt production.
	EnterNilLiteralAlt(c *NilLiteralAltContext)

	// EnterGstring is called when entering the gstring production.
	EnterGstring(c *GstringContext)

	// EnterGstringValue is called when entering the gstringValue production.
	EnterGstringValue(c *GstringValueContext)

	// EnterGstringPath is called when entering the gstringPath production.
	EnterGstringPath(c *GstringPathContext)

	// EnterLambdaExpression is called when entering the lambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterStandardLambdaExpression is called when entering the standardLambdaExpression production.
	EnterStandardLambdaExpression(c *StandardLambdaExpressionContext)

	// EnterLambdaParameters is called when entering the lambdaParameters production.
	EnterLambdaParameters(c *LambdaParametersContext)

	// EnterStandardLambdaParameters is called when entering the standardLambdaParameters production.
	EnterStandardLambdaParameters(c *StandardLambdaParametersContext)

	// EnterLambdaBody is called when entering the lambdaBody production.
	EnterLambdaBody(c *LambdaBodyContext)

	// EnterClosure is called when entering the closure production.
	EnterClosure(c *ClosureContext)

	// EnterClosureOrLambdaExpression is called when entering the closureOrLambdaExpression production.
	EnterClosureOrLambdaExpression(c *ClosureOrLambdaExpressionContext)

	// EnterBlockStatementsOpt is called when entering the blockStatementsOpt production.
	EnterBlockStatementsOpt(c *BlockStatementsOptContext)

	// EnterBlockStatements is called when entering the blockStatements production.
	EnterBlockStatements(c *BlockStatementsContext)

	// EnterAnnotationsOpt is called when entering the annotationsOpt production.
	EnterAnnotationsOpt(c *AnnotationsOptContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterElementValues is called when entering the elementValues production.
	EnterElementValues(c *ElementValuesContext)

	// EnterAnnotationName is called when entering the annotationName production.
	EnterAnnotationName(c *AnnotationNameContext)

	// EnterElementValuePairs is called when entering the elementValuePairs production.
	EnterElementValuePairs(c *ElementValuePairsContext)

	// EnterElementValuePair is called when entering the elementValuePair production.
	EnterElementValuePair(c *ElementValuePairContext)

	// EnterElementValuePairName is called when entering the elementValuePairName production.
	EnterElementValuePairName(c *ElementValuePairNameContext)

	// EnterElementValue is called when entering the elementValue production.
	EnterElementValue(c *ElementValueContext)

	// EnterElementValueArrayInitializer is called when entering the elementValueArrayInitializer production.
	EnterElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterClassifiedModifiers is called when entering the classifiedModifiers production.
	EnterClassifiedModifiers(c *ClassifiedModifiersContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterTypeNamePairs is called when entering the typeNamePairs production.
	EnterTypeNamePairs(c *TypeNamePairsContext)

	// EnterTypeNamePair is called when entering the typeNamePair production.
	EnterTypeNamePair(c *TypeNamePairContext)

	// EnterVariableNames is called when entering the variableNames production.
	EnterVariableNames(c *VariableNamesContext)

	// EnterConditionalStatement is called when entering the conditionalStatement production.
	EnterConditionalStatement(c *ConditionalStatementContext)

	// EnterIfElseStatement is called when entering the ifElseStatement production.
	EnterIfElseStatement(c *IfElseStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterForStmtAlt is called when entering the forStmtAlt production.
	EnterForStmtAlt(c *ForStmtAltContext)

	// EnterWhileStmtAlt is called when entering the whileStmtAlt production.
	EnterWhileStmtAlt(c *WhileStmtAltContext)

	// EnterDoWhileStmtAlt is called when entering the doWhileStmtAlt production.
	EnterDoWhileStmtAlt(c *DoWhileStmtAltContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterTryCatchStatement is called when entering the tryCatchStatement production.
	EnterTryCatchStatement(c *TryCatchStatementContext)

	// EnterAssertStatement is called when entering the assertStatement production.
	EnterAssertStatement(c *AssertStatementContext)

	// EnterBlockStmtAlt is called when entering the blockStmtAlt production.
	EnterBlockStmtAlt(c *BlockStmtAltContext)

	// EnterConditionalStmtAlt is called when entering the conditionalStmtAlt production.
	EnterConditionalStmtAlt(c *ConditionalStmtAltContext)

	// EnterLoopStmtAlt is called when entering the loopStmtAlt production.
	EnterLoopStmtAlt(c *LoopStmtAltContext)

	// EnterTryCatchStmtAlt is called when entering the tryCatchStmtAlt production.
	EnterTryCatchStmtAlt(c *TryCatchStmtAltContext)

	// EnterSynchronizedStmtAlt is called when entering the synchronizedStmtAlt production.
	EnterSynchronizedStmtAlt(c *SynchronizedStmtAltContext)

	// EnterReturnStmtAlt is called when entering the returnStmtAlt production.
	EnterReturnStmtAlt(c *ReturnStmtAltContext)

	// EnterThrowStmtAlt is called when entering the throwStmtAlt production.
	EnterThrowStmtAlt(c *ThrowStmtAltContext)

	// EnterBreakStmtAlt is called when entering the breakStmtAlt production.
	EnterBreakStmtAlt(c *BreakStmtAltContext)

	// EnterContinueStmtAlt is called when entering the continueStmtAlt production.
	EnterContinueStmtAlt(c *ContinueStmtAltContext)

	// EnterLabeledStmtAlt is called when entering the labeledStmtAlt production.
	EnterLabeledStmtAlt(c *LabeledStmtAltContext)

	// EnterAssertStmtAlt is called when entering the assertStmtAlt production.
	EnterAssertStmtAlt(c *AssertStmtAltContext)

	// EnterLocalVariableDeclarationStmtAlt is called when entering the localVariableDeclarationStmtAlt production.
	EnterLocalVariableDeclarationStmtAlt(c *LocalVariableDeclarationStmtAltContext)

	// EnterMethodDeclarationStmtAlt is called when entering the methodDeclarationStmtAlt production.
	EnterMethodDeclarationStmtAlt(c *MethodDeclarationStmtAltContext)

	// EnterExpressionStmtAlt is called when entering the expressionStmtAlt production.
	EnterExpressionStmtAlt(c *ExpressionStmtAltContext)

	// EnterEmptyStmtAlt is called when entering the emptyStmtAlt production.
	EnterEmptyStmtAlt(c *EmptyStmtAltContext)

	// EnterCatchClause is called when entering the catchClause production.
	EnterCatchClause(c *CatchClauseContext)

	// EnterCatchType is called when entering the catchType production.
	EnterCatchType(c *CatchTypeContext)

	// EnterFinallyBlock is called when entering the finallyBlock production.
	EnterFinallyBlock(c *FinallyBlockContext)

	// EnterResources is called when entering the resources production.
	EnterResources(c *ResourcesContext)

	// EnterResourceList is called when entering the resourceList production.
	EnterResourceList(c *ResourceListContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterSwitchBlockStatementGroup is called when entering the switchBlockStatementGroup production.
	EnterSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// EnterSwitchLabel is called when entering the switchLabel production.
	EnterSwitchLabel(c *SwitchLabelContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterEnhancedForControl is called when entering the enhancedForControl production.
	EnterEnhancedForControl(c *EnhancedForControlContext)

	// EnterClassicalForControl is called when entering the classicalForControl production.
	EnterClassicalForControl(c *ClassicalForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterForUpdate is called when entering the forUpdate production.
	EnterForUpdate(c *ForUpdateContext)

	// EnterCastParExpression is called when entering the castParExpression production.
	EnterCastParExpression(c *CastParExpressionContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterExpressionInPar is called when entering the expressionInPar production.
	EnterExpressionInPar(c *ExpressionInParContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterExpressionListElement is called when entering the expressionListElement production.
	EnterExpressionListElement(c *ExpressionListElementContext)

	// EnterEnhancedStatementExpression is called when entering the enhancedStatementExpression production.
	EnterEnhancedStatementExpression(c *EnhancedStatementExpressionContext)

	// EnterCommandExprAlt is called when entering the commandExprAlt production.
	EnterCommandExprAlt(c *CommandExprAltContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterCastExprAltForExpr is called when entering the castExprAltForExpr production.
	EnterCastExprAltForExpr(c *CastExprAltForExprContext)

	// EnterShiftExprAlt is called when entering the shiftExprAlt production.
	EnterShiftExprAlt(c *ShiftExprAltContext)

	// EnterMultipleAssignmentExprAlt is called when entering the multipleAssignmentExprAlt production.
	EnterMultipleAssignmentExprAlt(c *MultipleAssignmentExprAltContext)

	// EnterExclusiveOrExprAlt is called when entering the exclusiveOrExprAlt production.
	EnterExclusiveOrExprAlt(c *ExclusiveOrExprAltContext)

	// EnterAdditiveExprAlt is called when entering the additiveExprAlt production.
	EnterAdditiveExprAlt(c *AdditiveExprAltContext)

	// EnterRegexExprAlt is called when entering the regexExprAlt production.
	EnterRegexExprAlt(c *RegexExprAltContext)

	// EnterUnaryNotExprAltForExpr is called when entering the unaryNotExprAltForExpr production.
	EnterUnaryNotExprAltForExpr(c *UnaryNotExprAltForExprContext)

	// EnterConditionalExprAlt is called when entering the conditionalExprAlt production.
	EnterConditionalExprAlt(c *ConditionalExprAltContext)

	// EnterPowerExprAlt is called when entering the powerExprAlt production.
	EnterPowerExprAlt(c *PowerExprAltContext)

	// EnterRelationalExprAlt is called when entering the relationalExprAlt production.
	EnterRelationalExprAlt(c *RelationalExprAltContext)

	// EnterLogicalAndExprAlt is called when entering the logicalAndExprAlt production.
	EnterLogicalAndExprAlt(c *LogicalAndExprAltContext)

	// EnterAssignmentExprAlt is called when entering the assignmentExprAlt production.
	EnterAssignmentExprAlt(c *AssignmentExprAltContext)

	// EnterPostfixExprAltForExpr is called when entering the postfixExprAltForExpr production.
	EnterPostfixExprAltForExpr(c *PostfixExprAltForExprContext)

	// EnterMultiplicativeExprAlt is called when entering the multiplicativeExprAlt production.
	EnterMultiplicativeExprAlt(c *MultiplicativeExprAltContext)

	// EnterInclusiveOrExprAlt is called when entering the inclusiveOrExprAlt production.
	EnterInclusiveOrExprAlt(c *InclusiveOrExprAltContext)

	// EnterLogicalOrExprAlt is called when entering the logicalOrExprAlt production.
	EnterLogicalOrExprAlt(c *LogicalOrExprAltContext)

	// EnterEqualityExprAlt is called when entering the equalityExprAlt production.
	EnterEqualityExprAlt(c *EqualityExprAltContext)

	// EnterAndExprAlt is called when entering the andExprAlt production.
	EnterAndExprAlt(c *AndExprAltContext)

	// EnterUnaryAddExprAltForExpr is called when entering the unaryAddExprAltForExpr production.
	EnterUnaryAddExprAltForExpr(c *UnaryAddExprAltForExprContext)

	// EnterCastExprAlt is called when entering the castExprAlt production.
	EnterCastExprAlt(c *CastExprAltContext)

	// EnterPostfixExprAlt is called when entering the postfixExprAlt production.
	EnterPostfixExprAlt(c *PostfixExprAltContext)

	// EnterUnaryNotExprAlt is called when entering the unaryNotExprAlt production.
	EnterUnaryNotExprAlt(c *UnaryNotExprAltContext)

	// EnterUnaryAddExprAlt is called when entering the unaryAddExprAlt production.
	EnterUnaryAddExprAlt(c *UnaryAddExprAltContext)

	// EnterCommandExpression is called when entering the commandExpression production.
	EnterCommandExpression(c *CommandExpressionContext)

	// EnterCommandArgument is called when entering the commandArgument production.
	EnterCommandArgument(c *CommandArgumentContext)

	// EnterPathExpression is called when entering the pathExpression production.
	EnterPathExpression(c *PathExpressionContext)

	// EnterPathElement is called when entering the pathElement production.
	EnterPathElement(c *PathElementContext)

	// EnterNamePart is called when entering the namePart production.
	EnterNamePart(c *NamePartContext)

	// EnterDynamicMemberName is called when entering the dynamicMemberName production.
	EnterDynamicMemberName(c *DynamicMemberNameContext)

	// EnterIndexPropertyArgs is called when entering the indexPropertyArgs production.
	EnterIndexPropertyArgs(c *IndexPropertyArgsContext)

	// EnterNamedPropertyArgs is called when entering the namedPropertyArgs production.
	EnterNamedPropertyArgs(c *NamedPropertyArgsContext)

	// EnterIdentifierPrmrAlt is called when entering the identifierPrmrAlt production.
	EnterIdentifierPrmrAlt(c *IdentifierPrmrAltContext)

	// EnterLiteralPrmrAlt is called when entering the literalPrmrAlt production.
	EnterLiteralPrmrAlt(c *LiteralPrmrAltContext)

	// EnterGstringPrmrAlt is called when entering the gstringPrmrAlt production.
	EnterGstringPrmrAlt(c *GstringPrmrAltContext)

	// EnterNewPrmrAlt is called when entering the newPrmrAlt production.
	EnterNewPrmrAlt(c *NewPrmrAltContext)

	// EnterThisPrmrAlt is called when entering the thisPrmrAlt production.
	EnterThisPrmrAlt(c *ThisPrmrAltContext)

	// EnterSuperPrmrAlt is called when entering the superPrmrAlt production.
	EnterSuperPrmrAlt(c *SuperPrmrAltContext)

	// EnterParenPrmrAlt is called when entering the parenPrmrAlt production.
	EnterParenPrmrAlt(c *ParenPrmrAltContext)

	// EnterClosureOrLambdaExpressionPrmrAlt is called when entering the closureOrLambdaExpressionPrmrAlt production.
	EnterClosureOrLambdaExpressionPrmrAlt(c *ClosureOrLambdaExpressionPrmrAltContext)

	// EnterListPrmrAlt is called when entering the listPrmrAlt production.
	EnterListPrmrAlt(c *ListPrmrAltContext)

	// EnterMapPrmrAlt is called when entering the mapPrmrAlt production.
	EnterMapPrmrAlt(c *MapPrmrAltContext)

	// EnterBuiltInTypePrmrAlt is called when entering the builtInTypePrmrAlt production.
	EnterBuiltInTypePrmrAlt(c *BuiltInTypePrmrAltContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterNormalMap is called when entering the normalMap production.
	EnterNormalMap(c *NormalMapContext)

	// EnterMapEntryList is called when entering the mapEntryList production.
	EnterMapEntryList(c *MapEntryListContext)

	// EnterMapEntry is called when entering the mapEntry production.
	EnterMapEntry(c *MapEntryContext)

	// EnterMapEntryLabel is called when entering the mapEntryLabel production.
	EnterMapEntryLabel(c *MapEntryLabelContext)

	// EnterCreator is called when entering the creator production.
	EnterCreator(c *CreatorContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterAnonymousInnerClassDeclaration is called when entering the anonymousInnerClassDeclaration production.
	EnterAnonymousInnerClassDeclaration(c *AnonymousInnerClassDeclarationContext)

	// EnterCreatedName is called when entering the createdName production.
	EnterCreatedName(c *CreatedNameContext)

	// EnterNonWildcardTypeArguments is called when entering the nonWildcardTypeArguments production.
	EnterNonWildcardTypeArguments(c *NonWildcardTypeArgumentsContext)

	// EnterTypeArgumentsOrDiamond is called when entering the typeArgumentsOrDiamond production.
	EnterTypeArgumentsOrDiamond(c *TypeArgumentsOrDiamondContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterEnhancedArgumentList is called when entering the enhancedArgumentList production.
	EnterEnhancedArgumentList(c *EnhancedArgumentListContext)

	// EnterArgumentListElement is called when entering the argumentListElement production.
	EnterArgumentListElement(c *ArgumentListElementContext)

	// EnterEnhancedArgumentListElement is called when entering the enhancedArgumentListElement production.
	EnterEnhancedArgumentListElement(c *EnhancedArgumentListElementContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterClassName is called when entering the className production.
	EnterClassName(c *ClassNameContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterBuiltInType is called when entering the builtInType production.
	EnterBuiltInType(c *BuiltInTypeContext)

	// EnterKeywords is called when entering the keywords production.
	EnterKeywords(c *KeywordsContext)

	// EnterRparen is called when entering the rparen production.
	EnterRparen(c *RparenContext)

	// EnterNls is called when entering the nls production.
	EnterNls(c *NlsContext)

	// EnterSep is called when entering the sep production.
	EnterSep(c *SepContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitScriptStatements is called when exiting the scriptStatements production.
	ExitScriptStatements(c *ScriptStatementsContext)

	// ExitScriptStatement is called when exiting the scriptStatement production.
	ExitScriptStatement(c *ScriptStatementContext)

	// ExitPackageDeclaration is called when exiting the packageDeclaration production.
	ExitPackageDeclaration(c *PackageDeclarationContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitModifiersOpt is called when exiting the modifiersOpt production.
	ExitModifiersOpt(c *ModifiersOptContext)

	// ExitModifiers is called when exiting the modifiers production.
	ExitModifiers(c *ModifiersContext)

	// ExitClassOrInterfaceModifiersOpt is called when exiting the classOrInterfaceModifiersOpt production.
	ExitClassOrInterfaceModifiersOpt(c *ClassOrInterfaceModifiersOptContext)

	// ExitClassOrInterfaceModifiers is called when exiting the classOrInterfaceModifiers production.
	ExitClassOrInterfaceModifiers(c *ClassOrInterfaceModifiersContext)

	// ExitClassOrInterfaceModifier is called when exiting the classOrInterfaceModifier production.
	ExitClassOrInterfaceModifier(c *ClassOrInterfaceModifierContext)

	// ExitVariableModifier is called when exiting the variableModifier production.
	ExitVariableModifier(c *VariableModifierContext)

	// ExitVariableModifiersOpt is called when exiting the variableModifiersOpt production.
	ExitVariableModifiersOpt(c *VariableModifiersOptContext)

	// ExitVariableModifiers is called when exiting the variableModifiers production.
	ExitVariableModifiers(c *VariableModifiersContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitTypeBound is called when exiting the typeBound production.
	ExitTypeBound(c *TypeBoundContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitEnumConstants is called when exiting the enumConstants production.
	ExitEnumConstants(c *EnumConstantsContext)

	// ExitEnumConstant is called when exiting the enumConstant production.
	ExitEnumConstant(c *EnumConstantContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitMemberDeclaration is called when exiting the memberDeclaration production.
	ExitMemberDeclaration(c *MemberDeclarationContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitMethodName is called when exiting the methodName production.
	ExitMethodName(c *MethodNameContext)

	// ExitReturnType is called when exiting the returnType production.
	ExitReturnType(c *ReturnTypeContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableDeclaratorId is called when exiting the variableDeclaratorId production.
	ExitVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitVariableInitializers is called when exiting the variableInitializers production.
	ExitVariableInitializers(c *VariableInitializersContext)

	// ExitDims is called when exiting the dims production.
	ExitDims(c *DimsContext)

	// ExitDimsOpt is called when exiting the dimsOpt production.
	ExitDimsOpt(c *DimsOptContext)

	// ExitStandardType is called when exiting the standardType production.
	ExitStandardType(c *StandardTypeContext)

	// ExitTypeType is called when exiting the typeType production.
	ExitTypeType(c *TypeTypeContext)

	// ExitClassOrInterfaceType is called when exiting the classOrInterfaceType production.
	ExitClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// ExitGeneralClassOrInterfaceType is called when exiting the generalClassOrInterfaceType production.
	ExitGeneralClassOrInterfaceType(c *GeneralClassOrInterfaceTypeContext)

	// ExitStandardClassOrInterfaceType is called when exiting the standardClassOrInterfaceType production.
	ExitStandardClassOrInterfaceType(c *StandardClassOrInterfaceTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeArgument is called when exiting the typeArgument production.
	ExitTypeArgument(c *TypeArgumentContext)

	// ExitAnnotatedQualifiedClassName is called when exiting the annotatedQualifiedClassName production.
	ExitAnnotatedQualifiedClassName(c *AnnotatedQualifiedClassNameContext)

	// ExitQualifiedClassNameList is called when exiting the qualifiedClassNameList production.
	ExitQualifiedClassNameList(c *QualifiedClassNameListContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitThisFormalParameter is called when exiting the thisFormalParameter production.
	ExitThisFormalParameter(c *ThisFormalParameterContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitQualifiedNameElement is called when exiting the qualifiedNameElement production.
	ExitQualifiedNameElement(c *QualifiedNameElementContext)

	// ExitQualifiedNameElements is called when exiting the qualifiedNameElements production.
	ExitQualifiedNameElements(c *QualifiedNameElementsContext)

	// ExitQualifiedClassName is called when exiting the qualifiedClassName production.
	ExitQualifiedClassName(c *QualifiedClassNameContext)

	// ExitQualifiedStandardClassName is called when exiting the qualifiedStandardClassName production.
	ExitQualifiedStandardClassName(c *QualifiedStandardClassNameContext)

	// ExitIntegerLiteralAlt is called when exiting the integerLiteralAlt production.
	ExitIntegerLiteralAlt(c *IntegerLiteralAltContext)

	// ExitFloatingPointLiteralAlt is called when exiting the floatingPointLiteralAlt production.
	ExitFloatingPointLiteralAlt(c *FloatingPointLiteralAltContext)

	// ExitStringLiteralAlt is called when exiting the stringLiteralAlt production.
	ExitStringLiteralAlt(c *StringLiteralAltContext)

	// ExitBoolLiteralAlt is called when exiting the boolLiteralAlt production.
	ExitBoolLiteralAlt(c *BoolLiteralAltContext)

	// ExitNilLiteralAlt is called when exiting the nilLiteralAlt production.
	ExitNilLiteralAlt(c *NilLiteralAltContext)

	// ExitGstring is called when exiting the gstring production.
	ExitGstring(c *GstringContext)

	// ExitGstringValue is called when exiting the gstringValue production.
	ExitGstringValue(c *GstringValueContext)

	// ExitGstringPath is called when exiting the gstringPath production.
	ExitGstringPath(c *GstringPathContext)

	// ExitLambdaExpression is called when exiting the lambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitStandardLambdaExpression is called when exiting the standardLambdaExpression production.
	ExitStandardLambdaExpression(c *StandardLambdaExpressionContext)

	// ExitLambdaParameters is called when exiting the lambdaParameters production.
	ExitLambdaParameters(c *LambdaParametersContext)

	// ExitStandardLambdaParameters is called when exiting the standardLambdaParameters production.
	ExitStandardLambdaParameters(c *StandardLambdaParametersContext)

	// ExitLambdaBody is called when exiting the lambdaBody production.
	ExitLambdaBody(c *LambdaBodyContext)

	// ExitClosure is called when exiting the closure production.
	ExitClosure(c *ClosureContext)

	// ExitClosureOrLambdaExpression is called when exiting the closureOrLambdaExpression production.
	ExitClosureOrLambdaExpression(c *ClosureOrLambdaExpressionContext)

	// ExitBlockStatementsOpt is called when exiting the blockStatementsOpt production.
	ExitBlockStatementsOpt(c *BlockStatementsOptContext)

	// ExitBlockStatements is called when exiting the blockStatements production.
	ExitBlockStatements(c *BlockStatementsContext)

	// ExitAnnotationsOpt is called when exiting the annotationsOpt production.
	ExitAnnotationsOpt(c *AnnotationsOptContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitElementValues is called when exiting the elementValues production.
	ExitElementValues(c *ElementValuesContext)

	// ExitAnnotationName is called when exiting the annotationName production.
	ExitAnnotationName(c *AnnotationNameContext)

	// ExitElementValuePairs is called when exiting the elementValuePairs production.
	ExitElementValuePairs(c *ElementValuePairsContext)

	// ExitElementValuePair is called when exiting the elementValuePair production.
	ExitElementValuePair(c *ElementValuePairContext)

	// ExitElementValuePairName is called when exiting the elementValuePairName production.
	ExitElementValuePairName(c *ElementValuePairNameContext)

	// ExitElementValue is called when exiting the elementValue production.
	ExitElementValue(c *ElementValueContext)

	// ExitElementValueArrayInitializer is called when exiting the elementValueArrayInitializer production.
	ExitElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitClassifiedModifiers is called when exiting the classifiedModifiers production.
	ExitClassifiedModifiers(c *ClassifiedModifiersContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitTypeNamePairs is called when exiting the typeNamePairs production.
	ExitTypeNamePairs(c *TypeNamePairsContext)

	// ExitTypeNamePair is called when exiting the typeNamePair production.
	ExitTypeNamePair(c *TypeNamePairContext)

	// ExitVariableNames is called when exiting the variableNames production.
	ExitVariableNames(c *VariableNamesContext)

	// ExitConditionalStatement is called when exiting the conditionalStatement production.
	ExitConditionalStatement(c *ConditionalStatementContext)

	// ExitIfElseStatement is called when exiting the ifElseStatement production.
	ExitIfElseStatement(c *IfElseStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitForStmtAlt is called when exiting the forStmtAlt production.
	ExitForStmtAlt(c *ForStmtAltContext)

	// ExitWhileStmtAlt is called when exiting the whileStmtAlt production.
	ExitWhileStmtAlt(c *WhileStmtAltContext)

	// ExitDoWhileStmtAlt is called when exiting the doWhileStmtAlt production.
	ExitDoWhileStmtAlt(c *DoWhileStmtAltContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitTryCatchStatement is called when exiting the tryCatchStatement production.
	ExitTryCatchStatement(c *TryCatchStatementContext)

	// ExitAssertStatement is called when exiting the assertStatement production.
	ExitAssertStatement(c *AssertStatementContext)

	// ExitBlockStmtAlt is called when exiting the blockStmtAlt production.
	ExitBlockStmtAlt(c *BlockStmtAltContext)

	// ExitConditionalStmtAlt is called when exiting the conditionalStmtAlt production.
	ExitConditionalStmtAlt(c *ConditionalStmtAltContext)

	// ExitLoopStmtAlt is called when exiting the loopStmtAlt production.
	ExitLoopStmtAlt(c *LoopStmtAltContext)

	// ExitTryCatchStmtAlt is called when exiting the tryCatchStmtAlt production.
	ExitTryCatchStmtAlt(c *TryCatchStmtAltContext)

	// ExitSynchronizedStmtAlt is called when exiting the synchronizedStmtAlt production.
	ExitSynchronizedStmtAlt(c *SynchronizedStmtAltContext)

	// ExitReturnStmtAlt is called when exiting the returnStmtAlt production.
	ExitReturnStmtAlt(c *ReturnStmtAltContext)

	// ExitThrowStmtAlt is called when exiting the throwStmtAlt production.
	ExitThrowStmtAlt(c *ThrowStmtAltContext)

	// ExitBreakStmtAlt is called when exiting the breakStmtAlt production.
	ExitBreakStmtAlt(c *BreakStmtAltContext)

	// ExitContinueStmtAlt is called when exiting the continueStmtAlt production.
	ExitContinueStmtAlt(c *ContinueStmtAltContext)

	// ExitLabeledStmtAlt is called when exiting the labeledStmtAlt production.
	ExitLabeledStmtAlt(c *LabeledStmtAltContext)

	// ExitAssertStmtAlt is called when exiting the assertStmtAlt production.
	ExitAssertStmtAlt(c *AssertStmtAltContext)

	// ExitLocalVariableDeclarationStmtAlt is called when exiting the localVariableDeclarationStmtAlt production.
	ExitLocalVariableDeclarationStmtAlt(c *LocalVariableDeclarationStmtAltContext)

	// ExitMethodDeclarationStmtAlt is called when exiting the methodDeclarationStmtAlt production.
	ExitMethodDeclarationStmtAlt(c *MethodDeclarationStmtAltContext)

	// ExitExpressionStmtAlt is called when exiting the expressionStmtAlt production.
	ExitExpressionStmtAlt(c *ExpressionStmtAltContext)

	// ExitEmptyStmtAlt is called when exiting the emptyStmtAlt production.
	ExitEmptyStmtAlt(c *EmptyStmtAltContext)

	// ExitCatchClause is called when exiting the catchClause production.
	ExitCatchClause(c *CatchClauseContext)

	// ExitCatchType is called when exiting the catchType production.
	ExitCatchType(c *CatchTypeContext)

	// ExitFinallyBlock is called when exiting the finallyBlock production.
	ExitFinallyBlock(c *FinallyBlockContext)

	// ExitResources is called when exiting the resources production.
	ExitResources(c *ResourcesContext)

	// ExitResourceList is called when exiting the resourceList production.
	ExitResourceList(c *ResourceListContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitSwitchBlockStatementGroup is called when exiting the switchBlockStatementGroup production.
	ExitSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// ExitSwitchLabel is called when exiting the switchLabel production.
	ExitSwitchLabel(c *SwitchLabelContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitEnhancedForControl is called when exiting the enhancedForControl production.
	ExitEnhancedForControl(c *EnhancedForControlContext)

	// ExitClassicalForControl is called when exiting the classicalForControl production.
	ExitClassicalForControl(c *ClassicalForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitForUpdate is called when exiting the forUpdate production.
	ExitForUpdate(c *ForUpdateContext)

	// ExitCastParExpression is called when exiting the castParExpression production.
	ExitCastParExpression(c *CastParExpressionContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitExpressionInPar is called when exiting the expressionInPar production.
	ExitExpressionInPar(c *ExpressionInParContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpressionListElement is called when exiting the expressionListElement production.
	ExitExpressionListElement(c *ExpressionListElementContext)

	// ExitEnhancedStatementExpression is called when exiting the enhancedStatementExpression production.
	ExitEnhancedStatementExpression(c *EnhancedStatementExpressionContext)

	// ExitCommandExprAlt is called when exiting the commandExprAlt production.
	ExitCommandExprAlt(c *CommandExprAltContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitCastExprAltForExpr is called when exiting the castExprAltForExpr production.
	ExitCastExprAltForExpr(c *CastExprAltForExprContext)

	// ExitShiftExprAlt is called when exiting the shiftExprAlt production.
	ExitShiftExprAlt(c *ShiftExprAltContext)

	// ExitMultipleAssignmentExprAlt is called when exiting the multipleAssignmentExprAlt production.
	ExitMultipleAssignmentExprAlt(c *MultipleAssignmentExprAltContext)

	// ExitExclusiveOrExprAlt is called when exiting the exclusiveOrExprAlt production.
	ExitExclusiveOrExprAlt(c *ExclusiveOrExprAltContext)

	// ExitAdditiveExprAlt is called when exiting the additiveExprAlt production.
	ExitAdditiveExprAlt(c *AdditiveExprAltContext)

	// ExitRegexExprAlt is called when exiting the regexExprAlt production.
	ExitRegexExprAlt(c *RegexExprAltContext)

	// ExitUnaryNotExprAltForExpr is called when exiting the unaryNotExprAltForExpr production.
	ExitUnaryNotExprAltForExpr(c *UnaryNotExprAltForExprContext)

	// ExitConditionalExprAlt is called when exiting the conditionalExprAlt production.
	ExitConditionalExprAlt(c *ConditionalExprAltContext)

	// ExitPowerExprAlt is called when exiting the powerExprAlt production.
	ExitPowerExprAlt(c *PowerExprAltContext)

	// ExitRelationalExprAlt is called when exiting the relationalExprAlt production.
	ExitRelationalExprAlt(c *RelationalExprAltContext)

	// ExitLogicalAndExprAlt is called when exiting the logicalAndExprAlt production.
	ExitLogicalAndExprAlt(c *LogicalAndExprAltContext)

	// ExitAssignmentExprAlt is called when exiting the assignmentExprAlt production.
	ExitAssignmentExprAlt(c *AssignmentExprAltContext)

	// ExitPostfixExprAltForExpr is called when exiting the postfixExprAltForExpr production.
	ExitPostfixExprAltForExpr(c *PostfixExprAltForExprContext)

	// ExitMultiplicativeExprAlt is called when exiting the multiplicativeExprAlt production.
	ExitMultiplicativeExprAlt(c *MultiplicativeExprAltContext)

	// ExitInclusiveOrExprAlt is called when exiting the inclusiveOrExprAlt production.
	ExitInclusiveOrExprAlt(c *InclusiveOrExprAltContext)

	// ExitLogicalOrExprAlt is called when exiting the logicalOrExprAlt production.
	ExitLogicalOrExprAlt(c *LogicalOrExprAltContext)

	// ExitEqualityExprAlt is called when exiting the equalityExprAlt production.
	ExitEqualityExprAlt(c *EqualityExprAltContext)

	// ExitAndExprAlt is called when exiting the andExprAlt production.
	ExitAndExprAlt(c *AndExprAltContext)

	// ExitUnaryAddExprAltForExpr is called when exiting the unaryAddExprAltForExpr production.
	ExitUnaryAddExprAltForExpr(c *UnaryAddExprAltForExprContext)

	// ExitCastExprAlt is called when exiting the castExprAlt production.
	ExitCastExprAlt(c *CastExprAltContext)

	// ExitPostfixExprAlt is called when exiting the postfixExprAlt production.
	ExitPostfixExprAlt(c *PostfixExprAltContext)

	// ExitUnaryNotExprAlt is called when exiting the unaryNotExprAlt production.
	ExitUnaryNotExprAlt(c *UnaryNotExprAltContext)

	// ExitUnaryAddExprAlt is called when exiting the unaryAddExprAlt production.
	ExitUnaryAddExprAlt(c *UnaryAddExprAltContext)

	// ExitCommandExpression is called when exiting the commandExpression production.
	ExitCommandExpression(c *CommandExpressionContext)

	// ExitCommandArgument is called when exiting the commandArgument production.
	ExitCommandArgument(c *CommandArgumentContext)

	// ExitPathExpression is called when exiting the pathExpression production.
	ExitPathExpression(c *PathExpressionContext)

	// ExitPathElement is called when exiting the pathElement production.
	ExitPathElement(c *PathElementContext)

	// ExitNamePart is called when exiting the namePart production.
	ExitNamePart(c *NamePartContext)

	// ExitDynamicMemberName is called when exiting the dynamicMemberName production.
	ExitDynamicMemberName(c *DynamicMemberNameContext)

	// ExitIndexPropertyArgs is called when exiting the indexPropertyArgs production.
	ExitIndexPropertyArgs(c *IndexPropertyArgsContext)

	// ExitNamedPropertyArgs is called when exiting the namedPropertyArgs production.
	ExitNamedPropertyArgs(c *NamedPropertyArgsContext)

	// ExitIdentifierPrmrAlt is called when exiting the identifierPrmrAlt production.
	ExitIdentifierPrmrAlt(c *IdentifierPrmrAltContext)

	// ExitLiteralPrmrAlt is called when exiting the literalPrmrAlt production.
	ExitLiteralPrmrAlt(c *LiteralPrmrAltContext)

	// ExitGstringPrmrAlt is called when exiting the gstringPrmrAlt production.
	ExitGstringPrmrAlt(c *GstringPrmrAltContext)

	// ExitNewPrmrAlt is called when exiting the newPrmrAlt production.
	ExitNewPrmrAlt(c *NewPrmrAltContext)

	// ExitThisPrmrAlt is called when exiting the thisPrmrAlt production.
	ExitThisPrmrAlt(c *ThisPrmrAltContext)

	// ExitSuperPrmrAlt is called when exiting the superPrmrAlt production.
	ExitSuperPrmrAlt(c *SuperPrmrAltContext)

	// ExitParenPrmrAlt is called when exiting the parenPrmrAlt production.
	ExitParenPrmrAlt(c *ParenPrmrAltContext)

	// ExitClosureOrLambdaExpressionPrmrAlt is called when exiting the closureOrLambdaExpressionPrmrAlt production.
	ExitClosureOrLambdaExpressionPrmrAlt(c *ClosureOrLambdaExpressionPrmrAltContext)

	// ExitListPrmrAlt is called when exiting the listPrmrAlt production.
	ExitListPrmrAlt(c *ListPrmrAltContext)

	// ExitMapPrmrAlt is called when exiting the mapPrmrAlt production.
	ExitMapPrmrAlt(c *MapPrmrAltContext)

	// ExitBuiltInTypePrmrAlt is called when exiting the builtInTypePrmrAlt production.
	ExitBuiltInTypePrmrAlt(c *BuiltInTypePrmrAltContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitNormalMap is called when exiting the normalMap production.
	ExitNormalMap(c *NormalMapContext)

	// ExitMapEntryList is called when exiting the mapEntryList production.
	ExitMapEntryList(c *MapEntryListContext)

	// ExitMapEntry is called when exiting the mapEntry production.
	ExitMapEntry(c *MapEntryContext)

	// ExitMapEntryLabel is called when exiting the mapEntryLabel production.
	ExitMapEntryLabel(c *MapEntryLabelContext)

	// ExitCreator is called when exiting the creator production.
	ExitCreator(c *CreatorContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitAnonymousInnerClassDeclaration is called when exiting the anonymousInnerClassDeclaration production.
	ExitAnonymousInnerClassDeclaration(c *AnonymousInnerClassDeclarationContext)

	// ExitCreatedName is called when exiting the createdName production.
	ExitCreatedName(c *CreatedNameContext)

	// ExitNonWildcardTypeArguments is called when exiting the nonWildcardTypeArguments production.
	ExitNonWildcardTypeArguments(c *NonWildcardTypeArgumentsContext)

	// ExitTypeArgumentsOrDiamond is called when exiting the typeArgumentsOrDiamond production.
	ExitTypeArgumentsOrDiamond(c *TypeArgumentsOrDiamondContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitEnhancedArgumentList is called when exiting the enhancedArgumentList production.
	ExitEnhancedArgumentList(c *EnhancedArgumentListContext)

	// ExitArgumentListElement is called when exiting the argumentListElement production.
	ExitArgumentListElement(c *ArgumentListElementContext)

	// ExitEnhancedArgumentListElement is called when exiting the enhancedArgumentListElement production.
	ExitEnhancedArgumentListElement(c *EnhancedArgumentListElementContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitClassName is called when exiting the className production.
	ExitClassName(c *ClassNameContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitBuiltInType is called when exiting the builtInType production.
	ExitBuiltInType(c *BuiltInTypeContext)

	// ExitKeywords is called when exiting the keywords production.
	ExitKeywords(c *KeywordsContext)

	// ExitRparen is called when exiting the rparen production.
	ExitRparen(c *RparenContext)

	// ExitNls is called when exiting the nls production.
	ExitNls(c *NlsContext)

	// ExitSep is called when exiting the sep production.
	ExitSep(c *SepContext)
}
