package tequila

import (
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

func (fullGraph *FullGraph) MergeHeaderFile(merge func(string) string) *FullGraph {
	result := &FullGraph{
		NodeList:     make(map[string]string),
		RelationList: make(map[string]*Relation),
	}
	nodes := make(map[string]string)

	for key := range fullGraph.NodeList {
		mergedKey := merge(key)
		nodes[key] = mergedKey
		result.NodeList[mergedKey] = mergedKey
	}
	for key := range fullGraph.RelationList {
		relation := fullGraph.RelationList[key]
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

func (fullGraph *FullGraph) SortedByFan(merge func(string) string) []*Fan {
	mergedGraph := fullGraph.MergeHeaderFile(merge)
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

func (fullGraph *FullGraph) ToMapDot(include func(string) bool) *gographviz.Graph {
	node := fullGraph.BuildMapTree(include)
	dot := fullGraph.MapToGraph(node)
	return dot
}

func (fullGraph *FullGraph) MapToGraph(trie *PathTrie) *gographviz.Graph {
	graph := gographviz.NewGraph()
	_ = graph.SetName("G")

	nodes := make(map[string]string)
	fullGraph.layerIndex = 1
	fullGraph.nodeIndex = 1

	for _, child := range trie.Children {
		fullGraph.buildGraphNode("G", child, graph, nodes, "")
	}

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

func (fullGraph *FullGraph) buildGraphNode(subgraph string, current *PathTrie, graph *gographviz.Graph, nodes map[string]string, s string) {
	if s != "" {
		s = s + "." + current.Value
	} else {
		s = s + current.Value
	}

	layerAttr, layerName := buildLayerAttr(current.Value, fullGraph.layerIndex)
	_ = graph.AddSubGraph(subgraph, layerName, layerAttr)
	fullGraph.layerIndex++

	if len(current.Children) > 0 {
		for _, child := range current.Children {
			fullGraph.buildGraphNode(layerName, child, graph, nodes, s)
		}
	} else {
		_ = graph.AddNode(subgraph, "node"+strconv.Itoa(fullGraph.nodeIndex), fullGraph.buildRelationAttr(current.Value))
		nodes[s] = "node" + strconv.Itoa(fullGraph.nodeIndex)
		fullGraph.nodeIndex++
	}
}

func (fullGraph *FullGraph) BuildMapTree(include func(key string) bool) *PathTrie {
	pkgTrie := NewPathTrie()
	for nodeKey := range fullGraph.NodeList {
		if include(nodeKey) || include(fullGraph.NodeList[nodeKey]) {
			pkgTrie.Put(strings.ReplaceAll(nodeKey, ".", "/"))
		}
	}

	return pkgTrie
}
