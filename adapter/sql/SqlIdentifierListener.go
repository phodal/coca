package sql

import (
	"fmt"
	. "coca/language/sql"
)

type SqlIdentifierListener struct {
	BaseSqlListener
}

func NewSqlIdentifierListener() *SqlIdentifierListener {
	return &SqlIdentifierListener{}
}

func (s *SqlIdentifierListener) EnterSelect_stmt(ctx *Select_stmtContext) {

}

func (s *SqlIdentifierListener) EnterSelect_or_values(ctx *Select_or_valuesContext) {

}

func (s *SqlIdentifierListener) EnterSelect_core(ctx *Select_coreContext) {
	columns := ctx.AllResult_column()
	for _, col := range columns {
		column := col.(*Result_columnContext)
		column.GetText()
	}
	if ctx.K_FROM() != nil {
		subqueries := ctx.AllTable_or_subquery()
		for _, subquery := range subqueries {
			subqueryCtx := subquery.(*Table_or_subqueryContext)
			if subqueryCtx.Table_name() != nil {
				fmt.Println(subqueryCtx.Table_name().GetText())
			}
			if subqueryCtx.Database_name() != nil {
				fmt.Println(subqueryCtx.Database_name().GetText())
			}
		}
	}
}
