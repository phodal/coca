package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/application/analysis/app_concept"
	"github.com/phodal/coca/pkg/application/analysis/goapp"
	"github.com/phodal/coca/pkg/application/analysis/javaapp"
	"github.com/phodal/coca/pkg/application/analysis/pyapp"
	"github.com/phodal/coca/pkg/application/analysis/tsapp"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
	"io/ioutil"
	"unicode"
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
	return CommentAnalysis(analysisCmdConfig.Path, new(tsapp.TypeScriptIdentApp), cocafile.TypeScriptFileFilter, true)
}

func AnalysisPython() []core_domain.CodeDataStruct {
	return CommentAnalysis(analysisCmdConfig.Path, new(pyapp.PythonIdentApp), cocafile.PythonFileFilter, true)
}

func AnalysisGo() []core_domain.CodeDataStruct {
	return CommentAnalysis(analysisCmdConfig.Path, new(goapp.GoIdentApp), cocafile.GoFileFilter, true)
}

func CommentAnalysis(path string, app app_concept.AbstractAnalysisApp, filter func(path string) bool, isFunctionBase bool) []core_domain.CodeDataStruct {
	var results []core_domain.CodeContainer
	files := cocafile.GetFilesWithFilter(path, filter)

	var codeMembers []core_domain.CodeMember

	app.AnalysisPackageManager(path)

	for _, file := range files {
		content, _ := ioutil.ReadFile(file)
		members := app.IdentAnalysis(string(content), file)
		codeMembers = append(codeMembers, members...)

		identModel, _ := json.MarshalIndent(codeMembers, "", "\t")
		cmd_util.WriteToCocaFile("members.json", string(identModel))
	}

	for _, file := range files {
		fmt.Fprintf(output, "Process file: %s\n", file)
		content, _ := ioutil.ReadFile(file)
		app.SetExtensions(codeMembers)
		result := app.Analysis(string(content), file)
		results = append(results, result)
	}

	var ds []core_domain.CodeDataStruct
	for _, result := range results {
		ds = append(ds, result.DataStructures...)

		if isFunctionBase {
			methodDs := BuildMethodDs(result)
			ds = append(ds, methodDs...)
		}
	}

	return ds
}

func BuildMethodDs(result core_domain.CodeContainer) []core_domain.CodeDataStruct {
	var methodsDs []core_domain.CodeDataStruct
	for _, member := range result.Members {
		for _, node := range member.FunctionNodes {
			if unicode.IsUpper(rune(node.Name[0])) {
				methodDs := core_domain.CodeDataStruct{
					NodeName:      node.Name,
					Package:       result.PackageName,
					FunctionCalls: node.FunctionCalls,
				}
				methodsDs = append(methodsDs, methodDs)
			}
		}
	}

	return methodsDs
}

func AnalysisJava() []core_domain.CodeDataStruct {
	importPath := analysisCmdConfig.Path
	identifierApp := javaapp.NewJavaIdentifierApp()
	iNodes := identifierApp.AnalysisPath(importPath)

	identModel, _ := json.MarshalIndent(iNodes, "", "\t")
	cmd_util.WriteToCocaFile("identify.json", string(identModel))

	callApp := javaapp.NewJavaFullApp()

	callNodes := callApp.AnalysisPath(importPath, iNodes)
	return callNodes
}

func init() {
	rootCmd.AddCommand(analysisCmd)

	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Path, "path", "p", ".", "example -p core/main")
	analysisCmd.PersistentFlags().StringVarP(&analysisCmdConfig.Lang, "lang", "l", "java", "coca analysis -l java")
	analysisCmd.PersistentFlags().BoolVarP(&analysisCmdConfig.ForceUpdate, "force", "f", false, "force update -f")
}
