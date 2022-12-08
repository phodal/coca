// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // TypeScriptParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// TypeScriptParserListener is a complete listener for a parse tree produced by TypeScriptParser.
type TypeScriptParserListener interface {
	antlr.ParseTreeListener

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterBindingPattern is called when entering the bindingPattern production.
	EnterBindingPattern(c *BindingPatternContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterTypeParameterList is called when entering the typeParameterList production.
	EnterTypeParameterList(c *TypeParameterListContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterConstraint is called when entering the constraint production.
	EnterConstraint(c *ConstraintContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterTypeArgumentList is called when entering the typeArgumentList production.
	EnterTypeArgumentList(c *TypeArgumentListContext)

	// EnterTypeArgument is called when entering the typeArgument production.
	EnterTypeArgument(c *TypeArgumentContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterIntersection is called when entering the Intersection production.
	EnterIntersection(c *IntersectionContext)

	// EnterPrimary is called when entering the Primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterUnion is called when entering the Union production.
	EnterUnion(c *UnionContext)

	// EnterRedefinitionOfType is called when entering the RedefinitionOfType production.
	EnterRedefinitionOfType(c *RedefinitionOfTypeContext)

	// EnterPredefinedPrimType is called when entering the PredefinedPrimType production.
	EnterPredefinedPrimType(c *PredefinedPrimTypeContext)

	// EnterArrayPrimType is called when entering the ArrayPrimType production.
	EnterArrayPrimType(c *ArrayPrimTypeContext)

	// EnterParenthesizedPrimType is called when entering the ParenthesizedPrimType production.
	EnterParenthesizedPrimType(c *ParenthesizedPrimTypeContext)

	// EnterThisPrimType is called when entering the ThisPrimType production.
	EnterThisPrimType(c *ThisPrimTypeContext)

	// EnterTuplePrimType is called when entering the TuplePrimType production.
	EnterTuplePrimType(c *TuplePrimTypeContext)

	// EnterObjectPrimType is called when entering the ObjectPrimType production.
	EnterObjectPrimType(c *ObjectPrimTypeContext)

	// EnterReferencePrimType is called when entering the ReferencePrimType production.
	EnterReferencePrimType(c *ReferencePrimTypeContext)

	// EnterQueryPrimType is called when entering the QueryPrimType production.
	EnterQueryPrimType(c *QueryPrimTypeContext)

	// EnterPredefinedType is called when entering the predefinedType production.
	EnterPredefinedType(c *PredefinedTypeContext)

	// EnterTypeReference is called when entering the typeReference production.
	EnterTypeReference(c *TypeReferenceContext)

	// EnterTypeGeneric is called when entering the typeGeneric production.
	EnterTypeGeneric(c *TypeGenericContext)

	// EnterTypeIncludeGeneric is called when entering the typeIncludeGeneric production.
	EnterTypeIncludeGeneric(c *TypeIncludeGenericContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterObjectType is called when entering the objectType production.
	EnterObjectType(c *ObjectTypeContext)

	// EnterTypeBody is called when entering the typeBody production.
	EnterTypeBody(c *TypeBodyContext)

	// EnterTypeMemberList is called when entering the typeMemberList production.
	EnterTypeMemberList(c *TypeMemberListContext)

	// EnterTypeMember is called when entering the typeMember production.
	EnterTypeMember(c *TypeMemberContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterTupleType is called when entering the tupleType production.
	EnterTupleType(c *TupleTypeContext)

	// EnterTupleElementTypes is called when entering the tupleElementTypes production.
	EnterTupleElementTypes(c *TupleElementTypesContext)

	// EnterFunctionType is called when entering the functionType production.
	EnterFunctionType(c *FunctionTypeContext)

	// EnterConstructorType is called when entering the constructorType production.
	EnterConstructorType(c *ConstructorTypeContext)

	// EnterTypeQuery is called when entering the typeQuery production.
	EnterTypeQuery(c *TypeQueryContext)

	// EnterTypeQueryExpression is called when entering the typeQueryExpression production.
	EnterTypeQueryExpression(c *TypeQueryExpressionContext)

	// EnterPropertySignature is called when entering the propertySignature production.
	EnterPropertySignature(c *PropertySignatureContext)

	// EnterTypeAnnotation is called when entering the typeAnnotation production.
	EnterTypeAnnotation(c *TypeAnnotationContext)

	// EnterCallSignature is called when entering the callSignature production.
	EnterCallSignature(c *CallSignatureContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterRequiredParameterList is called when entering the requiredParameterList production.
	EnterRequiredParameterList(c *RequiredParameterListContext)

	// EnterRequiredParameter is called when entering the requiredParameter production.
	EnterRequiredParameter(c *RequiredParameterContext)

	// EnterAccessibilityModifier is called when entering the accessibilityModifier production.
	EnterAccessibilityModifier(c *AccessibilityModifierContext)

	// EnterIdentifierOrPattern is called when entering the identifierOrPattern production.
	EnterIdentifierOrPattern(c *IdentifierOrPatternContext)

	// EnterOptionalParameterList is called when entering the optionalParameterList production.
	EnterOptionalParameterList(c *OptionalParameterListContext)

	// EnterOptionalParameter is called when entering the optionalParameter production.
	EnterOptionalParameter(c *OptionalParameterContext)

	// EnterRestParameter is called when entering the restParameter production.
	EnterRestParameter(c *RestParameterContext)

	// EnterConstructSignature is called when entering the constructSignature production.
	EnterConstructSignature(c *ConstructSignatureContext)

	// EnterIndexSignature is called when entering the indexSignature production.
	EnterIndexSignature(c *IndexSignatureContext)

	// EnterMethodSignature is called when entering the methodSignature production.
	EnterMethodSignature(c *MethodSignatureContext)

	// EnterTypeAliasDeclaration is called when entering the typeAliasDeclaration production.
	EnterTypeAliasDeclaration(c *TypeAliasDeclarationContext)

	// EnterConstructorDeclaration is called when entering the constructorDeclaration production.
	EnterConstructorDeclaration(c *ConstructorDeclarationContext)

	// EnterInterfaceDeclaration is called when entering the interfaceDeclaration production.
	EnterInterfaceDeclaration(c *InterfaceDeclarationContext)

	// EnterInterfaceExtendsClause is called when entering the interfaceExtendsClause production.
	EnterInterfaceExtendsClause(c *InterfaceExtendsClauseContext)

	// EnterClassOrInterfaceTypeList is called when entering the classOrInterfaceTypeList production.
	EnterClassOrInterfaceTypeList(c *ClassOrInterfaceTypeListContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumBody is called when entering the enumBody production.
	EnterEnumBody(c *EnumBodyContext)

	// EnterEnumMemberList is called when entering the enumMemberList production.
	EnterEnumMemberList(c *EnumMemberListContext)

	// EnterEnumMember is called when entering the enumMember production.
	EnterEnumMember(c *EnumMemberContext)

	// EnterNamespaceDeclaration is called when entering the namespaceDeclaration production.
	EnterNamespaceDeclaration(c *NamespaceDeclarationContext)

	// EnterNamespaceName is called when entering the namespaceName production.
	EnterNamespaceName(c *NamespaceNameContext)

	// EnterImportAliasDeclaration is called when entering the importAliasDeclaration production.
	EnterImportAliasDeclaration(c *ImportAliasDeclarationContext)

	// EnterImportAll is called when entering the importAll production.
	EnterImportAll(c *ImportAllContext)

	// EnterDecoratorList is called when entering the decoratorList production.
	EnterDecoratorList(c *DecoratorListContext)

	// EnterDecorator is called when entering the decorator production.
	EnterDecorator(c *DecoratorContext)

	// EnterDecoratorMemberExpression is called when entering the decoratorMemberExpression production.
	EnterDecoratorMemberExpression(c *DecoratorMemberExpressionContext)

	// EnterDecoratorCallExpression is called when entering the decoratorCallExpression production.
	EnterDecoratorCallExpression(c *DecoratorCallExpressionContext)

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

	// EnterAbstractDeclaration is called when entering the abstractDeclaration production.
	EnterAbstractDeclaration(c *AbstractDeclarationContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterImportFromBlock is called when entering the importFromBlock production.
	EnterImportFromBlock(c *ImportFromBlockContext)

	// EnterMultipleImportStatement is called when entering the multipleImportStatement production.
	EnterMultipleImportStatement(c *MultipleImportStatementContext)

	// EnterExportStatement is called when entering the exportStatement production.
	EnterExportStatement(c *ExportStatementContext)

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

	// EnterForVarStatement is called when entering the ForVarStatement production.
	EnterForVarStatement(c *ForVarStatementContext)

	// EnterForInStatement is called when entering the ForInStatement production.
	EnterForInStatement(c *ForInStatementContext)

	// EnterForVarInStatement is called when entering the ForVarInStatement production.
	EnterForVarInStatement(c *ForVarInStatementContext)

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

	// EnterAssignable is called when entering the assignable production.
	EnterAssignable(c *AssignableContext)

	// EnterFinallyProduction is called when entering the finallyProduction production.
	EnterFinallyProduction(c *FinallyProductionContext)

	// EnterDebuggerStatement is called when entering the debuggerStatement production.
	EnterDebuggerStatement(c *DebuggerStatementContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassHeritage is called when entering the classHeritage production.
	EnterClassHeritage(c *ClassHeritageContext)

	// EnterClassTail is called when entering the classTail production.
	EnterClassTail(c *ClassTailContext)

	// EnterClassExtendsClause is called when entering the classExtendsClause production.
	EnterClassExtendsClause(c *ClassExtendsClauseContext)

	// EnterImplementsClause is called when entering the implementsClause production.
	EnterImplementsClause(c *ImplementsClauseContext)

	// EnterClassElement is called when entering the classElement production.
	EnterClassElement(c *ClassElementContext)

	// EnterPropertyMemberDeclaration is called when entering the propertyMemberDeclaration production.
	EnterPropertyMemberDeclaration(c *PropertyMemberDeclarationContext)

	// EnterPropertyMemberBase is called when entering the propertyMemberBase production.
	EnterPropertyMemberBase(c *PropertyMemberBaseContext)

	// EnterIndexMemberDeclaration is called when entering the indexMemberDeclaration production.
	EnterIndexMemberDeclaration(c *IndexMemberDeclarationContext)

	// EnterGeneratorMethod is called when entering the generatorMethod production.
	EnterGeneratorMethod(c *GeneratorMethodContext)

	// EnterGeneratorFunctionDeclaration is called when entering the generatorFunctionDeclaration production.
	EnterGeneratorFunctionDeclaration(c *GeneratorFunctionDeclarationContext)

	// EnterGeneratorBlock is called when entering the generatorBlock production.
	EnterGeneratorBlock(c *GeneratorBlockContext)

	// EnterGeneratorDefinition is called when entering the generatorDefinition production.
	EnterGeneratorDefinition(c *GeneratorDefinitionContext)

	// EnterIteratorBlock is called when entering the iteratorBlock production.
	EnterIteratorBlock(c *IteratorBlockContext)

	// EnterIteratorDefinition is called when entering the iteratorDefinition production.
	EnterIteratorDefinition(c *IteratorDefinitionContext)

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

	// EnterLastElement is called when entering the lastElement production.
	EnterLastElement(c *LastElementContext)

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

	// EnterMethodProperty is called when entering the MethodProperty production.
	EnterMethodProperty(c *MethodPropertyContext)

	// EnterPropertyShorthand is called when entering the PropertyShorthand production.
	EnterPropertyShorthand(c *PropertyShorthandContext)

	// EnterRestParameterInObject is called when entering the RestParameterInObject production.
	EnterRestParameterInObject(c *RestParameterInObjectContext)

	// EnterGetAccessor is called when entering the getAccessor production.
	EnterGetAccessor(c *GetAccessorContext)

	// EnterSetAccessor is called when entering the setAccessor production.
	EnterSetAccessor(c *SetAccessorContext)

	// EnterPropertyName is called when entering the propertyName production.
	EnterPropertyName(c *PropertyNameContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterLastArgument is called when entering the lastArgument production.
	EnterLastArgument(c *LastArgumentContext)

	// EnterExpressionSequence is called when entering the expressionSequence production.
	EnterExpressionSequence(c *ExpressionSequenceContext)

	// EnterFunctionExpressionDeclaration is called when entering the functionExpressionDeclaration production.
	EnterFunctionExpressionDeclaration(c *FunctionExpressionDeclarationContext)

	// EnterTemplateStringExpression is called when entering the TemplateStringExpression production.
	EnterTemplateStringExpression(c *TemplateStringExpressionContext)

	// EnterGeneratorsExpression is called when entering the GeneratorsExpression production.
	EnterGeneratorsExpression(c *GeneratorsExpressionContext)

	// EnterPowerExpression is called when entering the PowerExpression production.
	EnterPowerExpression(c *PowerExpressionContext)

	// EnterInExpression is called when entering the InExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterGenericTypes is called when entering the GenericTypes production.
	EnterGenericTypes(c *GenericTypesContext)

	// EnterArgumentsExpression is called when entering the ArgumentsExpression production.
	EnterArgumentsExpression(c *ArgumentsExpressionContext)

	// EnterThisExpression is called when entering the ThisExpression production.
	EnterThisExpression(c *ThisExpressionContext)

	// EnterTypeofExpression is called when entering the TypeofExpression production.
	EnterTypeofExpression(c *TypeofExpressionContext)

	// EnterGeneratorsFunctionExpression is called when entering the GeneratorsFunctionExpression production.
	EnterGeneratorsFunctionExpression(c *GeneratorsFunctionExpressionContext)

	// EnterEqualityExpression is called when entering the EqualityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterBitXOrExpression is called when entering the BitXOrExpression production.
	EnterBitXOrExpression(c *BitXOrExpressionContext)

	// EnterMultiplicativeExpression is called when entering the MultiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterBitShiftExpression is called when entering the BitShiftExpression production.
	EnterBitShiftExpression(c *BitShiftExpressionContext)

	// EnterAdditiveExpression is called when entering the AdditiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterRelationalExpression is called when entering the RelationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

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

	// EnterMemberIndexExpression is called when entering the MemberIndexExpression production.
	EnterMemberIndexExpression(c *MemberIndexExpressionContext)

	// EnterBitAndExpression is called when entering the BitAndExpression production.
	EnterBitAndExpression(c *BitAndExpressionContext)

	// EnterBitOrExpression is called when entering the BitOrExpression production.
	EnterBitOrExpression(c *BitOrExpressionContext)

	// EnterAssignmentOperatorExpression is called when entering the AssignmentOperatorExpression production.
	EnterAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext)

	// EnterVoidExpression is called when entering the VoidExpression production.
	EnterVoidExpression(c *VoidExpressionContext)

	// EnterTernaryExpression is called when entering the TernaryExpression production.
	EnterTernaryExpression(c *TernaryExpressionContext)

	// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterPreIncrementExpression is called when entering the PreIncrementExpression production.
	EnterPreIncrementExpression(c *PreIncrementExpressionContext)

	// EnterObjectLiteralExpression is called when entering the ObjectLiteralExpression production.
	EnterObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterNotExpression is called when entering the NotExpression production.
	EnterNotExpression(c *NotExpressionContext)

	// EnterPreDecreaseExpression is called when entering the PreDecreaseExpression production.
	EnterPreDecreaseExpression(c *PreDecreaseExpressionContext)

	// EnterAwaitExpression is called when entering the AwaitExpression production.
	EnterAwaitExpression(c *AwaitExpressionContext)

	// EnterFunctionExpression is called when entering the FunctionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterUnaryMinusExpression is called when entering the UnaryMinusExpression production.
	EnterUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// EnterAssignmentExpression is called when entering the AssignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterPostDecreaseExpression is called when entering the PostDecreaseExpression production.
	EnterPostDecreaseExpression(c *PostDecreaseExpressionContext)

	// EnterInstanceofExpression is called when entering the InstanceofExpression production.
	EnterInstanceofExpression(c *InstanceofExpressionContext)

	// EnterUnaryPlusExpression is called when entering the UnaryPlusExpression production.
	EnterUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// EnterDeleteExpression is called when entering the DeleteExpression production.
	EnterDeleteExpression(c *DeleteExpressionContext)

	// EnterArrowFunctionExpression is called when entering the ArrowFunctionExpression production.
	EnterArrowFunctionExpression(c *ArrowFunctionExpressionContext)

	// EnterIteratorsExpression is called when entering the IteratorsExpression production.
	EnterIteratorsExpression(c *IteratorsExpressionContext)

	// EnterSuperExpression is called when entering the SuperExpression production.
	EnterSuperExpression(c *SuperExpressionContext)

	// EnterParenthesizedExpression is called when entering the ParenthesizedExpression production.
	EnterParenthesizedExpression(c *ParenthesizedExpressionContext)

	// EnterPostIncrementExpression is called when entering the PostIncrementExpression production.
	EnterPostIncrementExpression(c *PostIncrementExpressionContext)

	// EnterYieldExpression is called when entering the YieldExpression production.
	EnterYieldExpression(c *YieldExpressionContext)

	// EnterClassExpression is called when entering the ClassExpression production.
	EnterClassExpression(c *ClassExpressionContext)

	// EnterIdentifierExpression is called when entering the IdentifierExpression production.
	EnterIdentifierExpression(c *IdentifierExpressionContext)

	// EnterCoalesceExpression is called when entering the CoalesceExpression production.
	EnterCoalesceExpression(c *CoalesceExpressionContext)

	// EnterArrowFunctionDeclaration is called when entering the arrowFunctionDeclaration production.
	EnterArrowFunctionDeclaration(c *ArrowFunctionDeclarationContext)

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

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitBindingPattern is called when exiting the bindingPattern production.
	ExitBindingPattern(c *BindingPatternContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitTypeParameterList is called when exiting the typeParameterList production.
	ExitTypeParameterList(c *TypeParameterListContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitConstraint is called when exiting the constraint production.
	ExitConstraint(c *ConstraintContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitTypeArgumentList is called when exiting the typeArgumentList production.
	ExitTypeArgumentList(c *TypeArgumentListContext)

	// ExitTypeArgument is called when exiting the typeArgument production.
	ExitTypeArgument(c *TypeArgumentContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitIntersection is called when exiting the Intersection production.
	ExitIntersection(c *IntersectionContext)

	// ExitPrimary is called when exiting the Primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitUnion is called when exiting the Union production.
	ExitUnion(c *UnionContext)

	// ExitRedefinitionOfType is called when exiting the RedefinitionOfType production.
	ExitRedefinitionOfType(c *RedefinitionOfTypeContext)

	// ExitPredefinedPrimType is called when exiting the PredefinedPrimType production.
	ExitPredefinedPrimType(c *PredefinedPrimTypeContext)

	// ExitArrayPrimType is called when exiting the ArrayPrimType production.
	ExitArrayPrimType(c *ArrayPrimTypeContext)

	// ExitParenthesizedPrimType is called when exiting the ParenthesizedPrimType production.
	ExitParenthesizedPrimType(c *ParenthesizedPrimTypeContext)

	// ExitThisPrimType is called when exiting the ThisPrimType production.
	ExitThisPrimType(c *ThisPrimTypeContext)

	// ExitTuplePrimType is called when exiting the TuplePrimType production.
	ExitTuplePrimType(c *TuplePrimTypeContext)

	// ExitObjectPrimType is called when exiting the ObjectPrimType production.
	ExitObjectPrimType(c *ObjectPrimTypeContext)

	// ExitReferencePrimType is called when exiting the ReferencePrimType production.
	ExitReferencePrimType(c *ReferencePrimTypeContext)

	// ExitQueryPrimType is called when exiting the QueryPrimType production.
	ExitQueryPrimType(c *QueryPrimTypeContext)

	// ExitPredefinedType is called when exiting the predefinedType production.
	ExitPredefinedType(c *PredefinedTypeContext)

	// ExitTypeReference is called when exiting the typeReference production.
	ExitTypeReference(c *TypeReferenceContext)

	// ExitTypeGeneric is called when exiting the typeGeneric production.
	ExitTypeGeneric(c *TypeGenericContext)

	// ExitTypeIncludeGeneric is called when exiting the typeIncludeGeneric production.
	ExitTypeIncludeGeneric(c *TypeIncludeGenericContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitObjectType is called when exiting the objectType production.
	ExitObjectType(c *ObjectTypeContext)

	// ExitTypeBody is called when exiting the typeBody production.
	ExitTypeBody(c *TypeBodyContext)

	// ExitTypeMemberList is called when exiting the typeMemberList production.
	ExitTypeMemberList(c *TypeMemberListContext)

	// ExitTypeMember is called when exiting the typeMember production.
	ExitTypeMember(c *TypeMemberContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitTupleType is called when exiting the tupleType production.
	ExitTupleType(c *TupleTypeContext)

	// ExitTupleElementTypes is called when exiting the tupleElementTypes production.
	ExitTupleElementTypes(c *TupleElementTypesContext)

	// ExitFunctionType is called when exiting the functionType production.
	ExitFunctionType(c *FunctionTypeContext)

	// ExitConstructorType is called when exiting the constructorType production.
	ExitConstructorType(c *ConstructorTypeContext)

	// ExitTypeQuery is called when exiting the typeQuery production.
	ExitTypeQuery(c *TypeQueryContext)

	// ExitTypeQueryExpression is called when exiting the typeQueryExpression production.
	ExitTypeQueryExpression(c *TypeQueryExpressionContext)

	// ExitPropertySignature is called when exiting the propertySignature production.
	ExitPropertySignature(c *PropertySignatureContext)

	// ExitTypeAnnotation is called when exiting the typeAnnotation production.
	ExitTypeAnnotation(c *TypeAnnotationContext)

	// ExitCallSignature is called when exiting the callSignature production.
	ExitCallSignature(c *CallSignatureContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitRequiredParameterList is called when exiting the requiredParameterList production.
	ExitRequiredParameterList(c *RequiredParameterListContext)

	// ExitRequiredParameter is called when exiting the requiredParameter production.
	ExitRequiredParameter(c *RequiredParameterContext)

	// ExitAccessibilityModifier is called when exiting the accessibilityModifier production.
	ExitAccessibilityModifier(c *AccessibilityModifierContext)

	// ExitIdentifierOrPattern is called when exiting the identifierOrPattern production.
	ExitIdentifierOrPattern(c *IdentifierOrPatternContext)

	// ExitOptionalParameterList is called when exiting the optionalParameterList production.
	ExitOptionalParameterList(c *OptionalParameterListContext)

	// ExitOptionalParameter is called when exiting the optionalParameter production.
	ExitOptionalParameter(c *OptionalParameterContext)

	// ExitRestParameter is called when exiting the restParameter production.
	ExitRestParameter(c *RestParameterContext)

	// ExitConstructSignature is called when exiting the constructSignature production.
	ExitConstructSignature(c *ConstructSignatureContext)

	// ExitIndexSignature is called when exiting the indexSignature production.
	ExitIndexSignature(c *IndexSignatureContext)

	// ExitMethodSignature is called when exiting the methodSignature production.
	ExitMethodSignature(c *MethodSignatureContext)

	// ExitTypeAliasDeclaration is called when exiting the typeAliasDeclaration production.
	ExitTypeAliasDeclaration(c *TypeAliasDeclarationContext)

	// ExitConstructorDeclaration is called when exiting the constructorDeclaration production.
	ExitConstructorDeclaration(c *ConstructorDeclarationContext)

	// ExitInterfaceDeclaration is called when exiting the interfaceDeclaration production.
	ExitInterfaceDeclaration(c *InterfaceDeclarationContext)

	// ExitInterfaceExtendsClause is called when exiting the interfaceExtendsClause production.
	ExitInterfaceExtendsClause(c *InterfaceExtendsClauseContext)

	// ExitClassOrInterfaceTypeList is called when exiting the classOrInterfaceTypeList production.
	ExitClassOrInterfaceTypeList(c *ClassOrInterfaceTypeListContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumBody is called when exiting the enumBody production.
	ExitEnumBody(c *EnumBodyContext)

	// ExitEnumMemberList is called when exiting the enumMemberList production.
	ExitEnumMemberList(c *EnumMemberListContext)

	// ExitEnumMember is called when exiting the enumMember production.
	ExitEnumMember(c *EnumMemberContext)

	// ExitNamespaceDeclaration is called when exiting the namespaceDeclaration production.
	ExitNamespaceDeclaration(c *NamespaceDeclarationContext)

	// ExitNamespaceName is called when exiting the namespaceName production.
	ExitNamespaceName(c *NamespaceNameContext)

	// ExitImportAliasDeclaration is called when exiting the importAliasDeclaration production.
	ExitImportAliasDeclaration(c *ImportAliasDeclarationContext)

	// ExitImportAll is called when exiting the importAll production.
	ExitImportAll(c *ImportAllContext)

	// ExitDecoratorList is called when exiting the decoratorList production.
	ExitDecoratorList(c *DecoratorListContext)

	// ExitDecorator is called when exiting the decorator production.
	ExitDecorator(c *DecoratorContext)

	// ExitDecoratorMemberExpression is called when exiting the decoratorMemberExpression production.
	ExitDecoratorMemberExpression(c *DecoratorMemberExpressionContext)

	// ExitDecoratorCallExpression is called when exiting the decoratorCallExpression production.
	ExitDecoratorCallExpression(c *DecoratorCallExpressionContext)

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

	// ExitAbstractDeclaration is called when exiting the abstractDeclaration production.
	ExitAbstractDeclaration(c *AbstractDeclarationContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitImportFromBlock is called when exiting the importFromBlock production.
	ExitImportFromBlock(c *ImportFromBlockContext)

	// ExitMultipleImportStatement is called when exiting the multipleImportStatement production.
	ExitMultipleImportStatement(c *MultipleImportStatementContext)

	// ExitExportStatement is called when exiting the exportStatement production.
	ExitExportStatement(c *ExportStatementContext)

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

	// ExitForVarStatement is called when exiting the ForVarStatement production.
	ExitForVarStatement(c *ForVarStatementContext)

	// ExitForInStatement is called when exiting the ForInStatement production.
	ExitForInStatement(c *ForInStatementContext)

	// ExitForVarInStatement is called when exiting the ForVarInStatement production.
	ExitForVarInStatement(c *ForVarInStatementContext)

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

	// ExitAssignable is called when exiting the assignable production.
	ExitAssignable(c *AssignableContext)

	// ExitFinallyProduction is called when exiting the finallyProduction production.
	ExitFinallyProduction(c *FinallyProductionContext)

	// ExitDebuggerStatement is called when exiting the debuggerStatement production.
	ExitDebuggerStatement(c *DebuggerStatementContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassHeritage is called when exiting the classHeritage production.
	ExitClassHeritage(c *ClassHeritageContext)

	// ExitClassTail is called when exiting the classTail production.
	ExitClassTail(c *ClassTailContext)

	// ExitClassExtendsClause is called when exiting the classExtendsClause production.
	ExitClassExtendsClause(c *ClassExtendsClauseContext)

	// ExitImplementsClause is called when exiting the implementsClause production.
	ExitImplementsClause(c *ImplementsClauseContext)

	// ExitClassElement is called when exiting the classElement production.
	ExitClassElement(c *ClassElementContext)

	// ExitPropertyMemberDeclaration is called when exiting the propertyMemberDeclaration production.
	ExitPropertyMemberDeclaration(c *PropertyMemberDeclarationContext)

	// ExitPropertyMemberBase is called when exiting the propertyMemberBase production.
	ExitPropertyMemberBase(c *PropertyMemberBaseContext)

	// ExitIndexMemberDeclaration is called when exiting the indexMemberDeclaration production.
	ExitIndexMemberDeclaration(c *IndexMemberDeclarationContext)

	// ExitGeneratorMethod is called when exiting the generatorMethod production.
	ExitGeneratorMethod(c *GeneratorMethodContext)

	// ExitGeneratorFunctionDeclaration is called when exiting the generatorFunctionDeclaration production.
	ExitGeneratorFunctionDeclaration(c *GeneratorFunctionDeclarationContext)

	// ExitGeneratorBlock is called when exiting the generatorBlock production.
	ExitGeneratorBlock(c *GeneratorBlockContext)

	// ExitGeneratorDefinition is called when exiting the generatorDefinition production.
	ExitGeneratorDefinition(c *GeneratorDefinitionContext)

	// ExitIteratorBlock is called when exiting the iteratorBlock production.
	ExitIteratorBlock(c *IteratorBlockContext)

	// ExitIteratorDefinition is called when exiting the iteratorDefinition production.
	ExitIteratorDefinition(c *IteratorDefinitionContext)

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

	// ExitLastElement is called when exiting the lastElement production.
	ExitLastElement(c *LastElementContext)

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

	// ExitMethodProperty is called when exiting the MethodProperty production.
	ExitMethodProperty(c *MethodPropertyContext)

	// ExitPropertyShorthand is called when exiting the PropertyShorthand production.
	ExitPropertyShorthand(c *PropertyShorthandContext)

	// ExitRestParameterInObject is called when exiting the RestParameterInObject production.
	ExitRestParameterInObject(c *RestParameterInObjectContext)

	// ExitGetAccessor is called when exiting the getAccessor production.
	ExitGetAccessor(c *GetAccessorContext)

	// ExitSetAccessor is called when exiting the setAccessor production.
	ExitSetAccessor(c *SetAccessorContext)

	// ExitPropertyName is called when exiting the propertyName production.
	ExitPropertyName(c *PropertyNameContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitLastArgument is called when exiting the lastArgument production.
	ExitLastArgument(c *LastArgumentContext)

	// ExitExpressionSequence is called when exiting the expressionSequence production.
	ExitExpressionSequence(c *ExpressionSequenceContext)

	// ExitFunctionExpressionDeclaration is called when exiting the functionExpressionDeclaration production.
	ExitFunctionExpressionDeclaration(c *FunctionExpressionDeclarationContext)

	// ExitTemplateStringExpression is called when exiting the TemplateStringExpression production.
	ExitTemplateStringExpression(c *TemplateStringExpressionContext)

	// ExitGeneratorsExpression is called when exiting the GeneratorsExpression production.
	ExitGeneratorsExpression(c *GeneratorsExpressionContext)

	// ExitPowerExpression is called when exiting the PowerExpression production.
	ExitPowerExpression(c *PowerExpressionContext)

	// ExitInExpression is called when exiting the InExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitGenericTypes is called when exiting the GenericTypes production.
	ExitGenericTypes(c *GenericTypesContext)

	// ExitArgumentsExpression is called when exiting the ArgumentsExpression production.
	ExitArgumentsExpression(c *ArgumentsExpressionContext)

	// ExitThisExpression is called when exiting the ThisExpression production.
	ExitThisExpression(c *ThisExpressionContext)

	// ExitTypeofExpression is called when exiting the TypeofExpression production.
	ExitTypeofExpression(c *TypeofExpressionContext)

	// ExitGeneratorsFunctionExpression is called when exiting the GeneratorsFunctionExpression production.
	ExitGeneratorsFunctionExpression(c *GeneratorsFunctionExpressionContext)

	// ExitEqualityExpression is called when exiting the EqualityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitBitXOrExpression is called when exiting the BitXOrExpression production.
	ExitBitXOrExpression(c *BitXOrExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the MultiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitBitShiftExpression is called when exiting the BitShiftExpression production.
	ExitBitShiftExpression(c *BitShiftExpressionContext)

	// ExitAdditiveExpression is called when exiting the AdditiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitRelationalExpression is called when exiting the RelationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

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

	// ExitMemberIndexExpression is called when exiting the MemberIndexExpression production.
	ExitMemberIndexExpression(c *MemberIndexExpressionContext)

	// ExitBitAndExpression is called when exiting the BitAndExpression production.
	ExitBitAndExpression(c *BitAndExpressionContext)

	// ExitBitOrExpression is called when exiting the BitOrExpression production.
	ExitBitOrExpression(c *BitOrExpressionContext)

	// ExitAssignmentOperatorExpression is called when exiting the AssignmentOperatorExpression production.
	ExitAssignmentOperatorExpression(c *AssignmentOperatorExpressionContext)

	// ExitVoidExpression is called when exiting the VoidExpression production.
	ExitVoidExpression(c *VoidExpressionContext)

	// ExitTernaryExpression is called when exiting the TernaryExpression production.
	ExitTernaryExpression(c *TernaryExpressionContext)

	// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitPreIncrementExpression is called when exiting the PreIncrementExpression production.
	ExitPreIncrementExpression(c *PreIncrementExpressionContext)

	// ExitObjectLiteralExpression is called when exiting the ObjectLiteralExpression production.
	ExitObjectLiteralExpression(c *ObjectLiteralExpressionContext)

	// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitNotExpression is called when exiting the NotExpression production.
	ExitNotExpression(c *NotExpressionContext)

	// ExitPreDecreaseExpression is called when exiting the PreDecreaseExpression production.
	ExitPreDecreaseExpression(c *PreDecreaseExpressionContext)

	// ExitAwaitExpression is called when exiting the AwaitExpression production.
	ExitAwaitExpression(c *AwaitExpressionContext)

	// ExitFunctionExpression is called when exiting the FunctionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitUnaryMinusExpression is called when exiting the UnaryMinusExpression production.
	ExitUnaryMinusExpression(c *UnaryMinusExpressionContext)

	// ExitAssignmentExpression is called when exiting the AssignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitPostDecreaseExpression is called when exiting the PostDecreaseExpression production.
	ExitPostDecreaseExpression(c *PostDecreaseExpressionContext)

	// ExitInstanceofExpression is called when exiting the InstanceofExpression production.
	ExitInstanceofExpression(c *InstanceofExpressionContext)

	// ExitUnaryPlusExpression is called when exiting the UnaryPlusExpression production.
	ExitUnaryPlusExpression(c *UnaryPlusExpressionContext)

	// ExitDeleteExpression is called when exiting the DeleteExpression production.
	ExitDeleteExpression(c *DeleteExpressionContext)

	// ExitArrowFunctionExpression is called when exiting the ArrowFunctionExpression production.
	ExitArrowFunctionExpression(c *ArrowFunctionExpressionContext)

	// ExitIteratorsExpression is called when exiting the IteratorsExpression production.
	ExitIteratorsExpression(c *IteratorsExpressionContext)

	// ExitSuperExpression is called when exiting the SuperExpression production.
	ExitSuperExpression(c *SuperExpressionContext)

	// ExitParenthesizedExpression is called when exiting the ParenthesizedExpression production.
	ExitParenthesizedExpression(c *ParenthesizedExpressionContext)

	// ExitPostIncrementExpression is called when exiting the PostIncrementExpression production.
	ExitPostIncrementExpression(c *PostIncrementExpressionContext)

	// ExitYieldExpression is called when exiting the YieldExpression production.
	ExitYieldExpression(c *YieldExpressionContext)

	// ExitClassExpression is called when exiting the ClassExpression production.
	ExitClassExpression(c *ClassExpressionContext)

	// ExitIdentifierExpression is called when exiting the IdentifierExpression production.
	ExitIdentifierExpression(c *IdentifierExpressionContext)

	// ExitCoalesceExpression is called when exiting the CoalesceExpression production.
	ExitCoalesceExpression(c *CoalesceExpressionContext)

	// ExitArrowFunctionDeclaration is called when exiting the arrowFunctionDeclaration production.
	ExitArrowFunctionDeclaration(c *ArrowFunctionDeclarationContext)

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
