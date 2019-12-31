package sql

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/infrastructure/ast/sql"
	"github.com/phodal/coca/core/infrastructure/xmlparse"
	parser2 "github.com/phodal/coca/languages/sql"
	"os"
	"path/filepath"
	"strings"
)

type SqlIdentifierApp struct {

}

func NewSqlIdentifierApp() SqlIdentifierApp {
	return *&SqlIdentifierApp{}
}

func (j *SqlIdentifierApp) AnalysisPath(codeDir string) []sql.SqlNode {
	xmlFiles := (*SqlIdentifierApp)(nil).xmlFiles(codeDir)
	for _, xmlFile := range xmlFiles {
		xmlFile, err := os.Open(xmlFile)
		if err != nil {
			fmt.Println(err)
		}

		parsedXml := xmlparse.ParseXml(xmlFile)
		for _, attr := range parsedXml.Attrs {
			if strings.Contains(attr.Name.Local, "namespace") {
				fmt.Println(attr.Value)
			}
		}
	}

	var infos []sql.SqlNode
	files := (*SqlIdentifierApp)(nil).sqlFiles(codeDir)
	for index := range files {
		file := files[index]

		parser := (*SqlIdentifierApp)(nil).processFile(file)
		context := parser.Parse()

		listener := sql.NewSqlIdentifierListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		info := listener.GetNodeInfo()
		infos = append(infos, info)
	}

	return infos
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
