package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/application/analysis/javaapp"
	"github.com/phodal/coca/pkg/application/analysis/pyapp"
	"github.com/phodal/coca/pkg/application/analysis/tsapp"
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
		var outputName string
		var ds []core_domain.CodeDataStruct
		switch analysisCmdConfig.Lang {
		case "go":
			ds = AnalysisGo()
			outputName = "godeps.json"
		case "py", "python":
			ds = AnalysisPython()
			outputName = "pydeps.json"
		case "ts", "typescript":
			ds = AnalysisTypeScript()
			outputName = "tsdeps.json"
		default:
			ds = AnalysisJava()
			outputName = "deps.json"
		}

		cModel, _ := json.MarshalIndent(ds, "", "\t")
		cmd_util.WriteToCocaFile(outputName, string(cModel))
	},
}

func AnalysisTypeScript() []core_domain.CodeDataStruct {
	importPath := analysisCmdConfig.Path

	var results []core_domain.CodeFile
	files := cocafile.GetFilesWithFilter(importPath, cocafile.TypeScriptFileFilter)
	fmt.Println(files)
	for _, file := range files {
		fmt.Fprintf(output, "Process TypeScript file: %s\n", file)
		app := new(tsapp.TypeScriptApiApp)
		content, _ := ioutil.ReadFile(file)
		result := app.Analysis(string(content), "")
		results = append(results, result)
	}

	var ds []core_domain.CodeDataStruct
	for _, result := range results {
		ds = append(ds, result.DataStructures...)
	}

	return ds
}

func AnalysisPython() []core_domain.CodeDataStruct {
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

	return ds
}

func AnalysisGo() []core_domain.CodeDataStruct {
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

	return ds
}

func AnalysisJava() []core_domain.CodeDataStruct {
	importPath := analysisCmdConfig.Path
	identifierApp := javaapp.NewJavaIdentifierApp()
	iNodes := identifierApp.AnalysisPath(importPath)

	identModel, _ := json.MarshalIndent(iNodes, "", "\t")
	cmd_util.WriteToCocaFile("identify.json", string(identModel))

	var classes []string = nil

	for _, node := range iNodes {
		classes = append(classes, node.Package+"."+node.NodeName)
	}

	callApp := javaapp.NewJavaFullApp()

	callNodes := callApp.AnalysisPath(importPath, classes, iNodes)
	return callNodes
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Path, "path", "p", ".", "example -p core/main")
	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Lang, "lang", "l", "java", "coca analysis -l java")
	analysisCmd.PersistentFlags().BoolVarP(&analysisCmdConfig.ForceUpdate, "force", "f", false, "force update -f")
}
