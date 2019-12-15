package cmd

import (
	"coca/src/adapter/models"
	. "coca/src/domain"
	. "coca/src/utils"
	"encoding/json"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
)


var callGraphCmd *cobra.Command = &cobra.Command{
	Use:   "call",
	Short: "call graph api",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var parsedDeps []models.JClassNode
		className := cmd.Flag("className").Value.String()
		dependence := cmd.Flag("dependence").Value.String()

		if dependence != "" {
			analyser := NewCallGraph()
			file := ReadFile(dependence)
			if file == nil {
				log.Fatal("lost file:" + dependence)
			}

 			_ = json.Unmarshal(file, &parsedDeps)

			content := analyser.Analysis(className, *&parsedDeps)
			WriteToFile("call.dot", content)

			cmd := exec.Command("dot", []string{"-Tsvg", "call.dot", "-o", "call.svg"}...)
			_, err := cmd.CombinedOutput()
			if err != nil {
				log.Fatalf("cmd.Run() failed with %s\n", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(callGraphCmd)

	callGraphCmd.PersistentFlags().StringP("className", "c", "", "path")
	callGraphCmd.PersistentFlags().StringP("dependence", "d", "", "get dependence file")
}
