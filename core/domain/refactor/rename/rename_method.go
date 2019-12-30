package unused

import (
	support3 "github.com/phodal/coca/core/domain/refactor/rename/support"
	. "github.com/phodal/coca/core/models"
	"github.com/phodal/coca/core/infrastructure"
	"io/ioutil"
	"log"
	"strings"
)

var parsedChange []support3.RefactorChangeRelate

type RemoveMethodApp struct {
}

var configPath string
var conf string
var parsedDeps []JClassNode

func RenameMethodApp(deps []JClassNode, p string) *RemoveMethodApp {
	parsedDeps = deps
	configPath = p
	return &RemoveMethodApp{}
}

func (j *RemoveMethodApp) Start() {
	configBytes := infrastructure.ReadFile(configPath)
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

			for _, method := range pkgNode.Methods {
				for _, methodCall := range method.MethodCalls {
					if methodCall.Package+methodCall.Class == oldInfo.Package+oldInfo.Class {
						if methodCall.MethodName == oldInfo.Method {
							updateSelfRefs(pkgNode, methodCallToMethodModel(methodCall), newInfo)
						}
					}
				}
			}
		}
	}
}

func methodCallToMethodModel(call JMethodCall) JMethod {
	return *&JMethod{Name: call.MethodName, Type: call.Type, StartLine: call.StartLine, StartLinePosition: call.StartLinePosition, StopLine: call.StopLine, StopLinePosition: call.StopLinePosition}
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
