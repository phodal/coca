package app

import (
	"fmt"
	"github.com/modernizing/coca/cmd/cmd_util"
	"github.com/modernizing/coca/pkg/adapter/cocafile"
	"github.com/modernizing/coca/pkg/application/analysis/javaapp"
	"github.com/modernizing/coca/pkg/application/deps"
	"github.com/modernizing/coca/pkg/domain/core_domain"
	"github.com/spf13/cobra"
	"path/filepath"
)

type DepCmdConfig struct {
	Path string
}

var (
	depCmdConfig DepCmdConfig
)

type DepApp interface {
	AnalysisPath(path string, nodes []core_domain.CodeDataStruct) []core_domain.CodeDependency
}

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "evaluate dependencies",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := depCmdConfig.Path

		path, _ = filepath.Abs(path)
		files := cocafile.GetFilesWithFilter(path, cocafile.JavaFileFilter)

		identifierApp := javaapp.NewJavaIdentifierApp()
		iNodes := identifierApp.AnalysisFiles(files)

		callApp := javaapp.NewJavaFullApp()
		classNodes := callApp.AnalysisFiles(iNodes, files)

		app := deps.NewDepApp()

		results := app.AnalysisPath(path, classNodes)
		fmt.Fprintln(output, "unused")
		table := cmd_util.NewOutput(output)
		table.SetHeader([]string{"GroupId", "ArtifactId", "Scope"})
		for _, dep := range results {
			table.Append([]string{dep.GroupId, dep.ArtifactId, dep.Scope})
		}
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(depsCmd)

	depsCmd.PersistentFlags().StringVarP(&depCmdConfig.Path, "path", "p", ".", "example -p core/main")
}
