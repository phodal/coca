package tequila

import (
	. "github.com/onsi/gomega"
	"github.com/phodal/coca/cmd/cmd_util"
	"strings"
	"testing"
)

func createBasicMap() (*PathTrie, *FullGraph) {
	fullGraph, nodeFilter := createGraph()

	node := fullGraph.BuildMapTree(nodeFilter)
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

	g.Expect(len(node.Children["com"].Children)).To(Equal(2))
}

func Test_ShouldMergeSameMap(t *testing.T) {
	g := NewGomegaWithT(t)
	fullGraph, nodeFilter := createGraph()
	fullGraph.NodeList["com.phodal.coca"] = "com.phodal.coca"
	node := fullGraph.BuildMapTree(nodeFilter)

	g.Expect(len(node.Children["com"].Children)).To(Equal(2))
}

func Test_BuildNodeDot(t *testing.T) {
	g := NewGomegaWithT(t)
	graph, nodeFilter := createGraph()
	graph.NodeList["com.phodal.coca"] = "com.phodal.coca"
	node := graph.BuildMapTree(nodeFilter)
	dot := graph.MapToGraph(node)

	result := dot.String()
	cmd_util.WriteToCocaFile("demo.dot", result)

	g.Expect(len(dot.SubGraphs.SubGraphs)).To(Equal(6))
}

func Test_ShouldShowPackageOnly(t *testing.T) {
	g := NewGomegaWithT(t)

	fullGraph, nodeFilter := createGraph()

	fullGraph= fullGraph.MergeHeaderFile(MergeHeaderFunc)
	node := fullGraph.ToMapDot(nodeFilter)

	g.Expect(strings.Contains(node.String(), "Ledge")).To(Equal(false))
	g.Expect(strings.Contains(node.String(), "Boot")).To(Equal(false))
}

func Test_ShouldShowMergePackage(t *testing.T) {
	g := NewGomegaWithT(t)

	fullGraph, nodeFilter := createGraph()

	fullGraph= fullGraph.MergeHeaderFile(MergePackageFunc)
	node := fullGraph.ToMapDot(nodeFilter)

	g.Expect(strings.Contains(node.String(), "Ledge")).To(Equal(false))
	g.Expect(strings.Contains(node.String(), "Boot")).To(Equal(false))
}

func Test_ShouldConvertDot(t *testing.T) {
	g := NewGomegaWithT(t)

	fullGraph, nodeFilter := createGraph()

	node := fullGraph.ToDot(".", nodeFilter)

	g.Expect(strings.Contains(node.String(), "Ledge")).To(Equal(true))
}
