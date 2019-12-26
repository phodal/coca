package evaluate

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/core/adapter/call"
	"github.com/phodal/coca/core/adapter/identifier"
	"github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/support"
	"testing"
)

func TestAnalyser_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewEvaluateAnalyser()
	file := support.ReadFile("../../../_fixtures/evaluate/service.json")
	_ = json.Unmarshal(file, &parsedDeps)

	analyser.Analysis(parsedDeps, nil)
	g.Expect(true).To(Equal(true))
}

func Test_Service_LifeCycle(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewEvaluateAnalyser()
	file := support.ReadFile("../../../_fixtures/evaluate/service_lifecycle.json")
	_ = json.Unmarshal(file, &parsedDeps)

	analyser.Analysis(parsedDeps, nil)
	// Todo Refactoring use read/write file
	g.Expect(true).To(Equal(true))
}

func Test_Service_Same_Return_Type(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewEvaluateAnalyser()
	file := support.ReadFile("../../../_fixtures/evaluate/service_same_return_type.json")
	_ = json.Unmarshal(file, &parsedDeps)

	analyser.Analysis(parsedDeps, nil)

	// Todo Refactoring use read/write file
	g.Expect(true).To(Equal(true))
}

func Test_Long_Parameters(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewEvaluateAnalyser()
	file := support.ReadFile("../../../_fixtures/evaluate/service_long_parameters.json")
	_ = json.Unmarshal(file, &parsedDeps)

	analyser.Analysis(parsedDeps, nil)

	// Todo Refactoring use read/write file
	g.Expect(true).To(Equal(true))
}

func TestNullPointException(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/evaluate/null"
	identifierApp := new(identifier.JavaIdentifierApp)
	identifiers := identifierApp.AnalysisPath(codePath)
	var classes []string = nil
	for _, node := range identifiers {
		classes = append(classes, node.Package+"."+node.ClassName)
	}

	callApp := call.NewJavaCallApp()
	callNodes := callApp.AnalysisPath(codePath, classes, identifiers)

	analyser := NewEvaluateAnalyser()
	result := analyser.Analysis(callNodes, identifiers)

	g.Expect(result.Nullable.Items[0]).To(Equal("nonnull.Name.orElseNull"))
}
