package sql

import (
	"coca/core/adapter/sql/parse"
	parser2 "coca/core/language/sql"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strings"
)

type SqlIdentifierApp struct {

}

func (j *SqlIdentifierApp) AnalysisPath(codeDir string) {
	xmlFiles := (*SqlIdentifierApp)(nil).xmlFiles(codeDir)
	for _, xmlFile := range xmlFiles {
		xmlFile, err := os.Open(xmlFile)
		if err != nil {
			fmt.Println(err)
		}

		parsedXml := parse.ParseXml(xmlFile)
		for _, attr := range parsedXml.Attrs {
			if strings.Contains(attr.Name.Local, "namespace") {
				fmt.Println(attr.Value)
			}
		}
	}

	files := (*SqlIdentifierApp)(nil).sqlFiles(codeDir)
	for index := range files {
		file := files[index]

		parser := (*SqlIdentifierApp)(nil).processFile(file)
		context := parser.Parse()

		listener := NewSqlIdentifierListener()

		antlr.NewParseTreeWalker().Walk(listener, context)
	}
}

func (j *SqlIdentifierApp) xmlFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, "Mapper.xml") {
			files = append(files, path)
		}
		return nil
	})
	return files
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

func (j *SqlIdentifierApp) processFile(path string) *parser2.SqlParser {
	is, _ := antlr.NewFileStream(path)
	lexer := parser2.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := parser2.NewSqlParser(stream)
	return parser
}
