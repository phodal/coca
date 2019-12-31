package bs

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/core/ast/bs"
	"github.com/phodal/coca/core/domain/bs_domain"
	"github.com/phodal/coca/core/infrastructure/coca_file"
	"path/filepath"
)

var nodeInfos []bs_domain.BsJClass

type BadSmellApp struct {
}

func NewBadSmellApp() *BadSmellApp {
	return &BadSmellApp{}
}

func (j *BadSmellApp) AnalysisPath(codeDir string, ignoreRules []string) []bs_domain.BadSmellModel {
	nodeInfos = nil
	files := coca_file.GetJavaFiles(codeDir)
	for index := range files {
		nodeInfo := bs_domain.NewJFullClassNode()
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := coca_file.ProcessFile(file)
		context := parser.CompilationUnit()

		listener := bs.NewBadSmellListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo = listener.GetNodeInfo()
		nodeInfo.Path = file
		nodeInfos = append(nodeInfos, nodeInfo)
	}

	bsModel, _ := json.MarshalIndent(nodeInfos, "", "\t")
	coca_file.WriteToCocaFile("nodeInfos.json", string(bsModel))

	bsList := AnalysisBadSmell(nodeInfos)

	mapIgnoreRules := make(map[string]bool)
	for _, ignore := range ignoreRules {
		mapIgnoreRules[ignore] = true
	}

	filteredBsList := bs_domain.FilterBadSmellList(bsList, mapIgnoreRules)
	return filteredBsList
}
