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

func (s *SqlIdentifierListener) EnterCompound_select_stmt(ctx *Compound_select_stmtContext) {
	fmt.Println(ctx.GetText())
}
