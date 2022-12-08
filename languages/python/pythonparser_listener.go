// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // PythonParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// PythonParserListener is a complete listener for a parse tree produced by PythonParser.
type PythonParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterSingle_input is called when entering the single_input production.
	EnterSingle_input(c *Single_inputContext)

	// EnterFile_input is called when entering the file_input production.
	EnterFile_input(c *File_inputContext)

	// EnterEval_input is called when entering the eval_input production.
	EnterEval_input(c *Eval_inputContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterWhile_stmt is called when entering the while_stmt production.
	EnterWhile_stmt(c *While_stmtContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterTry_stmt is called when entering the try_stmt production.
	EnterTry_stmt(c *Try_stmtContext)

	// EnterWith_stmt is called when entering the with_stmt production.
	EnterWith_stmt(c *With_stmtContext)

	// EnterClass_or_func_def_stmt is called when entering the class_or_func_def_stmt production.
	EnterClass_or_func_def_stmt(c *Class_or_func_def_stmtContext)

	// EnterSuite is called when entering the suite production.
	EnterSuite(c *SuiteContext)

	// EnterDecorator is called when entering the decorator production.
	EnterDecorator(c *DecoratorContext)

	// EnterElif_clause is called when entering the elif_clause production.
	EnterElif_clause(c *Elif_clauseContext)

	// EnterElse_clause is called when entering the else_clause production.
	EnterElse_clause(c *Else_clauseContext)

	// EnterFinally_clause is called when entering the finally_clause production.
	EnterFinally_clause(c *Finally_clauseContext)

	// EnterWith_item is called when entering the with_item production.
	EnterWith_item(c *With_itemContext)

	// EnterExcept_clause is called when entering the except_clause production.
	EnterExcept_clause(c *Except_clauseContext)

	// EnterClassdef is called when entering the classdef production.
	EnterClassdef(c *ClassdefContext)

	// EnterFuncdef is called when entering the funcdef production.
	EnterFuncdef(c *FuncdefContext)

	// EnterTypedargslist is called when entering the typedargslist production.
	EnterTypedargslist(c *TypedargslistContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterKwargs is called when entering the kwargs production.
	EnterKwargs(c *KwargsContext)

	// EnterDef_parameters is called when entering the def_parameters production.
	EnterDef_parameters(c *Def_parametersContext)

	// EnterDef_parameter is called when entering the def_parameter production.
	EnterDef_parameter(c *Def_parameterContext)

	// EnterNamed_parameter is called when entering the named_parameter production.
	EnterNamed_parameter(c *Named_parameterContext)

	// EnterSimple_stmt is called when entering the simple_stmt production.
	EnterSimple_stmt(c *Simple_stmtContext)

	// EnterExpr_stmt is called when entering the expr_stmt production.
	EnterExpr_stmt(c *Expr_stmtContext)

	// EnterPrint_stmt is called when entering the print_stmt production.
	EnterPrint_stmt(c *Print_stmtContext)

	// EnterDel_stmt is called when entering the del_stmt production.
	EnterDel_stmt(c *Del_stmtContext)

	// EnterPass_stmt is called when entering the pass_stmt production.
	EnterPass_stmt(c *Pass_stmtContext)

	// EnterBreak_stmt is called when entering the break_stmt production.
	EnterBreak_stmt(c *Break_stmtContext)

	// EnterContinue_stmt is called when entering the continue_stmt production.
	EnterContinue_stmt(c *Continue_stmtContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterRaise_stmt is called when entering the raise_stmt production.
	EnterRaise_stmt(c *Raise_stmtContext)

	// EnterYield_stmt is called when entering the yield_stmt production.
	EnterYield_stmt(c *Yield_stmtContext)

	// EnterImport_stmt is called when entering the import_stmt production.
	EnterImport_stmt(c *Import_stmtContext)

	// EnterFrom_stmt is called when entering the from_stmt production.
	EnterFrom_stmt(c *From_stmtContext)

	// EnterGlobal_stmt is called when entering the global_stmt production.
	EnterGlobal_stmt(c *Global_stmtContext)

	// EnterExec_stmt is called when entering the exec_stmt production.
	EnterExec_stmt(c *Exec_stmtContext)

	// EnterAssert_stmt is called when entering the assert_stmt production.
	EnterAssert_stmt(c *Assert_stmtContext)

	// EnterNonlocal_stmt is called when entering the nonlocal_stmt production.
	EnterNonlocal_stmt(c *Nonlocal_stmtContext)

	// EnterFrom_stmt_source is called when entering the from_stmt_source production.
	EnterFrom_stmt_source(c *From_stmt_sourceContext)

	// EnterFrom_stmt_as_names is called when entering the from_stmt_as_names production.
	EnterFrom_stmt_as_names(c *From_stmt_as_namesContext)

	// EnterTestlist_star_expr is called when entering the testlist_star_expr production.
	EnterTestlist_star_expr(c *Testlist_star_exprContext)

	// EnterStar_expr is called when entering the star_expr production.
	EnterStar_expr(c *Star_exprContext)

	// EnterAssign_part is called when entering the assign_part production.
	EnterAssign_part(c *Assign_partContext)

	// EnterExprlist is called when entering the exprlist production.
	EnterExprlist(c *ExprlistContext)

	// EnterImport_as_names is called when entering the import_as_names production.
	EnterImport_as_names(c *Import_as_namesContext)

	// EnterImport_as_name is called when entering the import_as_name production.
	EnterImport_as_name(c *Import_as_nameContext)

	// EnterDotted_as_names is called when entering the dotted_as_names production.
	EnterDotted_as_names(c *Dotted_as_namesContext)

	// EnterDotted_as_name is called when entering the dotted_as_name production.
	EnterDotted_as_name(c *Dotted_as_nameContext)

	// EnterTest is called when entering the test production.
	EnterTest(c *TestContext)

	// EnterVarargslist is called when entering the varargslist production.
	EnterVarargslist(c *VarargslistContext)

	// EnterVardef_parameters is called when entering the vardef_parameters production.
	EnterVardef_parameters(c *Vardef_parametersContext)

	// EnterVardef_parameter is called when entering the vardef_parameter production.
	EnterVardef_parameter(c *Vardef_parameterContext)

	// EnterVarargs is called when entering the varargs production.
	EnterVarargs(c *VarargsContext)

	// EnterVarkwargs is called when entering the varkwargs production.
	EnterVarkwargs(c *VarkwargsContext)

	// EnterLogical_test is called when entering the logical_test production.
	EnterLogical_test(c *Logical_testContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterAtom is called when entering the atom production.
	EnterAtom(c *AtomContext)

	// EnterDictorsetmaker is called when entering the dictorsetmaker production.
	EnterDictorsetmaker(c *DictorsetmakerContext)

	// EnterTestlist_comp is called when entering the testlist_comp production.
	EnterTestlist_comp(c *Testlist_compContext)

	// EnterTestlist is called when entering the testlist production.
	EnterTestlist(c *TestlistContext)

	// EnterDotted_name is called when entering the dotted_name production.
	EnterDotted_name(c *Dotted_nameContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterYield_expr is called when entering the yield_expr production.
	EnterYield_expr(c *Yield_exprContext)

	// EnterYield_arg is called when entering the yield_arg production.
	EnterYield_arg(c *Yield_argContext)

	// EnterTrailer is called when entering the trailer production.
	EnterTrailer(c *TrailerContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterArglist is called when entering the arglist production.
	EnterArglist(c *ArglistContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterSubscriptlist is called when entering the subscriptlist production.
	EnterSubscriptlist(c *SubscriptlistContext)

	// EnterSubscript is called when entering the subscript production.
	EnterSubscript(c *SubscriptContext)

	// EnterSliceop is called when entering the sliceop production.
	EnterSliceop(c *SliceopContext)

	// EnterComp_for is called when entering the comp_for production.
	EnterComp_for(c *Comp_forContext)

	// EnterComp_iter is called when entering the comp_iter production.
	EnterComp_iter(c *Comp_iterContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitSingle_input is called when exiting the single_input production.
	ExitSingle_input(c *Single_inputContext)

	// ExitFile_input is called when exiting the file_input production.
	ExitFile_input(c *File_inputContext)

	// ExitEval_input is called when exiting the eval_input production.
	ExitEval_input(c *Eval_inputContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitWhile_stmt is called when exiting the while_stmt production.
	ExitWhile_stmt(c *While_stmtContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitTry_stmt is called when exiting the try_stmt production.
	ExitTry_stmt(c *Try_stmtContext)

	// ExitWith_stmt is called when exiting the with_stmt production.
	ExitWith_stmt(c *With_stmtContext)

	// ExitClass_or_func_def_stmt is called when exiting the class_or_func_def_stmt production.
	ExitClass_or_func_def_stmt(c *Class_or_func_def_stmtContext)

	// ExitSuite is called when exiting the suite production.
	ExitSuite(c *SuiteContext)

	// ExitDecorator is called when exiting the decorator production.
	ExitDecorator(c *DecoratorContext)

	// ExitElif_clause is called when exiting the elif_clause production.
	ExitElif_clause(c *Elif_clauseContext)

	// ExitElse_clause is called when exiting the else_clause production.
	ExitElse_clause(c *Else_clauseContext)

	// ExitFinally_clause is called when exiting the finally_clause production.
	ExitFinally_clause(c *Finally_clauseContext)

	// ExitWith_item is called when exiting the with_item production.
	ExitWith_item(c *With_itemContext)

	// ExitExcept_clause is called when exiting the except_clause production.
	ExitExcept_clause(c *Except_clauseContext)

	// ExitClassdef is called when exiting the classdef production.
	ExitClassdef(c *ClassdefContext)

	// ExitFuncdef is called when exiting the funcdef production.
	ExitFuncdef(c *FuncdefContext)

	// ExitTypedargslist is called when exiting the typedargslist production.
	ExitTypedargslist(c *TypedargslistContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitKwargs is called when exiting the kwargs production.
	ExitKwargs(c *KwargsContext)

	// ExitDef_parameters is called when exiting the def_parameters production.
	ExitDef_parameters(c *Def_parametersContext)

	// ExitDef_parameter is called when exiting the def_parameter production.
	ExitDef_parameter(c *Def_parameterContext)

	// ExitNamed_parameter is called when exiting the named_parameter production.
	ExitNamed_parameter(c *Named_parameterContext)

	// ExitSimple_stmt is called when exiting the simple_stmt production.
	ExitSimple_stmt(c *Simple_stmtContext)

	// ExitExpr_stmt is called when exiting the expr_stmt production.
	ExitExpr_stmt(c *Expr_stmtContext)

	// ExitPrint_stmt is called when exiting the print_stmt production.
	ExitPrint_stmt(c *Print_stmtContext)

	// ExitDel_stmt is called when exiting the del_stmt production.
	ExitDel_stmt(c *Del_stmtContext)

	// ExitPass_stmt is called when exiting the pass_stmt production.
	ExitPass_stmt(c *Pass_stmtContext)

	// ExitBreak_stmt is called when exiting the break_stmt production.
	ExitBreak_stmt(c *Break_stmtContext)

	// ExitContinue_stmt is called when exiting the continue_stmt production.
	ExitContinue_stmt(c *Continue_stmtContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitRaise_stmt is called when exiting the raise_stmt production.
	ExitRaise_stmt(c *Raise_stmtContext)

	// ExitYield_stmt is called when exiting the yield_stmt production.
	ExitYield_stmt(c *Yield_stmtContext)

	// ExitImport_stmt is called when exiting the import_stmt production.
	ExitImport_stmt(c *Import_stmtContext)

	// ExitFrom_stmt is called when exiting the from_stmt production.
	ExitFrom_stmt(c *From_stmtContext)

	// ExitGlobal_stmt is called when exiting the global_stmt production.
	ExitGlobal_stmt(c *Global_stmtContext)

	// ExitExec_stmt is called when exiting the exec_stmt production.
	ExitExec_stmt(c *Exec_stmtContext)

	// ExitAssert_stmt is called when exiting the assert_stmt production.
	ExitAssert_stmt(c *Assert_stmtContext)

	// ExitNonlocal_stmt is called when exiting the nonlocal_stmt production.
	ExitNonlocal_stmt(c *Nonlocal_stmtContext)

	// ExitFrom_stmt_source is called when exiting the from_stmt_source production.
	ExitFrom_stmt_source(c *From_stmt_sourceContext)

	// ExitFrom_stmt_as_names is called when exiting the from_stmt_as_names production.
	ExitFrom_stmt_as_names(c *From_stmt_as_namesContext)

	// ExitTestlist_star_expr is called when exiting the testlist_star_expr production.
	ExitTestlist_star_expr(c *Testlist_star_exprContext)

	// ExitStar_expr is called when exiting the star_expr production.
	ExitStar_expr(c *Star_exprContext)

	// ExitAssign_part is called when exiting the assign_part production.
	ExitAssign_part(c *Assign_partContext)

	// ExitExprlist is called when exiting the exprlist production.
	ExitExprlist(c *ExprlistContext)

	// ExitImport_as_names is called when exiting the import_as_names production.
	ExitImport_as_names(c *Import_as_namesContext)

	// ExitImport_as_name is called when exiting the import_as_name production.
	ExitImport_as_name(c *Import_as_nameContext)

	// ExitDotted_as_names is called when exiting the dotted_as_names production.
	ExitDotted_as_names(c *Dotted_as_namesContext)

	// ExitDotted_as_name is called when exiting the dotted_as_name production.
	ExitDotted_as_name(c *Dotted_as_nameContext)

	// ExitTest is called when exiting the test production.
	ExitTest(c *TestContext)

	// ExitVarargslist is called when exiting the varargslist production.
	ExitVarargslist(c *VarargslistContext)

	// ExitVardef_parameters is called when exiting the vardef_parameters production.
	ExitVardef_parameters(c *Vardef_parametersContext)

	// ExitVardef_parameter is called when exiting the vardef_parameter production.
	ExitVardef_parameter(c *Vardef_parameterContext)

	// ExitVarargs is called when exiting the varargs production.
	ExitVarargs(c *VarargsContext)

	// ExitVarkwargs is called when exiting the varkwargs production.
	ExitVarkwargs(c *VarkwargsContext)

	// ExitLogical_test is called when exiting the logical_test production.
	ExitLogical_test(c *Logical_testContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitAtom is called when exiting the atom production.
	ExitAtom(c *AtomContext)

	// ExitDictorsetmaker is called when exiting the dictorsetmaker production.
	ExitDictorsetmaker(c *DictorsetmakerContext)

	// ExitTestlist_comp is called when exiting the testlist_comp production.
	ExitTestlist_comp(c *Testlist_compContext)

	// ExitTestlist is called when exiting the testlist production.
	ExitTestlist(c *TestlistContext)

	// ExitDotted_name is called when exiting the dotted_name production.
	ExitDotted_name(c *Dotted_nameContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitYield_expr is called when exiting the yield_expr production.
	ExitYield_expr(c *Yield_exprContext)

	// ExitYield_arg is called when exiting the yield_arg production.
	ExitYield_arg(c *Yield_argContext)

	// ExitTrailer is called when exiting the trailer production.
	ExitTrailer(c *TrailerContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitArglist is called when exiting the arglist production.
	ExitArglist(c *ArglistContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitSubscriptlist is called when exiting the subscriptlist production.
	ExitSubscriptlist(c *SubscriptlistContext)

	// ExitSubscript is called when exiting the subscript production.
	ExitSubscript(c *SubscriptContext)

	// ExitSliceop is called when exiting the sliceop production.
	ExitSliceop(c *SliceopContext)

	// ExitComp_for is called when exiting the comp_for production.
	ExitComp_for(c *Comp_forContext)

	// ExitComp_iter is called when exiting the comp_iter production.
	ExitComp_iter(c *Comp_iterContext)
}
