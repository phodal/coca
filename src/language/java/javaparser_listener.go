// Code generated from JavaParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // JavaParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JavaParserListener is a complete listener for a parse tree produced by JavaParser.
type JavaParserListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterPackageDeclaration is called when entering the packageDeclaration production.
	EnterPackageDeclaration(c *PackageDeclarationContext)

	// EnterImportDeclaration is called when entering the importDeclaration production.
	EnterImportDeclaration(c *ImportDeclarationContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterModifier is called when entering the modifier production.
	EnterModifier(c *ModifierContext)

	// EnterClassOrInterfaceModifier is called when entering the classOrInterfaceModifier production.
	EnterClassOrInterfaceModifier(c *ClassOrInterfaceModifierContext)

	// EnterVariableModifier is called when entering the variableModifier production.
	EnterVariableModifier(c *VariableModifierContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterTypeParameters is called when entering the typeParameters production.
	EnterTypeParameters(c *TypeParametersContext)

	// EnterTypeParameter is called when entering the typeParameter production.
	EnterTypeParameter(c *TypeParameterContext)

	// EnterTypeBound is called when entering the typeBound production.
	EnterTypeBound(c *TypeBoundContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumConstants is called when entering the enumConstants production.
	EnterEnumConstants(c *EnumConstantsContext)

	// EnterEnumConstant is called when entering the enumConstant production.
	EnterEnumConstant(c *EnumConstantContext)

	// EnterEnumBodyDeclarations is called when entering the enumBodyDeclarations production.
	EnterEnumBodyDeclarations(c *EnumBodyDeclarationsContext)

	// EnterInterfaceDeclaration is called when entering the interfaceDeclaration production.
	EnterInterfaceDeclaration(c *InterfaceDeclarationContext)

	// EnterClassBody is called when entering the classBody production.
	EnterClassBody(c *ClassBodyContext)

	// EnterInterfaceBody is called when entering the interfaceBody production.
	EnterInterfaceBody(c *InterfaceBodyContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterMemberDeclaration is called when entering the memberDeclaration production.
	EnterMemberDeclaration(c *MemberDeclarationContext)

	// EnterMethodDeclaration is called when entering the methodDeclaration production.
	EnterMethodDeclaration(c *MethodDeclarationContext)

	// EnterMethodBody is called when entering the methodBody production.
	EnterMethodBody(c *MethodBodyContext)

	// EnterTypeTypeOrVoid is called when entering the typeTypeOrVoid production.
	EnterTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// EnterGenericMethodDeclaration is called when entering the genericMethodDeclaration production.
	EnterGenericMethodDeclaration(c *GenericMethodDeclarationContext)

	// EnterGenericConstructorDeclaration is called when entering the genericConstructorDeclaration production.
	EnterGenericConstructorDeclaration(c *GenericConstructorDeclarationContext)

	// EnterConstructorDeclaration is called when entering the constructorDeclaration production.
	EnterConstructorDeclaration(c *ConstructorDeclarationContext)

	// EnterFieldDeclaration is called when entering the fieldDeclaration production.
	EnterFieldDeclaration(c *FieldDeclarationContext)

	// EnterInterfaceBodyDeclaration is called when entering the interfaceBodyDeclaration production.
	EnterInterfaceBodyDeclaration(c *InterfaceBodyDeclarationContext)

	// EnterInterfaceMemberDeclaration is called when entering the interfaceMemberDeclaration production.
	EnterInterfaceMemberDeclaration(c *InterfaceMemberDeclarationContext)

	// EnterConstDeclaration is called when entering the constDeclaration production.
	EnterConstDeclaration(c *ConstDeclarationContext)

	// EnterConstantDeclarator is called when entering the constantDeclarator production.
	EnterConstantDeclarator(c *ConstantDeclaratorContext)

	// EnterInterfaceMethodDeclaration is called when entering the interfaceMethodDeclaration production.
	EnterInterfaceMethodDeclaration(c *InterfaceMethodDeclarationContext)

	// EnterInterfaceMethodModifier is called when entering the interfaceMethodModifier production.
	EnterInterfaceMethodModifier(c *InterfaceMethodModifierContext)

	// EnterGenericInterfaceMethodDeclaration is called when entering the genericInterfaceMethodDeclaration production.
	EnterGenericInterfaceMethodDeclaration(c *GenericInterfaceMethodDeclarationContext)

	// EnterVariableDeclarators is called when entering the variableDeclarators production.
	EnterVariableDeclarators(c *VariableDeclaratorsContext)

	// EnterVariableDeclarator is called when entering the variableDeclarator production.
	EnterVariableDeclarator(c *VariableDeclaratorContext)

	// EnterVariableDeclaratorId is called when entering the variableDeclaratorId production.
	EnterVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// EnterVariableInitializer is called when entering the variableInitializer production.
	EnterVariableInitializer(c *VariableInitializerContext)

	// EnterArrayInitializer is called when entering the arrayInitializer production.
	EnterArrayInitializer(c *ArrayInitializerContext)

	// EnterClassOrInterfaceType is called when entering the classOrInterfaceType production.
	EnterClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// EnterTypeArgument is called when entering the typeArgument production.
	EnterTypeArgument(c *TypeArgumentContext)

	// EnterQualifiedNameList is called when entering the qualifiedNameList production.
	EnterQualifiedNameList(c *QualifiedNameListContext)

	// EnterFormalParameters is called when entering the formalParameters production.
	EnterFormalParameters(c *FormalParametersContext)

	// EnterFormalParameterList is called when entering the formalParameterList production.
	EnterFormalParameterList(c *FormalParameterListContext)

	// EnterFormalParameter is called when entering the formalParameter production.
	EnterFormalParameter(c *FormalParameterContext)

	// EnterLastFormalParameter is called when entering the lastFormalParameter production.
	EnterLastFormalParameter(c *LastFormalParameterContext)

	// EnterQualifiedName is called when entering the qualifiedName production.
	EnterQualifiedName(c *QualifiedNameContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIntegerLiteral is called when entering the integerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatLiteral is called when entering the floatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterElementValuePairs is called when entering the elementValuePairs production.
	EnterElementValuePairs(c *ElementValuePairsContext)

	// EnterElementValuePair is called when entering the elementValuePair production.
	EnterElementValuePair(c *ElementValuePairContext)

	// EnterElementValue is called when entering the elementValue production.
	EnterElementValue(c *ElementValueContext)

	// EnterElementValueArrayInitializer is called when entering the elementValueArrayInitializer production.
	EnterElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// EnterAnnotationTypeDeclaration is called when entering the annotationTypeDeclaration production.
	EnterAnnotationTypeDeclaration(c *AnnotationTypeDeclarationContext)

	// EnterAnnotationTypeBody is called when entering the annotationTypeBody production.
	EnterAnnotationTypeBody(c *AnnotationTypeBodyContext)

	// EnterAnnotationTypeElementDeclaration is called when entering the annotationTypeElementDeclaration production.
	EnterAnnotationTypeElementDeclaration(c *AnnotationTypeElementDeclarationContext)

	// EnterAnnotationTypeElementRest is called when entering the annotationTypeElementRest production.
	EnterAnnotationTypeElementRest(c *AnnotationTypeElementRestContext)

	// EnterAnnotationMethodOrConstantRest is called when entering the annotationMethodOrConstantRest production.
	EnterAnnotationMethodOrConstantRest(c *AnnotationMethodOrConstantRestContext)

	// EnterAnnotationMethodRest is called when entering the annotationMethodRest production.
	EnterAnnotationMethodRest(c *AnnotationMethodRestContext)

	// EnterAnnotationConstantRest is called when entering the annotationConstantRest production.
	EnterAnnotationConstantRest(c *AnnotationConstantRestContext)

	// EnterDefaultValue is called when entering the defaultValue production.
	EnterDefaultValue(c *DefaultValueContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBlockStatement is called when entering the blockStatement production.
	EnterBlockStatement(c *BlockStatementContext)

	// EnterLocalVariableDeclaration is called when entering the localVariableDeclaration production.
	EnterLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// EnterLocalTypeDeclaration is called when entering the localTypeDeclaration production.
	EnterLocalTypeDeclaration(c *LocalTypeDeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterCatchClause is called when entering the catchClause production.
	EnterCatchClause(c *CatchClauseContext)

	// EnterCatchType is called when entering the catchType production.
	EnterCatchType(c *CatchTypeContext)

	// EnterFinallyBlock is called when entering the finallyBlock production.
	EnterFinallyBlock(c *FinallyBlockContext)

	// EnterResourceSpecification is called when entering the resourceSpecification production.
	EnterResourceSpecification(c *ResourceSpecificationContext)

	// EnterResources is called when entering the resources production.
	EnterResources(c *ResourcesContext)

	// EnterResource is called when entering the resource production.
	EnterResource(c *ResourceContext)

	// EnterSwitchBlockStatementGroup is called when entering the switchBlockStatementGroup production.
	EnterSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// EnterSwitchLabel is called when entering the switchLabel production.
	EnterSwitchLabel(c *SwitchLabelContext)

	// EnterForControl is called when entering the forControl production.
	EnterForControl(c *ForControlContext)

	// EnterForInit is called when entering the forInit production.
	EnterForInit(c *ForInitContext)

	// EnterEnhancedForControl is called when entering the enhancedForControl production.
	EnterEnhancedForControl(c *EnhancedForControlContext)

	// EnterParExpression is called when entering the parExpression production.
	EnterParExpression(c *ParExpressionContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterMethodCall is called when entering the methodCall production.
	EnterMethodCall(c *MethodCallContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLambdaExpression is called when entering the lambdaExpression production.
	EnterLambdaExpression(c *LambdaExpressionContext)

	// EnterLambdaParameters is called when entering the lambdaParameters production.
	EnterLambdaParameters(c *LambdaParametersContext)

	// EnterLambdaBody is called when entering the lambdaBody production.
	EnterLambdaBody(c *LambdaBodyContext)

	// EnterPrimary is called when entering the primary production.
	EnterPrimary(c *PrimaryContext)

	// EnterClassType is called when entering the classType production.
	EnterClassType(c *ClassTypeContext)

	// EnterCreator is called when entering the creator production.
	EnterCreator(c *CreatorContext)

	// EnterCreatedName is called when entering the createdName production.
	EnterCreatedName(c *CreatedNameContext)

	// EnterInnerCreator is called when entering the innerCreator production.
	EnterInnerCreator(c *InnerCreatorContext)

	// EnterArrayCreatorRest is called when entering the arrayCreatorRest production.
	EnterArrayCreatorRest(c *ArrayCreatorRestContext)

	// EnterClassCreatorRest is called when entering the classCreatorRest production.
	EnterClassCreatorRest(c *ClassCreatorRestContext)

	// EnterExplicitGenericInvocation is called when entering the explicitGenericInvocation production.
	EnterExplicitGenericInvocation(c *ExplicitGenericInvocationContext)

	// EnterTypeArgumentsOrDiamond is called when entering the typeArgumentsOrDiamond production.
	EnterTypeArgumentsOrDiamond(c *TypeArgumentsOrDiamondContext)

	// EnterNonWildcardTypeArgumentsOrDiamond is called when entering the nonWildcardTypeArgumentsOrDiamond production.
	EnterNonWildcardTypeArgumentsOrDiamond(c *NonWildcardTypeArgumentsOrDiamondContext)

	// EnterNonWildcardTypeArguments is called when entering the nonWildcardTypeArguments production.
	EnterNonWildcardTypeArguments(c *NonWildcardTypeArgumentsContext)

	// EnterTypeList is called when entering the typeList production.
	EnterTypeList(c *TypeListContext)

	// EnterTypeType is called when entering the typeType production.
	EnterTypeType(c *TypeTypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterTypeArguments is called when entering the typeArguments production.
	EnterTypeArguments(c *TypeArgumentsContext)

	// EnterSuperSuffix is called when entering the superSuffix production.
	EnterSuperSuffix(c *SuperSuffixContext)

	// EnterExplicitGenericInvocationSuffix is called when entering the explicitGenericInvocationSuffix production.
	EnterExplicitGenericInvocationSuffix(c *ExplicitGenericInvocationSuffixContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitPackageDeclaration is called when exiting the packageDeclaration production.
	ExitPackageDeclaration(c *PackageDeclarationContext)

	// ExitImportDeclaration is called when exiting the importDeclaration production.
	ExitImportDeclaration(c *ImportDeclarationContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitModifier is called when exiting the modifier production.
	ExitModifier(c *ModifierContext)

	// ExitClassOrInterfaceModifier is called when exiting the classOrInterfaceModifier production.
	ExitClassOrInterfaceModifier(c *ClassOrInterfaceModifierContext)

	// ExitVariableModifier is called when exiting the variableModifier production.
	ExitVariableModifier(c *VariableModifierContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitTypeParameters is called when exiting the typeParameters production.
	ExitTypeParameters(c *TypeParametersContext)

	// ExitTypeParameter is called when exiting the typeParameter production.
	ExitTypeParameter(c *TypeParameterContext)

	// ExitTypeBound is called when exiting the typeBound production.
	ExitTypeBound(c *TypeBoundContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumConstants is called when exiting the enumConstants production.
	ExitEnumConstants(c *EnumConstantsContext)

	// ExitEnumConstant is called when exiting the enumConstant production.
	ExitEnumConstant(c *EnumConstantContext)

	// ExitEnumBodyDeclarations is called when exiting the enumBodyDeclarations production.
	ExitEnumBodyDeclarations(c *EnumBodyDeclarationsContext)

	// ExitInterfaceDeclaration is called when exiting the interfaceDeclaration production.
	ExitInterfaceDeclaration(c *InterfaceDeclarationContext)

	// ExitClassBody is called when exiting the classBody production.
	ExitClassBody(c *ClassBodyContext)

	// ExitInterfaceBody is called when exiting the interfaceBody production.
	ExitInterfaceBody(c *InterfaceBodyContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitMemberDeclaration is called when exiting the memberDeclaration production.
	ExitMemberDeclaration(c *MemberDeclarationContext)

	// ExitMethodDeclaration is called when exiting the methodDeclaration production.
	ExitMethodDeclaration(c *MethodDeclarationContext)

	// ExitMethodBody is called when exiting the methodBody production.
	ExitMethodBody(c *MethodBodyContext)

	// ExitTypeTypeOrVoid is called when exiting the typeTypeOrVoid production.
	ExitTypeTypeOrVoid(c *TypeTypeOrVoidContext)

	// ExitGenericMethodDeclaration is called when exiting the genericMethodDeclaration production.
	ExitGenericMethodDeclaration(c *GenericMethodDeclarationContext)

	// ExitGenericConstructorDeclaration is called when exiting the genericConstructorDeclaration production.
	ExitGenericConstructorDeclaration(c *GenericConstructorDeclarationContext)

	// ExitConstructorDeclaration is called when exiting the constructorDeclaration production.
	ExitConstructorDeclaration(c *ConstructorDeclarationContext)

	// ExitFieldDeclaration is called when exiting the fieldDeclaration production.
	ExitFieldDeclaration(c *FieldDeclarationContext)

	// ExitInterfaceBodyDeclaration is called when exiting the interfaceBodyDeclaration production.
	ExitInterfaceBodyDeclaration(c *InterfaceBodyDeclarationContext)

	// ExitInterfaceMemberDeclaration is called when exiting the interfaceMemberDeclaration production.
	ExitInterfaceMemberDeclaration(c *InterfaceMemberDeclarationContext)

	// ExitConstDeclaration is called when exiting the constDeclaration production.
	ExitConstDeclaration(c *ConstDeclarationContext)

	// ExitConstantDeclarator is called when exiting the constantDeclarator production.
	ExitConstantDeclarator(c *ConstantDeclaratorContext)

	// ExitInterfaceMethodDeclaration is called when exiting the interfaceMethodDeclaration production.
	ExitInterfaceMethodDeclaration(c *InterfaceMethodDeclarationContext)

	// ExitInterfaceMethodModifier is called when exiting the interfaceMethodModifier production.
	ExitInterfaceMethodModifier(c *InterfaceMethodModifierContext)

	// ExitGenericInterfaceMethodDeclaration is called when exiting the genericInterfaceMethodDeclaration production.
	ExitGenericInterfaceMethodDeclaration(c *GenericInterfaceMethodDeclarationContext)

	// ExitVariableDeclarators is called when exiting the variableDeclarators production.
	ExitVariableDeclarators(c *VariableDeclaratorsContext)

	// ExitVariableDeclarator is called when exiting the variableDeclarator production.
	ExitVariableDeclarator(c *VariableDeclaratorContext)

	// ExitVariableDeclaratorId is called when exiting the variableDeclaratorId production.
	ExitVariableDeclaratorId(c *VariableDeclaratorIdContext)

	// ExitVariableInitializer is called when exiting the variableInitializer production.
	ExitVariableInitializer(c *VariableInitializerContext)

	// ExitArrayInitializer is called when exiting the arrayInitializer production.
	ExitArrayInitializer(c *ArrayInitializerContext)

	// ExitClassOrInterfaceType is called when exiting the classOrInterfaceType production.
	ExitClassOrInterfaceType(c *ClassOrInterfaceTypeContext)

	// ExitTypeArgument is called when exiting the typeArgument production.
	ExitTypeArgument(c *TypeArgumentContext)

	// ExitQualifiedNameList is called when exiting the qualifiedNameList production.
	ExitQualifiedNameList(c *QualifiedNameListContext)

	// ExitFormalParameters is called when exiting the formalParameters production.
	ExitFormalParameters(c *FormalParametersContext)

	// ExitFormalParameterList is called when exiting the formalParameterList production.
	ExitFormalParameterList(c *FormalParameterListContext)

	// ExitFormalParameter is called when exiting the formalParameter production.
	ExitFormalParameter(c *FormalParameterContext)

	// ExitLastFormalParameter is called when exiting the lastFormalParameter production.
	ExitLastFormalParameter(c *LastFormalParameterContext)

	// ExitQualifiedName is called when exiting the qualifiedName production.
	ExitQualifiedName(c *QualifiedNameContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIntegerLiteral is called when exiting the integerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatLiteral is called when exiting the floatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitElementValuePairs is called when exiting the elementValuePairs production.
	ExitElementValuePairs(c *ElementValuePairsContext)

	// ExitElementValuePair is called when exiting the elementValuePair production.
	ExitElementValuePair(c *ElementValuePairContext)

	// ExitElementValue is called when exiting the elementValue production.
	ExitElementValue(c *ElementValueContext)

	// ExitElementValueArrayInitializer is called when exiting the elementValueArrayInitializer production.
	ExitElementValueArrayInitializer(c *ElementValueArrayInitializerContext)

	// ExitAnnotationTypeDeclaration is called when exiting the annotationTypeDeclaration production.
	ExitAnnotationTypeDeclaration(c *AnnotationTypeDeclarationContext)

	// ExitAnnotationTypeBody is called when exiting the annotationTypeBody production.
	ExitAnnotationTypeBody(c *AnnotationTypeBodyContext)

	// ExitAnnotationTypeElementDeclaration is called when exiting the annotationTypeElementDeclaration production.
	ExitAnnotationTypeElementDeclaration(c *AnnotationTypeElementDeclarationContext)

	// ExitAnnotationTypeElementRest is called when exiting the annotationTypeElementRest production.
	ExitAnnotationTypeElementRest(c *AnnotationTypeElementRestContext)

	// ExitAnnotationMethodOrConstantRest is called when exiting the annotationMethodOrConstantRest production.
	ExitAnnotationMethodOrConstantRest(c *AnnotationMethodOrConstantRestContext)

	// ExitAnnotationMethodRest is called when exiting the annotationMethodRest production.
	ExitAnnotationMethodRest(c *AnnotationMethodRestContext)

	// ExitAnnotationConstantRest is called when exiting the annotationConstantRest production.
	ExitAnnotationConstantRest(c *AnnotationConstantRestContext)

	// ExitDefaultValue is called when exiting the defaultValue production.
	ExitDefaultValue(c *DefaultValueContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBlockStatement is called when exiting the blockStatement production.
	ExitBlockStatement(c *BlockStatementContext)

	// ExitLocalVariableDeclaration is called when exiting the localVariableDeclaration production.
	ExitLocalVariableDeclaration(c *LocalVariableDeclarationContext)

	// ExitLocalTypeDeclaration is called when exiting the localTypeDeclaration production.
	ExitLocalTypeDeclaration(c *LocalTypeDeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitCatchClause is called when exiting the catchClause production.
	ExitCatchClause(c *CatchClauseContext)

	// ExitCatchType is called when exiting the catchType production.
	ExitCatchType(c *CatchTypeContext)

	// ExitFinallyBlock is called when exiting the finallyBlock production.
	ExitFinallyBlock(c *FinallyBlockContext)

	// ExitResourceSpecification is called when exiting the resourceSpecification production.
	ExitResourceSpecification(c *ResourceSpecificationContext)

	// ExitResources is called when exiting the resources production.
	ExitResources(c *ResourcesContext)

	// ExitResource is called when exiting the resource production.
	ExitResource(c *ResourceContext)

	// ExitSwitchBlockStatementGroup is called when exiting the switchBlockStatementGroup production.
	ExitSwitchBlockStatementGroup(c *SwitchBlockStatementGroupContext)

	// ExitSwitchLabel is called when exiting the switchLabel production.
	ExitSwitchLabel(c *SwitchLabelContext)

	// ExitForControl is called when exiting the forControl production.
	ExitForControl(c *ForControlContext)

	// ExitForInit is called when exiting the forInit production.
	ExitForInit(c *ForInitContext)

	// ExitEnhancedForControl is called when exiting the enhancedForControl production.
	ExitEnhancedForControl(c *EnhancedForControlContext)

	// ExitParExpression is called when exiting the parExpression production.
	ExitParExpression(c *ParExpressionContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitMethodCall is called when exiting the methodCall production.
	ExitMethodCall(c *MethodCallContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLambdaExpression is called when exiting the lambdaExpression production.
	ExitLambdaExpression(c *LambdaExpressionContext)

	// ExitLambdaParameters is called when exiting the lambdaParameters production.
	ExitLambdaParameters(c *LambdaParametersContext)

	// ExitLambdaBody is called when exiting the lambdaBody production.
	ExitLambdaBody(c *LambdaBodyContext)

	// ExitPrimary is called when exiting the primary production.
	ExitPrimary(c *PrimaryContext)

	// ExitClassType is called when exiting the classType production.
	ExitClassType(c *ClassTypeContext)

	// ExitCreator is called when exiting the creator production.
	ExitCreator(c *CreatorContext)

	// ExitCreatedName is called when exiting the createdName production.
	ExitCreatedName(c *CreatedNameContext)

	// ExitInnerCreator is called when exiting the innerCreator production.
	ExitInnerCreator(c *InnerCreatorContext)

	// ExitArrayCreatorRest is called when exiting the arrayCreatorRest production.
	ExitArrayCreatorRest(c *ArrayCreatorRestContext)

	// ExitClassCreatorRest is called when exiting the classCreatorRest production.
	ExitClassCreatorRest(c *ClassCreatorRestContext)

	// ExitExplicitGenericInvocation is called when exiting the explicitGenericInvocation production.
	ExitExplicitGenericInvocation(c *ExplicitGenericInvocationContext)

	// ExitTypeArgumentsOrDiamond is called when exiting the typeArgumentsOrDiamond production.
	ExitTypeArgumentsOrDiamond(c *TypeArgumentsOrDiamondContext)

	// ExitNonWildcardTypeArgumentsOrDiamond is called when exiting the nonWildcardTypeArgumentsOrDiamond production.
	ExitNonWildcardTypeArgumentsOrDiamond(c *NonWildcardTypeArgumentsOrDiamondContext)

	// ExitNonWildcardTypeArguments is called when exiting the nonWildcardTypeArguments production.
	ExitNonWildcardTypeArguments(c *NonWildcardTypeArgumentsContext)

	// ExitTypeList is called when exiting the typeList production.
	ExitTypeList(c *TypeListContext)

	// ExitTypeType is called when exiting the typeType production.
	ExitTypeType(c *TypeTypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitTypeArguments is called when exiting the typeArguments production.
	ExitTypeArguments(c *TypeArgumentsContext)

	// ExitSuperSuffix is called when exiting the superSuffix production.
	ExitSuperSuffix(c *SuperSuffixContext)

	// ExitExplicitGenericInvocationSuffix is called when exiting the explicitGenericInvocationSuffix production.
	ExitExplicitGenericInvocationSuffix(c *ExplicitGenericInvocationSuffixContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)
}
