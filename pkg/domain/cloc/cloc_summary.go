package cloc

import (
	"github.com/boyter/scc/processor"
	"strconv"
)

func BuildClocCsvData(languageMap map[string]map[string]processor.LanguageSummary, keys []string) [][]string {
	var data [][]string
	baseKey := []string{"package", "summary"}
	data = append(data, append(baseKey, keys...))

	for dirName, dirSummary := range languageMap {
		var column []string
		column = append(column, dirName)

		var codes []string
		var summary int64

		for _, key := range keys {
			lang := dirSummary[key]
			summary = summary + lang.Code
			codes = append(codes, strconv.Itoa(int(lang.Code)))
		}

		column = append(column, strconv.Itoa(int(summary)))
		column = append(column, codes...)
		data = append(data, column)
	}
	return data
}

