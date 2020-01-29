package deps

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"github.com/phodal/coca/cocatest/testhelper"
	"testing"
)

func Test_ShouldReturnGradleDep(t *testing.T) {
	g := NewGomegaWithT(t)

	pluginsStr := `dependencies {
    implementation 'org.springframework.boot:spring-boot-starter-web'
    developmentOnly 'org.springframework.boot:spring-boot-devtools'
}`

	results := AnalysisGradleString(pluginsStr)

	g.Expect(len(results)).To(Equal(2))
	g.Expect(results[0].ArtifactId).To(Equal("spring-boot-starter-web"))
}

func Test_ShouldReturnCorrectGradleDepsFroFile(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/deps/gradle/build.gradle"
	bytes := cmd_util.ReadFile(codePath)

	mavenDeps := AnalysisGradleString(string(bytes))

	g.Expect(len(mavenDeps)).To(Equal(14))
}

func Test_ShouldHandleExclude(t *testing.T) {
	g := NewGomegaWithT(t)

	pluginsStr := `dependencies {
	testImplementation('org.springframework.boot:spring-boot-starter-test') {
		exclude group: 'org.junit.vintage', module: 'junit-vintage-engine'
		exclude module: 'junit'
    }
}`

	results := AnalysisGradleString(pluginsStr)

	g.Expect(len(results)).To(Equal(1))
	g.Expect(results[0].ArtifactId).To(Equal("spring-boot-starter-test"))
	g.Expect(results[0].GroupId).To(Equal("org.springframework.boot"))
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

	codePath := "../../../_fixtures/grammar/java/examples/api/"
	classNodes, _, _ := testhelper.BuildAnalysisDeps(codePath)

	depApp := NewDepApp()
	importMap := depApp.BuildImportMap(classNodes)

	g.Expect(len(importMap)).To(Equal(25))
}

func Test_ListUnusedImportForOneGradleFile(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/deps/maven_sample/"
	classNodes, _, _ := testhelper.BuildAnalysisDeps(codePath)

	mavenDeps := AnalysisMaven(codePath + "pom.xml")
	g.Expect(len(mavenDeps)).To(Equal(6))

	depApp := NewDepApp()
	deps := depApp.AnalysisPath(codePath, classNodes)

	g.Expect(len(deps)).To(Equal(3))
	g.Expect(deps[0].GroupId).To(Equal("org.flywaydb"))
	g.Expect(deps[1].GroupId).To(Equal("mysql"))
	g.Expect(deps[2].GroupId).To(Equal("org.springframework.cloud"))
}
