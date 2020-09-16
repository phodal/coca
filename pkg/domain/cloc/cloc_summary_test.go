package cloc


import (
	"fmt"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
	"log"
	"testing"
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