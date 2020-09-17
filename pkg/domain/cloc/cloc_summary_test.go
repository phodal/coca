package cloc

import (
	"encoding/json"
	"fmt"
	"github.com/boyter/scc/processor"
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func Test_parser_json_languages(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var data = `[
  {"Name":"Java","Bytes":21169200,"CodeBytes":0,"Lines":540043,"Code":381028,"Comment":93196,"Blank":65819,"Complexity":43899,"Count":4435,"WeightedComplexity":0,"Files":[]},
  {"Name":"Kotlin","Bytes":6961705,"CodeBytes":0,"Lines":168900,"Code":118448,"Comment":30743,"Blank":19709,"Complexity":7636,"Count":1315,"WeightedComplexity":0,"Files":[]}
]
`

	var f []processor.LanguageSummary
	err := json.Unmarshal([]byte(data), &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	fmt.Println(f)

	g.Expect(len(f)).To(Equal(2))
}

func TestGenerateCsv(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var languageMap = make(map[string]map[string]processor.LanguageSummary)
	var keys = []string{"Java", "Kotlin"}
	languageMap["home"] = make(map[string]processor.LanguageSummary)
	languageMap["home"]["Java"] = processor.LanguageSummary{}

	data := BuildClocCsvData(languageMap, keys)

	g.Expect(data[0]).To(Equal([]string{"package", "summary", "Java", "Kotlin"}))
	g.Expect(data[1]).To(Equal([]string{"home", "0", "0", "0"}))
}

func TestShouldBuildLanguageMap(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/cloc/summary/android.json"
	codePath = filepath.FromSlash(codePath)

	var keys = []string{"Java", "Kotlin", "Phodal"}
	languageMap := make(map[string]map[string]processor.LanguageSummary)

	BuildLanguageMap(languageMap, keys, codePath)

	g.Expect(len(languageMap["android"])).To(Equal(3))
	g.Expect(languageMap["android"]["Java"].Code).To(Equal(int64(381639)))
	g.Expect(languageMap["android"]["Kotlin"].Code).To(Equal(int64(123963)))
	g.Expect(languageMap["android"]["Phodal"].Code).To(Equal(int64(0)))
}
