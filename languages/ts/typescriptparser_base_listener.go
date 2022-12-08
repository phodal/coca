// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // TypeScriptParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseTypeScriptParserListener is a complete listener for a parse tree produced by TypeScriptParser.
type BaseTypeScriptParserListener struct{}

var _ TypeScriptParserListener = &BaseTypeScriptParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTypeScriptParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTypeScriptParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTypeScriptParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTypeScriptParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseTypeScriptParserListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseTypeScriptParserListener) ExitInitializer(ctx *InitializerContext) {}

// EnterBindingPattern is called when production bindingPattern is entered.
func (s *BaseTypeScriptParserListener) EnterBindingPattern(ctx *BindingPatternContext) {}

// ExitBindingPattern is called when production bindingPattern is exited.
func (s *BaseTypeScriptParserListener) ExitBindingPattern(ctx *BindingPatternContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseTypeScriptParserListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseTypeScriptParserListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameterList is called when production typeParameterList is entered.
func (s *BaseTypeScriptParserListener) EnterTypeParameterList(ctx *TypeParameterListContext) {}

// ExitTypeParameterList is called when production typeParameterList is exited.
func (s *BaseTypeScriptParserListener) ExitTypeParameterList(ctx *TypeParameterListContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseTypeScriptParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseTypeScriptParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterConstraint is called when production constraint is entered.
func (s *BaseTypeScriptParserListener) EnterConstraint(ctx *ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (s *BaseTypeScriptParserListener) ExitConstraint(ctx *ConstraintContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseTypeScriptParserListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseTypeScriptParserListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeArgumentList is called when production typeArgumentList is entered.
func (s *BaseTypeScriptParserListener) EnterTypeArgumentList(ctx *TypeArgumentListContext) {}

// ExitTypeArgumentList is called when production typeArgumentList is exited.
func (s *BaseTypeScriptParserListener) ExitTypeArgumentList(ctx *TypeArgumentListContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BaseTypeScriptParserListener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BaseTypeScriptParserListener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseTypeScriptParserListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseTypeScriptParserListener) ExitType_(ctx *Type_Context) {}

// EnterIntersection is called when production Intersection is entered.
func (s *BaseTypeScriptParserListener) EnterIntersection(ctx *IntersectionContext) {}

// ExitIntersection is called when production Intersection is exited.
func (s *BaseTypeScriptParserListener) ExitIntersection(ctx *IntersectionContext) {}

// EnterPrimary is called when production Primary is entered.
func (s *BaseTypeScriptParserListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production Primary is exited.
func (s *BaseTypeScriptParserListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterUnion is called when production Union is entered.
func (s *BaseTypeScriptParserListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production Union is exited.
func (s *BaseTypeScriptParserListener) ExitUnion(ctx *UnionContext) {}

// EnterRedefinitionOfType is called when production RedefinitionOfType is entered.
func (s *BaseTypeScriptParserListener) EnterRedefinitionOfType(ctx *RedefinitionOfTypeContext) {}

// ExitRedefinitionOfType is called when production RedefinitionOfType is exited.
func (s *BaseTypeScriptParserListener) ExitRedefinitionOfType(ctx *RedefinitionOfTypeContext) {}

// EnterPredefinedPrimType is called when production PredefinedPrimType is entered.
func (s *BaseTypeScriptParserListener) EnterPredefinedPrimType(ctx *PredefinedPrimTypeContext) {}

// ExitPredefinedPrimType is called when production PredefinedPrimType is exited.
func (s *BaseTypeScriptParserListener) ExitPredefinedPrimType(ctx *PredefinedPrimTypeContext) {}

// EnterArrayPrimType is called when production ArrayPrimType is entered.
func (s *BaseTypeScriptParserListener) EnterArrayPrimType(ctx *ArrayPrimTypeContext) {}

// ExitArrayPrimType is called when production ArrayPrimType is exited.
func (s *BaseTypeScriptParserListener) ExitArrayPrimType(ctx *ArrayPrimTypeContext) {}

// EnterParenthesizedPrimType is called when production ParenthesizedPrimType is entered.
func (s *BaseTypeScriptParserListener) EnterParenthesizedPrimType(ctx *ParenthesizedPrimTypeContext) {
}

// ExitParenthesizedPrimType is called when production ParenthesizedPrimType is exited.
func (s *BaseTypeScriptParserListener) ExitParenthesizedPrimType(ctx *ParenthesizedPrimTypeContext) {}

// EnterThisPrimType is called when production ThisPrimType is entered.
func (s *BaseTypeScriptParserListener) EnterThisPrimType(ctx *ThisPrimTypeContext) {}

// ExitThisPrimType is called when production ThisPrimType is exited.
func (s *BaseTypeScriptParserListener) ExitThisPrimType(ctx *ThisPrimTypeContext) {}

// EnterTuplePrimType is called when production TuplePrimType is entered.
func (s *BaseTypeScriptParserListener) EnterTuplePrimType(ctx *TuplePrimTypeContext) {}

// ExitTuplePrimType is called when production TuplePrimType is exited.
func (s *BaseTypeScriptParserListener) ExitTuplePrimType(ctx *TuplePrimTypeContext) {}

// EnterObjectPrimType is called when production ObjectPrimType is entered.
func (s *BaseTypeScriptParserListener) EnterObjectPrimType(ctx *ObjectPrimTypeContext) {}

// ExitObjectPrimType is called when production ObjectPrimType is exited.
func (s *BaseTypeScriptParserListener) ExitObjectPrimType(ctx *ObjectPrimTypeContext) {}

// EnterReferencePrimType is called when production ReferencePrimType is entered.
func (s *BaseTypeScriptParserListener) EnterReferencePrimType(ctx *ReferencePrimTypeContext) {}

// ExitReferencePrimType is called when production ReferencePrimType is exited.
func (s *BaseTypeScriptParserListener) ExitReferencePrimType(ctx *ReferencePrimTypeContext) {}

// EnterQueryPrimType is called when production QueryPrimType is entered.
func (s *BaseTypeScriptParserListener) EnterQueryPrimType(ctx *QueryPrimTypeContext) {}

// ExitQueryPrimType is called when production QueryPrimType is exited.
func (s *BaseTypeScriptParserListener) ExitQueryPrimType(ctx *QueryPrimTypeContext) {}

// EnterPredefinedType is called when production predefinedType is entered.
func (s *BaseTypeScriptParserListener) EnterPredefinedType(ctx *PredefinedTypeContext) {}

// ExitPredefinedType is called when production predefinedType is exited.
func (s *BaseTypeScriptParserListener) ExitPredefinedType(ctx *PredefinedTypeContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *BaseTypeScriptParserListener) EnterTypeReference(ctx *TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *BaseTypeScriptParserListener) ExitTypeReference(ctx *TypeReferenceContext) {}

// EnterTypeGeneric is called when production typeGeneric is entered.
func (s *BaseTypeScriptParserListener) EnterTypeGeneric(ctx *TypeGenericContext) {}

// ExitTypeGeneric is called when production typeGeneric is exited.
func (s *BaseTypeScriptParserListener) ExitTypeGeneric(ctx *TypeGenericContext) {}

// EnterTypeIncludeGeneric is called when production typeIncludeGeneric is entered.
func (s *BaseTypeScriptParserListener) EnterTypeIncludeGeneric(ctx *TypeIncludeGenericContext) {}

// ExitTypeIncludeGeneric is called when production typeIncludeGeneric is exited.
func (s *BaseTypeScriptParserListener) ExitTypeIncludeGeneric(ctx *TypeIncludeGenericContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseTypeScriptParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseTypeScriptParserListener) ExitTypeName(ctx *TypeNameContext) {}

// EnterObjectType is called when production objectType is entered.
func (s *BaseTypeScriptParserListener) EnterObjectType(ctx *ObjectTypeContext) {}

// ExitObjectType is called when production objectType is exited.
func (s *BaseTypeScriptParserListener) ExitObjectType(ctx *ObjectTypeContext) {}

// EnterTypeBody is called when production typeBody is entered.
func (s *BaseTypeScriptParserListener) EnterTypeBody(ctx *TypeBodyContext) {}

// ExitTypeBody is called when production typeBody is exited.
func (s *BaseTypeScriptParserListener) ExitTypeBody(ctx *TypeBodyContext) {}

// EnterTypeMemberList is called when production typeMemberList is entered.
func (s *BaseTypeScriptParserListener) EnterTypeMemberList(ctx *TypeMemberListContext) {}

// ExitTypeMemberList is called when production typeMemberList is exited.
func (s *BaseTypeScriptParserListener) ExitTypeMemberList(ctx *TypeMemberListContext) {}

// EnterTypeMember is called when production typeMember is entered.
func (s *BaseTypeScriptParserListener) EnterTypeMember(ctx *TypeMemberContext) {}

// ExitTypeMember is called when production typeMember is exited.
func (s *BaseTypeScriptParserListener) ExitTypeMember(ctx *TypeMemberContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseTypeScriptParserListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseTypeScriptParserListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterTupleType is called when production tupleType is entered.
func (s *BaseTypeScriptParserListener) EnterTupleType(ctx *TupleTypeContext) {}

// ExitTupleType is called when production tupleType is exited.
func (s *BaseTypeScriptParserListener) ExitTupleType(ctx *TupleTypeContext) {}

// EnterTupleElementTypes is called when production tupleElementTypes is entered.
func (s *BaseTypeScriptParserListener) EnterTupleElementTypes(ctx *TupleElementTypesContext) {}

// ExitTupleElementTypes is called when production tupleElementTypes is exited.
func (s *BaseTypeScriptParserListener) ExitTupleElementTypes(ctx *TupleElementTypesContext) {}

// EnterFunctionType is called when production functionType is entered.
func (s *BaseTypeScriptParserListener) EnterFunctionType(ctx *FunctionTypeContext) {}

// ExitFunctionType is called when production functionType is exited.
func (s *BaseTypeScriptParserListener) ExitFunctionType(ctx *FunctionTypeContext) {}

// EnterConstructorType is called when production constructorType is entered.
func (s *BaseTypeScriptParserListener) EnterConstructorType(ctx *ConstructorTypeContext) {}

// ExitConstructorType is called when production constructorType is exited.
func (s *BaseTypeScriptParserListener) ExitConstructorType(ctx *ConstructorTypeContext) {}

// EnterTypeQuery is called when production typeQuery is entered.
func (s *BaseTypeScriptParserListener) EnterTypeQuery(ctx *TypeQueryContext) {}

// ExitTypeQuery is called when production typeQuery is exited.
func (s *BaseTypeScriptParserListener) ExitTypeQuery(ctx *TypeQueryContext) {}

// EnterTypeQueryExpression is called when production typeQueryExpression is entered.
func (s *BaseTypeScriptParserListener) EnterTypeQueryExpression(ctx *TypeQueryExpressionContext) {}

// ExitTypeQueryExpression is called when production typeQueryExpression is exited.
func (s *BaseTypeScriptParserListener) ExitTypeQueryExpression(ctx *TypeQueryExpressionContext) {}

// EnterPropertySignature is called when production propertySignature is entered.
func (s *BaseTypeScriptParserListener) EnterPropertySignature(ctx *PropertySignatureContext) {}

// ExitPropertySignature is called when production propertySignature is exited.
func (s *BaseTypeScriptParserListener) ExitPropertySignature(ctx *PropertySignatureContext) {}

// EnterTypeAnnotation is called when production typeAnnotation is entered.
func (s *BaseTypeScriptParserListener) EnterTypeAnnotation(ctx *TypeAnnotationContext) {}

// ExitTypeAnnotation is called when production typeAnnotation is exited.
func (s *BaseTypeScriptParserListener) ExitTypeAnnotation(ctx *TypeAnnotationContext) {}

// EnterCallSignature is called when production callSignature is entered.
func (s *BaseTypeScriptParserListener) EnterCallSignature(ctx *CallSignatureContext) {}

// ExitCallSignature is called when production callSignature is exited.
func (s *BaseTypeScriptParserListener) ExitCallSignature(ctx *CallSignatureContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseTypeScriptParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseTypeScriptParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterRequiredParameterList is called when production requiredParameterList is entered.
func (s *BaseTypeScriptParserListener) EnterRequiredParameterList(ctx *RequiredParameterListContext) {
}

// ExitRequiredParameterList is called when production requiredParameterList is exited.
func (s *BaseTypeScriptParserListener) ExitRequiredParameterList(ctx *RequiredParameterListContext) {}

// EnterRequiredParameter is called when production requiredParameter is entered.
func (s *BaseTypeScriptParserListener) EnterRequiredParameter(ctx *RequiredParameterContext) {}

// ExitRequiredParameter is called when production requiredParameter is exited.
func (s *BaseTypeScriptParserListener) ExitRequiredParameter(ctx *RequiredParameterContext) {}

// EnterAccessibilityModifier is called when production accessibilityModifier is entered.
func (s *BaseTypeScriptParserListener) EnterAccessibilityModifier(ctx *AccessibilityModifierContext) {
}

// ExitAccessibilityModifier is called when production accessibilityModifier is exited.
func (s *BaseTypeScriptParserListener) ExitAccessibilityModifier(ctx *AccessibilityModifierContext) {}

// EnterIdentifierOrPattern is called when production identifierOrPattern is entered.
func (s *BaseTypeScriptParserListener) EnterIdentifierOrPattern(ctx *IdentifierOrPatternContext) {}

// ExitIdentifierOrPattern is called when production identifierOrPattern is exited.
func (s *BaseTypeScriptParserListener) ExitIdentifierOrPattern(ctx *IdentifierOrPatternContext) {}

// EnterOptionalParameterList is called when production optionalParameterList is entered.
func (s *BaseTypeScriptParserListener) EnterOptionalParameterList(ctx *OptionalParameterListContext) {
}

// ExitOptionalParameterList is called when production optionalParameterList is exited.
func (s *BaseTypeScriptParserListener) ExitOptionalParameterList(ctx *OptionalParameterListContext) {}

// EnterOptionalParameter is called when production optionalParameter is entered.
func (s *BaseTypeScriptParserListener) EnterOptionalParameter(ctx *OptionalParameterContext) {}

// ExitOptionalParameter is called when production optionalParameter is exited.
func (s *BaseTypeScriptParserListener) ExitOptionalParameter(ctx *OptionalParameterContext) {}

// EnterRestParameter is called when production restParameter is entered.
func (s *BaseTypeScriptParserListener) EnterRestParameter(ctx *RestParameterContext) {}

// ExitRestParameter is called when production restParameter is exited.
func (s *BaseTypeScriptParserListener) ExitRestParameter(ctx *RestParameterContext) {}

// EnterConstructSignature is called when production constructSignature is entered.
func (s *BaseTypeScriptParserListener) EnterConstructSignature(ctx *ConstructSignatureContext) {}

// ExitConstructSignature is called when production constructSignature is exited.
func (s *BaseTypeScriptParserListener) ExitConstructSignature(ctx *ConstructSignatureContext) {}

// EnterIndexSignature is called when production indexSignature is entered.
func (s *BaseTypeScriptParserListener) EnterIndexSignature(ctx *IndexSignatureContext) {}

// ExitIndexSignature is called when production indexSignature is exited.
func (s *BaseTypeScriptParserListener) ExitIndexSignature(ctx *IndexSignatureContext) {}

// EnterMethodSignature is called when production methodSignature is entered.
func (s *BaseTypeScriptParserListener) EnterMethodSignature(ctx *MethodSignatureContext) {}

// ExitMethodSignature is called when production methodSignature is exited.
func (s *BaseTypeScriptParserListener) ExitMethodSignature(ctx *MethodSignatureContext) {}

// EnterTypeAliasDeclaration is called when production typeAliasDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) {}

// ExitTypeAliasDeclaration is called when production typeAliasDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitTypeAliasDeclaration(ctx *TypeAliasDeclarationContext) {}

// EnterConstructorDeclaration is called when production constructorDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterConstructorDeclaration(ctx *ConstructorDeclarationContext) {
}

// ExitConstructorDeclaration is called when production constructorDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitConstructorDeclaration(ctx *ConstructorDeclarationContext) {
}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterInterfaceExtendsClause is called when production interfaceExtendsClause is entered.
func (s *BaseTypeScriptParserListener) EnterInterfaceExtendsClause(ctx *InterfaceExtendsClauseContext) {
}

// ExitInterfaceExtendsClause is called when production interfaceExtendsClause is exited.
func (s *BaseTypeScriptParserListener) ExitInterfaceExtendsClause(ctx *InterfaceExtendsClauseContext) {
}

// EnterClassOrInterfaceTypeList is called when production classOrInterfaceTypeList is entered.
func (s *BaseTypeScriptParserListener) EnterClassOrInterfaceTypeList(ctx *ClassOrInterfaceTypeListContext) {
}

// ExitClassOrInterfaceTypeList is called when production classOrInterfaceTypeList is exited.
func (s *BaseTypeScriptParserListener) ExitClassOrInterfaceTypeList(ctx *ClassOrInterfaceTypeListContext) {
}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumBody is called when production enumBody is entered.
func (s *BaseTypeScriptParserListener) EnterEnumBody(ctx *EnumBodyContext) {}

// ExitEnumBody is called when production enumBody is exited.
func (s *BaseTypeScriptParserListener) ExitEnumBody(ctx *EnumBodyContext) {}

// EnterEnumMemberList is called when production enumMemberList is entered.
func (s *BaseTypeScriptParserListener) EnterEnumMemberList(ctx *EnumMemberListContext) {}

// ExitEnumMemberList is called when production enumMemberList is exited.
func (s *BaseTypeScriptParserListener) ExitEnumMemberList(ctx *EnumMemberListContext) {}

// EnterEnumMember is called when production enumMember is entered.
func (s *BaseTypeScriptParserListener) EnterEnumMember(ctx *EnumMemberContext) {}

// ExitEnumMember is called when production enumMember is exited.
func (s *BaseTypeScriptParserListener) ExitEnumMember(ctx *EnumMemberContext) {}

// EnterNamespaceDeclaration is called when production namespaceDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// ExitNamespaceDeclaration is called when production namespaceDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitNamespaceDeclaration(ctx *NamespaceDeclarationContext) {}

// EnterNamespaceName is called when production namespaceName is entered.
func (s *BaseTypeScriptParserListener) EnterNamespaceName(ctx *NamespaceNameContext) {}

// ExitNamespaceName is called when production namespaceName is exited.
func (s *BaseTypeScriptParserListener) ExitNamespaceName(ctx *NamespaceNameContext) {}

// EnterImportAliasDeclaration is called when production importAliasDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterImportAliasDeclaration(ctx *ImportAliasDeclarationContext) {
}

// ExitImportAliasDeclaration is called when production importAliasDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitImportAliasDeclaration(ctx *ImportAliasDeclarationContext) {
}

// EnterImportAll is called when production importAll is entered.
func (s *BaseTypeScriptParserListener) EnterImportAll(ctx *ImportAllContext) {}

// ExitImportAll is called when production importAll is exited.
func (s *BaseTypeScriptParserListener) ExitImportAll(ctx *ImportAllContext) {}

// EnterDecoratorList is called when production decoratorList is entered.
func (s *BaseTypeScriptParserListener) EnterDecoratorList(ctx *DecoratorListContext) {}

// ExitDecoratorList is called when production decoratorList is exited.
func (s *BaseTypeScriptParserListener) ExitDecoratorList(ctx *DecoratorListContext) {}

// EnterDecorator is called when production decorator is entered.
func (s *BaseTypeScriptParserListener) EnterDecorator(ctx *DecoratorContext) {}

// ExitDecorator is called when production decorator is exited.
func (s *BaseTypeScriptParserListener) ExitDecorator(ctx *DecoratorContext) {}

// EnterDecoratorMemberExpression is called when production decoratorMemberExpression is entered.
func (s *BaseTypeScriptParserListener) EnterDecoratorMemberExpression(ctx *DecoratorMemberExpressionContext) {
}

// ExitDecoratorMemberExpression is called when production decoratorMemberExpression is exited.
func (s *BaseTypeScriptParserListener) ExitDecoratorMemberExpression(ctx *DecoratorMemberExpressionContext) {
}

// EnterDecoratorCallExpression is called when production decoratorCallExpression is entered.
func (s *BaseTypeScriptParserListener) EnterDecoratorCallExpression(ctx *DecoratorCallExpressionContext) {
}

// ExitDecoratorCallExpression is called when production decoratorCallExpression is exited.
func (s *BaseTypeScriptParserListener) ExitDecoratorCallExpression(ctx *DecoratorCallExpressionContext) {
}

// EnterProgram is called when production program is entered.
func (s *BaseTypeScriptParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseTypeScriptParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterSourceElement is called when production sourceElement is entered.
func (s *BaseTypeScriptParserListener) EnterSourceElement(ctx *SourceElementContext) {}

// ExitSourceElement is called when production sourceElement is exited.
func (s *BaseTypeScriptParserListener) ExitSourceElement(ctx *SourceElementContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseTypeScriptParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseTypeScriptParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseTypeScriptParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseTypeScriptParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseTypeScriptParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseTypeScriptParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterAbstractDeclaration is called when production abstractDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterAbstractDeclaration(ctx *AbstractDeclarationContext) {}

// ExitAbstractDeclaration is called when production abstractDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitAbstractDeclaration(ctx *AbstractDeclarationContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseTypeScriptParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseTypeScriptParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterImportFromBlock is called when production importFromBlock is entered.
func (s *BaseTypeScriptParserListener) EnterImportFromBlock(ctx *ImportFromBlockContext) {}

// ExitImportFromBlock is called when production importFromBlock is exited.
func (s *BaseTypeScriptParserListener) ExitImportFromBlock(ctx *ImportFromBlockContext) {}

// EnterMultipleImportStatement is called when production multipleImportStatement is entered.
func (s *BaseTypeScriptParserListener) EnterMultipleImportStatement(ctx *MultipleImportStatementContext) {
}

// ExitMultipleImportStatement is called when production multipleImportStatement is exited.
func (s *BaseTypeScriptParserListener) ExitMultipleImportStatement(ctx *MultipleImportStatementContext) {
}

// EnterExportStatement is called when production exportStatement is entered.
func (s *BaseTypeScriptParserListener) EnterExportStatement(ctx *ExportStatementContext) {}

// ExitExportStatement is called when production exportStatement is exited.
func (s *BaseTypeScriptParserListener) ExitExportStatement(ctx *ExportStatementContext) {}

// EnterVariableStatement is called when production variableStatement is entered.
func (s *BaseTypeScriptParserListener) EnterVariableStatement(ctx *VariableStatementContext) {}

// ExitVariableStatement is called when production variableStatement is exited.
func (s *BaseTypeScriptParserListener) ExitVariableStatement(ctx *VariableStatementContext) {}

// EnterVariableDeclarationList is called when production variableDeclarationList is entered.
func (s *BaseTypeScriptParserListener) EnterVariableDeclarationList(ctx *VariableDeclarationListContext) {
}

// ExitVariableDeclarationList is called when production variableDeclarationList is exited.
func (s *BaseTypeScriptParserListener) ExitVariableDeclarationList(ctx *VariableDeclarationListContext) {
}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterEmptyStatement_ is called when production emptyStatement_ is entered.
func (s *BaseTypeScriptParserListener) EnterEmptyStatement_(ctx *EmptyStatement_Context) {}

// ExitEmptyStatement_ is called when production emptyStatement_ is exited.
func (s *BaseTypeScriptParserListener) ExitEmptyStatement_(ctx *EmptyStatement_Context) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BaseTypeScriptParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BaseTypeScriptParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseTypeScriptParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseTypeScriptParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterDoStatement is called when production DoStatement is entered.
func (s *BaseTypeScriptParserListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production DoStatement is exited.
func (s *BaseTypeScriptParserListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterWhileStatement is called when production WhileStatement is entered.
func (s *BaseTypeScriptParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production WhileStatement is exited.
func (s *BaseTypeScriptParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production ForStatement is entered.
func (s *BaseTypeScriptParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production ForStatement is exited.
func (s *BaseTypeScriptParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForVarStatement is called when production ForVarStatement is entered.
func (s *BaseTypeScriptParserListener) EnterForVarStatement(ctx *ForVarStatementContext) {}

// ExitForVarStatement is called when production ForVarStatement is exited.
func (s *BaseTypeScriptParserListener) ExitForVarStatement(ctx *ForVarStatementContext) {}

// EnterForInStatement is called when production ForInStatement is entered.
func (s *BaseTypeScriptParserListener) EnterForInStatement(ctx *ForInStatementContext) {}

// ExitForInStatement is called when production ForInStatement is exited.
func (s *BaseTypeScriptParserListener) ExitForInStatement(ctx *ForInStatementContext) {}

// EnterForVarInStatement is called when production ForVarInStatement is entered.
func (s *BaseTypeScriptParserListener) EnterForVarInStatement(ctx *ForVarInStatementContext) {}

// ExitForVarInStatement is called when production ForVarInStatement is exited.
func (s *BaseTypeScriptParserListener) ExitForVarInStatement(ctx *ForVarInStatementContext) {}

// EnterVarModifier is called when production varModifier is entered.
func (s *BaseTypeScriptParserListener) EnterVarModifier(ctx *VarModifierContext) {}

// ExitVarModifier is called when production varModifier is exited.
func (s *BaseTypeScriptParserListener) ExitVarModifier(ctx *VarModifierContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseTypeScriptParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseTypeScriptParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseTypeScriptParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseTypeScriptParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseTypeScriptParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseTypeScriptParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterYieldStatement is called when production yieldStatement is entered.
func (s *BaseTypeScriptParserListener) EnterYieldStatement(ctx *YieldStatementContext) {}

// ExitYieldStatement is called when production yieldStatement is exited.
func (s *BaseTypeScriptParserListener) ExitYieldStatement(ctx *YieldStatementContext) {}

// EnterWithStatement is called when production withStatement is entered.
func (s *BaseTypeScriptParserListener) EnterWithStatement(ctx *WithStatementContext) {}

// ExitWithStatement is called when production withStatement is exited.
func (s *BaseTypeScriptParserListener) ExitWithStatement(ctx *WithStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseTypeScriptParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseTypeScriptParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseTypeScriptParserListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseTypeScriptParserListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterCaseClauses is called when production caseClauses is entered.
func (s *BaseTypeScriptParserListener) EnterCaseClauses(ctx *CaseClausesContext) {}

// ExitCaseClauses is called when production caseClauses is exited.
func (s *BaseTypeScriptParserListener) ExitCaseClauses(ctx *CaseClausesContext) {}

// EnterCaseClause is called when production caseClause is entered.
func (s *BaseTypeScriptParserListener) EnterCaseClause(ctx *CaseClauseContext) {}

// ExitCaseClause is called when production caseClause is exited.
func (s *BaseTypeScriptParserListener) ExitCaseClause(ctx *CaseClauseContext) {}

// EnterDefaultClause is called when production defaultClause is entered.
func (s *BaseTypeScriptParserListener) EnterDefaultClause(ctx *DefaultClauseContext) {}

// ExitDefaultClause is called when production defaultClause is exited.
func (s *BaseTypeScriptParserListener) ExitDefaultClause(ctx *DefaultClauseContext) {}

// EnterLabelledStatement is called when production labelledStatement is entered.
func (s *BaseTypeScriptParserListener) EnterLabelledStatement(ctx *LabelledStatementContext) {}

// ExitLabelledStatement is called when production labelledStatement is exited.
func (s *BaseTypeScriptParserListener) ExitLabelledStatement(ctx *LabelledStatementContext) {}

// EnterThrowStatement is called when production throwStatement is entered.
func (s *BaseTypeScriptParserListener) EnterThrowStatement(ctx *ThrowStatementContext) {}

// ExitThrowStatement is called when production throwStatement is exited.
func (s *BaseTypeScriptParserListener) ExitThrowStatement(ctx *ThrowStatementContext) {}

// EnterTryStatement is called when production tryStatement is entered.
func (s *BaseTypeScriptParserListener) EnterTryStatement(ctx *TryStatementContext) {}

// ExitTryStatement is called when production tryStatement is exited.
func (s *BaseTypeScriptParserListener) ExitTryStatement(ctx *TryStatementContext) {}

// EnterCatchProduction is called when production catchProduction is entered.
func (s *BaseTypeScriptParserListener) EnterCatchProduction(ctx *CatchProductionContext) {}

// ExitCatchProduction is called when production catchProduction is exited.
func (s *BaseTypeScriptParserListener) ExitCatchProduction(ctx *CatchProductionContext) {}

// EnterAssignable is called when production assignable is entered.
func (s *BaseTypeScriptParserListener) EnterAssignable(ctx *AssignableContext) {}

// ExitAssignable is called when production assignable is exited.
func (s *BaseTypeScriptParserListener) ExitAssignable(ctx *AssignableContext) {}

// EnterFinallyProduction is called when production finallyProduction is entered.
func (s *BaseTypeScriptParserListener) EnterFinallyProduction(ctx *FinallyProductionContext) {}

// ExitFinallyProduction is called when production finallyProduction is exited.
func (s *BaseTypeScriptParserListener) ExitFinallyProduction(ctx *FinallyProductionContext) {}

// EnterDebuggerStatement is called when production debuggerStatement is entered.
func (s *BaseTypeScriptParserListener) EnterDebuggerStatement(ctx *DebuggerStatementContext) {}

// ExitDebuggerStatement is called when production debuggerStatement is exited.
func (s *BaseTypeScriptParserListener) ExitDebuggerStatement(ctx *DebuggerStatementContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassHeritage is called when production classHeritage is entered.
func (s *BaseTypeScriptParserListener) EnterClassHeritage(ctx *ClassHeritageContext) {}

// ExitClassHeritage is called when production classHeritage is exited.
func (s *BaseTypeScriptParserListener) ExitClassHeritage(ctx *ClassHeritageContext) {}

// EnterClassTail is called when production classTail is entered.
func (s *BaseTypeScriptParserListener) EnterClassTail(ctx *ClassTailContext) {}

// ExitClassTail is called when production classTail is exited.
func (s *BaseTypeScriptParserListener) ExitClassTail(ctx *ClassTailContext) {}

// EnterClassExtendsClause is called when production classExtendsClause is entered.
func (s *BaseTypeScriptParserListener) EnterClassExtendsClause(ctx *ClassExtendsClauseContext) {}

// ExitClassExtendsClause is called when production classExtendsClause is exited.
func (s *BaseTypeScriptParserListener) ExitClassExtendsClause(ctx *ClassExtendsClauseContext) {}

// EnterImplementsClause is called when production implementsClause is entered.
func (s *BaseTypeScriptParserListener) EnterImplementsClause(ctx *ImplementsClauseContext) {}

// ExitImplementsClause is called when production implementsClause is exited.
func (s *BaseTypeScriptParserListener) ExitImplementsClause(ctx *ImplementsClauseContext) {}

// EnterClassElement is called when production classElement is entered.
func (s *BaseTypeScriptParserListener) EnterClassElement(ctx *ClassElementContext) {}

// ExitClassElement is called when production classElement is exited.
func (s *BaseTypeScriptParserListener) ExitClassElement(ctx *ClassElementContext) {}

// EnterPropertyMemberDeclaration is called when production propertyMemberDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterPropertyMemberDeclaration(ctx *PropertyMemberDeclarationContext) {
}

// ExitPropertyMemberDeclaration is called when production propertyMemberDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitPropertyMemberDeclaration(ctx *PropertyMemberDeclarationContext) {
}

// EnterPropertyMemberBase is called when production propertyMemberBase is entered.
func (s *BaseTypeScriptParserListener) EnterPropertyMemberBase(ctx *PropertyMemberBaseContext) {}

// ExitPropertyMemberBase is called when production propertyMemberBase is exited.
func (s *BaseTypeScriptParserListener) ExitPropertyMemberBase(ctx *PropertyMemberBaseContext) {}

// EnterIndexMemberDeclaration is called when production indexMemberDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterIndexMemberDeclaration(ctx *IndexMemberDeclarationContext) {
}

// ExitIndexMemberDeclaration is called when production indexMemberDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitIndexMemberDeclaration(ctx *IndexMemberDeclarationContext) {
}

// EnterGeneratorMethod is called when production generatorMethod is entered.
func (s *BaseTypeScriptParserListener) EnterGeneratorMethod(ctx *GeneratorMethodContext) {}

// ExitGeneratorMethod is called when production generatorMethod is exited.
func (s *BaseTypeScriptParserListener) ExitGeneratorMethod(ctx *GeneratorMethodContext) {}

// EnterGeneratorFunctionDeclaration is called when production generatorFunctionDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterGeneratorFunctionDeclaration(ctx *GeneratorFunctionDeclarationContext) {
}

// ExitGeneratorFunctionDeclaration is called when production generatorFunctionDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitGeneratorFunctionDeclaration(ctx *GeneratorFunctionDeclarationContext) {
}

// EnterGeneratorBlock is called when production generatorBlock is entered.
func (s *BaseTypeScriptParserListener) EnterGeneratorBlock(ctx *GeneratorBlockContext) {}

// ExitGeneratorBlock is called when production generatorBlock is exited.
func (s *BaseTypeScriptParserListener) ExitGeneratorBlock(ctx *GeneratorBlockContext) {}

// EnterGeneratorDefinition is called when production generatorDefinition is entered.
func (s *BaseTypeScriptParserListener) EnterGeneratorDefinition(ctx *GeneratorDefinitionContext) {}

// ExitGeneratorDefinition is called when production generatorDefinition is exited.
func (s *BaseTypeScriptParserListener) ExitGeneratorDefinition(ctx *GeneratorDefinitionContext) {}

// EnterIteratorBlock is called when production iteratorBlock is entered.
func (s *BaseTypeScriptParserListener) EnterIteratorBlock(ctx *IteratorBlockContext) {}

// ExitIteratorBlock is called when production iteratorBlock is exited.
func (s *BaseTypeScriptParserListener) ExitIteratorBlock(ctx *IteratorBlockContext) {}

// EnterIteratorDefinition is called when production iteratorDefinition is entered.
func (s *BaseTypeScriptParserListener) EnterIteratorDefinition(ctx *IteratorDefinitionContext) {}

// ExitIteratorDefinition is called when production iteratorDefinition is exited.
func (s *BaseTypeScriptParserListener) ExitIteratorDefinition(ctx *IteratorDefinitionContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseTypeScriptParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseTypeScriptParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterFormalParameterArg is called when production formalParameterArg is entered.
func (s *BaseTypeScriptParserListener) EnterFormalParameterArg(ctx *FormalParameterArgContext) {}

// ExitFormalParameterArg is called when production formalParameterArg is exited.
func (s *BaseTypeScriptParserListener) ExitFormalParameterArg(ctx *FormalParameterArgContext) {}

// EnterLastFormalParameterArg is called when production lastFormalParameterArg is entered.
func (s *BaseTypeScriptParserListener) EnterLastFormalParameterArg(ctx *LastFormalParameterArgContext) {
}

// ExitLastFormalParameterArg is called when production lastFormalParameterArg is exited.
func (s *BaseTypeScriptParserListener) ExitLastFormalParameterArg(ctx *LastFormalParameterArgContext) {
}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseTypeScriptParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseTypeScriptParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterSourceElements is called when production sourceElements is entered.
func (s *BaseTypeScriptParserListener) EnterSourceElements(ctx *SourceElementsContext) {}

// ExitSourceElements is called when production sourceElements is exited.
func (s *BaseTypeScriptParserListener) ExitSourceElements(ctx *SourceElementsContext) {}

// EnterArrayLiteral is called when production arrayLiteral is entered.
func (s *BaseTypeScriptParserListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

// ExitArrayLiteral is called when production arrayLiteral is exited.
func (s *BaseTypeScriptParserListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}

// EnterElementList is called when production elementList is entered.
func (s *BaseTypeScriptParserListener) EnterElementList(ctx *ElementListContext) {}

// ExitElementList is called when production elementList is exited.
func (s *BaseTypeScriptParserListener) ExitElementList(ctx *ElementListContext) {}

// EnterLastElement is called when production lastElement is entered.
func (s *BaseTypeScriptParserListener) EnterLastElement(ctx *LastElementContext) {}

// ExitLastElement is called when production lastElement is exited.
func (s *BaseTypeScriptParserListener) ExitLastElement(ctx *LastElementContext) {}

// EnterObjectLiteral is called when production objectLiteral is entered.
func (s *BaseTypeScriptParserListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

// ExitObjectLiteral is called when production objectLiteral is exited.
func (s *BaseTypeScriptParserListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

// EnterPropertyExpressionAssignment is called when production PropertyExpressionAssignment is entered.
func (s *BaseTypeScriptParserListener) EnterPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// ExitPropertyExpressionAssignment is called when production PropertyExpressionAssignment is exited.
func (s *BaseTypeScriptParserListener) ExitPropertyExpressionAssignment(ctx *PropertyExpressionAssignmentContext) {
}

// EnterComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is entered.
func (s *BaseTypeScriptParserListener) EnterComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// ExitComputedPropertyExpressionAssignment is called when production ComputedPropertyExpressionAssignment is exited.
func (s *BaseTypeScriptParserListener) ExitComputedPropertyExpressionAssignment(ctx *ComputedPropertyExpressionAssignmentContext) {
}

// EnterFunctionProperty is called when production FunctionProperty is entered.
func (s *BaseTypeScriptParserListener) EnterFunctionProperty(ctx *FunctionPropertyContext) {}

// ExitFunctionProperty is called when production FunctionProperty is exited.
func (s *BaseTypeScriptParserListener) ExitFunctionProperty(ctx *FunctionPropertyContext) {}

// EnterPropertyGetter is called when production PropertyGetter is entered.
func (s *BaseTypeScriptParserListener) EnterPropertyGetter(ctx *PropertyGetterContext) {}

// ExitPropertyGetter is called when production PropertyGetter is exited.
func (s *BaseTypeScriptParserListener) ExitPropertyGetter(ctx *PropertyGetterContext) {}

// EnterPropertySetter is called when production PropertySetter is entered.
func (s *BaseTypeScriptParserListener) EnterPropertySetter(ctx *PropertySetterContext) {}

// ExitPropertySetter is called when production PropertySetter is exited.
func (s *BaseTypeScriptParserListener) ExitPropertySetter(ctx *PropertySetterContext) {}

// EnterMethodProperty is called when production MethodProperty is entered.
func (s *BaseTypeScriptParserListener) EnterMethodProperty(ctx *MethodPropertyContext) {}

// ExitMethodProperty is called when production MethodProperty is exited.
func (s *BaseTypeScriptParserListener) ExitMethodProperty(ctx *MethodPropertyContext) {}

// EnterPropertyShorthand is called when production PropertyShorthand is entered.
func (s *BaseTypeScriptParserListener) EnterPropertyShorthand(ctx *PropertyShorthandContext) {}

// ExitPropertyShorthand is called when production PropertyShorthand is exited.
func (s *BaseTypeScriptParserListener) ExitPropertyShorthand(ctx *PropertyShorthandContext) {}

// EnterRestParameterInObject is called when production RestParameterInObject is entered.
func (s *BaseTypeScriptParserListener) EnterRestParameterInObject(ctx *RestParameterInObjectContext) {
}

// ExitRestParameterInObject is called when production RestParameterInObject is exited.
func (s *BaseTypeScriptParserListener) ExitRestParameterInObject(ctx *RestParameterInObjectContext) {}

// EnterGetAccessor is called when production getAccessor is entered.
func (s *BaseTypeScriptParserListener) EnterGetAccessor(ctx *GetAccessorContext) {}

// ExitGetAccessor is called when production getAccessor is exited.
func (s *BaseTypeScriptParserListener) ExitGetAccessor(ctx *GetAccessorContext) {}

// EnterSetAccessor is called when production setAccessor is entered.
func (s *BaseTypeScriptParserListener) EnterSetAccessor(ctx *SetAccessorContext) {}

// ExitSetAccessor is called when production setAccessor is exited.
func (s *BaseTypeScriptParserListener) ExitSetAccessor(ctx *SetAccessorContext) {}

// EnterPropertyName is called when production propertyName is entered.
func (s *BaseTypeScriptParserListener) EnterPropertyName(ctx *PropertyNameContext) {}

// ExitPropertyName is called when production propertyName is exited.
func (s *BaseTypeScriptParserListener) ExitPropertyName(ctx *PropertyNameContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseTypeScriptParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseTypeScriptParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterLastArgument is called when production lastArgument is entered.
func (s *BaseTypeScriptParserListener) EnterLastArgument(ctx *LastArgumentContext) {}

// ExitLastArgument is called when production lastArgument is exited.
func (s *BaseTypeScriptParserListener) ExitLastArgument(ctx *LastArgumentContext) {}

// EnterExpressionSequence is called when production expressionSequence is entered.
func (s *BaseTypeScriptParserListener) EnterExpressionSequence(ctx *ExpressionSequenceContext) {}

// ExitExpressionSequence is called when production expressionSequence is exited.
func (s *BaseTypeScriptParserListener) ExitExpressionSequence(ctx *ExpressionSequenceContext) {}

// EnterFunctionExpressionDeclaration is called when production functionExpressionDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterFunctionExpressionDeclaration(ctx *FunctionExpressionDeclarationContext) {
}

// ExitFunctionExpressionDeclaration is called when production functionExpressionDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitFunctionExpressionDeclaration(ctx *FunctionExpressionDeclarationContext) {
}

// EnterTemplateStringExpression is called when production TemplateStringExpression is entered.
func (s *BaseTypeScriptParserListener) EnterTemplateStringExpression(ctx *TemplateStringExpressionContext) {
}

// ExitTemplateStringExpression is called when production TemplateStringExpression is exited.
func (s *BaseTypeScriptParserListener) ExitTemplateStringExpression(ctx *TemplateStringExpressionContext) {
}

// EnterGeneratorsExpression is called when production GeneratorsExpression is entered.
func (s *BaseTypeScriptParserListener) EnterGeneratorsExpression(ctx *GeneratorsExpressionContext) {}

// ExitGeneratorsExpression is called when production GeneratorsExpression is exited.
func (s *BaseTypeScriptParserListener) ExitGeneratorsExpression(ctx *GeneratorsExpressionContext) {}

// EnterPowerExpression is called when production PowerExpression is entered.
func (s *BaseTypeScriptParserListener) EnterPowerExpression(ctx *PowerExpressionContext) {}

// ExitPowerExpression is called when production PowerExpression is exited.
func (s *BaseTypeScriptParserListener) ExitPowerExpression(ctx *PowerExpressionContext) {}

// EnterInExpression is called when production InExpression is entered.
func (s *BaseTypeScriptParserListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production InExpression is exited.
func (s *BaseTypeScriptParserListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterGenericTypes is called when production GenericTypes is entered.
func (s *BaseTypeScriptParserListener) EnterGenericTypes(ctx *GenericTypesContext) {}

// ExitGenericTypes is called when production GenericTypes is exited.
func (s *BaseTypeScriptParserListener) ExitGenericTypes(ctx *GenericTypesContext) {}

// EnterArgumentsExpression is called when production ArgumentsExpression is entered.
func (s *BaseTypeScriptParserListener) EnterArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// ExitArgumentsExpression is called when production ArgumentsExpression is exited.
func (s *BaseTypeScriptParserListener) ExitArgumentsExpression(ctx *ArgumentsExpressionContext) {}

// EnterThisExpression is called when production ThisExpression is entered.
func (s *BaseTypeScriptParserListener) EnterThisExpression(ctx *ThisExpressionContext) {}

// ExitThisExpression is called when production ThisExpression is exited.
func (s *BaseTypeScriptParserListener) ExitThisExpression(ctx *ThisExpressionContext) {}

// EnterTypeofExpression is called when production TypeofExpression is entered.
func (s *BaseTypeScriptParserListener) EnterTypeofExpression(ctx *TypeofExpressionContext) {}

// ExitTypeofExpression is called when production TypeofExpression is exited.
func (s *BaseTypeScriptParserListener) ExitTypeofExpression(ctx *TypeofExpressionContext) {}

// EnterGeneratorsFunctionExpression is called when production GeneratorsFunctionExpression is entered.
func (s *BaseTypeScriptParserListener) EnterGeneratorsFunctionExpression(ctx *GeneratorsFunctionExpressionContext) {
}

// ExitGeneratorsFunctionExpression is called when production GeneratorsFunctionExpression is exited.
func (s *BaseTypeScriptParserListener) ExitGeneratorsFunctionExpression(ctx *GeneratorsFunctionExpressionContext) {
}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseTypeScriptParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseTypeScriptParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterBitXOrExpression is called when production BitXOrExpression is entered.
func (s *BaseTypeScriptParserListener) EnterBitXOrExpression(ctx *BitXOrExpressionContext) {}

// ExitBitXOrExpression is called when production BitXOrExpression is exited.
func (s *BaseTypeScriptParserListener) ExitBitXOrExpression(ctx *BitXOrExpressionContext) {}

// EnterMultiplicativeExpression is called when production MultiplicativeExpression is entered.
func (s *BaseTypeScriptParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production MultiplicativeExpression is exited.
func (s *BaseTypeScriptParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterBitShiftExpression is called when production BitShiftExpression is entered.
func (s *BaseTypeScriptParserListener) EnterBitShiftExpression(ctx *BitShiftExpressionContext) {}

// ExitBitShiftExpression is called when production BitShiftExpression is exited.
func (s *BaseTypeScriptParserListener) ExitBitShiftExpression(ctx *BitShiftExpressionContext) {}

// EnterAdditiveExpression is called when production AdditiveExpression is entered.
func (s *BaseTypeScriptParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production AdditiveExpression is exited.
func (s *BaseTypeScriptParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterRelationalExpression is called when production RelationalExpression is entered.
func (s *BaseTypeScriptParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production RelationalExpression is exited.
func (s *BaseTypeScriptParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterBitNotExpression is called when production BitNotExpression is entered.
func (s *BaseTypeScriptParserListener) EnterBitNotExpression(ctx *BitNotExpressionContext) {}

// ExitBitNotExpression is called when production BitNotExpression is exited.
func (s *BaseTypeScriptParserListener) ExitBitNotExpression(ctx *BitNotExpressionContext) {}

// EnterNewExpression is called when production NewExpression is entered.
func (s *BaseTypeScriptParserListener) EnterNewExpression(ctx *NewExpressionContext) {}

// ExitNewExpression is called when production NewExpression is exited.
func (s *BaseTypeScriptParserListener) ExitNewExpression(ctx *NewExpressionContext) {}

// EnterLiteralExpression is called when production LiteralExpression is entered.
func (s *BaseTypeScriptParserListener) EnterLiteralExpression(ctx *LiteralExpressionContext) {}

// ExitLiteralExpression is called when production LiteralExpression is exited.
func (s *BaseTypeScriptParserListener) ExitLiteralExpression(ctx *LiteralExpressionContext) {}

// EnterArrayLiteralExpression is called when production ArrayLiteralExpression is entered.
func (s *BaseTypeScriptParserListener) EnterArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {
}

// ExitArrayLiteralExpression is called when production ArrayLiteralExpression is exited.
func (s *BaseTypeScriptParserListener) ExitArrayLiteralExpression(ctx *ArrayLiteralExpressionContext) {
}

// EnterMemberDotExpression is called when production MemberDotExpression is entered.
func (s *BaseTypeScriptParserListener) EnterMemberDotExpression(ctx *MemberDotExpressionContext) {}

// ExitMemberDotExpression is called when production MemberDotExpression is exited.
func (s *BaseTypeScriptParserListener) ExitMemberDotExpression(ctx *MemberDotExpressionContext) {}

// EnterMemberIndexExpression is called when production MemberIndexExpression is entered.
func (s *BaseTypeScriptParserListener) EnterMemberIndexExpression(ctx *MemberIndexExpressionContext) {
}

// ExitMemberIndexExpression is called when production MemberIndexExpression is exited.
func (s *BaseTypeScriptParserListener) ExitMemberIndexExpression(ctx *MemberIndexExpressionContext) {}

// EnterBitAndExpression is called when production BitAndExpression is entered.
func (s *BaseTypeScriptParserListener) EnterBitAndExpression(ctx *BitAndExpressionContext) {}

// ExitBitAndExpression is called when production BitAndExpression is exited.
func (s *BaseTypeScriptParserListener) ExitBitAndExpression(ctx *BitAndExpressionContext) {}

// EnterBitOrExpression is called when production BitOrExpression is entered.
func (s *BaseTypeScriptParserListener) EnterBitOrExpression(ctx *BitOrExpressionContext) {}

// ExitBitOrExpression is called when production BitOrExpression is exited.
func (s *BaseTypeScriptParserListener) ExitBitOrExpression(ctx *BitOrExpressionContext) {}

// EnterAssignmentOperatorExpression is called when production AssignmentOperatorExpression is entered.
func (s *BaseTypeScriptParserListener) EnterAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// ExitAssignmentOperatorExpression is called when production AssignmentOperatorExpression is exited.
func (s *BaseTypeScriptParserListener) ExitAssignmentOperatorExpression(ctx *AssignmentOperatorExpressionContext) {
}

// EnterVoidExpression is called when production VoidExpression is entered.
func (s *BaseTypeScriptParserListener) EnterVoidExpression(ctx *VoidExpressionContext) {}

// ExitVoidExpression is called when production VoidExpression is exited.
func (s *BaseTypeScriptParserListener) ExitVoidExpression(ctx *VoidExpressionContext) {}

// EnterTernaryExpression is called when production TernaryExpression is entered.
func (s *BaseTypeScriptParserListener) EnterTernaryExpression(ctx *TernaryExpressionContext) {}

// ExitTernaryExpression is called when production TernaryExpression is exited.
func (s *BaseTypeScriptParserListener) ExitTernaryExpression(ctx *TernaryExpressionContext) {}

// EnterLogicalAndExpression is called when production LogicalAndExpression is entered.
func (s *BaseTypeScriptParserListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production LogicalAndExpression is exited.
func (s *BaseTypeScriptParserListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterPreIncrementExpression is called when production PreIncrementExpression is entered.
func (s *BaseTypeScriptParserListener) EnterPreIncrementExpression(ctx *PreIncrementExpressionContext) {
}

// ExitPreIncrementExpression is called when production PreIncrementExpression is exited.
func (s *BaseTypeScriptParserListener) ExitPreIncrementExpression(ctx *PreIncrementExpressionContext) {
}

// EnterObjectLiteralExpression is called when production ObjectLiteralExpression is entered.
func (s *BaseTypeScriptParserListener) EnterObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {
}

// ExitObjectLiteralExpression is called when production ObjectLiteralExpression is exited.
func (s *BaseTypeScriptParserListener) ExitObjectLiteralExpression(ctx *ObjectLiteralExpressionContext) {
}

// EnterLogicalOrExpression is called when production LogicalOrExpression is entered.
func (s *BaseTypeScriptParserListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production LogicalOrExpression is exited.
func (s *BaseTypeScriptParserListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterNotExpression is called when production NotExpression is entered.
func (s *BaseTypeScriptParserListener) EnterNotExpression(ctx *NotExpressionContext) {}

// ExitNotExpression is called when production NotExpression is exited.
func (s *BaseTypeScriptParserListener) ExitNotExpression(ctx *NotExpressionContext) {}

// EnterPreDecreaseExpression is called when production PreDecreaseExpression is entered.
func (s *BaseTypeScriptParserListener) EnterPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {
}

// ExitPreDecreaseExpression is called when production PreDecreaseExpression is exited.
func (s *BaseTypeScriptParserListener) ExitPreDecreaseExpression(ctx *PreDecreaseExpressionContext) {}

// EnterAwaitExpression is called when production AwaitExpression is entered.
func (s *BaseTypeScriptParserListener) EnterAwaitExpression(ctx *AwaitExpressionContext) {}

// ExitAwaitExpression is called when production AwaitExpression is exited.
func (s *BaseTypeScriptParserListener) ExitAwaitExpression(ctx *AwaitExpressionContext) {}

// EnterFunctionExpression is called when production FunctionExpression is entered.
func (s *BaseTypeScriptParserListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {}

// ExitFunctionExpression is called when production FunctionExpression is exited.
func (s *BaseTypeScriptParserListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterUnaryMinusExpression is called when production UnaryMinusExpression is entered.
func (s *BaseTypeScriptParserListener) EnterUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// ExitUnaryMinusExpression is called when production UnaryMinusExpression is exited.
func (s *BaseTypeScriptParserListener) ExitUnaryMinusExpression(ctx *UnaryMinusExpressionContext) {}

// EnterAssignmentExpression is called when production AssignmentExpression is entered.
func (s *BaseTypeScriptParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production AssignmentExpression is exited.
func (s *BaseTypeScriptParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterPostDecreaseExpression is called when production PostDecreaseExpression is entered.
func (s *BaseTypeScriptParserListener) EnterPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {
}

// ExitPostDecreaseExpression is called when production PostDecreaseExpression is exited.
func (s *BaseTypeScriptParserListener) ExitPostDecreaseExpression(ctx *PostDecreaseExpressionContext) {
}

// EnterInstanceofExpression is called when production InstanceofExpression is entered.
func (s *BaseTypeScriptParserListener) EnterInstanceofExpression(ctx *InstanceofExpressionContext) {}

// ExitInstanceofExpression is called when production InstanceofExpression is exited.
func (s *BaseTypeScriptParserListener) ExitInstanceofExpression(ctx *InstanceofExpressionContext) {}

// EnterUnaryPlusExpression is called when production UnaryPlusExpression is entered.
func (s *BaseTypeScriptParserListener) EnterUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// ExitUnaryPlusExpression is called when production UnaryPlusExpression is exited.
func (s *BaseTypeScriptParserListener) ExitUnaryPlusExpression(ctx *UnaryPlusExpressionContext) {}

// EnterDeleteExpression is called when production DeleteExpression is entered.
func (s *BaseTypeScriptParserListener) EnterDeleteExpression(ctx *DeleteExpressionContext) {}

// ExitDeleteExpression is called when production DeleteExpression is exited.
func (s *BaseTypeScriptParserListener) ExitDeleteExpression(ctx *DeleteExpressionContext) {}

// EnterArrowFunctionExpression is called when production ArrowFunctionExpression is entered.
func (s *BaseTypeScriptParserListener) EnterArrowFunctionExpression(ctx *ArrowFunctionExpressionContext) {
}

// ExitArrowFunctionExpression is called when production ArrowFunctionExpression is exited.
func (s *BaseTypeScriptParserListener) ExitArrowFunctionExpression(ctx *ArrowFunctionExpressionContext) {
}

// EnterIteratorsExpression is called when production IteratorsExpression is entered.
func (s *BaseTypeScriptParserListener) EnterIteratorsExpression(ctx *IteratorsExpressionContext) {}

// ExitIteratorsExpression is called when production IteratorsExpression is exited.
func (s *BaseTypeScriptParserListener) ExitIteratorsExpression(ctx *IteratorsExpressionContext) {}

// EnterSuperExpression is called when production SuperExpression is entered.
func (s *BaseTypeScriptParserListener) EnterSuperExpression(ctx *SuperExpressionContext) {}

// ExitSuperExpression is called when production SuperExpression is exited.
func (s *BaseTypeScriptParserListener) ExitSuperExpression(ctx *SuperExpressionContext) {}

// EnterParenthesizedExpression is called when production ParenthesizedExpression is entered.
func (s *BaseTypeScriptParserListener) EnterParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// ExitParenthesizedExpression is called when production ParenthesizedExpression is exited.
func (s *BaseTypeScriptParserListener) ExitParenthesizedExpression(ctx *ParenthesizedExpressionContext) {
}

// EnterPostIncrementExpression is called when production PostIncrementExpression is entered.
func (s *BaseTypeScriptParserListener) EnterPostIncrementExpression(ctx *PostIncrementExpressionContext) {
}

// ExitPostIncrementExpression is called when production PostIncrementExpression is exited.
func (s *BaseTypeScriptParserListener) ExitPostIncrementExpression(ctx *PostIncrementExpressionContext) {
}

// EnterYieldExpression is called when production YieldExpression is entered.
func (s *BaseTypeScriptParserListener) EnterYieldExpression(ctx *YieldExpressionContext) {}

// ExitYieldExpression is called when production YieldExpression is exited.
func (s *BaseTypeScriptParserListener) ExitYieldExpression(ctx *YieldExpressionContext) {}

// EnterClassExpression is called when production ClassExpression is entered.
func (s *BaseTypeScriptParserListener) EnterClassExpression(ctx *ClassExpressionContext) {}

// ExitClassExpression is called when production ClassExpression is exited.
func (s *BaseTypeScriptParserListener) ExitClassExpression(ctx *ClassExpressionContext) {}

// EnterIdentifierExpression is called when production IdentifierExpression is entered.
func (s *BaseTypeScriptParserListener) EnterIdentifierExpression(ctx *IdentifierExpressionContext) {}

// ExitIdentifierExpression is called when production IdentifierExpression is exited.
func (s *BaseTypeScriptParserListener) ExitIdentifierExpression(ctx *IdentifierExpressionContext) {}

// EnterCoalesceExpression is called when production CoalesceExpression is entered.
func (s *BaseTypeScriptParserListener) EnterCoalesceExpression(ctx *CoalesceExpressionContext) {}

// ExitCoalesceExpression is called when production CoalesceExpression is exited.
func (s *BaseTypeScriptParserListener) ExitCoalesceExpression(ctx *CoalesceExpressionContext) {}

// EnterArrowFunctionDeclaration is called when production arrowFunctionDeclaration is entered.
func (s *BaseTypeScriptParserListener) EnterArrowFunctionDeclaration(ctx *ArrowFunctionDeclarationContext) {
}

// ExitArrowFunctionDeclaration is called when production arrowFunctionDeclaration is exited.
func (s *BaseTypeScriptParserListener) ExitArrowFunctionDeclaration(ctx *ArrowFunctionDeclarationContext) {
}

// EnterArrowFunctionParameters is called when production arrowFunctionParameters is entered.
func (s *BaseTypeScriptParserListener) EnterArrowFunctionParameters(ctx *ArrowFunctionParametersContext) {
}

// ExitArrowFunctionParameters is called when production arrowFunctionParameters is exited.
func (s *BaseTypeScriptParserListener) ExitArrowFunctionParameters(ctx *ArrowFunctionParametersContext) {
}

// EnterArrowFunctionBody is called when production arrowFunctionBody is entered.
func (s *BaseTypeScriptParserListener) EnterArrowFunctionBody(ctx *ArrowFunctionBodyContext) {}

// ExitArrowFunctionBody is called when production arrowFunctionBody is exited.
func (s *BaseTypeScriptParserListener) ExitArrowFunctionBody(ctx *ArrowFunctionBodyContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseTypeScriptParserListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseTypeScriptParserListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseTypeScriptParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseTypeScriptParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseTypeScriptParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseTypeScriptParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterIdentifierName is called when production identifierName is entered.
func (s *BaseTypeScriptParserListener) EnterIdentifierName(ctx *IdentifierNameContext) {}

// ExitIdentifierName is called when production identifierName is exited.
func (s *BaseTypeScriptParserListener) ExitIdentifierName(ctx *IdentifierNameContext) {}

// EnterReservedWord is called when production reservedWord is entered.
func (s *BaseTypeScriptParserListener) EnterReservedWord(ctx *ReservedWordContext) {}

// ExitReservedWord is called when production reservedWord is exited.
func (s *BaseTypeScriptParserListener) ExitReservedWord(ctx *ReservedWordContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseTypeScriptParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseTypeScriptParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterGetter is called when production getter is entered.
func (s *BaseTypeScriptParserListener) EnterGetter(ctx *GetterContext) {}

// ExitGetter is called when production getter is exited.
func (s *BaseTypeScriptParserListener) ExitGetter(ctx *GetterContext) {}

// EnterSetter is called when production setter is entered.
func (s *BaseTypeScriptParserListener) EnterSetter(ctx *SetterContext) {}

// ExitSetter is called when production setter is exited.
func (s *BaseTypeScriptParserListener) ExitSetter(ctx *SetterContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseTypeScriptParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseTypeScriptParserListener) ExitEos(ctx *EosContext) {}
