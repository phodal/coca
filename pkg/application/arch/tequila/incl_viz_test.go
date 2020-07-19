package tequila

import (
	. "github.com/onsi/gomega"
	"testing"
)

func Test_VisualDemo(t *testing.T) {
	g := NewGomegaWithT(t)
	fullGraph := &FullGraph{
		NodeList:     make(map[string]string),
		RelationList: make(map[string]*Relation),
	}
	from := "com.phodal.Ledge"
	to := "com.spring.Boot"

	fullGraph.NodeList[from] = from
	fullGraph.NodeList[to] = to

	relation := Relation{
		From:  from,
		To:    to,
		Style: "\"solid\"",
	}
	fullGraph.RelationList["com.phodal.Ledge->com.spring.Boot"] = &relation

	var nodeFilter = func(key string) bool {
		return true
	}

	node := fullGraph.ToMapDot(".", nodeFilter)

	g.Expect(node.text).To(Equal("com"))
	children := node.children
	g.Expect(len(children)).To(Equal(2))
	g.Expect(children[0].text).To(Equal("phodal"))
	g.Expect(children[1].text).To(Equal("spring"))
}
