package cmd

import (
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
			callApp.AnalysisPath(importPath)

			identifierApp := new(JavaIdentifierApp)
			identifierApp.AnalysisPath(importPath)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("path", "p", "Code Path", "example -p src/main")
	viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))
}
