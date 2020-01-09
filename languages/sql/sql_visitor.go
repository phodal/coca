// Code generated from Sql.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Sql

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SqlParser.
type SqlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SqlParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Visit a parse tree produced by SqlParser#error.
	VisitError(ctx *ErrorContext) interface{}

	// Visit a parse tree produced by SqlParser#sql_stmt_list.
	VisitSql_stmt_list(ctx *Sql_stmt_listContext) interface{}

	// Visit a parse tree produced by SqlParser#sql_stmt.
	VisitSql_stmt(ctx *Sql_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#alter_table_stmt.
	VisitAlter_table_stmt(ctx *Alter_table_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#alter_table_add_constraint.
	VisitAlter_table_add_constraint(ctx *Alter_table_add_constraintContext) interface{}

	// Visit a parse tree produced by SqlParser#alter_table_add.
	VisitAlter_table_add(ctx *Alter_table_addContext) interface{}

	// Visit a parse tree produced by SqlParser#analyze_stmt.
	VisitAnalyze_stmt(ctx *Analyze_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#attach_stmt.
	VisitAttach_stmt(ctx *Attach_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#begin_stmt.
	VisitBegin_stmt(ctx *Begin_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#commit_stmt.
	VisitCommit_stmt(ctx *Commit_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#compound_select_stmt.
	VisitCompound_select_stmt(ctx *Compound_select_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#create_index_stmt.
	VisitCreate_index_stmt(ctx *Create_index_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#create_table_stmt.
	VisitCreate_table_stmt(ctx *Create_table_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#create_trigger_stmt.
	VisitCreate_trigger_stmt(ctx *Create_trigger_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#create_view_stmt.
	VisitCreate_view_stmt(ctx *Create_view_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#create_force_view_stmt.
	VisitCreate_force_view_stmt(ctx *Create_force_view_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#create_virtual_table_stmt.
	VisitCreate_virtual_table_stmt(ctx *Create_virtual_table_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#delete_stmt.
	VisitDelete_stmt(ctx *Delete_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#delete_stmt_limited.
	VisitDelete_stmt_limited(ctx *Delete_stmt_limitedContext) interface{}

	// Visit a parse tree produced by SqlParser#detach_stmt.
	VisitDetach_stmt(ctx *Detach_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#drop_index_stmt.
	VisitDrop_index_stmt(ctx *Drop_index_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#drop_table_stmt.
	VisitDrop_table_stmt(ctx *Drop_table_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#drop_trigger_stmt.
	VisitDrop_trigger_stmt(ctx *Drop_trigger_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#drop_view_stmt.
	VisitDrop_view_stmt(ctx *Drop_view_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#factored_select_stmt.
	VisitFactored_select_stmt(ctx *Factored_select_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#insert_stmt.
	VisitInsert_stmt(ctx *Insert_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#pragma_stmt.
	VisitPragma_stmt(ctx *Pragma_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#reindex_stmt.
	VisitReindex_stmt(ctx *Reindex_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#release_stmt.
	VisitRelease_stmt(ctx *Release_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#rollback_stmt.
	VisitRollback_stmt(ctx *Rollback_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#savepoint_stmt.
	VisitSavepoint_stmt(ctx *Savepoint_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#simple_select_stmt.
	VisitSimple_select_stmt(ctx *Simple_select_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#select_stmt.
	VisitSelect_stmt(ctx *Select_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#select_or_values.
	VisitSelect_or_values(ctx *Select_or_valuesContext) interface{}

	// Visit a parse tree produced by SqlParser#update_stmt.
	VisitUpdate_stmt(ctx *Update_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#update_stmt_limited.
	VisitUpdate_stmt_limited(ctx *Update_stmt_limitedContext) interface{}

	// Visit a parse tree produced by SqlParser#vacuum_stmt.
	VisitVacuum_stmt(ctx *Vacuum_stmtContext) interface{}

	// Visit a parse tree produced by SqlParser#column_def.
	VisitColumn_def(ctx *Column_defContext) interface{}

	// Visit a parse tree produced by SqlParser#type_name.
	VisitType_name(ctx *Type_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#column_constraint.
	VisitColumn_constraint(ctx *Column_constraintContext) interface{}

	// Visit a parse tree produced by SqlParser#column_constraint_primary_key.
	VisitColumn_constraint_primary_key(ctx *Column_constraint_primary_keyContext) interface{}

	// Visit a parse tree produced by SqlParser#column_constraint_foreign_key.
	VisitColumn_constraint_foreign_key(ctx *Column_constraint_foreign_keyContext) interface{}

	// Visit a parse tree produced by SqlParser#column_constraint_not_null.
	VisitColumn_constraint_not_null(ctx *Column_constraint_not_nullContext) interface{}

	// Visit a parse tree produced by SqlParser#column_constraint_null.
	VisitColumn_constraint_null(ctx *Column_constraint_nullContext) interface{}

	// Visit a parse tree produced by SqlParser#column_default.
	VisitColumn_default(ctx *Column_defaultContext) interface{}

	// Visit a parse tree produced by SqlParser#column_default_value.
	VisitColumn_default_value(ctx *Column_default_valueContext) interface{}

	// Visit a parse tree produced by SqlParser#conflict_clause.
	VisitConflict_clause(ctx *Conflict_clauseContext) interface{}

	// Visit a parse tree produced by SqlParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SqlParser#foreign_key_clause.
	VisitForeign_key_clause(ctx *Foreign_key_clauseContext) interface{}

	// Visit a parse tree produced by SqlParser#fk_target_column_name.
	VisitFk_target_column_name(ctx *Fk_target_column_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#raise_function.
	VisitRaise_function(ctx *Raise_functionContext) interface{}

	// Visit a parse tree produced by SqlParser#indexed_column.
	VisitIndexed_column(ctx *Indexed_columnContext) interface{}

	// Visit a parse tree produced by SqlParser#table_constraint.
	VisitTable_constraint(ctx *Table_constraintContext) interface{}

	// Visit a parse tree produced by SqlParser#table_constraint_primary_key.
	VisitTable_constraint_primary_key(ctx *Table_constraint_primary_keyContext) interface{}

	// Visit a parse tree produced by SqlParser#table_constraint_foreign_key.
	VisitTable_constraint_foreign_key(ctx *Table_constraint_foreign_keyContext) interface{}

	// Visit a parse tree produced by SqlParser#table_constraint_unique.
	VisitTable_constraint_unique(ctx *Table_constraint_uniqueContext) interface{}

	// Visit a parse tree produced by SqlParser#table_constraint_key.
	VisitTable_constraint_key(ctx *Table_constraint_keyContext) interface{}

	// Visit a parse tree produced by SqlParser#fk_origin_column_name.
	VisitFk_origin_column_name(ctx *Fk_origin_column_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#with_clause.
	VisitWith_clause(ctx *With_clauseContext) interface{}

	// Visit a parse tree produced by SqlParser#qualified_table_name.
	VisitQualified_table_name(ctx *Qualified_table_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#ordering_term.
	VisitOrdering_term(ctx *Ordering_termContext) interface{}

	// Visit a parse tree produced by SqlParser#pragma_value.
	VisitPragma_value(ctx *Pragma_valueContext) interface{}

	// Visit a parse tree produced by SqlParser#common_table_expression.
	VisitCommon_table_expression(ctx *Common_table_expressionContext) interface{}

	// Visit a parse tree produced by SqlParser#result_column.
	VisitResult_column(ctx *Result_columnContext) interface{}

	// Visit a parse tree produced by SqlParser#table_or_subquery.
	VisitTable_or_subquery(ctx *Table_or_subqueryContext) interface{}

	// Visit a parse tree produced by SqlParser#join_clause.
	VisitJoin_clause(ctx *Join_clauseContext) interface{}

	// Visit a parse tree produced by SqlParser#join_operator.
	VisitJoin_operator(ctx *Join_operatorContext) interface{}

	// Visit a parse tree produced by SqlParser#join_constraint.
	VisitJoin_constraint(ctx *Join_constraintContext) interface{}

	// Visit a parse tree produced by SqlParser#select_core.
	VisitSelect_core(ctx *Select_coreContext) interface{}

	// Visit a parse tree produced by SqlParser#compound_operator.
	VisitCompound_operator(ctx *Compound_operatorContext) interface{}

	// Visit a parse tree produced by SqlParser#cte_table_name.
	VisitCte_table_name(ctx *Cte_table_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#signed_number.
	VisitSigned_number(ctx *Signed_numberContext) interface{}

	// Visit a parse tree produced by SqlParser#literal_value.
	VisitLiteral_value(ctx *Literal_valueContext) interface{}

	// Visit a parse tree produced by SqlParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}

	// Visit a parse tree produced by SqlParser#error_message.
	VisitError_message(ctx *Error_messageContext) interface{}

	// Visit a parse tree produced by SqlParser#module_argument.
	VisitModule_argument(ctx *Module_argumentContext) interface{}

	// Visit a parse tree produced by SqlParser#column_alias.
	VisitColumn_alias(ctx *Column_aliasContext) interface{}

	// Visit a parse tree produced by SqlParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by SqlParser#unknown.
	VisitUnknown(ctx *UnknownContext) interface{}

	// Visit a parse tree produced by SqlParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by SqlParser#function_name.
	VisitFunction_name(ctx *Function_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#database_name.
	VisitDatabase_name(ctx *Database_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#source_table_name.
	VisitSource_table_name(ctx *Source_table_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#table_or_index_name.
	VisitTable_or_index_name(ctx *Table_or_index_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#new_table_name.
	VisitNew_table_name(ctx *New_table_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#collation_name.
	VisitCollation_name(ctx *Collation_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#foreign_table.
	VisitForeign_table(ctx *Foreign_tableContext) interface{}

	// Visit a parse tree produced by SqlParser#index_name.
	VisitIndex_name(ctx *Index_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#trigger_name.
	VisitTrigger_name(ctx *Trigger_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#view_name.
	VisitView_name(ctx *View_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#module_name.
	VisitModule_name(ctx *Module_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#pragma_name.
	VisitPragma_name(ctx *Pragma_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#savepoint_name.
	VisitSavepoint_name(ctx *Savepoint_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#table_alias.
	VisitTable_alias(ctx *Table_aliasContext) interface{}

	// Visit a parse tree produced by SqlParser#transaction_name.
	VisitTransaction_name(ctx *Transaction_nameContext) interface{}

	// Visit a parse tree produced by SqlParser#any_name.
	VisitAny_name(ctx *Any_nameContext) interface{}
}
