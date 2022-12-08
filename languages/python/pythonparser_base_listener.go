// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // PythonParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BasePythonParserListener is a complete listener for a parse tree produced by PythonParser.
type BasePythonParserListener struct{}

var _ PythonParserListener = &BasePythonParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePythonParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePythonParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePythonParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePythonParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BasePythonParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BasePythonParserListener) ExitRoot(ctx *RootContext) {}

// EnterSingle_input is called when production single_input is entered.
func (s *BasePythonParserListener) EnterSingle_input(ctx *Single_inputContext) {}

// ExitSingle_input is called when production single_input is exited.
func (s *BasePythonParserListener) ExitSingle_input(ctx *Single_inputContext) {}

// EnterFile_input is called when production file_input is entered.
func (s *BasePythonParserListener) EnterFile_input(ctx *File_inputContext) {}

// ExitFile_input is called when production file_input is exited.
func (s *BasePythonParserListener) ExitFile_input(ctx *File_inputContext) {}

// EnterEval_input is called when production eval_input is entered.
func (s *BasePythonParserListener) EnterEval_input(ctx *Eval_inputContext) {}

// ExitEval_input is called when production eval_input is exited.
func (s *BasePythonParserListener) ExitEval_input(ctx *Eval_inputContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BasePythonParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BasePythonParserListener) ExitStmt(ctx *StmtContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BasePythonParserListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BasePythonParserListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BasePythonParserListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BasePythonParserListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BasePythonParserListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BasePythonParserListener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterTry_stmt is called when production try_stmt is entered.
func (s *BasePythonParserListener) EnterTry_stmt(ctx *Try_stmtContext) {}

// ExitTry_stmt is called when production try_stmt is exited.
func (s *BasePythonParserListener) ExitTry_stmt(ctx *Try_stmtContext) {}

// EnterWith_stmt is called when production with_stmt is entered.
func (s *BasePythonParserListener) EnterWith_stmt(ctx *With_stmtContext) {}

// ExitWith_stmt is called when production with_stmt is exited.
func (s *BasePythonParserListener) ExitWith_stmt(ctx *With_stmtContext) {}

// EnterClass_or_func_def_stmt is called when production class_or_func_def_stmt is entered.
func (s *BasePythonParserListener) EnterClass_or_func_def_stmt(ctx *Class_or_func_def_stmtContext) {}

// ExitClass_or_func_def_stmt is called when production class_or_func_def_stmt is exited.
func (s *BasePythonParserListener) ExitClass_or_func_def_stmt(ctx *Class_or_func_def_stmtContext) {}

// EnterSuite is called when production suite is entered.
func (s *BasePythonParserListener) EnterSuite(ctx *SuiteContext) {}

// ExitSuite is called when production suite is exited.
func (s *BasePythonParserListener) ExitSuite(ctx *SuiteContext) {}

// EnterDecorator is called when production decorator is entered.
func (s *BasePythonParserListener) EnterDecorator(ctx *DecoratorContext) {}

// ExitDecorator is called when production decorator is exited.
func (s *BasePythonParserListener) ExitDecorator(ctx *DecoratorContext) {}

// EnterElif_clause is called when production elif_clause is entered.
func (s *BasePythonParserListener) EnterElif_clause(ctx *Elif_clauseContext) {}

// ExitElif_clause is called when production elif_clause is exited.
func (s *BasePythonParserListener) ExitElif_clause(ctx *Elif_clauseContext) {}

// EnterElse_clause is called when production else_clause is entered.
func (s *BasePythonParserListener) EnterElse_clause(ctx *Else_clauseContext) {}

// ExitElse_clause is called when production else_clause is exited.
func (s *BasePythonParserListener) ExitElse_clause(ctx *Else_clauseContext) {}

// EnterFinally_clause is called when production finally_clause is entered.
func (s *BasePythonParserListener) EnterFinally_clause(ctx *Finally_clauseContext) {}

// ExitFinally_clause is called when production finally_clause is exited.
func (s *BasePythonParserListener) ExitFinally_clause(ctx *Finally_clauseContext) {}

// EnterWith_item is called when production with_item is entered.
func (s *BasePythonParserListener) EnterWith_item(ctx *With_itemContext) {}

// ExitWith_item is called when production with_item is exited.
func (s *BasePythonParserListener) ExitWith_item(ctx *With_itemContext) {}

// EnterExcept_clause is called when production except_clause is entered.
func (s *BasePythonParserListener) EnterExcept_clause(ctx *Except_clauseContext) {}

// ExitExcept_clause is called when production except_clause is exited.
func (s *BasePythonParserListener) ExitExcept_clause(ctx *Except_clauseContext) {}

// EnterClassdef is called when production classdef is entered.
func (s *BasePythonParserListener) EnterClassdef(ctx *ClassdefContext) {}

// ExitClassdef is called when production classdef is exited.
func (s *BasePythonParserListener) ExitClassdef(ctx *ClassdefContext) {}

// EnterFuncdef is called when production funcdef is entered.
func (s *BasePythonParserListener) EnterFuncdef(ctx *FuncdefContext) {}

// ExitFuncdef is called when production funcdef is exited.
func (s *BasePythonParserListener) ExitFuncdef(ctx *FuncdefContext) {}

// EnterTypedargslist is called when production typedargslist is entered.
func (s *BasePythonParserListener) EnterTypedargslist(ctx *TypedargslistContext) {}

// ExitTypedargslist is called when production typedargslist is exited.
func (s *BasePythonParserListener) ExitTypedargslist(ctx *TypedargslistContext) {}

// EnterArgs is called when production args is entered.
func (s *BasePythonParserListener) EnterArgs(ctx *ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *BasePythonParserListener) ExitArgs(ctx *ArgsContext) {}

// EnterKwargs is called when production kwargs is entered.
func (s *BasePythonParserListener) EnterKwargs(ctx *KwargsContext) {}

// ExitKwargs is called when production kwargs is exited.
func (s *BasePythonParserListener) ExitKwargs(ctx *KwargsContext) {}

// EnterDef_parameters is called when production def_parameters is entered.
func (s *BasePythonParserListener) EnterDef_parameters(ctx *Def_parametersContext) {}

// ExitDef_parameters is called when production def_parameters is exited.
func (s *BasePythonParserListener) ExitDef_parameters(ctx *Def_parametersContext) {}

// EnterDef_parameter is called when production def_parameter is entered.
func (s *BasePythonParserListener) EnterDef_parameter(ctx *Def_parameterContext) {}

// ExitDef_parameter is called when production def_parameter is exited.
func (s *BasePythonParserListener) ExitDef_parameter(ctx *Def_parameterContext) {}

// EnterNamed_parameter is called when production named_parameter is entered.
func (s *BasePythonParserListener) EnterNamed_parameter(ctx *Named_parameterContext) {}

// ExitNamed_parameter is called when production named_parameter is exited.
func (s *BasePythonParserListener) ExitNamed_parameter(ctx *Named_parameterContext) {}

// EnterSimple_stmt is called when production simple_stmt is entered.
func (s *BasePythonParserListener) EnterSimple_stmt(ctx *Simple_stmtContext) {}

// ExitSimple_stmt is called when production simple_stmt is exited.
func (s *BasePythonParserListener) ExitSimple_stmt(ctx *Simple_stmtContext) {}

// EnterExpr_stmt is called when production expr_stmt is entered.
func (s *BasePythonParserListener) EnterExpr_stmt(ctx *Expr_stmtContext) {}

// ExitExpr_stmt is called when production expr_stmt is exited.
func (s *BasePythonParserListener) ExitExpr_stmt(ctx *Expr_stmtContext) {}

// EnterPrint_stmt is called when production print_stmt is entered.
func (s *BasePythonParserListener) EnterPrint_stmt(ctx *Print_stmtContext) {}

// ExitPrint_stmt is called when production print_stmt is exited.
func (s *BasePythonParserListener) ExitPrint_stmt(ctx *Print_stmtContext) {}

// EnterDel_stmt is called when production del_stmt is entered.
func (s *BasePythonParserListener) EnterDel_stmt(ctx *Del_stmtContext) {}

// ExitDel_stmt is called when production del_stmt is exited.
func (s *BasePythonParserListener) ExitDel_stmt(ctx *Del_stmtContext) {}

// EnterPass_stmt is called when production pass_stmt is entered.
func (s *BasePythonParserListener) EnterPass_stmt(ctx *Pass_stmtContext) {}

// ExitPass_stmt is called when production pass_stmt is exited.
func (s *BasePythonParserListener) ExitPass_stmt(ctx *Pass_stmtContext) {}

// EnterBreak_stmt is called when production break_stmt is entered.
func (s *BasePythonParserListener) EnterBreak_stmt(ctx *Break_stmtContext) {}

// ExitBreak_stmt is called when production break_stmt is exited.
func (s *BasePythonParserListener) ExitBreak_stmt(ctx *Break_stmtContext) {}

// EnterContinue_stmt is called when production continue_stmt is entered.
func (s *BasePythonParserListener) EnterContinue_stmt(ctx *Continue_stmtContext) {}

// ExitContinue_stmt is called when production continue_stmt is exited.
func (s *BasePythonParserListener) ExitContinue_stmt(ctx *Continue_stmtContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BasePythonParserListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BasePythonParserListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterRaise_stmt is called when production raise_stmt is entered.
func (s *BasePythonParserListener) EnterRaise_stmt(ctx *Raise_stmtContext) {}

// ExitRaise_stmt is called when production raise_stmt is exited.
func (s *BasePythonParserListener) ExitRaise_stmt(ctx *Raise_stmtContext) {}

// EnterYield_stmt is called when production yield_stmt is entered.
func (s *BasePythonParserListener) EnterYield_stmt(ctx *Yield_stmtContext) {}

// ExitYield_stmt is called when production yield_stmt is exited.
func (s *BasePythonParserListener) ExitYield_stmt(ctx *Yield_stmtContext) {}

// EnterImport_stmt is called when production import_stmt is entered.
func (s *BasePythonParserListener) EnterImport_stmt(ctx *Import_stmtContext) {}

// ExitImport_stmt is called when production import_stmt is exited.
func (s *BasePythonParserListener) ExitImport_stmt(ctx *Import_stmtContext) {}

// EnterFrom_stmt is called when production from_stmt is entered.
func (s *BasePythonParserListener) EnterFrom_stmt(ctx *From_stmtContext) {}

// ExitFrom_stmt is called when production from_stmt is exited.
func (s *BasePythonParserListener) ExitFrom_stmt(ctx *From_stmtContext) {}

// EnterGlobal_stmt is called when production global_stmt is entered.
func (s *BasePythonParserListener) EnterGlobal_stmt(ctx *Global_stmtContext) {}

// ExitGlobal_stmt is called when production global_stmt is exited.
func (s *BasePythonParserListener) ExitGlobal_stmt(ctx *Global_stmtContext) {}

// EnterExec_stmt is called when production exec_stmt is entered.
func (s *BasePythonParserListener) EnterExec_stmt(ctx *Exec_stmtContext) {}

// ExitExec_stmt is called when production exec_stmt is exited.
func (s *BasePythonParserListener) ExitExec_stmt(ctx *Exec_stmtContext) {}

// EnterAssert_stmt is called when production assert_stmt is entered.
func (s *BasePythonParserListener) EnterAssert_stmt(ctx *Assert_stmtContext) {}

// ExitAssert_stmt is called when production assert_stmt is exited.
func (s *BasePythonParserListener) ExitAssert_stmt(ctx *Assert_stmtContext) {}

// EnterNonlocal_stmt is called when production nonlocal_stmt is entered.
func (s *BasePythonParserListener) EnterNonlocal_stmt(ctx *Nonlocal_stmtContext) {}

// ExitNonlocal_stmt is called when production nonlocal_stmt is exited.
func (s *BasePythonParserListener) ExitNonlocal_stmt(ctx *Nonlocal_stmtContext) {}

// EnterFrom_stmt_source is called when production from_stmt_source is entered.
func (s *BasePythonParserListener) EnterFrom_stmt_source(ctx *From_stmt_sourceContext) {}

// ExitFrom_stmt_source is called when production from_stmt_source is exited.
func (s *BasePythonParserListener) ExitFrom_stmt_source(ctx *From_stmt_sourceContext) {}

// EnterFrom_stmt_as_names is called when production from_stmt_as_names is entered.
func (s *BasePythonParserListener) EnterFrom_stmt_as_names(ctx *From_stmt_as_namesContext) {}

// ExitFrom_stmt_as_names is called when production from_stmt_as_names is exited.
func (s *BasePythonParserListener) ExitFrom_stmt_as_names(ctx *From_stmt_as_namesContext) {}

// EnterTestlist_star_expr is called when production testlist_star_expr is entered.
func (s *BasePythonParserListener) EnterTestlist_star_expr(ctx *Testlist_star_exprContext) {}

// ExitTestlist_star_expr is called when production testlist_star_expr is exited.
func (s *BasePythonParserListener) ExitTestlist_star_expr(ctx *Testlist_star_exprContext) {}

// EnterStar_expr is called when production star_expr is entered.
func (s *BasePythonParserListener) EnterStar_expr(ctx *Star_exprContext) {}

// ExitStar_expr is called when production star_expr is exited.
func (s *BasePythonParserListener) ExitStar_expr(ctx *Star_exprContext) {}

// EnterAssign_part is called when production assign_part is entered.
func (s *BasePythonParserListener) EnterAssign_part(ctx *Assign_partContext) {}

// ExitAssign_part is called when production assign_part is exited.
func (s *BasePythonParserListener) ExitAssign_part(ctx *Assign_partContext) {}

// EnterExprlist is called when production exprlist is entered.
func (s *BasePythonParserListener) EnterExprlist(ctx *ExprlistContext) {}

// ExitExprlist is called when production exprlist is exited.
func (s *BasePythonParserListener) ExitExprlist(ctx *ExprlistContext) {}

// EnterImport_as_names is called when production import_as_names is entered.
func (s *BasePythonParserListener) EnterImport_as_names(ctx *Import_as_namesContext) {}

// ExitImport_as_names is called when production import_as_names is exited.
func (s *BasePythonParserListener) ExitImport_as_names(ctx *Import_as_namesContext) {}

// EnterImport_as_name is called when production import_as_name is entered.
func (s *BasePythonParserListener) EnterImport_as_name(ctx *Import_as_nameContext) {}

// ExitImport_as_name is called when production import_as_name is exited.
func (s *BasePythonParserListener) ExitImport_as_name(ctx *Import_as_nameContext) {}

// EnterDotted_as_names is called when production dotted_as_names is entered.
func (s *BasePythonParserListener) EnterDotted_as_names(ctx *Dotted_as_namesContext) {}

// ExitDotted_as_names is called when production dotted_as_names is exited.
func (s *BasePythonParserListener) ExitDotted_as_names(ctx *Dotted_as_namesContext) {}

// EnterDotted_as_name is called when production dotted_as_name is entered.
func (s *BasePythonParserListener) EnterDotted_as_name(ctx *Dotted_as_nameContext) {}

// ExitDotted_as_name is called when production dotted_as_name is exited.
func (s *BasePythonParserListener) ExitDotted_as_name(ctx *Dotted_as_nameContext) {}

// EnterTest is called when production test is entered.
func (s *BasePythonParserListener) EnterTest(ctx *TestContext) {}

// ExitTest is called when production test is exited.
func (s *BasePythonParserListener) ExitTest(ctx *TestContext) {}

// EnterVarargslist is called when production varargslist is entered.
func (s *BasePythonParserListener) EnterVarargslist(ctx *VarargslistContext) {}

// ExitVarargslist is called when production varargslist is exited.
func (s *BasePythonParserListener) ExitVarargslist(ctx *VarargslistContext) {}

// EnterVardef_parameters is called when production vardef_parameters is entered.
func (s *BasePythonParserListener) EnterVardef_parameters(ctx *Vardef_parametersContext) {}

// ExitVardef_parameters is called when production vardef_parameters is exited.
func (s *BasePythonParserListener) ExitVardef_parameters(ctx *Vardef_parametersContext) {}

// EnterVardef_parameter is called when production vardef_parameter is entered.
func (s *BasePythonParserListener) EnterVardef_parameter(ctx *Vardef_parameterContext) {}

// ExitVardef_parameter is called when production vardef_parameter is exited.
func (s *BasePythonParserListener) ExitVardef_parameter(ctx *Vardef_parameterContext) {}

// EnterVarargs is called when production varargs is entered.
func (s *BasePythonParserListener) EnterVarargs(ctx *VarargsContext) {}

// ExitVarargs is called when production varargs is exited.
func (s *BasePythonParserListener) ExitVarargs(ctx *VarargsContext) {}

// EnterVarkwargs is called when production varkwargs is entered.
func (s *BasePythonParserListener) EnterVarkwargs(ctx *VarkwargsContext) {}

// ExitVarkwargs is called when production varkwargs is exited.
func (s *BasePythonParserListener) ExitVarkwargs(ctx *VarkwargsContext) {}

// EnterLogical_test is called when production logical_test is entered.
func (s *BasePythonParserListener) EnterLogical_test(ctx *Logical_testContext) {}

// ExitLogical_test is called when production logical_test is exited.
func (s *BasePythonParserListener) ExitLogical_test(ctx *Logical_testContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BasePythonParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BasePythonParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasePythonParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasePythonParserListener) ExitExpr(ctx *ExprContext) {}

// EnterAtom is called when production atom is entered.
func (s *BasePythonParserListener) EnterAtom(ctx *AtomContext) {}

// ExitAtom is called when production atom is exited.
func (s *BasePythonParserListener) ExitAtom(ctx *AtomContext) {}

// EnterDictorsetmaker is called when production dictorsetmaker is entered.
func (s *BasePythonParserListener) EnterDictorsetmaker(ctx *DictorsetmakerContext) {}

// ExitDictorsetmaker is called when production dictorsetmaker is exited.
func (s *BasePythonParserListener) ExitDictorsetmaker(ctx *DictorsetmakerContext) {}

// EnterTestlist_comp is called when production testlist_comp is entered.
func (s *BasePythonParserListener) EnterTestlist_comp(ctx *Testlist_compContext) {}

// ExitTestlist_comp is called when production testlist_comp is exited.
func (s *BasePythonParserListener) ExitTestlist_comp(ctx *Testlist_compContext) {}

// EnterTestlist is called when production testlist is entered.
func (s *BasePythonParserListener) EnterTestlist(ctx *TestlistContext) {}

// ExitTestlist is called when production testlist is exited.
func (s *BasePythonParserListener) ExitTestlist(ctx *TestlistContext) {}

// EnterDotted_name is called when production dotted_name is entered.
func (s *BasePythonParserListener) EnterDotted_name(ctx *Dotted_nameContext) {}

// ExitDotted_name is called when production dotted_name is exited.
func (s *BasePythonParserListener) ExitDotted_name(ctx *Dotted_nameContext) {}

// EnterName is called when production name is entered.
func (s *BasePythonParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BasePythonParserListener) ExitName(ctx *NameContext) {}

// EnterNumber is called when production number is entered.
func (s *BasePythonParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BasePythonParserListener) ExitNumber(ctx *NumberContext) {}

// EnterInteger is called when production integer is entered.
func (s *BasePythonParserListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BasePythonParserListener) ExitInteger(ctx *IntegerContext) {}

// EnterYield_expr is called when production yield_expr is entered.
func (s *BasePythonParserListener) EnterYield_expr(ctx *Yield_exprContext) {}

// ExitYield_expr is called when production yield_expr is exited.
func (s *BasePythonParserListener) ExitYield_expr(ctx *Yield_exprContext) {}

// EnterYield_arg is called when production yield_arg is entered.
func (s *BasePythonParserListener) EnterYield_arg(ctx *Yield_argContext) {}

// ExitYield_arg is called when production yield_arg is exited.
func (s *BasePythonParserListener) ExitYield_arg(ctx *Yield_argContext) {}

// EnterTrailer is called when production trailer is entered.
func (s *BasePythonParserListener) EnterTrailer(ctx *TrailerContext) {}

// ExitTrailer is called when production trailer is exited.
func (s *BasePythonParserListener) ExitTrailer(ctx *TrailerContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BasePythonParserListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BasePythonParserListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterArglist is called when production arglist is entered.
func (s *BasePythonParserListener) EnterArglist(ctx *ArglistContext) {}

// ExitArglist is called when production arglist is exited.
func (s *BasePythonParserListener) ExitArglist(ctx *ArglistContext) {}

// EnterArgument is called when production argument is entered.
func (s *BasePythonParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BasePythonParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterSubscriptlist is called when production subscriptlist is entered.
func (s *BasePythonParserListener) EnterSubscriptlist(ctx *SubscriptlistContext) {}

// ExitSubscriptlist is called when production subscriptlist is exited.
func (s *BasePythonParserListener) ExitSubscriptlist(ctx *SubscriptlistContext) {}

// EnterSubscript is called when production subscript is entered.
func (s *BasePythonParserListener) EnterSubscript(ctx *SubscriptContext) {}

// ExitSubscript is called when production subscript is exited.
func (s *BasePythonParserListener) ExitSubscript(ctx *SubscriptContext) {}

// EnterSliceop is called when production sliceop is entered.
func (s *BasePythonParserListener) EnterSliceop(ctx *SliceopContext) {}

// ExitSliceop is called when production sliceop is exited.
func (s *BasePythonParserListener) ExitSliceop(ctx *SliceopContext) {}

// EnterComp_for is called when production comp_for is entered.
func (s *BasePythonParserListener) EnterComp_for(ctx *Comp_forContext) {}

// ExitComp_for is called when production comp_for is exited.
func (s *BasePythonParserListener) ExitComp_for(ctx *Comp_forContext) {}

// EnterComp_iter is called when production comp_iter is entered.
func (s *BasePythonParserListener) EnterComp_iter(ctx *Comp_iterContext) {}

// ExitComp_iter is called when production comp_iter is exited.
func (s *BasePythonParserListener) ExitComp_iter(ctx *Comp_iterContext) {}
