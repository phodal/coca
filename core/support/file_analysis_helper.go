package support

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	ignore "github.com/sabhiram/go-gitignore"
	"os"
	"path/filepath"
	"strings"

	. "github.com/phodal/coca/languages/java"
)

func GetJavaFiles(codeDir string) []string {
	files := make([]string, 0)
	gitIgnore, err := ignore.CompileIgnoreFile(".gitIgnore")
	if err != nil {
		fmt.Println(err)
	}

	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if gitIgnore != nil {
			if gitIgnore.MatchesPath(path) {
				return nil
			}
		}

		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java")&& !strings.Contains(path, "Tests.java"){
			files = append(files, path)
		}
		return nil
	})
	return files
}

func ProcessFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
