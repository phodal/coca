package tequila

import (
	. "github.com/onsi/gomega"
	"testing"
)

func createBasicMap() (*GraphNode, *FullGraph) {
	fullGraph, nodeFilter := createGraph()

	node := fullGraph.BuildMapTree(".", nodeFilter)
	return node, fullGraph
}

func createGraph() (*FullGraph, func(key string) bool) {
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
	return fullGraph, nodeFilter
}

func Test_BuildGraphNode(t *testing.T) {
	g := NewGomegaWithT(t)
	node, _ := createBasicMap()

	g.Expect(node.text).To(Equal("com"))
	children := node.children
	g.Expect(len(children)).To(Equal(2))
}

func Test_ShouldMergeSameMap(t *testing.T) {
	g := NewGomegaWithT(t)
	fullGraph, nodeFilter := createGraph()
	fullGraph.NodeList["com.phodal.coca"] = "com.phodal.coca"
	node := fullGraph.BuildMapTree(".", nodeFilter)

	g.Expect(node.text).To(Equal("com"))
	children := node.children
	g.Expect(len(children)).To(Equal(3))
}

func Test_BuildNodeDot(t *testing.T) {
	g := NewGomegaWithT(t)
	node, graph := createBasicMap()
	dot := graph.ToMapDot(node)

	//result := dot.String()
	//cmd_util.WriteToCocaFile("demo.dot", result)

	g.Expect(len(dot.SubGraphs.SubGraphs)).To(Equal(5))
	g.Expect(len(dot.Nodes.Nodes)).To(Equal(2))
}
