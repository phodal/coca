package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/application/analysis"
	"github.com/phodal/coca/pkg/application/pyapp"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/infrastructure/ast/cocago"
	"github.com/spf13/cobra"
	"io/ioutil"
)

type AnalysisCmdConfig struct {
	Path        string
	ForceUpdate bool
	Lang        string
}

var (
	analysisCmdConfig AnalysisCmdConfig
)

var analysisCmd = &cobra.Command{
	Use:   "analysis",
	Short: "analysis code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch analysisCmdConfig.Lang {
		case "go":
			analysisGo()
		case "py":
		case "python":
			analysisPython()
		default:
			analysisJava()
		}
	},
}

func analysisPython() {
	importPath := analysisCmdConfig.Path

	var results []core_domain.CodeFile
	files := cocafile.GetFilesWithFilter(importPath, cocafile.PythonFileFilter)
	for _, file := range files {
		fmt.Fprintf(output, "Process Python file: %s\n", file)
		app := new(pyapp.PythonApiApp)
		content, _ := ioutil.ReadFile(file)
		result := app.Analysis(string(content), "")
		results = append(results, result)
	}

	var ds []core_domain.CodeDataStruct
	for _, result := range results {
		ds = append(ds, result.DataStructures...)
	}

	cModel, _ := json.MarshalIndent(ds, "", "\t")
	cmd_util.WriteToCocaFile("pydeps.json", string(cModel))
}

func analysisGo() {
	importPath := analysisCmdConfig.Path

	var results []core_domain.CodeFile
	files := cocafile.GetFilesWithFilter(importPath, cocafile.GoFileFilter)
	for _, file := range files {
		parser := cocago.NewCocagoParser()
		parser.SetOutput(output)
		result := parser.ProcessFile(file)

		results = append(results, result)
	}

	var ds []core_domain.CodeDataStruct
	for _, result := range results {
		ds = append(ds, result.DataStructures...)
	}

	cModel, _ := json.MarshalIndent(ds, "", "\t")
	cmd_util.WriteToCocaFile("godeps.json", string(cModel))
}

func analysisJava() {
	importPath := analysisCmdConfig.Path
	identifierApp := analysis.NewJavaIdentifierApp()
	iNodes := identifierApp.AnalysisPath(importPath)

	identModel, _ := json.MarshalIndent(iNodes, "", "\t")
	cmd_util.WriteToCocaFile("identify.json", string(identModel))

	var classes []string = nil

	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	callApp := analysis.NewJavaFullApp()

	callNodes := callApp.AnalysisPath(importPath, classes, iNodes)
	cModel, _ := json.MarshalIndent(callNodes, "", "\t")
	cmd_util.WriteToCocaFile("deps.json", string(cModel))
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Path, "path", "p", ".", "example -p core/main")
	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Lang, "lang", "l", "java", "coca analysis -l java")
	analysisCmd.PersistentFlags().BoolVarP(&analysisCmdConfig.ForceUpdate, "force", "f", false, "force update -f")
}
