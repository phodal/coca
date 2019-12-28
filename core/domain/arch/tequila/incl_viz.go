package tequila

import (
	"bufio"
	"fmt"
	"github.com/awalterschulze/gographviz"
	"io/ioutil"
	"os"
	"path/filepath"
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

func (f *FullGraph) FindCrossRef(merge func(string) string) []string {
	mergedRelationMap := make(map[string]string)
	result := make([]string, 0)
	for key := range f.RelationList {
		relation := f.RelationList[key]
		mergedFrom := merge(relation.From)
		mergedTo := merge(relation.To)
		if mergedFrom == mergedTo {
			continue
		}
		if _, ok := mergedRelationMap[mergedTo+mergedFrom]; ok {
			result = append(result, mergedFrom+" <-> "+mergedTo)
		}
		mergedRelationMap[mergedFrom+mergedTo] = ""
	}
	return result
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

func (f *FullGraph) EntryPoints(merge func(string) string) []string {
	mergedGraph := f.MergeHeaderFile(merge)
	fromMap := make(map[string]bool)
	toMap := make(map[string]bool)
	for key := range mergedGraph.RelationList {
		relation := mergedGraph.RelationList[key]
		if relation.From == "main" {
			continue
		}
		fromMap[relation.From] = true
		toMap[relation.To] = true
	}
	result := make([]string, 0)
	for key := range fromMap {
		if _, ok := toMap[key]; !ok {
			result = append(result, key)
		}
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

var fullGraph *FullGraph

func parseRelation(edge *gographviz.Edge, nodes map[string]string) {
	if _, ok := nodes[edge.Src]; ok {
		if _, ok := nodes[edge.Dst]; ok {
			dst := nodes[edge.Dst]
			src := nodes[edge.Src]
			dst = strings.ToLower(dst)
			src = strings.ToLower(src)
			relation := &Relation{
				From:  dst,
				To:    src,
				Style: "\"solid\"",
			}
			fullGraph.RelationList[relation.From+"->"+relation.To] = relation
		}
	}
}

func filterDirectory(fullMethodName string) bool {
	if strings.Contains(fullMethodName, "_test") {
		return true
	}

	if strings.Contains(fullMethodName, "Test") {
		return true
	}

	if strings.Contains(fullMethodName, "/Library/") {
		return true
	}
	return false
}

func parseDotFile(codeDotfile string) {
	fbuf, _ := ioutil.ReadFile(codeDotfile)
	parseFromBuffer(fbuf)
}
func parseFromBuffer(fbuf []byte) {
	g, err := gographviz.Read(fbuf)
	if err != nil {
		fmt.Println(string(fbuf))
	}
	nodes := make(map[string]string)
	for _, node := range g.Nodes.Nodes {
		fullMethodName := strings.Replace(node.Attrs["label"], "\"", "", 2)
		if strings.Contains(fullMethodName, " ") {
			tmp := strings.Split(fullMethodName, " ")
			fullMethodName = tmp[len(tmp)-1]
		}
		if filterDirectory(fullMethodName) {
			continue
		}

		methodName := formatMethodName(fullMethodName)
		fullGraph.NodeList[methodName] = methodName
		nodes[node.Name] = methodName
	}
	for key := range g.Edges.DstToSrcs {
		for edgesKey := range g.Edges.DstToSrcs[key] {
			for _, edge := range g.Edges.DstToSrcs[key][edgesKey] {
				parseRelation(edge, nodes)
			}
		}
	}
}
func formatMethodName(fullMethodName string) string {
	methodName := strings.Replace(fullMethodName, "\\l", "", -1)
	methodName = strings.Replace(methodName, "src/", "", -1)
	methodName = strings.Replace(methodName, "include/", "", -1)
	methodName = strings.ToLower(methodName)
	return methodName
}

func codeDotFiles(codeDir string, fileFilter func(string) bool) []string {
	codeDotFiles := make([]string, 0)
	filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".dot") {
			if fileFilter(path) {
				//return nil
				if strings.Contains(path, "_test_") {
					return nil
				}

				codeDotFiles = append(codeDotFiles, path)
			}
		}

		return nil
	})

	return codeDotFiles
}

func ParseInclude(codeDir string) *FullGraph {
	fullGraph = &FullGraph{
		NodeList:     make(map[string]string),
		RelationList: make(map[string]*Relation),
	}
	codeDotFiles := codeDotFiles(codeDir, func(path string) bool {

		return strings.HasSuffix(path, "_dep__incl.dot")
	})

	for _, codeDotfile := range codeDotFiles {
		parseDotFile(codeDotfile)
	}

	return fullGraph
}

func (fullGraph *FullGraph) ToDot(fileName string, split string, filter func(string) bool) {
	graph := gographviz.NewGraph()
	graph.SetName("G")

	nodeIndex := 1
	layerIndex := 1
	nodes := make(map[string]string)

	layerMap := make(map[string][]string)

	for nodeKey := range fullGraph.NodeList {
		if filter(nodeKey) {
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
		graph.AddSubGraph("G", layerName, layerAttr)
		layerIndex++
		for _, node := range layerMap[layer] {
			attrs := make(map[string]string)
			fileName := strings.Replace(node, layer+split, "", -1)
			attrs["label"] = "\"" + fileName + "\""
			attrs["shape"] = "box"
			graph.AddNode(layerName, "node"+strconv.Itoa(nodeIndex), attrs)
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

			graph.AddEdge(fromNode, toNode, true, attrs)
		}
	}

	f, _ := os.Create(fileName)
	w := bufio.NewWriter(f)
	w.WriteString("di" + graph.String())
	w.Flush()
}

var Foo = func() string {
	return ""
}

func (fullGraph *FullGraph) ToDataSet(fileName string, split string, filter func(string) bool) {
	nodes := make(map[string]string)

	for nodeKey := range fullGraph.NodeList {
		if filter(nodeKey) {
			continue
		}

		nodes[nodeKey] = nodeKey
	}

	relMap := make(map[string][]string)
	for key := range fullGraph.RelationList {
		relation := fullGraph.RelationList[key]

		if nodes[relation.From] == "" && nodes[relation.To] != "" {
			if _, ok := relMap[relation.From]; !ok {
				relMap[relation.From] = make([]string, 0)
			}
			relMap[relation.From] = append(relMap[relation.From], relation.To)
		}
	}

	for key := range relMap {
		tos := relMap[key]
		fmt.Print("['" + strings.Join(tos, "','") + "'],")
	}
}
