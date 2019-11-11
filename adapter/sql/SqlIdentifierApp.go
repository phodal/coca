package sql

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/phodal/coca/language/sql"
	"os"
	"path/filepath"
	"strings"
)

type SqlIdentifierApp struct {

}

func (j *SqlIdentifierApp) AnalysisPath(codeDir string) {
	files := (*SqlIdentifierApp)(nil).sqlFiles(codeDir)
	for index := range files {
		file := files[index]

		parser := (*SqlIdentifierApp)(nil).processFile(file)
		context := parser.Parse()

		listener := NewSqlIdentifierListener()

		antlr.NewParseTreeWalker().Walk(listener, context)
	}
}

func (j *SqlIdentifierApp) sqlFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".sql") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *SqlIdentifierApp) processFile(path string) *parser.SqlParser {
	is, _ := antlr.NewFileStream(path)
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := parser.NewSqlParser(stream)
	return parser
}
