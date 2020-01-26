package ast_sql

import (
	"fmt"
	"github.com/phodal/coca/languages/sql"
)

type SQLIdentifierListener struct {
	parser.BaseSqlListener
}

type SQLNode struct {
}

var sqlNode SQLNode

func init() {
	sqlNode = SQLNode{}
}

func NewSqlIdentifierListener() *SQLIdentifierListener {
	return &SQLIdentifierListener{}
}

func (s *SQLIdentifierListener) EnterSelect_core(ctx *parser.Select_coreContext) {
	columns := ctx.AllResult_column()
	for _, col := range columns {
		column := col.(*parser.Result_columnContext)
		column.GetText()
	}
	if ctx.K_FROM() != nil {
		subqueries := ctx.AllTable_or_subquery()
		for _, subquery := range subqueries {
			subqueryCtx := subquery.(*parser.Table_or_subqueryContext)
			if subqueryCtx.Table_name() != nil {
				fmt.Println(subqueryCtx.Table_name().GetText())
			}
			if subqueryCtx.Database_name() != nil {
				fmt.Println(subqueryCtx.Database_name().GetText())
			}
		}
	}
}

func (s *SQLIdentifierListener) GetNodeInfo() SQLNode {
	return sqlNode
}
