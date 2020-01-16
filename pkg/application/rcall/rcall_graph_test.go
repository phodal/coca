package rcall

import (
	"encoding/json"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"log"
	"testing"

	. "github.com/onsi/gomega"
)

func MockWriteCallMap(rcallMap map[string][]string) {

}

func TestRCallGraph_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := NewRCallGraph()
	file := cmd_util.ReadFile("../../../_fixtures/call/call_api_test.json")
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	content := analyser.Analysis("com.phodal.pholedge.book.BookService.createBook", parsedDeps, MockWriteCallMap)

	g.Expect(content).To(Equal(`digraph G {
rankdir = LR;
edge [dir="back"];
"com.phodal.pholedge.book.BookService.createBook" -> "com.phodal.pholedge.book.BookController.createBook";
}
`))
}

func TestRCallGraph_Constructor(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := NewRCallGraph()
	file := cmd_util.ReadFile("../../../_fixtures/rcall/constructor_call.json")
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	content := analyser.Analysis("com.phodal.coca.analysis.JavaCallApp.parse", parsedDeps, MockWriteCallMap)

	// Todo bug: to be fix
	g.Expect(content).To(Equal(`digraph G {
rankdir = LR;
edge [dir="back"];
"com.phodal.coca.analysis.JavaCallApp.parse" -> "com.phodal.coca.analysis.JavaCallApp.analysisDir";
}
`))
}
