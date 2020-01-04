package deps

import (
	. "github.com/onsi/gomega"
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
