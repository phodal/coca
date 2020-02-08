package unused

import (
	"github.com/phodal/coca/pkg/application/refactor/rename/support"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"io/ioutil"
	"log"
	"strings"
)

var parsedChange []support.RefactorChangeRelate

type RemoveMethodApp struct {
}

var parsedDeps []core_domain.CodeDataStruct

func RenameMethodApp(deps []core_domain.CodeDataStruct) *RemoveMethodApp {
	parsedDeps = deps
	return &RemoveMethodApp{}
}

func (j *RemoveMethodApp) Refactoring(conf string) {
	parsedChange = support.ParseRelates(conf)
	startParse(parsedDeps, parsedChange)
}

func startParse(nodes []core_domain.CodeDataStruct, relates []support.RefactorChangeRelate) {
	for _, pkgNode := range nodes {
		for _, related := range relates {
			oldInfo := support.BuildMethodPackageInfo(related.OldObj)
			newInfo := support.BuildMethodPackageInfo(related.NewObj)

			if pkgNode.Package+pkgNode.NodeName == oldInfo.Package+oldInfo.Class {
				for _, method := range pkgNode.Functions {
					if method.Name == oldInfo.Method {
						updateSelfRefs(pkgNode, method, newInfo)
					}
				}
			}

			for _, method := range pkgNode.Functions {
				for _, methodCall := range method.FunctionCalls {
					if methodCall.Package+methodCall.NodeName == oldInfo.Package+oldInfo.Class {
						if methodCall.FunctionName == oldInfo.Method {
							updateSelfRefs(pkgNode, methodCallToMethodModel(methodCall), newInfo)
						}
					}
				}
			}
		}
	}
}

func methodCallToMethodModel(call core_domain.CodeCall) core_domain.CodeFunction {
	position := core_domain.CodePosition{
		StartLine:         call.Position.StartLine,
		StartLinePosition: call.Position.StartLinePosition,
		StopLine:          call.Position.StopLine,
		StopLinePosition:  call.Position.StopLinePosition,
	}
	return core_domain.CodeFunction{
		Name:       call.FunctionName,
		ReturnType: call.Type,
		Position:   position,
	}
}

func updateSelfRefs(node core_domain.CodeDataStruct, method core_domain.CodeFunction, info *support.PackageClassInfo) {
	path := node.FilePath
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if i == method.Position.StartLine-1 {
			newLine := line[:method.Position.StartLinePosition] + info.Method + line[method.Position.StopLinePosition:]
			lines[i] = newLine
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
