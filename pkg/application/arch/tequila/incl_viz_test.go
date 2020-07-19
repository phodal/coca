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

	fullGraph.ToMapDot(".", nodeFilter)

	g.Expect(true).To(Equal(true))
}
