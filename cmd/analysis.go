package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"

	. "../adapter/call"
	. "../adapter/identifier"
)

var collCmd *cobra.Command = &cobra.Command{
	Use:   "analysis",
	Short: "analysis package",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		importPath := cmd.Flag("path").Value.String()

		if importPath != "" {
			callApp := new(JavaCallApp)
			callNodes := callApp.AnalysisPath(importPath)

			cModel, _ := json.Marshal(callNodes)

			identifierApp := new(JavaIdentifierApp)
			identNodes := identifierApp.AnalysisPath(importPath)
			iNodes, _ := json.Marshal(identNodes)

			fmt.Println(string(cModel))
			fmt.Println(string(iNodes))
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("path", "p", "Code Path", "example -p src/main")
}
