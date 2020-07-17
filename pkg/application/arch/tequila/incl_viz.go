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

func (fullGraph *FullGraph) ToDot(split string, include func(string) bool) *gographviz.Graph {
	graph := gographviz.NewGraph()
	_ = graph.SetName("G")

	nodeIndex := 1
	layerIndex := 1
	nodes := make(map[string]string)

	layerMap := make(map[string][]string)

	for nodeKey := range fullGraph.NodeList {
		if !include(nodeKey) && !include(fullGraph.NodeList[nodeKey]) {
			continue
		}

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

	for layer := range layerMap {
		layerAttr := make(map[string]string)
		layerAttr["label"] = "\"" + layer + "\""
		layerName := "cluster" + strconv.Itoa(layerIndex)
		_ = graph.AddSubGraph("G", layerName, layerAttr)
		layerIndex++
		for _, node := range layerMap[layer] {
			attrs := make(map[string]string)
			fileName := strings.Replace(node, layer+split, "", -1)
			attrs["label"] = "\"" + fileName + "\""
			attrs["shape"] = "box"
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
