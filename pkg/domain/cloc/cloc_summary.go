package cloc

import (
	"encoding/json"
	"fmt"
	"github.com/boyter/scc/processor"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
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

func BuildLanguageMap(languageMap map[string]map[string]processor.LanguageSummary, keys []string, filePath string) {
	var dirLangSummary []processor.LanguageSummary
	contents, _ := ioutil.ReadFile(filePath)
	err := json.Unmarshal(contents, &dirLangSummary)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	dirName := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	languageMap[dirName] = make(map[string]processor.LanguageSummary)

	for _, key := range keys {
		var hasSet = false
		for _, langSummary := range dirLangSummary {
			if key == langSummary.Name {
				hasSet = true
				langSummary.Name = key
				languageMap[dirName][key] = langSummary
				break
			}
		}
		if !hasSet {
			languageMap[dirName][key] = processor.LanguageSummary{}
		}
	}
}
