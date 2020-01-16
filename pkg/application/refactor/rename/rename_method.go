package unused

import (
	"github.com/phodal/coca/pkg/application/refactor/rename/support"
	"github.com/phodal/coca/pkg/domain/core_domain"
	"github.com/phodal/coca/pkg/domain/jdomain"
	"io/ioutil"
	"log"
	"strings"
)

var parsedChange []support.RefactorChangeRelate

type RemoveMethodApp struct {
}

var parsedDeps []jdomain.JClassNode

func RenameMethodApp(deps []jdomain.JClassNode) *RemoveMethodApp {
	parsedDeps = deps
	return &RemoveMethodApp{}
}

func (j *RemoveMethodApp) Refactoring(conf string) {
	parsedChange = support.ParseRelates(conf)
	startParse(parsedDeps, parsedChange)
}

func startParse(nodes []jdomain.JClassNode, relates []support.RefactorChangeRelate) {
	for _, pkgNode := range nodes {
		for _, related := range relates {
			oldInfo := support.BuildMethodPackageInfo(related.OldObj)
			newInfo := support.BuildMethodPackageInfo(related.NewObj)

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

func methodCallToMethodModel(call core_domain.CodeCall) core_domain.JMethod {
	position := core_domain.CodePosition{
		StartLine:         call.Position.StartLine,
		StartLinePosition: call.Position.StartLinePosition,
		StopLine:          call.Position.StopLine,
		StopLinePosition:  call.Position.StopLinePosition,
	}
	return core_domain.JMethod{
		Name:     call.MethodName,
		Type:     call.Type,
		Position: position,
	}
}

func updateSelfRefs(node jdomain.JClassNode, method core_domain.JMethod, info *support.PackageClassInfo) {
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
