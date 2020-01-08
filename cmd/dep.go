package cmd

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/phodal/coca/core/adapter/cocafile"
	"github.com/phodal/coca/core/context/analysis"
	"github.com/phodal/coca/core/context/deps"
	"github.com/spf13/cobra"
	"path/filepath"
)

type DepCmdConfig struct {
	Path string
}

var (
	depCmdConfig DepCmdConfig
)

var depCmd = &cobra.Command{
	Use:   "deps",
	Short: "evaluate dependencies",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := depCmdConfig.Path

		path, _ = filepath.Abs(path)
		files := cocafile.GetFilesWithFilter(path, cocafile.JavaFileFilter)

		fmt.Println(path, files)
		identifierApp := analysis.NewJavaIdentifierApp()
		iNodes := identifierApp.AnalysisFiles(files)

		var classes []string = nil

		for _, node := range iNodes {
			classes = append(classes, node.Package+"."+node.ClassName)
		}

		callApp := analysis.NewJavaFullApp()
		classNodes := callApp.AnalysisFiles(iNodes, files, classes)

		depApp := deps.NewDepApp()
		deps := depApp.AnalysisPath(path, classNodes)

		fmt.Fprintln(output, "unused")
		table := tablewriter.NewWriter(output)
		table.SetHeader([]string{"GroupId", "ArtifactId", "Scope"})
		for _, dep := range deps {
			table.Append([]string{dep.GroupId, dep.ArtifactId, dep.Scope})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(depCmd)

	depCmd.PersistentFlags().StringVarP(&depCmdConfig.Path, "path", "p", ".", "example -p core/main")
}
