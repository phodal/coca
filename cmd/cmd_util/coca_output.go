package cmd_util

import (
	"github.com/olekukonko/tablewriter"
	"io"
	"strings"
)

func NewOutput(output io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(output)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetColWidth(80)
	return table
}

func NewCsv() (*tablewriter.Table, *strings.Builder) {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetCenterSeparator("")
	table.SetRowSeparator("")
	table.SetColumnSeparator(",")
	table.SetColWidth(80)
	return table, tableString
}
