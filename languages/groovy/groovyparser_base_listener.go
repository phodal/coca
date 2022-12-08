// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // GroovyParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseGroovyParserListener is a complete listener for a parse tree produced by GroovyParser.
type BaseGroovyParserListener struct{}

var _ GroovyParserListener = &BaseGroovyParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGroovyParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGroovyParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGroovyParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGroovyParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCompilationUnit is called when production compilationUnit is entered.
func (s *BaseGroovyParserListener) EnterCompilationUnit(ctx *CompilationUnitContext) {}

// ExitCompilationUnit is called when production compilationUnit is exited.
func (s *BaseGroovyParserListener) ExitCompilationUnit(ctx *CompilationUnitContext) {}

// EnterScriptStatements is called when production scriptStatements is entered.
func (s *BaseGroovyParserListener) EnterScriptStatements(ctx *ScriptStatementsContext) {}

// ExitScriptStatements is called when production scriptStatements is exited.
func (s *BaseGroovyParserListener) ExitScriptStatements(ctx *ScriptStatementsContext) {}

// EnterScriptStatement is called when production scriptStatement is entered.
func (s *BaseGroovyParserListener) EnterScriptStatement(ctx *ScriptStatementContext) {}

// ExitScriptStatement is called when production scriptStatement is exited.
func (s *BaseGroovyParserListener) ExitScriptStatement(ctx *ScriptStatementContext) {}

// EnterPackageDeclaration is called when production packageDeclaration is entered.
func (s *BaseGroovyParserListener) EnterPackageDeclaration(ctx *PackageDeclarationContext) {}

// ExitPackageDeclaration is called when production packageDeclaration is exited.
func (s *BaseGroovyParserListener) ExitPackageDeclaration(ctx *PackageDeclarationContext) {}

// EnterImportDeclaration is called when production importDeclaration is entered.
func (s *BaseGroovyParserListener) EnterImportDeclaration(ctx *ImportDeclarationContext) {}

// ExitImportDeclaration is called when production importDeclaration is exited.
func (s *BaseGroovyParserListener) ExitImportDeclaration(ctx *ImportDeclarationContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseGroovyParserListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseGroovyParserListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterModifier is called when production modifier is entered.
func (s *BaseGroovyParserListener) EnterModifier(ctx *ModifierContext) {}

// ExitModifier is called when production modifier is exited.
func (s *BaseGroovyParserListener) ExitModifier(ctx *ModifierContext) {}

// EnterModifiersOpt is called when production modifiersOpt is entered.
func (s *BaseGroovyParserListener) EnterModifiersOpt(ctx *ModifiersOptContext) {}

// ExitModifiersOpt is called when production modifiersOpt is exited.
func (s *BaseGroovyParserListener) ExitModifiersOpt(ctx *ModifiersOptContext) {}

// EnterModifiers is called when production modifiers is entered.
func (s *BaseGroovyParserListener) EnterModifiers(ctx *ModifiersContext) {}

// ExitModifiers is called when production modifiers is exited.
func (s *BaseGroovyParserListener) ExitModifiers(ctx *ModifiersContext) {}

// EnterClassOrInterfaceModifiersOpt is called when production classOrInterfaceModifiersOpt is entered.
func (s *BaseGroovyParserListener) EnterClassOrInterfaceModifiersOpt(ctx *ClassOrInterfaceModifiersOptContext) {
}

// ExitClassOrInterfaceModifiersOpt is called when production classOrInterfaceModifiersOpt is exited.
func (s *BaseGroovyParserListener) ExitClassOrInterfaceModifiersOpt(ctx *ClassOrInterfaceModifiersOptContext) {
}

// EnterClassOrInterfaceModifiers is called when production classOrInterfaceModifiers is entered.
func (s *BaseGroovyParserListener) EnterClassOrInterfaceModifiers(ctx *ClassOrInterfaceModifiersContext) {
}

// ExitClassOrInterfaceModifiers is called when production classOrInterfaceModifiers is exited.
func (s *BaseGroovyParserListener) ExitClassOrInterfaceModifiers(ctx *ClassOrInterfaceModifiersContext) {
}

// EnterClassOrInterfaceModifier is called when production classOrInterfaceModifier is entered.
func (s *BaseGroovyParserListener) EnterClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) {
}

// ExitClassOrInterfaceModifier is called when production classOrInterfaceModifier is exited.
func (s *BaseGroovyParserListener) ExitClassOrInterfaceModifier(ctx *ClassOrInterfaceModifierContext) {
}

// EnterVariableModifier is called when production variableModifier is entered.
func (s *BaseGroovyParserListener) EnterVariableModifier(ctx *VariableModifierContext) {}

// ExitVariableModifier is called when production variableModifier is exited.
func (s *BaseGroovyParserListener) ExitVariableModifier(ctx *VariableModifierContext) {}

// EnterVariableModifiersOpt is called when production variableModifiersOpt is entered.
func (s *BaseGroovyParserListener) EnterVariableModifiersOpt(ctx *VariableModifiersOptContext) {}

// ExitVariableModifiersOpt is called when production variableModifiersOpt is exited.
func (s *BaseGroovyParserListener) ExitVariableModifiersOpt(ctx *VariableModifiersOptContext) {}

// EnterVariableModifiers is called when production variableModifiers is entered.
func (s *BaseGroovyParserListener) EnterVariableModifiers(ctx *VariableModifiersContext) {}

// ExitVariableModifiers is called when production variableModifiers is exited.
func (s *BaseGroovyParserListener) ExitVariableModifiers(ctx *VariableModifiersContext) {}

// EnterTypeParameters is called when production typeParameters is entered.
func (s *BaseGroovyParserListener) EnterTypeParameters(ctx *TypeParametersContext) {}

// ExitTypeParameters is called when production typeParameters is exited.
func (s *BaseGroovyParserListener) ExitTypeParameters(ctx *TypeParametersContext) {}

// EnterTypeParameter is called when production typeParameter is entered.
func (s *BaseGroovyParserListener) EnterTypeParameter(ctx *TypeParameterContext) {}

// ExitTypeParameter is called when production typeParameter is exited.
func (s *BaseGroovyParserListener) ExitTypeParameter(ctx *TypeParameterContext) {}

// EnterTypeBound is called when production typeBound is entered.
func (s *BaseGroovyParserListener) EnterTypeBound(ctx *TypeBoundContext) {}

// ExitTypeBound is called when production typeBound is exited.
func (s *BaseGroovyParserListener) ExitTypeBound(ctx *TypeBoundContext) {}

// EnterTypeList is called when production typeList is entered.
func (s *BaseGroovyParserListener) EnterTypeList(ctx *TypeListContext) {}

// ExitTypeList is called when production typeList is exited.
func (s *BaseGroovyParserListener) ExitTypeList(ctx *TypeListContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseGroovyParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseGroovyParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBody is called when production classBody is entered.
func (s *BaseGroovyParserListener) EnterClassBody(ctx *ClassBodyContext) {}

// ExitClassBody is called when production classBody is exited.
func (s *BaseGroovyParserListener) ExitClassBody(ctx *ClassBodyContext) {}

// EnterEnumConstants is called when production enumConstants is entered.
func (s *BaseGroovyParserListener) EnterEnumConstants(ctx *EnumConstantsContext) {}

// ExitEnumConstants is called when production enumConstants is exited.
func (s *BaseGroovyParserListener) ExitEnumConstants(ctx *EnumConstantsContext) {}

// EnterEnumConstant is called when production enumConstant is entered.
func (s *BaseGroovyParserListener) EnterEnumConstant(ctx *EnumConstantContext) {}

// ExitEnumConstant is called when production enumConstant is exited.
func (s *BaseGroovyParserListener) ExitEnumConstant(ctx *EnumConstantContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseGroovyParserListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseGroovyParserListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterMemberDeclaration is called when production memberDeclaration is entered.
func (s *BaseGroovyParserListener) EnterMemberDeclaration(ctx *MemberDeclarationContext) {}

// ExitMemberDeclaration is called when production memberDeclaration is exited.
func (s *BaseGroovyParserListener) ExitMemberDeclaration(ctx *MemberDeclarationContext) {}

// EnterMethodDeclaration is called when production methodDeclaration is entered.
func (s *BaseGroovyParserListener) EnterMethodDeclaration(ctx *MethodDeclarationContext) {}

// ExitMethodDeclaration is called when production methodDeclaration is exited.
func (s *BaseGroovyParserListener) ExitMethodDeclaration(ctx *MethodDeclarationContext) {}

// EnterMethodName is called when production methodName is entered.
func (s *BaseGroovyParserListener) EnterMethodName(ctx *MethodNameContext) {}

// ExitMethodName is called when production methodName is exited.
func (s *BaseGroovyParserListener) ExitMethodName(ctx *MethodNameContext) {}

// EnterReturnType is called when production returnType is entered.
func (s *BaseGroovyParserListener) EnterReturnType(ctx *ReturnTypeContext) {}

// ExitReturnType is called when production returnType is exited.
func (s *BaseGroovyParserListener) ExitReturnType(ctx *ReturnTypeContext) {}

// EnterFieldDeclaration is called when production fieldDeclaration is entered.
func (s *BaseGroovyParserListener) EnterFieldDeclaration(ctx *FieldDeclarationContext) {}

// ExitFieldDeclaration is called when production fieldDeclaration is exited.
func (s *BaseGroovyParserListener) ExitFieldDeclaration(ctx *FieldDeclarationContext) {}

// EnterVariableDeclarators is called when production variableDeclarators is entered.
func (s *BaseGroovyParserListener) EnterVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// ExitVariableDeclarators is called when production variableDeclarators is exited.
func (s *BaseGroovyParserListener) ExitVariableDeclarators(ctx *VariableDeclaratorsContext) {}

// EnterVariableDeclarator is called when production variableDeclarator is entered.
func (s *BaseGroovyParserListener) EnterVariableDeclarator(ctx *VariableDeclaratorContext) {}

// ExitVariableDeclarator is called when production variableDeclarator is exited.
func (s *BaseGroovyParserListener) ExitVariableDeclarator(ctx *VariableDeclaratorContext) {}

// EnterVariableDeclaratorId is called when production variableDeclaratorId is entered.
func (s *BaseGroovyParserListener) EnterVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// ExitVariableDeclaratorId is called when production variableDeclaratorId is exited.
func (s *BaseGroovyParserListener) ExitVariableDeclaratorId(ctx *VariableDeclaratorIdContext) {}

// EnterVariableInitializer is called when production variableInitializer is entered.
func (s *BaseGroovyParserListener) EnterVariableInitializer(ctx *VariableInitializerContext) {}

// ExitVariableInitializer is called when production variableInitializer is exited.
func (s *BaseGroovyParserListener) ExitVariableInitializer(ctx *VariableInitializerContext) {}

// EnterVariableInitializers is called when production variableInitializers is entered.
func (s *BaseGroovyParserListener) EnterVariableInitializers(ctx *VariableInitializersContext) {}

// ExitVariableInitializers is called when production variableInitializers is exited.
func (s *BaseGroovyParserListener) ExitVariableInitializers(ctx *VariableInitializersContext) {}

// EnterDims is called when production dims is entered.
func (s *BaseGroovyParserListener) EnterDims(ctx *DimsContext) {}

// ExitDims is called when production dims is exited.
func (s *BaseGroovyParserListener) ExitDims(ctx *DimsContext) {}

// EnterDimsOpt is called when production dimsOpt is entered.
func (s *BaseGroovyParserListener) EnterDimsOpt(ctx *DimsOptContext) {}

// ExitDimsOpt is called when production dimsOpt is exited.
func (s *BaseGroovyParserListener) ExitDimsOpt(ctx *DimsOptContext) {}

// EnterStandardType is called when production standardType is entered.
func (s *BaseGroovyParserListener) EnterStandardType(ctx *StandardTypeContext) {}

// ExitStandardType is called when production standardType is exited.
func (s *BaseGroovyParserListener) ExitStandardType(ctx *StandardTypeContext) {}

// EnterTypeType is called when production typeType is entered.
func (s *BaseGroovyParserListener) EnterTypeType(ctx *TypeTypeContext) {}

// ExitTypeType is called when production typeType is exited.
func (s *BaseGroovyParserListener) ExitTypeType(ctx *TypeTypeContext) {}

// EnterClassOrInterfaceType is called when production classOrInterfaceType is entered.
func (s *BaseGroovyParserListener) EnterClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// ExitClassOrInterfaceType is called when production classOrInterfaceType is exited.
func (s *BaseGroovyParserListener) ExitClassOrInterfaceType(ctx *ClassOrInterfaceTypeContext) {}

// EnterGeneralClassOrInterfaceType is called when production generalClassOrInterfaceType is entered.
func (s *BaseGroovyParserListener) EnterGeneralClassOrInterfaceType(ctx *GeneralClassOrInterfaceTypeContext) {
}

// ExitGeneralClassOrInterfaceType is called when production generalClassOrInterfaceType is exited.
func (s *BaseGroovyParserListener) ExitGeneralClassOrInterfaceType(ctx *GeneralClassOrInterfaceTypeContext) {
}

// EnterStandardClassOrInterfaceType is called when production standardClassOrInterfaceType is entered.
func (s *BaseGroovyParserListener) EnterStandardClassOrInterfaceType(ctx *StandardClassOrInterfaceTypeContext) {
}

// ExitStandardClassOrInterfaceType is called when production standardClassOrInterfaceType is exited.
func (s *BaseGroovyParserListener) ExitStandardClassOrInterfaceType(ctx *StandardClassOrInterfaceTypeContext) {
}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseGroovyParserListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseGroovyParserListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterTypeArguments is called when production typeArguments is entered.
func (s *BaseGroovyParserListener) EnterTypeArguments(ctx *TypeArgumentsContext) {}

// ExitTypeArguments is called when production typeArguments is exited.
func (s *BaseGroovyParserListener) ExitTypeArguments(ctx *TypeArgumentsContext) {}

// EnterTypeArgument is called when production typeArgument is entered.
func (s *BaseGroovyParserListener) EnterTypeArgument(ctx *TypeArgumentContext) {}

// ExitTypeArgument is called when production typeArgument is exited.
func (s *BaseGroovyParserListener) ExitTypeArgument(ctx *TypeArgumentContext) {}

// EnterAnnotatedQualifiedClassName is called when production annotatedQualifiedClassName is entered.
func (s *BaseGroovyParserListener) EnterAnnotatedQualifiedClassName(ctx *AnnotatedQualifiedClassNameContext) {
}

// ExitAnnotatedQualifiedClassName is called when production annotatedQualifiedClassName is exited.
func (s *BaseGroovyParserListener) ExitAnnotatedQualifiedClassName(ctx *AnnotatedQualifiedClassNameContext) {
}

// EnterQualifiedClassNameList is called when production qualifiedClassNameList is entered.
func (s *BaseGroovyParserListener) EnterQualifiedClassNameList(ctx *QualifiedClassNameListContext) {}

// ExitQualifiedClassNameList is called when production qualifiedClassNameList is exited.
func (s *BaseGroovyParserListener) ExitQualifiedClassNameList(ctx *QualifiedClassNameListContext) {}

// EnterFormalParameters is called when production formalParameters is entered.
func (s *BaseGroovyParserListener) EnterFormalParameters(ctx *FormalParametersContext) {}

// ExitFormalParameters is called when production formalParameters is exited.
func (s *BaseGroovyParserListener) ExitFormalParameters(ctx *FormalParametersContext) {}

// EnterFormalParameterList is called when production formalParameterList is entered.
func (s *BaseGroovyParserListener) EnterFormalParameterList(ctx *FormalParameterListContext) {}

// ExitFormalParameterList is called when production formalParameterList is exited.
func (s *BaseGroovyParserListener) ExitFormalParameterList(ctx *FormalParameterListContext) {}

// EnterThisFormalParameter is called when production thisFormalParameter is entered.
func (s *BaseGroovyParserListener) EnterThisFormalParameter(ctx *ThisFormalParameterContext) {}

// ExitThisFormalParameter is called when production thisFormalParameter is exited.
func (s *BaseGroovyParserListener) ExitThisFormalParameter(ctx *ThisFormalParameterContext) {}

// EnterFormalParameter is called when production formalParameter is entered.
func (s *BaseGroovyParserListener) EnterFormalParameter(ctx *FormalParameterContext) {}

// ExitFormalParameter is called when production formalParameter is exited.
func (s *BaseGroovyParserListener) ExitFormalParameter(ctx *FormalParameterContext) {}

// EnterMethodBody is called when production methodBody is entered.
func (s *BaseGroovyParserListener) EnterMethodBody(ctx *MethodBodyContext) {}

// ExitMethodBody is called when production methodBody is exited.
func (s *BaseGroovyParserListener) ExitMethodBody(ctx *MethodBodyContext) {}

// EnterQualifiedName is called when production qualifiedName is entered.
func (s *BaseGroovyParserListener) EnterQualifiedName(ctx *QualifiedNameContext) {}

// ExitQualifiedName is called when production qualifiedName is exited.
func (s *BaseGroovyParserListener) ExitQualifiedName(ctx *QualifiedNameContext) {}

// EnterQualifiedNameElement is called when production qualifiedNameElement is entered.
func (s *BaseGroovyParserListener) EnterQualifiedNameElement(ctx *QualifiedNameElementContext) {}

// ExitQualifiedNameElement is called when production qualifiedNameElement is exited.
func (s *BaseGroovyParserListener) ExitQualifiedNameElement(ctx *QualifiedNameElementContext) {}

// EnterQualifiedNameElements is called when production qualifiedNameElements is entered.
func (s *BaseGroovyParserListener) EnterQualifiedNameElements(ctx *QualifiedNameElementsContext) {}

// ExitQualifiedNameElements is called when production qualifiedNameElements is exited.
func (s *BaseGroovyParserListener) ExitQualifiedNameElements(ctx *QualifiedNameElementsContext) {}

// EnterQualifiedClassName is called when production qualifiedClassName is entered.
func (s *BaseGroovyParserListener) EnterQualifiedClassName(ctx *QualifiedClassNameContext) {}

// ExitQualifiedClassName is called when production qualifiedClassName is exited.
func (s *BaseGroovyParserListener) ExitQualifiedClassName(ctx *QualifiedClassNameContext) {}

// EnterQualifiedStandardClassName is called when production qualifiedStandardClassName is entered.
func (s *BaseGroovyParserListener) EnterQualifiedStandardClassName(ctx *QualifiedStandardClassNameContext) {
}

// ExitQualifiedStandardClassName is called when production qualifiedStandardClassName is exited.
func (s *BaseGroovyParserListener) ExitQualifiedStandardClassName(ctx *QualifiedStandardClassNameContext) {
}

// EnterIntegerLiteralAlt is called when production integerLiteralAlt is entered.
func (s *BaseGroovyParserListener) EnterIntegerLiteralAlt(ctx *IntegerLiteralAltContext) {}

// ExitIntegerLiteralAlt is called when production integerLiteralAlt is exited.
func (s *BaseGroovyParserListener) ExitIntegerLiteralAlt(ctx *IntegerLiteralAltContext) {}

// EnterFloatingPointLiteralAlt is called when production floatingPointLiteralAlt is entered.
func (s *BaseGroovyParserListener) EnterFloatingPointLiteralAlt(ctx *FloatingPointLiteralAltContext) {
}

// ExitFloatingPointLiteralAlt is called when production floatingPointLiteralAlt is exited.
func (s *BaseGroovyParserListener) ExitFloatingPointLiteralAlt(ctx *FloatingPointLiteralAltContext) {}

// EnterStringLiteralAlt is called when production stringLiteralAlt is entered.
func (s *BaseGroovyParserListener) EnterStringLiteralAlt(ctx *StringLiteralAltContext) {}

// ExitStringLiteralAlt is called when production stringLiteralAlt is exited.
func (s *BaseGroovyParserListener) ExitStringLiteralAlt(ctx *StringLiteralAltContext) {}

// EnterBoolLiteralAlt is called when production boolLiteralAlt is entered.
func (s *BaseGroovyParserListener) EnterBoolLiteralAlt(ctx *BoolLiteralAltContext) {}

// ExitBoolLiteralAlt is called when production boolLiteralAlt is exited.
func (s *BaseGroovyParserListener) ExitBoolLiteralAlt(ctx *BoolLiteralAltContext) {}

// EnterNilLiteralAlt is called when production nilLiteralAlt is entered.
func (s *BaseGroovyParserListener) EnterNilLiteralAlt(ctx *NilLiteralAltContext) {}

// ExitNilLiteralAlt is called when production nilLiteralAlt is exited.
func (s *BaseGroovyParserListener) ExitNilLiteralAlt(ctx *NilLiteralAltContext) {}

// EnterGstring is called when production gstring is entered.
func (s *BaseGroovyParserListener) EnterGstring(ctx *GstringContext) {}

// ExitGstring is called when production gstring is exited.
func (s *BaseGroovyParserListener) ExitGstring(ctx *GstringContext) {}

// EnterGstringValue is called when production gstringValue is entered.
func (s *BaseGroovyParserListener) EnterGstringValue(ctx *GstringValueContext) {}

// ExitGstringValue is called when production gstringValue is exited.
func (s *BaseGroovyParserListener) ExitGstringValue(ctx *GstringValueContext) {}

// EnterGstringPath is called when production gstringPath is entered.
func (s *BaseGroovyParserListener) EnterGstringPath(ctx *GstringPathContext) {}

// ExitGstringPath is called when production gstringPath is exited.
func (s *BaseGroovyParserListener) ExitGstringPath(ctx *GstringPathContext) {}

// EnterLambdaExpression is called when production lambdaExpression is entered.
func (s *BaseGroovyParserListener) EnterLambdaExpression(ctx *LambdaExpressionContext) {}

// ExitLambdaExpression is called when production lambdaExpression is exited.
func (s *BaseGroovyParserListener) ExitLambdaExpression(ctx *LambdaExpressionContext) {}

// EnterStandardLambdaExpression is called when production standardLambdaExpression is entered.
func (s *BaseGroovyParserListener) EnterStandardLambdaExpression(ctx *StandardLambdaExpressionContext) {
}

// ExitStandardLambdaExpression is called when production standardLambdaExpression is exited.
func (s *BaseGroovyParserListener) ExitStandardLambdaExpression(ctx *StandardLambdaExpressionContext) {
}

// EnterLambdaParameters is called when production lambdaParameters is entered.
func (s *BaseGroovyParserListener) EnterLambdaParameters(ctx *LambdaParametersContext) {}

// ExitLambdaParameters is called when production lambdaParameters is exited.
func (s *BaseGroovyParserListener) ExitLambdaParameters(ctx *LambdaParametersContext) {}

// EnterStandardLambdaParameters is called when production standardLambdaParameters is entered.
func (s *BaseGroovyParserListener) EnterStandardLambdaParameters(ctx *StandardLambdaParametersContext) {
}

// ExitStandardLambdaParameters is called when production standardLambdaParameters is exited.
func (s *BaseGroovyParserListener) ExitStandardLambdaParameters(ctx *StandardLambdaParametersContext) {
}

// EnterLambdaBody is called when production lambdaBody is entered.
func (s *BaseGroovyParserListener) EnterLambdaBody(ctx *LambdaBodyContext) {}

// ExitLambdaBody is called when production lambdaBody is exited.
func (s *BaseGroovyParserListener) ExitLambdaBody(ctx *LambdaBodyContext) {}

// EnterClosure is called when production closure is entered.
func (s *BaseGroovyParserListener) EnterClosure(ctx *ClosureContext) {}

// ExitClosure is called when production closure is exited.
func (s *BaseGroovyParserListener) ExitClosure(ctx *ClosureContext) {}

// EnterClosureOrLambdaExpression is called when production closureOrLambdaExpression is entered.
func (s *BaseGroovyParserListener) EnterClosureOrLambdaExpression(ctx *ClosureOrLambdaExpressionContext) {
}

// ExitClosureOrLambdaExpression is called when production closureOrLambdaExpression is exited.
func (s *BaseGroovyParserListener) ExitClosureOrLambdaExpression(ctx *ClosureOrLambdaExpressionContext) {
}

// EnterBlockStatementsOpt is called when production blockStatementsOpt is entered.
func (s *BaseGroovyParserListener) EnterBlockStatementsOpt(ctx *BlockStatementsOptContext) {}

// ExitBlockStatementsOpt is called when production blockStatementsOpt is exited.
func (s *BaseGroovyParserListener) ExitBlockStatementsOpt(ctx *BlockStatementsOptContext) {}

// EnterBlockStatements is called when production blockStatements is entered.
func (s *BaseGroovyParserListener) EnterBlockStatements(ctx *BlockStatementsContext) {}

// ExitBlockStatements is called when production blockStatements is exited.
func (s *BaseGroovyParserListener) ExitBlockStatements(ctx *BlockStatementsContext) {}

// EnterAnnotationsOpt is called when production annotationsOpt is entered.
func (s *BaseGroovyParserListener) EnterAnnotationsOpt(ctx *AnnotationsOptContext) {}

// ExitAnnotationsOpt is called when production annotationsOpt is exited.
func (s *BaseGroovyParserListener) ExitAnnotationsOpt(ctx *AnnotationsOptContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseGroovyParserListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseGroovyParserListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterElementValues is called when production elementValues is entered.
func (s *BaseGroovyParserListener) EnterElementValues(ctx *ElementValuesContext) {}

// ExitElementValues is called when production elementValues is exited.
func (s *BaseGroovyParserListener) ExitElementValues(ctx *ElementValuesContext) {}

// EnterAnnotationName is called when production annotationName is entered.
func (s *BaseGroovyParserListener) EnterAnnotationName(ctx *AnnotationNameContext) {}

// ExitAnnotationName is called when production annotationName is exited.
func (s *BaseGroovyParserListener) ExitAnnotationName(ctx *AnnotationNameContext) {}

// EnterElementValuePairs is called when production elementValuePairs is entered.
func (s *BaseGroovyParserListener) EnterElementValuePairs(ctx *ElementValuePairsContext) {}

// ExitElementValuePairs is called when production elementValuePairs is exited.
func (s *BaseGroovyParserListener) ExitElementValuePairs(ctx *ElementValuePairsContext) {}

// EnterElementValuePair is called when production elementValuePair is entered.
func (s *BaseGroovyParserListener) EnterElementValuePair(ctx *ElementValuePairContext) {}

// ExitElementValuePair is called when production elementValuePair is exited.
func (s *BaseGroovyParserListener) ExitElementValuePair(ctx *ElementValuePairContext) {}

// EnterElementValuePairName is called when production elementValuePairName is entered.
func (s *BaseGroovyParserListener) EnterElementValuePairName(ctx *ElementValuePairNameContext) {}

// ExitElementValuePairName is called when production elementValuePairName is exited.
func (s *BaseGroovyParserListener) ExitElementValuePairName(ctx *ElementValuePairNameContext) {}

// EnterElementValue is called when production elementValue is entered.
func (s *BaseGroovyParserListener) EnterElementValue(ctx *ElementValueContext) {}

// ExitElementValue is called when production elementValue is exited.
func (s *BaseGroovyParserListener) ExitElementValue(ctx *ElementValueContext) {}

// EnterElementValueArrayInitializer is called when production elementValueArrayInitializer is entered.
func (s *BaseGroovyParserListener) EnterElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// ExitElementValueArrayInitializer is called when production elementValueArrayInitializer is exited.
func (s *BaseGroovyParserListener) ExitElementValueArrayInitializer(ctx *ElementValueArrayInitializerContext) {
}

// EnterBlock is called when production block is entered.
func (s *BaseGroovyParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGroovyParserListener) ExitBlock(ctx *BlockContext) {}

// EnterBlockStatement is called when production blockStatement is entered.
func (s *BaseGroovyParserListener) EnterBlockStatement(ctx *BlockStatementContext) {}

// ExitBlockStatement is called when production blockStatement is exited.
func (s *BaseGroovyParserListener) ExitBlockStatement(ctx *BlockStatementContext) {}

// EnterLocalVariableDeclaration is called when production localVariableDeclaration is entered.
func (s *BaseGroovyParserListener) EnterLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
}

// ExitLocalVariableDeclaration is called when production localVariableDeclaration is exited.
func (s *BaseGroovyParserListener) ExitLocalVariableDeclaration(ctx *LocalVariableDeclarationContext) {
}

// EnterClassifiedModifiers is called when production classifiedModifiers is entered.
func (s *BaseGroovyParserListener) EnterClassifiedModifiers(ctx *ClassifiedModifiersContext) {}

// ExitClassifiedModifiers is called when production classifiedModifiers is exited.
func (s *BaseGroovyParserListener) ExitClassifiedModifiers(ctx *ClassifiedModifiersContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseGroovyParserListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseGroovyParserListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterTypeNamePairs is called when production typeNamePairs is entered.
func (s *BaseGroovyParserListener) EnterTypeNamePairs(ctx *TypeNamePairsContext) {}

// ExitTypeNamePairs is called when production typeNamePairs is exited.
func (s *BaseGroovyParserListener) ExitTypeNamePairs(ctx *TypeNamePairsContext) {}

// EnterTypeNamePair is called when production typeNamePair is entered.
func (s *BaseGroovyParserListener) EnterTypeNamePair(ctx *TypeNamePairContext) {}

// ExitTypeNamePair is called when production typeNamePair is exited.
func (s *BaseGroovyParserListener) ExitTypeNamePair(ctx *TypeNamePairContext) {}

// EnterVariableNames is called when production variableNames is entered.
func (s *BaseGroovyParserListener) EnterVariableNames(ctx *VariableNamesContext) {}

// ExitVariableNames is called when production variableNames is exited.
func (s *BaseGroovyParserListener) ExitVariableNames(ctx *VariableNamesContext) {}

// EnterConditionalStatement is called when production conditionalStatement is entered.
func (s *BaseGroovyParserListener) EnterConditionalStatement(ctx *ConditionalStatementContext) {}

// ExitConditionalStatement is called when production conditionalStatement is exited.
func (s *BaseGroovyParserListener) ExitConditionalStatement(ctx *ConditionalStatementContext) {}

// EnterIfElseStatement is called when production ifElseStatement is entered.
func (s *BaseGroovyParserListener) EnterIfElseStatement(ctx *IfElseStatementContext) {}

// ExitIfElseStatement is called when production ifElseStatement is exited.
func (s *BaseGroovyParserListener) ExitIfElseStatement(ctx *IfElseStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BaseGroovyParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BaseGroovyParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterForStmtAlt is called when production forStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterForStmtAlt(ctx *ForStmtAltContext) {}

// ExitForStmtAlt is called when production forStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitForStmtAlt(ctx *ForStmtAltContext) {}

// EnterWhileStmtAlt is called when production whileStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterWhileStmtAlt(ctx *WhileStmtAltContext) {}

// ExitWhileStmtAlt is called when production whileStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitWhileStmtAlt(ctx *WhileStmtAltContext) {}

// EnterDoWhileStmtAlt is called when production doWhileStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterDoWhileStmtAlt(ctx *DoWhileStmtAltContext) {}

// ExitDoWhileStmtAlt is called when production doWhileStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitDoWhileStmtAlt(ctx *DoWhileStmtAltContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseGroovyParserListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseGroovyParserListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseGroovyParserListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseGroovyParserListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterTryCatchStatement is called when production tryCatchStatement is entered.
func (s *BaseGroovyParserListener) EnterTryCatchStatement(ctx *TryCatchStatementContext) {}

// ExitTryCatchStatement is called when production tryCatchStatement is exited.
func (s *BaseGroovyParserListener) ExitTryCatchStatement(ctx *TryCatchStatementContext) {}

// EnterAssertStatement is called when production assertStatement is entered.
func (s *BaseGroovyParserListener) EnterAssertStatement(ctx *AssertStatementContext) {}

// ExitAssertStatement is called when production assertStatement is exited.
func (s *BaseGroovyParserListener) ExitAssertStatement(ctx *AssertStatementContext) {}

// EnterBlockStmtAlt is called when production blockStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterBlockStmtAlt(ctx *BlockStmtAltContext) {}

// ExitBlockStmtAlt is called when production blockStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitBlockStmtAlt(ctx *BlockStmtAltContext) {}

// EnterConditionalStmtAlt is called when production conditionalStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterConditionalStmtAlt(ctx *ConditionalStmtAltContext) {}

// ExitConditionalStmtAlt is called when production conditionalStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitConditionalStmtAlt(ctx *ConditionalStmtAltContext) {}

// EnterLoopStmtAlt is called when production loopStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterLoopStmtAlt(ctx *LoopStmtAltContext) {}

// ExitLoopStmtAlt is called when production loopStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitLoopStmtAlt(ctx *LoopStmtAltContext) {}

// EnterTryCatchStmtAlt is called when production tryCatchStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterTryCatchStmtAlt(ctx *TryCatchStmtAltContext) {}

// ExitTryCatchStmtAlt is called when production tryCatchStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitTryCatchStmtAlt(ctx *TryCatchStmtAltContext) {}

// EnterSynchronizedStmtAlt is called when production synchronizedStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterSynchronizedStmtAlt(ctx *SynchronizedStmtAltContext) {}

// ExitSynchronizedStmtAlt is called when production synchronizedStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitSynchronizedStmtAlt(ctx *SynchronizedStmtAltContext) {}

// EnterReturnStmtAlt is called when production returnStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterReturnStmtAlt(ctx *ReturnStmtAltContext) {}

// ExitReturnStmtAlt is called when production returnStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitReturnStmtAlt(ctx *ReturnStmtAltContext) {}

// EnterThrowStmtAlt is called when production throwStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterThrowStmtAlt(ctx *ThrowStmtAltContext) {}

// ExitThrowStmtAlt is called when production throwStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitThrowStmtAlt(ctx *ThrowStmtAltContext) {}

// EnterBreakStmtAlt is called when production breakStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterBreakStmtAlt(ctx *BreakStmtAltContext) {}

// ExitBreakStmtAlt is called when production breakStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitBreakStmtAlt(ctx *BreakStmtAltContext) {}

// EnterContinueStmtAlt is called when production continueStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterContinueStmtAlt(ctx *ContinueStmtAltContext) {}

// ExitContinueStmtAlt is called when production continueStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitContinueStmtAlt(ctx *ContinueStmtAltContext) {}

// EnterLabeledStmtAlt is called when production labeledStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterLabeledStmtAlt(ctx *LabeledStmtAltContext) {}

// ExitLabeledStmtAlt is called when production labeledStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitLabeledStmtAlt(ctx *LabeledStmtAltContext) {}

// EnterAssertStmtAlt is called when production assertStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterAssertStmtAlt(ctx *AssertStmtAltContext) {}

// ExitAssertStmtAlt is called when production assertStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitAssertStmtAlt(ctx *AssertStmtAltContext) {}

// EnterLocalVariableDeclarationStmtAlt is called when production localVariableDeclarationStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterLocalVariableDeclarationStmtAlt(ctx *LocalVariableDeclarationStmtAltContext) {
}

// ExitLocalVariableDeclarationStmtAlt is called when production localVariableDeclarationStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitLocalVariableDeclarationStmtAlt(ctx *LocalVariableDeclarationStmtAltContext) {
}

// EnterMethodDeclarationStmtAlt is called when production methodDeclarationStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterMethodDeclarationStmtAlt(ctx *MethodDeclarationStmtAltContext) {
}

// ExitMethodDeclarationStmtAlt is called when production methodDeclarationStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitMethodDeclarationStmtAlt(ctx *MethodDeclarationStmtAltContext) {
}

// EnterExpressionStmtAlt is called when production expressionStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterExpressionStmtAlt(ctx *ExpressionStmtAltContext) {}

// ExitExpressionStmtAlt is called when production expressionStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitExpressionStmtAlt(ctx *ExpressionStmtAltContext) {}

// EnterEmptyStmtAlt is called when production emptyStmtAlt is entered.
func (s *BaseGroovyParserListener) EnterEmptyStmtAlt(ctx *EmptyStmtAltContext) {}

// ExitEmptyStmtAlt is called when production emptyStmtAlt is exited.
func (s *BaseGroovyParserListener) ExitEmptyStmtAlt(ctx *EmptyStmtAltContext) {}

// EnterCatchClause is called when production catchClause is entered.
func (s *BaseGroovyParserListener) EnterCatchClause(ctx *CatchClauseContext) {}

// ExitCatchClause is called when production catchClause is exited.
func (s *BaseGroovyParserListener) ExitCatchClause(ctx *CatchClauseContext) {}

// EnterCatchType is called when production catchType is entered.
func (s *BaseGroovyParserListener) EnterCatchType(ctx *CatchTypeContext) {}

// ExitCatchType is called when production catchType is exited.
func (s *BaseGroovyParserListener) ExitCatchType(ctx *CatchTypeContext) {}

// EnterFinallyBlock is called when production finallyBlock is entered.
func (s *BaseGroovyParserListener) EnterFinallyBlock(ctx *FinallyBlockContext) {}

// ExitFinallyBlock is called when production finallyBlock is exited.
func (s *BaseGroovyParserListener) ExitFinallyBlock(ctx *FinallyBlockContext) {}

// EnterResources is called when production resources is entered.
func (s *BaseGroovyParserListener) EnterResources(ctx *ResourcesContext) {}

// ExitResources is called when production resources is exited.
func (s *BaseGroovyParserListener) ExitResources(ctx *ResourcesContext) {}

// EnterResourceList is called when production resourceList is entered.
func (s *BaseGroovyParserListener) EnterResourceList(ctx *ResourceListContext) {}

// ExitResourceList is called when production resourceList is exited.
func (s *BaseGroovyParserListener) ExitResourceList(ctx *ResourceListContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseGroovyParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseGroovyParserListener) ExitResource(ctx *ResourceContext) {}

// EnterSwitchBlockStatementGroup is called when production switchBlockStatementGroup is entered.
func (s *BaseGroovyParserListener) EnterSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// ExitSwitchBlockStatementGroup is called when production switchBlockStatementGroup is exited.
func (s *BaseGroovyParserListener) ExitSwitchBlockStatementGroup(ctx *SwitchBlockStatementGroupContext) {
}

// EnterSwitchLabel is called when production switchLabel is entered.
func (s *BaseGroovyParserListener) EnterSwitchLabel(ctx *SwitchLabelContext) {}

// ExitSwitchLabel is called when production switchLabel is exited.
func (s *BaseGroovyParserListener) ExitSwitchLabel(ctx *SwitchLabelContext) {}

// EnterForControl is called when production forControl is entered.
func (s *BaseGroovyParserListener) EnterForControl(ctx *ForControlContext) {}

// ExitForControl is called when production forControl is exited.
func (s *BaseGroovyParserListener) ExitForControl(ctx *ForControlContext) {}

// EnterEnhancedForControl is called when production enhancedForControl is entered.
func (s *BaseGroovyParserListener) EnterEnhancedForControl(ctx *EnhancedForControlContext) {}

// ExitEnhancedForControl is called when production enhancedForControl is exited.
func (s *BaseGroovyParserListener) ExitEnhancedForControl(ctx *EnhancedForControlContext) {}

// EnterClassicalForControl is called when production classicalForControl is entered.
func (s *BaseGroovyParserListener) EnterClassicalForControl(ctx *ClassicalForControlContext) {}

// ExitClassicalForControl is called when production classicalForControl is exited.
func (s *BaseGroovyParserListener) ExitClassicalForControl(ctx *ClassicalForControlContext) {}

// EnterForInit is called when production forInit is entered.
func (s *BaseGroovyParserListener) EnterForInit(ctx *ForInitContext) {}

// ExitForInit is called when production forInit is exited.
func (s *BaseGroovyParserListener) ExitForInit(ctx *ForInitContext) {}

// EnterForUpdate is called when production forUpdate is entered.
func (s *BaseGroovyParserListener) EnterForUpdate(ctx *ForUpdateContext) {}

// ExitForUpdate is called when production forUpdate is exited.
func (s *BaseGroovyParserListener) ExitForUpdate(ctx *ForUpdateContext) {}

// EnterCastParExpression is called when production castParExpression is entered.
func (s *BaseGroovyParserListener) EnterCastParExpression(ctx *CastParExpressionContext) {}

// ExitCastParExpression is called when production castParExpression is exited.
func (s *BaseGroovyParserListener) ExitCastParExpression(ctx *CastParExpressionContext) {}

// EnterParExpression is called when production parExpression is entered.
func (s *BaseGroovyParserListener) EnterParExpression(ctx *ParExpressionContext) {}

// ExitParExpression is called when production parExpression is exited.
func (s *BaseGroovyParserListener) ExitParExpression(ctx *ParExpressionContext) {}

// EnterExpressionInPar is called when production expressionInPar is entered.
func (s *BaseGroovyParserListener) EnterExpressionInPar(ctx *ExpressionInParContext) {}

// ExitExpressionInPar is called when production expressionInPar is exited.
func (s *BaseGroovyParserListener) ExitExpressionInPar(ctx *ExpressionInParContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseGroovyParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseGroovyParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpressionListElement is called when production expressionListElement is entered.
func (s *BaseGroovyParserListener) EnterExpressionListElement(ctx *ExpressionListElementContext) {}

// ExitExpressionListElement is called when production expressionListElement is exited.
func (s *BaseGroovyParserListener) ExitExpressionListElement(ctx *ExpressionListElementContext) {}

// EnterEnhancedStatementExpression is called when production enhancedStatementExpression is entered.
func (s *BaseGroovyParserListener) EnterEnhancedStatementExpression(ctx *EnhancedStatementExpressionContext) {
}

// ExitEnhancedStatementExpression is called when production enhancedStatementExpression is exited.
func (s *BaseGroovyParserListener) ExitEnhancedStatementExpression(ctx *EnhancedStatementExpressionContext) {
}

// EnterCommandExprAlt is called when production commandExprAlt is entered.
func (s *BaseGroovyParserListener) EnterCommandExprAlt(ctx *CommandExprAltContext) {}

// ExitCommandExprAlt is called when production commandExprAlt is exited.
func (s *BaseGroovyParserListener) ExitCommandExprAlt(ctx *CommandExprAltContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseGroovyParserListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseGroovyParserListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterCastExprAltForExpr is called when production castExprAltForExpr is entered.
func (s *BaseGroovyParserListener) EnterCastExprAltForExpr(ctx *CastExprAltForExprContext) {}

// ExitCastExprAltForExpr is called when production castExprAltForExpr is exited.
func (s *BaseGroovyParserListener) ExitCastExprAltForExpr(ctx *CastExprAltForExprContext) {}

// EnterShiftExprAlt is called when production shiftExprAlt is entered.
func (s *BaseGroovyParserListener) EnterShiftExprAlt(ctx *ShiftExprAltContext) {}

// ExitShiftExprAlt is called when production shiftExprAlt is exited.
func (s *BaseGroovyParserListener) ExitShiftExprAlt(ctx *ShiftExprAltContext) {}

// EnterMultipleAssignmentExprAlt is called when production multipleAssignmentExprAlt is entered.
func (s *BaseGroovyParserListener) EnterMultipleAssignmentExprAlt(ctx *MultipleAssignmentExprAltContext) {
}

// ExitMultipleAssignmentExprAlt is called when production multipleAssignmentExprAlt is exited.
func (s *BaseGroovyParserListener) ExitMultipleAssignmentExprAlt(ctx *MultipleAssignmentExprAltContext) {
}

// EnterExclusiveOrExprAlt is called when production exclusiveOrExprAlt is entered.
func (s *BaseGroovyParserListener) EnterExclusiveOrExprAlt(ctx *ExclusiveOrExprAltContext) {}

// ExitExclusiveOrExprAlt is called when production exclusiveOrExprAlt is exited.
func (s *BaseGroovyParserListener) ExitExclusiveOrExprAlt(ctx *ExclusiveOrExprAltContext) {}

// EnterAdditiveExprAlt is called when production additiveExprAlt is entered.
func (s *BaseGroovyParserListener) EnterAdditiveExprAlt(ctx *AdditiveExprAltContext) {}

// ExitAdditiveExprAlt is called when production additiveExprAlt is exited.
func (s *BaseGroovyParserListener) ExitAdditiveExprAlt(ctx *AdditiveExprAltContext) {}

// EnterRegexExprAlt is called when production regexExprAlt is entered.
func (s *BaseGroovyParserListener) EnterRegexExprAlt(ctx *RegexExprAltContext) {}

// ExitRegexExprAlt is called when production regexExprAlt is exited.
func (s *BaseGroovyParserListener) ExitRegexExprAlt(ctx *RegexExprAltContext) {}

// EnterUnaryNotExprAltForExpr is called when production unaryNotExprAltForExpr is entered.
func (s *BaseGroovyParserListener) EnterUnaryNotExprAltForExpr(ctx *UnaryNotExprAltForExprContext) {}

// ExitUnaryNotExprAltForExpr is called when production unaryNotExprAltForExpr is exited.
func (s *BaseGroovyParserListener) ExitUnaryNotExprAltForExpr(ctx *UnaryNotExprAltForExprContext) {}

// EnterConditionalExprAlt is called when production conditionalExprAlt is entered.
func (s *BaseGroovyParserListener) EnterConditionalExprAlt(ctx *ConditionalExprAltContext) {}

// ExitConditionalExprAlt is called when production conditionalExprAlt is exited.
func (s *BaseGroovyParserListener) ExitConditionalExprAlt(ctx *ConditionalExprAltContext) {}

// EnterPowerExprAlt is called when production powerExprAlt is entered.
func (s *BaseGroovyParserListener) EnterPowerExprAlt(ctx *PowerExprAltContext) {}

// ExitPowerExprAlt is called when production powerExprAlt is exited.
func (s *BaseGroovyParserListener) ExitPowerExprAlt(ctx *PowerExprAltContext) {}

// EnterRelationalExprAlt is called when production relationalExprAlt is entered.
func (s *BaseGroovyParserListener) EnterRelationalExprAlt(ctx *RelationalExprAltContext) {}

// ExitRelationalExprAlt is called when production relationalExprAlt is exited.
func (s *BaseGroovyParserListener) ExitRelationalExprAlt(ctx *RelationalExprAltContext) {}

// EnterLogicalAndExprAlt is called when production logicalAndExprAlt is entered.
func (s *BaseGroovyParserListener) EnterLogicalAndExprAlt(ctx *LogicalAndExprAltContext) {}

// ExitLogicalAndExprAlt is called when production logicalAndExprAlt is exited.
func (s *BaseGroovyParserListener) ExitLogicalAndExprAlt(ctx *LogicalAndExprAltContext) {}

// EnterAssignmentExprAlt is called when production assignmentExprAlt is entered.
func (s *BaseGroovyParserListener) EnterAssignmentExprAlt(ctx *AssignmentExprAltContext) {}

// ExitAssignmentExprAlt is called when production assignmentExprAlt is exited.
func (s *BaseGroovyParserListener) ExitAssignmentExprAlt(ctx *AssignmentExprAltContext) {}

// EnterPostfixExprAltForExpr is called when production postfixExprAltForExpr is entered.
func (s *BaseGroovyParserListener) EnterPostfixExprAltForExpr(ctx *PostfixExprAltForExprContext) {}

// ExitPostfixExprAltForExpr is called when production postfixExprAltForExpr is exited.
func (s *BaseGroovyParserListener) ExitPostfixExprAltForExpr(ctx *PostfixExprAltForExprContext) {}

// EnterMultiplicativeExprAlt is called when production multiplicativeExprAlt is entered.
func (s *BaseGroovyParserListener) EnterMultiplicativeExprAlt(ctx *MultiplicativeExprAltContext) {}

// ExitMultiplicativeExprAlt is called when production multiplicativeExprAlt is exited.
func (s *BaseGroovyParserListener) ExitMultiplicativeExprAlt(ctx *MultiplicativeExprAltContext) {}

// EnterInclusiveOrExprAlt is called when production inclusiveOrExprAlt is entered.
func (s *BaseGroovyParserListener) EnterInclusiveOrExprAlt(ctx *InclusiveOrExprAltContext) {}

// ExitInclusiveOrExprAlt is called when production inclusiveOrExprAlt is exited.
func (s *BaseGroovyParserListener) ExitInclusiveOrExprAlt(ctx *InclusiveOrExprAltContext) {}

// EnterLogicalOrExprAlt is called when production logicalOrExprAlt is entered.
func (s *BaseGroovyParserListener) EnterLogicalOrExprAlt(ctx *LogicalOrExprAltContext) {}

// ExitLogicalOrExprAlt is called when production logicalOrExprAlt is exited.
func (s *BaseGroovyParserListener) ExitLogicalOrExprAlt(ctx *LogicalOrExprAltContext) {}

// EnterEqualityExprAlt is called when production equalityExprAlt is entered.
func (s *BaseGroovyParserListener) EnterEqualityExprAlt(ctx *EqualityExprAltContext) {}

// ExitEqualityExprAlt is called when production equalityExprAlt is exited.
func (s *BaseGroovyParserListener) ExitEqualityExprAlt(ctx *EqualityExprAltContext) {}

// EnterAndExprAlt is called when production andExprAlt is entered.
func (s *BaseGroovyParserListener) EnterAndExprAlt(ctx *AndExprAltContext) {}

// ExitAndExprAlt is called when production andExprAlt is exited.
func (s *BaseGroovyParserListener) ExitAndExprAlt(ctx *AndExprAltContext) {}

// EnterUnaryAddExprAltForExpr is called when production unaryAddExprAltForExpr is entered.
func (s *BaseGroovyParserListener) EnterUnaryAddExprAltForExpr(ctx *UnaryAddExprAltForExprContext) {}

// ExitUnaryAddExprAltForExpr is called when production unaryAddExprAltForExpr is exited.
func (s *BaseGroovyParserListener) ExitUnaryAddExprAltForExpr(ctx *UnaryAddExprAltForExprContext) {}

// EnterCastExprAlt is called when production castExprAlt is entered.
func (s *BaseGroovyParserListener) EnterCastExprAlt(ctx *CastExprAltContext) {}

// ExitCastExprAlt is called when production castExprAlt is exited.
func (s *BaseGroovyParserListener) ExitCastExprAlt(ctx *CastExprAltContext) {}

// EnterPostfixExprAlt is called when production postfixExprAlt is entered.
func (s *BaseGroovyParserListener) EnterPostfixExprAlt(ctx *PostfixExprAltContext) {}

// ExitPostfixExprAlt is called when production postfixExprAlt is exited.
func (s *BaseGroovyParserListener) ExitPostfixExprAlt(ctx *PostfixExprAltContext) {}

// EnterUnaryNotExprAlt is called when production unaryNotExprAlt is entered.
func (s *BaseGroovyParserListener) EnterUnaryNotExprAlt(ctx *UnaryNotExprAltContext) {}

// ExitUnaryNotExprAlt is called when production unaryNotExprAlt is exited.
func (s *BaseGroovyParserListener) ExitUnaryNotExprAlt(ctx *UnaryNotExprAltContext) {}

// EnterUnaryAddExprAlt is called when production unaryAddExprAlt is entered.
func (s *BaseGroovyParserListener) EnterUnaryAddExprAlt(ctx *UnaryAddExprAltContext) {}

// ExitUnaryAddExprAlt is called when production unaryAddExprAlt is exited.
func (s *BaseGroovyParserListener) ExitUnaryAddExprAlt(ctx *UnaryAddExprAltContext) {}

// EnterCommandExpression is called when production commandExpression is entered.
func (s *BaseGroovyParserListener) EnterCommandExpression(ctx *CommandExpressionContext) {}

// ExitCommandExpression is called when production commandExpression is exited.
func (s *BaseGroovyParserListener) ExitCommandExpression(ctx *CommandExpressionContext) {}

// EnterCommandArgument is called when production commandArgument is entered.
func (s *BaseGroovyParserListener) EnterCommandArgument(ctx *CommandArgumentContext) {}

// ExitCommandArgument is called when production commandArgument is exited.
func (s *BaseGroovyParserListener) ExitCommandArgument(ctx *CommandArgumentContext) {}

// EnterPathExpression is called when production pathExpression is entered.
func (s *BaseGroovyParserListener) EnterPathExpression(ctx *PathExpressionContext) {}

// ExitPathExpression is called when production pathExpression is exited.
func (s *BaseGroovyParserListener) ExitPathExpression(ctx *PathExpressionContext) {}

// EnterPathElement is called when production pathElement is entered.
func (s *BaseGroovyParserListener) EnterPathElement(ctx *PathElementContext) {}

// ExitPathElement is called when production pathElement is exited.
func (s *BaseGroovyParserListener) ExitPathElement(ctx *PathElementContext) {}

// EnterNamePart is called when production namePart is entered.
func (s *BaseGroovyParserListener) EnterNamePart(ctx *NamePartContext) {}

// ExitNamePart is called when production namePart is exited.
func (s *BaseGroovyParserListener) ExitNamePart(ctx *NamePartContext) {}

// EnterDynamicMemberName is called when production dynamicMemberName is entered.
func (s *BaseGroovyParserListener) EnterDynamicMemberName(ctx *DynamicMemberNameContext) {}

// ExitDynamicMemberName is called when production dynamicMemberName is exited.
func (s *BaseGroovyParserListener) ExitDynamicMemberName(ctx *DynamicMemberNameContext) {}

// EnterIndexPropertyArgs is called when production indexPropertyArgs is entered.
func (s *BaseGroovyParserListener) EnterIndexPropertyArgs(ctx *IndexPropertyArgsContext) {}

// ExitIndexPropertyArgs is called when production indexPropertyArgs is exited.
func (s *BaseGroovyParserListener) ExitIndexPropertyArgs(ctx *IndexPropertyArgsContext) {}

// EnterNamedPropertyArgs is called when production namedPropertyArgs is entered.
func (s *BaseGroovyParserListener) EnterNamedPropertyArgs(ctx *NamedPropertyArgsContext) {}

// ExitNamedPropertyArgs is called when production namedPropertyArgs is exited.
func (s *BaseGroovyParserListener) ExitNamedPropertyArgs(ctx *NamedPropertyArgsContext) {}

// EnterIdentifierPrmrAlt is called when production identifierPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterIdentifierPrmrAlt(ctx *IdentifierPrmrAltContext) {}

// ExitIdentifierPrmrAlt is called when production identifierPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitIdentifierPrmrAlt(ctx *IdentifierPrmrAltContext) {}

// EnterLiteralPrmrAlt is called when production literalPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterLiteralPrmrAlt(ctx *LiteralPrmrAltContext) {}

// ExitLiteralPrmrAlt is called when production literalPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitLiteralPrmrAlt(ctx *LiteralPrmrAltContext) {}

// EnterGstringPrmrAlt is called when production gstringPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterGstringPrmrAlt(ctx *GstringPrmrAltContext) {}

// ExitGstringPrmrAlt is called when production gstringPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitGstringPrmrAlt(ctx *GstringPrmrAltContext) {}

// EnterNewPrmrAlt is called when production newPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterNewPrmrAlt(ctx *NewPrmrAltContext) {}

// ExitNewPrmrAlt is called when production newPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitNewPrmrAlt(ctx *NewPrmrAltContext) {}

// EnterThisPrmrAlt is called when production thisPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterThisPrmrAlt(ctx *ThisPrmrAltContext) {}

// ExitThisPrmrAlt is called when production thisPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitThisPrmrAlt(ctx *ThisPrmrAltContext) {}

// EnterSuperPrmrAlt is called when production superPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterSuperPrmrAlt(ctx *SuperPrmrAltContext) {}

// ExitSuperPrmrAlt is called when production superPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitSuperPrmrAlt(ctx *SuperPrmrAltContext) {}

// EnterParenPrmrAlt is called when production parenPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterParenPrmrAlt(ctx *ParenPrmrAltContext) {}

// ExitParenPrmrAlt is called when production parenPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitParenPrmrAlt(ctx *ParenPrmrAltContext) {}

// EnterClosureOrLambdaExpressionPrmrAlt is called when production closureOrLambdaExpressionPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterClosureOrLambdaExpressionPrmrAlt(ctx *ClosureOrLambdaExpressionPrmrAltContext) {
}

// ExitClosureOrLambdaExpressionPrmrAlt is called when production closureOrLambdaExpressionPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitClosureOrLambdaExpressionPrmrAlt(ctx *ClosureOrLambdaExpressionPrmrAltContext) {
}

// EnterListPrmrAlt is called when production listPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterListPrmrAlt(ctx *ListPrmrAltContext) {}

// ExitListPrmrAlt is called when production listPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitListPrmrAlt(ctx *ListPrmrAltContext) {}

// EnterMapPrmrAlt is called when production mapPrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterMapPrmrAlt(ctx *MapPrmrAltContext) {}

// ExitMapPrmrAlt is called when production mapPrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitMapPrmrAlt(ctx *MapPrmrAltContext) {}

// EnterBuiltInTypePrmrAlt is called when production builtInTypePrmrAlt is entered.
func (s *BaseGroovyParserListener) EnterBuiltInTypePrmrAlt(ctx *BuiltInTypePrmrAltContext) {}

// ExitBuiltInTypePrmrAlt is called when production builtInTypePrmrAlt is exited.
func (s *BaseGroovyParserListener) ExitBuiltInTypePrmrAlt(ctx *BuiltInTypePrmrAltContext) {}

// EnterList is called when production list is entered.
func (s *BaseGroovyParserListener) EnterList(ctx *ListContext) {}

// ExitList is called when production list is exited.
func (s *BaseGroovyParserListener) ExitList(ctx *ListContext) {}

// EnterNormalMap is called when production normalMap is entered.
func (s *BaseGroovyParserListener) EnterNormalMap(ctx *NormalMapContext) {}

// ExitNormalMap is called when production normalMap is exited.
func (s *BaseGroovyParserListener) ExitNormalMap(ctx *NormalMapContext) {}

// EnterMapEntryList is called when production mapEntryList is entered.
func (s *BaseGroovyParserListener) EnterMapEntryList(ctx *MapEntryListContext) {}

// ExitMapEntryList is called when production mapEntryList is exited.
func (s *BaseGroovyParserListener) ExitMapEntryList(ctx *MapEntryListContext) {}

// EnterMapEntry is called when production mapEntry is entered.
func (s *BaseGroovyParserListener) EnterMapEntry(ctx *MapEntryContext) {}

// ExitMapEntry is called when production mapEntry is exited.
func (s *BaseGroovyParserListener) ExitMapEntry(ctx *MapEntryContext) {}

// EnterMapEntryLabel is called when production mapEntryLabel is entered.
func (s *BaseGroovyParserListener) EnterMapEntryLabel(ctx *MapEntryLabelContext) {}

// ExitMapEntryLabel is called when production mapEntryLabel is exited.
func (s *BaseGroovyParserListener) ExitMapEntryLabel(ctx *MapEntryLabelContext) {}

// EnterCreator is called when production creator is entered.
func (s *BaseGroovyParserListener) EnterCreator(ctx *CreatorContext) {}

// ExitCreator is called when production creator is exited.
func (s *BaseGroovyParserListener) ExitCreator(ctx *CreatorContext) {}

// EnterArrayInitializer is called when production arrayInitializer is entered.
func (s *BaseGroovyParserListener) EnterArrayInitializer(ctx *ArrayInitializerContext) {}

// ExitArrayInitializer is called when production arrayInitializer is exited.
func (s *BaseGroovyParserListener) ExitArrayInitializer(ctx *ArrayInitializerContext) {}

// EnterAnonymousInnerClassDeclaration is called when production anonymousInnerClassDeclaration is entered.
func (s *BaseGroovyParserListener) EnterAnonymousInnerClassDeclaration(ctx *AnonymousInnerClassDeclarationContext) {
}

// ExitAnonymousInnerClassDeclaration is called when production anonymousInnerClassDeclaration is exited.
func (s *BaseGroovyParserListener) ExitAnonymousInnerClassDeclaration(ctx *AnonymousInnerClassDeclarationContext) {
}

// EnterCreatedName is called when production createdName is entered.
func (s *BaseGroovyParserListener) EnterCreatedName(ctx *CreatedNameContext) {}

// ExitCreatedName is called when production createdName is exited.
func (s *BaseGroovyParserListener) ExitCreatedName(ctx *CreatedNameContext) {}

// EnterNonWildcardTypeArguments is called when production nonWildcardTypeArguments is entered.
func (s *BaseGroovyParserListener) EnterNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) {
}

// ExitNonWildcardTypeArguments is called when production nonWildcardTypeArguments is exited.
func (s *BaseGroovyParserListener) ExitNonWildcardTypeArguments(ctx *NonWildcardTypeArgumentsContext) {
}

// EnterTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is entered.
func (s *BaseGroovyParserListener) EnterTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// ExitTypeArgumentsOrDiamond is called when production typeArgumentsOrDiamond is exited.
func (s *BaseGroovyParserListener) ExitTypeArgumentsOrDiamond(ctx *TypeArgumentsOrDiamondContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGroovyParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGroovyParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseGroovyParserListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseGroovyParserListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterEnhancedArgumentList is called when production enhancedArgumentList is entered.
func (s *BaseGroovyParserListener) EnterEnhancedArgumentList(ctx *EnhancedArgumentListContext) {}

// ExitEnhancedArgumentList is called when production enhancedArgumentList is exited.
func (s *BaseGroovyParserListener) ExitEnhancedArgumentList(ctx *EnhancedArgumentListContext) {}

// EnterArgumentListElement is called when production argumentListElement is entered.
func (s *BaseGroovyParserListener) EnterArgumentListElement(ctx *ArgumentListElementContext) {}

// ExitArgumentListElement is called when production argumentListElement is exited.
func (s *BaseGroovyParserListener) ExitArgumentListElement(ctx *ArgumentListElementContext) {}

// EnterEnhancedArgumentListElement is called when production enhancedArgumentListElement is entered.
func (s *BaseGroovyParserListener) EnterEnhancedArgumentListElement(ctx *EnhancedArgumentListElementContext) {
}

// ExitEnhancedArgumentListElement is called when production enhancedArgumentListElement is exited.
func (s *BaseGroovyParserListener) ExitEnhancedArgumentListElement(ctx *EnhancedArgumentListElementContext) {
}

// EnterStringLiteral is called when production stringLiteral is entered.
func (s *BaseGroovyParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production stringLiteral is exited.
func (s *BaseGroovyParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterClassName is called when production className is entered.
func (s *BaseGroovyParserListener) EnterClassName(ctx *ClassNameContext) {}

// ExitClassName is called when production className is exited.
func (s *BaseGroovyParserListener) ExitClassName(ctx *ClassNameContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseGroovyParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseGroovyParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterBuiltInType is called when production builtInType is entered.
func (s *BaseGroovyParserListener) EnterBuiltInType(ctx *BuiltInTypeContext) {}

// ExitBuiltInType is called when production builtInType is exited.
func (s *BaseGroovyParserListener) ExitBuiltInType(ctx *BuiltInTypeContext) {}

// EnterKeywords is called when production keywords is entered.
func (s *BaseGroovyParserListener) EnterKeywords(ctx *KeywordsContext) {}

// ExitKeywords is called when production keywords is exited.
func (s *BaseGroovyParserListener) ExitKeywords(ctx *KeywordsContext) {}

// EnterRparen is called when production rparen is entered.
func (s *BaseGroovyParserListener) EnterRparen(ctx *RparenContext) {}

// ExitRparen is called when production rparen is exited.
func (s *BaseGroovyParserListener) ExitRparen(ctx *RparenContext) {}

// EnterNls is called when production nls is entered.
func (s *BaseGroovyParserListener) EnterNls(ctx *NlsContext) {}

// ExitNls is called when production nls is exited.
func (s *BaseGroovyParserListener) ExitNls(ctx *NlsContext) {}

// EnterSep is called when production sep is entered.
func (s *BaseGroovyParserListener) EnterSep(ctx *SepContext) {}

// ExitSep is called when production sep is exited.
func (s *BaseGroovyParserListener) ExitSep(ctx *SepContext) {}
