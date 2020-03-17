package javaapp

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/ast_java"
	"path/filepath"
)

type JavaFullApp struct {
}

func NewJavaFullApp() JavaFullApp {
	return JavaFullApp{}
}

func (j *JavaFullApp) AnalysisPath(codeDir string, identNodes []core_domain.CodeDataStruct) []core_domain.CodeDataStruct {
	files := cocafile.GetJavaFiles(codeDir)
	return j.AnalysisFiles(identNodes, files)
}

func (j *JavaFullApp) AnalysisFiles(identNodes []core_domain.CodeDataStruct, files []string) []core_domain.CodeDataStruct {
	var nodeInfos []core_domain.CodeDataStruct

	var identMap = make(map[string]core_domain.CodeDataStruct)
	for _, ident := range identNodes {
		identMap[ident.GetClassFullName()] = ident
	}

	var classes []string = nil
	for _, node := range identNodes {
		classes = append(classes, node.GetClassFullName())
	}

	for _, file := range files {
		displayName := filepath.Base(file)
		fmt.Println("parse java call: " + displayName)

		parser := ast_java.ProcessJavaFile(file)
		context := parser.CompilationUnit()

		listener := ast_java.NewJavaFullListener(identMap, file)
		listener.AppendClasses(classes)

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodes := listener.GetNodeInfo()
		nodeInfos = append(nodeInfos, nodes...)
	}

	return nodeInfos
}
