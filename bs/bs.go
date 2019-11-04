package bs
import (
	"encoding/json"
	"fmt"
	. "github.com/phodal/coca/adapter/models"
	. "github.com/phodal/coca/refactor/base/models"
	. "github.com/phodal/coca/utils"
)

var nodes []JMoveStruct

type BadSmellApp struct {
}

var depsFile string
var parsedDeps []JClassNode

func NewBadSmellApp(depPath string) *BadSmellApp {
	depsFile = depPath
	return &BadSmellApp{}
}

func (j *BadSmellApp) Start() {
	file := ReadFile(depsFile)
	if file == nil {
		return
	}

	_ = json.Unmarshal(file, &parsedDeps)

	fmt.Println(parsedDeps)
}
