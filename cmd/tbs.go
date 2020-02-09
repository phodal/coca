package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/adapter/cocafile"
	"github.com/phodal/coca/pkg/application/analysis/javaapp"
	"github.com/phodal/coca/pkg/application/tbs"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
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
		files := cocafile.GetJavaTestFiles(tbsCmdConfig.Path)

		identifiers := cmd_util.LoadTestIdentify(files)
		identifiersMap := core_domain.BuildIdentifierMap(identifiers)

		analysisApp := javaapp.NewJavaFullApp()
		classNodes := analysisApp.AnalysisFiles(identifiers, files)

		nodeContent, _ := json.MarshalIndent(classNodes, "", "\t")
		cmd_util.WriteToCocaFile("tdeps.json", string(nodeContent))

		app := tbs.NewTbsApp()
		result := app.AnalysisPath(classNodes, identifiersMap)

		fmt.Fprintf(output, "Test Bad Smell nums:  %d\n", len(result))
		resultContent, _ := json.MarshalIndent(result, "", "\t")

		if tbsCmdConfig.IsSort {
			var tbsMap = make(map[string][]tbs.TestBadSmell)
			for _, tbs := range result {
				tbsMap[tbs.Type] = append(tbsMap[tbs.Type], tbs)
			}

			resultContent, _ = json.MarshalIndent(tbsMap, "", "\t")
		}

		cmd_util.WriteToCocaFile("tbs.json", string(resultContent))
		if len(result) <= 20 {
			table := cmd_util.NewOutput(output)
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
