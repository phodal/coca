package cmd_util

import (
	"github.com/olekukonko/tablewriter"
	"io"
)

func NewOutput(output io.Writer) *tablewriter.Table {
	table := tablewriter.NewWriter(output)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetColWidth(80)
	return table
}
