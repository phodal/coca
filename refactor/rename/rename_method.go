package unused

import (
	. "../../adapter/models"
	. "../../utils"
	. "../../utils/models"
	. "../base/models"
	"encoding/json"
	"fmt"
)

var parsedChange []RefactorChangeRelate
var nodes []JMoveStruct

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
	file := ReadFile(depsFile)
	if file == nil {
		return
	}

	_ = json.Unmarshal(file, &parsedDeps)

	configBytes := ReadFile(configPath)
	if configBytes == nil {
		return
	}

	conf = string(configBytes)

	parsedChange = ParseRelates(conf)

	startParse(parsedDeps, parsedChange)
}

func startParse(nodes []JClassNode, relates []RefactorChangeRelate) {
	for _, related := range relates {
		oldInfo := BuildMethodPackageInfo(related.OldObj)
		newInfo := BuildMethodPackageInfo(related.NewObj)

		fmt.Print(oldInfo, newInfo)
	}
}
