package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/core/ast"
	"github.com/phodal/coca/core/ast/full"
	"github.com/phodal/coca/core/context/tbs"
	"github.com/phodal/coca/core/domain"
	"github.com/phodal/coca/core/infrastructure"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

type TbsCmdConfig struct {
	Path   string
	IsSort bool
}

var (
	tbsCmdConfig TbsCmdConfig
)

var tbsCmd = &cobra.Command{
	Use:   "tbs",
	Short: "generate tests bad smell",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		files := infrastructure.GetJavaTestFiles(tbsCmdConfig.Path)
		var identifiers []domain.JIdentifier

		identifiers = ast.LoadTestIdentify(files)
		identifiersMap := ast.BuildIdentifierMap(identifiers)

		var classes []string = nil
		for _, node := range identifiers {
			classes = append(classes, node.Package+"."+node.ClassName)
		}

		analysisApp := full.NewJavaFullApp()
		classNodes := analysisApp.AnalysisFiles(identifiers, files, classes)

		nodeContent, _ := json.MarshalIndent(classNodes, "", "\t")
		infrastructure.WriteToCocaFile("tdeps.json", string(nodeContent))

		app := tbs.NewTbsApp()
		result := app.AnalysisPath(classNodes, identifiersMap)

		fmt.Println("Test Bad Smell nums: ", len(result))
		resultContent, _ := json.MarshalIndent(result, "", "\t")

		if tbsCmdConfig.IsSort {
			var tbsMap = make(map[string][]tbs.TestBadSmell)
			for _, tbs := range result {
				tbsMap[tbs.Type] = append(tbsMap[tbs.Type], tbs)
			}

			resultContent, _ = json.MarshalIndent(tbsMap, "", "\t")
		}

		infrastructure.WriteToCocaFile("tbs.json", string(resultContent))

		if len(result) <= 20 {
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Type", "FileName", "Line"})

			for _, result := range result {
				table.Append([]string{result.Type, result.FileName, strconv.Itoa(result.Line)})
			}

			table.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(tbsCmd)

	tbsCmd.PersistentFlags().StringVarP(&tbsCmdConfig.Path, "path", "p", ".", "example -p core/main")
	tbsCmd.PersistentFlags().BoolVarP(&tbsCmdConfig.IsSort, "sort", "s", false, "-s")
}
