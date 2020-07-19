package tequila

import (
	"fmt"
	"github.com/awalterschulze/gographviz"
	"sort"
	"strconv"
	"strings"
)

type Relation struct {
	From  string
	To    string
	Style string
}

type FullGraph struct {
	layerIndex   int
	nodeIndex    int
	NodeList     map[string]string
	RelationList map[string]*Relation
}

type Fan struct {
	Name   string
	FanIn  int
	FanOut int
}

func (f *FullGraph) MergeHeaderFile(merge func(string) string) *FullGraph {
	result := &FullGraph{
		NodeList:     make(map[string]string),
		RelationList: make(map[string]*Relation),
	}
	nodes := make(map[string]string)

	for key := range f.NodeList {
		mergedKey := merge(key)
		nodes[key] = mergedKey
		result.NodeList[mergedKey] = mergedKey
	}
	for key := range f.RelationList {
		relation := f.RelationList[key]
		mergedFrom := merge(relation.From)
		mergedTo := merge(relation.To)
		if mergedFrom == mergedTo {
			continue
		}

		mergedRelation := &Relation{
			From:  mergedFrom,
			To:    mergedTo,
			Style: "\"solid\"",
		}

		result.RelationList[mergedRelation.From+mergedRelation.To] = mergedRelation
	}
	return result
}

func (f *FullGraph) SortedByFan(merge func(string) string) []*Fan {
	mergedGraph := f.MergeHeaderFile(merge)
	result := make([]*Fan, len(mergedGraph.NodeList))
	index := 0
	fanMap := make(map[string]*Fan)
	for key := range mergedGraph.NodeList {
		fan := &Fan{Name: key}
		result[index] = fan
		fanMap[key] = fan
		index++
	}
	for key := range mergedGraph.RelationList {
		relation := mergedGraph.RelationList[key]
		fanMap[relation.From].FanOut++
		fanMap[relation.To].FanIn++
	}
	sort.Slice(result, func(i, j int) bool {
		return (result[i].FanIn + result[i].FanOut) > (result[j].FanIn + result[j].FanOut)
	})
	return result
}

func buildLayerAttr(layer string, layerIndex int) (map[string]string, string) {
	layerAttr := make(map[string]string)
	layerAttr["label"] = "\"" + layer + "\""
	layerName := "cluster" + strconv.Itoa(layerIndex)
	return layerAttr, layerName
}

func (fullGraph *FullGraph) buildRelationAttr(text string) map[string]string {
	attrs := make(map[string]string)
	attrs["label"] = "\"" + text + "\""
	attrs["shape"] = "box"
	return attrs
}

func (fullGraph *FullGraph) ToDot(split string, include func(string) bool) *gographviz.Graph {
	graph := gographviz.NewGraph()
	_ = graph.SetName("G")

	nodeIndex := 1
	layerIndex := 1
	nodes := make(map[string]string)

	layerMap := make(map[string][]string)

	for nodeKey := range fullGraph.NodeList {
		if include(nodeKey) || include(fullGraph.NodeList[nodeKey]) {
			tmp := strings.Split(nodeKey, split)
			packageName := tmp[0]
			if packageName == nodeKey {
				packageName = "main"
			}
			if len(tmp) > 2 {
				packageName = strings.Join(tmp[0:len(tmp)-1], split)
			}

			if _, ok := layerMap[packageName]; !ok {
				layerMap[packageName] = make([]string, 0)
			}
			layerMap[packageName] = append(layerMap[packageName], nodeKey)
		}
	}

	for layer := range layerMap {
		layerAttr, layerName := buildLayerAttr(layer, layerIndex)
		_ = graph.AddSubGraph("G", layerName, layerAttr)
		layerIndex++
		for _, node := range layerMap[layer] {
			fileName := strings.Replace(node, layer+split, "", -1)
			attrs := fullGraph.buildRelationAttr(fileName)
			_ = graph.AddNode(layerName, "node"+strconv.Itoa(nodeIndex), attrs)
			nodes[node] = "node" + strconv.Itoa(nodeIndex)
			nodeIndex++
		}
	}

	cross := make(map[string]bool) // mapping from strings to ints
	for key := range fullGraph.RelationList {
		relation := fullGraph.RelationList[key]
		if nodes[relation.From] != "" && nodes[relation.To] != "" {
			fromNode := nodes[relation.From]
			toNode := nodes[relation.To]

			cross[fromNode+toNode] = true
			attrs := make(map[string]string)
			attrs["style"] = relation.Style

			_ = graph.AddEdge(fromNode, toNode, true, attrs)
		}
	}

	return graph
}

type GraphNode struct {
	text     string
	children []*GraphNode
}

func (fullGraph *FullGraph) BuildMapTree(split string, include func(key string) bool) *GraphNode {
	graphNode := &GraphNode{}

	for nodeKey := range fullGraph.NodeList {
		tmp := strings.Split(nodeKey, split)
		graphNode.text = tmp[0]
		graphNode = buildNode(tmp[1:], graphNode)
	}

	return graphNode
}

func (fullGraph *FullGraph) ToMapDot(node *GraphNode) *gographviz.Graph {
	graph := gographviz.NewGraph()
	_ = graph.SetName("G")

	nodes := make(map[string]string)
	fullGraph.layerIndex = 1
	fullGraph.nodeIndex = 1

	fullGraph.buildGraphNode("G", node, graph, nodes)

	for key := range fullGraph.RelationList {
		relation := fullGraph.RelationList[key]
		if nodes[relation.From] != "" && nodes[relation.To] != "" {
			fromNode := nodes[relation.From]
			toNode := nodes[relation.To]

			attrs := make(map[string]string)
			attrs["style"] = relation.Style

			_ = graph.AddEdge(fromNode, toNode, true, attrs)
		}
	}

	return graph
}

func (fullGraph *FullGraph) buildGraphNode(subgraph string, current *GraphNode, graph *gographviz.Graph, nodes map[string]string) {
	layerAttr, layerName := buildLayerAttr(current.text, fullGraph.layerIndex)
	_ = graph.AddSubGraph(subgraph, layerName, layerAttr)
	fullGraph.layerIndex++

	if len(current.children) > 0 {
		for _, child := range current.children {
			fmt.Println(current.text)
			fullGraph.buildGraphNode(layerName, child, graph, nodes)
		}
	} else {
		fmt.Println(layerName, current.text, fullGraph.nodeIndex)
		_ = graph.AddNode(layerName, "node"+strconv.Itoa(fullGraph.nodeIndex), fullGraph.buildRelationAttr(current.text))
		nodes[current.text] = "node" + strconv.Itoa(fullGraph.nodeIndex)
		fullGraph.nodeIndex++
	}
}

func buildNode(arr []string, node *GraphNode) *GraphNode {
	if node.text == arr[0] {
		return node
	}

	child := &GraphNode{}
	if len(arr) == 1 {
		child.text = arr[0]
		node.children = append(node.children, child)
	} else {
		child.text = arr[0]
		graphNode := buildNode(arr[1:], child)
		node.children = append(node.children, graphNode)
	}

	return node
}
