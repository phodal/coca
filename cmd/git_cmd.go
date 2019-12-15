package cmd

import (
	. "coca/src/gitt"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	relatedConfig string
)

var gitCmd *cobra.Command = &cobra.Command{
	Use:   "ga",
	Short: "git analysis",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		message := BuildCommitMessage()
		isFullMessage := cmd.Flag("basic").Value.String() == "true"

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

			if len(teamSummary) > 20 && isFullMessage {
				teamSummary = teamSummary[:20]
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

			if len(age) > 20 && isFullMessage {
				age = age[:20]
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

			if len(authors) > 20 && isFullMessage {
				authors = authors[:20]
			}
			for _, v := range authors {
				table.Append([]string{v.Name, strconv.Itoa(v.CommitCount), strconv.Itoa(v.LineCount)})
			}
			table.Render()
		}

		if relatedConfig != "" {
			config, err := ioutil.ReadFile(relatedConfig)
			if err != nil {
				_ = fmt.Errorf("🥯  🦆 🦉 🥓 🦄 🦀 🖕 🍣 🍤 🍥 , lost related json", err)
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
	gitCmd.PersistentFlags().StringVar(&relatedConfig, "r", "", "related")
}
