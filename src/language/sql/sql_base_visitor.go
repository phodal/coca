// Code generated from Sql.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSqlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSqlVisitor) VisitParse(ctx *ParseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitError(ctx *ErrorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSql_stmt_list(ctx *Sql_stmt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSql_stmt(ctx *Sql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitAlter_table_stmt(ctx *Alter_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitAlter_table_add_constraint(ctx *Alter_table_add_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitAlter_table_add(ctx *Alter_table_addContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitAnalyze_stmt(ctx *Analyze_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitAttach_stmt(ctx *Attach_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitBegin_stmt(ctx *Begin_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCommit_stmt(ctx *Commit_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCompound_select_stmt(ctx *Compound_select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCreate_index_stmt(ctx *Create_index_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCreate_table_stmt(ctx *Create_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCreate_trigger_stmt(ctx *Create_trigger_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCreate_force_view_stmt(ctx *Create_force_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCreate_virtual_table_stmt(ctx *Create_virtual_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDelete_stmt(ctx *Delete_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDelete_stmt_limited(ctx *Delete_stmt_limitedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDetach_stmt(ctx *Detach_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDrop_index_stmt(ctx *Drop_index_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDrop_table_stmt(ctx *Drop_table_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDrop_trigger_stmt(ctx *Drop_trigger_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDrop_view_stmt(ctx *Drop_view_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitFactored_select_stmt(ctx *Factored_select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitInsert_stmt(ctx *Insert_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitPragma_stmt(ctx *Pragma_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitReindex_stmt(ctx *Reindex_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitRelease_stmt(ctx *Release_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitRollback_stmt(ctx *Rollback_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSavepoint_stmt(ctx *Savepoint_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSimple_select_stmt(ctx *Simple_select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSelect_stmt(ctx *Select_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSelect_or_values(ctx *Select_or_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitUpdate_stmt(ctx *Update_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitUpdate_stmt_limited(ctx *Update_stmt_limitedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitVacuum_stmt(ctx *Vacuum_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_def(ctx *Column_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitType_name(ctx *Type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_constraint(ctx *Column_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_constraint_primary_key(ctx *Column_constraint_primary_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_constraint_foreign_key(ctx *Column_constraint_foreign_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_constraint_not_null(ctx *Column_constraint_not_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_constraint_null(ctx *Column_constraint_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_default(ctx *Column_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_default_value(ctx *Column_default_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitConflict_clause(ctx *Conflict_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitForeign_key_clause(ctx *Foreign_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitFk_target_column_name(ctx *Fk_target_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitRaise_function(ctx *Raise_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitIndexed_column(ctx *Indexed_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_constraint(ctx *Table_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_constraint_primary_key(ctx *Table_constraint_primary_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_constraint_foreign_key(ctx *Table_constraint_foreign_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_constraint_unique(ctx *Table_constraint_uniqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_constraint_key(ctx *Table_constraint_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitFk_origin_column_name(ctx *Fk_origin_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitWith_clause(ctx *With_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitQualified_table_name(ctx *Qualified_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitOrdering_term(ctx *Ordering_termContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitPragma_value(ctx *Pragma_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCommon_table_expression(ctx *Common_table_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitResult_column(ctx *Result_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_or_subquery(ctx *Table_or_subqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitJoin_clause(ctx *Join_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitJoin_operator(ctx *Join_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitJoin_constraint(ctx *Join_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSelect_core(ctx *Select_coreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCompound_operator(ctx *Compound_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCte_table_name(ctx *Cte_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSigned_number(ctx *Signed_numberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitLiteral_value(ctx *Literal_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitError_message(ctx *Error_messageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitModule_argument(ctx *Module_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_alias(ctx *Column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitUnknown(ctx *UnknownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitDatabase_name(ctx *Database_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSource_table_name(ctx *Source_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_or_index_name(ctx *Table_or_index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitNew_table_name(ctx *New_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitCollation_name(ctx *Collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitForeign_table(ctx *Foreign_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitIndex_name(ctx *Index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTrigger_name(ctx *Trigger_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitView_name(ctx *View_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitModule_name(ctx *Module_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitPragma_name(ctx *Pragma_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitSavepoint_name(ctx *Savepoint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTable_alias(ctx *Table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitTransaction_name(ctx *Transaction_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlVisitor) VisitAny_name(ctx *Any_nameContext) interface{} {
	return v.VisitChildren(ctx)
}
