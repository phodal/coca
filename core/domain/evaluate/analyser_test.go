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

	result := analyser.Analysis(parsedDeps, nil)
	g.Expect(len(result.ServiceIssues.LifecycleMap["do"])).To(Equal(2))
	g.Expect(result.ServiceIssues.LifecycleMap["do"][0]).To(Equal("doSave"))
	g.Expect(result.ServiceIssues.LifecycleMap["do"][1]).To(Equal("doUpdate"))
}

func Test_Service_Same_Return_Type(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewEvaluateAnalyser()
	file := support.ReadFile("../../../_fixtures/evaluate/service_same_return_type.json")
	_ = json.Unmarshal(file, &parsedDeps)

	results := analyser.Analysis(parsedDeps, nil)
	g.Expect(len(results.ServiceIssues.ReturnTypeMap)).To(Equal(1))
}

func Test_Long_Parameters(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []models.JClassNode
	analyser := NewEvaluateAnalyser()
	file := support.ReadFile("../../../_fixtures/evaluate/service_long_parameters.json")
	_ = json.Unmarshal(file, &parsedDeps)

	result := analyser.Analysis(parsedDeps, nil)

	g.Expect(result.ServiceIssues.RelatedMethod[0]).To(Equal("address"))
	g.Expect(result.ServiceIssues.RelatedMethod[1]).To(Equal("age"))
	g.Expect(result.ServiceIssues.RelatedMethod[2]).To(Equal("firstname"))
	g.Expect(result.ServiceIssues.RelatedMethod[3]).To(Equal("lastname"))
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
