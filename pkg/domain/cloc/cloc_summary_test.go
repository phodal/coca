package cloc

import (
	"encoding/json"
	"fmt"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
	"log"
	"testing"
	"github.com/boyter/scc/processor"
)

func Test_Yaml_Parse_Model(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var data = `
header:
  url: https://github.com/boyter/scc/
  version: 2.13.0
  elapsed_seconds: 0.006
  n_files: 25
  n_lines: 4045
  files_per_second: 4166.666666666667
  lines_per_second: 674166.6666666666

`

	var header = LanguageReportStart{}
	err := yaml.Unmarshal([]byte(data), &header)
	if err != nil {
		log.Fatalf("error: %v", header)
	}

	g.Expect(header.Header.Version).To(Equal("2.13.0"))
}

func Test_Cloc_Yaml_File_Parse_Model(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	var data = `
# https://github.com/boyter/scc/
header:
  url: https://github.com/boyter/scc/
  version: 2.13.0
  elapsed_seconds: 0.006
  n_files: 25
  n_lines: 4045
  files_per_second: 4166.666666666667
  lines_per_second: 674166.6666666666
Java:
  name: Java
  code: 3010
  comment: 516
  blank: 519
  nFiles: 25
SUM:
  code: 3010
  comment: 516
  blank: 519
  nFiles: 25

`

	var header = ClocSummary{}
	err := yaml.Unmarshal([]byte(data), &header)
	if err != nil {
		log.Fatalf("error: %v", header)
	}

	g.Expect(header.Header.Version).To(Equal("2.13.0"))
	g.Expect(header.Sum.Code).To(Equal(int64(3010)))
}

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
