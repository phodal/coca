package bs

import (
	"coca/src/support"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	. "coca/src/bs/models"
	. "coca/src/language/java"
)

var nodeInfos []JFullClassNode

type BadSmellModel struct {
	File        string `json:"File,omitempty"`
	Line        string `json:"Line,omitempty"`
	Bs          string `json:"BS,omitempty"`
	Description string `json:"Description,omitempty"`
}

type BadSmellApp struct {
}

func (j *BadSmellApp) AnalysisPath(codeDir string, ignoreRules []string) []BadSmellModel {
	nodeInfos = nil
	files := (*BadSmellApp)(nil).javaFiles(codeDir)
	for index := range files {
		nodeInfo := NewJFullClassNode()
		file := files[index]

		displayName := filepath.Base(file)
		fmt.Println("Start parse java call: " + displayName)

		parser := (*BadSmellApp)(nil).processFile(file)
		context := parser.CompilationUnit()

		listener := NewBadSmellListener()

		antlr.NewParseTreeWalker().Walk(listener, context)

		nodeInfo = listener.getNodeInfo()
		nodeInfo.Path = file
		nodeInfos = append(nodeInfos, *nodeInfo)
	}

	bsModel, _ := json.MarshalIndent(nodeInfos, "", "\t")
	support.WriteToFile("nodeInfos.json", string(bsModel))

	bsList := analysisBadSmell(nodeInfos)

	mapIgnoreRules := make(map[string]bool)
	for _, ignore := range ignoreRules {
		mapIgnoreRules[ignore] = true
	}

	filteredBsList := filterBadsmellList(bsList, mapIgnoreRules)
	return filteredBsList
}

func filterBadsmellList(models []BadSmellModel, ignoreRules map[string]bool) []BadSmellModel {
	var results []BadSmellModel
	for _, model := range models {
		if !ignoreRules[model.Bs] {
			results = append(results, model)
		}
	}
	return results
}

func analysisBadSmell(nodes []JFullClassNode) []BadSmellModel {
	var badSmellList []BadSmellModel
	for _, node := range nodes {
		// To be Defined number
		if node.Type == "Class" && len(node.Methods) < 1 {
			badSmellList = append(badSmellList, *&BadSmellModel{node.Path, "", "lazyElement", ""})
		}

		onlyHaveGetterAndSetter := true
		// Long Method
		for _, method := range node.Methods {
			methodLength := method.StopLine - method.StartLine
			if methodLength > 30 {
				description := "method length: " + strconv.Itoa(methodLength)
				longMethod := &BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longMethod", description}
				badSmellList = append(badSmellList, *longMethod)
			}

			if strings.Contains(method.Name, "get") && strings.Contains(method.Name, "set") {
				onlyHaveGetterAndSetter = false
			}

			// longParameterList
			if len(method.Parameters) > 6 {
				paramsJson, _ := json.Marshal(method.Parameters)
				str := string(paramsJson[:])
				longParams := &BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longParameterList", str}
				badSmellList = append(badSmellList, *longParams)
			}

			// repeatedSwitches
			if method.MethodBs.IfSize > 8 || method.MethodBs.SwitchSize > 8 {
				longParams := &BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "repeatedSwitches", ""}
				badSmellList = append(badSmellList, *longParams)
			}
		}

		// dataClass
		if onlyHaveGetterAndSetter && node.Type == "Class" && len(node.Methods) > 0 {
			dataClass := &BadSmellModel{node.Path, "", "dataClass", ""}
			badSmellList = append(badSmellList, *dataClass)
		}

		//Refused Bequest
		if node.Extends != "" {
			hasCallParentMethod := false
			for _, methodCall := range node.MethodCalls {
				if methodCall.Class == node.Extends {
					hasCallParentMethod = true
				}
			}

			if !hasCallParentMethod {
				badSmellList = append(badSmellList, *&BadSmellModel{node.Path, "", "refusedBequest", ""})
			}
		}

		// LargeClass
		normalClassLength := withOutGetterSetterClass(node.Methods)
		if node.Type == "Class" && normalClassLength > 20 {
			description := "methods number (without getter/setter): " + strconv.Itoa(normalClassLength)
			badSmellList = append(badSmellList, *&BadSmellModel{node.Path, "", "largeClass", description})
		}
	}

	return badSmellList
}

func withOutGetterSetterClass(fullMethods []JFullMethod) int {
	var normalMethodSize = 0
	for _, method := range fullMethods {
		if !strings.Contains(method.Name, "get") && !strings.Contains(method.Name, "set") {
			normalMethodSize++
		}
	}

	return normalMethodSize
}

func (j *BadSmellApp) javaFiles(codeDir string) []string {
	files := make([]string, 0)
	_ = filepath.Walk(codeDir, func(path string, fi os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".java") && !strings.Contains(path, "Test.java") {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func (j *BadSmellApp) processFile(path string) *JavaParser {
	is, _ := antlr.NewFileStream(path)
	lexer := NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0);
	parser := NewJavaParser(stream)
	return parser
}
