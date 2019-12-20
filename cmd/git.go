package cmd

import (
	. "coca/core/domain/gitt"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"
)

type GitCmdConfig struct {
	Size int
}

var (
	relatedConfig string
	gitCmdConfig  GitCmdConfig
)

var gitCmd *cobra.Command = &cobra.Command{
	Use:   "git",
	Short: "git analysis",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		message := BuildCommitMessage()
		isFullMessage := cmd.Flag("basic").Value.String() == "true"
		size := gitCmdConfig.Size

		if cmd.Flag("basic").Value.String() == "true" {
			table := tablewriter.NewWriter(os.Stdout)

			basicSummary := BasicSummary(message)
			table.SetHeader([]string{"Statistic", "Number"})
			table.Append([]string{"Commits", strconv.Itoa(basicSummary.Commits)})
			table.Append([]string{"Entities", strconv.Itoa(basicSummary.Entities)})
			table.Append([]string{"Changes", strconv.Itoa(basicSummary.Changes)})
			table.Append([]string{"Authors", strconv.Itoa(basicSummary.Authors)})
			table.Render()
		}

		if cmd.Flag("team").Value.String() == "true" {
			table := tablewriter.NewWriter(os.Stdout)

			teamSummary := GetTeamSummary(message)
			table.SetHeader([]string{"EntityName", "RevsCount", "AuthorCount"})

			if len(teamSummary) > size && isFullMessage {
				teamSummary = teamSummary[:size]
			}
			for _, v := range teamSummary {
				table.Append([]string{v.EntityName, strconv.Itoa(v.RevsCount), strconv.Itoa(v.AuthorCount)})
			}
			table.Render()
		}

		if cmd.Flag("age").Value.String() == "true" {
			table := tablewriter.NewWriter(os.Stdout)

			age := CalculateCodeAge(message)
			table.SetHeader([]string{"File", "Month"})

			if len(age) > size && isFullMessage {
				age = age[:size]
			}
			for _, v := range age {
				table.Append([]string{v.File, v.Month})
			}
			table.Render()
		}

		if cmd.Flag("top").Value.String() == "true" {
			table := tablewriter.NewWriter(os.Stdout)

			authors := GetTopAuthors(message)
			table.SetHeader([]string{"Author", "CommitCount", "LineCount"})

			if len(authors) > size && isFullMessage {
				authors = authors[:size]
			}
			for _, v := range authors {
				table.Append([]string{v.Name, strconv.Itoa(v.CommitCount), strconv.Itoa(v.LineCount)})
			}
			table.Render()
		}

		if relatedConfig != "" {
			config, err := ioutil.ReadFile(relatedConfig)
			if err != nil {
				_ = fmt.Errorf("lost related json %s", err)
				return
			}

			GetRelatedFiles(message, config)
			//results := GetRelatedFiles(message, config)
			//fmt.Println(results)
		}
	},
}

func init() {
	rootCmd.AddCommand(gitCmd)

	gitCmd.PersistentFlags().BoolP("basic", "b", false, "Basic Summary")
	gitCmd.PersistentFlags().BoolP("team", "t", false, "Team Summary")
	gitCmd.PersistentFlags().BoolP("age", "a", false, "Code Age")
	gitCmd.PersistentFlags().BoolP("top", "o", false, "Top Authors")
	gitCmd.PersistentFlags().BoolP("full", "f", false, "full")
	gitCmd.PersistentFlags().IntVarP(&gitCmdConfig.Size, "size", "s", 20, "full")
	gitCmd.PersistentFlags().StringVar(&relatedConfig, "r", "", "related")
}
