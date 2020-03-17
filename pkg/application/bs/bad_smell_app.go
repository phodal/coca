package bs

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/bs_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java"
	"github.com/phodal/coca/pkg/infrastructure/ast/bs_java"
	"path/filepath"
)

var nodeInfos []bs_domain.BSDataStruct

type BadSmellApp struct {
}

func NewBadSmellApp() *BadSmellApp {
	return &BadSmellApp{}
}

func (j *BadSmellApp) AnalysisPath(codeDir string) *[]bs_domain.BSDataStruct {
	nodeInfos = nil
	files := cocafile.GetJavaFiles(codeDir)
	for index := range files {
		nodeInfo := bs_domain.NewJFullClassNode()
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("parse java call: " + displayName)

		parser := ast_java.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := bs_java.NewBadSmellListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo = listener.GetNodeInfo()
		nodeInfo.FilePath = file
		nodeInfos = append(nodeInfos, nodeInfo)
	}

	return &nodeInfos
}

func (j *BadSmellApp) IdentifyBadSmell(nodeInfos *[]bs_domain.BSDataStruct, ignoreRules []string) []bs_domain.BadSmellModel {
	bsList := AnalysisBadSmell(*nodeInfos)

	mapIgnoreRules := make(map[string]bool)
	for _, ignore := range ignoreRules {
		mapIgnoreRules[ignore] = true
	}

	filteredBsList := bs_domain.FilterBadSmellList(bsList, mapIgnoreRules)
	return filteredBsList
}
