package unused

import (
	models2 "coca/core/domain/refactor/base/models"
	support3 "coca/core/domain/refactor/rename/support"
	. "coca/core/models"
	"coca/core/support"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

var parsedChange []support3.RefactorChangeRelate
var nodes []models2.JMoveStruct

type RemoveMethodApp struct {
}

var depsFile string
var configPath string
var conf string
var parsedDeps []JClassNode

func RenameMethodApp(dep string, p string) *RemoveMethodApp {
	nodes = nil
	depsFile = dep
	configPath = p
	return &RemoveMethodApp{}
}

func (j *RemoveMethodApp) Start() {
	file := support.ReadFile(depsFile)
	if file == nil {
		return
	}

	_ = json.Unmarshal(file, &parsedDeps)

	configBytes := support.ReadFile(configPath)
	if configBytes == nil {
		return
	}

	conf = string(configBytes)

	parsedChange = support3.ParseRelates(conf)

	startParse(parsedDeps, parsedChange)
}

func startParse(nodes []JClassNode, relates []support3.RefactorChangeRelate) {

	for _, pkgNode := range nodes {
		for _, related := range relates {
			oldInfo := support3.BuildMethodPackageInfo(related.OldObj)
			newInfo := support3.BuildMethodPackageInfo(related.NewObj)

			if pkgNode.Package+pkgNode.Class == oldInfo.Package+oldInfo.Class {
				for _, method := range pkgNode.Methods {
					if method.Name == oldInfo.Method {
						updateSelfRefs(pkgNode, method, newInfo)
					}
				}
			}

			for _, methodCall := range pkgNode.MethodCalls {
				if methodCall.Package+methodCall.Class == oldInfo.Package+oldInfo.Class {
					if methodCall.MethodName == oldInfo.Method {
						updateSelfRefs(pkgNode, *methodCallToMethodModel(methodCall), newInfo)
					}
				}
			}
		}
	}
}

func methodCallToMethodModel(call JMethodCall) *JMethod {
	return &JMethod{call.MethodName, call.Type, call.StartLine, call.StartLinePosition, call.StopLine, call.StopLinePosition, nil, nil, false}
}

func updateSelfRefs(node JClassNode, method JMethod, info *support3.PackageClassInfo) {
	path := node.Path
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if i == method.StartLine-1 {
			newLine := line[:method.StartLinePosition] + info.Method + line[method.StopLinePosition:]
			lines[i] = newLine
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
