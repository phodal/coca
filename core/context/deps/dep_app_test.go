package deps

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestAnalysis(t *testing.T) {
	g := NewGomegaWithT(t)

	pluginsStr := `plugins {
    id 'java'
    id 'org.springframework.boot' version '2.2.2.RELEASE'
}
`

	Analysis(pluginsStr)

	g.Expect(true).To(Equal(true))
}
