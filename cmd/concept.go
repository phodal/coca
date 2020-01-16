package cmd

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/pkg/application/concept"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
	"strconv"
)

var parsedDeps []core_domain.CodeDataStruct

var conceptCmd = &cobra.Command{
	Use:   "concept",
	Short: "build domain concept from source code",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := cmd.Flag("dependence").Value.String()

		if dependence != "" {
			analyser := concept.NewConceptAnalyser()
			file := cmd_util.ReadFile(dependence)
			_ = json.Unmarshal(file, &parsedDeps)

			wordCounts := analyser.Analysis(&parsedDeps)

			table := cmd_util.NewOutput(output)
			table.SetHeader([]string{"Words", "Counts"})

			for _, word := range wordCounts {
				if word.Value > 0 {
					table.Append([]string{word.Key, strconv.Itoa(word.Value)})
				}
			}

			table.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(conceptCmd)

	conceptCmd.PersistentFlags().StringP("dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
