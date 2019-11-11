package sql

import (
	"fmt"
	. "github.com/phodal/coca/language/sql"
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
	fmt.Println(ctx.GetText())
}
