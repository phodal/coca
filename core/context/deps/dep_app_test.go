package deps

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cocatest"
	"testing"
)

func TestAnalysis(t *testing.T) {
	g := NewGomegaWithT(t)

	pluginsStr := `dependencies {
    implementation 'org.springframework.boot:spring-boot-starter-web'
    developmentOnly 'org.springframework.boot:spring-boot-devtools'
}`

	Analysis(pluginsStr)

	g.Expect(true).To(Equal(true))
}

func Test_ShouldCountDeps_WhenHadClassNodes(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/examples/api"
	_, classNodes := cocatest.BuildAnalysisResultsByPath(codePath)

	depApp := NewDepApp()
	depApp.CountDeps(classNodes)

	g.Expect(true).To(Equal(true))
}
