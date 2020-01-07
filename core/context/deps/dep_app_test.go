package deps

import (
	"fmt"
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cocatest"
	"testing"
)

func Test_ShouldReturnGradleDep(t *testing.T) {
	g := NewGomegaWithT(t)

	pluginsStr := `dependencies {
    implementation 'org.springframework.boot:spring-boot-starter-web'
    developmentOnly 'org.springframework.boot:spring-boot-devtools'
}`

	results := AnalysisGradle(pluginsStr)

	g.Expect(len(results)).To(Equal(2))
	g.Expect(results[0].ArtifactId).To(Equal("spring-boot-starter-web"))
}

func Test_ShouldReturnCorrectGradleDepsFroFile(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/deps/gradle/build.gradle"
	bytes := cmd_util.ReadFile(codePath)

	mavenDeps := AnalysisGradle(string(bytes))

	g.Expect(len(mavenDeps)).To(Equal(13))
}

func Test_ShouldReturnCorrectMavenDeps(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/deps/maven/pom.xml"
	mavenDeps := AnalysisMaven(codePath)

	g.Expect(len(mavenDeps)).To(Equal(12))
}

func Test_ShouldReturnNilWhenErrorPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/god_know_it"
	mavenDeps := AnalysisMaven(codePath)

	g.Expect(len(mavenDeps)).To(Equal(0))
}

func Test_ShouldCountDeps_WhenHadClassNodes(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/examples/api/"
	classNodes, _, _ := cocatest.BuildAnalysisDeps(codePath)

	depApp := NewDepApp()
	importMap := depApp.BuildImportMap(classNodes)

	g.Expect(len(importMap)).To(Equal(25))
}
