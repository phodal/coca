package bs

import (
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	. "github.com/phodal/coca/bs/models"
	. "github.com/phodal/coca/language/java"
	. "github.com/phodal/coca/utils"
)

var nodeInfos []JFullClassNode

type BadSmellModel struct {
	File string
	Line string
	Bs   string
}

type BadSmellApp struct {
}

func (j *BadSmellApp) AnalysisPath(codeDir string) []BadSmellModel {
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
	WriteToFile("nodeInfos.json", string(bsModel))

	bsList := analysisBadSmell(nodeInfos)

	return bsList
}

func analysisBadSmell(nodes []JFullClassNode) []BadSmellModel {
	var badSmellList []BadSmellModel
	for _, node := range nodes {
		// To be Defined number
		if node.Type == "Class" && len(node.Methods) < 1 {
			badSmellList = append(badSmellList, *&BadSmellModel{node.Path, "", "lazyElement"})
		}

		onlyHaveGetterAndSetter := true
		// Long Method
		for _, method := range node.Methods {
			if method.StopLine-method.StartLine > 50 {
				longMethod := &BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longMethod"}
				badSmellList = append(badSmellList, *longMethod)
			}

			if strings.Contains(method.Name, "get") && strings.Contains(method.Name, "set") {
				onlyHaveGetterAndSetter = false
			}

			// longParameterList
			if len(method.Parameters) > 6 {
				longParams := &BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "longParameterList"}
				badSmellList = append(badSmellList, *longParams)
			}

			// longParameterList
			if method.MethodBs.IfSize > 8 || method.MethodBs.SwitchSize > 8 {
				longParams := &BadSmellModel{node.Path, strconv.Itoa(method.StartLine), "repeatedSwitches"}
				badSmellList = append(badSmellList, *longParams)
			}
		}

		fmt.Println(onlyHaveGetterAndSetter, node.Type, len(node.Methods))
		if onlyHaveGetterAndSetter && node.Type == "Class" && len(node.Methods) > 0 {
			dataClass := &BadSmellModel{node.Path, "", "dataClass"}
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
				badSmellList = append(badSmellList, *&BadSmellModel{node.Path, "", "refusedBequest"})
			}
		}

		// LargeClass
		if node.Type == "Class" && len(node.Methods) > 30 {
			badSmellList = append(badSmellList, *&BadSmellModel{node.Path, "", "largeClass"})
		}
	}

	return badSmellList
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
