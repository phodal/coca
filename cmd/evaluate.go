package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/pkg/application/evaluate"
	"github.com/phodal/coca/pkg/application/evaluate/evaluator"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
	"log"
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
	Short: "evaluate code situation and refactor effort",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dependence := evaluateConfig.DependencePath

		analyser := evaluate.NewEvaluateAnalyser()
		file := cmd_util.ReadFile(dependence)
		if file == nil {
			log.Fatal("lost file:" + dependence)
		}

		var identifiers []core_domain.CodeDataStruct
		identContent := cmd_util.ReadCocaFile("identify.json")

		_ = json.Unmarshal(identContent, &identifiers)
		_ = json.Unmarshal(file, &parsedDeps)

		result := analyser.Analysis(parsedDeps, identifiers)

		cModel, _ := json.MarshalIndent(result, "", "\t")
		cmd_util.WriteToCocaFile("evaluate.json", string(cModel))

		buildOutput(result)
	},
}

func buildOutput(result evaluator.EvaluateModel) {
	table := cmd_util.NewOutput(output)
	table.SetHeader([]string{"Type", "Count", "Level", "Total", "Rate"})

	classCount := result.Summary.ClassCount
	methodCount := result.Summary.MethodCount

	nullItemsLength := len(result.Nullable.Items)
	table.Append([]string{"Nullable / Return Null", strconv.Itoa(nullItemsLength), "Method", strconv.Itoa(methodCount), Percent(nullItemsLength, methodCount)})

	utilsCount := result.Summary.UtilsCount
	table.Append([]string{"Utils", strconv.Itoa(utilsCount), "Class", strconv.Itoa(classCount), Percent(utilsCount, classCount)})

	staticCount := result.Summary.StaticMethodCount
	table.Append([]string{"Static Method", strconv.Itoa(staticCount), "Method", strconv.Itoa(methodCount), Percent(utilsCount, methodCount)})

	table.Append([]string{"Average Method Num.", strconv.Itoa(methodCount), "Method/Class", strconv.Itoa(classCount), Rate(methodCount, classCount)})
	table.Append([]string{"Method Num. Std Dev / 标准差", strconv.Itoa(methodCount), "Class", "-", fmt.Sprintf("%f", result.Summary.MethodNumStdDeviation)})

	totalLength := result.Summary.TotalMethodLength
	normalMethodCount := result.Summary.NormalMethodCount
	table.Append([]string{"Average Method Length", strconv.Itoa(totalLength), "Without Getter/Setter", strconv.Itoa(normalMethodCount), Rate(totalLength, normalMethodCount)})

	table.Append([]string{"Method Length Std Dev / 标准差", strconv.Itoa(methodCount), "Method", "-", fmt.Sprintf("%f", result.Summary.MethodLengthStdDeviation)})

	table.Render()
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
