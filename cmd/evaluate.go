package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/config"
	"github.com/phodal/coca/core/domain/evaluate"
	"github.com/phodal/coca/core/models"
	. "github.com/phodal/coca/core/support"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

type EvaluateConfig struct {
	DependencePath string
}

var (
	evaluateConfig EvaluateConfig
)

var evaluateCmd = &cobra.Command{
	Use:   "evaluate",
	Short: "evaluate refactor effort",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := *&evaluateConfig.DependencePath

		analyser := evaluate.NewEvaluateAnalyser()
		file := ReadFile(dependence)
		if file == nil {
			log.Fatal("lost file:" + dependence)
		}

		var identifiers []models.JIdentifier
		identContent := ReadCocaFile("identify.json")

		_ = json.Unmarshal(identContent, &identifiers)
		_ = json.Unmarshal(file, &parsedDeps)

		result := analyser.Analysis(parsedDeps, identifiers)

		cModel, _ := json.MarshalIndent(result, "", "\t")
		WriteToCocaFile("evaluate.json", string(cModel))

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Type", "Count", "Level", "Total", "Rate"})

		classCount := result.Summary.ClassCount
		methodCount := result.Summary.MethodCount

		nullItemsLength := len(result.Nullable.Items)
		table.Append([]string{"Nullable / Return Null", strconv.Itoa(nullItemsLength), "Method", strconv.Itoa(methodCount), Percent(nullItemsLength, methodCount)})

		utilsCount := result.Summary.UtilsCount
		table.Append([]string{"Utils", strconv.Itoa(utilsCount), "Class", strconv.Itoa(classCount), Percent(utilsCount, classCount)})

		staticCount := result.Summary.StaticMethodCount
		table.Append([]string{"Static Method", strconv.Itoa(staticCount), "Method", strconv.Itoa(methodCount), Percent(utilsCount, methodCount)})

		table.Append([]string{"Average Method Num", strconv.Itoa(methodCount), "Method/Class", strconv.Itoa(classCount), Rate(methodCount, classCount)})

		totalLength := result.Summary.TotalMethodLength
		normalMethodCount := result.Summary.NormalMethodCount
		table.Append([]string{"Average Method Length", strconv.Itoa(totalLength), "Without Getter/Setter", strconv.Itoa(normalMethodCount), Rate(totalLength, normalMethodCount)})

		table.Render()
	},
}

func Percent(pcent int, all int) string {
	percent := 100.0 * float64(pcent) / float64(all)
	return fmt.Sprintf("%3.2f%%", percent)
}

func Rate(pcent int, all int) string {
	percent := float64(pcent) / float64(all)
	return fmt.Sprintf("%f", percent)
}

func init() {
	rootCmd.AddCommand(evaluateCmd)

	evaluateCmd.PersistentFlags().StringVarP(&evaluateConfig.DependencePath, "dependence", "d", config.CocaConfig.ReporterPath+"/deps.json", "get dependence file")
}
