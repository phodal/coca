package cloc

import (
	"encoding/json"
	"fmt"
	"github.com/boyter/scc/processor"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cmd/config"
	"github.com/phodal/coca/pkg/domain/cloc"
	"io/ioutil"
	"os"
	"sort"
)

func ConvertToCsv(outputFiles []string, keys []string) [][]string {
	var basemap = make(map[string]processor.LanguageSummary)
	for _, key := range keys {
		basemap[key] = processor.LanguageSummary{}
	}

	var languageMap = make(map[string]map[string]processor.LanguageSummary)
	for _, filePath := range outputFiles {
		cloc.BuildLanguageMap(languageMap, keys, filePath)
	}

	deb, _ := json.Marshal(languageMap)
	cmd_util.WriteToCocaFile("debug_cloc.json", string(deb))

	csvData := cloc.BuildClocCsvData(languageMap, keys)
	return csvData
}

func BuildBaseKey(baseDir string) []string {
	contents, _ := ioutil.ReadFile(baseDir)
	var languages []processor.LanguageSummary
	err := json.Unmarshal(contents, &languages)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	var keys []string
	for _, data := range languages {
		keys = append(keys, data.Name)
	}

	return keys
}

func CreateClocDir() error {
	os.Mkdir(config.CocaConfig.ReporterPath, os.ModePerm)
	return os.Mkdir(config.CocaConfig.ReporterPath+"/cloc/", os.ModePerm)
}

func IsIgnoreDir(baseName string) bool {
	dirs := []string{".git", ".svn", ".hg", ".idea"}
	for _, dir := range dirs {
		if dir == baseName {
			return true
		}
	}
	return false
}

func SortLangeByCode(languageSummaries []processor.LanguageSummary) {
	for _, langSummary := range languageSummaries {
		files := langSummary.Files
		sort.Slice(files, func(i, j int) bool {
			return files[i].Code > files[j].Code
		})

		langSummary.Files = files
	}
}

