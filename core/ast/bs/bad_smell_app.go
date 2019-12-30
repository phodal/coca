package bs

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/context/bs"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure"
	"os"
	"path/filepath"
	"strings"

	. "github.com/phodal/coca/languages/java"
)

var nodeInfos []domain.BsJClass

type BadSmellApp struct {
}

func NewBadSmellApp() *BadSmellApp {
	return &BadSmellApp{}
}

func (j *BadSmellApp) AnalysisPath(codeDir string, ignoreRules []string) []domain.BadSmellModel {
	nodeInfos = nil
	files := (*BadSmellApp)(nil).javaFiles(codeDir)
	for index := range files {
		nodeInfo := domain.NewJFullClassNode()
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*BadSmellApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		listener := NewBadSmellListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo = listener.getNodeInfo()
		nodeInfo.Path = file
		nodeInfos = append(nodeInfos, nodeInfo)
	}

	bsModel, _ := json.MarshalIndent(nodeInfos, "", "\t")
	infrastructure.WriteToCocaFile("nodeInfos.json", string(bsModel))

	bsList := bs.AnalysisBadSmell(nodeInfos)

	mapIgnoreRules := make(map[string]bool)
	for _, ignore := range ignoreRules {
		mapIgnoreRules[ignore] = true
	}

	filteredBsList := filterBadsmellList(bsList, mapIgnoreRules)
	return filteredBsList
}

func filterBadsmellList(models []domain.BadSmellModel, ignoreRules map[string]bool) []domain.BadSmellModel {
	var results []domain.BadSmellModel
	for _, model := range models {
		if !ignoreRules[model.Bs] {
			results = append(results, model)
		}
	}
	return results
}

func (j *BadSmellApp) javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *BadSmellApp) processFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := NewJavaParser(stream)
	return parser
}
