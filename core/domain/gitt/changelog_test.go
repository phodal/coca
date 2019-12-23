package gitt

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestBuildChangeMap(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BuildMessageByInput(`
[c24069b] Phodal HUANG 2019-10-25 fix: fix test
7       0       README.md

[c24069b] Phodal HUANG 2019-10-25 feat: add README.md
7       0       README.md

`)

	buildChangeMap := BuildChangeMap(result)
	g.Expect(buildChangeMap["feat"]["README.md"]).To(Equal(1))
	g.Expect(buildChangeMap["fix"]["README.md"]).To(Equal(1))
}