package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("path", "p", "Code Path", "example -p src/main")
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))
}
