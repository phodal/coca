// Generated from TypeScriptParser.g4 by ANTLR 4.7.2
package tsantlr;

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TypeScriptParser extends TypeScriptBaseParser {
	static { RuntimeMetaData.checkVersion("4.7.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		HashBangLine=1, MultiLineComment=2, SingleLineComment=3, RegularExpressionLiteral=4, 
		OpenBracket=5, CloseBracket=6, OpenParen=7, CloseParen=8, OpenBrace=9, 
		CloseBrace=10, SemiColon=11, Comma=12, Assign=13, QuestionMark=14, Colon=15, 
		Ellipsis=16, Dot=17, PlusPlus=18, MinusMinus=19, Plus=20, Minus=21, BitNot=22, 
		Not=23, Multiply=24, Lodash=25, Dollar=26, Divide=27, Modulus=28, Power=29, 
		NullCoalesce=30, Hashtag=31, RightShiftArithmetic=32, LeftShiftArithmetic=33, 
		RightShiftLogical=34, LessThan=35, MoreThan=36, LessThanEquals=37, GreaterThanEquals=38, 
		Equals_=39, NotEquals=40, IdentityEquals=41, IdentityNotEquals=42, BitAnd=43, 
		BitXOr=44, BitOr=45, And=46, Or=47, MultiplyAssign=48, DivideAssign=49, 
		ModulusAssign=50, PlusAssign=51, MinusAssign=52, LeftShiftArithmeticAssign=53, 
		RightShiftArithmeticAssign=54, RightShiftLogicalAssign=55, BitAndAssign=56, 
		BitXorAssign=57, BitOrAssign=58, ARROW=59, PowerAssign=60, NullLiteral=61, 
		BooleanLiteral=62, DecimalLiteral=63, HexIntegerLiteral=64, OctalIntegerLiteral=65, 
		OctalIntegerLiteral2=66, BinaryIntegerLiteral=67, Break=68, Do=69, Instanceof=70, 
		Typeof=71, Case=72, Else=73, New=74, Var=75, Catch=76, Finally=77, Return=78, 
		Void=79, Continue=80, For=81, Switch=82, While=83, Debugger=84, Function=85, 
		This=86, With=87, Default=88, If=89, Throw=90, Delete=91, In=92, Try=93, 
		As=94, From=95, ReadOnly=96, Async=97, Class=98, Enum=99, Extends=100, 
		Super=101, Const=102, Export=103, Import=104, Await=105, Implements=106, 
		Let=107, Private=108, Public=109, Interface=110, Package=111, Protected=112, 
		Static=113, Yield=114, ANY=115, NUMBER=116, BOOLEAN=117, STRING=118, SYMBOL=119, 
		Type=120, Get=121, Set=122, Constructor=123, Namespace=124, Require=125, 
		Module=126, Declare=127, Abstract=128, Is=129, At=130, Identifier=131, 
		StringLiteral=132, TemplateStringLiteral=133, WhiteSpaces=134, LineTerminator=135, 
		HtmlComment=136, CDataComment=137, UnexpectedCharacter=138;
	public static final int
		RULE_initializer = 0, RULE_bindingPattern = 1, RULE_typeParameters = 2, 
		RULE_typeParameterList = 3, RULE_typeParameter = 4, RULE_constraint = 5, 
		RULE_typeArguments = 6, RULE_typeArgumentList = 7, RULE_typeArgument = 8, 
		RULE_type_ = 9, RULE_unionOrIntersectionOrPrimaryType = 10, RULE_primaryType = 11, 
		RULE_predefinedType = 12, RULE_typeReference = 13, RULE_typeGeneric = 14, 
		RULE_typeIncludeGeneric = 15, RULE_typeName = 16, RULE_objectType = 17, 
		RULE_typeBody = 18, RULE_typeMemberList = 19, RULE_typeMember = 20, RULE_arrayType = 21, 
		RULE_tupleType = 22, RULE_tupleElementTypes = 23, RULE_functionType = 24, 
		RULE_constructorType = 25, RULE_typeQuery = 26, RULE_typeQueryExpression = 27, 
		RULE_propertySignature = 28, RULE_typeAnnotation = 29, RULE_callSignature = 30, 
		RULE_parameterList = 31, RULE_requiredParameterList = 32, RULE_requiredParameter = 33, 
		RULE_accessibilityModifier = 34, RULE_identifierOrPattern = 35, RULE_optionalParameterList = 36, 
		RULE_optionalParameter = 37, RULE_restParameter = 38, RULE_constructSignature = 39, 
		RULE_indexSignature = 40, RULE_methodSignature = 41, RULE_typeAliasDeclaration = 42, 
		RULE_constructorDeclaration = 43, RULE_interfaceDeclaration = 44, RULE_interfaceExtendsClause = 45, 
		RULE_classOrInterfaceTypeList = 46, RULE_enumDeclaration = 47, RULE_enumBody = 48, 
		RULE_enumMemberList = 49, RULE_enumMember = 50, RULE_namespaceDeclaration = 51, 
		RULE_namespaceName = 52, RULE_importAliasDeclaration = 53, RULE_importAll = 54, 
		RULE_decoratorList = 55, RULE_decorator = 56, RULE_decoratorMemberExpression = 57, 
		RULE_decoratorCallExpression = 58, RULE_program = 59, RULE_sourceElement = 60, 
		RULE_statement = 61, RULE_block = 62, RULE_statementList = 63, RULE_abstractDeclaration = 64, 
		RULE_importStatement = 65, RULE_importFromBlock = 66, RULE_multipleImportStatement = 67, 
		RULE_exportStatement = 68, RULE_variableStatement = 69, RULE_variableDeclarationList = 70, 
		RULE_variableDeclaration = 71, RULE_emptyStatement_ = 72, RULE_expressionStatement = 73, 
		RULE_ifStatement = 74, RULE_iterationStatement = 75, RULE_varModifier = 76, 
		RULE_continueStatement = 77, RULE_breakStatement = 78, RULE_returnStatement = 79, 
		RULE_yieldStatement = 80, RULE_withStatement = 81, RULE_switchStatement = 82, 
		RULE_caseBlock = 83, RULE_caseClauses = 84, RULE_caseClause = 85, RULE_defaultClause = 86, 
		RULE_labelledStatement = 87, RULE_throwStatement = 88, RULE_tryStatement = 89, 
		RULE_catchProduction = 90, RULE_assignable = 91, RULE_finallyProduction = 92, 
		RULE_debuggerStatement = 93, RULE_functionDeclaration = 94, RULE_classDeclaration = 95, 
		RULE_classHeritage = 96, RULE_classTail = 97, RULE_classExtendsClause = 98, 
		RULE_implementsClause = 99, RULE_classElement = 100, RULE_propertyMemberDeclaration = 101, 
		RULE_propertyMemberBase = 102, RULE_indexMemberDeclaration = 103, RULE_generatorMethod = 104, 
		RULE_generatorFunctionDeclaration = 105, RULE_generatorBlock = 106, RULE_generatorDefinition = 107, 
		RULE_iteratorBlock = 108, RULE_iteratorDefinition = 109, RULE_formalParameterList = 110, 
		RULE_formalParameterArg = 111, RULE_lastFormalParameterArg = 112, RULE_functionBody = 113, 
		RULE_sourceElements = 114, RULE_arrayLiteral = 115, RULE_elementList = 116, 
		RULE_lastElement = 117, RULE_objectLiteral = 118, RULE_propertyAssignment = 119, 
		RULE_getAccessor = 120, RULE_setAccessor = 121, RULE_propertyName = 122, 
		RULE_arguments = 123, RULE_lastArgument = 124, RULE_expressionSequence = 125, 
		RULE_functionExpressionDeclaration = 126, RULE_singleExpression = 127, 
		RULE_arrowFunctionDeclaration = 128, RULE_arrowFunctionParameters = 129, 
		RULE_arrowFunctionBody = 130, RULE_assignmentOperator = 131, RULE_literal = 132, 
		RULE_numericLiteral = 133, RULE_identifierName = 134, RULE_reservedWord = 135, 
		RULE_keyword = 136, RULE_getter = 137, RULE_setter = 138, RULE_eos = 139;
	private static String[] makeRuleNames() {
		return new String[] {
			"initializer", "bindingPattern", "typeParameters", "typeParameterList", 
			"typeParameter", "constraint", "typeArguments", "typeArgumentList", "typeArgument", 
			"type_", "unionOrIntersectionOrPrimaryType", "primaryType", "predefinedType", 
			"typeReference", "typeGeneric", "typeIncludeGeneric", "typeName", "objectType", 
			"typeBody", "typeMemberList", "typeMember", "arrayType", "tupleType", 
			"tupleElementTypes", "functionType", "constructorType", "typeQuery", 
			"typeQueryExpression", "propertySignature", "typeAnnotation", "callSignature", 
			"parameterList", "requiredParameterList", "requiredParameter", "accessibilityModifier", 
			"identifierOrPattern", "optionalParameterList", "optionalParameter", 
			"restParameter", "constructSignature", "indexSignature", "methodSignature", 
			"typeAliasDeclaration", "constructorDeclaration", "interfaceDeclaration", 
			"interfaceExtendsClause", "classOrInterfaceTypeList", "enumDeclaration", 
			"enumBody", "enumMemberList", "enumMember", "namespaceDeclaration", "namespaceName", 
			"importAliasDeclaration", "importAll", "decoratorList", "decorator", 
			"decoratorMemberExpression", "decoratorCallExpression", "program", "sourceElement", 
			"statement", "block", "statementList", "abstractDeclaration", "importStatement", 
			"importFromBlock", "multipleImportStatement", "exportStatement", "variableStatement", 
			"variableDeclarationList", "variableDeclaration", "emptyStatement_", 
			"expressionStatement", "ifStatement", "iterationStatement", "varModifier", 
			"continueStatement", "breakStatement", "returnStatement", "yieldStatement", 
			"withStatement", "switchStatement", "caseBlock", "caseClauses", "caseClause", 
			"defaultClause", "labelledStatement", "throwStatement", "tryStatement", 
			"catchProduction", "assignable", "finallyProduction", "debuggerStatement", 
			"functionDeclaration", "classDeclaration", "classHeritage", "classTail", 
			"classExtendsClause", "implementsClause", "classElement", "propertyMemberDeclaration", 
			"propertyMemberBase", "indexMemberDeclaration", "generatorMethod", "generatorFunctionDeclaration", 
			"generatorBlock", "generatorDefinition", "iteratorBlock", "iteratorDefinition", 
			"formalParameterList", "formalParameterArg", "lastFormalParameterArg", 
			"functionBody", "sourceElements", "arrayLiteral", "elementList", "lastElement", 
			"objectLiteral", "propertyAssignment", "getAccessor", "setAccessor", 
			"propertyName", "arguments", "lastArgument", "expressionSequence", "functionExpressionDeclaration", 
			"singleExpression", "arrowFunctionDeclaration", "arrowFunctionParameters", 
			"arrowFunctionBody", "assignmentOperator", "literal", "numericLiteral", 
			"identifierName", "reservedWord", "keyword", "getter", "setter", "eos"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, "'['", "']'", "'('", "')'", "'{'", "'}'", 
			"';'", "','", "'='", "'?'", "':'", "'...'", "'.'", "'++'", "'--'", "'+'", 
			"'-'", "'~'", "'!'", "'*'", "'_'", "'$'", "'/'", "'%'", "'**'", "'??'", 
			"'#'", "'>>'", "'<<'", "'>>>'", "'<'", "'>'", "'<='", "'>='", "'=='", 
			"'!='", "'==='", "'!=='", "'&'", "'^'", "'|'", "'&&'", "'||'", "'*='", 
			"'/='", "'%='", "'+='", "'-='", "'<<='", "'>>='", "'>>>='", "'&='", "'^='", 
			"'|='", "'=>'", "'**='", "'null'", null, null, null, null, null, null, 
			"'break'", "'do'", "'instanceof'", "'typeof'", "'case'", "'else'", "'new'", 
			"'var'", "'catch'", "'finally'", "'return'", "'void'", "'continue'", 
			"'for'", "'switch'", "'while'", "'debugger'", "'function'", "'this'", 
			"'with'", "'default'", "'if'", "'throw'", "'delete'", "'in'", "'try'", 
			"'as'", "'from'", "'readonly'", "'async'", "'class'", "'enum'", "'extends'", 
			"'super'", "'const'", "'export'", "'import'", "'await'", "'implements'", 
			"'let'", "'private'", "'public'", "'interface'", "'package'", "'protected'", 
			"'static'", "'yield'", "'any'", "'number'", "'boolean'", "'string'", 
			"'symbol'", "'type'", "'get '", "'set '", "'constructor'", "'namespace'", 
			"'require'", "'module'", "'declare'", "'abstract'", "'is'", "'@'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "HashBangLine", "MultiLineComment", "SingleLineComment", "RegularExpressionLiteral", 
			"OpenBracket", "CloseBracket", "OpenParen", "CloseParen", "OpenBrace", 
			"CloseBrace", "SemiColon", "Comma", "Assign", "QuestionMark", "Colon", 
			"Ellipsis", "Dot", "PlusPlus", "MinusMinus", "Plus", "Minus", "BitNot", 
			"Not", "Multiply", "Lodash", "Dollar", "Divide", "Modulus", "Power", 
			"NullCoalesce", "Hashtag", "RightShiftArithmetic", "LeftShiftArithmetic", 
			"RightShiftLogical", "LessThan", "MoreThan", "LessThanEquals", "GreaterThanEquals", 
			"Equals_", "NotEquals", "IdentityEquals", "IdentityNotEquals", "BitAnd", 
			"BitXOr", "BitOr", "And", "Or", "MultiplyAssign", "DivideAssign", "ModulusAssign", 
			"PlusAssign", "MinusAssign", "LeftShiftArithmeticAssign", "RightShiftArithmeticAssign", 
			"RightShiftLogicalAssign", "BitAndAssign", "BitXorAssign", "BitOrAssign", 
			"ARROW", "PowerAssign", "NullLiteral", "BooleanLiteral", "DecimalLiteral", 
			"HexIntegerLiteral", "OctalIntegerLiteral", "OctalIntegerLiteral2", "BinaryIntegerLiteral", 
			"Break", "Do", "Instanceof", "Typeof", "Case", "Else", "New", "Var", 
			"Catch", "Finally", "Return", "Void", "Continue", "For", "Switch", "While", 
			"Debugger", "Function", "This", "With", "Default", "If", "Throw", "Delete", 
			"In", "Try", "As", "From", "ReadOnly", "Async", "Class", "Enum", "Extends", 
			"Super", "Const", "Export", "Import", "Await", "Implements", "Let", "Private", 
			"Public", "Interface", "Package", "Protected", "Static", "Yield", "ANY", 
			"NUMBER", "BOOLEAN", "STRING", "SYMBOL", "Type", "Get", "Set", "Constructor", 
			"Namespace", "Require", "Module", "Declare", "Abstract", "Is", "At", 
			"Identifier", "StringLiteral", "TemplateStringLiteral", "WhiteSpaces", 
			"LineTerminator", "HtmlComment", "CDataComment", "UnexpectedCharacter"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "TypeScriptParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public TypeScriptParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class InitializerContext extends ParserRuleContext {
		public TerminalNode Assign() { return getToken(TypeScriptParser.Assign, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public InitializerContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_initializer; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterInitializer(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitInitializer(this);
		}
	}

	public final InitializerContext initializer() throws RecognitionException {
		InitializerContext _localctx = new InitializerContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_initializer);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(280);
			match(Assign);
			setState(281);
			singleExpression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BindingPatternContext extends ParserRuleContext {
		public ArrayLiteralContext arrayLiteral() {
			return getRuleContext(ArrayLiteralContext.class,0);
		}
		public ObjectLiteralContext objectLiteral() {
			return getRuleContext(ObjectLiteralContext.class,0);
		}
		public BindingPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bindingPattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterBindingPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitBindingPattern(this);
		}
	}

	public final BindingPatternContext bindingPattern() throws RecognitionException {
		BindingPatternContext _localctx = new BindingPatternContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_bindingPattern);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OpenBracket:
				{
				setState(283);
				arrayLiteral();
				}
				break;
			case OpenBrace:
				{
				setState(284);
				objectLiteral();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeParametersContext extends ParserRuleContext {
		public TerminalNode LessThan() { return getToken(TypeScriptParser.LessThan, 0); }
		public TerminalNode MoreThan() { return getToken(TypeScriptParser.MoreThan, 0); }
		public TypeParameterListContext typeParameterList() {
			return getRuleContext(TypeParameterListContext.class,0);
		}
		public TypeParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeParameters; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeParameters(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeParameters(this);
		}
	}

	public final TypeParametersContext typeParameters() throws RecognitionException {
		TypeParametersContext _localctx = new TypeParametersContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_typeParameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(287);
			match(LessThan);
			setState(289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LessThan || _la==Identifier) {
				{
				setState(288);
				typeParameterList();
				}
			}

			setState(291);
			match(MoreThan);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeParameterListContext extends ParserRuleContext {
		public List<TypeParameterContext> typeParameter() {
			return getRuleContexts(TypeParameterContext.class);
		}
		public TypeParameterContext typeParameter(int i) {
			return getRuleContext(TypeParameterContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public TypeParameterListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeParameterList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeParameterList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeParameterList(this);
		}
	}

	public final TypeParameterListContext typeParameterList() throws RecognitionException {
		TypeParameterListContext _localctx = new TypeParameterListContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_typeParameterList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(293);
			typeParameter();
			setState(298);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Comma) {
				{
				{
				setState(294);
				match(Comma);
				setState(295);
				typeParameter();
				}
				}
				setState(300);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeParameterContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public ConstraintContext constraint() {
			return getRuleContext(ConstraintContext.class,0);
		}
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public TypeParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeParameter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeParameter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeParameter(this);
		}
	}

	public final TypeParameterContext typeParameter() throws RecognitionException {
		TypeParameterContext _localctx = new TypeParameterContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_typeParameter);
		int _la;
		try {
			setState(306);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(301);
				match(Identifier);
				setState(303);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Extends) {
					{
					setState(302);
					constraint();
					}
				}

				}
				break;
			case LessThan:
				enterOuterAlt(_localctx, 2);
				{
				setState(305);
				typeParameters();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstraintContext extends ParserRuleContext {
		public TerminalNode Extends() { return getToken(TypeScriptParser.Extends, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public ConstraintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constraint; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterConstraint(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitConstraint(this);
		}
	}

	public final ConstraintContext constraint() throws RecognitionException {
		ConstraintContext _localctx = new ConstraintContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_constraint);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(308);
			match(Extends);
			setState(309);
			type_();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeArgumentsContext extends ParserRuleContext {
		public TerminalNode LessThan() { return getToken(TypeScriptParser.LessThan, 0); }
		public TerminalNode MoreThan() { return getToken(TypeScriptParser.MoreThan, 0); }
		public TypeArgumentListContext typeArgumentList() {
			return getRuleContext(TypeArgumentListContext.class,0);
		}
		public TypeArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeArguments; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeArguments(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeArguments(this);
		}
	}

	public final TypeArgumentsContext typeArguments() throws RecognitionException {
		TypeArgumentsContext _localctx = new TypeArgumentsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_typeArguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(311);
			match(LessThan);
			setState(313);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenParen) | (1L << OpenBrace) | (1L << LessThan))) != 0) || ((((_la - 71)) & ~0x3f) == 0 && ((1L << (_la - 71)) & ((1L << (Typeof - 71)) | (1L << (New - 71)) | (1L << (Void - 71)) | (1L << (This - 71)) | (1L << (ANY - 71)) | (1L << (NUMBER - 71)) | (1L << (BOOLEAN - 71)) | (1L << (STRING - 71)) | (1L << (SYMBOL - 71)) | (1L << (Identifier - 71)) | (1L << (StringLiteral - 71)))) != 0)) {
				{
				setState(312);
				typeArgumentList();
				}
			}

			setState(315);
			match(MoreThan);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeArgumentListContext extends ParserRuleContext {
		public List<TypeArgumentContext> typeArgument() {
			return getRuleContexts(TypeArgumentContext.class);
		}
		public TypeArgumentContext typeArgument(int i) {
			return getRuleContext(TypeArgumentContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public TypeArgumentListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeArgumentList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeArgumentList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeArgumentList(this);
		}
	}

	public final TypeArgumentListContext typeArgumentList() throws RecognitionException {
		TypeArgumentListContext _localctx = new TypeArgumentListContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_typeArgumentList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(317);
			typeArgument();
			setState(322);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Comma) {
				{
				{
				setState(318);
				match(Comma);
				setState(319);
				typeArgument();
				}
				}
				setState(324);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeArgumentContext extends ParserRuleContext {
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TypeArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeArgument; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeArgument(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeArgument(this);
		}
	}

	public final TypeArgumentContext typeArgument() throws RecognitionException {
		TypeArgumentContext _localctx = new TypeArgumentContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_typeArgument);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(325);
			type_();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Type_Context extends ParserRuleContext {
		public UnionOrIntersectionOrPrimaryTypeContext unionOrIntersectionOrPrimaryType() {
			return getRuleContext(UnionOrIntersectionOrPrimaryTypeContext.class,0);
		}
		public FunctionTypeContext functionType() {
			return getRuleContext(FunctionTypeContext.class,0);
		}
		public ConstructorTypeContext constructorType() {
			return getRuleContext(ConstructorTypeContext.class,0);
		}
		public TypeGenericContext typeGeneric() {
			return getRuleContext(TypeGenericContext.class,0);
		}
		public TerminalNode StringLiteral() { return getToken(TypeScriptParser.StringLiteral, 0); }
		public Type_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterType_(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitType_(this);
		}
	}

	public final Type_Context type_() throws RecognitionException {
		Type_Context _localctx = new Type_Context(_ctx, getState());
		enterRule(_localctx, 18, RULE_type_);
		try {
			setState(332);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(327);
				unionOrIntersectionOrPrimaryType(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(328);
				functionType();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(329);
				constructorType();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(330);
				typeGeneric();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(331);
				match(StringLiteral);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UnionOrIntersectionOrPrimaryTypeContext extends ParserRuleContext {
		public UnionOrIntersectionOrPrimaryTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unionOrIntersectionOrPrimaryType; }
	 
		public UnionOrIntersectionOrPrimaryTypeContext() { }
		public void copyFrom(UnionOrIntersectionOrPrimaryTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class IntersectionContext extends UnionOrIntersectionOrPrimaryTypeContext {
		public List<UnionOrIntersectionOrPrimaryTypeContext> unionOrIntersectionOrPrimaryType() {
			return getRuleContexts(UnionOrIntersectionOrPrimaryTypeContext.class);
		}
		public UnionOrIntersectionOrPrimaryTypeContext unionOrIntersectionOrPrimaryType(int i) {
			return getRuleContext(UnionOrIntersectionOrPrimaryTypeContext.class,i);
		}
		public TerminalNode BitAnd() { return getToken(TypeScriptParser.BitAnd, 0); }
		public IntersectionContext(UnionOrIntersectionOrPrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIntersection(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIntersection(this);
		}
	}
	public static class PrimaryContext extends UnionOrIntersectionOrPrimaryTypeContext {
		public PrimaryTypeContext primaryType() {
			return getRuleContext(PrimaryTypeContext.class,0);
		}
		public PrimaryContext(UnionOrIntersectionOrPrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPrimary(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPrimary(this);
		}
	}
	public static class UnionContext extends UnionOrIntersectionOrPrimaryTypeContext {
		public List<UnionOrIntersectionOrPrimaryTypeContext> unionOrIntersectionOrPrimaryType() {
			return getRuleContexts(UnionOrIntersectionOrPrimaryTypeContext.class);
		}
		public UnionOrIntersectionOrPrimaryTypeContext unionOrIntersectionOrPrimaryType(int i) {
			return getRuleContext(UnionOrIntersectionOrPrimaryTypeContext.class,i);
		}
		public TerminalNode BitOr() { return getToken(TypeScriptParser.BitOr, 0); }
		public UnionContext(UnionOrIntersectionOrPrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterUnion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitUnion(this);
		}
	}

	public final UnionOrIntersectionOrPrimaryTypeContext unionOrIntersectionOrPrimaryType() throws RecognitionException {
		return unionOrIntersectionOrPrimaryType(0);
	}

	private UnionOrIntersectionOrPrimaryTypeContext unionOrIntersectionOrPrimaryType(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		UnionOrIntersectionOrPrimaryTypeContext _localctx = new UnionOrIntersectionOrPrimaryTypeContext(_ctx, _parentState);
		UnionOrIntersectionOrPrimaryTypeContext _prevctx = _localctx;
		int _startState = 20;
		enterRecursionRule(_localctx, 20, RULE_unionOrIntersectionOrPrimaryType, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new PrimaryContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(335);
			primaryType(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(345);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(343);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new UnionContext(new UnionOrIntersectionOrPrimaryTypeContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_unionOrIntersectionOrPrimaryType);
						setState(337);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(338);
						match(BitOr);
						setState(339);
						unionOrIntersectionOrPrimaryType(4);
						}
						break;
					case 2:
						{
						_localctx = new IntersectionContext(new UnionOrIntersectionOrPrimaryTypeContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_unionOrIntersectionOrPrimaryType);
						setState(340);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(341);
						match(BitAnd);
						setState(342);
						unionOrIntersectionOrPrimaryType(3);
						}
						break;
					}
					} 
				}
				setState(347);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PrimaryTypeContext extends ParserRuleContext {
		public PrimaryTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primaryType; }
	 
		public PrimaryTypeContext() { }
		public void copyFrom(PrimaryTypeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class RedefinitionOfTypeContext extends PrimaryTypeContext {
		public TypeReferenceContext typeReference() {
			return getRuleContext(TypeReferenceContext.class,0);
		}
		public TerminalNode Is() { return getToken(TypeScriptParser.Is, 0); }
		public PrimaryTypeContext primaryType() {
			return getRuleContext(PrimaryTypeContext.class,0);
		}
		public RedefinitionOfTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterRedefinitionOfType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitRedefinitionOfType(this);
		}
	}
	public static class PredefinedPrimTypeContext extends PrimaryTypeContext {
		public PredefinedTypeContext predefinedType() {
			return getRuleContext(PredefinedTypeContext.class,0);
		}
		public PredefinedPrimTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPredefinedPrimType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPredefinedPrimType(this);
		}
	}
	public static class ArrayPrimTypeContext extends PrimaryTypeContext {
		public PrimaryTypeContext primaryType() {
			return getRuleContext(PrimaryTypeContext.class,0);
		}
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public ArrayPrimTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArrayPrimType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArrayPrimType(this);
		}
	}
	public static class ParenthesizedPrimTypeContext extends PrimaryTypeContext {
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public ParenthesizedPrimTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterParenthesizedPrimType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitParenthesizedPrimType(this);
		}
	}
	public static class ThisPrimTypeContext extends PrimaryTypeContext {
		public TerminalNode This() { return getToken(TypeScriptParser.This, 0); }
		public ThisPrimTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterThisPrimType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitThisPrimType(this);
		}
	}
	public static class TuplePrimTypeContext extends PrimaryTypeContext {
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public TupleElementTypesContext tupleElementTypes() {
			return getRuleContext(TupleElementTypesContext.class,0);
		}
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public TuplePrimTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTuplePrimType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTuplePrimType(this);
		}
	}
	public static class ObjectPrimTypeContext extends PrimaryTypeContext {
		public ObjectTypeContext objectType() {
			return getRuleContext(ObjectTypeContext.class,0);
		}
		public ObjectPrimTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterObjectPrimType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitObjectPrimType(this);
		}
	}
	public static class ReferencePrimTypeContext extends PrimaryTypeContext {
		public TypeReferenceContext typeReference() {
			return getRuleContext(TypeReferenceContext.class,0);
		}
		public ReferencePrimTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterReferencePrimType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitReferencePrimType(this);
		}
	}
	public static class QueryPrimTypeContext extends PrimaryTypeContext {
		public TypeQueryContext typeQuery() {
			return getRuleContext(TypeQueryContext.class,0);
		}
		public QueryPrimTypeContext(PrimaryTypeContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterQueryPrimType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitQueryPrimType(this);
		}
	}

	public final PrimaryTypeContext primaryType() throws RecognitionException {
		return primaryType(0);
	}

	private PrimaryTypeContext primaryType(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		PrimaryTypeContext _localctx = new PrimaryTypeContext(_ctx, _parentState);
		PrimaryTypeContext _prevctx = _localctx;
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_primaryType, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(366);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				{
				_localctx = new ParenthesizedPrimTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(349);
				match(OpenParen);
				setState(350);
				type_();
				setState(351);
				match(CloseParen);
				}
				break;
			case 2:
				{
				_localctx = new PredefinedPrimTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(353);
				predefinedType();
				}
				break;
			case 3:
				{
				_localctx = new ReferencePrimTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(354);
				typeReference();
				}
				break;
			case 4:
				{
				_localctx = new ObjectPrimTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(355);
				objectType();
				}
				break;
			case 5:
				{
				_localctx = new TuplePrimTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(356);
				match(OpenBracket);
				setState(357);
				tupleElementTypes();
				setState(358);
				match(CloseBracket);
				}
				break;
			case 6:
				{
				_localctx = new QueryPrimTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(360);
				typeQuery();
				}
				break;
			case 7:
				{
				_localctx = new ThisPrimTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(361);
				match(This);
				}
				break;
			case 8:
				{
				_localctx = new RedefinitionOfTypeContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(362);
				typeReference();
				setState(363);
				match(Is);
				setState(364);
				primaryType(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(374);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ArrayPrimTypeContext(new PrimaryTypeContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_primaryType);
					setState(368);
					if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
					setState(369);
					if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
					setState(370);
					match(OpenBracket);
					setState(371);
					match(CloseBracket);
					}
					} 
				}
				setState(376);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PredefinedTypeContext extends ParserRuleContext {
		public TerminalNode ANY() { return getToken(TypeScriptParser.ANY, 0); }
		public TerminalNode NUMBER() { return getToken(TypeScriptParser.NUMBER, 0); }
		public TerminalNode BOOLEAN() { return getToken(TypeScriptParser.BOOLEAN, 0); }
		public TerminalNode STRING() { return getToken(TypeScriptParser.STRING, 0); }
		public TerminalNode SYMBOL() { return getToken(TypeScriptParser.SYMBOL, 0); }
		public TerminalNode Void() { return getToken(TypeScriptParser.Void, 0); }
		public PredefinedTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_predefinedType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPredefinedType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPredefinedType(this);
		}
	}

	public final PredefinedTypeContext predefinedType() throws RecognitionException {
		PredefinedTypeContext _localctx = new PredefinedTypeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_predefinedType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(377);
			_la = _input.LA(1);
			if ( !(((((_la - 79)) & ~0x3f) == 0 && ((1L << (_la - 79)) & ((1L << (Void - 79)) | (1L << (ANY - 79)) | (1L << (NUMBER - 79)) | (1L << (BOOLEAN - 79)) | (1L << (STRING - 79)) | (1L << (SYMBOL - 79)))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeReferenceContext extends ParserRuleContext {
		public TypeNameContext typeName() {
			return getRuleContext(TypeNameContext.class,0);
		}
		public TypeIncludeGenericContext typeIncludeGeneric() {
			return getRuleContext(TypeIncludeGenericContext.class,0);
		}
		public TypeGenericContext typeGeneric() {
			return getRuleContext(TypeGenericContext.class,0);
		}
		public TypeReferenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeReference; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeReference(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeReference(this);
		}
	}

	public final TypeReferenceContext typeReference() throws RecognitionException {
		TypeReferenceContext _localctx = new TypeReferenceContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_typeReference);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(379);
			typeName();
			setState(382);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(380);
				typeIncludeGeneric();
				}
				break;
			case 2:
				{
				setState(381);
				typeGeneric();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeGenericContext extends ParserRuleContext {
		public TerminalNode LessThan() { return getToken(TypeScriptParser.LessThan, 0); }
		public TypeArgumentListContext typeArgumentList() {
			return getRuleContext(TypeArgumentListContext.class,0);
		}
		public TerminalNode MoreThan() { return getToken(TypeScriptParser.MoreThan, 0); }
		public TypeGenericContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeGeneric; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeGeneric(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeGeneric(this);
		}
	}

	public final TypeGenericContext typeGeneric() throws RecognitionException {
		TypeGenericContext _localctx = new TypeGenericContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_typeGeneric);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(384);
			match(LessThan);
			setState(385);
			typeArgumentList();
			setState(386);
			match(MoreThan);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeIncludeGenericContext extends ParserRuleContext {
		public List<TerminalNode> LessThan() { return getTokens(TypeScriptParser.LessThan); }
		public TerminalNode LessThan(int i) {
			return getToken(TypeScriptParser.LessThan, i);
		}
		public List<TypeArgumentListContext> typeArgumentList() {
			return getRuleContexts(TypeArgumentListContext.class);
		}
		public TypeArgumentListContext typeArgumentList(int i) {
			return getRuleContext(TypeArgumentListContext.class,i);
		}
		public List<TerminalNode> MoreThan() { return getTokens(TypeScriptParser.MoreThan); }
		public TerminalNode MoreThan(int i) {
			return getToken(TypeScriptParser.MoreThan, i);
		}
		public BindingPatternContext bindingPattern() {
			return getRuleContext(BindingPatternContext.class,0);
		}
		public TerminalNode RightShiftArithmetic() { return getToken(TypeScriptParser.RightShiftArithmetic, 0); }
		public TypeIncludeGenericContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeIncludeGeneric; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeIncludeGeneric(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeIncludeGeneric(this);
		}
	}

	public final TypeIncludeGenericContext typeIncludeGeneric() throws RecognitionException {
		TypeIncludeGenericContext _localctx = new TypeIncludeGenericContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_typeIncludeGeneric);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(388);
			match(LessThan);
			setState(389);
			typeArgumentList();
			setState(390);
			match(LessThan);
			setState(391);
			typeArgumentList();
			setState(397);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MoreThan:
				{
				setState(392);
				match(MoreThan);
				setState(393);
				bindingPattern();
				setState(394);
				match(MoreThan);
				}
				break;
			case RightShiftArithmetic:
				{
				setState(396);
				match(RightShiftArithmetic);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeNameContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public NamespaceNameContext namespaceName() {
			return getRuleContext(NamespaceNameContext.class,0);
		}
		public TypeNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeName(this);
		}
	}

	public final TypeNameContext typeName() throws RecognitionException {
		TypeNameContext _localctx = new TypeNameContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_typeName);
		try {
			setState(401);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(399);
				match(Identifier);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(400);
				namespaceName();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjectTypeContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public TypeBodyContext typeBody() {
			return getRuleContext(TypeBodyContext.class,0);
		}
		public ObjectTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterObjectType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitObjectType(this);
		}
	}

	public final ObjectTypeContext objectType() throws RecognitionException {
		ObjectTypeContext _localctx = new ObjectTypeContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_objectType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(403);
			match(OpenBrace);
			setState(405);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & ((1L << (OpenBracket - 5)) | (1L << (OpenParen - 5)) | (1L << (LessThan - 5)) | (1L << (NullLiteral - 5)) | (1L << (BooleanLiteral - 5)) | (1L << (DecimalLiteral - 5)) | (1L << (HexIntegerLiteral - 5)) | (1L << (OctalIntegerLiteral - 5)) | (1L << (OctalIntegerLiteral2 - 5)) | (1L << (BinaryIntegerLiteral - 5)) | (1L << (Break - 5)))) != 0) || ((((_la - 69)) & ~0x3f) == 0 && ((1L << (_la - 69)) & ((1L << (Do - 69)) | (1L << (Instanceof - 69)) | (1L << (Typeof - 69)) | (1L << (Case - 69)) | (1L << (Else - 69)) | (1L << (New - 69)) | (1L << (Var - 69)) | (1L << (Catch - 69)) | (1L << (Finally - 69)) | (1L << (Return - 69)) | (1L << (Void - 69)) | (1L << (Continue - 69)) | (1L << (For - 69)) | (1L << (Switch - 69)) | (1L << (While - 69)) | (1L << (Debugger - 69)) | (1L << (Function - 69)) | (1L << (This - 69)) | (1L << (With - 69)) | (1L << (Default - 69)) | (1L << (If - 69)) | (1L << (Throw - 69)) | (1L << (Delete - 69)) | (1L << (In - 69)) | (1L << (Try - 69)) | (1L << (As - 69)) | (1L << (From - 69)) | (1L << (ReadOnly - 69)) | (1L << (Async - 69)) | (1L << (Class - 69)) | (1L << (Enum - 69)) | (1L << (Extends - 69)) | (1L << (Super - 69)) | (1L << (Const - 69)) | (1L << (Export - 69)) | (1L << (Import - 69)) | (1L << (Await - 69)) | (1L << (Implements - 69)) | (1L << (Let - 69)) | (1L << (Private - 69)) | (1L << (Public - 69)) | (1L << (Interface - 69)) | (1L << (Package - 69)) | (1L << (Protected - 69)) | (1L << (Static - 69)) | (1L << (Yield - 69)) | (1L << (Identifier - 69)) | (1L << (StringLiteral - 69)))) != 0)) {
				{
				setState(404);
				typeBody();
				}
			}

			setState(407);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeBodyContext extends ParserRuleContext {
		public TypeMemberListContext typeMemberList() {
			return getRuleContext(TypeMemberListContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public TerminalNode Comma() { return getToken(TypeScriptParser.Comma, 0); }
		public TypeBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeBody; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeBody(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeBody(this);
		}
	}

	public final TypeBodyContext typeBody() throws RecognitionException {
		TypeBodyContext _localctx = new TypeBodyContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_typeBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(409);
			typeMemberList();
			setState(411);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SemiColon || _la==Comma) {
				{
				setState(410);
				_la = _input.LA(1);
				if ( !(_la==SemiColon || _la==Comma) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeMemberListContext extends ParserRuleContext {
		public List<TypeMemberContext> typeMember() {
			return getRuleContexts(TypeMemberContext.class);
		}
		public TypeMemberContext typeMember(int i) {
			return getRuleContext(TypeMemberContext.class,i);
		}
		public List<TerminalNode> SemiColon() { return getTokens(TypeScriptParser.SemiColon); }
		public TerminalNode SemiColon(int i) {
			return getToken(TypeScriptParser.SemiColon, i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public TypeMemberListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeMemberList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeMemberList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeMemberList(this);
		}
	}

	public final TypeMemberListContext typeMemberList() throws RecognitionException {
		TypeMemberListContext _localctx = new TypeMemberListContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_typeMemberList);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(413);
			typeMember();
			setState(418);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(414);
					_la = _input.LA(1);
					if ( !(_la==SemiColon || _la==Comma) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(415);
					typeMember();
					}
					} 
				}
				setState(420);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeMemberContext extends ParserRuleContext {
		public PropertySignatureContext propertySignature() {
			return getRuleContext(PropertySignatureContext.class,0);
		}
		public CallSignatureContext callSignature() {
			return getRuleContext(CallSignatureContext.class,0);
		}
		public ConstructSignatureContext constructSignature() {
			return getRuleContext(ConstructSignatureContext.class,0);
		}
		public IndexSignatureContext indexSignature() {
			return getRuleContext(IndexSignatureContext.class,0);
		}
		public MethodSignatureContext methodSignature() {
			return getRuleContext(MethodSignatureContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(TypeScriptParser.ARROW, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TypeMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeMember; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeMember(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeMember(this);
		}
	}

	public final TypeMemberContext typeMember() throws RecognitionException {
		TypeMemberContext _localctx = new TypeMemberContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_typeMember);
		int _la;
		try {
			setState(430);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(421);
				propertySignature();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(422);
				callSignature();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(423);
				constructSignature();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(424);
				indexSignature();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(425);
				methodSignature();
				setState(428);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ARROW) {
					{
					setState(426);
					match(ARROW);
					setState(427);
					type_();
					}
				}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrayTypeContext extends ParserRuleContext {
		public PrimaryTypeContext primaryType() {
			return getRuleContext(PrimaryTypeContext.class,0);
		}
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public ArrayTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArrayType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArrayType(this);
		}
	}

	public final ArrayTypeContext arrayType() throws RecognitionException {
		ArrayTypeContext _localctx = new ArrayTypeContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_arrayType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(432);
			primaryType(0);
			setState(433);
			if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
			setState(434);
			match(OpenBracket);
			setState(435);
			match(CloseBracket);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TupleTypeContext extends ParserRuleContext {
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public TupleElementTypesContext tupleElementTypes() {
			return getRuleContext(TupleElementTypesContext.class,0);
		}
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public TupleTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTupleType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTupleType(this);
		}
	}

	public final TupleTypeContext tupleType() throws RecognitionException {
		TupleTypeContext _localctx = new TupleTypeContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_tupleType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(437);
			match(OpenBracket);
			setState(438);
			tupleElementTypes();
			setState(439);
			match(CloseBracket);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TupleElementTypesContext extends ParserRuleContext {
		public List<Type_Context> type_() {
			return getRuleContexts(Type_Context.class);
		}
		public Type_Context type_(int i) {
			return getRuleContext(Type_Context.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public TupleElementTypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tupleElementTypes; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTupleElementTypes(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTupleElementTypes(this);
		}
	}

	public final TupleElementTypesContext tupleElementTypes() throws RecognitionException {
		TupleElementTypesContext _localctx = new TupleElementTypesContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_tupleElementTypes);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(441);
			type_();
			setState(446);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Comma) {
				{
				{
				setState(442);
				match(Comma);
				setState(443);
				type_();
				}
				}
				setState(448);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionTypeContext extends ParserRuleContext {
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode ARROW() { return getToken(TypeScriptParser.ARROW, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public ParameterListContext parameterList() {
			return getRuleContext(ParameterListContext.class,0);
		}
		public FunctionTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFunctionType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFunctionType(this);
		}
	}

	public final FunctionTypeContext functionType() throws RecognitionException {
		FunctionTypeContext _localctx = new FunctionTypeContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_functionType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(450);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LessThan) {
				{
				setState(449);
				typeParameters();
				}
			}

			setState(452);
			match(OpenParen);
			setState(454);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis) | (1L << NullLiteral) | (1L << BooleanLiteral))) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & ((1L << (Break - 68)) | (1L << (Do - 68)) | (1L << (Instanceof - 68)) | (1L << (Typeof - 68)) | (1L << (Case - 68)) | (1L << (Else - 68)) | (1L << (New - 68)) | (1L << (Var - 68)) | (1L << (Catch - 68)) | (1L << (Finally - 68)) | (1L << (Return - 68)) | (1L << (Void - 68)) | (1L << (Continue - 68)) | (1L << (For - 68)) | (1L << (Switch - 68)) | (1L << (While - 68)) | (1L << (Debugger - 68)) | (1L << (Function - 68)) | (1L << (This - 68)) | (1L << (With - 68)) | (1L << (Default - 68)) | (1L << (If - 68)) | (1L << (Throw - 68)) | (1L << (Delete - 68)) | (1L << (In - 68)) | (1L << (Try - 68)) | (1L << (As - 68)) | (1L << (From - 68)) | (1L << (ReadOnly - 68)) | (1L << (Async - 68)) | (1L << (Class - 68)) | (1L << (Enum - 68)) | (1L << (Extends - 68)) | (1L << (Super - 68)) | (1L << (Const - 68)) | (1L << (Export - 68)) | (1L << (Import - 68)) | (1L << (Await - 68)) | (1L << (Implements - 68)) | (1L << (Let - 68)) | (1L << (Private - 68)) | (1L << (Public - 68)) | (1L << (Interface - 68)) | (1L << (Package - 68)) | (1L << (Protected - 68)) | (1L << (Static - 68)) | (1L << (Yield - 68)) | (1L << (ANY - 68)) | (1L << (NUMBER - 68)) | (1L << (BOOLEAN - 68)) | (1L << (STRING - 68)) | (1L << (SYMBOL - 68)) | (1L << (At - 68)) | (1L << (Identifier - 68)))) != 0)) {
				{
				setState(453);
				parameterList();
				}
			}

			setState(456);
			match(CloseParen);
			setState(457);
			match(ARROW);
			setState(458);
			type_();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstructorTypeContext extends ParserRuleContext {
		public TerminalNode New() { return getToken(TypeScriptParser.New, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode ARROW() { return getToken(TypeScriptParser.ARROW, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public ParameterListContext parameterList() {
			return getRuleContext(ParameterListContext.class,0);
		}
		public ConstructorTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constructorType; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterConstructorType(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitConstructorType(this);
		}
	}

	public final ConstructorTypeContext constructorType() throws RecognitionException {
		ConstructorTypeContext _localctx = new ConstructorTypeContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_constructorType);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(460);
			match(New);
			setState(462);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LessThan) {
				{
				setState(461);
				typeParameters();
				}
			}

			setState(464);
			match(OpenParen);
			setState(466);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis) | (1L << NullLiteral) | (1L << BooleanLiteral))) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & ((1L << (Break - 68)) | (1L << (Do - 68)) | (1L << (Instanceof - 68)) | (1L << (Typeof - 68)) | (1L << (Case - 68)) | (1L << (Else - 68)) | (1L << (New - 68)) | (1L << (Var - 68)) | (1L << (Catch - 68)) | (1L << (Finally - 68)) | (1L << (Return - 68)) | (1L << (Void - 68)) | (1L << (Continue - 68)) | (1L << (For - 68)) | (1L << (Switch - 68)) | (1L << (While - 68)) | (1L << (Debugger - 68)) | (1L << (Function - 68)) | (1L << (This - 68)) | (1L << (With - 68)) | (1L << (Default - 68)) | (1L << (If - 68)) | (1L << (Throw - 68)) | (1L << (Delete - 68)) | (1L << (In - 68)) | (1L << (Try - 68)) | (1L << (As - 68)) | (1L << (From - 68)) | (1L << (ReadOnly - 68)) | (1L << (Async - 68)) | (1L << (Class - 68)) | (1L << (Enum - 68)) | (1L << (Extends - 68)) | (1L << (Super - 68)) | (1L << (Const - 68)) | (1L << (Export - 68)) | (1L << (Import - 68)) | (1L << (Await - 68)) | (1L << (Implements - 68)) | (1L << (Let - 68)) | (1L << (Private - 68)) | (1L << (Public - 68)) | (1L << (Interface - 68)) | (1L << (Package - 68)) | (1L << (Protected - 68)) | (1L << (Static - 68)) | (1L << (Yield - 68)) | (1L << (ANY - 68)) | (1L << (NUMBER - 68)) | (1L << (BOOLEAN - 68)) | (1L << (STRING - 68)) | (1L << (SYMBOL - 68)) | (1L << (At - 68)) | (1L << (Identifier - 68)))) != 0)) {
				{
				setState(465);
				parameterList();
				}
			}

			setState(468);
			match(CloseParen);
			setState(469);
			match(ARROW);
			setState(470);
			type_();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeQueryContext extends ParserRuleContext {
		public TerminalNode Typeof() { return getToken(TypeScriptParser.Typeof, 0); }
		public TypeQueryExpressionContext typeQueryExpression() {
			return getRuleContext(TypeQueryExpressionContext.class,0);
		}
		public TypeQueryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeQuery; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeQuery(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeQuery(this);
		}
	}

	public final TypeQueryContext typeQuery() throws RecognitionException {
		TypeQueryContext _localctx = new TypeQueryContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_typeQuery);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(472);
			match(Typeof);
			setState(473);
			typeQueryExpression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeQueryExpressionContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public List<IdentifierNameContext> identifierName() {
			return getRuleContexts(IdentifierNameContext.class);
		}
		public IdentifierNameContext identifierName(int i) {
			return getRuleContext(IdentifierNameContext.class,i);
		}
		public List<TerminalNode> Dot() { return getTokens(TypeScriptParser.Dot); }
		public TerminalNode Dot(int i) {
			return getToken(TypeScriptParser.Dot, i);
		}
		public TypeQueryExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeQueryExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeQueryExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeQueryExpression(this);
		}
	}

	public final TypeQueryExpressionContext typeQueryExpression() throws RecognitionException {
		TypeQueryExpressionContext _localctx = new TypeQueryExpressionContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_typeQueryExpression);
		try {
			int _alt;
			setState(485);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(475);
				match(Identifier);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(479); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(476);
						identifierName();
						setState(477);
						match(Dot);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(481); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				setState(483);
				identifierName();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PropertySignatureContext extends ParserRuleContext {
		public PropertyNameContext propertyName() {
			return getRuleContext(PropertyNameContext.class,0);
		}
		public TerminalNode ReadOnly() { return getToken(TypeScriptParser.ReadOnly, 0); }
		public TerminalNode QuestionMark() { return getToken(TypeScriptParser.QuestionMark, 0); }
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(TypeScriptParser.ARROW, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public PropertySignatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertySignature; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPropertySignature(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPropertySignature(this);
		}
	}

	public final PropertySignatureContext propertySignature() throws RecognitionException {
		PropertySignatureContext _localctx = new PropertySignatureContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_propertySignature);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(488);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				{
				setState(487);
				match(ReadOnly);
				}
				break;
			}
			setState(490);
			propertyName();
			setState(492);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==QuestionMark) {
				{
				setState(491);
				match(QuestionMark);
				}
			}

			setState(495);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(494);
				typeAnnotation();
				}
			}

			setState(499);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ARROW) {
				{
				setState(497);
				match(ARROW);
				setState(498);
				type_();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeAnnotationContext extends ParserRuleContext {
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TypeAnnotationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAnnotation; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeAnnotation(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeAnnotation(this);
		}
	}

	public final TypeAnnotationContext typeAnnotation() throws RecognitionException {
		TypeAnnotationContext _localctx = new TypeAnnotationContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_typeAnnotation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(501);
			match(Colon);
			setState(502);
			type_();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CallSignatureContext extends ParserRuleContext {
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public ParameterListContext parameterList() {
			return getRuleContext(ParameterListContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public CallSignatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_callSignature; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterCallSignature(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitCallSignature(this);
		}
	}

	public final CallSignatureContext callSignature() throws RecognitionException {
		CallSignatureContext _localctx = new CallSignatureContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_callSignature);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(505);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LessThan) {
				{
				setState(504);
				typeParameters();
				}
			}

			setState(507);
			match(OpenParen);
			setState(509);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis) | (1L << NullLiteral) | (1L << BooleanLiteral))) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & ((1L << (Break - 68)) | (1L << (Do - 68)) | (1L << (Instanceof - 68)) | (1L << (Typeof - 68)) | (1L << (Case - 68)) | (1L << (Else - 68)) | (1L << (New - 68)) | (1L << (Var - 68)) | (1L << (Catch - 68)) | (1L << (Finally - 68)) | (1L << (Return - 68)) | (1L << (Void - 68)) | (1L << (Continue - 68)) | (1L << (For - 68)) | (1L << (Switch - 68)) | (1L << (While - 68)) | (1L << (Debugger - 68)) | (1L << (Function - 68)) | (1L << (This - 68)) | (1L << (With - 68)) | (1L << (Default - 68)) | (1L << (If - 68)) | (1L << (Throw - 68)) | (1L << (Delete - 68)) | (1L << (In - 68)) | (1L << (Try - 68)) | (1L << (As - 68)) | (1L << (From - 68)) | (1L << (ReadOnly - 68)) | (1L << (Async - 68)) | (1L << (Class - 68)) | (1L << (Enum - 68)) | (1L << (Extends - 68)) | (1L << (Super - 68)) | (1L << (Const - 68)) | (1L << (Export - 68)) | (1L << (Import - 68)) | (1L << (Await - 68)) | (1L << (Implements - 68)) | (1L << (Let - 68)) | (1L << (Private - 68)) | (1L << (Public - 68)) | (1L << (Interface - 68)) | (1L << (Package - 68)) | (1L << (Protected - 68)) | (1L << (Static - 68)) | (1L << (Yield - 68)) | (1L << (ANY - 68)) | (1L << (NUMBER - 68)) | (1L << (BOOLEAN - 68)) | (1L << (STRING - 68)) | (1L << (SYMBOL - 68)) | (1L << (At - 68)) | (1L << (Identifier - 68)))) != 0)) {
				{
				setState(508);
				parameterList();
				}
			}

			setState(511);
			match(CloseParen);
			setState(513);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				setState(512);
				typeAnnotation();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParameterListContext extends ParserRuleContext {
		public RestParameterContext restParameter() {
			return getRuleContext(RestParameterContext.class,0);
		}
		public List<PredefinedTypeContext> predefinedType() {
			return getRuleContexts(PredefinedTypeContext.class);
		}
		public PredefinedTypeContext predefinedType(int i) {
			return getRuleContext(PredefinedTypeContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public OptionalParameterListContext optionalParameterList() {
			return getRuleContext(OptionalParameterListContext.class,0);
		}
		public RequiredParameterListContext requiredParameterList() {
			return getRuleContext(RequiredParameterListContext.class,0);
		}
		public ParameterListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterParameterList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitParameterList(this);
		}
	}

	public final ParameterListContext parameterList() throws RecognitionException {
		ParameterListContext _localctx = new ParameterListContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_parameterList);
		int _la;
		try {
			setState(541);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(515);
				restParameter();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(516);
				predefinedType();
				setState(521);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==Comma) {
					{
					{
					setState(517);
					match(Comma);
					setState(518);
					predefinedType();
					}
					}
					setState(523);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(524);
				optionalParameterList();
				setState(527);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Comma) {
					{
					setState(525);
					match(Comma);
					setState(526);
					restParameter();
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(529);
				requiredParameterList();
				setState(539);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Comma) {
					{
					setState(530);
					match(Comma);
					setState(537);
					_errHandler.sync(this);
					switch (_input.LA(1)) {
					case OpenBracket:
					case OpenBrace:
					case NullLiteral:
					case BooleanLiteral:
					case Break:
					case Do:
					case Instanceof:
					case Typeof:
					case Case:
					case Else:
					case New:
					case Var:
					case Catch:
					case Finally:
					case Return:
					case Void:
					case Continue:
					case For:
					case Switch:
					case While:
					case Debugger:
					case Function:
					case This:
					case With:
					case Default:
					case If:
					case Throw:
					case Delete:
					case In:
					case Try:
					case As:
					case From:
					case ReadOnly:
					case Async:
					case Class:
					case Enum:
					case Extends:
					case Super:
					case Const:
					case Export:
					case Import:
					case Await:
					case Implements:
					case Let:
					case Private:
					case Public:
					case Interface:
					case Package:
					case Protected:
					case Static:
					case Yield:
					case At:
					case Identifier:
						{
						setState(531);
						optionalParameterList();
						setState(534);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==Comma) {
							{
							setState(532);
							match(Comma);
							setState(533);
							restParameter();
							}
						}

						}
						break;
					case Ellipsis:
						{
						setState(536);
						restParameter();
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					}
				}

				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RequiredParameterListContext extends ParserRuleContext {
		public List<RequiredParameterContext> requiredParameter() {
			return getRuleContexts(RequiredParameterContext.class);
		}
		public RequiredParameterContext requiredParameter(int i) {
			return getRuleContext(RequiredParameterContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public RequiredParameterListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_requiredParameterList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterRequiredParameterList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitRequiredParameterList(this);
		}
	}

	public final RequiredParameterListContext requiredParameterList() throws RecognitionException {
		RequiredParameterListContext _localctx = new RequiredParameterListContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_requiredParameterList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(543);
			requiredParameter();
			setState(548);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(544);
					match(Comma);
					setState(545);
					requiredParameter();
					}
					} 
				}
				setState(550);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RequiredParameterContext extends ParserRuleContext {
		public IdentifierOrPatternContext identifierOrPattern() {
			return getRuleContext(IdentifierOrPatternContext.class,0);
		}
		public DecoratorListContext decoratorList() {
			return getRuleContext(DecoratorListContext.class,0);
		}
		public AccessibilityModifierContext accessibilityModifier() {
			return getRuleContext(AccessibilityModifierContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public RequiredParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_requiredParameter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterRequiredParameter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitRequiredParameter(this);
		}
	}

	public final RequiredParameterContext requiredParameter() throws RecognitionException {
		RequiredParameterContext _localctx = new RequiredParameterContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_requiredParameter);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(552);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==At) {
				{
				setState(551);
				decoratorList();
				}
			}

			setState(555);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
			case 1:
				{
				setState(554);
				accessibilityModifier();
				}
				break;
			}
			setState(557);
			identifierOrPattern();
			setState(559);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(558);
				typeAnnotation();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AccessibilityModifierContext extends ParserRuleContext {
		public TerminalNode Public() { return getToken(TypeScriptParser.Public, 0); }
		public TerminalNode Private() { return getToken(TypeScriptParser.Private, 0); }
		public TerminalNode Protected() { return getToken(TypeScriptParser.Protected, 0); }
		public AccessibilityModifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_accessibilityModifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterAccessibilityModifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitAccessibilityModifier(this);
		}
	}

	public final AccessibilityModifierContext accessibilityModifier() throws RecognitionException {
		AccessibilityModifierContext _localctx = new AccessibilityModifierContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_accessibilityModifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(561);
			_la = _input.LA(1);
			if ( !(((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IdentifierOrPatternContext extends ParserRuleContext {
		public IdentifierNameContext identifierName() {
			return getRuleContext(IdentifierNameContext.class,0);
		}
		public BindingPatternContext bindingPattern() {
			return getRuleContext(BindingPatternContext.class,0);
		}
		public IdentifierOrPatternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierOrPattern; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIdentifierOrPattern(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIdentifierOrPattern(this);
		}
	}

	public final IdentifierOrPatternContext identifierOrPattern() throws RecognitionException {
		IdentifierOrPatternContext _localctx = new IdentifierOrPatternContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_identifierOrPattern);
		try {
			setState(565);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NullLiteral:
			case BooleanLiteral:
			case Break:
			case Do:
			case Instanceof:
			case Typeof:
			case Case:
			case Else:
			case New:
			case Var:
			case Catch:
			case Finally:
			case Return:
			case Void:
			case Continue:
			case For:
			case Switch:
			case While:
			case Debugger:
			case Function:
			case This:
			case With:
			case Default:
			case If:
			case Throw:
			case Delete:
			case In:
			case Try:
			case As:
			case From:
			case ReadOnly:
			case Async:
			case Class:
			case Enum:
			case Extends:
			case Super:
			case Const:
			case Export:
			case Import:
			case Await:
			case Implements:
			case Let:
			case Private:
			case Public:
			case Interface:
			case Package:
			case Protected:
			case Static:
			case Yield:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(563);
				identifierName();
				}
				break;
			case OpenBracket:
			case OpenBrace:
				enterOuterAlt(_localctx, 2);
				{
				setState(564);
				bindingPattern();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OptionalParameterListContext extends ParserRuleContext {
		public List<OptionalParameterContext> optionalParameter() {
			return getRuleContexts(OptionalParameterContext.class);
		}
		public OptionalParameterContext optionalParameter(int i) {
			return getRuleContext(OptionalParameterContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public OptionalParameterListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_optionalParameterList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterOptionalParameterList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitOptionalParameterList(this);
		}
	}

	public final OptionalParameterListContext optionalParameterList() throws RecognitionException {
		OptionalParameterListContext _localctx = new OptionalParameterListContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_optionalParameterList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(567);
			optionalParameter();
			setState(572);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(568);
					match(Comma);
					setState(569);
					optionalParameter();
					}
					} 
				}
				setState(574);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OptionalParameterContext extends ParserRuleContext {
		public IdentifierOrPatternContext identifierOrPattern() {
			return getRuleContext(IdentifierOrPatternContext.class,0);
		}
		public DecoratorListContext decoratorList() {
			return getRuleContext(DecoratorListContext.class,0);
		}
		public TerminalNode QuestionMark() { return getToken(TypeScriptParser.QuestionMark, 0); }
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public AccessibilityModifierContext accessibilityModifier() {
			return getRuleContext(AccessibilityModifierContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public OptionalParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_optionalParameter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterOptionalParameter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitOptionalParameter(this);
		}
	}

	public final OptionalParameterContext optionalParameter() throws RecognitionException {
		OptionalParameterContext _localctx = new OptionalParameterContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_optionalParameter);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(576);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==At) {
				{
				setState(575);
				decoratorList();
				}
			}

			{
			setState(579);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,47,_ctx) ) {
			case 1:
				{
				setState(578);
				accessibilityModifier();
				}
				break;
			}
			setState(581);
			identifierOrPattern();
			setState(590);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case QuestionMark:
				{
				setState(582);
				match(QuestionMark);
				setState(584);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Colon) {
					{
					setState(583);
					typeAnnotation();
					}
				}

				}
				break;
			case Assign:
			case Colon:
				{
				setState(587);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Colon) {
					{
					setState(586);
					typeAnnotation();
					}
				}

				setState(589);
				initializer();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RestParameterContext extends ParserRuleContext {
		public TerminalNode Ellipsis() { return getToken(TypeScriptParser.Ellipsis, 0); }
		public RequiredParameterContext requiredParameter() {
			return getRuleContext(RequiredParameterContext.class,0);
		}
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public RestParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_restParameter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterRestParameter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitRestParameter(this);
		}
	}

	public final RestParameterContext restParameter() throws RecognitionException {
		RestParameterContext _localctx = new RestParameterContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_restParameter);
		try {
			setState(596);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,51,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(592);
				match(Ellipsis);
				setState(593);
				requiredParameter();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(594);
				match(Ellipsis);
				setState(595);
				singleExpression(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstructSignatureContext extends ParserRuleContext {
		public TerminalNode New() { return getToken(TypeScriptParser.New, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public ParameterListContext parameterList() {
			return getRuleContext(ParameterListContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public ConstructSignatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constructSignature; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterConstructSignature(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitConstructSignature(this);
		}
	}

	public final ConstructSignatureContext constructSignature() throws RecognitionException {
		ConstructSignatureContext _localctx = new ConstructSignatureContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_constructSignature);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(598);
			match(New);
			setState(600);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LessThan) {
				{
				setState(599);
				typeParameters();
				}
			}

			setState(602);
			match(OpenParen);
			setState(604);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis) | (1L << NullLiteral) | (1L << BooleanLiteral))) != 0) || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & ((1L << (Break - 68)) | (1L << (Do - 68)) | (1L << (Instanceof - 68)) | (1L << (Typeof - 68)) | (1L << (Case - 68)) | (1L << (Else - 68)) | (1L << (New - 68)) | (1L << (Var - 68)) | (1L << (Catch - 68)) | (1L << (Finally - 68)) | (1L << (Return - 68)) | (1L << (Void - 68)) | (1L << (Continue - 68)) | (1L << (For - 68)) | (1L << (Switch - 68)) | (1L << (While - 68)) | (1L << (Debugger - 68)) | (1L << (Function - 68)) | (1L << (This - 68)) | (1L << (With - 68)) | (1L << (Default - 68)) | (1L << (If - 68)) | (1L << (Throw - 68)) | (1L << (Delete - 68)) | (1L << (In - 68)) | (1L << (Try - 68)) | (1L << (As - 68)) | (1L << (From - 68)) | (1L << (ReadOnly - 68)) | (1L << (Async - 68)) | (1L << (Class - 68)) | (1L << (Enum - 68)) | (1L << (Extends - 68)) | (1L << (Super - 68)) | (1L << (Const - 68)) | (1L << (Export - 68)) | (1L << (Import - 68)) | (1L << (Await - 68)) | (1L << (Implements - 68)) | (1L << (Let - 68)) | (1L << (Private - 68)) | (1L << (Public - 68)) | (1L << (Interface - 68)) | (1L << (Package - 68)) | (1L << (Protected - 68)) | (1L << (Static - 68)) | (1L << (Yield - 68)) | (1L << (ANY - 68)) | (1L << (NUMBER - 68)) | (1L << (BOOLEAN - 68)) | (1L << (STRING - 68)) | (1L << (SYMBOL - 68)) | (1L << (At - 68)) | (1L << (Identifier - 68)))) != 0)) {
				{
				setState(603);
				parameterList();
				}
			}

			setState(606);
			match(CloseParen);
			setState(608);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(607);
				typeAnnotation();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IndexSignatureContext extends ParserRuleContext {
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public TerminalNode NUMBER() { return getToken(TypeScriptParser.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(TypeScriptParser.STRING, 0); }
		public IndexSignatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_indexSignature; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIndexSignature(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIndexSignature(this);
		}
	}

	public final IndexSignatureContext indexSignature() throws RecognitionException {
		IndexSignatureContext _localctx = new IndexSignatureContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_indexSignature);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(610);
			match(OpenBracket);
			setState(611);
			match(Identifier);
			setState(612);
			match(Colon);
			setState(613);
			_la = _input.LA(1);
			if ( !(_la==NUMBER || _la==STRING) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(614);
			match(CloseBracket);
			setState(615);
			typeAnnotation();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MethodSignatureContext extends ParserRuleContext {
		public PropertyNameContext propertyName() {
			return getRuleContext(PropertyNameContext.class,0);
		}
		public CallSignatureContext callSignature() {
			return getRuleContext(CallSignatureContext.class,0);
		}
		public TerminalNode QuestionMark() { return getToken(TypeScriptParser.QuestionMark, 0); }
		public MethodSignatureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_methodSignature; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterMethodSignature(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitMethodSignature(this);
		}
	}

	public final MethodSignatureContext methodSignature() throws RecognitionException {
		MethodSignatureContext _localctx = new MethodSignatureContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_methodSignature);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(617);
			propertyName();
			setState(619);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==QuestionMark) {
				{
				setState(618);
				match(QuestionMark);
				}
			}

			setState(621);
			callSignature();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeAliasDeclarationContext extends ParserRuleContext {
		public TerminalNode Type() { return getToken(TypeScriptParser.Type, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode Assign() { return getToken(TypeScriptParser.Assign, 0); }
		public Type_Context type_() {
			return getRuleContext(Type_Context.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public TypeAliasDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeAliasDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeAliasDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeAliasDeclaration(this);
		}
	}

	public final TypeAliasDeclarationContext typeAliasDeclaration() throws RecognitionException {
		TypeAliasDeclarationContext _localctx = new TypeAliasDeclarationContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_typeAliasDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(623);
			match(Type);
			setState(624);
			match(Identifier);
			setState(626);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LessThan) {
				{
				setState(625);
				typeParameters();
				}
			}

			setState(628);
			match(Assign);
			setState(629);
			type_();
			setState(630);
			match(SemiColon);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstructorDeclarationContext extends ParserRuleContext {
		public TerminalNode Constructor() { return getToken(TypeScriptParser.Constructor, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public AccessibilityModifierContext accessibilityModifier() {
			return getRuleContext(AccessibilityModifierContext.class,0);
		}
		public FormalParameterListContext formalParameterList() {
			return getRuleContext(FormalParameterListContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public ConstructorDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constructorDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterConstructorDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitConstructorDeclaration(this);
		}
	}

	public final ConstructorDeclarationContext constructorDeclaration() throws RecognitionException {
		ConstructorDeclarationContext _localctx = new ConstructorDeclarationContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_constructorDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(633);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)))) != 0)) {
				{
				setState(632);
				accessibilityModifier();
				}
			}

			setState(635);
			match(Constructor);
			setState(636);
			match(OpenParen);
			setState(638);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis))) != 0) || ((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)) | (1L << (Identifier - 108)))) != 0)) {
				{
				setState(637);
				formalParameterList();
				}
			}

			setState(640);
			match(CloseParen);
			setState(646);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,59,_ctx) ) {
			case 1:
				{
				{
				setState(641);
				match(OpenBrace);
				setState(642);
				functionBody();
				setState(643);
				match(CloseBrace);
				}
				}
				break;
			case 2:
				{
				setState(645);
				match(SemiColon);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InterfaceDeclarationContext extends ParserRuleContext {
		public TerminalNode Interface() { return getToken(TypeScriptParser.Interface, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public ObjectTypeContext objectType() {
			return getRuleContext(ObjectTypeContext.class,0);
		}
		public TerminalNode Export() { return getToken(TypeScriptParser.Export, 0); }
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public InterfaceExtendsClauseContext interfaceExtendsClause() {
			return getRuleContext(InterfaceExtendsClauseContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public InterfaceDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterInterfaceDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitInterfaceDeclaration(this);
		}
	}

	public final InterfaceDeclarationContext interfaceDeclaration() throws RecognitionException {
		InterfaceDeclarationContext _localctx = new InterfaceDeclarationContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_interfaceDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(649);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Export) {
				{
				setState(648);
				match(Export);
				}
			}

			setState(651);
			match(Interface);
			setState(652);
			match(Identifier);
			setState(654);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LessThan) {
				{
				setState(653);
				typeParameters();
				}
			}

			setState(657);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Extends) {
				{
				setState(656);
				interfaceExtendsClause();
				}
			}

			setState(659);
			objectType();
			setState(661);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,63,_ctx) ) {
			case 1:
				{
				setState(660);
				match(SemiColon);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InterfaceExtendsClauseContext extends ParserRuleContext {
		public TerminalNode Extends() { return getToken(TypeScriptParser.Extends, 0); }
		public ClassOrInterfaceTypeListContext classOrInterfaceTypeList() {
			return getRuleContext(ClassOrInterfaceTypeListContext.class,0);
		}
		public InterfaceExtendsClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceExtendsClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterInterfaceExtendsClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitInterfaceExtendsClause(this);
		}
	}

	public final InterfaceExtendsClauseContext interfaceExtendsClause() throws RecognitionException {
		InterfaceExtendsClauseContext _localctx = new InterfaceExtendsClauseContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_interfaceExtendsClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(663);
			match(Extends);
			setState(664);
			classOrInterfaceTypeList();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ClassOrInterfaceTypeListContext extends ParserRuleContext {
		public List<TypeReferenceContext> typeReference() {
			return getRuleContexts(TypeReferenceContext.class);
		}
		public TypeReferenceContext typeReference(int i) {
			return getRuleContext(TypeReferenceContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public ClassOrInterfaceTypeListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classOrInterfaceTypeList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterClassOrInterfaceTypeList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitClassOrInterfaceTypeList(this);
		}
	}

	public final ClassOrInterfaceTypeListContext classOrInterfaceTypeList() throws RecognitionException {
		ClassOrInterfaceTypeListContext _localctx = new ClassOrInterfaceTypeListContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_classOrInterfaceTypeList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(666);
			typeReference();
			setState(671);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Comma) {
				{
				{
				setState(667);
				match(Comma);
				setState(668);
				typeReference();
				}
				}
				setState(673);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EnumDeclarationContext extends ParserRuleContext {
		public TerminalNode Enum() { return getToken(TypeScriptParser.Enum, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public TerminalNode Const() { return getToken(TypeScriptParser.Const, 0); }
		public EnumBodyContext enumBody() {
			return getRuleContext(EnumBodyContext.class,0);
		}
		public EnumDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterEnumDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitEnumDeclaration(this);
		}
	}

	public final EnumDeclarationContext enumDeclaration() throws RecognitionException {
		EnumDeclarationContext _localctx = new EnumDeclarationContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_enumDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(675);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Const) {
				{
				setState(674);
				match(Const);
				}
			}

			setState(677);
			match(Enum);
			setState(678);
			match(Identifier);
			setState(679);
			match(OpenBrace);
			setState(681);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & ((1L << (OpenBracket - 5)) | (1L << (NullLiteral - 5)) | (1L << (BooleanLiteral - 5)) | (1L << (DecimalLiteral - 5)) | (1L << (HexIntegerLiteral - 5)) | (1L << (OctalIntegerLiteral - 5)) | (1L << (OctalIntegerLiteral2 - 5)) | (1L << (BinaryIntegerLiteral - 5)) | (1L << (Break - 5)))) != 0) || ((((_la - 69)) & ~0x3f) == 0 && ((1L << (_la - 69)) & ((1L << (Do - 69)) | (1L << (Instanceof - 69)) | (1L << (Typeof - 69)) | (1L << (Case - 69)) | (1L << (Else - 69)) | (1L << (New - 69)) | (1L << (Var - 69)) | (1L << (Catch - 69)) | (1L << (Finally - 69)) | (1L << (Return - 69)) | (1L << (Void - 69)) | (1L << (Continue - 69)) | (1L << (For - 69)) | (1L << (Switch - 69)) | (1L << (While - 69)) | (1L << (Debugger - 69)) | (1L << (Function - 69)) | (1L << (This - 69)) | (1L << (With - 69)) | (1L << (Default - 69)) | (1L << (If - 69)) | (1L << (Throw - 69)) | (1L << (Delete - 69)) | (1L << (In - 69)) | (1L << (Try - 69)) | (1L << (As - 69)) | (1L << (From - 69)) | (1L << (ReadOnly - 69)) | (1L << (Async - 69)) | (1L << (Class - 69)) | (1L << (Enum - 69)) | (1L << (Extends - 69)) | (1L << (Super - 69)) | (1L << (Const - 69)) | (1L << (Export - 69)) | (1L << (Import - 69)) | (1L << (Await - 69)) | (1L << (Implements - 69)) | (1L << (Let - 69)) | (1L << (Private - 69)) | (1L << (Public - 69)) | (1L << (Interface - 69)) | (1L << (Package - 69)) | (1L << (Protected - 69)) | (1L << (Static - 69)) | (1L << (Yield - 69)) | (1L << (Identifier - 69)) | (1L << (StringLiteral - 69)))) != 0)) {
				{
				setState(680);
				enumBody();
				}
			}

			setState(683);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EnumBodyContext extends ParserRuleContext {
		public EnumMemberListContext enumMemberList() {
			return getRuleContext(EnumMemberListContext.class,0);
		}
		public TerminalNode Comma() { return getToken(TypeScriptParser.Comma, 0); }
		public EnumBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumBody; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterEnumBody(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitEnumBody(this);
		}
	}

	public final EnumBodyContext enumBody() throws RecognitionException {
		EnumBodyContext _localctx = new EnumBodyContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_enumBody);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(685);
			enumMemberList();
			setState(687);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Comma) {
				{
				setState(686);
				match(Comma);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EnumMemberListContext extends ParserRuleContext {
		public List<EnumMemberContext> enumMember() {
			return getRuleContexts(EnumMemberContext.class);
		}
		public EnumMemberContext enumMember(int i) {
			return getRuleContext(EnumMemberContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public EnumMemberListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumMemberList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterEnumMemberList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitEnumMemberList(this);
		}
	}

	public final EnumMemberListContext enumMemberList() throws RecognitionException {
		EnumMemberListContext _localctx = new EnumMemberListContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_enumMemberList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(689);
			enumMember();
			setState(694);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,68,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(690);
					match(Comma);
					setState(691);
					enumMember();
					}
					} 
				}
				setState(696);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,68,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EnumMemberContext extends ParserRuleContext {
		public PropertyNameContext propertyName() {
			return getRuleContext(PropertyNameContext.class,0);
		}
		public TerminalNode Assign() { return getToken(TypeScriptParser.Assign, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public EnumMemberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumMember; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterEnumMember(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitEnumMember(this);
		}
	}

	public final EnumMemberContext enumMember() throws RecognitionException {
		EnumMemberContext _localctx = new EnumMemberContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_enumMember);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(697);
			propertyName();
			setState(700);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Assign) {
				{
				setState(698);
				match(Assign);
				setState(699);
				singleExpression(0);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NamespaceDeclarationContext extends ParserRuleContext {
		public TerminalNode Namespace() { return getToken(TypeScriptParser.Namespace, 0); }
		public NamespaceNameContext namespaceName() {
			return getRuleContext(NamespaceNameContext.class,0);
		}
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public NamespaceDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_namespaceDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterNamespaceDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitNamespaceDeclaration(this);
		}
	}

	public final NamespaceDeclarationContext namespaceDeclaration() throws RecognitionException {
		NamespaceDeclarationContext _localctx = new NamespaceDeclarationContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_namespaceDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(702);
			match(Namespace);
			setState(703);
			namespaceName();
			setState(704);
			match(OpenBrace);
			setState(706);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,70,_ctx) ) {
			case 1:
				{
				setState(705);
				statementList();
				}
				break;
			}
			setState(708);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NamespaceNameContext extends ParserRuleContext {
		public List<TerminalNode> Identifier() { return getTokens(TypeScriptParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(TypeScriptParser.Identifier, i);
		}
		public List<TerminalNode> Dot() { return getTokens(TypeScriptParser.Dot); }
		public TerminalNode Dot(int i) {
			return getToken(TypeScriptParser.Dot, i);
		}
		public NamespaceNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_namespaceName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterNamespaceName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitNamespaceName(this);
		}
	}

	public final NamespaceNameContext namespaceName() throws RecognitionException {
		NamespaceNameContext _localctx = new NamespaceNameContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_namespaceName);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(710);
			match(Identifier);
			setState(719);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,72,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(712); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(711);
						match(Dot);
						}
						}
						setState(714); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==Dot );
					setState(716);
					match(Identifier);
					}
					} 
				}
				setState(721);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,72,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportAliasDeclarationContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode Assign() { return getToken(TypeScriptParser.Assign, 0); }
		public NamespaceNameContext namespaceName() {
			return getRuleContext(NamespaceNameContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public TerminalNode Require() { return getToken(TypeScriptParser.Require, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode StringLiteral() { return getToken(TypeScriptParser.StringLiteral, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public ImportAliasDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importAliasDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterImportAliasDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitImportAliasDeclaration(this);
		}
	}

	public final ImportAliasDeclarationContext importAliasDeclaration() throws RecognitionException {
		ImportAliasDeclarationContext _localctx = new ImportAliasDeclarationContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_importAliasDeclaration);
		try {
			setState(734);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,73,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(722);
				match(Identifier);
				setState(723);
				match(Assign);
				setState(724);
				namespaceName();
				setState(725);
				match(SemiColon);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(727);
				match(Identifier);
				setState(728);
				match(Assign);
				setState(729);
				match(Require);
				setState(730);
				match(OpenParen);
				setState(731);
				match(StringLiteral);
				setState(732);
				match(CloseParen);
				setState(733);
				match(SemiColon);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportAllContext extends ParserRuleContext {
		public TerminalNode StringLiteral() { return getToken(TypeScriptParser.StringLiteral, 0); }
		public ImportAllContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importAll; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterImportAll(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitImportAll(this);
		}
	}

	public final ImportAllContext importAll() throws RecognitionException {
		ImportAllContext _localctx = new ImportAllContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_importAll);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(736);
			match(StringLiteral);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DecoratorListContext extends ParserRuleContext {
		public List<DecoratorContext> decorator() {
			return getRuleContexts(DecoratorContext.class);
		}
		public DecoratorContext decorator(int i) {
			return getRuleContext(DecoratorContext.class,i);
		}
		public DecoratorListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decoratorList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterDecoratorList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitDecoratorList(this);
		}
	}

	public final DecoratorListContext decoratorList() throws RecognitionException {
		DecoratorListContext _localctx = new DecoratorListContext(_ctx, getState());
		enterRule(_localctx, 110, RULE_decoratorList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(739); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(738);
				decorator();
				}
				}
				setState(741); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==At );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DecoratorContext extends ParserRuleContext {
		public TerminalNode At() { return getToken(TypeScriptParser.At, 0); }
		public DecoratorMemberExpressionContext decoratorMemberExpression() {
			return getRuleContext(DecoratorMemberExpressionContext.class,0);
		}
		public DecoratorCallExpressionContext decoratorCallExpression() {
			return getRuleContext(DecoratorCallExpressionContext.class,0);
		}
		public DecoratorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decorator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterDecorator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitDecorator(this);
		}
	}

	public final DecoratorContext decorator() throws RecognitionException {
		DecoratorContext _localctx = new DecoratorContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_decorator);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(743);
			match(At);
			setState(746);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,75,_ctx) ) {
			case 1:
				{
				setState(744);
				decoratorMemberExpression(0);
				}
				break;
			case 2:
				{
				setState(745);
				decoratorCallExpression();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DecoratorMemberExpressionContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public DecoratorMemberExpressionContext decoratorMemberExpression() {
			return getRuleContext(DecoratorMemberExpressionContext.class,0);
		}
		public TerminalNode Dot() { return getToken(TypeScriptParser.Dot, 0); }
		public IdentifierNameContext identifierName() {
			return getRuleContext(IdentifierNameContext.class,0);
		}
		public DecoratorMemberExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decoratorMemberExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterDecoratorMemberExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitDecoratorMemberExpression(this);
		}
	}

	public final DecoratorMemberExpressionContext decoratorMemberExpression() throws RecognitionException {
		return decoratorMemberExpression(0);
	}

	private DecoratorMemberExpressionContext decoratorMemberExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		DecoratorMemberExpressionContext _localctx = new DecoratorMemberExpressionContext(_ctx, _parentState);
		DecoratorMemberExpressionContext _prevctx = _localctx;
		int _startState = 114;
		enterRecursionRule(_localctx, 114, RULE_decoratorMemberExpression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(754);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				{
				setState(749);
				match(Identifier);
				}
				break;
			case OpenParen:
				{
				setState(750);
				match(OpenParen);
				setState(751);
				singleExpression(0);
				setState(752);
				match(CloseParen);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(761);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,77,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new DecoratorMemberExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_decoratorMemberExpression);
					setState(756);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(757);
					match(Dot);
					setState(758);
					identifierName();
					}
					} 
				}
				setState(763);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,77,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class DecoratorCallExpressionContext extends ParserRuleContext {
		public DecoratorMemberExpressionContext decoratorMemberExpression() {
			return getRuleContext(DecoratorMemberExpressionContext.class,0);
		}
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public DecoratorCallExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decoratorCallExpression; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterDecoratorCallExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitDecoratorCallExpression(this);
		}
	}

	public final DecoratorCallExpressionContext decoratorCallExpression() throws RecognitionException {
		DecoratorCallExpressionContext _localctx = new DecoratorCallExpressionContext(_ctx, getState());
		enterRule(_localctx, 116, RULE_decoratorCallExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(764);
			decoratorMemberExpression(0);
			setState(765);
			arguments();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ProgramContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(TypeScriptParser.EOF, 0); }
		public SourceElementsContext sourceElements() {
			return getRuleContext(SourceElementsContext.class,0);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterProgram(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitProgram(this);
		}
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 118, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(768);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,78,_ctx) ) {
			case 1:
				{
				setState(767);
				sourceElements();
				}
				break;
			}
			setState(770);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SourceElementContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode Export() { return getToken(TypeScriptParser.Export, 0); }
		public SourceElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sourceElement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterSourceElement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitSourceElement(this);
		}
	}

	public final SourceElementContext sourceElement() throws RecognitionException {
		SourceElementContext _localctx = new SourceElementContext(_ctx, getState());
		enterRule(_localctx, 120, RULE_sourceElement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(773);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,79,_ctx) ) {
			case 1:
				{
				setState(772);
				match(Export);
				}
				break;
			}
			setState(775);
			statement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public VariableStatementContext variableStatement() {
			return getRuleContext(VariableStatementContext.class,0);
		}
		public ImportStatementContext importStatement() {
			return getRuleContext(ImportStatementContext.class,0);
		}
		public ExportStatementContext exportStatement() {
			return getRuleContext(ExportStatementContext.class,0);
		}
		public EmptyStatement_Context emptyStatement_() {
			return getRuleContext(EmptyStatement_Context.class,0);
		}
		public AbstractDeclarationContext abstractDeclaration() {
			return getRuleContext(AbstractDeclarationContext.class,0);
		}
		public ClassDeclarationContext classDeclaration() {
			return getRuleContext(ClassDeclarationContext.class,0);
		}
		public InterfaceDeclarationContext interfaceDeclaration() {
			return getRuleContext(InterfaceDeclarationContext.class,0);
		}
		public NamespaceDeclarationContext namespaceDeclaration() {
			return getRuleContext(NamespaceDeclarationContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public IterationStatementContext iterationStatement() {
			return getRuleContext(IterationStatementContext.class,0);
		}
		public ContinueStatementContext continueStatement() {
			return getRuleContext(ContinueStatementContext.class,0);
		}
		public BreakStatementContext breakStatement() {
			return getRuleContext(BreakStatementContext.class,0);
		}
		public ReturnStatementContext returnStatement() {
			return getRuleContext(ReturnStatementContext.class,0);
		}
		public YieldStatementContext yieldStatement() {
			return getRuleContext(YieldStatementContext.class,0);
		}
		public WithStatementContext withStatement() {
			return getRuleContext(WithStatementContext.class,0);
		}
		public LabelledStatementContext labelledStatement() {
			return getRuleContext(LabelledStatementContext.class,0);
		}
		public SwitchStatementContext switchStatement() {
			return getRuleContext(SwitchStatementContext.class,0);
		}
		public ThrowStatementContext throwStatement() {
			return getRuleContext(ThrowStatementContext.class,0);
		}
		public TryStatementContext tryStatement() {
			return getRuleContext(TryStatementContext.class,0);
		}
		public DebuggerStatementContext debuggerStatement() {
			return getRuleContext(DebuggerStatementContext.class,0);
		}
		public FunctionDeclarationContext functionDeclaration() {
			return getRuleContext(FunctionDeclarationContext.class,0);
		}
		public ArrowFunctionDeclarationContext arrowFunctionDeclaration() {
			return getRuleContext(ArrowFunctionDeclarationContext.class,0);
		}
		public GeneratorFunctionDeclarationContext generatorFunctionDeclaration() {
			return getRuleContext(GeneratorFunctionDeclarationContext.class,0);
		}
		public TypeAliasDeclarationContext typeAliasDeclaration() {
			return getRuleContext(TypeAliasDeclarationContext.class,0);
		}
		public EnumDeclarationContext enumDeclaration() {
			return getRuleContext(EnumDeclarationContext.class,0);
		}
		public ExpressionStatementContext expressionStatement() {
			return getRuleContext(ExpressionStatementContext.class,0);
		}
		public TerminalNode Export() { return getToken(TypeScriptParser.Export, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitStatement(this);
		}
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 122, RULE_statement);
		try {
			setState(806);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,80,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(777);
				block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(778);
				variableStatement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(779);
				importStatement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(780);
				exportStatement();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(781);
				emptyStatement_();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(782);
				abstractDeclaration();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(783);
				classDeclaration();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(784);
				interfaceDeclaration();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(785);
				namespaceDeclaration();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(786);
				ifStatement();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(787);
				iterationStatement();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(788);
				continueStatement();
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(789);
				breakStatement();
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(790);
				returnStatement();
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(791);
				yieldStatement();
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(792);
				withStatement();
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(793);
				labelledStatement();
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(794);
				switchStatement();
				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(795);
				throwStatement();
				}
				break;
			case 20:
				enterOuterAlt(_localctx, 20);
				{
				setState(796);
				tryStatement();
				}
				break;
			case 21:
				enterOuterAlt(_localctx, 21);
				{
				setState(797);
				debuggerStatement();
				}
				break;
			case 22:
				enterOuterAlt(_localctx, 22);
				{
				setState(798);
				functionDeclaration();
				}
				break;
			case 23:
				enterOuterAlt(_localctx, 23);
				{
				setState(799);
				arrowFunctionDeclaration();
				}
				break;
			case 24:
				enterOuterAlt(_localctx, 24);
				{
				setState(800);
				generatorFunctionDeclaration();
				}
				break;
			case 25:
				enterOuterAlt(_localctx, 25);
				{
				setState(801);
				typeAliasDeclaration();
				}
				break;
			case 26:
				enterOuterAlt(_localctx, 26);
				{
				setState(802);
				enumDeclaration();
				}
				break;
			case 27:
				enterOuterAlt(_localctx, 27);
				{
				setState(803);
				expressionStatement();
				}
				break;
			case 28:
				enterOuterAlt(_localctx, 28);
				{
				setState(804);
				match(Export);
				setState(805);
				statement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitBlock(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 124, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(808);
			match(OpenBrace);
			setState(810);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,81,_ctx) ) {
			case 1:
				{
				setState(809);
				statementList();
				}
				break;
			}
			setState(812);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementListContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statementList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterStatementList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitStatementList(this);
		}
	}

	public final StatementListContext statementList() throws RecognitionException {
		StatementListContext _localctx = new StatementListContext(_ctx, getState());
		enterRule(_localctx, 126, RULE_statementList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(815); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(814);
					statement();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(817); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,82,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AbstractDeclarationContext extends ParserRuleContext {
		public TerminalNode Abstract() { return getToken(TypeScriptParser.Abstract, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public CallSignatureContext callSignature() {
			return getRuleContext(CallSignatureContext.class,0);
		}
		public VariableStatementContext variableStatement() {
			return getRuleContext(VariableStatementContext.class,0);
		}
		public AbstractDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_abstractDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterAbstractDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitAbstractDeclaration(this);
		}
	}

	public final AbstractDeclarationContext abstractDeclaration() throws RecognitionException {
		AbstractDeclarationContext _localctx = new AbstractDeclarationContext(_ctx, getState());
		enterRule(_localctx, 128, RULE_abstractDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(819);
			match(Abstract);
			setState(823);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,83,_ctx) ) {
			case 1:
				{
				setState(820);
				match(Identifier);
				setState(821);
				callSignature();
				}
				break;
			case 2:
				{
				setState(822);
				variableStatement();
				}
				break;
			}
			setState(825);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportStatementContext extends ParserRuleContext {
		public TerminalNode Import() { return getToken(TypeScriptParser.Import, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public ImportFromBlockContext importFromBlock() {
			return getRuleContext(ImportFromBlockContext.class,0);
		}
		public ImportAliasDeclarationContext importAliasDeclaration() {
			return getRuleContext(ImportAliasDeclarationContext.class,0);
		}
		public ImportAllContext importAll() {
			return getRuleContext(ImportAllContext.class,0);
		}
		public ImportStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterImportStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitImportStatement(this);
		}
	}

	public final ImportStatementContext importStatement() throws RecognitionException {
		ImportStatementContext _localctx = new ImportStatementContext(_ctx, getState());
		enterRule(_localctx, 130, RULE_importStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(827);
			match(Import);
			setState(831);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,84,_ctx) ) {
			case 1:
				{
				setState(828);
				importFromBlock();
				}
				break;
			case 2:
				{
				setState(829);
				importAliasDeclaration();
				}
				break;
			case 3:
				{
				setState(830);
				importAll();
				}
				break;
			}
			setState(833);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportFromBlockContext extends ParserRuleContext {
		public TerminalNode From() { return getToken(TypeScriptParser.From, 0); }
		public TerminalNode StringLiteral() { return getToken(TypeScriptParser.StringLiteral, 0); }
		public TerminalNode Dollar() { return getToken(TypeScriptParser.Dollar, 0); }
		public TerminalNode Lodash() { return getToken(TypeScriptParser.Lodash, 0); }
		public TerminalNode Multiply() { return getToken(TypeScriptParser.Multiply, 0); }
		public MultipleImportStatementContext multipleImportStatement() {
			return getRuleContext(MultipleImportStatementContext.class,0);
		}
		public List<IdentifierNameContext> identifierName() {
			return getRuleContexts(IdentifierNameContext.class);
		}
		public IdentifierNameContext identifierName(int i) {
			return getRuleContext(IdentifierNameContext.class,i);
		}
		public TerminalNode As() { return getToken(TypeScriptParser.As, 0); }
		public ImportFromBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importFromBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterImportFromBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitImportFromBlock(this);
		}
	}

	public final ImportFromBlockContext importFromBlock() throws RecognitionException {
		ImportFromBlockContext _localctx = new ImportFromBlockContext(_ctx, getState());
		enterRule(_localctx, 132, RULE_importFromBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(840);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,85,_ctx) ) {
			case 1:
				{
				setState(835);
				match(Dollar);
				}
				break;
			case 2:
				{
				setState(836);
				match(Lodash);
				}
				break;
			case 3:
				{
				setState(837);
				match(Multiply);
				}
				break;
			case 4:
				{
				setState(838);
				multipleImportStatement();
				}
				break;
			case 5:
				{
				setState(839);
				identifierName();
				}
				break;
			}
			setState(844);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==As) {
				{
				setState(842);
				match(As);
				setState(843);
				identifierName();
				}
			}

			setState(846);
			match(From);
			setState(847);
			match(StringLiteral);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MultipleImportStatementContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public List<IdentifierNameContext> identifierName() {
			return getRuleContexts(IdentifierNameContext.class);
		}
		public IdentifierNameContext identifierName(int i) {
			return getRuleContext(IdentifierNameContext.class,i);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public MultipleImportStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_multipleImportStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterMultipleImportStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitMultipleImportStatement(this);
		}
	}

	public final MultipleImportStatementContext multipleImportStatement() throws RecognitionException {
		MultipleImportStatementContext _localctx = new MultipleImportStatementContext(_ctx, getState());
		enterRule(_localctx, 134, RULE_multipleImportStatement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(852);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==NullLiteral || _la==BooleanLiteral || ((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & ((1L << (Break - 68)) | (1L << (Do - 68)) | (1L << (Instanceof - 68)) | (1L << (Typeof - 68)) | (1L << (Case - 68)) | (1L << (Else - 68)) | (1L << (New - 68)) | (1L << (Var - 68)) | (1L << (Catch - 68)) | (1L << (Finally - 68)) | (1L << (Return - 68)) | (1L << (Void - 68)) | (1L << (Continue - 68)) | (1L << (For - 68)) | (1L << (Switch - 68)) | (1L << (While - 68)) | (1L << (Debugger - 68)) | (1L << (Function - 68)) | (1L << (This - 68)) | (1L << (With - 68)) | (1L << (Default - 68)) | (1L << (If - 68)) | (1L << (Throw - 68)) | (1L << (Delete - 68)) | (1L << (In - 68)) | (1L << (Try - 68)) | (1L << (As - 68)) | (1L << (From - 68)) | (1L << (ReadOnly - 68)) | (1L << (Async - 68)) | (1L << (Class - 68)) | (1L << (Enum - 68)) | (1L << (Extends - 68)) | (1L << (Super - 68)) | (1L << (Const - 68)) | (1L << (Export - 68)) | (1L << (Import - 68)) | (1L << (Await - 68)) | (1L << (Implements - 68)) | (1L << (Let - 68)) | (1L << (Private - 68)) | (1L << (Public - 68)) | (1L << (Interface - 68)) | (1L << (Package - 68)) | (1L << (Protected - 68)) | (1L << (Static - 68)) | (1L << (Yield - 68)) | (1L << (Identifier - 68)))) != 0)) {
				{
				setState(849);
				identifierName();
				setState(850);
				match(Comma);
				}
			}

			setState(854);
			match(OpenBrace);
			setState(855);
			identifierName();
			setState(860);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==Comma) {
				{
				{
				setState(856);
				match(Comma);
				setState(857);
				identifierName();
				}
				}
				setState(862);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(863);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExportStatementContext extends ParserRuleContext {
		public TerminalNode Export() { return getToken(TypeScriptParser.Export, 0); }
		public ImportFromBlockContext importFromBlock() {
			return getRuleContext(ImportFromBlockContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode Default() { return getToken(TypeScriptParser.Default, 0); }
		public ExportStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exportStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterExportStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitExportStatement(this);
		}
	}

	public final ExportStatementContext exportStatement() throws RecognitionException {
		ExportStatementContext _localctx = new ExportStatementContext(_ctx, getState());
		enterRule(_localctx, 136, RULE_exportStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(865);
			match(Export);
			setState(867);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,89,_ctx) ) {
			case 1:
				{
				setState(866);
				match(Default);
				}
				break;
			}
			setState(871);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,90,_ctx) ) {
			case 1:
				{
				setState(869);
				importFromBlock();
				}
				break;
			case 2:
				{
				setState(870);
				statement();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariableStatementContext extends ParserRuleContext {
		public BindingPatternContext bindingPattern() {
			return getRuleContext(BindingPatternContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public VariableDeclarationListContext variableDeclarationList() {
			return getRuleContext(VariableDeclarationListContext.class,0);
		}
		public AccessibilityModifierContext accessibilityModifier() {
			return getRuleContext(AccessibilityModifierContext.class,0);
		}
		public VarModifierContext varModifier() {
			return getRuleContext(VarModifierContext.class,0);
		}
		public TerminalNode ReadOnly() { return getToken(TypeScriptParser.ReadOnly, 0); }
		public VariableStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterVariableStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitVariableStatement(this);
		}
	}

	public final VariableStatementContext variableStatement() throws RecognitionException {
		VariableStatementContext _localctx = new VariableStatementContext(_ctx, getState());
		enterRule(_localctx, 138, RULE_variableStatement);
		int _la;
		try {
			setState(894);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,97,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(873);
				bindingPattern();
				setState(875);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Colon) {
					{
					setState(874);
					typeAnnotation();
					}
				}

				setState(877);
				initializer();
				setState(879);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,92,_ctx) ) {
				case 1:
					{
					setState(878);
					match(SemiColon);
					}
					break;
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(882);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)))) != 0)) {
					{
					setState(881);
					accessibilityModifier();
					}
				}

				setState(885);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 75)) & ~0x3f) == 0 && ((1L << (_la - 75)) & ((1L << (Var - 75)) | (1L << (Const - 75)) | (1L << (Let - 75)))) != 0)) {
					{
					setState(884);
					varModifier();
					}
				}

				setState(888);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ReadOnly) {
					{
					setState(887);
					match(ReadOnly);
					}
				}

				setState(890);
				variableDeclarationList();
				setState(892);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,96,_ctx) ) {
				case 1:
					{
					setState(891);
					match(SemiColon);
					}
					break;
				}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariableDeclarationListContext extends ParserRuleContext {
		public List<VariableDeclarationContext> variableDeclaration() {
			return getRuleContexts(VariableDeclarationContext.class);
		}
		public VariableDeclarationContext variableDeclaration(int i) {
			return getRuleContext(VariableDeclarationContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public VariableDeclarationListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDeclarationList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterVariableDeclarationList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitVariableDeclarationList(this);
		}
	}

	public final VariableDeclarationListContext variableDeclarationList() throws RecognitionException {
		VariableDeclarationListContext _localctx = new VariableDeclarationListContext(_ctx, getState());
		enterRule(_localctx, 140, RULE_variableDeclarationList);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(896);
			variableDeclaration();
			setState(901);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,98,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(897);
					match(Comma);
					setState(898);
					variableDeclaration();
					}
					} 
				}
				setState(903);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,98,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariableDeclarationContext extends ParserRuleContext {
		public AssignableContext assignable() {
			return getRuleContext(AssignableContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode Assign() { return getToken(TypeScriptParser.Assign, 0); }
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public VariableDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variableDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterVariableDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitVariableDeclaration(this);
		}
	}

	public final VariableDeclarationContext variableDeclaration() throws RecognitionException {
		VariableDeclarationContext _localctx = new VariableDeclarationContext(_ctx, getState());
		enterRule(_localctx, 142, RULE_variableDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(904);
			assignable();
			setState(906);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,99,_ctx) ) {
			case 1:
				{
				setState(905);
				typeAnnotation();
				}
				break;
			}
			setState(909);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,100,_ctx) ) {
			case 1:
				{
				setState(908);
				singleExpression(0);
				}
				break;
			}
			setState(916);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,102,_ctx) ) {
			case 1:
				{
				setState(911);
				match(Assign);
				setState(913);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,101,_ctx) ) {
				case 1:
					{
					setState(912);
					typeParameters();
					}
					break;
				}
				setState(915);
				singleExpression(0);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EmptyStatement_Context extends ParserRuleContext {
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public EmptyStatement_Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_emptyStatement_; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterEmptyStatement_(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitEmptyStatement_(this);
		}
	}

	public final EmptyStatement_Context emptyStatement_() throws RecognitionException {
		EmptyStatement_Context _localctx = new EmptyStatement_Context(_ctx, getState());
		enterRule(_localctx, 144, RULE_emptyStatement_);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(918);
			match(SemiColon);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionStatementContext extends ParserRuleContext {
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public ExpressionStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterExpressionStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitExpressionStatement(this);
		}
	}

	public final ExpressionStatementContext expressionStatement() throws RecognitionException {
		ExpressionStatementContext _localctx = new ExpressionStatementContext(_ctx, getState());
		enterRule(_localctx, 146, RULE_expressionStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(920);
			if (!(this.notOpenBraceAndNotFunction())) throw new FailedPredicateException(this, "this.notOpenBraceAndNotFunction()");
			setState(921);
			expressionSequence();
			setState(923);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,103,_ctx) ) {
			case 1:
				{
				setState(922);
				match(SemiColon);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfStatementContext extends ParserRuleContext {
		public TerminalNode If() { return getToken(TypeScriptParser.If, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public TerminalNode Else() { return getToken(TypeScriptParser.Else, 0); }
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIfStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIfStatement(this);
		}
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 148, RULE_ifStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(925);
			match(If);
			setState(926);
			match(OpenParen);
			setState(927);
			expressionSequence();
			setState(928);
			match(CloseParen);
			setState(929);
			statement();
			setState(932);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,104,_ctx) ) {
			case 1:
				{
				setState(930);
				match(Else);
				setState(931);
				statement();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IterationStatementContext extends ParserRuleContext {
		public IterationStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iterationStatement; }
	 
		public IterationStatementContext() { }
		public void copyFrom(IterationStatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DoStatementContext extends IterationStatementContext {
		public TerminalNode Do() { return getToken(TypeScriptParser.Do, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode While() { return getToken(TypeScriptParser.While, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public DoStatementContext(IterationStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterDoStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitDoStatement(this);
		}
	}
	public static class ForVarStatementContext extends IterationStatementContext {
		public TerminalNode For() { return getToken(TypeScriptParser.For, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public VarModifierContext varModifier() {
			return getRuleContext(VarModifierContext.class,0);
		}
		public VariableDeclarationListContext variableDeclarationList() {
			return getRuleContext(VariableDeclarationListContext.class,0);
		}
		public List<TerminalNode> SemiColon() { return getTokens(TypeScriptParser.SemiColon); }
		public TerminalNode SemiColon(int i) {
			return getToken(TypeScriptParser.SemiColon, i);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public List<ExpressionSequenceContext> expressionSequence() {
			return getRuleContexts(ExpressionSequenceContext.class);
		}
		public ExpressionSequenceContext expressionSequence(int i) {
			return getRuleContext(ExpressionSequenceContext.class,i);
		}
		public ForVarStatementContext(IterationStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterForVarStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitForVarStatement(this);
		}
	}
	public static class ForVarInStatementContext extends IterationStatementContext {
		public TerminalNode For() { return getToken(TypeScriptParser.For, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public VarModifierContext varModifier() {
			return getRuleContext(VarModifierContext.class,0);
		}
		public VariableDeclarationContext variableDeclaration() {
			return getRuleContext(VariableDeclarationContext.class,0);
		}
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode In() { return getToken(TypeScriptParser.In, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode Await() { return getToken(TypeScriptParser.Await, 0); }
		public ForVarInStatementContext(IterationStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterForVarInStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitForVarInStatement(this);
		}
	}
	public static class WhileStatementContext extends IterationStatementContext {
		public TerminalNode While() { return getToken(TypeScriptParser.While, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public WhileStatementContext(IterationStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterWhileStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitWhileStatement(this);
		}
	}
	public static class ForStatementContext extends IterationStatementContext {
		public TerminalNode For() { return getToken(TypeScriptParser.For, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public List<TerminalNode> SemiColon() { return getTokens(TypeScriptParser.SemiColon); }
		public TerminalNode SemiColon(int i) {
			return getToken(TypeScriptParser.SemiColon, i);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public List<ExpressionSequenceContext> expressionSequence() {
			return getRuleContexts(ExpressionSequenceContext.class);
		}
		public ExpressionSequenceContext expressionSequence(int i) {
			return getRuleContext(ExpressionSequenceContext.class,i);
		}
		public ForStatementContext(IterationStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterForStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitForStatement(this);
		}
	}
	public static class ForInStatementContext extends IterationStatementContext {
		public TerminalNode For() { return getToken(TypeScriptParser.For, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode In() { return getToken(TypeScriptParser.In, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode Await() { return getToken(TypeScriptParser.Await, 0); }
		public ForInStatementContext(IterationStatementContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterForInStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitForInStatement(this);
		}
	}

	public final IterationStatementContext iterationStatement() throws RecognitionException {
		IterationStatementContext _localctx = new IterationStatementContext(_ctx, getState());
		enterRule(_localctx, 150, RULE_iterationStatement);
		int _la;
		try {
			setState(1009);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,114,_ctx) ) {
			case 1:
				_localctx = new DoStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(934);
				match(Do);
				setState(935);
				statement();
				setState(936);
				match(While);
				setState(937);
				match(OpenParen);
				setState(938);
				expressionSequence();
				setState(939);
				match(CloseParen);
				setState(940);
				eos();
				}
				break;
			case 2:
				_localctx = new WhileStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(942);
				match(While);
				setState(943);
				match(OpenParen);
				setState(944);
				expressionSequence();
				setState(945);
				match(CloseParen);
				setState(946);
				statement();
				}
				break;
			case 3:
				_localctx = new ForStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(948);
				match(For);
				setState(949);
				match(OpenParen);
				setState(951);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RegularExpressionLiteral) | (1L << OpenBracket) | (1L << OpenParen) | (1L << OpenBrace) | (1L << PlusPlus) | (1L << MinusMinus) | (1L << Plus) | (1L << Minus) | (1L << BitNot) | (1L << Not) | (1L << LessThan) | (1L << NullLiteral) | (1L << BooleanLiteral) | (1L << DecimalLiteral))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (HexIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral2 - 64)) | (1L << (BinaryIntegerLiteral - 64)) | (1L << (Break - 64)) | (1L << (Do - 64)) | (1L << (Instanceof - 64)) | (1L << (Typeof - 64)) | (1L << (Case - 64)) | (1L << (Else - 64)) | (1L << (New - 64)) | (1L << (Var - 64)) | (1L << (Catch - 64)) | (1L << (Finally - 64)) | (1L << (Return - 64)) | (1L << (Void - 64)) | (1L << (Continue - 64)) | (1L << (For - 64)) | (1L << (Switch - 64)) | (1L << (While - 64)) | (1L << (Debugger - 64)) | (1L << (Function - 64)) | (1L << (This - 64)) | (1L << (With - 64)) | (1L << (Default - 64)) | (1L << (If - 64)) | (1L << (Throw - 64)) | (1L << (Delete - 64)) | (1L << (In - 64)) | (1L << (Try - 64)) | (1L << (As - 64)) | (1L << (From - 64)) | (1L << (ReadOnly - 64)) | (1L << (Async - 64)) | (1L << (Class - 64)) | (1L << (Enum - 64)) | (1L << (Extends - 64)) | (1L << (Super - 64)) | (1L << (Const - 64)) | (1L << (Export - 64)) | (1L << (Import - 64)) | (1L << (Await - 64)) | (1L << (Implements - 64)) | (1L << (Let - 64)) | (1L << (Private - 64)) | (1L << (Public - 64)) | (1L << (Interface - 64)) | (1L << (Package - 64)) | (1L << (Protected - 64)) | (1L << (Static - 64)) | (1L << (Yield - 64)))) != 0) || ((((_la - 131)) & ~0x3f) == 0 && ((1L << (_la - 131)) & ((1L << (Identifier - 131)) | (1L << (StringLiteral - 131)) | (1L << (TemplateStringLiteral - 131)))) != 0)) {
					{
					setState(950);
					expressionSequence();
					}
				}

				setState(953);
				match(SemiColon);
				setState(955);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RegularExpressionLiteral) | (1L << OpenBracket) | (1L << OpenParen) | (1L << OpenBrace) | (1L << PlusPlus) | (1L << MinusMinus) | (1L << Plus) | (1L << Minus) | (1L << BitNot) | (1L << Not) | (1L << LessThan) | (1L << NullLiteral) | (1L << BooleanLiteral) | (1L << DecimalLiteral))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (HexIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral2 - 64)) | (1L << (BinaryIntegerLiteral - 64)) | (1L << (Break - 64)) | (1L << (Do - 64)) | (1L << (Instanceof - 64)) | (1L << (Typeof - 64)) | (1L << (Case - 64)) | (1L << (Else - 64)) | (1L << (New - 64)) | (1L << (Var - 64)) | (1L << (Catch - 64)) | (1L << (Finally - 64)) | (1L << (Return - 64)) | (1L << (Void - 64)) | (1L << (Continue - 64)) | (1L << (For - 64)) | (1L << (Switch - 64)) | (1L << (While - 64)) | (1L << (Debugger - 64)) | (1L << (Function - 64)) | (1L << (This - 64)) | (1L << (With - 64)) | (1L << (Default - 64)) | (1L << (If - 64)) | (1L << (Throw - 64)) | (1L << (Delete - 64)) | (1L << (In - 64)) | (1L << (Try - 64)) | (1L << (As - 64)) | (1L << (From - 64)) | (1L << (ReadOnly - 64)) | (1L << (Async - 64)) | (1L << (Class - 64)) | (1L << (Enum - 64)) | (1L << (Extends - 64)) | (1L << (Super - 64)) | (1L << (Const - 64)) | (1L << (Export - 64)) | (1L << (Import - 64)) | (1L << (Await - 64)) | (1L << (Implements - 64)) | (1L << (Let - 64)) | (1L << (Private - 64)) | (1L << (Public - 64)) | (1L << (Interface - 64)) | (1L << (Package - 64)) | (1L << (Protected - 64)) | (1L << (Static - 64)) | (1L << (Yield - 64)))) != 0) || ((((_la - 131)) & ~0x3f) == 0 && ((1L << (_la - 131)) & ((1L << (Identifier - 131)) | (1L << (StringLiteral - 131)) | (1L << (TemplateStringLiteral - 131)))) != 0)) {
					{
					setState(954);
					expressionSequence();
					}
				}

				setState(957);
				match(SemiColon);
				setState(959);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RegularExpressionLiteral) | (1L << OpenBracket) | (1L << OpenParen) | (1L << OpenBrace) | (1L << PlusPlus) | (1L << MinusMinus) | (1L << Plus) | (1L << Minus) | (1L << BitNot) | (1L << Not) | (1L << LessThan) | (1L << NullLiteral) | (1L << BooleanLiteral) | (1L << DecimalLiteral))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (HexIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral2 - 64)) | (1L << (BinaryIntegerLiteral - 64)) | (1L << (Break - 64)) | (1L << (Do - 64)) | (1L << (Instanceof - 64)) | (1L << (Typeof - 64)) | (1L << (Case - 64)) | (1L << (Else - 64)) | (1L << (New - 64)) | (1L << (Var - 64)) | (1L << (Catch - 64)) | (1L << (Finally - 64)) | (1L << (Return - 64)) | (1L << (Void - 64)) | (1L << (Continue - 64)) | (1L << (For - 64)) | (1L << (Switch - 64)) | (1L << (While - 64)) | (1L << (Debugger - 64)) | (1L << (Function - 64)) | (1L << (This - 64)) | (1L << (With - 64)) | (1L << (Default - 64)) | (1L << (If - 64)) | (1L << (Throw - 64)) | (1L << (Delete - 64)) | (1L << (In - 64)) | (1L << (Try - 64)) | (1L << (As - 64)) | (1L << (From - 64)) | (1L << (ReadOnly - 64)) | (1L << (Async - 64)) | (1L << (Class - 64)) | (1L << (Enum - 64)) | (1L << (Extends - 64)) | (1L << (Super - 64)) | (1L << (Const - 64)) | (1L << (Export - 64)) | (1L << (Import - 64)) | (1L << (Await - 64)) | (1L << (Implements - 64)) | (1L << (Let - 64)) | (1L << (Private - 64)) | (1L << (Public - 64)) | (1L << (Interface - 64)) | (1L << (Package - 64)) | (1L << (Protected - 64)) | (1L << (Static - 64)) | (1L << (Yield - 64)))) != 0) || ((((_la - 131)) & ~0x3f) == 0 && ((1L << (_la - 131)) & ((1L << (Identifier - 131)) | (1L << (StringLiteral - 131)) | (1L << (TemplateStringLiteral - 131)))) != 0)) {
					{
					setState(958);
					expressionSequence();
					}
				}

				setState(961);
				match(CloseParen);
				setState(962);
				statement();
				}
				break;
			case 4:
				_localctx = new ForVarStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(963);
				match(For);
				setState(964);
				match(OpenParen);
				setState(965);
				varModifier();
				setState(966);
				variableDeclarationList();
				setState(967);
				match(SemiColon);
				setState(969);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RegularExpressionLiteral) | (1L << OpenBracket) | (1L << OpenParen) | (1L << OpenBrace) | (1L << PlusPlus) | (1L << MinusMinus) | (1L << Plus) | (1L << Minus) | (1L << BitNot) | (1L << Not) | (1L << LessThan) | (1L << NullLiteral) | (1L << BooleanLiteral) | (1L << DecimalLiteral))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (HexIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral2 - 64)) | (1L << (BinaryIntegerLiteral - 64)) | (1L << (Break - 64)) | (1L << (Do - 64)) | (1L << (Instanceof - 64)) | (1L << (Typeof - 64)) | (1L << (Case - 64)) | (1L << (Else - 64)) | (1L << (New - 64)) | (1L << (Var - 64)) | (1L << (Catch - 64)) | (1L << (Finally - 64)) | (1L << (Return - 64)) | (1L << (Void - 64)) | (1L << (Continue - 64)) | (1L << (For - 64)) | (1L << (Switch - 64)) | (1L << (While - 64)) | (1L << (Debugger - 64)) | (1L << (Function - 64)) | (1L << (This - 64)) | (1L << (With - 64)) | (1L << (Default - 64)) | (1L << (If - 64)) | (1L << (Throw - 64)) | (1L << (Delete - 64)) | (1L << (In - 64)) | (1L << (Try - 64)) | (1L << (As - 64)) | (1L << (From - 64)) | (1L << (ReadOnly - 64)) | (1L << (Async - 64)) | (1L << (Class - 64)) | (1L << (Enum - 64)) | (1L << (Extends - 64)) | (1L << (Super - 64)) | (1L << (Const - 64)) | (1L << (Export - 64)) | (1L << (Import - 64)) | (1L << (Await - 64)) | (1L << (Implements - 64)) | (1L << (Let - 64)) | (1L << (Private - 64)) | (1L << (Public - 64)) | (1L << (Interface - 64)) | (1L << (Package - 64)) | (1L << (Protected - 64)) | (1L << (Static - 64)) | (1L << (Yield - 64)))) != 0) || ((((_la - 131)) & ~0x3f) == 0 && ((1L << (_la - 131)) & ((1L << (Identifier - 131)) | (1L << (StringLiteral - 131)) | (1L << (TemplateStringLiteral - 131)))) != 0)) {
					{
					setState(968);
					expressionSequence();
					}
				}

				setState(971);
				match(SemiColon);
				setState(973);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RegularExpressionLiteral) | (1L << OpenBracket) | (1L << OpenParen) | (1L << OpenBrace) | (1L << PlusPlus) | (1L << MinusMinus) | (1L << Plus) | (1L << Minus) | (1L << BitNot) | (1L << Not) | (1L << LessThan) | (1L << NullLiteral) | (1L << BooleanLiteral) | (1L << DecimalLiteral))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (HexIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral2 - 64)) | (1L << (BinaryIntegerLiteral - 64)) | (1L << (Break - 64)) | (1L << (Do - 64)) | (1L << (Instanceof - 64)) | (1L << (Typeof - 64)) | (1L << (Case - 64)) | (1L << (Else - 64)) | (1L << (New - 64)) | (1L << (Var - 64)) | (1L << (Catch - 64)) | (1L << (Finally - 64)) | (1L << (Return - 64)) | (1L << (Void - 64)) | (1L << (Continue - 64)) | (1L << (For - 64)) | (1L << (Switch - 64)) | (1L << (While - 64)) | (1L << (Debugger - 64)) | (1L << (Function - 64)) | (1L << (This - 64)) | (1L << (With - 64)) | (1L << (Default - 64)) | (1L << (If - 64)) | (1L << (Throw - 64)) | (1L << (Delete - 64)) | (1L << (In - 64)) | (1L << (Try - 64)) | (1L << (As - 64)) | (1L << (From - 64)) | (1L << (ReadOnly - 64)) | (1L << (Async - 64)) | (1L << (Class - 64)) | (1L << (Enum - 64)) | (1L << (Extends - 64)) | (1L << (Super - 64)) | (1L << (Const - 64)) | (1L << (Export - 64)) | (1L << (Import - 64)) | (1L << (Await - 64)) | (1L << (Implements - 64)) | (1L << (Let - 64)) | (1L << (Private - 64)) | (1L << (Public - 64)) | (1L << (Interface - 64)) | (1L << (Package - 64)) | (1L << (Protected - 64)) | (1L << (Static - 64)) | (1L << (Yield - 64)))) != 0) || ((((_la - 131)) & ~0x3f) == 0 && ((1L << (_la - 131)) & ((1L << (Identifier - 131)) | (1L << (StringLiteral - 131)) | (1L << (TemplateStringLiteral - 131)))) != 0)) {
					{
					setState(972);
					expressionSequence();
					}
				}

				setState(975);
				match(CloseParen);
				setState(976);
				statement();
				}
				break;
			case 5:
				_localctx = new ForInStatementContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(978);
				match(For);
				setState(980);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Await) {
					{
					setState(979);
					match(Await);
					}
				}

				setState(982);
				match(OpenParen);
				setState(983);
				singleExpression(0);
				setState(987);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case In:
					{
					setState(984);
					match(In);
					}
					break;
				case Identifier:
					{
					setState(985);
					match(Identifier);
					setState(986);
					if (!(this.p("of"))) throw new FailedPredicateException(this, "this.p(\"of\")");
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(989);
				expressionSequence();
				setState(990);
				match(CloseParen);
				setState(991);
				statement();
				}
				break;
			case 6:
				_localctx = new ForVarInStatementContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(993);
				match(For);
				setState(995);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Await) {
					{
					setState(994);
					match(Await);
					}
				}

				setState(997);
				match(OpenParen);
				setState(998);
				varModifier();
				setState(999);
				variableDeclaration();
				setState(1003);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case In:
					{
					setState(1000);
					match(In);
					}
					break;
				case Identifier:
					{
					setState(1001);
					match(Identifier);
					setState(1002);
					if (!(this.p("of"))) throw new FailedPredicateException(this, "this.p(\"of\")");
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(1005);
				expressionSequence();
				setState(1006);
				match(CloseParen);
				setState(1007);
				statement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VarModifierContext extends ParserRuleContext {
		public TerminalNode Var() { return getToken(TypeScriptParser.Var, 0); }
		public TerminalNode Let() { return getToken(TypeScriptParser.Let, 0); }
		public TerminalNode Const() { return getToken(TypeScriptParser.Const, 0); }
		public VarModifierContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varModifier; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterVarModifier(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitVarModifier(this);
		}
	}

	public final VarModifierContext varModifier() throws RecognitionException {
		VarModifierContext _localctx = new VarModifierContext(_ctx, getState());
		enterRule(_localctx, 152, RULE_varModifier);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1011);
			_la = _input.LA(1);
			if ( !(((((_la - 75)) & ~0x3f) == 0 && ((1L << (_la - 75)) & ((1L << (Var - 75)) | (1L << (Const - 75)) | (1L << (Let - 75)))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ContinueStatementContext extends ParserRuleContext {
		public TerminalNode Continue() { return getToken(TypeScriptParser.Continue, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public ContinueStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continueStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterContinueStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitContinueStatement(this);
		}
	}

	public final ContinueStatementContext continueStatement() throws RecognitionException {
		ContinueStatementContext _localctx = new ContinueStatementContext(_ctx, getState());
		enterRule(_localctx, 154, RULE_continueStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1013);
			match(Continue);
			setState(1016);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,115,_ctx) ) {
			case 1:
				{
				setState(1014);
				if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
				setState(1015);
				match(Identifier);
				}
				break;
			}
			setState(1018);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BreakStatementContext extends ParserRuleContext {
		public TerminalNode Break() { return getToken(TypeScriptParser.Break, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public BreakStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterBreakStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitBreakStatement(this);
		}
	}

	public final BreakStatementContext breakStatement() throws RecognitionException {
		BreakStatementContext _localctx = new BreakStatementContext(_ctx, getState());
		enterRule(_localctx, 156, RULE_breakStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1020);
			match(Break);
			setState(1023);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,116,_ctx) ) {
			case 1:
				{
				setState(1021);
				if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
				setState(1022);
				match(Identifier);
				}
				break;
			}
			setState(1025);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReturnStatementContext extends ParserRuleContext {
		public TerminalNode Return() { return getToken(TypeScriptParser.Return, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public ReturnStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterReturnStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitReturnStatement(this);
		}
	}

	public final ReturnStatementContext returnStatement() throws RecognitionException {
		ReturnStatementContext _localctx = new ReturnStatementContext(_ctx, getState());
		enterRule(_localctx, 158, RULE_returnStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1027);
			match(Return);
			setState(1030);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,117,_ctx) ) {
			case 1:
				{
				setState(1028);
				if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
				setState(1029);
				expressionSequence();
				}
				break;
			}
			setState(1032);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class YieldStatementContext extends ParserRuleContext {
		public TerminalNode Yield() { return getToken(TypeScriptParser.Yield, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public YieldStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_yieldStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterYieldStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitYieldStatement(this);
		}
	}

	public final YieldStatementContext yieldStatement() throws RecognitionException {
		YieldStatementContext _localctx = new YieldStatementContext(_ctx, getState());
		enterRule(_localctx, 160, RULE_yieldStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1034);
			match(Yield);
			setState(1037);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,118,_ctx) ) {
			case 1:
				{
				setState(1035);
				if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
				setState(1036);
				expressionSequence();
				}
				break;
			}
			setState(1039);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WithStatementContext extends ParserRuleContext {
		public TerminalNode With() { return getToken(TypeScriptParser.With, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public WithStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_withStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterWithStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitWithStatement(this);
		}
	}

	public final WithStatementContext withStatement() throws RecognitionException {
		WithStatementContext _localctx = new WithStatementContext(_ctx, getState());
		enterRule(_localctx, 162, RULE_withStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1041);
			match(With);
			setState(1042);
			match(OpenParen);
			setState(1043);
			expressionSequence();
			setState(1044);
			match(CloseParen);
			setState(1045);
			statement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SwitchStatementContext extends ParserRuleContext {
		public TerminalNode Switch() { return getToken(TypeScriptParser.Switch, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public CaseBlockContext caseBlock() {
			return getRuleContext(CaseBlockContext.class,0);
		}
		public SwitchStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterSwitchStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitSwitchStatement(this);
		}
	}

	public final SwitchStatementContext switchStatement() throws RecognitionException {
		SwitchStatementContext _localctx = new SwitchStatementContext(_ctx, getState());
		enterRule(_localctx, 164, RULE_switchStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1047);
			match(Switch);
			setState(1048);
			match(OpenParen);
			setState(1049);
			expressionSequence();
			setState(1050);
			match(CloseParen);
			setState(1051);
			caseBlock();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CaseBlockContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public List<CaseClausesContext> caseClauses() {
			return getRuleContexts(CaseClausesContext.class);
		}
		public CaseClausesContext caseClauses(int i) {
			return getRuleContext(CaseClausesContext.class,i);
		}
		public DefaultClauseContext defaultClause() {
			return getRuleContext(DefaultClauseContext.class,0);
		}
		public CaseBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterCaseBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitCaseBlock(this);
		}
	}

	public final CaseBlockContext caseBlock() throws RecognitionException {
		CaseBlockContext _localctx = new CaseBlockContext(_ctx, getState());
		enterRule(_localctx, 166, RULE_caseBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1053);
			match(OpenBrace);
			setState(1055);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Case) {
				{
				setState(1054);
				caseClauses();
				}
			}

			setState(1061);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Default) {
				{
				setState(1057);
				defaultClause();
				setState(1059);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Case) {
					{
					setState(1058);
					caseClauses();
					}
				}

				}
			}

			setState(1063);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CaseClausesContext extends ParserRuleContext {
		public List<CaseClauseContext> caseClause() {
			return getRuleContexts(CaseClauseContext.class);
		}
		public CaseClauseContext caseClause(int i) {
			return getRuleContext(CaseClauseContext.class,i);
		}
		public CaseClausesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseClauses; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterCaseClauses(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitCaseClauses(this);
		}
	}

	public final CaseClausesContext caseClauses() throws RecognitionException {
		CaseClausesContext _localctx = new CaseClausesContext(_ctx, getState());
		enterRule(_localctx, 168, RULE_caseClauses);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1066); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(1065);
				caseClause();
				}
				}
				setState(1068); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==Case );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CaseClauseContext extends ParserRuleContext {
		public TerminalNode Case() { return getToken(TypeScriptParser.Case, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public CaseClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterCaseClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitCaseClause(this);
		}
	}

	public final CaseClauseContext caseClause() throws RecognitionException {
		CaseClauseContext _localctx = new CaseClauseContext(_ctx, getState());
		enterRule(_localctx, 170, RULE_caseClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1070);
			match(Case);
			setState(1071);
			expressionSequence();
			setState(1072);
			match(Colon);
			setState(1074);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,123,_ctx) ) {
			case 1:
				{
				setState(1073);
				statementList();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DefaultClauseContext extends ParserRuleContext {
		public TerminalNode Default() { return getToken(TypeScriptParser.Default, 0); }
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public StatementListContext statementList() {
			return getRuleContext(StatementListContext.class,0);
		}
		public DefaultClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterDefaultClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitDefaultClause(this);
		}
	}

	public final DefaultClauseContext defaultClause() throws RecognitionException {
		DefaultClauseContext _localctx = new DefaultClauseContext(_ctx, getState());
		enterRule(_localctx, 172, RULE_defaultClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1076);
			match(Default);
			setState(1077);
			match(Colon);
			setState(1079);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,124,_ctx) ) {
			case 1:
				{
				setState(1078);
				statementList();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LabelledStatementContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public LabelledStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_labelledStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterLabelledStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitLabelledStatement(this);
		}
	}

	public final LabelledStatementContext labelledStatement() throws RecognitionException {
		LabelledStatementContext _localctx = new LabelledStatementContext(_ctx, getState());
		enterRule(_localctx, 174, RULE_labelledStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1081);
			match(Identifier);
			setState(1082);
			match(Colon);
			setState(1083);
			statement();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ThrowStatementContext extends ParserRuleContext {
		public TerminalNode Throw() { return getToken(TypeScriptParser.Throw, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public ThrowStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_throwStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterThrowStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitThrowStatement(this);
		}
	}

	public final ThrowStatementContext throwStatement() throws RecognitionException {
		ThrowStatementContext _localctx = new ThrowStatementContext(_ctx, getState());
		enterRule(_localctx, 176, RULE_throwStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1085);
			match(Throw);
			setState(1086);
			if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
			setState(1087);
			expressionSequence();
			setState(1088);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TryStatementContext extends ParserRuleContext {
		public TerminalNode Try() { return getToken(TypeScriptParser.Try, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public CatchProductionContext catchProduction() {
			return getRuleContext(CatchProductionContext.class,0);
		}
		public FinallyProductionContext finallyProduction() {
			return getRuleContext(FinallyProductionContext.class,0);
		}
		public TryStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tryStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTryStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTryStatement(this);
		}
	}

	public final TryStatementContext tryStatement() throws RecognitionException {
		TryStatementContext _localctx = new TryStatementContext(_ctx, getState());
		enterRule(_localctx, 178, RULE_tryStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1090);
			match(Try);
			setState(1091);
			block();
			setState(1097);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Catch:
				{
				setState(1092);
				catchProduction();
				setState(1094);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,125,_ctx) ) {
				case 1:
					{
					setState(1093);
					finallyProduction();
					}
					break;
				}
				}
				break;
			case Finally:
				{
				setState(1096);
				finallyProduction();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CatchProductionContext extends ParserRuleContext {
		public TerminalNode Catch() { return getToken(TypeScriptParser.Catch, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public AssignableContext assignable() {
			return getRuleContext(AssignableContext.class,0);
		}
		public CatchProductionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_catchProduction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterCatchProduction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitCatchProduction(this);
		}
	}

	public final CatchProductionContext catchProduction() throws RecognitionException {
		CatchProductionContext _localctx = new CatchProductionContext(_ctx, getState());
		enterRule(_localctx, 180, RULE_catchProduction);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1099);
			match(Catch);
			setState(1105);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OpenParen) {
				{
				setState(1100);
				match(OpenParen);
				setState(1102);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==OpenBracket || _la==OpenBrace || _la==Identifier) {
					{
					setState(1101);
					assignable();
					}
				}

				setState(1104);
				match(CloseParen);
				}
			}

			setState(1107);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignableContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public ArrayLiteralContext arrayLiteral() {
			return getRuleContext(ArrayLiteralContext.class,0);
		}
		public ObjectLiteralContext objectLiteral() {
			return getRuleContext(ObjectLiteralContext.class,0);
		}
		public AssignableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignable; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterAssignable(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitAssignable(this);
		}
	}

	public final AssignableContext assignable() throws RecognitionException {
		AssignableContext _localctx = new AssignableContext(_ctx, getState());
		enterRule(_localctx, 182, RULE_assignable);
		try {
			setState(1112);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(1109);
				match(Identifier);
				}
				break;
			case OpenBracket:
				enterOuterAlt(_localctx, 2);
				{
				setState(1110);
				arrayLiteral();
				}
				break;
			case OpenBrace:
				enterOuterAlt(_localctx, 3);
				{
				setState(1111);
				objectLiteral();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FinallyProductionContext extends ParserRuleContext {
		public TerminalNode Finally() { return getToken(TypeScriptParser.Finally, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FinallyProductionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_finallyProduction; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFinallyProduction(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFinallyProduction(this);
		}
	}

	public final FinallyProductionContext finallyProduction() throws RecognitionException {
		FinallyProductionContext _localctx = new FinallyProductionContext(_ctx, getState());
		enterRule(_localctx, 184, RULE_finallyProduction);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1114);
			match(Finally);
			setState(1115);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DebuggerStatementContext extends ParserRuleContext {
		public TerminalNode Debugger() { return getToken(TypeScriptParser.Debugger, 0); }
		public EosContext eos() {
			return getRuleContext(EosContext.class,0);
		}
		public DebuggerStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_debuggerStatement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterDebuggerStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitDebuggerStatement(this);
		}
	}

	public final DebuggerStatementContext debuggerStatement() throws RecognitionException {
		DebuggerStatementContext _localctx = new DebuggerStatementContext(_ctx, getState());
		enterRule(_localctx, 186, RULE_debuggerStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1117);
			match(Debugger);
			setState(1118);
			eos();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionDeclarationContext extends ParserRuleContext {
		public TerminalNode Function() { return getToken(TypeScriptParser.Function, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public CallSignatureContext callSignature() {
			return getRuleContext(CallSignatureContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public TerminalNode Async() { return getToken(TypeScriptParser.Async, 0); }
		public TerminalNode Multiply() { return getToken(TypeScriptParser.Multiply, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public FunctionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFunctionDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFunctionDeclaration(this);
		}
	}

	public final FunctionDeclarationContext functionDeclaration() throws RecognitionException {
		FunctionDeclarationContext _localctx = new FunctionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 188, RULE_functionDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1121);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Async) {
				{
				setState(1120);
				match(Async);
				}
			}

			setState(1123);
			match(Function);
			setState(1125);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Multiply) {
				{
				setState(1124);
				match(Multiply);
				}
			}

			setState(1127);
			match(Identifier);
			setState(1128);
			callSignature();
			setState(1134);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OpenBrace:
				{
				{
				setState(1129);
				match(OpenBrace);
				setState(1130);
				functionBody();
				setState(1131);
				match(CloseBrace);
				}
				}
				break;
			case SemiColon:
				{
				setState(1133);
				match(SemiColon);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ClassDeclarationContext extends ParserRuleContext {
		public TerminalNode Class() { return getToken(TypeScriptParser.Class, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public ClassHeritageContext classHeritage() {
			return getRuleContext(ClassHeritageContext.class,0);
		}
		public ClassTailContext classTail() {
			return getRuleContext(ClassTailContext.class,0);
		}
		public TerminalNode Abstract() { return getToken(TypeScriptParser.Abstract, 0); }
		public TypeParametersContext typeParameters() {
			return getRuleContext(TypeParametersContext.class,0);
		}
		public ClassDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterClassDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitClassDeclaration(this);
		}
	}

	public final ClassDeclarationContext classDeclaration() throws RecognitionException {
		ClassDeclarationContext _localctx = new ClassDeclarationContext(_ctx, getState());
		enterRule(_localctx, 190, RULE_classDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Abstract) {
				{
				setState(1136);
				match(Abstract);
				}
			}

			setState(1139);
			match(Class);
			setState(1140);
			match(Identifier);
			setState(1142);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LessThan) {
				{
				setState(1141);
				typeParameters();
				}
			}

			setState(1144);
			classHeritage();
			setState(1145);
			classTail();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ClassHeritageContext extends ParserRuleContext {
		public ClassExtendsClauseContext classExtendsClause() {
			return getRuleContext(ClassExtendsClauseContext.class,0);
		}
		public ImplementsClauseContext implementsClause() {
			return getRuleContext(ImplementsClauseContext.class,0);
		}
		public ClassHeritageContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classHeritage; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterClassHeritage(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitClassHeritage(this);
		}
	}

	public final ClassHeritageContext classHeritage() throws RecognitionException {
		ClassHeritageContext _localctx = new ClassHeritageContext(_ctx, getState());
		enterRule(_localctx, 192, RULE_classHeritage);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1148);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Extends) {
				{
				setState(1147);
				classExtendsClause();
				}
			}

			setState(1151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Implements) {
				{
				setState(1150);
				implementsClause();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ClassTailContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public List<ClassElementContext> classElement() {
			return getRuleContexts(ClassElementContext.class);
		}
		public ClassElementContext classElement(int i) {
			return getRuleContext(ClassElementContext.class,i);
		}
		public ClassTailContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classTail; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterClassTail(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitClassTail(this);
		}
	}

	public final ClassTailContext classTail() throws RecognitionException {
		ClassTailContext _localctx = new ClassTailContext(_ctx, getState());
		enterRule(_localctx, 194, RULE_classTail);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1153);
			match(OpenBrace);
			setState(1157);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,137,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1154);
					classElement();
					}
					} 
				}
				setState(1159);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,137,_ctx);
			}
			setState(1160);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ClassExtendsClauseContext extends ParserRuleContext {
		public TerminalNode Extends() { return getToken(TypeScriptParser.Extends, 0); }
		public TypeReferenceContext typeReference() {
			return getRuleContext(TypeReferenceContext.class,0);
		}
		public ClassExtendsClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classExtendsClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterClassExtendsClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitClassExtendsClause(this);
		}
	}

	public final ClassExtendsClauseContext classExtendsClause() throws RecognitionException {
		ClassExtendsClauseContext _localctx = new ClassExtendsClauseContext(_ctx, getState());
		enterRule(_localctx, 196, RULE_classExtendsClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1162);
			match(Extends);
			setState(1163);
			typeReference();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImplementsClauseContext extends ParserRuleContext {
		public TerminalNode Implements() { return getToken(TypeScriptParser.Implements, 0); }
		public ClassOrInterfaceTypeListContext classOrInterfaceTypeList() {
			return getRuleContext(ClassOrInterfaceTypeListContext.class,0);
		}
		public ImplementsClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_implementsClause; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterImplementsClause(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitImplementsClause(this);
		}
	}

	public final ImplementsClauseContext implementsClause() throws RecognitionException {
		ImplementsClauseContext _localctx = new ImplementsClauseContext(_ctx, getState());
		enterRule(_localctx, 198, RULE_implementsClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1165);
			match(Implements);
			setState(1166);
			classOrInterfaceTypeList();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ClassElementContext extends ParserRuleContext {
		public ConstructorDeclarationContext constructorDeclaration() {
			return getRuleContext(ConstructorDeclarationContext.class,0);
		}
		public PropertyMemberDeclarationContext propertyMemberDeclaration() {
			return getRuleContext(PropertyMemberDeclarationContext.class,0);
		}
		public IndexMemberDeclarationContext indexMemberDeclaration() {
			return getRuleContext(IndexMemberDeclarationContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public ClassElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_classElement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterClassElement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitClassElement(this);
		}
	}

	public final ClassElementContext classElement() throws RecognitionException {
		ClassElementContext _localctx = new ClassElementContext(_ctx, getState());
		enterRule(_localctx, 200, RULE_classElement);
		try {
			setState(1172);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,138,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1168);
				constructorDeclaration();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1169);
				propertyMemberDeclaration();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1170);
				indexMemberDeclaration();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1171);
				statement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PropertyMemberDeclarationContext extends ParserRuleContext {
		public PropertyMemberBaseContext propertyMemberBase() {
			return getRuleContext(PropertyMemberBaseContext.class,0);
		}
		public PropertyNameContext propertyName() {
			return getRuleContext(PropertyNameContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public TerminalNode Multiply() { return getToken(TypeScriptParser.Multiply, 0); }
		public TerminalNode Hashtag() { return getToken(TypeScriptParser.Hashtag, 0); }
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public InitializerContext initializer() {
			return getRuleContext(InitializerContext.class,0);
		}
		public CallSignatureContext callSignature() {
			return getRuleContext(CallSignatureContext.class,0);
		}
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public GetAccessorContext getAccessor() {
			return getRuleContext(GetAccessorContext.class,0);
		}
		public SetAccessorContext setAccessor() {
			return getRuleContext(SetAccessorContext.class,0);
		}
		public AbstractDeclarationContext abstractDeclaration() {
			return getRuleContext(AbstractDeclarationContext.class,0);
		}
		public PropertyMemberDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertyMemberDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPropertyMemberDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPropertyMemberDeclaration(this);
		}
	}

	public final PropertyMemberDeclarationContext propertyMemberDeclaration() throws RecognitionException {
		PropertyMemberDeclarationContext _localctx = new PropertyMemberDeclarationContext(_ctx, getState());
		enterRule(_localctx, 202, RULE_propertyMemberDeclaration);
		int _la;
		try {
			setState(1218);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,149,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1174);
				propertyMemberBase();
				setState(1176);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Multiply) {
					{
					setState(1175);
					match(Multiply);
					}
				}

				setState(1179);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Hashtag) {
					{
					setState(1178);
					match(Hashtag);
					}
				}

				setState(1181);
				propertyName();
				setState(1183);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Colon) {
					{
					setState(1182);
					typeAnnotation();
					}
				}

				setState(1186);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Assign) {
					{
					setState(1185);
					initializer();
					}
				}

				setState(1188);
				match(SemiColon);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1190);
				propertyMemberBase();
				setState(1192);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Multiply) {
					{
					setState(1191);
					match(Multiply);
					}
				}

				setState(1195);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Hashtag) {
					{
					setState(1194);
					match(Hashtag);
					}
				}

				setState(1197);
				propertyName();
				setState(1198);
				callSignature();
				setState(1204);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case OpenBrace:
					{
					{
					setState(1199);
					match(OpenBrace);
					setState(1200);
					functionBody();
					setState(1201);
					match(CloseBrace);
					}
					}
					break;
				case SemiColon:
					{
					setState(1203);
					match(SemiColon);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1206);
				propertyMemberBase();
				setState(1208);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Multiply) {
					{
					setState(1207);
					match(Multiply);
					}
				}

				setState(1211);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Hashtag) {
					{
					setState(1210);
					match(Hashtag);
					}
				}

				setState(1215);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,148,_ctx) ) {
				case 1:
					{
					setState(1213);
					getAccessor();
					}
					break;
				case 2:
					{
					setState(1214);
					setAccessor();
					}
					break;
				}
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1217);
				abstractDeclaration();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PropertyMemberBaseContext extends ParserRuleContext {
		public TerminalNode Async() { return getToken(TypeScriptParser.Async, 0); }
		public AccessibilityModifierContext accessibilityModifier() {
			return getRuleContext(AccessibilityModifierContext.class,0);
		}
		public TerminalNode Static() { return getToken(TypeScriptParser.Static, 0); }
		public TerminalNode ReadOnly() { return getToken(TypeScriptParser.ReadOnly, 0); }
		public PropertyMemberBaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertyMemberBase; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPropertyMemberBase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPropertyMemberBase(this);
		}
	}

	public final PropertyMemberBaseContext propertyMemberBase() throws RecognitionException {
		PropertyMemberBaseContext _localctx = new PropertyMemberBaseContext(_ctx, getState());
		enterRule(_localctx, 204, RULE_propertyMemberBase);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1221);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,150,_ctx) ) {
			case 1:
				{
				setState(1220);
				match(Async);
				}
				break;
			}
			setState(1224);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,151,_ctx) ) {
			case 1:
				{
				setState(1223);
				accessibilityModifier();
				}
				break;
			}
			setState(1227);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,152,_ctx) ) {
			case 1:
				{
				setState(1226);
				match(Static);
				}
				break;
			}
			setState(1230);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,153,_ctx) ) {
			case 1:
				{
				setState(1229);
				match(ReadOnly);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IndexMemberDeclarationContext extends ParserRuleContext {
		public IndexSignatureContext indexSignature() {
			return getRuleContext(IndexSignatureContext.class,0);
		}
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public IndexMemberDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_indexMemberDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIndexMemberDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIndexMemberDeclaration(this);
		}
	}

	public final IndexMemberDeclarationContext indexMemberDeclaration() throws RecognitionException {
		IndexMemberDeclarationContext _localctx = new IndexMemberDeclarationContext(_ctx, getState());
		enterRule(_localctx, 206, RULE_indexMemberDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1232);
			indexSignature();
			setState(1233);
			match(SemiColon);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GeneratorMethodContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public TerminalNode Multiply() { return getToken(TypeScriptParser.Multiply, 0); }
		public FormalParameterListContext formalParameterList() {
			return getRuleContext(FormalParameterListContext.class,0);
		}
		public GeneratorMethodContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generatorMethod; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGeneratorMethod(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGeneratorMethod(this);
		}
	}

	public final GeneratorMethodContext generatorMethod() throws RecognitionException {
		GeneratorMethodContext _localctx = new GeneratorMethodContext(_ctx, getState());
		enterRule(_localctx, 208, RULE_generatorMethod);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1236);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Multiply) {
				{
				setState(1235);
				match(Multiply);
				}
			}

			setState(1238);
			match(Identifier);
			setState(1239);
			match(OpenParen);
			setState(1241);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis))) != 0) || ((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)) | (1L << (Identifier - 108)))) != 0)) {
				{
				setState(1240);
				formalParameterList();
				}
			}

			setState(1243);
			match(CloseParen);
			setState(1244);
			match(OpenBrace);
			setState(1245);
			functionBody();
			setState(1246);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GeneratorFunctionDeclarationContext extends ParserRuleContext {
		public TerminalNode Function() { return getToken(TypeScriptParser.Function, 0); }
		public TerminalNode Multiply() { return getToken(TypeScriptParser.Multiply, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public FormalParameterListContext formalParameterList() {
			return getRuleContext(FormalParameterListContext.class,0);
		}
		public GeneratorFunctionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generatorFunctionDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGeneratorFunctionDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGeneratorFunctionDeclaration(this);
		}
	}

	public final GeneratorFunctionDeclarationContext generatorFunctionDeclaration() throws RecognitionException {
		GeneratorFunctionDeclarationContext _localctx = new GeneratorFunctionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 210, RULE_generatorFunctionDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1248);
			match(Function);
			setState(1249);
			match(Multiply);
			setState(1251);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(1250);
				match(Identifier);
				}
			}

			setState(1253);
			match(OpenParen);
			setState(1255);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis))) != 0) || ((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)) | (1L << (Identifier - 108)))) != 0)) {
				{
				setState(1254);
				formalParameterList();
				}
			}

			setState(1257);
			match(CloseParen);
			setState(1258);
			match(OpenBrace);
			setState(1259);
			functionBody();
			setState(1260);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GeneratorBlockContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public List<GeneratorDefinitionContext> generatorDefinition() {
			return getRuleContexts(GeneratorDefinitionContext.class);
		}
		public GeneratorDefinitionContext generatorDefinition(int i) {
			return getRuleContext(GeneratorDefinitionContext.class,i);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public GeneratorBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generatorBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGeneratorBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGeneratorBlock(this);
		}
	}

	public final GeneratorBlockContext generatorBlock() throws RecognitionException {
		GeneratorBlockContext _localctx = new GeneratorBlockContext(_ctx, getState());
		enterRule(_localctx, 212, RULE_generatorBlock);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1262);
			match(OpenBrace);
			setState(1263);
			generatorDefinition();
			setState(1268);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,158,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1264);
					match(Comma);
					setState(1265);
					generatorDefinition();
					}
					} 
				}
				setState(1270);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,158,_ctx);
			}
			setState(1272);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Comma) {
				{
				setState(1271);
				match(Comma);
				}
			}

			setState(1274);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GeneratorDefinitionContext extends ParserRuleContext {
		public TerminalNode Multiply() { return getToken(TypeScriptParser.Multiply, 0); }
		public IteratorDefinitionContext iteratorDefinition() {
			return getRuleContext(IteratorDefinitionContext.class,0);
		}
		public GeneratorDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_generatorDefinition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGeneratorDefinition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGeneratorDefinition(this);
		}
	}

	public final GeneratorDefinitionContext generatorDefinition() throws RecognitionException {
		GeneratorDefinitionContext _localctx = new GeneratorDefinitionContext(_ctx, getState());
		enterRule(_localctx, 214, RULE_generatorDefinition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1276);
			match(Multiply);
			setState(1277);
			iteratorDefinition();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IteratorBlockContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public List<IteratorDefinitionContext> iteratorDefinition() {
			return getRuleContexts(IteratorDefinitionContext.class);
		}
		public IteratorDefinitionContext iteratorDefinition(int i) {
			return getRuleContext(IteratorDefinitionContext.class,i);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public IteratorBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iteratorBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIteratorBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIteratorBlock(this);
		}
	}

	public final IteratorBlockContext iteratorBlock() throws RecognitionException {
		IteratorBlockContext _localctx = new IteratorBlockContext(_ctx, getState());
		enterRule(_localctx, 216, RULE_iteratorBlock);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1279);
			match(OpenBrace);
			setState(1280);
			iteratorDefinition();
			setState(1285);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,160,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1281);
					match(Comma);
					setState(1282);
					iteratorDefinition();
					}
					} 
				}
				setState(1287);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,160,_ctx);
			}
			setState(1289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Comma) {
				{
				setState(1288);
				match(Comma);
				}
			}

			setState(1291);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IteratorDefinitionContext extends ParserRuleContext {
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public FormalParameterListContext formalParameterList() {
			return getRuleContext(FormalParameterListContext.class,0);
		}
		public IteratorDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iteratorDefinition; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIteratorDefinition(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIteratorDefinition(this);
		}
	}

	public final IteratorDefinitionContext iteratorDefinition() throws RecognitionException {
		IteratorDefinitionContext _localctx = new IteratorDefinitionContext(_ctx, getState());
		enterRule(_localctx, 218, RULE_iteratorDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1293);
			match(OpenBracket);
			setState(1294);
			singleExpression(0);
			setState(1295);
			match(CloseBracket);
			setState(1296);
			match(OpenParen);
			setState(1298);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis))) != 0) || ((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)) | (1L << (Identifier - 108)))) != 0)) {
				{
				setState(1297);
				formalParameterList();
				}
			}

			setState(1300);
			match(CloseParen);
			setState(1301);
			match(OpenBrace);
			setState(1302);
			functionBody();
			setState(1303);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FormalParameterListContext extends ParserRuleContext {
		public List<FormalParameterArgContext> formalParameterArg() {
			return getRuleContexts(FormalParameterArgContext.class);
		}
		public FormalParameterArgContext formalParameterArg(int i) {
			return getRuleContext(FormalParameterArgContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public LastFormalParameterArgContext lastFormalParameterArg() {
			return getRuleContext(LastFormalParameterArgContext.class,0);
		}
		public ArrayLiteralContext arrayLiteral() {
			return getRuleContext(ArrayLiteralContext.class,0);
		}
		public ObjectLiteralContext objectLiteral() {
			return getRuleContext(ObjectLiteralContext.class,0);
		}
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public FormalParameterListContext formalParameterList() {
			return getRuleContext(FormalParameterListContext.class,0);
		}
		public FormalParameterListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formalParameterList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFormalParameterList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFormalParameterList(this);
		}
	}

	public final FormalParameterListContext formalParameterList() throws RecognitionException {
		FormalParameterListContext _localctx = new FormalParameterListContext(_ctx, getState());
		enterRule(_localctx, 220, RULE_formalParameterList);
		int _la;
		try {
			int _alt;
			setState(1324);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Private:
			case Public:
			case Protected:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(1305);
				formalParameterArg();
				setState(1310);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,163,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1306);
						match(Comma);
						setState(1307);
						formalParameterArg();
						}
						} 
					}
					setState(1312);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,163,_ctx);
				}
				setState(1315);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Comma) {
					{
					setState(1313);
					match(Comma);
					setState(1314);
					lastFormalParameterArg();
					}
				}

				}
				break;
			case Ellipsis:
				enterOuterAlt(_localctx, 2);
				{
				setState(1317);
				lastFormalParameterArg();
				}
				break;
			case OpenBracket:
				enterOuterAlt(_localctx, 3);
				{
				setState(1318);
				arrayLiteral();
				}
				break;
			case OpenBrace:
				enterOuterAlt(_localctx, 4);
				{
				setState(1319);
				objectLiteral();
				setState(1322);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Colon) {
					{
					setState(1320);
					match(Colon);
					setState(1321);
					formalParameterList();
					}
				}

				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FormalParameterArgContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public AccessibilityModifierContext accessibilityModifier() {
			return getRuleContext(AccessibilityModifierContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public TerminalNode Assign() { return getToken(TypeScriptParser.Assign, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public FormalParameterArgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formalParameterArg; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFormalParameterArg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFormalParameterArg(this);
		}
	}

	public final FormalParameterArgContext formalParameterArg() throws RecognitionException {
		FormalParameterArgContext _localctx = new FormalParameterArgContext(_ctx, getState());
		enterRule(_localctx, 222, RULE_formalParameterArg);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1327);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)))) != 0)) {
				{
				setState(1326);
				accessibilityModifier();
				}
			}

			setState(1329);
			match(Identifier);
			setState(1331);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(1330);
				typeAnnotation();
				}
			}

			setState(1335);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Assign) {
				{
				setState(1333);
				match(Assign);
				setState(1334);
				singleExpression(0);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LastFormalParameterArgContext extends ParserRuleContext {
		public TerminalNode Ellipsis() { return getToken(TypeScriptParser.Ellipsis, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public LastFormalParameterArgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lastFormalParameterArg; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterLastFormalParameterArg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitLastFormalParameterArg(this);
		}
	}

	public final LastFormalParameterArgContext lastFormalParameterArg() throws RecognitionException {
		LastFormalParameterArgContext _localctx = new LastFormalParameterArgContext(_ctx, getState());
		enterRule(_localctx, 224, RULE_lastFormalParameterArg);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1337);
			match(Ellipsis);
			setState(1338);
			match(Identifier);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionBodyContext extends ParserRuleContext {
		public SourceElementsContext sourceElements() {
			return getRuleContext(SourceElementsContext.class,0);
		}
		public FunctionBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionBody; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFunctionBody(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFunctionBody(this);
		}
	}

	public final FunctionBodyContext functionBody() throws RecognitionException {
		FunctionBodyContext _localctx = new FunctionBodyContext(_ctx, getState());
		enterRule(_localctx, 226, RULE_functionBody);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1341);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,170,_ctx) ) {
			case 1:
				{
				setState(1340);
				sourceElements();
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SourceElementsContext extends ParserRuleContext {
		public List<SourceElementContext> sourceElement() {
			return getRuleContexts(SourceElementContext.class);
		}
		public SourceElementContext sourceElement(int i) {
			return getRuleContext(SourceElementContext.class,i);
		}
		public SourceElementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sourceElements; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterSourceElements(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitSourceElements(this);
		}
	}

	public final SourceElementsContext sourceElements() throws RecognitionException {
		SourceElementsContext _localctx = new SourceElementsContext(_ctx, getState());
		enterRule(_localctx, 228, RULE_sourceElements);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1344); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(1343);
					sourceElement();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(1346); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,171,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrayLiteralContext extends ParserRuleContext {
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public ElementListContext elementList() {
			return getRuleContext(ElementListContext.class,0);
		}
		public ArrayLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayLiteral; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArrayLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArrayLiteral(this);
		}
	}

	public final ArrayLiteralContext arrayLiteral() throws RecognitionException {
		ArrayLiteralContext _localctx = new ArrayLiteralContext(_ctx, getState());
		enterRule(_localctx, 230, RULE_arrayLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(1348);
			match(OpenBracket);
			setState(1350);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RegularExpressionLiteral) | (1L << OpenBracket) | (1L << OpenParen) | (1L << OpenBrace) | (1L << Ellipsis) | (1L << PlusPlus) | (1L << MinusMinus) | (1L << Plus) | (1L << Minus) | (1L << BitNot) | (1L << Not) | (1L << LessThan) | (1L << NullLiteral) | (1L << BooleanLiteral) | (1L << DecimalLiteral))) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & ((1L << (HexIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral - 64)) | (1L << (OctalIntegerLiteral2 - 64)) | (1L << (BinaryIntegerLiteral - 64)) | (1L << (Break - 64)) | (1L << (Do - 64)) | (1L << (Instanceof - 64)) | (1L << (Typeof - 64)) | (1L << (Case - 64)) | (1L << (Else - 64)) | (1L << (New - 64)) | (1L << (Var - 64)) | (1L << (Catch - 64)) | (1L << (Finally - 64)) | (1L << (Return - 64)) | (1L << (Void - 64)) | (1L << (Continue - 64)) | (1L << (For - 64)) | (1L << (Switch - 64)) | (1L << (While - 64)) | (1L << (Debugger - 64)) | (1L << (Function - 64)) | (1L << (This - 64)) | (1L << (With - 64)) | (1L << (Default - 64)) | (1L << (If - 64)) | (1L << (Throw - 64)) | (1L << (Delete - 64)) | (1L << (In - 64)) | (1L << (Try - 64)) | (1L << (As - 64)) | (1L << (From - 64)) | (1L << (ReadOnly - 64)) | (1L << (Async - 64)) | (1L << (Class - 64)) | (1L << (Enum - 64)) | (1L << (Extends - 64)) | (1L << (Super - 64)) | (1L << (Const - 64)) | (1L << (Export - 64)) | (1L << (Import - 64)) | (1L << (Await - 64)) | (1L << (Implements - 64)) | (1L << (Let - 64)) | (1L << (Private - 64)) | (1L << (Public - 64)) | (1L << (Interface - 64)) | (1L << (Package - 64)) | (1L << (Protected - 64)) | (1L << (Static - 64)) | (1L << (Yield - 64)))) != 0) || ((((_la - 131)) & ~0x3f) == 0 && ((1L << (_la - 131)) & ((1L << (Identifier - 131)) | (1L << (StringLiteral - 131)) | (1L << (TemplateStringLiteral - 131)))) != 0)) {
				{
				setState(1349);
				elementList();
				}
			}

			setState(1352);
			match(CloseBracket);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ElementListContext extends ParserRuleContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public LastElementContext lastElement() {
			return getRuleContext(LastElementContext.class,0);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public ElementListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elementList; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterElementList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitElementList(this);
		}
	}

	public final ElementListContext elementList() throws RecognitionException {
		ElementListContext _localctx = new ElementListContext(_ctx, getState());
		enterRule(_localctx, 232, RULE_elementList);
		int _la;
		try {
			int _alt;
			setState(1375);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RegularExpressionLiteral:
			case OpenBracket:
			case OpenParen:
			case OpenBrace:
			case PlusPlus:
			case MinusMinus:
			case Plus:
			case Minus:
			case BitNot:
			case Not:
			case LessThan:
			case NullLiteral:
			case BooleanLiteral:
			case DecimalLiteral:
			case HexIntegerLiteral:
			case OctalIntegerLiteral:
			case OctalIntegerLiteral2:
			case BinaryIntegerLiteral:
			case Break:
			case Do:
			case Instanceof:
			case Typeof:
			case Case:
			case Else:
			case New:
			case Var:
			case Catch:
			case Finally:
			case Return:
			case Void:
			case Continue:
			case For:
			case Switch:
			case While:
			case Debugger:
			case Function:
			case This:
			case With:
			case Default:
			case If:
			case Throw:
			case Delete:
			case In:
			case Try:
			case As:
			case From:
			case ReadOnly:
			case Async:
			case Class:
			case Enum:
			case Extends:
			case Super:
			case Const:
			case Export:
			case Import:
			case Await:
			case Implements:
			case Let:
			case Private:
			case Public:
			case Interface:
			case Package:
			case Protected:
			case Static:
			case Yield:
			case Identifier:
			case StringLiteral:
			case TemplateStringLiteral:
				enterOuterAlt(_localctx, 1);
				{
				setState(1354);
				singleExpression(0);
				setState(1363);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,174,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1356); 
						_errHandler.sync(this);
						_la = _input.LA(1);
						do {
							{
							{
							setState(1355);
							match(Comma);
							}
							}
							setState(1358); 
							_errHandler.sync(this);
							_la = _input.LA(1);
						} while ( _la==Comma );
						setState(1360);
						singleExpression(0);
						}
						} 
					}
					setState(1365);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,174,_ctx);
				}
				setState(1372);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Comma) {
					{
					setState(1367); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(1366);
						match(Comma);
						}
						}
						setState(1369); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==Comma );
					setState(1371);
					lastElement();
					}
				}

				}
				break;
			case Ellipsis:
				enterOuterAlt(_localctx, 2);
				{
				setState(1374);
				lastElement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LastElementContext extends ParserRuleContext {
		public TerminalNode Ellipsis() { return getToken(TypeScriptParser.Ellipsis, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public LastElementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lastElement; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterLastElement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitLastElement(this);
		}
	}

	public final LastElementContext lastElement() throws RecognitionException {
		LastElementContext _localctx = new LastElementContext(_ctx, getState());
		enterRule(_localctx, 234, RULE_lastElement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1377);
			match(Ellipsis);
			setState(1380);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,178,_ctx) ) {
			case 1:
				{
				setState(1378);
				match(Identifier);
				}
				break;
			case 2:
				{
				setState(1379);
				singleExpression(0);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ObjectLiteralContext extends ParserRuleContext {
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public List<PropertyAssignmentContext> propertyAssignment() {
			return getRuleContexts(PropertyAssignmentContext.class);
		}
		public PropertyAssignmentContext propertyAssignment(int i) {
			return getRuleContext(PropertyAssignmentContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public ObjectLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_objectLiteral; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterObjectLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitObjectLiteral(this);
		}
	}

	public final ObjectLiteralContext objectLiteral() throws RecognitionException {
		ObjectLiteralContext _localctx = new ObjectLiteralContext(_ctx, getState());
		enterRule(_localctx, 236, RULE_objectLiteral);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1382);
			match(OpenBrace);
			setState(1391);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & ((1L << (OpenBracket - 5)) | (1L << (Ellipsis - 5)) | (1L << (Multiply - 5)) | (1L << (NullLiteral - 5)) | (1L << (BooleanLiteral - 5)) | (1L << (DecimalLiteral - 5)) | (1L << (HexIntegerLiteral - 5)) | (1L << (OctalIntegerLiteral - 5)) | (1L << (OctalIntegerLiteral2 - 5)) | (1L << (BinaryIntegerLiteral - 5)) | (1L << (Break - 5)))) != 0) || ((((_la - 69)) & ~0x3f) == 0 && ((1L << (_la - 69)) & ((1L << (Do - 69)) | (1L << (Instanceof - 69)) | (1L << (Typeof - 69)) | (1L << (Case - 69)) | (1L << (Else - 69)) | (1L << (New - 69)) | (1L << (Var - 69)) | (1L << (Catch - 69)) | (1L << (Finally - 69)) | (1L << (Return - 69)) | (1L << (Void - 69)) | (1L << (Continue - 69)) | (1L << (For - 69)) | (1L << (Switch - 69)) | (1L << (While - 69)) | (1L << (Debugger - 69)) | (1L << (Function - 69)) | (1L << (This - 69)) | (1L << (With - 69)) | (1L << (Default - 69)) | (1L << (If - 69)) | (1L << (Throw - 69)) | (1L << (Delete - 69)) | (1L << (In - 69)) | (1L << (Try - 69)) | (1L << (As - 69)) | (1L << (From - 69)) | (1L << (ReadOnly - 69)) | (1L << (Async - 69)) | (1L << (Class - 69)) | (1L << (Enum - 69)) | (1L << (Extends - 69)) | (1L << (Super - 69)) | (1L << (Const - 69)) | (1L << (Export - 69)) | (1L << (Import - 69)) | (1L << (Await - 69)) | (1L << (Implements - 69)) | (1L << (Let - 69)) | (1L << (Private - 69)) | (1L << (Public - 69)) | (1L << (Interface - 69)) | (1L << (Package - 69)) | (1L << (Protected - 69)) | (1L << (Static - 69)) | (1L << (Yield - 69)) | (1L << (Identifier - 69)) | (1L << (StringLiteral - 69)))) != 0)) {
				{
				setState(1383);
				propertyAssignment();
				setState(1388);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,179,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1384);
						match(Comma);
						setState(1385);
						propertyAssignment();
						}
						} 
					}
					setState(1390);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,179,_ctx);
				}
				}
			}

			setState(1394);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Comma) {
				{
				setState(1393);
				match(Comma);
				}
			}

			setState(1396);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PropertyAssignmentContext extends ParserRuleContext {
		public PropertyAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertyAssignment; }
	 
		public PropertyAssignmentContext() { }
		public void copyFrom(PropertyAssignmentContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class PropertyExpressionAssignmentContext extends PropertyAssignmentContext {
		public PropertyNameContext propertyName() {
			return getRuleContext(PropertyNameContext.class,0);
		}
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public TerminalNode Assign() { return getToken(TypeScriptParser.Assign, 0); }
		public PropertyExpressionAssignmentContext(PropertyAssignmentContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPropertyExpressionAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPropertyExpressionAssignment(this);
		}
	}
	public static class ComputedPropertyExpressionAssignmentContext extends PropertyAssignmentContext {
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public ComputedPropertyExpressionAssignmentContext(PropertyAssignmentContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterComputedPropertyExpressionAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitComputedPropertyExpressionAssignment(this);
		}
	}
	public static class PropertyShorthandContext extends PropertyAssignmentContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public PropertyShorthandContext(PropertyAssignmentContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPropertyShorthand(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPropertyShorthand(this);
		}
	}
	public static class PropertySetterContext extends PropertyAssignmentContext {
		public SetAccessorContext setAccessor() {
			return getRuleContext(SetAccessorContext.class,0);
		}
		public PropertySetterContext(PropertyAssignmentContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPropertySetter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPropertySetter(this);
		}
	}
	public static class PropertyGetterContext extends PropertyAssignmentContext {
		public GetAccessorContext getAccessor() {
			return getRuleContext(GetAccessorContext.class,0);
		}
		public PropertyGetterContext(PropertyAssignmentContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPropertyGetter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPropertyGetter(this);
		}
	}
	public static class FunctionPropertyContext extends PropertyAssignmentContext {
		public PropertyNameContext propertyName() {
			return getRuleContext(PropertyNameContext.class,0);
		}
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public TerminalNode Async() { return getToken(TypeScriptParser.Async, 0); }
		public TerminalNode Multiply() { return getToken(TypeScriptParser.Multiply, 0); }
		public FormalParameterListContext formalParameterList() {
			return getRuleContext(FormalParameterListContext.class,0);
		}
		public FunctionPropertyContext(PropertyAssignmentContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFunctionProperty(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFunctionProperty(this);
		}
	}
	public static class RestParameterInObjectContext extends PropertyAssignmentContext {
		public RestParameterContext restParameter() {
			return getRuleContext(RestParameterContext.class,0);
		}
		public RestParameterInObjectContext(PropertyAssignmentContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterRestParameterInObject(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitRestParameterInObject(this);
		}
	}
	public static class MethodPropertyContext extends PropertyAssignmentContext {
		public GeneratorMethodContext generatorMethod() {
			return getRuleContext(GeneratorMethodContext.class,0);
		}
		public MethodPropertyContext(PropertyAssignmentContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterMethodProperty(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitMethodProperty(this);
		}
	}

	public final PropertyAssignmentContext propertyAssignment() throws RecognitionException {
		PropertyAssignmentContext _localctx = new PropertyAssignmentContext(_ctx, getState());
		enterRule(_localctx, 238, RULE_propertyAssignment);
		int _la;
		try {
			setState(1429);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,185,_ctx) ) {
			case 1:
				_localctx = new PropertyExpressionAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(1398);
				propertyName();
				setState(1399);
				_la = _input.LA(1);
				if ( !(_la==Assign || _la==Colon) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(1400);
				singleExpression(0);
				}
				break;
			case 2:
				_localctx = new ComputedPropertyExpressionAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(1402);
				match(OpenBracket);
				setState(1403);
				singleExpression(0);
				setState(1404);
				match(CloseBracket);
				setState(1405);
				match(Colon);
				setState(1406);
				singleExpression(0);
				}
				break;
			case 3:
				_localctx = new FunctionPropertyContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(1409);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,182,_ctx) ) {
				case 1:
					{
					setState(1408);
					match(Async);
					}
					break;
				}
				setState(1412);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Multiply) {
					{
					setState(1411);
					match(Multiply);
					}
				}

				setState(1414);
				propertyName();
				setState(1415);
				match(OpenParen);
				setState(1417);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis))) != 0) || ((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)) | (1L << (Identifier - 108)))) != 0)) {
					{
					setState(1416);
					formalParameterList();
					}
				}

				setState(1419);
				match(CloseParen);
				setState(1420);
				match(OpenBrace);
				setState(1421);
				functionBody();
				setState(1422);
				match(CloseBrace);
				}
				break;
			case 4:
				_localctx = new PropertyGetterContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(1424);
				getAccessor();
				}
				break;
			case 5:
				_localctx = new PropertySetterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(1425);
				setAccessor();
				}
				break;
			case 6:
				_localctx = new MethodPropertyContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(1426);
				generatorMethod();
				}
				break;
			case 7:
				_localctx = new PropertyShorthandContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(1427);
				match(Identifier);
				}
				break;
			case 8:
				_localctx = new RestParameterInObjectContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(1428);
				restParameter();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GetAccessorContext extends ParserRuleContext {
		public GetterContext getter() {
			return getRuleContext(GetterContext.class,0);
		}
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public GetAccessorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_getAccessor; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGetAccessor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGetAccessor(this);
		}
	}

	public final GetAccessorContext getAccessor() throws RecognitionException {
		GetAccessorContext _localctx = new GetAccessorContext(_ctx, getState());
		enterRule(_localctx, 240, RULE_getAccessor);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1431);
			getter();
			setState(1432);
			match(OpenParen);
			setState(1433);
			match(CloseParen);
			setState(1435);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(1434);
				typeAnnotation();
				}
			}

			setState(1437);
			match(OpenBrace);
			setState(1438);
			functionBody();
			setState(1439);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SetAccessorContext extends ParserRuleContext {
		public SetterContext setter() {
			return getRuleContext(SetterContext.class,0);
		}
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public BindingPatternContext bindingPattern() {
			return getRuleContext(BindingPatternContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public SetAccessorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_setAccessor; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterSetAccessor(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitSetAccessor(this);
		}
	}

	public final SetAccessorContext setAccessor() throws RecognitionException {
		SetAccessorContext _localctx = new SetAccessorContext(_ctx, getState());
		enterRule(_localctx, 242, RULE_setAccessor);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1441);
			setter();
			setState(1442);
			match(OpenParen);
			setState(1445);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				{
				setState(1443);
				match(Identifier);
				}
				break;
			case OpenBracket:
			case OpenBrace:
				{
				setState(1444);
				bindingPattern();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(1448);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(1447);
				typeAnnotation();
				}
			}

			setState(1450);
			match(CloseParen);
			setState(1451);
			match(OpenBrace);
			setState(1452);
			functionBody();
			setState(1453);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PropertyNameContext extends ParserRuleContext {
		public IdentifierNameContext identifierName() {
			return getRuleContext(IdentifierNameContext.class,0);
		}
		public TerminalNode StringLiteral() { return getToken(TypeScriptParser.StringLiteral, 0); }
		public NumericLiteralContext numericLiteral() {
			return getRuleContext(NumericLiteralContext.class,0);
		}
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public PropertyNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertyName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPropertyName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPropertyName(this);
		}
	}

	public final PropertyNameContext propertyName() throws RecognitionException {
		PropertyNameContext _localctx = new PropertyNameContext(_ctx, getState());
		enterRule(_localctx, 244, RULE_propertyName);
		try {
			setState(1462);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NullLiteral:
			case BooleanLiteral:
			case Break:
			case Do:
			case Instanceof:
			case Typeof:
			case Case:
			case Else:
			case New:
			case Var:
			case Catch:
			case Finally:
			case Return:
			case Void:
			case Continue:
			case For:
			case Switch:
			case While:
			case Debugger:
			case Function:
			case This:
			case With:
			case Default:
			case If:
			case Throw:
			case Delete:
			case In:
			case Try:
			case As:
			case From:
			case ReadOnly:
			case Async:
			case Class:
			case Enum:
			case Extends:
			case Super:
			case Const:
			case Export:
			case Import:
			case Await:
			case Implements:
			case Let:
			case Private:
			case Public:
			case Interface:
			case Package:
			case Protected:
			case Static:
			case Yield:
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(1455);
				identifierName();
				}
				break;
			case StringLiteral:
				enterOuterAlt(_localctx, 2);
				{
				setState(1456);
				match(StringLiteral);
				}
				break;
			case DecimalLiteral:
			case HexIntegerLiteral:
			case OctalIntegerLiteral:
			case OctalIntegerLiteral2:
			case BinaryIntegerLiteral:
				enterOuterAlt(_localctx, 3);
				{
				setState(1457);
				numericLiteral();
				}
				break;
			case OpenBracket:
				enterOuterAlt(_localctx, 4);
				{
				setState(1458);
				match(OpenBracket);
				setState(1459);
				singleExpression(0);
				setState(1460);
				match(CloseBracket);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentsContext extends ParserRuleContext {
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public LastArgumentContext lastArgument() {
			return getRuleContext(LastArgumentContext.class,0);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArguments(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArguments(this);
		}
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 246, RULE_arguments);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1464);
			match(OpenParen);
			setState(1478);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RegularExpressionLiteral:
			case OpenBracket:
			case OpenParen:
			case OpenBrace:
			case PlusPlus:
			case MinusMinus:
			case Plus:
			case Minus:
			case BitNot:
			case Not:
			case LessThan:
			case NullLiteral:
			case BooleanLiteral:
			case DecimalLiteral:
			case HexIntegerLiteral:
			case OctalIntegerLiteral:
			case OctalIntegerLiteral2:
			case BinaryIntegerLiteral:
			case Break:
			case Do:
			case Instanceof:
			case Typeof:
			case Case:
			case Else:
			case New:
			case Var:
			case Catch:
			case Finally:
			case Return:
			case Void:
			case Continue:
			case For:
			case Switch:
			case While:
			case Debugger:
			case Function:
			case This:
			case With:
			case Default:
			case If:
			case Throw:
			case Delete:
			case In:
			case Try:
			case As:
			case From:
			case ReadOnly:
			case Async:
			case Class:
			case Enum:
			case Extends:
			case Super:
			case Const:
			case Export:
			case Import:
			case Await:
			case Implements:
			case Let:
			case Private:
			case Public:
			case Interface:
			case Package:
			case Protected:
			case Static:
			case Yield:
			case Identifier:
			case StringLiteral:
			case TemplateStringLiteral:
				{
				setState(1465);
				singleExpression(0);
				setState(1470);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,190,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(1466);
						match(Comma);
						setState(1467);
						singleExpression(0);
						}
						} 
					}
					setState(1472);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,190,_ctx);
				}
				setState(1475);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Comma) {
					{
					setState(1473);
					match(Comma);
					setState(1474);
					lastArgument();
					}
				}

				}
				break;
			case Ellipsis:
				{
				setState(1477);
				lastArgument();
				}
				break;
			case CloseParen:
				break;
			default:
				break;
			}
			setState(1480);
			match(CloseParen);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LastArgumentContext extends ParserRuleContext {
		public TerminalNode Ellipsis() { return getToken(TypeScriptParser.Ellipsis, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public LastArgumentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lastArgument; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterLastArgument(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitLastArgument(this);
		}
	}

	public final LastArgumentContext lastArgument() throws RecognitionException {
		LastArgumentContext _localctx = new LastArgumentContext(_ctx, getState());
		enterRule(_localctx, 248, RULE_lastArgument);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1482);
			match(Ellipsis);
			setState(1483);
			match(Identifier);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionSequenceContext extends ParserRuleContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public List<TerminalNode> Comma() { return getTokens(TypeScriptParser.Comma); }
		public TerminalNode Comma(int i) {
			return getToken(TypeScriptParser.Comma, i);
		}
		public ExpressionSequenceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionSequence; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterExpressionSequence(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitExpressionSequence(this);
		}
	}

	public final ExpressionSequenceContext expressionSequence() throws RecognitionException {
		ExpressionSequenceContext _localctx = new ExpressionSequenceContext(_ctx, getState());
		enterRule(_localctx, 250, RULE_expressionSequence);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1485);
			singleExpression(0);
			setState(1490);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,193,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(1486);
					match(Comma);
					setState(1487);
					singleExpression(0);
					}
					} 
				}
				setState(1492);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,193,_ctx);
			}
			setState(1494);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,194,_ctx) ) {
			case 1:
				{
				setState(1493);
				match(Comma);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionExpressionDeclarationContext extends ParserRuleContext {
		public TerminalNode Function() { return getToken(TypeScriptParser.Function, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public TerminalNode Async() { return getToken(TypeScriptParser.Async, 0); }
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public FormalParameterListContext formalParameterList() {
			return getRuleContext(FormalParameterListContext.class,0);
		}
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public FunctionExpressionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionExpressionDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFunctionExpressionDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFunctionExpressionDeclaration(this);
		}
	}

	public final FunctionExpressionDeclarationContext functionExpressionDeclaration() throws RecognitionException {
		FunctionExpressionDeclarationContext _localctx = new FunctionExpressionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 252, RULE_functionExpressionDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1497);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Async) {
				{
				setState(1496);
				match(Async);
				}
			}

			setState(1499);
			match(Function);
			setState(1501);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(1500);
				match(Identifier);
				}
			}

			setState(1503);
			match(OpenParen);
			setState(1505);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis))) != 0) || ((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)) | (1L << (Identifier - 108)))) != 0)) {
				{
				setState(1504);
				formalParameterList();
				}
			}

			setState(1507);
			match(CloseParen);
			setState(1509);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(1508);
				typeAnnotation();
				}
			}

			setState(1511);
			match(OpenBrace);
			setState(1512);
			functionBody();
			setState(1513);
			match(CloseBrace);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SingleExpressionContext extends ParserRuleContext {
		public SingleExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_singleExpression; }
	 
		public SingleExpressionContext() { }
		public void copyFrom(SingleExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class TemplateStringExpressionContext extends SingleExpressionContext {
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode TemplateStringLiteral() { return getToken(TypeScriptParser.TemplateStringLiteral, 0); }
		public TemplateStringExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTemplateStringExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTemplateStringExpression(this);
		}
	}
	public static class GeneratorsExpressionContext extends SingleExpressionContext {
		public GeneratorBlockContext generatorBlock() {
			return getRuleContext(GeneratorBlockContext.class,0);
		}
		public GeneratorsExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGeneratorsExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGeneratorsExpression(this);
		}
	}
	public static class PowerExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode Power() { return getToken(TypeScriptParser.Power, 0); }
		public PowerExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPowerExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPowerExpression(this);
		}
	}
	public static class InExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode In() { return getToken(TypeScriptParser.In, 0); }
		public InExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterInExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitInExpression(this);
		}
	}
	public static class GenericTypesContext extends SingleExpressionContext {
		public TypeArgumentsContext typeArguments() {
			return getRuleContext(TypeArgumentsContext.class,0);
		}
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public GenericTypesContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGenericTypes(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGenericTypes(this);
		}
	}
	public static class ArgumentsExpressionContext extends SingleExpressionContext {
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public ArgumentsExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArgumentsExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArgumentsExpression(this);
		}
	}
	public static class ThisExpressionContext extends SingleExpressionContext {
		public TerminalNode This() { return getToken(TypeScriptParser.This, 0); }
		public ThisExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterThisExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitThisExpression(this);
		}
	}
	public static class TypeofExpressionContext extends SingleExpressionContext {
		public TerminalNode Typeof() { return getToken(TypeScriptParser.Typeof, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TypeofExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTypeofExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTypeofExpression(this);
		}
	}
	public static class GeneratorsFunctionExpressionContext extends SingleExpressionContext {
		public GeneratorFunctionDeclarationContext generatorFunctionDeclaration() {
			return getRuleContext(GeneratorFunctionDeclarationContext.class,0);
		}
		public GeneratorsFunctionExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGeneratorsFunctionExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGeneratorsFunctionExpression(this);
		}
	}
	public static class EqualityExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode Equals_() { return getToken(TypeScriptParser.Equals_, 0); }
		public TerminalNode NotEquals() { return getToken(TypeScriptParser.NotEquals, 0); }
		public TerminalNode IdentityEquals() { return getToken(TypeScriptParser.IdentityEquals, 0); }
		public TerminalNode IdentityNotEquals() { return getToken(TypeScriptParser.IdentityNotEquals, 0); }
		public EqualityExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterEqualityExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitEqualityExpression(this);
		}
	}
	public static class BitXOrExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode BitXOr() { return getToken(TypeScriptParser.BitXOr, 0); }
		public BitXOrExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterBitXOrExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitBitXOrExpression(this);
		}
	}
	public static class MultiplicativeExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode Multiply() { return getToken(TypeScriptParser.Multiply, 0); }
		public TerminalNode Divide() { return getToken(TypeScriptParser.Divide, 0); }
		public TerminalNode Modulus() { return getToken(TypeScriptParser.Modulus, 0); }
		public MultiplicativeExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterMultiplicativeExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitMultiplicativeExpression(this);
		}
	}
	public static class BitShiftExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode LeftShiftArithmetic() { return getToken(TypeScriptParser.LeftShiftArithmetic, 0); }
		public TerminalNode RightShiftArithmetic() { return getToken(TypeScriptParser.RightShiftArithmetic, 0); }
		public TerminalNode RightShiftLogical() { return getToken(TypeScriptParser.RightShiftLogical, 0); }
		public BitShiftExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterBitShiftExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitBitShiftExpression(this);
		}
	}
	public static class AdditiveExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode Plus() { return getToken(TypeScriptParser.Plus, 0); }
		public TerminalNode Minus() { return getToken(TypeScriptParser.Minus, 0); }
		public AdditiveExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterAdditiveExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitAdditiveExpression(this);
		}
	}
	public static class RelationalExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode LessThan() { return getToken(TypeScriptParser.LessThan, 0); }
		public TerminalNode MoreThan() { return getToken(TypeScriptParser.MoreThan, 0); }
		public TerminalNode LessThanEquals() { return getToken(TypeScriptParser.LessThanEquals, 0); }
		public TerminalNode GreaterThanEquals() { return getToken(TypeScriptParser.GreaterThanEquals, 0); }
		public RelationalExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterRelationalExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitRelationalExpression(this);
		}
	}
	public static class BitNotExpressionContext extends SingleExpressionContext {
		public TerminalNode BitNot() { return getToken(TypeScriptParser.BitNot, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public BitNotExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterBitNotExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitBitNotExpression(this);
		}
	}
	public static class NewExpressionContext extends SingleExpressionContext {
		public TerminalNode New() { return getToken(TypeScriptParser.New, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TypeArgumentsContext typeArguments() {
			return getRuleContext(TypeArgumentsContext.class,0);
		}
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public NewExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterNewExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitNewExpression(this);
		}
	}
	public static class LiteralExpressionContext extends SingleExpressionContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public LiteralExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterLiteralExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitLiteralExpression(this);
		}
	}
	public static class ArrayLiteralExpressionContext extends SingleExpressionContext {
		public ArrayLiteralContext arrayLiteral() {
			return getRuleContext(ArrayLiteralContext.class,0);
		}
		public ArrayLiteralExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArrayLiteralExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArrayLiteralExpression(this);
		}
	}
	public static class MemberDotExpressionContext extends SingleExpressionContext {
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode Dot() { return getToken(TypeScriptParser.Dot, 0); }
		public IdentifierNameContext identifierName() {
			return getRuleContext(IdentifierNameContext.class,0);
		}
		public TerminalNode QuestionMark() { return getToken(TypeScriptParser.QuestionMark, 0); }
		public TerminalNode Hashtag() { return getToken(TypeScriptParser.Hashtag, 0); }
		public MemberDotExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterMemberDotExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitMemberDotExpression(this);
		}
	}
	public static class MemberIndexExpressionContext extends SingleExpressionContext {
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode OpenBracket() { return getToken(TypeScriptParser.OpenBracket, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseBracket() { return getToken(TypeScriptParser.CloseBracket, 0); }
		public MemberIndexExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterMemberIndexExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitMemberIndexExpression(this);
		}
	}
	public static class BitAndExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode BitAnd() { return getToken(TypeScriptParser.BitAnd, 0); }
		public BitAndExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterBitAndExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitBitAndExpression(this);
		}
	}
	public static class BitOrExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode BitOr() { return getToken(TypeScriptParser.BitOr, 0); }
		public BitOrExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterBitOrExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitBitOrExpression(this);
		}
	}
	public static class AssignmentOperatorExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public AssignmentOperatorContext assignmentOperator() {
			return getRuleContext(AssignmentOperatorContext.class,0);
		}
		public AssignmentOperatorExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterAssignmentOperatorExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitAssignmentOperatorExpression(this);
		}
	}
	public static class VoidExpressionContext extends SingleExpressionContext {
		public TerminalNode Void() { return getToken(TypeScriptParser.Void, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public VoidExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterVoidExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitVoidExpression(this);
		}
	}
	public static class TernaryExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode QuestionMark() { return getToken(TypeScriptParser.QuestionMark, 0); }
		public TerminalNode Colon() { return getToken(TypeScriptParser.Colon, 0); }
		public TernaryExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterTernaryExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitTernaryExpression(this);
		}
	}
	public static class LogicalAndExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode And() { return getToken(TypeScriptParser.And, 0); }
		public LogicalAndExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterLogicalAndExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitLogicalAndExpression(this);
		}
	}
	public static class PreIncrementExpressionContext extends SingleExpressionContext {
		public TerminalNode PlusPlus() { return getToken(TypeScriptParser.PlusPlus, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public PreIncrementExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPreIncrementExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPreIncrementExpression(this);
		}
	}
	public static class ObjectLiteralExpressionContext extends SingleExpressionContext {
		public ObjectLiteralContext objectLiteral() {
			return getRuleContext(ObjectLiteralContext.class,0);
		}
		public ObjectLiteralExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterObjectLiteralExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitObjectLiteralExpression(this);
		}
	}
	public static class LogicalOrExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode Or() { return getToken(TypeScriptParser.Or, 0); }
		public LogicalOrExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterLogicalOrExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitLogicalOrExpression(this);
		}
	}
	public static class NotExpressionContext extends SingleExpressionContext {
		public TerminalNode Not() { return getToken(TypeScriptParser.Not, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public NotExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterNotExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitNotExpression(this);
		}
	}
	public static class PreDecreaseExpressionContext extends SingleExpressionContext {
		public TerminalNode MinusMinus() { return getToken(TypeScriptParser.MinusMinus, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public PreDecreaseExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPreDecreaseExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPreDecreaseExpression(this);
		}
	}
	public static class AwaitExpressionContext extends SingleExpressionContext {
		public TerminalNode Await() { return getToken(TypeScriptParser.Await, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public AwaitExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterAwaitExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitAwaitExpression(this);
		}
	}
	public static class FunctionExpressionContext extends SingleExpressionContext {
		public FunctionExpressionDeclarationContext functionExpressionDeclaration() {
			return getRuleContext(FunctionExpressionDeclarationContext.class,0);
		}
		public FunctionExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterFunctionExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitFunctionExpression(this);
		}
	}
	public static class UnaryMinusExpressionContext extends SingleExpressionContext {
		public TerminalNode Minus() { return getToken(TypeScriptParser.Minus, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public UnaryMinusExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterUnaryMinusExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitUnaryMinusExpression(this);
		}
	}
	public static class AssignmentExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode Assign() { return getToken(TypeScriptParser.Assign, 0); }
		public AssignmentExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterAssignmentExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitAssignmentExpression(this);
		}
	}
	public static class PostDecreaseExpressionContext extends SingleExpressionContext {
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode MinusMinus() { return getToken(TypeScriptParser.MinusMinus, 0); }
		public PostDecreaseExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPostDecreaseExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPostDecreaseExpression(this);
		}
	}
	public static class InstanceofExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode Instanceof() { return getToken(TypeScriptParser.Instanceof, 0); }
		public InstanceofExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterInstanceofExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitInstanceofExpression(this);
		}
	}
	public static class UnaryPlusExpressionContext extends SingleExpressionContext {
		public TerminalNode Plus() { return getToken(TypeScriptParser.Plus, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public UnaryPlusExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterUnaryPlusExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitUnaryPlusExpression(this);
		}
	}
	public static class DeleteExpressionContext extends SingleExpressionContext {
		public TerminalNode Delete() { return getToken(TypeScriptParser.Delete, 0); }
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public DeleteExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterDeleteExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitDeleteExpression(this);
		}
	}
	public static class ArrowFunctionExpressionContext extends SingleExpressionContext {
		public ArrowFunctionDeclarationContext arrowFunctionDeclaration() {
			return getRuleContext(ArrowFunctionDeclarationContext.class,0);
		}
		public ArrowFunctionExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArrowFunctionExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArrowFunctionExpression(this);
		}
	}
	public static class IteratorsExpressionContext extends SingleExpressionContext {
		public IteratorBlockContext iteratorBlock() {
			return getRuleContext(IteratorBlockContext.class,0);
		}
		public IteratorsExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIteratorsExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIteratorsExpression(this);
		}
	}
	public static class SuperExpressionContext extends SingleExpressionContext {
		public TerminalNode Super() { return getToken(TypeScriptParser.Super, 0); }
		public SuperExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterSuperExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitSuperExpression(this);
		}
	}
	public static class ParenthesizedExpressionContext extends SingleExpressionContext {
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public ExpressionSequenceContext expressionSequence() {
			return getRuleContext(ExpressionSequenceContext.class,0);
		}
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public ParenthesizedExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterParenthesizedExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitParenthesizedExpression(this);
		}
	}
	public static class PostIncrementExpressionContext extends SingleExpressionContext {
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode PlusPlus() { return getToken(TypeScriptParser.PlusPlus, 0); }
		public PostIncrementExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterPostIncrementExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitPostIncrementExpression(this);
		}
	}
	public static class YieldExpressionContext extends SingleExpressionContext {
		public YieldStatementContext yieldStatement() {
			return getRuleContext(YieldStatementContext.class,0);
		}
		public YieldExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterYieldExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitYieldExpression(this);
		}
	}
	public static class ClassExpressionContext extends SingleExpressionContext {
		public TerminalNode Class() { return getToken(TypeScriptParser.Class, 0); }
		public ClassTailContext classTail() {
			return getRuleContext(ClassTailContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public ClassExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterClassExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitClassExpression(this);
		}
	}
	public static class IdentifierExpressionContext extends SingleExpressionContext {
		public IdentifierNameContext identifierName() {
			return getRuleContext(IdentifierNameContext.class,0);
		}
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public IdentifierExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIdentifierExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIdentifierExpression(this);
		}
	}
	public static class CoalesceExpressionContext extends SingleExpressionContext {
		public List<SingleExpressionContext> singleExpression() {
			return getRuleContexts(SingleExpressionContext.class);
		}
		public SingleExpressionContext singleExpression(int i) {
			return getRuleContext(SingleExpressionContext.class,i);
		}
		public TerminalNode NullCoalesce() { return getToken(TypeScriptParser.NullCoalesce, 0); }
		public CoalesceExpressionContext(SingleExpressionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterCoalesceExpression(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitCoalesceExpression(this);
		}
	}

	public final SingleExpressionContext singleExpression() throws RecognitionException {
		return singleExpression(0);
	}

	private SingleExpressionContext singleExpression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		SingleExpressionContext _localctx = new SingleExpressionContext(_ctx, _parentState);
		SingleExpressionContext _prevctx = _localctx;
		int _startState = 254;
		enterRecursionRule(_localctx, 254, RULE_singleExpression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(1572);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,204,_ctx) ) {
			case 1:
				{
				_localctx = new FunctionExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(1516);
				functionExpressionDeclaration();
				}
				break;
			case 2:
				{
				_localctx = new ArrowFunctionExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1517);
				arrowFunctionDeclaration();
				}
				break;
			case 3:
				{
				_localctx = new ClassExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1518);
				match(Class);
				setState(1520);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Identifier) {
					{
					setState(1519);
					match(Identifier);
					}
				}

				setState(1522);
				classTail();
				}
				break;
			case 4:
				{
				_localctx = new NewExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1523);
				match(New);
				setState(1524);
				singleExpression(0);
				setState(1526);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,200,_ctx) ) {
				case 1:
					{
					setState(1525);
					typeArguments();
					}
					break;
				}
				setState(1529);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,201,_ctx) ) {
				case 1:
					{
					setState(1528);
					arguments();
					}
					break;
				}
				}
				break;
			case 5:
				{
				_localctx = new DeleteExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1531);
				match(Delete);
				setState(1532);
				singleExpression(40);
				}
				break;
			case 6:
				{
				_localctx = new VoidExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1533);
				match(Void);
				setState(1534);
				singleExpression(39);
				}
				break;
			case 7:
				{
				_localctx = new TypeofExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1535);
				match(Typeof);
				setState(1536);
				singleExpression(38);
				}
				break;
			case 8:
				{
				_localctx = new PreIncrementExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1537);
				match(PlusPlus);
				setState(1538);
				singleExpression(37);
				}
				break;
			case 9:
				{
				_localctx = new PreDecreaseExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1539);
				match(MinusMinus);
				setState(1540);
				singleExpression(36);
				}
				break;
			case 10:
				{
				_localctx = new UnaryPlusExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1541);
				match(Plus);
				setState(1542);
				singleExpression(35);
				}
				break;
			case 11:
				{
				_localctx = new UnaryMinusExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1543);
				match(Minus);
				setState(1544);
				singleExpression(34);
				}
				break;
			case 12:
				{
				_localctx = new BitNotExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1545);
				match(BitNot);
				setState(1546);
				singleExpression(33);
				}
				break;
			case 13:
				{
				_localctx = new NotExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1547);
				match(Not);
				setState(1548);
				singleExpression(32);
				}
				break;
			case 14:
				{
				_localctx = new AwaitExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1549);
				match(Await);
				setState(1550);
				singleExpression(31);
				}
				break;
			case 15:
				{
				_localctx = new IteratorsExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1551);
				iteratorBlock();
				}
				break;
			case 16:
				{
				_localctx = new GeneratorsExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1552);
				generatorBlock();
				}
				break;
			case 17:
				{
				_localctx = new GeneratorsFunctionExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1553);
				generatorFunctionDeclaration();
				}
				break;
			case 18:
				{
				_localctx = new YieldExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1554);
				yieldStatement();
				}
				break;
			case 19:
				{
				_localctx = new ThisExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1555);
				match(This);
				}
				break;
			case 20:
				{
				_localctx = new IdentifierExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1556);
				identifierName();
				setState(1558);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,202,_ctx) ) {
				case 1:
					{
					setState(1557);
					singleExpression(0);
					}
					break;
				}
				}
				break;
			case 21:
				{
				_localctx = new SuperExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1560);
				match(Super);
				}
				break;
			case 22:
				{
				_localctx = new LiteralExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1561);
				literal();
				}
				break;
			case 23:
				{
				_localctx = new ArrayLiteralExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1562);
				arrayLiteral();
				}
				break;
			case 24:
				{
				_localctx = new ObjectLiteralExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1563);
				objectLiteral();
				}
				break;
			case 25:
				{
				_localctx = new ParenthesizedExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1564);
				match(OpenParen);
				setState(1565);
				expressionSequence();
				setState(1566);
				match(CloseParen);
				}
				break;
			case 26:
				{
				_localctx = new GenericTypesContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(1568);
				typeArguments();
				setState(1570);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,203,_ctx) ) {
				case 1:
					{
					setState(1569);
					expressionSequence();
					}
					break;
				}
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(1655);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,208,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(1653);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,207,_ctx) ) {
					case 1:
						{
						_localctx = new PowerExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1574);
						if (!(precpred(_ctx, 30))) throw new FailedPredicateException(this, "precpred(_ctx, 30)");
						setState(1575);
						match(Power);
						setState(1576);
						singleExpression(30);
						}
						break;
					case 2:
						{
						_localctx = new MultiplicativeExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1577);
						if (!(precpred(_ctx, 29))) throw new FailedPredicateException(this, "precpred(_ctx, 29)");
						setState(1578);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Multiply) | (1L << Divide) | (1L << Modulus))) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(1579);
						singleExpression(30);
						}
						break;
					case 3:
						{
						_localctx = new AdditiveExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1580);
						if (!(precpred(_ctx, 28))) throw new FailedPredicateException(this, "precpred(_ctx, 28)");
						setState(1581);
						_la = _input.LA(1);
						if ( !(_la==Plus || _la==Minus) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(1582);
						singleExpression(29);
						}
						break;
					case 4:
						{
						_localctx = new BitShiftExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1583);
						if (!(precpred(_ctx, 27))) throw new FailedPredicateException(this, "precpred(_ctx, 27)");
						setState(1584);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RightShiftArithmetic) | (1L << LeftShiftArithmetic) | (1L << RightShiftLogical))) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(1585);
						singleExpression(28);
						}
						break;
					case 5:
						{
						_localctx = new RelationalExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1586);
						if (!(precpred(_ctx, 26))) throw new FailedPredicateException(this, "precpred(_ctx, 26)");
						setState(1587);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LessThan) | (1L << MoreThan) | (1L << LessThanEquals) | (1L << GreaterThanEquals))) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(1588);
						singleExpression(27);
						}
						break;
					case 6:
						{
						_localctx = new InstanceofExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1589);
						if (!(precpred(_ctx, 25))) throw new FailedPredicateException(this, "precpred(_ctx, 25)");
						setState(1590);
						match(Instanceof);
						setState(1591);
						singleExpression(26);
						}
						break;
					case 7:
						{
						_localctx = new InExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1592);
						if (!(precpred(_ctx, 24))) throw new FailedPredicateException(this, "precpred(_ctx, 24)");
						setState(1593);
						match(In);
						setState(1594);
						singleExpression(25);
						}
						break;
					case 8:
						{
						_localctx = new EqualityExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1595);
						if (!(precpred(_ctx, 23))) throw new FailedPredicateException(this, "precpred(_ctx, 23)");
						setState(1596);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << Equals_) | (1L << NotEquals) | (1L << IdentityEquals) | (1L << IdentityNotEquals))) != 0)) ) {
						_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(1597);
						singleExpression(24);
						}
						break;
					case 9:
						{
						_localctx = new BitAndExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1598);
						if (!(precpred(_ctx, 22))) throw new FailedPredicateException(this, "precpred(_ctx, 22)");
						setState(1599);
						match(BitAnd);
						setState(1600);
						singleExpression(23);
						}
						break;
					case 10:
						{
						_localctx = new BitXOrExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1601);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(1602);
						match(BitXOr);
						setState(1603);
						singleExpression(22);
						}
						break;
					case 11:
						{
						_localctx = new BitOrExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1604);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(1605);
						match(BitOr);
						setState(1606);
						singleExpression(21);
						}
						break;
					case 12:
						{
						_localctx = new LogicalAndExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1607);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(1608);
						match(And);
						setState(1609);
						singleExpression(20);
						}
						break;
					case 13:
						{
						_localctx = new LogicalOrExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1610);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(1611);
						match(Or);
						setState(1612);
						singleExpression(19);
						}
						break;
					case 14:
						{
						_localctx = new TernaryExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1613);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(1614);
						match(QuestionMark);
						setState(1615);
						singleExpression(0);
						setState(1616);
						match(Colon);
						setState(1617);
						singleExpression(18);
						}
						break;
					case 15:
						{
						_localctx = new CoalesceExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1619);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(1620);
						match(NullCoalesce);
						setState(1621);
						singleExpression(17);
						}
						break;
					case 16:
						{
						_localctx = new AssignmentExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1622);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(1623);
						match(Assign);
						setState(1624);
						singleExpression(15);
						}
						break;
					case 17:
						{
						_localctx = new AssignmentOperatorExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1625);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(1626);
						assignmentOperator();
						setState(1627);
						singleExpression(14);
						}
						break;
					case 18:
						{
						_localctx = new MemberIndexExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1629);
						if (!(precpred(_ctx, 46))) throw new FailedPredicateException(this, "precpred(_ctx, 46)");
						setState(1630);
						match(OpenBracket);
						setState(1631);
						expressionSequence();
						setState(1632);
						match(CloseBracket);
						}
						break;
					case 19:
						{
						_localctx = new MemberDotExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1634);
						if (!(precpred(_ctx, 45))) throw new FailedPredicateException(this, "precpred(_ctx, 45)");
						setState(1636);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==QuestionMark) {
							{
							setState(1635);
							match(QuestionMark);
							}
						}

						setState(1638);
						match(Dot);
						setState(1640);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if (_la==Hashtag) {
							{
							setState(1639);
							match(Hashtag);
							}
						}

						setState(1642);
						identifierName();
						}
						break;
					case 20:
						{
						_localctx = new ArgumentsExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1643);
						if (!(precpred(_ctx, 44))) throw new FailedPredicateException(this, "precpred(_ctx, 44)");
						setState(1644);
						arguments();
						}
						break;
					case 21:
						{
						_localctx = new PostIncrementExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1645);
						if (!(precpred(_ctx, 42))) throw new FailedPredicateException(this, "precpred(_ctx, 42)");
						setState(1646);
						if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
						setState(1647);
						match(PlusPlus);
						}
						break;
					case 22:
						{
						_localctx = new PostDecreaseExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1648);
						if (!(precpred(_ctx, 41))) throw new FailedPredicateException(this, "precpred(_ctx, 41)");
						setState(1649);
						if (!(this.notLineTerminator())) throw new FailedPredicateException(this, "this.notLineTerminator()");
						setState(1650);
						match(MinusMinus);
						}
						break;
					case 23:
						{
						_localctx = new TemplateStringExpressionContext(new SingleExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_singleExpression);
						setState(1651);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(1652);
						match(TemplateStringLiteral);
						}
						break;
					}
					} 
				}
				setState(1657);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,208,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ArrowFunctionDeclarationContext extends ParserRuleContext {
		public ArrowFunctionParametersContext arrowFunctionParameters() {
			return getRuleContext(ArrowFunctionParametersContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(TypeScriptParser.ARROW, 0); }
		public ArrowFunctionBodyContext arrowFunctionBody() {
			return getRuleContext(ArrowFunctionBodyContext.class,0);
		}
		public TerminalNode Async() { return getToken(TypeScriptParser.Async, 0); }
		public TypeAnnotationContext typeAnnotation() {
			return getRuleContext(TypeAnnotationContext.class,0);
		}
		public ArrowFunctionDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrowFunctionDeclaration; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArrowFunctionDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArrowFunctionDeclaration(this);
		}
	}

	public final ArrowFunctionDeclarationContext arrowFunctionDeclaration() throws RecognitionException {
		ArrowFunctionDeclarationContext _localctx = new ArrowFunctionDeclarationContext(_ctx, getState());
		enterRule(_localctx, 256, RULE_arrowFunctionDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1659);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Async) {
				{
				setState(1658);
				match(Async);
				}
			}

			setState(1661);
			arrowFunctionParameters();
			setState(1663);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Colon) {
				{
				setState(1662);
				typeAnnotation();
				}
			}

			setState(1665);
			match(ARROW);
			setState(1666);
			arrowFunctionBody();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrowFunctionParametersContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public TerminalNode OpenParen() { return getToken(TypeScriptParser.OpenParen, 0); }
		public TerminalNode CloseParen() { return getToken(TypeScriptParser.CloseParen, 0); }
		public FormalParameterListContext formalParameterList() {
			return getRuleContext(FormalParameterListContext.class,0);
		}
		public ArrowFunctionParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrowFunctionParameters; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArrowFunctionParameters(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArrowFunctionParameters(this);
		}
	}

	public final ArrowFunctionParametersContext arrowFunctionParameters() throws RecognitionException {
		ArrowFunctionParametersContext _localctx = new ArrowFunctionParametersContext(_ctx, getState());
		enterRule(_localctx, 258, RULE_arrowFunctionParameters);
		int _la;
		try {
			setState(1674);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(1668);
				match(Identifier);
				}
				break;
			case OpenParen:
				enterOuterAlt(_localctx, 2);
				{
				setState(1669);
				match(OpenParen);
				setState(1671);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << OpenBracket) | (1L << OpenBrace) | (1L << Ellipsis))) != 0) || ((((_la - 108)) & ~0x3f) == 0 && ((1L << (_la - 108)) & ((1L << (Private - 108)) | (1L << (Public - 108)) | (1L << (Protected - 108)) | (1L << (Identifier - 108)))) != 0)) {
					{
					setState(1670);
					formalParameterList();
					}
				}

				setState(1673);
				match(CloseParen);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrowFunctionBodyContext extends ParserRuleContext {
		public SingleExpressionContext singleExpression() {
			return getRuleContext(SingleExpressionContext.class,0);
		}
		public TerminalNode OpenBrace() { return getToken(TypeScriptParser.OpenBrace, 0); }
		public FunctionBodyContext functionBody() {
			return getRuleContext(FunctionBodyContext.class,0);
		}
		public TerminalNode CloseBrace() { return getToken(TypeScriptParser.CloseBrace, 0); }
		public ArrowFunctionBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrowFunctionBody; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterArrowFunctionBody(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitArrowFunctionBody(this);
		}
	}

	public final ArrowFunctionBodyContext arrowFunctionBody() throws RecognitionException {
		ArrowFunctionBodyContext _localctx = new ArrowFunctionBodyContext(_ctx, getState());
		enterRule(_localctx, 260, RULE_arrowFunctionBody);
		try {
			setState(1681);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,213,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1676);
				singleExpression(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1677);
				match(OpenBrace);
				setState(1678);
				functionBody();
				setState(1679);
				match(CloseBrace);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentOperatorContext extends ParserRuleContext {
		public TerminalNode MultiplyAssign() { return getToken(TypeScriptParser.MultiplyAssign, 0); }
		public TerminalNode DivideAssign() { return getToken(TypeScriptParser.DivideAssign, 0); }
		public TerminalNode ModulusAssign() { return getToken(TypeScriptParser.ModulusAssign, 0); }
		public TerminalNode PlusAssign() { return getToken(TypeScriptParser.PlusAssign, 0); }
		public TerminalNode MinusAssign() { return getToken(TypeScriptParser.MinusAssign, 0); }
		public TerminalNode LeftShiftArithmeticAssign() { return getToken(TypeScriptParser.LeftShiftArithmeticAssign, 0); }
		public TerminalNode RightShiftArithmeticAssign() { return getToken(TypeScriptParser.RightShiftArithmeticAssign, 0); }
		public TerminalNode RightShiftLogicalAssign() { return getToken(TypeScriptParser.RightShiftLogicalAssign, 0); }
		public TerminalNode BitAndAssign() { return getToken(TypeScriptParser.BitAndAssign, 0); }
		public TerminalNode BitXorAssign() { return getToken(TypeScriptParser.BitXorAssign, 0); }
		public TerminalNode BitOrAssign() { return getToken(TypeScriptParser.BitOrAssign, 0); }
		public TerminalNode PowerAssign() { return getToken(TypeScriptParser.PowerAssign, 0); }
		public AssignmentOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignmentOperator; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterAssignmentOperator(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitAssignmentOperator(this);
		}
	}

	public final AssignmentOperatorContext assignmentOperator() throws RecognitionException {
		AssignmentOperatorContext _localctx = new AssignmentOperatorContext(_ctx, getState());
		enterRule(_localctx, 262, RULE_assignmentOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1683);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MultiplyAssign) | (1L << DivideAssign) | (1L << ModulusAssign) | (1L << PlusAssign) | (1L << MinusAssign) | (1L << LeftShiftArithmeticAssign) | (1L << RightShiftArithmeticAssign) | (1L << RightShiftLogicalAssign) | (1L << BitAndAssign) | (1L << BitXorAssign) | (1L << BitOrAssign) | (1L << PowerAssign))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LiteralContext extends ParserRuleContext {
		public TerminalNode NullLiteral() { return getToken(TypeScriptParser.NullLiteral, 0); }
		public TerminalNode BooleanLiteral() { return getToken(TypeScriptParser.BooleanLiteral, 0); }
		public TerminalNode StringLiteral() { return getToken(TypeScriptParser.StringLiteral, 0); }
		public TerminalNode TemplateStringLiteral() { return getToken(TypeScriptParser.TemplateStringLiteral, 0); }
		public TerminalNode RegularExpressionLiteral() { return getToken(TypeScriptParser.RegularExpressionLiteral, 0); }
		public NumericLiteralContext numericLiteral() {
			return getRuleContext(NumericLiteralContext.class,0);
		}
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitLiteral(this);
		}
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 264, RULE_literal);
		try {
			setState(1691);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NullLiteral:
				enterOuterAlt(_localctx, 1);
				{
				setState(1685);
				match(NullLiteral);
				}
				break;
			case BooleanLiteral:
				enterOuterAlt(_localctx, 2);
				{
				setState(1686);
				match(BooleanLiteral);
				}
				break;
			case StringLiteral:
				enterOuterAlt(_localctx, 3);
				{
				setState(1687);
				match(StringLiteral);
				}
				break;
			case TemplateStringLiteral:
				enterOuterAlt(_localctx, 4);
				{
				setState(1688);
				match(TemplateStringLiteral);
				}
				break;
			case RegularExpressionLiteral:
				enterOuterAlt(_localctx, 5);
				{
				setState(1689);
				match(RegularExpressionLiteral);
				}
				break;
			case DecimalLiteral:
			case HexIntegerLiteral:
			case OctalIntegerLiteral:
			case OctalIntegerLiteral2:
			case BinaryIntegerLiteral:
				enterOuterAlt(_localctx, 6);
				{
				setState(1690);
				numericLiteral();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NumericLiteralContext extends ParserRuleContext {
		public TerminalNode DecimalLiteral() { return getToken(TypeScriptParser.DecimalLiteral, 0); }
		public TerminalNode HexIntegerLiteral() { return getToken(TypeScriptParser.HexIntegerLiteral, 0); }
		public TerminalNode OctalIntegerLiteral() { return getToken(TypeScriptParser.OctalIntegerLiteral, 0); }
		public TerminalNode OctalIntegerLiteral2() { return getToken(TypeScriptParser.OctalIntegerLiteral2, 0); }
		public TerminalNode BinaryIntegerLiteral() { return getToken(TypeScriptParser.BinaryIntegerLiteral, 0); }
		public NumericLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_numericLiteral; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterNumericLiteral(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitNumericLiteral(this);
		}
	}

	public final NumericLiteralContext numericLiteral() throws RecognitionException {
		NumericLiteralContext _localctx = new NumericLiteralContext(_ctx, getState());
		enterRule(_localctx, 266, RULE_numericLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1693);
			_la = _input.LA(1);
			if ( !(((((_la - 63)) & ~0x3f) == 0 && ((1L << (_la - 63)) & ((1L << (DecimalLiteral - 63)) | (1L << (HexIntegerLiteral - 63)) | (1L << (OctalIntegerLiteral - 63)) | (1L << (OctalIntegerLiteral2 - 63)) | (1L << (BinaryIntegerLiteral - 63)))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IdentifierNameContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public ReservedWordContext reservedWord() {
			return getRuleContext(ReservedWordContext.class,0);
		}
		public IdentifierNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifierName; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterIdentifierName(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitIdentifierName(this);
		}
	}

	public final IdentifierNameContext identifierName() throws RecognitionException {
		IdentifierNameContext _localctx = new IdentifierNameContext(_ctx, getState());
		enterRule(_localctx, 268, RULE_identifierName);
		try {
			setState(1697);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(1695);
				match(Identifier);
				}
				break;
			case NullLiteral:
			case BooleanLiteral:
			case Break:
			case Do:
			case Instanceof:
			case Typeof:
			case Case:
			case Else:
			case New:
			case Var:
			case Catch:
			case Finally:
			case Return:
			case Void:
			case Continue:
			case For:
			case Switch:
			case While:
			case Debugger:
			case Function:
			case This:
			case With:
			case Default:
			case If:
			case Throw:
			case Delete:
			case In:
			case Try:
			case As:
			case From:
			case ReadOnly:
			case Async:
			case Class:
			case Enum:
			case Extends:
			case Super:
			case Const:
			case Export:
			case Import:
			case Await:
			case Implements:
			case Let:
			case Private:
			case Public:
			case Interface:
			case Package:
			case Protected:
			case Static:
			case Yield:
				enterOuterAlt(_localctx, 2);
				{
				setState(1696);
				reservedWord();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReservedWordContext extends ParserRuleContext {
		public KeywordContext keyword() {
			return getRuleContext(KeywordContext.class,0);
		}
		public TerminalNode NullLiteral() { return getToken(TypeScriptParser.NullLiteral, 0); }
		public TerminalNode BooleanLiteral() { return getToken(TypeScriptParser.BooleanLiteral, 0); }
		public ReservedWordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_reservedWord; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterReservedWord(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitReservedWord(this);
		}
	}

	public final ReservedWordContext reservedWord() throws RecognitionException {
		ReservedWordContext _localctx = new ReservedWordContext(_ctx, getState());
		enterRule(_localctx, 270, RULE_reservedWord);
		try {
			setState(1702);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Break:
			case Do:
			case Instanceof:
			case Typeof:
			case Case:
			case Else:
			case New:
			case Var:
			case Catch:
			case Finally:
			case Return:
			case Void:
			case Continue:
			case For:
			case Switch:
			case While:
			case Debugger:
			case Function:
			case This:
			case With:
			case Default:
			case If:
			case Throw:
			case Delete:
			case In:
			case Try:
			case As:
			case From:
			case ReadOnly:
			case Async:
			case Class:
			case Enum:
			case Extends:
			case Super:
			case Const:
			case Export:
			case Import:
			case Await:
			case Implements:
			case Let:
			case Private:
			case Public:
			case Interface:
			case Package:
			case Protected:
			case Static:
			case Yield:
				enterOuterAlt(_localctx, 1);
				{
				setState(1699);
				keyword();
				}
				break;
			case NullLiteral:
				enterOuterAlt(_localctx, 2);
				{
				setState(1700);
				match(NullLiteral);
				}
				break;
			case BooleanLiteral:
				enterOuterAlt(_localctx, 3);
				{
				setState(1701);
				match(BooleanLiteral);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeywordContext extends ParserRuleContext {
		public TerminalNode Break() { return getToken(TypeScriptParser.Break, 0); }
		public TerminalNode Do() { return getToken(TypeScriptParser.Do, 0); }
		public TerminalNode Instanceof() { return getToken(TypeScriptParser.Instanceof, 0); }
		public TerminalNode Typeof() { return getToken(TypeScriptParser.Typeof, 0); }
		public TerminalNode Case() { return getToken(TypeScriptParser.Case, 0); }
		public TerminalNode Else() { return getToken(TypeScriptParser.Else, 0); }
		public TerminalNode New() { return getToken(TypeScriptParser.New, 0); }
		public TerminalNode Var() { return getToken(TypeScriptParser.Var, 0); }
		public TerminalNode Catch() { return getToken(TypeScriptParser.Catch, 0); }
		public TerminalNode Finally() { return getToken(TypeScriptParser.Finally, 0); }
		public TerminalNode Return() { return getToken(TypeScriptParser.Return, 0); }
		public TerminalNode Void() { return getToken(TypeScriptParser.Void, 0); }
		public TerminalNode Continue() { return getToken(TypeScriptParser.Continue, 0); }
		public TerminalNode For() { return getToken(TypeScriptParser.For, 0); }
		public TerminalNode Switch() { return getToken(TypeScriptParser.Switch, 0); }
		public TerminalNode While() { return getToken(TypeScriptParser.While, 0); }
		public TerminalNode Debugger() { return getToken(TypeScriptParser.Debugger, 0); }
		public TerminalNode Function() { return getToken(TypeScriptParser.Function, 0); }
		public TerminalNode This() { return getToken(TypeScriptParser.This, 0); }
		public TerminalNode With() { return getToken(TypeScriptParser.With, 0); }
		public TerminalNode Default() { return getToken(TypeScriptParser.Default, 0); }
		public TerminalNode If() { return getToken(TypeScriptParser.If, 0); }
		public TerminalNode Throw() { return getToken(TypeScriptParser.Throw, 0); }
		public TerminalNode Delete() { return getToken(TypeScriptParser.Delete, 0); }
		public TerminalNode In() { return getToken(TypeScriptParser.In, 0); }
		public TerminalNode Try() { return getToken(TypeScriptParser.Try, 0); }
		public TerminalNode ReadOnly() { return getToken(TypeScriptParser.ReadOnly, 0); }
		public TerminalNode Async() { return getToken(TypeScriptParser.Async, 0); }
		public TerminalNode From() { return getToken(TypeScriptParser.From, 0); }
		public TerminalNode Class() { return getToken(TypeScriptParser.Class, 0); }
		public TerminalNode Enum() { return getToken(TypeScriptParser.Enum, 0); }
		public TerminalNode Extends() { return getToken(TypeScriptParser.Extends, 0); }
		public TerminalNode Super() { return getToken(TypeScriptParser.Super, 0); }
		public TerminalNode Const() { return getToken(TypeScriptParser.Const, 0); }
		public TerminalNode Export() { return getToken(TypeScriptParser.Export, 0); }
		public TerminalNode Import() { return getToken(TypeScriptParser.Import, 0); }
		public TerminalNode Implements() { return getToken(TypeScriptParser.Implements, 0); }
		public TerminalNode Let() { return getToken(TypeScriptParser.Let, 0); }
		public TerminalNode Private() { return getToken(TypeScriptParser.Private, 0); }
		public TerminalNode Public() { return getToken(TypeScriptParser.Public, 0); }
		public TerminalNode Interface() { return getToken(TypeScriptParser.Interface, 0); }
		public TerminalNode Package() { return getToken(TypeScriptParser.Package, 0); }
		public TerminalNode Protected() { return getToken(TypeScriptParser.Protected, 0); }
		public TerminalNode Static() { return getToken(TypeScriptParser.Static, 0); }
		public TerminalNode Yield() { return getToken(TypeScriptParser.Yield, 0); }
		public TerminalNode Await() { return getToken(TypeScriptParser.Await, 0); }
		public TerminalNode As() { return getToken(TypeScriptParser.As, 0); }
		public KeywordContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_keyword; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterKeyword(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitKeyword(this);
		}
	}

	public final KeywordContext keyword() throws RecognitionException {
		KeywordContext _localctx = new KeywordContext(_ctx, getState());
		enterRule(_localctx, 272, RULE_keyword);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1704);
			_la = _input.LA(1);
			if ( !(((((_la - 68)) & ~0x3f) == 0 && ((1L << (_la - 68)) & ((1L << (Break - 68)) | (1L << (Do - 68)) | (1L << (Instanceof - 68)) | (1L << (Typeof - 68)) | (1L << (Case - 68)) | (1L << (Else - 68)) | (1L << (New - 68)) | (1L << (Var - 68)) | (1L << (Catch - 68)) | (1L << (Finally - 68)) | (1L << (Return - 68)) | (1L << (Void - 68)) | (1L << (Continue - 68)) | (1L << (For - 68)) | (1L << (Switch - 68)) | (1L << (While - 68)) | (1L << (Debugger - 68)) | (1L << (Function - 68)) | (1L << (This - 68)) | (1L << (With - 68)) | (1L << (Default - 68)) | (1L << (If - 68)) | (1L << (Throw - 68)) | (1L << (Delete - 68)) | (1L << (In - 68)) | (1L << (Try - 68)) | (1L << (As - 68)) | (1L << (From - 68)) | (1L << (ReadOnly - 68)) | (1L << (Async - 68)) | (1L << (Class - 68)) | (1L << (Enum - 68)) | (1L << (Extends - 68)) | (1L << (Super - 68)) | (1L << (Const - 68)) | (1L << (Export - 68)) | (1L << (Import - 68)) | (1L << (Await - 68)) | (1L << (Implements - 68)) | (1L << (Let - 68)) | (1L << (Private - 68)) | (1L << (Public - 68)) | (1L << (Interface - 68)) | (1L << (Package - 68)) | (1L << (Protected - 68)) | (1L << (Static - 68)) | (1L << (Yield - 68)))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GetterContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public PropertyNameContext propertyName() {
			return getRuleContext(PropertyNameContext.class,0);
		}
		public GetterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_getter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterGetter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitGetter(this);
		}
	}

	public final GetterContext getter() throws RecognitionException {
		GetterContext _localctx = new GetterContext(_ctx, getState());
		enterRule(_localctx, 274, RULE_getter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1706);
			match(Identifier);
			setState(1707);
			if (!(this.p("get"))) throw new FailedPredicateException(this, "this.p(\"get\")");
			setState(1708);
			propertyName();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SetterContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(TypeScriptParser.Identifier, 0); }
		public PropertyNameContext propertyName() {
			return getRuleContext(PropertyNameContext.class,0);
		}
		public SetterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_setter; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterSetter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitSetter(this);
		}
	}

	public final SetterContext setter() throws RecognitionException {
		SetterContext _localctx = new SetterContext(_ctx, getState());
		enterRule(_localctx, 276, RULE_setter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(1710);
			match(Identifier);
			setState(1711);
			if (!(this.p("set"))) throw new FailedPredicateException(this, "this.p(\"set\")");
			setState(1712);
			propertyName();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EosContext extends ParserRuleContext {
		public TerminalNode SemiColon() { return getToken(TypeScriptParser.SemiColon, 0); }
		public TerminalNode EOF() { return getToken(TypeScriptParser.EOF, 0); }
		public EosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eos; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).enterEos(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof TypeScriptParserListener ) ((TypeScriptParserListener)listener).exitEos(this);
		}
	}

	public final EosContext eos() throws RecognitionException {
		EosContext _localctx = new EosContext(_ctx, getState());
		enterRule(_localctx, 278, RULE_eos);
		try {
			setState(1718);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,217,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(1714);
				match(SemiColon);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(1715);
				match(EOF);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(1716);
				if (!(this.lineTerminatorAhead())) throw new FailedPredicateException(this, "this.lineTerminatorAhead()");
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(1717);
				if (!(this.closeBrace())) throw new FailedPredicateException(this, "this.closeBrace()");
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 10:
			return unionOrIntersectionOrPrimaryType_sempred((UnionOrIntersectionOrPrimaryTypeContext)_localctx, predIndex);
		case 11:
			return primaryType_sempred((PrimaryTypeContext)_localctx, predIndex);
		case 21:
			return arrayType_sempred((ArrayTypeContext)_localctx, predIndex);
		case 57:
			return decoratorMemberExpression_sempred((DecoratorMemberExpressionContext)_localctx, predIndex);
		case 73:
			return expressionStatement_sempred((ExpressionStatementContext)_localctx, predIndex);
		case 75:
			return iterationStatement_sempred((IterationStatementContext)_localctx, predIndex);
		case 77:
			return continueStatement_sempred((ContinueStatementContext)_localctx, predIndex);
		case 78:
			return breakStatement_sempred((BreakStatementContext)_localctx, predIndex);
		case 79:
			return returnStatement_sempred((ReturnStatementContext)_localctx, predIndex);
		case 80:
			return yieldStatement_sempred((YieldStatementContext)_localctx, predIndex);
		case 88:
			return throwStatement_sempred((ThrowStatementContext)_localctx, predIndex);
		case 127:
			return singleExpression_sempred((SingleExpressionContext)_localctx, predIndex);
		case 137:
			return getter_sempred((GetterContext)_localctx, predIndex);
		case 138:
			return setter_sempred((SetterContext)_localctx, predIndex);
		case 139:
			return eos_sempred((EosContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean unionOrIntersectionOrPrimaryType_sempred(UnionOrIntersectionOrPrimaryTypeContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean primaryType_sempred(PrimaryTypeContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return this.notLineTerminator();
		}
		return true;
	}
	private boolean arrayType_sempred(ArrayTypeContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return this.notLineTerminator();
		}
		return true;
	}
	private boolean decoratorMemberExpression_sempred(DecoratorMemberExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expressionStatement_sempred(ExpressionStatementContext _localctx, int predIndex) {
		switch (predIndex) {
		case 6:
			return this.notOpenBraceAndNotFunction();
		}
		return true;
	}
	private boolean iterationStatement_sempred(IterationStatementContext _localctx, int predIndex) {
		switch (predIndex) {
		case 7:
			return this.p("of");
		case 8:
			return this.p("of");
		}
		return true;
	}
	private boolean continueStatement_sempred(ContinueStatementContext _localctx, int predIndex) {
		switch (predIndex) {
		case 9:
			return this.notLineTerminator();
		}
		return true;
	}
	private boolean breakStatement_sempred(BreakStatementContext _localctx, int predIndex) {
		switch (predIndex) {
		case 10:
			return this.notLineTerminator();
		}
		return true;
	}
	private boolean returnStatement_sempred(ReturnStatementContext _localctx, int predIndex) {
		switch (predIndex) {
		case 11:
			return this.notLineTerminator();
		}
		return true;
	}
	private boolean yieldStatement_sempred(YieldStatementContext _localctx, int predIndex) {
		switch (predIndex) {
		case 12:
			return this.notLineTerminator();
		}
		return true;
	}
	private boolean throwStatement_sempred(ThrowStatementContext _localctx, int predIndex) {
		switch (predIndex) {
		case 13:
			return this.notLineTerminator();
		}
		return true;
	}
	private boolean singleExpression_sempred(SingleExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 14:
			return precpred(_ctx, 30);
		case 15:
			return precpred(_ctx, 29);
		case 16:
			return precpred(_ctx, 28);
		case 17:
			return precpred(_ctx, 27);
		case 18:
			return precpred(_ctx, 26);
		case 19:
			return precpred(_ctx, 25);
		case 20:
			return precpred(_ctx, 24);
		case 21:
			return precpred(_ctx, 23);
		case 22:
			return precpred(_ctx, 22);
		case 23:
			return precpred(_ctx, 21);
		case 24:
			return precpred(_ctx, 20);
		case 25:
			return precpred(_ctx, 19);
		case 26:
			return precpred(_ctx, 18);
		case 27:
			return precpred(_ctx, 17);
		case 28:
			return precpred(_ctx, 16);
		case 29:
			return precpred(_ctx, 15);
		case 30:
			return precpred(_ctx, 14);
		case 31:
			return precpred(_ctx, 46);
		case 32:
			return precpred(_ctx, 45);
		case 33:
			return precpred(_ctx, 44);
		case 34:
			return precpred(_ctx, 42);
		case 35:
			return this.notLineTerminator();
		case 36:
			return precpred(_ctx, 41);
		case 37:
			return this.notLineTerminator();
		case 38:
			return precpred(_ctx, 13);
		}
		return true;
	}
	private boolean getter_sempred(GetterContext _localctx, int predIndex) {
		switch (predIndex) {
		case 39:
			return this.p("get");
		}
		return true;
	}
	private boolean setter_sempred(SetterContext _localctx, int predIndex) {
		switch (predIndex) {
		case 40:
			return this.p("set");
		}
		return true;
	}
	private boolean eos_sempred(EosContext _localctx, int predIndex) {
		switch (predIndex) {
		case 41:
			return this.lineTerminatorAhead();
		case 42:
			return this.closeBrace();
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\u008c\u06bb\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\4K\tK\4L\tL\4M\tM\4N\tN\4O\tO\4P\tP\4Q\tQ\4R\tR\4S\tS\4T\tT"+
		"\4U\tU\4V\tV\4W\tW\4X\tX\4Y\tY\4Z\tZ\4[\t[\4\\\t\\\4]\t]\4^\t^\4_\t_\4"+
		"`\t`\4a\ta\4b\tb\4c\tc\4d\td\4e\te\4f\tf\4g\tg\4h\th\4i\ti\4j\tj\4k\t"+
		"k\4l\tl\4m\tm\4n\tn\4o\to\4p\tp\4q\tq\4r\tr\4s\ts\4t\tt\4u\tu\4v\tv\4"+
		"w\tw\4x\tx\4y\ty\4z\tz\4{\t{\4|\t|\4}\t}\4~\t~\4\177\t\177\4\u0080\t\u0080"+
		"\4\u0081\t\u0081\4\u0082\t\u0082\4\u0083\t\u0083\4\u0084\t\u0084\4\u0085"+
		"\t\u0085\4\u0086\t\u0086\4\u0087\t\u0087\4\u0088\t\u0088\4\u0089\t\u0089"+
		"\4\u008a\t\u008a\4\u008b\t\u008b\4\u008c\t\u008c\4\u008d\t\u008d\3\2\3"+
		"\2\3\2\3\3\3\3\5\3\u0120\n\3\3\4\3\4\5\4\u0124\n\4\3\4\3\4\3\5\3\5\3\5"+
		"\7\5\u012b\n\5\f\5\16\5\u012e\13\5\3\6\3\6\5\6\u0132\n\6\3\6\5\6\u0135"+
		"\n\6\3\7\3\7\3\7\3\b\3\b\5\b\u013c\n\b\3\b\3\b\3\t\3\t\3\t\7\t\u0143\n"+
		"\t\f\t\16\t\u0146\13\t\3\n\3\n\3\13\3\13\3\13\3\13\3\13\5\13\u014f\n\13"+
		"\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\7\f\u015a\n\f\f\f\16\f\u015d\13\f"+
		"\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\r\5\r\u0171\n\r\3\r\3\r\3\r\3\r\7\r\u0177\n\r\f\r\16\r\u017a\13\r\3\16"+
		"\3\16\3\17\3\17\3\17\5\17\u0181\n\17\3\20\3\20\3\20\3\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u0190\n\21\3\22\3\22\5\22\u0194\n"+
		"\22\3\23\3\23\5\23\u0198\n\23\3\23\3\23\3\24\3\24\5\24\u019e\n\24\3\25"+
		"\3\25\3\25\7\25\u01a3\n\25\f\25\16\25\u01a6\13\25\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\5\26\u01af\n\26\5\26\u01b1\n\26\3\27\3\27\3\27\3\27\3"+
		"\27\3\30\3\30\3\30\3\30\3\31\3\31\3\31\7\31\u01bf\n\31\f\31\16\31\u01c2"+
		"\13\31\3\32\5\32\u01c5\n\32\3\32\3\32\5\32\u01c9\n\32\3\32\3\32\3\32\3"+
		"\32\3\33\3\33\5\33\u01d1\n\33\3\33\3\33\5\33\u01d5\n\33\3\33\3\33\3\33"+
		"\3\33\3\34\3\34\3\34\3\35\3\35\3\35\3\35\6\35\u01e2\n\35\r\35\16\35\u01e3"+
		"\3\35\3\35\5\35\u01e8\n\35\3\36\5\36\u01eb\n\36\3\36\3\36\5\36\u01ef\n"+
		"\36\3\36\5\36\u01f2\n\36\3\36\3\36\5\36\u01f6\n\36\3\37\3\37\3\37\3 \5"+
		" \u01fc\n \3 \3 \5 \u0200\n \3 \3 \5 \u0204\n \3!\3!\3!\3!\7!\u020a\n"+
		"!\f!\16!\u020d\13!\3!\3!\3!\5!\u0212\n!\3!\3!\3!\3!\3!\5!\u0219\n!\3!"+
		"\5!\u021c\n!\5!\u021e\n!\5!\u0220\n!\3\"\3\"\3\"\7\"\u0225\n\"\f\"\16"+
		"\"\u0228\13\"\3#\5#\u022b\n#\3#\5#\u022e\n#\3#\3#\5#\u0232\n#\3$\3$\3"+
		"%\3%\5%\u0238\n%\3&\3&\3&\7&\u023d\n&\f&\16&\u0240\13&\3\'\5\'\u0243\n"+
		"\'\3\'\5\'\u0246\n\'\3\'\3\'\3\'\5\'\u024b\n\'\3\'\5\'\u024e\n\'\3\'\5"+
		"\'\u0251\n\'\3(\3(\3(\3(\5(\u0257\n(\3)\3)\5)\u025b\n)\3)\3)\5)\u025f"+
		"\n)\3)\3)\5)\u0263\n)\3*\3*\3*\3*\3*\3*\3*\3+\3+\5+\u026e\n+\3+\3+\3,"+
		"\3,\3,\5,\u0275\n,\3,\3,\3,\3,\3-\5-\u027c\n-\3-\3-\3-\5-\u0281\n-\3-"+
		"\3-\3-\3-\3-\3-\5-\u0289\n-\3.\5.\u028c\n.\3.\3.\3.\5.\u0291\n.\3.\5."+
		"\u0294\n.\3.\3.\5.\u0298\n.\3/\3/\3/\3\60\3\60\3\60\7\60\u02a0\n\60\f"+
		"\60\16\60\u02a3\13\60\3\61\5\61\u02a6\n\61\3\61\3\61\3\61\3\61\5\61\u02ac"+
		"\n\61\3\61\3\61\3\62\3\62\5\62\u02b2\n\62\3\63\3\63\3\63\7\63\u02b7\n"+
		"\63\f\63\16\63\u02ba\13\63\3\64\3\64\3\64\5\64\u02bf\n\64\3\65\3\65\3"+
		"\65\3\65\5\65\u02c5\n\65\3\65\3\65\3\66\3\66\6\66\u02cb\n\66\r\66\16\66"+
		"\u02cc\3\66\7\66\u02d0\n\66\f\66\16\66\u02d3\13\66\3\67\3\67\3\67\3\67"+
		"\3\67\3\67\3\67\3\67\3\67\3\67\3\67\3\67\5\67\u02e1\n\67\38\38\39\69\u02e6"+
		"\n9\r9\169\u02e7\3:\3:\3:\5:\u02ed\n:\3;\3;\3;\3;\3;\3;\5;\u02f5\n;\3"+
		";\3;\3;\7;\u02fa\n;\f;\16;\u02fd\13;\3<\3<\3<\3=\5=\u0303\n=\3=\3=\3>"+
		"\5>\u0308\n>\3>\3>\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?"+
		"\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\3?\5?\u0329\n?\3@\3@\5@\u032d\n@\3@"+
		"\3@\3A\6A\u0332\nA\rA\16A\u0333\3B\3B\3B\3B\5B\u033a\nB\3B\3B\3C\3C\3"+
		"C\3C\5C\u0342\nC\3C\3C\3D\3D\3D\3D\3D\5D\u034b\nD\3D\3D\5D\u034f\nD\3"+
		"D\3D\3D\3E\3E\3E\5E\u0357\nE\3E\3E\3E\3E\7E\u035d\nE\fE\16E\u0360\13E"+
		"\3E\3E\3F\3F\5F\u0366\nF\3F\3F\5F\u036a\nF\3G\3G\5G\u036e\nG\3G\3G\5G"+
		"\u0372\nG\3G\5G\u0375\nG\3G\5G\u0378\nG\3G\5G\u037b\nG\3G\3G\5G\u037f"+
		"\nG\5G\u0381\nG\3H\3H\3H\7H\u0386\nH\fH\16H\u0389\13H\3I\3I\5I\u038d\n"+
		"I\3I\5I\u0390\nI\3I\3I\5I\u0394\nI\3I\5I\u0397\nI\3J\3J\3K\3K\3K\5K\u039e"+
		"\nK\3L\3L\3L\3L\3L\3L\3L\5L\u03a7\nL\3M\3M\3M\3M\3M\3M\3M\3M\3M\3M\3M"+
		"\3M\3M\3M\3M\3M\3M\5M\u03ba\nM\3M\3M\5M\u03be\nM\3M\3M\5M\u03c2\nM\3M"+
		"\3M\3M\3M\3M\3M\3M\3M\5M\u03cc\nM\3M\3M\5M\u03d0\nM\3M\3M\3M\3M\3M\5M"+
		"\u03d7\nM\3M\3M\3M\3M\3M\5M\u03de\nM\3M\3M\3M\3M\3M\3M\5M\u03e6\nM\3M"+
		"\3M\3M\3M\3M\3M\5M\u03ee\nM\3M\3M\3M\3M\5M\u03f4\nM\3N\3N\3O\3O\3O\5O"+
		"\u03fb\nO\3O\3O\3P\3P\3P\5P\u0402\nP\3P\3P\3Q\3Q\3Q\5Q\u0409\nQ\3Q\3Q"+
		"\3R\3R\3R\5R\u0410\nR\3R\3R\3S\3S\3S\3S\3S\3S\3T\3T\3T\3T\3T\3T\3U\3U"+
		"\5U\u0422\nU\3U\3U\5U\u0426\nU\5U\u0428\nU\3U\3U\3V\6V\u042d\nV\rV\16"+
		"V\u042e\3W\3W\3W\3W\5W\u0435\nW\3X\3X\3X\5X\u043a\nX\3Y\3Y\3Y\3Y\3Z\3"+
		"Z\3Z\3Z\3Z\3[\3[\3[\3[\5[\u0449\n[\3[\5[\u044c\n[\3\\\3\\\3\\\5\\\u0451"+
		"\n\\\3\\\5\\\u0454\n\\\3\\\3\\\3]\3]\3]\5]\u045b\n]\3^\3^\3^\3_\3_\3_"+
		"\3`\5`\u0464\n`\3`\3`\5`\u0468\n`\3`\3`\3`\3`\3`\3`\3`\5`\u0471\n`\3a"+
		"\5a\u0474\na\3a\3a\3a\5a\u0479\na\3a\3a\3a\3b\5b\u047f\nb\3b\5b\u0482"+
		"\nb\3c\3c\7c\u0486\nc\fc\16c\u0489\13c\3c\3c\3d\3d\3d\3e\3e\3e\3f\3f\3"+
		"f\3f\5f\u0497\nf\3g\3g\5g\u049b\ng\3g\5g\u049e\ng\3g\3g\5g\u04a2\ng\3"+
		"g\5g\u04a5\ng\3g\3g\3g\3g\5g\u04ab\ng\3g\5g\u04ae\ng\3g\3g\3g\3g\3g\3"+
		"g\3g\5g\u04b7\ng\3g\3g\5g\u04bb\ng\3g\5g\u04be\ng\3g\3g\5g\u04c2\ng\3"+
		"g\5g\u04c5\ng\3h\5h\u04c8\nh\3h\5h\u04cb\nh\3h\5h\u04ce\nh\3h\5h\u04d1"+
		"\nh\3i\3i\3i\3j\5j\u04d7\nj\3j\3j\3j\5j\u04dc\nj\3j\3j\3j\3j\3j\3k\3k"+
		"\3k\5k\u04e6\nk\3k\3k\5k\u04ea\nk\3k\3k\3k\3k\3k\3l\3l\3l\3l\7l\u04f5"+
		"\nl\fl\16l\u04f8\13l\3l\5l\u04fb\nl\3l\3l\3m\3m\3m\3n\3n\3n\3n\7n\u0506"+
		"\nn\fn\16n\u0509\13n\3n\5n\u050c\nn\3n\3n\3o\3o\3o\3o\3o\5o\u0515\no\3"+
		"o\3o\3o\3o\3o\3p\3p\3p\7p\u051f\np\fp\16p\u0522\13p\3p\3p\5p\u0526\np"+
		"\3p\3p\3p\3p\3p\5p\u052d\np\5p\u052f\np\3q\5q\u0532\nq\3q\3q\5q\u0536"+
		"\nq\3q\3q\5q\u053a\nq\3r\3r\3r\3s\5s\u0540\ns\3t\6t\u0543\nt\rt\16t\u0544"+
		"\3u\3u\5u\u0549\nu\3u\3u\3v\3v\6v\u054f\nv\rv\16v\u0550\3v\7v\u0554\n"+
		"v\fv\16v\u0557\13v\3v\6v\u055a\nv\rv\16v\u055b\3v\5v\u055f\nv\3v\5v\u0562"+
		"\nv\3w\3w\3w\5w\u0567\nw\3x\3x\3x\3x\7x\u056d\nx\fx\16x\u0570\13x\5x\u0572"+
		"\nx\3x\5x\u0575\nx\3x\3x\3y\3y\3y\3y\3y\3y\3y\3y\3y\3y\3y\5y\u0584\ny"+
		"\3y\5y\u0587\ny\3y\3y\3y\5y\u058c\ny\3y\3y\3y\3y\3y\3y\3y\3y\3y\3y\5y"+
		"\u0598\ny\3z\3z\3z\3z\5z\u059e\nz\3z\3z\3z\3z\3{\3{\3{\3{\5{\u05a8\n{"+
		"\3{\5{\u05ab\n{\3{\3{\3{\3{\3{\3|\3|\3|\3|\3|\3|\3|\5|\u05b9\n|\3}\3}"+
		"\3}\3}\7}\u05bf\n}\f}\16}\u05c2\13}\3}\3}\5}\u05c6\n}\3}\5}\u05c9\n}\3"+
		"}\3}\3~\3~\3~\3\177\3\177\3\177\7\177\u05d3\n\177\f\177\16\177\u05d6\13"+
		"\177\3\177\5\177\u05d9\n\177\3\u0080\5\u0080\u05dc\n\u0080\3\u0080\3\u0080"+
		"\5\u0080\u05e0\n\u0080\3\u0080\3\u0080\5\u0080\u05e4\n\u0080\3\u0080\3"+
		"\u0080\5\u0080\u05e8\n\u0080\3\u0080\3\u0080\3\u0080\3\u0080\3\u0081\3"+
		"\u0081\3\u0081\3\u0081\3\u0081\5\u0081\u05f3\n\u0081\3\u0081\3\u0081\3"+
		"\u0081\3\u0081\5\u0081\u05f9\n\u0081\3\u0081\5\u0081\u05fc\n\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\5\u0081"+
		"\u0619\n\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\5\u0081\u0625\n\u0081\5\u0081\u0627\n\u0081\3"+
		"\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\5\u0081"+
		"\u0667\n\u0081\3\u0081\3\u0081\5\u0081\u066b\n\u0081\3\u0081\3\u0081\3"+
		"\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081\3\u0081"+
		"\7\u0081\u0678\n\u0081\f\u0081\16\u0081\u067b\13\u0081\3\u0082\5\u0082"+
		"\u067e\n\u0082\3\u0082\3\u0082\5\u0082\u0682\n\u0082\3\u0082\3\u0082\3"+
		"\u0082\3\u0083\3\u0083\3\u0083\5\u0083\u068a\n\u0083\3\u0083\5\u0083\u068d"+
		"\n\u0083\3\u0084\3\u0084\3\u0084\3\u0084\3\u0084\5\u0084\u0694\n\u0084"+
		"\3\u0085\3\u0085\3\u0086\3\u0086\3\u0086\3\u0086\3\u0086\3\u0086\5\u0086"+
		"\u069e\n\u0086\3\u0087\3\u0087\3\u0088\3\u0088\5\u0088\u06a4\n\u0088\3"+
		"\u0089\3\u0089\3\u0089\5\u0089\u06a9\n\u0089\3\u008a\3\u008a\3\u008b\3"+
		"\u008b\3\u008b\3\u008b\3\u008c\3\u008c\3\u008c\3\u008c\3\u008d\3\u008d"+
		"\3\u008d\3\u008d\5\u008d\u06b9\n\u008d\3\u008d\2\6\26\30t\u0100\u008e"+
		"\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFH"+
		"JLNPRTVXZ\\^`bdfhjlnprtvxz|~\u0080\u0082\u0084\u0086\u0088\u008a\u008c"+
		"\u008e\u0090\u0092\u0094\u0096\u0098\u009a\u009c\u009e\u00a0\u00a2\u00a4"+
		"\u00a6\u00a8\u00aa\u00ac\u00ae\u00b0\u00b2\u00b4\u00b6\u00b8\u00ba\u00bc"+
		"\u00be\u00c0\u00c2\u00c4\u00c6\u00c8\u00ca\u00cc\u00ce\u00d0\u00d2\u00d4"+
		"\u00d6\u00d8\u00da\u00dc\u00de\u00e0\u00e2\u00e4\u00e6\u00e8\u00ea\u00ec"+
		"\u00ee\u00f0\u00f2\u00f4\u00f6\u00f8\u00fa\u00fc\u00fe\u0100\u0102\u0104"+
		"\u0106\u0108\u010a\u010c\u010e\u0110\u0112\u0114\u0116\u0118\2\20\4\2"+
		"QQuy\3\2\r\16\4\2norr\4\2vvxx\5\2MMhhmm\4\2\17\17\21\21\4\2\32\32\35\36"+
		"\3\2\26\27\3\2\"$\3\2%(\3\2),\4\2\62<>>\3\2AE\3\2Ft\2\u077e\2\u011a\3"+
		"\2\2\2\4\u011f\3\2\2\2\6\u0121\3\2\2\2\b\u0127\3\2\2\2\n\u0134\3\2\2\2"+
		"\f\u0136\3\2\2\2\16\u0139\3\2\2\2\20\u013f\3\2\2\2\22\u0147\3\2\2\2\24"+
		"\u014e\3\2\2\2\26\u0150\3\2\2\2\30\u0170\3\2\2\2\32\u017b\3\2\2\2\34\u017d"+
		"\3\2\2\2\36\u0182\3\2\2\2 \u0186\3\2\2\2\"\u0193\3\2\2\2$\u0195\3\2\2"+
		"\2&\u019b\3\2\2\2(\u019f\3\2\2\2*\u01b0\3\2\2\2,\u01b2\3\2\2\2.\u01b7"+
		"\3\2\2\2\60\u01bb\3\2\2\2\62\u01c4\3\2\2\2\64\u01ce\3\2\2\2\66\u01da\3"+
		"\2\2\28\u01e7\3\2\2\2:\u01ea\3\2\2\2<\u01f7\3\2\2\2>\u01fb\3\2\2\2@\u021f"+
		"\3\2\2\2B\u0221\3\2\2\2D\u022a\3\2\2\2F\u0233\3\2\2\2H\u0237\3\2\2\2J"+
		"\u0239\3\2\2\2L\u0242\3\2\2\2N\u0256\3\2\2\2P\u0258\3\2\2\2R\u0264\3\2"+
		"\2\2T\u026b\3\2\2\2V\u0271\3\2\2\2X\u027b\3\2\2\2Z\u028b\3\2\2\2\\\u0299"+
		"\3\2\2\2^\u029c\3\2\2\2`\u02a5\3\2\2\2b\u02af\3\2\2\2d\u02b3\3\2\2\2f"+
		"\u02bb\3\2\2\2h\u02c0\3\2\2\2j\u02c8\3\2\2\2l\u02e0\3\2\2\2n\u02e2\3\2"+
		"\2\2p\u02e5\3\2\2\2r\u02e9\3\2\2\2t\u02f4\3\2\2\2v\u02fe\3\2\2\2x\u0302"+
		"\3\2\2\2z\u0307\3\2\2\2|\u0328\3\2\2\2~\u032a\3\2\2\2\u0080\u0331\3\2"+
		"\2\2\u0082\u0335\3\2\2\2\u0084\u033d\3\2\2\2\u0086\u034a\3\2\2\2\u0088"+
		"\u0356\3\2\2\2\u008a\u0363\3\2\2\2\u008c\u0380\3\2\2\2\u008e\u0382\3\2"+
		"\2\2\u0090\u038a\3\2\2\2\u0092\u0398\3\2\2\2\u0094\u039a\3\2\2\2\u0096"+
		"\u039f\3\2\2\2\u0098\u03f3\3\2\2\2\u009a\u03f5\3\2\2\2\u009c\u03f7\3\2"+
		"\2\2\u009e\u03fe\3\2\2\2\u00a0\u0405\3\2\2\2\u00a2\u040c\3\2\2\2\u00a4"+
		"\u0413\3\2\2\2\u00a6\u0419\3\2\2\2\u00a8\u041f\3\2\2\2\u00aa\u042c\3\2"+
		"\2\2\u00ac\u0430\3\2\2\2\u00ae\u0436\3\2\2\2\u00b0\u043b\3\2\2\2\u00b2"+
		"\u043f\3\2\2\2\u00b4\u0444\3\2\2\2\u00b6\u044d\3\2\2\2\u00b8\u045a\3\2"+
		"\2\2\u00ba\u045c\3\2\2\2\u00bc\u045f\3\2\2\2\u00be\u0463\3\2\2\2\u00c0"+
		"\u0473\3\2\2\2\u00c2\u047e\3\2\2\2\u00c4\u0483\3\2\2\2\u00c6\u048c\3\2"+
		"\2\2\u00c8\u048f\3\2\2\2\u00ca\u0496\3\2\2\2\u00cc\u04c4\3\2\2\2\u00ce"+
		"\u04c7\3\2\2\2\u00d0\u04d2\3\2\2\2\u00d2\u04d6\3\2\2\2\u00d4\u04e2\3\2"+
		"\2\2\u00d6\u04f0\3\2\2\2\u00d8\u04fe\3\2\2\2\u00da\u0501\3\2\2\2\u00dc"+
		"\u050f\3\2\2\2\u00de\u052e\3\2\2\2\u00e0\u0531\3\2\2\2\u00e2\u053b\3\2"+
		"\2\2\u00e4\u053f\3\2\2\2\u00e6\u0542\3\2\2\2\u00e8\u0546\3\2\2\2\u00ea"+
		"\u0561\3\2\2\2\u00ec\u0563\3\2\2\2\u00ee\u0568\3\2\2\2\u00f0\u0597\3\2"+
		"\2\2\u00f2\u0599\3\2\2\2\u00f4\u05a3\3\2\2\2\u00f6\u05b8\3\2\2\2\u00f8"+
		"\u05ba\3\2\2\2\u00fa\u05cc\3\2\2\2\u00fc\u05cf\3\2\2\2\u00fe\u05db\3\2"+
		"\2\2\u0100\u0626\3\2\2\2\u0102\u067d\3\2\2\2\u0104\u068c\3\2\2\2\u0106"+
		"\u0693\3\2\2\2\u0108\u0695\3\2\2\2\u010a\u069d\3\2\2\2\u010c\u069f\3\2"+
		"\2\2\u010e\u06a3\3\2\2\2\u0110\u06a8\3\2\2\2\u0112\u06aa\3\2\2\2\u0114"+
		"\u06ac\3\2\2\2\u0116\u06b0\3\2\2\2\u0118\u06b8\3\2\2\2\u011a\u011b\7\17"+
		"\2\2\u011b\u011c\5\u0100\u0081\2\u011c\3\3\2\2\2\u011d\u0120\5\u00e8u"+
		"\2\u011e\u0120\5\u00eex\2\u011f\u011d\3\2\2\2\u011f\u011e\3\2\2\2\u0120"+
		"\5\3\2\2\2\u0121\u0123\7%\2\2\u0122\u0124\5\b\5\2\u0123\u0122\3\2\2\2"+
		"\u0123\u0124\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0126\7&\2\2\u0126\7\3"+
		"\2\2\2\u0127\u012c\5\n\6\2\u0128\u0129\7\16\2\2\u0129\u012b\5\n\6\2\u012a"+
		"\u0128\3\2\2\2\u012b\u012e\3\2\2\2\u012c\u012a\3\2\2\2\u012c\u012d\3\2"+
		"\2\2\u012d\t\3\2\2\2\u012e\u012c\3\2\2\2\u012f\u0131\7\u0085\2\2\u0130"+
		"\u0132\5\f\7\2\u0131\u0130\3\2\2\2\u0131\u0132\3\2\2\2\u0132\u0135\3\2"+
		"\2\2\u0133\u0135\5\6\4\2\u0134\u012f\3\2\2\2\u0134\u0133\3\2\2\2\u0135"+
		"\13\3\2\2\2\u0136\u0137\7f\2\2\u0137\u0138\5\24\13\2\u0138\r\3\2\2\2\u0139"+
		"\u013b\7%\2\2\u013a\u013c\5\20\t\2\u013b\u013a\3\2\2\2\u013b\u013c\3\2"+
		"\2\2\u013c\u013d\3\2\2\2\u013d\u013e\7&\2\2\u013e\17\3\2\2\2\u013f\u0144"+
		"\5\22\n\2\u0140\u0141\7\16\2\2\u0141\u0143\5\22\n\2\u0142\u0140\3\2\2"+
		"\2\u0143\u0146\3\2\2\2\u0144\u0142\3\2\2\2\u0144\u0145\3\2\2\2\u0145\21"+
		"\3\2\2\2\u0146\u0144\3\2\2\2\u0147\u0148\5\24\13\2\u0148\23\3\2\2\2\u0149"+
		"\u014f\5\26\f\2\u014a\u014f\5\62\32\2\u014b\u014f\5\64\33\2\u014c\u014f"+
		"\5\36\20\2\u014d\u014f\7\u0086\2\2\u014e\u0149\3\2\2\2\u014e\u014a\3\2"+
		"\2\2\u014e\u014b\3\2\2\2\u014e\u014c\3\2\2\2\u014e\u014d\3\2\2\2\u014f"+
		"\25\3\2\2\2\u0150\u0151\b\f\1\2\u0151\u0152\5\30\r\2\u0152\u015b\3\2\2"+
		"\2\u0153\u0154\f\5\2\2\u0154\u0155\7/\2\2\u0155\u015a\5\26\f\6\u0156\u0157"+
		"\f\4\2\2\u0157\u0158\7-\2\2\u0158\u015a\5\26\f\5\u0159\u0153\3\2\2\2\u0159"+
		"\u0156\3\2\2\2\u015a\u015d\3\2\2\2\u015b\u0159\3\2\2\2\u015b\u015c\3\2"+
		"\2\2\u015c\27\3\2\2\2\u015d\u015b\3\2\2\2\u015e\u015f\b\r\1\2\u015f\u0160"+
		"\7\t\2\2\u0160\u0161\5\24\13\2\u0161\u0162\7\n\2\2\u0162\u0171\3\2\2\2"+
		"\u0163\u0171\5\32\16\2\u0164\u0171\5\34\17\2\u0165\u0171\5$\23\2\u0166"+
		"\u0167\7\7\2\2\u0167\u0168\5\60\31\2\u0168\u0169\7\b\2\2\u0169\u0171\3"+
		"\2\2\2\u016a\u0171\5\66\34\2\u016b\u0171\7X\2\2\u016c\u016d\5\34\17\2"+
		"\u016d\u016e\7\u0083\2\2\u016e\u016f\5\30\r\3\u016f\u0171\3\2\2\2\u0170"+
		"\u015e\3\2\2\2\u0170\u0163\3\2\2\2\u0170\u0164\3\2\2\2\u0170\u0165\3\2"+
		"\2\2\u0170\u0166\3\2\2\2\u0170\u016a\3\2\2\2\u0170\u016b\3\2\2\2\u0170"+
		"\u016c\3\2\2\2\u0171\u0178\3\2\2\2\u0172\u0173\f\7\2\2\u0173\u0174\6\r"+
		"\5\2\u0174\u0175\7\7\2\2\u0175\u0177\7\b\2\2\u0176\u0172\3\2\2\2\u0177"+
		"\u017a\3\2\2\2\u0178\u0176\3\2\2\2\u0178\u0179\3\2\2\2\u0179\31\3\2\2"+
		"\2\u017a\u0178\3\2\2\2\u017b\u017c\t\2\2\2\u017c\33\3\2\2\2\u017d\u0180"+
		"\5\"\22\2\u017e\u0181\5 \21\2\u017f\u0181\5\36\20\2\u0180\u017e\3\2\2"+
		"\2\u0180\u017f\3\2\2\2\u0180\u0181\3\2\2\2\u0181\35\3\2\2\2\u0182\u0183"+
		"\7%\2\2\u0183\u0184\5\20\t\2\u0184\u0185\7&\2\2\u0185\37\3\2\2\2\u0186"+
		"\u0187\7%\2\2\u0187\u0188\5\20\t\2\u0188\u0189\7%\2\2\u0189\u018f\5\20"+
		"\t\2\u018a\u018b\7&\2\2\u018b\u018c\5\4\3\2\u018c\u018d\7&\2\2\u018d\u0190"+
		"\3\2\2\2\u018e\u0190\7\"\2\2\u018f\u018a\3\2\2\2\u018f\u018e\3\2\2\2\u0190"+
		"!\3\2\2\2\u0191\u0194\7\u0085\2\2\u0192\u0194\5j\66\2\u0193\u0191\3\2"+
		"\2\2\u0193\u0192\3\2\2\2\u0194#\3\2\2\2\u0195\u0197\7\13\2\2\u0196\u0198"+
		"\5&\24\2\u0197\u0196\3\2\2\2\u0197\u0198\3\2\2\2\u0198\u0199\3\2\2\2\u0199"+
		"\u019a\7\f\2\2\u019a%\3\2\2\2\u019b\u019d\5(\25\2\u019c\u019e\t\3\2\2"+
		"\u019d\u019c\3\2\2\2\u019d\u019e\3\2\2\2\u019e\'\3\2\2\2\u019f\u01a4\5"+
		"*\26\2\u01a0\u01a1\t\3\2\2\u01a1\u01a3\5*\26\2\u01a2\u01a0\3\2\2\2\u01a3"+
		"\u01a6\3\2\2\2\u01a4\u01a2\3\2\2\2\u01a4\u01a5\3\2\2\2\u01a5)\3\2\2\2"+
		"\u01a6\u01a4\3\2\2\2\u01a7\u01b1\5:\36\2\u01a8\u01b1\5> \2\u01a9\u01b1"+
		"\5P)\2\u01aa\u01b1\5R*\2\u01ab\u01ae\5T+\2\u01ac\u01ad\7=\2\2\u01ad\u01af"+
		"\5\24\13\2\u01ae\u01ac\3\2\2\2\u01ae\u01af\3\2\2\2\u01af\u01b1\3\2\2\2"+
		"\u01b0\u01a7\3\2\2\2\u01b0\u01a8\3\2\2\2\u01b0\u01a9\3\2\2\2\u01b0\u01aa"+
		"\3\2\2\2\u01b0\u01ab\3\2\2\2\u01b1+\3\2\2\2\u01b2\u01b3\5\30\r\2\u01b3"+
		"\u01b4\6\27\6\2\u01b4\u01b5\7\7\2\2\u01b5\u01b6\7\b\2\2\u01b6-\3\2\2\2"+
		"\u01b7\u01b8\7\7\2\2\u01b8\u01b9\5\60\31\2\u01b9\u01ba\7\b\2\2\u01ba/"+
		"\3\2\2\2\u01bb\u01c0\5\24\13\2\u01bc\u01bd\7\16\2\2\u01bd\u01bf\5\24\13"+
		"\2\u01be\u01bc\3\2\2\2\u01bf\u01c2\3\2\2\2\u01c0\u01be\3\2\2\2\u01c0\u01c1"+
		"\3\2\2\2\u01c1\61\3\2\2\2\u01c2\u01c0\3\2\2\2\u01c3\u01c5\5\6\4\2\u01c4"+
		"\u01c3\3\2\2\2\u01c4\u01c5\3\2\2\2\u01c5\u01c6\3\2\2\2\u01c6\u01c8\7\t"+
		"\2\2\u01c7\u01c9\5@!\2\u01c8\u01c7\3\2\2\2\u01c8\u01c9\3\2\2\2\u01c9\u01ca"+
		"\3\2\2\2\u01ca\u01cb\7\n\2\2\u01cb\u01cc\7=\2\2\u01cc\u01cd\5\24\13\2"+
		"\u01cd\63\3\2\2\2\u01ce\u01d0\7L\2\2\u01cf\u01d1\5\6\4\2\u01d0\u01cf\3"+
		"\2\2\2\u01d0\u01d1\3\2\2\2\u01d1\u01d2\3\2\2\2\u01d2\u01d4\7\t\2\2\u01d3"+
		"\u01d5\5@!\2\u01d4\u01d3\3\2\2\2\u01d4\u01d5\3\2\2\2\u01d5\u01d6\3\2\2"+
		"\2\u01d6\u01d7\7\n\2\2\u01d7\u01d8\7=\2\2\u01d8\u01d9\5\24\13\2\u01d9"+
		"\65\3\2\2\2\u01da\u01db\7I\2\2\u01db\u01dc\58\35\2\u01dc\67\3\2\2\2\u01dd"+
		"\u01e8\7\u0085\2\2\u01de\u01df\5\u010e\u0088\2\u01df\u01e0\7\23\2\2\u01e0"+
		"\u01e2\3\2\2\2\u01e1\u01de\3\2\2\2\u01e2\u01e3\3\2\2\2\u01e3\u01e1\3\2"+
		"\2\2\u01e3\u01e4\3\2\2\2\u01e4\u01e5\3\2\2\2\u01e5\u01e6\5\u010e\u0088"+
		"\2\u01e6\u01e8\3\2\2\2\u01e7\u01dd\3\2\2\2\u01e7\u01e1\3\2\2\2\u01e89"+
		"\3\2\2\2\u01e9\u01eb\7b\2\2\u01ea\u01e9\3\2\2\2\u01ea\u01eb\3\2\2\2\u01eb"+
		"\u01ec\3\2\2\2\u01ec\u01ee\5\u00f6|\2\u01ed\u01ef\7\20\2\2\u01ee\u01ed"+
		"\3\2\2\2\u01ee\u01ef\3\2\2\2\u01ef\u01f1\3\2\2\2\u01f0\u01f2\5<\37\2\u01f1"+
		"\u01f0\3\2\2\2\u01f1\u01f2\3\2\2\2\u01f2\u01f5\3\2\2\2\u01f3\u01f4\7="+
		"\2\2\u01f4\u01f6\5\24\13\2\u01f5\u01f3\3\2\2\2\u01f5\u01f6\3\2\2\2\u01f6"+
		";\3\2\2\2\u01f7\u01f8\7\21\2\2\u01f8\u01f9\5\24\13\2\u01f9=\3\2\2\2\u01fa"+
		"\u01fc\5\6\4\2\u01fb\u01fa\3\2\2\2\u01fb\u01fc\3\2\2\2\u01fc\u01fd\3\2"+
		"\2\2\u01fd\u01ff\7\t\2\2\u01fe\u0200\5@!\2\u01ff\u01fe\3\2\2\2\u01ff\u0200"+
		"\3\2\2\2\u0200\u0201\3\2\2\2\u0201\u0203\7\n\2\2\u0202\u0204\5<\37\2\u0203"+
		"\u0202\3\2\2\2\u0203\u0204\3\2\2\2\u0204?\3\2\2\2\u0205\u0220\5N(\2\u0206"+
		"\u020b\5\32\16\2\u0207\u0208\7\16\2\2\u0208\u020a\5\32\16\2\u0209\u0207"+
		"\3\2\2\2\u020a\u020d\3\2\2\2\u020b\u0209\3\2\2\2\u020b\u020c\3\2\2\2\u020c"+
		"\u0220\3\2\2\2\u020d\u020b\3\2\2\2\u020e\u0211\5J&\2\u020f\u0210\7\16"+
		"\2\2\u0210\u0212\5N(\2\u0211\u020f\3\2\2\2\u0211\u0212\3\2\2\2\u0212\u0220"+
		"\3\2\2\2\u0213\u021d\5B\"\2\u0214\u021b\7\16\2\2\u0215\u0218\5J&\2\u0216"+
		"\u0217\7\16\2\2\u0217\u0219\5N(\2\u0218\u0216\3\2\2\2\u0218\u0219\3\2"+
		"\2\2\u0219\u021c\3\2\2\2\u021a\u021c\5N(\2\u021b\u0215\3\2\2\2\u021b\u021a"+
		"\3\2\2\2\u021c\u021e\3\2\2\2\u021d\u0214\3\2\2\2\u021d\u021e\3\2\2\2\u021e"+
		"\u0220\3\2\2\2\u021f\u0205\3\2\2\2\u021f\u0206\3\2\2\2\u021f\u020e\3\2"+
		"\2\2\u021f\u0213\3\2\2\2\u0220A\3\2\2\2\u0221\u0226\5D#\2\u0222\u0223"+
		"\7\16\2\2\u0223\u0225\5D#\2\u0224\u0222\3\2\2\2\u0225\u0228\3\2\2\2\u0226"+
		"\u0224\3\2\2\2\u0226\u0227\3\2\2\2\u0227C\3\2\2\2\u0228\u0226\3\2\2\2"+
		"\u0229\u022b\5p9\2\u022a\u0229\3\2\2\2\u022a\u022b\3\2\2\2\u022b\u022d"+
		"\3\2\2\2\u022c\u022e\5F$\2\u022d\u022c\3\2\2\2\u022d\u022e\3\2\2\2\u022e"+
		"\u022f\3\2\2\2\u022f\u0231\5H%\2\u0230\u0232\5<\37\2\u0231\u0230\3\2\2"+
		"\2\u0231\u0232\3\2\2\2\u0232E\3\2\2\2\u0233\u0234\t\4\2\2\u0234G\3\2\2"+
		"\2\u0235\u0238\5\u010e\u0088\2\u0236\u0238\5\4\3\2\u0237\u0235\3\2\2\2"+
		"\u0237\u0236\3\2\2\2\u0238I\3\2\2\2\u0239\u023e\5L\'\2\u023a\u023b\7\16"+
		"\2\2\u023b\u023d\5L\'\2\u023c\u023a\3\2\2\2\u023d\u0240\3\2\2\2\u023e"+
		"\u023c\3\2\2\2\u023e\u023f\3\2\2\2\u023fK\3\2\2\2\u0240\u023e\3\2\2\2"+
		"\u0241\u0243\5p9\2\u0242\u0241\3\2\2\2\u0242\u0243\3\2\2\2\u0243\u0245"+
		"\3\2\2\2\u0244\u0246\5F$\2\u0245\u0244\3\2\2\2\u0245\u0246\3\2\2\2\u0246"+
		"\u0247\3\2\2\2\u0247\u0250\5H%\2\u0248\u024a\7\20\2\2\u0249\u024b\5<\37"+
		"\2\u024a\u0249\3\2\2\2\u024a\u024b\3\2\2\2\u024b\u0251\3\2\2\2\u024c\u024e"+
		"\5<\37\2\u024d\u024c\3\2\2\2\u024d\u024e\3\2\2\2\u024e\u024f\3\2\2\2\u024f"+
		"\u0251\5\2\2\2\u0250\u0248\3\2\2\2\u0250\u024d\3\2\2\2\u0251M\3\2\2\2"+
		"\u0252\u0253\7\22\2\2\u0253\u0257\5D#\2\u0254\u0255\7\22\2\2\u0255\u0257"+
		"\5\u0100\u0081\2\u0256\u0252\3\2\2\2\u0256\u0254\3\2\2\2\u0257O\3\2\2"+
		"\2\u0258\u025a\7L\2\2\u0259\u025b\5\6\4\2\u025a\u0259\3\2\2\2\u025a\u025b"+
		"\3\2\2\2\u025b\u025c\3\2\2\2\u025c\u025e\7\t\2\2\u025d\u025f\5@!\2\u025e"+
		"\u025d\3\2\2\2\u025e\u025f\3\2\2\2\u025f\u0260\3\2\2\2\u0260\u0262\7\n"+
		"\2\2\u0261\u0263\5<\37\2\u0262\u0261\3\2\2\2\u0262\u0263\3\2\2\2\u0263"+
		"Q\3\2\2\2\u0264\u0265\7\7\2\2\u0265\u0266\7\u0085\2\2\u0266\u0267\7\21"+
		"\2\2\u0267\u0268\t\5\2\2\u0268\u0269\7\b\2\2\u0269\u026a\5<\37\2\u026a"+
		"S\3\2\2\2\u026b\u026d\5\u00f6|\2\u026c\u026e\7\20\2\2\u026d\u026c\3\2"+
		"\2\2\u026d\u026e\3\2\2\2\u026e\u026f\3\2\2\2\u026f\u0270\5> \2\u0270U"+
		"\3\2\2\2\u0271\u0272\7z\2\2\u0272\u0274\7\u0085\2\2\u0273\u0275\5\6\4"+
		"\2\u0274\u0273\3\2\2\2\u0274\u0275\3\2\2\2\u0275\u0276\3\2\2\2\u0276\u0277"+
		"\7\17\2\2\u0277\u0278\5\24\13\2\u0278\u0279\7\r\2\2\u0279W\3\2\2\2\u027a"+
		"\u027c\5F$\2\u027b\u027a\3\2\2\2\u027b\u027c\3\2\2\2\u027c\u027d\3\2\2"+
		"\2\u027d\u027e\7}\2\2\u027e\u0280\7\t\2\2\u027f\u0281\5\u00dep\2\u0280"+
		"\u027f\3\2\2\2\u0280\u0281\3\2\2\2\u0281\u0282\3\2\2\2\u0282\u0288\7\n"+
		"\2\2\u0283\u0284\7\13\2\2\u0284\u0285\5\u00e4s\2\u0285\u0286\7\f\2\2\u0286"+
		"\u0289\3\2\2\2\u0287\u0289\7\r\2\2\u0288\u0283\3\2\2\2\u0288\u0287\3\2"+
		"\2\2\u0288\u0289\3\2\2\2\u0289Y\3\2\2\2\u028a\u028c\7i\2\2\u028b\u028a"+
		"\3\2\2\2\u028b\u028c\3\2\2\2\u028c\u028d\3\2\2\2\u028d\u028e\7p\2\2\u028e"+
		"\u0290\7\u0085\2\2\u028f\u0291\5\6\4\2\u0290\u028f\3\2\2\2\u0290\u0291"+
		"\3\2\2\2\u0291\u0293\3\2\2\2\u0292\u0294\5\\/\2\u0293\u0292\3\2\2\2\u0293"+
		"\u0294\3\2\2\2\u0294\u0295\3\2\2\2\u0295\u0297\5$\23\2\u0296\u0298\7\r"+
		"\2\2\u0297\u0296\3\2\2\2\u0297\u0298\3\2\2\2\u0298[\3\2\2\2\u0299\u029a"+
		"\7f\2\2\u029a\u029b\5^\60\2\u029b]\3\2\2\2\u029c\u02a1\5\34\17\2\u029d"+
		"\u029e\7\16\2\2\u029e\u02a0\5\34\17\2\u029f\u029d\3\2\2\2\u02a0\u02a3"+
		"\3\2\2\2\u02a1\u029f\3\2\2\2\u02a1\u02a2\3\2\2\2\u02a2_\3\2\2\2\u02a3"+
		"\u02a1\3\2\2\2\u02a4\u02a6\7h\2\2\u02a5\u02a4\3\2\2\2\u02a5\u02a6\3\2"+
		"\2\2\u02a6\u02a7\3\2\2\2\u02a7\u02a8\7e\2\2\u02a8\u02a9\7\u0085\2\2\u02a9"+
		"\u02ab\7\13\2\2\u02aa\u02ac\5b\62\2\u02ab\u02aa\3\2\2\2\u02ab\u02ac\3"+
		"\2\2\2\u02ac\u02ad\3\2\2\2\u02ad\u02ae\7\f\2\2\u02aea\3\2\2\2\u02af\u02b1"+
		"\5d\63\2\u02b0\u02b2\7\16\2\2\u02b1\u02b0\3\2\2\2\u02b1\u02b2\3\2\2\2"+
		"\u02b2c\3\2\2\2\u02b3\u02b8\5f\64\2\u02b4\u02b5\7\16\2\2\u02b5\u02b7\5"+
		"f\64\2\u02b6\u02b4\3\2\2\2\u02b7\u02ba\3\2\2\2\u02b8\u02b6\3\2\2\2\u02b8"+
		"\u02b9\3\2\2\2\u02b9e\3\2\2\2\u02ba\u02b8\3\2\2\2\u02bb\u02be\5\u00f6"+
		"|\2\u02bc\u02bd\7\17\2\2\u02bd\u02bf\5\u0100\u0081\2\u02be\u02bc\3\2\2"+
		"\2\u02be\u02bf\3\2\2\2\u02bfg\3\2\2\2\u02c0\u02c1\7~\2\2\u02c1\u02c2\5"+
		"j\66\2\u02c2\u02c4\7\13\2\2\u02c3\u02c5\5\u0080A\2\u02c4\u02c3\3\2\2\2"+
		"\u02c4\u02c5\3\2\2\2\u02c5\u02c6\3\2\2\2\u02c6\u02c7\7\f\2\2\u02c7i\3"+
		"\2\2\2\u02c8\u02d1\7\u0085\2\2\u02c9\u02cb\7\23\2\2\u02ca\u02c9\3\2\2"+
		"\2\u02cb\u02cc\3\2\2\2\u02cc\u02ca\3\2\2\2\u02cc\u02cd\3\2\2\2\u02cd\u02ce"+
		"\3\2\2\2\u02ce\u02d0\7\u0085\2\2\u02cf\u02ca\3\2\2\2\u02d0\u02d3\3\2\2"+
		"\2\u02d1\u02cf\3\2\2\2\u02d1\u02d2\3\2\2\2\u02d2k\3\2\2\2\u02d3\u02d1"+
		"\3\2\2\2\u02d4\u02d5\7\u0085\2\2\u02d5\u02d6\7\17\2\2\u02d6\u02d7\5j\66"+
		"\2\u02d7\u02d8\7\r\2\2\u02d8\u02e1\3\2\2\2\u02d9\u02da\7\u0085\2\2\u02da"+
		"\u02db\7\17\2\2\u02db\u02dc\7\177\2\2\u02dc\u02dd\7\t\2\2\u02dd\u02de"+
		"\7\u0086\2\2\u02de\u02df\7\n\2\2\u02df\u02e1\7\r\2\2\u02e0\u02d4\3\2\2"+
		"\2\u02e0\u02d9\3\2\2\2\u02e1m\3\2\2\2\u02e2\u02e3\7\u0086\2\2\u02e3o\3"+
		"\2\2\2\u02e4\u02e6\5r:\2\u02e5\u02e4\3\2\2\2\u02e6\u02e7\3\2\2\2\u02e7"+
		"\u02e5\3\2\2\2\u02e7\u02e8\3\2\2\2\u02e8q\3\2\2\2\u02e9\u02ec\7\u0084"+
		"\2\2\u02ea\u02ed\5t;\2\u02eb\u02ed\5v<\2\u02ec\u02ea\3\2\2\2\u02ec\u02eb"+
		"\3\2\2\2\u02eds\3\2\2\2\u02ee\u02ef\b;\1\2\u02ef\u02f5\7\u0085\2\2\u02f0"+
		"\u02f1\7\t\2\2\u02f1\u02f2\5\u0100\u0081\2\u02f2\u02f3\7\n\2\2\u02f3\u02f5"+
		"\3\2\2\2\u02f4\u02ee\3\2\2\2\u02f4\u02f0\3\2\2\2\u02f5\u02fb\3\2\2\2\u02f6"+
		"\u02f7\f\4\2\2\u02f7\u02f8\7\23\2\2\u02f8\u02fa\5\u010e\u0088\2\u02f9"+
		"\u02f6\3\2\2\2\u02fa\u02fd\3\2\2\2\u02fb\u02f9\3\2\2\2\u02fb\u02fc\3\2"+
		"\2\2\u02fcu\3\2\2\2\u02fd\u02fb\3\2\2\2\u02fe\u02ff\5t;\2\u02ff\u0300"+
		"\5\u00f8}\2\u0300w\3\2\2\2\u0301\u0303\5\u00e6t\2\u0302\u0301\3\2\2\2"+
		"\u0302\u0303\3\2\2\2\u0303\u0304\3\2\2\2\u0304\u0305\7\2\2\3\u0305y\3"+
		"\2\2\2\u0306\u0308\7i\2\2\u0307\u0306\3\2\2\2\u0307\u0308\3\2\2\2\u0308"+
		"\u0309\3\2\2\2\u0309\u030a\5|?\2\u030a{\3\2\2\2\u030b\u0329\5~@\2\u030c"+
		"\u0329\5\u008cG\2\u030d\u0329\5\u0084C\2\u030e\u0329\5\u008aF\2\u030f"+
		"\u0329\5\u0092J\2\u0310\u0329\5\u0082B\2\u0311\u0329\5\u00c0a\2\u0312"+
		"\u0329\5Z.\2\u0313\u0329\5h\65\2\u0314\u0329\5\u0096L\2\u0315\u0329\5"+
		"\u0098M\2\u0316\u0329\5\u009cO\2\u0317\u0329\5\u009eP\2\u0318\u0329\5"+
		"\u00a0Q\2\u0319\u0329\5\u00a2R\2\u031a\u0329\5\u00a4S\2\u031b\u0329\5"+
		"\u00b0Y\2\u031c\u0329\5\u00a6T\2\u031d\u0329\5\u00b2Z\2\u031e\u0329\5"+
		"\u00b4[\2\u031f\u0329\5\u00bc_\2\u0320\u0329\5\u00be`\2\u0321\u0329\5"+
		"\u0102\u0082\2\u0322\u0329\5\u00d4k\2\u0323\u0329\5V,\2\u0324\u0329\5"+
		"`\61\2\u0325\u0329\5\u0094K\2\u0326\u0327\7i\2\2\u0327\u0329\5|?\2\u0328"+
		"\u030b\3\2\2\2\u0328\u030c\3\2\2\2\u0328\u030d\3\2\2\2\u0328\u030e\3\2"+
		"\2\2\u0328\u030f\3\2\2\2\u0328\u0310\3\2\2\2\u0328\u0311\3\2\2\2\u0328"+
		"\u0312\3\2\2\2\u0328\u0313\3\2\2\2\u0328\u0314\3\2\2\2\u0328\u0315\3\2"+
		"\2\2\u0328\u0316\3\2\2\2\u0328\u0317\3\2\2\2\u0328\u0318\3\2\2\2\u0328"+
		"\u0319\3\2\2\2\u0328\u031a\3\2\2\2\u0328\u031b\3\2\2\2\u0328\u031c\3\2"+
		"\2\2\u0328\u031d\3\2\2\2\u0328\u031e\3\2\2\2\u0328\u031f\3\2\2\2\u0328"+
		"\u0320\3\2\2\2\u0328\u0321\3\2\2\2\u0328\u0322\3\2\2\2\u0328\u0323\3\2"+
		"\2\2\u0328\u0324\3\2\2\2\u0328\u0325\3\2\2\2\u0328\u0326\3\2\2\2\u0329"+
		"}\3\2\2\2\u032a\u032c\7\13\2\2\u032b\u032d\5\u0080A\2\u032c\u032b\3\2"+
		"\2\2\u032c\u032d\3\2\2\2\u032d\u032e\3\2\2\2\u032e\u032f\7\f\2\2\u032f"+
		"\177\3\2\2\2\u0330\u0332\5|?\2\u0331\u0330\3\2\2\2\u0332\u0333\3\2\2\2"+
		"\u0333\u0331\3\2\2\2\u0333\u0334\3\2\2\2\u0334\u0081\3\2\2\2\u0335\u0339"+
		"\7\u0082\2\2\u0336\u0337\7\u0085\2\2\u0337\u033a\5> \2\u0338\u033a\5\u008c"+
		"G\2\u0339\u0336\3\2\2\2\u0339\u0338\3\2\2\2\u033a\u033b\3\2\2\2\u033b"+
		"\u033c\5\u0118\u008d\2\u033c\u0083\3\2\2\2\u033d\u0341\7j\2\2\u033e\u0342"+
		"\5\u0086D\2\u033f\u0342\5l\67\2\u0340\u0342\5n8\2\u0341\u033e\3\2\2\2"+
		"\u0341\u033f\3\2\2\2\u0341\u0340\3\2\2\2\u0342\u0343\3\2\2\2\u0343\u0344"+
		"\5\u0118\u008d\2\u0344\u0085\3\2\2\2\u0345\u034b\7\34\2\2\u0346\u034b"+
		"\7\33\2\2\u0347\u034b\7\32\2\2\u0348\u034b\5\u0088E\2\u0349\u034b\5\u010e"+
		"\u0088\2\u034a\u0345\3\2\2\2\u034a\u0346\3\2\2\2\u034a\u0347\3\2\2\2\u034a"+
		"\u0348\3\2\2\2\u034a\u0349\3\2\2\2\u034b\u034e\3\2\2\2\u034c\u034d\7`"+
		"\2\2\u034d\u034f\5\u010e\u0088\2\u034e\u034c\3\2\2\2\u034e\u034f\3\2\2"+
		"\2\u034f\u0350\3\2\2\2\u0350\u0351\7a\2\2\u0351\u0352\7\u0086\2\2\u0352"+
		"\u0087\3\2\2\2\u0353\u0354\5\u010e\u0088\2\u0354\u0355\7\16\2\2\u0355"+
		"\u0357\3\2\2\2\u0356\u0353\3\2\2\2\u0356\u0357\3\2\2\2\u0357\u0358\3\2"+
		"\2\2\u0358\u0359\7\13\2\2\u0359\u035e\5\u010e\u0088\2\u035a\u035b\7\16"+
		"\2\2\u035b\u035d\5\u010e\u0088\2\u035c\u035a\3\2\2\2\u035d\u0360\3\2\2"+
		"\2\u035e\u035c\3\2\2\2\u035e\u035f\3\2\2\2\u035f\u0361\3\2\2\2\u0360\u035e"+
		"\3\2\2\2\u0361\u0362\7\f\2\2\u0362\u0089\3\2\2\2\u0363\u0365\7i\2\2\u0364"+
		"\u0366\7Z\2\2\u0365\u0364\3\2\2\2\u0365\u0366\3\2\2\2\u0366\u0369\3\2"+
		"\2\2\u0367\u036a\5\u0086D\2\u0368\u036a\5|?\2\u0369\u0367\3\2\2\2\u0369"+
		"\u0368\3\2\2\2\u036a\u008b\3\2\2\2\u036b\u036d\5\4\3\2\u036c\u036e\5<"+
		"\37\2\u036d\u036c\3\2\2\2\u036d\u036e\3\2\2\2\u036e\u036f\3\2\2\2\u036f"+
		"\u0371\5\2\2\2\u0370\u0372\7\r\2\2\u0371\u0370\3\2\2\2\u0371\u0372\3\2"+
		"\2\2\u0372\u0381\3\2\2\2\u0373\u0375\5F$\2\u0374\u0373\3\2\2\2\u0374\u0375"+
		"\3\2\2\2\u0375\u0377\3\2\2\2\u0376\u0378\5\u009aN\2\u0377\u0376\3\2\2"+
		"\2\u0377\u0378\3\2\2\2\u0378\u037a\3\2\2\2\u0379\u037b\7b\2\2\u037a\u0379"+
		"\3\2\2\2\u037a\u037b\3\2\2\2\u037b\u037c\3\2\2\2\u037c\u037e\5\u008eH"+
		"\2\u037d\u037f\7\r\2\2\u037e\u037d\3\2\2\2\u037e\u037f\3\2\2\2\u037f\u0381"+
		"\3\2\2\2\u0380\u036b\3\2\2\2\u0380\u0374\3\2\2\2\u0381\u008d\3\2\2\2\u0382"+
		"\u0387\5\u0090I\2\u0383\u0384\7\16\2\2\u0384\u0386\5\u0090I\2\u0385\u0383"+
		"\3\2\2\2\u0386\u0389\3\2\2\2\u0387\u0385\3\2\2\2\u0387\u0388\3\2\2\2\u0388"+
		"\u008f\3\2\2\2\u0389\u0387\3\2\2\2\u038a\u038c\5\u00b8]\2\u038b\u038d"+
		"\5<\37\2\u038c\u038b\3\2\2\2\u038c\u038d\3\2\2\2\u038d\u038f\3\2\2\2\u038e"+
		"\u0390\5\u0100\u0081\2\u038f\u038e\3\2\2\2\u038f\u0390\3\2\2\2\u0390\u0396"+
		"\3\2\2\2\u0391\u0393\7\17\2\2\u0392\u0394\5\6\4\2\u0393\u0392\3\2\2\2"+
		"\u0393\u0394\3\2\2\2\u0394\u0395\3\2\2\2\u0395\u0397\5\u0100\u0081\2\u0396"+
		"\u0391\3\2\2\2\u0396\u0397\3\2\2\2\u0397\u0091\3\2\2\2\u0398\u0399\7\r"+
		"\2\2\u0399\u0093\3\2\2\2\u039a\u039b\6K\b\2\u039b\u039d\5\u00fc\177\2"+
		"\u039c\u039e\7\r\2\2\u039d\u039c\3\2\2\2\u039d\u039e\3\2\2\2\u039e\u0095"+
		"\3\2\2\2\u039f\u03a0\7[\2\2\u03a0\u03a1\7\t\2\2\u03a1\u03a2\5\u00fc\177"+
		"\2\u03a2\u03a3\7\n\2\2\u03a3\u03a6\5|?\2\u03a4\u03a5\7K\2\2\u03a5\u03a7"+
		"\5|?\2\u03a6\u03a4\3\2\2\2\u03a6\u03a7\3\2\2\2\u03a7\u0097\3\2\2\2\u03a8"+
		"\u03a9\7G\2\2\u03a9\u03aa\5|?\2\u03aa\u03ab\7U\2\2\u03ab\u03ac\7\t\2\2"+
		"\u03ac\u03ad\5\u00fc\177\2\u03ad\u03ae\7\n\2\2\u03ae\u03af\5\u0118\u008d"+
		"\2\u03af\u03f4\3\2\2\2\u03b0\u03b1\7U\2\2\u03b1\u03b2\7\t\2\2\u03b2\u03b3"+
		"\5\u00fc\177\2\u03b3\u03b4\7\n\2\2\u03b4\u03b5\5|?\2\u03b5\u03f4\3\2\2"+
		"\2\u03b6\u03b7\7S\2\2\u03b7\u03b9\7\t\2\2\u03b8\u03ba\5\u00fc\177\2\u03b9"+
		"\u03b8\3\2\2\2\u03b9\u03ba\3\2\2\2\u03ba\u03bb\3\2\2\2\u03bb\u03bd\7\r"+
		"\2\2\u03bc\u03be\5\u00fc\177\2\u03bd\u03bc\3\2\2\2\u03bd\u03be\3\2\2\2"+
		"\u03be\u03bf\3\2\2\2\u03bf\u03c1\7\r\2\2\u03c0\u03c2\5\u00fc\177\2\u03c1"+
		"\u03c0\3\2\2\2\u03c1\u03c2\3\2\2\2\u03c2\u03c3\3\2\2\2\u03c3\u03c4\7\n"+
		"\2\2\u03c4\u03f4\5|?\2\u03c5\u03c6\7S\2\2\u03c6\u03c7\7\t\2\2\u03c7\u03c8"+
		"\5\u009aN\2\u03c8\u03c9\5\u008eH\2\u03c9\u03cb\7\r\2\2\u03ca\u03cc\5\u00fc"+
		"\177\2\u03cb\u03ca\3\2\2\2\u03cb\u03cc\3\2\2\2\u03cc\u03cd\3\2\2\2\u03cd"+
		"\u03cf\7\r\2\2\u03ce\u03d0\5\u00fc\177\2\u03cf\u03ce\3\2\2\2\u03cf\u03d0"+
		"\3\2\2\2\u03d0\u03d1\3\2\2\2\u03d1\u03d2\7\n\2\2\u03d2\u03d3\5|?\2\u03d3"+
		"\u03f4\3\2\2\2\u03d4\u03d6\7S\2\2\u03d5\u03d7\7k\2\2\u03d6\u03d5\3\2\2"+
		"\2\u03d6\u03d7\3\2\2\2\u03d7\u03d8\3\2\2\2\u03d8\u03d9\7\t\2\2\u03d9\u03dd"+
		"\5\u0100\u0081\2\u03da\u03de\7^\2\2\u03db\u03dc\7\u0085\2\2\u03dc\u03de"+
		"\6M\t\2\u03dd\u03da\3\2\2\2\u03dd\u03db\3\2\2\2\u03de\u03df\3\2\2\2\u03df"+
		"\u03e0\5\u00fc\177\2\u03e0\u03e1\7\n\2\2\u03e1\u03e2\5|?\2\u03e2\u03f4"+
		"\3\2\2\2\u03e3\u03e5\7S\2\2\u03e4\u03e6\7k\2\2\u03e5\u03e4\3\2\2\2\u03e5"+
		"\u03e6\3\2\2\2\u03e6\u03e7\3\2\2\2\u03e7\u03e8\7\t\2\2\u03e8\u03e9\5\u009a"+
		"N\2\u03e9\u03ed\5\u0090I\2\u03ea\u03ee\7^\2\2\u03eb\u03ec\7\u0085\2\2"+
		"\u03ec\u03ee\6M\n\2\u03ed\u03ea\3\2\2\2\u03ed\u03eb\3\2\2\2\u03ee\u03ef"+
		"\3\2\2\2\u03ef\u03f0\5\u00fc\177\2\u03f0\u03f1\7\n\2\2\u03f1\u03f2\5|"+
		"?\2\u03f2\u03f4\3\2\2\2\u03f3\u03a8\3\2\2\2\u03f3\u03b0\3\2\2\2\u03f3"+
		"\u03b6\3\2\2\2\u03f3\u03c5\3\2\2\2\u03f3\u03d4\3\2\2\2\u03f3\u03e3\3\2"+
		"\2\2\u03f4\u0099\3\2\2\2\u03f5\u03f6\t\6\2\2\u03f6\u009b\3\2\2\2\u03f7"+
		"\u03fa\7R\2\2\u03f8\u03f9\6O\13\2\u03f9\u03fb\7\u0085\2\2\u03fa\u03f8"+
		"\3\2\2\2\u03fa\u03fb\3\2\2\2\u03fb\u03fc\3\2\2\2\u03fc\u03fd\5\u0118\u008d"+
		"\2\u03fd\u009d\3\2\2\2\u03fe\u0401\7F\2\2\u03ff\u0400\6P\f\2\u0400\u0402"+
		"\7\u0085\2\2\u0401\u03ff\3\2\2\2\u0401\u0402\3\2\2\2\u0402\u0403\3\2\2"+
		"\2\u0403\u0404\5\u0118\u008d\2\u0404\u009f\3\2\2\2\u0405\u0408\7P\2\2"+
		"\u0406\u0407\6Q\r\2\u0407\u0409\5\u00fc\177\2\u0408\u0406\3\2\2\2\u0408"+
		"\u0409\3\2\2\2\u0409\u040a\3\2\2\2\u040a\u040b\5\u0118\u008d\2\u040b\u00a1"+
		"\3\2\2\2\u040c\u040f\7t\2\2\u040d\u040e\6R\16\2\u040e\u0410\5\u00fc\177"+
		"\2\u040f\u040d\3\2\2\2\u040f\u0410\3\2\2\2\u0410\u0411\3\2\2\2\u0411\u0412"+
		"\5\u0118\u008d\2\u0412\u00a3\3\2\2\2\u0413\u0414\7Y\2\2\u0414\u0415\7"+
		"\t\2\2\u0415\u0416\5\u00fc\177\2\u0416\u0417\7\n\2\2\u0417\u0418\5|?\2"+
		"\u0418\u00a5\3\2\2\2\u0419\u041a\7T\2\2\u041a\u041b\7\t\2\2\u041b\u041c"+
		"\5\u00fc\177\2\u041c\u041d\7\n\2\2\u041d\u041e\5\u00a8U\2\u041e\u00a7"+
		"\3\2\2\2\u041f\u0421\7\13\2\2\u0420\u0422\5\u00aaV\2\u0421\u0420\3\2\2"+
		"\2\u0421\u0422\3\2\2\2\u0422\u0427\3\2\2\2\u0423\u0425\5\u00aeX\2\u0424"+
		"\u0426\5\u00aaV\2\u0425\u0424\3\2\2\2\u0425\u0426\3\2\2\2\u0426\u0428"+
		"\3\2\2\2\u0427\u0423\3\2\2\2\u0427\u0428\3\2\2\2\u0428\u0429\3\2\2\2\u0429"+
		"\u042a\7\f\2\2\u042a\u00a9\3\2\2\2\u042b\u042d\5\u00acW\2\u042c\u042b"+
		"\3\2\2\2\u042d\u042e\3\2\2\2\u042e\u042c\3\2\2\2\u042e\u042f\3\2\2\2\u042f"+
		"\u00ab\3\2\2\2\u0430\u0431\7J\2\2\u0431\u0432\5\u00fc\177\2\u0432\u0434"+
		"\7\21\2\2\u0433\u0435\5\u0080A\2\u0434\u0433\3\2\2\2\u0434\u0435\3\2\2"+
		"\2\u0435\u00ad\3\2\2\2\u0436\u0437\7Z\2\2\u0437\u0439\7\21\2\2\u0438\u043a"+
		"\5\u0080A\2\u0439\u0438\3\2\2\2\u0439\u043a\3\2\2\2\u043a\u00af\3\2\2"+
		"\2\u043b\u043c\7\u0085\2\2\u043c\u043d\7\21\2\2\u043d\u043e\5|?\2\u043e"+
		"\u00b1\3\2\2\2\u043f\u0440\7\\\2\2\u0440\u0441\6Z\17\2\u0441\u0442\5\u00fc"+
		"\177\2\u0442\u0443\5\u0118\u008d\2\u0443\u00b3\3\2\2\2\u0444\u0445\7_"+
		"\2\2\u0445\u044b\5~@\2\u0446\u0448\5\u00b6\\\2\u0447\u0449\5\u00ba^\2"+
		"\u0448\u0447\3\2\2\2\u0448\u0449\3\2\2\2\u0449\u044c\3\2\2\2\u044a\u044c"+
		"\5\u00ba^\2\u044b\u0446\3\2\2\2\u044b\u044a\3\2\2\2\u044c\u00b5\3\2\2"+
		"\2\u044d\u0453\7N\2\2\u044e\u0450\7\t\2\2\u044f\u0451\5\u00b8]\2\u0450"+
		"\u044f\3\2\2\2\u0450\u0451\3\2\2\2\u0451\u0452\3\2\2\2\u0452\u0454\7\n"+
		"\2\2\u0453\u044e\3\2\2\2\u0453\u0454\3\2\2\2\u0454\u0455\3\2\2\2\u0455"+
		"\u0456\5~@\2\u0456\u00b7\3\2\2\2\u0457\u045b\7\u0085\2\2\u0458\u045b\5"+
		"\u00e8u\2\u0459\u045b\5\u00eex\2\u045a\u0457\3\2\2\2\u045a\u0458\3\2\2"+
		"\2\u045a\u0459\3\2\2\2\u045b\u00b9\3\2\2\2\u045c\u045d\7O\2\2\u045d\u045e"+
		"\5~@\2\u045e\u00bb\3\2\2\2\u045f\u0460\7V\2\2\u0460\u0461\5\u0118\u008d"+
		"\2\u0461\u00bd\3\2\2\2\u0462\u0464\7c\2\2\u0463\u0462\3\2\2\2\u0463\u0464"+
		"\3\2\2\2\u0464\u0465\3\2\2\2\u0465\u0467\7W\2\2\u0466\u0468\7\32\2\2\u0467"+
		"\u0466\3\2\2\2\u0467\u0468\3\2\2\2\u0468\u0469\3\2\2\2\u0469\u046a\7\u0085"+
		"\2\2\u046a\u0470\5> \2\u046b\u046c\7\13\2\2\u046c\u046d\5\u00e4s\2\u046d"+
		"\u046e\7\f\2\2\u046e\u0471\3\2\2\2\u046f\u0471\7\r\2\2\u0470\u046b\3\2"+
		"\2\2\u0470\u046f\3\2\2\2\u0471\u00bf\3\2\2\2\u0472\u0474\7\u0082\2\2\u0473"+
		"\u0472\3\2\2\2\u0473\u0474\3\2\2\2\u0474\u0475\3\2\2\2\u0475\u0476\7d"+
		"\2\2\u0476\u0478\7\u0085\2\2\u0477\u0479\5\6\4\2\u0478\u0477\3\2\2\2\u0478"+
		"\u0479\3\2\2\2\u0479\u047a\3\2\2\2\u047a\u047b\5\u00c2b\2\u047b\u047c"+
		"\5\u00c4c\2\u047c\u00c1\3\2\2\2\u047d\u047f\5\u00c6d\2\u047e\u047d\3\2"+
		"\2\2\u047e\u047f\3\2\2\2\u047f\u0481\3\2\2\2\u0480\u0482\5\u00c8e\2\u0481"+
		"\u0480\3\2\2\2\u0481\u0482\3\2\2\2\u0482\u00c3\3\2\2\2\u0483\u0487\7\13"+
		"\2\2\u0484\u0486\5\u00caf\2\u0485\u0484\3\2\2\2\u0486\u0489\3\2\2\2\u0487"+
		"\u0485\3\2\2\2\u0487\u0488\3\2\2\2\u0488\u048a\3\2\2\2\u0489\u0487\3\2"+
		"\2\2\u048a\u048b\7\f\2\2\u048b\u00c5\3\2\2\2\u048c\u048d\7f\2\2\u048d"+
		"\u048e\5\34\17\2\u048e\u00c7\3\2\2\2\u048f\u0490\7l\2\2\u0490\u0491\5"+
		"^\60\2\u0491\u00c9\3\2\2\2\u0492\u0497\5X-\2\u0493\u0497\5\u00ccg\2\u0494"+
		"\u0497\5\u00d0i\2\u0495\u0497\5|?\2\u0496\u0492\3\2\2\2\u0496\u0493\3"+
		"\2\2\2\u0496\u0494\3\2\2\2\u0496\u0495\3\2\2\2\u0497\u00cb\3\2\2\2\u0498"+
		"\u049a\5\u00ceh\2\u0499\u049b\7\32\2\2\u049a\u0499\3\2\2\2\u049a\u049b"+
		"\3\2\2\2\u049b\u049d\3\2\2\2\u049c\u049e\7!\2\2\u049d\u049c\3\2\2\2\u049d"+
		"\u049e\3\2\2\2\u049e\u049f\3\2\2\2\u049f\u04a1\5\u00f6|\2\u04a0\u04a2"+
		"\5<\37\2\u04a1\u04a0\3\2\2\2\u04a1\u04a2\3\2\2\2\u04a2\u04a4\3\2\2\2\u04a3"+
		"\u04a5\5\2\2\2\u04a4\u04a3\3\2\2\2\u04a4\u04a5\3\2\2\2\u04a5\u04a6\3\2"+
		"\2\2\u04a6\u04a7\7\r\2\2\u04a7\u04c5\3\2\2\2\u04a8\u04aa\5\u00ceh\2\u04a9"+
		"\u04ab\7\32\2\2\u04aa\u04a9\3\2\2\2\u04aa\u04ab\3\2\2\2\u04ab\u04ad\3"+
		"\2\2\2\u04ac\u04ae\7!\2\2\u04ad\u04ac\3\2\2\2\u04ad\u04ae\3\2\2\2\u04ae"+
		"\u04af\3\2\2\2\u04af\u04b0\5\u00f6|\2\u04b0\u04b6\5> \2\u04b1\u04b2\7"+
		"\13\2\2\u04b2\u04b3\5\u00e4s\2\u04b3\u04b4\7\f\2\2\u04b4\u04b7\3\2\2\2"+
		"\u04b5\u04b7\7\r\2\2\u04b6\u04b1\3\2\2\2\u04b6\u04b5\3\2\2\2\u04b7\u04c5"+
		"\3\2\2\2\u04b8\u04ba\5\u00ceh\2\u04b9\u04bb\7\32\2\2\u04ba\u04b9\3\2\2"+
		"\2\u04ba\u04bb\3\2\2\2\u04bb\u04bd\3\2\2\2\u04bc\u04be\7!\2\2\u04bd\u04bc"+
		"\3\2\2\2\u04bd\u04be\3\2\2\2\u04be\u04c1\3\2\2\2\u04bf\u04c2\5\u00f2z"+
		"\2\u04c0\u04c2\5\u00f4{\2\u04c1\u04bf\3\2\2\2\u04c1\u04c0\3\2\2\2\u04c2"+
		"\u04c5\3\2\2\2\u04c3\u04c5\5\u0082B\2\u04c4\u0498\3\2\2\2\u04c4\u04a8"+
		"\3\2\2\2\u04c4\u04b8\3\2\2\2\u04c4\u04c3\3\2\2\2\u04c5\u00cd\3\2\2\2\u04c6"+
		"\u04c8\7c\2\2\u04c7\u04c6\3\2\2\2\u04c7\u04c8\3\2\2\2\u04c8\u04ca\3\2"+
		"\2\2\u04c9\u04cb\5F$\2\u04ca\u04c9\3\2\2\2\u04ca\u04cb\3\2\2\2\u04cb\u04cd"+
		"\3\2\2\2\u04cc\u04ce\7s\2\2\u04cd\u04cc\3\2\2\2\u04cd\u04ce\3\2\2\2\u04ce"+
		"\u04d0\3\2\2\2\u04cf\u04d1\7b\2\2\u04d0\u04cf\3\2\2\2\u04d0\u04d1\3\2"+
		"\2\2\u04d1\u00cf\3\2\2\2\u04d2\u04d3\5R*\2\u04d3\u04d4\7\r\2\2\u04d4\u00d1"+
		"\3\2\2\2\u04d5\u04d7\7\32\2\2\u04d6\u04d5\3\2\2\2\u04d6\u04d7\3\2\2\2"+
		"\u04d7\u04d8\3\2\2\2\u04d8\u04d9\7\u0085\2\2\u04d9\u04db\7\t\2\2\u04da"+
		"\u04dc\5\u00dep\2\u04db\u04da\3\2\2\2\u04db\u04dc\3\2\2\2\u04dc\u04dd"+
		"\3\2\2\2\u04dd\u04de\7\n\2\2\u04de\u04df\7\13\2\2\u04df\u04e0\5\u00e4"+
		"s\2\u04e0\u04e1\7\f\2\2\u04e1\u00d3\3\2\2\2\u04e2\u04e3\7W\2\2\u04e3\u04e5"+
		"\7\32\2\2\u04e4\u04e6\7\u0085\2\2\u04e5\u04e4\3\2\2\2\u04e5\u04e6\3\2"+
		"\2\2\u04e6\u04e7\3\2\2\2\u04e7\u04e9\7\t\2\2\u04e8\u04ea\5\u00dep\2\u04e9"+
		"\u04e8\3\2\2\2\u04e9\u04ea\3\2\2\2\u04ea\u04eb\3\2\2\2\u04eb\u04ec\7\n"+
		"\2\2\u04ec\u04ed\7\13\2\2\u04ed\u04ee\5\u00e4s\2\u04ee\u04ef\7\f\2\2\u04ef"+
		"\u00d5\3\2\2\2\u04f0\u04f1\7\13\2\2\u04f1\u04f6\5\u00d8m\2\u04f2\u04f3"+
		"\7\16\2\2\u04f3\u04f5\5\u00d8m\2\u04f4\u04f2\3\2\2\2\u04f5\u04f8\3\2\2"+
		"\2\u04f6\u04f4\3\2\2\2\u04f6\u04f7\3\2\2\2\u04f7\u04fa\3\2\2\2\u04f8\u04f6"+
		"\3\2\2\2\u04f9\u04fb\7\16\2\2\u04fa\u04f9\3\2\2\2\u04fa\u04fb\3\2\2\2"+
		"\u04fb\u04fc\3\2\2\2\u04fc\u04fd\7\f\2\2\u04fd\u00d7\3\2\2\2\u04fe\u04ff"+
		"\7\32\2\2\u04ff\u0500\5\u00dco\2\u0500\u00d9\3\2\2\2\u0501\u0502\7\13"+
		"\2\2\u0502\u0507\5\u00dco\2\u0503\u0504\7\16\2\2\u0504\u0506\5\u00dco"+
		"\2\u0505\u0503\3\2\2\2\u0506\u0509\3\2\2\2\u0507\u0505\3\2\2\2\u0507\u0508"+
		"\3\2\2\2\u0508\u050b\3\2\2\2\u0509\u0507\3\2\2\2\u050a\u050c\7\16\2\2"+
		"\u050b\u050a\3\2\2\2\u050b\u050c\3\2\2\2\u050c\u050d\3\2\2\2\u050d\u050e"+
		"\7\f\2\2\u050e\u00db\3\2\2\2\u050f\u0510\7\7\2\2\u0510\u0511\5\u0100\u0081"+
		"\2\u0511\u0512\7\b\2\2\u0512\u0514\7\t\2\2\u0513\u0515\5\u00dep\2\u0514"+
		"\u0513\3\2\2\2\u0514\u0515\3\2\2\2\u0515\u0516\3\2\2\2\u0516\u0517\7\n"+
		"\2\2\u0517\u0518\7\13\2\2\u0518\u0519\5\u00e4s\2\u0519\u051a\7\f\2\2\u051a"+
		"\u00dd\3\2\2\2\u051b\u0520\5\u00e0q\2\u051c\u051d\7\16\2\2\u051d\u051f"+
		"\5\u00e0q\2\u051e\u051c\3\2\2\2\u051f\u0522\3\2\2\2\u0520\u051e\3\2\2"+
		"\2\u0520\u0521\3\2\2\2\u0521\u0525\3\2\2\2\u0522\u0520\3\2\2\2\u0523\u0524"+
		"\7\16\2\2\u0524\u0526\5\u00e2r\2\u0525\u0523\3\2\2\2\u0525\u0526\3\2\2"+
		"\2\u0526\u052f\3\2\2\2\u0527\u052f\5\u00e2r\2\u0528\u052f\5\u00e8u\2\u0529"+
		"\u052c\5\u00eex\2\u052a\u052b\7\21\2\2\u052b\u052d\5\u00dep\2\u052c\u052a"+
		"\3\2\2\2\u052c\u052d\3\2\2\2\u052d\u052f\3\2\2\2\u052e\u051b\3\2\2\2\u052e"+
		"\u0527\3\2\2\2\u052e\u0528\3\2\2\2\u052e\u0529\3\2\2\2\u052f\u00df\3\2"+
		"\2\2\u0530\u0532\5F$\2\u0531\u0530\3\2\2\2\u0531\u0532\3\2\2\2\u0532\u0533"+
		"\3\2\2\2\u0533\u0535\7\u0085\2\2\u0534\u0536\5<\37\2\u0535\u0534\3\2\2"+
		"\2\u0535\u0536\3\2\2\2\u0536\u0539\3\2\2\2\u0537\u0538\7\17\2\2\u0538"+
		"\u053a\5\u0100\u0081\2\u0539\u0537\3\2\2\2\u0539\u053a\3\2\2\2\u053a\u00e1"+
		"\3\2\2\2\u053b\u053c\7\22\2\2\u053c\u053d\7\u0085\2\2\u053d\u00e3\3\2"+
		"\2\2\u053e\u0540\5\u00e6t\2\u053f\u053e\3\2\2\2\u053f\u0540\3\2\2\2\u0540"+
		"\u00e5\3\2\2\2\u0541\u0543\5z>\2\u0542\u0541\3\2\2\2\u0543\u0544\3\2\2"+
		"\2\u0544\u0542\3\2\2\2\u0544\u0545\3\2\2\2\u0545\u00e7\3\2\2\2\u0546\u0548"+
		"\7\7\2\2\u0547\u0549\5\u00eav\2\u0548\u0547\3\2\2\2\u0548\u0549\3\2\2"+
		"\2\u0549\u054a\3\2\2\2\u054a\u054b\7\b\2\2\u054b\u00e9\3\2\2\2\u054c\u0555"+
		"\5\u0100\u0081\2\u054d\u054f\7\16\2\2\u054e\u054d\3\2\2\2\u054f\u0550"+
		"\3\2\2\2\u0550\u054e\3\2\2\2\u0550\u0551\3\2\2\2\u0551\u0552\3\2\2\2\u0552"+
		"\u0554\5\u0100\u0081\2\u0553\u054e\3\2\2\2\u0554\u0557\3\2\2\2\u0555\u0553"+
		"\3\2\2\2\u0555\u0556\3\2\2\2\u0556\u055e\3\2\2\2\u0557\u0555\3\2\2\2\u0558"+
		"\u055a\7\16\2\2\u0559\u0558\3\2\2\2\u055a\u055b\3\2\2\2\u055b\u0559\3"+
		"\2\2\2\u055b\u055c\3\2\2\2\u055c\u055d\3\2\2\2\u055d\u055f\5\u00ecw\2"+
		"\u055e\u0559\3\2\2\2\u055e\u055f\3\2\2\2\u055f\u0562\3\2\2\2\u0560\u0562"+
		"\5\u00ecw\2\u0561\u054c\3\2\2\2\u0561\u0560\3\2\2\2\u0562\u00eb\3\2\2"+
		"\2\u0563\u0566\7\22\2\2\u0564\u0567\7\u0085\2\2\u0565\u0567\5\u0100\u0081"+
		"\2\u0566\u0564\3\2\2\2\u0566\u0565\3\2\2\2\u0567\u00ed\3\2\2\2\u0568\u0571"+
		"\7\13\2\2\u0569\u056e\5\u00f0y\2\u056a\u056b\7\16\2\2\u056b\u056d\5\u00f0"+
		"y\2\u056c\u056a\3\2\2\2\u056d\u0570\3\2\2\2\u056e\u056c\3\2\2\2\u056e"+
		"\u056f\3\2\2\2\u056f\u0572\3\2\2\2\u0570\u056e\3\2\2\2\u0571\u0569\3\2"+
		"\2\2\u0571\u0572\3\2\2\2\u0572\u0574\3\2\2\2\u0573\u0575\7\16\2\2\u0574"+
		"\u0573\3\2\2\2\u0574\u0575\3\2\2\2\u0575\u0576\3\2\2\2\u0576\u0577\7\f"+
		"\2\2\u0577\u00ef\3\2\2\2\u0578\u0579\5\u00f6|\2\u0579\u057a\t\7\2\2\u057a"+
		"\u057b\5\u0100\u0081\2\u057b\u0598\3\2\2\2\u057c\u057d\7\7\2\2\u057d\u057e"+
		"\5\u0100\u0081\2\u057e\u057f\7\b\2\2\u057f\u0580\7\21\2\2\u0580\u0581"+
		"\5\u0100\u0081\2\u0581\u0598\3\2\2\2\u0582\u0584\7c\2\2\u0583\u0582\3"+
		"\2\2\2\u0583\u0584\3\2\2\2\u0584\u0586\3\2\2\2\u0585\u0587\7\32\2\2\u0586"+
		"\u0585\3\2\2\2\u0586\u0587\3\2\2\2\u0587\u0588\3\2\2\2\u0588\u0589\5\u00f6"+
		"|\2\u0589\u058b\7\t\2\2\u058a\u058c\5\u00dep\2\u058b\u058a\3\2\2\2\u058b"+
		"\u058c\3\2\2\2\u058c\u058d\3\2\2\2\u058d\u058e\7\n\2\2\u058e\u058f\7\13"+
		"\2\2\u058f\u0590\5\u00e4s\2\u0590\u0591\7\f\2\2\u0591\u0598\3\2\2\2\u0592"+
		"\u0598\5\u00f2z\2\u0593\u0598\5\u00f4{\2\u0594\u0598\5\u00d2j\2\u0595"+
		"\u0598\7\u0085\2\2\u0596\u0598\5N(\2\u0597\u0578\3\2\2\2\u0597\u057c\3"+
		"\2\2\2\u0597\u0583\3\2\2\2\u0597\u0592\3\2\2\2\u0597\u0593\3\2\2\2\u0597"+
		"\u0594\3\2\2\2\u0597\u0595\3\2\2\2\u0597\u0596\3\2\2\2\u0598\u00f1\3\2"+
		"\2\2\u0599\u059a\5\u0114\u008b\2\u059a\u059b\7\t\2\2\u059b\u059d\7\n\2"+
		"\2\u059c\u059e\5<\37\2\u059d\u059c\3\2\2\2\u059d\u059e\3\2\2\2\u059e\u059f"+
		"\3\2\2\2\u059f\u05a0\7\13\2\2\u05a0\u05a1\5\u00e4s\2\u05a1\u05a2\7\f\2"+
		"\2\u05a2\u00f3\3\2\2\2\u05a3\u05a4\5\u0116\u008c\2\u05a4\u05a7\7\t\2\2"+
		"\u05a5\u05a8\7\u0085\2\2\u05a6\u05a8\5\4\3\2\u05a7\u05a5\3\2\2\2\u05a7"+
		"\u05a6\3\2\2\2\u05a8\u05aa\3\2\2\2\u05a9\u05ab\5<\37\2\u05aa\u05a9\3\2"+
		"\2\2\u05aa\u05ab\3\2\2\2\u05ab\u05ac\3\2\2\2\u05ac\u05ad\7\n\2\2\u05ad"+
		"\u05ae\7\13\2\2\u05ae\u05af\5\u00e4s\2\u05af\u05b0\7\f\2\2\u05b0\u00f5"+
		"\3\2\2\2\u05b1\u05b9\5\u010e\u0088\2\u05b2\u05b9\7\u0086\2\2\u05b3\u05b9"+
		"\5\u010c\u0087\2\u05b4\u05b5\7\7\2\2\u05b5\u05b6\5\u0100\u0081\2\u05b6"+
		"\u05b7\7\b\2\2\u05b7\u05b9\3\2\2\2\u05b8\u05b1\3\2\2\2\u05b8\u05b2\3\2"+
		"\2\2\u05b8\u05b3\3\2\2\2\u05b8\u05b4\3\2\2\2\u05b9\u00f7\3\2\2\2\u05ba"+
		"\u05c8\7\t\2\2\u05bb\u05c0\5\u0100\u0081\2\u05bc\u05bd\7\16\2\2\u05bd"+
		"\u05bf\5\u0100\u0081\2\u05be\u05bc\3\2\2\2\u05bf\u05c2\3\2\2\2\u05c0\u05be"+
		"\3\2\2\2\u05c0\u05c1\3\2\2\2\u05c1\u05c5\3\2\2\2\u05c2\u05c0\3\2\2\2\u05c3"+
		"\u05c4\7\16\2\2\u05c4\u05c6\5\u00fa~\2\u05c5\u05c3\3\2\2\2\u05c5\u05c6"+
		"\3\2\2\2\u05c6\u05c9\3\2\2\2\u05c7\u05c9\5\u00fa~\2\u05c8\u05bb\3\2\2"+
		"\2\u05c8\u05c7\3\2\2\2\u05c8\u05c9\3\2\2\2\u05c9\u05ca\3\2\2\2\u05ca\u05cb"+
		"\7\n\2\2\u05cb\u00f9\3\2\2\2\u05cc\u05cd\7\22\2\2\u05cd\u05ce\7\u0085"+
		"\2\2\u05ce\u00fb\3\2\2\2\u05cf\u05d4\5\u0100\u0081\2\u05d0\u05d1\7\16"+
		"\2\2\u05d1\u05d3\5\u0100\u0081\2\u05d2\u05d0\3\2\2\2\u05d3\u05d6\3\2\2"+
		"\2\u05d4\u05d2\3\2\2\2\u05d4\u05d5\3\2\2\2\u05d5\u05d8\3\2\2\2\u05d6\u05d4"+
		"\3\2\2\2\u05d7\u05d9\7\16\2\2\u05d8\u05d7\3\2\2\2\u05d8\u05d9\3\2\2\2"+
		"\u05d9\u00fd\3\2\2\2\u05da\u05dc\7c\2\2\u05db\u05da\3\2\2\2\u05db\u05dc"+
		"\3\2\2\2\u05dc\u05dd\3\2\2\2\u05dd\u05df\7W\2\2\u05de\u05e0\7\u0085\2"+
		"\2\u05df\u05de\3\2\2\2\u05df\u05e0\3\2\2\2\u05e0\u05e1\3\2\2\2\u05e1\u05e3"+
		"\7\t\2\2\u05e2\u05e4\5\u00dep\2\u05e3\u05e2\3\2\2\2\u05e3\u05e4\3\2\2"+
		"\2\u05e4\u05e5\3\2\2\2\u05e5\u05e7\7\n\2\2\u05e6\u05e8\5<\37\2\u05e7\u05e6"+
		"\3\2\2\2\u05e7\u05e8\3\2\2\2\u05e8\u05e9\3\2\2\2\u05e9\u05ea\7\13\2\2"+
		"\u05ea\u05eb\5\u00e4s\2\u05eb\u05ec\7\f\2\2\u05ec\u00ff\3\2\2\2\u05ed"+
		"\u05ee\b\u0081\1\2\u05ee\u0627\5\u00fe\u0080\2\u05ef\u0627\5\u0102\u0082"+
		"\2\u05f0\u05f2\7d\2\2\u05f1\u05f3\7\u0085\2\2\u05f2\u05f1\3\2\2\2\u05f2"+
		"\u05f3\3\2\2\2\u05f3\u05f4\3\2\2\2\u05f4\u0627\5\u00c4c\2\u05f5\u05f6"+
		"\7L\2\2\u05f6\u05f8\5\u0100\u0081\2\u05f7\u05f9\5\16\b\2\u05f8\u05f7\3"+
		"\2\2\2\u05f8\u05f9\3\2\2\2\u05f9\u05fb\3\2\2\2\u05fa\u05fc\5\u00f8}\2"+
		"\u05fb\u05fa\3\2\2\2\u05fb\u05fc\3\2\2\2\u05fc\u0627\3\2\2\2\u05fd\u05fe"+
		"\7]\2\2\u05fe\u0627\5\u0100\u0081*\u05ff\u0600\7Q\2\2\u0600\u0627\5\u0100"+
		"\u0081)\u0601\u0602\7I\2\2\u0602\u0627\5\u0100\u0081(\u0603\u0604\7\24"+
		"\2\2\u0604\u0627\5\u0100\u0081\'\u0605\u0606\7\25\2\2\u0606\u0627\5\u0100"+
		"\u0081&\u0607\u0608\7\26\2\2\u0608\u0627\5\u0100\u0081%\u0609\u060a\7"+
		"\27\2\2\u060a\u0627\5\u0100\u0081$\u060b\u060c\7\30\2\2\u060c\u0627\5"+
		"\u0100\u0081#\u060d\u060e\7\31\2\2\u060e\u0627\5\u0100\u0081\"\u060f\u0610"+
		"\7k\2\2\u0610\u0627\5\u0100\u0081!\u0611\u0627\5\u00dan\2\u0612\u0627"+
		"\5\u00d6l\2\u0613\u0627\5\u00d4k\2\u0614\u0627\5\u00a2R\2\u0615\u0627"+
		"\7X\2\2\u0616\u0618\5\u010e\u0088\2\u0617\u0619\5\u0100\u0081\2\u0618"+
		"\u0617\3\2\2\2\u0618\u0619\3\2\2\2\u0619\u0627\3\2\2\2\u061a\u0627\7g"+
		"\2\2\u061b\u0627\5\u010a\u0086\2\u061c\u0627\5\u00e8u\2\u061d\u0627\5"+
		"\u00eex\2\u061e\u061f\7\t\2\2\u061f\u0620\5\u00fc\177\2\u0620\u0621\7"+
		"\n\2\2\u0621\u0627\3\2\2\2\u0622\u0624\5\16\b\2\u0623\u0625\5\u00fc\177"+
		"\2\u0624\u0623\3\2\2\2\u0624\u0625\3\2\2\2\u0625\u0627\3\2\2\2\u0626\u05ed"+
		"\3\2\2\2\u0626\u05ef\3\2\2\2\u0626\u05f0\3\2\2\2\u0626\u05f5\3\2\2\2\u0626"+
		"\u05fd\3\2\2\2\u0626\u05ff\3\2\2\2\u0626\u0601\3\2\2\2\u0626\u0603\3\2"+
		"\2\2\u0626\u0605\3\2\2\2\u0626\u0607\3\2\2\2\u0626\u0609\3\2\2\2\u0626"+
		"\u060b\3\2\2\2\u0626\u060d\3\2\2\2\u0626\u060f\3\2\2\2\u0626\u0611\3\2"+
		"\2\2\u0626\u0612\3\2\2\2\u0626\u0613\3\2\2\2\u0626\u0614\3\2\2\2\u0626"+
		"\u0615\3\2\2\2\u0626\u0616\3\2\2\2\u0626\u061a\3\2\2\2\u0626\u061b\3\2"+
		"\2\2\u0626\u061c\3\2\2\2\u0626\u061d\3\2\2\2\u0626\u061e\3\2\2\2\u0626"+
		"\u0622\3\2\2\2\u0627\u0679\3\2\2\2\u0628\u0629\f \2\2\u0629\u062a\7\37"+
		"\2\2\u062a\u0678\5\u0100\u0081 \u062b\u062c\f\37\2\2\u062c\u062d\t\b\2"+
		"\2\u062d\u0678\5\u0100\u0081 \u062e\u062f\f\36\2\2\u062f\u0630\t\t\2\2"+
		"\u0630\u0678\5\u0100\u0081\37\u0631\u0632\f\35\2\2\u0632\u0633\t\n\2\2"+
		"\u0633\u0678\5\u0100\u0081\36\u0634\u0635\f\34\2\2\u0635\u0636\t\13\2"+
		"\2\u0636\u0678\5\u0100\u0081\35\u0637\u0638\f\33\2\2\u0638\u0639\7H\2"+
		"\2\u0639\u0678\5\u0100\u0081\34\u063a\u063b\f\32\2\2\u063b\u063c\7^\2"+
		"\2\u063c\u0678\5\u0100\u0081\33\u063d\u063e\f\31\2\2\u063e\u063f\t\f\2"+
		"\2\u063f\u0678\5\u0100\u0081\32\u0640\u0641\f\30\2\2\u0641\u0642\7-\2"+
		"\2\u0642\u0678\5\u0100\u0081\31\u0643\u0644\f\27\2\2\u0644\u0645\7.\2"+
		"\2\u0645\u0678\5\u0100\u0081\30\u0646\u0647\f\26\2\2\u0647\u0648\7/\2"+
		"\2\u0648\u0678\5\u0100\u0081\27\u0649\u064a\f\25\2\2\u064a\u064b\7\60"+
		"\2\2\u064b\u0678\5\u0100\u0081\26\u064c\u064d\f\24\2\2\u064d\u064e\7\61"+
		"\2\2\u064e\u0678\5\u0100\u0081\25\u064f\u0650\f\23\2\2\u0650\u0651\7\20"+
		"\2\2\u0651\u0652\5\u0100\u0081\2\u0652\u0653\7\21\2\2\u0653\u0654\5\u0100"+
		"\u0081\24\u0654\u0678\3\2\2\2\u0655\u0656\f\22\2\2\u0656\u0657\7 \2\2"+
		"\u0657\u0678\5\u0100\u0081\23\u0658\u0659\f\21\2\2\u0659\u065a\7\17\2"+
		"\2\u065a\u0678\5\u0100\u0081\21\u065b\u065c\f\20\2\2\u065c\u065d\5\u0108"+
		"\u0085\2\u065d\u065e\5\u0100\u0081\20\u065e\u0678\3\2\2\2\u065f\u0660"+
		"\f\60\2\2\u0660\u0661\7\7\2\2\u0661\u0662\5\u00fc\177\2\u0662\u0663\7"+
		"\b\2\2\u0663\u0678\3\2\2\2\u0664\u0666\f/\2\2\u0665\u0667\7\20\2\2\u0666"+
		"\u0665\3\2\2\2\u0666\u0667\3\2\2\2\u0667\u0668\3\2\2\2\u0668\u066a\7\23"+
		"\2\2\u0669\u066b\7!\2\2\u066a\u0669\3\2\2\2\u066a\u066b\3\2\2\2\u066b"+
		"\u066c\3\2\2\2\u066c\u0678\5\u010e\u0088\2\u066d\u066e\f.\2\2\u066e\u0678"+
		"\5\u00f8}\2\u066f\u0670\f,\2\2\u0670\u0671\6\u0081%\2\u0671\u0678\7\24"+
		"\2\2\u0672\u0673\f+\2\2\u0673\u0674\6\u0081\'\2\u0674\u0678\7\25\2\2\u0675"+
		"\u0676\f\17\2\2\u0676\u0678\7\u0087\2\2\u0677\u0628\3\2\2\2\u0677\u062b"+
		"\3\2\2\2\u0677\u062e\3\2\2\2\u0677\u0631\3\2\2\2\u0677\u0634\3\2\2\2\u0677"+
		"\u0637\3\2\2\2\u0677\u063a\3\2\2\2\u0677\u063d\3\2\2\2\u0677\u0640\3\2"+
		"\2\2\u0677\u0643\3\2\2\2\u0677\u0646\3\2\2\2\u0677\u0649\3\2\2\2\u0677"+
		"\u064c\3\2\2\2\u0677\u064f\3\2\2\2\u0677\u0655\3\2\2\2\u0677\u0658\3\2"+
		"\2\2\u0677\u065b\3\2\2\2\u0677\u065f\3\2\2\2\u0677\u0664\3\2\2\2\u0677"+
		"\u066d\3\2\2\2\u0677\u066f\3\2\2\2\u0677\u0672\3\2\2\2\u0677\u0675\3\2"+
		"\2\2\u0678\u067b\3\2\2\2\u0679\u0677\3\2\2\2\u0679\u067a\3\2\2\2\u067a"+
		"\u0101\3\2\2\2\u067b\u0679\3\2\2\2\u067c\u067e\7c\2\2\u067d\u067c\3\2"+
		"\2\2\u067d\u067e\3\2\2\2\u067e\u067f\3\2\2\2\u067f\u0681\5\u0104\u0083"+
		"\2\u0680\u0682\5<\37\2\u0681\u0680\3\2\2\2\u0681\u0682\3\2\2\2\u0682\u0683"+
		"\3\2\2\2\u0683\u0684\7=\2\2\u0684\u0685\5\u0106\u0084\2\u0685\u0103\3"+
		"\2\2\2\u0686\u068d\7\u0085\2\2\u0687\u0689\7\t\2\2\u0688\u068a\5\u00de"+
		"p\2\u0689\u0688\3\2\2\2\u0689\u068a\3\2\2\2\u068a\u068b\3\2\2\2\u068b"+
		"\u068d\7\n\2\2\u068c\u0686\3\2\2\2\u068c\u0687\3\2\2\2\u068d\u0105\3\2"+
		"\2\2\u068e\u0694\5\u0100\u0081\2\u068f\u0690\7\13\2\2\u0690\u0691\5\u00e4"+
		"s\2\u0691\u0692\7\f\2\2\u0692\u0694\3\2\2\2\u0693\u068e\3\2\2\2\u0693"+
		"\u068f\3\2\2\2\u0694\u0107\3\2\2\2\u0695\u0696\t\r\2\2\u0696\u0109\3\2"+
		"\2\2\u0697\u069e\7?\2\2\u0698\u069e\7@\2\2\u0699\u069e\7\u0086\2\2\u069a"+
		"\u069e\7\u0087\2\2\u069b\u069e\7\6\2\2\u069c\u069e\5\u010c\u0087\2\u069d"+
		"\u0697\3\2\2\2\u069d\u0698\3\2\2\2\u069d\u0699\3\2\2\2\u069d\u069a\3\2"+
		"\2\2\u069d\u069b\3\2\2\2\u069d\u069c\3\2\2\2\u069e\u010b\3\2\2\2\u069f"+
		"\u06a0\t\16\2\2\u06a0\u010d\3\2\2\2\u06a1\u06a4\7\u0085\2\2\u06a2\u06a4"+
		"\5\u0110\u0089\2\u06a3\u06a1\3\2\2\2\u06a3\u06a2\3\2\2\2\u06a4\u010f\3"+
		"\2\2\2\u06a5\u06a9\5\u0112\u008a\2\u06a6\u06a9\7?\2\2\u06a7\u06a9\7@\2"+
		"\2\u06a8\u06a5\3\2\2\2\u06a8\u06a6\3\2\2\2\u06a8\u06a7\3\2\2\2\u06a9\u0111"+
		"\3\2\2\2\u06aa\u06ab\t\17\2\2\u06ab\u0113\3\2\2\2\u06ac\u06ad\7\u0085"+
		"\2\2\u06ad\u06ae\6\u008b)\2\u06ae\u06af\5\u00f6|\2\u06af\u0115\3\2\2\2"+
		"\u06b0\u06b1\7\u0085\2\2\u06b1\u06b2\6\u008c*\2\u06b2\u06b3\5\u00f6|\2"+
		"\u06b3\u0117\3\2\2\2\u06b4\u06b9\7\r\2\2\u06b5\u06b9\7\2\2\3\u06b6\u06b9"+
		"\6\u008d+\2\u06b7\u06b9\6\u008d,\2\u06b8\u06b4\3\2\2\2\u06b8\u06b5\3\2"+
		"\2\2\u06b8\u06b6\3\2\2\2\u06b8\u06b7\3\2\2\2\u06b9\u0119\3\2\2\2\u00dc"+
		"\u011f\u0123\u012c\u0131\u0134\u013b\u0144\u014e\u0159\u015b\u0170\u0178"+
		"\u0180\u018f\u0193\u0197\u019d\u01a4\u01ae\u01b0\u01c0\u01c4\u01c8\u01d0"+
		"\u01d4\u01e3\u01e7\u01ea\u01ee\u01f1\u01f5\u01fb\u01ff\u0203\u020b\u0211"+
		"\u0218\u021b\u021d\u021f\u0226\u022a\u022d\u0231\u0237\u023e\u0242\u0245"+
		"\u024a\u024d\u0250\u0256\u025a\u025e\u0262\u026d\u0274\u027b\u0280\u0288"+
		"\u028b\u0290\u0293\u0297\u02a1\u02a5\u02ab\u02b1\u02b8\u02be\u02c4\u02cc"+
		"\u02d1\u02e0\u02e7\u02ec\u02f4\u02fb\u0302\u0307\u0328\u032c\u0333\u0339"+
		"\u0341\u034a\u034e\u0356\u035e\u0365\u0369\u036d\u0371\u0374\u0377\u037a"+
		"\u037e\u0380\u0387\u038c\u038f\u0393\u0396\u039d\u03a6\u03b9\u03bd\u03c1"+
		"\u03cb\u03cf\u03d6\u03dd\u03e5\u03ed\u03f3\u03fa\u0401\u0408\u040f\u0421"+
		"\u0425\u0427\u042e\u0434\u0439\u0448\u044b\u0450\u0453\u045a\u0463\u0467"+
		"\u0470\u0473\u0478\u047e\u0481\u0487\u0496\u049a\u049d\u04a1\u04a4\u04aa"+
		"\u04ad\u04b6\u04ba\u04bd\u04c1\u04c4\u04c7\u04ca\u04cd\u04d0\u04d6\u04db"+
		"\u04e5\u04e9\u04f6\u04fa\u0507\u050b\u0514\u0520\u0525\u052c\u052e\u0531"+
		"\u0535\u0539\u053f\u0544\u0548\u0550\u0555\u055b\u055e\u0561\u0566\u056e"+
		"\u0571\u0574\u0583\u0586\u058b\u0597\u059d\u05a7\u05aa\u05b8\u05c0\u05c5"+
		"\u05c8\u05d4\u05d8\u05db\u05df\u05e3\u05e7\u05f2\u05f8\u05fb\u0618\u0624"+
		"\u0626\u0666\u066a\u0677\u0679\u067d\u0681\u0689\u068c\u0693\u069d\u06a3"+
		"\u06a8\u06b8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
